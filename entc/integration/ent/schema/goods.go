// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "github.com/facebook/ent"

// Goods holds the schema definition for the Goods entity.
type Goods struct {
	ent.Schema
}

// Fields of the Goods.
func (Goods) Fields() []ent.Field {
	return nil
}

// Edges of the Goods.
func (Goods) Edges() []ent.Edge {
	return nil
}
