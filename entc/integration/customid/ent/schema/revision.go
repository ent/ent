// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

// User holds the schema definition for the User entity.
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Revision holds the schema definition for the Revision entity.
type Revision struct {
	ent.Schema
}

// Fields of the Revision.
func (Revision) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
	}
}
