// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// ColumnScanner is the interface that wraps the
// four sql.Rows methods used for scanning.
type ColumnScanner interface {
	Next() bool
	Scan(...interface{}) error
	Columns() ([]string, error)
	Err() error
}

// ScanOne scans one row to the given value. It fails if the rows holds more than 1 row.
func ScanOne(rows ColumnScanner, v interface{}) error {
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("sql/scan: failed getting column names: %v", err)
	}
	if n := len(columns); n != 1 {
		return fmt.Errorf("sql/scan: unexpected number of columns: %d", n)
	}
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	if err := rows.Scan(v); err != nil {
		return err
	}
	if rows.Next() {
		return fmt.Errorf("sql/scan: expect exactly one row in result set")
	}
	return rows.Err()
}

// ScanInt64 scans and returns an int64 from the rows columns.
func ScanInt64(rows ColumnScanner) (int64, error) {
	var n int64
	if err := ScanOne(rows, &n); err != nil {
		return 0, err
	}
	return n, nil
}

// ScanInt scans and returns an int from the rows columns.
func ScanInt(rows ColumnScanner) (int, error) {
	n, err := ScanInt64(rows)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// ScanString scans and returns a string from the rows columns.
func ScanString(rows ColumnScanner) (string, error) {
	var s string
	if err := ScanOne(rows, &s); err != nil {
		return "", err
	}
	return s, nil
}

// ScanSlice scans the given ColumnScanner (basically, sql.Row or sql.Rows) into the given slice.
func ScanSlice(rows ColumnScanner, v interface{}) error {
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("sql/scan: failed getting column names: %v", err)
	}
	rv := reflect.Indirect(reflect.ValueOf(v))
	if k := rv.Kind(); k != reflect.Slice {
		return fmt.Errorf("sql/scan: invalid type %s. expected slice as an argument", k)
	}
	scan, err := scanType(rv.Type().Elem(), columns)
	if err != nil {
		return err
	}
	if n, m := len(columns), len(scan.columns); n > m {
		return fmt.Errorf("sql/scan: columns do not match (%d > %d)", n, m)
	}
	for rows.Next() {
		values := scan.values()
		if err := rows.Scan(values...); err != nil {
			return fmt.Errorf("sql/scan: failed scanning rows: %v", err)
		}
		vv := reflect.Append(rv, scan.value(values...))
		rv.Set(vv)
	}
	return rows.Err()
}

// rowScan is the configuration for scanning one sql.Row.
type rowScan struct {
	// column types of a row.
	columns []reflect.Type
	// value functions that converts the row columns (result) to a reflect.Value.
	value func(v ...interface{}) reflect.Value
}

// values returns a []interface{} from the configured column types.
func (r *rowScan) values() []interface{} {
	values := make([]interface{}, len(r.columns))
	for i := range r.columns {
		values[i] = reflect.New(r.columns[i]).Interface()
	}
	return values
}

// scanType returns rowScan for the given reflect.Type.
func scanType(typ reflect.Type, columns []string) (*rowScan, error) {
	switch k := typ.Kind(); {
	case assignable(typ):
		return &rowScan{
			columns: []reflect.Type{typ},
			value: func(v ...interface{}) reflect.Value {
				return reflect.Indirect(reflect.ValueOf(v[0]))
			},
		}, nil
	case k == reflect.Ptr:
		return scanPtr(typ, columns)
	case k == reflect.Struct:
		return scanStruct(typ, columns)
	default:
		return nil, fmt.Errorf("sql/scan: unsupported type ([]%s)", k)
	}
}

var scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()

// assignable reports if the given type can be assigned directly by `Rows.Scan`.
func assignable(typ reflect.Type) bool {
	switch k := typ.Kind(); {
	case typ.Implements(scannerType):
	case k == reflect.Interface && typ.NumMethod() == 0:
	case k == reflect.String || k >= reflect.Bool && k <= reflect.Float64:
	case (k == reflect.Slice || k == reflect.Array) && typ.Elem().Kind() == reflect.Uint8:
	default:
		return false
	}
	return true
}

// scanStruct returns the a configuration for scanning an sql.Row into a struct.
func scanStruct(typ reflect.Type, columns []string) (*rowScan, error) {
	var (
		scan  = &rowScan{}
		names = make(map[string]int)
		idx   = make([]int, 0, typ.NumField())
	)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		name := strings.ToLower(f.Name)
		if tag, ok := f.Tag.Lookup("sql"); ok {
			name = tag
		} else if tag, ok := f.Tag.Lookup("json"); ok {
			name = strings.Split(tag, ",")[0]
		}
		names[name] = i
	}
	for _, c := range columns {
		// normalize columns if necessary, for example: COUNT(*) => count.
		name := strings.ToLower(strings.Split(c, "(")[0])
		i, ok := names[name]
		if !ok {
			return nil, fmt.Errorf("sql/scan: missing struct field for column: %s (%s)", c, name)
		}
		idx = append(idx, i)
		scan.columns = append(scan.columns, typ.Field(i).Type)
	}
	scan.value = func(vs ...interface{}) reflect.Value {
		st := reflect.New(typ).Elem()
		for i, v := range vs {
			st.Field(idx[i]).Set(reflect.Indirect(reflect.ValueOf(v)))
		}
		return st
	}
	return scan, nil
}

// scanPtr wraps the underlying type with rowScan.
func scanPtr(typ reflect.Type, columns []string) (*rowScan, error) {
	typ = typ.Elem()
	scan, err := scanType(typ, columns)
	if err != nil {
		return nil, err
	}
	wrap := scan.value
	scan.value = func(vs ...interface{}) reflect.Value {
		v := wrap(vs...)
		pt := reflect.PtrTo(v.Type())
		pv := reflect.New(pt.Elem())
		pv.Elem().Set(v)
		return pv
	}
	return scan, nil
}
