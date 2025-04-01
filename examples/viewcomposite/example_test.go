// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"log"
	"os"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/examples/viewcomposite/ent"
	"entgo.io/ent/examples/viewcomposite/ent/petusername"

	"ariga.io/atlas-go-sdk/atlasexec"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestViews(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	ctx := context.Background()
	client, err := ent.Open(dialect.Postgres, os.Getenv("DB_URL"))
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
		URL:         os.Getenv("DB_URL"),
		Env:         "local",
		AutoApprove: true,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		client.User.Delete().ExecX(ctx)
		client.Pet.Delete().ExecX(ctx)
	})
	u1 := client.User.Create().SetName("a8m").SetPrivateInfo("secret").SetPublicInfo("public").SaveX(ctx)
	v1 := client.CleanUser.Query().OnlyX(ctx)
	require.Equal(t, u1.ID, v1.ID)
	require.Equal(t, u1.Name, v1.Name)
	require.Equal(t, u1.PublicInfo, v1.PublicInfo)

	p1 := client.Pet.Create().SetName("pedro").SaveX(ctx)
	names := client.PetUserName.Query().Order(petusername.ByName()).AllX(ctx)
	require.Len(t, names, 2)
	require.Equal(t, names[0].Name, u1.Name)
	require.Equal(t, names[1].Name, p1.Name)
}
