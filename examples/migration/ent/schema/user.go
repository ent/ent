package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Float("age"),
		field.String("name"),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// In case schema.ModeInspect is used without a dev-database, unnamed check constraints
		// should be normalized (i.e. identical to their definition in the database). In this
		// case, it is entsql.Check("(`age` > 0)"). See: https://atlasgo.io/concepts/dev-database.
		entsql.Check("age > 0"),

		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		entsql.Checks(map[string]string{
			"name_not_empty": "name <> ''",
		}),
	}
}
