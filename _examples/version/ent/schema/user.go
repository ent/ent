// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		VersionMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("online", "offline"),
	}
}

// VersionMixin provides an optimistic concurrency
// control mechanism using a "version" field.
type VersionMixin struct {
	mixin.Schema
}

// Fields of the VersionMixin.
func (VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("version").
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}).
			Comment("Unix time of when the latest update occurred"),
	}
}
