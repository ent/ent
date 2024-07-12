// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"log"
	"os"
	"testing"

	"ariga.io/atlas-go-sdk/atlasexec"
	"entgo.io/ent/dialect"
	"entgo.io/ent/examples/domaintypes/ent"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestDomainTypes(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	ctx := context.Background()
	client, err := ent.Open(dialect.Postgres, "postgres://postgres:pass@:5429/dev?search_path=public&sslmode=disable")
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
		URL: os.Getenv("DB_URL"),
		Env: "local",
	})
	require.NoError(t, err)
	t.Cleanup(func() { client.User.Delete().ExecX(ctx) })
	err = client.User.Create().SetPostalCode("foo").Exec(ctx)
	require.EqualError(t, err, `pq: value for domain us_postal_code violates check constraint "us_postal_code_check"`)
	err = client.User.Create().SetPostalCode("12345").Exec(ctx)
	require.NoError(t, err)
}
