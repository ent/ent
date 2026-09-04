// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"crypto"
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// DocPayload is a custom type for blob fields.
type DocPayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Blob("content").
			Lazy().
			HashKey(crypto.SHA256).
			CheckRefs(),
		field.Blob("thumbnail").
			Lazy().
			CheckRefs(),
		field.Blob("attachment").
			DualWrite().
			CheckRefs(),
		// "metadata" and the fields below deliberately skip CheckRefs, covering the
		// default: cleanup removes the object without looking for other holders.
		field.Blob("metadata").
			Optional(),
		field.Blob("payload").
			DualWrite(map[string]string{
				dialect.MySQL:    "longblob",
				dialect.Postgres: "jsonb",
				dialect.SQLite:   "json",
			}).
			GoType(&DocPayload{}).
			ValueScanner(field.ValueScannerFunc[*DocPayload, *sql.NullString]{
				V: func(v *DocPayload) (driver.Value, error) {
					return json.Marshal(v)
				},
				S: func(s *sql.NullString) (*DocPayload, error) {
					if !s.Valid {
						return nil, nil
					}
					var p DocPayload
					if err := json.Unmarshal([]byte(s.String), &p); err != nil {
						return nil, err
					}
					return &p, nil
				},
			}).
			Optional(),
		field.Blob("description").
			GoType("").
			Optional(),
		field.Blob("archive").
			Lazy().
			DualWrite().
			Optional(),
	}
}
