// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dsl

import (
	"fmt"
	"strings"
)

// Traversal mimics the TinkerPop graph traversal.
type Traversal struct {
	// nodes holds the dsl nodes. first element is the reference name
	// of the TinkerGraph. defaults to "g".
	nodes []Node
	errs  []error
}

// NewTraversal returns a new default traversal with "g" as a reference name to the Graph.
func NewTraversal() *Traversal {
	return &Traversal{nodes: []Node{G}}
}

// Group groups a list of traversals into one. all traversals are assigned into a temporary
// variables named by their index. The last variable functions as a return value of the query.
// Note that, this "temporary hack" is not perfect and may not work in some cases because of
// the limitation of evaluation order.
func Group(trs ...*Traversal) *Traversal {
	var (
		b     = Block{}
		names = make(map[*Traversal]Token)
	)
	for i, tr := range trs {
		if _, ok := names[tr]; ok {
			continue
		}
		v := &Var{Name: fmt.Sprintf("t%d", i), Elem: &Traversal{nodes: tr.nodes}}
		b.Nodes = append(b.Nodes, v)
		names[tr] = Token(v.Name)
	}
	for _, tr := range trs {
		tr.nodes = []Node{names[tr]}
	}
	b.Nodes = append(b.Nodes, names[trs[len(trs)-1]])
	return &Traversal{nodes: []Node{b}}
}

// Join joins a list of traversals with a semicolon separator.
func Join(trs ...*Traversal) *Traversal {
	b := Block{}
	for _, tr := range trs {
		b.Nodes = append(b.Nodes, &Traversal{nodes: tr.nodes})
	}
	return &Traversal{nodes: []Node{b}}
}

// AddError adds an error to the traversal.
func (t *Traversal) AddError(err error) *Traversal {
	t.errs = append(t.errs, err)
	return t
}

// Err returns a concatenated error of all errors encountered during
// the query-building, or were added manually by calling AddError.
func (t *Traversal) Err() error {
	if len(t.errs) == 0 {
		return nil
	}
	br := strings.Builder{}
	for i := range t.errs {
		if i > 0 {
			br.WriteString("; ")
		}
		br.WriteString(t.errs[i].Error())
	}
	return fmt.Errorf(br.String())
}

// V step is usually used to start a traversal but it may also be used mid-traversal.
func (t *Traversal) V(args ...any) *Traversal {
	t.Add(Dot, NewFunc("V", args...))
	return t
}

// OtherV maps the Edge to the incident vertex that was not just traversed from in the path history.
func (t *Traversal) OtherV() *Traversal {
	t.Add(Dot, NewFunc("otherV"))
	return t
}

// E step is usually used to start a traversal but it may also be used mid-traversal.
func (t *Traversal) E(args ...any) *Traversal {
	t.Add(Dot, NewFunc("E", args...))
	return t
}

// AddV adds a vertex.
func (t *Traversal) AddV(args ...any) *Traversal {
	t.Add(Dot, NewFunc("addV", args...))
	return t
}

// AddE adds an edge.
func (t *Traversal) AddE(args ...any) *Traversal {
	t.Add(Dot, NewFunc("addE", args...))
	return t
}

// Next gets the next n-number of results from the traversal.
func (t *Traversal) Next() *Traversal {
	return t.Add(Dot, NewFunc("next"))
}

// Drop removes elements and properties from the graph.
func (t *Traversal) Drop() *Traversal {
	return t.Add(Dot, NewFunc("drop"))
}

// Property sets a Property value and related meta properties if supplied,
// if supported by the Graph and if the Element is a VertexProperty.
func (t *Traversal) Property(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("property", args...))
}

// Both maps the Vertex to its adjacent vertices given the edge labels.
func (t *Traversal) Both(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("both", args...))
}

// BothE maps the Vertex to its incident edges given the edge labels.
func (t *Traversal) BothE(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("bothE", args...))
}

// Has filters vertices, edges and vertex properties based on their properties.
// See: http://tinkerpop.apache.org/docs/current/reference/#has-step.
func (t *Traversal) Has(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("has", args...))
}

// HasNot filters vertices, edges and vertex properties based on the non-existence of properties.
// See: http://tinkerpop.apache.org/docs/current/reference/#has-step.
func (t *Traversal) HasNot(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("hasNot", args...))
}

// HasID filters vertices, edges and vertex properties based on their identifier.
func (t *Traversal) HasID(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("hasId", args...))
}

// HasLabel filters vertices, edges and vertex properties based on their label.
func (t *Traversal) HasLabel(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("hasLabel", args...))
}

// HasNext returns true if the iteration has more elements.
func (t *Traversal) HasNext() *Traversal {
	return t.Add(Dot, NewFunc("hasNext"))
}

// Match maps the Traverser to a Map of bindings as specified by the provided match traversals.
func (t *Traversal) Match(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("match", args...))
}

// Choose routes the current traverser to a particular traversal branch option which allows the creation of if-then-else like semantics within a traversal.
func (t *Traversal) Choose(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("choose", args...))
}

// Select arbitrary values from the traversal.
func (t *Traversal) Select(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("select", args...))
}

// Group organizes objects in the stream into a Map.Calls to group() are typically accompanied with by() modulators which help specify how the grouping should occur.
func (t *Traversal) Group() *Traversal {
	return t.Add(Dot, NewFunc("group"))
}

// Values maps the Element to the values of the associated properties given the provide property keys.
func (t *Traversal) Values(args ...string) *Traversal {
	return t.Add(Dot, NewFunc("values", sface(args)...))
}

// ValueMap maps the Element to a Map of the property values key'd according to their Property.key().
func (t *Traversal) ValueMap(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("valueMap", args...))
}

// Properties maps the Element to its associated properties given the provide property keys.
func (t *Traversal) Properties(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("properties", args...))
}

// Range filters the objects in the traversal by the number of them to pass through the stream.
func (t *Traversal) Range(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("range", args...))
}

// Limit filters the objects in the traversal by the number of them to pass through the stream, where only the first n objects are allowed as defined by the limit argument.
func (t *Traversal) Limit(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("limit", args...))
}

// ID maps the Element to its Element.id().
func (t *Traversal) ID() *Traversal {
	return t.Add(Dot, NewFunc("id"))
}

// Label maps the Element to its Element.label().
func (t *Traversal) Label() *Traversal {
	return t.Add(Dot, NewFunc("label"))
}

// From provides from()-modulation to respective steps.
func (t *Traversal) From(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("from", args...))
}

// To used as a modifier to addE(String) this method specifies the traversal to use for selecting the incoming vertex of the newly added Edge.
func (t *Traversal) To(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("to", args...))
}

// As provides a label to the step that can be accessed later in the traversal by other steps.
func (t *Traversal) As(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("as", args...))
}

// Or ensures that at least one of the provided traversals yield a result.
func (t *Traversal) Or(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("or", args...))
}

// And ensures that all of the provided traversals yield a result.
func (t *Traversal) And(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("and", args...))
}

// Is filters the E object if it is not P.eq(V) to the provided value.
func (t *Traversal) Is(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("is", args...))
}

// Not removes objects from the traversal stream when the traversal provided as an argument does not return any objects.
func (t *Traversal) Not(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("not", args...))
}

// In maps the Vertex to its incoming adjacent vertices given the edge labels.
func (t *Traversal) In(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("in", args...))
}

// Where filters the current object based on the object itself or the path history.
func (t *Traversal) Where(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("where", args...))
}

// Out maps the Vertex to its outgoing adjacent vertices given the edge labels.
func (t *Traversal) Out(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("out", args...))
}

// OutE maps the Vertex to its outgoing incident edges given the edge labels.
func (t *Traversal) OutE(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("outE", args...))
}

// InE maps the Vertex to its incoming incident edges given the edge labels.
func (t *Traversal) InE(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("inE", args...))
}

// OutV maps the Edge to its outgoing/tail incident Vertex.
func (t *Traversal) OutV(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("outV", args...))
}

// InV maps the Edge to its incoming/head incident Vertex.
func (t *Traversal) InV(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("inV", args...))
}

// ToList puts all the results into a Groovy list.
func (t *Traversal) ToList() *Traversal {
	return t.Add(Dot, NewFunc("toList"))
}

// Iterate iterates the traversal presumably for the generation of side-effects.
func (t *Traversal) Iterate() *Traversal {
	return t.Add(Dot, NewFunc("iterate"))
}

// Count maps the traversal stream to its reduction as a sum of the Traverser.bulk() values
// (i.e. count the number of traversers up to this point).
func (t *Traversal) Count(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("count", args...))
}

// Order all the objects in the traversal up to this point and then emit them one-by-one in their ordered sequence.
func (t *Traversal) Order(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("order", args...))
}

// By can be applied to a number of different step to alter their behaviors.
// This form is essentially an identity() modulation.
func (t *Traversal) By(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("by", args...))
}

// Fold rolls up objects in the stream into an aggregate list..
func (t *Traversal) Fold() *Traversal {
	return t.Add(Dot, NewFunc("fold"))
}

// Unfold unrolls a Iterator, Iterable or Map into a linear form or simply emits the object if it is not one of those types.
func (t *Traversal) Unfold() *Traversal {
	return t.Add(Dot, NewFunc("unfold"))
}

// Sum maps the traversal stream to its reduction as a sum of the Traverser.get() values multiplied by their Traverser.bulk().
func (t *Traversal) Sum(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("sum", args...))
}

// Mean determines the mean value in the stream.
func (t *Traversal) Mean(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("mean", args...))
}

// Min determines the smallest value in the stream.
func (t *Traversal) Min(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("min", args...))
}

// Max determines the greatest value in the stream.
func (t *Traversal) Max(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("max", args...))
}

// Coalesce evaluates the provided traversals and returns the result of the first traversal to emit at least one object.
func (t *Traversal) Coalesce(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("coalesce", args...))
}

// Dedup removes all duplicates in the traversal stream up to this point.
func (t *Traversal) Dedup(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("dedup", args...))
}

// Constant maps any object to a fixed E value.
func (t *Traversal) Constant(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("constant", args...))
}

// Union merges the results of an arbitrary number of traversals.
func (t *Traversal) Union(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("union", args...))
}

// SideEffect allows the traverser to proceed unchanged, but yield some computational
// sideEffect in the process.
func (t *Traversal) SideEffect(args ...any) *Traversal {
	return t.Add(Dot, NewFunc("sideEffect", args...))
}

// Each is a Groovy each-loop function.
func Each(v any, cb func(it *Traversal) *Traversal) *Traversal {
	t := &Traversal{}
	switch v := v.(type) {
	case *Traversal:
		t.Add(&Var{Elem: v})
	case []any:
		t.Add(NewList(v...))
	default:
		t.Add(Token("undefined"))
	}
	t.Add(Dot, Token("each"), Token(" { "))
	t.Add(cb(&Traversal{nodes: []Node{Token("it")}}).nodes...)
	t.Add(Token(" }"))
	return t
}

// Add is the public API for adding new nodes to the traversal by its sub packages.
func (t *Traversal) Add(n ...Node) *Traversal {
	t.nodes = append(t.nodes, n...)
	return t
}

// Query returns the query-representation and its binding of this traversal object.
func (t *Traversal) Query() (string, Bindings) {
	var (
		names    []any
		query    strings.Builder
		bindings = Bindings{}
	)
	for _, n := range t.nodes {
		code, args := n.Code()
		query.WriteString(code)
		for _, arg := range args {
			names = append(names, bindings.Add(arg))
		}
	}
	return fmt.Sprintf(query.String(), names...), bindings
}

// Clone creates a deep copy of an existing traversal.
func (t *Traversal) Clone() *Traversal {
	if t == nil {
		return nil
	}
	return &Traversal{nodes: append(make([]Node, 0, len(t.nodes)), t.nodes...)}
}

// Undo reverts the last-step of the traversal.
func (t *Traversal) Undo() *Traversal {
	if n := len(t.nodes); n > 2 {
		t.nodes = t.nodes[:n-2]
	}
	return t
}
