// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// FileType is the model entity for the FileType schema.
type FileType struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into FileType.
func (ft *FileType) FromRows(rows *sql.Rows) error {
	var scanft struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `filetype.Columns`.
	if err := rows.Scan(
		&scanft.ID,
		&scanft.Name,
	); err != nil {
		return err
	}
	ft.ID = strconv.Itoa(scanft.ID)
	ft.Name = scanft.Name.String
	return nil
}

// QueryFiles queries the files edge of the FileType.
func (ft *FileType) QueryFiles() *FileQuery {
	return (&FileTypeClient{ft.config}).QueryFiles(ft)
}

// Update returns a builder for updating this FileType.
// Note that, you need to call FileType.Unwrap() before calling this method, if this FileType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FileType) Update() *FileTypeUpdateOne {
	return (&FileTypeClient{ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ft *FileType) Unwrap() *FileType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FileType) String() string {
	var builder strings.Builder
	builder.WriteString("FileType(")
	builder.WriteString(fmt.Sprintf("id=%v", ft.ID))
	builder.WriteString(", name=")
	builder.WriteString(ft.Name)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (ft *FileType) id() int {
	id, _ := strconv.Atoi(ft.ID)
	return id
}

// FileTypes is a parsable slice of FileType.
type FileTypes []*FileType

// FromRows scans the sql response data into FileTypes.
func (ft *FileTypes) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanft := &FileType{}
		if err := scanft.FromRows(rows); err != nil {
			return err
		}
		*ft = append(*ft, scanft)
	}
	return nil
}

func (ft FileTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
