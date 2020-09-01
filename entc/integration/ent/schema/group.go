// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"regexp"
	"strings"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// CheckError is returned by the validators.
type CheckError struct {
	msg string
}

func (c CheckError) Error() string {
	return c.msg
}

// Group holds the schema for the group entity.
type Group struct {
	ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active").
			Default(true),
		field.Time("expire"),
		field.String("type").
			Optional().
			Nillable().
			MinLen(3).
			MaxLen(255),
		field.Int("max_users").
			Optional().
			Positive().
			Default(10),
		field.String("name").
			Comment("field with multiple validators").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return CheckError{msg: "last name must begin with uppercase"}
				}
				return nil
			}),
	}
}

// Edges of the group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type),
		edge.To("blocked", User.Type),
		edge.From("users", User.Type).Ref("groups"),
		edge.To("info", GroupInfo.Type).Unique().Required(),
	}
}
