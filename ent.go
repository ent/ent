// Package ent is an interface package for the schemas that use entc.
package ent

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type (
	// The Interface type describes the requirements for an exported type defined in the schema package.
	// It functions as the interface between the user's schema types and codegen loader.
	// Users should use the Schema type for embedding as follows:
	//
	//	type T struct {
	//		ent.Schema
	//	}
	//
	Interface interface {
		Type()
		Fields() []Field
		Edges() []Edge
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
		IsImmutable() bool
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
		Edges() []string
		Fields() []string
	}

	// Schema is the default implementation Interface.
	Schema struct {
		Interface
	}
)

// Fields of the schema.
func (Schema) Fields() []Field { return nil }

// Edges of the schema.
func (Schema) Edges() []Edge { return nil }

// Indexes of the schema.
func (Schema) Indexes() []Index { return nil }
