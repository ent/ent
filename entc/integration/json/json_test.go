// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package json

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/entc/integration/json/ent"
	"entgo.io/ent/entc/integration/json/ent/migrate"
	"entgo.io/ent/entc/integration/json/ent/schema"
	"entgo.io/ent/entc/integration/json/ent/user"

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
			err = db.Exec(ctx, "CREATE DATABASE IF NOT EXISTS json", []any{}, nil)
			require.NoError(t, err, "creating database")
			defer db.Exec(ctx, "DROP DATABASE IF EXISTS json", []any{}, nil)
			client, err := ent.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/json", port))
			require.NoError(t, err, "connecting to json database")
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)

			URL(t, client)
			Dirs(t, client)
			Floats(t, client)
			NetAddr(t, client)
			RawMessage(t, client)
			Any(t, client)
			// Skip tests with JSON functions for old MySQL versions.
			if version != "56" {
				URLs(t, client)
				Ints(t, client)
				Strings(t, client)
				IntsValidate(t, client)
				FloatsValidate(t, client)
				StringsValidate(t, client)
				Predicates(t, client)
				Order(t, client)
			}
			Scan(t, client)
		})
	}
}

func TestMaria(t *testing.T) {
	for version, port := range map[string]int{"105": 4306, "102": 4307} {
		t.Run(version, func(t *testing.T) {
			db, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/", port))
			require.NoError(t, err)
			defer db.Close()
			ctx := context.Background()
			err = db.Exec(ctx, "CREATE DATABASE IF NOT EXISTS json", []any{}, nil)
			require.NoError(t, err, "creating database")
			defer db.Exec(ctx, "DROP DATABASE IF EXISTS json", []any{}, nil)
			client, err := ent.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/json", port))
			require.NoError(t, err, "connecting to json database")
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)
			// We run the migration twice to check that migration handles
			// the JSON columns, since MariaDB stores them as longtext.
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)

			URL(t, client)
			Dirs(t, client)
			Ints(t, client)
			Floats(t, client)
			Strings(t, client)
			IntsValidate(t, client)
			FloatsValidate(t, client)
			StringsValidate(t, client)
			NetAddr(t, client)
			RawMessage(t, client)
			Any(t, client)
			Predicates(t, client)
			Scan(t, client)
			Order(t, client)
		})
	}
}

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5433, "13": 5434} {
		t.Run(version, func(t *testing.T) {
			dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", port)
			db, err := sql.Open(dialect.Postgres, dsn)
			require.NoError(t, err)
			defer db.Close()
			ctx := context.Background()
			err = db.Exec(ctx, "CREATE DATABASE json", []any{}, nil)
			require.NoError(t, err, "creating database")
			defer db.Exec(ctx, "DROP DATABASE IF EXISTS json", []any{}, nil)

			client, err := ent.Open(dialect.Postgres, dsn+" dbname=json")
			require.NoError(t, err, "connecting to json database")
			defer client.Close()
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)

			URL(t, client)
			URLs(t, client)
			Dirs(t, client)
			Ints(t, client)
			Floats(t, client)
			Strings(t, client)
			IntsValidate(t, client)
			FloatsValidate(t, client)
			StringsValidate(t, client)
			NetAddr(t, client)
			RawMessage(t, client)
			Any(t, client)
			Predicates(t, client)
			Scan(t, client)
			Order(t, client)
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
	URLs(t, client)
	Dirs(t, client)
	Ints(t, client)
	Floats(t, client)
	Strings(t, client)
	IntsValidate(t, client)
	FloatsValidate(t, client)
	StringsValidate(t, client)
	NetAddr(t, client)
	RawMessage(t, client)
	Any(t, client)
	Predicates(t, client)
	Scan(t, client)
	Order(t, client)
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

	usr = client.User.Create().SaveX(ctx)
	require.Equal(t, []int{1, 2, 3}, usr.Ints)
	usr = client.User.GetX(ctx, usr.ID)
	require.Equal(t, []int{1, 2, 3}, usr.Ints)
	usr = usr.Update().AppendInts([]int{4, 5, 6}).SaveX(ctx)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, usr.Ints)
}

func IntsValidate(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	xs := []int{1, 2, 3}
	err := client.User.Create().SetIntsValidate(xs).Exec(ctx)
	require.ErrorIs(t, err, schema.ErrValidate)
	err = client.User.Create().Exec(ctx)
	require.NoError(t, err, schema.ErrValidate)
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

func FloatsValidate(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	xs := []float64{1, 2, 3}
	err := client.User.Create().SetFloatsValidate(xs).Exec(ctx)
	require.ErrorIs(t, err, schema.ErrValidate)
	err = client.User.Create().Exec(ctx)
	require.NoError(t, err, schema.ErrValidate)
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

	t.Run("Modifier API", func(t *testing.T) {
		// Append to an empty array.
		usr.Update().SetStrings([]string{}).SetT(&schema.T{Ls: []string{}}).ExecX(ctx)
		usr = usr.Update().Modify(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, []string{"foo"})
			sqljson.Append(u, user.FieldT, []string{"foo"}, sqljson.Path("ls"))
		}).SaveX(ctx)
		require.Equal(t, []string{"foo"}, usr.Strings)
		require.Equal(t, []string{"foo"}, usr.T.Ls)

		// Set a 'null' (or an undefined) value.
		usr.Update().ClearStrings().ClearT().ExecX(ctx)
		usr.Update().SetStrings(nil).SetT(&schema.T{Ls: nil}).ExecX(ctx)
		usr = client.User.GetX(ctx, usr.ID)
		usr = usr.Update().Modify(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, []string{"foo"})
			sqljson.Append(u, user.FieldT, []string{"foo"}, sqljson.Path("ls"))
		}).SaveX(ctx)
		require.Equal(t, []string{"foo"}, usr.Strings)
		require.Equal(t, []string{"foo"}, usr.T.Ls)
		usr = usr.Update().Modify(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, []string{"bar", "baz"})
			sqljson.Append(u, user.FieldT, []string{"bar", "baz"}, sqljson.Path("ls"))
		}).SaveX(ctx)
		require.Equal(t, []string{"foo", "bar", "baz"}, usr.Strings)
		require.Equal(t, []string{"foo", "bar", "baz"}, usr.T.Ls)

		// Set a NULL (or an undefined) value.
		usr.Update().ClearStrings().ExecX(ctx)
		usr = usr.Update().Modify(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, []string{"foo"})
		}).SaveX(ctx)
		require.Equal(t, []string{"foo"}, usr.Strings)
	})

	t.Run("Fluent API", func(t *testing.T) {
		// Append to an empty array.
		usr.Update().SetStrings([]string{}).SetInts([]int{}).ExecX(ctx)
		usr = usr.Update().AppendStrings([]string{"foo"}).AppendInts([]int{1}).SaveX(ctx)
		require.Equal(t, []int{1}, usr.Ints)
		require.Equal(t, []string{"foo"}, usr.Strings)
		usr = client.User.GetX(ctx, usr.ID)
		require.Equal(t, []int{1}, usr.Ints)
		require.Equal(t, []string{"foo"}, usr.Strings)
		usr = usr.Update().AppendStrings([]string{"bar", "baz"}).AppendInts([]int{2, 3}).SaveX(ctx)
		require.Equal(t, []int{1, 2, 3}, usr.Ints)
		require.Equal(t, []string{"foo", "bar", "baz"}, usr.Strings)

		// Set a 'null' (or an undefined) value.
		usr.Update().ClearStrings().SetInts(nil).SetDirs(nil).ExecX(ctx)
		usr = client.User.GetX(ctx, usr.ID)
		require.Empty(t, usr.Ints)
		require.Empty(t, usr.Strings)
		usr = usr.Update().AppendStrings([]string{"foo"}).AppendInts([]int{1}).SaveX(ctx)
		require.Equal(t, []int{1}, usr.Ints)
		require.Equal(t, []string{"foo"}, usr.Strings)

		usr.Update().AppendStrings([]string{"bar"}).SetStrings([]string{"baz"}).ExecX(ctx)
		require.Equal(t, []string{"baz"}, client.User.GetX(ctx, usr.ID).Strings)
		usr.Update().AppendStrings([]string{"bar"}).SetStrings([]string{"baz"}).ExecX(ctx)
		require.Equal(t, []string{"baz"}, client.User.GetX(ctx, usr.ID).Strings)
		usr.Update().AppendStrings([]string{"bar"}).ClearStrings().AppendDirs([]http.Dir{"/etc", "/dev"}).ExecX(ctx)
		usr = client.User.GetX(ctx, usr.ID)
		require.Empty(t, usr.Strings)
		require.Equal(t, []http.Dir{"/etc", "/dev"}, usr.Dirs)
	})
}

func StringsValidate(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	xs := []string{"a", "b", "c"}
	err := client.User.Create().SetStringsValidate(xs).Exec(ctx)
	require.ErrorIs(t, err, schema.ErrValidate)
	err = client.User.Create().Exec(ctx)
	require.NoError(t, err, schema.ErrValidate)
}

func Any(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	u := client.User.Create().SetUnknown("string").SaveX(ctx)
	require.Equal(t, "string", u.Unknown)
	u = u.Update().SetUnknown([]any{1, 2, 3}).SaveX(ctx)
	require.Equal(t, []any{1.0, 2.0, 3.0}, u.Unknown)
	require.Equal(t, []any{1.0, 2.0, 3.0}, client.User.GetX(ctx, u.ID).Unknown)
}

func RawMessage(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	raw := json.RawMessage("{}")
	usr := client.User.Create().SetRaw(raw).SaveX(ctx)
	require.Equal(t, raw, usr.Raw)
	require.Equal(t, raw, client.User.GetX(ctx, usr.ID).Raw)
}

func NetAddr(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	ip := net.ParseIP("127.0.0.1")
	usr := client.User.Create().SetAddr(schema.Addr{Addr: &net.TCPAddr{IP: ip, Port: 80}}).SaveX(ctx)
	require.Equal(t, "127.0.0.1:80", client.User.GetX(ctx, usr.ID).Addr.String())
	usr.Update().SetAddr(schema.Addr{Addr: &net.UDPAddr{IP: ip, Port: 1812}}).ExecX(ctx)
	require.Equal(t, "127.0.0.1:1812", client.User.GetX(ctx, usr.ID).Addr.String())

	// Ensure sensitive fields are not marshalled.
	f, ok := reflect.TypeOf(ent.User{}).FieldByName("Addr")
	require.True(t, ok)
	require.Equal(t, "-", f.Tag.Get("json"))
}

func Dirs(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	dirs := []http.Dir{"dev", "usr"}
	usr := client.User.Create().SetDirs(dirs).SaveX(ctx)
	require.Equal(t, dirs, usr.Dirs)
	require.Equal(t, dirs, client.User.GetX(ctx, usr.ID).Dirs)

	usr = client.User.Create().SaveX(ctx)
	require.Equal(t, []http.Dir{"/tmp"}, usr.Dirs)
	usr = client.User.GetX(ctx, usr.ID)
	require.Equal(t, []http.Dir{"/tmp"}, usr.Dirs)
}

func URL(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	usr := client.User.Create().SaveX(ctx)
	require.Nil(t, usr.URL, "url field should be nil")
	u, err := url.Parse("https://github.com/a8m")
	require.NoError(t, err)
	usr = client.User.Create().SetURL(u).SaveX(ctx)
	require.Equal(t, u, usr.URL)
	require.Equal(t, u, client.User.GetX(ctx, usr.ID).URL)
}

func URLs(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	u1, err := url.Parse("https://github.com/a8m")
	require.NoError(t, err)
	u2, err := url.Parse("https://github.com/ent")
	require.NoError(t, err)
	usr := client.User.Create().SetURLs([]*url.URL{u1}).SaveX(ctx)
	require.NoError(t, err)
	require.Len(t, usr.URLs, 1)
	require.Equal(t, u1, usr.URLs[0])
	usr = client.User.GetX(ctx, usr.ID)
	require.Equal(t, u1, usr.URLs[0])
	usr = usr.Update().AppendURLs([]*url.URL{u2}).SaveX(ctx)
	require.Len(t, usr.URLs, 2)
	require.Equal(t, u1, usr.URLs[0])
	require.Equal(t, u2, usr.URLs[1])
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

	t.Run("ValueIn", func(t *testing.T) {
		count, err = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueIn(user.FieldURL, []any{"https", "http"}, sqljson.Path("Scheme")))
		}).Count(ctx)
		require.NoError(t, err)
		require.Equal(t, 1, count)

		count, err = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueIn(user.FieldURL, []any{"https", "ftp"}, sqljson.Path("Scheme")))
		}).Count(ctx)
		require.NoError(t, err)
		require.Equal(t, 2, count)

		count, err = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueIn(user.FieldURL, []any{"a", "b"}, sqljson.Path("Scheme")))
		}).Count(ctx)
		require.NoError(t, err)
		require.Zero(t, count)
	})

	t.Run("ValueNotIn", func(t *testing.T) {
		count, err = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueNotIn(user.FieldURL, []any{"https", "http"}, sqljson.Path("Scheme")))
		}).Count(ctx)
		require.NoError(t, err)
		require.Equal(t, 1, count)

		count, err = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueNotIn(user.FieldURL, []any{"https", "ftp"}, sqljson.Path("Scheme")))
		}).Count(ctx)
		require.NoError(t, err)
		require.Equal(t, 0, count)

		count, err = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueNotIn(user.FieldURL, []any{"a", "b"}, sqljson.Path("Scheme")))
		}).Count(ctx)
		require.NoError(t, err)
		require.Equal(t, 2, count)
	})

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

	t.Run("NullLiteral", func(t *testing.T) {
		client.User.Delete().ExecX(ctx)
		users := client.User.CreateBulk(
			client.User.Create().SetURL(u1),
			client.User.Create().SetURL(u2),
		).SaveX(ctx)
		require.Nil(t, users[0].URL.User)
		require.NotNil(t, users[1].URL.User)

		u1 := client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueIsNull(user.FieldURL, sqljson.Path("User")))
		}).OnlyX(ctx)
		require.Equal(t, users[0].ID, u1.ID)

		u2 := client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.ValueIsNotNull(user.FieldURL, sqljson.Path("User")))
		}).OnlyX(ctx)
		require.Equal(t, users[1].ID, u2.ID)

		n := client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.HasKey(user.FieldURL, sqljson.Path("User")))
		}).CountX(ctx)
		require.Equal(t, 2, n, "both u1 and u2 have a 'User' key")
	})

	t.Run("Strings", func(t *testing.T) {
		client.User.Delete().ExecX(ctx)
		u, err := url.Parse("https://github.com/a8m")
		require.NoError(t, err)
		dirs := []http.Dir{"/dev/null"}
		client.User.CreateBulk(
			client.User.Create().SetURL(u),
			client.User.Create().SetDirs(dirs),
			client.User.Create().SetT(&schema.T{S: "foobar", Ls: []string{"foo", "bar"}}),
		).ExecX(ctx)
		require.NoError(t, err)

		ps := []*sql.Predicate{
			sqljson.StringContains(user.FieldDirs, "dev", sqljson.Path("[0]")),
			sqljson.StringHasPrefix(user.FieldDirs, "/dev", sqljson.Path("[0]")),
			sqljson.StringHasSuffix(user.FieldDirs, "/null", sqljson.Path("[0]")),
		}
		for _, p := range ps {
			r = client.User.Query().Where(func(s *sql.Selector) { s.Where(p) }).OnlyX(ctx)
			require.Equal(t, dirs, r.Dirs)
		}
		r = client.User.Query().Where(func(s *sql.Selector) { s.Where(sql.And(ps...)) }).OnlyX(ctx)
		require.Equal(t, dirs, r.Dirs)

		ps = []*sql.Predicate{
			sqljson.StringContains(user.FieldURL, "hub", sqljson.Path("Host")),
			sqljson.StringHasPrefix(user.FieldURL, "github", sqljson.Path("Host")),
			sqljson.StringHasSuffix(user.FieldURL, "hub.com", sqljson.Path("Host")),
		}
		for _, p := range ps {
			r = client.User.Query().Where(func(s *sql.Selector) { s.Where(p) }).OnlyX(ctx)
			require.Equal(t, u, r.URL)
		}

		ps = []*sql.Predicate{
			sqljson.StringHasPrefix(user.FieldT, "foo", sqljson.Path("ls", "[0]")),
			sqljson.StringHasSuffix(user.FieldT, "bar", sqljson.DotPath("ls[1]")),
			sql.And(
				sql.Or(
					sqljson.StringContains(user.FieldT, "foo", sqljson.DotPath("ls[0]")),
					sqljson.StringContains(user.FieldT, "foo", sqljson.DotPath("ls[1]")),
				),
				sql.Or(
					sqljson.StringContains(user.FieldT, "bar", sqljson.DotPath("ls[0]")),
					sqljson.StringContains(user.FieldT, "bar", sqljson.DotPath("ls[1]")),
				),
			),
		}
		for _, p := range ps {
			r = client.User.Query().Where(func(s *sql.Selector) { s.Where(p) }).OnlyX(ctx)
			require.Equal(t, []string{"foo", "bar"}, r.T.Ls)
		}
	})

	t.Run("HasKey", func(t *testing.T) {
		client.User.Delete().ExecX(ctx)
		client.User.CreateBulk(
			client.User.Create(),
			client.User.Create().SetT(&schema.T{}),
			client.User.Create().SetT(&schema.T{M: map[string]any{}}),
			client.User.Create().SetT(&schema.T{M: map[string]any{"a": nil}}),
			client.User.Create().SetT(&schema.T{M: map[string]any{"a": map[string]any{"b": nil, "c": "c"}}}),
		).ExecX(ctx)
		require.NoError(t, err)

		n := client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.HasKey(user.FieldT, sqljson.Path("m")))
		}).CountX(ctx)
		require.Equal(t, 4, n, "take all 'm', including empty and null as omitempty is not set")

		n = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.HasKey(user.FieldT, sqljson.DotPath("m.a")))
		}).CountX(ctx)
		require.Equal(t, 2, n)
		n = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(
				sql.Not(
					sqljson.HasKey(user.FieldT, sqljson.DotPath("m.a")),
				),
			)
		}).CountX(ctx)
		require.Equal(t, 3, n)

		n = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.HasKey(user.FieldT, sqljson.DotPath("m.a.b")))
		}).CountX(ctx)
		require.Equal(t, 1, n)
		n = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(
				sql.Not(
					sqljson.HasKey(user.FieldT, sqljson.DotPath("m.a.b")),
				),
			)
		}).CountX(ctx)
		require.Equal(t, 4, n)

		n = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sqljson.HasKey(user.FieldT, sqljson.DotPath("m.a.c")))
		}).CountX(ctx)
		require.Equal(t, 1, n)
		n = client.User.Query().Where(func(s *sql.Selector) {
			s.Where(
				sql.Not(
					sqljson.HasKey(user.FieldT, sqljson.DotPath("m.a.c")),
				),
			)
		}).CountX(ctx)
		require.Equal(t, 4, n)
	})

	t.Run("Boolean", func(t *testing.T) {
		users := client.User.Query().
			Where(func(s *sql.Selector) {
				s.Where(sqljson.ValueEQ(user.FieldT, true, sqljson.Path("b")))
			}).
			AllX(ctx)
		require.Empty(t, users)
		client.User.Create().SetT(&schema.T{B: true}).ExecX(ctx)
		u1 := client.User.Query().
			Where(func(s *sql.Selector) {
				s.Where(sqljson.ValueEQ(user.FieldT, true, sqljson.Path("b")))
			}).
			OnlyX(ctx)
		require.True(t, u1.T.B)
	})
}

func Order(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	client.User.Delete().ExecX(ctx)
	client.User.CreateBulk(
		client.User.Create().SetT(&schema.T{I: 1, Li: []int{1, 1, 1}}),
		client.User.Create().SetT(&schema.T{I: 2, Li: []int{2, 2}}),
		client.User.Create().SetT(&schema.T{I: 3, Li: []int{3}}),
	).ExecX(ctx)

	users := client.User.Query().
		Order(
			sqljson.OrderValue(user.FieldT, sqljson.Path("i")),
		).
		// PostgreSQL doesn't support ORDER BY
		// expressions with SELECT DISTINCT.
		Unique(false).
		AllX(ctx)
	require.Equal(t, 1, users[0].T.I)
	require.Equal(t, 2, users[1].T.I)
	require.Equal(t, 3, users[2].T.I)

	users = client.User.Query().
		Order(
			sqljson.OrderValueDesc(user.FieldT, sqljson.Path("i")),
		).
		Unique(false).
		AllX(ctx)
	require.Equal(t, 3, users[0].T.I)
	require.Equal(t, 2, users[1].T.I)
	require.Equal(t, 1, users[2].T.I)

	// Order by array length.
	users = client.User.Query().
		Order(
			sqljson.OrderLenDesc(user.FieldT, sqljson.Path("li")),
		).
		Unique(false).
		AllX(ctx)
	require.Len(t, users[0].T.Li, 3)
	require.Len(t, users[1].T.Li, 2)
	require.Len(t, users[2].T.Li, 1)
}

func Scan(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	all := client.User.Query().Order(ent.Asc(user.FieldID)).AllX(ctx)
	require.NotEmpty(t, all)
	var scanned []*ent.User
	// Select all non-sensitive fields.
	client.User.Query().Order(ent.Asc(user.FieldID)).Select(user.Columns[:len(user.Columns)-2]...).ScanX(ctx, &scanned)
	require.Equal(t, len(all), len(scanned))
	for i := range all {
		require.Equal(t, all[i].ID, scanned[i].ID)
		require.Equal(t, all[i].T, scanned[i].T)
		require.Equal(t, all[i].URL, scanned[i].URL)
		require.Equal(t, all[i].URLs, scanned[i].URLs)
		require.Equal(t, all[i].Dirs, scanned[i].Dirs)
		require.Equal(t, all[i].Raw, scanned[i].Raw)
		require.Equal(t, all[i].Ints, scanned[i].Ints)
		require.Equal(t, all[i].Floats, scanned[i].Floats)
		require.Equal(t, all[i].Strings, scanned[i].Strings)
	}
}
