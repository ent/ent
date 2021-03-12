package upsert

import (
	"context"
	"testing"

	"entgo.io/ent/entc/integration/upsert/ent/enttest"
	"entgo.io/ent/entc/integration/upsert/ent/migrate"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func Test_PostgresUserUpsert(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
	uc, err := client.User.Create().
		SetEmail("alex-test@entgo.io").
		OnConflict("email").
		Save(ctx)
	require.NoError(t, err)
	require.NotNil(t, uc, "User record was inserted")
}

func Test_PostgresUserUpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
	uc1 := client.User.Create().
		SetEmail("alex-test@entgo.io").
		OnConflict("email")
	uc2 := client.User.Create().
		SetEmail("roger-test@entgo.io").
		OnConflict("email")
	uc3 := client.User.Create().
		SetEmail("alex-test@entgo.io").
		OnConflict("email")
		// OnConflict().UpdateEmail().UpdateName()

	users, err := client.User.CreateBulk(uc1, uc2, uc3).Save(ctx)

	require.NoError(t, err)
	require.NotNil(t, users, "User record was inserted")
}
