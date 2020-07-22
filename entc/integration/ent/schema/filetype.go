// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
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
			ValueMap(map[string]string{
				"PNG": "png",
				"SVG": "svg",
				"JPG": "jpg",
			}).
			Default("png"),
	}
}

// Edges of the FileType.
func (FileType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type),
	}
}
