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
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/file"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	size  *int
	name  *string
	user  *string
	group *string
	owner map[string]struct{}
	_type map[string]struct{}
}

// SetSize sets the size field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.size = &i
	return fc
}

// SetNillableSize sets the size field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetName sets the name field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.name = &s
	return fc
}

// SetUser sets the user field.
func (fc *FileCreate) SetUser(s string) *FileCreate {
	fc.user = &s
	return fc
}

// SetNillableUser sets the user field if the given value is not nil.
func (fc *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		fc.SetUser(*s)
	}
	return fc
}

// SetGroup sets the group field.
func (fc *FileCreate) SetGroup(s string) *FileCreate {
	fc.group = &s
	return fc
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fc *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		fc.SetGroup(*s)
	}
	return fc
}

// SetOwnerID sets the owner edge to User by id.
func (fc *FileCreate) SetOwnerID(id string) *FileCreate {
	if fc.owner == nil {
		fc.owner = make(map[string]struct{})
	}
	fc.owner[id] = struct{}{}
	return fc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the owner edge to User.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetTypeID sets the type edge to FileType by id.
func (fc *FileCreate) SetTypeID(id string) *FileCreate {
	if fc._type == nil {
		fc._type = make(map[string]struct{})
	}
	fc._type[id] = struct{}{}
	return fc
}

// SetNillableTypeID sets the type edge to FileType by id if the given value is not nil.
func (fc *FileCreate) SetNillableTypeID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetTypeID(*id)
	}
	return fc
}

// SetType sets the type edge to FileType.
func (fc *FileCreate) SetType(f *FileType) *FileCreate {
	return fc.SetTypeID(f.ID)
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if fc.size == nil {
		v := file.DefaultSize
		fc.size = &v
	}
	if err := file.SizeValidator(*fc.size); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
	}
	if fc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(fc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if len(fc._type) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"type\"")
	}
	return fc.gremlinSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FileCreate) gremlinSave(ctx context.Context) (*File, error) {
	res := &gremlin.Response{}
	query, bindings := fc.gremlin().Query()
	if err := fc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	f := &File{config: fc.config}
	if err := f.FromResponse(res); err != nil {
		return nil, err
	}
	return f, nil
}

func (fc *FileCreate) gremlin() *dsl.Traversal {
	v := g.AddV(file.Label)
	if fc.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fc.size)
	}
	if fc.name != nil {
		v.Property(dsl.Single, file.FieldName, *fc.name)
	}
	if fc.user != nil {
		v.Property(dsl.Single, file.FieldUser, *fc.user)
	}
	if fc.group != nil {
		v.Property(dsl.Single, file.FieldGroup, *fc.group)
	}
	for id := range fc.owner {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	for id := range fc._type {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	return v.ValueMap(true)
}
