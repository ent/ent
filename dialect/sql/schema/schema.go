// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package schema contains all schema migration logic for SQL dialects.
package schema

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
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
	Schema      string
	Columns     []*Column
	columns     map[string]*Column
	Indexes     []*Index
	PrimaryKey  []*Column
	ForeignKeys []*ForeignKey
	Annotation  *entsql.Annotation
	Comment     string
}

// NewTable returns a new table with the given name.
func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		columns: make(map[string]*Column),
	}
}

// SetComment sets the table comment.
func (t *Table) SetComment(c string) *Table {
	t.Comment = c
	return t
}

// SetSchema sets the table schema.
func (t *Table) SetSchema(s string) *Table {
	t.Schema = s
	return t
}

// AddPrimary adds a new primary key to the table.
func (t *Table) AddPrimary(c *Column) *Table {
	c.Key = PrimaryKey
	t.AddColumn(c)
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

// HasColumn reports if the table contains a column with the given name.
func (t *Table) HasColumn(name string) bool {
	_, ok := t.Column(name)
	return ok
}

// Column returns the column with the given name. If exists.
func (t *Table) Column(name string) (*Column, bool) {
	if c, ok := t.columns[name]; ok {
		return c, true
	}
	// In case the column was added
	// directly to the Columns field.
	for _, c := range t.Columns {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}

// SetAnnotation the entsql.Annotation on the table.
func (t *Table) SetAnnotation(ant *entsql.Annotation) *Table {
	t.Annotation = ant
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

// Index returns a table index by its exact name.
func (t *Table) Index(name string) (*Index, bool) {
	idx, ok := t.index(name)
	if ok && idx.Name == name {
		return idx, ok
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

// hasIndex reports if the table has at least one index that matches the given names.
func (t *Table) hasIndex(names ...string) bool {
	for i := range names {
		if names[i] == "" {
			continue
		}
		if _, ok := t.index(names[i]); ok {
			return true
		}
	}
	return false
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

// CopyTables returns a deep-copy of the given tables. This utility function is
// useful for copying the generated schema tables (i.e. migrate.Tables) before
// running schema migration when there is a need for execute multiple migrations
// concurrently. e.g. running parallel unit-tests using the generated enttest package.
func CopyTables(tables []*Table) ([]*Table, error) {
	var (
		copyT  = make([]*Table, len(tables))
		byName = make(map[string]*Table)
	)
	for i, t := range tables {
		copyT[i] = &Table{
			Name:        t.Name,
			Columns:     make([]*Column, len(t.Columns)),
			Indexes:     make([]*Index, len(t.Indexes)),
			ForeignKeys: make([]*ForeignKey, len(t.ForeignKeys)),
		}
		for j, c := range t.Columns {
			cc := *c
			// SchemaType and Enums are read-only fields.
			cc.indexes = nil
			cc.foreign = nil
			copyT[i].Columns[j] = &cc
		}
		if at := t.Annotation; at != nil {
			cat := *at
			copyT[i].Annotation = &cat
		}
		byName[t.Name] = copyT[i]
	}
	for i, t := range tables {
		ct := copyT[i]
		for _, c := range t.PrimaryKey {
			cc, ok := ct.column(c.Name)
			if !ok {
				return nil, fmt.Errorf("sql/schema: missing primary key column %q", c.Name)
			}
			ct.PrimaryKey = append(ct.PrimaryKey, cc)
		}
		for j, idx := range t.Indexes {
			cidx := &Index{
				Name:    idx.Name,
				Unique:  idx.Unique,
				Columns: make([]*Column, len(idx.Columns)),
			}
			if at := idx.Annotation; at != nil {
				cat := *at
				cidx.Annotation = &cat
			}
			for k, c := range idx.Columns {
				cc, ok := ct.column(c.Name)
				if !ok {
					return nil, fmt.Errorf("sql/schema: missing index column %q", c.Name)
				}
				cidx.Columns[k] = cc
			}
			ct.Indexes[j] = cidx
		}
		for j, fk := range t.ForeignKeys {
			cfk := &ForeignKey{
				Symbol:     fk.Symbol,
				OnUpdate:   fk.OnUpdate,
				OnDelete:   fk.OnDelete,
				Columns:    make([]*Column, len(fk.Columns)),
				RefColumns: make([]*Column, len(fk.RefColumns)),
			}
			for k, c := range fk.Columns {
				cc, ok := ct.column(c.Name)
				if !ok {
					return nil, fmt.Errorf("sql/schema: missing foreign-key column %q", c.Name)
				}
				cfk.Columns[k] = cc
			}
			cref, ok := byName[fk.RefTable.Name]
			if !ok {
				return nil, fmt.Errorf("sql/schema: missing foreign-key ref-table %q", fk.RefTable.Name)
			}
			cfk.RefTable = cref
			for k, c := range fk.RefColumns {
				cc, ok := cref.column(c.Name)
				if !ok {
					return nil, fmt.Errorf("sql/schema: missing foreign-key ref-column %q", c.Name)
				}
				cfk.RefColumns[k] = cc
			}
			ct.ForeignKeys[j] = cfk
		}
	}
	return copyT, nil
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
	Default    any               // default value.
	Enums      []string          // enum values.
	Collation  string            // collation type (utf8mb4_unicode_ci, utf8mb4_general_ci)
	typ        string            // row column type (used for Rows.Scan).
	indexes    Indexes           // linked indexes.
	foreign    *ForeignKey       // linked foreign-key.
	Comment    string            // optional column comment.
}

// Expr represents a raw expression. It is used to distinguish between
// literal values and raw expressions when defining default values.
type Expr string

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
	case c.Type.Integer() && d.Type == field.TypeString:
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
			return fmt.Errorf("scanning int value for column %q: %w", c.Name, err)
		}
		c.Default = v.Int64
	case c.UintType():
		v := &sql.NullInt64{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning uint value for column %q: %w", c.Name, err)
		}
		c.Default = uint64(v.Int64)
	case c.FloatType():
		v := &sql.NullFloat64{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning float value for column %q: %w", c.Name, err)
		}
		c.Default = v.Float64
	case c.Type == field.TypeBool:
		v := &sql.NullBool{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning bool value for column %q: %w", c.Name, err)
		}
		c.Default = v.Bool
	case c.Type == field.TypeString || c.Type == field.TypeEnum:
		v := &sql.NullString{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning string value for column %q: %w", c.Name, err)
		}
		c.Default = v.String
	case c.Type == field.TypeJSON:
		v := &sql.NullString{}
		if err := v.Scan(value); err != nil {
			return fmt.Errorf("scanning json value for column %q: %w", c.Name, err)
		}
		c.Default = v.String
	case c.Type == field.TypeBytes:
		c.Default = []byte(value)
	case c.Type == field.TypeUUID:
		// skip function
		if !strings.Contains(value, "()") {
			c.Default = value
		}
	default:
		return fmt.Errorf("unsupported default type: %v default to %q", c.Type, value)
	}
	return nil
}

// defaultValue adds the `DEFAULT` attribute to the column.
// Note that, in SQLite if a NOT NULL constraint is specified,
// then the column must have a default value which not NULL.
func (c *Column) defaultValue(b *sql.ColumnBuilder) {
	if c.Default == nil || !c.supportDefault() {
		return
	}
	// Has default and the database supports adding this default.
	attr := fmt.Sprint(c.Default)
	switch v := c.Default.(type) {
	case bool:
		attr = strconv.FormatBool(v)
	case string:
		if t := c.Type; t != field.TypeUUID && t != field.TypeTime {
			// Escape single quote by replacing each with 2.
			attr = fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''"))
		}
	}
	b.Attr("DEFAULT " + attr)
}

// supportDefault reports if the column type supports default value.
func (c Column) supportDefault() bool {
	switch t := c.Type; t {
	case field.TypeString, field.TypeEnum:
		return c.Size < 1<<16 // not a text.
	case field.TypeBool, field.TypeTime, field.TypeUUID:
		return true
	default:
		return t.Numeric()
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

// nullable adds the `NULL`/`NOT NULL` attribute to the column if it exists in
// a different function to share the common declaration between the two dialects.
func (c *Column) nullable(b *sql.ColumnBuilder) {
	attr := Null
	if !c.Nullable {
		attr = "NOT " + attr
	}
	b.Attr(attr)
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

func (fk ForeignKey) column(name string) (*Column, bool) {
	for _, c := range fk.Columns {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}

func (fk ForeignKey) refColumn(name string) (*Column, bool) {
	for _, c := range fk.RefColumns {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
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
	return strings.ReplaceAll(strings.Title(strings.ToLower(string(r))), " ", "")
}

// Index definition for table index.
type Index struct {
	Name       string                  // index name.
	Unique     bool                    // uniqueness.
	Columns    []*Column               // actual table columns.
	Annotation *entsql.IndexAnnotation // index annotation.
	columns    []string                // columns loaded from query scan.
	primary    bool                    // primary key index.
	realname   string                  // real name in the database (Postgres only).
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

// addChecks appends the CHECK clauses from the entsql.Annotation.
func addChecks(t *sql.TableBuilder, ant *entsql.Annotation) {
	if check := ant.Check; check != "" {
		t.Checks(func(b *sql.Builder) {
			b.WriteString("CHECK " + checkExpr(check))
		})
	}
	if checks := ant.Checks; len(ant.Checks) > 0 {
		names := make([]string, 0, len(checks))
		for name := range checks {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			name := name
			t.Checks(func(b *sql.Builder) {
				b.WriteString("CONSTRAINT ").Ident(name).WriteString(" CHECK " + checkExpr(checks[name]))
			})
		}
	}
}

// checkExpr formats the CHECK expression.
func checkExpr(expr string) string {
	expr = strings.TrimSpace(expr)
	if !strings.HasPrefix(expr, "(") && !strings.HasSuffix(expr, ")") {
		expr = "(" + expr + ")"
	}
	return expr
}
