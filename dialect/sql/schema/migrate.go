// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

const (
	// TypeTable defines the table name holding the type information.
	TypeTable = "ent_types"

	// MaxTypes defines the max number of types can be created when
	// defining universal ids. The left 16-bits are reserved.
	MaxTypes = math.MaxUint16
)

// NewTypesTable returns a new table for holding the global-id information.
func NewTypesTable() *Table {
	return NewTable(TypeTable).
		AddPrimary(&Column{Name: "id", Type: field.TypeUint, Increment: true}).
		AddColumn(&Column{Name: "type", Type: field.TypeString, Unique: true})
}

// MigrateOption allows configuring Atlas using functional arguments.
type MigrateOption func(*Atlas)

// WithGlobalUniqueID sets the universal ids options to the migration.
// Defaults to false.
func WithGlobalUniqueID(b bool) MigrateOption {
	return func(a *Atlas) {
		a.universalID = b
	}
}

// WithIndent sets Atlas to generate SQL statements with indentation.
// An empty string indicates no indentation.
func WithIndent(indent string) MigrateOption {
	return func(a *Atlas) {
		a.indent = indent
	}
}

// WithErrNoPlan sets Atlas to returns a migrate.ErrNoPlan in case
// the migration plan is empty. Defaults to false.
func WithErrNoPlan(b bool) MigrateOption {
	return func(a *Atlas) {
		a.errNoPlan = b
	}
}

// WithSchemaName sets the database schema for the migration.
// If not set, the CURRENT_SCHEMA() is used.
func WithSchemaName(ns string) MigrateOption {
	return func(a *Atlas) {
		a.schema = ns
	}
}

// WithDropColumn sets the columns dropping option to the migration.
// Defaults to false.
func WithDropColumn(b bool) MigrateOption {
	return func(a *Atlas) {
		a.dropColumns = b
	}
}

// WithDropIndex sets the indexes dropping option to the migration.
// Defaults to false.
func WithDropIndex(b bool) MigrateOption {
	return func(a *Atlas) {
		a.dropIndexes = b
	}
}

// WithForeignKeys enables creating foreign-key in ddl. Defaults to true.
func WithForeignKeys(b bool) MigrateOption {
	return func(a *Atlas) {
		a.withForeignKeys = b
	}
}

// WithHooks adds a list of hooks to the schema migration.
func WithHooks(hooks ...Hook) MigrateOption {
	return func(a *Atlas) {
		a.hooks = append(a.hooks, hooks...)
	}
}

type (
	// Creator is the interface that wraps the Create method.
	Creator interface {
		// Create creates the given tables in the database. See Migrate.Create for more details.
		Create(context.Context, ...*Table) error
	}

	// The CreateFunc type is an adapter to allow the use of ordinary function as Creator.
	// If f is a function with the appropriate signature, CreateFunc(f) is a Creator that calls f.
	CreateFunc func(context.Context, ...*Table) error

	// Hook defines the "create middleware". A function that gets a Creator and returns a Creator.
	// For example:
	//
	//	hook := func(next schema.Creator) schema.Creator {
	//		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
	//			fmt.Println("Tables:", tables)
	//			return next.Create(ctx, tables...)
	//		})
	//	}
	//
	Hook func(Creator) Creator
)

// Create calls f(ctx, tables...).
func (f CreateFunc) Create(ctx context.Context, tables ...*Table) error {
	return f(ctx, tables...)
}

// exist checks if the given COUNT query returns a value >= 1.
func exist(ctx context.Context, conn dialect.ExecQuerier, query string, args ...any) (bool, error) {
	rows := &sql.Rows{}
	if err := conn.Query(ctx, query, args, rows); err != nil {
		return false, fmt.Errorf("reading schema information %w", err)
	}
	defer rows.Close()
	n, err := sql.ScanInt(rows)
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func indexOf(a []string, s string) int {
	for i := range a {
		if a[i] == s {
			return i
		}
	}
	return -1
}

type sqlDialect interface {
	atBuilder
	dialect.Driver
	init(context.Context) error
	tableExist(context.Context, dialect.ExecQuerier, string) (bool, error)
}

// verifyRanger wraps the method for verifying global-id range correctness.
type verifyRanger interface {
	verifyRange(context.Context, dialect.ExecQuerier, *Table, int64) error
}
