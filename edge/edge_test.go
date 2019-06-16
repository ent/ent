package edge_test

import (
	"testing"

	"fbc/ent"
	"fbc/ent/edge"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
	assert := assert.New(t)
	type User struct{ ent.Schema }
	e := edge.To("friends", User.Type).Required()
	assert.True(e.IsAssoc())
	assert.Equal("User", e.Type())
	assert.Equal("friends", e.Name())
	assert.True(e.IsRequired())

	type Node struct{ ent.Schema }
	e = edge.To("parent", Node.Type).Unique()
	assert.True(e.IsAssoc())
	assert.True(e.IsUnique())
	assert.Equal("Node", e.Type())
	assert.Equal("parent", e.Name())
	assert.False(e.IsRequired())

	t.Log("m2m relation of the same type")
	From := edge.To("following", User.Type).From("followers")
	assert.False(From.IsAssoc())
	assert.True(From.IsInverse())
	assert.False(From.IsUnique())
	assert.Equal("followers", From.Name())
	assert.NotNil(From.Assoc())
	assert.Equal("following", From.Assoc().Name())
	assert.False(From.Assoc().IsUnique())

	t.Log("o2m relation of the same type")
	From = edge.To("following", User.Type).Unique().From("followers")
	assert.False(From.IsUnique())
	assert.True(From.Assoc().IsUnique())
	From = edge.To("following", User.Type).From("followers").Unique()
	assert.True(From.IsUnique())
	assert.False(From.Assoc().IsUnique())

	t.Log("o2o relation of the same type")
	From = edge.To("following", User.Type).Unique().From("followers").Unique()
	assert.True(From.IsUnique())
	assert.True(From.Assoc().IsUnique())

	e = edge.To("user", User.Type).StructTag(`json:"user_name,omitempty"`)
	assert.Equal(`json:"user_name,omitempty"`, e.Tag())
	From = edge.To("following", User.Type).StructTag("following").From("followers").StructTag("followers")
	assert.Equal("followers", From.Tag())
	assert.Equal("following", From.Assoc().Tag())
}
