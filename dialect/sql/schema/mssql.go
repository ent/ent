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
	"math"
	"regexp"
	"strings"
)

// MSSQL is a MSSQL migration driver.
type MSSQL struct {
	dialect.Driver
	version string
}

// init loads the MSSQL version from the database for later use in the migration process.
func (d *MSSQL) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SELECT SERVERPROPERTY('productversion')", []interface{}{}, rows); err != nil {
		return fmt.Errorf("MSSQL: querying MSSQL version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("MSSQL: version variable was not found")
	}
	if err := rows.Scan(&d.version); err != nil {
		return fmt.Errorf("MSSQL: scanning MSSQL version: %v", err)
	}
	return nil
}

func (d *MSSQL) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.MSSQL).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", d.tableSchema()).And().EQ("TABLE_NAME", name)).Query()
	return exist(ctx, tx, query, args...)
}

func (d *MSSQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.MSSQL).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", d.tableSchema()).And().EQ("CONSTRAINT_TYPE", "FOREIGN KEY").And().EQ("CONSTRAINT_NAME", name)).Query()
	return exist(ctx, tx, query, args...)
}

// table loads the current table description from the database.
func (d *MSSQL) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query := `
select 
		columns.name                                 as col_name,
       	types.name                                   as type_name,
		columns.is_nullable,
       	IIF(columns.default_object_id = 0, null, object_definition(columns.default_object_id)) as "default",
	    columns.is_identity,
       	columns.max_length,
       	columns.precision,
       	columns.scale
from sys.columns
         inner join sys.types on columns.system_type_id = types.system_type_id
         inner join sys.objects on objects.object_id = columns.object_id
where objects.schema_id = SCHEMA_ID() and objects.name = @p1`

	if err := tx.Query(ctx, query, []interface{}{name}, rows); err != nil {
		return nil, fmt.Errorf("MSSQL: reading table description %v", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()

	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, fmt.Errorf("MSSQL: %v", err)
		}
		t.AddColumn(c)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("MSSQL: closing rows %v", err)
	}

	idxs, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}

	for _, idx := range idxs {
		switch {
		case idx.primary:
			for _, name := range idx.columns {
				c, ok := t.column(name)
				if !ok {
					return nil, fmt.Errorf("index %q column %q was not found in table %q", idx.Name, name, t.Name)
				}
				c.Key = PrimaryKey
				t.PrimaryKey = append(t.PrimaryKey, c)
			}
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
func (d *MSSQL) indexes(ctx context.Context, tx dialect.Tx, name string) ([]*Index, error) {
	rows := &sql.Rows{}
	query := `
SELECT 
	ind.name                                   as index_name,
    col.name                                   as col_name,
    (ind.is_unique | ind.is_unique_constraint) as is_unique,
	is_primary_key,
	key_ordinal
FROM sys.indexes ind
         INNER JOIN
     sys.index_columns ic ON ind.object_id = ic.object_id and ind.index_id = ic.index_id
         INNER JOIN
     sys.columns col ON ic.object_id = col.object_id and ic.column_id = col.column_id
         INNER JOIN
     sys.tables t ON ind.object_id = t.object_id
WHERE t.is_ms_shipped = 0
  AND t.name = @p1
ORDER BY t.name, ind.name, key_ordinal`

	if err := tx.Query(ctx, query, []interface{}{name}, rows); err != nil {
		return nil, fmt.Errorf("MSSQL: reading index description %v", err)
	}
	defer rows.Close()
	idx, err := d.scanIndexes(rows)
	if err != nil {
		return nil, fmt.Errorf("MSSQL: %v", err)
	}
	return idx, nil
}

func (d *MSSQL) setRange(ctx context.Context, tx dialect.Tx, t *Table, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE [%s] AUTO_INCREMENT = %d", t.Name, value), []interface{}{}, nil)
}

func (d *MSSQL) verifyRange(ctx context.Context, tx dialect.Tx, t *Table, expected int) error {
	if expected == 0 {
		return nil
	}
	rows := &sql.Rows{}
	query, args := sql.Dialect(dialect.MSSQL).
		Select("AUTO_INCREMENT").
		From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", d.tableSchema()).And().EQ("TABLE_NAME", t.Name)).
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return fmt.Errorf("MSSQL: query auto_increment %v", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	actual := &sql.NullInt64{}
	if err := sql.ScanOne(rows, actual); err != nil {
		return fmt.Errorf("MSSQL: scan auto_increment %v", err)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	// Table is empty and auto-increment is not configured. This can happen
	// because MSSQL (< 8.0) stores the auto-increment counter in main memory
	// (not persistent), and the value is reset on restart (if table is empty).
	if actual.Int64 <= 1 {
		return d.setRange(ctx, tx, t, expected)
	}
	return nil
}

// tBuilder returns the MSSQL DSL query for table creation.
func (d *MSSQL) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.Dialect(dialect.MSSQL).CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(d.addColumn(c))
	}

	if len(t.PrimaryKey) > 0 {
		b.Constraints(d.primaryKeyConstraint(t.Name, t.PrimaryKey))
	}

	return b
}

func (d *MSSQL) primaryKeyConstraint(table string, cols []*Column) *sql.TableConstraintBuilder {
	var names []string

	for _, c := range cols {
		names = append(names, c.Name)
	}

	return sql.TableConstraint(fmt.Sprintf("%s_%s_pk", table, strings.Join(names, "_"))).
		Columns(names...).Primary()
}

// cType returns the MSSQL string type for the given column.
func (d *MSSQL) cType(c *Column) (t string) {
	if c.SchemaType != nil && c.SchemaType[dialect.MSSQL] != "" {
		return c.SchemaType[dialect.MSSQL]
	}
	switch c.Type {
	case field.TypeBool:
		t = "bit"
	case field.TypeInt8, field.TypeUint8, field.TypeInt16, field.TypeUint16:
		t = "smallint"
	case field.TypeInt32, field.TypeUint32:
		t = "int"
	case field.TypeInt, field.TypeInt64, field.TypeUint, field.TypeUint64:
		t = "bigint"
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
		// TODO: what value is correct here
		// t = "varchar(max)"
		// max breaks int64 conversion
		t = "varchar(8000)"
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
		t = "float"
	case field.TypeTime:
		t = "datetime2"
	case field.TypeEnum:
		// Currently, the support for enums is weak (application level only.
		// like SQLite). Dialect needs to create and maintain its enum type.
		// The default varchar is 1
		t = "varchar(255)"
	case field.TypeUUID:
		t = "char(36) binary"
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
	return t
}

// addColumn returns the DSL query for adding the given column to a table.
// The syntax/order is: datatype [Charset] [Unique|Increment] [Collation] [Nullable].
func (d *MSSQL) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Dialect(dialect.MSSQL).Column(c.Name).Type(d.cType(c)).Attr(c.Attr)

	// Unique is handled by indexes
	// c.unique(b)

	if c.Increment {
		b.Attr("IDENTITY(1,1)")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// addIndex returns the querying for adding an index to MSSQL.
func (d *MSSQL) addIndex(i *Index, table string) *sql.IndexBuilder {
	// return i.Builder(table)

	idx := sql.Dialect(dialect.MSSQL).
		CreateIndex(i.Name).Table(table)
	if i.Unique {
		idx.Unique()

		// Exclude nulls by default
		if len(i.Columns) == 1 {
			b := sql.Builder{}
			b.SetDialect(d.Dialect())

			b.Ident(i.Columns[0].Name)
			b.WriteString(" is not NULL")

			idx.Filter(b.String())
		}
	}
	for _, c := range i.Columns {
		idx.Column(c.Name)
	}
	return idx
}

// dropIndex drops a MSSQL index.
func (d *MSSQL) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	query, args := idx.DropBuilder(table).Query()
	return tx.Exec(ctx, query, args, nil)
}

// prepare runs preparation work that needs to be done to apply the change-set.
func (d *MSSQL) prepare(ctx context.Context, tx dialect.Tx, change *changes, table string) error {
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
			names, err := fkNames(ctx, tx, d.tableSchema(), table, col.Name)
			if err != nil {
				return err
			}
			if len(names) == 1 {
				qr = sql.Dialect(dialect.MSSQL).AlterTable(table).DropForeignKey(names[0])
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
			names, err := fkNames(ctx, tx, d.tableSchema(), table, col.Name)
			if err != nil {
				return err
			}
			if len(names) == 1 {
				qr = sql.Dialect(dialect.MSSQL).CreateIndex(names[0]).Table(table).Columns(col.Name)
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

var defaultsRe = regexp.MustCompile(`\(\((.+)\)\)`)

func (d *MSSQL) parseDefaults(defaults string) string {
	match := defaultsRe.FindStringSubmatch(defaults)

	if len(match) > 0 {
		return match[1]
	}

	return defaults
}

// scanColumn scans the column information from MSSQL column description.
func (d *MSSQL) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		nullable   sql.NullBool
		defaults   sql.NullString
		isIdentity sql.NullBool
		maxLength  int64
		precision  int64
		scale      int64
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &defaults, &isIdentity, &maxLength, &precision, &scale); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}

	if isIdentity.Valid && isIdentity.Bool {
		c.Increment = true
	}

	if defaults.Valid {
		defaults.String = d.parseDefaults(defaults.String)
	}

	c.Unique = c.UniqueKey()
	if nullable.Valid {
		c.Nullable = nullable.Bool
	}
	size := d.getColumnSize(c.typ, maxLength, precision, scale)

	switch c.typ {
	case "int":
		c.Type = field.TypeInt32
	case "smallint":
		c.Type = field.TypeInt16
	case "bigint":
		c.Type = field.TypeInt64
	case "tinyint":
		c.Type = field.TypeInt8
	case "bit":
		c.Type = field.TypeBool
	case "float":
		c.Type = field.TypeFloat64
		// TODO: use precision and scale to calculate correct type?
	case "timestamp", "datetime", "datetime2":
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
		c.Size = size
	case "varchar", "nvarchar":
		c.Type = field.TypeString
		c.Size = size
	case "longtext":
		c.Size = math.MaxInt32
		c.Type = field.TypeString
	case "json":
		c.Type = field.TypeJSON
	//case "enum":
	//	c.Type = field.TypeEnum
	//	c.Enums = make([]string, len(parts)-1)
	//	for i, e := range parts[1:] {
	//		c.Enums[i] = strings.Trim(e, "'")
	//	}
	case "char":
		c.Type = field.TypeString
	default:
		return fmt.Errorf("unknown column type %q for version %q", c.typ, d.version)
	}
	if defaults.Valid {
		return c.ScanDefault(defaults.String)
	}
	return nil
}

// scanIndexes scans sql.Rows into an Indexes list. The query for returning the rows,
// should return the following 4 columns: INDEX_NAME, COLUMN_NAME, NON_UNIQUE, SEQ_IN_INDEX.
// SEQ_IN_INDEX specifies the position of the column in the index columns.
func (d *MSSQL) scanIndexes(rows *sql.Rows) (Indexes, error) {
	var (
		i     Indexes
		names = make(map[string]*Index)
	)
	for rows.Next() {
		var (
			name     string
			column   string
			uniq     bool
			primary  bool
			seqindex int
		)
		if err := rows.Scan(&name, &column, &uniq, &primary, &seqindex); err != nil {
			return nil, fmt.Errorf("scanning index description: %v", err)
		}

		idx, ok := names[name]
		if !ok {
			idx = &Index{Name: name, Unique: uniq, primary: primary}
			i = append(i, idx)
			names[name] = idx
		}
		idx.columns = append(idx.columns, column)
	}
	return i, nil
}

// tableSchema returns the query for getting the table schema.
func (d *MSSQL) tableSchema() sql.Querier {
	return sql.Raw("(SELECT SCHEMA_NAME())")
}

// alterColumns returns the queries for applying the columns change-set.
func (d *MSSQL) alterColumns(table string, add, modify, drop []*Column) sql.Queries {
	b := sql.Dialect(dialect.MSSQL).AlterTable(table)
	for _, c := range add {
		b.AddColumn(d.addColumn(c))
	}
	for _, c := range modify {
		b.ModifyColumn(d.addColumn(c))
	}
	for _, c := range drop {
		b.DropColumn(sql.Dialect(dialect.MSSQL).Column(c.Name))
	}
	if len(b.Queries) == 0 {
		return nil
	}
	return sql.Queries{b}
}

// parseColumn returns column parts, size and signedness by mysql type
func (d MSSQL) getColumnSize(typ string, maxLength int64, precision int64, scale int64) int64 {
	var (
		size int64
	)

	switch typ {
	case "int", "smallint", "bigint", "tinyint":
		size = precision // todo should this be scale

	case "varbinary", "varchar", "char":
		size = maxLength
	}

	return size
}
