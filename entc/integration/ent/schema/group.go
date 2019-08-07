package schema

import (
	"errors"
	"regexp"
	"strings"

	"fbc/ent"
	"fbc/ent/edge"
	"fbc/ent/field"
)

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
			MinLen(3),
		field.Int("max_users").
			Optional().
			Positive().
			Default(10),
		field.String("name").
			Comment("field with multiple validators").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("last name must begin with uppercase")
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
