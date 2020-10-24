// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqljson

import (
	"encoding/json"
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
		ValuePath(b, column, opts...)
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
		opts, arg = normalizePG(b, arg, opts)
		ValuePath(b, column, opts...)
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
		opts, arg = normalizePG(b, arg, opts)
		ValuePath(b, column, opts...)
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
		opts, arg = normalizePG(b, arg, opts)
		ValuePath(b, column, opts...)
		b.WriteOp(sql.OpGT).Arg(arg)
	})
}

// ValueGTE return a predicate for checking that a JSON value
// (returned by the path) is greater than or equal to the given
// argument.
//
//	sqljson.ValueGTE("a", 1, sqljson.Path("b"))
//
func ValueGTE(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts, arg = normalizePG(b, arg, opts)
		ValuePath(b, column, opts...)
		b.WriteOp(sql.OpGTE).Arg(arg)
	})
}

// ValueLT return a predicate for checking that a JSON value
// (returned by the path) is less than the given argument.
//
//	sqljson.ValueLT("a", 1, sqljson.Path("b"))
//
func ValueLT(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts, arg = normalizePG(b, arg, opts)
		ValuePath(b, column, opts...)
		b.WriteOp(sql.OpLT).Arg(arg)
	})
}

// ValueLTE return a predicate for checking that a JSON value
// (returned by the path) is less than or equal to the given
// argument.
//
//	sqljson.ValueLTE("a", 1, sqljson.Path("b"))
//
func ValueLTE(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts, arg = normalizePG(b, arg, opts)
		ValuePath(b, column, opts...)
		b.WriteOp(sql.OpLTE).Arg(arg)
	})
}

// ValueContains return a predicate for checking that a JSON
// value (returned by the path) contains the given argument.
//
//	sqljson.ValueContains("a", 1, sqljson.Path("b"))
//
func ValueContains(column string, arg interface{}, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		path := &PathOptions{Ident: column}
		for i := range opts {
			opts[i](path)
		}
		switch b.Dialect() {
		case dialect.MySQL:
			b.WriteString("JSON_CONTAINS").Nested(func(b *sql.Builder) {
				b.Ident(column).Comma()
				b.Arg(marshal(arg)).Comma()
				path.mysqlPath(b)
			})
			b.WriteOp(sql.OpEQ).Arg(1)
		case dialect.SQLite:
			b.WriteString("EXISTS").Nested(func(b *sql.Builder) {
				b.WriteString("SELECT * FROM JSON_EACH").Nested(func(b *sql.Builder) {
					b.Ident(column).Comma()
					path.mysqlPath(b)
				})
				b.WriteString(" WHERE ").Ident("value").WriteOp(sql.OpEQ).Arg(arg)
			})
		case dialect.Postgres:
			opts, arg = normalizePG(b, arg, opts)
			path.Cast = "jsonb"
			path.value(b)
			b.WriteString(" @> ").Arg(marshal(arg))
		}
	})
}

// LenEQ return a predicate for checking that an array length
// of a JSON (returned by the path) is equal to the given argument.
//
//	sqljson.LenEQ("a", 1, sqljson.Path("b"))
//
func LenEQ(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		LenPath(b, column, opts...)
		b.WriteOp(sql.OpEQ).Arg(size)
	})
}

// LenNEQ return a predicate for checking that an array length
// of a JSON (returned by the path) is not equal to the given argument.
//
//	sqljson.LenEQ("a", 1, sqljson.Path("b"))
//
func LenNEQ(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		LenPath(b, column, opts...)
		b.WriteOp(sql.OpNEQ).Arg(size)
	})
}

// LenGT return a predicate for checking that an array length
// of a JSON (returned by the path) is greater than the given
// argument.
//
//	sqljson.LenGT("a", 1, sqljson.Path("b"))
//
func LenGT(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		LenPath(b, column, opts...)
		b.WriteOp(sql.OpGT).Arg(size)
	})
}

// LenGTE return a predicate for checking that an array length
// of a JSON (returned by the path) is greater than or equal to
// the given argument.
//
//	sqljson.LenGTE("a", 1, sqljson.Path("b"))
//
func LenGTE(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		LenPath(b, column, opts...)
		b.WriteOp(sql.OpGTE).Arg(size)
	})
}

// LenLT return a predicate for checking that an array length
// of a JSON (returned by the path) is less than the given
// argument.
//
//	sqljson.LenLT("a", 1, sqljson.Path("b"))
//
func LenLT(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		LenPath(b, column, opts...)
		b.WriteOp(sql.OpLT).Arg(size)
	})
}

// LenLTE return a predicate for checking that an array length
// of a JSON (returned by the path) is less than or equal to
// the given argument.
//
//	sqljson.LenLTE("a", 1, sqljson.Path("b"))
//
func LenLTE(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		LenPath(b, column, opts...)
		b.WriteOp(sql.OpLTE).Arg(size)
	})
}

// ValuePath writes to the given SQL builder the JSON path for
// getting the value of a given JSON path.
//
//	sqljson.ValuePath(b, Path("a", "b", "[1]", "c"), Cast("int"))
//
func ValuePath(b *sql.Builder, column string, opts ...Option) {
	path := &PathOptions{Ident: column}
	for i := range opts {
		opts[i](path)
	}
	path.value(b)
}

// LenPath writes to the given SQL builder the JSON path for
// getting the length of a given JSON path.
//
//	sqljson.LenPath(b, Path("a", "b", "[1]", "c"))
//
func LenPath(b *sql.Builder, column string, opts ...Option) {
	path := &PathOptions{Ident: column}
	for i := range opts {
		opts[i](path)
	}
	path.length(b)
}

// Option allows for calling database JSON paths with functional options.
type Option func(*PathOptions)

// Path sets the path to the JSON value of a column.
//
//	ValuePath(b, "column", Path("a", "b", "[1]", "c"))
//
func Path(path ...string) Option {
	return func(p *PathOptions) {
		p.Path = path
	}
}

// DotPath is similar to Path, but accepts string with dot format.
//
//	ValuePath(b, "column", DotPath("a.b.c"))
//	ValuePath(b, "column", DotPath("a.b[2].c"))
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
//	ValuePath(b, "column", Path("a", "b", "[1]", "c"), Unquote(true))
//
func Unquote(unquote bool) Option {
	return func(p *PathOptions) {
		p.Unquote = unquote
	}
}

// Cast indicates that the result value should be casted to the given type.
//
//	ValuePath(b, "column", Path("a", "b", "[1]", "c"), Cast("int"))
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

// value writes the path for getting the JSON value.
func (p *PathOptions) value(b *sql.Builder) {
	switch {
	case len(p.Path) == 0:
		b.Ident(p.Ident)
	case b.Dialect() == dialect.Postgres:
		if p.Cast != "" {
			b.WriteByte('(')
			defer b.WriteString(")::" + p.Cast)
		}
		p.pgPath(b)
	default:
		if p.Unquote && b.Dialect() == dialect.MySQL {
			b.WriteString("JSON_UNQUOTE(")
			defer b.WriteByte(')')
		}
		p.mysqlFunc("JSON_EXTRACT", b)
	}
}

// value writes the path for getting the length of a JSON value.
func (p *PathOptions) length(b *sql.Builder) {
	switch {
	case b.Dialect() == dialect.Postgres:
		b.WriteString("JSONB_ARRAY_LENGTH(")
		p.pgPath(b)
		b.WriteByte(')')
	case b.Dialect() == dialect.MySQL:
		p.mysqlFunc("JSON_LENGTH", b)
	default:
		p.mysqlFunc("JSON_ARRAY_LENGTH", b)
	}
}

// mysqlFunc writes the JSON path in MySQL format for the
// the given function. `JSON_EXTRACT("a", '$.b.c')`.
func (p *PathOptions) mysqlFunc(fn string, b *sql.Builder) {
	b.WriteString(fn).WriteByte('(')
	b.Ident(p.Ident).Comma()
	p.mysqlPath(b)
	b.WriteByte(')')
}

// mysqlPath writes the JSON path in MySQL (or SQLite) format.
func (p *PathOptions) mysqlPath(b *sql.Builder) {
	b.WriteString(`"$`)
	for _, p := range p.Path {
		if _, ok := isJSONIdx(p); ok {
			b.WriteString(p)
		} else {
			b.WriteString("." + p)
		}
	}
	b.WriteByte('"')
}

// pgPath writes the JSON path in Postgres format `"a"->'b'->>'c'`.
func (p *PathOptions) pgPath(b *sql.Builder) {
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

// normalizePG adds cast option to the JSON path is the argument type is
// not string, in order to avoid "missing type casts" error in Postgres.
func normalizePG(b *sql.Builder, arg interface{}, opts []Option) ([]Option, interface{}) {
	if b.Dialect() != dialect.Postgres {
		return opts, arg
	}
	base := []Option{Unquote(true)}
	switch arg.(type) {
	case string:
	case bool:
		base = append(base, Cast("bool"))
	case float32, float64:
		base = append(base, Cast("float"))
	case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64:
		base = append(base, Cast("int"))
	default: // convert unknown types to text.
		arg = marshal(arg)
	}
	return append(base, opts...), arg
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

// marshal stringifies the given argument to a valid JSON document.
func marshal(arg interface{}) interface{} {
	if buf, err := json.Marshal(arg); err == nil {
		arg = string(buf)
	}
	return arg
}
