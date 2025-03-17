// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	schemadir "entgo.io/ent/entc/integration/ent/schema/dir"
	"entgo.io/ent/entc/integration/gremlin/ent/comment"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetUniqueInt sets the "unique_int" field.
func (m *CommentCreate) SetUniqueInt(v int) *CommentCreate {
	m.mutation.SetUniqueInt(v)
	return m
}

// SetUniqueFloat sets the "unique_float" field.
func (m *CommentCreate) SetUniqueFloat(v float64) *CommentCreate {
	m.mutation.SetUniqueFloat(v)
	return m
}

// SetNillableInt sets the "nillable_int" field.
func (m *CommentCreate) SetNillableInt(v int) *CommentCreate {
	m.mutation.SetNillableInt(v)
	return m
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (m *CommentCreate) SetNillableNillableInt(v *int) *CommentCreate {
	if v != nil {
		m.SetNillableInt(*v)
	}
	return m
}

// SetTable sets the "table" field.
func (m *CommentCreate) SetTable(v string) *CommentCreate {
	m.mutation.SetTable(v)
	return m
}

// SetNillableTable sets the "table" field if the given value is not nil.
func (m *CommentCreate) SetNillableTable(v *string) *CommentCreate {
	if v != nil {
		m.SetTable(*v)
	}
	return m
}

// SetDir sets the "dir" field.
func (m *CommentCreate) SetDir(v schemadir.Dir) *CommentCreate {
	m.mutation.SetDir(v)
	return m
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (m *CommentCreate) SetNillableDir(v *schemadir.Dir) *CommentCreate {
	if v != nil {
		m.SetDir(*v)
	}
	return m
}

// SetClient sets the "client" field.
func (m *CommentCreate) SetClient(v string) *CommentCreate {
	m.mutation.SetClient(v)
	return m
}

// SetNillableClient sets the "client" field if the given value is not nil.
func (m *CommentCreate) SetNillableClient(v *string) *CommentCreate {
	if v != nil {
		m.SetClient(*v)
	}
	return m
}

// Mutation returns the CommentMutation object of the builder.
func (m *CommentCreate) Mutation() *CommentMutation {
	return m.mutation
}

// Save creates the Comment in the database.
func (c *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	return withHooks(ctx, c.gremlinSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *CommentCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *CommentCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *CommentCreate) check() error {
	if _, ok := c.mutation.UniqueInt(); !ok {
		return &ValidationError{Name: "unique_int", err: errors.New(`ent: missing required field "Comment.unique_int"`)}
	}
	if _, ok := c.mutation.UniqueFloat(); !ok {
		return &ValidationError{Name: "unique_float", err: errors.New(`ent: missing required field "Comment.unique_float"`)}
	}
	return nil
}

func (c *CommentCreate) gremlinSave(ctx context.Context) (*Comment, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := c.gremlin().Query()
	if err := c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Comment{config: c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	c.mutation.id = &rnode.ID
	c.mutation.done = true
	return rnode, nil
}

func (c *CommentCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(comment.Label)
	if value, ok := c.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := c.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := c.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
	}
	if value, ok := c.mutation.Table(); ok {
		v.Property(dsl.Single, comment.FieldTable, value)
	}
	if value, ok := c.mutation.Dir(); ok {
		v.Property(dsl.Single, comment.FieldDir, value)
	}
	if value, ok := c.mutation.GetClient(); ok {
		v.Property(dsl.Single, comment.FieldClient, value)
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}
