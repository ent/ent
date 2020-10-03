// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package edge_test

import (
	"testing"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEdge(t *testing.T) {
	assert := assert.New(t)
	type User struct{ ent.Schema }
	e := edge.To("friends", User.Type).
		Required().
		Descriptor()
	assert.False(e.Inverse)
	assert.Equal("User", e.Type)
	assert.Equal("friends", e.Name)
	assert.True(e.Required)

	type Node struct{ ent.Schema }
	e = edge.To("parent", Node.Type).
		Unique().
		Descriptor()
	assert.False(e.Inverse)
	assert.True(e.Unique)
	assert.Equal("Node", e.Type)
	assert.Equal("parent", e.Name)
	assert.False(e.Required)

	t.Log("m2m relation of the same type")
	from := edge.To("following", User.Type).
		From("followers").
		Descriptor()

	assert.True(from.Inverse)
	assert.False(from.Unique)
	assert.Equal("followers", from.Name)
	assert.NotNil(from.Ref)
	assert.Equal("following", from.Ref.Name)
	assert.False(from.Ref.Unique)

	t.Log("o2m relation of the same type")
	from = edge.To("following", User.Type).
		Unique().
		From("followers").
		Descriptor()
	assert.False(from.Unique)
	assert.True(from.Ref.Unique)
	from = edge.To("following", User.Type).
		From("followers").
		Unique().
		Descriptor()
	assert.True(from.Unique)
	assert.False(from.Ref.Unique)

	t.Log("o2o relation of the same type")
	from = edge.To("following", User.Type).
		Unique().
		From("followers").
		Unique().
		Descriptor()
	assert.True(from.Unique)
	assert.True(from.Ref.Unique)

	e = edge.To("user", User.Type).
		StructTag(`json:"user_name,omitempty"`).
		Descriptor()
	assert.Equal(`json:"user_name,omitempty"`, e.Tag)

	from = edge.To("following", User.Type).
		StructTag("following").
		StorageKey(edge.Table("user_followers"), edge.Columns("following_id", "followers_id")).
		From("followers").
		StructTag("followers").
		Descriptor()
	assert.Equal("followers", from.Tag)
	assert.Equal("following", from.Ref.Tag)
	assert.Equal(edge.StorageKey{Table: "user_followers", Columns: []string{"following_id", "followers_id"}}, *from.Ref.StorageKey)
}

type GQL struct {
	Field string
}

func (GQL) Name() string {
	return "GQL"
}

func TestAnnotations(t *testing.T) {
	type User struct{ ent.Schema }
	to := edge.To("user", User.Type).
		Annotations(GQL{Field: "to"}).
		Descriptor()
	require.Equal(t, []schema.Annotation{GQL{Field: "to"}}, to.Annotations)
	from := edge.From("user", User.Type).
		Annotations(GQL{Field: "from"}).
		Descriptor()
	require.Equal(t, []schema.Annotation{GQL{Field: "from"}}, from.Annotations)
	bidi := edge.To("following", User.Type).
		Annotations(GQL{Field: "to"}).
		From("followers").
		Annotations(GQL{Field: "from"}).
		Descriptor()
	require.Equal(t, []schema.Annotation{GQL{Field: "from"}}, bidi.Annotations)
	require.Equal(t, []schema.Annotation{GQL{Field: "to"}}, bidi.Ref.Annotations)
}
