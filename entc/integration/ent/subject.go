// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/subject"
	"github.com/google/uuid"
)

// Subject is the model entity for the Subject schema.
type Subject struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectQuery when eager-loading is set.
	Edges        SubjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SubjectEdges holds the relations/edges for other nodes in the graph.
type SubjectEdges struct {
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// SubjectStudents holds the value of the subject_students edge.
	SubjectStudents []*SubjectStudent `json:"subject_students,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes          [2]bool
	namedStudents        map[string][]*Student
	namedSubjectStudents map[string][]*SubjectStudent
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[0] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// SubjectStudentsOrErr returns the SubjectStudents value or an error if the edge
// was not loaded in eager-loading.
func (e SubjectEdges) SubjectStudentsOrErr() ([]*SubjectStudent, error) {
	if e.loadedTypes[1] {
		return e.SubjectStudents, nil
	}
	return nil, &NotLoadedError{edge: "subject_students"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subject.FieldName:
			values[i] = new(sql.NullString)
		case subject.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subject fields.
func (s *Subject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subject.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subject.
// This includes values selected through modifiers, order, etc.
func (s *Subject) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryStudents queries the "students" edge of the Subject entity.
func (s *Subject) QueryStudents() *StudentQuery {
	return NewSubjectClient(s.config).QueryStudents(s)
}

// QuerySubjectStudents queries the "subject_students" edge of the Subject entity.
func (s *Subject) QuerySubjectStudents() *SubjectStudentQuery {
	return NewSubjectClient(s.config).QuerySubjectStudents(s)
}

// Update returns a builder for updating this Subject.
// Note that you need to call Subject.Unwrap() before calling this method if this Subject
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subject) Update() *SubjectUpdateOne {
	return NewSubjectClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subject) Unwrap() *Subject {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subject is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subject) String() string {
	var builder strings.Builder
	builder.WriteString("Subject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedStudents returns the Students named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Subject) NamedStudents(name string) ([]*Student, error) {
	if s.Edges.namedStudents == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedStudents[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Subject) appendNamedStudents(name string, edges ...*Student) {
	if s.Edges.namedStudents == nil {
		s.Edges.namedStudents = make(map[string][]*Student)
	}
	if len(edges) == 0 {
		s.Edges.namedStudents[name] = []*Student{}
	} else {
		s.Edges.namedStudents[name] = append(s.Edges.namedStudents[name], edges...)
	}
}

// NamedSubjectStudents returns the SubjectStudents named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *Subject) NamedSubjectStudents(name string) ([]*SubjectStudent, error) {
	if s.Edges.namedSubjectStudents == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedSubjectStudents[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *Subject) appendNamedSubjectStudents(name string, edges ...*SubjectStudent) {
	if s.Edges.namedSubjectStudents == nil {
		s.Edges.namedSubjectStudents = make(map[string][]*SubjectStudent)
	}
	if len(edges) == 0 {
		s.Edges.namedSubjectStudents[name] = []*SubjectStudent{}
	} else {
		s.Edges.namedSubjectStudents[name] = append(s.Edges.namedSubjectStudents[name], edges...)
	}
}

// Subjects is a parsable slice of Subject.
type Subjects []*Subject
