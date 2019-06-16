// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
)

// Boring is the model entity for the Boring schema.
type Boring struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Boring.
func (b *Boring) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vb struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&vb); err != nil {
		return err
	}
	b.ID = vb.ID
	return nil
}

// FromRows scans the sql response data into Boring.
func (b *Boring) FromRows(rows *sql.Rows) error {
	var vb struct {
		ID int
	}
	// the order here should be the same as in the `boring.Columns`.
	if err := rows.Scan(
		&vb.ID,
	); err != nil {
		return err
	}
	b.ID = strconv.Itoa(vb.ID)
	return nil
}

// Update returns a builder for updating this Boring.
// Note that, you need to call Boring.Unwrap() before calling this method, if this Boring
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Boring) Update() *BoringUpdateOne {
	return (&BoringClient{b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Boring) Unwrap() *Boring {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Boring is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Boring) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Boring(")
	buf.WriteString(fmt.Sprintf("id=%v,", b.ID))
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (b *Boring) id() int {
	id, _ := strconv.Atoi(b.ID)
	return id
}

// Borings is a parsable slice of Boring.
type Borings []*Boring

// FromResponse scans the gremlin response data into Borings.
func (b *Borings) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vb []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&vb); err != nil {
		return err
	}
	for _, v := range vb {
		*b = append(*b, &Boring{
			ID: v.ID,
		})
	}
	return nil
}

// FromRows scans the sql response data into Borings.
func (b *Borings) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vb := &Boring{}
		if err := vb.FromRows(rows); err != nil {
			return err
		}
		*b = append(*b, vb)
	}
	return nil
}

func (b Borings) config(cfg config) {
	for i := range b {
		b[i].config = cfg
	}
}
