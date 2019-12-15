// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/file"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	size         *int
	addsize      *int
	name         *string
	user         *string
	clearuser    bool
	group        *string
	cleargroup   bool
	owner        map[string]struct{}
	_type        map[string]struct{}
	clearedOwner bool
	clearedType  bool
	predicates   []predicate.File
}

// Where adds a new predicate for the builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetSize sets the size field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.size = &i
	fu.addsize = nil
	return fu
}

// SetNillableSize sets the size field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to size.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	if fu.addsize == nil {
		fu.addsize = &i
	} else {
		*fu.addsize += i
	}
	return fu
}

// SetName sets the name field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.name = &s
	return fu
}

// SetUser sets the user field.
func (fu *FileUpdate) SetUser(s string) *FileUpdate {
	fu.user = &s
	return fu
}

// SetNillableUser sets the user field if the given value is not nil.
func (fu *FileUpdate) SetNillableUser(s *string) *FileUpdate {
	if s != nil {
		fu.SetUser(*s)
	}
	return fu
}

// ClearUser clears the value of user.
func (fu *FileUpdate) ClearUser() *FileUpdate {
	fu.user = nil
	fu.clearuser = true
	return fu
}

// SetGroup sets the group field.
func (fu *FileUpdate) SetGroup(s string) *FileUpdate {
	fu.group = &s
	return fu
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fu *FileUpdate) SetNillableGroup(s *string) *FileUpdate {
	if s != nil {
		fu.SetGroup(*s)
	}
	return fu
}

// ClearGroup clears the value of group.
func (fu *FileUpdate) ClearGroup() *FileUpdate {
	fu.group = nil
	fu.cleargroup = true
	return fu
}

// SetOwnerID sets the owner edge to User by id.
func (fu *FileUpdate) SetOwnerID(id string) *FileUpdate {
	if fu.owner == nil {
		fu.owner = make(map[string]struct{})
	}
	fu.owner[id] = struct{}{}
	return fu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fu *FileUpdate) SetNillableOwnerID(id *string) *FileUpdate {
	if id != nil {
		fu = fu.SetOwnerID(*id)
	}
	return fu
}

// SetOwner sets the owner edge to User.
func (fu *FileUpdate) SetOwner(u *User) *FileUpdate {
	return fu.SetOwnerID(u.ID)
}

// SetTypeID sets the type edge to FileType by id.
func (fu *FileUpdate) SetTypeID(id string) *FileUpdate {
	if fu._type == nil {
		fu._type = make(map[string]struct{})
	}
	fu._type[id] = struct{}{}
	return fu
}

// SetNillableTypeID sets the type edge to FileType by id if the given value is not nil.
func (fu *FileUpdate) SetNillableTypeID(id *string) *FileUpdate {
	if id != nil {
		fu = fu.SetTypeID(*id)
	}
	return fu
}

// SetType sets the type edge to FileType.
func (fu *FileUpdate) SetType(f *FileType) *FileUpdate {
	return fu.SetTypeID(f.ID)
}

// ClearOwner clears the owner edge to User.
func (fu *FileUpdate) ClearOwner() *FileUpdate {
	fu.clearedOwner = true
	return fu
}

// ClearType clears the type edge to FileType.
func (fu *FileUpdate) ClearType() *FileUpdate {
	fu.clearedType = true
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	if fu.size != nil {
		if err := file.SizeValidator(*fu.size); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	if len(fu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if len(fu._type) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"type\"")
	}
	return fu.gremlinSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := fu.gremlin().Query()
	if err := fu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (fu *FileUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(file.Label)
	for _, p := range fu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := fu.size; value != nil {
		v.Property(dsl.Single, file.FieldSize, *value)
	}
	if value := fu.addsize; value != nil {
		v.Property(dsl.Single, file.FieldSize, __.Union(__.Values(file.FieldSize), __.Constant(*value)).Sum())
	}
	if value := fu.name; value != nil {
		v.Property(dsl.Single, file.FieldName, *value)
	}
	if value := fu.user; value != nil {
		v.Property(dsl.Single, file.FieldUser, *value)
	}
	if value := fu.group; value != nil {
		v.Property(dsl.Single, file.FieldGroup, *value)
	}
	var properties []interface{}
	if fu.clearuser {
		properties = append(properties, file.FieldUser)
	}
	if fu.cleargroup {
		properties = append(properties, file.FieldGroup)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if fu.clearedOwner {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range fu.owner {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	if fu.clearedType {
		tr := rv.Clone().InE(filetype.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range fu._type {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	id           string
	size         *int
	addsize      *int
	name         *string
	user         *string
	clearuser    bool
	group        *string
	cleargroup   bool
	owner        map[string]struct{}
	_type        map[string]struct{}
	clearedOwner bool
	clearedType  bool
}

// SetSize sets the size field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.size = &i
	fuo.addsize = nil
	return fuo
}

// SetNillableSize sets the size field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to size.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	if fuo.addsize == nil {
		fuo.addsize = &i
	} else {
		*fuo.addsize += i
	}
	return fuo
}

// SetName sets the name field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.name = &s
	return fuo
}

// SetUser sets the user field.
func (fuo *FileUpdateOne) SetUser(s string) *FileUpdateOne {
	fuo.user = &s
	return fuo
}

// SetNillableUser sets the user field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUser(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetUser(*s)
	}
	return fuo
}

// ClearUser clears the value of user.
func (fuo *FileUpdateOne) ClearUser() *FileUpdateOne {
	fuo.user = nil
	fuo.clearuser = true
	return fuo
}

// SetGroup sets the group field.
func (fuo *FileUpdateOne) SetGroup(s string) *FileUpdateOne {
	fuo.group = &s
	return fuo
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableGroup(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetGroup(*s)
	}
	return fuo
}

// ClearGroup clears the value of group.
func (fuo *FileUpdateOne) ClearGroup() *FileUpdateOne {
	fuo.group = nil
	fuo.cleargroup = true
	return fuo
}

// SetOwnerID sets the owner edge to User by id.
func (fuo *FileUpdateOne) SetOwnerID(id string) *FileUpdateOne {
	if fuo.owner == nil {
		fuo.owner = make(map[string]struct{})
	}
	fuo.owner[id] = struct{}{}
	return fuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableOwnerID(id *string) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetOwnerID(*id)
	}
	return fuo
}

// SetOwner sets the owner edge to User.
func (fuo *FileUpdateOne) SetOwner(u *User) *FileUpdateOne {
	return fuo.SetOwnerID(u.ID)
}

// SetTypeID sets the type edge to FileType by id.
func (fuo *FileUpdateOne) SetTypeID(id string) *FileUpdateOne {
	if fuo._type == nil {
		fuo._type = make(map[string]struct{})
	}
	fuo._type[id] = struct{}{}
	return fuo
}

// SetNillableTypeID sets the type edge to FileType by id if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableTypeID(id *string) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetTypeID(*id)
	}
	return fuo
}

// SetType sets the type edge to FileType.
func (fuo *FileUpdateOne) SetType(f *FileType) *FileUpdateOne {
	return fuo.SetTypeID(f.ID)
}

// ClearOwner clears the owner edge to User.
func (fuo *FileUpdateOne) ClearOwner() *FileUpdateOne {
	fuo.clearedOwner = true
	return fuo
}

// ClearType clears the type edge to FileType.
func (fuo *FileUpdateOne) ClearType() *FileUpdateOne {
	fuo.clearedType = true
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	if fuo.size != nil {
		if err := file.SizeValidator(*fuo.size); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	if len(fuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if len(fuo._type) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"type\"")
	}
	return fuo.gremlinSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) gremlinSave(ctx context.Context) (*File, error) {
	res := &gremlin.Response{}
	query, bindings := fuo.gremlin(fuo.id).Query()
	if err := fuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	f := &File{config: fuo.config}
	if err := f.FromResponse(res); err != nil {
		return nil, err
	}
	return f, nil
}

func (fuo *FileUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := fuo.size; value != nil {
		v.Property(dsl.Single, file.FieldSize, *value)
	}
	if value := fuo.addsize; value != nil {
		v.Property(dsl.Single, file.FieldSize, __.Union(__.Values(file.FieldSize), __.Constant(*value)).Sum())
	}
	if value := fuo.name; value != nil {
		v.Property(dsl.Single, file.FieldName, *value)
	}
	if value := fuo.user; value != nil {
		v.Property(dsl.Single, file.FieldUser, *value)
	}
	if value := fuo.group; value != nil {
		v.Property(dsl.Single, file.FieldGroup, *value)
	}
	var properties []interface{}
	if fuo.clearuser {
		properties = append(properties, file.FieldUser)
	}
	if fuo.cleargroup {
		properties = append(properties, file.FieldGroup)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if fuo.clearedOwner {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range fuo.owner {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	if fuo.clearedType {
		tr := rv.Clone().InE(filetype.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range fuo._type {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
