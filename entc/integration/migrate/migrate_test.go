// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package migrate

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/migrate/entv1"
	migratev1 "entgo.io/ent/entc/integration/migrate/entv1/migrate"
	userv1 "entgo.io/ent/entc/integration/migrate/entv1/user"
	"entgo.io/ent/entc/integration/migrate/entv2"
	"entgo.io/ent/entc/integration/migrate/entv2/conversion"
	"entgo.io/ent/entc/integration/migrate/entv2/customtype"
	migratev2 "entgo.io/ent/entc/integration/migrate/entv2/migrate"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/entc/integration/migrate/entv2/user"
	"entgo.io/ent/entc/integration/migrate/versioned"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		t.Run(version, func(t *testing.T) {
			root, err := sql.Open(dialect.MySQL, fmt.Sprintf("root:pass@tcp(localhost:%d)/", port))
			require.NoError(t, err)
			defer root.Close()
			ctx := context.Background()
			err = root.Exec(ctx, "CREATE DATABASE IF NOT EXISTS migrate", []interface{}{}, new(sql.Result))
			require.NoError(t, err, "creating database")
			defer root.Exec(ctx, "DROP DATABASE IF EXISTS migrate", []interface{}{}, new(sql.Result))

			drv, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/migrate?parseTime=True", port))
			require.NoError(t, err, "connecting to migrate database")

			clientv1 := entv1.NewClient(entv1.Driver(drv))
			clientv2 := entv2.NewClient(entv2.Driver(drv))
			V1ToV2(t, drv.Dialect(), clientv1, clientv2)
			if version == "8" {
				CheckConstraint(t, clientv2)
			}
			NicknameSearch(t, clientv2)
			TimePrecision(t, drv, "SELECT datetime_precision FROM information_schema.columns WHERE table_name = ? AND column_name = ?")

			require.NoError(t, err, root.Exec(ctx, "DROP DATABASE IF EXISTS migrate", []interface{}{}, new(sql.Result)))
			require.NoError(t, root.Exec(ctx, "CREATE DATABASE IF NOT EXISTS migrate", []interface{}{}, new(sql.Result)))
			Versioned(t, versioned.NewClient(versioned.Driver(drv)))
		})
	}
}

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", port)
			root, err := sql.Open(dialect.Postgres, dsn)
			require.NoError(t, err)
			defer root.Close()
			ctx := context.Background()
			err = root.Exec(ctx, "DROP DATABASE IF EXISTS migrate", []interface{}{}, new(sql.Result))
			require.NoError(t, err)
			err = root.Exec(ctx, "CREATE DATABASE migrate", []interface{}{}, new(sql.Result))
			require.NoError(t, err, "creating database")
			defer root.Exec(ctx, "DROP DATABASE migrate", []interface{}{}, new(sql.Result))

			drv, err := sql.Open(dialect.Postgres, dsn+" dbname=migrate")
			require.NoError(t, err, "connecting to migrate database")
			defer drv.Close()

			err = drv.Exec(ctx, "CREATE TYPE customtype as range (subtype = time)", []interface{}{}, new(sql.Result))
			require.NoError(t, err, "creating custom type")

			clientv1 := entv1.NewClient(entv1.Driver(drv))
			clientv2 := entv2.NewClient(entv2.Driver(drv))
			V1ToV2(t, drv.Dialect(), clientv1, clientv2)
			CheckConstraint(t, clientv2)
			TimePrecision(t, drv, "SELECT datetime_precision FROM information_schema.columns WHERE table_name = $1 AND column_name = $2")
		})
	}
}

func TestSQLite(t *testing.T) {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer drv.Close()

	ctx := context.Background()
	client := entv2.NewClient(entv2.Driver(drv))
	require.NoError(
		t,
		client.Schema.Create(
			ctx,
			migratev2.WithGlobalUniqueID(true),
			migratev2.WithDropIndex(true),
			migratev2.WithDropColumn(true),
			schema.WithDiffHook(renameTokenColumn),
			schema.WithDiffHook(func(next schema.Differ) schema.Differ {
				return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
					// Example to hook into the diff process.
					changes, err := next.Diff(current, desired)
					if err != nil {
						return nil, err
					}
					// After diff, you can filter
					// changes or return new ones.
					return changes, nil
				})
			}),
			schema.WithApplyHook(func(next schema.Applier) schema.Applier {
				return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
					// Example to hook into the apply process, or implement
					// a custom applier. For example, write to a file.
					//
					//	for _, c := range plan.Changes {
					//		fmt.Printf("%s: %s", c.Comment, c.Cmd)
					//		if err := conn.Exec(ctx, c.Cmd, c.Args, nil); err != nil {
					//			return err
					//		}
					//	}
					//
					return next.Apply(ctx, conn, plan)
				})
			}),
		),
	)

	SanityV2(t, drv.Dialect(), client)
	u := client.User.Create().SetAge(1).SetName("x").SetNickname("x'").SetPhone("y").SaveX(ctx)
	idRange(t, client.Car.Create().SetOwner(u).SaveX(ctx).ID, 0, 1<<32)
	idRange(t, client.Conversion.Create().SaveX(ctx).ID, 1<<32-1, 2<<32)
	idRange(t, client.CustomType.Create().SaveX(ctx).ID, 2<<32-1, 3<<32)
	idRange(t, client.Group.Create().SaveX(ctx).ID, 3<<32-1, 4<<32)
	idRange(t, client.Media.Create().SaveX(ctx).ID, 4<<32-1, 5<<32)
	idRange(t, client.Pet.Create().SaveX(ctx).ID, 5<<32-1, 6<<32)
	idRange(t, u.ID, 6<<32-1, 7<<32)

	// Override the default behavior of LIKE in SQLite.
	// https://www.sqlite.org/pragma.html#pragma_case_sensitive_like
	_, err = drv.ExecContext(ctx, "PRAGMA case_sensitive_like=1")
	require.NoError(t, err)
	EqualFold(t, client)
	ContainsFold(t, client)
	CheckConstraint(t, client)
}

func TestStorageKey(t *testing.T) {
	require.Equal(t, "user_pet_id", migratev2.PetsTable.ForeignKeys[0].Symbol)
	require.Equal(t, "user_friend_id1", migratev2.FriendsTable.ForeignKeys[0].Symbol)
	require.Equal(t, "user_friend_id2", migratev2.FriendsTable.ForeignKeys[1].Symbol)
}

func Versioned(t *testing.T, client *versioned.Client) {
	ctx := context.Background()

	p := t.TempDir()
	dir, err := migrate.NewLocalDir(p)
	require.NoError(t, err)
	require.NoError(t, client.Schema.Diff(ctx, schema.WithDir(dir)))
	require.Equal(t, 2, countFiles(t, dir))

	p = t.TempDir()
	dir, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	require.NoError(t, client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithFormatter(sqltool.GooseFormatter)))
	require.Equal(t, 1, countFiles(t, dir))

	p = t.TempDir()
	dir, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	require.NoError(t, client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithFormatter(sqltool.FlywayFormatter)))
	require.Equal(t, 2, countFiles(t, dir))

	p = t.TempDir()
	dir, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	require.NoError(t, client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithFormatter(sqltool.LiquibaseFormatter)))
	require.Equal(t, 1, countFiles(t, dir))

	p = t.TempDir()
	dir, err = migrate.NewLocalDir(p)
	require.NoError(t, err)
	require.NoError(t, client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithSumFile()))
	require.Equal(t, 3, countFiles(t, dir))
	require.FileExists(t, filepath.Join(p, "atlas.sum"))
}

func V1ToV2(t *testing.T, dialect string, clientv1 *entv1.Client, clientv2 *entv2.Client) {
	ctx := context.Background()

	// Run migration and execute queries on v1.
	require.NoError(t, clientv1.Schema.Create(ctx, migratev1.WithGlobalUniqueID(true), schema.WithAtlas(true)))
	SanityV1(t, dialect, clientv1)

	// Run migration and execute queries on v2.
	require.NoError(t, clientv2.Schema.Create(ctx, migratev2.WithGlobalUniqueID(true), migratev2.WithDropIndex(true), migratev2.WithDropColumn(true), schema.WithAtlas(true), schema.WithDiffHook(renameTokenColumn), schema.WithApplyHook(fillNulls(dialect))))
	require.NoError(t, clientv2.Schema.Create(ctx, migratev2.WithGlobalUniqueID(true), migratev2.WithDropIndex(true), migratev2.WithDropColumn(true), schema.WithAtlas(true)), "should not create additional resources on multiple runs")
	SanityV2(t, dialect, clientv2)

	u := clientv2.User.Create().SetAge(1).SetName("foo").SetNickname("nick_foo").SetPhone("phone").SaveX(ctx)
	idRange(t, clientv2.Car.Create().SetOwner(u).SaveX(ctx).ID, 0, 1<<32)
	idRange(t, clientv2.Conversion.Create().SaveX(ctx).ID, 1<<32-1, 2<<32)
	// Since "users" created in the migration of v1, it will occupy the range of 1<<32-1 ... 2<<32-1,
	// even though they are ordered differently in the migration of v2 (groups, pets, users).
	idRange(t, u.ID, 3<<32-1, 4<<32)
	idRange(t, clientv2.Group.Create().SaveX(ctx).ID, 4<<32-1, 5<<32)
	idRange(t, clientv2.Media.Create().SaveX(ctx).ID, 5<<32-1, 6<<32)
	idRange(t, clientv2.Pet.Create().SaveX(ctx).ID, 6<<32-1, 7<<32)

	// SQL specific predicates.
	EqualFold(t, clientv2)
	ContainsFold(t, clientv2)

	// "renamed" field was renamed to "new_name".
	exist := clientv2.User.Query().Where(user.NewName("renamed")).ExistX(ctx)
	require.True(t, exist, "expect renamed column to have previous values")
}

func SanityV1(t *testing.T, dbdialect string, client *entv1.Client) {
	ctx := context.Background()
	u := client.User.Create().SetAge(1).SetName("foo").SetNickname("nick_foo").SetRenamed("renamed").SaveX(ctx)
	require.EqualValues(t, 1, u.Age)
	require.Equal(t, "foo", u.Name)

	err := client.User.Create().SetAge(2).SetName("foobarbazqux").Exec(ctx)
	require.Error(t, err, "name is limited to 10 chars")

	// Unique index on (name, address).
	client.User.Create().SetAge(3).SetName("foo").SetNickname("nick_foo_2").SetAddress("tlv").SetState(userv1.StateLoggedIn).SaveX(ctx)
	err = client.User.Create().SetAge(4).SetName("foo").SetAddress("tlv").Exec(ctx)
	require.Error(t, err)

	// Blob type limited to 255.
	u = u.Update().SetBlob([]byte("hello")).SaveX(ctx)
	require.Equal(t, "hello", string(u.Blob))
	err = u.Update().SetBlob(make([]byte, 256)).Exec(ctx)
	require.True(t, strings.Contains(t.Name(), "Postgres") || err != nil, "blob should be limited on SQLite and MySQL")

	// Invalid enum value.
	err = client.User.Create().SetAge(1).SetName("bar").SetNickname("nick_bar").SetState("unknown").Exec(ctx)
	require.Error(t, err)

	// Conversions
	client.Conversion.Create().
		SetName("zero").
		SetInt8ToString(0).
		SetUint8ToString(0).
		SetInt16ToString(0).
		SetUint16ToString(0).
		SetInt32ToString(0).
		SetUint32ToString(0).
		SetInt64ToString(0).
		SetUint64ToString(0).
		SaveX(ctx)

	client.Conversion.Create().
		SetName("min").
		SetInt8ToString(math.MinInt8).
		SetUint8ToString(0).
		SetInt16ToString(math.MinInt16).
		SetUint16ToString(0).
		SetInt32ToString(math.MinInt32).
		SetUint32ToString(0).
		SetInt64ToString(math.MinInt64).
		SetUint64ToString(0).
		SaveX(ctx)

	creator := client.Conversion.Create().
		SetName("max").
		SetInt8ToString(math.MaxInt8).
		SetUint8ToString(math.MaxUint8).
		SetInt16ToString(math.MaxInt16).
		SetUint16ToString(math.MaxUint16).
		SetInt32ToString(math.MaxInt32).
		SetUint32ToString(math.MaxUint32).
		SetInt64ToString(math.MaxInt64).
		SetUint64ToString(math.MaxUint64)
	if dbdialect == dialect.Postgres {
		// Postgres does not support unsigned types.
		creator.SetInt8ToString(math.MaxInt8).
			SetUint8ToString(math.MaxInt8).
			SetUint16ToString(math.MaxInt16).
			SetUint32ToString(math.MaxInt32).
			SetUint32ToString(math.MaxInt32).
			SetUint64ToString(math.MaxInt64)
	}
	creator.SaveX(ctx)
}

func SanityV2(t *testing.T, dbdialect string, client *entv2.Client) {
	ctx := context.Background()
	for _, u := range client.User.Query().AllX(ctx) {
		require.NotEmpty(t, u.NewToken, "old_token column should be renamed to new_token")
	}
	if dbdialect != dialect.SQLite {
		require.True(t, client.User.Query().ExistX(ctx), "table 'users' should contain rows after running the migration")
		users := client.User.Query().Select(user.FieldCreatedAt).AllX(ctx)
		for i := range users {
			require.False(t, users[i].CreatedAt.IsZero(), "default 'CURRENT_TIMESTAMP' should fill previous rows")
		}
	}
	u := client.User.Create().SetAge(1).SetName("bar").SetNickname("nick_bar").SetPhone("100").SetBuffer([]byte("{}")).SetState(user.StateLoggedOut).SaveX(ctx)
	require.Equal(t, 1, u.Age)
	require.Equal(t, "bar", u.Name)
	require.Equal(t, []byte("{}"), u.Buffer)
	u = u.Update().SetBuffer([]byte("[]")).SaveX(ctx)
	require.Equal(t, []byte("[]"), u.Buffer)
	require.Equal(t, user.StateLoggedOut, u.State)

	err := u.Update().SetState(user.State("boring")).Exec(ctx)
	require.Error(t, err, "invalid enum value")
	u = u.Update().SetState(user.StateOnline).SaveX(ctx)
	require.Equal(t, user.StateOnline, u.State)

	err = client.User.Create().SetAge(1).SetName("foobarbazqux").SetNickname("nick_bar").SetPhone("200").Exec(ctx)
	require.NoError(t, err, "name is not limited to 10 chars and nickname is not unique")

	// New unique index was added to (age, phone).
	err = client.User.Create().SetAge(1).SetName("foo").SetPhone("200").SetNickname("nick_bar").Exec(ctx)
	require.Error(t, err)
	require.True(t, entv2.IsConstraintError(err))

	// Ensure all rows in the database have the same default for the `title` column.
	require.Equal(
		t,
		client.User.Query().CountX(ctx),
		client.User.Query().Where(user.Title(user.DefaultTitle)).CountX(ctx),
	)

	// Blob type was extended.
	u, err = u.Update().SetBlob(make([]byte, 256)).SetState(user.StateLoggedOut).Save(ctx)
	require.NoError(t, err, "data type blob was extended in v2")
	require.Equal(t, make([]byte, 256), u.Blob)

	if dbdialect != dialect.SQLite {
		// Conversions
		zero := client.Conversion.Query().Where(conversion.Name("zero")).OnlyX(ctx)
		require.Equal(t, strconv.Itoa(0), zero.Int8ToString)
		require.Equal(t, strconv.Itoa(0), zero.Uint8ToString)
		require.Equal(t, strconv.Itoa(0), zero.Int16ToString)
		require.Equal(t, strconv.Itoa(0), zero.Uint16ToString)
		require.Equal(t, strconv.Itoa(0), zero.Int32ToString)
		require.Equal(t, strconv.Itoa(0), zero.Uint32ToString)
		require.Equal(t, strconv.Itoa(0), zero.Int64ToString)
		require.Equal(t, strconv.Itoa(0), zero.Uint64ToString)

		min := client.Conversion.Query().Where(conversion.Name("min")).OnlyX(ctx)
		require.Equal(t, strconv.Itoa(math.MinInt8), min.Int8ToString)
		require.Equal(t, strconv.Itoa(0), min.Uint8ToString)
		require.Equal(t, strconv.Itoa(math.MinInt16), min.Int16ToString)
		require.Equal(t, strconv.Itoa(0), min.Uint16ToString)
		require.Equal(t, strconv.Itoa(math.MinInt32), min.Int32ToString)
		require.Equal(t, strconv.Itoa(0), min.Uint32ToString)
		require.Equal(t, strconv.Itoa(math.MinInt64), min.Int64ToString)
		require.Equal(t, strconv.Itoa(0), min.Uint64ToString)

		max := client.Conversion.Query().Where(conversion.Name("max")).OnlyX(ctx)
		require.Equal(t, strconv.Itoa(math.MaxInt8), max.Int8ToString)
		require.Equal(t, strconv.Itoa(math.MaxInt16), max.Int16ToString)
		require.Equal(t, strconv.Itoa(math.MaxInt32), max.Int32ToString)
		require.Equal(t, strconv.Itoa(math.MaxInt64), max.Int64ToString)

		if dbdialect == dialect.Postgres {
			require.Equal(t, strconv.Itoa(math.MaxInt8), max.Uint8ToString)
			require.Equal(t, strconv.Itoa(math.MaxInt16), max.Uint16ToString)
			require.Equal(t, strconv.Itoa(math.MaxInt32), max.Uint32ToString)
			require.Equal(t, strconv.Itoa(math.MaxInt64), max.Uint64ToString)
		} else {
			require.Equal(t, strconv.Itoa(math.MaxUint8), max.Uint8ToString)
			require.Equal(t, strconv.Itoa(math.MaxUint16), max.Uint16ToString)
			require.Equal(t, strconv.Itoa(math.MaxUint32), max.Uint32ToString)
			require.Equal(t, strconv.FormatUint(math.MaxUint64, 10), max.Uint64ToString)
		}
	}
}

func CheckConstraint(t *testing.T, client *entv2.Client) {
	ctx := context.Background()
	t.Log("testing check constraints")
	err := client.Media.Create().SetText("boring").Exec(ctx)
	require.Error(t, err)
	err = client.Media.Create().SetSourceURI("entgo.io").Exec(ctx)
	require.Error(t, err)
}

func NicknameSearch(t *testing.T, client *entv2.Client) {
	ctx := context.Background()
	names := client.User.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.P(func(b *sql.Builder) {
				b.WriteString("MATCH(").Ident(user.FieldNickname).WriteString(") AGAINST(").Arg("nick_bar | nick_foo").WriteString(")")
			}))
		}).
		Unique(true).
		Order(entv2.Asc(user.FieldNickname)).
		Select(user.FieldNickname).
		StringsX(ctx)
	require.Equal(t, []string{"nick_bar", "nick_foo"}, names)
}

func EqualFold(t *testing.T, client *entv2.Client) {
	ctx := context.Background()
	t.Log("testing equal-fold on sql specific dialects")
	client.User.Create().SetAge(37).SetName("Alex").SetNickname("alexsn").SetPhone("123456789").SaveX(ctx)
	require.False(t, client.User.Query().Where(user.NameEQ("alex")).ExistX(ctx))
	require.True(t, client.User.Query().Where(user.NameEqualFold("alex")).ExistX(ctx))
}

func ContainsFold(t *testing.T, client *entv2.Client) {
	ctx := context.Background()
	t.Log("testing contains-fold on sql specific dialects")
	client.User.Create().SetAge(30).SetName("Mashraki").SetNickname("a8m").SetPhone("102030").SaveX(ctx)
	require.Zero(t, client.User.Query().Where(user.NameContains("mash")).CountX(ctx))
	require.Equal(t, 1, client.User.Query().Where(user.NameContainsFold("mash")).CountX(ctx))
	require.Equal(t, 1, client.User.Query().Where(user.NameContainsFold("Raki")).CountX(ctx))
}

func TimePrecision(t *testing.T, drv *sql.Driver, query string) {
	ctx := context.Background()
	rows, err := drv.QueryContext(ctx, query, customtype.Table, customtype.FieldTz0)
	require.NoError(t, err)
	p, err := sql.ScanInt(rows)
	require.NoError(t, err)
	require.Zerof(t, p, "custom_types field %q", customtype.FieldTz0)
	require.NoError(t, rows.Close())
	rows, err = drv.QueryContext(ctx, query, customtype.Table, customtype.FieldTz3)
	require.NoError(t, err)
	p, err = sql.ScanInt(rows)
	require.NoError(t, err)
	require.Equalf(t, 3, p, "custom_types field %q", customtype.FieldTz3)
	require.NoError(t, rows.Close())
}

func idRange(t *testing.T, id, l, h int) {
	require.Truef(t, id > l && id < h, "id %s should be between %d to %d", id, l, h)
}

func countFiles(t *testing.T, d migrate.Dir) int {
	files, err := fs.ReadDir(d, "")
	require.NoError(t, err)
	return len(files)
}

func renameTokenColumn(next schema.Differ) schema.Differ {
	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
		// Example to hook into the diff process.
		changes, err := next.Diff(current, desired)
		if err != nil {
			return nil, err
		}
		for _, c := range changes {
			m, ok := c.(*atlas.ModifyTable)
			if !ok || m.T.Name != user.Table {
				continue
			}
			changes := atlas.Changes(m.Changes)
			switch i, j := changes.IndexDropColumn("old_token"), changes.IndexAddColumn("new_token"); {
			case i != -1 && j != -1:
				rename := &atlas.RenameColumn{From: changes[i].(*atlas.DropColumn).C, To: changes[j].(*atlas.AddColumn).C}
				changes.RemoveIndex(i, j)
				m.Changes = append(changes, rename)
			case i != -1 || j != -1:
				return nil, errors.New("old_token and new_token must be present or absent")
			}
		}
		return changes, nil
	})
}

func fillNulls(dbdialect string) schema.ApplyHook {
	return func(next schema.Applier) schema.Applier {
		return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
			// There are three ways to UPDATE the NULL values to "Unknown" in this stage.
			// Append a custom migrate.Change to the plan, execute an SQL statement directly
			// on the dialect.ExecQuerier, or use the ent.Client used by the project.
			drv := sql.NewDriver(dbdialect, sql.Conn{ExecQuerier: conn.(*sql.Tx)})
			client := entv2.NewClient(entv2.Driver(drv))
			if err := client.User.
				Update().
				SetDropOptional("Unknown").
				Where(predicate.User(userv1.DropOptionalIsNil())).
				Exec(ctx); err != nil {
				return fmt.Errorf("fix default values to uppercase: %w", err)
			}
			return next.Apply(ctx, conn, plan)
		})
	}
}
