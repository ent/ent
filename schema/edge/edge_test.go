package edge_test

import (
	"testing"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
	assert := assert.New(t)
	type User struct{ ent.Schema }
	e := edge.To("friends", User.Type).Required()
	assert.False(e.IsInverse())
	assert.Equal("User", e.Type())
	assert.Equal("friends", e.Name())
	assert.True(e.IsRequired())

	type Node struct{ ent.Schema }
	e = edge.To("parent", Node.Type).Unique()
	assert.False(e.IsInverse())
	assert.True(e.IsUnique())
	assert.Equal("Node", e.Type())
	assert.Equal("parent", e.Name())
	assert.False(e.IsRequired())

	t.Log("m2m relation of the same type")
	from := edge.To("following", User.Type).From("followers")

	assert.True(from.IsInverse())
	assert.False(from.IsUnique())
	assert.Equal("followers", from.Name())
	assert.NotNil(from.Assoc())
	assert.Equal("following", from.Assoc().Name())
	assert.False(from.Assoc().IsUnique())

	t.Log("o2m relation of the same type")
	from = edge.To("following", User.Type).Unique().From("followers")
	assert.False(from.IsUnique())
	assert.True(from.Assoc().IsUnique())
	from = edge.To("following", User.Type).From("followers").Unique()
	assert.True(from.IsUnique())
	assert.False(from.Assoc().IsUnique())

	t.Log("o2o relation of the same type")
	from = edge.To("following", User.Type).Unique().From("followers").Unique()
	assert.True(from.IsUnique())
	assert.True(from.Assoc().IsUnique())

	e = edge.To("user", User.Type).StructTag(`json:"user_name,omitempty"`)
	assert.Equal(`json:"user_name,omitempty"`, e.Tag())
	from = edge.To("following", User.Type).StructTag("following").From("followers").StructTag("followers")
	assert.Equal("followers", from.Tag())
	assert.Equal("following", from.Assoc().Tag())
}
