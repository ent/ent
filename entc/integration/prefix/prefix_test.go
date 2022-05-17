// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package template

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/integration/prefix/ent/enttest"
	"entgo.io/ent/entc/integration/prefix/ent/entuser"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestCustomPrefix(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	_ = client.User.Create().SaveX(ctx)

	result := client.User.Query().Where(entuser.ID(1)).AllX(ctx)

	require.Equal(t, 1, len(result))
}
