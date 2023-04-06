// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

import (
	"fmt"
	"testing"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVertexCreation(t *testing.T) {
	v := NewVertex(45, "person")
	assert.Equal(t, 45, v.ID)
	assert.Equal(t, "person", v.Label)
	v = NewVertex(46, "")
	assert.Equal(t, "vertex", v.Label)
}

func TestVertexString(t *testing.T) {
	v := NewVertex(42, "")
	assert.Equal(t, "v[42]", fmt.Sprint(v))
}

func TestVertexEncoding(t *testing.T) {
	t.Parallel()

	v := NewVertex(1, "user")
	got, err := graphson.MarshalToString(v)
	require.NoError(t, err)

	want := `{
		"@type" : "g:Vertex",
		"@value" : {
			"id" : {
				"@type" : "g:Int64",
				"@value" : 1
			},
			"label" : "user"
		}
	}`
	assert.JSONEq(t, want, got)

	v = Vertex{}
	err = graphson.UnmarshalFromString(got, &v)
	require.NoError(t, err)

	assert.Equal(t, int64(1), v.ID)
	assert.Equal(t, "user", v.Label)
}

func TestVertexPropertyEncoding(t *testing.T) {
	t.Parallel()

	vp := NewVertexProperty("46ab60c2-918c-4cc4-a13b-350510e8908a", "name", "alex")
	got, err := graphson.MarshalToString(vp)
	require.NoError(t, err)

	want := `{
		"@type" : "g:VertexProperty",
		"@value" : {
			"id" : "46ab60c2-918c-4cc4-a13b-350510e8908a",
			"label": "name",
			"value": "alex"
		}
	}`
	assert.JSONEq(t, want, got)

	vp = VertexProperty{}
	err = graphson.UnmarshalFromString(got, &vp)
	require.NoError(t, err)

	assert.Equal(t, "46ab60c2-918c-4cc4-a13b-350510e8908a", vp.ID)
	assert.Equal(t, "name", vp.Key)
	assert.Equal(t, "alex", vp.Value)
}

func TestVertexPropertyString(t *testing.T) {
	vp := NewVertexProperty(55, "country", "israel")
	assert.Equal(t, "vp[country->israel]", fmt.Sprint(vp))
}
