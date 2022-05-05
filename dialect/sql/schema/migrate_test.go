// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestMigrateHookOmitTable(t *testing.T) {
	db, mk, err := sqlmock.New()
	require.NoError(t, err)

	tables := []*Table{{Name: "users"}, {Name: "pets"}}
	mock := mysqlMock{mk}
	mock.start("5.7.23")
	mock.tableExists("pets", false)
	mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`() CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	m, err := NewMigrate(sql.OpenDB("mysql", db), WithHooks(func(next Creator) Creator {
		return CreateFunc(func(ctx context.Context, tables ...*Table) error {
			return next.Create(ctx, tables[1])
		})
	}))
	require.NoError(t, err)
	err = m.Create(context.Background(), tables...)
	require.NoError(t, err)
}

func TestMigrateHookAddTable(t *testing.T) {
	db, mk, err := sqlmock.New()
	require.NoError(t, err)

	tables := []*Table{{Name: "users"}}
	mock := mysqlMock{mk}
	mock.start("5.7.23")
	mock.tableExists("users", false)
	mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`() CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.tableExists("pets", false)
	mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`() CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	m, err := NewMigrate(sql.OpenDB("mysql", db), WithHooks(func(next Creator) Creator {
		return CreateFunc(func(ctx context.Context, tables ...*Table) error {
			return next.Create(ctx, tables[0], &Table{Name: "pets"})
		})
	}))
	require.NoError(t, err)
	err = m.Create(context.Background(), tables...)
	require.NoError(t, err)
}

func TestMigrate_Diff(t *testing.T) {
	db, err := sql.Open(dialect.SQLite, "file:test?mode=memory&_fk=1")
	require.NoError(t, err)

	p := t.TempDir()
	d, err := migrate.NewLocalDir(p)
	require.NoError(t, err)

	m, err := NewMigrate(db, WithDir(d))
	require.NoError(t, err)
	require.NoError(t, m.Diff(context.Background(), &Table{Name: "users"}))
	v := time.Now().Format("20060102150405")
	requireFileEqual(t, filepath.Join(p, v+"_changes.up.sql"), "-- create \"users\" table\nCREATE TABLE `users` (, PRIMARY KEY ());\n")
	requireFileEqual(t, filepath.Join(p, v+"_changes.down.sql"), "-- reverse: create \"users\" table\nDROP TABLE `users`;\n")
	require.NoFileExists(t, filepath.Join(p, "atlas.sum"))

	// Test integrity file.
	p = t.TempDir()
	d, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	m, err = NewMigrate(db, WithDir(d), WithSumFile())
	require.NoError(t, err)
	require.NoError(t, m.Diff(context.Background(), &Table{Name: "users"}))
	requireFileEqual(t, filepath.Join(p, v+"_changes.up.sql"), "-- create \"users\" table\nCREATE TABLE `users` (, PRIMARY KEY ());\n")
	requireFileEqual(t, filepath.Join(p, v+"_changes.down.sql"), "-- reverse: create \"users\" table\nDROP TABLE `users`;\n")
	require.FileExists(t, filepath.Join(p, "atlas.sum"))
	require.NoError(t, d.WriteFile("tmp.sql", nil))
	require.ErrorIs(t, m.Diff(context.Background(), &Table{Name: "users"}), migrate.ErrChecksumMismatch)
}

func requireFileEqual(t *testing.T, name, contents string) {
	c, err := os.ReadFile(name)
	require.NoError(t, err)
	require.Equal(t, contents, string(c))
}

func TestMigrateWithoutForeignKeys(t *testing.T) {
	tbl := &schema.Table{
		Name: "tbl",
		Columns: []*schema.Column{
			{Name: "id", Type: &schema.ColumnType{Type: &schema.IntegerType{T: "bigint"}}},
		},
	}
	fk := &schema.ForeignKey{
		Symbol:     "fk",
		Table:      tbl,
		Columns:    tbl.Columns[1:],
		RefTable:   tbl,
		RefColumns: tbl.Columns[:1],
		OnUpdate:   schema.NoAction,
		OnDelete:   schema.Cascade,
	}
	tbl.ForeignKeys = append(tbl.ForeignKeys, fk)
	t.Run("AddTable", func(t *testing.T) {
		mdiff := DiffFunc(func(_, _ *schema.Schema) ([]schema.Change, error) {
			return []schema.Change{
				&schema.AddTable{
					T: tbl,
				},
			}, nil
		})
		df, err := withoutForeignKeys(mdiff).Diff(nil, nil)
		require.NoError(t, err)
		require.Len(t, df, 1)
		actual, ok := df[0].(*schema.AddTable)
		require.True(t, ok)
		require.Nil(t, actual.T.ForeignKeys)
	})
	t.Run("ModifyTable", func(t *testing.T) {
		mdiff := DiffFunc(func(_, _ *schema.Schema) ([]schema.Change, error) {
			return []schema.Change{
				&schema.ModifyTable{
					T: tbl,
					Changes: []schema.Change{
						&schema.AddIndex{
							I: &schema.Index{
								Name: "id_key",
								Parts: []*schema.IndexPart{
									{C: tbl.Columns[0]},
								},
							},
						},
						&schema.DropForeignKey{
							F: fk,
						},
						&schema.AddForeignKey{
							F: fk,
						},
						&schema.ModifyForeignKey{
							From:   fk,
							To:     fk,
							Change: schema.ChangeRefColumn,
						},
						&schema.AddColumn{
							C: &schema.Column{Name: "name", Type: &schema.ColumnType{Type: &schema.StringType{T: "varchar(255)"}}},
						},
					},
				},
			}, nil
		})
		df, err := withoutForeignKeys(mdiff).Diff(nil, nil)
		require.NoError(t, err)
		require.Len(t, df, 1)
		actual, ok := df[0].(*schema.ModifyTable)
		require.True(t, ok)
		require.Len(t, actual.Changes, 2)
		addIndex, ok := actual.Changes[0].(*schema.AddIndex)
		require.True(t, ok)
		require.EqualValues(t, "id_key", addIndex.I.Name)
		addColumn, ok := actual.Changes[1].(*schema.AddColumn)
		require.True(t, ok)
		require.EqualValues(t, "name", addColumn.C.Name)
	})
}
