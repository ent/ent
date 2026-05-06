// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

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
	}
}
