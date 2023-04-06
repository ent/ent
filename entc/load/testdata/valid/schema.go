// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package valid

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the user schema.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Group holds the group schema.
type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expired_at"),
		field.String("organization"),
	}
}

// Tag holds the tag schema.
type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
	}
}
