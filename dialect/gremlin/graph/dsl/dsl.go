// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package dsl provide an API for writing gremlin dsl queries almost as-is
// in Go without using strings in the code.
//
// Note that, the API is not type-safe and assume the provided query and
// its arguments are valid.
package dsl

import (
	"fmt"
	"strings"
	"time"
)

// Node represents a DSL step in the traversal.
type Node interface {
	// Code returns the code representation of the element and its bindings (if any).
	Code() (string, []interface{})
}

type (
	// Token holds a simple token, like assignment.
	Token string
	// List represents a list of elements.
	List struct {
		Elements []interface{}
	}
	// Func represents a function call.
	Func struct {
		Name string
		Args []interface{}
	}
	// Block represents a block/group of nodes.
	Block struct {
		Nodes []interface{}
	}
	// Var represents a variable assignment and usage.
	Var struct {
		Name string
		Elem interface{}
	}
)

// Code stringified the token.
func (t Token) Code() (string, []interface{}) { return string(t), nil }

// Code returns the code representation of a list.
func (l List) Code() (string, []interface{}) {
	c, args := codeList(", ", l.Elements...)
	return fmt.Sprintf("[%s]", c), args
}

// Code returns the code representation of a function call.
func (f Func) Code() (string, []interface{}) {
	c, args := codeList(", ", f.Args...)
	return fmt.Sprintf("%s(%s)", f.Name, c), args
}

// Code returns the code representation of group/block of nodes.
func (b Block) Code() (string, []interface{}) {
	return codeList("; ", b.Nodes...)
}

// Code returns the code representation of variable declaration or its identifier.
func (v Var) Code() (string, []interface{}) {
	c, args := code(v.Elem)
	if v.Name == "" {
		return c, args
	}
	return fmt.Sprintf("%s = %s", v.Name, c), args
}

// predefined nodes.
var (
	G   = Token("g")
	Dot = Token(".")
)

// NewFunc returns a new function node.
func NewFunc(name string, args ...interface{}) *Func {
	return &Func{Name: name, Args: args}
}

// NewList returns a new list node.
func NewList(args ...interface{}) *List {
	return &List{Elements: args}
}

// Querier is the interface that wraps the Query method.
type Querier interface {
	// Query returns the query-string (similar to the Gremlin byte-code) and its bindings.
	Query() (string, Bindings)
}

// Bindings are used to associate a variable with a value.
type Bindings map[string]interface{}

// Add adds new value to the bindings map, formats it if needed, and returns its generated name.
func (b Bindings) Add(v interface{}) string {
	k := fmt.Sprintf("$%x", len(b))
	switch v := v.(type) {
	case time.Time:
		b[k] = v.UnixNano()
	default:
		b[k] = v
	}
	return k
}

// Cardinality of vertex properties.
type Cardinality string

// Cardinality options.
const (
	Set    Cardinality = "set"
	Single Cardinality = "single"
)

// Code implements the Node interface.
func (c Cardinality) Code() (string, []interface{}) { return string(c), nil }

// Order of vertex properties.
type Order string

// Order options.
const (
	Incr    Order = "incr"
	Decr    Order = "decr"
	Shuffle Order = "shuffle"
)

// Code implements the Node interface.
func (o Order) Code() (string, []interface{}) { return string(o), nil }

// Column references a particular type of column in a complex data structure such as a Map, a Map.Entry, or a Path.
type Column string

// Column options.
const (
	Keys   Column = "keys"
	Values Column = "values"
)

// Code implements the Node interface.
func (o Column) Code() (string, []interface{}) { return string(o), nil }

// Scope used for steps that have a variable scope which alter the manner in which the step will behave in relation to how the traverses are processed.
type Scope string

// Scope options.
const (
	Local  Scope = "local"
	Global Scope = "global"
)

// Code implements the Node interface.
func (s Scope) Code() (string, []interface{}) { return string(s), nil }

func codeList(sep string, vs ...interface{}) (string, []interface{}) {
	var (
		br   strings.Builder
		args []interface{}
	)
	for i, node := range vs {
		if i > 0 {
			br.WriteString(sep)
		}
		c, nargs := code(node)
		br.WriteString(c)
		args = append(args, nargs...)
	}
	return br.String(), args
}

func code(v interface{}) (string, []interface{}) {
	switch n := v.(type) {
	case Node:
		return n.Code()
	case *Traversal:
		var (
			b    strings.Builder
			args []interface{}
		)
		for i := range n.nodes {
			code, nargs := n.nodes[i].Code()
			b.WriteString(code)
			args = append(args, nargs...)
		}
		return b.String(), args
	default:
		return "%s", []interface{}{v}
	}
}

func sface(args []string) (v []interface{}) {
	for _, s := range args {
		v = append(v, s)
	}
	return
}
