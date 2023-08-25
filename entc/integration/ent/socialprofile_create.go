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
	"entgo.io/ent/entc/integration/ent/socialprofile"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
)

// SocialProfileCreate is the builder for creating a SocialProfile entity.
type SocialProfileCreate struct {
	config
	mutation *SocialProfileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDesc sets the "desc" field.
func (spc *SocialProfileCreate) SetDesc(s string) *SocialProfileCreate {
	spc.mutation.SetDesc(s)
	return spc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (spc *SocialProfileCreate) SetUserID(id int) *SocialProfileCreate {
	spc.mutation.SetUserID(id)
	return spc
}

// SetUser sets the "user" edge to the User entity.
func (spc *SocialProfileCreate) SetUser(u *User) *SocialProfileCreate {
	return spc.SetUserID(u.ID)
}

// Mutation returns the SocialProfileMutation object of the builder.
func (spc *SocialProfileCreate) Mutation() *SocialProfileMutation {
	return spc.mutation
}

// Save creates the SocialProfile in the database.
func (spc *SocialProfileCreate) Save(ctx context.Context) (*SocialProfile, error) {
	return withHooks(ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SocialProfileCreate) SaveX(ctx context.Context) *SocialProfile {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SocialProfileCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SocialProfileCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SocialProfileCreate) check() error {
	if _, ok := spc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "SocialProfile.desc"`)}
	}
	if _, ok := spc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "SocialProfile.user"`)}
	}
	return nil
}

func (spc *SocialProfileCreate) sqlSave(ctx context.Context) (*SocialProfile, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *SocialProfileCreate) createSpec() (*SocialProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &SocialProfile{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(socialprofile.Table, sqlgraph.NewFieldSpec(socialprofile.FieldID, field.TypeInt))
	)
	_spec.OnConflict = spc.conflict
	if value, ok := spc.mutation.Desc(); ok {
		_spec.SetField(socialprofile.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if nodes := spc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   socialprofile.UserTable,
			Columns: []string{socialprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_social_profiles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SocialProfile.Create().
//		SetDesc(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SocialProfileUpsert) {
//			SetDesc(v+v).
//		}).
//		Exec(ctx)
func (spc *SocialProfileCreate) OnConflict(opts ...sql.ConflictOption) *SocialProfileUpsertOne {
	spc.conflict = opts
	return &SocialProfileUpsertOne{
		create: spc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SocialProfile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (spc *SocialProfileCreate) OnConflictColumns(columns ...string) *SocialProfileUpsertOne {
	spc.conflict = append(spc.conflict, sql.ConflictColumns(columns...))
	return &SocialProfileUpsertOne{
		create: spc,
	}
}

type (
	// SocialProfileUpsertOne is the builder for "upsert"-ing
	//  one SocialProfile node.
	SocialProfileUpsertOne struct {
		create *SocialProfileCreate
	}

	// SocialProfileUpsert is the "OnConflict" setter.
	SocialProfileUpsert struct {
		*sql.UpdateSet
	}
)

// SetDesc sets the "desc" field.
func (u *SocialProfileUpsert) SetDesc(v string) *SocialProfileUpsert {
	u.Set(socialprofile.FieldDesc, v)
	return u
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *SocialProfileUpsert) UpdateDesc() *SocialProfileUpsert {
	u.SetExcluded(socialprofile.FieldDesc)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.SocialProfile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SocialProfileUpsertOne) UpdateNewValues() *SocialProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SocialProfile.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SocialProfileUpsertOne) Ignore() *SocialProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SocialProfileUpsertOne) DoNothing() *SocialProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SocialProfileCreate.OnConflict
// documentation for more info.
func (u *SocialProfileUpsertOne) Update(set func(*SocialProfileUpsert)) *SocialProfileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SocialProfileUpsert{UpdateSet: update})
	}))
	return u
}

// SetDesc sets the "desc" field.
func (u *SocialProfileUpsertOne) SetDesc(v string) *SocialProfileUpsertOne {
	return u.Update(func(s *SocialProfileUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *SocialProfileUpsertOne) UpdateDesc() *SocialProfileUpsertOne {
	return u.Update(func(s *SocialProfileUpsert) {
		s.UpdateDesc()
	})
}

// Exec executes the query.
func (u *SocialProfileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SocialProfileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SocialProfileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SocialProfileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SocialProfileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SocialProfileCreateBulk is the builder for creating many SocialProfile entities in bulk.
type SocialProfileCreateBulk struct {
	config
	err      error
	builders []*SocialProfileCreate
	conflict []sql.ConflictOption
}

// Save creates the SocialProfile entities in the database.
func (spcb *SocialProfileCreateBulk) Save(ctx context.Context) ([]*SocialProfile, error) {
	if spcb.err != nil {
		return nil, spcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*SocialProfile, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SocialProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = spcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *SocialProfileCreateBulk) SaveX(ctx context.Context) []*SocialProfile {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *SocialProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *SocialProfileCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SocialProfile.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SocialProfileUpsert) {
//			SetDesc(v+v).
//		}).
//		Exec(ctx)
func (spcb *SocialProfileCreateBulk) OnConflict(opts ...sql.ConflictOption) *SocialProfileUpsertBulk {
	spcb.conflict = opts
	return &SocialProfileUpsertBulk{
		create: spcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SocialProfile.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (spcb *SocialProfileCreateBulk) OnConflictColumns(columns ...string) *SocialProfileUpsertBulk {
	spcb.conflict = append(spcb.conflict, sql.ConflictColumns(columns...))
	return &SocialProfileUpsertBulk{
		create: spcb,
	}
}

// SocialProfileUpsertBulk is the builder for "upsert"-ing
// a bulk of SocialProfile nodes.
type SocialProfileUpsertBulk struct {
	create *SocialProfileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SocialProfile.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SocialProfileUpsertBulk) UpdateNewValues() *SocialProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SocialProfile.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SocialProfileUpsertBulk) Ignore() *SocialProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SocialProfileUpsertBulk) DoNothing() *SocialProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SocialProfileCreateBulk.OnConflict
// documentation for more info.
func (u *SocialProfileUpsertBulk) Update(set func(*SocialProfileUpsert)) *SocialProfileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SocialProfileUpsert{UpdateSet: update})
	}))
	return u
}

// SetDesc sets the "desc" field.
func (u *SocialProfileUpsertBulk) SetDesc(v string) *SocialProfileUpsertBulk {
	return u.Update(func(s *SocialProfileUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *SocialProfileUpsertBulk) UpdateDesc() *SocialProfileUpsertBulk {
	return u.Update(func(s *SocialProfileUpsert) {
		s.UpdateDesc()
	})
}

// Exec executes the query.
func (u *SocialProfileUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SocialProfileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SocialProfileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SocialProfileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
