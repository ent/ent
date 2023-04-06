// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// InspectOption allows for managing schema configuration using functional options.
type InspectOption func(inspect *Inspector)

// WithSchema provides a schema (named-database) for reading the tables from.
func WithSchema(schema string) InspectOption {
	return func(m *Inspector) {
		m.schema = schema
	}
}

// An Inspector provides methods for inspecting database tables.
type Inspector struct {
	sqlDialect
	schema string
}

// NewInspect returns an inspector for the given SQL driver.
func NewInspect(d dialect.Driver, opts ...InspectOption) (*Inspector, error) {
	i := &Inspector{}
	for _, opt := range opts {
		opt(i)
	}
	switch d.Dialect() {
	case dialect.MySQL:
		i.sqlDialect = &MySQL{Driver: d, schema: i.schema}
	case dialect.SQLite:
		i.sqlDialect = &SQLite{Driver: d}
	case dialect.Postgres:
		i.sqlDialect = &Postgres{Driver: d, schema: i.schema}
	default:
		return nil, fmt.Errorf("sql/schema: unsupported dialect %q", d.Dialect())
	}
	return i, nil
}

// Tables returns the tables in the schema.
func (i *Inspector) Tables(ctx context.Context) ([]*Table, error) {
	names, err := i.tables(ctx)
	if err != nil {
		return nil, err
	}
	tx := dialect.NopTx(i.sqlDialect)
	tables := make([]*Table, 0, len(names))
	for _, name := range names {
		t, err := i.table(ctx, tx, name)
		if err != nil {
			return nil, err
		}
		tables = append(tables, t)
	}

	fki, ok := i.sqlDialect.(interface {
		foreignKeys(context.Context, dialect.Tx, []*Table) error
	})
	if ok {
		if err := fki.foreignKeys(ctx, tx, tables); err != nil {
			return nil, err
		}
	}
	return tables, nil
}

func (i *Inspector) tables(ctx context.Context) ([]string, error) {
	t, ok := i.sqlDialect.(interface{ tables() sql.Querier })
	if !ok {
		return nil, fmt.Errorf("sql/schema: %q driver does not support inspection", i.Dialect())
	}
	query, args := t.tables().Query()
	var (
		names []string
		rows  = &sql.Rows{}
	)
	if err := i.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("%q driver: reading table names %w", i.Dialect(), err)
	}
	defer rows.Close()
	if err := sql.ScanSlice(rows, &names); err != nil {
		return nil, err
	}
	return names, nil
}
