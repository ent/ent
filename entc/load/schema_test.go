// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"

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
			Edges("parent").
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
	for _, u := range []ent.Interface{User{}, &User{}} {
		buf, err := MarshalSchema(u)
		require.NoError(t, err)

		schema := &Schema{}
		require.NoError(t, json.Unmarshal(buf, schema))
		require.Equal(t, "User", schema.Name)
		require.Len(t, schema.Fields, 4)
		require.Equal(t, "age", schema.Fields[0].Name)
		require.Equal(t, field.TypeInt, schema.Fields[0].Info.Type)

		require.Equal(t, "name", schema.Fields[1].Name)
		require.Equal(t, field.TypeString, schema.Fields[1].Info.Type)

		require.Equal(t, "nillable", schema.Fields[2].Name)
		require.Equal(t, field.TypeString, schema.Fields[2].Info.Type)
		require.True(t, schema.Fields[2].Nillable)
		require.False(t, schema.Fields[2].Optional)

		require.Equal(t, "optional", schema.Fields[3].Name)
		require.Equal(t, field.TypeString, schema.Fields[3].Info.Type)
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
		require.Equal(t, []string{"parent"}, schema.Indexes[1].Edges)
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

type WithDefaults struct {
	ent.Schema
}

func (WithDefaults) Fields() []ent.Field {
	return []ent.Field{
		field.Int("int").
			Default(1),
		field.Float("float").
			Default(math.Pi),
		field.String("string").
			Default("foo"),
		field.Bool("string").
			Default(true),
		field.Time("updated_at").
			UpdateDefault(time.Now),
	}
}

func (WithDefaults) Edges() []ent.Edge {
	return nil
}

func (WithDefaults) Indexes() []ent.Index {
	return nil
}

func TestMarshalDefaults(t *testing.T) {
	d := WithDefaults{}
	buf, err := MarshalSchema(d)
	require.NoError(t, err)

	schema := &Schema{}
	err = json.Unmarshal(buf, schema)
	require.NoError(t, err)

	require.Equal(t, "WithDefaults", schema.Name)
	require.True(t, schema.Fields[0].Default)
	require.True(t, schema.Fields[1].Default)
	require.True(t, schema.Fields[2].Default)
	require.True(t, schema.Fields[3].Default)
	require.False(t, schema.Fields[4].Default)
	require.True(t, schema.Fields[4].UpdateDefault)
}
