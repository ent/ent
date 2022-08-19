// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package __

import "entgo.io/ent/dialect/gremlin/graph/dsl"

// As is the api for calling __.As().
func As(args ...any) *dsl.Traversal { return New().As(args...) }

// Is is the api for calling __.Is().
func Is(args ...any) *dsl.Traversal { return New().Is(args...) }

// Not is the api for calling __.Not().
func Not(args ...any) *dsl.Traversal { return New().Not(args...) }

// Has is the api for calling __.Has().
func Has(args ...any) *dsl.Traversal { return New().Has(args...) }

// HasNot is the api for calling __.HasNot().
func HasNot(args ...any) *dsl.Traversal { return New().HasNot(args...) }

// Or is the api for calling __.Or().
func Or(args ...any) *dsl.Traversal { return New().Or(args...) }

// And is the api for calling __.And().
func And(args ...any) *dsl.Traversal { return New().And(args...) }

// In is the api for calling __.In().
func In(args ...any) *dsl.Traversal { return New().In(args...) }

// Out is the api for calling __.Out().
func Out(args ...any) *dsl.Traversal { return New().Out(args...) }

// OutE is the api for calling __.OutE().
func OutE(args ...any) *dsl.Traversal { return New().OutE(args...) }

// InE is the api for calling __.InE().
func InE(args ...any) *dsl.Traversal { return New().InE(args...) }

// InV is the api for calling __.InV().
func InV(args ...any) *dsl.Traversal { return New().InV(args...) }

// V is the api for calling __.V().
func V(args ...any) *dsl.Traversal { return New().V(args...) }

// OutV is the api for calling __.OutV().
func OutV(args ...any) *dsl.Traversal { return New().OutV(args...) }

// Values is the api for calling __.Values().
func Values(args ...string) *dsl.Traversal { return New().Values(args...) }

// Union is the api for calling __.Union().
func Union(args ...any) *dsl.Traversal { return New().Union(args...) }

// Constant is the api for calling __.Constant().
func Constant(args ...any) *dsl.Traversal { return New().Constant(args...) }

// Properties is the api for calling __.Properties().
func Properties(args ...any) *dsl.Traversal { return New().Properties(args...) }

// OtherV is the api for calling __.OtherV().
func OtherV() *dsl.Traversal { return New().OtherV() }

// Count is the api for calling __.Count().
func Count() *dsl.Traversal { return New().Count() }

// Drop is the api for calling __.Drop().
func Drop() *dsl.Traversal { return New().Drop() }

// Fold is the api for calling __.Fold().
func Fold() *dsl.Traversal { return New().Fold() }

func New() *dsl.Traversal { return new(dsl.Traversal).Add(dsl.Token("__")) }
