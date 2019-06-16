// Package ent is an interface package for the schemas that use entc.
package ent

import (
	"context"

	"fbc/ent/dialect/sql"
	"fbc/ent/edge"
	"fbc/ent/field"
	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
)

type (
	// Schema is the interface for describing an entity schema for entc.
	Schema interface {
		Type()
		Edges() []Edge
		Fields() []Field
	}

	// Field is the interface for vertex and edges fields used by the code generation.
	Field interface {
		Tag() string
		Name() string
		Type() field.Type
		IsUnique() bool
		IsNullable() bool
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
		IsAssoc() bool
		IsUnique() bool
		IsInverse() bool
		IsRequired() bool
		Assoc() *edge.Edge
	}

	// Execer is the driver for executing gremlin requests.
	Execer interface {
		Exec(context.Context, string, interface{}) (*gremlin.Response, error)
	}

	// Predicate applies condition changes on either graph traversal or sql selector.
	Predicate struct {
		SQL     func(*sql.Selector)  // common sql.
		Gremlin func(*dsl.Traversal) // common gremlin.
	}
)

// DefaultSchema holds the default schema implementation.
var DefaultSchema = defaultSchema{}

// defaultSchema is the default implementation for the schema.
type defaultSchema struct{ Schema }

func (defaultSchema) Edges() []Edge { return nil }

func (defaultSchema) Fields() []Field { return nil }
