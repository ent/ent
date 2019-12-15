// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/comment"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	unique_int        *int
	addunique_int     *int
	unique_float      *float64
	addunique_float   *float64
	nillable_int      *int
	addnillable_int   *int
	clearnillable_int bool
	predicates        []predicate.Comment
}

// Where adds a new predicate for the builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetUniqueInt sets the unique_int field.
func (cu *CommentUpdate) SetUniqueInt(i int) *CommentUpdate {
	cu.unique_int = &i
	cu.addunique_int = nil
	return cu
}

// AddUniqueInt adds i to unique_int.
func (cu *CommentUpdate) AddUniqueInt(i int) *CommentUpdate {
	if cu.addunique_int == nil {
		cu.addunique_int = &i
	} else {
		*cu.addunique_int += i
	}
	return cu
}

// SetUniqueFloat sets the unique_float field.
func (cu *CommentUpdate) SetUniqueFloat(f float64) *CommentUpdate {
	cu.unique_float = &f
	cu.addunique_float = nil
	return cu
}

// AddUniqueFloat adds f to unique_float.
func (cu *CommentUpdate) AddUniqueFloat(f float64) *CommentUpdate {
	if cu.addunique_float == nil {
		cu.addunique_float = &f
	} else {
		*cu.addunique_float += f
	}
	return cu
}

// SetNillableInt sets the nillable_int field.
func (cu *CommentUpdate) SetNillableInt(i int) *CommentUpdate {
	cu.nillable_int = &i
	cu.addnillable_int = nil
	return cu
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (cu *CommentUpdate) SetNillableNillableInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetNillableInt(*i)
	}
	return cu
}

// AddNillableInt adds i to nillable_int.
func (cu *CommentUpdate) AddNillableInt(i int) *CommentUpdate {
	if cu.addnillable_int == nil {
		cu.addnillable_int = &i
	} else {
		*cu.addnillable_int += i
	}
	return cu
}

// ClearNillableInt clears the value of nillable_int.
func (cu *CommentUpdate) ClearNillableInt() *CommentUpdate {
	cu.nillable_int = nil
	cu.clearnillable_int = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(cu.driver.Dialect())
		selector = builder.Select(comment.FieldID).From(builder.Table(comment.Table))
	)
	for _, p := range cu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := cu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(comment.Table)
	)
	updater = updater.Where(sql.InInts(comment.FieldID, ids...))
	if value := cu.unique_int; value != nil {
		updater.Set(comment.FieldUniqueInt, *value)
	}
	if value := cu.addunique_int; value != nil {
		updater.Add(comment.FieldUniqueInt, *value)
	}
	if value := cu.unique_float; value != nil {
		updater.Set(comment.FieldUniqueFloat, *value)
	}
	if value := cu.addunique_float; value != nil {
		updater.Add(comment.FieldUniqueFloat, *value)
	}
	if value := cu.nillable_int; value != nil {
		updater.Set(comment.FieldNillableInt, *value)
	}
	if value := cu.addnillable_int; value != nil {
		updater.Add(comment.FieldNillableInt, *value)
	}
	if cu.clearnillable_int {
		updater.SetNull(comment.FieldNillableInt)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	id                string
	unique_int        *int
	addunique_int     *int
	unique_float      *float64
	addunique_float   *float64
	nillable_int      *int
	addnillable_int   *int
	clearnillable_int bool
}

// SetUniqueInt sets the unique_int field.
func (cuo *CommentUpdateOne) SetUniqueInt(i int) *CommentUpdateOne {
	cuo.unique_int = &i
	cuo.addunique_int = nil
	return cuo
}

// AddUniqueInt adds i to unique_int.
func (cuo *CommentUpdateOne) AddUniqueInt(i int) *CommentUpdateOne {
	if cuo.addunique_int == nil {
		cuo.addunique_int = &i
	} else {
		*cuo.addunique_int += i
	}
	return cuo
}

// SetUniqueFloat sets the unique_float field.
func (cuo *CommentUpdateOne) SetUniqueFloat(f float64) *CommentUpdateOne {
	cuo.unique_float = &f
	cuo.addunique_float = nil
	return cuo
}

// AddUniqueFloat adds f to unique_float.
func (cuo *CommentUpdateOne) AddUniqueFloat(f float64) *CommentUpdateOne {
	if cuo.addunique_float == nil {
		cuo.addunique_float = &f
	} else {
		*cuo.addunique_float += f
	}
	return cuo
}

// SetNillableInt sets the nillable_int field.
func (cuo *CommentUpdateOne) SetNillableInt(i int) *CommentUpdateOne {
	cuo.nillable_int = &i
	cuo.addnillable_int = nil
	return cuo
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableNillableInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetNillableInt(*i)
	}
	return cuo
}

// AddNillableInt adds i to nillable_int.
func (cuo *CommentUpdateOne) AddNillableInt(i int) *CommentUpdateOne {
	if cuo.addnillable_int == nil {
		cuo.addnillable_int = &i
	} else {
		*cuo.addnillable_int += i
	}
	return cuo
}

// ClearNillableInt clears the value of nillable_int.
func (cuo *CommentUpdateOne) ClearNillableInt() *CommentUpdateOne {
	cuo.nillable_int = nil
	cuo.clearnillable_int = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (c *Comment, err error) {
	var (
		builder  = sql.Dialect(cuo.driver.Dialect())
		selector = builder.Select(comment.Columns...).From(builder.Table(comment.Table))
	)
	comment.ID(cuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		c = &Comment{config: cuo.config}
		if err := c.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Comment: %v", err)
		}
		id = c.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Comment with id: %v", cuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Comment with the same id: %v", cuo.id)
	}

	tx, err := cuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(comment.Table)
	)
	updater = updater.Where(sql.InInts(comment.FieldID, ids...))
	if value := cuo.unique_int; value != nil {
		updater.Set(comment.FieldUniqueInt, *value)
		c.UniqueInt = *value
	}
	if value := cuo.addunique_int; value != nil {
		updater.Add(comment.FieldUniqueInt, *value)
		c.UniqueInt += *value
	}
	if value := cuo.unique_float; value != nil {
		updater.Set(comment.FieldUniqueFloat, *value)
		c.UniqueFloat = *value
	}
	if value := cuo.addunique_float; value != nil {
		updater.Add(comment.FieldUniqueFloat, *value)
		c.UniqueFloat += *value
	}
	if value := cuo.nillable_int; value != nil {
		updater.Set(comment.FieldNillableInt, *value)
		c.NillableInt = value
	}
	if value := cuo.addnillable_int; value != nil {
		updater.Add(comment.FieldNillableInt, *value)
		if c.NillableInt != nil {
			*c.NillableInt += *value
		} else {
			c.NillableInt = value
		}
	}
	if cuo.clearnillable_int {
		c.NillableInt = nil
		updater.SetNull(comment.FieldNillableInt)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
