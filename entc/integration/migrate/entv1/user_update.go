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
func (_u *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetAge sets the "age" field.
func (_u *UserUpdate) SetAge(v int32) *UserUpdate {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(v)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdate) SetNillableAge(v *int32) *UserUpdate {
	if v != nil {
		_u.SetAge(*v)
	}
	return _u
}

// AddAge adds value to the "age" field.
func (_u *UserUpdate) AddAge(v int32) *UserUpdate {
	_u.mutation.AddAge(v)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdate) SetName(v string) *UserUpdate {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdate) SetNillableName(v *string) *UserUpdate {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetDescription sets the "description" field.
func (_u *UserUpdate) SetDescription(v string) *UserUpdate {
	_u.mutation.SetDescription(v)
	return _u
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (_u *UserUpdate) SetNillableDescription(v *string) *UserUpdate {
	if v != nil {
		_u.SetDescription(*v)
	}
	return _u
}

// ClearDescription clears the value of the "description" field.
func (_u *UserUpdate) ClearDescription() *UserUpdate {
	_u.mutation.ClearDescription()
	return _u
}

// SetNickname sets the "nickname" field.
func (_u *UserUpdate) SetNickname(v string) *UserUpdate {
	_u.mutation.SetNickname(v)
	return _u
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (_u *UserUpdate) SetNillableNickname(v *string) *UserUpdate {
	if v != nil {
		_u.SetNickname(*v)
	}
	return _u
}

// SetAddress sets the "address" field.
func (_u *UserUpdate) SetAddress(v string) *UserUpdate {
	_u.mutation.SetAddress(v)
	return _u
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (_u *UserUpdate) SetNillableAddress(v *string) *UserUpdate {
	if v != nil {
		_u.SetAddress(*v)
	}
	return _u
}

// ClearAddress clears the value of the "address" field.
func (_u *UserUpdate) ClearAddress() *UserUpdate {
	_u.mutation.ClearAddress()
	return _u
}

// SetRenamed sets the "renamed" field.
func (_u *UserUpdate) SetRenamed(v string) *UserUpdate {
	_u.mutation.SetRenamed(v)
	return _u
}

// SetNillableRenamed sets the "renamed" field if the given value is not nil.
func (_u *UserUpdate) SetNillableRenamed(v *string) *UserUpdate {
	if v != nil {
		_u.SetRenamed(*v)
	}
	return _u
}

// ClearRenamed clears the value of the "renamed" field.
func (_u *UserUpdate) ClearRenamed() *UserUpdate {
	_u.mutation.ClearRenamed()
	return _u
}

// SetOldToken sets the "old_token" field.
func (_u *UserUpdate) SetOldToken(v string) *UserUpdate {
	_u.mutation.SetOldToken(v)
	return _u
}

// SetNillableOldToken sets the "old_token" field if the given value is not nil.
func (_u *UserUpdate) SetNillableOldToken(v *string) *UserUpdate {
	if v != nil {
		_u.SetOldToken(*v)
	}
	return _u
}

// SetBlob sets the "blob" field.
func (_u *UserUpdate) SetBlob(v []byte) *UserUpdate {
	_u.mutation.SetBlob(v)
	return _u
}

// ClearBlob clears the value of the "blob" field.
func (_u *UserUpdate) ClearBlob() *UserUpdate {
	_u.mutation.ClearBlob()
	return _u
}

// SetState sets the "state" field.
func (_u *UserUpdate) SetState(v user.State) *UserUpdate {
	_u.mutation.SetState(v)
	return _u
}

// SetNillableState sets the "state" field if the given value is not nil.
func (_u *UserUpdate) SetNillableState(v *user.State) *UserUpdate {
	if v != nil {
		_u.SetState(*v)
	}
	return _u
}

// ClearState clears the value of the "state" field.
func (_u *UserUpdate) ClearState() *UserUpdate {
	_u.mutation.ClearState()
	return _u
}

// SetStatus sets the "status" field.
func (_u *UserUpdate) SetStatus(v string) *UserUpdate {
	_u.mutation.SetStatus(v)
	return _u
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (_u *UserUpdate) SetNillableStatus(v *string) *UserUpdate {
	if v != nil {
		_u.SetStatus(*v)
	}
	return _u
}

// ClearStatus clears the value of the "status" field.
func (_u *UserUpdate) ClearStatus() *UserUpdate {
	_u.mutation.ClearStatus()
	return _u
}

// SetWorkplace sets the "workplace" field.
func (_u *UserUpdate) SetWorkplace(v string) *UserUpdate {
	_u.mutation.SetWorkplace(v)
	return _u
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (_u *UserUpdate) SetNillableWorkplace(v *string) *UserUpdate {
	if v != nil {
		_u.SetWorkplace(*v)
	}
	return _u
}

// ClearWorkplace clears the value of the "workplace" field.
func (_u *UserUpdate) ClearWorkplace() *UserUpdate {
	_u.mutation.ClearWorkplace()
	return _u
}

// SetDropOptional sets the "drop_optional" field.
func (_u *UserUpdate) SetDropOptional(v string) *UserUpdate {
	_u.mutation.SetDropOptional(v)
	return _u
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (_u *UserUpdate) SetNillableDropOptional(v *string) *UserUpdate {
	if v != nil {
		_u.SetDropOptional(*v)
	}
	return _u
}

// ClearDropOptional clears the value of the "drop_optional" field.
func (_u *UserUpdate) ClearDropOptional() *UserUpdate {
	_u.mutation.ClearDropOptional()
	return _u
}

// SetParentID sets the "parent" edge to the User entity by ID.
func (_u *UserUpdate) SetParentID(id int) *UserUpdate {
	_u.mutation.SetParentID(id)
	return _u
}

// SetNillableParentID sets the "parent" edge to the User entity by ID if the given value is not nil.
func (_u *UserUpdate) SetNillableParentID(id *int) *UserUpdate {
	if id != nil {
		_u = _u.SetParentID(*id)
	}
	return _u
}

// SetParent sets the "parent" edge to the User entity.
func (_u *UserUpdate) SetParent(v *User) *UserUpdate {
	return _u.SetParentID(v.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (_u *UserUpdate) AddChildIDs(ids ...int) *UserUpdate {
	_u.mutation.AddChildIDs(ids...)
	return _u
}

// AddChildren adds the "children" edges to the User entity.
func (_u *UserUpdate) AddChildren(v ...*User) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddChildIDs(ids...)
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (_u *UserUpdate) SetSpouseID(id int) *UserUpdate {
	_u.mutation.SetSpouseID(id)
	return _u
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (_u *UserUpdate) SetNillableSpouseID(id *int) *UserUpdate {
	if id != nil {
		_u = _u.SetSpouseID(*id)
	}
	return _u
}

// SetSpouse sets the "spouse" edge to the User entity.
func (_u *UserUpdate) SetSpouse(v *User) *UserUpdate {
	return _u.SetSpouseID(v.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (_u *UserUpdate) SetCarID(id int) *UserUpdate {
	_u.mutation.SetCarID(id)
	return _u
}

// SetNillableCarID sets the "car" edge to the Car entity by ID if the given value is not nil.
func (_u *UserUpdate) SetNillableCarID(id *int) *UserUpdate {
	if id != nil {
		_u = _u.SetCarID(*id)
	}
	return _u
}

// SetCar sets the "car" edge to the Car entity.
func (_u *UserUpdate) SetCar(v *Car) *UserUpdate {
	return _u.SetCarID(v.ID)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdate) Mutation() *UserMutation {
	return _u.mutation
}

// ClearParent clears the "parent" edge to the User entity.
func (_u *UserUpdate) ClearParent() *UserUpdate {
	_u.mutation.ClearParent()
	return _u
}

// ClearChildren clears all "children" edges to the User entity.
func (_u *UserUpdate) ClearChildren() *UserUpdate {
	_u.mutation.ClearChildren()
	return _u
}

// RemoveChildIDs removes the "children" edge to User entities by IDs.
func (_u *UserUpdate) RemoveChildIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveChildIDs(ids...)
	return _u
}

// RemoveChildren removes "children" edges to User entities.
func (_u *UserUpdate) RemoveChildren(v ...*User) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveChildIDs(ids...)
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (_u *UserUpdate) ClearSpouse() *UserUpdate {
	_u.mutation.ClearSpouse()
	return _u
}

// ClearCar clears the "car" edge to the Car entity.
func (_u *UserUpdate) ClearCar() *UserUpdate {
	_u.mutation.ClearCar()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *UserUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *UserUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *UserUpdate) check() error {
	if v, ok := _u.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entv1: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`entv1: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := _u.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`entv1: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Workplace(); ok {
		if err := user.WorkplaceValidator(v); err != nil {
			return &ValidationError{Name: "workplace", err: fmt.Errorf(`entv1: validator failed for field "User.workplace": %w`, err)}
		}
	}
	return nil
}

func (_u *UserUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := _u.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if _u.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := _u.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := _u.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if _u.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := _u.mutation.Renamed(); ok {
		_spec.SetField(user.FieldRenamed, field.TypeString, value)
	}
	if _u.mutation.RenamedCleared() {
		_spec.ClearField(user.FieldRenamed, field.TypeString)
	}
	if value, ok := _u.mutation.OldToken(); ok {
		_spec.SetField(user.FieldOldToken, field.TypeString, value)
	}
	if value, ok := _u.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
	}
	if _u.mutation.BlobCleared() {
		_spec.ClearField(user.FieldBlob, field.TypeBytes)
	}
	if value, ok := _u.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
	}
	if _u.mutation.StateCleared() {
		_spec.ClearField(user.FieldState, field.TypeEnum)
	}
	if value, ok := _u.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if _u.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeString)
	}
	if value, ok := _u.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
	}
	if _u.mutation.WorkplaceCleared() {
		_spec.ClearField(user.FieldWorkplace, field.TypeString)
	}
	if value, ok := _u.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
	}
	if _u.mutation.DropOptionalCleared() {
		_spec.ClearField(user.FieldDropOptional, field.TypeString)
	}
	if _u.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
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
	if _u.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !_u.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
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
	if _u.mutation.SpouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (_u *UserUpdateOne) SetAge(v int32) *UserUpdateOne {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(v)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableAge(v *int32) *UserUpdateOne {
	if v != nil {
		_u.SetAge(*v)
	}
	return _u
}

// AddAge adds value to the "age" field.
func (_u *UserUpdateOne) AddAge(v int32) *UserUpdateOne {
	_u.mutation.AddAge(v)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdateOne) SetName(v string) *UserUpdateOne {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableName(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetDescription sets the "description" field.
func (_u *UserUpdateOne) SetDescription(v string) *UserUpdateOne {
	_u.mutation.SetDescription(v)
	return _u
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableDescription(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetDescription(*v)
	}
	return _u
}

// ClearDescription clears the value of the "description" field.
func (_u *UserUpdateOne) ClearDescription() *UserUpdateOne {
	_u.mutation.ClearDescription()
	return _u
}

// SetNickname sets the "nickname" field.
func (_u *UserUpdateOne) SetNickname(v string) *UserUpdateOne {
	_u.mutation.SetNickname(v)
	return _u
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableNickname(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetNickname(*v)
	}
	return _u
}

// SetAddress sets the "address" field.
func (_u *UserUpdateOne) SetAddress(v string) *UserUpdateOne {
	_u.mutation.SetAddress(v)
	return _u
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableAddress(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetAddress(*v)
	}
	return _u
}

// ClearAddress clears the value of the "address" field.
func (_u *UserUpdateOne) ClearAddress() *UserUpdateOne {
	_u.mutation.ClearAddress()
	return _u
}

// SetRenamed sets the "renamed" field.
func (_u *UserUpdateOne) SetRenamed(v string) *UserUpdateOne {
	_u.mutation.SetRenamed(v)
	return _u
}

// SetNillableRenamed sets the "renamed" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableRenamed(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetRenamed(*v)
	}
	return _u
}

// ClearRenamed clears the value of the "renamed" field.
func (_u *UserUpdateOne) ClearRenamed() *UserUpdateOne {
	_u.mutation.ClearRenamed()
	return _u
}

// SetOldToken sets the "old_token" field.
func (_u *UserUpdateOne) SetOldToken(v string) *UserUpdateOne {
	_u.mutation.SetOldToken(v)
	return _u
}

// SetNillableOldToken sets the "old_token" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableOldToken(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetOldToken(*v)
	}
	return _u
}

// SetBlob sets the "blob" field.
func (_u *UserUpdateOne) SetBlob(v []byte) *UserUpdateOne {
	_u.mutation.SetBlob(v)
	return _u
}

// ClearBlob clears the value of the "blob" field.
func (_u *UserUpdateOne) ClearBlob() *UserUpdateOne {
	_u.mutation.ClearBlob()
	return _u
}

// SetState sets the "state" field.
func (_u *UserUpdateOne) SetState(v user.State) *UserUpdateOne {
	_u.mutation.SetState(v)
	return _u
}

// SetNillableState sets the "state" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableState(v *user.State) *UserUpdateOne {
	if v != nil {
		_u.SetState(*v)
	}
	return _u
}

// ClearState clears the value of the "state" field.
func (_u *UserUpdateOne) ClearState() *UserUpdateOne {
	_u.mutation.ClearState()
	return _u
}

// SetStatus sets the "status" field.
func (_u *UserUpdateOne) SetStatus(v string) *UserUpdateOne {
	_u.mutation.SetStatus(v)
	return _u
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableStatus(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetStatus(*v)
	}
	return _u
}

// ClearStatus clears the value of the "status" field.
func (_u *UserUpdateOne) ClearStatus() *UserUpdateOne {
	_u.mutation.ClearStatus()
	return _u
}

// SetWorkplace sets the "workplace" field.
func (_u *UserUpdateOne) SetWorkplace(v string) *UserUpdateOne {
	_u.mutation.SetWorkplace(v)
	return _u
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableWorkplace(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetWorkplace(*v)
	}
	return _u
}

// ClearWorkplace clears the value of the "workplace" field.
func (_u *UserUpdateOne) ClearWorkplace() *UserUpdateOne {
	_u.mutation.ClearWorkplace()
	return _u
}

// SetDropOptional sets the "drop_optional" field.
func (_u *UserUpdateOne) SetDropOptional(v string) *UserUpdateOne {
	_u.mutation.SetDropOptional(v)
	return _u
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableDropOptional(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetDropOptional(*v)
	}
	return _u
}

// ClearDropOptional clears the value of the "drop_optional" field.
func (_u *UserUpdateOne) ClearDropOptional() *UserUpdateOne {
	_u.mutation.ClearDropOptional()
	return _u
}

// SetParentID sets the "parent" edge to the User entity by ID.
func (_u *UserUpdateOne) SetParentID(id int) *UserUpdateOne {
	_u.mutation.SetParentID(id)
	return _u
}

// SetNillableParentID sets the "parent" edge to the User entity by ID if the given value is not nil.
func (_u *UserUpdateOne) SetNillableParentID(id *int) *UserUpdateOne {
	if id != nil {
		_u = _u.SetParentID(*id)
	}
	return _u
}

// SetParent sets the "parent" edge to the User entity.
func (_u *UserUpdateOne) SetParent(v *User) *UserUpdateOne {
	return _u.SetParentID(v.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (_u *UserUpdateOne) AddChildIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddChildIDs(ids...)
	return _u
}

// AddChildren adds the "children" edges to the User entity.
func (_u *UserUpdateOne) AddChildren(v ...*User) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddChildIDs(ids...)
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (_u *UserUpdateOne) SetSpouseID(id int) *UserUpdateOne {
	_u.mutation.SetSpouseID(id)
	return _u
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (_u *UserUpdateOne) SetNillableSpouseID(id *int) *UserUpdateOne {
	if id != nil {
		_u = _u.SetSpouseID(*id)
	}
	return _u
}

// SetSpouse sets the "spouse" edge to the User entity.
func (_u *UserUpdateOne) SetSpouse(v *User) *UserUpdateOne {
	return _u.SetSpouseID(v.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (_u *UserUpdateOne) SetCarID(id int) *UserUpdateOne {
	_u.mutation.SetCarID(id)
	return _u
}

// SetNillableCarID sets the "car" edge to the Car entity by ID if the given value is not nil.
func (_u *UserUpdateOne) SetNillableCarID(id *int) *UserUpdateOne {
	if id != nil {
		_u = _u.SetCarID(*id)
	}
	return _u
}

// SetCar sets the "car" edge to the Car entity.
func (_u *UserUpdateOne) SetCar(v *Car) *UserUpdateOne {
	return _u.SetCarID(v.ID)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdateOne) Mutation() *UserMutation {
	return _u.mutation
}

// ClearParent clears the "parent" edge to the User entity.
func (_u *UserUpdateOne) ClearParent() *UserUpdateOne {
	_u.mutation.ClearParent()
	return _u
}

// ClearChildren clears all "children" edges to the User entity.
func (_u *UserUpdateOne) ClearChildren() *UserUpdateOne {
	_u.mutation.ClearChildren()
	return _u
}

// RemoveChildIDs removes the "children" edge to User entities by IDs.
func (_u *UserUpdateOne) RemoveChildIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveChildIDs(ids...)
	return _u
}

// RemoveChildren removes "children" edges to User entities.
func (_u *UserUpdateOne) RemoveChildren(v ...*User) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveChildIDs(ids...)
}

// ClearSpouse clears the "spouse" edge to the User entity.
func (_u *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	_u.mutation.ClearSpouse()
	return _u
}

// ClearCar clears the "car" edge to the Car entity.
func (_u *UserUpdateOne) ClearCar() *UserUpdateOne {
	_u.mutation.ClearCar()
	return _u
}

// Where appends a list predicates to the UserUpdate builder.
func (_u *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated User entity.
func (_u *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *UserUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *UserUpdateOne) check() error {
	if v, ok := _u.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entv1: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`entv1: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := _u.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`entv1: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := _u.mutation.Workplace(); ok {
		if err := user.WorkplaceValidator(v); err != nil {
			return &ValidationError{Name: "workplace", err: fmt.Errorf(`entv1: validator failed for field "User.workplace": %w`, err)}
		}
	}
	return nil
}

func (_u *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv1: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
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
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt32, value)
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := _u.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
	}
	if _u.mutation.DescriptionCleared() {
		_spec.ClearField(user.FieldDescription, field.TypeString)
	}
	if value, ok := _u.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := _u.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if _u.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := _u.mutation.Renamed(); ok {
		_spec.SetField(user.FieldRenamed, field.TypeString, value)
	}
	if _u.mutation.RenamedCleared() {
		_spec.ClearField(user.FieldRenamed, field.TypeString)
	}
	if value, ok := _u.mutation.OldToken(); ok {
		_spec.SetField(user.FieldOldToken, field.TypeString, value)
	}
	if value, ok := _u.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
	}
	if _u.mutation.BlobCleared() {
		_spec.ClearField(user.FieldBlob, field.TypeBytes)
	}
	if value, ok := _u.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
	}
	if _u.mutation.StateCleared() {
		_spec.ClearField(user.FieldState, field.TypeEnum)
	}
	if value, ok := _u.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if _u.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeString)
	}
	if value, ok := _u.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
	}
	if _u.mutation.WorkplaceCleared() {
		_spec.ClearField(user.FieldWorkplace, field.TypeString)
	}
	if value, ok := _u.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
	}
	if _u.mutation.DropOptionalCleared() {
		_spec.ClearField(user.FieldDropOptional, field.TypeString)
	}
	if _u.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
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
	if _u.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !_u.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
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
	if _u.mutation.SpouseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
