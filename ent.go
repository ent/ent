// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package ent is the interface between end-user schemas and entc (ent codegen).
package ent

import (
	"context"

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
		// Hooks returns an optional list of Hook to apply on
		// mutations.
		Hooks() []Hook
		// Policy returns the privacy policy of the schema.
		Policy() Policy
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
		// Fields returns a slice of fields to add to the schema.
		Fields() []Field
		// Edges returns a slice of edges to add to the schema.
		Edges() []Edge
		// Indexes returns a slice of indexes to add to the schema.
		Indexes() []Index
		// Hooks returns a slice of hooks to add to the schema.
		// Note that mixin hooks are executed before schema hooks.
		Hooks() []Hook
	}

	// The Policy type defines the write privacy policy of an entity.
	// The usage for the interface is as follows:
	//
	// type T struct {
	//   ent.Schema
	// }
	//
	// func(T) Policy() ent.Policy {
	//     return privacy.AlwaysAllowReadWrite()
	// }
	//
	//
	Policy interface {
		EvalMutation(context.Context, Mutation) error
		EvalQuery(context.Context, Query) error
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

// Hooks of the schema.
func (Schema) Hooks() []Hook { return nil }

// Policy of the schema.
func (Schema) Policy() Policy { return nil }

type (
	// Value represents a value returned by ent.
	Value interface{}
	// Query represents an ent query builder.
	Query interface{}
	// Mutation represents an operation that mutate the graph.
	// For example, adding a new node, updating many, or dropping
	// data. The implementation is generated by entc (ent codegen).
	Mutation interface {
		// Op returns the operation name generated by entc.
		Op() Op
		// Type returns the schema type for this mutation.
		Type() string

		// Fields returns all fields that were changed during
		// this mutation. Note that, in order to get all numeric
		// fields that were in/decremented, call AddedFields().
		Fields() []string
		// Field returns the value of a field with the given name.
		// The second boolean value indicates that this field was
		// not set, or was not defined in the schema.
		Field(name string) (Value, bool)
		// SetField sets the value for the given name. It returns an
		// error if the field is not defined in the schema, or if the
		// type mismatch the field type.
		SetField(name string, value Value) error

		// AddedFields returns all numeric fields that were incremented
		// or decremented during this mutation.
		AddedFields() []string
		// AddedField returns the numeric value that was in/decremented
		// from a field with the given name. The second value indicates
		// that this field was not set, or was not define in the schema.
		AddedField(name string) (Value, bool)
		// AddField adds the value for the given name. It returns an
		// error if the field is not defined in the schema, or if the
		// type mismatch the field type.
		AddField(name string, value Value) error

		// ClearedFields returns all nullable fields that were cleared
		// during this mutation.
		ClearedFields() []string
		// FieldCleared returns a boolean indicates if this field was
		// cleared in this mutation.
		FieldCleared(name string) bool
		// ClearField clears the value for the given name. It returns an
		// error if the field is not defined in the schema.
		ClearField(name string) error

		// ResetField resets all changes in the mutation regarding the
		// given field name. It returns an error if the field is not
		// defined in the schema.
		ResetField(name string) error

		// AddedEdges returns all edge names that were set/added in this
		// mutation.
		AddedEdges() []string
		// AddedIDs returns all ids (to other nodes) that were added for
		// the given edge name.
		AddedIDs(name string) []Value

		// RemovedEdges returns all edge names that were removed in this
		// mutation.
		RemovedEdges() []string
		// RemovedIDs returns all ids (to other nodes) that were removed for
		// the given edge name.
		RemovedIDs(name string) []Value

		// ClearedEdges returns all edge names that were cleared in this
		// mutation.
		ClearedEdges() []string
		// EdgeCleared returns a boolean indicates if this edge was
		// cleared in this mutation.
		EdgeCleared(name string) bool
		// ClearEdge clears the value for the given name. It returns an
		// error if the edge name is not defined in the schema.
		ClearEdge(name string) error

		// ResetEdge resets all changes in the mutation regarding the
		// given edge name. It returns an error if the edge is not
		// defined in the schema.
		ResetEdge(name string) error

		// In order to not break users code, we release the codegen part
		// first, and uncomment the new method after a minor version release.
		//
		// OldField returns the old value of the field from the database.
		// An error is returned if the mutation operation is not UpdateOne,
		// or the query to the database was failed.
		//
		// OldField(ctx context.Context, name string) (Value, error)
	}

	// Mutator is the interface that wraps the Mutate method.
	Mutator interface {
		// Mutate apply the given mutation on the graph.
		Mutate(context.Context, Mutation) (Value, error)
	}

	// The MutateFunc type is an adapter to allow the use of ordinary
	// function as mutator. If f is a function with the appropriate signature,
	// MutateFunc(f) is a Mutator that calls f.
	MutateFunc func(context.Context, Mutation) (Value, error)

	// Hook defines the "mutation middleware". A function that gets a Mutator
	// and returns a Mutator. For example:
	//
	//	hook := func(next ent.Mutator) ent.Mutator {
	//		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	//			fmt.Printf("Type: %s, Operation: %s, ConcreteType: %T\n", m.Type(), m.Op(), m)
	//			return next.Mutate(ctx, m)
	//		})
	//	}
	//
	Hook func(Mutator) Mutator
)

// Mutate calls f(ctx, m).
func (f MutateFunc) Mutate(ctx context.Context, m Mutation) (Value, error) {
	return f(ctx, m)
}

// An Op represents a mutation operation.
type Op uint

// Mutation operations.
const (
	OpCreate    Op = 1 << iota // node creation.
	OpUpdate                   // update nodes by predicate (if any).
	OpUpdateOne                // update one node.
	OpDelete                   // delete nodes by predicate (if any).
	OpDeleteOne                // delete one one.
)

// Is reports whether o is match the given operation.
func (i Op) Is(o Op) bool { return i&o != 0 }

//go:generate go run golang.org/x/tools/cmd/stringer -type Op
