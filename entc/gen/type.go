// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"fmt"
	"go/token"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/facebookincubator/ent"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/entc/load"
	"github.com/facebookincubator/ent/schema/field"
)

type (
	// Type represents one node-type in the graph, its relations and
	// the information it holds.
	Type struct {
		*Config
		// schema definition.
		schema *load.Schema
		// Name holds the type/ent name.
		Name string
		// ID holds the ID field of this type.
		ID *Field
		// Fields holds all the primitive fields of this type.
		Fields []*Field
		fields map[string]*Field
		// Edge holds all the edges of this type.
		Edges []*Edge
		// Indexes are the configured indexes for this type.
		Indexes []*Index
		// ForeignKeys are the foreign-keys that resides in the type table.
		ForeignKeys []*ForeignKey
		foreignkeys map[string]*ForeignKey
	}

	// Field holds the information of a type field used for the templates.
	Field struct {
		// field definition.
		def *load.Field
		// Name is the name of this field in the database schema.
		Name string
		// Type holds the type information of the field.
		Type *field.TypeInfo
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
		// Position info of the field.
		Position *load.Position
		// UserDefined indicates that this field was defined by the loaded schema.
		// Unlike default id field, which is defined by the generator.
		UserDefined bool
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
		// Inverse holds the name of the reference edge declared in the schema.
		Inverse string
		// Owner holds the type of the edge-owner. For assoc-edges it's the
		// type that holds the edge, for inverse-edges, it's the assoc type.
		Owner *Type
		// StructTag of the edge-field in the struct. default to "json".
		StructTag string
		// Relation holds the relation info of an edge.
		Rel Relation
		// Bidi indicates if this edge is a bidirectional edge. A self-reference
		// to the same type with the same name (symmetric relation). For example,
		// a User type have one of following edges:
		//
		//	edge.To("friends", User.Type)           // many 2 many.
		//	edge.To("spouse", User.Type).Unique()   // one 2 one.
		//
		Bidi bool
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

	// ForeignKey holds the information for foreign-key columns of types.
	// It's exported only because it's used by the codegen templates and
	// should not be used beside that.
	ForeignKey struct {
		// Field information for the foreign-key column.
		Field *Field
		// Edge that is associated with this foreign-key.
		Edge *Edge
	}
)

// NewType creates a new type and its fields from the given schema.
func NewType(c *Config, schema *load.Schema) (*Type, error) {
	typ := &Type{
		Config: c,
		ID: &Field{
			Name:      "id",
			Type:      c.IDType,
			StructTag: `json:"id,omitempty"`,
		},
		schema:      schema,
		Name:        schema.Name,
		Fields:      make([]*Field, 0, len(schema.Fields)),
		fields:      make(map[string]*Field, len(schema.Fields)),
		foreignkeys: make(map[string]*ForeignKey),
	}
	for _, f := range schema.Fields {
		switch {
		case f.Name == "":
			return nil, fmt.Errorf("field name cannot be empty")
		case f.Info == nil || !f.Info.Valid():
			return nil, fmt.Errorf("invalid type for field %s", f.Name)
		case f.Nillable && !f.Optional:
			return nil, fmt.Errorf("nillable field %q must be optional", f.Name)
		case f.Unique && f.Default:
			return nil, fmt.Errorf("unique field %q cannot have default value", f.Name)
		case typ.fields[f.Name] != nil:
			return nil, fmt.Errorf("field %q redeclared for type %q", f.Name, typ.Name)
		case f.Sensitive && f.Tag != "":
			return nil, fmt.Errorf("sensitive field %q cannot have struct tags", f.Name)
		case f.Info.Type == field.TypeEnum:
			if err := validEnums(f); err != nil {
				return nil, err
			}
			// Enum types should be named as follows: typepkg.Field.
			f.Info.Ident = fmt.Sprintf("%s.%s", typ.Package(), pascal(f.Name))
		}
		tf := &Field{
			def:           f,
			Name:          f.Name,
			Type:          f.Info,
			Unique:        f.Unique,
			Position:      f.Position,
			Nillable:      f.Nillable,
			Optional:      f.Optional,
			Default:       f.Default,
			UpdateDefault: f.UpdateDefault,
			Immutable:     f.Immutable,
			StructTag:     structTag(f.Name, f.Tag),
			Validators:    f.Validators,
			UserDefined:   true,
		}
		// user defined id field.
		if tf.Name == typ.ID.Name {
			typ.ID = tf
		} else {
			typ.Fields = append(typ.Fields, tf)
			typ.fields[f.Name] = tf
		}
	}
	if typ.NumHooks() > 0 {
		typ.noSchemaImport = true
	}
	return typ, nil
}

// Label returns Gremlin label name of the node/type.
func (t Type) Label() string { return snake(t.Name) }

// Table returns SQL table name of the node/type.
func (t Type) Table() string {
	if t.schema != nil && t.schema.Config.Table != "" {
		return t.schema.Config.Table
	}
	return snake(rules.Pluralize(t.Name))
}

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

// HasValidators reports if any of the type's field has validators.
func (t Type) HasValidators() bool {
	for _, f := range t.Fields {
		if f.Validators > 0 {
			return true
		}
	}
	return false
}

// HasDefault reports if any of this type's fields has default value on creation.
func (t Type) HasDefault() bool {
	for _, f := range t.Fields {
		if f.Default {
			return true
		}
	}
	return false
}

// HasUpdateDefault reports if any of this type's fields has default value on update.
func (t Type) HasUpdateDefault() bool {
	for _, f := range t.Fields {
		if f.UpdateDefault {
			return true
		}
	}
	return false
}

// HasOptional reports if this type has an optional field.
func (t Type) HasOptional() bool {
	for _, f := range t.Fields {
		if f.Optional {
			return true
		}
	}
	return false
}

// HasNumeric reports if this type has a numeric field.
func (t Type) HasNumeric() bool {
	for _, f := range t.Fields {
		if f.Type.Numeric() {
			return true
		}
	}
	return false
}

// FKEdges returns all edges that reside on the type table as foreign-keys.
func (t Type) FKEdges() (edges []*Edge) {
	for _, e := range t.Edges {
		if e.OwnFK() {
			edges = append(edges, e)
		}
	}
	return
}

// MixedInWithDefaultOrValidator returns all mixed-in fields with default values for creation or update.
func (t Type) MixedInWithDefaultOrValidator() (fields []*Field) {
	for _, f := range t.Fields {
		if f.Position != nil && f.Position.MixedIn && (f.Default || f.UpdateDefault || f.Validators > 0) {
			fields = append(fields, f)
		}
	}
	return
}

// NumMixin returns the type's mixin count.
func (t Type) NumMixin() int {
	m := make(map[int]struct{})
	for _, f := range t.Fields {
		if p := f.Position; p != nil && p.MixedIn {
			m[p.MixinIndex] = struct{}{}
		}
	}
	return len(m)
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

// TagTypes returns all struct-tag types of the type fields.
func (t Type) TagTypes() []string {
	tags := make(map[string]bool)
	for _, f := range t.Fields {
		tag := reflect.StructTag(f.StructTag)
		fields := strings.FieldsFunc(f.StructTag, func(r rune) bool {
			return r == ':' || unicode.IsSpace(r)
		})
		for _, name := range fields {
			_, ok := tag.Lookup(name)
			if ok && !tags[name] {
				tags[name] = true
			}
		}
	}
	r := make([]string, 0, len(tags))
	for tag := range tags {
		r = append(r, tag)
	}
	sort.Strings(r)
	return r
}

// AddIndex adds a new index for the type.
// It fails if the schema index is invalid.
func (t *Type) AddIndex(idx *load.Index) error {
	index := &Index{Name: idx.StorageKey, Unique: idx.Unique}
	if len(idx.Fields) == 0 && len(idx.Edges) == 0 {
		return fmt.Errorf("missing fields or edges")
	}
	for _, name := range idx.Fields {
		f, ok := t.fields[name]
		if !ok {
			return fmt.Errorf("unknown index field %q", name)
		}
		if f.def.Size != nil && *f.def.Size > schema.DefaultStringLen {
			return fmt.Errorf("field %q exceeds the index size limit (%d)", name, schema.DefaultStringLen)
		}
		index.Columns = append(index.Columns, f.StorageKey())
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
	// If no storage-key was defined for this index, generate one.
	if idx.StorageKey == "" {
		// Add the type name as a prefix to the index parts, because
		// multiple types can share the same index attributes.
		parts := append([]string{strings.ToLower(t.Name)}, index.Columns...)
		index.Name = strings.Join(parts, "_")
	}
	t.Indexes = append(t.Indexes, index)
	return nil
}

// resolveFKs makes sure all edge-fks are created for the types.
func (t *Type) resolveFKs() {
	for _, e := range t.Edges {
		if e.IsInverse() || e.M2M() {
			continue
		}
		typ := t.ID.Type
		if e.OwnFK() {
			typ = e.Type.ID.Type
		}
		fk := &ForeignKey{
			Edge: e,
			Field: &Field{
				Name:        e.Rel.Column(),
				Type:        typ,
				Nillable:    true,
				Optional:    true,
				Unique:      e.Unique,
				UserDefined: e.Type.ID.UserDefined,
			},
		}
		if e.OwnFK() {
			t.addFK(fk)
		} else {
			e.Type.addFK(fk)
		}
	}
}

// AddForeignKey adds a foreign-key for the type if it doesn't exist.
func (t *Type) addFK(fk *ForeignKey) {
	if _, ok := t.foreignkeys[fk.Field.Name]; ok {
		return
	}
	t.foreignkeys[fk.Field.Name] = fk
	t.ForeignKeys = append(t.ForeignKeys, fk)
}

// QueryName returns the struct name of the query builder for this type.
func (t Type) QueryName() string {
	return pascal(t.Name) + "Query"
}

// MutationName returns the struct name of the mutation builder for this type.
func (t Type) MutationName() string {
	return pascal(t.Name) + "Mutation"
}

// SiblingImports returns all sibling packages that are needed for the different builders.
func (t Type) SiblingImports() []string {
	var (
		paths = []string{path.Join(t.Config.Package, t.Package())}
		seen  = map[string]bool{paths[0]: true}
	)
	for _, e := range t.Edges {
		name := path.Join(t.Config.Package, e.Type.Package())
		if !seen[name] {
			seen[name] = true
			paths = append(paths, name)
		}
	}
	return paths
}

// NumHooks returns the number of hooks declared in the type schema.
func (t Type) NumHooks() int {
	if t.schema != nil {
		return t.schema.Hooks
	}
	return 0
}

// ImportSchema reports if the type-package need to import the schema.
func (t Type) ImportSchema() bool {
	return !t.state.noSchemaImport
}

// Constant returns the constant name of the field.
func (f Field) Constant() string {
	return "Field" + pascal(f.Name)
}

// DefaultName returns the variable name of the default value of this field.
func (f Field) DefaultName() string { return "Default" + pascal(f.Name) }

// UpdateDefaultName returns the variable name of the update default value of this field.
func (f Field) UpdateDefaultName() string { return "Update" + f.DefaultName() }

// DefaultValue returns the default value of the field. Invoked by the template.
func (f Field) DefaultValue() interface{} { return f.def.DefaultValue }

// BuilderField returns the struct member of the field in the builder.
func (f Field) BuilderField() string {
	return builderField(f.Name)
}

// StructField returns the struct member of the field in the model.
func (f Field) StructField() string {
	return pascal(f.Name)
}

// Enums returns the enum values of a field.
func (f Field) Enums() []string {
	if f.IsEnum() {
		return f.def.Enums
	}
	return nil
}

// EnumName returns the constant name of the enum value.
func (f Field) EnumName(enum string) string {
	return pascal(f.Name) + pascal(enum)
}

// Validator returns the validator name.
func (f Field) Validator() string { return pascal(f.Name) + "Validator" }

// mutMethods returns the method names of mutation interface.
var mutMethods = func() map[string]struct{} {
	t := reflect.TypeOf(new(ent.Mutation)).Elem()
	names := make(map[string]struct{})
	for i := 0; i < t.NumMethod(); i++ {
		names[t.Method(i).Name] = struct{}{}
	}
	return names
}()

// MutationGet returns the method name for getting the field value.
// The default name is just a pascal format. If the the method conflicts
// with the mutation methods, prefix the method with "Get".
func (f Field) MutationGet() string {
	name := pascal(f.Name)
	if _, ok := mutMethods[name]; ok {
		name = "Get" + name
	}
	return name
}

// IsTime returns true if the field is a timestamp field.
func (f Field) IsTime() bool { return f.Type != nil && f.Type.Type == field.TypeTime }

// IsJSON returns true if the field is a JSON field.
func (f Field) IsJSON() bool { return f.Type != nil && f.Type.Type == field.TypeJSON }

// IsString returns true if the field is a string field.
func (f Field) IsString() bool { return f.Type != nil && f.Type.Type == field.TypeString }

// IsUUID returns true if the field is a UUID field.
func (f Field) IsUUID() bool { return f.Type != nil && f.Type.Type == field.TypeUUID }

// IsInt returns true if the field is an int field.
func (f Field) IsInt() bool { return f.Type != nil && f.Type.Type == field.TypeInt }

// IsEnum returns true if the field is an enum field.
func (f Field) IsEnum() bool { return f.Type != nil && f.Type.Type == field.TypeEnum }

// Sensitive returns true if the field is a sensitive field.
func (f Field) Sensitive() bool { return f.def != nil && f.def.Sensitive }

// NullType returns the sql null-type for optional and nullable fields.
func (f Field) NullType() string {
	switch f.Type.Type {
	case field.TypeJSON:
		return "[]byte"
	case field.TypeString, field.TypeEnum:
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
	switch f.Type.Type {
	case field.TypeEnum:
		return fmt.Sprintf("%s(%s.String)", f.Type, rec)
	case field.TypeString, field.TypeBool, field.TypeInt64, field.TypeFloat64:
		return fmt.Sprintf("%s.%s", rec, strings.Title(f.Type.String()))
	case field.TypeTime:
		return fmt.Sprintf("%s.Time", rec)
	case field.TypeFloat32:
		return fmt.Sprintf("%s(%s.Float64)", f.Type, rec)
	case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32,
		field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64:
		return fmt.Sprintf("%s(%s.Int64)", f.Type, rec)
	}
	return rec
}

// Column returns the table column. It sets it as a primary key (auto_increment) in case of ID field.
func (f Field) Column() *schema.Column {
	f.Enums()
	c := &schema.Column{
		Name:     f.StorageKey(),
		Type:     f.Type.Type,
		Unique:   f.Unique,
		Nullable: f.Optional,
		Enums:    f.Enums(),
	}
	if f.def != nil && f.def.Size != nil {
		c.Size = *f.def.Size
	}
	switch {
	case f.Default && (f.Type.Numeric() || f.Type.Type == field.TypeBool):
		c.Default = f.DefaultValue()
	case f.Default && (f.IsString() || f.IsEnum()):
		if s, ok := f.DefaultValue().(string); ok {
			c.Default = strconv.Quote(s)
		}
	}
	return c
}

// PK is like Column, but for table primary key.
func (f Field) PK() *schema.Column {
	c := &schema.Column{
		Name:      f.StorageKey(),
		Type:      field.TypeInt,
		Key:       schema.PrimaryKey,
		Increment: true,
	}
	// If the PK was defined by the user and it's UUID or string.
	if f.UserDefined && !f.Type.Numeric() {
		c.Increment = false
		c.Type = f.Type.Type
		c.Unique = f.Unique
		if f.def != nil && f.def.Size != nil {
			c.Size = *f.def.Size
		}
	}
	return c
}

// StorageKey returns the storage name of the field.
// SQL columns or Gremlin property.
func (f Field) StorageKey() string {
	if f.def != nil && f.def.StorageKey != "" {
		return f.def.StorageKey
	}
	return snake(f.Name)
}

// Label returns the Gremlin label name of the edge.
// If the edge is inverse
func (e Edge) Label() string {
	if e.IsInverse() {
		return fmt.Sprintf("%s_%s", e.Owner.Label(), snake(e.Inverse))
	}
	return fmt.Sprintf("%s_%s", e.Owner.Label(), snake(e.Name))
}

// Constant returns the constant name of the edge.
func (e Edge) Constant() string {
	return "Edge" + pascal(e.Name)
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

// Constant returns the constant name of the edge for the gremlin dialect.
// If the edge is inverse, it returns the constant name of the owner-edge (assoc-edge).
func (e Edge) LabelConstant() string {
	name := e.Name
	if e.IsInverse() {
		name = e.Inverse
	}
	return pascal(name) + "Label"
}

// InverseConstant returns the inverse constant name of the edge.
func (e Edge) InverseLabelConstant() string { return pascal(e.Name) + "InverseLabel" }

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
// Used by the Gremlin storage-layer.
func (e Edge) HasConstraint() bool {
	return e.Rel.Type == O2O || e.Rel.Type == O2M
}

// BuilderField returns the struct member of the edge in the builder.
func (e Edge) BuilderField() string {
	return builderField(e.Name)
}

// StructField returns the struct member of the edge in the model.
func (e Edge) StructField() string {
	return pascal(e.Name)
}

// StructFKField returns the struct member for holding the edge
// foreign-key in the model.
func (e Edge) StructFKField() string {
	return e.Rel.Column()
}

// OwnFK indicates if the foreign-key of this edge is owned by the edge
// column (reside in the type's table). Used by the SQL storage-driver.
func (e Edge) OwnFK() bool {
	switch {
	case e.M2O():
		return true
	case e.O2O() && (e.IsInverse() || e.Bidi):
		return true
	}
	return false
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

func validEnums(f *load.Field) error {
	if len(f.Enums) == 0 {
		return fmt.Errorf("missing values for enum field %q", f.Name)
	}
	values := make(map[string]bool, len(f.Enums))
	for _, e := range f.Enums {
		switch {
		case e == "":
			return fmt.Errorf("%q field value cannot be empty", f.Name)
		case values[e]:
			return fmt.Errorf("duplicate values %q for enum field %q", e, f.Name)
		default:
			values[e] = true
		}
	}
	if value := f.DefaultValue; value != nil {
		if value, ok := value.(string); !ok || !values[value] {
			return fmt.Errorf("invalid default value for enum field %q", f.Name)
		}
	}
	return nil
}

// builderField returns the struct field for the given name
// and ensures it doesn't conflict with Go keywords and other
// private fields.
func builderField(name string) string {
	if token.Lookup(name).IsKeyword() || name == "config" {
		return "_" + name
	}
	return name
}
