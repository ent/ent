// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// UserAuditLog holds the schema definition for the UserAuditLog entity.
type UserAuditLog struct {
	ent.Schema
}

// Fields of the UserAuditLog.
func (UserAuditLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("operation_type"),
		field.String("operation_time"),
		field.String("old_value").
			Optional(),
		field.String("new_value").
			Optional(),
	}
}
