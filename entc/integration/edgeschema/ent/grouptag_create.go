// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/grouptag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/schema/field"
)

// GroupTagCreate is the builder for creating a GroupTag entity.
type GroupTagCreate struct {
	config
	mutation *GroupTagMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTagID sets the "tag_id" field.
func (gtc *GroupTagCreate) SetTagID(i int) *GroupTagCreate {
	gtc.mutation.SetTagID(i)
	return gtc
}

// SetGroupID sets the "group_id" field.
func (gtc *GroupTagCreate) SetGroupID(i int) *GroupTagCreate {
	gtc.mutation.SetGroupID(i)
	return gtc
}

// SetTag sets the "tag" edge to the Tag entity.
func (gtc *GroupTagCreate) SetTag(t *Tag) *GroupTagCreate {
	return gtc.SetTagID(t.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (gtc *GroupTagCreate) SetGroup(g *Group) *GroupTagCreate {
	return gtc.SetGroupID(g.ID)
}

// Mutation returns the GroupTagMutation object of the builder.
func (gtc *GroupTagCreate) Mutation() *GroupTagMutation {
	return gtc.mutation
}

// Save creates the GroupTag in the database.
func (gtc *GroupTagCreate) Save(ctx context.Context) (*GroupTag, error) {
	return withHooks(ctx, gtc.sqlSave, gtc.mutation, gtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gtc *GroupTagCreate) SaveX(ctx context.Context) *GroupTag {
	v, err := gtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gtc *GroupTagCreate) Exec(ctx context.Context) error {
	_, err := gtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtc *GroupTagCreate) ExecX(ctx context.Context) {
	if err := gtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtc *GroupTagCreate) check() error {
	if _, ok := gtc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag_id", err: errors.New(`ent: missing required field "GroupTag.tag_id"`)}
	}
	if _, ok := gtc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "GroupTag.group_id"`)}
	}
	if _, ok := gtc.mutation.TagID(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required edge "GroupTag.tag"`)}
	}
	if _, ok := gtc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupTag.group"`)}
	}
	return nil
}

func (gtc *GroupTagCreate) sqlSave(ctx context.Context) (*GroupTag, error) {
	if err := gtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gtc.mutation.id = &_node.ID
	gtc.mutation.done = true
	return _node, nil
}

func (gtc *GroupTagCreate) createSpec() (*GroupTag, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupTag{config: gtc.config}
		_spec = sqlgraph.NewCreateSpec(grouptag.Table, sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt))
	)
	_spec.OnConflict = gtc.conflict
	if nodes := gtc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.TagTable,
			Columns: []string{grouptag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gtc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.GroupTable,
			Columns: []string{grouptag.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupTag.Create().
//		SetTagID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupTagUpsert) {
//			SetTagID(v+v).
//		}).
//		Exec(ctx)
func (gtc *GroupTagCreate) OnConflict(opts ...sql.ConflictOption) *GroupTagUpsertOne {
	gtc.conflict = opts
	return &GroupTagUpsertOne{
		create: gtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gtc *GroupTagCreate) OnConflictColumns(columns ...string) *GroupTagUpsertOne {
	gtc.conflict = append(gtc.conflict, sql.ConflictColumns(columns...))
	return &GroupTagUpsertOne{
		create: gtc,
	}
}

type (
	// GroupTagUpsertOne is the builder for "upsert"-ing
	//  one GroupTag node.
	GroupTagUpsertOne struct {
		create *GroupTagCreate
	}

	// GroupTagUpsert is the "OnConflict" setter.
	GroupTagUpsert struct {
		*sql.UpdateSet
	}
)

// SetTagID sets the "tag_id" field.
func (u *GroupTagUpsert) SetTagID(v int) *GroupTagUpsert {
	u.Set(grouptag.FieldTagID, v)
	return u
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *GroupTagUpsert) UpdateTagID() *GroupTagUpsert {
	u.SetExcluded(grouptag.FieldTagID)
	return u
}

// SetGroupID sets the "group_id" field.
func (u *GroupTagUpsert) SetGroupID(v int) *GroupTagUpsert {
	u.Set(grouptag.FieldGroupID, v)
	return u
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *GroupTagUpsert) UpdateGroupID() *GroupTagUpsert {
	u.SetExcluded(grouptag.FieldGroupID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GroupTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupTagUpsertOne) UpdateNewValues() *GroupTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupTag.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupTagUpsertOne) Ignore() *GroupTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupTagUpsertOne) DoNothing() *GroupTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupTagCreate.OnConflict
// documentation for more info.
func (u *GroupTagUpsertOne) Update(set func(*GroupTagUpsert)) *GroupTagUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetTagID sets the "tag_id" field.
func (u *GroupTagUpsertOne) SetTagID(v int) *GroupTagUpsertOne {
	return u.Update(func(s *GroupTagUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *GroupTagUpsertOne) UpdateTagID() *GroupTagUpsertOne {
	return u.Update(func(s *GroupTagUpsert) {
		s.UpdateTagID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *GroupTagUpsertOne) SetGroupID(v int) *GroupTagUpsertOne {
	return u.Update(func(s *GroupTagUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *GroupTagUpsertOne) UpdateGroupID() *GroupTagUpsertOne {
	return u.Update(func(s *GroupTagUpsert) {
		s.UpdateGroupID()
	})
}

// Exec executes the query.
func (u *GroupTagUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupTagCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupTagUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupTagUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupTagUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupTagCreateBulk is the builder for creating many GroupTag entities in bulk.
type GroupTagCreateBulk struct {
	config
	err      error
	builders []*GroupTagCreate
	conflict []sql.ConflictOption
}

// Save creates the GroupTag entities in the database.
func (gtcb *GroupTagCreateBulk) Save(ctx context.Context) ([]*GroupTag, error) {
	if gtcb.err != nil {
		return nil, gtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gtcb.builders))
	nodes := make([]*GroupTag, len(gtcb.builders))
	mutators := make([]Mutator, len(gtcb.builders))
	for i := range gtcb.builders {
		func(i int, root context.Context) {
			builder := gtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupTagMutation)
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
					_, err = mutators[i+1].Mutate(root, gtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gtcb *GroupTagCreateBulk) SaveX(ctx context.Context) []*GroupTag {
	v, err := gtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gtcb *GroupTagCreateBulk) Exec(ctx context.Context) error {
	_, err := gtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtcb *GroupTagCreateBulk) ExecX(ctx context.Context) {
	if err := gtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupTag.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupTagUpsert) {
//			SetTagID(v+v).
//		}).
//		Exec(ctx)
func (gtcb *GroupTagCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupTagUpsertBulk {
	gtcb.conflict = opts
	return &GroupTagUpsertBulk{
		create: gtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupTag.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gtcb *GroupTagCreateBulk) OnConflictColumns(columns ...string) *GroupTagUpsertBulk {
	gtcb.conflict = append(gtcb.conflict, sql.ConflictColumns(columns...))
	return &GroupTagUpsertBulk{
		create: gtcb,
	}
}

// GroupTagUpsertBulk is the builder for "upsert"-ing
// a bulk of GroupTag nodes.
type GroupTagUpsertBulk struct {
	create *GroupTagCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GroupTag.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupTagUpsertBulk) UpdateNewValues() *GroupTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupTag.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupTagUpsertBulk) Ignore() *GroupTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupTagUpsertBulk) DoNothing() *GroupTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupTagCreateBulk.OnConflict
// documentation for more info.
func (u *GroupTagUpsertBulk) Update(set func(*GroupTagUpsert)) *GroupTagUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupTagUpsert{UpdateSet: update})
	}))
	return u
}

// SetTagID sets the "tag_id" field.
func (u *GroupTagUpsertBulk) SetTagID(v int) *GroupTagUpsertBulk {
	return u.Update(func(s *GroupTagUpsert) {
		s.SetTagID(v)
	})
}

// UpdateTagID sets the "tag_id" field to the value that was provided on create.
func (u *GroupTagUpsertBulk) UpdateTagID() *GroupTagUpsertBulk {
	return u.Update(func(s *GroupTagUpsert) {
		s.UpdateTagID()
	})
}

// SetGroupID sets the "group_id" field.
func (u *GroupTagUpsertBulk) SetGroupID(v int) *GroupTagUpsertBulk {
	return u.Update(func(s *GroupTagUpsert) {
		s.SetGroupID(v)
	})
}

// UpdateGroupID sets the "group_id" field to the value that was provided on create.
func (u *GroupTagUpsertBulk) UpdateGroupID() *GroupTagUpsertBulk {
	return u.Update(func(s *GroupTagUpsert) {
		s.UpdateGroupID()
	})
}

// Exec executes the query.
func (u *GroupTagUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupTagCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupTagCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupTagUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
