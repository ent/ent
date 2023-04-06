// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package failure

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type User struct {
	ent.Schema
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panic", User{}.Type),
	}
}
