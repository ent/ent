// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entql"
)

type (
	// A Schema holds a representation of ent/schema at runtime. Each Node
	// represents a single schema-type and its relations in the graph (storage).
	//
	// It is used for translating common graph traversal operations to the
	// underlying SQL storage. For example, an operation like `has_edge(E)`,
	// will be translated to an SQL lookup based on the relation type and the
	// FK configuration.
	Schema struct {
		Nodes []*Node
	}

	// A Node in the graph holds the SQL information for an ent/schema.
	Node struct {
		NodeSpec

		// Type holds the node type (schema name).
		Type string

		// Fields maps from field names to their spec.
		Fields map[string]*FieldSpec

		// Edges maps from edge names to their spec.
		Edges map[string]struct {
			To   *Node
			Spec *EdgeSpec
		}
	}
)

// AddEdge adds an edge to the graph. It fails, if one of the node
// types is missing.
//
//	g.AddE("pets", spec, "user", "pet")
//	g.AddE("friends", spec, "user", "user")
//
func (g *Schema) AddE(name string, spec *EdgeSpec, from, to string) error {
	var fromT, toT *Node
	for i := range g.Nodes {
		t := g.Nodes[i].Type
		if t == from {
			fromT = g.Nodes[i]
		}
		if t == to {
			toT = g.Nodes[i]
		}
	}
	if fromT == nil || toT == nil {
		return fmt.Errorf("from/to type was not found")
	}
	if fromT.Edges == nil {
		fromT.Edges = make(map[string]struct {
			To   *Node
			Spec *EdgeSpec
		})
	}
	fromT.Edges[name] = struct {
		To   *Node
		Spec *EdgeSpec
	}{
		To:   toT,
		Spec: spec,
	}
	return nil
}

// MustAddE is like AddE but panics if the edge can be added to the graph.
func (g *Schema) MustAddE(name string, spec *EdgeSpec, from, to string) {
	if err := g.AddE(name, spec, from, to); err != nil {
		panic(err)
	}
}

// EvalP evaluates the entql predicate on the given selector (query builder).
func (g *Schema) EvalP(nodeType string, p entql.P, selector *sql.Selector) error {
	var node *Node
	for i := range g.Nodes {
		if g.Nodes[i].Type == nodeType {
			node = g.Nodes[i]
			break
		}
	}
	if node == nil {
		return fmt.Errorf("node %s was not found in the graph schema", nodeType)
	}
	pr, err := evalExpr(node, selector, p)
	if err != nil {
		return err
	}
	selector.Where(pr)
	return nil
}

// FuncSelector represents a selector function to be used as an entql foreign-function.
const FuncSelector entql.Func = "func_selector"

// wrappedFunc wraps the selector-function to an ent-expression.
type wrappedFunc struct {
	entql.Expr
	Func func(*sql.Selector)
}

// WrapFunc wraps a selector-func with an entql call expression.
func WrapFunc(s func(*sql.Selector)) *entql.CallExpr {
	return &entql.CallExpr{
		Func: FuncSelector,
		Args: []entql.Expr{wrappedFunc{Func: s}},
	}
}

var (
	binary = [...]sql.Op{
		entql.OpEQ:    sql.OpEQ,
		entql.OpNEQ:   sql.OpNEQ,
		entql.OpGT:    sql.OpGT,
		entql.OpGTE:   sql.OpGTE,
		entql.OpLT:    sql.OpLT,
		entql.OpLTE:   sql.OpLTE,
		entql.OpIn:    sql.OpIn,
		entql.OpNotIn: sql.OpNotIn,
	}
	nary = [...]func(...*sql.Predicate) *sql.Predicate{
		entql.OpAnd: sql.And,
		entql.OpOr:  sql.Or,
	}
	strFunc = map[entql.Func]func(string, string) *sql.Predicate{
		entql.FuncContains:     sql.Contains,
		entql.FuncContainsFold: sql.ContainsFold,
		entql.FuncEqualFold:    sql.EqualFold,
		entql.FuncHasPrefix:    sql.HasPrefix,
		entql.FuncHasSuffix:    sql.HasSuffix,
	}
	nullFunc = [...]func(string) *sql.Predicate{
		entql.OpEQ:  sql.IsNull,
		entql.OpNEQ: sql.NotNull,
	}
)

// state represents the state of a predicate evaluation.
// Note that, the evaluation output is a predicate to be
// applied on the database.
type state struct {
	sql.Builder
	context  *Node
	selector *sql.Selector
}

// evalExpr evaluates the entql expression and returns a new SQL predicate to be applied on the database.
func evalExpr(context *Node, selector *sql.Selector, expr entql.Expr) (p *sql.Predicate, err error) {
	ex := &state{
		context:  context,
		selector: selector,
	}
	defer catch(&err)
	p = ex.evalExpr(expr)
	return
}

// evalExpr evaluates any expression.
func (e *state) evalExpr(expr entql.Expr) *sql.Predicate {
	switch expr := expr.(type) {
	case *entql.BinaryExpr:
		return e.evalBinary(expr)
	case *entql.UnaryExpr:
		return sql.Not(e.evalExpr(expr.X))
	case *entql.NaryExpr:
		ps := make([]*sql.Predicate, len(expr.Xs))
		for i, x := range expr.Xs {
			ps[i] = e.evalExpr(x)
		}
		return nary[expr.Op](ps...)
	case *entql.CallExpr:
		switch expr.Func {
		case entql.FuncHasPrefix, entql.FuncHasSuffix, entql.FuncContains, entql.FuncEqualFold, entql.FuncContainsFold:
			expect(len(expr.Args) == 2, "invalid number of arguments for %s", expr.Func)
			f, ok := expr.Args[0].(*entql.Field)
			expect(ok, "*entql.Field, got %T", expr.Args[0])
			v, ok := expr.Args[1].(*entql.Value)
			expect(ok, "*entql.Value, got %T", expr.Args[1])
			s, ok := v.V.(string)
			expect(ok, "string value, got %T", v.V)
			return strFunc[expr.Func](e.field(f), s)
		case entql.FuncHasEdge:
			expect(len(expr.Args) > 0, "invalid number of arguments for %s", expr.Func)
			edge, ok := expr.Args[0].(*entql.Edge)
			expect(ok, "*entql.Edge, got %T", expr.Args[0])
			return e.evalEdge(edge.Name, expr.Args[1:]...)
		}
	}
	panic("invalid")
}

// evalBinary evaluates binary expressions.
func (e *state) evalBinary(expr *entql.BinaryExpr) *sql.Predicate {
	switch expr.Op {
	case entql.OpOr:
		return sql.Or(e.evalExpr(expr.X), e.evalExpr(expr.Y))
	case entql.OpAnd:
		return sql.And(e.evalExpr(expr.X), e.evalExpr(expr.Y))
	case entql.OpEQ, entql.OpNEQ:
		if expr.Y == (*entql.Value)(nil) {
			f, ok := expr.X.(*entql.Field)
			expect(ok, "*entql.Field, got %T", expr.Y)
			return nullFunc[expr.Op](e.field(f))
		}
		fallthrough
	default:
		field, ok := expr.X.(*entql.Field)
		expect(ok, "expr.X to be *entql.Field (got %T)", expr.X)
		_, ok = expr.Y.(*entql.Field)
		if !ok {
			_, ok = expr.Y.(*entql.Value)
		}
		expect(ok, "expr.Y to be *entql.Field or *entql.Value (got %T)", expr.X)
		return sql.P(func(b *sql.Builder) {
			b.Ident(e.field(field))
			b.WriteOp(binary[expr.Op])
			switch x := expr.Y.(type) {
			case *entql.Field:
				b.Ident(e.field(x))
			case *entql.Value:
				args(b, x)
			}
		})
	}
}

// evalEdge evaluates has-edge and has-edge-with calls.
func (e *state) evalEdge(name string, exprs ...entql.Expr) *sql.Predicate {
	edge, ok := e.context.Edges[name]
	expect(ok, "edge %q was not found for node %q", name, e.context.Type)
	step := NewStep(
		From(e.context.Table, e.context.ID.Column),
		To(edge.To.Table, edge.To.ID.Column),
		Edge(edge.Spec.Rel, edge.Spec.Inverse, edge.Spec.Table, edge.Spec.Columns...),
	)
	selector := e.selector.Clone().SetP(nil)
	selector.SetTotal(e.Total())
	if len(exprs) == 0 {
		HasNeighbors(selector, step)
		return selector.P()
	}
	HasNeighborsWith(selector, step, func(s *sql.Selector) {
		for _, expr := range exprs {
			if cx, ok := expr.(*entql.CallExpr); ok && cx.Func == FuncSelector {
				expect(len(cx.Args) == 1, "invalid number of arguments for %s", FuncSelector)
				wrapped, ok := cx.Args[0].(wrappedFunc)
				expect(ok, "invalid argument for %s: %T", FuncSelector, cx.Args[0])
				wrapped.Func(s)
			} else {
				p, err := evalExpr(edge.To, s, expr)
				expect(err == nil, "edge evaluation failed for %s->%s: %s", e.context.Type, name, err)
				s.Where(p)
			}
		}
	})
	return selector.P()
}

func (e *state) field(f *entql.Field) string {
	_, ok := e.context.Fields[f.Name]
	expect(ok || e.context.ID.Column == f.Name, "field %q was not found for node %q", f.Name, e.context.Type)
	return f.Name
}

func args(b *sql.Builder, v *entql.Value) {
	vs, ok := v.V.([]interface{})
	if !ok {
		b.Arg(v.V)
		return
	}
	b.Args(vs...)
}

// expect panics if the condition is false.
func expect(cond bool, msg string, args ...interface{}) {
	if !cond {
		panic(evalError{fmt.Sprintf("expect "+msg, args...)})
	}
}

type evalError struct {
	msg string
}

func (p evalError) Error() string {
	return fmt.Sprintf("sqlgraph: %s", p.msg)
}

func catch(err *error) {
	if e := recover(); e != nil {
		xerr, ok := e.(evalError)
		if !ok {
			panic(e)
		}
		*err = xerr
	}
}
