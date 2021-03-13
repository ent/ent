package upsert

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/integration/upsert/ent"
	"entgo.io/ent/entc/integration/upsert/ent/enttest"
	"entgo.io/ent/entc/integration/upsert/ent/migrate"
	"entgo.io/ent/entc/integration/upsert/ent/user"
	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func Test_SQLIte_UserUpsert(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
	user1, err := client.User.Create().
		SetEmail("alex-test@entgo.io").
		OnConflict("email").
		Save(ctx)
	require.NoError(t, err)
	require.NotNil(t, user1, "User record was inserted")

	user2, err := client.Debug().User.Create().
		SetEmail("alex-test@entgo.io").
		SetUpdateCount(2).
		OnConflict("email").
		Save(ctx)
	require.NoError(t, err)
	require.NotNil(t, user2, "User 2 record was inserted")
}

func Test_Postgres_UserUpsertBulk(t *testing.T) {
	ctx := context.Background()

	dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", 5433)
	db, err := sql.Open(dialect.Postgres, dsn)
	require.NoError(t, err)
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE upsert_test")
	require.NoError(t, err, "creating database")
	defer db.Exec("DROP DATABASE upsert_test")

	client := enttest.Open(t, dialect.Postgres, dsn+" dbname=upsert_test", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()

	t.Run(bulkUpsert(ctx, client))
}

func Test_MySQL_UserUpsertBulk(t *testing.T) {
	for version, port := range map[string]int{"57": 3307, "8": 3308} {
		t.Run(version, func(t *testing.T) {
			ctx := context.Background()
			root, err := entsql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/", port))
			require.NoError(t, err)
			defer root.Close()
			err = root.Exec(ctx, "CREATE DATABASE IF NOT EXISTS upsert_test", []interface{}{}, new(sql.Result))
			require.NoError(t, err, "creating database")
			defer root.Exec(ctx, "DROP DATABASE IF EXISTS upsert_test", []interface{}{}, new(sql.Result))

			drv, err := entsql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/upsert_test?parseTime=True", port))
			require.NoError(t, err, "connecting to migrate database")

			client := ent.NewClient(ent.Driver(drv))
			// Run schema creation.
			require.NoError(t, client.Schema.Create(ctx))

			t.Run(bulkUpsert(ctx, client))
		})
	}
}

func Test_SQLLite_UserUpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()

	t.Run(bulkUpsert(ctx, client))
}

func bulkUpsert(ctx context.Context, client *ent.Client) (string, func(*testing.T)) {
	return "Bulk Upsert", func(t *testing.T) {
		_, err := client.Debug().User.Create().
			SetEmail("roger-test@entgo.io").
			OnConflict("email").
			Save(ctx)
		require.NoError(t, err)

		_, err = client.Debug().User.Create().
			SetEmail("alex-test@entgo.io").
			OnConflict("email").
			Save(ctx)
		require.NoError(t, err)

		users, err := client.Debug().User.CreateBulk(
			client.User.Create().SetEmail("roger-test@entgo.io").SetUpdateCount(2).OnConflict("email"),
			client.User.Create().SetEmail("alex-test@entgo.io").SetUpdateCount(2).OnConflict("email"),
		).Save(ctx)

		require.NoError(t, err)
		require.NotNil(t, users, "User record was inserted")

		user1 := client.User.Query().Where(user.Email(("alex-test@entgo.io"))).OnlyX(ctx)
		require.Equal(t, 2, user1.UpdateCount, "User was updated correctly on conflict")

		user2 := client.User.Query().Where(user.Email(("roger-test@entgo.io"))).OnlyX(ctx)
		require.Equal(t, 2, user2.UpdateCount, "User was updated correctly on conflict")
	}
}
