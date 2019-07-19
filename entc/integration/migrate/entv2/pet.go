// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/sql"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Pet.
func (pe *Pet) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vpe struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&vpe); err != nil {
		return err
	}
	pe.ID = vpe.ID
	return nil
}

// FromRows scans the sql response data into Pet.
func (pe *Pet) FromRows(rows *sql.Rows) error {
	var vpe struct {
		ID int
	}
	// the order here should be the same as in the `pet.Columns`.
	if err := rows.Scan(
		&vpe.ID,
	); err != nil {
		return err
	}
	pe.ID = strconv.Itoa(vpe.ID)
	return nil
}

// Update returns a builder for updating this Pet.
// Note that, you need to call Pet.Unwrap() before calling this method, if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Pet(")
	buf.WriteString(fmt.Sprintf("id=%v", pe.ID))
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (pe *Pet) id() int {
	id, _ := strconv.Atoi(pe.ID)
	return id
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

// FromResponse scans the gremlin response data into Pets.
func (pe *Pets) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vpe []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&vpe); err != nil {
		return err
	}
	for _, v := range vpe {
		*pe = append(*pe, &Pet{
			ID: v.ID,
		})
	}
	return nil
}

// FromRows scans the sql response data into Pets.
func (pe *Pets) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vpe := &Pet{}
		if err := vpe.FromRows(rows); err != nil {
			return err
		}
		*pe = append(*pe, vpe)
	}
	return nil
}

func (pe Pets) config(cfg config) {
	for i := range pe {
		pe[i].config = cfg
	}
}
