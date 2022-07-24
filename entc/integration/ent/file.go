// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/user"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// User holds the value of the "user" field.
	User *string `json:"user,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// Op holds the value of the "op" field.
	Op bool `json:"op,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges           FileEdges `json:"file_edges"`
	file_type_files *int
	group_files     *int
	user_files      *int
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Type holds the value of the type edge.
	Type *FileType `json:"type,omitempty"`
	// Field holds the value of the field edge.
	Field []*FieldType `json:"field,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	namedField  map[string][]*FieldType
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TypeOrErr returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) TypeOrErr() (*FileType, error) {
	if e.loadedTypes[1] {
		if e.Type == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: filetype.Label}
		}
		return e.Type, nil
	}
	return nil, &NotLoadedError{edge: "type"}
}

// FieldOrErr returns the Field value or an error if the edge
// was not loaded in eager-loading.
func (e FileEdges) FieldOrErr() ([]*FieldType, error) {
	if e.loadedTypes[2] {
		return e.Field, nil
	}
	return nil, &NotLoadedError{edge: "field"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldOp:
			values[i] = new(sql.NullBool)
		case file.FieldID, file.FieldSize:
			values[i] = new(sql.NullInt64)
		case file.FieldName, file.FieldUser, file.FieldGroup:
			values[i] = new(sql.NullString)
		case file.ForeignKeys[0]: // file_type_files
			values[i] = new(sql.NullInt64)
		case file.ForeignKeys[1]: // group_files
			values[i] = new(sql.NullInt64)
		case file.ForeignKeys[2]: // user_files
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type File", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = int(value.Int64)
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case file.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				f.User = new(string)
				*f.User = value.String
			}
		case file.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				f.Group = value.String
			}
		case file.FieldOp:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field op", values[i])
			} else if value.Valid {
				f.Op = value.Bool
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field file_type_files", value)
			} else if value.Valid {
				f.file_type_files = new(int)
				*f.file_type_files = int(value.Int64)
			}
		case file.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_files", value)
			} else if value.Valid {
				f.group_files = new(int)
				*f.group_files = int(value.Int64)
			}
		case file.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_files", value)
			} else if value.Valid {
				f.user_files = new(int)
				*f.user_files = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the File entity.
func (f *File) QueryOwner() *UserQuery {
	return (&FileClient{config: f.config}).QueryOwner(f)
}

// QueryType queries the "type" edge of the File entity.
func (f *File) QueryType() *FileTypeQuery {
	return (&FileClient{config: f.config}).QueryType(f)
}

// QueryField queries the "field" edge of the File entity.
func (f *File) QueryField() *FieldTypeQuery {
	return (&FileClient{config: f.config}).QueryField(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	if v := f.User; v != nil {
		builder.WriteString("user=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(f.Group)
	builder.WriteString(", ")
	builder.WriteString("op=")
	builder.WriteString(fmt.Sprintf("%v", f.Op))
	builder.WriteByte(')')
	return builder.String()
}

// NamedField returns the Field named value or an error if the edge was not
// loaded in eager-loading with this name.
func (f *File) NamedField(name string) ([]*FieldType, error) {
	if f.Edges.namedField == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := f.Edges.namedField[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (f *File) appendNamedField(name string, edges ...*FieldType) {
	if f.Edges.namedField == nil {
		f.Edges.namedField = make(map[string][]*FieldType)
	}
	if len(edges) == 0 {
		f.Edges.namedField[name] = []*FieldType{}
	} else {
		f.Edges.namedField[name] = append(f.Edges.namedField[name], edges...)
	}
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
