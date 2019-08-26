// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Group is the model entity for the Group schema.
type Group struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into Group.
func (gr *Group) FromRows(rows *sql.Rows) error {
	var vgr struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `group.Columns`.
	if err := rows.Scan(
		&vgr.ID,
		&vgr.Name,
	); err != nil {
		return err
	}
	gr.ID = vgr.ID
	gr.Name = vgr.Name.String
	return nil
}

// QueryUsers queries the users edge of the Group.
func (gr *Group) QueryUsers() *UserQuery {
	return (&GroupClient{gr.config}).QueryUsers(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Group(")
	buf.WriteString(fmt.Sprintf("id=%v", gr.ID))
	buf.WriteString(fmt.Sprintf(", name=%v", gr.Name))
	buf.WriteString(")")
	return buf.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

// FromRows scans the sql response data into Groups.
func (gr *Groups) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vgr := &Group{}
		if err := vgr.FromRows(rows); err != nil {
			return err
		}
		*gr = append(*gr, vgr)
	}
	return nil
}

func (gr Groups) config(cfg config) {
	for i := range gr {
		gr[i].config = cfg
	}
}
