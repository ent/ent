// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqljson

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// HasKey return a predicate for checking that a JSON key
// exists and not NULL.
//
//	sqljson.HasKey("column", sql.DotPath("a.b[2].c"))
//
func HasKey(column string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		WritePath(b, column, opts...)
		b.WriteOp(sql.OpNotNull)
	})
}

// ValueEQ return a predicate for checking that a JSON value
// (returned by the path) is equal to the given argument.
//
//	sqljson.ValueEQ("a", 1, sqljson.Path("b"))
//
func ValueEQ(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		WritePath(b, column, opts...)
		b.WriteOp(sql.OpEQ).Arg(arg)
	})
}

// ValueNEQ return a predicate for checking that a JSON value
// (returned by the path) is not equal to the given argument.
//
//	sqljson.ValueNEQ("a", 1, sqljson.Path("b"))
//
func ValueNEQ(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		WritePath(b, column, opts...)
		b.WriteOp(sql.OpNEQ).Arg(arg)
	})
}

// ValueGT return a predicate for checking that a JSON value
// (returned by the path) is greater than the given argument.
//
//	sqljson.ValueGT("a", 1, sqljson.Path("b"))
//
func ValueGT(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		WritePath(b, column, opts...)
		b.WriteOp(sql.OpGT).Arg(arg)
	})
}

// ValueGTE return a predicate for checking that a JSON value
// (returned by the path) is greater than the given argument.
//
//	sqljson.ValueGTE("a", 1, sqljson.Path("b"))
//
func ValueGTE(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		WritePath(b, column, opts...)
		b.WriteOp(sql.OpGTE).Arg(arg)
	})
}

// WritePath writes the JSON path from the given options to the SQL builder.
//
//	sqljson.WritePath(b, Path("a", "b", "[1]", "c"), Cast("int"))
//
func WritePath(b *sql.Builder, column string, opts ...Option) {
	path := &PathOptions{Ident: column}
	for i := range opts {
		opts[i](path)
	}
	path.WriteTo(b)
}

// Option allows for calling database JSON paths with functional options.
type Option func(*PathOptions)

// Path sets the path to the JSON value of a column.
//
//	WritePath(b, "column", Path("a", "b", "[1]", "c"))
//
func Path(path ...string) Option {
	return func(p *PathOptions) {
		p.Path = path
	}
}

// DotPath is similar to Path, but accepts string with dot format.
//
//	WritePath(b, "column", DotPath("a.b.c"))
//	WritePath(b, "column", DotPath("a.b[2].c"))
//
// Note that DotPath is ignored if the input is invalid.
func DotPath(dotpath string) Option {
	path, _ := ParsePath(dotpath)
	return func(p *PathOptions) {
		p.Path = path
	}
}

// Unquote indicates that the result value should be unquoted.
//
//	WritePath(b, "column", Path("a", "b", "[1]", "c"), Unquote(true))
//
func Unquote(unquote bool) Option {
	return func(p *PathOptions) {
		p.Unquote = unquote
	}
}

// Cast indicates that the result value should be casted to the given type.
//
//	WritePath(b, "column", Path("a", "b", "[1]", "c"), Cast("int"))
//
func Cast(typ string) Option {
	return func(p *PathOptions) {
		p.Cast = typ
	}
}

// PathOptions holds the options for accessing a JSON value from an identifier.
type PathOptions struct {
	Ident   string
	Path    []string
	Cast    string
	Unquote bool
}

// WriteTo writes the JSON path to the sql.Builder.
func (p *PathOptions) WriteTo(b *sql.Builder) {
	switch {
	case len(p.Path) == 0:
		b.Ident(p.Ident)
	case b.Dialect() == dialect.Postgres:
		if p.Cast != "" {
			b.WriteString("CAST(")
			defer b.WriteString(" AS " + p.Cast + ")")
		}
		b.Ident(p.Ident)
		for i, s := range p.Path {
			b.WriteString("->")
			if p.Unquote && i == len(p.Path)-1 {
				b.WriteString(">")
			}
			if idx, ok := isJSONIdx(s); ok {
				b.WriteString(idx)
			} else {
				b.WriteString("'" + s + "'")
			}
		}
	default:
		if p.Unquote && b.Dialect() == dialect.MySQL {
			b.WriteString("JSON_UNQUOTE(")
			defer b.WriteByte(')')
		}
		b.WriteString("JSON_EXTRACT(")
		b.Ident(p.Ident).Comma()
		b.WriteString(`"$`)
		for _, p := range p.Path {
			if _, ok := isJSONIdx(p); ok {
				b.WriteString(p)
			} else {
				b.WriteString("." + p)
			}
		}
		b.WriteString(`")`)
	}
}

// ParsePath parses the "dotpath" for the DotPath option.
//
//	"a.b"		=> ["a", "b"]
//	"a[1][2]"	=> ["a", "[1]", "[2]"]
//	"a.\"b.c\"	=> ["a", "\"b.c\""]
//
func ParsePath(dotpath string) ([]string, error) {
	var (
		i, p int
		path []string
	)
	for i < len(dotpath) {
		switch r := dotpath[i]; {
		case r == '"':
			if i == len(dotpath)-1 {
				return nil, fmt.Errorf("unexpected quote")
			}
			idx := strings.IndexRune(dotpath[i+1:], '"')
			if idx == -1 || idx == 0 {
				return nil, fmt.Errorf("unbalanced quote")
			}
			i += idx + 2
		case r == '[':
			if p != i {
				path = append(path, dotpath[p:i])
			}
			p = i
			if i == len(dotpath)-1 {
				return nil, fmt.Errorf("unexpected bracket")
			}
			idx := strings.IndexRune(dotpath[i:], ']')
			if idx == -1 || idx == 1 {
				return nil, fmt.Errorf("unbalanced bracket")
			}
			if !isNumber(dotpath[i+1 : i+idx]) {
				return nil, fmt.Errorf("invalid index %q", dotpath[i:i+idx+1])
			}
			i += idx + 1
		case r == '.' || r == ']':
			if p != i {
				path = append(path, dotpath[p:i])
			}
			i++
			p = i
		default:
			i++
		}
	}
	if p != i {
		path = append(path, dotpath[p:i])
	}
	return path, nil
}

// isJSONIdx reports whether the string represents a JSON index.
func isJSONIdx(s string) (string, bool) {
	if len(s) > 2 && s[0] == '[' && s[len(s)-1] == ']' && isNumber(s[1:len(s)-1]) {
		return s[1 : len(s)-1], true
	}
	return "", false
}

// isNumber reports whether the string is a number (category N).
func isNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}
