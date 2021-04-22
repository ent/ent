// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
	"github.com/stretchr/testify/require"
	"testing"
)

func Inspect(t *testing.T, dataSourceName string) {
	db, err := sql.Open("postgres", dataSourceName)
	require.NoError(t, err)
	defer db.Close()
	inspect, err := schema.NewInspect(db)
	require.NoError(t, err)

	tables, err := inspect.Tables(context.Background())
	require.NoError(t, err)

	var cardTable, specCardTable, userTable *schema.Table

	for _, t := range tables {
		switch t.Name {
		case "cards":
			cardTable = t
		case "spec_card":
			specCardTable = t
		case "users":
			userTable = t
		}
	}

	require.NotNil(t, cardTable)
	require.NotNil(t, specCardTable)
	require.NotNil(t, userTable)

	require.Equal(t, inspectTable{
		Name: "spec_card",
		Columns: []inspectColumn{
			{Name: "spec_id", Type: field.TypeInt64, Key: "PRI"},
			{Name: "card_id", Type: field.TypeInt64, Key: "PRI"},
		},
		ForeignKeys: []inspectForeignKey{
			{Symbol: "spec_card_card_id", Columns: []string{"card_id"}, RefTable: "cards", RefColumns: []string{"id"}},
			{Symbol: "spec_card_spec_id", Columns: []string{"spec_id"}, RefTable: "specs", RefColumns: []string{"id"}},
		},
		Indexes: []inspectIndex(nil),
	}, tableToInspectTable(specCardTable))

	require.Equal(t, inspectTable{
		Name: "cards", Columns: []inspectColumn{
			{Name: "id", Key: "PRI", Type: field.TypeInt64},
			{Name: "create_time", Type: field.TypeTime},
			{Name: "update_time", Type: field.TypeTime},
			{Name: "balance", Type: field.TypeFloat64},
			{Name: "number", Type: field.TypeString},
			{Name: "name", Type: field.TypeString, Nullable: true},
			{Name: "user_card", Type: field.TypeInt64, Key: "UNI", Nullable: true},
		},
		ForeignKeys: []inspectForeignKey{
			{Symbol: "cards_users_card", Columns: []string{"user_card"}, RefTable: "users", RefColumns: []string{"id"}},
		},
		Indexes: []inspectIndex{
			{Name: "card_id", Columns: []string{"id"}},
			{Name: "card_id_name_number", Columns: []string{"id", "name", "number"}},
			{Name: "card_number", Columns: []string{"number"}},
			{Name: "user_card_key", Unique: true, Columns: []string{"user_card"}},
		},
	}, tableToInspectTable(cardTable))

	require.Equal(t, inspectTable{
		Name: "users", Columns: []inspectColumn{
			{Name: "id", Type: field.TypeInt64, Key: "PRI"},
			{Name: "optional_int", Type: field.TypeInt64, Nullable: true},
			{Name: "age", Type: field.TypeInt64},
			{Name: "name", Type: field.TypeString},
			{Name: "last", Type: field.TypeString},
			{Name: "nickname", Type: field.TypeString, Key: "UNI", Nullable: true},
			{Name: "address", Type: field.TypeString, Nullable: true},
			{Name: "phone", Type: field.TypeString, Key: "UNI", Nullable: true},
			{Name: "password", Type: field.TypeString, Nullable: true},
			{Name: "role", Type: field.TypeString},
			{Name: "sso_cert", Type: field.TypeString, Nullable: true},
			{Name: "group_blocked", Type: field.TypeInt64, Nullable: true},
			{Name: "user_spouse", Type: field.TypeInt64, Key: "UNI", Nullable: true},
			{Name: "user_parent", Type: field.TypeInt64, Nullable: true},
		},
		ForeignKeys: []inspectForeignKey{
			{Symbol: "users_groups_blocked", Columns: []string{"group_blocked"}, RefTable: "groups", RefColumns: []string{"id"}},
			{Symbol: "users_users_parent", Columns: []string{"user_parent"}, RefTable: "users", RefColumns: []string{"id"}},
			{Symbol: "users_users_spouse", Columns: []string{"user_spouse"}, RefTable: "users", RefColumns: []string{"id"}},
		}, Indexes: []inspectIndex{
			{Name: "nickname_key", Unique: true, Columns: []string{"nickname"}},
			{Name: "phone_key", Unique: true, Columns: []string{"phone"}},
			{Name: "user_spouse_key", Unique: true, Columns: []string{"user_spouse"}},
		},
	}, tableToInspectTable(userTable))
}

func tableToInspectTable(table *schema.Table) inspectTable {
	var (
		columns     []inspectColumn
		foreignKeys []inspectForeignKey
		indexes     []inspectIndex
	)
	for _, c := range table.Columns {
		columns = append(columns, inspectColumn{
			Name:      c.Name,
			Key:       c.Key,
			Increment: c.Increment,
			Nullable:  c.Nullable,
			Enums:     c.Enums,
			Type:      c.Type,
		})
	}
	for _, index := range table.Indexes {
		var indexColumns []string

		for _, c := range index.Columns {
			indexColumns = append(indexColumns, c.Name)
		}

		indexes = append(indexes, inspectIndex{
			Name:    index.Name,
			Unique:  index.Unique,
			Columns: indexColumns,
		})
	}
	for _, fk := range table.ForeignKeys {
		var fkColumns, fkRefColumns []string

		for _, c := range fk.Columns {
			fkColumns = append(fkColumns, c.Name)
		}

		for _, c := range fk.RefColumns {
			fkRefColumns = append(fkRefColumns, c.Name)
		}

		foreignKeys = append(foreignKeys, inspectForeignKey{
			Symbol:     fk.Symbol,
			RefTable:   fk.RefTable.Name,
			Columns:    fkColumns,
			RefColumns: fkRefColumns,
		})
	}
	return inspectTable{
		Name:        table.Name,
		Columns:     columns,
		ForeignKeys: foreignKeys,
		Indexes:     indexes,
	}
}

type inspectTable struct {
	Name        string
	Columns     []inspectColumn
	ForeignKeys []inspectForeignKey
	Indexes     []inspectIndex
}

type inspectColumn struct {
	Name      string
	Key       string
	Increment bool
	Unique    bool
	Nullable  bool
	Enums     []string
	Type      field.Type
}

type inspectForeignKey struct {
	Symbol     string
	Columns    []string
	RefTable   string
	RefColumns []string
}

type inspectIndex struct {
	Name    string
	Unique  bool
	Columns []string
}
