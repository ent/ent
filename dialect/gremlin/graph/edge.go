// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

import (
	"fmt"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"
)

type (
	// An Edge between two vertices.
	Edge struct {
		Element
		OutV, InV Vertex
	}

	// graphson edge repr.
	edge struct {
		Element
		OutV      any    `json:"outV"`
		OutVLabel string `json:"outVLabel"`
		InV       any    `json:"inV"`
		InVLabel  string `json:"inVLabel"`
	}
)

// NewEdge create a new graph edge.
func NewEdge(id any, label string, outV, inV Vertex) Edge {
	return Edge{
		Element: NewElement(id, label),
		OutV:    outV,
		InV:     inV,
	}
}

// String implements fmt.Stringer interface.
func (e Edge) String() string {
	return fmt.Sprintf("e[%v][%v-%s->%v]", e.ID, e.OutV.ID, e.Label, e.InV.ID)
}

// MarshalGraphson implements graphson.Marshaler interface.
func (e Edge) MarshalGraphson() ([]byte, error) {
	return graphson.Marshal(edge{
		Element:   e.Element,
		OutV:      e.OutV.ID,
		OutVLabel: e.OutV.Label,
		InV:       e.InV.ID,
		InVLabel:  e.InV.Label,
	})
}

// UnmarshalGraphson implements graphson.Unmarshaler interface.
func (e *Edge) UnmarshalGraphson(data []byte) error {
	var edge edge
	if err := graphson.Unmarshal(data, &edge); err != nil {
		return fmt.Errorf("unmarshalling edge: %w", err)
	}

	*e = NewEdge(
		edge.ID, edge.Label,
		NewVertex(edge.OutV, edge.OutVLabel),
		NewVertex(edge.InV, edge.InVLabel),
	)
	return nil
}

// GraphsonType implements graphson.Typer interface.
func (edge) GraphsonType() graphson.Type {
	return "g:Edge"
}

// Property denotes a key/value pair associated with an edge.
type Property struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

// NewProperty create a new graph edge property.
func NewProperty(key string, value any) Property {
	return Property{key, value}
}

// GraphsonType implements graphson.Typer interface.
func (Property) GraphsonType() graphson.Type {
	return "g:Property"
}

// String implements fmt.Stringer interface.
func (p Property) String() string {
	return fmt.Sprintf("p[%s->%v]", p.Key, p.Value)
}
