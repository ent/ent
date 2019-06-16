package schema

import (
	"fmt"
	"strings"

	"fbc/ent/dialect/sql"
	"fbc/ent/field"
)

// Table schema definition for SQL dialects.
type Table struct {
	Name        string
	Columns     []*Column
	Indexes     []*Index
	PrimaryKey  []*Column
	ForeignKeys []*ForeignKey
}

// NewTable returns a new table with the given name.
func NewTable(name string) *Table { return &Table{Name: name} }

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

// DSL returns the default DSL query for table creation.
func (t *Table) DSL() *sql.TableBuilder {
	b := sql.CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(c.DSL())
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
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

// Column schema definition for SQL dialects.
type Column struct {
	Name      string     // column name.
	Type      field.Type // column type.
	Attr      string     // extra attributes.
	Default   string     // default value.
	Nullable  *bool      // null or not null attribute.
	Size      int        // max size parameter for string, blob, etc.
	Key       string     // key definition (PRI, UNI or MUL).
	Unique    bool       // column with unique constraint.
	Increment bool       // auto increment attribute.
}

// UniqueKey returns boolean indicates if this column is a unique key.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) UniqueKey() bool { return c.Key == "UNI" }

// PrimaryKey returns boolean indicates if this column is on of the primary key columns.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) PrimaryKey() bool { return c.Key == "PRI" }

// DSL returns the default DSL query for table creation.
func (c *Column) DSL() *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(c.MySQLType()).Attr(c.Attr)
	c.unique(b)
	if c.Increment {
		b.Attr("AUTO_INCREMENT")
	}
	c.nullable(b)
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
	return b
}

// MySQLType returns the MySQL string type for this column.
func (c *Column) MySQLType() (t string) {
	switch c.Type {
	case field.TypeBool:
		t = "boolean"
	case field.TypeInt8:
		t = "tinyint"
	case field.TypeUint8:
		t = "tinyint unsigned"
	case field.TypeInt64:
		t = "bigint"
	case field.TypeUint64:
		t = "bigint unsigned"
	case field.TypeInt, field.TypeInt16, field.TypeInt32:
		t = "int"
	case field.TypeUint, field.TypeUint16, field.TypeUint32:
		t = "int unsigned"
	case field.TypeString:
		size := c.Size
		if size == 0 {
			size = 255
		}
		if size < 1<<16 {
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
		if c.Nullable == nil {
			nullable := true
			c.Nullable = &nullable
		}
	default:
		panic("unsupported type " + c.Type.String())
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
	case field.TypeString:
		size := c.Size
		if size == 0 {
			size = 255
		}
		// sqlite has no size limit on varchar.
		t = fmt.Sprintf("varchar(%d)", size)
	case field.TypeFloat32, field.TypeFloat64:
		t = "real"
	case field.TypeTime:
		t = "datetime"
	default:
		panic("unsupported type " + c.Type.String())
	}
	return t
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
	if c.Nullable != nil {
		attr := "NULL"
		if !*c.Nullable {
			attr = "NOT " + attr
		}
		b.Attr(attr)
	}
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
	Key    string // key name.
	Column string // column name.
}

// Primary indicates if this index is a primary key.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (i *Index) Primary() bool { return i.Key == "PRIMARY" }
