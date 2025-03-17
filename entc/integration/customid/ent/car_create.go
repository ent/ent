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
	"entgo.io/ent/entc/integration/customid/ent/car"
	"entgo.io/ent/entc/integration/customid/ent/pet"
	"entgo.io/ent/schema/field"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBeforeID sets the "before_id" field.
func (m *CarCreate) SetBeforeID(v float64) *CarCreate {
	m.mutation.SetBeforeID(v)
	return m
}

// SetNillableBeforeID sets the "before_id" field if the given value is not nil.
func (m *CarCreate) SetNillableBeforeID(v *float64) *CarCreate {
	if v != nil {
		m.SetBeforeID(*v)
	}
	return m
}

// SetAfterID sets the "after_id" field.
func (m *CarCreate) SetAfterID(v float64) *CarCreate {
	m.mutation.SetAfterID(v)
	return m
}

// SetNillableAfterID sets the "after_id" field if the given value is not nil.
func (m *CarCreate) SetNillableAfterID(v *float64) *CarCreate {
	if v != nil {
		m.SetAfterID(*v)
	}
	return m
}

// SetModel sets the "model" field.
func (m *CarCreate) SetModel(v string) *CarCreate {
	m.mutation.SetModel(v)
	return m
}

// SetID sets the "id" field.
func (m *CarCreate) SetID(v int) *CarCreate {
	m.mutation.SetID(v)
	return m
}

// SetOwnerID sets the "owner" edge to the Pet entity by ID.
func (m *CarCreate) SetOwnerID(id string) *CarCreate {
	m.mutation.SetOwnerID(id)
	return m
}

// SetNillableOwnerID sets the "owner" edge to the Pet entity by ID if the given value is not nil.
func (m *CarCreate) SetNillableOwnerID(id *string) *CarCreate {
	if id != nil {
		m = m.SetOwnerID(*id)
	}
	return m
}

// SetOwner sets the "owner" edge to the Pet entity.
func (m *CarCreate) SetOwner(v *Pet) *CarCreate {
	return m.SetOwnerID(v.ID)
}

// Mutation returns the CarMutation object of the builder.
func (m *CarCreate) Mutation() *CarMutation {
	return m.mutation
}

// Save creates the Car in the database.
func (c *CarCreate) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, c.sqlSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *CarCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *CarCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *CarCreate) check() error {
	if v, ok := c.mutation.BeforeID(); ok {
		if err := car.BeforeIDValidator(v); err != nil {
			return &ValidationError{Name: "before_id", err: fmt.Errorf(`ent: validator failed for field "Car.before_id": %w`, err)}
		}
	}
	if v, ok := c.mutation.AfterID(); ok {
		if err := car.AfterIDValidator(v); err != nil {
			return &ValidationError{Name: "after_id", err: fmt.Errorf(`ent: validator failed for field "Car.after_id": %w`, err)}
		}
	}
	if _, ok := c.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Car.model"`)}
	}
	if v, ok := c.mutation.ID(); ok {
		if err := car.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Car.id": %w`, err)}
		}
	}
	return nil
}

func (c *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	_node, _spec := c.createSpec()
	if err := sqlgraph.CreateNode(ctx, c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	c.mutation.id = &_node.ID
	c.mutation.done = true
	return _node, nil
}

func (c *CarCreate) createSpec() (*Car, *sqlgraph.CreateSpec) {
	var (
		_node = &Car{config: c.config}
		_spec = sqlgraph.NewCreateSpec(car.Table, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	)
	_spec.OnConflict = c.conflict
	if id, ok := c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := c.mutation.BeforeID(); ok {
		_spec.SetField(car.FieldBeforeID, field.TypeFloat64, value)
		_node.BeforeID = value
	}
	if value, ok := c.mutation.AfterID(); ok {
		_spec.SetField(car.FieldAfterID, field.TypeFloat64, value)
		_node.AfterID = value
	}
	if value, ok := c.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if nodes := c.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pet_cars = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Car.Create().
//		SetBeforeID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CarUpsert) {
//			SetBeforeID(v+v).
//		}).
//		Exec(ctx)
func (c *CarCreate) OnConflict(opts ...sql.ConflictOption) *CarUpsertOne {
	c.conflict = opts
	return &CarUpsertOne{create: c}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (c *CarCreate) OnConflictColumns(columns ...string) *CarUpsertOne {
	c.conflict = append(c.conflict, sql.ConflictColumns(columns...))
	return &CarUpsertOne{create: c}
}

type (
	// CarUpsertOne is the builder for "upsert"-ing
	//  one Car node.
	CarUpsertOne struct {
		create *CarCreate
	}

	// CarUpsert is the "OnConflict" setter.
	CarUpsert struct {
		*sql.UpdateSet
	}
)

// SetBeforeID sets the "before_id" field.
func (u *CarUpsert) SetBeforeID(v float64) *CarUpsert {
	u.Set(car.FieldBeforeID, v)
	return u
}

// UpdateBeforeID sets the "before_id" field to the value that was provided on create.
func (u *CarUpsert) UpdateBeforeID() *CarUpsert {
	u.SetExcluded(car.FieldBeforeID)
	return u
}

// AddBeforeID adds v to the "before_id" field.
func (u *CarUpsert) AddBeforeID(v float64) *CarUpsert {
	u.Add(car.FieldBeforeID, v)
	return u
}

// ClearBeforeID clears the value of the "before_id" field.
func (u *CarUpsert) ClearBeforeID() *CarUpsert {
	u.SetNull(car.FieldBeforeID)
	return u
}

// SetAfterID sets the "after_id" field.
func (u *CarUpsert) SetAfterID(v float64) *CarUpsert {
	u.Set(car.FieldAfterID, v)
	return u
}

// UpdateAfterID sets the "after_id" field to the value that was provided on create.
func (u *CarUpsert) UpdateAfterID() *CarUpsert {
	u.SetExcluded(car.FieldAfterID)
	return u
}

// AddAfterID adds v to the "after_id" field.
func (u *CarUpsert) AddAfterID(v float64) *CarUpsert {
	u.Add(car.FieldAfterID, v)
	return u
}

// ClearAfterID clears the value of the "after_id" field.
func (u *CarUpsert) ClearAfterID() *CarUpsert {
	u.SetNull(car.FieldAfterID)
	return u
}

// SetModel sets the "model" field.
func (u *CarUpsert) SetModel(v string) *CarUpsert {
	u.Set(car.FieldModel, v)
	return u
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *CarUpsert) UpdateModel() *CarUpsert {
	u.SetExcluded(car.FieldModel)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(car.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CarUpsertOne) UpdateNewValues() *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(car.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Car.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CarUpsertOne) Ignore() *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CarUpsertOne) DoNothing() *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CarCreate.OnConflict
// documentation for more info.
func (u *CarUpsertOne) Update(set func(*CarUpsert)) *CarUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CarUpsert{UpdateSet: update})
	}))
	return u
}

// SetBeforeID sets the "before_id" field.
func (u *CarUpsertOne) SetBeforeID(v float64) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetBeforeID(v)
	})
}

// AddBeforeID adds v to the "before_id" field.
func (u *CarUpsertOne) AddBeforeID(v float64) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.AddBeforeID(v)
	})
}

// UpdateBeforeID sets the "before_id" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateBeforeID() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateBeforeID()
	})
}

// ClearBeforeID clears the value of the "before_id" field.
func (u *CarUpsertOne) ClearBeforeID() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.ClearBeforeID()
	})
}

// SetAfterID sets the "after_id" field.
func (u *CarUpsertOne) SetAfterID(v float64) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetAfterID(v)
	})
}

// AddAfterID adds v to the "after_id" field.
func (u *CarUpsertOne) AddAfterID(v float64) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.AddAfterID(v)
	})
}

// UpdateAfterID sets the "after_id" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateAfterID() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateAfterID()
	})
}

// ClearAfterID clears the value of the "after_id" field.
func (u *CarUpsertOne) ClearAfterID() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.ClearAfterID()
	})
}

// SetModel sets the "model" field.
func (u *CarUpsertOne) SetModel(v string) *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *CarUpsertOne) UpdateModel() *CarUpsertOne {
	return u.Update(func(s *CarUpsert) {
		s.UpdateModel()
	})
}

// Exec executes the query.
func (u *CarUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CarCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CarUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CarUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CarUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CarCreateBulk is the builder for creating many Car entities in bulk.
type CarCreateBulk struct {
	config
	err      error
	builders []*CarCreate
	conflict []sql.ConflictOption
}

// Save creates the Car entities in the database.
func (c *CarCreateBulk) Save(ctx context.Context) ([]*Car, error) {
	if c.err != nil {
		return nil, c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(c.builders))
	nodes := make([]*Car, len(c.builders))
	mutators := make([]Mutator, len(c.builders))
	for i := range c.builders {
		func(i int, root context.Context) {
			builder := c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarMutation)
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
					_, err = mutators[i+1].Mutate(root, c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (c *CarCreateBulk) SaveX(ctx context.Context) []*Car {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *CarCreateBulk) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *CarCreateBulk) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Car.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CarUpsert) {
//			SetBeforeID(v+v).
//		}).
//		Exec(ctx)
func (c *CarCreateBulk) OnConflict(opts ...sql.ConflictOption) *CarUpsertBulk {
	c.conflict = opts
	return &CarUpsertBulk{create: c}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (c *CarCreateBulk) OnConflictColumns(columns ...string) *CarUpsertBulk {
	c.conflict = append(c.conflict, sql.ConflictColumns(columns...))
	return &CarUpsertBulk{create: c}
}

// CarUpsertBulk is the builder for "upsert"-ing
// a bulk of Car nodes.
type CarUpsertBulk struct {
	create *CarCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(car.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CarUpsertBulk) UpdateNewValues() *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(car.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Car.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CarUpsertBulk) Ignore() *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CarUpsertBulk) DoNothing() *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CarCreateBulk.OnConflict
// documentation for more info.
func (u *CarUpsertBulk) Update(set func(*CarUpsert)) *CarUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CarUpsert{UpdateSet: update})
	}))
	return u
}

// SetBeforeID sets the "before_id" field.
func (u *CarUpsertBulk) SetBeforeID(v float64) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetBeforeID(v)
	})
}

// AddBeforeID adds v to the "before_id" field.
func (u *CarUpsertBulk) AddBeforeID(v float64) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.AddBeforeID(v)
	})
}

// UpdateBeforeID sets the "before_id" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateBeforeID() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateBeforeID()
	})
}

// ClearBeforeID clears the value of the "before_id" field.
func (u *CarUpsertBulk) ClearBeforeID() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.ClearBeforeID()
	})
}

// SetAfterID sets the "after_id" field.
func (u *CarUpsertBulk) SetAfterID(v float64) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetAfterID(v)
	})
}

// AddAfterID adds v to the "after_id" field.
func (u *CarUpsertBulk) AddAfterID(v float64) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.AddAfterID(v)
	})
}

// UpdateAfterID sets the "after_id" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateAfterID() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateAfterID()
	})
}

// ClearAfterID clears the value of the "after_id" field.
func (u *CarUpsertBulk) ClearAfterID() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.ClearAfterID()
	})
}

// SetModel sets the "model" field.
func (u *CarUpsertBulk) SetModel(v string) *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *CarUpsertBulk) UpdateModel() *CarUpsertBulk {
	return u.Update(func(s *CarUpsert) {
		s.UpdateModel()
	})
}

// Exec executes the query.
func (u *CarUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CarCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CarCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CarUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
