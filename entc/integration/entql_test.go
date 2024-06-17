// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"testing"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent"
	"entgo.io/ent/entc/integration/ent/comment"
	"entgo.io/ent/entc/integration/ent/pet"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/entql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func EntQL(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()

	a8m := client.User.Create().SetName("a8m").SetAge(30).SaveX(ctx)
	nati := client.User.Create().SetName("nati").SetAge(30).AddFriends(a8m).SaveX(ctx)

	uq := client.User.Query()
	uq.Filter().Where(entql.HasEdge("friends"))
	require.Equal(2, uq.CountX(ctx))

	uq = client.User.Query()
	uq.Filter().Where(
		entql.And(
			entql.FieldEQ("name", "nati"),
			entql.HasEdge("friends"),
		),
	)
	require.Equal(nati.ID, uq.OnlyIDX(ctx))

	u1, u2 := uuid.New(), uuid.New()
	xabi := client.Pet.Create().SetName("xabi").SetOwner(a8m).SetUUID(u1).SaveX(ctx)
	luna := client.Pet.Create().SetName("luna").SetOwner(nati).SetUUID(u2).SaveX(ctx)
	uq = client.User.Query()
	uq.Filter().Where(
		entql.And(
			entql.HasEdge("pets"),
			entql.HasEdgeWith("friends", entql.FieldEQ("name", "nati")),
			entql.HasEdgeWith("friends", entql.FieldIn("name", "nati")),
			entql.HasEdgeWith("friends", entql.FieldIn("name", "nati", "a8m")),
		),
	)
	require.Equal(a8m.ID, uq.OnlyIDX(ctx))
	uq = client.User.Query()
	uq.Filter().Where(
		entql.And(
			entql.HasEdgeWith("pets", entql.FieldEQ("name", "luna")),
			entql.HasEdge("friends"),
		),
	)
	require.Equal(nati.ID, uq.OnlyIDX(ctx))

	pq := client.Pet.Query()
	pq.Filter().WhereUUID(entql.ValueEQ(u1))
	require.Equal(xabi.ID, pq.OnlyIDX(ctx))
	pq = client.Pet.Query()
	pq.Filter().WhereUUID(entql.ValueEQ(u2))
	require.Equal(luna.ID, pq.OnlyIDX(ctx))

	uq = client.User.Query()
	uq.Filter().WhereName(entql.StringEQ("a8m"))
	require.Equal(a8m.ID, uq.OnlyIDX(ctx))
	pq = client.Pet.Query()
	pq.Filter().WhereName(entql.StringOr(entql.StringEQ("xabi"), entql.StringEQ("luna")))
	require.Equal([]int{luna.ID, xabi.ID}, pq.Order(ent.Asc(pet.FieldName)).IDsX(ctx))

	pq = client.Pet.Query()
	pq.Where(pet.Name(luna.Name)).Filter().WhereID(entql.IntEQ(luna.ID))
	require.Equal(luna.ID, pq.Order(ent.Asc(pet.FieldName)).OnlyIDX(ctx))
	pq = client.Pet.Query()
	pq.Where(pet.Name(luna.Name)).Filter().WhereID(entql.IntEQ(xabi.ID))
	require.False(pq.ExistX(ctx))

	update := client.User.Update().SetRole(user.RoleAdmin)
	update.Mutation().Filter().WhereName(entql.StringEQ(a8m.Name))
	updated := update.SaveX(ctx)
	require.Equal(1, updated)
	uq = client.User.Query()
	uq.Filter().WhereRole(entql.StringEQ(string(user.RoleAdmin)))
	require.Equal(a8m.ID, uq.OnlyIDX(ctx))

	uq = client.User.Query()
	uq.Filter().WhereName(entql.StringEQ(a8m.Name))
	uq = uq.QueryFriends()
	uq.Filter().WhereName(entql.StringEQ(nati.Name))
	require.Equal(luna.ID, uq.QueryPets().OnlyIDX(ctx))

	client.Comment.Create().SetUniqueInt(1).SetUniqueFloat(1.0).SaveX(ctx)
	client.Comment.Create().SetUniqueInt(2).SetUniqueFloat(2.0).SetNillableInt(1).SaveX(ctx)

	comments := client.Comment.Query().Order(comment.ByNillableInt(sql.OrderNullsFirst())).AllX(ctx)
	require.True(comments[0].NillableInt == nil)
	require.True(*comments[1].NillableInt == 1)

	comments = client.Comment.Query().Order(comment.ByNillableInt(sql.OrderNullsLast())).AllX(ctx)
	require.True(*comments[0].NillableInt == 1)
	require.True(comments[1].NillableInt == nil)

}
