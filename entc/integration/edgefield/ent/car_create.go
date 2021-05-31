// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/schema/field"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
}

// SetNumber sets the "number" field.
func (cc *CarCreate) SetNumber(s string) *CarCreate {
	cc.mutation.SetNumber(s)
	return cc
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cc *CarCreate) SetNillableNumber(s *string) *CarCreate {
	if s != nil {
		cc.SetNumber(*s)
	}
	return cc
}

// AddRentalIDs adds the "rentals" edge to the Rental entity by IDs.
func (cc *CarCreate) AddRentalIDs(ids ...int) *CarCreate {
	cc.mutation.AddRentalIDs(ids...)
	return cc
}

// AddRentals adds the "rentals" edges to the Rental entity.
func (cc *CarCreate) AddRentals(r ...*Rental) *CarCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRentalIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (cc *CarCreate) Mutation() *CarMutation {
	return cc.mutation
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	var (
		err  error
		node *Car
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CarCreate) check() error {
	return nil
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CarCreate) createSpec() (*Car, *sqlgraph.CreateSpec) {
	var (
		_node = &Car{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: car.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Number(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: car.FieldNumber,
		})
		_node.Number = value
	}
	if nodes := cc.mutation.RentalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rental.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CarCreateBulk is the builder for creating many Car entities in bulk.
type CarCreateBulk struct {
	config
	builders []*CarCreate
}

// Save creates the Car entities in the database.
func (ccb *CarCreateBulk) Save(ctx context.Context) ([]*Car, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Car, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CarCreateBulk) SaveX(ctx context.Context) []*Car {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
