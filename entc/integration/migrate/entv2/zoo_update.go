// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/entc/integration/migrate/entv2/zoo"
	"entgo.io/ent/schema/field"
)

// ZooUpdate is the builder for updating Zoo entities.
type ZooUpdate struct {
	config
	hooks    []Hook
	mutation *ZooMutation
}

// Where appends a list predicates to the ZooUpdate builder.
func (zu *ZooUpdate) Where(ps ...predicate.Zoo) *ZooUpdate {
	zu.mutation.Where(ps...)
	return zu
}

// Mutation returns the ZooMutation object of the builder.
func (zu *ZooUpdate) Mutation() *ZooMutation {
	return zu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (zu *ZooUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, zu.sqlSave, zu.mutation, zu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (zu *ZooUpdate) SaveX(ctx context.Context) int {
	affected, err := zu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (zu *ZooUpdate) Exec(ctx context.Context) error {
	_, err := zu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (zu *ZooUpdate) ExecX(ctx context.Context) {
	if err := zu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (zu *ZooUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(zoo.Table, zoo.Columns, sqlgraph.NewFieldSpec(zoo.FieldID, field.TypeInt))
	if ps := zu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, zu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{zoo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	zu.mutation.done = true
	return n, nil
}

// ZooUpdateOne is the builder for updating a single Zoo entity.
type ZooUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ZooMutation
}

// Mutation returns the ZooMutation object of the builder.
func (zuo *ZooUpdateOne) Mutation() *ZooMutation {
	return zuo.mutation
}

// Where appends a list predicates to the ZooUpdate builder.
func (zuo *ZooUpdateOne) Where(ps ...predicate.Zoo) *ZooUpdateOne {
	zuo.mutation.Where(ps...)
	return zuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (zuo *ZooUpdateOne) Select(field string, fields ...string) *ZooUpdateOne {
	zuo.fields = append([]string{field}, fields...)
	return zuo
}

// Save executes the query and returns the updated Zoo entity.
func (zuo *ZooUpdateOne) Save(ctx context.Context) (*Zoo, error) {
	return withHooks(ctx, zuo.sqlSave, zuo.mutation, zuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (zuo *ZooUpdateOne) SaveX(ctx context.Context) *Zoo {
	node, err := zuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (zuo *ZooUpdateOne) Exec(ctx context.Context) error {
	_, err := zuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (zuo *ZooUpdateOne) ExecX(ctx context.Context) {
	if err := zuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (zuo *ZooUpdateOne) sqlSave(ctx context.Context) (_node *Zoo, err error) {
	_spec := sqlgraph.NewUpdateSpec(zoo.Table, zoo.Columns, sqlgraph.NewFieldSpec(zoo.FieldID, field.TypeInt))
	id, ok := zuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv2: missing "Zoo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := zuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, zoo.FieldID)
		for _, f := range fields {
			if !zoo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
			}
			if f != zoo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := zuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Zoo{config: zuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, zuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{zoo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	zuo.mutation.done = true
	return _node, nil
}
