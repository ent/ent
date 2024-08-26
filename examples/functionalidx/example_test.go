// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/examples/functionalidx/ent"
	"entgo.io/ent/examples/functionalidx/ent/user"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestTriggersTypes(t *testing.T) {
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
		URL: os.Getenv("DB_URL"),
		Env: "local",
	})
	require.NoError(t, err)

	// Clean up tables before running assertion.
	client.User.Delete().ExecX(ctx)

	// Test that the unique index is enforced.
	client.User.Create().SetName("Ariel").SaveX(ctx)
	err = client.User.Create().SetName("ariel").Exec(ctx)
	require.EqualError(t, err, `ent: constraint failed: pq: duplicate key value violates unique constraint "unique_name"`)
	// Type-assert returned error.
	var pqerr *pq.Error
	require.True(t, errors.As(err, &pqerr))
	require.Equal(t, `duplicate key value violates unique constraint "unique_name"`, pqerr.Message)
	require.Equal(t, user.Table, pqerr.Table)
	require.Equal(t, "unique_name", pqerr.Constraint)
	require.Equal(t, pq.ErrorCode("23505"), pqerr.Code, "unique violation")
}
