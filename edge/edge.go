package edge

import (
	"reflect"

	"fbc/ent/field"
)

// Edge represents an edge in the graph.
type Edge struct {
	typ      string
	tag      string
	ref      string
	name     string
	unique   bool
	required bool
	inverse  bool
	parent   *Edge
	fields   []*field.Field
}

// To defines an association edge between two vertices.
func To(name string, t interface{}) *assocBuilder {
	return &assocBuilder{&Edge{name: name, typ: typ(t)}}
}

// From represents a reversed-edge between two vertices that has a back-reference to its source edge.
func From(name string, t interface{}) *inverseBuilder {
	return &inverseBuilder{&Edge{name: name, typ: typ(t), inverse: true}}
}

// Type returns the type of the edge.
func (e Edge) Type() string { return e.typ }

// IsUnique returns is the edge is unique.
func (e Edge) IsUnique() bool { return e.unique }

// AssocName returns the edge name.
func (e Edge) Name() string { return e.name }

// IsAssoc returns is the edge is assoc type.
func (e Edge) IsAssoc() bool { return !e.inverse }

// IsInverse returns is the edge is inverse type.
func (e Edge) IsInverse() bool { return e.inverse }

// Assoc returns the assoc edge of the inverse edge.
func (e Edge) Assoc() *Edge { return e.parent }

// RefName returns the reference edge name.
func (e Edge) RefName() string { return e.ref }

// GetFields returns the edge fields.
func (e Edge) GetFields() []*field.Field { return e.fields }

// Tag returns the struct tag of the edge.
func (e Edge) Tag() string { return e.tag }

// IsRequired returns is this edge is an optional edge.
func (e Edge) IsRequired() bool { return e.required }

func typ(t interface{}) string {
	if rt := reflect.TypeOf(t); rt.NumIn() > 0 {
		return rt.In(0).Name()
	}
	return ""
}

// assocBuilder is the builder for assoc edges.
type assocBuilder struct {
	*Edge
}

// Fields sets the fields of the edge.
func (b *assocBuilder) Fields(f ...*field.Field) *assocBuilder {
	b.fields = f
	return b
}

// Unique sets the edge type to be unique. Basically, it's limited the ent to be one of the two:
// one2one or one2many. one2one applied if the inverse-edge is also unique.
func (b *assocBuilder) Unique() *assocBuilder {
	b.unique = true
	return b
}

// Required indicates that this edge is a required field on creation.
// Unlike fields, edges are optional by default.
func (b *assocBuilder) Required() *assocBuilder {
	b.required = true
	return b
}

// StructTag sets the struct tag of the assoc edge.
func (b *assocBuilder) StructTag(s string) *assocBuilder {
	b.tag = s
	return b
}

// Assoc creates an inverse-edge with the same type.
func (b *assocBuilder) From(name string) *inverseBuilder {
	return &inverseBuilder{&Edge{name: name, typ: b.typ, inverse: true, parent: b.Edge}}
}

// Comment used to put annotations on the schema.
func (b *assocBuilder) Comment(string) *assocBuilder {
	return b
}

// assocBuilder is the builder for inverse edges.
type inverseBuilder struct {
	*Edge
}

// Ref sets the referenced-edge of this inverse edge.
func (b *inverseBuilder) Ref(ref string) *inverseBuilder {
	b.ref = ref
	return b
}

// Fields sets the fields of the edge.
func (b *inverseBuilder) Fields(f ...*field.Field) *inverseBuilder {
	b.fields = f
	return b
}

// Unique sets the edge type to be unique. Basically, it's limited the ent to be one of the two:
// one2one or one2many. one2one applied if the inverse-edge is also unique.
func (b *inverseBuilder) Unique() *inverseBuilder {
	b.unique = true
	return b
}

// Required indicates that this edge is a required field on creation.
// Unlike fields, edges are optional by default.
func (b *inverseBuilder) Required() *inverseBuilder {
	b.required = true
	return b
}

// StructTag sets the struct tag of the inverse edge.
func (b *inverseBuilder) StructTag(s string) *inverseBuilder {
	b.tag = s
	return b
}

// Comment used to put annotations on the schema.
func (b *inverseBuilder) Comment(string) *inverseBuilder {
	return b
}
