// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/facebook/ent/examples/entcpkg/ent"

	_ "github.com/mattn/go-sqlite3"
)

func Example_EntcPkg() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	usr := client.User.Create().SaveX(ctx)
	fmt.Println("boring user:", usr)
	// Output: boring user: User(id=1)
}
