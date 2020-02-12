// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"github.com/facebookincubator/ent/examples/customfield/ent/predicate"
	"github.com/facebookincubator/ent/examples/customfield/ent/user"
	"github.com/lib/pq/hstore"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/customfield/ent"
	_ "github.com/lib/pq"
)

func Example_EntcPkg() {
	drv, err := sql.Open("postgres", "host=localhost port=5430 user=postgres dbname=test password=pass sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	db := drv.DB()

	_, _ = db.Exec("CREATE EXTENSION hstore")

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

	usr := client.User.Create().
		SetAge(10).
		SetName("John").
		SetCustom(hstore.Hstore{Map: map[string]sql.NullString{
			"key1": {String: "value1", Valid: true},
		}}).SaveX(ctx)

	_, err = client.User.Query().Where(test()).All(context.Background())

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	fmt.Println("boring user:", usr)
	// Output: boring user: User(id=1)
}

func test() predicate.User {
	return func(s *sql.Selector) {
		s.Where(sql.EQ(fmt.Sprintf("%s->'%s'", user.FieldCustom, "key1"), sql.Raw("'value1'")))
	}
}
