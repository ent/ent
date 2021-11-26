// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/spec"
	"entgo.io/ent/entc/integration/ent/user"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation *CardMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (cc *CardCreate) SetCreateTime(t time.Time) *CardCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableCreateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CardCreate) SetUpdateTime(t time.Time) *CardCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetBalance sets the "balance" field.
func (cc *CardCreate) SetBalance(f float64) *CardCreate {
	cc.mutation.SetBalance(f)
	return cc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (cc *CardCreate) SetNillableBalance(f *float64) *CardCreate {
	if f != nil {
		cc.SetBalance(*f)
	}
	return cc
}

// SetNumber sets the "number" field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.mutation.SetNumber(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CardCreate) SetName(s string) *CardCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CardCreate) SetNillableName(s *string) *CardCreate {
	if s != nil {
		cc.SetName(*s)
	}
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

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (cc *CardCreate) AddSpecIDs(ids ...int) *CardCreate {
	cc.mutation.AddSpecIDs(ids...)
	return cc
}

// AddSpec adds the "spec" edges to the Spec entity.
func (cc *CardCreate) AddSpec(s ...*Spec) *CardCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSpecIDs(ids...)
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
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
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
			if node, err = cc.sqlSave(ctx); err != nil {
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
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
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

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := card.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := card.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Balance(); !ok {
		v := card.DefaultBalance
		cc.mutation.SetBalance(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Card.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Card.update_time"`)}
	}
	if _, ok := cc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New(`ent: missing required field "Card.balance"`)}
	}
	if _, ok := cc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Card.number"`)}
	}
	if v, ok := cc.mutation.Number(); ok {
		if err := card.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Card.number": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	return nil
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CardCreate) createSpec() (*Card, *sqlgraph.CreateSpec) {
	var (
		_node = &Card{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: card.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: card.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: card.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: card.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Balance(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: card.FieldBalance,
		})
		_node.Balance = value
	}
	if value, ok := cc.mutation.Number(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: card.FieldNumber,
		})
		_node.Number = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: card.FieldName,
		})
		_node.Name = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_card = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: spec.FieldID,
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
//	client.Card.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CardUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CardCreate) OnConflict(opts ...sql.ConflictOption) *CardUpsertOne {
	cc.conflict = opts
	return &CardUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CardCreate) OnConflictColumns(columns ...string) *CardUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CardUpsertOne{
		create: cc,
	}
}

type (
	// CardUpsertOne is the builder for "upsert"-ing
	//  one Card node.
	CardUpsertOne struct {
		create *CardCreate
	}

	// CardUpsert is the "OnConflict" setter.
	CardUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *CardUpsert) SetCreateTime(v time.Time) *CardUpsert {
	u.Set(card.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CardUpsert) UpdateCreateTime() *CardUpsert {
	u.SetExcluded(card.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *CardUpsert) SetUpdateTime(v time.Time) *CardUpsert {
	u.Set(card.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CardUpsert) UpdateUpdateTime() *CardUpsert {
	u.SetExcluded(card.FieldUpdateTime)
	return u
}

// SetBalance sets the "balance" field.
func (u *CardUpsert) SetBalance(v float64) *CardUpsert {
	u.Set(card.FieldBalance, v)
	return u
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CardUpsert) UpdateBalance() *CardUpsert {
	u.SetExcluded(card.FieldBalance)
	return u
}

// AddBalance adds v to the "balance" field.
func (u *CardUpsert) AddBalance(v float64) *CardUpsert {
	u.Add(card.FieldBalance, v)
	return u
}

// SetNumber sets the "number" field.
func (u *CardUpsert) SetNumber(v string) *CardUpsert {
	u.Set(card.FieldNumber, v)
	return u
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *CardUpsert) UpdateNumber() *CardUpsert {
	u.SetExcluded(card.FieldNumber)
	return u
}

// SetName sets the "name" field.
func (u *CardUpsert) SetName(v string) *CardUpsert {
	u.Set(card.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CardUpsert) UpdateName() *CardUpsert {
	u.SetExcluded(card.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *CardUpsert) ClearName() *CardUpsert {
	u.SetNull(card.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CardUpsertOne) UpdateNewValues() *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(card.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Number(); exists {
			s.SetIgnore(card.FieldNumber)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Card.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CardUpsertOne) Ignore() *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CardUpsertOne) DoNothing() *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CardCreate.OnConflict
// documentation for more info.
func (u *CardUpsertOne) Update(set func(*CardUpsert)) *CardUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CardUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CardUpsertOne) SetCreateTime(v time.Time) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateCreateTime() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CardUpsertOne) SetUpdateTime(v time.Time) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateUpdateTime() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetBalance sets the "balance" field.
func (u *CardUpsertOne) SetBalance(v float64) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetBalance(v)
	})
}

// AddBalance adds v to the "balance" field.
func (u *CardUpsertOne) AddBalance(v float64) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.AddBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateBalance() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateBalance()
	})
}

// SetNumber sets the "number" field.
func (u *CardUpsertOne) SetNumber(v string) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateNumber() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateNumber()
	})
}

// SetName sets the "name" field.
func (u *CardUpsertOne) SetName(v string) *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CardUpsertOne) UpdateName() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *CardUpsertOne) ClearName() *CardUpsertOne {
	return u.Update(func(s *CardUpsert) {
		s.ClearName()
	})
}

// Exec executes the query.
func (u *CardUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CardCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CardUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CardUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	builders []*CardCreate
	conflict []sql.ConflictOption
}

// Save creates the Card entities in the database.
func (ccb *CardCreateBulk) Save(ctx context.Context) ([]*Card, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Card, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CardCreateBulk) SaveX(ctx context.Context) []*Card {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CardCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CardCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Card.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CardUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CardCreateBulk) OnConflict(opts ...sql.ConflictOption) *CardUpsertBulk {
	ccb.conflict = opts
	return &CardUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CardCreateBulk) OnConflictColumns(columns ...string) *CardUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CardUpsertBulk{
		create: ccb,
	}
}

// CardUpsertBulk is the builder for "upsert"-ing
// a bulk of Card nodes.
type CardUpsertBulk struct {
	create *CardCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *CardUpsertBulk) UpdateNewValues() *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(card.FieldCreateTime)
			}
			if _, exists := b.mutation.Number(); exists {
				s.SetIgnore(card.FieldNumber)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Card.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CardUpsertBulk) Ignore() *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CardUpsertBulk) DoNothing() *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CardCreateBulk.OnConflict
// documentation for more info.
func (u *CardUpsertBulk) Update(set func(*CardUpsert)) *CardUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CardUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *CardUpsertBulk) SetCreateTime(v time.Time) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateCreateTime() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *CardUpsertBulk) SetUpdateTime(v time.Time) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateUpdateTime() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetBalance sets the "balance" field.
func (u *CardUpsertBulk) SetBalance(v float64) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetBalance(v)
	})
}

// AddBalance adds v to the "balance" field.
func (u *CardUpsertBulk) AddBalance(v float64) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.AddBalance(v)
	})
}

// UpdateBalance sets the "balance" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateBalance() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateBalance()
	})
}

// SetNumber sets the "number" field.
func (u *CardUpsertBulk) SetNumber(v string) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateNumber() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateNumber()
	})
}

// SetName sets the "name" field.
func (u *CardUpsertBulk) SetName(v string) *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CardUpsertBulk) UpdateName() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *CardUpsertBulk) ClearName() *CardUpsertBulk {
	return u.Update(func(s *CardUpsert) {
		s.ClearName()
	})
}

// Exec executes the query.
func (u *CardUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CardCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CardCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
