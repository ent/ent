// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv1/car"
	"entgo.io/ent/entc/integration/migrate/entv1/predicate"
	"entgo.io/ent/entc/integration/migrate/entv1/user"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int32) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int32) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetDescription sets the "description" field.
func (uu *UserUpdate) SetDescription(s string) *UserUpdate {
	uu.mutation.SetDescription(s)
	return uu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDescription(s *string) *UserUpdate {
	if s != nil {
		uu.SetDescription(*s)
	}
	return uu
}

// ClearDescription clears the value of the "description" field.
func (uu *UserUpdate) ClearDescription() *UserUpdate {
	uu.mutation.ClearDescription()
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of the "address" field.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.mutation.ClearAddress()
	return uu
}

// SetRenamed sets the "renamed" field.
func (uu *UserUpdate) SetRenamed(s string) *UserUpdate {
	uu.mutation.SetRenamed(s)
	return uu
}

// SetNillableRenamed sets the "renamed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRenamed(s *string) *UserUpdate {
	if s != nil {
		uu.SetRenamed(*s)
	}
	return uu
}

// ClearRenamed clears the value of the "renamed" field.
func (uu *UserUpdate) ClearRenamed() *UserUpdate {
	uu.mutation.ClearRenamed()
	return uu
}

// SetOldToken sets the "old_token" field.
func (uu *UserUpdate) SetOldToken(s string) *UserUpdate {
	uu.mutation.SetOldToken(s)
	return uu
}

// SetNillableOldToken sets the "old_token" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOldToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetOldToken(*s)
	}
	return uu
}

// SetBlob sets the "blob" field.
func (uu *UserUpdate) SetBlob(b []byte) *UserUpdate {
	uu.mutation.SetBlob(b)
	return uu
}

// ClearBlob clears the value of the "blob" field.
func (uu *UserUpdate) ClearBlob() *UserUpdate {
	uu.mutation.ClearBlob()
	return uu
}

// SetState sets the "state" field.
func (uu *UserUpdate) SetState(u user.State) *UserUpdate {
	uu.mutation.SetState(u)
	return uu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(u *user.State) *UserUpdate {
	if u != nil {
		uu.SetState(*u)
	}
	return uu
}

// ClearState clears the value of the "state" field.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.mutation.ClearState()
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(s string) *UserUpdate {
	uu.mutation.SetStatus(s)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(s *string) *UserUpdate {
	if s != nil {
		uu.SetStatus(*s)
	}
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// SetWorkplace sets the "workplace" field.
func (uu *UserUpdate) SetWorkplace(s string) *UserUpdate {
	uu.mutation.SetWorkplace(s)
	return uu
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWorkplace(s *string) *UserUpdate {
	if s != nil {
		uu.SetWorkplace(*s)
	}
	return uu
}

// ClearWorkplace clears the value of the "workplace" field.
func (uu *UserUpdate) ClearWorkplace() *UserUpdate {
	uu.mutation.ClearWorkplace()
	return uu
}

// SetDropOptional sets the "drop_optional" field.
func (uu *UserUpdate) SetDropOptional(s string) *UserUpdate {
	uu.mutation.SetDropOptional(s)
	return uu
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDropOptional(s *string) *UserUpdate {
	if s != nil {
		uu.SetDropOptional(*s)
	}
	return uu
}

// ClearDropOptional clears the value of the "drop_optional" field.
func (uu *UserUpdate) ClearDropOptional() *UserUpdate {
	uu.mutation.ClearDropOptional()
	return uu
}

// SetParentID sets the "parent" edge to the User entity by ID.
func (uu *UserUpdate) SetParentID(id int) *UserUpdate {
	uu.mutation.SetParentID(id)
	return uu
}

// SetNillableParentID sets the "parent" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableParentID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetParentID(*id)
	}
	return uu
}

// SetParent sets the "parent" edge to the User entity.
func (uu *UserUpdate) SetParent(u *User) *UserUpdate {
	return uu.SetParentID(u.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (uu *UserUpdate) AddChildIDs(ids ...int) *UserUpdate {
	uu.mutation.AddChildIDs(ids...)
	return uu
}

// AddChildren adds the "children" edges to the User entity.
func (uu *UserUpdate) AddChildren(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddChildIDs(ids...)
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (uu *UserUpdate) SetSpouseID(id int) *UserUpdate {
	uu.mutation.SetSpouseID(id)
	return uu
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableSpouseID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetSpouseID(*id)
	}
	return uu
}

// SetSpouse sets the "spouse" edge to the User entity.
func (uu *UserUpdate) SetSpouse(u *User) *UserUpdate {
	return uu.SetSpouseID(u.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (uu *UserUpdate) SetCarID(id int) *UserUpdate {
	uu.mutation.SetCarID(id)
	return uu
}

// SetNillableCarID sets the "car" edge to the Car entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableCarID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetCarID(*id)
	}
	return uu
}

// SetCar sets the "car" edge to the Car entity.
func (uu *UserUpdate) SetCar(c *Car) *UserUpdate {
	return uu.SetCarID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearParent clears the "parent" edge to the User entity.
func (uu *UserUpdate) ClearParent() *UserUpdate {
	uu.mutation.ClearParent()
	return uu
}

// ClearChildren clears all "children" edges to the User entity.
func (uu *UserUpdate) ClearChildren() *UserUpdate {
	uu.mutation.ClearChildren()
	return uu
}

// RemoveChildIDs removes the "children" edge to User entities by IDs.
func (uu *UserUpdate) RemoveChildIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveChildIDs(ids...)
	return uu
}

// RemoveChildren removes "children" edges to User entities.
func (uu *UserUpdate) RemoveChildren(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveChildIDs(ids...)
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (uu *UserUpdate) ClearSpouse() *UserUpdate {
	uu.mutation.ClearSpouse()
	return uu
}

// ClearCar clears the "car" edge to the Car entity.
func (uu *UserUpdate) ClearCar() *UserUpdate {
	uu.mutation.ClearCar()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entv1: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`entv1: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := uu.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`entv1: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Workplace(); ok {
		if err := user.WorkplaceValidator(v); err != nil {
			return &ValidationError{Name: "workplace", err: fmt.Errorf(`entv1: validator failed for field "User.workplace": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uu.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uu.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uu.mutation.Renamed(); ok {
		_spec.SetField(user.FieldRenamed, field.TypeString, value)
	}
	if uu.mutation.RenamedCleared() {
		_spec.ClearField(user.FieldRenamed, field.TypeString)
	}
	if value, ok := uu.mutation.OldToken(); ok {
		_spec.SetField(user.FieldOldToken, field.TypeString, value)
	}
	if value, ok := uu.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
	}
	if uu.mutation.BlobCleared() {
		_spec.ClearField(user.FieldBlob, field.TypeBytes)
	}
	if value, ok := uu.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
	}
	if uu.mutation.StateCleared() {
		_spec.ClearField(user.FieldState, field.TypeEnum)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if uu.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeString)
	}
	if value, ok := uu.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
	}
	if uu.mutation.WorkplaceCleared() {
		_spec.ClearField(user.FieldWorkplace, field.TypeString)
	}
	if value, ok := uu.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
	}
	if uu.mutation.DropOptionalCleared() {
		_spec.ClearField(user.FieldDropOptional, field.TypeString)
	}
	if uu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !uu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.SpouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int32) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int32) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetDescription sets the "description" field.
func (uuo *UserUpdateOne) SetDescription(s string) *UserUpdateOne {
	uuo.mutation.SetDescription(s)
	return uuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDescription(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDescription(*s)
	}
	return uuo
}

// ClearDescription clears the value of the "description" field.
func (uuo *UserUpdateOne) ClearDescription() *UserUpdateOne {
	uuo.mutation.ClearDescription()
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of the "address" field.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.mutation.ClearAddress()
	return uuo
}

// SetRenamed sets the "renamed" field.
func (uuo *UserUpdateOne) SetRenamed(s string) *UserUpdateOne {
	uuo.mutation.SetRenamed(s)
	return uuo
}

// SetNillableRenamed sets the "renamed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRenamed(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRenamed(*s)
	}
	return uuo
}

// ClearRenamed clears the value of the "renamed" field.
func (uuo *UserUpdateOne) ClearRenamed() *UserUpdateOne {
	uuo.mutation.ClearRenamed()
	return uuo
}

// SetOldToken sets the "old_token" field.
func (uuo *UserUpdateOne) SetOldToken(s string) *UserUpdateOne {
	uuo.mutation.SetOldToken(s)
	return uuo
}

// SetNillableOldToken sets the "old_token" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOldToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetOldToken(*s)
	}
	return uuo
}

// SetBlob sets the "blob" field.
func (uuo *UserUpdateOne) SetBlob(b []byte) *UserUpdateOne {
	uuo.mutation.SetBlob(b)
	return uuo
}

// ClearBlob clears the value of the "blob" field.
func (uuo *UserUpdateOne) ClearBlob() *UserUpdateOne {
	uuo.mutation.ClearBlob()
	return uuo
}

// SetState sets the "state" field.
func (uuo *UserUpdateOne) SetState(u user.State) *UserUpdateOne {
	uuo.mutation.SetState(u)
	return uuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(u *user.State) *UserUpdateOne {
	if u != nil {
		uuo.SetState(*u)
	}
	return uuo
}

// ClearState clears the value of the "state" field.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.mutation.ClearState()
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(s string) *UserUpdateOne {
	uuo.mutation.SetStatus(s)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetStatus(*s)
	}
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// SetWorkplace sets the "workplace" field.
func (uuo *UserUpdateOne) SetWorkplace(s string) *UserUpdateOne {
	uuo.mutation.SetWorkplace(s)
	return uuo
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWorkplace(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetWorkplace(*s)
	}
	return uuo
}

// ClearWorkplace clears the value of the "workplace" field.
func (uuo *UserUpdateOne) ClearWorkplace() *UserUpdateOne {
	uuo.mutation.ClearWorkplace()
	return uuo
}

// SetDropOptional sets the "drop_optional" field.
func (uuo *UserUpdateOne) SetDropOptional(s string) *UserUpdateOne {
	uuo.mutation.SetDropOptional(s)
	return uuo
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDropOptional(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDropOptional(*s)
	}
	return uuo
}

// ClearDropOptional clears the value of the "drop_optional" field.
func (uuo *UserUpdateOne) ClearDropOptional() *UserUpdateOne {
	uuo.mutation.ClearDropOptional()
	return uuo
}

// SetParentID sets the "parent" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetParentID(id int) *UserUpdateOne {
	uuo.mutation.SetParentID(id)
	return uuo
}

// SetNillableParentID sets the "parent" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableParentID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetParentID(*id)
	}
	return uuo
}

// SetParent sets the "parent" edge to the User entity.
func (uuo *UserUpdateOne) SetParent(u *User) *UserUpdateOne {
	return uuo.SetParentID(u.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddChildIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddChildIDs(ids...)
	return uuo
}

// AddChildren adds the "children" edges to the User entity.
func (uuo *UserUpdateOne) AddChildren(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddChildIDs(ids...)
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (uuo *UserUpdateOne) SetSpouseID(id int) *UserUpdateOne {
	uuo.mutation.SetSpouseID(id)
	return uuo
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSpouseID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSpouseID(*id)
	}
	return uuo
}

// SetSpouse sets the "spouse" edge to the User entity.
func (uuo *UserUpdateOne) SetSpouse(u *User) *UserUpdateOne {
	return uuo.SetSpouseID(u.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (uuo *UserUpdateOne) SetCarID(id int) *UserUpdateOne {
	uuo.mutation.SetCarID(id)
	return uuo
}

// SetNillableCarID sets the "car" edge to the Car entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCarID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCarID(*id)
	}
	return uuo
}

// SetCar sets the "car" edge to the Car entity.
func (uuo *UserUpdateOne) SetCar(c *Car) *UserUpdateOne {
	return uuo.SetCarID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearParent clears the "parent" edge to the User entity.
func (uuo *UserUpdateOne) ClearParent() *UserUpdateOne {
	uuo.mutation.ClearParent()
	return uuo
}

// ClearChildren clears all "children" edges to the User entity.
func (uuo *UserUpdateOne) ClearChildren() *UserUpdateOne {
	uuo.mutation.ClearChildren()
	return uuo
}

// RemoveChildIDs removes the "children" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveChildIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveChildIDs(ids...)
	return uuo
}

// RemoveChildren removes "children" edges to User entities.
func (uuo *UserUpdateOne) RemoveChildren(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveChildIDs(ids...)
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (uuo *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	uuo.mutation.ClearSpouse()
	return uuo
}

// ClearCar clears the "car" edge to the Car entity.
func (uuo *UserUpdateOne) ClearCar() *UserUpdateOne {
	uuo.mutation.ClearCar()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entv1: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`entv1: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`entv1: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Workplace(); ok {
		if err := user.WorkplaceValidator(v); err != nil {
			return &ValidationError{Name: "workplace", err: fmt.Errorf(`entv1: validator failed for field "User.workplace": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv1: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv1: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if uuo.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uuo.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uuo.mutation.Renamed(); ok {
		_spec.SetField(user.FieldRenamed, field.TypeString, value)
	}
	if uuo.mutation.RenamedCleared() {
		_spec.ClearField(user.FieldRenamed, field.TypeString)
	}
	if value, ok := uuo.mutation.OldToken(); ok {
		_spec.SetField(user.FieldOldToken, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
	}
	if uuo.mutation.BlobCleared() {
		_spec.ClearField(user.FieldBlob, field.TypeBytes)
	}
	if value, ok := uuo.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
	}
	if uuo.mutation.StateCleared() {
		_spec.ClearField(user.FieldState, field.TypeEnum)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if uuo.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeString)
	}
	if value, ok := uuo.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
	}
	if uuo.mutation.WorkplaceCleared() {
		_spec.ClearField(user.FieldWorkplace, field.TypeString)
	}
	if value, ok := uuo.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
	}
	if uuo.mutation.DropOptionalCleared() {
		_spec.ClearField(user.FieldDropOptional, field.TypeString)
	}
	if uuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !uuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.SpouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
