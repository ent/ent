// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	mutation *BlobMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUUID sets the "uuid" field.
func (_c *BlobCreate) SetUUID(v uuid.UUID) *BlobCreate {
	_c.mutation.SetUUID(v)
	return _c
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (_c *BlobCreate) SetNillableUUID(v *uuid.UUID) *BlobCreate {
	if v != nil {
		_c.SetUUID(*v)
	}
	return _c
}

// SetCount sets the "count" field.
func (_c *BlobCreate) SetCount(v int) *BlobCreate {
	_c.mutation.SetCount(v)
	return _c
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (_c *BlobCreate) SetNillableCount(v *int) *BlobCreate {
	if v != nil {
		_c.SetCount(*v)
	}
	return _c
}

// SetID sets the "id" field.
func (_c *BlobCreate) SetID(v uuid.UUID) *BlobCreate {
	_c.mutation.SetID(v)
	return _c
}

// SetNillableID sets the "id" field if the given value is not nil.
func (_c *BlobCreate) SetNillableID(v *uuid.UUID) *BlobCreate {
	if v != nil {
		_c.SetID(*v)
	}
	return _c
}

// SetParentID sets the "parent" edge to the Blob entity by ID.
func (_c *BlobCreate) SetParentID(id uuid.UUID) *BlobCreate {
	_c.mutation.SetParentID(id)
	return _c
}

// SetNillableParentID sets the "parent" edge to the Blob entity by ID if the given value is not nil.
func (_c *BlobCreate) SetNillableParentID(id *uuid.UUID) *BlobCreate {
	if id != nil {
		_c = _c.SetParentID(*id)
	}
	return _c
}

// SetParent sets the "parent" edge to the Blob entity.
func (_c *BlobCreate) SetParent(v *Blob) *BlobCreate {
	return _c.SetParentID(v.ID)
}

// AddLinkIDs adds the "links" edge to the Blob entity by IDs.
func (_c *BlobCreate) AddLinkIDs(ids ...uuid.UUID) *BlobCreate {
	_c.mutation.AddLinkIDs(ids...)
	return _c
}

// AddLinks adds the "links" edges to the Blob entity.
func (_c *BlobCreate) AddLinks(v ...*Blob) *BlobCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddLinkIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (_c *BlobCreate) Mutation() *BlobMutation {
	return _c.mutation
}

// Save creates the Blob in the database.
func (_c *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *BlobCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *BlobCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *BlobCreate) defaults() {
	if _, ok := _c.mutation.UUID(); !ok {
		v := blob.DefaultUUID()
		_c.mutation.SetUUID(v)
	}
	if _, ok := _c.mutation.Count(); !ok {
		v := blob.DefaultCount
		_c.mutation.SetCount(v)
	}
	if _, ok := _c.mutation.ID(); !ok {
		v := blob.DefaultID()
		_c.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *BlobCreate) check() error {
	if _, ok := _c.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Blob.uuid"`)}
	}
	if _, ok := _c.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`ent: missing required field "Blob.count"`)}
	}
	return nil
}

func (_c *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *BlobCreate) createSpec() (*Blob, *sqlgraph.CreateSpec) {
	var (
		_node = &Blob{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(blob.Table, sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = _c.conflict
	if id, ok := _c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := _c.mutation.UUID(); ok {
		_spec.SetField(blob.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := _c.mutation.Count(); ok {
		_spec.SetField(blob.FieldCount, field.TypeInt, value)
		_node.Count = value
	}
	if nodes := _c.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.blob_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BlobLinkCreate{config: _c.config, mutation: newBlobLinkMutation(_c.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Blob.Create().
//		SetUUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlobUpsert) {
//			SetUUID(v+v).
//		}).
//		Exec(ctx)
func (_c *BlobCreate) OnConflict(opts ...sql.ConflictOption) *BlobUpsertOne {
	_c.conflict = opts
	return &BlobUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *BlobCreate) OnConflictColumns(columns ...string) *BlobUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &BlobUpsertOne{
		create: _c,
	}
}

type (
	// BlobUpsertOne is the builder for "upsert"-ing
	//  one Blob node.
	BlobUpsertOne struct {
		create *BlobCreate
	}

	// BlobUpsert is the "OnConflict" setter.
	BlobUpsert struct {
		*sql.UpdateSet
	}
)

// SetUUID sets the "uuid" field.
func (u *BlobUpsert) SetUUID(v uuid.UUID) *BlobUpsert {
	u.Set(blob.FieldUUID, v)
	return u
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *BlobUpsert) UpdateUUID() *BlobUpsert {
	u.SetExcluded(blob.FieldUUID)
	return u
}

// SetCount sets the "count" field.
func (u *BlobUpsert) SetCount(v int) *BlobUpsert {
	u.Set(blob.FieldCount, v)
	return u
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *BlobUpsert) UpdateCount() *BlobUpsert {
	u.SetExcluded(blob.FieldCount)
	return u
}

// AddCount adds v to the "count" field.
func (u *BlobUpsert) AddCount(v int) *BlobUpsert {
	u.Add(blob.FieldCount, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(blob.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BlobUpsertOne) UpdateNewValues() *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(blob.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Blob.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BlobUpsertOne) Ignore() *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlobUpsertOne) DoNothing() *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlobCreate.OnConflict
// documentation for more info.
func (u *BlobUpsertOne) Update(set func(*BlobUpsert)) *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlobUpsert{UpdateSet: update})
	}))
	return u
}

// SetUUID sets the "uuid" field.
func (u *BlobUpsertOne) SetUUID(v uuid.UUID) *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *BlobUpsertOne) UpdateUUID() *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateUUID()
	})
}

// SetCount sets the "count" field.
func (u *BlobUpsertOne) SetCount(v int) *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *BlobUpsertOne) AddCount(v int) *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *BlobUpsertOne) UpdateCount() *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateCount()
	})
}

// Exec executes the query.
func (u *BlobUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlobCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlobUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BlobUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BlobUpsertOne.ID is not supported by MySQL driver. Use BlobUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BlobUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BlobCreateBulk is the builder for creating many Blob entities in bulk.
type BlobCreateBulk struct {
	config
	err      error
	builders []*BlobCreate
	conflict []sql.ConflictOption
}

// Save creates the Blob entities in the database.
func (_c *BlobCreateBulk) Save(ctx context.Context) ([]*Blob, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Blob, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlobMutation)
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
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = _c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *BlobCreateBulk) SaveX(ctx context.Context) []*Blob {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *BlobCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *BlobCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Blob.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlobUpsert) {
//			SetUUID(v+v).
//		}).
//		Exec(ctx)
func (_c *BlobCreateBulk) OnConflict(opts ...sql.ConflictOption) *BlobUpsertBulk {
	_c.conflict = opts
	return &BlobUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *BlobCreateBulk) OnConflictColumns(columns ...string) *BlobUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &BlobUpsertBulk{
		create: _c,
	}
}

// BlobUpsertBulk is the builder for "upsert"-ing
// a bulk of Blob nodes.
type BlobUpsertBulk struct {
	create *BlobCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(blob.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BlobUpsertBulk) UpdateNewValues() *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(blob.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BlobUpsertBulk) Ignore() *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlobUpsertBulk) DoNothing() *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlobCreateBulk.OnConflict
// documentation for more info.
func (u *BlobUpsertBulk) Update(set func(*BlobUpsert)) *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlobUpsert{UpdateSet: update})
	}))
	return u
}

// SetUUID sets the "uuid" field.
func (u *BlobUpsertBulk) SetUUID(v uuid.UUID) *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *BlobUpsertBulk) UpdateUUID() *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateUUID()
	})
}

// SetCount sets the "count" field.
func (u *BlobUpsertBulk) SetCount(v int) *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *BlobUpsertBulk) AddCount(v int) *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *BlobUpsertBulk) UpdateCount() *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateCount()
	})
}

// Exec executes the query.
func (u *BlobUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if b == nil {
			return fmt.Errorf("ent: missing builder at index %d, unexpected nil builder passed to CreateBulk", i)
		}
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BlobCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlobCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlobUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
