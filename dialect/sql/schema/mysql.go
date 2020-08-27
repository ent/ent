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

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"
)

// MySQL is a MySQL migration driver.
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
		if err := rows.Err(); err != nil {
			return err
		}
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
		Where(sql.And(
			sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
			sql.EQ("TABLE_NAME", name),
		)).Query()
	return exist(ctx, tx, query, args...)
}

func (d *MySQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.And(
			sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
			sql.EQ("CONSTRAINT_TYPE", "FOREIGN KEY"),
			sql.EQ("CONSTRAINT_NAME", name),
		)).Query()
	return exist(ctx, tx, query, args...)
}

// table loads the current table description from the database.
func (d *MySQL) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.And(
			sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
			sql.EQ("TABLE_NAME", name)),
		).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading table description %v", err)
	}
	// Call Close in cases of failures (Close is idempotent).
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("mysql: closing rows %v", err)
	}
	indexes, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	// Add and link indexes to table columns.
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
		Where(sql.And(
			sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
			sql.EQ("TABLE_NAME", name),
		)).
		OrderBy("index_name", "seq_in_index").
		Query()
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

func (d *MySQL) setRange(ctx context.Context, tx dialect.Tx, t *Table, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE `%s` AUTO_INCREMENT = %d", t.Name, value), []interface{}{}, nil)
}

func (d *MySQL) verifyRange(ctx context.Context, tx dialect.Tx, t *Table, expected int) error {
	if expected == 0 {
		return nil
	}
	rows := &sql.Rows{}
	query, args := sql.Select("AUTO_INCREMENT").
		From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.And(
			sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
			sql.EQ("TABLE_NAME", t.Name),
		)).
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return fmt.Errorf("mysql: query auto_increment %v", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	actual := &sql.NullInt64{}
	if err := sql.ScanOne(rows, actual); err != nil {
		return fmt.Errorf("mysql: scan auto_increment %v", err)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	// Table is empty and auto-increment is not configured. This can happen
	// because MySQL (< 8.0) stores the auto-increment counter in main memory
	// (not persistent), and the value is reset on restart (if table is empty).
	if actual.Int64 <= 1 {
		return d.setRange(ctx, tx, t, expected)
	}
	return nil
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
	// Default charset / collation on MySQL table.
	// columns can be override using the Charset / Collate fields.
	b.Charset("utf8mb4").Collate("utf8mb4_bin")
	return b
}

// cType returns the MySQL string type for the given column.
func (d *MySQL) cType(c *Column) (t string) {
	if c.SchemaType != nil && c.SchemaType[dialect.MySQL] != "" {
		// MySQL returns the column type lower cased.
		return strings.ToLower(c.SchemaType[dialect.MySQL])
	}
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
		t = c.scanTypeOr("double")
	case field.TypeTime:
		t = c.scanTypeOr("timestamp")
		// In MySQL, timestamp columns are `NOT NULL` by default, and assigning NULL
		// assigns the current_timestamp(). We avoid this if not set otherwise.
		c.Nullable = c.Attr == ""
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

// addIndex returns the querying for adding an index to MySQL.
func (d *MySQL) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table)
}

// dropIndex drops a MySQL index.
func (d *MySQL) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	query, args := idx.DropBuilder(table).Query()
	return tx.Exec(ctx, query, args, nil)
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
			if err := tx.Exec(ctx, query, args, nil); err != nil {
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
	parts, size, unsigned, err := parseColumn(c.typ)
	if err != nil {
		return err
	}
	switch parts[0] {
	case "mediumint", "int":
		c.Type = field.TypeInt32
		if unsigned {
			c.Type = field.TypeUint32
		}
	case "smallint":
		c.Type = field.TypeInt16
		if unsigned {
			c.Type = field.TypeUint16
		}
	case "bigint":
		c.Type = field.TypeInt64
		if unsigned {
			c.Type = field.TypeUint64
		}
	case "tinyint":
		switch {
		case size == 1:
			c.Type = field.TypeBool
		case unsigned:
			c.Type = field.TypeUint8
		default:
			c.Type = field.TypeInt8
		}
	case "numeric", "decimal", "double":
		c.Type = field.TypeFloat64
	case "time", "timestamp", "date", "datetime":
		c.Type = field.TypeTime
		// The mapping from schema defaults to database
		// defaults is not supported for TypeTime fields.
		defaults = sql.NullString{}
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
		c.Size = size
	case "varchar":
		c.Type = field.TypeString
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
		// Ignore primary keys.
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return i, nil
}

// isImplicitIndex reports if the index was created implicitly for the unique column.
func (d *MySQL) isImplicitIndex(idx *Index, col *Column) bool {
	// We execute `CHANGE COLUMN` on older versions of MySQL (<8.0), which
	// auto create the new index. The old one, will be dropped in `changeSet`.
	if compareVersions(d.version, "8.0.0") >= 0 {
		return idx.Name == col.Name && col.Unique
	}
	return false
}

// renameColumn returns the statement for renaming a column in
// MySQL based on its version.
func (d *MySQL) renameColumn(t *Table, old, new *Column) sql.Querier {
	q := sql.AlterTable(t.Name)
	if compareVersions(d.version, "8.0.0") >= 0 {
		return q.RenameColumn(old.Name, new.Name)
	}
	return q.ChangeColumn(old.Name, d.addColumn(new))
}

// renameIndex returns the statement for renaming an index.
func (d *MySQL) renameIndex(t *Table, old, new *Index) sql.Querier {
	q := sql.AlterTable(t.Name)
	if compareVersions(d.version, "5.7.0") >= 0 {
		return q.RenameIndex(old.Name, new.Name)
	}
	return q.DropIndex(old.Name).AddIndex(new.Builder(t.Name))
}

// tableSchema returns the query for getting the table schema.
func (d *MySQL) tableSchema() sql.Querier {
	return sql.Raw("(SELECT DATABASE())")
}

// alterColumns returns the queries for applying the columns change-set.
func (d *MySQL) alterColumns(table string, add, modify, drop []*Column) sql.Queries {
	b := sql.Dialect(dialect.MySQL).AlterTable(table)
	for _, c := range add {
		b.AddColumn(d.addColumn(c))
	}
	for _, c := range modify {
		b.ModifyColumn(d.addColumn(c))
	}
	for _, c := range drop {
		b.DropColumn(sql.Dialect(dialect.MySQL).Column(c.Name))
	}
	if len(b.Queries) == 0 {
		return nil
	}
	return sql.Queries{b}
}

// parseColumn returns column parts, size and signed-info from a MySQL type.
func parseColumn(typ string) (parts []string, size int64, unsigned bool, err error) {
	switch parts = strings.FieldsFunc(typ, func(r rune) bool {
		return r == '(' || r == ')' || r == ' ' || r == ','
	}); parts[0] {
	case "tinyint", "smallint", "mediumint", "int", "bigint":
		switch {
		case len(parts) == 2 && parts[1] == "unsigned": // int unsigned
			unsigned = true
		case len(parts) == 3: // int(10) unsigned
			unsigned = true
			fallthrough
		case len(parts) == 2: // int(10)
			size, err = strconv.ParseInt(parts[1], 10, 0)
		}
	case "varbinary", "varchar", "char":
		size, err = strconv.ParseInt(parts[1], 10, 64)
	}
	if err != nil {
		return parts, size, unsigned, fmt.Errorf("converting %s size to int: %v", parts[0], err)
	}
	return parts, size, unsigned, nil
}

// fkNames returns the foreign-key names of a column.
func fkNames(ctx context.Context, tx dialect.Tx, table, column string) ([]string, error) {
	query, args := sql.Select("CONSTRAINT_NAME").From(sql.Table("INFORMATION_SCHEMA.KEY_COLUMN_USAGE").Unquote()).
		Where(sql.And(
			sql.EQ("TABLE_NAME", table),
			sql.EQ("COLUMN_NAME", column),
			// NULL for unique and primary-key constraints.
			sql.NotNull("POSITION_IN_UNIQUE_CONSTRAINT"),
			sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")),
		)).
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
