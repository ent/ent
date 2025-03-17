// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/customtype"
	"entgo.io/ent/schema/field"
)

// CustomTypeCreate is the builder for creating a CustomType entity.
type CustomTypeCreate struct {
	config
	mutation *CustomTypeMutation
	hooks    []Hook
}

// SetCustom sets the "custom" field.
func (_c *CustomTypeCreate) SetCustom(s string) *CustomTypeCreate {
	_c.mutation.SetCustom(s)
	return _c
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (_c *CustomTypeCreate) SetNillableCustom(s *string) *CustomTypeCreate {
	if s != nil {
		_c.SetCustom(*s)
	}
	return _c
}

// SetTz0 sets the "tz0" field.
func (_c *CustomTypeCreate) SetTz0(t time.Time) *CustomTypeCreate {
	_c.mutation.SetTz0(t)
	return _c
}

// SetNillableTz0 sets the "tz0" field if the given value is not nil.
func (_c *CustomTypeCreate) SetNillableTz0(t *time.Time) *CustomTypeCreate {
	if t != nil {
		_c.SetTz0(*t)
	}
	return _c
}

// SetTz3 sets the "tz3" field.
func (_c *CustomTypeCreate) SetTz3(t time.Time) *CustomTypeCreate {
	_c.mutation.SetTz3(t)
	return _c
}

// SetNillableTz3 sets the "tz3" field if the given value is not nil.
func (_c *CustomTypeCreate) SetNillableTz3(t *time.Time) *CustomTypeCreate {
	if t != nil {
		_c.SetTz3(*t)
	}
	return _c
}

// Mutation returns the CustomTypeMutation object of the builder.
func (_c *CustomTypeCreate) Mutation() *CustomTypeMutation {
	return _c.mutation
}

// Save creates the CustomType in the database.
func (_c *CustomTypeCreate) Save(ctx context.Context) (*CustomType, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *CustomTypeCreate) SaveX(ctx context.Context) *CustomType {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *CustomTypeCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *CustomTypeCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *CustomTypeCreate) check() error {
	return nil
}

func (_c *CustomTypeCreate) sqlSave(ctx context.Context) (*CustomType, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *CustomTypeCreate) createSpec() (*CustomType, *sqlgraph.CreateSpec) {
	var (
		_node = &CustomType{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(customtype.Table, sqlgraph.NewFieldSpec(customtype.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Custom(); ok {
		_spec.SetField(customtype.FieldCustom, field.TypeString, value)
		_node.Custom = value
	}
	if value, ok := _c.mutation.Tz0(); ok {
		_spec.SetField(customtype.FieldTz0, field.TypeTime, value)
		_node.Tz0 = value
	}
	if value, ok := _c.mutation.Tz3(); ok {
		_spec.SetField(customtype.FieldTz3, field.TypeTime, value)
		_node.Tz3 = value
	}
	return _node, _spec
}

// CustomTypeCreateBulk is the builder for creating many CustomType entities in bulk.
type CustomTypeCreateBulk struct {
	config
	err      error
	builders []*CustomTypeCreate
}

// Save creates the CustomType entities in the database.
func (_c *CustomTypeCreateBulk) Save(ctx context.Context) ([]*CustomType, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*CustomType, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *CustomTypeCreateBulk) SaveX(ctx context.Context) []*CustomType {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *CustomTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *CustomTypeCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
