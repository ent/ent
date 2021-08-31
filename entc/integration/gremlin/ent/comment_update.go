// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/comment"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUniqueInt sets the "unique_int" field.
func (cu *CommentUpdate) SetUniqueInt(i int) *CommentUpdate {
	cu.mutation.ResetUniqueInt()
	cu.mutation.SetUniqueInt(i)
	return cu
}

// AddUniqueInt adds i to the "unique_int" field.
func (cu *CommentUpdate) AddUniqueInt(i int) *CommentUpdate {
	cu.mutation.AddUniqueInt(i)
	return cu
}

// SetUniqueFloat sets the "unique_float" field.
func (cu *CommentUpdate) SetUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.ResetUniqueFloat()
	cu.mutation.SetUniqueFloat(f)
	return cu
}

// AddUniqueFloat adds f to the "unique_float" field.
func (cu *CommentUpdate) AddUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.AddUniqueFloat(f)
	return cu
}

// SetNillableInt sets the "nillable_int" field.
func (cu *CommentUpdate) SetNillableInt(i int) *CommentUpdate {
	cu.mutation.ResetNillableInt()
	cu.mutation.SetNillableInt(i)
	return cu
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableNillableInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetNillableInt(*i)
	}
	return cu
}

// AddNillableInt adds i to the "nillable_int" field.
func (cu *CommentUpdate) AddNillableInt(i int) *CommentUpdate {
	cu.mutation.AddNillableInt(i)
	return cu
}

// ClearNillableInt clears the value of the "nillable_int" field.
func (cu *CommentUpdate) ClearNillableInt() *CommentUpdate {
	cu.mutation.ClearNillableInt()
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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

func (cu *CommentUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cu.gremlin().Query()
	if err := cu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (cu *CommentUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(comment.Label)
	for _, p := range cu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cu.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := cu.mutation.AddedUniqueInt(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, __.Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum())
	}
	if value, ok := cu.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := cu.mutation.AddedUniqueFloat(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, __.Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum())
	}
	if value, ok := cu.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
	}
	if value, ok := cu.mutation.AddedNillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, __.Union(__.Values(comment.FieldNillableInt), __.Constant(value)).Sum())
	}
	var properties []interface{}
	if cu.mutation.NillableIntCleared() {
		properties = append(properties, comment.FieldNillableInt)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUniqueInt sets the "unique_int" field.
func (cuo *CommentUpdateOne) SetUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetUniqueInt()
	cuo.mutation.SetUniqueInt(i)
	return cuo
}

// AddUniqueInt adds i to the "unique_int" field.
func (cuo *CommentUpdateOne) AddUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.AddUniqueInt(i)
	return cuo
}

// SetUniqueFloat sets the "unique_float" field.
func (cuo *CommentUpdateOne) SetUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.ResetUniqueFloat()
	cuo.mutation.SetUniqueFloat(f)
	return cuo
}

// AddUniqueFloat adds f to the "unique_float" field.
func (cuo *CommentUpdateOne) AddUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.AddUniqueFloat(f)
	return cuo
}

// SetNillableInt sets the "nillable_int" field.
func (cuo *CommentUpdateOne) SetNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetNillableInt()
	cuo.mutation.SetNillableInt(i)
	return cuo
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableNillableInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetNillableInt(*i)
	}
	return cuo
}

// AddNillableInt adds i to the "nillable_int" field.
func (cuo *CommentUpdateOne) AddNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.AddNillableInt(i)
	return cuo
}

// ClearNillableInt clears the value of the "nillable_int" field.
func (cuo *CommentUpdateOne) ClearNillableInt() *CommentUpdateOne {
	cuo.mutation.ClearNillableInt()
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
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

func (cuo *CommentUpdateOne) gremlinSave(ctx context.Context) (*Comment, error) {
	res := &gremlin.Response{}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Comment.ID for update")}
	}
	query, bindings := cuo.gremlin(id).Query()
	if err := cuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Comment{config: cuo.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CommentUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cuo.mutation.UniqueInt(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, value)
	}
	if value, ok := cuo.mutation.AddedUniqueInt(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, __.Union(__.Values(comment.FieldUniqueInt), __.Constant(value)).Sum())
	}
	if value, ok := cuo.mutation.UniqueFloat(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, value)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, value)
	}
	if value, ok := cuo.mutation.AddedUniqueFloat(); ok {
		addValue := rv.Clone().Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum().Next()
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, addValue).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, fmt.Sprintf("+= %v", value))),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, __.Union(__.Values(comment.FieldUniqueFloat), __.Constant(value)).Sum())
	}
	if value, ok := cuo.mutation.NillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, value)
	}
	if value, ok := cuo.mutation.AddedNillableInt(); ok {
		v.Property(dsl.Single, comment.FieldNillableInt, __.Union(__.Values(comment.FieldNillableInt), __.Constant(value)).Sum())
	}
	var properties []interface{}
	if cuo.mutation.NillableIntCleared() {
		properties = append(properties, comment.FieldNillableInt)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if len(cuo.fields) > 0 {
		fields := make([]interface{}, 0, len(cuo.fields)+1)
		fields = append(fields, true)
		for _, f := range cuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
