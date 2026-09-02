// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package nativeuuid

import (
	"context"
	"testing"
	"uuid"

	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/integration/nativeuuid/ent"
	"entgo.io/ent/entc/integration/nativeuuid/ent/node"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

// End-to-end coverage of native uuid.UUID fields: create, query,
// predicates and NULL handling through the generated scan types.
func TestNativeUUID(t *testing.T) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx))

	zero := uuid.UUID{}
	n1 := client.Node.Create().SetID(zero).SaveX(ctx)
	require.Equal(t, zero, n1.Token)
	require.Nil(t, n1.Parent)

	n2 := client.Node.Create().
		SetID(uuid.UUID{2}).
		SetToken(uuid.UUID{3}).
		SetParent(uuid.UUID{4}).
		SaveX(ctx)
	require.Equal(t, uuid.UUID{3}, n2.Token)
	require.NotNil(t, n2.Parent)
	require.Equal(t, uuid.UUID{4}, *n2.Parent)

	// Predicates on the ID and optional fields.
	require.Equal(t, n2.ID, client.Node.Query().Where(node.ID(uuid.UUID{2})).OnlyIDX(ctx))
	require.Equal(t, n2.ID, client.Node.Query().Where(node.Token(uuid.UUID{3})).OnlyIDX(ctx))
	require.Equal(t, n1.ID, client.Node.Query().Where(node.TokenIsNil()).OnlyIDX(ctx))
	require.Len(t, client.Node.Query().Where(node.IDIn(zero, uuid.UUID{2})).AllX(ctx), 2)
	require.Equal(t, n1.ID, client.Node.Query().Where(node.ParentIsNil()).OnlyIDX(ctx))

	// Scan a non-NULL value into fields that were previously NULL.
	err = client.Node.UpdateOneID(zero).SetToken(uuid.UUID{5}).SetParent(uuid.UUID{6}).Exec(ctx)
	require.NoError(t, err)
	got := client.Node.Query().Where(node.ID(zero)).OnlyX(ctx)
	require.Equal(t, uuid.UUID{5}, got.Token)
	require.NotNil(t, got.Parent)
	require.Equal(t, uuid.UUID{6}, *got.Parent)

	// Clear back to NULL and verify the NULL scan path again.
	err = client.Node.UpdateOneID(zero).ClearToken().ClearParent().Exec(ctx)
	require.NoError(t, err)
	got = client.Node.GetX(ctx, zero)
	require.Equal(t, zero, got.Token)
	require.Nil(t, got.Parent)
}