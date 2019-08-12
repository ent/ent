package load

import (
	"encoding/json"
	"reflect"

	"fbc/ent"
	"fbc/ent/schema/field"
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
	Edge   string   `json:"edge,omitempty"`
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
func MarshalSchema(schema ent.Schema) ([]byte, error) {
	s := &Schema{Name: indirect(reflect.TypeOf(schema)).Name()}
	for _, f := range schema.Fields() {
		sf := &Field{
			Name:       f.Name(),
			Type:       f.Type(),
			Tag:        f.Tag(),
			Unique:     f.IsUnique(),
			Nillable:   f.IsNillable(),
			Optional:   f.IsOptional(),
			Default:    f.HasDefault(),
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
	for _, e := range schema.Edges() {
		s.Edges = append(s.Edges, NewEdge(e))
	}
	for _, idx := range schema.Indexes() {
		s.Indexes = append(s.Indexes, &Index{
			Edge:   idx.Edge(),
			Fields: idx.Fields(),
			Unique: idx.IsUnique(),
		})
	}
	return json.Marshal(s)
}

func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
