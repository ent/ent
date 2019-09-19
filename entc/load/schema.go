// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/facebookincubator/ent/schema/edge"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Schema represents an ent.Schema that was loaded from a complied user package.
type Schema struct {
	Name         string         `json:"name,omitempty"`
	Edges        []*Edge        `json:"edges,omitempty"`
	Fields       []*Field       `json:"fields,omitempty"`
	Indexes      []*Index       `json:"indexes,omitempty"`
	StructFields []*StructField `json:"struct_fields,omitempty"`
}

// Field represents an ent.Field that was loaded from a complied user package.
type Field struct {
	Name          string          `json:"name,omitempty"`
	Info          *field.TypeInfo `json:"type,omitempty"`
	Tag           string          `json:"tag,omitempty"`
	Size          *int            `json:"size,omitempty"`
	Unique        bool            `json:"unique,omitempty"`
	Nillable      bool            `json:"nillable,omitempty"`
	Optional      bool            `json:"optional,omitempty"`
	Default       bool            `json:"default,omitempty"`
	UpdateDefault bool            `json:"update_default,omitempty"`
	Immutable     bool            `json:"immutable,omitempty"`
	Validators    int             `json:"validators,omitempty"`
	StorageKey    string          `json:"storage_key,omitempty"`
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

// MarshalSchema encode the ent.Schema interface into a JSON
// that can be decoded into the Schema object object.
func MarshalSchema(schema ent.Interface) (b []byte, err error) {
	s := &Schema{Name: indirect(reflect.TypeOf(schema)).Name()}
	fields, err := safeFields(schema)
	if err != nil {
		return nil, fmt.Errorf("schema %q: %v", s.Name, err)
	}
	for _, f := range fields {
		fd := f.Descriptor()
		sf := &Field{
			Name:          fd.Name,
			Info:          fd.Info,
			Tag:           fd.Tag,
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
			return nil, fmt.Errorf("schema %q: missing type info for field %q", s.Name, sf.Name)
		}
		if fd.Size != 0 {
			sf.Size = &fd.Size
		}
		s.Fields = append(s.Fields, sf)
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
