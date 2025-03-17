// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
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
	// SetID holds the value of the "set_id" field.
	SetID int `json:"set_id,omitempty"`
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
	// FieldID holds the value of the "field_id" field.
	FieldID int `json:"field_id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges           FileEdges `json:"file_edges"`
	file_type_files *int
	group_files     *int
	user_files      *int
	selectValues    sql.SelectValues
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
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// TypeOrErr returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) TypeOrErr() (*FileType, error) {
	if e.Type != nil {
		return e.Type, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: filetype.Label}
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
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldOp:
			values[i] = new(sql.NullBool)
		case file.FieldID, file.FieldSetID, file.FieldSize, file.FieldFieldID:
			values[i] = new(sql.NullInt64)
		case file.FieldName, file.FieldUser, file.FieldGroup:
			values[i] = new(sql.NullString)
		case file.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case file.ForeignKeys[0]: // file_type_files
			values[i] = new(sql.NullInt64)
		case file.ForeignKeys[1]: // group_files
			values[i] = new(sql.NullInt64)
		case file.ForeignKeys[2]: // user_files
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (_m *File) assignValues(columns []string, values []any) error {
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
			_m.ID = int(value.Int64)
		case file.FieldSetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field set_id", values[i])
			} else if value.Valid {
				_m.SetID = int(value.Int64)
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				_m.Size = int(value.Int64)
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_m.Name = value.String
			}
		case file.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				_m.User = new(string)
				*_m.User = value.String
			}
		case file.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				_m.Group = value.String
			}
		case file.FieldOp:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field op", values[i])
			} else if value.Valid {
				_m.Op = value.Bool
			}
		case file.FieldFieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field field_id", values[i])
			} else if value.Valid {
				_m.FieldID = int(value.Int64)
			}
		case file.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				_m.CreateTime = value.Time
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field file_type_files", value)
			} else if value.Valid {
				_m.file_type_files = new(int)
				*_m.file_type_files = int(value.Int64)
			}
		case file.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_files", value)
			} else if value.Valid {
				_m.group_files = new(int)
				*_m.group_files = int(value.Int64)
			}
		case file.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_files", value)
			} else if value.Valid {
				_m.user_files = new(int)
				*_m.user_files = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (_m *File) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the File entity.
func (_m *File) QueryOwner() *UserQuery {
	return NewFileClient(_m.config).QueryOwner(_m)
}

// QueryType queries the "type" edge of the File entity.
func (_m *File) QueryType() *FileTypeQuery {
	return NewFileClient(_m.config).QueryType(_m)
}

// QueryField queries the "field" edge of the File entity.
func (_m *File) QueryField() *FieldTypeQuery {
	return NewFileClient(_m.config).QueryField(_m)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *File) Update() *FileUpdateOne {
	return NewFileClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *File) Unwrap() *File {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("set_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.SetID))
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", _m.Size))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteString(", ")
	if v := _m.User; v != nil {
		builder.WriteString("user=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(_m.Group)
	builder.WriteString(", ")
	builder.WriteString("op=")
	builder.WriteString(fmt.Sprintf("%v", _m.Op))
	builder.WriteString(", ")
	builder.WriteString("field_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.FieldID))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(_m.CreateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NamedField returns the Field named value or an error if the edge was not
// loaded in eager-loading with this name.
func (_m *File) NamedField(name string) ([]*FieldType, error) {
	if _m.Edges.namedField == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := _m.Edges.namedField[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (_m *File) appendNamedField(name string, edges ...*FieldType) {
	if _m.Edges.namedField == nil {
		_m.Edges.namedField = make(map[string][]*FieldType)
	}
	if len(edges) == 0 {
		_m.Edges.namedField[name] = []*FieldType{}
	} else {
		_m.Edges.namedField[name] = append(_m.Edges.namedField[name], edges...)
	}
}

// Files is a parsable slice of File.
type Files []*File
