// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package json

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/entc/integration/json/ent"
	"github.com/facebook/ent/entc/integration/json/ent/migrate"
	"github.com/facebook/ent/entc/integration/json/ent/user"

	"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			cfg := mysql.Config{
				User: "root", Passwd: "pass", Net: "tcp", Addr: addr,
				AllowNativePasswords: true, ParseTime: true,
			}
			db, err := sql.Open("mysql", cfg.FormatDSN())
			require.NoError(t, err)
			defer db.Close()
			_, err = db.Exec("CREATE DATABASE IF NOT EXISTS json")
			require.NoError(t, err, "creating database")
			defer db.Exec("DROP DATABASE IF EXISTS json")

			cfg.DBName = "json"
			client, err := ent.Open("mysql", cfg.FormatDSN())
			require.NoError(t, err, "connecting to json database")
			err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
			require.NoError(t, err)

			URL(t, client)
			Dirs(t, client)
			Ints(t, client)
			Floats(t, client)
			Strings(t, client)
			RawMessage(t, client)
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
			_, err = db.Exec("CREATE DATABASE json")
			require.NoError(t, err, "creating database")
			defer db.Exec("DROP DATABASE json")

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
		})
	}
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
