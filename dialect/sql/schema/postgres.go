// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
)

// Postgres is a postgres migration driver.
type Postgres struct {
	dialect.Driver
	version string
}

// init loads the Postgres version from the database for later use in the migration process.
// It returns an error if the server version is lower than v10.
func (d *Postgres) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW server_version_num", []interface{}{}, rows); err != nil {
		return fmt.Errorf("querying server version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("server_version_num variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("scanning version: %v", err)
	}
	if len(version) < 6 {
		return fmt.Errorf("malformed version: %s", version)
	}
	d.version = fmt.Sprintf("%s.%s.%s", version[:2], version[2:4], version[4:])
	if compareVersions(d.version, "10.0.0") == -1 {
		return fmt.Errorf("unsupported postgres version: %s", d.version)
	}
	return nil
}

// tableExist checks if a table exists in the database and current schema.
func (d *Postgres) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("table_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

// tableExist checks if a foreign-key exists in the current schema.
func (d *Postgres) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("constraint_type", "FOREIGN KEY").And().EQ("constraint_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

// setRange sets restart the identity column to the given offset. Used by the universal-id option.
func (d *Postgres) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	if value == 0 {
		value = 1 // RESTART value cannot be < 1.
	}
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE %s ALTER COLUMN id RESTART WITH %d", name, value), []interface{}{}, new(sql.Result))
}

// table loads the current table description from the database.
func (d *Postgres) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(dialect.Postgres).
		Select("column_name", "data_type", "is_nullable", "column_default").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("table_name", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("postgres: reading table description %v", err)
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
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("closing rows %v", err)
	}
	idxs, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	// Populate the index information to the table and its columns.
	// We do it manually, because PK and uniqueness information does
	// not exist when querying the INFORMATION_SCHEMA.COLUMNS above.
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
			c.indexes.append(idx)
			fallthrough
		default:
			t.AddIndex(idx.Name, idx.Unique, idx.columns)
		}
	}
	return t, nil
}

// indexesQuery holds a query format for retrieving
// table indexes of the current schema.
const indexesQuery = `
SELECT i.relname AS index_name,
       a.attname AS column_name,
       idx.indisprimary AS primary,
       idx.indisunique AS unique
FROM pg_class t,
     pg_class i,
     pg_index idx,
     pg_attribute a,
     pg_namespace n
WHERE t.oid = idx.indrelid
  AND i.oid = idx.indexrelid
  AND n.oid = t.relnamespace
  AND a.attrelid = t.oid
  AND a.attnum = ANY(idx.indkey)
  AND t.relkind = 'r'
  AND n.nspname = CURRENT_SCHEMA()
  AND t.relname = '%s';
`

func (d *Postgres) indexes(ctx context.Context, tx dialect.Tx, table string) (Indexes, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, fmt.Sprintf(indexesQuery, table), []interface{}{}, rows); err != nil {
		return nil, fmt.Errorf("querying indexes for table %s: %v", table, err)
	}
	defer rows.Close()
	var (
		idxs  Indexes
		names = make(map[string]*Index)
	)
	for rows.Next() {
		var (
			name, column    string
			unique, primary bool
		)
		if err := rows.Scan(&name, &column, &primary, &unique); err != nil {
			return nil, fmt.Errorf("scanning index description: %v", err)
		}
		// If the index is prefixed with the table, it's probably was
		// added by `addIndex` (and not entc) and it should be trimmed.
		name = strings.TrimPrefix(name, table+"_")
		idx, ok := names[name]
		if !ok {
			idx = &Index{Name: name, Unique: unique, primary: primary}
			idxs = append(idxs, idx)
			names[name] = idx
		}
		idx.columns = append(idx.columns, column)
	}
	return idxs, nil
}

// maxCharSize defines the maximum size of limited character types in Postgres (10 MB).
const maxCharSize = 10 << 20

// scanColumn scans the information a column from column description.
func (d *Postgres) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		nullable sql.NullString
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &defaults); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	switch c.typ {
	case "boolean":
		c.Type = field.TypeBool
	case "smallint":
		c.Type = field.TypeInt16
	case "integer":
		c.Type = field.TypeInt32
	case "bigint":
		c.Type = field.TypeInt64
	case "real":
		c.Type = field.TypeFloat32
	case "double precision":
		c.Type = field.TypeFloat64
	case "text":
		c.Type = field.TypeString
		c.Size = maxCharSize + 1
	case "character", "character varying":
		c.Type = field.TypeString
	case "timestamp with time zone":
		c.Type = field.TypeTime
	case "bytea":
		c.Type = field.TypeBytes
	case "jsonb":
		c.Type = field.TypeJSON
	case "uuid":
		c.Type = field.TypeUUID
	}
	switch {
	case !defaults.Valid:
		return nil
	case strings.Contains(defaults.String, "::"):
		parts := strings.Split(defaults.String, "::")
		defaults.String = strings.Trim(parts[0], "'")
		fallthrough
	default:
		return c.ScanDefault(defaults.String)
	}
}

// tBuilder returns the TableBuilder for the given table.
func (d *Postgres) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.Dialect(dialect.Postgres).
		CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(d.addColumn(c))
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	return b
}

// cType returns the PostgreSQL string type for this column.
func (d *Postgres) cType(c *Column) (t string) {
	switch c.Type {
	case field.TypeBool:
		t = "boolean"
	case field.TypeUint8, field.TypeInt8, field.TypeInt16, field.TypeUint16:
		t = "smallint"
	case field.TypeInt32, field.TypeUint32:
		t = "int"
	case field.TypeInt, field.TypeUint, field.TypeInt64, field.TypeUint64:
		t = "bigint"
	case field.TypeFloat32:
		t = "real"
	case field.TypeFloat64:
		t = "double precision"
	case field.TypeBytes:
		t = "bytea"
	case field.TypeJSON:
		t = "jsonb"
	case field.TypeUUID:
		t = "uuid"
	case field.TypeString:
		t = "varchar"
		if c.Size > maxCharSize {
			t = "text"
		}
	case field.TypeTime:
		t = "timestamp with time zone"
	case field.TypeEnum:
		// Currently, the support for enums is weak (application level only.
		// like SQLite). Dialect needs to create and maintain its enum type.
		t = "varchar"
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
	return t
}

// addColumn returns the ColumnBuilder for adding the given column to a table.
func (d *Postgres) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Dialect(dialect.Postgres).
		Column(c.Name).Type(d.cType(c)).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("GENERATED BY DEFAULT AS IDENTITY")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// alterColumn returns list of ColumnBuilder for applying in order to alter a column.
func (d *Postgres) alterColumn(c *Column) (ops []*sql.ColumnBuilder) {
	b := sql.Dialect(dialect.Postgres)
	ops = append(ops, b.Column(c.Name).Type(d.cType(c)))
	if c.Nullable {
		ops = append(ops, b.Column(c.Name).Attr("DROP NOT NULL"))
	} else {
		ops = append(ops, b.Column(c.Name).Attr("SET NOT NULL"))
	}
	return ops
}

// addIndex returns the querying for adding an index to PostgreSQL.
func (d *Postgres) addIndex(i *Index, table string) *sql.IndexBuilder {
	// Since index name should be unique in pg_class for schema,
	// we prefix it with the table name and remove on read.
	name := fmt.Sprintf("%s_%s", table, i.Name)
	idx := sql.Dialect(dialect.Postgres).
		CreateIndex(name).Table(table)
	if i.Unique {
		idx.Unique()
	}
	for _, c := range i.Columns {
		idx.Column(c.Name)
	}
	return idx
}

// dropIndex drops a Postgres index.
func (d *Postgres) dropIndex(ctx context.Context, tx dialect.Tx, idx *Index, table string) error {
	name := idx.Name
	build := sql.Dialect(dialect.Postgres)
	if prefix := table + "_"; !strings.HasPrefix(name, prefix) {
		name = prefix + name
	}
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("constraint_type", "UNIQUE").And().EQ("constraint_name", name)).
		Query()
	exists, err := exist(ctx, tx, query, args...)
	if err != nil {
		return err
	}
	query, args = build.DropIndex(name).Query()
	if exists {
		query, args = build.AlterTable(table).DropConstraint(name).Query()
	}
	return tx.Exec(ctx, query, args, new(sql.Result))
}
