// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RelationshipInfo holds the schema definition for the RelationshipInfo entity.
type RelationshipInfo struct {
	ent.Schema
}

// Fields of the RelationshipInfo.
func (RelationshipInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
	}
}
