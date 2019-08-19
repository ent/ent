package __

import "github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"

// As is the api for calling __.As().
func As(args ...interface{}) *dsl.Traversal { return New().As(args...) }

// Is is the api for calling __.Is().
func Is(args ...interface{}) *dsl.Traversal { return New().Is(args...) }

// Not is the api for calling __.Not().
func Not(args ...interface{}) *dsl.Traversal { return New().Not(args...) }

// Has is the api for calling __.Has().
func Has(args ...interface{}) *dsl.Traversal { return New().Has(args...) }

// HasNot is the api for calling __.HasNot().
func HasNot(args ...interface{}) *dsl.Traversal { return New().HasNot(args...) }

// Or is the api for calling __.Or().
func Or(args ...interface{}) *dsl.Traversal { return New().Or(args...) }

// And is the api for calling __.And().
func And(args ...interface{}) *dsl.Traversal { return New().And(args...) }

// In is the api for calling __.In().
func In(args ...interface{}) *dsl.Traversal { return New().In(args...) }

// Out is the api for calling __.Out().
func Out(args ...interface{}) *dsl.Traversal { return New().Out(args...) }

// OutE is the api for calling __.OutE().
func OutE(args ...interface{}) *dsl.Traversal { return New().OutE(args...) }

// InE is the api for calling __.InE().
func InE(args ...interface{}) *dsl.Traversal { return New().InE(args...) }

// InV is the api for calling __.InV().
func InV(args ...interface{}) *dsl.Traversal { return New().InV(args...) }

// V is the api for calling __.V().
func V(args ...interface{}) *dsl.Traversal { return New().V(args...) }

// OutV is the api for calling __.OutV().
func OutV(args ...interface{}) *dsl.Traversal { return New().OutV(args...) }

// Values is the api for calling __.Values().
func Values(args ...string) *dsl.Traversal { return New().Values(args...) }

// OtherV is the api for calling __.OtherV().
func OtherV() *dsl.Traversal { return New().OtherV() }

// Count is the api for calling __.Count().
func Count() *dsl.Traversal { return New().Count() }

// Fold is the api for calling __.Fold().
func Fold() *dsl.Traversal { return New().Fold() }

func New() *dsl.Traversal { return new(dsl.Traversal).Add(dsl.Token("__")) }
