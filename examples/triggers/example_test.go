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
	"entgo.io/ent/examples/triggers/ent"
	"entgo.io/ent/examples/triggers/ent/userauditlog"

	"ariga.io/atlas-go-sdk/atlasexec"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestTriggers(t *testing.T) {
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
		client.UserAuditLog.Delete().ExecX(ctx)
	})
	client.User.Create().SetName("a8m").ExecX(ctx)
	logs := client.UserAuditLog.Query().AllX(ctx)
	require.Len(t, logs, 1)
	require.Equal(t, "INSERT", logs[0].OperationType)
	require.Empty(t, logs[0].OldValue)
	require.Contains(t, logs[0].NewValue, "a8m")

	client.User.Update().SetName("Ariel").ExecX(ctx)
	logs = client.UserAuditLog.Query().Order(userauditlog.ByID()).AllX(ctx)
	require.Len(t, logs, 2)
	require.Equal(t, "UPDATE", logs[1].OperationType)
	require.Contains(t, logs[1].OldValue, "a8m")
	require.Contains(t, logs[1].NewValue, "Ariel")
}
