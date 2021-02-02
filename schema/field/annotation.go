// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field

import "entgo.io/ent/schema"

// Annotation is a builtin schema annotation for
// configuring the schema fields in codegen.
type Annotation struct {
	// The StructTag option allows overriding the struct-tag
	// of the fields in the generated entity. For example:
	//
	//	field.Annotation{
	//		StructTag: map[string]string{
	//			"id": `json:"id,omitempty" yaml:"-"`
	//		}
	//	}
	//
	StructTag map[string]string
}

// Name describes the annotation name.
func (Annotation) Name() string {
	return "Fields"
}

// Merge implements the schema.Merger interface.
func (a Annotation) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotation
	switch other := other.(type) {
	case Annotation:
		ant = other
	case *Annotation:
		if other != nil {
			ant = *other
		}
	default:
		return a
	}
	for k, v := range ant.StructTag {
		if a.StructTag == nil {
			a.StructTag = make(map[string]string)
		}
		a.StructTag[k] = v
	}
	return a
}

var (
	_ schema.Annotation = (*Annotation)(nil)
	_ schema.Merger     = (*Annotation)(nil)
)
