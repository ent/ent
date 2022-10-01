// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

import (
	"fmt"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"
)

// Vertex represents a graph vertex.
type Vertex struct {
	Element
}

// NewVertex create a new graph vertex.
func NewVertex(id any, label string) Vertex {
	if label == "" {
		label = "vertex"
	}
	return Vertex{
		Element: NewElement(id, label),
	}
}

// GraphsonType implements graphson.Typer interface.
func (Vertex) GraphsonType() graphson.Type {
	return "g:Vertex"
}

// String implements fmt.Stringer interface.
func (v Vertex) String() string {
	return fmt.Sprintf("v[%v]", v.ID)
}

// VertexProperty denotes a key/value pair associated with a vertex.
type VertexProperty struct {
	ID    any    `json:"id"`
	Key   string `json:"label"`
	Value any    `json:"value"`
}

// NewVertexProperty create a new graph vertex property.
func NewVertexProperty(id any, key string, value any) VertexProperty {
	return VertexProperty{
		ID:    id,
		Key:   key,
		Value: value,
	}
}

// GraphsonType implements graphson.Typer interface.
func (VertexProperty) GraphsonType() graphson.Type {
	return "g:VertexProperty"
}

// String implements fmt.Stringer interface.
func (vp VertexProperty) String() string {
	return fmt.Sprintf("vp[%s->%v]", vp.Key, vp.Value)
}
