package sqlgraph

import (
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entql"
)

type (
	// A Graph holds multiple ent/schemas and their relations in the graph.
	// It is used for translating common graph traversal operations to the
	// underlying SQL storage. For example, an operation like `has_edge(E)`,
	// will be translated to an SQL lookup based on the relation type and the
	// FK configuration.
	Graph struct {
		Nodes []*Node
	}

	// A Node in the graph holds the SQL information for an ent/schema.
	Node struct {
		NodeSpec

		// Type or label holds the
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
func (g *Graph) AddE(name string, spec *EdgeSpec, from, to string) error {
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

// EvalP evaluates the entql predicate on the query builder.
func (g *Graph) EvalP(nodeType string, p entql.P, selector *sql.Selector) error {
	var node *Node
	for i := range g.Nodes {
		if g.Nodes[i].Type == nodeType {
			node = g.Nodes[i]
			break
		}
	}
	if node == nil {
		return fmt.Errorf("node %s was not found in the graph", nodeType)
	}
	pr, err := execExpr(node, selector, p)
	if err != nil {
		return err
	}
	selector.Where(pr)
	return nil
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
	strFunc = [...]func(string, string) *sql.Predicate{
		entql.OpContains:     sql.Contains,
		entql.OpContainsFold: sql.ContainsFold,
		entql.OpEqualFold:    sql.EqualFold,
		entql.OpHasPrefix:    sql.HasPrefix,
		entql.OpHasSuffix:    sql.HasSuffix,
	}
)

type exec struct {
	sql.Builder
	context  *Node
	selector *sql.Selector
}

func execExpr(context *Node, selector *sql.Selector, expr entql.Expr) (p *sql.Predicate, err error) {
	ex := &exec{
		context:  context,
		selector: selector,
	}
	defer catch(&err)
	p = ex.evalExpr(expr)
	return
}

func (e *exec) evalExpr(expr entql.Expr) *sql.Predicate {
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
		switch expr.Op {
		case entql.OpHasPrefix, entql.OpHasSuffix, entql.OpContains, entql.OpEqualFold, entql.OpContainsFold:
			expect(len(expr.Args) == 2, "invalid number of arguments for %s", expr.Op)
			f, ok := expr.Args[0].(*entql.Field)
			expect(ok, "*entql.Field, got %T", expr.Args[0])
			v, ok := expr.Args[1].(*entql.Value)
			expect(ok, "*entql.Value, got %T", expr.Args[1])
			s, ok := v.V.(string)
			expect(ok, "string value, got %T", v.V)
			return strFunc[expr.Op](e.field(f), s)
		case entql.OpHasEdge:
			expect(len(expr.Args) > 0, "invalid number of arguments for %s", expr.Op)
			edge, ok := expr.Args[0].(*entql.Edge)
			expect(ok, "*entql.Edge, got %T", expr.Args[0])
			return e.evalEdge(edge.Name, expr.Args[1:]...)
		}
	}
	panic("invalid")
}

func (e *exec) evalBinary(expr *entql.BinaryExpr) *sql.Predicate {
	if (expr.Op == entql.OpEQ || expr.Op == entql.OpNEQ) && expr.Y == (*entql.Value)(nil) {
		f, ok := expr.X.(*entql.Field)
		expect(ok, "*entql.Field, got %T", expr.Y)
		if expr.Op == entql.OpEQ {
			return sql.IsNull(e.field(f))
		}
		return sql.NotNull(e.field(f))
	}
	switch expr.Op {
	case entql.OpOr:
		return sql.Or(e.evalExpr(expr.X), e.evalExpr(expr.Y))
	case entql.OpAnd:
		return sql.And(e.evalExpr(expr.X), e.evalExpr(expr.Y))
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

func (e *exec) field(f *entql.Field) string {
	_, ok := e.context.Fields[f.Name]
	expect(ok || e.context.ID.Column == f.Name, "field %q was not found for node %q", f.Name, e.context.Type)
	return f.Name
}

func (e *exec) evalEdge(name string, exprs ...entql.Expr) *sql.Predicate {
	edge, ok := e.context.Edges[name]
	expect(ok, "edge %q was not found for node %q", name, e.context.Type)
	step := NewStep(
		From(e.context.Table, e.context.ID.Column),
		To(edge.To.Table, edge.To.ID.Column),
		Edge(edge.Spec.Rel, edge.Spec.Inverse, edge.Spec.Table, edge.Spec.Columns...),
	)
	selector := e.selector.Clone().SetP(nil)
	selector.SetTotal(e.Total())
	if len(exprs) > 0 {
		HasNeighborsWith(selector, step, func(s *sql.Selector) {
			for i := range exprs {
				p, err := execExpr(edge.To, s, exprs[i])
				expect(err == nil, "edge evaluation failed for %s->%s: %s", e.context.Type, name, err)
				s.Where(p)
			}
		})
	} else {
		HasNeighbors(selector, step)
	}
	return selector.P()
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
		panic(execError{fmt.Sprintf("expect "+msg, args...)})
	}
}

type execError struct {
	msg string
}

func (p execError) Error() string { return fmt.Sprintf("sqlgraph: %s", p.msg) }

func catch(err *error) {
	if e := recover(); e != nil {
		xerr, ok := e.(execError)
		if !ok {
			panic(e)
		}
		*err = xerr
	}
}
