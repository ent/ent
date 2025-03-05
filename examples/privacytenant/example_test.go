// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"entgo.io/ent/examples/privacytenant/ent"
	"entgo.io/ent/examples/privacytenant/ent/privacy"
	_ "entgo.io/ent/examples/privacytenant/ent/runtime"
	"entgo.io/ent/examples/privacytenant/ent/user"
	"entgo.io/ent/examples/privacytenant/viewer"

	_ "github.com/mattn/go-sqlite3"
)

func Example_createTenants() {
	ctx := context.Background()
	client := open(ctx)
	defer client.Close()

	// Expect operation to fail in case viewer-context is missing.
	// First mutation privacy policy rule defined in BaseMixin.
	if err := client.Tenant.Create().Exec(ctx); !errors.Is(err, privacy.Deny) {
		log.Fatal("expect tenant creation to fail, but got:", err)
	}

	// Expect operation to fail in case the ent.User in the viewer-context
	// is not an admin user. Privacy policy defined in the Tenant schema.
	viewCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if err := client.Tenant.Create().Exec(viewCtx); !errors.Is(err, privacy.Deny) {
		log.Fatal("expect tenant creation to fail, but got:", err)
	}

	// Operation should pass successfully as the user in the viewer-context
	// is an admin user. First mutation privacy policy in Tenant schema.
	adminCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	hub, err := client.Tenant.Create().SetName("GitHub").Save(adminCtx)
	if err != nil {
		log.Fatal("expect tenant creation to pass, but got:", err)
	}
	fmt.Println(hub)

	lab, err := client.Tenant.Create().SetName("GitLab").Save(adminCtx)
	if err != nil {
		log.Fatal("expect tenant creation to pass, but got:", err)
	}
	fmt.Println(lab)

	// Output:
	// Tenant(id=1, name=GitHub)
	// Tenant(id=2, name=GitLab)
}

func Example_tenantView() {
	ctx := context.Background()
	client := open(ctx)
	defer client.Close()

	// Operation should pass successfully as the user in the viewer-context
	// is an admin user. First mutation privacy policy in Tenant schema.
	adminCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	hub := client.Tenant.Create().SetName("GitHub").SaveX(adminCtx)
	lab := client.Tenant.Create().SetName("GitLab").SaveX(adminCtx)

	// Create 2 tenant-specific viewer contexts.
	hubView := viewer.NewContext(ctx, viewer.UserViewer{T: hub})
	labView := viewer.NewContext(ctx, viewer.UserViewer{T: lab})

	// Create 2 users in each tenant.
	hubUsers := client.User.CreateBulk(
		client.User.Create().SetName("a8m").SetTenant(hub),
		client.User.Create().SetName("nati").SetTenant(hub),
	).SaveX(hubView)
	fmt.Println(hubUsers)

	labUsers := client.User.CreateBulk(
		client.User.Create().SetName("foo").SetTenant(lab),
		client.User.Create().SetName("bar").SetTenant(lab),
	).SaveX(labView)
	fmt.Println(labUsers)

	// Query users should fail in case viewer-context is missing.
	if _, err := client.User.Query().Count(ctx); !errors.Is(err, privacy.Deny) {
		log.Fatal("expect user query to fail, but got:", err)
	}

	// Ensure each tenant can see only its users.
	// First and only rule in TenantMixin.
	fmt.Println(client.User.Query().Select(user.FieldName).StringsX(hubView))
	fmt.Println(client.User.Query().CountX(hubView))
	fmt.Println(client.User.Query().Select(user.FieldName).StringsX(labView))
	fmt.Println(client.User.Query().CountX(labView))

	// Expect admin users to see everything. First
	// query privacy policy defined in BaseMixin.
	fmt.Println(client.User.Query().CountX(adminCtx)) // 4

	// Update operation with specific tenant-view should update
	// only the tenant in the viewer-context.
	client.User.Update().SetFoods([]string{"pizza"}).SaveX(hubView)
	fmt.Println(client.User.Query().AllX(hubView))
	fmt.Println(client.User.Query().AllX(labView))

	// Delete operation with specific tenant-view should delete
	// only the tenant in the viewer-context.
	client.User.Delete().ExecX(labView)
	fmt.Println(
		client.User.Query().CountX(hubView), // 2
		client.User.Query().CountX(labView), // 0
	)

	// DeleteOne with wrong viewer-context should cause the operation to fail with NotFoundError.
	err := client.User.DeleteOne(hubUsers[0]).Exec(labView)
	if !ent.IsNotFound(err) {
		log.Fatal("expect user deletion to fail, but got:", err)
	}
	fmt.Println(client.User.Query().CountX(hubView)) // 2

	// Unlike queries, admin users are not allowed to mutate tenant specific data.
	if err := client.User.DeleteOne(hubUsers[0]).Exec(adminCtx); !errors.Is(err, privacy.Deny) {
		log.Fatal("expect user deletion to fail, but got:", err)
	}

	// Output:
	// [User(id=1, tenant_id=1, name=a8m, foods=[]) User(id=2, tenant_id=1, name=nati, foods=[])]
	// [User(id=3, tenant_id=2, name=foo, foods=[]) User(id=4, tenant_id=2, name=bar, foods=[])]
	// [a8m nati]
	// 2
	// [foo bar]
	// 2
	// 4
	// [User(id=1, tenant_id=1, name=a8m, foods=[pizza]) User(id=2, tenant_id=1, name=nati, foods=[pizza])]
	// [User(id=3, tenant_id=2, name=foo, foods=[]) User(id=4, tenant_id=2, name=bar, foods=[])]
	// 2 0
	// 2
}

func Example_denyMismatchedTenants() {
	ctx := context.Background()
	client := open(ctx)
	defer client.Close()

	// Operation should pass successfully as the user in the viewer-context
	// is an admin user. First mutation privacy policy in Tenant schema.
	adminCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	hub := client.Tenant.Create().SetName("GitHub").SaveX(adminCtx)
	lab := client.Tenant.Create().SetName("GitLab").SaveX(adminCtx)

	// Create 2 tenant-specific viewer contexts.
	hubView := viewer.NewContext(ctx, viewer.UserViewer{T: hub})
	labView := viewer.NewContext(ctx, viewer.UserViewer{T: lab})

	// Create 2 users in each tenant.
	hubUsers := client.User.CreateBulk(
		client.User.Create().SetName("a8m").SetTenant(hub),
		client.User.Create().SetName("nati").SetTenant(hub),
	).SaveX(hubView)
	fmt.Println(hubUsers)

	labUsers := client.User.CreateBulk(
		client.User.Create().SetName("foo").SetTenant(lab),
		client.User.Create().SetName("bar").SetTenant(lab),
	).SaveX(labView)
	fmt.Println(labUsers)

	// Expect operation to fail as the DenyMismatchedTenants rule makes
	// sure the group and the users are connected to the same tenant.
	if err := client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(labUsers...).Exec(hubView); !errors.Is(err, privacy.Deny) {
		log.Fatal("expect operation to fail, since labUsers are not connected to the same tenant")
	}
	if err := client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(hubUsers[0], labUsers[0]).Exec(hubView); !errors.Is(err, privacy.Deny) {
		log.Fatal("expect operation to fail, since labUsers[0] is not connected to the same tenant")
	}
	// Expect mutation to pass as all users belong to the same tenant as the group.
	entgo := client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(hubUsers...).SaveX(hubView)
	fmt.Println(entgo)

	// Expect operation to fail, because the FilterTenantRule rule makes sure
	// that tenants can update and delete only their groups.
	if err := entgo.Update().SetName("fail.go").Exec(labView); !ent.IsNotFound(err) {
		log.Fatal("expect operation to fail, since the group (entgo) is managed by a different tenant (hub), but got:", err)
	}

	// Operation should pass in case it was applied with the right viewer-context.
	entgo = entgo.Update().SetName("entgo").SaveX(hubView)
	fmt.Println(entgo)

	// Output:
	// [User(id=1, tenant_id=1, name=a8m, foods=[]) User(id=2, tenant_id=1, name=nati, foods=[])]
	// [User(id=3, tenant_id=2, name=foo, foods=[]) User(id=4, tenant_id=2, name=bar, foods=[])]
	// Group(id=1, tenant_id=1, name=entgo.io)
	// Group(id=1, tenant_id=1, name=entgo)
}

func open(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
