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

// SQLite is an SQLite migration driver.
type SQLite struct {
	dialect.Driver
}

// init makes sure that foreign_keys support is enabled.
func (d *SQLite) init(ctx context.Context, tx dialect.Tx) error {
	on, err := exist(ctx, tx, "PRAGMA foreign_keys")
	if err != nil {
		return fmt.Errorf("sqlite: check foreign_keys pragma: %v", err)
	}
	if !on {
		// foreign_keys pragma is off, either enable it by execute "PRAGMA foreign_keys=ON"
		// or add the following parameter in the connection string "_fk=1".
		return fmt.Errorf("sqlite: foreign_keys pragma is off: missing %q is the connection string", "_fk=1")
	}
	return nil
}

func (d *SQLite) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_master")).
		Where(sql.EQ("type", "table").And().EQ("name", name)).
		Query()
	return exist(ctx, tx, query, args...)
}

// setRange sets the start value of table PK.
// SQLite tracks the AUTOINCREMENT in the "sqlite_sequence" table that is created and initialized automatically
// whenever a table that contains an AUTOINCREMENT column is created. However, it populates to it a rows (for tables)
// only after the first insertion. Therefore, we check. If a record (for the given table) already exists in the "sqlite_sequence"
// table, we updated it. Otherwise, we insert a new value.
func (d *SQLite) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_sequence")).
		Where(sql.EQ("name", name)).
		Query()
	exists, err := exist(ctx, tx, query, args...)
	switch {
	case err != nil:
		return err
	case exists:
		query, args = sql.Update("sqlite_sequence").Set("seq", value).Where(sql.EQ("name", name)).Query()
	default: // !exists
		query, args = sql.Insert("sqlite_sequence").Columns("name", "seq").Values(name, value).Query()
	}
	return tx.Exec(ctx, query, args, new(sql.Result))
}

// fkExist returns always true to disable foreign-keys creation after the table was created.
func (d *SQLite) fkExist(context.Context, dialect.Tx, string) (bool, error) { return true, nil }
func (d *SQLite) table(context.Context, dialect.Tx, string) (*Table, error) { return nil, nil }

func (*SQLite) cType(c *Column) string                     { return c.SQLiteType() }
func (*SQLite) tBuilder(t *Table) *sql.TableBuilder        { return t.SQLite() }
func (*SQLite) addColumn(c *Column) *sql.ColumnBuilder     { return c.SQLite() }
func (*SQLite) alterColumn(c *Column) []*sql.ColumnBuilder { return []*sql.ColumnBuilder{c.SQLite()} }

func (d *SQLite) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

func (d *SQLite) dropIndex(i *Index, _ string) *sql.DropIndexBuilder {
	return i.DropBuilder("")
}
