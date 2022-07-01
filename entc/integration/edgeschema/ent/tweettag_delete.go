// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
)

// TweetTagDelete is the builder for deleting a TweetTag entity.
type TweetTagDelete struct {
	config
	hooks    []Hook
	mutation *TweetTagMutation
}

// Where appends a list predicates to the TweetTagDelete builder.
func (ttd *TweetTagDelete) Where(ps ...predicate.TweetTag) *TweetTagDelete {
	ttd.mutation.Where(ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TweetTagDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttd.hooks) == 0 {
		affected, err = ttd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttd.mutation = mutation
			affected, err = ttd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttd.hooks) - 1; i >= 0; i-- {
			if ttd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ttd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TweetTagDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TweetTagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tweettag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tweettag.FieldID,
			},
		},
	}
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TweetTagDeleteOne is the builder for deleting a single TweetTag entity.
type TweetTagDeleteOne struct {
	ttd *TweetTagDelete
}

// Exec executes the deletion query.
func (ttdo *TweetTagDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tweettag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TweetTagDeleteOne) ExecX(ctx context.Context) {
	ttdo.ttd.ExecX(ctx)
}
