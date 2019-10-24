// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package schema contains all schema migration logic for SQL dialects.
package schema

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
)

const (
	// DefaultStringLen describes the default length for string/varchar types.
	DefaultStringLen int64 = 255
	// Null is the string representation of NULL in SQL.
	Null = "NULL"
	// PrimaryKey is the string representation of PKs in SQL.
	PrimaryKey = "PRI"
	// UniqueKey is the string representation of PKs in SQL.
	UniqueKey = "UNI"
)

// Table schema definition for SQL dialects.
type Table struct {
	Name        string
	Columns     []*Column
	columns     map[string]*Column
	Indexes     []*Index
	PrimaryKey  []*Column
	ForeignKeys []*ForeignKey
}

// NewTable returns a new table with the given name.
func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		columns: make(map[string]*Column),
	}
}

// AddPrimary adds a new primary key to the table.
func (t *Table) AddPrimary(c *Column) *Table {
	t.Columns = append(t.Columns, c)
	t.PrimaryKey = append(t.PrimaryKey, c)
	return t
}

// AddForeignKey adds a foreign key to the table.
func (t *Table) AddForeignKey(fk *ForeignKey) *Table {
	t.ForeignKeys = append(t.ForeignKeys, fk)
	return t
}

// AddColumn adds a new column to the table.
func (t *Table) AddColumn(c *Column) *Table {
	t.columns[c.Name] = c
	t.Columns = append(t.Columns, c)
	return t
}

// AddIndex creates and adds a new index to the table from the given options.
func (t *Table) AddIndex(name string, unique bool, columns []string) *Table {
	idx := &Index{
		Name:    name,
		Unique:  unique,
		columns: columns,
		Columns: make([]*Column, len(columns)),
	}
	for i, name := range columns {
		c, ok := t.columns[name]
		if ok {
			idx.Columns[i] = c
			c.indexes = append(c.indexes, idx)
		}
	}
	t.Indexes = append(t.Indexes, idx)
	return t
}

// setup ensures the table is configured properly, like table columns
// are linked to their indexes, and PKs columns are defined.
func (t *Table) setup() {
	if t.columns == nil {
		t.columns = make(map[string]*Column, len(t.Columns))
	}
	for _, c := range t.Columns {
		t.columns[c.Name] = c
	}
	for _, idx := range t.Indexes {
		for _, c := range idx.Columns {
			c.indexes.append(idx)
		}
	}
	for _, pk := range t.PrimaryKey {
		c := t.columns[pk.Name]
		c.Key = PrimaryKey
		pk.Key = PrimaryKey
	}
}

// MySQL returns the MySQL DSL query for table creation.
func (t *Table) MySQL(version string) *sql.TableBuilder {
	b := sql.CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(c.MySQL(version))
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	// default charset / collation on MySQL table.
	// columns can be override using the Charset / Collate fields.
	b.Charset("utf8mb4").Collate("utf8mb4_bin")
	return b
}

// SQLite returns the SQLite query for table creation.
func (t *Table) SQLite() *sql.TableBuilder {
	b := sql.CreateTable(t.Name)
	for _, c := range t.Columns {
		b.Column(c.SQLite())
	}
	// Unlike in MySQL, we're not able to add foreign-key constraints to table
	// after it was created, and adding them to the `CREATE TABLE` statement is
	// not always valid (because circular foreign-keys situation is possible).
	// We stay consistent by not using constraints at all, and just defining the
	// foreign keys in the `CREATE TABLE` statement.
	for _, fk := range t.ForeignKeys {
		b.ForeignKeys(fk.DSL())
	}
	// if it's an ID based primary key, we add the `PRIMARY KEY`
	// clause to the column declaration.
	if len(t.PrimaryKey) == 1 {
		return b
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	return b
}

// column returns a table column by its name.
// faster than map lookup for most cases.
func (t *Table) column(name string) (*Column, bool) {
	for _, c := range t.Columns {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}

// index returns a table index by its name.
// faster than map lookup for most cases.
func (t *Table) index(name string) (*Index, bool) {
	for _, idx := range t.Indexes {
		if idx.Name == name {
			return idx, true
		}
		// Same as below, there are cases where the index name
		// is unknown (created automatically on column constraint).
		if len(idx.Columns) == 1 && idx.Columns[0].Name == name {
			return idx, true
		}
	}
	// If it is an "implicit index" (unique constraint on
	// table creation) and it didn't load on table scanning.
	if c, ok := t.column(name); ok && c.Unique {
		return &Index{Name: name, Unique: c.Unique, Columns: []*Column{c}, columns: []string{c.Name}}, true
	}
	return nil, false
}

// Column schema definition for SQL dialects.
type Column struct {
	Name      string      // column name.
	Type      field.Type  // column type.
	typ       string      // row column type (used for Rows.Scan).
	Attr      string      // extra attributes.
	Size      int64       // max size parameter for string, blob, etc.
	Key       string      // key definition (PRI, UNI or MUL).
	Unique    bool        // column with unique constraint.
	Increment bool        // auto increment attribute.
	Nullable  bool        // null or not null attribute.
	Default   interface{} // default value.
	indexes   Indexes     // linked indexes.
	Enums     []string    // enum values.
}

// UniqueKey returns boolean indicates if this column is a unique key.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) UniqueKey() bool { return c.Key == UniqueKey }

// PrimaryKey returns boolean indicates if this column is on of the primary key columns.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) PrimaryKey() bool { return c.Key == PrimaryKey }

// MySQL returns the MySQL DSL query for table creation.
// The syntax/order is: datatype [Charset] [Unique|Increment] [Collation] [Nullable].
func (c *Column) MySQL(version string) *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(c.MySQLType(version)).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("AUTO_INCREMENT")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// SQLite returns a SQLite DSL node for this column.
func (c *Column) SQLite() *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(c.SQLiteType()).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("PRIMARY KEY AUTOINCREMENT")
	}
	c.nullable(b)
	c.defaultValue(b)
	return b
}

// MySQLType returns the MySQL string type for this column.
func (c *Column) MySQLType(version string) (t string) {
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
		if compareVersions(version, "5.7.8") == -1 {
			t = "longblob"
		}
	case field.TypeString:
		size := c.Size
		if size == 0 {
			size = c.defaultSize(version)
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
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
	return t
}

// SQLiteType returns the SQLite string type for this column.
func (c *Column) SQLiteType() (t string) {
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
	default:
		panic("unsupported type " + c.Type.String())
	}
	return t
}

// ScanMySQL scans the information from MySQL column description.
func (c *Column) ScanMySQL(rows *sql.Rows) error {
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
	}
	if defaults.Valid {
		return c.ScanDefault(defaults.String)
	}
	return nil
}

// ConvertibleTo reports whether a column can be converted to the new column without altering its data.
func (c *Column) ConvertibleTo(d *Column) bool {
	switch {
	case c.Type == d.Type:
		return c.Size <= d.Size
	case c.IntType() && d.IntType() || c.UintType() && d.UintType():
		return c.Type <= d.Type
	case c.UintType() && d.IntType():
		// uintX can not be converted to intY, when X > Y.
		return c.Type-field.TypeUint8 <= d.Type-field.TypeInt8
	}
	return c.FloatType() && d.FloatType()
}

// IntType reports whether the column is an int type (int8 ... int64).
func (c Column) IntType() bool { return c.Type >= field.TypeInt8 && c.Type <= field.TypeInt64 }

// UintType reports of the given type is a uint type (int8 ... int64).
func (c Column) UintType() bool { return c.Type >= field.TypeUint8 && c.Type <= field.TypeUint64 }

// FloatType reports of the given type is a float type (float32, float64).
func (c Column) FloatType() bool { return c.Type == field.TypeFloat32 || c.Type == field.TypeFloat64 }

// ScanDefault scans the default value string to its interface type.
func (c *Column) ScanDefault(value string) (err error) {
	switch {
	case value == Null: // ignore.
	case c.IntType():
		v := &sql.NullInt64{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning int value for column %q: %v", c.Name, err)
		}
		c.Default = v.Int64
	case c.UintType():
		v := &sql.NullInt64{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning uint value for column %q: %v", c.Name, err)
		}
		c.Default = uint64(v.Int64)
	case c.FloatType():
		v := &sql.NullFloat64{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning float value for column %q: %v", c.Name, err)
		}
		c.Default = v.Float64
	case c.Type == field.TypeBool:
		v := &sql.NullBool{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning bool value for column %q: %v", c.Name, err)
		}
		c.Default = v.Bool
	case c.Type == field.TypeString:
		v := &sql.NullString{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning string value for column %q: %v", c.Name, err)
		}
		c.Default = v.String
	default:
		return fmt.Errorf("unsupported type: %v", c.Type)
	}
	return nil
}

// defaultValue adds tge `DEFAULT` attribute the the column.
// Note that, in SQLite if a NOT NULL constraint is specified,
// then the column must have a default value which not NULL.
func (c *Column) defaultValue(b *sql.ColumnBuilder) {
	// has default, and it's supported in the database level.
	if c.Default != nil && c.supportDefault() {
		attr := "DEFAULT "
		switch v := c.Default.(type) {
		case bool:
			attr += strconv.FormatBool(v)
		case string:
			// escape single quote by replacing each with 2.
			attr += fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''"))
		default:
			attr += fmt.Sprint(v)
		}
		b.Attr(attr)
	}
}

// supportDefault reports if the column type supports default value.
func (c Column) supportDefault() bool {
	switch {
	case c.Type == field.TypeString:
		return c.Size < 1<<16 // not a text.
	case c.Type.Numeric(), c.Type == field.TypeBool:
		return true
	default:
		return false
	}
}

// unique adds the `UNIQUE` attribute if the column is a unique type.
// it is exist in a different function to share the common declaration
// between the two dialects.
func (c *Column) unique(b *sql.ColumnBuilder) {
	if c.Unique {
		b.Attr("UNIQUE")
	}
}

// nullable adds the `NULL`/`NOT NULL` attribute to the column. it is exist in
// a different function to share the common declaration between the two dialects.
func (c *Column) nullable(b *sql.ColumnBuilder) {
	attr := Null
	if !c.Nullable {
		attr = "NOT " + attr
	}
	b.Attr(attr)
}

// defaultSize returns the default size for MySQL varchar type based
// on column size, charset and table indexes, in order to avoid index
// prefix key limit (767).
func (c *Column) defaultSize(version string) int64 {
	size := DefaultStringLen
	switch {
	// version is >= 5.7.
	case compareVersions(version, "5.7.0") != -1:
	// non-unique, or not part of any index (reaching the error 1071).
	case !c.Unique && len(c.indexes) == 0:
	default:
		size = 191
	}
	return size
}

// ForeignKey definition for creation.
type ForeignKey struct {
	Symbol     string          // foreign-key name. Generated if empty.
	Columns    []*Column       // table column
	RefTable   *Table          // referenced table.
	RefColumns []*Column       // referenced columns.
	OnUpdate   ReferenceOption // action on update.
	OnDelete   ReferenceOption // action on delete.
}

// DSL returns a default DSL query for a foreign-key.
func (fk ForeignKey) DSL() *sql.ForeignKeyBuilder {
	cols := make([]string, len(fk.Columns))
	refs := make([]string, len(fk.RefColumns))
	for i, c := range fk.Columns {
		cols[i] = c.Name
	}
	for i, c := range fk.RefColumns {
		refs[i] = c.Name
	}
	dsl := sql.ForeignKey().Symbol(fk.Symbol).
		Columns(cols...).
		Reference(sql.Reference().Table(fk.RefTable.Name).Columns(refs...))
	if action := string(fk.OnDelete); action != "" {
		dsl.OnDelete(action)
	}
	if action := string(fk.OnUpdate); action != "" {
		dsl.OnUpdate(action)
	}
	return dsl
}

// ReferenceOption for constraint actions.
type ReferenceOption string

// Reference options.
const (
	NoAction   ReferenceOption = "NO ACTION"
	Restrict   ReferenceOption = "RESTRICT"
	Cascade    ReferenceOption = "CASCADE"
	SetNull    ReferenceOption = "SET NULL"
	SetDefault ReferenceOption = "SET DEFAULT"
)

// ConstName returns the constant name of a reference option. It's used by entc for printing the constant name in templates.
func (r ReferenceOption) ConstName() string {
	if r == NoAction {
		return ""
	}
	return strings.ReplaceAll(strings.Title(strings.ToLower(string(r))), " ", "")
}

// Index definition for table index.
type Index struct {
	Name    string    // index name.
	Unique  bool      // uniqueness.
	Columns []*Column // actual table columns.
	columns []string  // columns loaded from query scan.
	primary bool      // primary key index.
}

// Builder returns the query builder for index creation. The DSL is identical in all dialects.
func (i *Index) Builder(table string) *sql.IndexBuilder {
	idx := sql.CreateIndex(i.Name).Table(table)
	if i.Unique {
		idx.Unique()
	}
	for _, c := range i.Columns {
		idx.Column(c.Name)
	}
	return idx
}

// DropBuilder returns the query builder for the drop index.
func (i *Index) DropBuilder(table string) *sql.DropIndexBuilder {
	idx := sql.DropIndex(i.Name).Table(table)
	return idx
}

// Indexes used for scanning all sql.Rows into a list of indexes, because
// multiple sql rows can represent the same index (multi-columns indexes).
type Indexes []*Index

// append wraps the basic `append` function by filtering duplicates indexes.
func (i *Indexes) append(idx1 *Index) {
	for _, idx2 := range *i {
		if idx2.Name == idx1.Name {
			return
		}
	}
	*i = append(*i, idx1)
}

// ScanMySQL scans sql.Rows into an Indexes list. The query for returning the rows,
// should return the following 4 columns: INDEX_NAME, COLUMN_NAME, NON_UNIQUE, SEQ_IN_INDEX.
// SEQ_IN_INDEX specifies the position of the column in the index columns.
func (i *Indexes) ScanMySQL(rows *sql.Rows) error {
	names := make(map[string]*Index)
	for rows.Next() {
		var (
			name     string
			column   string
			nonuniq  bool
			seqindex int
		)
		if err := rows.Scan(&name, &column, &nonuniq, &seqindex); err != nil {
			return fmt.Errorf("scanning index description: %v", err)
		}
		// ignore primary keys.
		if name == "PRIMARY" {
			continue
		}
		idx, ok := names[name]
		if !ok {
			idx = &Index{Name: name, Unique: !nonuniq}
			*i = append(*i, idx)
			names[name] = idx
		}
		idx.columns = append(idx.columns, column)
	}
	return nil
}

// compareVersions returns an integer comparing the 2 versions.
func compareVersions(v1, v2 string) int {
	pv1, ok1 := parseVersion(v1)
	pv2, ok2 := parseVersion(v2)
	if !ok1 && !ok2 {
		return 0
	}
	if !ok1 {
		return -1
	}
	if !ok2 {
		return 1
	}
	if v := compare(pv1.major, pv2.major); v != 0 {
		return v
	}
	if v := compare(pv1.minor, pv2.minor); v != 0 {
		return v
	}
	return compare(pv1.patch, pv2.patch)
}

// version represents a parsed MySQL version.
type version struct {
	major int
	minor int
	patch int
}

// parseVersion returns an integer comparing the 2 versions.
func parseVersion(v string) (*version, bool) {
	parts := strings.Split(v, ".")
	if len(parts) == 0 {
		return nil, false
	}
	var (
		err error
		ver = &version{}
	)
	for i, e := range []*int{&ver.major, &ver.minor, &ver.patch} {
		if i == len(parts) {
			break
		}
		if *e, err = strconv.Atoi(strings.Split(parts[i], "-")[0]); err != nil {
			return nil, false
		}
	}
	return ver, true
}

func compare(v1, v2 int) int {
	if v1 == v2 {
		return 0
	}
	if v1 < v2 {
		return -1
	}
	return 1
}
