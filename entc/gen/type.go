// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go/token"
	"go/types"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/entc/load"
	"github.com/facebook/ent/schema/field"
)

// The following types and their exported methods used by the codegen
// to generate the assets.
type (
	// Type represents one node-type in the graph, its relations and
	// the information it holds.
	Type struct {
		*Config
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
		foreignKeys map[string]struct{}
		// Annotations that were defined for the field in the schema.
		// The mapping is from the Annotation.Name() to a JSON decoded object.
		Annotations map[string]interface{}
	}

	// Field holds the information of a type field used for the templates.
	Field struct {
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
		// Enums information for enum fields.
		Enums []Enum
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
		// Annotations that were defined for the field in the schema.
		// The mapping is from the Annotation.Name() to a JSON decoded object.
		Annotations map[string]interface{}
	}

	// Edge of a graph between two types.
	Edge struct {
		def *load.Edge
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
		// Annotations that were defined for the edge in the schema.
		// The mapping is from the Annotation.Name() to a JSON decoded object.
		Annotations map[string]interface{}
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
	// Enum holds the enum information for schema enums in codegen.
	Enum struct {
		// Name is the Go name of the enum.
		Name string
		// Value in the schema.
		Value string
	}
)

// NewType creates a new type and its fields from the given schema.
func NewType(c *Config, schema *load.Schema) (*Type, error) {
	idType := c.IDType
	if idType == nil {
		idType = defaultIDType
	}
	typ := &Type{
		Config: c,
		ID: &Field{
			Name: "id",
			def: &load.Field{
				Name: "id",
			},
			Type:      idType,
			StructTag: structTag("id", ""),
		},
		schema:      schema,
		Name:        schema.Name,
		Annotations: schema.Annotations,
		Fields:      make([]*Field, 0, len(schema.Fields)),
		fields:      make(map[string]*Field, len(schema.Fields)),
		foreignKeys: make(map[string]struct{}),
	}
	if err := typ.check(); err != nil {
		return nil, err
	}
	for _, f := range schema.Fields {
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
			Annotations:   f.Annotations,
		}
		if err := typ.checkField(tf, f); err != nil {
			return nil, err
		}
		// User defined id field.
		if tf.Name == typ.ID.Name {
			typ.ID = tf
		} else {
			typ.Fields = append(typ.Fields, tf)
			typ.fields[f.Name] = tf
		}
	}
	return typ, nil
}

// Label returns Gremlin label name of the node/type.
func (t Type) Label() string {
	return snake(t.Name)
}

// Table returns SQL table name of the node/type.
func (t Type) Table() string {
	if ant := t.EntSQL(); ant != nil && ant.Table != "" {
		return ant.Table
	}
	if t.schema != nil && t.schema.Config.Table != "" {
		return t.schema.Config.Table
	}
	return snake(rules.Pluralize(t.Name))
}

// EntSQL returns the EntSQL annotation if exists.
func (t Type) EntSQL() *entsql.Annotation {
	return entsqlAnnotate(t.Annotations)
}

// Package returns the package name of this node.
func (t Type) Package() string {
	return strings.ToLower(t.Name)
}

// Receiver returns the receiver name of this node. It makes sure the
// receiver names doesn't conflict with import names.
func (t Type) Receiver() string {
	return receiver(t.Name)
}

// HasAssoc returns true if this type has an assoc-edge (non-inverse)
// with the given name. faster than map access for most cases.
func (t Type) HasAssoc(name string) (*Edge, bool) {
	for _, e := range t.Edges {
		if name == e.Name && !e.IsInverse() {
			return e, true
		}
	}
	return nil, false
}

// HasValidators reports if any of the type's field has validators.
func (t Type) HasValidators() bool {
	fields := t.Fields
	if t.ID.UserDefined {
		fields = append(fields, t.ID)
	}
	for _, f := range fields {
		if f.Validators > 0 {
			return true
		}
	}
	return false
}

// HasDefault reports if any of this type's fields has default value on creation.
func (t Type) HasDefault() bool {
	fields := t.Fields
	if t.ID.UserDefined {
		fields = append(fields, t.ID)
	}
	for _, f := range fields {
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

// HasUpdateCheckers reports if this type has any checkers to run on update(one).
func (t Type) HasUpdateCheckers() bool {
	for _, f := range t.Fields {
		if (f.Validators > 0 || f.IsEnum()) && !f.Immutable {
			return true
		}
	}
	for _, e := range t.Edges {
		if e.Unique && !e.Optional {
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

// RuntimeMixin returns schema mixin that needs to be loaded at
// runtime. For example, for default values, validators or hooks.
func (t Type) RuntimeMixin() bool {
	return len(t.MixedInFields()) > 0 || len(t.MixedInHooks()) > 0 || len(t.MixedInPolicies()) > 0
}

// MixedInFields returns the indices of mixin holds runtime code.
func (t Type) MixedInFields() []int {
	idx := make(map[int]struct{})
	fields := t.Fields
	if t.ID.UserDefined {
		fields = append(fields, t.ID)
	}
	for _, f := range fields {
		if f.Position != nil && f.Position.MixedIn && (f.Default || f.UpdateDefault || f.Validators > 0) {
			idx[f.Position.MixinIndex] = struct{}{}
		}
	}
	return sortedKeys(idx)
}

// MixedInHooks returns the indices of mixin with hooks.
func (t Type) MixedInHooks() []int {
	if t.schema == nil {
		return nil
	}
	idx := make(map[int]struct{})
	for _, h := range t.schema.Hooks {
		if h.MixedIn {
			idx[h.MixinIndex] = struct{}{}
		}
	}
	return sortedKeys(idx)
}

// MixedInPolicies returns the indices of mixin with policies.
func (t Type) MixedInPolicies() []int {
	if t.schema == nil {
		return nil
	}
	idx := make(map[int]struct{})
	for _, h := range t.schema.Policy {
		if h.MixedIn {
			idx[h.MixinIndex] = struct{}{}
		}
	}
	return sortedKeys(idx)
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

// EnumFields returns the types's enum fields.
func (t Type) EnumFields() []*Field {
	var fields []*Field
	for _, f := range t.Fields {
		if f.IsEnum() {
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
		var f *Field
		if name == t.ID.Name {
			f = t.ID
		} else {
			var ok bool
			f, ok = t.fields[name]
			if !ok {
				return fmt.Errorf("unknown index field %q", name)
			}
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
func (t *Type) resolveFKs() error {
	for _, e := range t.Edges {
		if err := e.setStorageKey(); err != nil {
			return fmt.Errorf("%q edge: %v", e.Name, err)
		}
		if e.IsInverse() || e.M2M() {
			continue
		}
		refid := t.ID
		if e.OwnFK() {
			refid = e.Type.ID
		}
		fk := &ForeignKey{
			Edge: e,
			Field: &Field{
				Name:        builderField(e.Rel.Column()),
				Type:        refid.Type,
				Nillable:    true,
				Optional:    true,
				Unique:      e.Unique,
				UserDefined: refid.UserDefined,
			},
		}
		if e.OwnFK() {
			t.addFK(fk)
		} else {
			e.Type.addFK(fk)
		}
	}
	return nil
}

// AddForeignKey adds a foreign-key for the type if it doesn't exist.
func (t *Type) addFK(fk *ForeignKey) {
	if _, ok := t.foreignKeys[fk.Field.Name]; ok {
		return
	}
	t.foreignKeys[fk.Field.Name] = struct{}{}
	t.ForeignKeys = append(t.ForeignKeys, fk)
}

// QueryName returns the struct name denoting the query-builder for this type.
func (t Type) QueryName() string {
	return pascal(t.Name) + "Query"
}

// FilterName returns the struct name denoting the filter-builder for this type.
func (t Type) FilterName() string {
	return pascal(t.Name) + "Filter"
}

// CreateName returns the struct name denoting the create-builder for this type.
func (t Type) CreateName() string {
	return pascal(t.Name) + "Create"
}

// CreateBulkName returns the struct name denoting the create-bulk-builder for this type.
func (t Type) CreateBulkName() string {
	return pascal(t.Name) + "CreateBulk"
}

// UpdateName returns the struct name denoting the update-builder for this type.
func (t Type) UpdateName() string {
	return pascal(t.Name) + "Update"
}

// UpdateOneName returns the struct name denoting the update-one-builder for this type.
func (t Type) UpdateOneName() string {
	return pascal(t.Name) + "UpdateOne"
}

// DeleteName returns the struct name denoting the delete-builder for this type.
func (t Type) DeleteName() string {
	return pascal(t.Name) + "Delete"
}

// DeleteOneName returns the struct name denoting the delete-one-builder for this type.
func (t Type) DeleteOneName() string {
	return pascal(t.Name) + "DeleteOne"
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
		return len(t.schema.Hooks)
	}
	return 0
}

// HookPositions returns the position information of hooks declared in the type schema.
func (t Type) HookPositions() []*load.Position {
	if t.schema != nil {
		return t.schema.Hooks
	}
	return nil
}

// NumPolicy returns the number of privacy-policy declared in the type schema.
func (t Type) NumPolicy() int {
	if t.schema != nil {
		return len(t.schema.Policy)
	}
	return 0
}

// PolicyPositions returns the position information of privacy policy declared in the type schema.
func (t Type) PolicyPositions() []*load.Position {
	if t.schema != nil {
		return t.schema.Policy
	}
	return nil
}

// RelatedTypes returns all the types (nodes) that
// are related (with edges) to this type.
func (t Type) RelatedTypes() []*Type {
	seen := make(map[string]struct{})
	related := make([]*Type, 0, len(t.Edges))
	for _, e := range t.Edges {
		if _, ok := seen[e.Type.Name]; !ok {
			related = append(related, e.Type)
			seen[e.Type.Name] = struct{}{}
		}
	}
	return related
}

// check checks the schema type.
func (t *Type) check() error {
	pkg := t.Package()
	if token.Lookup(pkg).IsKeyword() {
		return fmt.Errorf("schema lowercase name conflicts with Go keyword %q", pkg)
	}
	if types.Universe.Lookup(pkg) != nil {
		return fmt.Errorf("schema lowercase name conflicts with Go predeclared identifier %q", pkg)
	}
	if _, ok := globalIdent[t.Name]; ok {
		return fmt.Errorf("schema name conflicts with ent predeclared identifier %q", t.Name)
	}
	return nil
}

// checkField checks the schema field.
func (t *Type) checkField(tf *Field, f *load.Field) (err error) {
	switch {
	case f.Name == "":
		err = fmt.Errorf("field name cannot be empty")
	case f.Info == nil || !f.Info.Valid():
		err = fmt.Errorf("invalid type for field %s", f.Name)
	case f.Nillable && !f.Optional:
		err = fmt.Errorf("nillable field %q must be optional", f.Name)
	case f.Unique && f.Default && f.Info.Type != field.TypeUUID:
		err = fmt.Errorf("unique field %q cannot have default value", f.Name)
	case t.fields[f.Name] != nil:
		err = fmt.Errorf("field %q redeclared for type %q", f.Name, t.Name)
	case f.Sensitive && f.Tag != "":
		err = fmt.Errorf("sensitive field %q cannot have struct tags", f.Name)
	case f.Info.Type == field.TypeEnum:
		if tf.Enums, err = tf.enums(f); err == nil && !tf.HasGoType() {
			// Enum types should be named as follows: typepkg.Field.
			f.Info.Ident = fmt.Sprintf("%s.%s", t.Package(), pascal(f.Name))
		}
	case tf.Validators > 0 && !tf.ConvertedToBasic():
		err = fmt.Errorf("GoType %q for field %q must be converted to the basic %q type for validators", tf.Type, f.Name, tf.Type.Type)
	}
	return err
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

// EnumNames returns the enum values of a field.
func (f Field) EnumNames() []string {
	names := make([]string, 0, len(f.def.Enums))
	for _, e := range f.Enums {
		names = append(names, e.Name)
	}
	return names
}

// EnumValues returns the values of the enum field.
func (f Field) EnumValues() []string {
	values := make([]string, 0, len(f.def.Enums))
	for _, e := range f.Enums {
		values = append(values, e.Value)
	}
	return values
}

// EnumName returns the constant name for the enum.
func (f Field) EnumName(enum string) string {
	if !token.IsExported(enum) {
		enum = pascal(enum)
	}
	return pascal(f.Name) + enum
}

// Validator returns the validator name.
func (f Field) Validator() string {
	return pascal(f.Name) + "Validator"
}

// EntSQL returns the EntSQL annotation if exists.
func (f Field) EntSQL() *entsql.Annotation {
	return entsqlAnnotate(f.Annotations)
}

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

// MutationGetOld returns the method name for getting the old value of a field.
func (f Field) MutationGetOld() string {
	name := "Old" + pascal(f.Name)
	if _, ok := mutMethods[name]; ok {
		name = "Get" + name
	}
	return name
}

// MutationReset returns the method name for resetting the field value.
// The default name is "Reset<FieldName>". If the the method conflicts
// with the mutation methods, suffix the method with "Field".
func (f Field) MutationReset() string {
	name := "Reset" + pascal(f.Name)
	if _, ok := mutMethods[name]; ok {
		name += "Field"
	}
	return name
}

// IsBool returns true if the field is a bool field.
func (f Field) IsBool() bool { return f.Type != nil && f.Type.Type == field.TypeBool }

// IsBytes returns true if the field is a bytes field.
func (f Field) IsBytes() bool { return f.Type != nil && f.Type.Type == field.TypeBytes }

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
	if f.Type.ValueScanner() {
		return f.Type.String()
	}
	switch f.Type.Type {
	case field.TypeJSON, field.TypeBytes:
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
	expr := rec
	switch f.Type.Type {
	case field.TypeEnum:
		expr = fmt.Sprintf("%s(%s.String)", f.Type, rec)
	case field.TypeString, field.TypeBool, field.TypeInt64, field.TypeFloat64:
		expr = f.goType(fmt.Sprintf("%s.%s", rec, strings.Title(f.Type.Type.String())))
	case field.TypeTime:
		expr = fmt.Sprintf("%s.Time", rec)
	case field.TypeFloat32:
		expr = fmt.Sprintf("%s(%s.Float64)", f.Type, rec)
	case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32,
		field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64:
		expr = fmt.Sprintf("%s(%s.Int64)", f.Type, rec)
	}
	return expr
}

// Column returns the table column. It sets it as a primary key (auto_increment) in case of ID field, unless stated
// otherwise.
func (f Field) Column() *schema.Column {
	c := &schema.Column{
		Name:     f.StorageKey(),
		Type:     f.Type.Type,
		Unique:   f.Unique,
		Nullable: f.Optional,
		Size:     f.size(),
		Enums:    f.EnumValues(),
	}
	switch {
	case f.Default && (f.Type.Numeric() || f.Type.Type == field.TypeBool):
		c.Default = f.DefaultValue()
	case f.Default && (f.IsString() || f.IsEnum()):
		if s, ok := f.DefaultValue().(string); ok {
			c.Default = strconv.Quote(s)
		}
	}
	if f.def != nil {
		c.SchemaType = f.def.SchemaType
	}
	return c
}

// incremental returns if the column has an incremental behavior.
// If no value is defined externally, we use a provided def flag
func (f Field) incremental(def bool) bool {
	if ant := f.EntSQL(); ant != nil && ant.Incremental != nil {
		return *ant.Incremental
	}
	return def
}

// size returns the the field size defined in the schema.
func (f Field) size() int64 {
	if ant := f.EntSQL(); ant != nil && ant.Size != 0 {
		return ant.Size
	}
	if f.def != nil && f.def.Size != nil {
		return *f.def.Size
	}
	return 0
}

// PK is like Column, but for table primary key.
func (f Field) PK() *schema.Column {
	c := &schema.Column{
		Name:      f.StorageKey(),
		Type:      f.Type.Type,
		Key:       schema.PrimaryKey,
		Increment: f.incremental(true),
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
	if f.def != nil {
		c.SchemaType = f.def.SchemaType
	}
	return c
}

// StorageKey returns the storage name of the field.
// SQL column or Gremlin property.
func (f Field) StorageKey() string {
	if f.def != nil && f.def.StorageKey != "" {
		return f.def.StorageKey
	}
	return snake(f.Name)
}

// HasGoType indicate if a basic field (like string or bool)
// has a custom GoType.
func (f Field) HasGoType() bool {
	return f.Type != nil && f.Type.RType != nil
}

// ConvertedToBasic indicates if the Go type of the field
// can be converted to basic type (string, int, etc).
func (f Field) ConvertedToBasic() bool {
	return !f.HasGoType() || f.BasicType("ident") != ""
}

var (
	nullBoolType   = reflect.TypeOf(sql.NullBool{})
	nullTimeType   = reflect.TypeOf(sql.NullTime{})
	nullStringType = reflect.TypeOf(sql.NullString{})
)

// BasicType returns a Go expression for the given identifier
// to convert it to a basic type. For example:
//
//	v (http.Dir)		=> string(v)
//	v (fmt.Stringer)	=> v.String()
//	v (sql.NullString)	=> v.String
//
func (f Field) BasicType(ident string) (expr string) {
	if !f.HasGoType() {
		return ident
	}
	t, rt := f.Type, f.Type.RType
	switch t.Type {
	case field.TypeEnum:
		expr = ident
	case field.TypeBool:
		switch {
		case rt.Kind == reflect.Bool:
			expr = fmt.Sprintf("bool(%s)", ident)
		case rt.TypeEqual(nullBoolType):
			expr = fmt.Sprintf("%s.Bool", ident)
		}
	case field.TypeBytes:
		if rt.Kind == reflect.Slice {
			expr = fmt.Sprintf("[]byte(%s)", ident)
		}
	case field.TypeTime:
		switch {
		case rt.TypeEqual(nullTimeType):
			expr = fmt.Sprintf("%s.Time", ident)
		case rt.Kind == reflect.Struct:
			expr = fmt.Sprintf("time.Time(%s)", ident)
		}
	case field.TypeString:
		switch {
		case rt.Kind == reflect.String:
			expr = fmt.Sprintf("string(%s)", ident)
		case t.Stringer():
			expr = fmt.Sprintf("%s.String()", ident)
		case rt.TypeEqual(nullStringType):
			expr = fmt.Sprintf("%s.String", ident)
		}
	default:
		if t.Numeric() && rt.Kind >= reflect.Int && rt.Kind <= reflect.Float64 {
			expr = fmt.Sprintf("%s(%s)", rt.Kind, ident)
		}
	}
	return expr
}

// goType returns the Go expression for the given basic-type
// identifier to covert it to the custom Go type.
func (f Field) goType(ident string) string {
	if !f.HasGoType() {
		return ident
	}
	return fmt.Sprintf("%s(%s)", f.Type, ident)
}

func (f Field) enums(lf *load.Field) ([]Enum, error) {
	if len(lf.Enums) == 0 {
		return nil, fmt.Errorf("missing values for enum field %q", f.Name)
	}
	enums := make([]Enum, 0, len(lf.Enums))
	values := make(map[string]bool, len(lf.Enums))
	for i := range lf.Enums {
		switch name, value := lf.Enums[i].N, lf.Enums[i].V; {
		case value == "":
			return nil, fmt.Errorf("%q field value cannot be empty", f.Name)
		case values[value]:
			return nil, fmt.Errorf("duplicate values %q for enum field %q", value, f.Name)
		case strings.IndexFunc(value, unicode.IsSpace) != -1:
			return nil, fmt.Errorf("enum value %q cannot contain spaces", value)
		default:
			values[value] = true
			enums = append(enums, Enum{Name: f.EnumName(name), Value: value})
		}
	}
	if value := lf.DefaultValue; value != nil {
		if value, ok := value.(string); !ok || !values[value] {
			return nil, fmt.Errorf("invalid default value for enum field %q", f.Name)
		}
	}
	return enums, nil
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

// LabelConstant returns the constant name of the edge for the gremlin dialect.
// If the edge is inverse, it returns the constant name of the owner-edge (assoc-edge).
func (e Edge) LabelConstant() string {
	name := e.Name
	if e.IsInverse() {
		name = e.Inverse
	}
	return pascal(name) + "Label"
}

// InverseLabelConstant returns the inverse constant name of the edge.
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

// EagerLoadField returns the struct field (of query builder) for storing the eager-loading info.
func (e Edge) EagerLoadField() string {
	return "with" + pascal(e.Name)
}

// StructField returns the struct member of the edge in the model.
func (e Edge) StructField() string {
	return pascal(e.Name)
}

// StructFKField returns the struct member for holding the edge
// foreign-key in the model.
func (e Edge) StructFKField() string {
	return builderField(e.Rel.Column())
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

// MutationSet returns the method name for setting the edge id.
func (e Edge) MutationSet() string {
	return "Set" + pascal(e.Name) + "ID"
}

// MutationAdd returns the method name for adding edge ids.
func (e Edge) MutationAdd() string {
	return "Add" + pascal(rules.Singularize(e.Name)) + "IDs"
}

// MutationReset returns the method name for resetting the edge value.
// The default name is "Reset<EdgeName>". If the the method conflicts
// with the mutation methods, suffix the method with "Edge".
func (e Edge) MutationReset() string {
	name := "Reset" + pascal(e.Name)
	if _, ok := mutMethods[name]; ok {
		name += "Edge"
	}
	return name
}

// MutationClear returns the method name for clearing the edge value.
// The default name is "Clear<EdgeName>". If the the method conflicts
// with the mutation methods, suffix the method with "Edge".
func (e Edge) MutationClear() string {
	name := "Clear" + pascal(e.Name)
	if _, ok := mutMethods[name]; ok {
		name += "Edge"
	}
	return name
}

// MutationCleared returns the method name for indicating if the edge
// was cleared in the mutation. The default name is "<EdgeName>Cleared".
// If the the method conflicts with the mutation methods, add "Edge" the
// after the edge name.
func (e Edge) MutationCleared() string {
	name := pascal(e.Name) + "Cleared"
	if _, ok := mutMethods[name]; ok {
		return pascal(e.Name) + "EdgeCleared"
	}
	return name
}

// setStorageKey sets the storage-key option in the schema or fail.
func (e *Edge) setStorageKey() error {
	rel := e.Rel
	key := e.def.StorageKey
	if e.IsInverse() {
		assoc, ok := e.Owner.HasAssoc(e.Inverse)
		if ok {
			key = assoc.def.StorageKey
		}
	}
	if key == nil {
		return nil
	}
	switch {
	case key.Table != "" && rel.Type != M2M:
		return fmt.Errorf("StorageKey.Table is allowed only for M2M edges (got %s)", e.Rel.Type)
	case len(key.Columns) == 1 && rel.Type == M2M:
		return fmt.Errorf("%s edge have 2 columns. Use edge.Columns(to, from) instead", e.Rel.Type)
	case len(key.Columns) > 1 && rel.Type != M2M:
		return fmt.Errorf("%s edge does not have 2 columns. Use edge.Column(%s) instead", e.Rel.Type, key.Columns[0])
	}
	if key.Table != "" {
		e.Rel.Table = key.Table
	}
	if len(key.Columns) > 0 {
		e.Rel.Columns[0] = key.Columns[0]
	}
	if len(key.Columns) > 1 {
		e.Rel.Columns[1] = key.Columns[1]
	}
	return nil
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

// builderField returns the struct field for the given name
// and ensures it doesn't conflict with Go keywords and other
// builder fields and it's not exported.
func builderField(name string) string {
	_, ok := privateField[name]
	if ok || token.Lookup(name).IsKeyword() || strings.ToUpper(name[:1]) == name[:1] {
		return "_" + name
	}
	return name
}

// entsqlAnnotate extracts the entsql annotation from a loaded annotation format.
func entsqlAnnotate(annotation map[string]interface{}) *entsql.Annotation {
	annotate := &entsql.Annotation{}
	if annotation == nil || annotation[annotate.Name()] == nil {
		return nil
	}
	if buf, err := json.Marshal(annotation[annotate.Name()]); err == nil {
		_ = json.Unmarshal(buf, &annotate)
	}
	return annotate
}

var (
	// global identifiers used by the generated package.
	globalIdent = names(
		"AggregateFunc",
		"As",
		"Asc",
		"Client",
		"Count",
		"Debug",
		"Desc",
		"Driver",
		"Hook",
		"Log",
		"MutateFunc",
		"Mutation",
		"Mutator",
		"Op",
		"Option",
		"OrderFunc",
		"Max",
		"Mean",
		"Min",
		"Sum",
		"Policy",
		"Query",
		"Value",
	)
	// private fields used by the different builders.
	privateField = names(
		"config",
		"done",
		"hooks",
		"limit",
		"mutation",
		"offset",
		"oldValue",
		"order",
		"op",
		"path",
		"predicates",
		"typ",
		"unique",
		"withFKs",
	)
)

func names(ids ...string) map[string]struct{} {
	m := make(map[string]struct{})
	for i := range ids {
		m[ids[i]] = struct{}{}
	}
	return m
}

func sortedKeys(m map[int]struct{}) []int {
	s := make([]int, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Ints(s)
	return s
}
