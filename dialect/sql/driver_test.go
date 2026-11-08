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
	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = E'bar'`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = DEFAULT`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
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

	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = E'bar'`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = E'baz'`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = DEFAULT`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
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
	mock.ExpectExec(regexp.QuoteMeta(`SET LOCAL "foo" = E'bar'`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
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

	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = E'qux'`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("INSERT INTO users DEFAULT VALUES").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(regexp.QuoteMeta(`SET "foo" = DEFAULT`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
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
		name, key, sqlName string
		open               func(*sql.DB) *Driver
		transaction        bool
	}{
		{"postgres", "app.user", `"app"."user"`, func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, false},
		{"mysql", "@user", "", func(db *sql.DB) *Driver { return OpenDB(dialect.MySQL, db) }, false},
		{"postgres-driver", "app.user", `"app"."user"`, func(db *sql.DB) *Driver { return NewDriver(dialect.Postgres, Conn{ExecQuerier: db}) }, false},
		{"postgres-wrapped-driver", "app.user", `"app"."user"`, func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, false},
		{"postgres-wrapped-transaction", "app.user", `"app"."user"`, func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, true},
		{"quoted-name", ` "app" . "User" `, `"app"."User"`, func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, false},
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
				command := "SET "
				if tt.transaction {
					command = "SET LOCAL "
				}
				mock.ExpectExec(regexp.QuoteMeta(command + tt.sqlName + ` = E'O''Reilly\\team; SELECT 1 --'`)).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			} else {
				mock.ExpectExec(regexp.QuoteMeta("SET @user = ?")).WithArgs(value).WillReturnResult(sqlmock.NewResult(0, 0))
			}
			mock.ExpectExec("SELECT 1").WillReturnResult(sqlmock.NewResult(0, 0))
			switch {
			case tx != nil:
				mock.ExpectCommit()
			case drv.Dialect() == dialect.Postgres:
				mock.ExpectExec(regexp.QuoteMeta("SET " + tt.sqlName + " = DEFAULT")).WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))
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

func TestPostgresVarName(t *testing.T) {
	for _, tt := range []struct{ name, want string }{
		{"app.user", `"app"."user"`},
		{"APP.USER", `"app"."user"`},
		{` "app" . "User" `, `"app"."User"`},
		{`"app"."us""er"`, `"app"."us""er"`},
		{`"app.user"`, `"app.user"`},
		{`"app"."some user"`, `"app"."some user"`},
		{" \f APP \n.\t USER \r", `"app"."user"`},
		{"ÄPP.USER", `"Äpp"."user"`},
		{"App_1.User$2", `"app_1"."user$2"`},
		{"app." + strings.Repeat("x", 100), `"app"."` + strings.Repeat("x", 100) + `"`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			name, err := postgresVarName(tt.name)
			require.NoError(t, err)
			require.Equal(t, tt.want, name)
		})
	}
	for _, name := range []string{"", " ", ".app", "app.", "app..user", "1app.user", "app.user name", `""`, `"app`, `"app"user`, "app.user; SELECT 1", "app.user--comment", "app.\vuser", "app.\x00user", "\"app\x00\".user"} {
		t.Run("invalid/"+name, func(t *testing.T) {
			_, err := postgresVarName(name)
			require.ErrorContains(t, err, "invalid PostgreSQL variable name")
		})
	}
}
