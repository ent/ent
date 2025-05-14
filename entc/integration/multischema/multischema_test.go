// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package multischema

import (
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"testing"

	"ariga.io/atlas-go-sdk/atlasexec"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/multischema/ent"
	"entgo.io/ent/entc/integration/multischema/ent/group"
	"entgo.io/ent/entc/integration/multischema/ent/migrate"
	"entgo.io/ent/entc/integration/multischema/ent/pet"
	"entgo.io/ent/entc/integration/multischema/ent/user"
	"entgo.io/ent/entc/integration/multischema/versioned"
	vgroup "entgo.io/ent/entc/integration/multischema/versioned/group"
	vpet "entgo.io/ent/entc/integration/multischema/versioned/pet"
	vuser "entgo.io/ent/entc/integration/multischema/versioned/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/?parseTime=true&multiStatements=true")
	require.NoError(t, err)
	ctx := context.Background()
	t.Cleanup(func() {
		db.ExecContext(ctx, "SET foreign_key_checks = 0")
		db.ExecContext(ctx, "DROP DATABASE IF EXISTS db1")
		db.ExecContext(ctx, "DROP DATABASE IF EXISTS db2")
		db.ExecContext(ctx, "SET foreign_key_checks = 1")
		db.Close()
	})

	migrate.ParentsTable.Schema = "db1"
	migrate.PetsTable.Schema = "db1"
	migrate.UsersTable.Schema = "db1"
	migrate.GroupsTable.Schema = "db2"
	migrate.GroupUsersTable.Schema = "db2"
	migrate.FriendshipsTable.Schema = "db2"

	pl, err := schema.Dump(ctx, dialect.MySQL, "8.0.19", migrate.Tables)
	require.NoError(t, err)
	_, err = db.ExecContext(ctx, pl)
	require.NoError(t, err)

	// Default schema for the connection is db1.
	db1, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/db1?parseTime=true")
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

		// The "groups", "group_users" and "friendship" reside in external schema (db2).
		Group:      "db2",
		GroupUsers: "db2",
		// An edge with the "Through" definition is set on its edge-schema.
		Friendship: "db2",
	}
	client := ent.NewClient(ent.Driver(db1), ent.AlternateSchema(cfg))
	pedro := client.Pet.Create().SetName("Pedro").SaveX(ctx)
	groups := client.Group.CreateBulk(
		client.Group.Create().SetName("GitHub"),
		client.Group.Create().SetName("GitLab"),
	).SaveX(ctx)
	a8m := client.User.Create().SetName("a8m").AddPets(pedro).AddGroups(groups...).SaveX(ctx)

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
		Where(group.HasUsersWith(user.ID(a8m.ID))).
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

	nat := client.User.Create().SetName("nati").AddFriends(a8m).SaveX(ctx)
	users := client.User.Query().WithFriends().WithFriendships().WithGroups().Order(ent.Asc(user.FieldName)).AllX(ctx)
	require.Len(t, users, 2)
	require.Equal(t, users[0].Name, a8m.Name)
	require.Equal(t, users[1].Name, nat.Name)
	require.Len(t, users[0].Edges.Groups, 1)
	require.Len(t, users[1].Edges.Groups, 0)
	require.Len(t, users[0].Edges.Friends, 1)
	require.Len(t, users[1].Edges.Friends, 1)
	require.Len(t, users[0].Edges.Friendships, 1)
	require.Len(t, users[1].Edges.Friendships, 1)

	ta := client.User.Create().SetName("ta").AddParents(a8m, nat).SaveX(ctx)
	el := client.User.Create().SetName("el").AddParents(a8m, nat).SaveX(ctx)
	jo := client.User.Create().SetName("be").AddParents(a8m, nat).SaveX(ctx)

	require.Equal(t, 3, client.User.Query().Where(user.HasParents()).CountX(ctx))
	require.Equal(t, 3, a8m.QueryChildren().CountX(ctx))

	sib := ta.QueryParents().QueryChildren().Where(user.NameNEQ(ta.Name)).AllX(ctx)
	require.Len(t, sib, 2)
	require.True(t, slices.ContainsFunc(sib, func(u *ent.User) bool { return u.Name == el.Name }))
	require.True(t, slices.ContainsFunc(sib, func(u *ent.User) bool { return u.Name == jo.Name }))
}

func TestVersionedMigration(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping on CI")
	}
	t.Cleanup(func() {
		db, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/")
		require.NoError(t, err)
		defer db.Close()
		for _, name := range []string{"db1", "db2", "db3", "atlas_schema_revisions"} {
			_, err := db.ExecContext(context.Background(), fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", name))
			require.NoError(t, err, "drop database")
		}
	})
	ac, err := atlasexec.NewClient(".", "atlas")
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	// Run `atlas migrate apply` on a SQLite database under /tmp.
	res, err := ac.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL:        "mysql://root:pass@:3308/",
		AllowDirty: true,
	})
	if err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}
	if len(res.Applied) > 0 {
		t.Logf("Applied %d migrations", len(res.Applied))
	}
	client, err := versioned.Open("mysql", "root:pass@tcp(localhost:3308)/?parseTime=true")
	require.NoError(t, err)
	defer client.Close()

	// Copy of the test above.
	ctx := context.Background()
	pedro := client.Pet.Create().SetName("Pedro").SaveX(ctx)
	groups := client.Group.CreateBulk(
		client.Group.Create().SetName("GitHub"),
		client.Group.Create().SetName("GitLab"),
	).SaveX(ctx)
	a8m := client.User.Create().SetName("a8m").AddPets(pedro).AddGroups(groups...).SaveX(ctx)

	// Custom modifier with schema config.
	var names []struct {
		User string `sql:"user_name"`
		Pet  string `sql:"pet_name"`
	}
	client.Pet.Query().
		Modify(func(s *sql.Selector) {
			// The below function is exported using a custom
			// template defined in ent/template/config.tmpl.
			cfg := versioned.DefaultSchemaConfig
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
		Where(vgroup.HasUsersWith(vuser.ID(a8m.ID))).
		Limit(1).
		QueryUsers().
		QueryPets().
		OnlyIDX(ctx)
	require.Equal(t, pedro.ID, id)

	affected := client.Group.
		Update().
		ClearUsers().
		Where(
			vgroup.And(
				vgroup.Name(groups[0].Name),
				vgroup.HasUsersWith(
					vuser.HasPetsWith(
						vpet.Name(pedro.Name),
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

	nat := client.User.Create().SetName("nati").AddFriends(a8m).SaveX(ctx)
	users := client.User.Query().WithFriends().WithFriendships().WithGroups().Order(ent.Asc(user.FieldName)).AllX(ctx)
	require.Len(t, users, 2)
	require.Equal(t, users[0].Name, a8m.Name)
	require.Equal(t, users[1].Name, nat.Name)
	require.Len(t, users[0].Edges.Groups, 1)
	require.Len(t, users[1].Edges.Groups, 0)
	require.Len(t, users[0].Edges.Friends, 1)
	require.Len(t, users[1].Edges.Friends, 1)
	require.Len(t, users[0].Edges.Friendships, 1)
	require.Len(t, users[1].Edges.Friendships, 1)
}
