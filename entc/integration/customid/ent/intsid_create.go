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
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/sid"
	"entgo.io/ent/schema/field"
)

// IntSIDCreate is the builder for creating a IntSID entity.
type IntSIDCreate struct {
	config
	mutation *IntSIDMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (isc *IntSIDCreate) SetID(s sid.ID) *IntSIDCreate {
	isc.mutation.SetID(s)
	return isc
}

// SetParentID sets the "parent" edge to the IntSID entity by ID.
func (isc *IntSIDCreate) SetParentID(id sid.ID) *IntSIDCreate {
	isc.mutation.SetParentID(id)
	return isc
}

// SetNillableParentID sets the "parent" edge to the IntSID entity by ID if the given value is not nil.
func (isc *IntSIDCreate) SetNillableParentID(id *sid.ID) *IntSIDCreate {
	if id != nil {
		isc = isc.SetParentID(*id)
	}
	return isc
}

// SetParent sets the "parent" edge to the IntSID entity.
func (isc *IntSIDCreate) SetParent(i *IntSID) *IntSIDCreate {
	return isc.SetParentID(i.ID)
}

// AddChildIDs adds the "children" edge to the IntSID entity by IDs.
func (isc *IntSIDCreate) AddChildIDs(ids ...sid.ID) *IntSIDCreate {
	isc.mutation.AddChildIDs(ids...)
	return isc
}

// AddChildren adds the "children" edges to the IntSID entity.
func (isc *IntSIDCreate) AddChildren(i ...*IntSID) *IntSIDCreate {
	ids := make([]sid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return isc.AddChildIDs(ids...)
}

// Mutation returns the IntSIDMutation object of the builder.
func (isc *IntSIDCreate) Mutation() *IntSIDMutation {
	return isc.mutation
}

// Save creates the IntSID in the database.
func (isc *IntSIDCreate) Save(ctx context.Context) (*IntSID, error) {
	return withHooks(ctx, isc.sqlSave, isc.mutation, isc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (isc *IntSIDCreate) SaveX(ctx context.Context) *IntSID {
	v, err := isc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (isc *IntSIDCreate) Exec(ctx context.Context) error {
	_, err := isc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isc *IntSIDCreate) ExecX(ctx context.Context) {
	if err := isc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (isc *IntSIDCreate) check() error {
	return nil
}

func (isc *IntSIDCreate) sqlSave(ctx context.Context) (*IntSID, error) {
	if err := isc.check(); err != nil {
		return nil, err
	}
	_node, _spec := isc.createSpec()
	if err := sqlgraph.CreateNode(ctx, isc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*sid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	isc.mutation.id = &_node.ID
	isc.mutation.done = true
	return _node, nil
}

func (isc *IntSIDCreate) createSpec() (*IntSID, *sqlgraph.CreateSpec) {
	var (
		_node = &IntSID{config: isc.config}
		_spec = sqlgraph.NewCreateSpec(intsid.Table, sqlgraph.NewFieldSpec(intsid.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = isc.conflict
	if id, ok := isc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if nodes := isc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   intsid.ParentTable,
			Columns: []string{intsid.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(intsid.FieldID, field.TypeInt64),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.int_sid_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := isc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(intsid.FieldID, field.TypeInt64),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.IntSID.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (isc *IntSIDCreate) OnConflict(opts ...sql.ConflictOption) *IntSIDUpsertOne {
	isc.conflict = opts
	return &IntSIDUpsertOne{
		create: isc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.IntSID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (isc *IntSIDCreate) OnConflictColumns(columns ...string) *IntSIDUpsertOne {
	isc.conflict = append(isc.conflict, sql.ConflictColumns(columns...))
	return &IntSIDUpsertOne{
		create: isc,
	}
}

type (
	// IntSIDUpsertOne is the builder for "upsert"-ing
	//  one IntSID node.
	IntSIDUpsertOne struct {
		create *IntSIDCreate
	}

	// IntSIDUpsert is the "OnConflict" setter.
	IntSIDUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.IntSID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(intsid.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *IntSIDUpsertOne) UpdateNewValues() *IntSIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(intsid.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.IntSID.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *IntSIDUpsertOne) Ignore() *IntSIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IntSIDUpsertOne) DoNothing() *IntSIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IntSIDCreate.OnConflict
// documentation for more info.
func (u *IntSIDUpsertOne) Update(set func(*IntSIDUpsert)) *IntSIDUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IntSIDUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *IntSIDUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IntSIDCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IntSIDUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *IntSIDUpsertOne) ID(ctx context.Context) (id sid.ID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *IntSIDUpsertOne) IDX(ctx context.Context) sid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// IntSIDCreateBulk is the builder for creating many IntSID entities in bulk.
type IntSIDCreateBulk struct {
	config
	err      error
	builders []*IntSIDCreate
	conflict []sql.ConflictOption
}

// Save creates the IntSID entities in the database.
func (iscb *IntSIDCreateBulk) Save(ctx context.Context) ([]*IntSID, error) {
	if iscb.err != nil {
		return nil, iscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(iscb.builders))
	nodes := make([]*IntSID, len(iscb.builders))
	mutators := make([]Mutator, len(iscb.builders))
	for i := range iscb.builders {
		func(i int, root context.Context) {
			builder := iscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IntSIDMutation)
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
					_, err = mutators[i+1].Mutate(root, iscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = iscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iscb.driver, spec); err != nil {
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
					if err := nodes[i].ID.Scan(specs[i].ID.Value); err != nil {
						return nil, err
					}
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
		if _, err := mutators[0].Mutate(ctx, iscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iscb *IntSIDCreateBulk) SaveX(ctx context.Context) []*IntSID {
	v, err := iscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iscb *IntSIDCreateBulk) Exec(ctx context.Context) error {
	_, err := iscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iscb *IntSIDCreateBulk) ExecX(ctx context.Context) {
	if err := iscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.IntSID.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (iscb *IntSIDCreateBulk) OnConflict(opts ...sql.ConflictOption) *IntSIDUpsertBulk {
	iscb.conflict = opts
	return &IntSIDUpsertBulk{
		create: iscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.IntSID.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (iscb *IntSIDCreateBulk) OnConflictColumns(columns ...string) *IntSIDUpsertBulk {
	iscb.conflict = append(iscb.conflict, sql.ConflictColumns(columns...))
	return &IntSIDUpsertBulk{
		create: iscb,
	}
}

// IntSIDUpsertBulk is the builder for "upsert"-ing
// a bulk of IntSID nodes.
type IntSIDUpsertBulk struct {
	create *IntSIDCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.IntSID.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(intsid.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *IntSIDUpsertBulk) UpdateNewValues() *IntSIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(intsid.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.IntSID.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *IntSIDUpsertBulk) Ignore() *IntSIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IntSIDUpsertBulk) DoNothing() *IntSIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IntSIDCreateBulk.OnConflict
// documentation for more info.
func (u *IntSIDUpsertBulk) Update(set func(*IntSIDUpsert)) *IntSIDUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IntSIDUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *IntSIDUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the IntSIDCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for IntSIDCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IntSIDUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
