// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/dynamodb/dynamodbgraph"
	"entgo.io/ent/examples/dynamodb/o2o2types/ent/card"
	"entgo.io/ent/examples/dynamodb/o2o2types/ent/user"
	"entgo.io/ent/schema/field"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
}

// SetExpired sets the "expired" field.
func (cc *CardCreate) SetExpired(t time.Time) *CardCreate {
	cc.mutation.SetExpired(t)
	return cc
}

// SetNumber sets the "number" field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.mutation.SetNumber(s)
	return cc
}

// SetID sets the "id" field.
func (cc *CardCreate) SetID(i int) *CardCreate {
	cc.mutation.SetID(i)
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CardCreate) SetOwnerID(id int) *CardCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(id *int) *CardCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	var (
		err  error
		node *Card
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.dynamodbSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.dynamodbSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Card)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CardCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CardCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.Expired(); !ok {
		return &ValidationError{Name: "expired", err: errors.New(`ent: missing required field "Card.expired"`)}
	}
	if _, ok := cc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Card.number"`)}
	}
	return nil
}

func (cc *CardCreate) dynamodbSave(ctx context.Context) (*Card, error) {
	_node, _spec := cc.createSpec()
	if err := dynamodbgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		return nil, err
	}
	return _node, nil
}

func (cc *CardCreate) createSpec() (*Card, *dynamodbgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = &dynamodbgraph.CreateSpec{
			Table: card.Table,
			ID: &dynamodbgraph.FieldSpec{
				Type: field.TypeInt,
				Key:  card.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Expired(); ok {
		_spec.Fields = append(_spec.Fields, &dynamodbgraph.FieldSpec{
			Type:  field.TypeTime,
			Value: value,
			Key:   card.FieldExpired,
		})
		_node.Expired = value
	}
	if value, ok := cc.mutation.Number(); ok {
		_spec.Fields = append(_spec.Fields, &dynamodbgraph.FieldSpec{
			Type:  field.TypeString,
			Value: value,
			Key:   card.FieldNumber,
		})
		_node.Number = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &dynamodbgraph.EdgeSpec{
			Rel:        dynamodbgraph.O2O,
			Inverse:    true,
			Table:      card.OwnerTable,
			Attributes: []string{card.OwnerAttribute},
			Bidi:       false,
			Target: &dynamodbgraph.EdgeTarget{
				IDSpec: &dynamodbgraph.FieldSpec{
					Type: field.TypeInt,
					Key:  user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_card = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}
