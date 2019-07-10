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
		// changing the type of the field.
		field.Int("age"),
		// extending name field to longtext.
		field.Text("name"),
		// adding new column.
		field.String("phone"),
		// deleting the address column.
	}
}
