// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package hooks

import (
	"context"
	"fmt"
	"sort"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/hooks/ent"
	"entgo.io/ent/entc/integration/hooks/ent/card"
	"entgo.io/ent/entc/integration/hooks/ent/enttest"
	"entgo.io/ent/entc/integration/hooks/ent/hook"
	"entgo.io/ent/entc/integration/hooks/ent/intercept"
	"entgo.io/ent/entc/integration/hooks/ent/migrate"
	"entgo.io/ent/entc/integration/hooks/ent/pet"
	"entgo.io/ent/entc/integration/hooks/ent/schema"
	"entgo.io/ent/entc/integration/hooks/ent/user"

	entgo "entgo.io/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSchemaHooks(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
	err := client.Card.Create().SetNumber("123").Exec(ctx)
	require.EqualError(t, err, "card number is too short", "error is returned from hook")
	crd := client.Card.Create().SetNumber("1234").SaveX(ctx)
	require.Equal(t, "unknown", crd.Name, "name was set by hook")
	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			name, ok := m.Name()
			require.True(t, !ok && name == "", "should be the first hook to execute")
			return next.Mutate(ctx, m)
		})
	})
	client.Card.Create().SetNumber("1234").SaveX(ctx)
	err = client.Card.Update().Exec(ctx)
	require.EqualError(t, err, "OpUpdate operation is not allowed")

	err = client.User.Update().SetPassword("pass").Exec(ctx)
	require.EqualError(t, err, "password cannot be edited on update-many")
	err = client.User.Update().ClearPassword().Exec(ctx)
	require.EqualError(t, err, "password cannot be edited on update-many")
}

func TestRuntimeHooks(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithOptions(ent.Log(t.Log)), enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
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
	require.Equal(t, 4, calls, "debug client should keep the same hooks")
}

func TestRuntimeChain(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
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
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
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

func TestMutatorClient(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
					op := ent.OpUpdate
					if _, exists := m.ID(); exists && m.Op().Is(ent.OpDeleteOne) {
						op = ent.OpUpdateOne
					}
					// Ensure card was not expired before.
					m.Where(card.ExpiredAtIsNil())
					// Count the number of affected records.
					ids, err := m.IDs(ctx)
					if err != nil {
						return nil, err
					}
					// Change the operation to update.
					m.SetOp(op)
					// Record when card was expired.
					m.SetExpiredAt(time.Now())
					// Execute the update operation.
					if _, err := m.Client().Mutate(ctx, m); err != nil {
						return nil, err
					}
					return len(ids), nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	)
	c1 := client.Card.Create().SetNumber("1234").SaveX(ctx)
	client.Card.DeleteOne(c1).ExecX(ctx)
	expired := client.Card.Query().OnlyX(ctx)
	require.False(t, expired.ExpiredAt.IsZero())

	client.Card.Create().SetNumber("4567").ExecX(ctx)
	client.Card.Create().SetNumber("7890").ExecX(ctx)
	client.Card.Delete().Where(card.Number("4567")).ExecX(ctx)
	cards := client.Card.Query().Order(ent.Asc(card.FieldNumber)).AllX(ctx)
	require.Len(t, cards, 3)
	require.True(t, cards[0].ExpiredAt.Equal(expired.ExpiredAt), "expired field should not be updated")
	require.False(t, cards[1].ExpiredAt.IsZero())
	require.True(t, cards[2].ExpiredAt.IsZero())

	client.Card.Delete().ExecX(ctx)
	require.False(t, client.Card.Query().Where(card.ExpiredAtIsNil()).ExistX(ctx))
}

func TestMutationTx(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
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
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
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

func TestMutationIDs(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	count := make(map[ent.Op]int)
	client.User.Use(
		hook.Unless(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
					count[m.Op()]++
					ids, err := m.IDs(ctx)
					require.NoError(t, err)
					require.Len(t, ids, 1)
					require.Equal(t, count[m.Op()], ids[0])
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	)
	for i := 0; i < 5; i++ {
		owner := client.User.Create().SetName(fmt.Sprintf("owner-%d", i)).SaveX(ctx)
		client.Card.Create().SetNumber(fmt.Sprintf("card-%d", i)).SetOwner(owner).ExecX(ctx)
	}
	for i := 0; i < 5; i++ {
		p := user.And(user.Name(fmt.Sprintf("owner-%d", i)), user.HasCardsWith(card.Number(fmt.Sprintf("card-%d", i))))
		client.User.Update().AddVersion(1).Where(p).ExecX(ctx)
		client.User.Delete().Where(p).ExecX(ctx)
	}
}

func TestPostCreation(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.Card.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			id, exists := m.ID()
			require.False(t, exists, "id should not exist pre mutation")
			require.Zero(t, id)
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			id, exists = m.ID()
			require.True(t, exists, "id should exist post mutation")
			require.NotZero(t, id)
			require.True(t, id == value.(*ent.Card).ID)
			return value, nil
		})
	}, ent.OpCreate))
	client.Card.Create().SetNumber("12345").SetName("a8m").SaveX(ctx)
	client.Card.CreateBulk(client.Card.Create().SetNumber("12345")).SaveX(ctx)
}

func TestUpdateAfterCreation(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.User.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			existingUser, ok := value.(*ent.User)
			require.Truef(t, ok, "value should be of type %T", existingUser)
			require.Equal(t, 1, existingUser.Version, "version does not match the original value")

			// After the user was created, return its updated version (a new object).
			newUser := m.Client().User.UpdateOne(existingUser).
				SetVersion(2).
				SaveX(ctx)
			return newUser, nil
		})
	}, ent.OpCreate))

	u := client.User.Create().SetName("a8m").SetVersion(1).SaveX(ctx)
	require.Equal(t, 2, u.Version, "version mutation in hook should have propagated back to call site")
}

func TestUpdateAfterUpdateOne(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.User.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			u, ok := value.(*ent.User)
			require.Truef(t, ok, "value should be of type %T", u)
			require.Equal(t, 2, u.Version, "version does not match the original value")

			// After the user was created, return its updated version (a new object).  Don't use UpdateOne because it
			// will cause recursive calls to this hook.
			m.Client().User.Update().
				Where(user.IDEQ(u.ID)).
				SetVersion(3).
				SaveX(ctx)

			return m.Client().User.Get(ctx, u.ID)
		})
	}, ent.OpUpdateOne))

	u := client.User.Create().SetName("a8m").SetVersion(1).SaveX(ctx)
	u = client.User.UpdateOne(u).SetVersion(2).SaveX(ctx)

	require.Equal(t, 3, u.Version, "version mutation in hook should have propagated back to call site")
}

func TestOldValues(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()

	// Querying old fields post mutation should fail.
	client.Card.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			require.NoError(t, err)
			_, err = m.OldNumber(ctx)
			require.Error(t, err)
			return value, nil
		})
	}, ent.OpUpdateOne))
	crd := client.Card.Create().SetNumber("1234").SetName("a8m").SaveX(ctx)
	client.Card.UpdateOneID(crd.ID).SetName("a8m").SaveX(ctx)

	// A typed hook.
	client.User.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			name, err := m.OldName(ctx)
			if err != nil {
				return nil, err
			}
			require.Equal(t, "a8m", name)
			return next.Mutate(ctx, m)
		})
	}, ent.OpUpdateOne))
	// A generic hook (executed on all types).
	client.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			namer, ok := m.(interface {
				OldName(context.Context) (string, error)
			})
			if !ok {
				// Skip if the mutation does not have
				// a method for getting the old name.
				return next.Mutate(ctx, m)
			}
			name, err := namer.OldName(ctx)
			if err != nil {
				return nil, err
			}
			require.Equal(t, "a8m", name)
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			_, err = namer.OldName(ctx)
			require.NoError(t, err)
			return value, nil
		})
	}, ent.OpUpdateOne))
	// A generic hook (executed on all types).
	client.Use(hook.Unless(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			namer, ok := m.(interface {
				OldName(context.Context) (string, error)
			})
			if !ok {
				// Skip if the mutation does not have
				// a method for getting the old name.
				return next.Mutate(ctx, m)
			}
			name, err := namer.OldName(ctx)
			if err != nil {
				return nil, err
			}
			require.Equal(t, "a8m", name)
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}
			_, err = namer.OldName(ctx)
			require.NoError(t, err)
			return value, nil
		})
	}, ^ent.OpUpdateOne))
	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	require.Equal(t, "a8m", a8m.Name)
	err := client.User.UpdateOne(a8m).SetName("Ariel").SetVersion(a8m.Version).Exec(ctx)
	require.EqualError(t, err, "version field must be incremented by 1")
	a8m = client.User.UpdateOne(a8m).SetName("Ariel").SetVersion(a8m.Version + 1).SaveX(ctx)
	require.Equal(t, "Ariel", a8m.Name)
}

func TestConditions(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()

	var calls int
	defer func() { require.Equal(t, 2, calls) }()
	client.Card.Use(hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			require.True(t, m.Op().Is(ent.OpUpdateOne))
			calls++
			return next.Mutate(ctx, m)
		})
	}, hook.Or(
		hook.HasFields(card.FieldName),
		hook.HasClearedFields(card.FieldName),
	)))
	client.User.Use(hook.If(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			require.True(t, m.Op().Is(ent.OpUpdate))
			incr, exists := m.AddedWorth()
			require.True(t, exists)
			require.EqualValues(t, 100, incr)
			return next.Mutate(ctx, m)
		})
	}, hook.HasAddedFields(user.FieldWorth)))

	ctx := context.Background()
	crd := client.Card.Create().SetNumber("9876").SaveX(ctx)
	crd = crd.Update().SetName("alexsn").SaveX(ctx)
	crd = crd.Update().ClearName().SaveX(ctx)
	client.Card.DeleteOne(crd).ExecX(ctx)

	alexsn := client.User.Create().SetName("alexsn").SaveX(ctx)
	client.User.Update().Where(user.ID(alexsn.ID)).AddWorth(100).SaveX(ctx)
	client.User.DeleteOne(alexsn).ExecX(ctx)
}

func TestRuntimeTx(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1", enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)))
	defer client.Close()
	client.Card.Use(func(next ent.Mutator) ent.Mutator {
		return hook.CardFunc(func(ctx context.Context, m *ent.CardMutation) (ent.Value, error) {
			v, err := next.Mutate(ctx, m)
			require.NoError(t, err)
			tx, err := m.Tx()
			require.NoError(t, err)
			tx.OnCommit(func(next ent.Committer) ent.Committer {
				return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
					// Ensure the transaction can see the created card.
					tx.Card.GetX(ctx, v.(*ent.Card).ID)
					// Cause the transaction to fail.
					require.NoError(t, tx.Rollback())
					return fmt.Errorf("fail")
				})
			})
			return v, nil
		})
	})
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	tx.Card.Create().SetNumber("9876").ExecX(ctx)
	require.EqualError(t, tx.Commit(), "fail")
	require.Zero(t, client.Card.Query().CountX(ctx), "database is empty")
}

func TestInterceptor_Sanity(t *testing.T) {
	ctx := context.Background()
	t.Run("All", func(t *testing.T) {
		var (
			calls  int
			client = enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
		)
		defer client.Close()
		client.Intercept(
			ent.InterceptFunc(func(next ent.Querier) ent.Querier {
				return intercept.UserFunc(func(ctx context.Context, query *ent.UserQuery) (ent.Value, error) {
					calls++
					nodes, err := next.Query(ctx, query)
					require.NoError(t, err)
					require.IsType(t, ([]*ent.User)(nil), nodes)
					return nodes, nil
				})
			}),
			intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
				calls++
				return nil
			}),
		)
		client.User.Query().AllX(ctx)
		require.Equal(t, 2, calls)
	})
	t.Run("Count", func(t *testing.T) {
		var (
			calls  int
			client = enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
		)
		defer client.Close()
		client.Intercept(
			ent.InterceptFunc(func(next ent.Querier) ent.Querier {
				return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
					calls++
					count, err := next.Query(ctx, query)
					require.NoError(t, err)
					require.Equal(t, 0, count)
					return count, nil
				})
			}),
			intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
				calls++
				return nil
			}),
		)
		client.User.Query().CountX(ctx)
		require.Equal(t, 2, calls)
	})
	t.Run("IDs", func(t *testing.T) {
		var (
			calls  int
			client = enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
			users  = client.User.CreateBulk(
				client.User.Create().SetName("a8m"),
				client.User.Create().SetName("nati"),
			).SaveX(ctx)
		)
		defer client.Close()
		client.Intercept(
			ent.InterceptFunc(func(next ent.Querier) ent.Querier {
				return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
					calls++
					ids, err := next.Query(ctx, query)
					require.NoError(t, err)
					// IDs() uses Select() under the hood, therefore,
					// scanned values are returned as values and not pointers.
					require.IsType(t, ([]int)(nil), ids)
					require.Equal(t, []int{users[0].ID}, ids)
					// Values can be changed and affect the return value.
					return append(ids.([]int), 10, 20), nil
				})
			}),
			intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
				calls++
				query.WhereP(user.ID(users[0].ID))
				return nil
			}),
		)
		require.Equal(t, []int{users[0].ID, 10, 20}, client.User.Query().IDsX(ctx))
		require.Equal(t, 2, calls)
	})
	t.Run("GroupBy", func(t *testing.T) {
		type n2c struct {
			N string `sql:"name"`
			C int    `sql:"count"`
		}
		var (
			calls  int
			client = enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
		)
		defer client.Close()
		client.Intercept(
			ent.InterceptFunc(func(next ent.Querier) ent.Querier {
				return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
					calls++
					vs, err := next.Query(ctx, query)
					require.NoError(t, err)
					return append(vs.([]n2c), n2c{N: "fake", C: 10}), nil
				})
			}),
			intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
				calls++
				query.WhereP(card.NameNEQ("a8m"))
				return nil
			}),
		)
		client.Card.CreateBulk(
			client.Card.Create().SetName("a8m").SetNumber("1234"),
			client.Card.Create().SetName("a8m").SetNumber("5678"),
			client.Card.Create().SetName("nati").SetNumber("9876"),
		).ExecX(ctx)
		var vs []n2c
		require.NoError(t, client.Card.Query().GroupBy(user.FieldName).Aggregate(ent.Count()).Scan(ctx, &vs))
		require.Equal(t, []n2c{{"nati", 1}, {"fake", 10}}, vs)
		require.Equal(t, 2, calls)
	})
	t.Run("Clone", func(t *testing.T) {
		var (
			calls  int
			client = enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
		)
		defer client.Close()
		client.Intercept(
			ent.InterceptFunc(func(next ent.Querier) ent.Querier {
				return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
					calls++
					count, err := next.Query(ctx, query)
					require.NoError(t, err)
					require.Equal(t, 0, count)
					return count, nil
				})
			}),
			intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
				calls++
				return nil
			}),
		)
		client.User.Query().CountX(ctx)
		require.Equal(t, 2, calls)
		client.User.Query().Clone().CountX(ctx)
		require.Equal(t, 4, calls)
	})
}

func TestSoftDelete(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()

	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	pets := client.Pet.CreateBulk(
		client.Pet.Create().SetName("a").SetOwner(a8m),
		client.Pet.Create().SetName("b").SetOwner(a8m),
		// Set delete_time manually.
		client.Pet.Create().SetName("c").SetOwner(a8m).SetDeleteTime(time.Now()),
	).SaveX(ctx)
	require.Equal(t, []int{pets[0].ID, pets[1].ID}, client.Pet.Query().Order(ent.Asc(pet.FieldID)).IDsX(ctx))

	// DeleteOne using the API.
	client.Pet.DeleteOne(pets[1]).ExecX(ctx)
	require.Equal(t, pets[0].ID, client.Pet.Query().OnlyIDX(ctx))

	// Delete using the API.
	n := client.Pet.Delete().ExecX(ctx)
	require.Equal(t, 1, n)
	require.False(t, client.Pet.Query().ExistX(ctx))

	// Query entities through the interceptor.
	require.Zero(t, client.Pet.Query().CountX(ctx))
	require.Empty(t, client.Pet.Query().AllX(ctx))
	// Skip soft-deleted interceptors.
	require.Equal(t, 3, client.Pet.Query().CountX(schema.SkipSoftDelete(ctx)))
	require.Len(t, client.Pet.Query().AllX(schema.SkipSoftDelete(ctx)), 3)

	client.Pet.CreateBulk(
		client.Pet.Create().SetName("d"),
		client.Pet.Create().SetName("d"),
		client.Pet.Create().SetName("d"),
	).ExecX(ctx)

	// Select entities through the interceptor.
	names := client.Pet.Query().Unique(true).Select(pet.FieldName).StringsX(ctx)
	require.Equal(t, []string{"d"}, names)
	names = client.Pet.Query().Unique(true).Order(ent.Asc(pet.FieldName)).Select(pet.FieldName).StringsX(schema.SkipSoftDelete(ctx))
	require.Equal(t, []string{"a", "b", "c", "d"}, names)

	// Group by.
	var n2c []struct {
		N string `sql:"name"`
		C int    `sql:"count"`
	}
	client.Pet.Query().GroupBy(pet.FieldName).Aggregate(ent.Count()).ScanX(ctx, &n2c)
	require.Equal(t, "d", n2c[0].N)
	require.Equal(t, 3, n2c[0].C)
	n2c = nil
	client.Pet.Query().GroupBy(pet.FieldName).Aggregate(ent.Count()).ScanX(schema.SkipSoftDelete(ctx), &n2c)
	require.Len(t, n2c, 4)

	// Edge traversals.
	require.False(t, a8m.QueryPets().ExistX(ctx), "interceptors should be applied on edge traversals")
	require.False(t, client.User.Query().QueryPets().ExistX(ctx))
	require.Equal(t, 3, a8m.QueryPets().CountX(schema.SkipSoftDelete(ctx)))

	// Eager-loading edges.
	a8m = client.User.Query().WithPets().OnlyX(ctx)
	require.Empty(t, a8m.Edges.Pets)
}

func TestTraverseUnique(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()

	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	client.Pet.CreateBulk(
		client.Pet.Create().SetName("a").SetOwner(a8m),
		client.Pet.Create().SetName("b").SetOwner(a8m),
	).ExecX(ctx)
	require.Equal(t, 1, client.Pet.Query().QueryOwner().CountX(ctx))

	// Disable unique traversal using interceptors.
	client.User.Intercept(
		intercept.Func(func(ctx context.Context, q intercept.Query) error {
			// Skip setting the Unique if the modifier was set explicitly.
			if entgo.QueryFromContext(ctx).Unique == nil {
				q.Unique(false)
			}
			return nil
		}),
	)
	// The JOIN with pets will return the same owner twice, one for each pet.
	require.Equal(t, 2, client.Pet.Query().QueryOwner().CountX(ctx))
	require.Equal(t, 1, client.Pet.Query().QueryOwner().Unique(true).CountX(ctx))
}

// The following example demonstrates how to write interceptors that
// can be used by multiple packages/projects using generics.
func TestSharedInterceptor(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.User.Create().SetName("a8m").ExecX(ctx)
	client.User.Create().SetName("nati").ExecX(ctx)
	require.Len(t, client.User.Query().AllX(ctx), 2)
	client.Intercept(SharedLimiter(intercept.NewQuery, 1))
	require.Len(t, client.User.Query().AllX(ctx), 1)
	require.Len(t, client.User.Query().Limit(10).AllX(ctx), 2)
}

// Project-level example. The usage of "entgo" package demonstrates how to
// write a generic interceptor that can be used by any ent-based project.
func SharedLimiter[Q interface{ Limit(int) }](f func(entgo.Query) (Q, error), limit int) entgo.Interceptor {
	return entgo.InterceptFunc(func(next entgo.Querier) entgo.Querier {
		return entgo.QuerierFunc(func(ctx context.Context, query entgo.Query) (ent.Value, error) {
			l, err := f(query)
			if err != nil {
				return nil, err
			}
			// LimitInterceptor limits the number of records returned from the
			// database to the configured one, in case Limit was not explicitly set.
			if entgo.QueryFromContext(ctx).Limit == nil {
				l.Limit(limit)
			}
			return next.Query(ctx, query)
		})
	})
}

func TestTypedTraverser(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	a8m, nat := client.User.Create().SetName("a8m").SaveX(ctx), client.User.Create().SetName("nati").SetActive(false).SaveX(ctx)
	client.Pet.CreateBulk(
		client.Pet.Create().SetName("a").SetOwner(a8m),
		client.Pet.Create().SetName("b").SetOwner(a8m),
		client.Pet.Create().SetName("c").SetOwner(nat),
	).ExecX(ctx)

	// Get all pets of all users.
	if n := client.User.Query().QueryPets().CountX(ctx); n != 3 {
		t.Errorf("got %d pets, want 3", n)
	}

	// Add an interceptor that filters out inactive users.
	client.User.Intercept(
		intercept.TraverseUser(func(ctx context.Context, q *ent.UserQuery) error {
			q.Where(user.Active(true))
			return nil
		}),
	)

	// Only pets of active users are returned.
	if n := client.User.Query().QueryPets().CountX(ctx); n != 2 {
		t.Errorf("got %d pets, want 2", n)
	}
}

func TestLimitInterceptor(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	client.User.Create().SetName("a8m").SaveX(ctx)
	client.User.Create().SetName("nati").SaveX(ctx)
	require.Len(t, client.User.Query().AllX(ctx), 2)
	client.Intercept(
		intercept.Func(func(ctx context.Context, q intercept.Query) error {
			q.Limit(1)
			return nil
		}),
	)
	require.Len(t, client.User.Query().AllX(ctx), 1)
}

func TestFilterTraverseFunc(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&_fk=1")
	defer client.Close()
	a8m, nat := client.User.Create().SetName("a8m").SaveX(ctx), client.User.Create().SetName("nati").SetActive(false).SaveX(ctx)
	client.Pet.CreateBulk(
		client.Pet.Create().SetName("a").SetOwner(a8m),
		client.Pet.Create().SetName("b").SetOwner(a8m),
		client.Pet.Create().SetName("c").SetOwner(nat),
	).ExecX(ctx)
	// Get all pets of all users.
	if n := client.User.Query().QueryPets().CountX(ctx); n != 3 {
		t.Errorf("got %d pets, want 3", n)
	}

	// Add an interceptor that filters out inactive users.
	client.User.Intercept(
		intercept.TraverseFunc(func(ctx context.Context, query intercept.Query) error {
			query.WhereP(sql.FieldEQ("active", true))
			return nil
		}),
	)
	// Only pets of active users are returned.
	if n := client.User.Query().QueryPets().CountX(ctx); n != 2 {
		t.Errorf("got %d pets, want 2", n)
	}
}
