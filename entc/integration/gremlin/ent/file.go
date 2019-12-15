// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/gremlin"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// User holds the value of the "user" field.
	User *string `json:"user,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
}

// FromResponse scans the gremlin response data into File.
func (f *File) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanf struct {
		ID    string  `json:"id,omitempty"`
		Size  int     `json:"fsize,omitempty"`
		Name  string  `json:"name,omitempty"`
		User  *string `json:"user,omitempty"`
		Group string  `json:"group,omitempty"`
	}
	if err := vmap.Decode(&scanf); err != nil {
		return err
	}
	f.ID = scanf.ID
	f.Size = scanf.Size
	f.Name = scanf.Name
	f.User = scanf.User
	f.Group = scanf.Group
	return nil
}

// QueryOwner queries the owner edge of the File.
func (f *File) QueryOwner() *UserQuery {
	return (&FileClient{f.config}).QueryOwner(f)
}

// QueryType queries the type edge of the File.
func (f *File) QueryType() *FileTypeQuery {
	return (&FileClient{f.config}).QueryType(f)
}

// Update returns a builder for updating this File.
// Note that, you need to call File.Unwrap() before calling this method, if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", name=")
	builder.WriteString(f.Name)
	if v := f.User; v != nil {
		builder.WriteString(", user=")
		builder.WriteString(*v)
	}
	builder.WriteString(", group=")
	builder.WriteString(f.Group)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (f *File) id() int {
	id, _ := strconv.Atoi(f.ID)
	return id
}

// Files is a parsable slice of File.
type Files []*File

// FromResponse scans the gremlin response data into Files.
func (f *Files) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanf []struct {
		ID    string  `json:"id,omitempty"`
		Size  int     `json:"fsize,omitempty"`
		Name  string  `json:"name,omitempty"`
		User  *string `json:"user,omitempty"`
		Group string  `json:"group,omitempty"`
	}
	if err := vmap.Decode(&scanf); err != nil {
		return err
	}
	for _, v := range scanf {
		*f = append(*f, &File{
			ID:    v.ID,
			Size:  v.Size,
			Name:  v.Name,
			User:  v.User,
			Group: v.Group,
		})
	}
	return nil
}

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
