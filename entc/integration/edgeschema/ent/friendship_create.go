// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"
	errors "errors"
	fmt "fmt"
	time "time"

	"entgo.io/ent/dialect/sql"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	friendship "entgo.io/ent/entc/integration/edgeschema/ent/friendship"
	user "entgo.io/ent/entc/integration/edgeschema/ent/user"
	field "entgo.io/ent/schema/field"
)

// FriendshipCreate is the builder for creating a Friendship entity.
type FriendshipCreate struct {
	config
	mutation *FriendshipMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetWeight sets the "weight" field.
func (fc *FriendshipCreate) SetWeight(i int) *FriendshipCreate {
	fc.mutation.SetWeight(i)
	return fc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableWeight(i *int) *FriendshipCreate {
	if i != nil {
		fc.SetWeight(*i)
	}
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FriendshipCreate) SetCreatedAt(t time.Time) *FriendshipCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FriendshipCreate) SetNillableCreatedAt(t *time.Time) *FriendshipCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUserID sets the "user_id" field.
func (fc *FriendshipCreate) SetUserID(i int) *FriendshipCreate {
	fc.mutation.SetUserID(i)
	return fc
}

// SetFriendID sets the "friend_id" field.
func (fc *FriendshipCreate) SetFriendID(i int) *FriendshipCreate {
	fc.mutation.SetFriendID(i)
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FriendshipCreate) SetUser(u *User) *FriendshipCreate {
	return fc.SetUserID(u.ID)
}

// SetFriend sets the "friend" edge to the User entity.
func (fc *FriendshipCreate) SetFriend(u *User) *FriendshipCreate {
	return fc.SetFriendID(u.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fc *FriendshipCreate) Mutation() *FriendshipMutation {
	return fc.mutation
}

// Save creates the Friendship in the database.
func (fc *FriendshipCreate) Save(ctx context.Context) (*Friendship, error) {
	fc.defaults()
	return withHooks[*Friendship, FriendshipMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FriendshipCreate) SaveX(ctx context.Context) *Friendship {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FriendshipCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FriendshipCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FriendshipCreate) defaults() {
	if _, ok := fc.mutation.Weight(); !ok {
		v := friendship.DefaultWeight
		fc.mutation.SetWeight(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := friendship.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FriendshipCreate) check() error {
	if _, ok := fc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "Friendship.weight"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Friendship.created_at"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Friendship.user_id"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend_id", err: errors.New(`ent: missing required field "Friendship.friend_id"`)}
	}
	if _, ok := fc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Friendship.user"`)}
	}
	if _, ok := fc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend", err: errors.New(`ent: missing required edge "Friendship.friend"`)}
	}
	return nil
}

func (fc *FriendshipCreate) sqlSave(ctx context.Context) (*Friendship, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FriendshipCreate) createSpec() (*Friendship, *sqlgraph.CreateSpec) {
	var (
		_node = &Friendship{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: friendship.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: friendship.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.Weight(); ok {
		_spec.SetField(friendship.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
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
		_node.FriendID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Friendship.Create().
//		SetWeight(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FriendshipUpsert) {
//			SetWeight(v+v).
//		}).
//		Exec(ctx)
func (fc *FriendshipCreate) OnConflict(opts ...sql.ConflictOption) *FriendshipUpsertOne {
	fc.conflict = opts
	return &FriendshipUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FriendshipCreate) OnConflictColumns(columns ...string) *FriendshipUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FriendshipUpsertOne{
		create: fc,
	}
}

type (
	// FriendshipUpsertOne is the builder for "upsert"-ing
	//  one Friendship node.
	FriendshipUpsertOne struct {
		create *FriendshipCreate
	}

	// FriendshipUpsert is the "OnConflict" setter.
	FriendshipUpsert struct {
		*sql.UpdateSet
	}
)

// SetWeight sets the "weight" field.
func (u *FriendshipUpsert) SetWeight(v int) *FriendshipUpsert {
	u.Set(friendship.FieldWeight, v)
	return u
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateWeight() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldWeight)
	return u
}

// AddWeight adds v to the "weight" field.
func (u *FriendshipUpsert) AddWeight(v int) *FriendshipUpsert {
	u.Add(friendship.FieldWeight, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FriendshipUpsert) SetCreatedAt(v time.Time) *FriendshipUpsert {
	u.Set(friendship.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FriendshipUpsert) UpdateCreatedAt() *FriendshipUpsert {
	u.SetExcluded(friendship.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FriendshipUpsertOne) UpdateNewValues() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(friendship.FieldUserID)
		}
		if _, exists := u.create.mutation.FriendID(); exists {
			s.SetIgnore(friendship.FieldFriendID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FriendshipUpsertOne) Ignore() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FriendshipUpsertOne) DoNothing() *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FriendshipCreate.OnConflict
// documentation for more info.
func (u *FriendshipUpsertOne) Update(set func(*FriendshipUpsert)) *FriendshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FriendshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetWeight sets the "weight" field.
func (u *FriendshipUpsertOne) SetWeight(v int) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetWeight(v)
	})
}

// AddWeight adds v to the "weight" field.
func (u *FriendshipUpsertOne) AddWeight(v int) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.AddWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateWeight() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateWeight()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FriendshipUpsertOne) SetCreatedAt(v time.Time) *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FriendshipUpsertOne) UpdateCreatedAt() *FriendshipUpsertOne {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FriendshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FriendshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FriendshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FriendshipUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FriendshipUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FriendshipCreateBulk is the builder for creating many Friendship entities in bulk.
type FriendshipCreateBulk struct {
	config
	builders []*FriendshipCreate
	conflict []sql.ConflictOption
}

// Save creates the Friendship entities in the database.
func (fcb *FriendshipCreateBulk) Save(ctx context.Context) ([]*Friendship, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Friendship, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) SaveX(ctx context.Context) []*Friendship {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FriendshipCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FriendshipCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Friendship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FriendshipUpsert) {
//			SetWeight(v+v).
//		}).
//		Exec(ctx)
func (fcb *FriendshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *FriendshipUpsertBulk {
	fcb.conflict = opts
	return &FriendshipUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FriendshipCreateBulk) OnConflictColumns(columns ...string) *FriendshipUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FriendshipUpsertBulk{
		create: fcb,
	}
}

// FriendshipUpsertBulk is the builder for "upsert"-ing
// a bulk of Friendship nodes.
type FriendshipUpsertBulk struct {
	create *FriendshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FriendshipUpsertBulk) UpdateNewValues() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(friendship.FieldUserID)
			}
			if _, exists := b.mutation.FriendID(); exists {
				s.SetIgnore(friendship.FieldFriendID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Friendship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FriendshipUpsertBulk) Ignore() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FriendshipUpsertBulk) DoNothing() *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FriendshipCreateBulk.OnConflict
// documentation for more info.
func (u *FriendshipUpsertBulk) Update(set func(*FriendshipUpsert)) *FriendshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FriendshipUpsert{UpdateSet: update})
	}))
	return u
}

// SetWeight sets the "weight" field.
func (u *FriendshipUpsertBulk) SetWeight(v int) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetWeight(v)
	})
}

// AddWeight adds v to the "weight" field.
func (u *FriendshipUpsertBulk) AddWeight(v int) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.AddWeight(v)
	})
}

// UpdateWeight sets the "weight" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateWeight() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateWeight()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FriendshipUpsertBulk) SetCreatedAt(v time.Time) *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FriendshipUpsertBulk) UpdateCreatedAt() *FriendshipUpsertBulk {
	return u.Update(func(s *FriendshipUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FriendshipUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FriendshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FriendshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FriendshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
