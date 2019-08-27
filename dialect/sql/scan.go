// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"fmt"
	"reflect"
	"strings"
)

// ColumnScanner is the interface that wraps the
// three sql.Rows methods used for scanning.
type ColumnScanner interface {
	Next() bool
	Scan(...interface{}) error
	Columns() ([]string, error)
}

// ScanSlice scans the given ColumnScanner (basically, sql.Rows or sql.Rows) into the given slice.
func ScanSlice(rows ColumnScanner, v interface{}) error {
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("sql/scan: failed getting column names: %v", err)
	}
	rv := reflect.Indirect(reflect.ValueOf(v))
	if k := rv.Kind(); k != reflect.Slice {
		return fmt.Errorf("sql/scan: invalid type %s. expected slice as an argument", k)
	}
	var (
		scan *rowScan
		typ  = rv.Type().Elem()
	)
	switch k := typ.Kind(); {
	case k == reflect.String || k >= reflect.Bool && k <= reflect.Float64:
		scan = &rowScan{
			columns: []reflect.Type{typ},
			value: func(v ...interface{}) reflect.Value {
				return reflect.Indirect(reflect.ValueOf(v[0]))
			},
		}
	case k == reflect.Ptr:
		typ = typ.Elem()
		if scan, err = scanStruct(typ, columns); err != nil {
			return err
		}
		wrap := scan.value
		scan.value = func(vs ...interface{}) reflect.Value {
			v := wrap(vs...)
			pt := reflect.PtrTo(v.Type())
			pv := reflect.New(pt.Elem())
			pv.Elem().Set(v)
			return pv
		}
	case k == reflect.Struct:
		if scan, err = scanStruct(typ, columns); err != nil {
			return err
		}
	default:
		return fmt.Errorf("sql/scan: unsupported type ([]%s)", k)
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
	return nil
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
		if tag, ok := f.Tag.Lookup("json"); ok {
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
