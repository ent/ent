// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/privacy"

	"github.com/facebookincubator/ent/entc/integration/privacy/rule"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
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
