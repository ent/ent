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

func TestEdgeString(t *testing.T) {
	e := NewEdge(
		13, "develops",
		NewVertex(1, ""),
		NewVertex(10, ""),
	)
	assert.Equal(t, "e[13][1-develops->10]", fmt.Sprint(e))
}

func TestEdgeEncoding(t *testing.T) {
	t.Parallel()

	e := NewEdge(13, "develops",
		NewVertex(1, "person"),
		NewVertex(10, "software"),
	)
	got, err := graphson.MarshalToString(e)
	require.NoError(t, err)

	want := `{
		"@type" : "g:Edge",
		"@value" : {
			"id" : {
			  "@type" : "g:Int64",
			  "@value" : 13
			},
			"label" : "develops",
			"inVLabel" : "software",
			"outVLabel" : "person",
			"inV" : {
			  "@type" : "g:Int64",
			  "@value" : 10
			},
			"outV" : {
			  "@type" : "g:Int64",
			  "@value" : 1
			}
		}
	}`
	assert.JSONEq(t, want, got)

	e = Edge{}
	err = graphson.UnmarshalFromString(got, &e)
	require.NoError(t, err)

	assert.Equal(t, NewElement(int64(13), "develops"), e.Element)
	assert.Equal(t, NewVertex(int64(1), "person"), e.OutV)
	assert.Equal(t, NewVertex(int64(10), "software"), e.InV)
}

func TestPropertyEncoding(t *testing.T) {
	t.Parallel()

	props := []Property{
		NewProperty("from", int32(2017)),
		NewProperty("to", int32(2019)),
	}
	got, err := graphson.MarshalToString(props)
	require.NoError(t, err)

	want := `{
		"@type" : "g:List",
		"@value" : [
			{
				"@type" : "g:Property",
				"@value" : {
					"key" : "from",
					"value" : {
						"@type" : "g:Int32",
						"@value" : 2017
					}
				}
			},
			{
				"@type" : "g:Property",
				"@value" : {
					"key" : "to",
					"value" : {
						"@type" : "g:Int32",
						"@value" : 2019
					}
				}
			}
		]
	}`
	assert.JSONEq(t, want, got)
}

func TestPropertyString(t *testing.T) {
	p := NewProperty("since", 2019)
	assert.Equal(t, "p[since->2019]", fmt.Sprint(p))
}
