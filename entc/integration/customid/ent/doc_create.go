// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/schema/field"
)

// DocCreate is the builder for creating a Doc entity.
type DocCreate struct {
	config
	mutation *DocMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetText sets the "text" field.
func (dc *DocCreate) SetText(s string) *DocCreate {
	dc.mutation.SetText(s)
	return dc
}

// SetNillableText sets the "text" field if the given value is not nil.
func (dc *DocCreate) SetNillableText(s *string) *DocCreate {
	if s != nil {
		dc.SetText(*s)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DocCreate) SetID(si schema.DocID) *DocCreate {
	dc.mutation.SetID(si)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DocCreate) SetNillableID(si *schema.DocID) *DocCreate {
	if si != nil {
		dc.SetID(*si)
	}
	return dc
}

// SetParentID sets the "parent" edge to the Doc entity by ID.
func (dc *DocCreate) SetParentID(id schema.DocID) *DocCreate {
	dc.mutation.SetParentID(id)
	return dc
}

// SetNillableParentID sets the "parent" edge to the Doc entity by ID if the given value is not nil.
func (dc *DocCreate) SetNillableParentID(id *schema.DocID) *DocCreate {
	if id != nil {
		dc = dc.SetParentID(*id)
	}
	return dc
}

// SetParent sets the "parent" edge to the Doc entity.
func (dc *DocCreate) SetParent(d *Doc) *DocCreate {
	return dc.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Doc entity by IDs.
func (dc *DocCreate) AddChildIDs(ids ...schema.DocID) *DocCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Doc entity.
func (dc *DocCreate) AddChildren(d ...*Doc) *DocCreate {
	ids := make([]schema.DocID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// Mutation returns the DocMutation object of the builder.
func (dc *DocCreate) Mutation() *DocMutation {
	return dc.mutation
}

// Save creates the Doc in the database.
func (dc *DocCreate) Save(ctx context.Context) (*Doc, error) {
	var (
		err  error
		node *Doc
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DocCreate) SaveX(ctx context.Context) *Doc {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DocCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DocCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DocCreate) defaults() {
	if _, ok := dc.mutation.ID(); !ok {
		v := doc.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DocCreate) check() error {
	if v, ok := dc.mutation.ID(); ok {
		if err := doc.IDValidator(string(v)); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Doc.id": %w`, err)}
		}
	}
	return nil
}

func (dc *DocCreate) sqlSave(ctx context.Context) (*Doc, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*schema.DocID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (dc *DocCreate) createSpec() (*Doc, *sqlgraph.CreateSpec) {
	var (
		_node = &Doc{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: doc.FieldID,
			},
		}
	)
	_spec.OnConflict = dc.conflict
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doc.FieldText,
		})
		_node.Text = value
	}
	if nodes := dc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   doc.ParentTable,
			Columns: []string{doc.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.doc_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doc.ChildrenTable,
			Columns: []string{doc.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: doc.FieldID,
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Doc.Create().
//		SetText(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DocUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
//
func (dc *DocCreate) OnConflict(opts ...sql.ConflictOption) *DocUpsertOne {
	dc.conflict = opts
	return &DocUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Doc.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dc *DocCreate) OnConflictColumns(columns ...string) *DocUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DocUpsertOne{
		create: dc,
	}
}

type (
	// DocUpsertOne is the builder for "upsert"-ing
	//  one Doc node.
	DocUpsertOne struct {
		create *DocCreate
	}

	// DocUpsert is the "OnConflict" setter.
	DocUpsert struct {
		*sql.UpdateSet
	}
)

// SetText sets the "text" field.
func (u *DocUpsert) SetText(v string) *DocUpsert {
	u.Set(doc.FieldText, v)
	return u
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *DocUpsert) UpdateText() *DocUpsert {
	u.SetExcluded(doc.FieldText)
	return u
}

// ClearText clears the value of the "text" field.
func (u *DocUpsert) ClearText() *DocUpsert {
	u.SetNull(doc.FieldText)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Doc.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(doc.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DocUpsertOne) UpdateNewValues() *DocUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(doc.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Doc.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DocUpsertOne) Ignore() *DocUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DocUpsertOne) DoNothing() *DocUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DocCreate.OnConflict
// documentation for more info.
func (u *DocUpsertOne) Update(set func(*DocUpsert)) *DocUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DocUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *DocUpsertOne) SetText(v string) *DocUpsertOne {
	return u.Update(func(s *DocUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *DocUpsertOne) UpdateText() *DocUpsertOne {
	return u.Update(func(s *DocUpsert) {
		s.UpdateText()
	})
}

// ClearText clears the value of the "text" field.
func (u *DocUpsertOne) ClearText() *DocUpsertOne {
	return u.Update(func(s *DocUpsert) {
		s.ClearText()
	})
}

// Exec executes the query.
func (u *DocUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DocCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DocUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DocUpsertOne) ID(ctx context.Context) (id schema.DocID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DocUpsertOne.ID is not supported by MySQL driver. Use DocUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DocUpsertOne) IDX(ctx context.Context) schema.DocID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DocCreateBulk is the builder for creating many Doc entities in bulk.
type DocCreateBulk struct {
	config
	builders []*DocCreate
	conflict []sql.ConflictOption
}

// Save creates the Doc entities in the database.
func (dcb *DocCreateBulk) Save(ctx context.Context) ([]*Doc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Doc, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DocCreateBulk) SaveX(ctx context.Context) []*Doc {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DocCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DocCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Doc.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DocUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
//
func (dcb *DocCreateBulk) OnConflict(opts ...sql.ConflictOption) *DocUpsertBulk {
	dcb.conflict = opts
	return &DocUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Doc.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (dcb *DocCreateBulk) OnConflictColumns(columns ...string) *DocUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DocUpsertBulk{
		create: dcb,
	}
}

// DocUpsertBulk is the builder for "upsert"-ing
// a bulk of Doc nodes.
type DocUpsertBulk struct {
	create *DocCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Doc.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(doc.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *DocUpsertBulk) UpdateNewValues() *DocUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(doc.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Doc.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DocUpsertBulk) Ignore() *DocUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DocUpsertBulk) DoNothing() *DocUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DocCreateBulk.OnConflict
// documentation for more info.
func (u *DocUpsertBulk) Update(set func(*DocUpsert)) *DocUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DocUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *DocUpsertBulk) SetText(v string) *DocUpsertBulk {
	return u.Update(func(s *DocUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *DocUpsertBulk) UpdateText() *DocUpsertBulk {
	return u.Update(func(s *DocUpsert) {
		s.UpdateText()
	})
}

// ClearText clears the value of the "text" field.
func (u *DocUpsertBulk) ClearText() *DocUpsertBulk {
	return u.Update(func(s *DocUpsert) {
		s.ClearText()
	})
}

// Exec executes the query.
func (u *DocUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DocCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DocCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DocUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
