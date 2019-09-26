// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package ent is the interface between end-user schemas and entc (ent codegen).
package ent

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
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
		// Type is a dummy method, that is used in edge declaration.
		//
		// The Type method should be used as follows:
		//
		//	type S struct { ent.Schema }
		//
		//	type T struct { ent.Schema }
		//
		//	func (T) Edges() []ent.Edge {
		//		return []ent.Edge{
		//			edge.To("S", S.Type),
		//		}
		//	}
		//
		Type()
		// Fields returns the fields of the schema.
		Fields() []Field
		// Edges returns the edges of the schema.
		Edges() []Edge
		// Indexes returns the indexes of the schema.
		Indexes() []Index
		// Config returns an optional config for the schema.
		Config() Config
		// Mixin returns an optional list of Mixin to extends
		// the schema.
		Mixin() []Mixin
	}

	// A Field interface returns a field descriptor for vertex fields/properties.
	// The usage for the interface is as follows:
	//
	//	func (T) Fields() []ent.Field {
	//		return []ent.Field{
	//			field.Int("int"),
	//		}
	//	}
	//
	Field interface {
		Descriptor() *field.Descriptor
	}

	// A Edge interface returns an edge descriptor for vertex edges.
	// The usage for the interface is as follows:
	//
	//	func (T) Edges() []ent.Edge {
	//		return []ent.Edge{
	//			edge.To("S", S.Type),
	//		}
	//	}
	//
	Edge interface {
		Descriptor() *edge.Descriptor
	}

	// A Index interface returns an index descriptor for vertex indexes.
	// The usage for the interface is as follows:
	//
	//	func (T) Indexes() []ent.Index {
	//		return []ent.Index{
	//			index.Fields("f1", "f2").
	//				Unique(),
	//		}
	//	}
	//
	Index interface {
		Descriptor() *index.Descriptor
	}

	// A Config structure is used to configure an entity schema.
	// The usage of this structure is as follows:
	//
	//	func (T) Config() ent.Config {
	//		return ent.Config{
	//			Table: "Name",
	//		}
	//	}
	//
	Config struct {
		// A Table is an optional table name defined for the schema.
		Table string
	}

	// The Mixin type describes a set of methods that can extend
	// other methods in the schema without calling them directly.
	//
	//	type TimeMixin struct {}
	//
	//	func (TimeMixin) Fields() []ent.Field {
	//		return []ent.Field{
	//			field.Time("created_at").
	//				Immutable().
	//				Default(time.Now),
	//			field.Time("updated_at").
	//				Default(time.Now).
	//				UpdateDefault(time.Now),
	//		}
	//	}
	//
	//	type T struct {
	//		ent.Schema
	//	}
	//
	// 	func(T) Mixin() []ent.Mixin {
	// 		return []ent.Mixin{
	//			TimeMixin{},
	// 		}
	// 	}
	//
	Mixin interface {
		// Fields returns a slice of fields to be added
		// to the schema fields.
		Fields() []Field
	}

	// Schema is the default implementation for the schema Interface.
	// It can be embedded in end-user schemas as follows:
	//
	//	type T struct {
	//		ent.Schema
	//	}
	//
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

// Config of the schema.
func (Schema) Config() Config { return Config{} }

// Mixin of the schema.
func (Schema) Mixin() []Mixin { return nil }
