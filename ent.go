// Package ent is an interface package for the schemas that use entc.
package ent

import (
	"fbc/ent/edge"
	"fbc/ent/field"
)

type (
	// Schema is the interface for describing an entity schema for entc.
	Schema interface {
		Type()
		Edges() []Edge
		Fields() []Field
		Indexes() []Index
	}

	// Field is the interface for vertex and edges fields used by the code generation.
	Field interface {
		Tag() string
		Name() string
		Type() field.Type
		IsUnique() bool
		IsNillable() bool
		IsOptional() bool
		HasDefault() bool
		Value() interface{}
		Validators() []interface{}
	}

	// Edge is the interface for graph edges in the schema. It is used by the code generation.
	Edge interface {
		Tag() string
		Type() string
		Name() string
		RefName() string
		Assoc() *edge.Edge
		IsUnique() bool
		IsInverse() bool
		IsRequired() bool
	}

	// Index is the interface for graph indexes in the schema. It is used by the code generation.
	Index interface {
		IsUnique() bool
		Edge() string
		Fields() []string
	}
)

// DefaultSchema holds the default schema implementation.
var DefaultSchema defaultSchema

// defaultSchema is the default implementation for the schema.
type defaultSchema struct{ Schema }

func (defaultSchema) Edges() []Edge { return nil }

func (defaultSchema) Fields() []Field { return nil }

func (defaultSchema) Indexes() []Index { return nil }
