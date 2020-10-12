// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"log"

	"github.com/facebook/ent/examples/privacyadmin/ent"
	_ "github.com/facebook/ent/examples/privacyadmin/ent/runtime"

	_ "github.com/mattn/go-sqlite3"
)

func Example_PrivacyTenant() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
}

func Do(ctx context.Context, client *ent.Client) error {
	return nil
}
