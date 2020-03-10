package hooks

import (
	"context"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/migrate"
	_ "github.com/facebookincubator/ent/entc/integration/hooks/ent/runtime"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSchemaHooks(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	_, err = client.Card.Create().SetNumber("123").Save(ctx)
	require.EqualError(t, err, "card number is too short", "error is returned from hook")
	crd := client.Card.Create().SetNumber("1234").SaveX(ctx)
	require.Equal(t, "unknown", crd.Name, "name was set by hook")
}
