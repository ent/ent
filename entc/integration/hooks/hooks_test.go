package hooks

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/hooks/ent"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/card"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/hook"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/migrate"
	_ "github.com/facebookincubator/ent/entc/integration/hooks/ent/runtime"
	"github.com/facebookincubator/ent/entc/integration/hooks/ent/user"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSchemaHooks(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	_, err = client.Card.Create().SetNumber("123").Save(ctx)
	require.EqualError(t, err, "card number is too short", "error is returned from hook")
	crd := client.Card.Create().SetNumber("1234").SaveX(ctx)
	require.Equal(t, "unknown", crd.Name, "name was set by hook")
	client.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			name, ok := m.Name()
			require.True(t, !ok && name == "", "should be the first hook to execute")
			return next.Mutate(ctx, m)
		})
	})
	client.Card.Create().SetNumber("1234").SaveX(ctx)
	_, err = client.Card.Update().Save(ctx)
	require.EqualError(t, err, "OpUpdate operation is not allowed")
}

func TestRuntimeHooks(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", ent.Log(t.Log))
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	var calls int
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			calls++
			return next.Mutate(ctx, m)
		})
	})
	client.Card.Create().SetNumber("1234").SaveX(ctx)
	client.User.Create().SetName("a8m").SaveX(ctx)
	require.Equal(t, 2, calls)
	client = client.Debug()
	client.Card.Create().SetNumber("1234").SaveX(ctx)
	client.User.Create().SetName("a8m").SaveX(ctx)
	require.Equal(t, 4, calls)
}

func TestRuntimeChain(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", ent.Log(t.Log))
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx))
	var (
		chain  hook.Chain
		values []int
	)
	for value := 0; value < 5; value++ {
		chain = chain.Append(func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				values = append(values, value)
				return next.Mutate(ctx, m)
			})
		})
	}
	client.User.Use(chain.Hook())
	client.User.Create().SetName("alexsn").SaveX(ctx)
	require.Len(t, values, 5)
	require.True(t, sort.IntsAreSorted(values))
}

func TestMutationClient(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			id, _ := m.OwnerID()
			usr := m.Client().User.GetX(ctx, id)
			m.SetName(usr.Name)
			return next.Mutate(ctx, m)
		})
	})
	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	crd := client.Card.Create().SetNumber("1234").SetOwner(a8m).SaveX(ctx)
	require.Equal(t, a8m.Name, crd.Name)
}

func TestMutationTx(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			tx, err := m.Tx()
			if err != nil {
				return nil, err
			}
			if err := tx.Rollback(); err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("rolled back")
		})
	})
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	a8m := tx.User.Create().SetName("a8m").SaveX(ctx)
	crd, err := tx.Card.Create().SetNumber("1234").SetOwner(a8m).Save(ctx)
	require.EqualError(t, err, "rolled back")
	require.Nil(t, crd)
	_, err = tx.Card.Query().All(ctx)
	require.Error(t, err, "tx already rolled back")
}

func TestDeletion(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	client.User.Use(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			if !m.Op().Is(ent.OpDeleteOne) {
				return next.Mutate(ctx, m)
			}
			id, ok := m.ID()
			if !ok {
				return nil, fmt.Errorf("missing id")
			}
			m.Client().Card.Delete().Where(card.HasOwnerWith(user.ID(id))).ExecX(ctx)
			return next.Mutate(ctx, m)
		})
	})
	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	for i := 0; i < 5; i++ {
		client.Card.Create().SetNumber(fmt.Sprintf("card-%d", i)).SetOwner(a8m).SaveX(ctx)
	}
	client.User.DeleteOne(a8m).ExecX(ctx)
	require.Zero(t, client.User.Query().CountX(ctx))
	require.Zero(t, client.Card.Query().CountX(ctx))
}
