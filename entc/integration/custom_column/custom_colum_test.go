// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package template

import (
	"context"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/custom_column/ent/user"
	"github.com/lib/pq/hstore"
	"github.com/stretchr/testify/require"
	"log"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/custom_column/ent"

	_ "github.com/lib/pq"
)

func TestCustomColumn(t *testing.T) {
	drv, err := sql.Open("postgres", "host=localhost port=5430 user=postgres dbname=test password=pass sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	db := drv.DB()

	_, err = db.Exec("drop schema public cascade ")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	_, err = db.Exec("create schema public")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	_, err = db.Exec("CREATE EXTENSION hstore")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	//if err != nil && err.Error() != `pg: extension "hstore" already exists` {
	//	log.Fatalf("failed opening connection to postgres: %v", err)
	//}

	client := ent.NewClient(ent.Driver(drv))

	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_ = client.User.Create().
		SetName("John").
		SetCustom(hstore.Hstore{Map: map[string]sql.NullString{
			"key1": {String: "value1", Valid: true},
		}}).SaveX(ctx)

	usr := client.User.Query().Where(user.CustomContainsKey("key1")).OnlyX(context.Background())

	require.NotNil(t, usr)

	usr = client.User.Query().Where(user.CustomContainsKey("key2")).FirstX(context.Background())

	require.Nil(t, usr)
}
