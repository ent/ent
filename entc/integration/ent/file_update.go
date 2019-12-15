// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
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
	return fu.sqlSave(ctx)
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

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(fu.driver.Dialect())
		selector = builder.Select(file.FieldID).From(builder.Table(file.Table))
	)
	for _, p := range fu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = fu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := fu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(file.Table)
	)
	updater = updater.Where(sql.InInts(file.FieldID, ids...))
	if value := fu.size; value != nil {
		updater.Set(file.FieldSize, *value)
	}
	if value := fu.addsize; value != nil {
		updater.Add(file.FieldSize, *value)
	}
	if value := fu.name; value != nil {
		updater.Set(file.FieldName, *value)
	}
	if value := fu.user; value != nil {
		updater.Set(file.FieldUser, *value)
	}
	if fu.clearuser {
		updater.SetNull(file.FieldUser)
	}
	if value := fu.group; value != nil {
		updater.Set(file.FieldGroup, *value)
	}
	if fu.cleargroup {
		updater.SetNull(file.FieldGroup)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if fu.clearedOwner {
		query, args := builder.Update(file.OwnerTable).
			SetNull(file.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(fu.owner) > 0 {
		for eid := range fu.owner {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(file.OwnerTable).
				Set(file.OwnerColumn, eid).
				Where(sql.InInts(file.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if fu.clearedType {
		query, args := builder.Update(file.TypeTable).
			SetNull(file.TypeColumn).
			Where(sql.InInts(filetype.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(fu._type) > 0 {
		for eid := range fu._type {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(file.TypeTable).
				Set(file.TypeColumn, eid).
				Where(sql.InInts(file.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
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
	return fuo.sqlSave(ctx)
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

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (f *File, err error) {
	var (
		builder  = sql.Dialect(fuo.driver.Dialect())
		selector = builder.Select(file.Columns...).From(builder.Table(file.Table))
	)
	file.ID(fuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = fuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		f = &File{config: fuo.config}
		if err := f.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into File: %v", err)
		}
		id = f.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("File with id: %v", fuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one File with the same id: %v", fuo.id)
	}

	tx, err := fuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(file.Table)
	)
	updater = updater.Where(sql.InInts(file.FieldID, ids...))
	if value := fuo.size; value != nil {
		updater.Set(file.FieldSize, *value)
		f.Size = *value
	}
	if value := fuo.addsize; value != nil {
		updater.Add(file.FieldSize, *value)
		f.Size += *value
	}
	if value := fuo.name; value != nil {
		updater.Set(file.FieldName, *value)
		f.Name = *value
	}
	if value := fuo.user; value != nil {
		updater.Set(file.FieldUser, *value)
		f.User = value
	}
	if fuo.clearuser {
		f.User = nil
		updater.SetNull(file.FieldUser)
	}
	if value := fuo.group; value != nil {
		updater.Set(file.FieldGroup, *value)
		f.Group = *value
	}
	if fuo.cleargroup {
		var value string
		f.Group = value
		updater.SetNull(file.FieldGroup)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if fuo.clearedOwner {
		query, args := builder.Update(file.OwnerTable).
			SetNull(file.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(fuo.owner) > 0 {
		for eid := range fuo.owner {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(file.OwnerTable).
				Set(file.OwnerColumn, eid).
				Where(sql.InInts(file.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if fuo.clearedType {
		query, args := builder.Update(file.TypeTable).
			SetNull(file.TypeColumn).
			Where(sql.InInts(filetype.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(fuo._type) > 0 {
		for eid := range fuo._type {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(file.TypeTable).
				Set(file.TypeColumn, eid).
				Where(sql.InInts(file.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return f, nil
}
