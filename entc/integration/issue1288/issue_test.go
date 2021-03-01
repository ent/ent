// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package issue1288

import (
	"context"
	"testing"

	"entgo.io/ent/entc/integration/issue1288/ent"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSchemaConfig(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx))
	a8m := client.User.Create().SetID(1).SetName("a8m").SaveX(ctx)
	client.Metadata.Create().SetUser(a8m).SetAge(10).SaveX(ctx)
	m1 := a8m.QueryMetadata().OnlyX(ctx)
	require.Equal(t, a8m.ID, m1.ID)
	_, err = client.Metadata.Create().SetID(a8m.ID).SetAge(10).Save(ctx)
	require.True(t, ent.IsConstraintError(err), "UNIQUE constraint failed: metadata.id")
}
