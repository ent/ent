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
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/customid/ent"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/ent/pet"
	"entgo.io/ent/entc/integration/customid/ent/token"
	"entgo.io/ent/entc/integration/customid/ent/user"
	"entgo.io/ent/entc/integration/customid/sid"
	"entgo.io/ent/schema/field"

	atlas "ariga.io/atlas/sql/schema"
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
			require.NoError(t, err, "connecting to custom_id database")
			err = client.Schema.Create(context.Background(), schema.WithHooks(clearDefault, skipBytesID), schema.WithAtlas(true))
			require.NoError(t, err)
			CustomID(t, client)
		})
	}
}

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5433, "13": 5434} {
		t.Run(version, func(t *testing.T) {
			dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable dbname=test", port)
			db, err := sql.Open(dialect.Postgres, dsn)
			require.NoError(t, err)
			defer db.Close()
			_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS custom_id")
			require.NoError(t, err, "creating schema")
			_, err = db.Exec("SET search_path TO custom_id")
			require.NoError(t, err, "setting schema")
			_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA custom_id`)
			require.NoError(t, err, "creating extension")
			defer db.Exec(`DROP EXTENSION "uuid-ossp"`)
			defer db.Exec("DROP SCHEMA custom_id CASCADE")

			client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))
			err = client.Schema.Create(context.Background(), schema.WithAtlas(true), schema.WithDiffHook(expectOnePetsIndex))
			require.NoError(t, err)
			CustomID(t, client)
			BytesID(t, client)
		})
	}
}

func TestSQLite(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(context.Background(), schema.WithHooks(clearDefault)), schema.WithAtlas(true))
	CustomID(t, client)
	BytesID(t, client)
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

	iSID := client.IntSID.Create().SaveX(ctx)
	require.Equal(t, sid.ID("1"), iSID.ID)
	iSIDChildID := sid.ID("100")
	iSIDChild := client.IntSID.Create().SetID(iSIDChildID).SetParent(iSID).SaveX(ctx)
	require.Equal(t, iSIDChildID, iSIDChild.ID)
	require.Equal(t, iSID.ID, iSIDChild.QueryParent().OnlyX(ctx).ID)
	iSIDBulk := make([]*ent.IntSIDCreate, 2)
	iSIDBulk[0] = client.IntSID.Create().SetParent(iSID)
	iSIDBulk[1] = client.IntSID.Create().SetParent(iSID)
	iSIDChildren := client.IntSID.CreateBulk(iSIDBulk...).SaveX(ctx)
	if strings.HasPrefix(t.Name(), "TestPostgres/") {
		require.Equal(t, sid.ID("2"), iSIDChildren[0].ID)
		require.Equal(t, sid.ID("3"), iSIDChildren[1].ID)
	} else {
		require.Equal(t, sid.ID("101"), iSIDChildren[0].ID)
		require.Equal(t, sid.ID("102"), iSIDChildren[1].ID)
	}
	iSIDArray := client.IntSID.Query().Where(intsid.ID(iSID.ID)).WithChildren().AllX(ctx)
	require.Equal(t, 1, len(iSIDArray))
	require.Equal(t, 3, len(iSIDArray[0].Edges.Children))

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

	pets = client.Pet.CreateBulk(
		client.Pet.Create().SetID("luna").SetOwner(a8m).AddFriends(xabi),
		client.Pet.Create().SetID("layla").SetOwner(a8m).AddFriendIDs(pedro.ID),
		client.Pet.Create().AddFriends(pedro, xabi),
	).SaveX(ctx)
	require.Equal(t, "luna", pets[0].ID)
	require.Equal(t, xabi.ID, pets[0].QueryFriends().OnlyIDX(ctx))
	require.Equal(t, "layla", pets[1].ID)
	require.Equal(t, pedro.ID, pets[1].QueryFriends().OnlyIDX(ctx))
	require.Equal(t, []string{"pedro", "xabi"}, pets[2].QueryFriends().Order(ent.Asc(pet.FieldID)).IDsX(ctx))

	u1, u2 := uuid.New(), uuid.New()
	blobs := client.Blob.CreateBulk(
		client.Blob.Create().SetID(u1),
		client.Blob.Create().SetID(u2),
	).SaveX(ctx)
	require.Equal(t, u1, blobs[0].ID)
	require.Equal(t, u2, blobs[1].ID)

	parent := client.Note.Create().SetText("parent").SaveX(ctx)
	require.NotEmpty(t, parent.ID)
	require.NotEmpty(t, parent.Text)
	child := client.Note.Create().SetText("child").SetParent(parent).SaveX(ctx)
	require.NotEmpty(t, child.QueryParent().OnlyIDX(ctx))

	pdoc := client.Doc.Create().SetText("parent").SaveX(ctx)
	require.NotEmpty(t, pdoc.ID)
	require.NotEmpty(t, pdoc.Text)
	cdoc := client.Doc.Create().SetText("child").SetParent(pdoc).SaveX(ctx)
	require.NotEmpty(t, cdoc.QueryParent().OnlyIDX(ctx))

	t.Run("Upsert", func(t *testing.T) {
		id := uuid.New()
		client.Blob.Create().
			SetID(id).
			OnConflictColumns(blob.FieldID).
			UpdateNewValues().
			ExecX(ctx)
		require.Zero(t, client.Blob.GetX(ctx, id).Count)
		client.Blob.Create().
			SetID(id).
			OnConflictColumns(blob.FieldID).
			Update(func(set *ent.BlobUpsert) {
				set.AddCount(1)
			}).
			ExecX(ctx)
		require.Equal(t, 1, client.Blob.GetX(ctx, id).Count)

		d := client.Doc.Create().SaveX(ctx)
		client.Doc.Create().
			SetID(d.ID).
			OnConflictColumns(doc.FieldID).
			SetText("Hello World").
			UpdateNewValues().
			ExecX(ctx)
		require.Equal(t, "Hello World", client.Doc.GetX(ctx, d.ID).Text)
	})

	t.Run("Other ID", func(t *testing.T) {
		o := client.Other.Create().SaveX(ctx)
		require.NotEmpty(t, o.ID.String())

		o = client.Other.Create().SetID(sid.NewLength(15)).SaveX(ctx)
		require.NotEmpty(t, o.ID.String())
	})

	t.Run("CustomID edge", func(t *testing.T) {
		a := client.Account.Create().SetEmail("test@example.org").SaveX(ctx)
		require.NotEmpty(t, a.ID)

		tk := client.Token.Create().SetAccountID(a.ID).SetBody("token").SaveX(ctx)
		require.NotEmpty(t, tk.ID)

		ta := client.Token.Query().Where(token.Body("token")).WithAccount().FirstX(ctx)
		require.Equal(t, tk.ID, ta.ID)
		require.NotNil(t, ta.Edges.Account)
		require.Equal(t, a.ID, ta.Edges.Account.ID)
	})
}

func BytesID(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	s := client.Session.Create().SaveX(ctx)
	require.NotEmpty(t, s.ID)
	client.Device.Create().SetActiveSession(s).AddSessionIDs(s.ID).SaveX(ctx)
	d := client.Device.Query().WithActiveSession().WithSessions().OnlyX(ctx)
	require.Equal(t, s.ID, d.Edges.ActiveSession.ID)
	require.Equal(t, s.ID, d.Edges.Sessions[0].ID)
}

// clearDefault clears the id's default for non-postgres dialects.
func clearDefault(c schema.Creator) schema.Creator {
	return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
		// Drop DEFAULT clause for MySQL without changing the tables.
		ct := make([]*schema.Table, len(tables))
		copy(ct, tables)
		*ct[1] = *tables[1]
		ct[1].Columns = append([]*schema.Column(nil), tables[1].Columns...)
		*ct[1].Columns[0] = *tables[1].Columns[0]
		ct[1].Columns[0].Default = nil
		return c.Create(ctx, ct...)
	})
}

// skipBytesID tables with blob ids from the migration.
func skipBytesID(c schema.Creator) schema.Creator {
	return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
		t := make([]*schema.Table, 0, len(tables))
		for i := range tables {
			if tables[i].PrimaryKey[0].Type == field.TypeBytes {
				continue
			}
			t = append(t, tables[i])
		}
		return c.Create(ctx, t...)
	})
}

// expectOnePetsIndex expects that pets table contains only one index.
func expectOnePetsIndex(next schema.Differ) schema.Differ {
	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
		changes, err := next.Diff(current, desired)
		for _, c := range changes {
			addT, ok := c.(*atlas.AddTable)
			if !ok || addT.T.Name != pet.Table {
				continue
			}
			if n := len(addT.T.Indexes); n != 1 {
				return nil, fmt.Errorf("expect only one index, but got: %d", n)
			}
		}
		return changes, err
	})
}
