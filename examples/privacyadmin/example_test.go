// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/facebook/ent/examples/privacyadmin/viewer"

	"github.com/facebook/ent/examples/privacyadmin/ent"
	"github.com/facebook/ent/examples/privacyadmin/ent/privacy"
	_ "github.com/facebook/ent/examples/privacyadmin/ent/runtime"

	_ "github.com/mattn/go-sqlite3"
)

func Example_PrivacyAdmin() {
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
	// 1
	// 1
	// 1
}

func Do(ctx context.Context, client *ent.Client) error {
	// Expect operation to fail, because viewer-context
	// is missing (first mutation rule check).
	if _, err := client.User.Create().Save(ctx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %v", err)
	}
	// Apply the same operation with "Admin" role.
	admin := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	if _, err := client.User.Create().Save(admin); err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
	// Apply the same operation with "ViewOnly" role.
	viewOnly := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if _, err := client.User.Create().Save(viewOnly); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %v", err)
	}
	// Allow all viewers to query users.
	for _, ctx := range []context.Context{ctx, viewOnly, admin} {
		// Operation should pass for all viewers.
		count := client.User.Query().CountX(ctx)
		fmt.Println(count)
	}
	// Bind a privacy decision to the context (bypass all other rules).
	allow := privacy.DecisionContext(ctx, privacy.Allow)
	if _, err := client.User.Create().Save(allow); err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
	return nil
}
