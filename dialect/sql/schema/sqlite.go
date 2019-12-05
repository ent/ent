// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pkg/errors"
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

func (d *SQLite) typeField(c *Column, str string) error {
	switch parts := typeFields(str); parts[0] {
	case "int", "integer":
		c.Type = field.TypeInt32
	case "smallint":
		c.Type = field.TypeInt16
	case "bigint":
		c.Type = field.TypeInt64
	case "tinyint":
		c.Type = field.TypeInt8
	case "double", "real":
		c.Type = field.TypeFloat64
	case "timestamp", "datetime":
		c.Type = field.TypeTime
	case "blob":
		c.Size = math.MaxUint32
		c.Type = field.TypeBytes
	case "text":
		c.Size = math.MaxInt32
		c.Type = field.TypeString
	case "varchar":
		c.Type = field.TypeString
		if len(parts) > 1 {
			size, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return fmt.Errorf("converting varchar size to int: %v", err)
			}
			c.Size = size
		}
	case "json":
		c.Type = field.TypeJSON
	case "uuid":
		c.Type = field.TypeUUID
	case "enum":
		c.Type = field.TypeEnum
		c.Enums = make([]string, len(parts)-1)
		for i, e := range parts[1:] {
			c.Enums[i] = strings.Trim(e, "'")
		}
	default:
		return fmt.Errorf("unknown column type %q", parts[0])
	}

	return nil
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
func (d *SQLite) fkExist(ctx context.Context, tx dialect.Tx, table *Table, fk *ForeignKey) (bool, error) {
	rows := &sql.Rows{}

	if err := tx.Query(ctx, fmt.Sprintf("pragma foreign_key_list('%s')", table.Name), []interface{}{}, rows); err != nil {
		return false, fmt.Errorf("sqlite3: reading table foreign keys description %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        int
			seq       int
			tableName string
			from      string
			to        string
			onUpdate  string
			onDelete  string
			match     string
		)
		if err := rows.Scan(&id, &seq, &tableName, &from, &to, &onUpdate, &onDelete, &match); err != nil {
			return false, errors.Wrap(err, "while querying foreign keys")
		}

		if tableName == fk.RefTable.Name && fk.Columns[0].Name == from && fk.RefColumns[0].Name == to {
			return true, rows.Close()
		}
	}

	return false, rows.Close()
}

func (d *SQLite) indexes(ctx context.Context, tx dialect.Tx, table string) (Indexes, error) {
	idxs := Indexes{}
	rows := &sql.Rows{}

	if err := tx.Query(ctx, fmt.Sprintf("pragma index_list('%s')", table), []interface{}{}, rows); err != nil {
		return nil, fmt.Errorf("sqlite3: reading table index description %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			seq     int
			name    string
			unique  int
			origin  string
			partial int
		)
		if err := rows.Scan(&seq, &name, &unique, &origin, &partial); err != nil {
			return nil, errors.Wrap(err, "while querying indexes")
		}

		if origin == "pk" {
			continue // we already know the primary key
		}

		idxs = append(idxs, &Index{
			Name:   name,
			Unique: unique == 1,
		})
	}
	rows.Close()

	// second loop to gather column info

	for _, idx := range idxs {
		if err := tx.Query(ctx, fmt.Sprintf("pragma index_info('%s')", idx.Name), []interface{}{}, rows); err != nil {
			return nil, fmt.Errorf("sqlite3: reading index description %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			var (
				seq        int
				columnID   int
				columnName string
			)
			if err := rows.Scan(&seq, &columnID, &columnName); err != nil {
				return nil, errors.Wrap(err, "while querying indexes")
			}

			idx.columns = append(idx.columns, columnName)
		}
		rows.Close()
	}

	return idxs, nil
}

func (d *SQLite) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		id         int
		name       string
		typ        string
		notnullInt int
		dflt       sql.NullString
		pkInt      int
	)

	if err := rows.Scan(&id, &name, &typ, &notnullInt, &dflt, &pkInt); err != nil {
		return err
	}

	c.Name = name
	c.Nullable = notnullInt == 0
	if pkInt > 0 {
		c.Key = PrimaryKey
	}

	return d.typeField(c, typ)
}

func (d *SQLite) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, fmt.Sprintf("pragma table_info('%s')", name), []interface{}{}, rows); err != nil {
		return nil, fmt.Errorf("sqlite3: reading table description %v", err)
	}
	// call `Close` in cases of failures (`Close` is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, err
		}
		t.AddColumn(c)
		if c.Key == PrimaryKey {
			t.PrimaryKey = append(t.PrimaryKey, c)
		}
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("closing rows %v", err)
	}
	idxs, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}

	return t, processIndexes(idxs, t)
}
