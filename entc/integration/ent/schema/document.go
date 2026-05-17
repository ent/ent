// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
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
		field.String("name"),
		field.Blob("content").
			Lazy().
			Key(func(context.Context) (string, error) {
				return fmt.Sprintf("documents/%s/content", uuid.NewString()), nil
			}),
		field.Blob("thumbnail").
			Lazy(),
		field.Blob("attachment").
			DualWrite(),
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
	}
}
