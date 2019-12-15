// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
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

// FromRows scans the sql response data into File.
func (f *File) FromRows(rows *sql.Rows) error {
	var scanf struct {
		ID    int
		Size  sql.NullInt64
		Name  sql.NullString
		User  sql.NullString
		Group sql.NullString
	}
	// the order here should be the same as in the `file.Columns`.
	if err := rows.Scan(
		&scanf.ID,
		&scanf.Size,
		&scanf.Name,
		&scanf.User,
		&scanf.Group,
	); err != nil {
		return err
	}
	f.ID = strconv.Itoa(scanf.ID)
	f.Size = int(scanf.Size.Int64)
	f.Name = scanf.Name.String
	if scanf.User.Valid {
		f.User = new(string)
		*f.User = scanf.User.String
	}
	f.Group = scanf.Group.String
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

// FromRows scans the sql response data into Files.
func (f *Files) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanf := &File{}
		if err := scanf.FromRows(rows); err != nil {
			return err
		}
		*f = append(*f, scanf)
	}
	return nil
}

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
