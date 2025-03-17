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
	"entgo.io/ent/examples/traversal/ent/group"
	"entgo.io/ent/examples/traversal/ent/pet"
	"entgo.io/ent/examples/traversal/ent/predicate"
	"entgo.io/ent/examples/traversal/ent/user"
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
func (_u *UserUpdate) SetAge(i int) *UserUpdate {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(i)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdate) SetNillableAge(i *int) *UserUpdate {
	if i != nil {
		_u.SetAge(*i)
	}
	return _u
}

// AddAge adds i to the "age" field.
func (_u *UserUpdate) AddAge(i int) *UserUpdate {
	_u.mutation.AddAge(i)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdate) SetName(s string) *UserUpdate {
	_u.mutation.SetName(s)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		_u.SetName(*s)
	}
	return _u
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (_u *UserUpdate) AddPetIDs(ids ...int) *UserUpdate {
	_u.mutation.AddPetIDs(ids...)
	return _u
}

// AddPets adds the "pets" edges to the Pet entity.
func (_u *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return _u.AddPetIDs(ids...)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (_u *UserUpdate) AddFriendIDs(ids ...int) *UserUpdate {
	_u.mutation.AddFriendIDs(ids...)
	return _u
}

// AddFriends adds the "friends" edges to the User entity.
func (_u *UserUpdate) AddFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.AddFriendIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_u *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	_u.mutation.AddGroupIDs(ids...)
	return _u
}

// AddGroups adds the "groups" edges to the Group entity.
func (_u *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.AddGroupIDs(ids...)
}

// AddManageIDs adds the "manage" edge to the Group entity by IDs.
func (_u *UserUpdate) AddManageIDs(ids ...int) *UserUpdate {
	_u.mutation.AddManageIDs(ids...)
	return _u
}

// AddManage adds the "manage" edges to the Group entity.
func (_u *UserUpdate) AddManage(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.AddManageIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdate) Mutation() *UserMutation {
	return _u.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (_u *UserUpdate) ClearPets() *UserUpdate {
	_u.mutation.ClearPets()
	return _u
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (_u *UserUpdate) RemovePetIDs(ids ...int) *UserUpdate {
	_u.mutation.RemovePetIDs(ids...)
	return _u
}

// RemovePets removes "pets" edges to Pet entities.
func (_u *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return _u.RemovePetIDs(ids...)
}

// ClearFriends clears all "friends" edges to the User entity.
func (_u *UserUpdate) ClearFriends() *UserUpdate {
	_u.mutation.ClearFriends()
	return _u
}

// RemoveFriendIDs removes the "friends" edge to User entities by IDs.
func (_u *UserUpdate) RemoveFriendIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveFriendIDs(ids...)
	return _u
}

// RemoveFriends removes "friends" edges to User entities.
func (_u *UserUpdate) RemoveFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.RemoveFriendIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (_u *UserUpdate) ClearGroups() *UserUpdate {
	_u.mutation.ClearGroups()
	return _u
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (_u *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveGroupIDs(ids...)
	return _u
}

// RemoveGroups removes "groups" edges to Group entities.
func (_u *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.RemoveGroupIDs(ids...)
}

// ClearManage clears all "manage" edges to the Group entity.
func (_u *UserUpdate) ClearManage() *UserUpdate {
	_u.mutation.ClearManage()
	return _u
}

// RemoveManageIDs removes the "manage" edge to Group entities by IDs.
func (_u *UserUpdate) RemoveManageIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveManageIDs(ids...)
	return _u
}

// RemoveManage removes "manage" edges to Group entities.
func (_u *UserUpdate) RemoveManage(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.RemoveManageIDs(ids...)
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

func (_u *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if _u.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedPetsIDs(); len(nodes) > 0 && !_u.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !_u.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
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
	if _u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !_u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.ManageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ManageTable,
			Columns: []string{user.ManageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedManageIDs(); len(nodes) > 0 && !_u.mutation.ManageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ManageTable,
			Columns: []string{user.ManageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.ManageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ManageTable,
			Columns: []string{user.ManageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
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
func (_u *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(i)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableAge(i *int) *UserUpdateOne {
	if i != nil {
		_u.SetAge(*i)
	}
	return _u
}

// AddAge adds i to the "age" field.
func (_u *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	_u.mutation.AddAge(i)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdateOne) SetName(s string) *UserUpdateOne {
	_u.mutation.SetName(s)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		_u.SetName(*s)
	}
	return _u
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (_u *UserUpdateOne) AddPetIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddPetIDs(ids...)
	return _u
}

// AddPets adds the "pets" edges to the Pet entity.
func (_u *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return _u.AddPetIDs(ids...)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (_u *UserUpdateOne) AddFriendIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddFriendIDs(ids...)
	return _u
}

// AddFriends adds the "friends" edges to the User entity.
func (_u *UserUpdateOne) AddFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.AddFriendIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_u *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddGroupIDs(ids...)
	return _u
}

// AddGroups adds the "groups" edges to the Group entity.
func (_u *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.AddGroupIDs(ids...)
}

// AddManageIDs adds the "manage" edge to the Group entity by IDs.
func (_u *UserUpdateOne) AddManageIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddManageIDs(ids...)
	return _u
}

// AddManage adds the "manage" edges to the Group entity.
func (_u *UserUpdateOne) AddManage(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.AddManageIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdateOne) Mutation() *UserMutation {
	return _u.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (_u *UserUpdateOne) ClearPets() *UserUpdateOne {
	_u.mutation.ClearPets()
	return _u
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (_u *UserUpdateOne) RemovePetIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemovePetIDs(ids...)
	return _u
}

// RemovePets removes "pets" edges to Pet entities.
func (_u *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return _u.RemovePetIDs(ids...)
}

// ClearFriends clears all "friends" edges to the User entity.
func (_u *UserUpdateOne) ClearFriends() *UserUpdateOne {
	_u.mutation.ClearFriends()
	return _u
}

// RemoveFriendIDs removes the "friends" edge to User entities by IDs.
func (_u *UserUpdateOne) RemoveFriendIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveFriendIDs(ids...)
	return _u
}

// RemoveFriends removes "friends" edges to User entities.
func (_u *UserUpdateOne) RemoveFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.RemoveFriendIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (_u *UserUpdateOne) ClearGroups() *UserUpdateOne {
	_u.mutation.ClearGroups()
	return _u
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (_u *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveGroupIDs(ids...)
	return _u
}

// RemoveGroups removes "groups" edges to Group entities.
func (_u *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.RemoveGroupIDs(ids...)
}

// ClearManage clears all "manage" edges to the Group entity.
func (_u *UserUpdateOne) ClearManage() *UserUpdateOne {
	_u.mutation.ClearManage()
	return _u
}

// RemoveManageIDs removes the "manage" edge to Group entities by IDs.
func (_u *UserUpdateOne) RemoveManageIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveManageIDs(ids...)
	return _u
}

// RemoveManage removes "manage" edges to Group entities.
func (_u *UserUpdateOne) RemoveManage(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _u.RemoveManageIDs(ids...)
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

func (_u *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
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
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if _u.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedPetsIDs(); len(nodes) > 0 && !_u.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !_u.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
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
	if _u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !_u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.ManageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ManageTable,
			Columns: []string{user.ManageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedManageIDs(); len(nodes) > 0 && !_u.mutation.ManageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ManageTable,
			Columns: []string{user.ManageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.ManageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.ManageTable,
			Columns: []string{user.ManageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
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
