// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"fmt"
	"go/token"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/entc/load"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/olekukonko/tablewriter"
)

type (
	// Type represents one node/type in the graph, its relations and the information it holds.
	Type struct {
		Config
		// Name holds the type/ent name.
		Name string
		// ID holds the ID field of this type.
		ID *Field
		// Fields holds all the primitive fields of this type.
		Fields []*Field
		// fields map for fast lookup.
		fields map[string]*Field
		// Edge holds all the edges of this type.
		Edges []*Edge
		// Indexes are the configured indexes for this type.
		Indexes []*Index
	}

	// Field holds the information of a type field used for the templates.
	Field struct {
		// field definition.
		def *load.Field
		// Name is the name of this field in the database schema.
		Name string
		// Type holds the type information of the field.
		Type field.Type
		// Unique indicate if this field is a unique field.
		Unique bool
		// Optional indicates is this field is optional on create.
		Optional bool
		// Nillable indicates that this field can be null in the
		// database and pointer in the generated entities.
		Nillable bool
		// Default indicates if this field has a default value for creation.
		Default bool
		// UpdateDefault indicates if this field has a default value for update.
		UpdateDefault bool
		// Immutable indicates is this field cannot be updated.
		Immutable bool
		// StructTag of the field. default to "json".
		StructTag string
		// Validators holds the number of validators this field have.
		Validators int
	}

	// Edge of a graph between two types.
	Edge struct {
		// Name holds the name of the edge.
		Name string
		// Type holds a reference to the type this edge is directed to.
		Type *Type
		// Optional indicates is this edge is optional on create.
		Optional bool
		// Unique indicates if this edge is a unique edge.
		Unique bool
		// Inverse holds the name of the inverse edge.
		Inverse string
		// Owner holds the type of the edge-owner. For assoc-edges it's the
		// type that holds the edge, for inverse-edges, it's the assoc type.
		Owner *Type
		// StructTag of the edge-field in the struct. default to "json".
		StructTag string
		// Relation holds the relation info of an edge.
		Rel Relation
		// SelfRef indicates if this edge is a self-reference to the same
		// type with the same name. For example, a User type have one of
		// following edges:
		//
		//	edge.To("friends", User.Type)			// many 2 many.
		//	edge.To("spouse", User.Type).Unique()	// one 2 one.
		//
		SelfRef bool
	}

	// Relation holds the relational database information for edges.
	Relation struct {
		// Type holds the relation type of the edge.
		Type Rel
		// Table holds the relation table for this edge.
		// For O2O and O2M, it's the table name of the type we're this edge point to.
		// For M2O, this is the owner's type, and for M2M this is the join table.
		Table string
		// Columns holds the relation column in the relation table above.
		// In O2M, M2O and O2O, this the first element.
		Columns []string
	}

	// Index represents a database index used for either increasing speed
	// on database operations or defining constraints such as "UNIQUE INDEX".
	// Note that some indexes are created implicitly like table foreign keys.
	Index struct {
		// Name of the index. One column index is simply the column name.
		Name string
		// Unique index or not.
		Unique bool
		// Columns are the table columns.
		Columns []string
	}
)

// NewType creates a new type and its fields from the given schema.
func NewType(c Config, schema *load.Schema) (*Type, error) {
	typ := &Type{
		Config: c,
		Name:   schema.Name,
		ID: &Field{
			Name:      "id",
			Type:      c.IDType,
			StructTag: `json:"id,omitempty"`,
		},
		Fields: make([]*Field, len(schema.Fields)),
		fields: make(map[string]*Field, len(schema.Fields)),
	}
	for i, f := range schema.Fields {
		switch {
		case !f.Type.Valid():
			return nil, fmt.Errorf("invalid type for field %s", f.Name)
		case f.Nillable && !f.Optional:
			return nil, fmt.Errorf("nillable field %q must be optional", f.Name)
		case f.Unique && f.Default:
			return nil, fmt.Errorf("unique field %q cannot have default value", f.Name)
		case typ.fields[f.Name] != nil:
			return nil, fmt.Errorf("field %q redeclared for type %q", f.Name, typ.Name)
		}
		typ.Fields[i] = &Field{
			def:           f,
			Name:          f.Name,
			Type:          f.Type,
			Unique:        f.Unique,
			Nillable:      f.Nillable,
			Optional:      f.Optional,
			Default:       f.Default,
			UpdateDefault: f.UpdateDefault,
			Immutable:     f.Immutable,
			StructTag:     structTag(f.Name, f.Tag),
			Validators:    f.Validators,
		}
		typ.fields[f.Name] = typ.Fields[i]
	}
	return typ, nil
}

// Label returns Gremlin label name of the node/type.
func (t Type) Label() string { return snake(t.Name) }

// Table returns SQL table name of the node/type.
func (t Type) Table() string { return snake(rules.Pluralize(t.Name)) }

// Package returns the package name of this node.
func (t Type) Package() string { return strings.ToLower(t.Name) }

// Receiver returns the receiver name of this node. It makes sure the
// receiver names doesn't conflict with import names.
func (t Type) Receiver() string {
	return receiver(t.Name)
}

// HasAssoc returns true if this type has an assoc edge with the given name.
// faster than map access for most cases.
func (t Type) HasAssoc(name string) (*Edge, bool) {
	for _, e := range t.Edges {
		if name == e.Name {
			return e, true
		}
	}
	return nil, false
}

// HasValidators indicates if any of the type's field has validators.
func (t Type) HasValidators() bool {
	for _, f := range t.Fields {
		if f.Validators > 0 {
			return true
		}
	}
	return false
}

// HasDefault indicates if any of this type's fields has default value on creation.
func (t Type) HasDefault() bool {
	for _, f := range t.Fields {
		if f.Default {
			return true
		}
	}
	return false
}

// HasUpdateDefault indicates if any of this type's fields has default value on update.
func (t Type) HasUpdateDefault() bool {
	for _, f := range t.Fields {
		if f.Default {
			return true
		}
	}
	return false
}

// NumConstraint returns the type's constraint count. Used for slice allocation.
func (t Type) NumConstraint() int {
	var n int
	for _, f := range t.Fields {
		if f.Unique {
			n++
		}
	}
	for _, e := range t.Edges {
		if e.HasConstraint() {
			n++
		}
	}
	return n
}

// MutableFields returns the types's mutable fields.
func (t Type) MutableFields() []*Field {
	var fields []*Field
	for _, f := range t.Fields {
		if !f.Immutable {
			fields = append(fields, f)
		}
	}
	return fields
}

// NumM2M returns the type's many-to-many edge count
func (t Type) NumM2M() int {
	var n int
	for _, e := range t.Edges {
		if e.M2M() {
			n++
		}
	}
	return n
}

// Describe returns description of a type. The format of the description is:
//
//	Type:
//			<Fields Table>
//
//			<Edges Table>
//
func (t Type) Describe(w io.Writer) {
	b := &strings.Builder{}
	b.WriteString(t.Name + ":\n")
	table := tablewriter.NewWriter(b)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Field", "Type", "Unique", "Optional", "Nillable", "Default", "UpdateDefault", "Immutable", "StructTag", "Validators"})
	for _, f := range append([]*Field{t.ID}, t.Fields...) {
		v := reflect.ValueOf(*f)
		row := make([]string, v.NumField()-1)
		for i := range row {
			row[i] = fmt.Sprint(v.Field(i + 1).Interface())
		}
		table.Append(row)
	}
	table.Render()
	table = tablewriter.NewWriter(b)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Edge", "Type", "Inverse", "BackRef", "Relation", "Unique", "Optional"})
	for _, e := range t.Edges {
		table.Append([]string{
			e.Name,
			e.Type.Name,
			strconv.FormatBool(e.IsInverse()),
			e.Inverse,
			e.Rel.Type.String(),
			strconv.FormatBool(e.Unique),
			strconv.FormatBool(e.Optional),
		})
	}
	if table.NumLines() > 0 {
		table.Render()
	}
	io.WriteString(w, strings.ReplaceAll(b.String(), "\n", "\n\t")+"\n")
}

// NewIndex adds a new index for the given type table.
// It fails if the schema index is invalid.
func (t *Type) AddIndex(idx *load.Index) error {
	index := &Index{Unique: idx.Unique}
	if len(idx.Fields) == 0 {
		return fmt.Errorf("missing fields")
	}
	for _, name := range idx.Fields {
		f, ok := t.fields[name]
		if !ok {
			return fmt.Errorf("unknown index field %q", name)
		}
		if f.def.Size != nil && *f.def.Size > schema.DefaultStringLen {
			return fmt.Errorf("field %q exceeds the index size limit (%d)", name, schema.DefaultStringLen)
		}
		index.Columns = append(index.Columns, snake(name))
	}
	for _, name := range idx.Edges {
		var edge *Edge
		for _, e := range t.Edges {
			if e.Name == name {
				edge = e
				break
			}
		}
		switch {
		case edge == nil:
			return fmt.Errorf("unknown index field %q", name)
		case edge.Rel.Type == O2O && !edge.IsInverse():
			return fmt.Errorf("non-inverse edge (edge.From) for index %q on O2O relation", name)
		case edge.Rel.Type != M2O && edge.Rel.Type != O2O:
			return fmt.Errorf("relation %s for inverse edge %q is not one of (O2O, M2O)", edge.Rel.Type, name)
		default:
			index.Columns = append(index.Columns, edge.Rel.Column())
		}
	}
	index.Name += strings.Join(index.Columns, "_")
	t.Indexes = append(t.Indexes, index)
	return nil
}

// Constant returns the constant name of the field.
func (f Field) Constant() string { return "Field" + pascal(f.Name) }

// DefaultName returns the variable name of the default value of this field.
func (f Field) DefaultName() string { return "Default" + pascal(f.Name) }

// UpdateDefaultName returns the variable name of the update default value of this field.
func (f Field) UpdateDefaultName() string { return "Update" + f.DefaultName() }

// StructField returns the struct member of the field.
func (f Field) StructField() string {
	if token.Lookup(f.Name).IsKeyword() {
		return "_" + f.Name
	}
	return f.Name
}

// Validator returns the validator name.
func (f Field) Validator() string { return pascal(f.Name) + "Validator" }

// IsTime returns true if the field is timestamp field.
func (f Field) IsTime() bool { return f.Type == field.TypeTime }

// IsString returns true if the field is a string field.
func (f Field) IsString() bool { return f.Type == field.TypeString }

// NullType returns the sql null-type for optional and nullable fields.
func (f Field) NullType() string {
	switch f.Type {
	case field.TypeString:
		return "sql.NullString"
	case field.TypeBool:
		return "sql.NullBool"
	case field.TypeTime:
		return "sql.NullTime"
	case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt64,
		field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64:
		return "sql.NullInt64"
	case field.TypeFloat32, field.TypeFloat64:
		return "sql.NullFloat64"
	}
	return f.Type.String()
}

// NullTypeField extracts the nullable type field (if exists) from the given receiver.
// It also does the type conversion if needed.
func (f Field) NullTypeField(rec string) string {
	switch f.Type {
	case field.TypeString, field.TypeBool, field.TypeInt64, field.TypeFloat64:
		return fmt.Sprintf("%s.%s", rec, strings.Title(f.Type.String()))
	case field.TypeTime:
		return fmt.Sprintf("%s.Time", rec)
	case field.TypeFloat32:
		return fmt.Sprintf("%s(%s.Float32)", f.Type.String(), rec)
	case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32,
		field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64:
		return fmt.Sprintf("%s(%s.Int64)", f.Type.String(), rec)
	}
	return rec
}

// Column returns the table column. It sets it as a primary key (auto_increment) in case of ID field.
func (f Field) Column() *schema.Column {
	pk := f.Name == "id"
	c := &schema.Column{
		Name:     f.Name,
		Type:     f.Type,
		Unique:   f.Unique,
		Nullable: f.Optional,
	}
	if pk {
		c.Type = field.TypeInt
		c.Increment = true
	}
	if f.def != nil {
		if f.def.Size != nil {
			c.Size = *f.def.Size
		}
		if f.def.Charset != nil {
			c.Charset = *f.def.Charset
		}
	}
	if f.Default && !f.IsTime() {
		// since this column is used only for codegen, the actual default
		// value is imported by the migrate package and used directly.
		c.Default = true
	}
	return c
}

// ExampleCode returns an example code of the field value for the example_test file.
func (f Field) ExampleCode() string {
	switch f.Type {
	case field.TypeString:
		return "\"string\""
	case field.TypeBytes:
		return "[]byte{}"
	case field.TypeBool:
		return "true"
	case field.TypeTime:
		return "time.Now()"
	default:
		return "1"
	}
}

// Label returns the Gremlin label name of the edge.
// If the edge is inverse
func (e Edge) Label() string {
	if e.IsInverse() {
		return fmt.Sprintf("%s_%s", e.Owner.Label(), snake(e.Inverse))
	}
	return fmt.Sprintf("%s_%s", e.Owner.Label(), snake(e.Name))
}

// M2M indicates if this edge is M2M edge.
func (e Edge) M2M() bool { return e.Rel.Type == M2M }

// M2O indicates if this edge is M2O edge.
func (e Edge) M2O() bool { return e.Rel.Type == M2O }

// O2M indicates if this edge is O2M edge.
func (e Edge) O2M() bool { return e.Rel.Type == O2M }

// O2O indicates if this edge is O2O edge.
func (e Edge) O2O() bool { return e.Rel.Type == O2O }

// IsInverse returns if this edge is an inverse edge.
func (e Edge) IsInverse() bool { return e.Inverse != "" }

// Constant returns the constant name of the edge.
// If the edge is inverse, it returns the constant name of the owner-edge (assoc-edge).
func (e Edge) Constant() string {
	name := e.Name
	if e.IsInverse() {
		name = e.Inverse
	}
	return pascal(name) + "Label"
}

// InverseConstant returns the inverse constant name of the edge.
func (e Edge) InverseConstant() string { return pascal(e.Name) + "InverseLabel" }

// TableConstant returns the constant name of the relation table.
func (e Edge) TableConstant() string { return pascal(e.Name) + "Table" }

// InverseTableConstant returns the constant name of the other/inverse type of the relation.
func (e Edge) InverseTableConstant() string { return pascal(e.Name) + "InverseTable" }

// ColumnConstant returns the constant name of the relation column.
func (e Edge) ColumnConstant() string { return pascal(e.Name) + "Column" }

// PKConstant returns the constant name of the primary key. Used for M2M edges.
func (e Edge) PKConstant() string { return pascal(e.Name) + "PrimaryKey" }

// HasConstraint indicates if this edge has a unique constraint check.
// We check uniqueness when both-directions are unique or one of them.
func (e Edge) HasConstraint() bool {
	return e.Rel.Type == O2O || e.Rel.Type == O2M
}

// StructField returns the struct member of the edge.
func (e Edge) StructField() string {
	if token.Lookup(e.Name).IsKeyword() {
		return "_" + e.Name
	}
	return e.Name
}

// Column returns the first element from the columns slice.
func (r Relation) Column() string {
	if len(r.Columns) == 0 {
		panic(fmt.Sprintf("missing column for Relation.Table: %s", r.Table))
	}
	return r.Columns[0]
}

// Rel is a relation type of an edge.
type Rel int

// Relation types.
const (
	Unk Rel = iota // Unknown.
	O2O            // One to one / has one.
	O2M            // One to many / has many.
	M2O            // Many to one (inverse perspective for O2M).
	M2M            // Many to many.
)

// String returns the relation name.
func (r Rel) String() string {
	s := "Unknown"
	switch r {
	case O2O:
		s = "O2O"
	case O2M:
		s = "O2M"
	case M2O:
		s = "M2O"
	case M2M:
		s = "M2M"
	}
	return s
}

func structTag(name, tag string) string {
	t := fmt.Sprintf(`json:"%s,omitempty"`, name)
	if tag == "" {
		return t
	}
	if _, ok := reflect.StructTag(tag).Lookup("json"); !ok {
		tag = t + " " + tag
	}
	return tag
}
