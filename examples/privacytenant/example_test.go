// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/facebook/ent/examples/privacytenant/ent"
	"github.com/facebook/ent/examples/privacytenant/ent/privacy"
	_ "github.com/facebook/ent/examples/privacytenant/ent/runtime"
	"github.com/facebook/ent/examples/privacytenant/viewer"

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
	// Tenant(id=1, name=GitHub)
	// Tenant(id=2, name=GitLab)
	// User(id=1, name=a8m)
	// User(id=2, name=nati)
	// Group(id=1, name=entgo.io)
}

func Do(ctx context.Context, client *ent.Client) error {
	// Expect operation to fail, because viewer-context
	// is missing (first mutation rule check).
	if _, err := client.Tenant.Create().Save(ctx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %v", err)
	}
	// Deny tenant creation if the viewer is not admin.
	viewOnly := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if _, err := client.Tenant.Create().Save(viewOnly); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %v", err)
	}
	// Apply the same operation with "Admin" role.
	admin := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	hub, err := client.Tenant.Create().SetName("GitHub").Save(admin)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
	fmt.Println(hub)
	lab, err := client.Tenant.Create().SetName("GitLab").Save(admin)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
	fmt.Println(lab)

	// Create 2 users connected to the 2 tenants we created above (a8m->GitHub, nati->GitLab).
	a8m := client.User.Create().SetName("a8m").SetTenant(hub).SaveX(admin)
	nati := client.User.Create().SetName("nati").SetTenant(lab).SaveX(admin)

	hubView := viewer.NewContext(ctx, viewer.UserViewer{T: hub})
	out := client.User.Query().OnlyX(hubView)
	// Expect that "GitHub" tenant to read only its users (i.e. a8m).
	if out.ID != a8m.ID {
		return fmt.Errorf("expect result for user query, got %v", out)
	}
	fmt.Println(out)

	labView := viewer.NewContext(ctx, viewer.UserViewer{T: lab})
	out = client.User.Query().OnlyX(labView)
	// Expect that "GitLab" tenant to read only its users (i.e. nati).
	if out.ID != nati.ID {
		return fmt.Errorf("expect result for user query, got %v", out)
	}
	fmt.Println(out)

	// Expect operation to fail, because the DenyMismatchedTenants rule makes sure
	// the group and the users are connected to the same tenant.
	_, err = client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(nati).Save(admin)
	if !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operatio to fail, since user (nati) is not connected to the same tenant")
	}
	_, err = client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(nati, a8m).Save(admin)
	if !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operatio to fail, since some users (nati) are not connected to the same tenant")
	}
	entgo, err := client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(a8m).Save(admin)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
	fmt.Println(entgo)

	return nil
}
