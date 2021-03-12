// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/mixinid"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MixinIDCreate is the builder for creating a MixinID entity.
type MixinIDCreate struct {
	config
	mutation *MixinIDMutation
	hooks    []Hook
}

// SetSomeField sets the "some_field" field.
func (mic *MixinIDCreate) SetSomeField(s string) *MixinIDCreate {
	mic.mutation.SetSomeField(s)
	return mic
}

// SetMixinField sets the "mixin_field" field.
func (mic *MixinIDCreate) SetMixinField(s string) *MixinIDCreate {
	mic.mutation.SetMixinField(s)
	return mic
}

// SetID sets the "id" field.
func (mic *MixinIDCreate) SetID(u uuid.UUID) *MixinIDCreate {
	mic.mutation.SetID(u)
	return mic
}

// Mutation returns the MixinIDMutation object of the builder.
func (mic *MixinIDCreate) Mutation() *MixinIDMutation {
	return mic.mutation
}

// Save creates the MixinID in the database.
func (mic *MixinIDCreate) Save(ctx context.Context) (*MixinID, error) {
	var (
		err  error
		node *MixinID
	)
	mic.defaults()
	if len(mic.hooks) == 0 {
		if err = mic.check(); err != nil {
			return nil, err
		}
		node, err = mic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MixinIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mic.check(); err != nil {
				return nil, err
			}
			mic.mutation = mutation
			node, err = mic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mic.hooks) - 1; i >= 0; i-- {
			mut = mic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MixinIDCreate) SaveX(ctx context.Context) *MixinID {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (mic *MixinIDCreate) defaults() {
	if _, ok := mic.mutation.ID(); !ok {
		v := mixinid.DefaultID()
		mic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mic *MixinIDCreate) check() error {
	if _, ok := mic.mutation.SomeField(); !ok {
		return &ValidationError{Name: "some_field", err: errors.New("ent: missing required field \"some_field\"")}
	}
	if _, ok := mic.mutation.MixinField(); !ok {
		return &ValidationError{Name: "mixin_field", err: errors.New("ent: missing required field \"mixin_field\"")}
	}
	return nil
}

// OnConflict specifies how to handle inserts that conflict with a unique constraint on MixinID entities in the database.
func (mic *MixinIDCreate) OnConflict(fields ...string) *MixinIDCreate {
	return mic
}

func (mic *MixinIDCreate) sqlSave(ctx context.Context) (*MixinID, error) {
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (mic *MixinIDCreate) createSpec() (*MixinID, *sqlgraph.CreateSpec) {
	var (
		_node = &MixinID{config: mic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mixinid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		}
	)
	if id, ok := mic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mic.mutation.SomeField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldSomeField,
		})
		_node.SomeField = value
	}
	if value, ok := mic.mutation.MixinField(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldMixinField,
		})
		_node.MixinField = value
	}
	return _node, _spec
}

// MixinIDCreateBulk is the builder for creating many MixinID entities in bulk.
type MixinIDCreateBulk struct {
	config
	builders []*MixinIDCreate
}

// OnConflict specifies how to handle bulk inserts that conflict with a unique constraint on MixinID entities in the database.
func (micb *MixinIDCreateBulk) OnConflict(fields ...string) *MixinIDCreateBulk {
	// for i := range micb.builders {

	// }

	return micb
}

// Save creates the MixinID entities in the database.
func (micb *MixinIDCreateBulk) Save(ctx context.Context) ([]*MixinID, error) {
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MixinID, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MixinIDMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MixinIDCreateBulk) SaveX(ctx context.Context) []*MixinID {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
