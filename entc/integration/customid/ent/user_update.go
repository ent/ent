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
	"entgo.io/ent/entc/integration/customid/ent/group"
	"entgo.io/ent/entc/integration/customid/ent/pet"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/user"
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

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
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

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddPetIDs(ids ...string) *UserUpdate {
	uu.mutation.AddPetIDs(ids...)
	return uu
}

// AddPets adds the "pets" edges to the Pet entity.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
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

// ClearPets clears all "pets" edges to the Pet entity.
func (uu *UserUpdate) ClearPets() *UserUpdate {
	uu.mutation.ClearPets()
	return uu
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uu *UserUpdate) RemovePetIDs(ids ...string) *UserUpdate {
	uu.mutation.RemovePetIDs(ids...)
	return uu
}

// RemovePets removes "pets" edges to Pet entities.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
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

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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
	if uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uu.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
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

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
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

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddPetIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.AddPetIDs(ids...)
	return uuo
}

// AddPets adds the "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
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

// ClearPets clears all "pets" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearPets() *UserUpdateOne {
	uuo.mutation.ClearPets()
	return uuo
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.RemovePetIDs(ids...)
	return uuo
}

// RemovePets removes "pets" edges to Pet entities.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
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

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
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
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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
	if uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !uuo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeString),
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
