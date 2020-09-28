// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package template

import (
	"context"
	"reflect"
	"testing"

	"github.com/facebook/ent/entc/integration/template/ent"
	"github.com/facebook/ent/entc/integration/template/ent/migrate"
	"github.com/facebook/ent/entc/integration/template/ent/user"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestCustomTemplate(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))

	p := client.Pet.Create().SetAge(1).SaveX(ctx)
	u := client.User.Create().SetName("a8m").AddPets(p).SaveX(ctx)
	g := client.Group.Create().SetMaxUsers(10).SaveX(ctx)

	node, err := client.Node(ctx, p.ID)
	require.NoError(t, err)
	require.Equal(t, p.ID, node.ID)
	require.Equal(t, &ent.Field{Type: "int", Name: "Age", Value: "1"}, node.Fields[0])
	require.Equal(t, &ent.Edge{Type: "User", Name: "Owner", IDs: []int{u.ID}}, node.Edges[0])

	node, err = client.Node(ctx, u.ID)
	require.NoError(t, err)
	require.Equal(t, u.ID, node.ID)
	require.Equal(t, &ent.Field{Type: "string", Name: "Name", Value: "\"a8m\""}, node.Fields[0])
	require.Equal(t, &ent.Edge{Type: "Pet", Name: "Pets", IDs: []int{p.ID}}, node.Edges[0])

	node, err = client.Node(ctx, g.ID)
	require.NoError(t, err)
	require.Equal(t, g.ID, node.ID)
	require.Equal(t, &ent.Field{Type: "int", Name: "MaxUsers", Value: "10"}, node.Fields[0])

	// check for client additional fields.
	require.True(t, reflect.ValueOf(client).Elem().FieldByName("tables").IsValid())

	result := client.User.Query().Where(user.NameGlob("a8*")).
		AllX(ctx)
	require.Equal(t, 1, len(result))
}
