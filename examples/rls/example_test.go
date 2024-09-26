// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"log"
	"os"
	"testing"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"entgo.io/ent/examples/rls/ent"

	"ariga.io/atlas-go-sdk/atlasexec"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestRowLevelSecurity(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	ctx := context.Background()
	// Note that APP_URL is used for the ent client, and ATLAS_URL is used for the atlas client.
	// Two different roles. The app role has access to the specific tables, and the atlas role in
	// this example is the default superuser role.
	client, err := ent.Open(dialect.Postgres, os.Getenv("APP_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	ac, err := atlasexec.NewClient(".", "atlas")
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	// Automatically update the database with the desired schema.
	// Another option, is to use 'migrate apply' or 'schema apply' manually.
	_, err = ac.SchemaApply(ctx, &atlasexec.SchemaApplyParams{
		// URL to your database. For example:
		// postgres://postgres:pass@localhost:5432/database?search_path=public&sslmode=disable
		URL: os.Getenv("ATLAS_URL"),
		Env: "local",
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		client.User.Delete().ExecX(ctx)
		client.Tenant.Delete().ExecX(ctx)
	})
	a8m, r3m := client.Tenant.Create().SetName("a8m").SaveX(ctx), client.Tenant.Create().SetName("r3m").SaveX(ctx)
	ctx1, ctx2 := sql.WithIntVar(ctx, "app.current_tenant", a8m.ID), sql.WithIntVar(ctx, "app.current_tenant", r3m.ID)
	u1 := client.User.Create().SetName("User: a8m").SetTenantID(a8m.ID).SaveX(ctx1)
	u2 := client.User.Create().SetName("User: r3m").SetTenantID(r3m.ID).SaveX(ctx2)
	users1 := client.User.Query().AllX(ctx1)
	require.Len(t, users1, 1)
	require.Equal(t, u1.ID, users1[0].ID)
	users2 := client.User.Query().AllX(ctx2)
	require.Len(t, users2, 1)
	require.Equal(t, u2.ID, users2[0].ID)
}
