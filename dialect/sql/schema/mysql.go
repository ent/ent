// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
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
		if err := d.scanColumn(c, rows); err != nil {
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
	idx, err := d.scanIndexes(rows)
	if err != nil {
		return nil, fmt.Errorf("mysql: %v", err)
	}
	return idx, nil
}

func (d *MySQL) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` AUTO_INCREMENT = %d", name, value), []interface{}{}, new(sql.Result))
}

// tBuilder returns the MySQL DSL query for table creation.
func (d *MySQL) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(d.addColumn(c))
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	// default charset / collation on MySQL table.
	// columns can be override using the Charset / Collate fields.
	b.Charset("utf8mb4").Collate("utf8mb4_bin")
	return b
}

// cType returns the MySQL string type for the given column.
func (d *MySQL) cType(c *Column) (t string) {
	switch c.Type {
	case field.TypeBool:
		t = "boolean"
	case field.TypeInt8:
		t = "tinyint"
	case field.TypeUint8:
		t = "tinyint unsigned"
	case field.TypeInt16:
		t = "smallint"
	case field.TypeUint16:
		t = "smallint unsigned"
	case field.TypeInt32:
		t = "int"
	case field.TypeUint32:
		t = "int unsigned"
	case field.TypeInt, field.TypeInt64:
		t = "bigint"
	case field.TypeUint, field.TypeUint64:
		t = "bigint unsigned"
	case field.TypeBytes:
		size := int64(math.MaxUint16)
		if c.Size > 0 {
			size = c.Size
		}
		switch {
		case size <= math.MaxUint8:
			t = "tinyblob"
		case size <= math.MaxUint16:
			t = "blob"
		case size < 1<<24:
			t = "mediumblob"
		case size <= math.MaxUint32:
			t = "longblob"
		}
	case field.TypeJSON:
		t = "json"
		if compareVersions(d.version, "5.7.8") == -1 {
			t = "longblob"
		}
	case field.TypeString:
		size := c.Size
		if size == 0 {
			size = c.defaultSize(d.version)
		}
		if size <= math.MaxUint16 {
			t = fmt.Sprintf("varchar(%d)", size)
		} else {
			t = "longtext"
		}
	case field.TypeFloat32, field.TypeFloat64:
		t = "double"
	case field.TypeTime:
		t = "timestamp"
		// in MySQL timestamp columns are `NOT NULL by default, and assigning NULL
		// assigns the current_timestamp(). We avoid this if not set otherwise.
		c.Nullable = true
	case field.TypeEnum:
		values := make([]string, len(c.Enums))
		for i, e := range c.Enums {
			values[i] = fmt.Sprintf("'%s'", e)
		}
		sort.Strings(values)
		t = fmt.Sprintf("enum(%s)", strings.Join(values, ", "))
	case field.TypeUUID:
		t = "char(36) binary"
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
	return t
}

// addColumn returns the DSL query for adding the given column to a table.
// The syntax/order is: datatype [Charset] [Unique|Increment] [Collation] [Nullable].
func (d *MySQL) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(d.cType(c)).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("AUTO_INCREMENT")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// alterColumn returns the DSL query for modifying the given column.
func (d *MySQL) alterColumn(c *Column) []*sql.ColumnBuilder {
	return []*sql.ColumnBuilder{d.addColumn(c)}
}

// addIndex returns the querying for adding an index to MySQL.
func (d *MySQL) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

// dropIndex drops a MySQL index.
func (d *MySQL) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	query, args := idx.DropBuilder(table).Query()
	return tx.Exec(ctx, query, args, new(sql.Result))
}

// prepare runs preparation work that needs to be done to apply the change-set.
func (d *MySQL) prepare(ctx context.Context, tx dialect.Tx, change *changes, table string) error {
	for _, idx := range change.index.drop {
		switch n := len(idx.columns); {
		case n == 0:
			return fmt.Errorf("index %q has no columns", idx.Name)
		case n > 1:
			continue // not a foreign-key index.
		}
		var qr sql.Querier
	Switch:
		switch col, ok := change.dropColumn(idx.columns[0]); {
		// If both the index and the column need to be dropped, the foreign-key
		// constraint that is associated with them need to be dropped as well.
		case ok:
			names, err := fkNames(ctx, tx, table, col.Name)
			if err != nil {
				return err
			}
			if len(names) == 1 {
				qr = sql.AlterTable(table).DropForeignKey(names[0])
			}
		// If the uniqueness was dropped from a foreign-key column,
		// create a "simple index" if no other index exist for it.
		case !ok && idx.Unique && len(idx.Columns) > 0:
			col := idx.Columns[0]
			for _, idx2 := range col.indexes {
				if idx2 != idx && len(idx2.columns) == 1 {
					break Switch
				}
			}
			names, err := fkNames(ctx, tx, table, col.Name)
			if err != nil {
				return err
			}
			if len(names) == 1 {
				qr = sql.CreateIndex(names[0]).Table(table).Columns(col.Name)
			}
		}
		if qr != nil {
			query, args := qr.Query()
			if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
				return err
			}
		}
	}
	return nil
}

// scanColumn scans the column information from MySQL column description.
func (d *MySQL) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		nullable sql.NullString
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &c.Key, &defaults, &c.Attr, &sql.NullString{}, &sql.NullString{}); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	c.Unique = c.UniqueKey()
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	switch parts := strings.FieldsFunc(c.typ, func(r rune) bool {
		return r == '(' || r == ')' || r == ' ' || r == ','
	}); parts[0] {
	case "int":
		c.Type = field.TypeInt32
	case "smallint":
		c.Type = field.TypeInt16
		if len(parts) == 3 { // smallint(5) unsigned.
			c.Type = field.TypeUint16
		}
	case "bigint":
		c.Type = field.TypeInt64
		if len(parts) == 3 { // bigint(20) unsigned.
			c.Type = field.TypeUint64
		}
	case "tinyint":
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("converting varchar size to int: %v", err)
		}
		switch {
		case size == 1:
			c.Type = field.TypeBool
		case len(parts) == 3: // tinyint(3) unsigned.
			c.Type = field.TypeUint8
		default:
			c.Type = field.TypeInt8
		}
	case "double":
		c.Type = field.TypeFloat64
	case "timestamp", "datetime":
		c.Type = field.TypeTime
	case "tinyblob":
		c.Size = math.MaxUint8
		c.Type = field.TypeBytes
	case "blob":
		c.Size = math.MaxUint16
		c.Type = field.TypeBytes
	case "mediumblob":
		c.Size = 1<<24 - 1
		c.Type = field.TypeBytes
	case "longblob":
		c.Size = math.MaxUint32
		c.Type = field.TypeBytes
	case "varbinary":
		c.Type = field.TypeBytes
		size, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return fmt.Errorf("converting varbinary size to int: %v", err)
		}
		c.Size = size
	case "varchar":
		c.Type = field.TypeString
		size, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return fmt.Errorf("converting varchar size to int: %v", err)
		}
		c.Size = size
	case "longtext":
		c.Size = math.MaxInt32
		c.Type = field.TypeString
	case "json":
		c.Type = field.TypeJSON
	case "enum":
		c.Type = field.TypeEnum
		c.Enums = make([]string, len(parts)-1)
		for i, e := range parts[1:] {
			c.Enums[i] = strings.Trim(e, "'")
		}
	case "char":
		size, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return fmt.Errorf("converting char size to int: %v", err)
		}
		// UUID field has length of 36 characters (32 alphanumeric characters and 4 hyphens).
		if size != 36 {
			return fmt.Errorf("unknown char(%d) type (not a uuid)", size)
		}
		c.Type = field.TypeUUID
	default:
		return fmt.Errorf("unknown column type %q for version %q", parts[0], d.version)
	}
	if defaults.Valid {
		return c.ScanDefault(defaults.String)
	}
	return nil
}

// scanIndexes scans sql.Rows into an Indexes list. The query for returning the rows,
// should return the following 4 columns: INDEX_NAME, COLUMN_NAME, NON_UNIQUE, SEQ_IN_INDEX.
// SEQ_IN_INDEX specifies the position of the column in the index columns.
func (d *MySQL) scanIndexes(rows *sql.Rows) (Indexes, error) {
	var (
		i     Indexes
		names = make(map[string]*Index)
	)
	for rows.Next() {
		var (
			name     string
			column   string
			nonuniq  bool
			seqindex int
		)
		if err := rows.Scan(&name, &column, &nonuniq, &seqindex); err != nil {
			return nil, fmt.Errorf("scanning index description: %v", err)
		}
		// ignore primary keys.
		if name == "PRIMARY" {
			continue
		}
		idx, ok := names[name]
		if !ok {
			idx = &Index{Name: name, Unique: !nonuniq}
			i = append(i, idx)
			names[name] = idx
		}
		idx.columns = append(idx.columns, column)
	}
	return i, nil
}

// fkNames returns the foreign-key names of a column.
func fkNames(ctx context.Context, tx dialect.Tx, table, column string) ([]string, error) {
	query, args := sql.Select("CONSTRAINT_NAME").From(sql.Table("INFORMATION_SCHEMA.KEY_COLUMN_USAGE").Unquote()).
		Where(sql.
			EQ("TABLE_NAME", table).
			And().EQ("COLUMN_NAME", column).
			// NULL for unique and primary-key constraints.
			And().NotNull("POSITION_IN_UNIQUE_CONSTRAINT").
			And().EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
		).
		Query()
	var (
		names []string
		rows  = &sql.Rows{}
	)
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading constraint names %v", err)
	}
	defer rows.Close()
	if err := sql.ScanSlice(rows, &names); err != nil {
		return nil, err
	}
	return names, nil
}
