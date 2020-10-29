// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package customid

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"strconv"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/entc/integration/customid/ent"
	"github.com/facebook/ent/entc/integration/customid/ent/pet"
	"github.com/facebook/ent/entc/integration/customid/ent/user"
	"github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			cfg := mysql.Config{
				User: "root", Passwd: "pass", Net: "tcp", Addr: addr,
				AllowNativePasswords: true, ParseTime: true,
			}
			db, err := sql.Open("mysql", cfg.FormatDSN())
			require.NoError(t, err)
			defer db.Close()
			_, err = db.Exec("CREATE DATABASE IF NOT EXISTS custom_id")
			require.NoError(t, err, "creating database")
			defer db.Exec("DROP DATABASE IF EXISTS custom_id")

			cfg.DBName = "custom_id"
			client, err := ent.Open("mysql", cfg.FormatDSN())
			require.NoError(t, err, "connecting to json database")
			err = client.Schema.Create(context.Background())
			require.NoError(t, err)
			CustomID(t, client)
		})
	}
}

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5433} {
		t.Run(version, func(t *testing.T) {
			dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", port)
			db, err := sql.Open(dialect.Postgres, dsn)
			require.NoError(t, err)
			defer db.Close()
			_, err = db.Exec("CREATE DATABASE custom_id")
			require.NoError(t, err, "creating database")
			defer db.Exec("DROP DATABASE custom_id")

			client, err := ent.Open(dialect.Postgres, dsn+" dbname=custom_id")
			require.NoError(t, err, "connecting to json database")
			defer client.Close()
			err = client.Schema.Create(context.Background())
			require.NoError(t, err)
			CustomID(t, client)
		})
	}
}

func TestSQLite(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(context.Background()))
	CustomID(t, client)
}

func CustomID(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	nat := client.User.Create().SaveX(ctx)
	require.Equal(t, 1, nat.ID)
	_, err := client.User.Create().SetID(1).Save(ctx)
	require.True(t, ent.IsConstraintError(err), "duplicate id")
	a8m := client.User.Create().SetID(5).SaveX(ctx)
	require.Equal(t, 5, a8m.ID)

	hub := client.Group.Create().SetID(3).AddUsers(a8m, nat).SaveX(ctx)
	require.Equal(t, 3, hub.ID)
	require.Equal(t, []int{1, 5}, hub.QueryUsers().Order(ent.Asc(user.FieldID)).IDsX(ctx))

	blb := client.Blob.Create().SaveX(ctx)
	require.NotEmpty(t, blb.ID, "use default value")
	id := uuid.New()
	chd := client.Blob.Create().SetID(id).SetParent(blb).SaveX(ctx)
	require.Equal(t, id, chd.ID, "use provided id")
	require.Equal(t, blb.ID, chd.QueryParent().OnlyX(ctx).ID)
	lnk := client.Blob.Create().SetID(uuid.New()).AddLinks(chd, blb).SaveX(ctx)
	require.Equal(t, 2, lnk.QueryLinks().CountX(ctx))
	require.Equal(t, lnk.ID, chd.QueryLinks().OnlyX(ctx).ID)
	require.Equal(t, lnk.ID, blb.QueryLinks().OnlyX(ctx).ID)
	require.Len(t, client.Blob.Query().IDsX(ctx), 3)

	pedro := client.Pet.Create().SetID("pedro").SetOwner(a8m).SaveX(ctx)
	require.Equal(t, a8m.ID, pedro.QueryOwner().OnlyIDX(ctx))
	require.Equal(t, pedro.ID, a8m.QueryPets().OnlyIDX(ctx))
	xabi := client.Pet.Create().SetID("xabi").AddFriends(pedro).SetBestFriend(pedro).SaveX(ctx)
	require.Equal(t, "xabi", xabi.ID)
	pedro = client.Pet.Query().Where(pet.HasOwnerWith(user.ID(a8m.ID))).OnlyX(ctx)
	require.Equal(t, "pedro", pedro.ID)

	pets := client.Pet.Query().WithFriends().WithBestFriend().Order(ent.Asc(pet.FieldID)).AllX(ctx)
	require.Len(t, pets, 2)

	require.Equal(t, pedro.ID, pets[0].ID)
	require.NotNil(t, pets[0].Edges.BestFriend)
	require.Equal(t, xabi.ID, pets[0].Edges.BestFriend.ID)
	require.Len(t, pets[0].Edges.Friends, 1)
	require.Equal(t, xabi.ID, pets[0].Edges.Friends[0].ID)

	require.Equal(t, xabi.ID, pets[1].ID)
	require.NotNil(t, pets[1].Edges.BestFriend)
	require.Equal(t, pedro.ID, pets[1].Edges.BestFriend.ID)
	require.Len(t, pets[1].Edges.Friends, 1)
	require.Equal(t, pedro.ID, pets[1].Edges.Friends[0].ID)

	bee := client.Car.Create().SetModel("Chevrolet Camaro").SetOwner(pedro).SaveX(ctx)
	require.NotNil(t, bee)
	bee = client.Car.Query().WithOwner().OnlyX(ctx)
	require.Equal(t, "Chevrolet Camaro", bee.Model)
	require.NotNil(t, bee.Edges.Owner)
	require.Equal(t, pedro.ID, bee.Edges.Owner.ID)
}
