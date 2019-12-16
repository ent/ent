// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/comment"
	"github.com/facebookincubator/ent/schema/field"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	unique_int   *int
	unique_float *float64
	nillable_int *int
}

// SetUniqueInt sets the unique_int field.
func (cc *CommentCreate) SetUniqueInt(i int) *CommentCreate {
	cc.unique_int = &i
	return cc
}

// SetUniqueFloat sets the unique_float field.
func (cc *CommentCreate) SetUniqueFloat(f float64) *CommentCreate {
	cc.unique_float = &f
	return cc
}

// SetNillableInt sets the nillable_int field.
func (cc *CommentCreate) SetNillableInt(i int) *CommentCreate {
	cc.nillable_int = &i
	return cc
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (cc *CommentCreate) SetNillableNillableInt(i *int) *CommentCreate {
	if i != nil {
		cc.SetNillableInt(*i)
	}
	return cc
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	if cc.unique_int == nil {
		return nil, errors.New("ent: missing required field \"unique_int\"")
	}
	if cc.unique_float == nil {
		return nil, errors.New("ent: missing required field \"unique_float\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	var (
		c    = &Comment{config: cc.config}
		spec = &sqlgraph.CreateSpec{
			Table: comment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: comment.FieldID,
			},
		}
	)
	if value := cc.unique_int; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: comment.FieldUniqueInt,
		})
		c.UniqueInt = *value
	}
	if value := cc.unique_float; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  *value,
			Column: comment.FieldUniqueFloat,
		})
		c.UniqueFloat = *value
	}
	if value := cc.nillable_int; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: comment.FieldNillableInt,
		})
		c.NillableInt = value
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	c.ID = strconv.FormatInt(id, 10)
	return c, nil
}
