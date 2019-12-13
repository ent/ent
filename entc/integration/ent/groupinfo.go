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

// GroupInfo is the model entity for the GroupInfo schema.
type GroupInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
}

// FromRows scans the sql response data into GroupInfo.
func (gi *GroupInfo) FromRows(rows *sql.Rows) error {
	var scangi struct {
		ID       int
		Desc     sql.NullString
		MaxUsers sql.NullInt64
	}
	// the order here should be the same as in the `groupinfo.Columns`.
	if err := rows.Scan(
		&scangi.ID,
		&scangi.Desc,
		&scangi.MaxUsers,
	); err != nil {
		return err
	}
	gi.ID = strconv.Itoa(scangi.ID)
	gi.Desc = scangi.Desc.String
	gi.MaxUsers = int(scangi.MaxUsers.Int64)
	return nil
}

// QueryGroups queries the groups edge of the GroupInfo.
func (gi *GroupInfo) QueryGroups() *GroupQuery {
	return (&GroupInfoClient{gi.config}).QueryGroups(gi)
}

// Update returns a builder for updating this GroupInfo.
// Note that, you need to call GroupInfo.Unwrap() before calling this method, if this GroupInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (gi *GroupInfo) Update() *GroupInfoUpdateOne {
	return (&GroupInfoClient{gi.config}).UpdateOne(gi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gi *GroupInfo) Unwrap() *GroupInfo {
	tx, ok := gi.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupInfo is not a transactional entity")
	}
	gi.config.driver = tx.drv
	return gi
}

// String implements the fmt.Stringer.
func (gi *GroupInfo) String() string {
	var builder strings.Builder
	builder.WriteString("GroupInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", gi.ID))
	builder.WriteString(", desc=")
	builder.WriteString(gi.Desc)
	builder.WriteString(", max_users=")
	builder.WriteString(fmt.Sprintf("%v", gi.MaxUsers))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (gi *GroupInfo) id() int {
	id, _ := strconv.Atoi(gi.ID)
	return id
}

// GroupInfos is a parsable slice of GroupInfo.
type GroupInfos []*GroupInfo

// FromRows scans the sql response data into GroupInfos.
func (gi *GroupInfos) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scangi := &GroupInfo{}
		if err := scangi.FromRows(rows); err != nil {
			return err
		}
		*gi = append(*gi, scangi)
	}
	return nil
}

func (gi GroupInfos) config(cfg config) {
	for _i := range gi {
		gi[_i].config = cfg
	}
}
