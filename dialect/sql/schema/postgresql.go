// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// PostgreSQL is a mysql migration driver.
type PostgreSQL struct {
	dialect.Driver
	version string
}

// init loads the MySQL version from the database for later use in the migration process.
func (d *PostgreSQL) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW server_version_num", []interface{}{}, rows); err != nil {
		return fmt.Errorf("postgresql: querying mysql version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("postgresql: version variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("postgresql: scanning postgresql version: %v", err)
	}

	major, err := strconv.Atoi(version[:2])
	if err != nil {
		return fmt.Errorf("postgresql: failed to parse major version: %v", err)
	}

	minor, err := strconv.Atoi(version[2:])
	if err != nil {
		return fmt.Errorf("postgresql: failed to parse minor version: %v", err)
	}

	// TODO: should we keep the version number as machine readable number?
	d.version = fmt.Sprintf("%d.%d", major, minor)

	return nil
}

func (d *PostgreSQL) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("(SELECT CURRENT_DATABASE())")).And().EQ("table_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

func (d *PostgreSQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("(SELECT CURRENT_DATABASE())")).And().EQ("CONSTRAINT_TYPE", "FOREIGN KEY").And().EQ("CONSTRAINT_NAME", name)).Query()
	return exist(ctx, tx, query, args...)
}

// table loads the current table description from the database.
func (d *PostgreSQL) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(dialect.Postgres).Select("column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("(SELECT DATABASE())")).And().EQ("table_name", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("postgresql: reading table description %v", err)
	}
	// call `Close` in cases of failures (`Close` is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := c.ScanPostgreSQL(rows); err != nil {
			return nil, fmt.Errorf("PostgreSQL: %v", err)
		}
		if c.PrimaryKey() {
			t.PrimaryKey = append(t.PrimaryKey, c)
		}
		t.AddColumn(c)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("postgresql: closing rows %v", err)
	}
	indexes, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	// add and link indexes to table columns.
	for _, idx := range indexes {
		t.AddIndex(idx.Name, idx.Unique, idx.columns)
	}
	return t, nil
}

// table loads the table indexes from the database.
func (d *PostgreSQL) indexes(ctx context.Context, tx dialect.Tx, name string) ([]*Index, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(dialect.Postgres).Select("index_name", "column_name", "non_unique", "seq_in_index").
		From(sql.Table("INFORMATION_SCHEMA.STATISTICS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("(SELECT DATABASE())")).And().EQ("table_name", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("postgresql: reading index description %v", err)
	}
	defer rows.Close()
	var idx Indexes
	if err := idx.ScanPostgreSQL(rows); err != nil {
		return nil, fmt.Errorf("postgresql: %v", err)
	}
	return idx, nil
}

func (d *PostgreSQL) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` AUTO_INCREMENT = %d", name, value), []interface{}{}, new(sql.Result))
}

func (d *PostgreSQL) cType(c *Column) string                { return c.PostgreSQLType() }
func (d *PostgreSQL) tBuilder(t *Table) *sql.TableBuilder   { return t.PostgreSQL() }
func (d *PostgreSQL) cBuilder(c *Column) *sql.ColumnBuilder { return c.PostgreSQL() }
