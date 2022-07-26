//go:build !hidegroups

// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package buildflags

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

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
