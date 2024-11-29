// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"testing"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type OrderConfig struct {
	FieldName string
}

func (OrderConfig) Name() string {
	return "order_config"
}

func (o OrderConfig) Merge(ant schema.Annotation) schema.Annotation {
	o.FieldName = ant.(OrderConfig).FieldName
	return o
}

type IDConfig struct {
	TagName string
}

func (IDConfig) Name() string {
	return "id_config"
}

type PartialIndex struct {
	WhereClause string
}

func (PartialIndex) Name() string {
	return "partial_index"
}

func (p PartialIndex) Merge(ant schema.Annotation) schema.Annotation {
	p.WhereClause = ant.(PartialIndex).WhereClause
	return p
}

type AnnotationMixin struct {
	mixin.Schema
}

func (AnnotationMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		IDConfig{TagName: "id tag"},
		OrderConfig{FieldName: "mixin annotations"},
	}
}

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AnnotationMixin{},
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		OrderConfig{FieldName: "type annotations"},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Comment("some comment"),
		field.String("name").
			Default("unknown").
			Annotations(&OrderConfig{FieldName: "name"}),
		field.String("nillable").
			Nillable(),
		field.String("optional").
			Optional(),
		field.Enum("state").
			Values("on", "off").
			Optional(),
		field.String("sensitive").
			Sensitive(),
		field.Time("creation_time").
			Default(time.Now),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.Int("parent_id").
			Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type).
			Annotations(&OrderConfig{FieldName: "name"}),
		edge.To("parent", User.Type).
			Unique().
			Required().
			Immutable().
			Field("parent_id").
			StorageKey(edge.Column("parent_id")).
			From("children"),
		edge.To("following", User.Type).
			Annotations(&OrderConfig{FieldName: "following"}).
			From("followers").
			Annotations(&OrderConfig{FieldName: "followers"}),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "address").
			Unique(),
		index.Fields("name").
			Edges("parent").
			StorageKey("user_parent_name").
			Annotations(&PartialIndex{
				WhereClause: "age > 20",
			}).
			Unique(),
	}
}

type Group struct{ ent.Schema }

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

		schema, err := UnmarshalSchema(buf)
		require.NoError(t, err)
		require.Equal(t, "User", schema.Name)
		require.Len(t, schema.Annotations, 2)
		ant := schema.Annotations["order_config"].(map[string]any)
		require.Equal(t, ant["FieldName"], "type annotations")

		require.Len(t, schema.Fields, 9)
		require.Equal(t, "age", schema.Fields[0].Name)
		require.Equal(t, field.TypeInt, schema.Fields[0].Info.Type)

		require.Equal(t, "name", schema.Fields[1].Name)
		require.Equal(t, field.TypeString, schema.Fields[1].Info.Type)
		require.Equal(t, "unknown", schema.Fields[1].DefaultValue)
		require.NotEmpty(t, schema.Fields[1].Annotations)
		ant = schema.Fields[1].Annotations["order_config"].(map[string]any)
		require.Equal(t, ant["FieldName"], "name")

		require.Equal(t, "nillable", schema.Fields[2].Name)
		require.Equal(t, field.TypeString, schema.Fields[2].Info.Type)
		require.True(t, schema.Fields[2].Nillable)
		require.False(t, schema.Fields[2].Optional)
		require.False(t, schema.Fields[2].Sensitive)

		require.Equal(t, "optional", schema.Fields[3].Name)
		require.Equal(t, field.TypeString, schema.Fields[3].Info.Type)
		require.False(t, schema.Fields[3].Nillable)
		require.True(t, schema.Fields[3].Optional)

		require.Equal(t, "state", schema.Fields[4].Name)
		require.Equal(t, field.TypeEnum, schema.Fields[4].Info.Type)
		require.Equal(t, "on", schema.Fields[4].Enums[0].V)
		require.Equal(t, "off", schema.Fields[4].Enums[1].V)

		require.Equal(t, "sensitive", schema.Fields[5].Name)
		require.Equal(t, field.TypeString, schema.Fields[5].Info.Type)
		require.True(t, schema.Fields[5].Sensitive)
		require.Equal(t, reflect.Invalid, schema.Fields[5].DefaultKind)

		require.Equal(t, "creation_time", schema.Fields[6].Name)
		require.Equal(t, field.TypeTime, schema.Fields[6].Info.Type)
		require.Nil(t, schema.Fields[6].DefaultValue)
		require.Equal(t, reflect.Func, schema.Fields[6].DefaultKind)

		require.Equal(t, "uuid", schema.Fields[7].Name)
		require.Equal(t, field.TypeUUID, schema.Fields[7].Info.Type)
		require.True(t, schema.Fields[7].Default)
		require.Equal(t, "github.com/google/uuid", schema.Fields[7].Info.PkgPath)

		require.Equal(t, "parent_id", schema.Fields[8].Name)
		require.Equal(t, field.TypeInt, schema.Fields[8].Info.Type)
		require.True(t, schema.Fields[8].Optional)

		require.Len(t, schema.Edges, 3)
		require.Equal(t, "groups", schema.Edges[0].Name)
		require.Equal(t, "Group", schema.Edges[0].Type)
		require.False(t, schema.Edges[0].Inverse)
		require.NotEmpty(t, schema.Edges[0].Annotations)
		ant = schema.Edges[0].Annotations["order_config"].(map[string]any)
		require.Equal(t, ant["FieldName"], "name")

		require.Equal(t, "children", schema.Edges[1].Name)
		require.Equal(t, "parent_id", schema.Edges[1].StorageKey.Columns[0])
		require.Equal(t, "User", schema.Edges[1].Type)
		require.True(t, schema.Edges[1].Inverse)
		require.Equal(t, "parent", schema.Edges[1].Ref.Name)
		require.True(t, schema.Edges[1].Ref.Unique)
		require.True(t, schema.Edges[1].Ref.Required)
		require.True(t, schema.Edges[1].Ref.Immutable)
		require.Equal(t, "parent_id", schema.Edges[1].Ref.StorageKey.Columns[0])

		ant = schema.Edges[2].Annotations["order_config"].(map[string]any)
		require.Equal(t, ant["FieldName"], "followers")
		ant = schema.Edges[2].Ref.Annotations["order_config"].(map[string]any)
		require.Equal(t, ant["FieldName"], "following")

		require.Equal(t, []string{"name", "address"}, schema.Indexes[0].Fields)
		require.True(t, schema.Indexes[0].Unique)
		require.Equal(t, []string{"name"}, schema.Indexes[1].Fields)
		require.Equal(t, []string{"parent"}, schema.Indexes[1].Edges)
		require.Equal(t, "user_parent_name", schema.Indexes[1].StorageKey)
		require.True(t, schema.Indexes[1].Unique)
		ant = schema.Indexes[1].Annotations["partial_index"].(map[string]any)
		require.Equal(t, "age > 20", ant["WhereClause"])

		require.Equal(t, "some comment", schema.Fields[0].Comment)
		require.Empty(t, schema.Fields[1].Comment)
	}
}

type InvalidEdge struct {
	ent.Schema
}

// Edge panics because the edge declaration is invalid.
func (InvalidEdge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invalid", InvalidEdge{}.Type),
	}
}

type InvalidUUID struct {
	ent.Schema
}

func (InvalidUUID) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("invalid", uuid.New()).
			Default(time.Now),
	}
}

func TestMarshalFails(t *testing.T) {
	i1 := InvalidEdge{}
	buf, err := MarshalSchema(i1)
	require.Error(t, err)
	require.Nil(t, buf)

	i2 := InvalidUUID{}
	buf, err = MarshalSchema(i2)
	require.Nil(t, buf)
	require.EqualError(t, err, `schema "InvalidUUID": field "invalid": expect type (func() uuid.UUID) for uuid default value`)
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
		// see issue #1146
		field.Int("int_default_func").
			DefaultFunc(func() int {
				return 1e9
			}),
		field.Float("balance").
			Default(0),
		field.JSON("dirs", []http.Dir{}).
			Default([]http.Dir{"/tmp"}),
		field.Float("float_default_func").
			DefaultFunc(func() float64 {
				return math.Pi
			}),
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
	require.True(t, schema.Fields[5].Default)
	require.Equal(t, schema.Fields[5].DefaultKind, reflect.Func)
	require.True(t, schema.Fields[6].Default)
	require.True(t, schema.Fields[7].Default)
	require.Equal(t, schema.Fields[8].DefaultKind, reflect.Func)
}

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

type HooksMixin struct {
	mixin.Schema
}

func (HooksMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("boring"),
	}
}

func (HooksMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique(),
	}
}

func (HooksMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("boring").
			Edges("user"),
	}
}

func (HooksMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(ent.Mutator) ent.Mutator { return nil },
		func(ent.Mutator) ent.Mutator { return nil },
	}
}

type BoringPolicy struct{}

func (BoringPolicy) EvalMutation(context.Context, ent.Mutation) error { return nil }
func (BoringPolicy) EvalQuery(context.Context, ent.Query) error       { return nil }

type PrivacyMixin struct {
	mixin.Schema
}

func (PrivacyMixin) Policy() ent.Policy {
	return BoringPolicy{}
}

type WithMixin struct {
	ent.Schema
}

func (WithMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		HooksMixin{},
		PrivacyMixin{},
	}
}

func (WithMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("field"),
	}
}

func (WithMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type),
	}
}

func (WithMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("field").
			Edges("owner").
			Unique(),
	}
}

func (WithMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(ent.Mutator) ent.Mutator { return nil },
	}
}

func (WithMixin) Policy() ent.Policy {
	return BoringPolicy{}
}

func TestMarshalMixin(t *testing.T) {
	d := WithMixin{}
	buf, err := MarshalSchema(d)
	require.NoError(t, err)

	schema := &Schema{}
	err = json.Unmarshal(buf, schema)
	require.NoError(t, err)

	t.Run("Fields", func(t *testing.T) {
		require.Equal(t, "WithMixin", schema.Name)
		require.Equal(t, "created_at", schema.Fields[0].Name)
		require.True(t, schema.Fields[0].Default)
		require.True(t, schema.Fields[0].Position.MixedIn)
		require.Equal(t, 0, schema.Fields[0].Position.MixinIndex)
		require.Equal(t, 0, schema.Fields[0].Position.Index)

		require.Equal(t, "updated_at", schema.Fields[1].Name)
		require.True(t, schema.Fields[1].Default)
		require.True(t, schema.Fields[1].UpdateDefault)
		require.True(t, schema.Fields[1].Position.MixedIn)
		require.Equal(t, 0, schema.Fields[1].Position.MixinIndex)
		require.Equal(t, 1, schema.Fields[1].Position.Index)

		require.Equal(t, "boring", schema.Fields[2].Name)
		require.False(t, schema.Fields[2].Default)
		require.False(t, schema.Fields[2].UpdateDefault)
		require.True(t, schema.Fields[2].Position.MixedIn)
		require.Equal(t, 1, schema.Fields[2].Position.MixinIndex)
		require.Equal(t, 0, schema.Fields[2].Position.Index)

		require.Equal(t, "field", schema.Fields[3].Name)
		require.False(t, schema.Fields[3].Default)
		require.False(t, schema.Fields[3].Position.MixedIn)
		require.Equal(t, 0, schema.Fields[3].Position.Index)
	})

	t.Run("Hooks", func(t *testing.T) {
		require.True(t, schema.Hooks[0].MixedIn)
		require.True(t, schema.Hooks[1].MixedIn)

		require.Equal(t, 1, schema.Hooks[0].MixinIndex)
		require.Equal(t, 1, schema.Hooks[1].MixinIndex)
		require.Equal(t, 0, schema.Hooks[0].Index)
		require.Equal(t, 1, schema.Hooks[1].Index)

		require.False(t, schema.Hooks[2].MixedIn)
		require.Equal(t, 0, schema.Hooks[2].Index)
		require.Equal(t, 0, schema.Hooks[2].MixinIndex)
	})

	t.Run("Edges", func(t *testing.T) {
		require.Len(t, schema.Edges, 2)
		require.Equal(t, "user", schema.Edges[0].Name)
		require.Equal(t, "User", schema.Edges[0].Type)
		require.True(t, schema.Edges[0].Unique)

		require.Equal(t, "owner", schema.Edges[1].Name)
		require.Equal(t, "User", schema.Edges[1].Type)
		require.False(t, schema.Edges[1].Unique)
	})

	t.Run("Indexes", func(t *testing.T) {
		require.Len(t, schema.Indexes, 2)
		require.Equal(t, []string{"boring"}, schema.Indexes[0].Fields)
		require.Equal(t, []string{"user"}, schema.Indexes[0].Edges)
		require.False(t, schema.Indexes[0].Unique)

		require.Equal(t, []string{"field"}, schema.Indexes[1].Fields)
		require.Equal(t, []string{"owner"}, schema.Indexes[1].Edges)
		require.True(t, schema.Indexes[1].Unique)
	})

	t.Run("Policy", func(t *testing.T) {
		require.Len(t, schema.Policy, 2)
		require.True(t, schema.Policy[0].MixedIn)
		require.False(t, schema.Policy[1].MixedIn)
	})
}
