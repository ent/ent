// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"strings"
	"testing"

	"entgo.io/ent/dialect"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestWithVars(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	db.SetMaxOpenConns(1)
	drv := OpenDB(dialect.Postgres, db)
	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, $2, $3)")).WithArgs("foo", "bar", false).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectExec("RESET foo").WillReturnResult(sqlmock.NewResult(0, 0))
	rows := &Rows{}
	err = drv.Query(
		WithVar(context.Background(), "foo", "bar"),
		"SELECT 1",
		[]any{},
		rows,
	)
	require.NoError(t, err)
	require.NoError(t, rows.Close(), "rows should be closed to release the connection")
	require.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, $2, $3)")).WithArgs("foo", "bar", false).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, $2, $3)")).WithArgs("foo", "baz", false).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectExec("RESET foo").WillReturnResult(sqlmock.NewResult(0, 0))
	err = drv.Query(
		WithVar(WithVar(context.Background(), "foo", "bar"), "foo", "baz"),
		"SELECT 1",
		[]any{},
		rows,
	)
	require.NoError(t, err)
	require.NoError(t, rows.Close(), "rows should be closed to release the connection")
	require.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, $2, $3)")).WithArgs("foo", "bar", true).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectCommit()
	tx, err := drv.Tx(context.Background())
	require.NoError(t, err)
	err = tx.Query(
		WithVar(context.Background(), "foo", "bar"),
		"SELECT 1",
		[]any{},
		rows,
	)
	require.NoError(t, err)
	require.NoError(t, tx.Commit())
	require.NoError(t, mock.ExpectationsWereMet())
	// Rows should not be closed to release the session,
	// as a transaction is always scoped to a single connection.

	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, $2, $3)")).WithArgs("foo", "qux", false).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("INSERT INTO users DEFAULT VALUES").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("RESET foo").WillReturnResult(sqlmock.NewResult(0, 0))
	err = drv.Exec(
		WithVar(context.Background(), "foo", "qux"),
		"INSERT INTO users DEFAULT VALUES",
		[]any{},
		nil,
	)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
	// No rows are returned, so no need to close them.
}

func TestWithVarValues(t *testing.T) {
	const value = "O'Reilly\\team; SELECT 1 --"
	tests := []struct {
		name        string
		key         string
		open        func(*sql.DB) *Driver
		query       string
		args        []driver.Value
		reset       string
		transaction bool
	}{
		{"postgres", "app.user", func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, "SELECT set_config($1, $2, $3)", []driver.Value{"app.user", value, false}, "RESET app.user", false},
		{"mysql", "@user", func(db *sql.DB) *Driver { return OpenDB(dialect.MySQL, db) }, "SET @user = ?", []driver.Value{value}, "SET @user = NULL", false},
		{"postgres-driver", "app.user", func(db *sql.DB) *Driver { return NewDriver(dialect.Postgres, Conn{ExecQuerier: db}) }, "SELECT set_config($1, $2, $3)", []driver.Value{"app.user", value, false}, "RESET app.user", false},
		{"postgres-wrapped-driver", "app.user", func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, "SELECT set_config($1, $2, $3)", []driver.Value{"app.user", value, false}, "RESET app.user", false},
		{"postgres-wrapped-transaction", "app.user", func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, "SELECT set_config($1, $2, $3)", []driver.Value{"app.user", value, true}, "", true},
		{"quoted-name", ` "app" . "User" `, func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, "SELECT set_config($1, $2, $3)", []driver.Value{"app.User", value, false}, `RESET  "app" . "User" `, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			t.Cleanup(func() { db.Close() })
			drv := tt.open(db)
			var ex dialect.ExecQuerier = drv
			finish := func() {}
			if tt.transaction {
				mock.ExpectBegin()
				tx, err := drv.Tx(context.Background())
				require.NoError(t, err)
				ex = tx
				finish = func() { require.NoError(t, tx.Commit()) }
			}
			setting := mock.ExpectExec(regexp.QuoteMeta(tt.query))
			setting.WithArgs(tt.args...)
			setting.WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(0, 0))
			if tt.transaction {
				mock.ExpectCommit()
			} else {
				mock.ExpectExec(regexp.QuoteMeta(tt.reset)).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			require.NoError(t, ex.Exec(WithVar(context.Background(), tt.key, value), "SELECT 1", []any{}, nil))
			finish()
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestPostgresVar(t *testing.T) {
	for _, tt := range []struct {
		name, key, value, wantKey, wantValue string
	}{
		{"plain", "app.user", "O'Reilly", "app.user", "O'Reilly"},
		{"unquoted-case", "APP.USER", "value", "app.user", "value"},
		{"quoted-case", `"app"."User"`, "value", "app.User", "value"},
		{"escaped-quote", `"app"."us""er"`, "value", `app.us"er`, "value"},
		{"quoted-dot", `"app.user"`, "value", "app.user", "value"},
		{"quoted-space", `"app"."some user"`, "value", "app.some user", "value"},
		{"whitespace", " \f APP \n.\t USER \r", "value", "app.user", "value"},
		{"unicode", "ÄPP.USER", "value", "Äpp.user", "value"},
		{"identifier-suffix", "App_1.User$2", "value", "app_1.user$2", "value"},
		{"long-name", "app." + strings.Repeat("x", 100), "value", "app." + strings.Repeat("x", 100), "value"},
		{"mixed-schema", "search_path", "TenantA", "search_path", `"TenantA"`},
		{"comma-schema", "search_path", "a,b", "search_path", `"a,b"`},
		{"quoted-schema", `"SEARCH_PATH"`, `a"b`, "SEARCH_PATH", `"a""b"`},
		{"empty-schema", "search_path", "", "search_path", `""`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			key, value, err := postgresVar(tt.key, tt.value)
			require.NoError(t, err)
			require.Equal(t, tt.wantKey, key)
			require.Equal(t, tt.wantValue, value)
		})
	}
	for _, name := range []string{
		"temp_tablespaces", "session_preload_libraries", "shared_preload_libraries",
		"local_preload_libraries", "unix_socket_directories", "output_plugin_libraries",
	} {
		t.Run(name, func(t *testing.T) {
			key, value, err := postgresVar(name, `One,Two`)
			require.NoError(t, err)
			require.Equal(t, name, key)
			require.Equal(t, `"One,Two"`, value)
		})
	}
	for _, name := range []string{
		"", " ", ".app", "app.", "app..user", "1app.user", "app.user name",
		`""`, `"app`, `"app"user`, "app.user; SELECT 1", "app.user--comment",
		"app.\vuser", "app.\x00user", "\"app\x00\".user",
	} {
		t.Run("invalid/"+name, func(t *testing.T) {
			_, _, err := postgresVar(name, "value")
			require.ErrorContains(t, err, "invalid PostgreSQL variable name")
		})
	}
}
