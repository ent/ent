// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/facebook/ent/dialect/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrateHookOmitTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	tables := []*Table{
		{Name: "users"},
		{Name: "pets"},
	}

	myMock := mysqlMock{mock}
	myMock.start("5.7.23")

	myMock.tableExists("pets", false)
	myMock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`() CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
		WillReturnResult(sqlmock.NewResult(0, 1))

	myMock.ExpectCommit()

	migrate, err := NewMigrate(sql.OpenDB("mysql", db), WithHook(func(next Creator) Creator {
		return CreateFunc(func(ctx context.Context, tables ...*Table) error {
			return next.Create(ctx, tables[1])
		})
	}))
	require.NoError(t, err)
	err = migrate.Create(context.Background(), tables...)
	require.NoError(t, err)
}

func TestMigrateHookAddTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	tables := []*Table{
		{Name: "users"},
		{Name: "pets"},
	}

	myMock := mysqlMock{mock}
	myMock.start("5.7.23")

	myMock.tableExists("users", false)
	myMock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`() CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
		WillReturnResult(sqlmock.NewResult(0, 1))

	myMock.tableExists("pets", false)
	myMock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`() CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
		WillReturnResult(sqlmock.NewResult(0, 1))

	myMock.ExpectCommit()

	migrate, err := NewMigrate(sql.OpenDB("mysql", db), WithHook(func(next Creator) Creator {
		return CreateFunc(func(ctx context.Context, tables ...*Table) error {
			return next.Create(ctx, tables[0], &Table{Name: "pets"})
		})
	}))
	require.NoError(t, err)
	err = migrate.Create(context.Background(), tables...)
	require.NoError(t, err)
}
