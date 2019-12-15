// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
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

func (d *SQLite) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.CreateTable(t.Name)
	for _, c := range t.Columns {
		b.Column(d.addColumn(c))
	}
	// Unlike in MySQL, we're not able to add foreign-key constraints to table
	// after it was created, and adding them to the `CREATE TABLE` statement is
	// not always valid (because circular foreign-keys situation is possible).
	// We stay consistent by not using constraints at all, and just defining the
	// foreign keys in the `CREATE TABLE` statement.
	for _, fk := range t.ForeignKeys {
		b.ForeignKeys(fk.DSL())
	}
	// If it's an ID based primary key with autoincrement, we add
	// the `PRIMARY KEY` clause to the column declaration. Otherwise,
	// we append it to the constraint clause.
	if len(t.PrimaryKey) == 1 && t.PrimaryKey[0].Increment {
		return b
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	return b
}

// cType returns the SQLite string type for the given column.
func (*SQLite) cType(c *Column) (t string) {
	switch c.Type {
	case field.TypeBool:
		t = "bool"
	case field.TypeInt8, field.TypeUint8, field.TypeInt, field.TypeInt16, field.TypeInt32, field.TypeUint, field.TypeUint16, field.TypeUint32:
		t = "integer"
	case field.TypeInt64, field.TypeUint64:
		t = "bigint"
	case field.TypeBytes:
		t = "blob"
	case field.TypeString, field.TypeEnum:
		size := c.Size
		if size == 0 {
			size = DefaultStringLen
		}
		// sqlite has no size limit on varchar.
		t = fmt.Sprintf("varchar(%d)", size)
	case field.TypeFloat32, field.TypeFloat64:
		t = "real"
	case field.TypeTime:
		t = "datetime"
	case field.TypeJSON:
		t = "json"
	case field.TypeUUID:
		t = "uuid"
	default:
		panic("unsupported type " + c.Type.String())
	}
	return t
}

// addColumn returns the DSL query for adding the given column to a table.
func (d *SQLite) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(d.cType(c)).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("PRIMARY KEY AUTOINCREMENT")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// alterColumn returns the DSL query for modifying the given column.
func (d *SQLite) alterColumn(c *Column) []*sql.ColumnBuilder {
	return []*sql.ColumnBuilder{d.addColumn(c)}
}

// addIndex returns the querying for adding an index to SQLite.
func (d *SQLite) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

// dropIndex drops a SQLite index.
func (d *SQLite) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	query, args := idx.DropBuilder("").Query()
	return tx.Exec(ctx, query, args, new(sql.Result))
}

// fkExist returns always true to disable foreign-keys creation after the table was created.
func (d *SQLite) fkExist(context.Context, dialect.Tx, string) (bool, error) { return true, nil }
func (d *SQLite) table(context.Context, dialect.Tx, string) (*Table, error) { return nil, nil }
