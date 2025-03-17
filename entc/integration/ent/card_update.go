// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/spec"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks     []Hook
	mutation  *CardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CardUpdate builder.
func (u *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (m *CardUpdate) SetUpdateTime(v time.Time) *CardUpdate {
	m.mutation.SetUpdateTime(v)
	return m
}

// SetBalance sets the "balance" field.
func (m *CardUpdate) SetBalance(v float64) *CardUpdate {
	m.mutation.ResetBalance()
	m.mutation.SetBalance(v)
	return m
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (m *CardUpdate) SetNillableBalance(v *float64) *CardUpdate {
	if v != nil {
		m.SetBalance(*v)
	}
	return m
}

// AddBalance adds value to the "balance" field.
func (m *CardUpdate) AddBalance(v float64) *CardUpdate {
	m.mutation.AddBalance(v)
	return m
}

// SetName sets the "name" field.
func (m *CardUpdate) SetName(v string) *CardUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *CardUpdate) SetNillableName(v *string) *CardUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// ClearName clears the value of the "name" field.
func (m *CardUpdate) ClearName() *CardUpdate {
	m.mutation.ClearName()
	return m
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (m *CardUpdate) SetOwnerID(id int) *CardUpdate {
	m.mutation.SetOwnerID(id)
	return m
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (m *CardUpdate) SetNillableOwnerID(id *int) *CardUpdate {
	if id != nil {
		m = m.SetOwnerID(*id)
	}
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *CardUpdate) SetOwner(v *User) *CardUpdate {
	return m.SetOwnerID(v.ID)
}

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (m *CardUpdate) AddSpecIDs(ids ...int) *CardUpdate {
	m.mutation.AddSpecIDs(ids...)
	return m
}

// AddSpec adds the "spec" edges to the Spec entity.
func (m *CardUpdate) AddSpec(v ...*Spec) *CardUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddSpecIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (m *CardUpdate) Mutation() *CardMutation {
	return m.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (u *CardUpdate) ClearOwner() *CardUpdate {
	u.mutation.ClearOwner()
	return u
}

// ClearSpec clears all "spec" edges to the Spec entity.
func (u *CardUpdate) ClearSpec() *CardUpdate {
	u.mutation.ClearSpec()
	return u
}

// RemoveSpecIDs removes the "spec" edge to Spec entities by IDs.
func (u *CardUpdate) RemoveSpecIDs(ids ...int) *CardUpdate {
	u.mutation.RemoveSpecIDs(ids...)
	return u
}

// RemoveSpec removes "spec" edges to Spec entities.
func (u *CardUpdate) RemoveSpec(v ...*Spec) *CardUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *CardUpdate) Save(ctx context.Context) (int, error) {
	u.defaults()
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *CardUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (u *CardUpdate) defaults() {
	if _, ok := u.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		u.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *CardUpdate) check() error {
	if v, ok := u.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (u *CardUpdate) Modify(modifiers ...func(*sql.UpdateBuilder)) *CardUpdate {
	u.modifiers = append(u.modifiers, modifiers...)
	return u
}

func (u *CardUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.UpdateTime(); ok {
		_spec.SetField(card.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := u.mutation.Balance(); ok {
		_spec.SetField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := u.mutation.AddedBalance(); ok {
		_spec.AddField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeString, value)
	}
	if u.mutation.NameCleared() {
		_spec.ClearField(card.FieldName, field.TypeString)
	}
	if u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.SpecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedSpecIDs(); len(nodes) > 0 && !u.mutation.SpecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.SpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(u.modifiers...)
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (m *CardUpdateOne) SetUpdateTime(v time.Time) *CardUpdateOne {
	m.mutation.SetUpdateTime(v)
	return m
}

// SetBalance sets the "balance" field.
func (m *CardUpdateOne) SetBalance(v float64) *CardUpdateOne {
	m.mutation.ResetBalance()
	m.mutation.SetBalance(v)
	return m
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (m *CardUpdateOne) SetNillableBalance(v *float64) *CardUpdateOne {
	if v != nil {
		m.SetBalance(*v)
	}
	return m
}

// AddBalance adds value to the "balance" field.
func (m *CardUpdateOne) AddBalance(v float64) *CardUpdateOne {
	m.mutation.AddBalance(v)
	return m
}

// SetName sets the "name" field.
func (m *CardUpdateOne) SetName(v string) *CardUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *CardUpdateOne) SetNillableName(v *string) *CardUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// ClearName clears the value of the "name" field.
func (m *CardUpdateOne) ClearName() *CardUpdateOne {
	m.mutation.ClearName()
	return m
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (m *CardUpdateOne) SetOwnerID(id int) *CardUpdateOne {
	m.mutation.SetOwnerID(id)
	return m
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (m *CardUpdateOne) SetNillableOwnerID(id *int) *CardUpdateOne {
	if id != nil {
		m = m.SetOwnerID(*id)
	}
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *CardUpdateOne) SetOwner(v *User) *CardUpdateOne {
	return m.SetOwnerID(v.ID)
}

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (m *CardUpdateOne) AddSpecIDs(ids ...int) *CardUpdateOne {
	m.mutation.AddSpecIDs(ids...)
	return m
}

// AddSpec adds the "spec" edges to the Spec entity.
func (m *CardUpdateOne) AddSpec(v ...*Spec) *CardUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddSpecIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (m *CardUpdateOne) Mutation() *CardMutation {
	return m.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (u *CardUpdateOne) ClearOwner() *CardUpdateOne {
	u.mutation.ClearOwner()
	return u
}

// ClearSpec clears all "spec" edges to the Spec entity.
func (u *CardUpdateOne) ClearSpec() *CardUpdateOne {
	u.mutation.ClearSpec()
	return u
}

// RemoveSpecIDs removes the "spec" edge to Spec entities by IDs.
func (u *CardUpdateOne) RemoveSpecIDs(ids ...int) *CardUpdateOne {
	u.mutation.RemoveSpecIDs(ids...)
	return u
}

// RemoveSpec removes "spec" edges to Spec entities.
func (u *CardUpdateOne) RemoveSpec(v ...*Spec) *CardUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveSpecIDs(ids...)
}

// Where appends a list predicates to the CardUpdate builder.
func (u *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Card entity.
func (u *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	u.defaults()
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (u *CardUpdateOne) defaults() {
	if _, ok := u.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		u.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *CardUpdateOne) check() error {
	if v, ok := u.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (u *CardUpdateOne) Modify(modifiers ...func(*sql.UpdateBuilder)) *CardUpdateOne {
	u.modifiers = append(u.modifiers, modifiers...)
	return u
}

func (u *CardUpdateOne) sqlSave(ctx context.Context) (_n *Card, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.UpdateTime(); ok {
		_spec.SetField(card.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := u.mutation.Balance(); ok {
		_spec.SetField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := u.mutation.AddedBalance(); ok {
		_spec.AddField(card.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeString, value)
	}
	if u.mutation.NameCleared() {
		_spec.ClearField(card.FieldName, field.TypeString)
	}
	if u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.SpecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedSpecIDs(); len(nodes) > 0 && !u.mutation.SpecCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.SpecIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   card.SpecTable,
			Columns: card.SpecPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(u.modifiers...)
	_n = &Card{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
