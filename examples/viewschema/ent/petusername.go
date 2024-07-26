// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/viewschema/ent/petusername"
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
func (pun *PetUserName) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case petusername.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pun.Name = value.String
			}
		default:
			pun.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PetUserName.
// This includes values selected through modifiers, order, etc.
func (pun *PetUserName) Value(name string) (ent.Value, error) {
	return pun.selectValues.Get(name)
}

// Unwrap unwraps the PetUserName entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pun *PetUserName) Unwrap() *PetUserName {
	_tx, ok := pun.config.driver.(*txDriver)
	if !ok {
		panic("ent: PetUserName is not a transactional entity")
	}
	pun.config.driver = _tx.drv
	return pun
}

// String implements the fmt.Stringer.
func (pun *PetUserName) String() string {
	var builder strings.Builder
	builder.WriteString("PetUserName(")
	builder.WriteString("name=")
	builder.WriteString(pun.Name)
	builder.WriteByte(')')
	return builder.String()
}

// PetUserNames is a parsable slice of PetUserName.
type PetUserNames []*PetUserName
