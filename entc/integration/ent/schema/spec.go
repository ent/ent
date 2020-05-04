// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type Spec struct {
	ent.Schema
}

// Edges of the Spec.
func (Spec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type),
	}
}
