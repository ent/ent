// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/mysql"
	"ariga.io/atlas/sql/schema"
)

// MySQL adapter for Atlas migration engine.
type MySQL struct {
	dialect.Driver
	schema  string
	version string
}

// init loads the MySQL version from the database for later use in the migration process.
func (d *MySQL) init(ctx context.Context) error {
	if d.version != "" {
		return nil // already initialized.
	}
	rows := &sql.Rows{}
	if err := d.Query(ctx, "SHOW VARIABLES LIKE 'version'", []any{}, rows); err != nil {
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

func (d *MySQL) tableExist(ctx context.Context, conn dialect.ExecQuerier, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("TABLES").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("TABLE_NAME", name),
		)).Query()
	return exist(ctx, conn, query, args...)
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

func (d *MySQL) atOpen(conn dialect.ExecQuerier) (migrate.Driver, error) {
	return mysql.Open(&db{ExecQuerier: conn})
}

func (d *MySQL) atTable(t1 *Table, t2 *schema.Table) {
	t2.SetCharset("utf8mb4").SetCollation("utf8mb4_bin")
	if t1.Annotation == nil {
		return
	}
	if charset := t1.Annotation.Charset; charset != "" {
		t2.SetCharset(charset)
	}
	if collate := t1.Annotation.Collation; collate != "" {
		t2.SetCollation(collate)
	}
	if opts := t1.Annotation.Options; opts != "" {
		t2.AddAttrs(&mysql.CreateOptions{
			V: opts,
		})
	}
	// Check if the connected database supports the CHECK clause.
	// For MySQL, is >= "8.0.16" and for MariaDB it is "10.2.1".
	v1, v2 := d.version, "8.0.16"
	if v, ok := d.mariadb(); ok {
		v1, v2 = v, "10.2.1"
	}
	if compareVersions(v1, v2) >= 0 {
		setAtChecks(t1, t2)
	}
}

func (d *MySQL) supportsDefault(c *Column) bool {
	_, maria := d.mariadb()
	switch c.Default.(type) {
	case Expr, map[string]Expr:
		if maria {
			return compareVersions(d.version, "10.2.0") >= 0
		}
		return c.supportDefault() && compareVersions(d.version, "8.0.0") >= 0
	default:
		return c.supportDefault() || maria
	}
}

func (d *MySQL) supportsUUID() bool {
	_, maria := d.mariadb()
	return maria && compareVersions(d.version, "10.7.0") >= 0
}

func (d *MySQL) atTypeC(c1 *Column, c2 *schema.Column) error {
	if c1.SchemaType != nil && c1.SchemaType[dialect.MySQL] != "" {
		t, err := mysql.ParseType(strings.ToLower(c1.SchemaType[dialect.MySQL]))
		if err != nil {
			return err
		}
		c2.Type.Type = t
		return nil
	}
	var t schema.Type
	switch c1.Type {
	case field.TypeBool:
		t = &schema.BoolType{T: "boolean"}
	case field.TypeInt8:
		t = &schema.IntegerType{T: mysql.TypeTinyInt}
	case field.TypeUint8:
		t = &schema.IntegerType{T: mysql.TypeTinyInt, Unsigned: true}
	case field.TypeInt16:
		t = &schema.IntegerType{T: mysql.TypeSmallInt}
	case field.TypeUint16:
		t = &schema.IntegerType{T: mysql.TypeSmallInt, Unsigned: true}
	case field.TypeInt32:
		t = &schema.IntegerType{T: mysql.TypeInt}
	case field.TypeUint32:
		t = &schema.IntegerType{T: mysql.TypeInt, Unsigned: true}
	case field.TypeInt, field.TypeInt64:
		t = &schema.IntegerType{T: mysql.TypeBigInt}
	case field.TypeUint, field.TypeUint64:
		t = &schema.IntegerType{T: mysql.TypeBigInt, Unsigned: true}
	case field.TypeBytes:
		size := int64(math.MaxUint16)
		if c1.Size > 0 {
			size = c1.Size
		}
		switch {
		case size <= math.MaxUint8:
			t = &schema.BinaryType{T: mysql.TypeTinyBlob}
		case size <= math.MaxUint16:
			t = &schema.BinaryType{T: mysql.TypeBlob}
		case size < 1<<24:
			t = &schema.BinaryType{T: mysql.TypeMediumBlob}
		case size <= math.MaxUint32:
			t = &schema.BinaryType{T: mysql.TypeLongBlob}
		}
	case field.TypeJSON:
		t = &schema.JSONType{T: mysql.TypeJSON}
		if compareVersions(d.version, "5.7.8") == -1 {
			t = &schema.BinaryType{T: mysql.TypeLongBlob}
		}
	case field.TypeString:
		size := c1.Size
		if size == 0 {
			size = d.defaultSize(c1)
		}
		switch {
		case c1.typ == "tinytext", c1.typ == "text":
			t = &schema.StringType{T: c1.typ}
		case size <= math.MaxUint16:
			t = &schema.StringType{T: mysql.TypeVarchar, Size: int(size)}
		case size == 1<<24-1:
			t = &schema.StringType{T: mysql.TypeMediumText}
		default:
			t = &schema.StringType{T: mysql.TypeLongText}
		}
	case field.TypeFloat32, field.TypeFloat64:
		t = &schema.FloatType{T: c1.scanTypeOr(mysql.TypeDouble)}
	case field.TypeTime:
		t = &schema.TimeType{T: c1.scanTypeOr(mysql.TypeTimestamp)}
		// In MariaDB or in MySQL < v8.0.2, the TIMESTAMP column has both `DEFAULT CURRENT_TIMESTAMP`
		// and `ON UPDATE CURRENT_TIMESTAMP` if neither is specified explicitly. this behavior is
		// suppressed if the column is defined with a `DEFAULT` clause or with the `NULL` attribute.
		if _, maria := d.mariadb(); maria || compareVersions(d.version, "8.0.2") == -1 && c1.Default == nil {
			c2.SetNull(c1.Attr == "")
		}
	case field.TypeEnum:
		t = &schema.EnumType{T: mysql.TypeEnum, Values: c1.Enums}
	case field.TypeUUID:
		if d.supportsUUID() {
			// Native support for the uuid type
			t = &schema.UUIDType{T: mysql.TypeUUID}
		} else {
			// "CHAR(X) BINARY" is treated as "CHAR(X) COLLATE latin1_bin", and in MySQL < 8,
			// and "COLLATE utf8mb4_bin" in MySQL >= 8. However we already set the table to
			t = &schema.StringType{T: mysql.TypeChar, Size: 36}
			c2.SetCollation("utf8mb4_bin")
		}
	default:
		t, err := mysql.ParseType(strings.ToLower(c1.typ))
		if err != nil {
			return err
		}
		c2.Type.Type = t
	}
	c2.Type.Type = t
	return nil
}

func (d *MySQL) atUniqueC(t1 *Table, c1 *Column, t2 *schema.Table, c2 *schema.Column) {
	// For UNIQUE columns, MySQL create an implicit index
	// named as the column with an extra index in case the
	// name is already taken (<e.g. c>, <c_2>, <c_3>, ...).
	for _, idx := range t1.Indexes {
		// Index also defined explicitly, and will be add in atIndexes.
		if idx.Unique && d.atImplicitIndexName(idx, c1) {
			return
		}
	}
	t2.AddIndexes(schema.NewUniqueIndex(c1.Name).AddColumns(c2))
}

func (d *MySQL) atIncrementC(t *schema.Table, c *schema.Column) {
	if c.Default != nil {
		t.Attrs = removeAttr(t.Attrs, reflect.TypeOf(&mysql.AutoIncrement{}))
	} else {
		c.AddAttrs(&mysql.AutoIncrement{})
	}
}

func (d *MySQL) atIncrementT(t *schema.Table, v int64) {
	t.AddAttrs(&mysql.AutoIncrement{V: v})
}

func (d *MySQL) atImplicitIndexName(idx *Index, c1 *Column) bool {
	if idx.Name == c1.Name {
		return true
	}
	if !strings.HasPrefix(idx.Name, c1.Name+"_") {
		return false
	}
	i, err := strconv.ParseInt(strings.TrimLeft(idx.Name, c1.Name+"_"), 10, 64)
	return err == nil && i > 1
}

func (d *MySQL) atIndex(idx1 *Index, t2 *schema.Table, idx2 *schema.Index) error {
	prefix := indexParts(idx1)
	for _, c1 := range idx1.Columns {
		c2, ok := t2.Column(c1.Name)
		if !ok {
			return fmt.Errorf("unexpected index %q column: %q", idx1.Name, c1.Name)
		}
		part := &schema.IndexPart{C: c2}
		if v, ok := prefix[c1.Name]; ok {
			part.AddAttrs(&mysql.SubPart{Len: int(v)})
		}
		idx2.AddParts(part)
	}
	if t, ok := indexType(idx1, dialect.MySQL); ok {
		idx2.AddAttrs(&mysql.IndexType{T: t})
	}
	return nil
}

func (*MySQL) atTypeRangeSQL(ts ...string) string {
	for i := range ts {
		ts[i] = fmt.Sprintf("('%s')", ts[i])
	}
	return fmt.Sprintf("INSERT INTO `%s` (`type`) VALUES %s", TypeTable, strings.Join(ts, ", "))
}

// mariadb reports if the migration runs on MariaDB and returns the semver string.
func (d *MySQL) mariadb() (string, bool) {
	idx := strings.Index(d.version, "MariaDB")
	if idx == -1 {
		return "", false
	}
	return d.version[:idx-1], true
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
	case !c.Unique && len(c.indexes) == 0 && !c.PrimaryKey():
	default:
		size = 191
	}
	return size
}

// indexParts returns a map holding the sub_part mapping if exists.
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
