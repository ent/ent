package schema

import (
	"fbc/ent"
	"fbc/ent/field"
	"fbc/ent/index"
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
		field.String("address").Optional(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "address").
			Unique(),
	}
}
