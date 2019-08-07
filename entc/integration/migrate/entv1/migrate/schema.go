// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"fbc/ent/dialect/sql/schema"
	"fbc/ent/field"
)

var (
	nullable = true
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString, Size: 10},
		{Name: "address", Type: field.TypeString, Nullable: &nullable},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
