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

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromResponse scans the gremlin response data into Pet.
func (pe *Pet) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpe struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scanpe); err != nil {
		return err
	}
	pe.ID = scanpe.ID
	pe.Name = scanpe.Name
	return nil
}

// QueryTeam queries the team edge of the Pet.
func (pe *Pet) QueryTeam() *UserQuery {
	return (&PetClient{pe.config}).QueryTeam(pe)
}

// QueryOwner queries the owner edge of the Pet.
func (pe *Pet) QueryOwner() *UserQuery {
	return (&PetClient{pe.config}).QueryOwner(pe)
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
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteByte(')')
	return builder.String()
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
	var scanpe []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&scanpe); err != nil {
		return err
	}
	for _, v := range scanpe {
		*pe = append(*pe, &Pet{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return nil
}

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
