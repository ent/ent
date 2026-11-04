// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
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
	mock.ExpectExec(regexp.QuoteMeta(postgresVarQuery)).WithArgs("foo", "bar", false).WillReturnResult(sqlmock.NewResult(0, 0))
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

	mock.ExpectExec(regexp.QuoteMeta(postgresVarQuery)).WithArgs("foo", "bar", false).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(regexp.QuoteMeta(postgresVarQuery)).WithArgs("foo", "baz", false).WillReturnResult(sqlmock.NewResult(0, 0))
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
	mock.ExpectExec(regexp.QuoteMeta(postgresVarQuery)).WithArgs("foo", "bar", true).WillReturnResult(sqlmock.NewResult(0, 0))
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

	mock.ExpectExec(regexp.QuoteMeta(postgresVarQuery)).WithArgs("foo", "qux", false).WillReturnResult(sqlmock.NewResult(0, 0))
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
		open        func(*sql.DB) *Driver
		query       string
		args        []driver.Value
		reset       string
		transaction bool
	}{
		{"postgres", func(db *sql.DB) *Driver { return OpenDB(dialect.Postgres, db) }, postgresVarQuery, []driver.Value{"app.user", value, false}, "RESET app.user", false},
		{"mysql", func(db *sql.DB) *Driver { return OpenDB(dialect.MySQL, db) }, "SET @user = ?", []driver.Value{value}, "SET @user = NULL", false},
		{"postgres-driver", func(db *sql.DB) *Driver { return NewDriver(dialect.Postgres, Conn{ExecQuerier: db}) }, postgresVarQuery, []driver.Value{"app.user", value, false}, "RESET app.user", false},
		{"postgres-wrapped-driver", func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, postgresVarQuery, []driver.Value{"app.user", value, false}, "RESET app.user", false},
		{"postgres-wrapped-transaction", func(db *sql.DB) *Driver { return NewDriver("postgres-traced", Conn{ExecQuerier: db}) }, postgresVarQuery, []driver.Value{"app.user", value, true}, "", true},
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
			name := "app.user"
			if drv.Dialect() == dialect.MySQL {
				name = "@user"
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
			require.NoError(t, ex.Exec(WithVar(context.Background(), name, value), "SELECT 1", []any{}, nil))
			finish()
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
