// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package entql provides an experimental API for interacting dynamically
// with ent queries. For more info, search for it in https://entgo.io.
package entql

import (
	"encoding/json"
	"fmt"
	"strings"
)

// An Op represents a predicate operator.
type Op int

// Builtin operators.
const (
	OpAnd   Op = iota // logical and.
	OpOr              // logical or.
	OpNot             // logical negation.
	OpEQ              // =
	OpNEQ             // <>
	OpGT              // >
	OpGTE             // >=
	OpLT              // <
	OpLTE             // <=
	OpIn              // IN
	OpNotIn           // NOT IN
)

var ops = [...]string{
	OpAnd:   "&&",
	OpOr:    "||",
	OpNot:   "!",
	OpEQ:    "==",
	OpNEQ:   "!=",
	OpGT:    ">",
	OpGTE:   ">=",
	OpLT:    "<",
	OpLTE:   "<=",
	OpIn:    "in",
	OpNotIn: "not in",
}

// String returns the text representation of an operator.
func (o Op) String() string {
	if o >= 0 && int(o) < len(ops) {
		return ops[o]
	}
	return "<invalid>"
}

// A Func represents a function expression.
type Func string

// Builtin functions.
const (
	FuncEqualFold    Func = "equal_fold"    // equals case-insensitive
	FuncContains     Func = "contains"      // containing
	FuncContainsFold Func = "contains_fold" // containing case-insensitive
	FuncHasPrefix    Func = "has_prefix"    // startingWith
	FuncHasSuffix    Func = "has_suffix"    // endingWith
	FuncHasEdge      Func = "has_edge"      // HasEdge
)

type (
	// Expr represents an entql expression. All expressions implement the Expr interface.
	Expr interface {
		expr()
		fmt.Stringer
	}

	// P represents an expression that returns a boolean value depending on its variables.
	P interface {
		Expr
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

	// A CallExpr represents a function call with its arguments.
	CallExpr struct {
		Func Func
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

// NEQ returns a predicate to check if the expressions are not equal.
func NEQ(x, y Expr) P {
	return &BinaryExpr{
		Op: OpNEQ,
		X:  x,
		Y:  y,
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

// GT returns a predicate to check if the expression x > than expression y.
func GT(x, y Expr) P {
	return &BinaryExpr{
		Op: OpGT,
		X:  x,
		Y:  y,
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

// GTE returns a predicate to check if the expression x >= than expression y.
func GTE(x, y Expr) P {
	return &BinaryExpr{
		Op: OpGTE,
		X:  x,
		Y:  y,
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

// LT returns a predicate to check if the expression x < than expression y.
func LT(x, y Expr) P {
	return &BinaryExpr{
		Op: OpLT,
		X:  x,
		Y:  y,
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

// LTE returns a predicate to check if the expression x <= than expression y.
func LTE(x, y Expr) P {
	return &BinaryExpr{
		Op: OpLTE,
		X:  x,
		Y:  y,
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
		Func: FuncContains,
		Args: []Expr{&Field{Name: name}, &Value{V: substr}},
	}
}

// FieldContainsFold returns a predicate to check if the field value contains a substr under case-folding.
func FieldContainsFold(name, substr string) P {
	return &CallExpr{
		Func: FuncContainsFold,
		Args: []Expr{&Field{Name: name}, &Value{V: substr}},
	}
}

// FieldEqualFold returns a predicate to check if the field is equal to the given string under case-folding.
func FieldEqualFold(name, v string) P {
	return &CallExpr{
		Func: FuncEqualFold,
		Args: []Expr{&Field{Name: name}, &Value{V: v}},
	}
}

// FieldHasPrefix returns a predicate to check if the field starts with the given prefix.
func FieldHasPrefix(name, prefix string) P {
	return &CallExpr{
		Func: FuncHasPrefix,
		Args: []Expr{&Field{Name: name}, &Value{V: prefix}},
	}
}

// FieldHasSuffix returns a predicate to check if the field ends with the given suffix.
func FieldHasSuffix(name, suffix string) P {
	return &CallExpr{
		Func: FuncHasSuffix,
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
		Func: FuncHasEdge,
		Args: []Expr{&Edge{Name: name}},
	}
}

// HasEdgeWith returns a predicate to check if the "other nodes" that are connected to the
// edge returns true on the provided predicate.
func HasEdgeWith(name string, p ...P) P {
	return &CallExpr{
		Func: FuncHasEdge,
		Args: append([]Expr{&Edge{Name: name}}, p2expr(p)...),
	}
}

// Negate negates the predicate.
func (e *BinaryExpr) Negate() P {
	return Not(e)
}

// Negate negates the predicate.
func (e *NaryExpr) Negate() P {
	return Not(e)
}

// Negate negates the predicate.
func (e *UnaryExpr) Negate() P {
	return Not(e)
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
	s.WriteString(string(e.Func))
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
