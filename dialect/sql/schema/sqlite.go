// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"
)

// SQLite is an SQLite migration driver.
type SQLite struct {
	dialect.Driver
	WithForeignKeys bool
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
		Where(sql.And(
			sql.EQ("type", "table"),
			sql.EQ("name", name),
		)).
		Query()
	return exist(ctx, tx, query, args...)
}

// setRange sets the start value of table PK.
// SQLite tracks the AUTOINCREMENT in the "sqlite_sequence" table that is created and initialized automatically
// whenever a table that contains an AUTOINCREMENT column is created. However, it populates to it a rows (for tables)
// only after the first insertion. Therefore, we check. If a record (for the given table) already exists in the "sqlite_sequence"
// table, we updated it. Otherwise, we insert a new value.
func (d *SQLite) setRange(ctx context.Context, tx dialect.Tx, t *Table, value int) error {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_sequence")).
		Where(sql.EQ("name", t.Name)).
		Query()
	exists, err := exist(ctx, tx, query, args...)
	switch {
	case err != nil:
		return err
	case exists:
		query, args = sql.Update("sqlite_sequence").Set("seq", value).Where(sql.EQ("name", t.Name)).Query()
	default: // !exists
		query, args = sql.Insert("sqlite_sequence").Columns("name", "seq").Values(t.Name, value).Query()
	}
	return tx.Exec(ctx, query, args, nil)
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
	if d.WithForeignKeys {
		for _, fk := range t.ForeignKeys {
			b.ForeignKeys(fk.DSL())
		}
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
	if c.SchemaType != nil && c.SchemaType[dialect.SQLite] != "" {
		return c.SchemaType[dialect.SQLite]
	}
	switch c.Type {
	case field.TypeBool:
		t = "bool"
	case field.TypeInt8, field.TypeUint8, field.TypeInt16, field.TypeUint16, field.TypeInt32,
		field.TypeUint32, field.TypeUint, field.TypeInt, field.TypeInt64, field.TypeUint64:
		t = "integer"
	case field.TypeBytes:
		t = "blob"
	case field.TypeString, field.TypeEnum:
		// SQLite does not impose any length restrictions on
		// the length of strings, BLOBs or numeric values.
		t = fmt.Sprintf("varchar(%d)", DefaultStringLen)
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

// addIndex returns the querying for adding an index to SQLite.
func (d *SQLite) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

// dropIndex drops a SQLite index.
func (d *SQLite) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	query, args := idx.DropBuilder("").Query()
	return tx.Exec(ctx, query, args, nil)
}

// fkExist returns always true to disable foreign-keys creation after the table was created.
func (d *SQLite) fkExist(context.Context, dialect.Tx, string) (bool, error) { return true, nil }

// table returns always error to indicate that SQLite dialect doesn't support incremental migration.
func (d *SQLite) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("name", "type", "notnull", "dflt_value", "pk").
		From(sql.Table(fmt.Sprintf("pragma_table_info('%s')", name)).Unquote()).
		OrderBy("pk").
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("sqlite: reading table description %v", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, fmt.Errorf("sqlite: %v", err)
		}
		if c.PrimaryKey() {
			t.PrimaryKey = append(t.PrimaryKey, c)
		}
		t.AddColumn(c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("sqlite: closing rows %v", err)
	}
	indexes, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	// Add and link indexes to table columns.
	for _, idx := range indexes {
		switch {
		case idx.primary:
		case idx.Unique && len(idx.columns) == 1:
			name := idx.columns[0]
			c, ok := t.column(name)
			if !ok {
				return nil, fmt.Errorf("index %q column %q was not found in table %q", idx.Name, name, t.Name)
			}
			c.Key = UniqueKey
			c.Unique = true
			fallthrough
		default:
			t.addIndex(idx)
		}
	}
	return t, nil
}

// table loads the table indexes from the database.
func (d *SQLite) indexes(ctx context.Context, tx dialect.Tx, name string) (Indexes, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("name", "unique", "origin").
		From(sql.Table(fmt.Sprintf("pragma_index_list('%s')", name)).Unquote()).
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("reading table indexes %v", err)
	}
	defer rows.Close()
	var idx Indexes
	for rows.Next() {
		i := &Index{}
		origin := sql.NullString{}
		if err := rows.Scan(&i.Name, &i.Unique, &origin); err != nil {
			return nil, fmt.Errorf("scanning index description %v", err)
		}
		i.primary = origin.String == "pk"
		idx = append(idx, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("closing rows %v", err)
	}
	for i := range idx {
		columns, err := d.indexColumns(ctx, tx, idx[i].Name)
		if err != nil {
			return nil, err
		}
		idx[i].columns = columns
		// Normalize implicit index names to ent naming convention. See:
		// https://github.com/sqlite/sqlite/blob/e937df8/src/build.c#L3583
		if len(columns) == 1 && strings.HasPrefix(idx[i].Name, "sqlite_autoindex_"+name) {
			idx[i].Name = columns[0]
		}
	}
	return idx, nil
}

// indexColumns loads index columns from index info.
func (d *SQLite) indexColumns(ctx context.Context, tx dialect.Tx, name string) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("name").
		From(sql.Table(fmt.Sprintf("pragma_index_info('%s')", name)).Unquote()).
		OrderBy("seqno").
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("reading table indexes %v", err)
	}
	defer rows.Close()
	var names []string
	if err := sql.ScanSlice(rows, &names); err != nil {
		return nil, err
	}
	return names, nil
}

// scanColumn scans the column information from SQLite column description.
func (d *SQLite) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		pk       sql.NullInt64
		notnull  sql.NullInt64
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &notnull, &defaults, &pk); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	c.Nullable = notnull.Int64 == 0
	if pk.Int64 > 0 {
		c.Key = PrimaryKey
	}
	parts, _, _, err := parseColumn(c.typ)
	if err != nil {
		return err
	}
	switch parts[0] {
	case "bool", "boolean":
		c.Type = field.TypeBool
	case "blob":
		c.Type = field.TypeBytes
	case "integer":
		// All integer types have the same "type affinity".
		c.Type = field.TypeInt
	case "real", "float", "double":
		c.Type = field.TypeFloat64
	case "datetime":
		c.Type = field.TypeTime
	case "json":
		c.Type = field.TypeJSON
	case "uuid":
		c.Type = field.TypeUUID
	case "varchar", "text":
		c.Size = DefaultStringLen
		c.Type = field.TypeString
	}
	if defaults.Valid {
		return c.ScanDefault(defaults.String)
	}
	return nil
}

// alterColumns returns the queries for applying the columns change-set.
func (d *SQLite) alterColumns(table string, add, _, _ []*Column) sql.Queries {
	queries := make(sql.Queries, 0, len(add))
	for i := range add {
		c := d.addColumn(add[i])
		if fk := add[i].foreign; fk != nil {
			c.Constraint(fk.DSL())
		}
		queries = append(queries, sql.Dialect(dialect.SQLite).AlterTable(table).AddColumn(c))
	}
	// Modifying and dropping columns is not supported and disabled until we
	// will support https://www.sqlite.org/lang_altertable.html#otheralter
	return queries
}
