// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BlobsColumns holds the columns for the "blobs" table.
	BlobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "uuid", Type: field.TypeUUID, Default: blob.DefaultUUID},
	}
	// BlobsTable holds the schema information for the "blobs" table.
	BlobsTable = &schema.Table{
		Name:        "blobs",
		Columns:     BlobsColumns,
		PrimaryKey:  []*schema.Column{BlobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:        "groups",
		Columns:     GroupsColumns,
		PrimaryKey:  []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupUsersColumns holds the columns for the "group_users" table.
	GroupUsersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupUsersTable holds the schema information for the "group_users" table.
	GroupUsersTable = &schema.Table{
		Name:       "group_users",
		Columns:    GroupUsersColumns,
		PrimaryKey: []*schema.Column{GroupUsersColumns[0], GroupUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "group_users_group_id",
				Columns: []*schema.Column{GroupUsersColumns[0]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "group_users_user_id",
				Columns: []*schema.Column{GroupUsersColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlobsTable,
		GroupsTable,
		UsersTable,
		GroupUsersTable,
	}
)

func init() {
	GroupUsersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupUsersTable.ForeignKeys[1].RefTable = UsersTable
}
