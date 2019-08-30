// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Street is the model entity for the Street schema.
type Street struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into Street.
func (s *Street) FromRows(rows *sql.Rows) error {
	var vs struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `street.Columns`.
	if err := rows.Scan(
		&vs.ID,
		&vs.Name,
	); err != nil {
		return err
	}
	s.ID = vs.ID
	s.Name = vs.Name.String
	return nil
}

// QueryCity queries the city edge of the Street.
func (s *Street) QueryCity() *CityQuery {
	return (&StreetClient{s.config}).QueryCity(s)
}

// Update returns a builder for updating this Street.
// Note that, you need to call Street.Unwrap() before calling this method, if this Street
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Street) Update() *StreetUpdateOne {
	return (&StreetClient{s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Street) Unwrap() *Street {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Street is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Street) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Street(")
	buf.WriteString(fmt.Sprintf("id=%v", s.ID))
	buf.WriteString(fmt.Sprintf(", name=%v", s.Name))
	buf.WriteString(")")
	return buf.String()
}

// Streets is a parsable slice of Street.
type Streets []*Street

// FromRows scans the sql response data into Streets.
func (s *Streets) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vs := &Street{}
		if err := vs.FromRows(rows); err != nil {
			return err
		}
		*s = append(*s, vs)
	}
	return nil
}

func (s Streets) config(cfg config) {
	for i := range s {
		s[i].config = cfg
	}
}
