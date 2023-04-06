// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqljson

import (
	"fmt"
	"reflect"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

type sqlite struct{}

// Append implements the driver.Append method.
func (d *sqlite) Append(u *sql.UpdateBuilder, column string, elems []any, opts ...Option) {
	setCase(u, column, when{
		Cond: func(b *sql.Builder) {
			typ := func(b *sql.Builder) *sql.Builder {
				return b.WriteString("JSON_TYPE").Wrap(func(b *sql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).mysqlPath(b)
				})
			}
			typ(b).WriteOp(sql.OpIsNull)
			b.WriteString(" OR ")
			typ(b).WriteOp(sql.OpEQ).WriteString("'null'")
		},
		Then: func(b *sql.Builder) {
			if len(opts) > 0 {
				b.WriteString("JSON_SET").Wrap(func(b *sql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).mysqlPath(b)
					b.Comma().Argf("JSON(?)", marshalArg(elems))
				})
			} else {
				b.Arg(marshalArg(elems))
			}
		},
		Else: func(b *sql.Builder) {
			b.WriteString("JSON_INSERT").Wrap(func(b *sql.Builder) {
				b.Ident(column).Comma()
				// If no path was provided the top-level value is
				// a JSON array. i.e. JSON_INSERT(c, '$[#]', ?).
				path := func(b *sql.Builder) { b.WriteString("'$[#]'") }
				if len(opts) > 0 {
					p := identPath(column, opts...)
					p.Path = append(p.Path, "[#]")
					path = p.mysqlPath
				}
				for i, e := range elems {
					if i > 0 {
						b.Comma()
					}
					path(b)
					b.Comma()
					d.appendArg(b, e)
				}
			})
		},
	})
}

func (d *sqlite) appendArg(b *sql.Builder, v any) {
	switch {
	case !isPrimitive(v):
		b.Argf("JSON(?)", marshalArg(v))
	default:
		b.Arg(v)
	}
}

type mysql struct{}

// Append implements the driver.Append method.
func (d *mysql) Append(u *sql.UpdateBuilder, column string, elems []any, opts ...Option) {
	setCase(u, column, when{
		Cond: func(b *sql.Builder) {
			typ := func(b *sql.Builder) *sql.Builder {
				b.WriteString("JSON_TYPE(JSON_EXTRACT(")
				b.Ident(column).Comma()
				identPath(column, opts...).mysqlPath(b)
				return b.WriteString("))")
			}
			typ(b).WriteOp(sql.OpIsNull)
			b.WriteString(" OR ")
			typ(b).WriteOp(sql.OpEQ).WriteString("'NULL'")
		},
		Then: func(b *sql.Builder) {
			if len(opts) > 0 {
				b.WriteString("JSON_SET").Wrap(func(b *sql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).mysqlPath(b)
					b.Comma().WriteString("JSON_ARRAY(").Args(d.marshalArgs(elems)...).WriteByte(')')
				})
			} else {
				b.WriteString("JSON_ARRAY(").Args(d.marshalArgs(elems)...).WriteByte(')')
			}
		},
		Else: func(b *sql.Builder) {
			b.WriteString("JSON_ARRAY_APPEND").Wrap(func(b *sql.Builder) {
				b.Ident(column).Comma()
				for i, e := range elems {
					if i > 0 {
						b.Comma()
					}
					identPath(column, opts...).mysqlPath(b)
					b.Comma()
					d.appendArg(b, e)
				}
			})
		},
	})
}

func (d *mysql) marshalArgs(args []any) []any {
	vs := make([]any, len(args))
	for i, v := range args {
		if !isPrimitive(v) {
			v = marshalArg(v)
		}
		vs[i] = v
	}
	return vs
}

func (d *mysql) appendArg(b *sql.Builder, v any) {
	switch {
	case !isPrimitive(v):
		b.Argf("CAST(? AS JSON)", marshalArg(v))
	default:
		b.Arg(v)
	}
}

type postgres struct{}

// Append implements the driver.Append method.
func (*postgres) Append(u *sql.UpdateBuilder, column string, elems []any, opts ...Option) {
	setCase(u, column, when{
		Cond: func(b *sql.Builder) {
			valuePath(b, column, append(opts, Cast("jsonb"))...)
			b.WriteOp(sql.OpIsNull)
			b.WriteString(" OR ")
			valuePath(b, column, append(opts, Cast("jsonb"))...)
			b.WriteOp(sql.OpEQ).WriteString("'null'::jsonb")
		},
		Then: func(b *sql.Builder) {
			if len(opts) > 0 {
				b.WriteString("jsonb_set").Wrap(func(b *sql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).pgArrayPath(b)
					b.Comma().Arg(marshalArg(elems))
					b.Comma().WriteString("true")
				})
			} else {
				b.Arg(marshalArg(elems))
			}
		},
		Else: func(b *sql.Builder) {
			if len(opts) > 0 {
				b.WriteString("jsonb_set").Wrap(func(b *sql.Builder) {
					b.Ident(column).Comma()
					identPath(column, opts...).pgArrayPath(b)
					b.Comma()
					path := identPath(column, opts...)
					path.value(b)
					b.WriteString(" || ").Arg(marshalArg(elems))
					b.Comma().WriteString("true")
				})
			} else {
				b.Ident(column).WriteString(" || ").Arg(marshalArg(elems))
			}
		},
	})
}

// driver groups all dialect-specific methods.
type driver interface {
	Append(u *sql.UpdateBuilder, column string, elems []any, opts ...Option)
}

func newDriver(name string) (driver, error) {
	switch name {
	case dialect.SQLite:
		return (*sqlite)(nil), nil
	case dialect.MySQL:
		return (*mysql)(nil), nil
	case dialect.Postgres:
		return (*postgres)(nil), nil
	default:
		return nil, fmt.Errorf("sqljson: unknown driver %q", name)
	}
}

type when struct{ Cond, Then, Else func(*sql.Builder) }

// setCase sets the column value using the "CASE WHEN" statement.
// The x defines the condition/predicate, t is the true (if) case,
// and 'f' defines the false (else).
func setCase(u *sql.UpdateBuilder, column string, w when) {
	u.Set(column, sql.ExprFunc(func(b *sql.Builder) {
		b.WriteString("CASE WHEN ").Wrap(func(b *sql.Builder) {
			w.Cond(b)
		})
		b.WriteString(" THEN ")
		w.Then(b)
		b.WriteString(" ELSE ")
		w.Else(b)
		b.WriteString(" END")
	}))
}

func isPrimitive(v any) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct, reflect.Ptr, reflect.Interface:
		return false
	}
	return true
}
