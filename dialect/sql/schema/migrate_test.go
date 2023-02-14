// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"time"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
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
	}), WithAtlas(false))
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
	}), WithAtlas(false))
	require.NoError(t, err)
	err = m.Create(context.Background(), tables...)
	require.NoError(t, err)
}

func TestMigrate_Formatter(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)

	// If no formatter is given it will be set according to the given migration directory implementation.
	for _, tt := range []struct {
		dir migrate.Dir
		fmt migrate.Formatter
	}{
		{&migrate.LocalDir{}, sqltool.GolangMigrateFormatter},
		{&sqltool.GolangMigrateDir{}, sqltool.GolangMigrateFormatter},
		{&sqltool.GooseDir{}, sqltool.GooseFormatter},
		{&sqltool.DBMateDir{}, sqltool.DBMateFormatter},
		{&sqltool.FlywayDir{}, sqltool.FlywayFormatter},
		{&sqltool.LiquibaseDir{}, sqltool.LiquibaseFormatter},
		{struct{ migrate.Dir }{}, sqltool.GolangMigrateFormatter}, // default one if migration dir is unknown
	} {
		m, err := NewMigrate(sql.OpenDB("", db), WithDir(tt.dir))
		require.NoError(t, err)
		require.Equal(t, tt.fmt, m.fmt)
	}

	// If a formatter is given, it is not overridden.
	m, err := NewMigrate(sql.OpenDB("", db), WithDir(&migrate.LocalDir{}), WithFormatter(migrate.DefaultFormatter))
	require.NoError(t, err)
	require.Equal(t, migrate.DefaultFormatter, m.fmt)
}

func TestMigrate_DiffJoinTableAllocationBC(t *testing.T) {
	// Due to a bug in previous versions, if the universal ID option was enabled and the schema did contain an M2M
	// relation, the join table would have had an entry for the join table in the types table. This test ensures,
	// that the PK range allocated for the join table stays in place, since it's removal would break existing projects
	// due to shifted ranges.

	db, err := sql.Open(dialect.SQLite, "file:test?mode=memory&_fk=1")
	require.NoError(t, err)

	// Mock an existing database with an allocation for a join table.
	for _, stmt := range []string{
		"CREATE TABLE `groups` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);",
		"CREATE INDEX `short` ON `groups` (`id`);",
		"CREATE INDEX `long____________________________1cb2e7e47a309191385af4ad320875b1` ON `groups` (`id`);",
		"CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);",
		"INSERT INTO sqlite_sequence (name, seq) VALUES (\"users\", 4294967296);",
		"CREATE TABLE `user_groups` (`user_id` integer NOT NULL, `group_id` integer NOT NULL, PRIMARY KEY (`user_id`, `group_id`), CONSTRAINT `user_groups_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_groups_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE);",
		"INSERT INTO sqlite_sequence (name, seq) VALUES (\"user_groups\", 8589934592);",
		"CREATE TABLE `ent_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `type` text NOT NULL);",
		"CREATE UNIQUE INDEX `ent_types_type_key` ON `ent_types` (`type`);",
		"INSERT INTO `ent_types` (`type`) VALUES ('groups'), ('users'), ('user_groups');",
		"INSERT INTO `groups` (`name`) VALUES ('seniors'), ('juniors')",
		"INSERT INTO `users` (`name`) VALUES ('masseelch'), ('a8m'), ('rotemtam')",
		"INSERT INTO `user_groups` (`user_id`, `group_id`) VALUES (4294967297, 1), (4294967298, 1), (4294967299, 2)",
	} {
		_, err := db.ExecContext(context.Background(), stmt)
		require.NoError(t, err)
	}

	// Expect to have no changes when migration runs with fix.
	m, err := NewMigrate(db, WithGlobalUniqueID(true), WithDiffHook(func(next Differ) Differ {
		return DiffFunc(func(current, desired *schema.Schema) ([]schema.Change, error) {
			changes, err := next.Diff(current, desired)
			if err != nil {
				return nil, err
			}
			require.Len(t, changes, 0)
			return changes, nil
		})
	}))
	require.NoError(t, err)
	require.NoError(t, m.Create(context.Background(), tables...))

	// Expect to have no changes to the allocation when the join table is dropped.
	m, err = NewMigrate(db, WithGlobalUniqueID(true))
	require.NoError(t, err)
	require.NoError(t, m.Create(context.Background(), groupsTable, usersTable))

	rows, err := db.QueryContext(context.Background(), "SELECT `type` from `ent_types` ORDER BY `id` ASC")
	require.NoError(t, err)
	var types []string
	for rows.Next() {
		var typ string
		require.NoError(t, rows.Scan(&typ))
		types = append(types, typ)
	}
	require.NoError(t, rows.Err())
	require.Equal(t, []string{"groups", "users", "user_groups"}, types)
}

var (
	groupsColumns = []*Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	groupsTable = &Table{
		Name:       "groups",
		Columns:    groupsColumns,
		PrimaryKey: []*Column{groupsColumns[0]},
		Indexes: []*Index{
			{
				Name:    "short",
				Columns: []*Column{groupsColumns[0]}},
			{
				Name:    "long_" + strings.Repeat("_", 60),
				Columns: []*Column{groupsColumns[0]},
			},
		},
	}
	usersColumns = []*Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	usersTable = &Table{
		Name:       "users",
		Columns:    usersColumns,
		PrimaryKey: []*Column{usersColumns[0]},
	}
	userGroupsColumns = []*Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	userGroupsTable = &Table{
		Name:       "user_groups",
		Columns:    userGroupsColumns,
		PrimaryKey: []*Column{userGroupsColumns[0], userGroupsColumns[1]},
		ForeignKeys: []*ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*Column{userGroupsColumns[0]},
				RefColumns: []*Column{usersColumns[0]},
				OnDelete:   Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*Column{userGroupsColumns[1]},
				RefColumns: []*Column{groupsColumns[0]},
				OnDelete:   Cascade,
			},
		},
	}
	tables = []*Table{
		groupsTable,
		usersTable,
		userGroupsTable,
	}
	petColumns = []*Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	petsTable = &Table{
		Name:       "pets",
		Columns:    petColumns,
		PrimaryKey: petColumns,
	}
)

func init() {
	userGroupsTable.ForeignKeys[0].RefTable = usersTable
	userGroupsTable.ForeignKeys[1].RefTable = groupsTable
}

func TestMigrate_Diff(t *testing.T) {
	ctx := context.Background()

	db, err := sql.Open(dialect.SQLite, "file:test?mode=memory&_fk=1")
	require.NoError(t, err)

	p := t.TempDir()
	d, err := migrate.NewLocalDir(p)
	require.NoError(t, err)

	m, err := NewMigrate(db, WithDir(d))
	require.NoError(t, err)
	require.NoError(t, m.Diff(ctx, &Table{Name: "users"}))
	v := time.Now().UTC().Format("20060102150405")
	requireFileEqual(t, filepath.Join(p, v+"_changes.up.sql"), "-- create \"users\" table\nCREATE TABLE `users` ();\n")
	requireFileEqual(t, filepath.Join(p, v+"_changes.down.sql"), "-- reverse: create \"users\" table\nDROP TABLE `users`;\n")
	require.FileExists(t, filepath.Join(p, migrate.HashFileName))

	// Test integrity file.
	p = t.TempDir()
	d, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	m, err = NewMigrate(db, WithDir(d))
	require.NoError(t, err)
	require.NoError(t, m.Diff(ctx, &Table{Name: "users"}))
	requireFileEqual(t, filepath.Join(p, v+"_changes.up.sql"), "-- create \"users\" table\nCREATE TABLE `users` ();\n")
	requireFileEqual(t, filepath.Join(p, v+"_changes.down.sql"), "-- reverse: create \"users\" table\nDROP TABLE `users`;\n")
	require.FileExists(t, filepath.Join(p, migrate.HashFileName))
	require.NoError(t, d.WriteFile("tmp.sql", nil))
	require.ErrorIs(t, m.Diff(ctx, &Table{Name: "users"}), migrate.ErrChecksumMismatch)

	p = t.TempDir()
	d, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	f, err := migrate.NewTemplateFormatter(
		template.Must(template.New("").Parse("{{ .Name }}.sql")),
		template.Must(template.New("").Parse(
			`{{ range .Changes }}{{ printf "%s;\n" .Cmd }}{{ end }}`,
		)),
	)
	require.NoError(t, err)

	// Join tables (mapping between user and group) will not result in an entry to the types table.
	m, err = NewMigrate(db, WithFormatter(f), WithDir(d), WithGlobalUniqueID(true))
	require.NoError(t, err)
	require.NoError(t, m.Diff(ctx, tables...))
	changesSQL := strings.Join([]string{
		"CREATE TABLE `groups` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);",
		"CREATE INDEX `short` ON `groups` (`id`);",
		"CREATE INDEX `long____________________________1cb2e7e47a309191385af4ad320875b1` ON `groups` (`id`);",
		"CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);",
		fmt.Sprintf("INSERT INTO sqlite_sequence (name, seq) VALUES (\"users\", %d);", 1<<32),
		"CREATE TABLE `user_groups` (`user_id` integer NOT NULL, `group_id` integer NOT NULL, PRIMARY KEY (`user_id`, `group_id`), CONSTRAINT `user_groups_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_groups_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE);",
		"CREATE TABLE `ent_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `type` text NOT NULL);",
		"CREATE UNIQUE INDEX `ent_types_type_key` ON `ent_types` (`type`);",
		"INSERT INTO `ent_types` (`type`) VALUES ('groups'), ('users');",
		"",
	}, "\n")
	requireFileEqual(t, filepath.Join(p, "changes.sql"), changesSQL)

	// Adding another node will result in a new entry to the TypeTable (without actually creating it).
	_, err = db.ExecContext(ctx, changesSQL, nil, nil)
	require.NoError(t, err)
	require.NoError(t, m.NamedDiff(ctx, "changes_2", petsTable))
	requireFileEqual(t,
		filepath.Join(p, "changes_2.sql"), strings.Join([]string{
			"CREATE TABLE `pets` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);",
			fmt.Sprintf("INSERT INTO sqlite_sequence (name, seq) VALUES (\"pets\", %d);", 2<<32),
			"INSERT INTO `ent_types` (`type`) VALUES ('pets');", "",
		}, "\n"))

	// Checksum will be updated as well.
	require.NoError(t, migrate.Validate(d))
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

func TestAtlas_StateReader(t *testing.T) {
	db, err := sql.Open(dialect.SQLite, "file:test?mode=memory&_fk=1")
	require.NoError(t, err)
	m, err := NewMigrate(db)
	require.NoError(t, err)
	realm, err := m.StateReader(&Table{
		Name: "users",
		Columns: []*Column{
			{Name: "name", Type: field.TypeString},
			{Name: "active", Type: field.TypeBool},
		},
	}).ReadState(context.Background())
	require.NoError(t, err)
	require.NotNil(t, realm)
	require.Len(t, realm.Schemas, 1)
	require.Len(t, realm.Schemas[0].Tables, 1)
	require.Equal(t, realm.Schemas[0].Tables[0].Name, "users")
	require.Equal(t,
		realm.Schemas[0].Tables[0].Columns,
		[]*schema.Column{
			schema.NewStringColumn("name", "text"),
			schema.NewBoolColumn("active", "bool"),
		},
	)
}
