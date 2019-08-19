package base

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
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
