// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent/entc/integration/privacy/ent/privacy"

	"github.com/facebook/ent/entc/integration/privacy/rule"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Galaxy defines the schema of a galaxy.
type Galaxy struct {
	ent.Schema
}

// Fields of the galaxy.
func (Galaxy) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique(),
		field.Enum("type").
			Values("spiral", "barred_spiral", "elliptical", "irregular"),
	}
}

// Edges of the galaxy.
func (Galaxy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("planets", Planet.Type),
	}
}

func (Galaxy) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterIrregularGalaxyRule(),
			privacy.AlwaysAllowRule(),
		},
	}
}
