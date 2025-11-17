// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package schema contains all schema migration logic for SQL dialects.
package schema

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/mysql"
	"ariga.io/atlas/sql/postgres"
	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqlite"
	entdialect "entgo.io/ent/dialect"
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
	View        bool   // Indicate the table is a view.
	Pos         string // filename:line of the ent schema definition.
}

// NewTable returns a new table with the given name.
func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		columns: make(map[string]*Column),
	}
}

// NewView returns a new view with the given name.
func NewView(name string) *Table {
	t := NewTable(name)
	t.View = true
	return t
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

// SetPos sets the table position.
func (t *Table) SetPos(p string) *Table {
	t.Pos = p
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
	realname   string                  // real name in the database (Postgres only).
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

func indexType(idx *Index, d string) (string, bool) {
	ant := idx.Annotation
	if ant == nil {
		return "", false
	}
	if ant.Types != nil && ant.Types[d] != "" {
		return ant.Types[d], true
	}
	if ant.Type != "" {
		return ant.Type, true
	}
	return "", false
}

type driver struct {
	sqlDialect
	schema.Differ
	migrate.PlanApplier
}

var drivers = func(v string) map[string]driver {
	return map[string]driver{
		entdialect.SQLite: {
			&SQLite{
				WithForeignKeys: true,
				Driver:          nopDriver{dialect: entdialect.SQLite},
			},
			sqlite.DefaultDiff,
			sqlite.DefaultPlan,
		},
		entdialect.MySQL: {
			&MySQL{
				version: v,
				Driver:  nopDriver{dialect: entdialect.MySQL},
			},
			mysql.DefaultDiff,
			mysql.DefaultPlan,
		},
		entdialect.Postgres: {
			&Postgres{
				version: v,
				Driver:  nopDriver{dialect: entdialect.Postgres},
			},
			postgres.DefaultDiff,
			postgres.DefaultPlan,
		},
	}
}

type DDLArgs struct {
	// Dialect and Version of the target database.
	Dialect, Version string
	// HashSymbols indicates whether to hash long symbols in the DDL.
	HashSymbols bool
	// Tables to dump.
	Tables []*Table
	// Options to pass to the migration plan engine.
	Options []migrate.PlanOption
}

// Dump the schema DDL for the given tables.
//
// Deprecated: use DDL instead.
func Dump(ctx context.Context, dialect, version string, tables []*Table, opts ...migrate.PlanOption) (string, error) {
	return DDL(ctx, DDLArgs{
		Dialect: dialect,
		Version: version,
		Tables:  tables,
		Options: opts,
	})
}

// DDL the schema DDL for the given tables.
func DDL(ctx context.Context, args DDLArgs) (string, error) {
	args.Options = append([]migrate.PlanOption{func(o *migrate.PlanOptions) {
		o.Mode = migrate.PlanModeDump
		o.Indent = "  "
	}}, args.Options...)
	d, ok := drivers(args.Version)[args.Dialect]
	if !ok {
		return "", fmt.Errorf("unsupported dialect %q", args.Dialect)
	}
	a := &Atlas{
		sqlDialect:  d,
		dialect:     args.Dialect,
		hashSymbols: args.HashSymbols,
	}
	r, err := a.StateReader(args.Tables...).ReadState(ctx)
	if err != nil {
		return "", err
	}
	// Since the Atlas version bundled with Ent does not support view management,
	// simply spit out the definition instead of letting Atlas plan them.
	var vs []*schema.View
	for _, s := range r.Schemas {
		vs = append(vs, s.Views...)
		s.Views = nil
	}
	var c schema.Changes
	if slices.ContainsFunc(args.Tables, func(t *Table) bool { return t.Schema != "" }) {
		c, err = d.RealmDiff(&schema.Realm{}, r)
	} else {
		c, err = d.SchemaDiff(&schema.Schema{}, r.Schemas[0])
	}
	if err != nil {
		return "", err
	}
	p, err := d.PlanChanges(ctx, "dump", c, args.Options...)
	if err != nil {
		return "", err
	}
	for _, v := range vs {
		q, _ := sql.Dialect(args.Dialect).
			CreateView(v.Name).
			Schema(v.Schema.Name).
			Columns(func(cols []*schema.Column) (bs []*sql.ColumnBuilder) {
				for _, c := range cols {
					bs = append(bs, sql.Dialect(args.Dialect).Column(c.Name).Type(c.Type.Raw))
				}
				return
			}(v.Columns)...).
			As(sql.Raw(v.Def)).
			Query()
		p.Changes = append(p.Changes, &migrate.Change{
			Cmd:     q,
			Comment: fmt.Sprintf("Add %q view", v.Name),
		})
	}
	for _, t := range args.Tables {
		p.Directives = append(p.Directives, fmt.Sprintf(
			"-- atlas:pos %s%s[type=%s] %s",
			func() string {
				if t.Schema != "" {
					return t.Schema + "[type=schema]."
				}
				return ""
			}(),
			t.Name,
			func() string {
				if t.View {
					return "view"
				}
				return "table"
			}(),
			t.Pos,
		))
	}
	f, err := migrate.DefaultFormatter.FormatFile(p)
	if err != nil {
		return "", err
	}
	return string(f.Bytes()), nil
}
