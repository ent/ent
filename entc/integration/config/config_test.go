// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package template

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/config/ent"
	"github.com/facebook/ent/entc/integration/config/ent/migrate"
	"github.com/facebook/ent/entc/integration/config/ent/schema"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func TestSchemaConfig(t *testing.T) {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer drv.Close()
	ctx := context.Background()
	client := ent.NewClient(ent.Driver(drv))
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))
	client.User.Create().SetID(1).SaveX(ctx)

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

	// Check that the table was created with the expected values.
	idIncremental := schema.User{}.Fields()[0].Descriptor().Annotations[0].(entsql.Annotation).Incremental
	require.Equal(t, *idIncremental, migrate.Tables[0].Columns[0].Increment)
	size := schema.User{}.Fields()[1].Descriptor().Annotations[0].(entsql.Annotation).Size
	require.Equal(t, size, migrate.Tables[0].Columns[1].Size)
}

func TestMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		t.Run(version, func(t *testing.T) {
			root, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/", port))
			require.NoError(t, err)
			defer root.Close()
			ctx := context.Background()
			err = root.Exec(ctx, "CREATE DATABASE IF NOT EXISTS config", []interface{}{}, new(sql.Result))
			require.NoError(t, err, "creating database")
			defer root.Exec(ctx, "DROP DATABASE IF EXISTS config", []interface{}{}, new(sql.Result))

			drv, err := sql.Open("mysql", fmt.Sprintf("root:pass@tcp(localhost:%d)/config?parseTime=True", port))
			require.NoError(t, err, "connecting to migrate database")

			client := ent.NewClient(ent.Driver(drv))
			// Run schema creation.
			require.NoError(t, client.Schema.Create(ctx))

			u, err := client.User.Create().SetID(200).Save(ctx)
			require.NoError(t, err)
			assert.Equal(t, 200, u.ID)
			_, err = client.User.Create().Save(ctx)
			assert.Error(t, err)
		})
	}
}
