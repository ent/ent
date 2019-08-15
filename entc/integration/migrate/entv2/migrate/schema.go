// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"fbc/ent/dialect/sql/schema"
	"fbc/ent/entc/integration/migrate/entv2/user"
	"fbc/ent/schema/field"
)

var (
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
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:        "pets",
		Columns:     PetsColumns,
		PrimaryKey:  []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "phone", Type: field.TypeString},
		{Name: "buffer", Type: field.TypeBytes, Default: user.DefaultBuffer},
		{Name: "title", Type: field.TypeString, Default: user.DefaultTitle},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "phone_age",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[3], UsersColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		PetsTable,
		UsersTable,
	}
)

func init() {
}
