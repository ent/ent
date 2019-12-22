// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package customid

import (
	"context"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/customid/ent"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestCustomID(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx))

	nat := client.User.Create().SaveX(ctx)
	require.Equal(t, 1, nat.ID)
	_, err = client.User.Create().SetID(1).Save(ctx)
	require.True(t, ent.IsConstraintError(err), "duplicate id")
	a8m := client.User.Create().SetID(5).SaveX(ctx)
	require.Equal(t, 5, a8m.ID)

	hub := client.Group.Create().SetID(3).AddUsers(a8m, nat).SaveX(ctx)
	require.Equal(t, 3, hub.ID)
	require.Equal(t, []int{1, 5}, hub.QueryUsers().IDsX(ctx))

	b := client.Blob.Create().SetID(uuid.New()).SaveX(ctx)
	require.NotEmpty(t, b.ID)
}
