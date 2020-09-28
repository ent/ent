// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// FileType holds the schema definition for the FileType entity.
type FileType struct {
	ent.Schema
}

// Fields of the FileType.
func (FileType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Enum("type").
			NamedValues(
				"PNG", "png",
				"SVG", "svg",
				"JPG", "jpg",
			).
			Default("png"),
		field.Enum("state").
			NamedValues(
				"On", "ON",
				"Off", "OFF",
			).
			Default("ON"),
	}
}

// Edges of the FileType.
func (FileType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type),
	}
}
