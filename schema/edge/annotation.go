// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package edge

import "github.com/facebook/ent/schema"

// Annotation is a builtin schema annotation for
// configuring the edges behavior in codegen.
type Annotation struct {
	// The StructTag option allows overriding the struct-tag
	// of the `Edges` field in the generated entity. For example:
	//
	//	edge.Annotation{
	//		StructTag: `json: "pet_edges"`
	//	}
	//
	StructTag string
}

// Name describes the annotation name.
func (Annotation) Name() string {
	return "Edges"
}

var _ schema.Annotation = (*Annotation)(nil)
