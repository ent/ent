// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package json

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/facebook/ent/entc/integration/json/ent/schema"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqljson"
	"github.com/facebook/ent/entc/integration/json/ent"
	"github.com/facebook/ent/entc/integration/json/ent/migrate"
	"github.com/facebook/ent/entc/integration/json/ent/user"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		t.Run(version, func(t *testing.T) {
			db, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/", port))
			require.NoError(t, err)
			defer db.Close()
			ctx := context.Background()
			err = db.Exec(ctx, "CREATE DATABASE IF NOT EXISTS json", []interface{}{}, nil)
			require.NoError(t, err, "creating database")
			defer db.Exec(ctx, "DROP DATABASE IF EXISTS json", []interface{}{}, nil)
			client, err := ent.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/json", port))
			require.NoError(t, err, "connecting to json database")
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)

			URL(t, client)
			Dirs(t, client)
			Ints(t, client)
			Floats(t, client)
			Strings(t, client)
			RawMessage(t, client)
			// Skip predicates test for MySQL old versions.
			if version != "56" {
				Predicates(t, client)
			}
		})
	}
}

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5433} {
		t.Run(version, func(t *testing.T) {
			dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", port)
			db, err := sql.Open(dialect.Postgres, dsn)
			require.NoError(t, err)
			defer db.Close()
			ctx := context.Background()
			err = db.Exec(ctx, "CREATE DATABASE json", []interface{}{}, nil)
			require.NoError(t, err, "creating database")
			defer db.Exec(ctx, "DROP DATABASE IF EXISTS json", []interface{}{}, nil)

			client, err := ent.Open(dialect.Postgres, dsn+" dbname=json")
			require.NoError(t, err, "connecting to json database")
			defer client.Close()
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)

			URL(t, client)
			Dirs(t, client)
			Ints(t, client)
			Floats(t, client)
			Strings(t, client)
			RawMessage(t, client)
			Predicates(t, client)
		})
	}
}

func TestSQLite(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))

	URL(t, client)
	Dirs(t, client)
	Ints(t, client)
	Floats(t, client)
	Strings(t, client)
	RawMessage(t, client)
	Predicates(t, client)
}

func Ints(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	ints := []int{1, 2, 3}
	usr := client.User.Create().SetInts(ints).SaveX(ctx)
	require.Equal(t, ints, usr.Ints)
	require.Equal(t, ints, client.User.GetX(ctx, usr.ID).Ints)
	usr = usr.Update().SetInts(ints[:1]).SaveX(ctx)
	require.Equal(t, ints[:1], usr.Ints)
	require.Equal(t, ints[:1], client.User.GetX(ctx, usr.ID).Ints)
	usr = usr.Update().ClearInts().SaveX(ctx)
	require.Empty(t, usr.Ints)
	require.Empty(t, client.User.GetX(ctx, usr.ID).Ints)
}

func Floats(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	flts := []float64{1, 2, 3}
	usr := client.User.Create().SetFloats(flts).SaveX(ctx)
	require.Equal(t, flts, usr.Floats)
	require.Equal(t, flts, client.User.GetX(ctx, usr.ID).Floats)
	usr = usr.Update().SetFloats(flts[:1]).SaveX(ctx)
	require.Equal(t, flts[:1], usr.Floats)
	require.Equal(t, flts[:1], client.User.GetX(ctx, usr.ID).Floats)
	usr = usr.Update().ClearFloats().SaveX(ctx)
	require.Empty(t, usr.Floats)
	require.Empty(t, client.User.GetX(ctx, usr.ID).Floats)
}

func Strings(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	str := []string{"a", "b", "c"}
	usr := client.User.Create().SetStrings(str).SaveX(ctx)
	require.Equal(t, str, usr.Strings)
	require.Equal(t, str, client.User.GetX(ctx, usr.ID).Strings)
	usr = usr.Update().SetStrings(str[:1]).SaveX(ctx)
	require.Equal(t, str[:1], usr.Strings)
	require.Equal(t, str[:1], client.User.GetX(ctx, usr.ID).Strings)
	require.Equal(t, 1, client.User.Query().Where(user.StringsNotNil()).CountX(ctx))
	usr = usr.Update().ClearStrings().SaveX(ctx)
	require.Empty(t, usr.Strings)
	require.Empty(t, client.User.GetX(ctx, usr.ID).Strings)
	require.Zero(t, client.User.Query().Where(user.StringsNotNil()).CountX(ctx))
}

func RawMessage(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	raw := json.RawMessage("{}")
	usr := client.User.Create().SetRaw(raw).SaveX(ctx)
	require.Equal(t, raw, usr.Raw)
	require.Equal(t, raw, client.User.GetX(ctx, usr.ID).Raw)
}

func Dirs(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	dirs := []http.Dir{"dev", "usr"}
	usr := client.User.Create().SetDirs(dirs).SaveX(ctx)
	require.Equal(t, dirs, usr.Dirs)
	require.Equal(t, dirs, client.User.GetX(ctx, usr.ID).Dirs)
}

func URL(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	u, err := url.Parse("https://github.com/a8m")
	require.NoError(t, err)
	usr := client.User.Create().SetURL(u).SaveX(ctx)
	require.Equal(t, u, usr.URL)
	require.Equal(t, u, client.User.GetX(ctx, usr.ID).URL)
}

func Predicates(t *testing.T, client *ent.Client) {
	ctx := context.Background()

	client.User.Delete().ExecX(ctx)
	u1, err := url.Parse("https://github.com/a8m/ent")
	require.NoError(t, err)
	u2, err := url.Parse("ftp://a8m@github.com/ent")
	require.NoError(t, err)
	users, err := client.User.CreateBulk(
		client.User.Create().SetURL(u1),
		client.User.Create().SetURL(u2),
	).Save(ctx)
	require.NoError(t, err)
	require.Len(t, users, 2)

	count, err := client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.HasKey(user.FieldURL, sqljson.Path("Scheme")))
	}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, count)

	count, err = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.Not(sqljson.HasKey(user.FieldURL, sqljson.Path("Scheme"))))
	}).Count(ctx)
	require.NoError(t, err)
	require.Zero(t, count)

	count, err = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueEQ(user.FieldURL, "https", sqljson.Path("Scheme")))
	}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueNEQ(user.FieldURL, "https", sqljson.Path("Scheme")))
	}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, count)

	client.User.Delete().ExecX(ctx)
	users, err = client.User.CreateBulk(
		client.User.Create().SetT(&schema.T{I: 1, F: 1.1, T: &schema.T{I: 10}}),
		client.User.Create().SetT(&schema.T{I: 2, F: 2.2, T: &schema.T{I: 20, T: &schema.T{I: 30}}}),
	).Save(ctx)
	require.NoError(t, err)
	require.Len(t, users, 2)

	count, err = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueGTE(user.FieldT, 1, sqljson.Path("i")))
	}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, count)

	count, err = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueLTE(user.FieldT, 30, sqljson.DotPath("t.t.i")))
	}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, count)

	count, err = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(
			sql.Or(
				sqljson.ValueEQ(user.FieldT, 1.1, sqljson.Path("f")),
				sqljson.ValueEQ(user.FieldT, 30, sqljson.DotPath("t.t.i")),
			),
		)
	}).Count(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, count)

	client.User.Delete().ExecX(ctx)
	users, err = client.User.CreateBulk(
		client.User.Create().SetInts([]int{1}),
		client.User.Create().SetInts([]int{1, 2}).SetT(&schema.T{Li: []int{1, 2}, Ls: []string{"a"}}),
		client.User.Create().SetInts([]int{1, 2, 3}).SetT(&schema.T{Li: []int{3, 4}, Ls: []string{"b"}}),
	).Save(ctx)
	require.NoError(t, err)

	for _, u := range users {
		r := client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.LenEQ(user.FieldInts, len(u.Ints)))
		}).OnlyX(ctx)
		require.Equal(t, u.Ints, r.Ints)
	}

	r := client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(user.FieldInts, 3))
	}).OnlyX(ctx)
	require.Contains(t, r.Ints, 3)
	r = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(user.FieldT, 3, sqljson.Path("li")))
	}).OnlyX(ctx)
	require.Contains(t, r.T.Li, 3)

	r = client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(user.FieldT, "a", sqljson.Path("ls")))
	}).OnlyX(ctx)
	require.Contains(t, r.T.Ls, "a")
}
