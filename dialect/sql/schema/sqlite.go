// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqlite"
)

type (
	// SQLite is an SQLite migration driver.
	SQLite struct {
		dialect.Driver
		WithForeignKeys bool
	}
	// SQLiteTx implements dialect.Tx.
	SQLiteTx struct {
		dialect.Tx
		commit   func() error // Override Commit to toggle foreign keys back on after Commit.
		rollback func() error // Override Rollback to toggle foreign keys back on after Rollback.
	}
)

// Tx implements opens a transaction.
func (d *SQLite) Tx(ctx context.Context) (dialect.Tx, error) {
	db := &db{d}
	if _, err := db.ExecContext(ctx, "PRAGMA foreign_keys = off"); err != nil {
		return nil, fmt.Errorf("sqlite: set 'foreign_keys = off': %w", err)
	}
	t, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	tx := &tx{t}
	cm, err := sqlite.CommitFunc(ctx, db, tx, true)
	if err != nil {
		return nil, err
	}
	return &SQLiteTx{Tx: t, commit: cm, rollback: sqlite.RollbackFunc(ctx, db, tx, true)}, nil
}

// Commit ensures foreign keys are toggled back on after commit.
func (tx *SQLiteTx) Commit() error {
	return tx.commit()
}

// Rollback ensures foreign keys are toggled back on after rollback.
func (tx *SQLiteTx) Rollback() error {
	return tx.rollback()
}

// init makes sure that foreign_keys support is enabled.
func (d *SQLite) init(ctx context.Context) error {
	on, err := exist(ctx, d, "PRAGMA foreign_keys")
	if err != nil {
		return fmt.Errorf("sqlite: check foreign_keys pragma: %w", err)
	}
	if !on {
		// foreign_keys pragma is off, either enable it by execute "PRAGMA foreign_keys=ON"
		// or add the following parameter in the connection string "_fk=1".
		return fmt.Errorf("sqlite: foreign_keys pragma is off: missing %q in the connection string", "_fk=1")
	}
	return nil
}

func (d *SQLite) tableExist(ctx context.Context, conn dialect.ExecQuerier, name string) (bool, error) {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_master")).
		Where(sql.And(
			sql.EQ("type", "table"),
			sql.EQ("name", name),
		)).
		Query()
	return exist(ctx, conn, query, args...)
}

// setRange sets the start value of table PK.
// SQLite tracks the AUTOINCREMENT in the "sqlite_sequence" table that is created and initialized automatically
// whenever a table that contains an AUTOINCREMENT column is created. However, it populates to it a rows (for tables)
// only after the first insertion. Therefore, we check. If a record (for the given table) already exists in the "sqlite_sequence"
// table, we updated it. Otherwise, we insert a new value.
func (d *SQLite) setRange(ctx context.Context, conn dialect.ExecQuerier, t *Table, value int64) error {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_sequence")).
		Where(sql.EQ("name", t.Name)).
		Query()
	exists, err := exist(ctx, conn, query, args...)
	switch {
	case err != nil:
		return err
	case exists:
		query, args = sql.Update("sqlite_sequence").Set("seq", value).Where(sql.EQ("name", t.Name)).Query()
	default: // !exists
		query, args = sql.Insert("sqlite_sequence").Columns("name", "seq").Values(t.Name, value).Query()
	}
	return conn.Exec(ctx, query, args, nil)
}

func (d *SQLite) tBuilder(t *Table) *sql.TableBuilder {
	b := sql.CreateTable(t.Name)
	for _, c := range t.Columns {
		b.Column(d.addColumn(c))
	}
	if t.Annotation != nil {
		addChecks(b, t.Annotation)
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
	case field.TypeOther:
		t = c.typ
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type, c.Name))
	}
	return t
}

// addColumn returns the DSL query for adding the given column to a table.
func (d *SQLite) addColumn(c *Column) *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(d.cType(c)).Attr(c.Attr)
	c.unique(b)
	if c.PrimaryKey() && c.Increment {
		b.Attr("PRIMARY KEY AUTOINCREMENT")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// addIndex returns the query for adding an index to SQLite.
func (d *SQLite) addIndex(i *Index, table string) *sql.IndexBuilder {
	return i.Builder(table).IfNotExists()
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
		return nil, fmt.Errorf("sqlite: reading table description %w", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, fmt.Errorf("sqlite: %w", err)
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
		return nil, fmt.Errorf("sqlite: closing rows %w", err)
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
		return nil, fmt.Errorf("reading table indexes %w", err)
	}
	defer rows.Close()
	var idx Indexes
	for rows.Next() {
		i := &Index{}
		origin := sql.NullString{}
		if err := rows.Scan(&i.Name, &i.Unique, &origin); err != nil {
			return nil, fmt.Errorf("scanning index description %w", err)
		}
		i.primary = origin.String == "pk"
		idx = append(idx, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("closing rows %w", err)
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
		return nil, fmt.Errorf("reading table indexes %w", err)
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
		return fmt.Errorf("scanning column description: %w", err)
	}
	c.Nullable = notnull.Int64 == 0
	if pk.Int64 > 0 {
		c.Key = PrimaryKey
	}
	if c.typ == "" {
		return fmt.Errorf("missing type information for column %q", c.Name)
	}
	parts, size, _, err := parseColumn(c.typ)
	if err != nil {
		return err
	}
	switch strings.ToLower(parts[0]) {
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
	case "varchar", "char", "text":
		c.Size = size
		c.Type = field.TypeString
	case "decimal", "numeric":
		c.Type = field.TypeOther
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

// tables returns the query for getting the in the schema.
func (d *SQLite) tables() sql.Querier {
	return sql.Select("name").
		From(sql.Table("sqlite_schema")).
		Where(sql.EQ("type", "table"))
}

// needsConversion reports if column "old" needs to be converted
// (by table altering) to column "new".
func (d *SQLite) needsConversion(old, new *Column) bool {
	c1, c2 := d.cType(old), d.cType(new)
	return c1 != c2 && old.typ != c2
}

// Atlas integration.

func (d *SQLite) atOpen(conn dialect.ExecQuerier) (migrate.Driver, error) {
	return sqlite.Open(&db{ExecQuerier: conn})
}

func (d *SQLite) atTable(t1 *Table, t2 *schema.Table) {
	if t1.Annotation != nil {
		setAtChecks(t1, t2)
	}
}

func (d *SQLite) supportsDefault(*Column) bool {
	// SQLite supports default values for all standard types.
	return true
}

func (d *SQLite) atTypeC(c1 *Column, c2 *schema.Column) error {
	if c1.SchemaType != nil && c1.SchemaType[dialect.SQLite] != "" {
		t, err := sqlite.ParseType(strings.ToLower(c1.SchemaType[dialect.SQLite]))
		if err != nil {
			return err
		}
		c2.Type.Type = t
		return nil
	}
	var t schema.Type
	switch c1.Type {
	case field.TypeBool:
		t = &schema.BoolType{T: "bool"}
	case field.TypeInt8, field.TypeUint8, field.TypeInt16, field.TypeUint16, field.TypeInt32,
		field.TypeUint32, field.TypeUint, field.TypeInt, field.TypeInt64, field.TypeUint64:
		t = &schema.IntegerType{T: sqlite.TypeInteger}
	case field.TypeBytes:
		t = &schema.BinaryType{T: sqlite.TypeBlob}
	case field.TypeString, field.TypeEnum:
		// SQLite does not impose any length restrictions on
		// the length of strings, BLOBs or numeric values.
		t = &schema.StringType{T: sqlite.TypeText}
	case field.TypeFloat32, field.TypeFloat64:
		t = &schema.FloatType{T: sqlite.TypeReal}
	case field.TypeTime:
		t = &schema.TimeType{T: "datetime"}
	case field.TypeJSON:
		t = &schema.JSONType{T: "json"}
	case field.TypeUUID:
		t = &sqlite.UUIDType{T: "uuid"}
	case field.TypeOther:
		t = &schema.UnsupportedType{T: c1.typ}
	default:
		t, err := sqlite.ParseType(strings.ToLower(c1.typ))
		if err != nil {
			return err
		}
		c2.Type.Type = t
	}
	c2.Type.Type = t
	return nil
}

func (d *SQLite) atUniqueC(t1 *Table, c1 *Column, t2 *schema.Table, c2 *schema.Column) {
	// For UNIQUE columns, SQLite create an implicit index named
	// "sqlite_autoindex_<table>_<i>". Ent uses the PostgreSQL approach
	// in its migration, and name these indexes as "<table>_<column>_key".
	for _, idx := range t1.Indexes {
		// Index also defined explicitly, and will be add in atIndexes.
		if idx.Unique && d.atImplicitIndexName(idx, t1, c1) {
			return
		}
	}
	t2.AddIndexes(schema.NewUniqueIndex(fmt.Sprintf("%s_%s_key", t2.Name, c1.Name)).AddColumns(c2))
}

func (d *SQLite) atImplicitIndexName(idx *Index, t1 *Table, c1 *Column) bool {
	if idx.Name == c1.Name {
		return true
	}
	p := fmt.Sprintf("sqlite_autoindex_%s_", t1.Name)
	if !strings.HasPrefix(idx.Name, p) {
		return false
	}
	i, err := strconv.ParseInt(strings.TrimPrefix(idx.Name, p), 10, 64)
	return err == nil && i > 0
}

func (d *SQLite) atIncrementC(t *schema.Table, c *schema.Column) {
	if c.Default != nil {
		t.Attrs = removeAttr(t.Attrs, reflect.TypeOf(&sqlite.AutoIncrement{}))
	} else {
		c.AddAttrs(&sqlite.AutoIncrement{})
	}
}

func (d *SQLite) atIncrementT(t *schema.Table, v int64) {
	t.AddAttrs(&sqlite.AutoIncrement{Seq: v})
}

func (d *SQLite) atIndex(idx1 *Index, t2 *schema.Table, idx2 *schema.Index) error {
	for _, c1 := range idx1.Columns {
		c2, ok := t2.Column(c1.Name)
		if !ok {
			return fmt.Errorf("unexpected index %q column: %q", idx1.Name, c1.Name)
		}
		idx2.AddParts(&schema.IndexPart{C: c2})
	}
	if idx1.Annotation != nil && idx1.Annotation.Where != "" {
		idx2.AddAttrs(&sqlite.IndexPredicate{P: idx1.Annotation.Where})
	}
	return nil
}

func (*SQLite) atTypeRangeSQL(ts ...string) string {
	for i := range ts {
		ts[i] = fmt.Sprintf("('%s')", ts[i])
	}
	return fmt.Sprintf("INSERT INTO `%s` (`type`) VALUES %s", TypeTable, strings.Join(ts, ", "))
}

type tx struct {
	dialect.Tx
}

func (tx *tx) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	return rows.ColumnScanner.(*stdsql.Rows), nil
}

func (tx *tx) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	var r stdsql.Result
	if err := tx.Exec(ctx, query, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}
