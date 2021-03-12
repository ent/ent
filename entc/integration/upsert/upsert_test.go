package upsert

import (
	"context"
	"testing"

	"entgo.io/ent/entc/integration/upsert/ent/enttest"
	"entgo.io/ent/entc/integration/upsert/ent/migrate"
	"entgo.io/ent/entc/integration/upsert/ent/user"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func Test_PostgresUserUpsert(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
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

func Test_PostgresUserUpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
	uc1 := client.User.Create().
		SetEmail("alex-test@entgo.io").
		SetUpdateCount(1).
		OnConflict("email")
	uc2 := client.User.Create().
		SetEmail("roger-test@entgo.io").
		OnConflict("email")
	uc3 := client.User.Create().
		SetEmail("alex-test@entgo.io").
		SetUpdateCount(2).
		OnConflict("email")
		// OnConflict().UpdateEmail().UpdateName()

	users, err := client.Debug().User.CreateBulk(uc1, uc2, uc3).Save(ctx)

	require.NoError(t, err)
	require.NotNil(t, users, "User record was inserted")

	user := client.User.Query().Where(user.Email(("alex-test@entgo.io"))).OnlyX(ctx)
	require.Equal(t, 2, user.UpdateCount, "User was updated correctly on conflict")
}
