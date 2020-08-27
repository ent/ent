// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Street holds the schema definition for the Street entity.
type Street struct {
	ent.Schema
}

// Fields of the Street.
func (Street) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Street.
func (Street) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).
			Ref("streets").
			Unique(),
	}
}

// Indexes of the Street.
func (Street) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Edges("city").
			Unique(),
	}
}
