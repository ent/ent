package schema

import (
	"fmt"
	"strconv"
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

// MySQL returns the MySQL DSL query for table creation.
func (t *Table) MySQL(version string) *sql.TableBuilder {
	b := sql.CreateTable(t.Name).IfNotExists()
	for _, c := range t.Columns {
		b.Column(c.MySQL(version))
	}
	for _, pk := range t.PrimaryKey {
		b.PrimaryKey(pk.Name)
	}
	// default character set to MySQL table.
	// columns can be override using the "Charset" field.
	b.Charset("utf8mb4")
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
	}
	return nil, false
}

// Column schema definition for SQL dialects.
type Column struct {
	Name      string     // column name.
	Type      field.Type // column type.
	typ       string     // row column type (used for Rows.Scan).
	Attr      string     // extra attributes.
	Size      int        // max size parameter for string, blob, etc.
	Key       string     // key definition (PRI, UNI or MUL).
	Unique    bool       // column with unique constraint.
	Increment bool       // auto increment attribute.
	Nullable  *bool      // null or not null attribute.
	Default   string     // default value.
	Charset   string     // column character set.
	Collation string     // column collation.
}

// UniqueKey returns boolean indicates if this column is a unique key.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) UniqueKey() bool { return c.Key == "UNI" }

// PrimaryKey returns boolean indicates if this column is on of the primary key columns.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (c *Column) PrimaryKey() bool { return c.Key == "PRI" }

// MySQL returns the MySQL DSL query for table creation.
// The syntax/order is: datatype [Charset] [Unique|Increment] [Collation] [Nullable].
func (c *Column) MySQL(version string) *sql.ColumnBuilder {
	b := sql.Column(c.Name).Type(c.MySQLType(version)).Attr(c.Attr)
	if c.Charset != "" {
		b.Attr("CHARSET " + c.Charset)
	}
	c.unique(b)
	if c.Increment {
		b.Attr("AUTO_INCREMENT")
	}
	if c.Collation != "" {
		b.Attr("COLLATE " + c.Collation)
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
	case field.TypeString:
		size := c.Size
		if size == 0 {
			size = c.defaultSize(version)
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

// ScanMySQL scans the information from MySQL column description.
func (c *Column) ScanMySQL(rows *sql.Rows) error {
	var (
		charset  sql.NullString
		collate  sql.NullString
		nullable sql.NullString
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &nullable, &c.Key, &defaults, &c.Attr, &charset, &collate); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	c.Unique = c.UniqueKey()
	c.Charset = charset.String
	c.Default = defaults.String
	c.Collation = collate.String
	if nullable.Valid {
		null := nullable.String == "YES"
		c.Nullable = &null
	}
	switch parts := strings.FieldsFunc(c.typ, func(r rune) bool {
		return r == '(' || r == ')' || r == ' '
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
	case "timestamp":
		c.Type = field.TypeTime
	case "varchar":
		c.Type = field.TypeString
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("converting varchar size to int: %v", err)
		}
		c.Size = size
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

// defaultSize returns the default size for MySQL varchar
// type based on column size, charset and table indexes.
func (c *Column) defaultSize(version string) int {
	size := 255
	parts := strings.Split(version, ".")
	// non-unique or invalid version.
	if !c.Unique || len(parts) == 1 || parts[0] == "" || parts[1] == "" {
		return size
	}
	if major, minor := parts[0], parts[1]; major > "5" || minor > "6" {
		return size
	}
	return 191
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
	Name    string
	Unique  bool
	Columns []*Column
}

// Primary indicates if this index is a primary key.
// Used by the migration tool when parsing the `DESCRIBE TABLE` output Go objects.
func (i *Index) Primary() bool { return i.Name == "PRIMARY" }
