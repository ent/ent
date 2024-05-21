// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/url"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

// ExValueScan holds the schema definition for the ExValueScan entity.
type ExValueScan struct {
	ent.Schema
}

// Fields of the ExValueScan.
func (ExValueScan) Fields() []ent.Field {
	return []ent.Field{
		field.String("binary").
			GoType(&url.URL{}).
			ValueScanner(field.BinaryValueScanner[*url.URL]{}),
		field.Bytes("binary_bytes").
			GoType(&url.URL{}).
			ValueScanner(field.BinaryValueScanner[*url.URL]{}),
		field.String("binary_optional").
			Optional().
			GoType(&url.URL{}).
			ValueScanner(field.BinaryValueScanner[*url.URL]{}),
		field.String("text").
			GoType(&big.Int{}).
			ValueScanner(field.TextValueScanner[*big.Int]{}),
		field.String("text_optional").
			Optional().
			GoType(&big.Int{}).
			ValueScanner(field.TextValueScanner[*big.Int]{}),
		field.String("base64").
			ValueScanner(field.ValueScannerFunc[string, *sql.NullString]{
				V: func(s string) (driver.Value, error) {
					return base64.StdEncoding.EncodeToString([]byte(s)), nil
				},
				S: func(ns *sql.NullString) (string, error) {
					if !ns.Valid {
						return "", nil
					}
					b, err := base64.StdEncoding.DecodeString(ns.String)
					if err != nil {
						return "", err
					}
					return string(b), nil
				},
			}),
		field.String("custom").
			ValueScanner(PrefixedHex{
				prefix: "0x",
			}),
		field.String("custom_optional").
			Optional().
			ValueScanner(PrefixedHex{
				prefix: "0X",
			}),
	}
}

// PrefixedHex is a custom type that implements the ValueScanner interface.
type PrefixedHex struct {
	prefix string
}

// Value implements the TypeValueScanner.Value method.
func (p PrefixedHex) Value(s string) (driver.Value, error) {
	return p.prefix + ":" + hex.EncodeToString([]byte(s)), nil
}

// ScanValue implements the TypeValueScanner.ScanValue method.
func (PrefixedHex) ScanValue() field.ValueScanner {
	return &sql.NullString{}
}

// FromValue implements the TypeValueScanner.FromValue method.
func (p PrefixedHex) FromValue(v driver.Value) (string, error) {
	s, ok := v.(*sql.NullString)
	if !ok {
		return "", fmt.Errorf("unexpected input for FromValue: %T", v)
	}
	if !s.Valid {
		return "", nil
	}
	d, err := hex.DecodeString(strings.TrimPrefix(s.String, p.prefix+":"))
	if err != nil {
		return "", err
	}
	return string(d), nil
}
