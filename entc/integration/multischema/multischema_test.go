// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package multischema

import (
	"context"
	"testing"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/multischema/ent"
	"entgo.io/ent/entc/integration/multischema/ent/group"
	"entgo.io/ent/entc/integration/multischema/ent/migrate"
	"entgo.io/ent/entc/integration/multischema/ent/pet"
	"entgo.io/ent/entc/integration/multischema/ent/user"

	atlas "ariga.io/atlas/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/")
	require.NoError(t, err)
	defer db.Close()
	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS db1")
	require.NoError(t, err, "creating database")
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS db2")
	require.NoError(t, err, "creating database")
	defer db.ExecContext(ctx, "DROP DATABASE IF EXISTS db1")
	defer db.ExecContext(ctx, "DROP DATABASE IF EXISTS db2")

	// Default schema for the connection is db1.
	db1, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/db1")
	require.NoError(t, err)
	defer db1.Close()

	cfg := ent.SchemaConfig{
		// The "users" and the "pets" table reside in the same schema
		// as the connection (the default search schema). Thus, there
		// is no need to set them explicitly both in the migration and
		// runtime statements.
		//
		// Pet:  "db1",
		// User: "db1",

		// The "groups" and "group_users" reside in external schema (db2).
		Group:      "db2",
		GroupUsers: "db2",
	}
	client := ent.NewClient(ent.Driver(db1), ent.AlternateSchema(cfg))
	setupSchema(t, client, cfg)
	pedro := client.Pet.Create().SetName("Pedro").SaveX(ctx)
	groups := client.Group.CreateBulk(
		client.Group.Create().SetName("GitHub"),
		client.Group.Create().SetName("GitLab"),
	).SaveX(ctx)
	usr := client.User.Create().SetName("a8m").AddPets(pedro).AddGroups(groups...).SaveX(ctx)

	// Custom modifier with schema config.
	var names []struct {
		User string `sql:"user_name"`
		Pet  string `sql:"pet_name"`
	}
	client.Pet.Query().
		Modify(func(s *sql.Selector) {
			// The below function is exported using a custom
			// template defined in ent/template/config.tmpl.
			cfg := ent.SchemaConfigFromContext(s.Context())
			t := sql.Table(user.Table).Schema(cfg.User)
			s.Join(t).On(s.C(pet.FieldOwnerID), t.C(user.FieldID))
			s.Select(
				sql.As(t.C(user.FieldName), "user_name"),
				sql.As(s.C(pet.FieldName), "pet_name"),
			)
		}).
		ScanX(ctx, &names)
	require.Len(t, names, 1)
	require.Equal(t, "a8m", names[0].User)
	require.Equal(t, "Pedro", names[0].Pet)

	id := client.Group.Query().
		Where(group.HasUsersWith(user.ID(usr.ID))).
		Limit(1).
		QueryUsers().
		QueryPets().
		OnlyIDX(ctx)
	require.Equal(t, pedro.ID, id)

	affected := client.Group.
		Update().
		ClearUsers().
		Where(
			group.And(
				group.Name(groups[0].Name),
				group.HasUsersWith(
					user.HasPetsWith(
						pet.Name(pedro.Name),
					),
				),
			),
		).
		SaveX(ctx)
	require.Equal(t, 1, affected)

	exist := groups[0].QueryUsers().ExistX(ctx)
	require.False(t, exist)
	exist = groups[1].QueryUsers().ExistX(ctx)
	require.True(t, exist)
	exist = pedro.QueryOwner().ExistX(ctx)
	require.True(t, exist)
	pedro = pedro.Update().ClearOwner().SaveX(ctx)
	exist = pedro.QueryOwner().ExistX(ctx)
	require.False(t, exist)

	require.Equal(t, client.User.Query().CountX(ctx), len(client.User.Query().AllX(ctx)))
	require.Equal(t, client.Pet.Query().CountX(ctx), len(client.Pet.Query().AllX(ctx)))
}

func setupSchema(t *testing.T, client *ent.Client, cfg ent.SchemaConfig) {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				for tt, s := range map[string]string{group.Table: cfg.Group, group.UsersTable: cfg.GroupUsers} {
					t1, ok := desired.Table(tt)
					require.True(t, ok)
					t1.SetSchema(atlas.New(s))
				}
				return next.Diff(current, desired)
			})
		}))
	require.NoError(t, err)
}
