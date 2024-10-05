// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ScanOne scans one row to the given value. It fails if the rows holds more than 1 row.
func ScanOne(rows ColumnScanner, v any) error {
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("sql/scan: failed getting column names: %w", err)
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

// ScanInt64 scans and returns an int64 from the rows.
func ScanInt64(rows ColumnScanner) (int64, error) {
	var n int64
	if err := ScanOne(rows, &n); err != nil {
		return 0, err
	}
	return n, nil
}

// ScanInt scans and returns an int from the rows.
func ScanInt(rows ColumnScanner) (int, error) {
	n, err := ScanInt64(rows)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// ScanBool scans and returns a boolean from the rows.
func ScanBool(rows ColumnScanner) (bool, error) {
	var b bool
	if err := ScanOne(rows, &b); err != nil {
		return false, err
	}
	return b, nil
}

// ScanString scans and returns a string from the rows.
func ScanString(rows ColumnScanner) (string, error) {
	var s string
	if err := ScanOne(rows, &s); err != nil {
		return "", err
	}
	return s, nil
}

// ScanValue scans and returns a driver.Value from the rows.
func ScanValue(rows ColumnScanner) (driver.Value, error) {
	var v driver.Value
	if err := ScanOne(rows, &v); err != nil {
		return "", err
	}
	return v, nil
}

// ScanSlice scans the given ColumnScanner (basically, sql.Row or sql.Rows) into the given slice.
func ScanSlice(rows ColumnScanner, v any) error {
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("sql/scan: failed getting column names: %w", err)
	}
	rv := reflect.ValueOf(v)
	switch {
	case rv.Kind() != reflect.Ptr:
		if t := reflect.TypeOf(v); t != nil {
			return fmt.Errorf("sql/scan: ScanSlice(non-pointer %s)", t)
		}
		fallthrough
	case rv.IsNil():
		return fmt.Errorf("sql/scan: ScanSlice(nil)")
	}
	rv = reflect.Indirect(rv)
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
			return fmt.Errorf("sql/scan: failed scanning rows: %w", err)
		}
		vv, err := scan.value(values...)
		if err != nil {
			return err
		}
		rv.Set(reflect.Append(rv, vv))
	}
	return rows.Err()
}

// rowScan is the configuration for scanning one sql.Row.
type rowScan struct {
	// column types of a row.
	columns []reflect.Type
	// value functions that converts the row columns (result) to a reflect.Value.
	value func(v ...any) (reflect.Value, error)
}

// values returns a []any from the configured column types.
func (r *rowScan) values() []any {
	values := make([]any, len(r.columns))
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
			value: func(v ...any) (reflect.Value, error) {
				return reflect.Indirect(reflect.ValueOf(v[0])), nil
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

var (
	timeType     = reflect.TypeOf(time.Time{})
	scannerType  = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	nullJSONType = reflect.TypeOf((*nullJSON)(nil)).Elem()
)

// nullJSON represents a json.RawMessage that may be NULL.
type nullJSON json.RawMessage

// Scan implements the sql.Scanner interface.
func (j *nullJSON) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	*j = v.([]byte)
	return nil
}

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

// scanStruct returns the configuration for scanning a sql.Row into a struct.
func scanStruct(typ reflect.Type, columns []string) (*rowScan, error) {
	var (
		scan  = &rowScan{}
		idxs  = make([][]int, 0, typ.NumField())
		names = make(map[string][]int, typ.NumField())
	)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		// Skip unexported fields.
		if f.PkgPath != "" {
			continue
		}
		// Support 1-level embedding to accept types as `type T struct {ent.T; V int}`.
		if typ := f.Type; f.Anonymous && typ.Kind() == reflect.Struct {
			for j := 0; j < typ.NumField(); j++ {
				names[columnName(typ.Field(j))] = []int{i, j}
			}
			continue
		}
		names[columnName(f)] = []int{i}
	}
	for _, c := range columns {
		var idx []int
		// Normalize columns if necessary,
		// for example: COUNT(*) => count.
		switch name := strings.Split(c, "(")[0]; {
		case names[name] != nil:
			idx = names[name]
		case names[strings.ToLower(name)] != nil:
			idx = names[strings.ToLower(name)]
		default:
			return nil, fmt.Errorf("sql/scan: missing struct field for column: %s (%s)", c, name)
		}
		idxs = append(idxs, idx)
		rtype := typ.Field(idx[0]).Type
		if len(idx) > 1 {
			rtype = rtype.Field(idx[1]).Type
		}
		switch {
		// If the field is not support by the standard
		// convertAssign, assume it is a JSON field.
		case !supportsScan(rtype):
			rtype = nullJSONType
		// Create a pointer to the actual reflect
		// types to accept optional struct fields.
		case !nillable(rtype):
			rtype = reflect.PtrTo(rtype)
		}
		scan.columns = append(scan.columns, rtype)
	}
	scan.value = func(vs ...any) (reflect.Value, error) {
		st := reflect.New(typ).Elem()
		for i, v := range vs {
			rv := reflect.Indirect(reflect.ValueOf(v))
			if rv.IsNil() {
				continue
			}
			idx := idxs[i]
			rvalue, ft := st.Field(idx[0]), st.Type().Field(idx[0])
			if len(idx) > 1 {
				// Embedded field.
				rvalue, ft = rvalue.Field(idx[1]), ft.Type.Field(idx[1])
			}
			switch {
			case rv.Type() == nullJSONType:
				if rv = reflect.Indirect(rv); rv.IsNil() {
					continue
				}
				if err := json.Unmarshal(rv.Bytes(), rvalue.Addr().Interface()); err != nil {
					return reflect.Value{}, fmt.Errorf("unmarshal field %q: %w", ft.Name, err)
				}
			case !nillable(rvalue.Type()):
				rv = reflect.Indirect(rv)
				fallthrough
			default:
				rvalue.Set(rv)
			}
		}
		return st, nil
	}
	return scan, nil
}

// columnName returns the column name of a struct-field.
func columnName(f reflect.StructField) string {
	name := strings.ToLower(f.Name)
	if tag, ok := f.Tag.Lookup("sql"); ok {
		name = tag
	} else if tag, ok := f.Tag.Lookup("json"); ok {
		name = strings.Split(tag, ",")[0]
	}
	return name
}

// nillable reports if the reflect-type can have nil value.
func nillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Interface, reflect.Slice, reflect.Map, reflect.Ptr, reflect.UnsafePointer:
		return true
	}
	return false
}

// scanPtr wraps the underlying type with rowScan.
func scanPtr(typ reflect.Type, columns []string) (*rowScan, error) {
	typ = typ.Elem()
	scan, err := scanType(typ, columns)
	if err != nil {
		return nil, err
	}
	wrap := scan.value
	scan.value = func(vs ...any) (reflect.Value, error) {
		v, err := wrap(vs...)
		if err != nil {
			return reflect.Value{}, err
		}
		pt := reflect.PtrTo(v.Type())
		pv := reflect.New(pt.Elem())
		pv.Elem().Set(v)
		return pv, nil
	}
	return scan, nil
}

func supportsScan(t reflect.Type) bool {
	if t.Implements(scannerType) || reflect.PtrTo(t).Implements(scannerType) {
		return true
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Pointer, reflect.String:
		return true
	case reflect.Slice:
		return t == reflect.TypeOf(sql.RawBytes(nil)) || t == reflect.TypeOf([]byte(nil))
	case reflect.Interface:
		return t == reflect.TypeOf((*any)(nil)).Elem()
	default:
		return t == reflect.TypeOf(time.Time{}) || t.Implements(scannerType)
	}
}

// UnknownType is a named type to any indicates the info
// needs to be extracted from the underlying rows.
type UnknownType any

// ScanTypeOf returns the type used for scanning column i from the database.
func ScanTypeOf(rows *Rows, i int) any {
	unknown := new(any)
	ct, err := rows.ColumnTypes()
	if err != nil || len(ct) <= i {
		return unknown
	}
	rt := ct[i].ScanType()
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
	}
	// Handle NULL values.
	switch k := rt.Kind(); k {
	case reflect.Bool:
		rt = reflect.TypeOf(sql.NullBool{})
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		rt = reflect.TypeOf(sql.NullInt64{})
	case reflect.Float32, reflect.Float64:
		rt = reflect.TypeOf(sql.NullFloat64{})
	case reflect.String:
		rt = reflect.TypeOf(sql.NullString{})
	default:
		if k == reflect.Struct && rt == timeType {
			rt = reflect.TypeOf(sql.NullTime{})
		}
	}
	return reflect.New(rt).Interface()
}

// SelectValues maps a selected column to its value.
// Used by the generated code for storing runtime selected columns/expressions.
type SelectValues map[string]any

// Set sets the value of the given column.
func (s *SelectValues) Set(name string, v any) {
	if *s == nil {
		*s = make(SelectValues)
	}
	if pv, ok := v.(*any); ok && pv != nil {
		v = *pv
	}
	(*s)[name] = v
}

// Get returns the value of the given column.
func (s SelectValues) Get(name string) (any, error) {
	v, ok := s[name]
	if !ok {
		return nil, fmt.Errorf("%s value was not selected", name)
	}
	if v == nil {
		return nil, nil
	}
	switch rv := reflect.Indirect(reflect.ValueOf(v)).Interface().(type) {
	case NullString:
		if rv.Valid {
			return rv.String, nil
		}
	case NullInt64:
		if rv.Valid {
			return rv.Int64, nil
		}
	case NullFloat64:
		if rv.Valid {
			return rv.Float64, nil
		}
	case NullBool:
		if rv.Valid {
			return rv.Bool, nil
		}
	case NullTime:
		if rv.Valid {
			return rv.Time, nil
		}
	case sql.RawBytes:
		return []byte(rv), nil
	default:
		return rv, nil
	}
	return nil, nil
}
