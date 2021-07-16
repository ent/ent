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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

// MySQL is a MySQL migration driver.
type MySQL struct {
	dialect.Driver
	schema  string
	version string
}

// init loads the MySQL version from the database for later use in the migration process.
func (d *MySQL) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW VARIABLES LIKE 'version'", []interface{}{}, rows); err != nil {
		return fmt.Errorf("mysql: querying mysql version %w", err)
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
		return fmt.Errorf("mysql: scanning mysql version: %w", err)
	}
	d.version = version[1]
	return nil
}

func (d *MySQL) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("TABLES").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("TABLE_NAME", name),
		)).Query()
	return exist(ctx, tx, query, args...)
}

func (d *MySQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("TABLE_CONSTRAINTS").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("CONSTRAINT_TYPE", "FOREIGN KEY"),
			sql.EQ("CONSTRAINT_NAME", name),
		)).Query()
	return exist(ctx, tx, query, args...)
}

// table loads the current table description from the database.
func (d *MySQL) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Select(
		"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name",
		"numeric_precision", "numeric_scale",
	).
		From(sql.Table("COLUMNS").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("TABLE_NAME", name)),
		).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading table description %w", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, fmt.Errorf("mysql: %w", err)
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
		return nil, fmt.Errorf("mysql: closing rows %w", err)
	}
	indexes, err := d.indexes(ctx, tx, name)
	if err != nil {
		return nil, err
	}
	// Add and link indexes to table columns.
	for _, idx := range indexes {
		t.addIndex(idx)
	}
	if _, ok := d.mariadb(); ok {
		if err := d.normalizeJSON(ctx, tx, t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

// table loads the table indexes from the database.
func (d *MySQL) indexes(ctx context.Context, tx dialect.Tx, name string) ([]*Index, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("index_name", "column_name", "sub_part", "non_unique", "seq_in_index").
		From(sql.Table("STATISTICS").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("TABLE_NAME", name),
		)).
		OrderBy("index_name", "seq_in_index").
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading index description %w", err)
	}
	defer rows.Close()
	idx, err := d.scanIndexes(rows)
	if err != nil {
		return nil, fmt.Errorf("mysql: %w", err)
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
		From(sql.Table("TABLES").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("TABLE_NAME", t.Name),
		)).
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return fmt.Errorf("mysql: query auto_increment %w", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	actual := &sql.NullInt64{}
	if err := sql.ScanOne(rows, actual); err != nil {
		return fmt.Errorf("mysql: scan auto_increment %w", err)
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
	// Charset and collation config on MySQL table.
	// These options can be overridden by the entsql annotation.
	b.Charset("utf8mb4").Collate("utf8mb4_bin")
	if t.Annotation != nil {
		if charset := t.Annotation.Charset; charset != "" {
			b.Charset(charset)
		}
		if collate := t.Annotation.Collation; collate != "" {
			b.Collate(collate)
		}
		if opts := t.Annotation.Options; opts != "" {
			b.Options(opts)
		}
		addChecks(b, t.Annotation)
	}
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
			size = d.defaultSize(c)
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
		// In MySQL < v8.0.2, the TIMESTAMP column has both `DEFAULT CURRENT_TIMESTAMP` and
		// `ON UPDATE CURRENT_TIMESTAMP` if neither is specified explicitly. this behavior is
		// suppressed if the column is defined with a `DEFAULT` clause or with the `NULL` attribute.
		if compareVersions(d.version, "8.0.2") == -1 && c.Default == nil {
			c.Nullable = c.Attr == ""
		}
	case field.TypeEnum:
		values := make([]string, len(c.Enums))
		for i, e := range c.Enums {
			values[i] = fmt.Sprintf("'%s'", e)
		}
		t = fmt.Sprintf("enum(%s)", strings.Join(values, ", "))
	case field.TypeUUID:
		t = "char(36) binary"
	case field.TypeOther:
		t = c.typ
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
	if c.Collation != "" {
		b.Attr("COLLATE " + c.Collation)
	}
	if c.Type == field.TypeJSON {
		// Manually add a `CHECK` clause for older versions of MariaDB for validating the
		// JSON documents. This constraint is automatically included from version 10.4.3.
		if version, ok := d.mariadb(); ok && compareVersions(version, "10.4.3") == -1 {
			b.Check(func(b *sql.Builder) {
				b.WriteString("JSON_VALID(").Ident(c.Name).WriteByte(')')
			})
		}
	}
	return b
}

// addIndex returns the querying for adding an index to MySQL.
func (d *MySQL) addIndex(i *Index, table string) *sql.IndexBuilder {
	idx := sql.CreateIndex(i.Name).Table(table)
	if i.Unique {
		idx.Unique()
	}
	parts := indexParts(i)
	for _, c := range i.Columns {
		part, ok := parts[c.Name]
		if !ok || part == 0 {
			idx.Column(c.Name)
		} else {
			idx.Column(fmt.Sprintf("%s(%d)", idx.Builder.Quote(c.Name), part))
		}
	}
	return idx
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
			names, err := d.fkNames(ctx, tx, table, col.Name)
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
			names, err := d.fkNames(ctx, tx, table, col.Name)
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
		nullable         sql.NullString
		defaults         sql.NullString
		numericPrecision sql.NullInt64
		numericScale     sql.NullInt64
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &c.Key, &defaults, &c.Attr, &sql.NullString{}, &sql.NullString{}, &numericPrecision, &numericScale); err != nil {
		return fmt.Errorf("scanning column description: %w", err)
	}
	c.Unique = c.UniqueKey()
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	if c.typ == "" {
		return fmt.Errorf("missing type information for column %q", c.Name)
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
	case "double":
		c.Type = field.TypeFloat64
	case "numeric", "decimal":
		c.Type = field.TypeFloat64
		// If precision is specified then we should take that into account.
		if numericPrecision.Valid {
			schemaType := fmt.Sprintf("%s(%d,%d)", parts[0], numericPrecision.Int64, numericScale.Int64)
			c.SchemaType = map[string]string{dialect.MySQL: schemaType}
		}
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
	case "binary", "varbinary":
		c.Type = field.TypeBytes
		c.Size = size
	case "varchar":
		c.Type = field.TypeString
		c.Size = size
	case "text":
		c.Size = math.MaxUint16
		c.Type = field.TypeString
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
		c.Type = field.TypeOther
		// UUID field has length of 36 characters (32 alphanumeric characters and 4 hyphens).
		if size == 36 {
			c.Type = field.TypeUUID
		}
	case "point", "geometry", "linestring", "polygon":
		c.Type = field.TypeOther
	default:
		return fmt.Errorf("unknown column type %q for version %q", parts[0], d.version)
	}
	if defaults.Valid {
		return c.ScanDefault(defaults.String)
	}
	return nil
}

// scanIndexes scans sql.Rows into an Indexes list. The query for returning the rows,
// should return the following 5 columns: INDEX_NAME, COLUMN_NAME, SUB_PART, NON_UNIQUE,
// SEQ_IN_INDEX. SEQ_IN_INDEX specifies the position of the column in the index columns.
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
			subpart  sql.NullInt64
		)
		if err := rows.Scan(&name, &column, &subpart, &nonuniq, &seqindex); err != nil {
			return nil, fmt.Errorf("scanning index description: %w", err)
		}
		// Ignore primary keys.
		if name == "PRIMARY" {
			continue
		}
		idx, ok := names[name]
		if !ok {
			idx = &Index{Name: name, Unique: !nonuniq, Annotation: &entsql.IndexAnnotation{}}
			i = append(i, idx)
			names[name] = idx
		}
		idx.columns = append(idx.columns, column)
		if subpart.Int64 > 0 {
			if idx.Annotation.PrefixColumns == nil {
				idx.Annotation.PrefixColumns = make(map[string]uint)
			}
			idx.Annotation.PrefixColumns[column] = uint(subpart.Int64)
		}
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

// matchSchema returns the predicate for matching table schema.
func (d *MySQL) matchSchema(columns ...string) *sql.Predicate {
	column := "TABLE_SCHEMA"
	if len(columns) > 0 {
		column = columns[0]
	}
	if d.schema != "" {
		return sql.EQ(column, d.schema)
	}
	return sql.EQ(column, sql.Raw("(SELECT DATABASE())"))
}

// tables returns the query for getting the in the schema.
func (d *MySQL) tables() sql.Querier {
	return sql.Select("TABLE_NAME").
		From(sql.Table("TABLES").Schema("INFORMATION_SCHEMA")).
		Where(d.matchSchema())
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

// normalizeJSON normalize MariaDB longtext columns to type JSON.
func (d *MySQL) normalizeJSON(ctx context.Context, tx dialect.Tx, t *Table) error {
	columns := make(map[string]*Column)
	for _, c := range t.Columns {
		if c.typ == "longtext" {
			columns[c.Name] = c
		}
	}
	if len(columns) == 0 {
		return nil
	}
	rows := &sql.Rows{}
	query, args := sql.Select("CONSTRAINT_NAME").
		From(sql.Table("CHECK_CONSTRAINTS").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema("CONSTRAINT_SCHEMA"),
			sql.EQ("TABLE_NAME", t.Name),
			sql.Like("CHECK_CLAUSE", "json_valid(%)"),
		)).
		Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return fmt.Errorf("mysql: query table constraints %w", err)
	}
	// Call Close in cases of failures (Close is idempotent).
	defer rows.Close()
	names := make([]string, 0, len(columns))
	if err := sql.ScanSlice(rows, &names); err != nil {
		return fmt.Errorf("mysql: scan table constraints: %w", err)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if err := rows.Close(); err != nil {
		return err
	}
	for _, name := range names {
		c, ok := columns[name]
		if ok {
			c.Type = field.TypeJSON
		}
	}
	return nil
}

// mariadb reports if the migration runs on MariaDB and returns the semver string.
func (d *MySQL) mariadb() (string, bool) {
	idx := strings.Index(d.version, "MariaDB")
	if idx == -1 {
		return "", false
	}
	return d.version[:idx-1], true
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
	case "varbinary", "varchar", "char", "binary":
		if len(parts) > 1 {
			size, err = strconv.ParseInt(parts[1], 10, 64)
		}
	}
	if err != nil {
		return parts, size, unsigned, fmt.Errorf("converting %s size to int: %w", parts[0], err)
	}
	return parts, size, unsigned, nil
}

// fkNames returns the foreign-key names of a column.
func (d *MySQL) fkNames(ctx context.Context, tx dialect.Tx, table, column string) ([]string, error) {
	query, args := sql.Select("CONSTRAINT_NAME").From(sql.Table("KEY_COLUMN_USAGE").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			sql.EQ("TABLE_NAME", table),
			sql.EQ("COLUMN_NAME", column),
			// NULL for unique and primary-key constraints.
			sql.NotNull("POSITION_IN_UNIQUE_CONSTRAINT"),
			d.matchSchema(),
		)).
		Query()
	var (
		names []string
		rows  = &sql.Rows{}
	)
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("mysql: reading constraint names %w", err)
	}
	defer rows.Close()
	if err := sql.ScanSlice(rows, &names); err != nil {
		return nil, err
	}
	return names, nil
}

// defaultSize returns the default size for MySQL/MariaDB varchar type
// based on column size, charset and table indexes, in order to avoid
// index prefix key limit (767) for older versions of MySQL/MariaDB.
func (d *MySQL) defaultSize(c *Column) int64 {
	size := DefaultStringLen
	version, checked := d.version, "5.7.0"
	if v, ok := d.mariadb(); ok {
		version, checked = v, "10.2.2"
	}
	switch {
	// Version is >= 5.7 for MySQL, or >= 10.2.2 for MariaDB.
	case compareVersions(version, checked) != -1:
	// Column is non-unique, or not part of any index (reaching
	// the error 1071).
	case !c.Unique && len(c.indexes) == 0:
	default:
		size = 191
	}
	return size
}

// needsConversion reports if column "old" needs to be converted
// (by table altering) to column "new".
func (d *MySQL) needsConversion(old, new *Column) bool {
	return d.cType(old) != d.cType(new)
}

// indexModified used by the migration differ to check if the index was modified.
func (d *MySQL) indexModified(old, new *Index) bool {
	oldParts, newParts := indexParts(old), indexParts(new)
	if len(oldParts) != len(newParts) {
		return true
	}
	for column, oldPart := range oldParts {
		newPart, ok := newParts[column]
		if !ok || oldPart != newPart {
			return true
		}
	}
	return false
}

// indexParts returns a map holding the sub_part mapping if exist.
func indexParts(idx *Index) map[string]uint {
	parts := make(map[string]uint)
	if idx.Annotation == nil {
		return parts
	}
	// If prefix (without a name) was defined on the
	// annotation, map it to the single column index.
	if idx.Annotation.Prefix > 0 && len(idx.Columns) == 1 {
		parts[idx.Columns[0].Name] = idx.Annotation.Prefix
	}
	for column, part := range idx.Annotation.PrefixColumns {
		parts[column] = part
	}
	return parts
}
