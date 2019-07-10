package schema

import (
	"fbc/ent"
	"fbc/ent/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("age"),
		field.String("name").MaxLen(10),
	}
}
