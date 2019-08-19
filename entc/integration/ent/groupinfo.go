// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/sql"
)

// GroupInfo is the model entity for the GroupInfo schema.
type GroupInfo struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
}

// FromRows scans the sql response data into GroupInfo.
func (gi *GroupInfo) FromRows(rows *sql.Rows) error {
	var vgi struct {
		ID       int
		Desc     sql.NullString
		MaxUsers sql.NullInt64
	}
	// the order here should be the same as in the `groupinfo.Columns`.
	if err := rows.Scan(
		&vgi.ID,
		&vgi.Desc,
		&vgi.MaxUsers,
	); err != nil {
		return err
	}
	gi.ID = strconv.Itoa(vgi.ID)
	gi.Desc = vgi.Desc.String
	gi.MaxUsers = int(vgi.MaxUsers.Int64)
	return nil
}

// FromResponse scans the gremlin response data into GroupInfo.
func (gi *GroupInfo) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vgi struct {
		ID       string `json:"id,omitempty"`
		Desc     string `json:"desc,omitempty"`
		MaxUsers int    `json:"max_users,omitempty"`
	}
	if err := vmap.Decode(&vgi); err != nil {
		return err
	}
	gi.ID = vgi.ID
	gi.Desc = vgi.Desc
	gi.MaxUsers = vgi.MaxUsers
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
	buf := bytes.NewBuffer(nil)
	buf.WriteString("GroupInfo(")
	buf.WriteString(fmt.Sprintf("id=%v", gi.ID))
	buf.WriteString(fmt.Sprintf(", desc=%v", gi.Desc))
	buf.WriteString(fmt.Sprintf(", max_users=%v", gi.MaxUsers))
	buf.WriteString(")")
	return buf.String()
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
		vgi := &GroupInfo{}
		if err := vgi.FromRows(rows); err != nil {
			return err
		}
		*gi = append(*gi, vgi)
	}
	return nil
}

// FromResponse scans the gremlin response data into GroupInfos.
func (gi *GroupInfos) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vgi []struct {
		ID       string `json:"id,omitempty"`
		Desc     string `json:"desc,omitempty"`
		MaxUsers int    `json:"max_users,omitempty"`
	}
	if err := vmap.Decode(&vgi); err != nil {
		return err
	}
	for _, v := range vgi {
		*gi = append(*gi, &GroupInfo{
			ID:       v.ID,
			Desc:     v.Desc,
			MaxUsers: v.MaxUsers,
		})
	}
	return nil
}

func (gi GroupInfos) config(cfg config) {
	for i := range gi {
		gi[i].config = cfg
	}
}
