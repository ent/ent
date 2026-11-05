// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
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
	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, NULL, false)")).WithArgs("foo").WillReturnResult(sqlmock.NewResult(0, 0))
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
	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, NULL, false)")).WithArgs("foo").WillReturnResult(sqlmock.NewResult(0, 0))
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
	mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, NULL, false)")).WithArgs("foo").WillReturnResult(sqlmock.NewResult(0, 0))
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
	for _, tt := range []struct {
		name, key, boundKey string
		open                func(*sql.DB) *Driver
		transaction         bool
	}{
		{"postgres", "app.user", "app.user", func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, false},
		{"mysql", "@user", "@user", func(db *sql.DB) *Driver { return OpenDB(dialect.MySQL, db) }, false},
		{"postgres-driver", "app.user", "app.user", func(db *sql.DB) *Driver { return NewDriver(dialect.Postgres, Conn{ExecQuerier: db}) }, false},
		{"postgres-wrapped-driver", "app.user", "app.user", func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, false},
		{"postgres-wrapped-transaction", "app.user", "app.user", func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, true},
		{"quoted-name", ` "app" . "User" `, "app.User", func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			t.Cleanup(func() { db.Close() })
			drv := tt.open(db)
			var ex dialect.ExecQuerier = drv
			var tx dialect.Tx
			if tt.transaction {
				mock.ExpectBegin()
				tx, err = drv.Tx(context.Background())
				require.NoError(t, err)
				ex = tx
			}
			if drv.Dialect() == dialect.Postgres {
				mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, $2, $3)")).WithArgs(tt.boundKey, value, tt.transaction).WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectExec(regexp.QuoteMeta("SET @user = ?")).WithArgs(value).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(0, 0))
			switch {
			case tx != nil:
				mock.ExpectCommit()
			case drv.Dialect() == dialect.Postgres:
				mock.ExpectExec(regexp.QuoteMeta("SELECT set_config($1, NULL, false)")).WithArgs(tt.boundKey).WillReturnResult(sqlmock.NewResult(0, 0))
			default:
				mock.ExpectExec(regexp.QuoteMeta("SET @user = NULL")).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			require.NoError(t, ex.Exec(WithVar(context.Background(), tt.key, value), "SELECT 1", []any{}, nil))
			if tx != nil {
				require.NoError(t, tx.Commit())
			}
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

func TestPostgresTxVar(t *testing.T) {
	for _, tt := range []struct {
		name, value, want string
		invalid           bool
	}{
		{"transaction_isolation", "SERIALIZABLE", "SET LOCAL transaction_isolation = 'serializable'", false},
		{"transaction_isolation", "repeatable read", "SET LOCAL transaction_isolation = 'repeatable read'", false},
		{"transaction_read_only", "YES", "SET LOCAL transaction_read_only = 'on'", false},
		{"transaction_read_only", "0", "SET LOCAL transaction_read_only = 'off'", false},
		{"transaction_deferrable", "of", "SET LOCAL transaction_deferrable = 'off'", false},
		{"transaction_deferrable", "1", "SET LOCAL transaction_deferrable = 'on'", false},
		{"app.user", "alice", "", false},
		{"transaction_isolation", "serializable'; SELECT 1 --", "", true},
		{"transaction_read_only", "on'; SELECT 1 --", "", true},
		{"transaction_deferrable", "o", "", true},
		{"transaction_deferrable", "", "", true},
	} {
		t.Run(tt.name+"/"+tt.value, func(t *testing.T) {
			query, err := postgresTxVar(tt.name, tt.value)
			if tt.invalid {
				require.Error(t, err)
				require.Empty(t, query)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, query)
		})
	}
}
