// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package g

import "entgo.io/ent/dialect/gremlin/graph/dsl"

// V is the api for calling g.V().
func V(args ...any) *dsl.Traversal { return dsl.NewTraversal().V(args...) }

// E is the api for calling g.E().
func E(args ...any) *dsl.Traversal { return dsl.NewTraversal().E(args...) }

// AddV is the api for calling g.AddV().
func AddV(args ...any) *dsl.Traversal { return dsl.NewTraversal().AddV(args...) }

// AddE is the api for calling g.AddE().
func AddE(args ...any) *dsl.Traversal { return dsl.NewTraversal().AddE(args...) }
