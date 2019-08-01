package load

import (
	"encoding/json"
	"testing"

	"fbc/ent"
	"fbc/ent/edge"
	"fbc/ent/field"

	"github.com/stretchr/testify/require"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.String("nullable").
			Nullable(),
		field.String("optional").
			Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("parent", User.Type).
			Unique().
			From("children"),
	}
}

type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field { return nil }

func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type),
	}
}

func TestMarshalSchema(t *testing.T) {
	for _, u := range []ent.Schema{User{}, &User{}} {
		buf, err := MarshalSchema(u)
		require.NoError(t, err)

		schema := &Schema{}
		require.NoError(t, json.Unmarshal(buf, schema))
		require.Equal(t, "User", schema.Name)
		require.Len(t, schema.Fields, 4)
		require.Equal(t, "age", schema.Fields[0].Name)
		require.Equal(t, field.TypeInt, schema.Fields[0].Type)

		require.Equal(t, "name", schema.Fields[1].Name)
		require.Equal(t, field.TypeString, schema.Fields[1].Type)

		require.Equal(t, "nullable", schema.Fields[2].Name)
		require.Equal(t, field.TypeString, schema.Fields[2].Type)
		require.True(t, schema.Fields[2].Nullable)
		require.False(t, schema.Fields[2].Optional)

		require.Equal(t, "optional", schema.Fields[3].Name)
		require.Equal(t, field.TypeString, schema.Fields[3].Type)
		require.False(t, schema.Fields[3].Nullable)
		require.True(t, schema.Fields[3].Optional)

		require.Len(t, schema.Edges, 2)
		require.Equal(t, "groups", schema.Edges[0].Name)
		require.Equal(t, "Group", schema.Edges[0].Type)
		require.False(t, schema.Edges[0].Inverse)
		require.Equal(t, "children", schema.Edges[1].Name)
		require.Equal(t, "User", schema.Edges[1].Type)
		require.True(t, schema.Edges[1].Inverse)
		require.Equal(t, "parent", schema.Edges[1].Ref.Name)
		require.True(t, schema.Edges[1].Ref.Unique)
	}
}
