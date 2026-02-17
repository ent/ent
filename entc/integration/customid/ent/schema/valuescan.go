// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"database/sql"
	"database/sql/driver"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ValueScan holds the schema definition for the ValueScan entity.
type ValueScan struct {
	ent.Schema
}

// ValueScanID is a custom ID type that relies on an external ValueScanner.
type ValueScanID struct {
	V int
}

// Fields of the ValueScan.
func (ValueScan) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			GoType(ValueScanID{}).
			ValueScanner(field.ValueScannerFunc[ValueScanID, *sql.NullInt64]{
				V: func(id ValueScanID) (driver.Value, error) {
					return int64(id.V), nil
				},
				S: func(id *sql.NullInt64) (ValueScanID, error) {
					if !id.Valid {
						return ValueScanID{}, nil
					}
					return ValueScanID{V: int(id.Int64)}, nil
				},
			}),
		field.String("name"),
	}
}
