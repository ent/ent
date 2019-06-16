// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"fbc/ent/dialect/sql/schema"
	"fbc/ent/field"
)

var (
	nullable = true
	// BoringsColumns holds the columns for the "borings" table.
	BoringsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// BoringsTable holds the schema information for the "borings" table.
	BoringsTable = &schema.Table{
		Name:        "borings",
		Columns:     BoringsColumns,
		PrimaryKey:  []*schema.Column{BoringsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BoringsTable,
	}
)

func init() {
}
