package load

import (
	"encoding/json"
	"testing"

	"fbc/ent"
	"fbc/ent/schema/edge"
	"fbc/ent/schema/field"
	"fbc/ent/schema/index"

	"github.com/stretchr/testify/require"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.String("nillable").
			Nillable(),
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

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "address").
			Unique(),
		index.Fields("name").
			FromEdge("parent").
			Unique(),
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

		require.Equal(t, "nillable", schema.Fields[2].Name)
		require.Equal(t, field.TypeString, schema.Fields[2].Type)
		require.True(t, schema.Fields[2].Nillable)
		require.False(t, schema.Fields[2].Optional)

		require.Equal(t, "optional", schema.Fields[3].Name)
		require.Equal(t, field.TypeString, schema.Fields[3].Type)
		require.False(t, schema.Fields[3].Nillable)
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

		require.Equal(t, []string{"name", "address"}, schema.Indexes[0].Fields)
		require.True(t, schema.Indexes[0].Unique)
		require.Equal(t, []string{"name"}, schema.Indexes[1].Fields)
		require.Equal(t, "parent", schema.Indexes[1].Edge)
		require.True(t, schema.Indexes[1].Unique)
	}
}

type Invalid struct {
	ent.Schema
}

// Edge panics because the edge declaration is invalid.
func (Invalid) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invalid", Invalid{}.Type),
	}
}

func TestMarshalFails(t *testing.T) {
	i := Invalid{}
	buf, err := MarshalSchema(i)
	require.Error(t, err)
	require.Nil(t, buf)
}
