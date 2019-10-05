// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Schema represents an ent.Schema that was loaded from a complied user package.
type Schema struct {
	Name         string         `json:"name,omitempty"`
	Config       ent.Config     `json:"config,omitempty"`
	Edges        []*Edge        `json:"edges,omitempty"`
	Fields       []*Field       `json:"fields,omitempty"`
	Indexes      []*Index       `json:"indexes,omitempty"`
	StructFields []*StructField `json:"struct_fields,omitempty"`
}

// Position describes a field position in the schema.
type Position struct {
	Index      int  // field index in the field list.
	MixedIn    bool // indicates if the field was mixed-in.
	MixinIndex int  // mixin index in the mixin list.
}

// Field represents an ent.Field that was loaded from a complied user package.
type Field struct {
	Name          string          `json:"name,omitempty"`
	Info          *field.TypeInfo `json:"type,omitempty"`
	Tag           string          `json:"tag,omitempty"`
	Size          *int64          `json:"size,omitempty"`
	Enums         []string        `json:"enums,omitempty"`
	Unique        bool            `json:"unique,omitempty"`
	Nillable      bool            `json:"nillable,omitempty"`
	Optional      bool            `json:"optional,omitempty"`
	Default       bool            `json:"default,omitempty"`
	UpdateDefault bool            `json:"update_default,omitempty"`
	Immutable     bool            `json:"immutable,omitempty"`
	Validators    int             `json:"validators,omitempty"`
	StorageKey    string          `json:"storage_key,omitempty"`
	Position      *Position       `json:"position,omitempty"`
}

// StructField represents an external struct field defined in the schema.
type StructField struct {
	Tag      string `json:"tag,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Comment  string `json:"comment,omitempty"`
	PkgPath  string `json:"pkg_path,omitempty"`
	Embedded bool   `json:"embedded,omitempty"`
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

// NewEdge creates an loaded edge from edge descriptor.
func NewEdge(ed *edge.Descriptor) *Edge {
	ne := &Edge{
		Tag:      ed.Tag,
		Type:     ed.Type,
		Name:     ed.Name,
		Unique:   ed.Unique,
		Inverse:  ed.Inverse,
		Required: ed.Required,
		RefName:  ed.RefName,
	}
	if ref := ed.Ref; ref != nil {
		ne.Ref = NewEdge(ref)
	}
	return ne
}

// NewField creates an loaded field from edge descriptor.
func NewField(fd *field.Descriptor) (*Field, error) {
	sf := &Field{
		Name:          fd.Name,
		Info:          fd.Info,
		Tag:           fd.Tag,
		Enums:         fd.Enums,
		Unique:        fd.Unique,
		Nillable:      fd.Nillable,
		Optional:      fd.Optional,
		Immutable:     fd.Immutable,
		StorageKey:    fd.StorageKey,
		Validators:    len(fd.Validators),
		Default:       fd.Default != nil,
		UpdateDefault: fd.UpdateDefault != nil,
	}
	if sf.Info == nil {
		return nil, fmt.Errorf("missing type info for field %q", sf.Name)
	}
	if size := int64(fd.Size); size != 0 {
		sf.Size = &size
	}
	return sf, nil
}

// MarshalSchema encode the ent.Schema interface into a JSON
// that can be decoded into the Schema object object.
func MarshalSchema(schema ent.Interface) (b []byte, err error) {
	s := &Schema{
		Config: schema.Config(),
		Name:   indirect(reflect.TypeOf(schema)).Name(),
	}
	if err := s.loadFields(schema); err != nil {
		return nil, err
	}
	edges, err := safeEdges(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for _, e := range edges {
		s.Edges = append(s.Edges, NewEdge(e.Descriptor()))
	}
	indexes, err := safeIndexes(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for _, idx := range indexes {
		idx := idx.Descriptor()
		s.Indexes = append(s.Indexes, &Index{
			Edges:  idx.Edges,
			Fields: idx.Fields,
			Unique: idx.Unique,
		})
	}
	return json.Marshal(s)
}

// loadFields loads field to schema from ent.Interface.
func (s *Schema) loadFields(schema ent.Interface) error {
	mixin, err := safeMixin(schema)
	if err != nil {
		return fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for i, mx := range mixin {
		fields, err := safeFields(mx)
		if err != nil {
			return fmt.Errorf("schema %q: %v", s.Name, err)
		}
		for j, f := range fields {
			sf, err := NewField(f.Descriptor())
			if err != nil {
				return fmt.Errorf("schema %q: %v", s.Name, err)
			}
			sf.Position = &Position{
				Index:      j,
				MixedIn:    true,
				MixinIndex: i,
			}
			s.Fields = append(s.Fields, sf)
		}
	}
	fields, err := safeFields(schema)
	if err != nil {
		return fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for i, f := range fields {
		sf, err := NewField(f.Descriptor())
		if err != nil {
			return fmt.Errorf("schema %q: %v", s.Name, err)
		}
		sf.Position = &Position{Index: i}
		s.Fields = append(s.Fields, sf)
	}
	return nil
}

// safeFields wraps the schema.Fields and mixin.Fields method with recover to ensure no panics in marshaling.
func safeFields(fd interface{ Fields() []ent.Field }) (fields []ent.Field, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("%T.Fields panics: %v", fd, v)
			fields = nil
		}
	}()
	return fd.Fields(), nil
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

// safeMixin wraps the schema.Mixin method with recover to ensure no panics in marshaling.
func safeMixin(schema ent.Interface) (mixin []ent.Mixin, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("schema.Mixin panics: %v", v)
			mixin = nil
		}
	}()
	return schema.Mixin(), nil
}

func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
