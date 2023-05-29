// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqljson_test

import (
	"strconv"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/stretchr/testify/require"
)

func TestWritePath(t *testing.T) {
	tests := []struct {
		input     sql.Querier
		wantQuery string
		wantArgs  []any
	}{
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", 1, sqljson.Path("b", "c", "[1]", "d"), sqljson.Cast("int"))),
			wantQuery: `SELECT * FROM "users" WHERE ("a"->'b'->'c'->1->>'d')::int = $1`,
			wantArgs:  []any{1},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", "a", sqljson.DotPath("b.c[1].d"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_EXTRACT(`a`, '$.b.c[1].d') = ?",
			wantArgs:  []any{"a"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", true, sqljson.DotPath("b.c[1].d"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_EXTRACT(`a`, '$.b.c[1].d') = true",
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", "a", sqljson.DotPath("b.\"c[1]\".d[1][2].e"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_EXTRACT(`a`, '$.b.\"c[1]\".d[1][2].e') = ?",
			wantArgs:  []any{"a"},
		},
		{
			input: sql.Select("*").
				From(sql.Table("test")).
				Where(sqljson.ValueEQ("j", sqljson.ValuePath("j", sqljson.DotPath("a.*.b")), sqljson.DotPath("a.*.c"))),
			wantQuery: "SELECT * FROM `test` WHERE JSON_EXTRACT(`j`, '$.a.*.c') = JSON_EXTRACT(`j`, '$.a.*.b')",
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("test")).
				Where(sqljson.ValueEQ("j", sqljson.ValuePath("j", sqljson.DotPath("a.*.b")), sqljson.DotPath("a.*.c"))),
			wantQuery: `SELECT * FROM "test" WHERE "j"->'a'->'*'->>'c' = "j"->'a'->'*'->'b'`,
		},
		{
			input: sql.Select("*").
				From(sql.Table("test")).
				Where(sqljson.HasKey("j", sqljson.DotPath("a.*.c"))),
			wantQuery: "SELECT * FROM `test` WHERE JSON_EXTRACT(`j`, '$.a.*.c') IS NOT NULL",
		},
		{
			input: sql.Select("*").
				From(sql.Table("test")).
				Where(sqljson.HasKey("j", sqljson.DotPath("a.*.c"))),
			wantQuery: "SELECT * FROM `test` WHERE JSON_EXTRACT(`j`, '$.a.*.c') IS NOT NULL",
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("test")).
				Where(sqljson.HasKey("j", sqljson.DotPath("attributes[1].body"))),
			wantQuery: "SELECT * FROM `test` WHERE JSON_TYPE(`j`, '$.attributes[1].body') IS NOT NULL",
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("test")).
				Where(sqljson.HasKey("j", sqljson.DotPath("a.*.c"))),
			wantQuery: "SELECT * FROM `test` WHERE JSON_TYPE(`j`, '$.a.*.c') IS NOT NULL",
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("test")).
				Where(
					sql.And(
						sql.GT("id", 100),
						sqljson.HasKey("j", sqljson.DotPath("a.*.c")),
						sql.EQ("active", true),
					),
				),
			wantQuery: "SELECT * FROM `test` WHERE `id` > ? AND JSON_TYPE(`j`, '$.a.*.c') IS NOT NULL AND `active`",
			wantArgs:  []any{100},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("test")).
				Where(sql.And(
					sql.EQ("e", 10),
					sqljson.ValueEQ("a", 1, sqljson.DotPath("b.c")),
				)),
			wantQuery: `SELECT * FROM "test" WHERE "e" = $1 AND ("a"->'b'->>'c')::int = $2`,
			wantArgs:  []any{10, 1},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", "a", sqljson.Path("b", "c", "[1]", "d"), sqljson.Unquote(true))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_UNQUOTE(JSON_EXTRACT(`a`, '$.b.c[1].d')) = ?",
			wantArgs:  []any{"a"},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", "a", sqljson.Path("b", "c", "[1]", "d"), sqljson.Unquote(true))),
			wantQuery: `SELECT * FROM "users" WHERE "a"->'b'->'c'->1->>'d' = $1`,
			wantArgs:  []any{"a"},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueEQ("a", 1, sqljson.Path("b", "c", "[1]", "d"), sqljson.Cast("int"))),
			wantQuery: `SELECT * FROM "users" WHERE ("a"->'b'->'c'->1->>'d')::int = $1`,
			wantArgs:  []any{1},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(
					sql.Or(
						sqljson.ValueNEQ("a", 1, sqljson.Path("b")),
						sqljson.ValueGT("a", 1, sqljson.Path("c")),
						sqljson.ValueGTE("a", 1.1, sqljson.Path("d")),
						sqljson.ValueLT("a", 1, sqljson.Path("e")),
						sqljson.ValueLTE("a", 1, sqljson.Path("f")),
					),
				),
			wantQuery: `SELECT * FROM "users" WHERE ("a"->>'b')::int <> $1 OR ("a"->>'c')::int > $2 OR ("a"->>'d')::float >= $3 OR ("a"->>'e')::int < $4 OR ("a"->>'f')::int <= $5`,
			wantArgs:  []any{1, 1, 1.1, 1, 1},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.LenEQ("a", 1)),
			wantQuery: `SELECT * FROM "users" WHERE JSONB_ARRAY_LENGTH("a") = $1`,
			wantArgs:  []any{1},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.LenEQ("a", 1)),
			wantQuery: "SELECT * FROM `users` WHERE JSON_LENGTH(`a`, '$') = ?",
			wantArgs:  []any{1},
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.LenEQ("a", 1)),
			wantQuery: "SELECT * FROM `users` WHERE JSON_ARRAY_LENGTH(`a`, '$') = ?",
			wantArgs:  []any{1},
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("users")).
				Where(
					sql.Or(
						sqljson.LenGT("a", 1, sqljson.Path("b")),
						sqljson.LenGTE("a", 1, sqljson.Path("c")),
						sqljson.LenLT("a", 1, sqljson.Path("d")),
						sqljson.LenLTE("a", 1, sqljson.Path("e")),
					),
				),
			wantQuery: "SELECT * FROM `users` WHERE JSON_ARRAY_LENGTH(`a`, '$.b') > ? OR JSON_ARRAY_LENGTH(`a`, '$.c') >= ? OR JSON_ARRAY_LENGTH(`a`, '$.d') < ? OR JSON_ARRAY_LENGTH(`a`, '$.e') <= ?",
			wantArgs:  []any{1, 1, 1, 1},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueContains("tags", "foo")),
			wantQuery: "SELECT * FROM `users` WHERE JSON_CONTAINS(`tags`, ?, '$') = ?",
			wantArgs:  []any{"\"foo\"", 1},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueContains("tags", 1, sqljson.Path("a"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_CONTAINS(`tags`, ?, '$.a') = ?",
			wantArgs:  []any{"1", 1},
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueContains("tags", "foo")),
			wantQuery: "SELECT * FROM `users` WHERE EXISTS(SELECT * FROM JSON_EACH(`tags`, '$') WHERE `value` = ?)",
			wantArgs:  []any{"foo"},
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueContains("tags", 1, sqljson.Path("a"))),
			wantQuery: "SELECT * FROM `users` WHERE EXISTS(SELECT * FROM JSON_EACH(`tags`, '$.a') WHERE `value` = ?)",
			wantArgs:  []any{1},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueContains("tags", "foo")),
			wantQuery: "SELECT * FROM \"users\" WHERE \"tags\" @> $1",
			wantArgs:  []any{"\"foo\""},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueContains("tags", 1, sqljson.Path("a"))),
			wantQuery: "SELECT * FROM \"users\" WHERE (\"tags\"->'a')::jsonb @> $1",
			wantArgs:  []any{"1"},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIsNull("c", sqljson.Path("a"))),
			wantQuery: `SELECT * FROM "users" WHERE ("c"->'a')::jsonb = 'null'::jsonb`,
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIsNull("c", sqljson.Path("a"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_CONTAINS(`c`, 'null', '$.a')",
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIsNull("c", sqljson.Path("a"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_TYPE(`c`, '$.a') = 'null'",
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIsNotNull("c", sqljson.Path("a"))),
			wantQuery: `SELECT * FROM "users" WHERE ("c"->'a')::jsonb <> 'null'::jsonb`,
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIsNotNull("c", sqljson.Path("a"))),
			wantQuery: "SELECT * FROM `users` WHERE NOT(JSON_CONTAINS(`c`, 'null', '$.a'))",
		},
		{
			input: sql.Dialect(dialect.SQLite).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIsNotNull("c", sqljson.Path("a"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_TYPE(`c`, '$.a') <> 'null'",
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.StringContains("a", "substr", sqljson.Path("b", "c", "[1]", "d"))),
			wantQuery: `SELECT * FROM "users" WHERE "a"->'b'->'c'->1->>'d' LIKE $1`,
			wantArgs:  []any{"%substr%"},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(
					sql.And(
						sqljson.StringContains("a", "c", sqljson.Path("a")),
						sqljson.StringContains("b", "d", sqljson.Path("b")),
					),
				),
			wantQuery: `SELECT * FROM "users" WHERE "a"->>'a' LIKE $1 AND "b"->>'b' LIKE $2`,
			wantArgs:  []any{"%c%", "%d%"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.StringContains("a", "substr", sqljson.Path("b", "c", "[1]", "d"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_UNQUOTE(JSON_EXTRACT(`a`, '$.b.c[1].d')) LIKE ?",
			wantArgs:  []any{"%substr%"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(
					sql.And(
						sqljson.StringContains("a", "c", sqljson.Path("a")),
						sqljson.StringContains("b", "d", sqljson.Path("b")),
					),
				),
			wantQuery: "SELECT * FROM `users` WHERE JSON_UNQUOTE(JSON_EXTRACT(`a`, '$.a')) LIKE ? AND JSON_UNQUOTE(JSON_EXTRACT(`b`, '$.b')) LIKE ?",
			wantArgs:  []any{"%c%", "%d%"},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.StringHasPrefix("a", "substr", sqljson.Path("b", "c", "[1]", "d"))),
			wantQuery: `SELECT * FROM "users" WHERE "a"->'b'->'c'->1->>'d' LIKE $1`,
			wantArgs:  []any{"substr%"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.StringHasPrefix("a", "substr", sqljson.Path("b", "c", "[1]", "d"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_UNQUOTE(JSON_EXTRACT(`a`, '$.b.c[1].d')) LIKE ?",
			wantArgs:  []any{"substr%"},
		},
		{
			input: sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.StringHasSuffix("a", "substr", sqljson.Path("b", "c", "[1]", "d"))),
			wantQuery: `SELECT * FROM "users" WHERE "a"->'b'->'c'->1->>'d' LIKE $1`,
			wantArgs:  []any{"%substr"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.StringHasSuffix("a", "substr", sqljson.Path("b", "c", "[1]", "d"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_UNQUOTE(JSON_EXTRACT(`a`, '$.b.c[1].d')) LIKE ?",
			wantArgs:  []any{"%substr"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIn("a", []any{"a", "b"}, sqljson.Path("b"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_UNQUOTE(JSON_EXTRACT(`a`, '$.b')) IN (?, ?)",
			wantArgs:  []any{"a", "b"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIn("a", []any{1, 2}, sqljson.Path("b"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_EXTRACT(`a`, '$.b') IN (?, ?)",
			wantArgs:  []any{1, 2},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIn("a", []any{1, "a"}, sqljson.Path("b"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_EXTRACT(`a`, '$.b') IN (?, ?)",
			wantArgs:  []any{1, "a"},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				Where(sqljson.ValueIn("a", []any{1, 2}, sqljson.Path("foo-bar", "3000"))),
			wantQuery: "SELECT * FROM `users` WHERE JSON_EXTRACT(`a`, '$.\"foo-bar\".\"3000\"') IN (?, ?)",
			wantArgs:  []any{1, 2},
		},
		{
			input: sql.Dialect(dialect.MySQL).
				Select("*").
				From(sql.Table("users")).
				OrderExpr(
					sqljson.LenPath("a", sqljson.Path("b")),
				),
			wantQuery: "SELECT * FROM `users` ORDER BY JSON_LENGTH(`a`, '$.b')",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			query, args := tt.input.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}

func TestParsePath(t *testing.T) {
	tests := []struct {
		input    string
		wantPath []string
		wantErr  bool
	}{
		{
			input:    "a.b.c",
			wantPath: []string{"a", "b", "c"},
		},
		{
			input:    "a[1][2]",
			wantPath: []string{"a", "[1]", "[2]"},
		},
		{
			input:    "a[1][2].b",
			wantPath: []string{"a", "[1]", "[2]", "b"},
		},
		{
			input:    `a."b.c[0]"`,
			wantPath: []string{"a", `"b.c[0]"`},
		},
		{
			input:    `a."b.c[0]".d`,
			wantPath: []string{"a", `"b.c[0]"`, "d"},
		},
		{
			input: `...`,
		},
		{
			input:    `.a.b.`,
			wantPath: []string{"a", "b"},
		},
		{
			input:   `a."`,
			wantErr: true,
		},
		{
			input:   `a[`,
			wantErr: true,
		},
		{
			input:   `a[a]`,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			path, err := sqljson.ParsePath(tt.input)
			require.Equal(t, tt.wantPath, path)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		input     sql.Querier
		wantQuery string
		wantArgs  []any
	}{
		{
			input: func() sql.Querier {
				u := sql.Dialect(dialect.Postgres).Update("t")
				sqljson.Append(u, "c", []string{"a"})
				return u
			}(),
			wantQuery: `UPDATE "t" SET "c" = CASE WHEN ("c" IS NULL OR "c" = 'null'::jsonb) THEN $1 ELSE "c" || $2 END`,
			wantArgs:  []any{`["a"]`, `["a"]`},
		},
		{
			input: func() sql.Querier {
				u := sql.Dialect(dialect.Postgres).Update("t")
				sqljson.Append(u, "c", []string{"a"}, sqljson.Path("a"))
				return u
			}(),
			wantQuery: `UPDATE "t" SET "c" = CASE WHEN (("c"->'a')::jsonb IS NULL OR ("c"->'a')::jsonb = 'null'::jsonb) THEN jsonb_set("c", '{a}', $1, true) ELSE jsonb_set("c", '{a}', "c"->'a' || $2, true) END`,
			wantArgs:  []any{`["a"]`, `["a"]`},
		},
		{
			input: func() sql.Querier {
				u := sql.Dialect(dialect.SQLite).Update("t")
				sqljson.Append(u, "c", []string{"a"})
				return u
			}(),
			wantQuery: "UPDATE `t` SET `c` = CASE WHEN (JSON_TYPE(`c`, '$') IS NULL OR JSON_TYPE(`c`, '$') = 'null') THEN ? ELSE JSON_INSERT(`c`, '$[#]', ?) END",
			wantArgs:  []any{`["a"]`, "a"},
		},
		{
			input: func() sql.Querier {
				u := sql.Dialect(dialect.SQLite).Update("t")
				sqljson.Append(u, "c", []any{"a", struct{}{}}, sqljson.Path("a"))
				return u
			}(),
			wantQuery: "UPDATE `t` SET `c` = CASE WHEN (JSON_TYPE(`c`, '$.a') IS NULL OR JSON_TYPE(`c`, '$.a') = 'null') THEN JSON_SET(`c`, '$.a', JSON(?)) ELSE JSON_INSERT(`c`, '$.a[#]', ?, '$.a[#]', JSON(?)) END",
			wantArgs:  []any{`["a",{}]`, "a", "{}"},
		},
		{
			input: func() sql.Querier {
				u := sql.Dialect(dialect.MySQL).Update("t")
				sqljson.Append(u, "c", []string{"a"})
				return u
			}(),
			wantQuery: "UPDATE `t` SET `c` = CASE WHEN (JSON_TYPE(JSON_EXTRACT(`c`, '$')) IS NULL OR JSON_TYPE(JSON_EXTRACT(`c`, '$')) = 'NULL') THEN JSON_ARRAY(?) ELSE JSON_ARRAY_APPEND(`c`, '$', ?) END",
			wantArgs:  []any{"a", "a"},
		},
		{
			input: func() sql.Querier {
				u := sql.Dialect(dialect.MySQL).Update("t")
				sqljson.Append(u, "c", []string{"a"}, sqljson.Path("a"))
				return u
			}(),
			wantQuery: "UPDATE `t` SET `c` = CASE WHEN (JSON_TYPE(JSON_EXTRACT(`c`, '$.a')) IS NULL OR JSON_TYPE(JSON_EXTRACT(`c`, '$.a')) = 'NULL') THEN JSON_SET(`c`, '$.a', JSON_ARRAY(?)) ELSE JSON_ARRAY_APPEND(`c`, '$.a', ?) END",
			wantArgs:  []any{"a", "a"},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			query, args := tt.input.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}
