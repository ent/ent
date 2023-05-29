// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqljson

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// HasKey return a predicate for checking that a JSON key
// exists and not NULL.
//
//	sqljson.HasKey("column", sql.DotPath("a.b[2].c"))
func HasKey(column string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		switch b.Dialect() {
		case dialect.SQLite:
			// JSON_TYPE returns NULL in case the path selects an element
			// that does not exist. See: https://sqlite.org/json1.html#jtype.
			path := identPath(column, opts...)
			path.mysqlFunc("JSON_TYPE", b)
			b.WriteOp(sql.OpNotNull)
		default:
			valuePath(b, column, opts...)
			b.WriteOp(sql.OpNotNull)
		}
	})
}

// ValueIsNull return a predicate for checking that a JSON value
// (returned by the path) is a null literal (JSON "null").
//
// In order to check if the column is NULL (database NULL), or if
// the JSON key exists, use sql.IsNull or sqljson.HasKey.
//
//	sqljson.ValueIsNull("a", sqljson.Path("b"))
func ValueIsNull(column string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		switch b.Dialect() {
		case dialect.MySQL:
			path := identPath(column, opts...)
			b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
				b.Ident(column).Comma()
				b.WriteString("'null'").Comma()
				path.mysqlPath(b)
			})
		case dialect.Postgres:
			valuePath(b, column, append(opts, Cast("jsonb"))...)
			b.WriteOp(sql.OpEQ).WriteString("'null'::jsonb")
		case dialect.SQLite:
			path := identPath(column, opts...)
			path.mysqlFunc("JSON_TYPE", b)
			b.WriteOp(sql.OpEQ).WriteString("'null'")
		}
	})
}

// ValueIsNotNull return a predicate for checking that a JSON value
// (returned by the path) is not null literal (JSON "null").
//
//	sqljson.ValueIsNotNull("a", sqljson.Path("b"))
func ValueIsNotNull(column string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		switch b.Dialect() {
		case dialect.Postgres:
			valuePath(b, column, append(opts, Cast("jsonb"))...)
			b.WriteOp(sql.OpNEQ).WriteString("'null'::jsonb")
		case dialect.SQLite:
			path := identPath(column, opts...)
			path.mysqlFunc("JSON_TYPE", b)
			b.WriteOp(sql.OpNEQ).WriteString("'null'")
		case dialect.MySQL:
			path := identPath(column, opts...)
			b.WriteString("NOT(JSON_CONTAINS").Wrap(func(b *sql.Builder) {
				b.Ident(column).Comma()
				b.WriteString("'null'").Comma()
				path.mysqlPath(b)
			}).WriteString(")")
		}
	})
}

// ValueEQ return a predicate for checking that a JSON value
// (returned by the path) is equal to the given argument.
//
//	sqljson.ValueEQ("a", 1, sqljson.Path("b"))
func ValueEQ(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = normalizePG(b, arg, opts)
		valuePath(b, column, opts...)
		b.WriteOp(sql.OpEQ)
		// Inline boolean values, as some drivers (e.g., MySQL) encode them as 0/1.
		if v, ok := arg.(bool); ok {
			b.WriteString(strconv.FormatBool(v))
		} else {
			b.Arg(arg)
		}
	})
}

// ValueNEQ return a predicate for checking that a JSON value
// (returned by the path) is not equal to the given argument.
//
//	sqljson.ValueNEQ("a", 1, sqljson.Path("b"))
func ValueNEQ(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = normalizePG(b, arg, opts)
		valuePath(b, column, opts...)
		b.WriteOp(sql.OpNEQ).Arg(arg)
	})
}

// ValueGT return a predicate for checking that a JSON value
// (returned by the path) is greater than the given argument.
//
//	sqljson.ValueGT("a", 1, sqljson.Path("b"))
func ValueGT(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = normalizePG(b, arg, opts)
		valuePath(b, column, opts...)
		b.WriteOp(sql.OpGT).Arg(arg)
	})
}

// ValueGTE return a predicate for checking that a JSON value
// (returned by the path) is greater than or equal to the given
// argument.
//
//	sqljson.ValueGTE("a", 1, sqljson.Path("b"))
func ValueGTE(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = normalizePG(b, arg, opts)
		valuePath(b, column, opts...)
		b.WriteOp(sql.OpGTE).Arg(arg)
	})
}

// ValueLT return a predicate for checking that a JSON value
// (returned by the path) is less than the given argument.
//
//	sqljson.ValueLT("a", 1, sqljson.Path("b"))
func ValueLT(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = normalizePG(b, arg, opts)
		valuePath(b, column, opts...)
		b.WriteOp(sql.OpLT).Arg(arg)
	})
}

// ValueLTE return a predicate for checking that a JSON value
// (returned by the path) is less than or equal to the given
// argument.
//
//	sqljson.ValueLTE("a", 1, sqljson.Path("b"))
func ValueLTE(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = normalizePG(b, arg, opts)
		valuePath(b, column, opts...)
		b.WriteOp(sql.OpLTE).Arg(arg)
	})
}

// ValueContains return a predicate for checking that a JSON
// value (returned by the path) contains the given argument.
//
//	sqljson.ValueContains("a", 1, sqljson.Path("b"))
func ValueContains(column string, arg any, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		path := identPath(column, opts...)
		switch b.Dialect() {
		case dialect.MySQL:
			b.WriteString("JSON_CONTAINS").Wrap(func(b *sql.Builder) {
				b.Ident(column).Comma()
				b.Arg(marshalArg(arg)).Comma()
				path.mysqlPath(b)
			})
			b.WriteOp(sql.OpEQ).Arg(1)
		case dialect.SQLite:
			b.WriteString("EXISTS").Wrap(func(b *sql.Builder) {
				b.WriteString("SELECT * FROM JSON_EACH").Wrap(func(b *sql.Builder) {
					b.Ident(column).Comma()
					path.mysqlPath(b)
				})
				b.WriteString(" WHERE ").Ident("value").WriteOp(sql.OpEQ).Arg(arg)
			})
		case dialect.Postgres:
			opts = normalizePG(b, arg, opts)
			path.Cast = "jsonb"
			path.value(b)
			b.WriteString(" @> ").Arg(marshalArg(arg))
		}
	})
}

// StringHasPrefix return a predicate for checking that a JSON string value
// (returned by the path) has the given substring as prefix
func StringHasPrefix(column string, prefix string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = append([]Option{Unquote(true)}, opts...)
		valuePath(b, column, opts...)
		b.Join(sql.HasPrefix("", prefix))
	})
}

// StringHasSuffix return a predicate for checking that a JSON string value
// (returned by the path) has the given substring as suffix
func StringHasSuffix(column string, suffix string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = append([]Option{Unquote(true)}, opts...)
		valuePath(b, column, opts...)
		b.Join(sql.HasSuffix("", suffix))
	})
}

// StringContains return a predicate for checking that a JSON string value
// (returned by the path) contains the given substring
func StringContains(column string, sub string, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		opts = append([]Option{Unquote(true)}, opts...)
		valuePath(b, column, opts...)
		b.Join(sql.Contains("", sub))
	})
}

// ValueIn return a predicate for checking that a JSON value
// (returned by the path) is IN the given arguments.
//
//	sqljson.ValueIn("a", []any{1, 2, 3}, sqljson.Path("b"))
func ValueIn(column string, args []any, opts ...Option) *sql.Predicate {
	return valueInOp(column, args, opts, sql.OpIn)
}

// ValueNotIn return a predicate for checking that a JSON value
// (returned by the path) is NOT IN the given arguments.
//
//	sqljson.ValueNotIn("a", []any{1, 2, 3}, sqljson.Path("b"))
func ValueNotIn(column string, args []any, opts ...Option) *sql.Predicate {
	if len(args) == 0 {
		return sql.NotIn(column)
	}
	return valueInOp(column, args, opts, sql.OpNotIn)
}

func valueInOp(column string, args []any, opts []Option, op sql.Op) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		if allString(args) {
			opts = append(opts, Unquote(true))
		}
		if len(args) > 0 {
			opts = normalizePG(b, args[0], opts)
		}
		valuePath(b, column, opts...)
		b.WriteOp(op)
		b.Wrap(func(b *sql.Builder) {
			if s, ok := args[0].(*sql.Selector); ok {
				b.Join(s)
			} else {
				b.Args(args...)
			}
		})
	})
}

// LenEQ return a predicate for checking that an array length
// of a JSON (returned by the path) is equal to the given argument.
//
//	sqljson.LenEQ("a", 1, sqljson.Path("b"))
func LenEQ(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		lenPath(b, column, opts...)
		b.WriteOp(sql.OpEQ).Arg(size)
	})
}

// LenNEQ return a predicate for checking that an array length
// of a JSON (returned by the path) is not equal to the given argument.
//
//	sqljson.LenEQ("a", 1, sqljson.Path("b"))
func LenNEQ(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		lenPath(b, column, opts...)
		b.WriteOp(sql.OpNEQ).Arg(size)
	})
}

// LenGT return a predicate for checking that an array length
// of a JSON (returned by the path) is greater than the given
// argument.
//
//	sqljson.LenGT("a", 1, sqljson.Path("b"))
func LenGT(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		lenPath(b, column, opts...)
		b.WriteOp(sql.OpGT).Arg(size)
	})
}

// LenGTE return a predicate for checking that an array length
// of a JSON (returned by the path) is greater than or equal to
// the given argument.
//
//	sqljson.LenGTE("a", 1, sqljson.Path("b"))
func LenGTE(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		lenPath(b, column, opts...)
		b.WriteOp(sql.OpGTE).Arg(size)
	})
}

// LenLT return a predicate for checking that an array length
// of a JSON (returned by the path) is less than the given
// argument.
//
//	sqljson.LenLT("a", 1, sqljson.Path("b"))
func LenLT(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		lenPath(b, column, opts...)
		b.WriteOp(sql.OpLT).Arg(size)
	})
}

// LenLTE return a predicate for checking that an array length
// of a JSON (returned by the path) is less than or equal to
// the given argument.
//
//	sqljson.LenLTE("a", 1, sqljson.Path("b"))
func LenLTE(column string, size int, opts ...Option) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		lenPath(b, column, opts...)
		b.WriteOp(sql.OpLTE).Arg(size)
	})
}

// LenPath returns an SQL expression for getting the length
// of a JSON value (returned by the path).
func LenPath(column string, opts ...Option) sql.Querier {
	return sql.ExprFunc(func(b *sql.Builder) {
		lenPath(b, column, opts...)
	})
}

// OrderLen returns a custom predicate function (as defined in the doc),
// that sets the result order by the length of the given JSON value.
func OrderLen(column string, opts ...Option) func(*sql.Selector) {
	return func(s *sql.Selector) {
		s.OrderExpr(LenPath(column, opts...))
	}
}

// OrderLenDesc returns a custom predicate function (as defined in the doc), that
// sets the result order by the length of the given JSON value, but in descending order.
func OrderLenDesc(column string, opts ...Option) func(*sql.Selector) {
	return func(s *sql.Selector) {
		s.OrderExpr(
			sql.DescExpr(LenPath(column, opts...)),
		)
	}
}

// LenPath writes to the given SQL builder the JSON path for
// getting the length of a given JSON path.
//
//	sqljson.LenPath(b, Path("a", "b", "[1]", "c"))
func lenPath(b *sql.Builder, column string, opts ...Option) {
	path := identPath(column, opts...)
	path.length(b)
}

// Append writes to the given SQL builder the SQL command for appending JSON values
// into the array, optionally defined as a key. Note, the generated SQL will use the
// Go semantics, the JSON column/key will be set to the given Array in case it is `null`
// or NULL. For example:
//
//	Append(u, column, []string{"a", "b"})
//	UPDATE "t" SET "c" = CASE
//		WHEN ("c" IS NULL OR "c" = 'null'::jsonb)
//		THEN $1 ELSE "c" || $2 END
//
//	Append(u, column, []any{"a", 1}, sqljson.Path("a"))
//	UPDATE "t" SET "c" = CASE
//		WHEN (("c"->'a')::jsonb IS NULL OR ("c"->'a')::jsonb = 'null'::jsonb)
//		THEN jsonb_set("c", '{a}', $1, true) ELSE jsonb_set("c", '{a}', "c"->'a' || $2, true) END
func Append[T any](u *sql.UpdateBuilder, column string, elems []T, opts ...Option) {
	if len(elems) == 0 {
		u.AddError(fmt.Errorf("sqljson: cannot append an empty array to column %q", column))
		return
	}
	drv, err := newDriver(u.Dialect())
	if err != nil {
		u.AddError(err)
		return
	}
	vs := make([]any, len(elems))
	for i, e := range elems {
		vs[i] = e
	}
	drv.Append(u, column, vs, opts...)
}

// Option allows for calling database JSON paths with functional options.
type Option func(*PathOptions)

// Path sets the path to the JSON value of a column.
//
//	ValuePath(b, "column", Path("a", "b", "[1]", "c"))
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
func Unquote(unquote bool) Option {
	return func(p *PathOptions) {
		p.Unquote = unquote
	}
}

// Cast indicates that the result value should be cast to the given type.
//
//	ValuePath(b, "column", Path("a", "b", "[1]", "c"), Cast("int"))
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

// identPath creates a PathOptions for the given identifier.
func identPath(ident string, opts ...Option) *PathOptions {
	path := &PathOptions{Ident: ident}
	for i := range opts {
		opts[i](path)
	}
	return path
}

func (p *PathOptions) Query() (string, []any) {
	return p.Ident, nil
}

// ValuePath returns an SQL expression for getting the JSON
// value of a column with an optional path and cast options.
//
//	sqljson.ValueEQ(
//		column,
//		sqljson.ValuePath(column, Path("a"), Cast("int")),
//		sqljson.Path("a"),
//	)
func ValuePath(column string, opts ...Option) sql.Querier {
	return sql.ExprFunc(func(b *sql.Builder) {
		valuePath(b, column, opts...)
	})
}

// OrderValue returns a custom predicate function (as defined in the doc),
// that sets the result order by the given JSON value.
func OrderValue(column string, opts ...Option) func(*sql.Selector) {
	return func(s *sql.Selector) {
		s.OrderExpr(ValuePath(column, opts...))
	}
}

// OrderValueDesc returns a custom predicate function (as defined in the doc),
// that sets the result order by the given JSON value, but in descending order.
func OrderValueDesc(column string, opts ...Option) func(*sql.Selector) {
	return func(s *sql.Selector) {
		s.OrderExpr(
			sql.DescExpr(ValuePath(column, opts...)),
		)
	}
}

// valuePath writes to the given SQL builder the JSON path for
// getting the value of a given JSON path.
// Use sqljson.ValuePath for using a JSON value as an argument.
func valuePath(b *sql.Builder, column string, opts ...Option) {
	path := identPath(column, opts...)
	path.value(b)
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
		p.pgTextPath(b)
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
		p.pgTextPath(b)
		b.WriteByte(')')
	case b.Dialect() == dialect.MySQL:
		p.mysqlFunc("JSON_LENGTH", b)
	default:
		p.mysqlFunc("JSON_ARRAY_LENGTH", b)
	}
}

// mysqlFunc writes the JSON path in MySQL format for the
// given function. `JSON_EXTRACT("a", '$.b.c')`.
func (p *PathOptions) mysqlFunc(fn string, b *sql.Builder) {
	b.WriteString(fn).WriteByte('(')
	b.Ident(p.Ident).Comma()
	p.mysqlPath(b)
	b.WriteByte(')')
}

// mysqlPath writes the JSON path in MySQL (or SQLite) format.
func (p *PathOptions) mysqlPath(b *sql.Builder) {
	b.WriteString(`'$`)
	for _, p := range p.Path {
		switch _, isIndex := isJSONIdx(p); {
		case isIndex:
			b.WriteString(p)
		case p == "*" || isQuoted(p) || isIdentifier(p):
			b.WriteString("." + p)
		default:
			b.WriteString(`."` + p + `"`)
		}
	}
	b.WriteByte('\'')
}

// pgTextPath writes the JSON path in PostgreSQL text format: `"a"->'b'->>'c'`.
func (p *PathOptions) pgTextPath(b *sql.Builder) {
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

// pgArrayPath writes the JSON path in PostgreSQL array text[] format: '{a,1,b}'.
func (p *PathOptions) pgArrayPath(b *sql.Builder) {
	b.WriteString("'{")
	for i, s := range p.Path {
		if i > 0 {
			b.Comma()
		}
		if idx, ok := isJSONIdx(s); ok {
			s = idx
		}
		b.WriteString(s)
	}
	b.WriteString("}'")
}

// ParsePath parses the "dotpath" for the DotPath option.
//
//	"a.b"		=> ["a", "b"]
//	"a[1][2]"	=> ["a", "[1]", "[2]"]
//	"a.\"b.c\"	=> ["a", "\"b.c\""]
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
func normalizePG(b *sql.Builder, arg any, opts []Option) []Option {
	if b.Dialect() != dialect.Postgres {
		return opts
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
	}
	return append(base, opts...)
}

func isIdentifier(name string) bool {
	if name == "" {
		return false
	}
	for i, c := range name {
		if !unicode.IsLetter(c) && c != '_' && (i == 0 || !unicode.IsDigit(c)) {
			return false
		}
	}
	return true
}

func isQuoted(s string) bool {
	if s == "" {
		return false
	}
	return s[0] == '"' && s[len(s)-1] == '"'
}

// isJSONIdx reports whether the string represents a JSON index.
func isJSONIdx(s string) (string, bool) {
	if len(s) > 2 && s[0] == '[' && s[len(s)-1] == ']' && (isNumber(s[1:len(s)-1]) || s[1] == '#' && isNumber(s[2:len(s)-1])) {
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

// allString reports if the slice contains only strings.
func allString(v []any) bool {
	for i := range v {
		if _, ok := v[i].(string); !ok {
			return false
		}
	}
	return true
}

// marshalArg stringifies the given argument to a valid JSON document.
func marshalArg(arg any) any {
	if buf, err := json.Marshal(arg); err == nil {
		arg = string(buf)
	}
	return arg
}
