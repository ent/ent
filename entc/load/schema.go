package load

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Schema represents an ent.Schema that was loaded from a complied user package.
type Schema struct {
	Name    string   `json:"name,omitempty"`
	Edges   []*Edge  `json:"edges,omitempty"`
	Fields  []*Field `json:"fields,omitempty"`
	Indexes []*Index `json:"indexes,omitempty"`
}

// Field represents an ent.Field that was loaded from a complied user package.
type Field struct {
	Name       string     `json:"name,omitempty"`
	Type       field.Type `json:"type,omitempty"`
	Tag        string     `json:"tag,omitempty"`
	Size       *int       `json:"size,omitempty"`
	Charset    *string    `json:"charset,omitempty"`
	Unique     bool       `json:"unique,omitempty"`
	Nillable   bool       `json:"nillable,omitempty"`
	Optional   bool       `json:"optional,omitempty"`
	Default    bool       `json:"default,omitempty"`
	Immutable  bool       `json:"immutable,omitempty"`
	Validators int        `json:"validators,omitempty"`
}

// Edge represents an ent.Edge that was loaded from a complied user package.
type Edge struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Tag      string `json:"tag,omitempty"`
	RefName  string `json:"ref_name,omitempty"`
	Ref      *Edge  `json:"ref,omitempty"`
	Unique   bool   `json:"unique,omitempty"`
	Inverse  bool   `json:"inverse,omitempty"`
	Required bool   `json:"required,omitempty"`
}

// Index represents an ent.Index that was loaded from a complied user package.
type Index struct {
	Unique bool     `json:"unique,omitempty"`
	Edges  []string `json:"edges,omitempty"`
	Fields []string `json:"fields,omitempty"`
}

// NewEdge creates an loaded edge from schema interface.
func NewEdge(e ent.Edge) *Edge {
	ne := &Edge{
		Name:     e.Name(),
		Type:     e.Type(),
		Tag:      e.Tag(),
		RefName:  e.RefName(),
		Unique:   e.IsUnique(),
		Inverse:  e.IsInverse(),
		Required: e.IsRequired(),
	}
	if e := e.Assoc(); e != nil {
		ne.Ref = NewEdge(e)
	}
	return ne
}

// MarshalSchema encode the ent.Schema interface into a JSON
// that can be decoded into the Schema object object.
func MarshalSchema(schema ent.Interface) (b []byte, err error) {
	s := &Schema{Name: indirect(reflect.TypeOf(schema)).Name()}
	fields, err := safeFields(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for _, f := range fields {
		sf := &Field{
			Name:       f.Name(),
			Type:       f.Type(),
			Tag:        f.Tag(),
			Unique:     f.IsUnique(),
			Default:    f.HasDefault(),
			Nillable:   f.IsNillable(),
			Optional:   f.IsOptional(),
			Immutable:  f.IsImmutable(),
			Validators: len(f.Validators()),
		}
		if s, ok := f.(field.Sizer); ok {
			size := s.Size()
			sf.Size = &size
		}
		if c, ok := f.(field.Charseter); ok {
			charset := c.Charset()
			sf.Charset = &charset
		}
		s.Fields = append(s.Fields, sf)
	}
	edges, err := safeEdges(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for _, e := range edges {
		s.Edges = append(s.Edges, NewEdge(e))
	}
	indexes, err := safeIndexes(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for _, idx := range indexes {
		s.Indexes = append(s.Indexes, &Index{
			Edges:  idx.EdgeNames(),
			Fields: idx.FieldNames(),
			Unique: idx.IsUnique(),
		})
	}
	return json.Marshal(s)
}

// safeFields wraps the schema.Fields method with recover to ensure no panics in marshaling.
func safeFields(schema ent.Interface) (fields []ent.Field, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("schema.Fields panics: %v", v)
			fields = nil
		}
	}()
	return schema.Fields(), nil
}

// safeEdges wraps the schema.Edges method with recover to ensure no panics in marshaling.
func safeEdges(schema ent.Interface) (edges []ent.Edge, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("schema.Edges panics: %v", v)
			edges = nil
		}
	}()
	return schema.Edges(), nil
}

// safeIndexes wraps the schema.Indexes method with recover to ensure no panics in marshaling.
func safeIndexes(schema ent.Interface) (indexes []ent.Index, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("schema.Indexes panics: %v", v)
			indexes = nil
		}
	}()
	return schema.Indexes(), nil
}

func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
