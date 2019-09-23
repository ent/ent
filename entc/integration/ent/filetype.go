// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/gremlin"
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
	var vft struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `filetype.Columns`.
	if err := rows.Scan(
		&vft.ID,
		&vft.Name,
	); err != nil {
		return err
	}
	ft.ID = strconv.Itoa(vft.ID)
	ft.Name = vft.Name.String
	return nil
}

// FromResponse scans the gremlin response data into FileType.
func (ft *FileType) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vft struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&vft); err != nil {
		return err
	}
	ft.ID = vft.ID
	ft.Name = vft.Name
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
	buf := bytes.NewBuffer(nil)
	buf.WriteString("FileType(")
	buf.WriteString(fmt.Sprintf("id=%v", ft.ID))
	buf.WriteString(fmt.Sprintf(", name=%v", ft.Name))
	buf.WriteString(")")
	return buf.String()
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
		vft := &FileType{}
		if err := vft.FromRows(rows); err != nil {
			return err
		}
		*ft = append(*ft, vft)
	}
	return nil
}

// FromResponse scans the gremlin response data into FileTypes.
func (ft *FileTypes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vft []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	if err := vmap.Decode(&vft); err != nil {
		return err
	}
	for _, v := range vft {
		*ft = append(*ft, &FileType{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return nil
}

func (ft FileTypes) config(cfg config) {
	for i := range ft {
		ft[i].config = cfg
	}
}
