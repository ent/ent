// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarColumns holds the columns for the "Car" table.
	CarColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_car", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CarTable holds the schema information for the "Car" table.
	CarTable = &schema.Table{
		Name:       "Car",
		Columns:    CarColumns,
		PrimaryKey: []*schema.Column{CarColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Car_users_car",
				Columns:    []*schema.Column{CarColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ConversionsColumns holds the columns for the "conversions" table.
	ConversionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "int8_to_string", Type: field.TypeInt8, Nullable: true},
		{Name: "uint8_to_string", Type: field.TypeUint8, Nullable: true},
		{Name: "int16_to_string", Type: field.TypeInt16, Nullable: true},
		{Name: "uint16_to_string", Type: field.TypeUint16, Nullable: true},
		{Name: "int32_to_string", Type: field.TypeInt32, Nullable: true},
		{Name: "uint32_to_string", Type: field.TypeUint32, Nullable: true},
		{Name: "int64_to_string", Type: field.TypeInt64, Nullable: true},
		{Name: "uint64_to_string", Type: field.TypeUint64, Nullable: true},
	}
	// ConversionsTable holds the schema information for the "conversions" table.
	ConversionsTable = &schema.Table{
		Name:       "conversions",
		Columns:    ConversionsColumns,
		PrimaryKey: []*schema.Column{ConversionsColumns[0]},
	}
	// CustomTypesColumns holds the columns for the "custom_types" table.
	CustomTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "custom", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "customtype"}},
	}
	// CustomTypesTable holds the schema information for the "custom_types" table.
	CustomTypesTable = &schema.Table{
		Name:       "custom_types",
		Columns:    CustomTypesColumns,
		PrimaryKey: []*schema.Column{CustomTypesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt32},
		{Name: "name", Type: field.TypeString, Size: 10},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "renamed", Type: field.TypeString, Nullable: true},
		{Name: "old_token", Type: field.TypeString},
		{Name: "blob", Type: field.TypeBytes, Nullable: true, Size: 255},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"logged_in", "logged_out"}, Default: "logged_in"},
		{Name: "status", Type: field.TypeString, Nullable: true},
		{Name: "workplace", Type: field.TypeString, Nullable: true, Size: 30},
		{Name: "drop_optional", Type: field.TypeString, Nullable: true},
		{Name: "user_children", Type: field.TypeInt, Nullable: true},
		{Name: "user_spouse", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_children",
				Columns:    []*schema.Column{UsersColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_spouse",
				Columns:    []*schema.Column{UsersColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_description",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
				Annotation: &entsql.IndexAnnotation{
					Prefix: 50,
				},
			},
			{
				Name:    "user_name_address",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2], UsersColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarTable,
		ConversionsTable,
		CustomTypesTable,
		UsersTable,
	}
)

func init() {
	CarTable.ForeignKeys[0].RefTable = UsersTable
	CarTable.Annotation = &entsql.Annotation{
		Table: "Car",
	}
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[1].RefTable = UsersTable
}
