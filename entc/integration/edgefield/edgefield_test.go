// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package edgefield

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/integration/edgefield/ent"
	"entgo.io/ent/entc/integration/edgefield/ent/migrate"
	"entgo.io/ent/entc/integration/edgefield/ent/node"
	"entgo.io/ent/entc/integration/edgefield/ent/pet"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/entc/integration/edgefield/ent/user"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestEdgeField(t *testing.T) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))

	a8m := client.User.Create().SaveX(ctx)
	p1 := client.Pet.Create().SetOwner(a8m).SaveX(ctx)
	require.Equal(t, a8m.ID, p1.OwnerID)
	f1 := client.Pet.Query().Where(pet.OwnerID(a8m.ID)).OnlyX(ctx)
	require.Equal(t, p1.ID, f1.ID)
	require.Equal(t, p1.OwnerID, f1.OwnerID)

	c1 := client.User.Create().SetParent(a8m).SaveX(ctx)
	require.Equal(t, c1.ParentID, a8m.ID)
	c2 := client.User.Create().SetParentID(a8m.ID).SaveX(ctx)
	require.Equal(t, c2.ParentID, a8m.ID)
	pid := a8m.QueryChildren().GroupBy(user.FieldParentID).IntX(ctx)
	require.Equal(t, pid, a8m.ID)
	c3 := client.User.Create().SetParentID(c2.ID).SaveX(ctx)
	require.Equal(t,
		client.User.Query().
			Where(
				user.HasParentWith(
					user.ParentID(a8m.ID),
				),
			).OnlyIDX(ctx),
		c3.ID,
	)

	ps1 := client.Post.Create().SetText("entgo.io").SaveX(ctx)
	require.Nil(t, ps1.AuthorID)
	ps1 = ps1.Update().SetAuthorID(a8m.ID).SaveX(ctx)
	require.NotNil(t, ps1.AuthorID)
	require.Equal(t, a8m.ID, *ps1.AuthorID)
	ps1 = client.Post.Query().WithAuthor().OnlyX(ctx)
	require.NotNil(t, ps1.AuthorID)
	require.Equal(t, a8m.ID, *ps1.AuthorID)
	require.Equal(t, a8m.ID, ps1.Edges.Author.ID)

	nati := client.User.Create().SetSpouse(a8m).SaveX(ctx)
	require.Equal(t, nati.SpouseID, a8m.ID)
	require.Equal(t, nati.ID, a8m.QuerySpouse().OnlyIDX(ctx))

	visa := client.Card.Create().SetOwnerID(a8m.ID).SaveX(ctx)
	require.Equal(t, a8m.ID, visa.OwnerID)
	require.Equal(t, nati.ID, visa.QueryOwner().QuerySpouse().OnlyIDX(ctx))
	require.Equal(t, nati.ID, client.Card.Query().QueryOwner().QuerySpouse().OnlyIDX(ctx))

	m1 := client.Metadata.Create().SetUser(a8m).SetAge(10).SaveX(ctx)
	require.Equal(t, a8m.ID, m1.ID)
	require.Equal(t, 10, m1.Age)
	m1 = a8m.QueryMetadata().OnlyX(ctx)
	require.Equal(t, a8m.ID, m1.ID)
	require.Equal(t, a8m.ID, m1.QueryUser().OnlyIDX(ctx))
	_, err = client.Metadata.Create().SetID(a8m.ID).SetAge(10).Save(ctx)
	require.True(t, ent.IsConstraintError(err), "UNIQUE constraint failed: metadata.id")
	err = m1.Update().ClearUser().Exec(ctx)
	require.Error(t, err, "clearing primary key is not allowed")

	client.Info.Create().SetUser(a8m).SetContent(json.RawMessage("{}")).SaveX(ctx)
	inf := a8m.QueryInfo().OnlyX(ctx)
	require.Equal(t, a8m.ID, inf.ID)
	_, err = client.Info.Create().SetID(a8m.ID).SetContent(json.RawMessage("10")).Save(ctx)
	require.True(t, ent.IsConstraintError(err), "UNIQUE constraint failed: metadata.id")

	require.NotZero(t, client.Pet.Query().QueryOwner().CountX(ctx))
	client.Pet.Update().ClearOwnerID().ExecX(ctx)
	require.Zero(t, client.Pet.Query().QueryOwner().CountX(ctx))

	require.False(t, client.Rental.Query().ExistX(ctx))
	car1 := client.Car.Create().SetNumber("102030").SaveX(ctx)
	car2 := client.Car.Create().SetNumber("102030").SaveX(ctx)
	client.Rental.Create().SetUserID(a8m.ID).SetCarID(car1.ID).SaveX(ctx)
	require.Equal(t, car1.ID, a8m.QueryRentals().QueryCar().OnlyIDX(ctx))
	dt, err := time.Parse(time.RFC3339, "1906-01-02T00:00:00+00:00")
	require.NoError(t, err)
	client.Rental.Create().SetUserID(a8m.ID).SetCarID(car2.ID).SetDate(dt).SaveX(ctx)
	require.Equal(t, 2, a8m.QueryRentals().QueryCar().CountX(ctx))
	require.Equal(t, car2.ID, a8m.QueryRentals().Where(rental.DateLTE(dt)).QueryCar().OnlyIDX(ctx))
	_, err = client.Rental.Create().SetUserID(a8m.ID).SetCarID(car2.ID).SetDate(dt).Save(ctx)
	require.Error(t, err)
	require.True(t, ent.IsConstraintError(err))

	curr := client.Node.Create().SaveX(ctx)
	for i := 0; i < 5; i++ {
		curr = client.Node.Create().SetPrevID(curr.ID).SetValue(curr.Value + 1).SaveX(ctx)
	}
	head := client.Node.Query().Where(node.Not(node.HasPrev())).OnlyX(ctx)
	for i := 0; i < 5; i++ {
		curr = head.QueryNext().OnlyX(ctx)
		require.Equal(t, head.Value+1, curr.Value)
		head = curr
	}
}

func TestNamedEdges(t *testing.T) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	u1 := client.User.Create().SaveX(ctx)
	client.Pet.Create().SetOwner(u1).SaveX(ctx)

	u1 = client.User.Query().
		WithPets(func(q *ent.PetQuery) {
			q.Select(pet.FieldID)
		}).
		WithNamedPets("Named", func(q *ent.PetQuery) {
			q.Select(pet.FieldID)
		}).
		OnlyX(ctx)
	require.Len(t, u1.Edges.Pets, 1)
	require.Equal(t, u1.Edges.Pets[0].OwnerID, u1.ID)
	pets, err := u1.NamedPets("Named")
	require.NoError(t, err)
	require.Len(t, pets, 1)
	require.Equal(t, pets[0].OwnerID, u1.ID)
}
