// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package base

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// base schema for sharing fields and edges.
type base struct {
	ent.Schema
}

func (base) Fields() []ent.Field {
	return []ent.Field{
		field.Int("base_field"),
	}
}

// User holds the user schema.
type User struct {
	base
}

func (u User) Fields() []ent.Field {
	return append(
		u.base.Fields(),
		field.String("user_field"),
	)
}
