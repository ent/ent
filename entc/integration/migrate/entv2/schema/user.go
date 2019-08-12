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
		// changing the type of the field.
		field.Int("age"),
		// extending name field to longtext.
		field.Text("name"),
		// adding new columns.
		field.String("phone"),
		field.Bytes("buffer").
			Default([]byte("{}")),
		// deleting the address column.
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// deleting old indexes (name, address),
		// and defining a new one.
		index.Fields("phone", "age").
			Unique(),
	}
}

// Additional types to be added to the schema.
type (
	// Pet schema.
	Pet struct{ ent.Schema }
	// Group schema.
	Group struct{ ent.Schema }
)
