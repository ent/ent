// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/viewcomposite/ent/petusername"
)

// PetUserName is the model entity for the PetUserName schema.
type PetUserName struct {
	config `json:"-"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PetUserName) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case petusername.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PetUserName fields.
func (m *PetUserName) assignValues(columns []string, values []any) error {
	if v, c := len(values), len(columns); v < c {
		return fmt.Errorf("mismatch number of scan values: %d != %d", v, c)
	}
	for i := range columns {
		switch columns[i] {
		case petusername.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PetUserName.
// This includes values selected through modifiers, order, etc.
func (m *PetUserName) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Unwrap unwraps the PetUserName entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *PetUserName) Unwrap() *PetUserName {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: PetUserName is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *PetUserName) String() string {
	var builder strings.Builder
	builder.WriteString("PetUserName(")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// PetUserNames is a parsable slice of PetUserName.
type PetUserNames []*PetUserName
