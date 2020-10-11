// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field

import "github.com/facebook/ent/schema"

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

// StructTag returns an annotation with the configured struct-tag mapping.
//
//	StructTag( "id", `json:"-" yaml:"-"`, "oid", `json:"oid,omitempty"`)
//
func StructTag(nametag ...string) Annotation {
	if len(nametag)%2 == 1 {
		panic("field.StructTag: odd argument count")
	}
	tags := make(map[string]string, len(nametag)/2)
	for i := 0; i < len(nametag); i += 2 {
		tags[nametag[i]] = nametag[i+1]
	}
	return Annotation{StructTag: tags}
}

var _ schema.Annotation = (*Annotation)(nil)
