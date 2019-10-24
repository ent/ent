// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// MySQL is a mysql migration driver.
type MySQL struct {
	dialect.Driver
	version string
}

// init loads the MySQL version from the database for later use in the migration process.
func (d *MySQL) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW VARIABLES LIKE 'version'", []interface{}{}, rows); err != nil {
		return fmt.Errorf("mysql: querying mysql version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("mysql: version variable was not found")
	}
	version := make([]string, 2)
	if err := rows.Scan(&version[0], &version[1]); err != nil {
		return fmt.Errorf("mysql: scanning mysql version: %v", err)
	}
	d.version = version[1]
	return nil
}

func (d *MySQL) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("TABLE_NAME", name)).Query()
	return exist(ctx, tx, query, args...)
}

func (d *MySQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("CONSTRAINT_TYPE", "FOREIGN KEY").And().EQ("CONSTRAINT_NAME", name)).Query()
	return exist(ctx, tx, query, args...)
}

// table loads the current table description from the database.
func (d *MySQL) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("TABLE_NAME", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading table description %v", err)
	}
	// call `Close` in cases of failures (`Close` is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := c.ScanMySQL(rows); err != nil {
			return nil, fmt.Errorf("mysql: %v", err)
		}
		if c.PrimaryKey() {
			t.PrimaryKey = append(t.PrimaryKey, c)
		}
		t.AddColumn(c)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("mysql: closing rows %v", err)
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
func (d *MySQL) indexes(ctx context.Context, tx dialect.Tx, name string) ([]*Index, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("index_name", "column_name", "non_unique", "seq_in_index").
		From(sql.Table("INFORMATION_SCHEMA.STATISTICS").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("TABLE_NAME", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading index description %v", err)
	}
	defer rows.Close()
	var idx Indexes
	if err := idx.ScanMySQL(rows); err != nil {
		return nil, fmt.Errorf("mysql: %v", err)
	}
	return idx, nil
}

func (d *MySQL) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` AUTO_INCREMENT = %d", name, value), []interface{}{}, new(sql.Result))
}

func (d *MySQL) cType(c *Column) string                 { return c.MySQLType(d.version) }
func (d *MySQL) tBuilder(t *Table) *sql.TableBuilder    { return t.MySQL(d.version) }
func (d *MySQL) addColumn(c *Column) *sql.ColumnBuilder { return c.MySQL(d.version) }
func (d *MySQL) alterColumn(c *Column) []*sql.ColumnBuilder {
	return []*sql.ColumnBuilder{c.MySQL(d.version)}
}

// addIndex returns the querying for adding an index to MySQL.
func (d *MySQL) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

// addIndex returns the querying for dropping an index in MySQL.
func (d *MySQL) dropIndex(i *Index, table string) *sql.DropIndexBuilder {
	return i.DropBuilder(table)
}
