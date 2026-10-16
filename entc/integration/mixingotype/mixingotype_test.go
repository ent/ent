package mixingotype

import (
	"context"
	"testing"

	"entgo.io/ent/entc/integration/ent/migrate"
	"entgo.io/ent/entc/integration/mixingotype/ent"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

// TestMixingoType tests the mixingotype feature, this is just a simple sanity test, as the original
// issue was happening on entc side.
func TestMixingoType(t *testing.T) {
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))

	user, err := client.User.Create().SetName("John Doe").SetAnnotations(map[string]interface{}{
		"age":  30,
		"city": "New York",
	}).Save(ctx)
	require.NoError(t, err)
	require.Equal(t, "John Doe", user.Name)
	require.Equal(t, map[string]interface{}{"age": 30, "city": "New York"}, user.Annotations)
}
