// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package schema contains all schema migration logic for SQL dialects.
package schema

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"
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
	return t.addIndex(&Index{
		Name:    name,
		Unique:  unique,
		columns: columns,
		Columns: make([]*Column, 0, len(columns)),
	})
}

// AddIndex creates and adds a new index to the table from the given options.
func (t *Table) addIndex(idx *Index) *Table {
	for _, name := range idx.columns {
		c, ok := t.columns[name]
		if ok {
			c.indexes.append(idx)
			idx.Columns = append(idx.Columns, c)
		}
	}
	t.Indexes = append(t.Indexes, idx)
	return t
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
func (t *Table) index(name string) (*Index, bool) {
	for _, idx := range t.Indexes {
		if name == idx.Name || name == idx.realname {
			return idx, true
		}
		// Same as below, there are cases where the index name
		// is unknown (created automatically on column constraint).
		if len(idx.Columns) == 1 && idx.Columns[0].Name == name {
			return idx, true
		}
	}
	// If it is an "implicit index" (unique constraint on
	// table creation) and it wasn't loaded in table scanning.
	c, ok := t.column(name)
	if !ok {
		// Postgres naming convention for unique constraint (<table>_<column>_key).
		name = strings.TrimPrefix(name, t.Name+"_")
		name = strings.TrimSuffix(name, "_key")
		c, ok = t.column(name)
	}
	if ok && c.Unique {
		return &Index{Name: name, Unique: c.Unique, Columns: []*Column{c}, columns: []string{c.Name}}, true
	}
	return nil, false
}

// fk returns a table foreign-key by its symbol.
// faster than map lookup for most cases.
func (t *Table) fk(symbol string) (*ForeignKey, bool) {
	for _, fk := range t.ForeignKeys {
		if fk.Symbol == symbol {
			return fk, true
		}
	}
	return nil, false
}

// Column schema definition for SQL dialects.
type Column struct {
	Name       string            // column name.
	Type       field.Type        // column type.
	SchemaType map[string]string // optional schema type per dialect.
	Attr       string            // extra attributes.
	Size       int64             // max size parameter for string, blob, etc.
	Key        string            // key definition (PRI, UNI or MUL).
	Unique     bool              // column with unique constraint.
	Increment  bool              // auto increment attribute.
	Nullable   bool              // null or not null attribute.
	Default    interface{}       // default value.
	Enums      []string          // enum values.
	typ        string            // row column type (used for Rows.Scan).
	indexes    Indexes           // linked indexes.
	foreign    *ForeignKey       // linked foreign-key.
}

// UniqueKey returns boolean indicates if this column is a unique key.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) UniqueKey() bool { return c.Key == UniqueKey }

// PrimaryKey returns boolean indicates if this column is on of the primary key columns.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) PrimaryKey() bool { return c.Key == PrimaryKey }

// ConvertibleTo reports whether a column can be converted to the new column without altering its data.
func (c *Column) ConvertibleTo(d *Column) bool {
	switch {
	case c.Type == d.Type:
		if c.Size != 0 && d.Size != 0 {
			// Types match and have a size constraint.
			return c.Size <= d.Size
		}
		return true
	case c.IntType() && d.IntType() || c.UintType() && d.UintType():
		return c.Type <= d.Type
	case c.UintType() && d.IntType():
		// uintX can not be converted to intY, when X > Y.
		return c.Type-field.TypeUint8 <= d.Type-field.TypeInt8
	case c.Type == field.TypeString && d.Type == field.TypeEnum ||
		c.Type == field.TypeEnum && d.Type == field.TypeString:
		return true
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
func (c *Column) ScanDefault(value string) error {
	switch {
	case strings.ToUpper(value) == Null: // ignore.
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
	case c.Type == field.TypeString || c.Type == field.TypeEnum:
		v := &sql.NullString{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning string value for column %q: %v", c.Name, err)
		}
		c.Default = v.String
	case c.Type == field.TypeJSON:
		v := &sql.NullString{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning json value for column %q: %v", c.Name, err)
		}
		c.Default = v.String
	default:
		return fmt.Errorf("unsupported default type: %v", c.Type)
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
			// Escape single quote by replacing each with 2.
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
	case c.Type == field.TypeString || c.Type == field.TypeEnum:
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

// scanTypeOr returns the scanning type or the given value.
func (c *Column) scanTypeOr(t string) string {
	if c.typ != "" {
		return strings.ToLower(c.typ)
	}
	return t
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
	Name     string    // index name.
	Unique   bool      // uniqueness.
	Columns  []*Column // actual table columns.
	columns  []string  // columns loaded from query scan.
	primary  bool      // primary key index.
	realname string    // real name in the database (Postgres only).
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

// sameAs reports if the index has the same properties
// as the given index (except the name).
func (i *Index) sameAs(idx *Index) bool {
	if i.Unique != idx.Unique || len(i.Columns) != len(idx.Columns) {
		return false
	}
	for j, c := range i.Columns {
		if c.Name != idx.Columns[j].Name {
			return false
		}
	}
	return true
}

// columnNames returns the names of the columns of the index.
func (i *Index) columnNames() []string {
	if len(i.columns) > 0 {
		return i.columns
	}
	columns := make([]string, 0, len(i.Columns))
	for _, c := range i.Columns {
		columns = append(columns, c.Name)
	}
	return columns
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
