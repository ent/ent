// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package ent is the interface between end-user schemas and entc (ent codegen).
package ent

import (
	"context"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		//
		// Deprecated: the Config method predates the Annotations method, and it
		// is planned be removed in v0.5.0. New code should use Annotations instead.
		//
		//	func (T) Annotations() []schema.Annotation {
		//		return []schema.Annotation{
		//			entsql.Annotation{Table: "Name"},
		//		}
		//	}
		//
		Config() Config
		// Mixin returns an optional list of Mixin to extends
		// the schema.
		Mixin() []Mixin
		// Hooks returns an optional list of Hook to apply on
		// the executed mutations.
		Hooks() []Hook
		// Interceptors returns an optional list of Interceptor
		// to apply on the executed queries.
		Interceptors() []Interceptor
		// Policy returns the privacy policy of the schema.
		Policy() Policy
		// Annotations returns a list of schema annotations to be used by
		// codegen extensions.
		Annotations() []schema.Annotation
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

	// An Edge interface returns an edge descriptor for vertex edges.
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

	// An Index interface returns an index descriptor for vertex indexes.
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
	// Deprecated: the Config object predates the schema.Annotation method and it
	// is planned be removed in v0.5.0. New code should use Annotations instead.
	//
	//	func (T) Annotations() []schema.Annotation {
	//		return []schema.Annotation{
	//			entsql.Annotation{Table: "Name"},
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
		// Interceptors returns a slice of interceptors to add to the schema.
		// Note that mixin interceptors are executed before schema interceptors.
		Interceptors() []Interceptor
		// Policy returns a privacy policy to add to the schema.
		// Note that mixin policy are executed before schema policy.
		Policy() Policy
		// Annotations returns a list of schema annotations to add
		// to the schema annotations.
		Annotations() []schema.Annotation
	}

	// The Policy type defines the privacy policy of an entity.
	// The usage for the interface is as follows:
	//
	//	type T struct {
	//		ent.Schema
	//	}
	//
	//	func(T) Policy() ent.Policy {
	//		return privacy.AlwaysAllowRule()
	//	}
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

// Interceptors of the schema.
func (Schema) Interceptors() []Interceptor { return nil }

// Policy of the schema.
func (Schema) Policy() Policy { return nil }

// Annotations of the schema.
func (Schema) Annotations() []schema.Annotation { return nil }

type (
	// Value represents a dynamic value returned by mutations or queries.
	Value any

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
		// FieldCleared returns a bool indicates if this field was
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
		// EdgeCleared returns a bool indicates if this edge was
		// cleared in this mutation.
		EdgeCleared(name string) bool
		// ClearEdge clears the value for the given name. It returns an
		// error if the edge name is not defined in the schema.
		ClearEdge(name string) error

		// ResetEdge resets all changes in the mutation regarding the
		// given edge name. It returns an error if the edge is not
		// defined in the schema.
		ResetEdge(name string) error

		// OldField returns the old value of the field from the database.
		// An error is returned if the mutation operation is not UpdateOne,
		// or the query to the database was failed.
		OldField(ctx context.Context, name string) (Value, error)
	}

	// Mutator is the interface that wraps the Mutate method.
	Mutator interface {
		// Mutate apply the given mutation on the graph. The returned
		// ent.Value is changing according to the mutation operation:
		//
		// OpCreate, the returned value is the created node (T).
		// OpUpdateOne, the returned value is the updated node (T).
		// OpUpdate, the returned value is the amount of updated nodes (int).
		// OpDeleteOne, OpDelete, the returned value is the amount of deleted nodes (int).
		//
		Mutate(context.Context, Mutation) (Value, error)
	}

	// The MutateFunc type is an adapter to allow the use of ordinary
	// function as Mutator. If f is a function with the appropriate signature,
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

type (
	// Query represents a query builder of an entity. It is
	// usually one of the following types: <T>Query.
	Query any

	// Querier is the interface that wraps the Query method.
	// Calling Querier.Query(ent.Query) triggers the execution
	// of the query.
	Querier interface {
		// Query runs the given query on the graph and returns its result.
		Query(context.Context, Query) (Value, error)
	}

	// The QuerierFunc type is an adapter to allow the use of ordinary
	// function as Querier. If f is a function with the appropriate signature,
	// QuerierFunc(f) is a Querier that calls f.
	QuerierFunc func(context.Context, Query) (Value, error)

	// Interceptor defines an execution middleware for various types of Ent queries.
	// Contrary to Hooks, Interceptors are implemented as interfaces, allows them to
	// intercept and modify the query at different stages, providing more fine-grained
	// control over its behavior. For example, see the Traverser interface.
	Interceptor interface {
		// Intercept is a function that gets a Querier and returns a Querier. For example:
		//
		//	ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		//		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
		//			// Do something before the query execution.
		//			value, err := next.Query(ctx, query)
		//			// Do something after the query execution.
		//			return value, err
		//		})
		//	})
		//
		// Note that unlike Traverse functions, which are called at each traversal stage, Intercept functions
		// are invoked before the query executions. This means that using Traverse functions is a better fit
		// for adding default filters, while using Intercept functions is a better fit for implementing logging
		// or caching.
		//
		//
		//	client.User.Query().
		//		QueryGroups().	// User traverse functions applied.
		//		QueryPosts().	// Group traverse functions applied.
		//		All(ctx)	// Post traverse and intercept functions applied.
		//
		Intercept(Querier) Querier
	}

	// The InterceptFunc type is an adapter to allow the use of ordinary function as Interceptor.
	// If f is a function with the appropriate signature, InterceptFunc(f) is an Interceptor that calls f.
	InterceptFunc func(Querier) Querier

	// Traverser defines a graph-traversal middleware for various types of Ent queries.
	// Contrary to Interceptors, the Traverse are executed on graph traversals before the
	// query is executed. For example:
	//
	//	ent.TraverseFunc(func(ctx context.Context, q ent.Query) error {
	//		// Filter out deleted pets.
	//		if pq, ok := q.(*gen.PetQuery); ok {
	//			pq.Where(pet.DeletedAtIsNil())
	//		}
	//		return nil
	//	})
	//
	//	client.Pet.Query().
	//		QueryOwner().	// Pet traverse functions are applied and filter deleted pets.
	//		All(ctx)	// User traverse and interceptor functions are applied.
	//
	Traverser interface {
		Traverse(context.Context, Query) error
	}

	// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
	// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
	TraverseFunc func(context.Context, Query) error
)

// Query calls f(ctx, q).
func (f QuerierFunc) Query(ctx context.Context, q Query) (Value, error) {
	return f(ctx, q)
}

// Intercept calls f(ctx, q).
func (f InterceptFunc) Intercept(next Querier) Querier {
	return f(next)
}

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next Querier) Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q Query) error {
	return f(ctx, q)
}

//go:generate go run golang.org/x/tools/cmd/stringer -type Op

// An Op represents a mutation operation.
type Op uint

// Mutation operations.
const (
	OpCreate    Op = 1 << iota // node creation.
	OpUpdate                   // update nodes by predicate (if any).
	OpUpdateOne                // update one node.
	OpDelete                   // delete nodes by predicate (if any).
	OpDeleteOne                // delete one node.
)

// Is reports whether o is match the given operation.
func (i Op) Is(o Op) bool { return i&o != 0 }

type (
	// QueryContext contains additional information about
	// the context in which the query is executed.
	QueryContext struct {
		// Op defines the operation name. e.g., First, All, Count, etc.
		Op string
		// Type defines the query type as defined in the generated code.
		Type string
		// Unique indicates if the Unique modifier was set on the query and
		// its value. Calling Unique(false) sets the value of Unique to false.
		Unique *bool
		// Limit indicates if the Limit modifier was set on the query and
		// its value. Calling Limit(10) sets the value of Limit to 10.
		Limit *int
		// Offset indicates if the Offset modifier was set on the query and
		// its value. Calling Offset(10) sets the value of Offset to 10.
		Offset *int
		// Fields specifies the fields that were selected in the query.
		Fields []string
	}
	queryCtxKey struct{}
)

// NewQueryContext returns a new context with the given QueryContext attached.
func NewQueryContext(parent context.Context, c *QueryContext) context.Context {
	return context.WithValue(parent, queryCtxKey{}, c)
}

// QueryFromContext returns the QueryContext value stored in ctx, if any.
func QueryFromContext(ctx context.Context) *QueryContext {
	c, _ := ctx.Value(queryCtxKey{}).(*QueryContext)
	return c
}

// Clone returns a deep copy of the query context.
func (q *QueryContext) Clone() *QueryContext {
	c := &QueryContext{
		Op:     q.Op,
		Type:   q.Type,
		Fields: append([]string(nil), q.Fields...),
	}
	if q.Unique != nil {
		v := *q.Unique
		c.Unique = &v
	}
	if q.Limit != nil {
		v := *q.Limit
		c.Limit = &v
	}
	if q.Offset != nil {
		v := *q.Offset
		c.Offset = &v
	}
	return c
}

// AppendFieldOnce adds the given field to the spec if it is not already present.
func (q *QueryContext) AppendFieldOnce(f string) *QueryContext {
	for _, f1 := range q.Fields {
		if f == f1 {
			return q
		}
	}
	q.Fields = append(q.Fields, f)
	return q
}
