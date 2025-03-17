// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/zoo"
	"entgo.io/ent/schema/field"
)

// ZooCreate is the builder for creating a Zoo entity.
type ZooCreate struct {
	config
	mutation *ZooMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (m *ZooCreate) SetID(v int) *ZooCreate {
	m.mutation.SetID(v)
	return m
}

// Mutation returns the ZooMutation object of the builder.
func (m *ZooCreate) Mutation() *ZooMutation {
	return m.mutation
}

// Save creates the Zoo in the database.
func (c *ZooCreate) Save(ctx context.Context) (*Zoo, error) {
	return withHooks(ctx, c.sqlSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *ZooCreate) SaveX(ctx context.Context) *Zoo {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *ZooCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *ZooCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *ZooCreate) check() error {
	return nil
}

func (c *ZooCreate) sqlSave(ctx context.Context) (*Zoo, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	_node, _spec := c.createSpec()
	if err := sqlgraph.CreateNode(ctx, c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	c.mutation.id = &_node.ID
	c.mutation.done = true
	return _node, nil
}

func (c *ZooCreate) createSpec() (*Zoo, *sqlgraph.CreateSpec) {
	var (
		_node = &Zoo{config: c.config}
		_spec = sqlgraph.NewCreateSpec(zoo.Table, sqlgraph.NewFieldSpec(zoo.FieldID, field.TypeInt))
	)
	if id, ok := c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	return _node, _spec
}

// ZooCreateBulk is the builder for creating many Zoo entities in bulk.
type ZooCreateBulk struct {
	config
	err      error
	builders []*ZooCreate
}

// Save creates the Zoo entities in the database.
func (c *ZooCreateBulk) Save(ctx context.Context) ([]*Zoo, error) {
	if c.err != nil {
		return nil, c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(c.builders))
	nodes := make([]*Zoo, len(c.builders))
	mutators := make([]Mutator, len(c.builders))
	for i := range c.builders {
		func(i int, root context.Context) {
			builder := c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ZooMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (c *ZooCreateBulk) SaveX(ctx context.Context) []*Zoo {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *ZooCreateBulk) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *ZooCreateBulk) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}
