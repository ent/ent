// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package template

import (
	"context"
	"testing"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/config/ent"
	"github.com/facebook/ent/entc/integration/config/ent/migrate"
	"github.com/facebook/ent/entc/integration/config/ent/schema"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSchemaConfig(t *testing.T) {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer drv.Close()
	ctx := context.Background()
	client := ent.NewClient(ent.Driver(drv))
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	client.User.Create().SaveX(ctx)

	// Check that the table was created with the given custom name.
	table := schema.User{}.Annotations()[0].(entsql.Annotation).Table
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_master")).
		Where(sql.And(sql.EQ("type", "table"), sql.EQ("name", table))).
		Query()
	rows := &sql.Rows{}
	require.NoError(t, drv.Query(ctx, query, args, rows))
	defer rows.Close()
	require.True(t, rows.Next(), "no rows returned")
	var n int
	require.NoError(t, rows.Scan(&n), "scanning count")
	require.Equalf(t, 1, n, "expecting table %q to be exist", table)
}
