package entql

import (
	"encoding/json"
	"fmt"
	"strings"
)

// An Op represents a predicate operator.
type Op int

const (
	OpAnd          Op = iota // logical and.
	OpOr                     // logical or.
	OpNot                    // logical negation.
	OpEQ                     // =
	OpNEQ                    // <>
	OpGT                     // >
	OpGTE                    // >=
	OpLT                     // <
	OpLTE                    // <=
	OpIn                     // IN
	OpNotIn                  // NOT IN
	OpEqualFold              // equals case-insensitive
	OpContains               // containing
	OpContainsFold           // containing case-insensitive
	OpHasPrefix              // startingWith
	OpHasSuffix              // endingWith
	OpHasEdge                // HasEdge
)

var ops = [...]string{
	OpAnd:          "&&",
	OpOr:           "||",
	OpNot:          "!",
	OpEQ:           "==",
	OpNEQ:          "!=",
	OpGT:           ">",
	OpGTE:          ">=",
	OpLT:           "<",
	OpLTE:          "<=",
	OpIn:           "in",
	OpNotIn:        "not in",
	OpEqualFold:    "equal_fold",
	OpContains:     "contains",
	OpContainsFold: "contains_fold",
	OpHasPrefix:    "has_prefix",
	OpHasSuffix:    "has_suffix",
	OpHasEdge:      "has_edge",
}

// String returns the text representation of an operator.
func (o Op) String() string {
	return ops[o]
}

type (
	// Expr represents an entql expression. All expressions implement the Expr interface.
	Expr interface {
		expr()
		fmt.Stringer
	}

	// P represents an expression that returns a boolean value depending on its variables.
	P interface {
		Expr
		Or(P) P
		And(P) P
		Negate() P
	}
)

type (
	// A UnaryExpr represents a unary expression.
	UnaryExpr struct {
		Op Op
		X  Expr
	}

	// A BinaryExpr represents a binary expression.
	BinaryExpr struct {
		Op   Op
		X, Y Expr
	}

	// A NaryExpr represents a n-ary expression.
	NaryExpr struct {
		Op Op
		Xs []Expr
	}

	// A CallExpr represents an expression followed by its arguments.
	CallExpr struct {
		Op   Op
		Args []Expr
	}

	// A Field represents a node field.
	Field struct {
		Name string
	}

	// An Edge represents an edge in the graph.
	Edge struct {
		Name string
	}

	// A Value represents an arbitrary value.
	Value struct {
		V interface{}
	}
)

// Not returns a predicate that represents the logical negation of the given predicate.
func Not(x P) P {
	return &UnaryExpr{
		Op: OpNot,
		X:  x,
	}
}

// And returns a composed predicate that represents the logical AND predicate.
func And(x, y P, z ...P) P {
	if len(z) == 0 {
		return &BinaryExpr{
			Op: OpAnd,
			X:  x,
			Y:  y,
		}
	}
	return &NaryExpr{
		Op: OpAnd,
		Xs: append([]Expr{x, y}, p2expr(z)...),
	}
}

// Or returns a composed predicate that represents the logical OR predicate.
func Or(x, y P, z ...P) P {
	if len(z) == 0 {
		return &BinaryExpr{
			Op: OpOr,
			X:  x,
			Y:  y,
		}
	}
	return &NaryExpr{
		Op: OpOr,
		Xs: append([]Expr{x, y}, p2expr(z)...),
	}
}

// F returns a field expression for the given name.
func F(name string) *Field {
	return &Field{Name: name}
}

// EQ returns a predicate to check if the expressions are equal.
func EQ(x, y Expr) P {
	return &BinaryExpr{
		Op: OpEQ,
		X:  x,
		Y:  y,
	}
}

// FieldEQ returns a predicate to check if a field is equivalent to a given value.
func FieldEQ(name string, v interface{}) P {
	return &BinaryExpr{
		Op: OpEQ,
		X:  &Field{Name: name},
		Y:  &Value{V: v},
	}
}

// FieldNEQ returns a predicate to check if a field is not equivalent to a given value.
func FieldNEQ(name string, v interface{}) P {
	return &BinaryExpr{
		Op: OpNEQ,
		X:  &Field{Name: name},
		Y:  &Value{V: v},
	}
}

// FieldGT returns a predicate to check if a field is > than the given value.
func FieldGT(name string, v interface{}) P {
	return &BinaryExpr{
		Op: OpGT,
		X:  &Field{Name: name},
		Y:  &Value{V: v},
	}
}

// FieldGTE returns a predicate to check if a field is >= than the given value.
func FieldGTE(name string, v interface{}) P {
	return &BinaryExpr{
		Op: OpGTE,
		X:  &Field{Name: name},
		Y:  &Value{V: v},
	}
}

// FieldLT returns a predicate to check if a field is < than the given value.
func FieldLT(name string, v interface{}) P {
	return &BinaryExpr{
		Op: OpLT,
		X:  &Field{Name: name},
		Y:  &Value{V: v},
	}
}

// FieldLTE returns a predicate to check if a field is <= >than the given value.
func FieldLTE(name string, v interface{}) P {
	return &BinaryExpr{
		Op: OpLTE,
		X:  &Field{Name: name},
		Y:  &Value{V: v},
	}
}

// FieldContains returns a predicate to check if the field value contains a substr.
func FieldContains(name, substr string) P {
	return &CallExpr{
		Op:   OpContains,
		Args: []Expr{&Field{Name: name}, &Value{V: substr}},
	}
}

// FieldContainsFold returns a predicate to check if the field value contains a substr under case-folding.
func FieldContainsFold(name, substr string) P {
	return &CallExpr{
		Op:   OpContainsFold,
		Args: []Expr{&Field{Name: name}, &Value{V: substr}},
	}
}

// FieldEqualFold returns a predicate to check if the field is equal to the given string under case-folding.
func FieldEqualFold(name, v string) P {
	return &CallExpr{
		Op:   OpEqualFold,
		Args: []Expr{&Field{Name: name}, &Value{V: v}},
	}
}

// FieldHasPrefix returns a predicate to check if the field starts with the given prefix.
func FieldHasPrefix(name, prefix string) P {
	return &CallExpr{
		Op:   OpHasPrefix,
		Args: []Expr{&Field{Name: name}, &Value{V: prefix}},
	}
}

// FieldHasSuffix returns a predicate to check if the field ends with the given suffix.
func FieldHasSuffix(name, suffix string) P {
	return &CallExpr{
		Op:   OpHasSuffix,
		Args: []Expr{&Field{Name: name}, &Value{V: suffix}},
	}
}

// FieldIn returns a predicate to check if the field value matches any value in the given list.
func FieldIn(name string, vs ...interface{}) P {
	return &BinaryExpr{
		Op: OpIn,
		X:  &Field{Name: name},
		Y:  &Value{V: vs},
	}
}

// FieldNotIn returns a predicate to check if the field value doesn't match any value in the given list.
func FieldNotIn(name string, vs ...interface{}) P {
	return &BinaryExpr{
		Op: OpNotIn,
		X:  &Field{Name: name},
		Y:  &Value{V: vs},
	}
}

// FieldNil returns a predicate to check if a field is nil (null in databases).
func FieldNil(name string) P {
	return &BinaryExpr{
		Op: OpEQ,
		X:  &Field{Name: name},
		Y:  (*Value)(nil),
	}
}

// FieldNotNil returns a predicate to check if a field is not nil (not null in databases).
func FieldNotNil(name string) P {
	return &BinaryExpr{
		Op: OpNEQ,
		X:  &Field{Name: name},
		Y:  (*Value)(nil),
	}
}

// HasEdge returns a predicate to check if an edge exists (not null in databases).
func HasEdge(name string) P {
	return &CallExpr{
		Op:   OpHasEdge,
		Args: []Expr{&Edge{Name: name}},
	}
}

// HasEdgeWith returns a predicate to check if the "other nodes" that are connected to the
// edge returns true on the provided predicate.
func HasEdgeWith(name string, p ...P) P {
	return &CallExpr{
		Op:   OpHasEdge,
		Args: append([]Expr{&Edge{Name: name}}, p2expr(p)...),
	}
}

// And returns a composed predicate that represents the logical AND predicate.
func (e *BinaryExpr) And(p P) P {
	return And(e, p)
}

// Or returns a composed predicate that represents the logical OR predicate.
func (e *BinaryExpr) Or(e1 P) P {
	return Or(e, e1)
}

// Negate negates the predicate.
func (e *BinaryExpr) Negate() P {
	return Not(e)
}

// And returns a composed predicate that represents the logical AND predicate.
func (e *NaryExpr) And(e1 P) P {
	return And(e, e1)
}

// Or returns a composed predicate that represents the logical OR predicate.
func (e *NaryExpr) Or(e1 P) P {
	return Or(e, e1)
}

// Negate negates the predicate.
func (e *NaryExpr) Negate() P {
	return Not(e)
}

// And returns a composed predicate that represents the logical AND predicate.
func (e *UnaryExpr) And(e1 P) P {
	return And(e, e1)
}

// Or returns a composed predicate that represents the logical OR predicate.
func (e *UnaryExpr) Or(e1 P) P {
	return Or(e, e1)
}

// Negate negates the predicate.
func (e *UnaryExpr) Negate() P {
	return Not(e)
}

// And returns a composed predicate that represents the logical AND predicate.
func (e *CallExpr) And(e1 P) P {
	return And(e, e1)
}

// Or returns a composed predicate that represents the logical OR predicate.
func (e *CallExpr) Or(e1 P) P {
	return Or(e, e1)
}

// Negate negates the predicate.
func (e *CallExpr) Negate() P {
	return Not(e)
}

// String returns the text representation of a binary expression.
func (e *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.X, e.Op, e.Y)
}

// String returns the text representation of a unary expression.
func (e *UnaryExpr) String() string {
	return fmt.Sprintf("%s(%s)", e.Op, e.X)
}

// String returns the text representation of an n-ary expression.
func (e *NaryExpr) String() string {
	var s strings.Builder
	s.WriteByte('(')
	for i, x := range e.Xs {
		if i > 0 {
			s.WriteByte(' ')
			s.WriteString(e.Op.String())
			s.WriteByte(' ')
		}
		s.WriteString(x.String())
	}
	s.WriteByte(')')
	return s.String()
}

// String returns the text representation of a call expression.
func (e *CallExpr) String() string {
	var s strings.Builder
	s.WriteString(e.Op.String())
	s.WriteByte('(')
	for i, x := range e.Args {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(x.String())
	}
	s.WriteByte(')')
	return s.String()
}

// String returns the text representation of a field.
func (f *Field) String() string {
	return f.Name
}

// String returns the text representation of an edge.
func (e *Edge) String() string {
	return e.Name
}

// String returns the text representation of a value.
func (v *Value) String() string {
	if v == nil {
		return "nil"
	}
	buf, err := json.Marshal(v.V)
	if err != nil {
		return fmt.Sprint(v.V)
	}
	return string(buf)
}

func p2expr(ps []P) []Expr {
	expr := make([]Expr, len(ps))
	for i := range ps {
		expr[i] = ps[i]
	}
	return expr
}

func (*Edge) expr()       {}
func (*Field) expr()      {}
func (*Value) expr()      {}
func (*CallExpr) expr()   {}
func (*NaryExpr) expr()   {}
func (*UnaryExpr) expr()  {}
func (*BinaryExpr) expr() {}
