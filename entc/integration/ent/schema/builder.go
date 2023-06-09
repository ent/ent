// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "entgo.io/ent"

// Builder holds the schema definition for the Builder entity.
type Builder struct {
	ent.Schema
}

// Fields of the Builder.
func (Builder) Fields() []ent.Field {
	return nil
}

// Edges of the Builder.
func (Builder) Edges() []ent.Edge {
	return nil
}
