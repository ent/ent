// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/facebook/ent/entc/integration/postgres/ent"
	"testing"

	"github.com/facebook/ent/dialect"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5433, "13": 5434} {
		t.Run(version, func(t *testing.T) {
			dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", port)
			db, err := sql.Open(dialect.Postgres, dsn)
			require.NoError(t, err)
			defer db.Close()
			_, err = db.Exec("CREATE DATABASE test_postgres")
			require.NoError(t, err, "creating database")
			defer db.Exec("DROP DATABASE test_postgres")

			client, err := ent.Open(dialect.Postgres, dsn+" dbname=test_postgres")
			require.NoError(t, err, "connecting to test_postgres database")
			defer client.Close()
			err = client.Schema.Create(context.Background())
			require.NoError(t, err)
			Postgres(t, client)
		})
	}
}

func Postgres(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	nat := client.User.Create().SetAge(24).SetName("Nat").SaveX(ctx)
	client.Group.Create().SetName("hub").AddUsers(nat).SaveX(ctx)
}
