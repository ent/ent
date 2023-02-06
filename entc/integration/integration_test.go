// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	stdsql "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	sqlschema "entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/enttest"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/entc/integration/ent/hook"
	"entgo.io/ent/entc/integration/ent/item"
	"entgo.io/ent/entc/integration/ent/license"
	"entgo.io/ent/entc/integration/ent/migrate"
	"entgo.io/ent/entc/integration/ent/node"
	"entgo.io/ent/entc/integration/ent/pet"
	"entgo.io/ent/entc/integration/ent/schema"
	"entgo.io/ent/entc/integration/ent/schema/task"
	enttask "entgo.io/ent/entc/integration/ent/task"
	"entgo.io/ent/entc/integration/ent/user"

	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSQLite(t *testing.T) {
	t.Parallel()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", opts)
	defer client.Close()
	for _, tt := range tests {
		name := runtime.FuncForPC(reflect.ValueOf(tt).Pointer()).Name()
		t.Run(name[strings.LastIndex(name, ".")+1:], func(t *testing.T) {
			drop(t, client)
			tt(t, client)
		})
	}
}

func TestMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			t.Parallel()
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr), opts)
			defer client.Close()
			for _, tt := range tests {
				name := runtime.FuncForPC(reflect.ValueOf(tt).Pointer()).Name()
				t.Run(name[strings.LastIndex(name, ".")+1:], func(t *testing.T) {
					drop(t, client)
					tt(t, client)
				})
			}
		})
	}
}

func TestMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			t.Parallel()
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr), opts)
			defer client.Close()
			for _, tt := range tests {
				name := runtime.FuncForPC(reflect.ValueOf(tt).Pointer()).Name()
				t.Run(name[strings.LastIndex(name, ".")+1:], func(t *testing.T) {
					drop(t, client)
					tt(t, client)
				})
			}
		})
	}
}

func TestPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434, "15": 5435} {
		addr := fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port)
		t.Run(version, func(t *testing.T) {
			t.Parallel()
			client := enttest.Open(t, dialect.Postgres, addr, opts)
			defer client.Close()
			for _, tt := range tests {
				name := runtime.FuncForPC(reflect.ValueOf(tt).Pointer()).Name()
				t.Run(name[strings.LastIndex(name, ".")+1:], func(t *testing.T) {
					drop(t, client)
					tt(t, client)
				})
			}
		})
	}
}

var (
	opts = enttest.WithMigrateOptions(
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	tests = [...]func(*testing.T, *ent.Client){
		Sanity,
		NoSchemaChanges,
		Tx,
		Lock,
		Indexes,
		Types,
		Clone,
		EntQL,
		Paging,
		Select,
		Aggregate,
		Delete,
		Upsert,
		Relation,
		ExecQuery,
		Predicate,
		AddValues,
		ClearEdges,
		ClearFields,
		UniqueConstraint,
		O2OTwoTypes,
		O2OSameType,
		O2OSelfRef,
		O2MTwoTypes,
		O2MSameType,
		M2MSelfRef,
		M2MSameType,
		M2MTwoTypes,
		DefaultValue,
		ImmutableValue,
		Sensitive,
		EagerLoading,
		NamedEagerLoading,
		Mutation,
		CreateBulk,
		ConstraintChecks,
		NillableRequired,
	}
)

func Sanity(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	usr := client.User.Create().SetName("foo").SetAge(20).SaveX(ctx)
	client.User.Update().ExecX(ctx)
	client.User.UpdateOne(usr).ExecX(ctx)
	require.Equal("foo", usr.Name)
	require.Equal(20, usr.Age)
	require.NotEmpty(usr.ID)
	client.User.Query().OnlyX(ctx)
	client.User.Delete().ExecX(ctx)
	require.Empty(client.User.Query().AllX(ctx))
	pt := client.Pet.Create().SetName("pedro").SaveX(ctx)
	usr = client.User.Create().SetName("foo").SetAge(20).AddPets(pt).SaveX(ctx)
	child := client.User.Create().SetName("bar").SetAge(20).AddChildren(usr).SaveX(ctx)
	inf := client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
	grp := client.Group.Create().SetName("Github").SetExpire(time.Now()).AddUsers(usr, child).SetInfo(inf).SaveX(ctx)
	require.Equal(1, client.Group.Query().CountX(ctx))
	require.Zero(client.Group.Query().Where(group.Active(false)).CountX(ctx))
	require.Len(grp.QueryUsers().AllX(ctx), 2)
	usr.QueryGroups().OnlyX(ctx)
	child.QueryGroups().OnlyX(ctx)
	usr2 := client.User.Create().SetName("qux").SetAge(20).SetSpouse(usr).SaveX(ctx)
	usr2.QuerySpouse().OnlyX(ctx)
	usr.QuerySpouse().OnlyX(ctx)
	require.Equal(usr.Name, usr.QueryPets().QueryOwner().OnlyX(ctx).Name)
	require.Equal(pt.Name, usr.QueryPets().QueryOwner().QueryPets().OnlyX(ctx).Name)
	require.Empty(usr.QuerySpouse().QueryPets().AllX(ctx))
	require.Equal(pt.Name, usr2.QuerySpouse().QueryPets().OnlyX(ctx).Name)
	require.Len(usr.QueryGroups().QueryUsers().AllX(ctx), 2)
	require.Len(usr.QueryGroups().QueryUsers().QueryGroups().AllX(ctx), 1, "should be unique by default")
	require.Len(usr.QueryGroups().AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.HasPets()).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.HasSpouse()).AllX(ctx), 2)
	require.Len(client.User.Query().Where(user.Not(user.HasSpouse())).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.HasGroups()).AllX(ctx), 2)
	require.Len(client.Group.Query().Where(group.HasUsers()).AllX(ctx), 1)
	require.Len(client.Group.Query().Where(group.HasUsersWith(user.Name("foo"))).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.HasGroupsWith(group.NameHasPrefix("G"))).AllX(ctx), 2)
	require.Equal(3, client.User.Query().CountX(ctx))
	require.Equal(client.Group.Query().Where(group.HasUsersWith(user.Name("foo"))).CountX(ctx), 1)
	require.True(client.User.Query().ExistX(ctx))
	require.True(client.User.Query().Where(user.HasPetsWith(pet.NameHasPrefix("ped"))).ExistX(ctx))
	require.False(client.User.Query().Where(user.HasPetsWith(pet.NameHasPrefix("pan"))).ExistX(ctx))
	require.Equal(child.Name, client.User.Query().Order(ent.Asc("name")).FirstX(ctx).Name)
	require.Equal(usr2.Name, client.User.Query().Order(ent.Desc("name")).FirstX(ctx).Name)
	// Update fields.
	client.User.Update().Where(user.ID(child.ID)).SetName("Ariel").SaveX(ctx)
	client.User.Query().Where(user.Name("Ariel")).OnlyX(ctx)
	// Update edges.
	require.Empty(child.QueryPets().AllX(ctx))
	require.NoError(client.Pet.UpdateOne(pt).ClearOwner().Exec(ctx))
	client.User.Update().Where(user.ID(child.ID)).AddPets(pt).SaveX(ctx)
	require.NotEmpty(child.QueryPets().AllX(ctx))
	client.User.Update().Where(user.ID(child.ID)).RemovePets(pt).SaveX(ctx)
	require.Empty(child.QueryPets().AllX(ctx))
	// Remove edges.
	client.User.Update().ClearSpouse().SaveX(ctx)
	require.Empty(client.User.Query().Where(user.HasSpouse()).AllX(ctx))
	client.User.Update().AddFriends(child).RemoveGroups(grp).Where(user.ID(usr.ID)).SaveX(ctx)
	require.NotEmpty(child.QueryGroups().AllX(ctx))
	require.Empty(usr.QueryGroups().AllX(ctx))
	require.Len(child.QueryFriends().AllX(ctx), 1)
	require.Len(usr.QueryFriends().AllX(ctx), 1)
	// Update one node.
	usr = client.User.UpdateOne(usr).SetName("baz").AddGroups(grp).SaveX(ctx)
	require.Equal("baz", usr.Name)
	require.NotEmpty(usr.QueryGroups().AllX(ctx))
	// Update unknown node.
	err := client.User.UpdateOneID(usr.ID + math.MaxInt8).SetName("foo").Exec(ctx)
	require.Error(err)
	require.True(ent.IsNotFound(err))
	// Update a vertex with filter.
	u := client.User.UpdateOneID(usr.ID).SetName("foo")
	u.Mutation().Where(user.Name(usr.Name))
	require.NoError(u.Exec(ctx))
	u = client.User.UpdateOneID(usr.ID).SetName("bar")
	u.Mutation().Where(user.Name("baz"))
	require.Error(u.Exec(ctx))
	require.True(ent.IsNotFound(err))

	var v []struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Sum   int    `json:"sum"`
		Count int    `json:"count"`
	}
	client.User.Query().
		GroupBy(user.FieldName, user.FieldAge).
		Aggregate(ent.Count(), ent.Sum(user.FieldAge)).
		ScanX(ctx, &v)
	require.NotEmpty(v)
	// IN predicates.
	ids := client.User.Query().IDsX(ctx)
	require.Len(ids, 3)
	client.User.Delete().Where(user.IDIn(ids...)).ExecX(ctx)
	ids = client.User.Query().IDsX(ctx)
	require.Empty(ids)
	// Nop.
	client.User.Delete().Where(user.IDIn(ids...)).ExecX(ctx)
	// Check the struct-tag annotation.
	fi, ok := reflect.TypeOf(ent.Card{}).FieldByName("Edges")
	require.True(ok)
	require.NotEmpty(fi.Tag.Get("mashraki"))
	fi, ok = reflect.TypeOf(ent.Card{}).FieldByName("ID")
	require.True(ok)
	require.Equal("-", fi.Tag.Get("json"))
	fi, ok = reflect.TypeOf(ent.Card{}).FieldByName("Number")
	require.True(ok)
	require.Equal("-", fi.Tag.Get("json"))
	client.User.Create().SetName("tarrence").SetAge(30).ExecX(ctx)

	t.Run("StringPredicates", func(t *testing.T) {
		client.Pet.Delete().ExecX(ctx)
		a := client.Pet.Create().SetName("a%").SaveX(ctx)
		require.True(client.Pet.Query().Where(pet.NameHasPrefix("a%")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameHasPrefix("%a%")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.Or(pet.NameHasPrefix("%a%"), pet.NameHasPrefix("%a%"))).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameHasSuffix("%")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameHasSuffix("a%%")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameContains("a")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameContains("a%")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameContains("%a%")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameContainsFold("A%")).ExistX(ctx))

		a.Update().SetName("a_\\").ExecX(ctx)
		require.True(client.Pet.Query().Where(pet.NameHasPrefix("a")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameHasPrefix("%a")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameHasPrefix("a_")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameHasSuffix("a_\\")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameHasSuffix("%a")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameHasSuffix("a%")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameContains("a")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameContains("%a")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameContains("a%")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameContainsFold("A")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameContainsFold("%A")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameContainsFold("A%")).ExistX(ctx))
		require.True(client.Pet.Query().Where(pet.NameEqualFold("A_\\")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameEqualFold("%A_\\")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameEqualFold("A_\\%")).ExistX(ctx))
		require.False(client.Pet.Query().Where(pet.NameEqualFold("A%")).ExistX(ctx))
	})
}

func Upsert(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	u := client.User.Create().SetName("Ariel").SetAge(30).SetPhone("0000").SaveX(ctx)
	require.Equal(t, "static", u.Address, "address was set by default func")
	err := client.User.Create().SetName("Mashraki").SetAge(30).SetPhone("0000").Exec(ctx)
	require.True(t, ent.IsConstraintError(err), "phone field is unique")
	err = client.User.Create().SetName("Mashraki").SetAge(30).SetPhone("0000").OnConflict().Exec(ctx)
	require.EqualError(t, err, "ent: missing options for UserCreate.OnConflict")

	client.User.Create().
		SetName("Mashraki").
		SetAge(30).
		SetPhone("0000").
		OnConflict(
			sql.ConflictColumns(user.FieldPhone),
		).
		// Update "name" to the value that was set on create ("Mashraki").
		UpdateName().
		ExecX(ctx)
	u = client.User.GetX(ctx, u.ID)
	require.Equal(t, "Mashraki", u.Name, "name was changed by the UPDATE clause")

	id := client.User.Create().
		SetName("Boring").
		SetAge(33).
		SetPhone("0000").
		OnConflictColumns(user.FieldPhone).
		// Override some fields with custom update.
		Update(func(u *ent.UserUpsert) {
			// Age was set to the new value (33).
			u.UpdateAge()
			// Update an additional field that was defined in `VALUES`.
			u.SetAddress("localhost")
		}).
		IDX(ctx)
	require.Equal(t, u.ID, id)
	u = client.User.GetX(ctx, u.ID)
	require.Equal(t, "Mashraki", u.Name)
	require.Equal(t, 33, u.Age, "age was modified by the UPDATE clause")
	require.Equal(t, "localhost", u.Address, "address was modified by the UPDATE clause")

	id = client.User.Create().
		SetName("Boring").
		SetAge(33).
		SetPhone("0000").
		OnConflictColumns(user.FieldPhone).
		// Override some fields with custom update.
		AddAge(-1).
		IDX(ctx)
	u = client.User.GetX(ctx, id)
	require.Equal(t, 32, u.Age, "age was modified by the UPDATE clause")

	builders := []*ent.UserCreate{
		client.User.Create().SetName("A").SetAge(1).SetPhone("0000"), // Duplicate
		client.User.Create().SetName("B").SetAge(1).SetPhone("1111"), // New row.
	}
	client.User.CreateBulk(builders...).
		OnConflictColumns(user.FieldPhone).
		UpdateNewValues().
		ExecX(ctx)
	users := client.User.Query().Order(ent.Asc(user.FieldPhone)).AllX(ctx)
	require.Equal(t, "0000", users[0].Phone)
	require.Equal(t, "A", users[0].Name)
	require.Equal(t, "1111", users[1].Phone)
	require.Equal(t, "B", users[1].Name)

	// Setting primary key manually.
	a := client.Item.Create().SetID("A").SaveX(ctx)
	require.Equal(t, "A", a.ID)
	if strings.Contains(t.Name(), "MySQL") || strings.Contains(t.Name(), "Maria") {
		// MySQL is skipped since it does not support the RETURNING clause. Maria is skipped
		// as well, because there's no way to distinguish between MySQL and Maria at runtime.
		client.Item.Create().SetID("A").OnConflict().Ignore().ExecX(ctx)
		require.Equal(t, 1, client.Item.Query().CountX(ctx))
		client.Item.Delete().ExecX(ctx)

		// Primary key is set by a default function.
		b := client.Item.Create().SetText("hello").SaveX(ctx)
		require.NotZero(t, b.ID)
		client.Item.Create().SetID(b.ID).SetText("world").OnConflict().UpdateNewValues().ExecX(ctx)
		cb := client.Item.Query().OnlyX(ctx)
		require.Equal(t, cb.ID, b.ID)
		require.Equal(t, "world", cb.Text)
	} else {
		aid := client.Item.Create().SetID("A").OnConflict(sql.ConflictColumns(item.FieldID)).Ignore().IDX(ctx)
		require.Equal(t, a.ID, aid)
		client.Item.Delete().ExecX(ctx)

		// Primary key is set by a default function.
		b := client.Item.Create().SetText("hello").SaveX(ctx)
		require.NotZero(t, b.ID)
		bid := client.Item.Create().SetID(b.ID).SetText("hello").OnConflictColumns(item.FieldText).Ignore().IDX(ctx)
		require.Equal(t, b.ID, bid)
		bid = client.Item.Create().SetText("hello").OnConflictColumns(item.FieldText).UpdateNewValues().IDX(ctx)
		require.Equal(t, bid, b.ID)
		require.Equal(t, bid, client.Item.Query().OnlyIDX(ctx))
		bid = client.Item.Create().SetID(bid).SetText("world").OnConflictColumns(item.FieldID).UpdateNewValues().IDX(ctx)
		require.Equal(t, bid, b.ID)
		b = client.Item.Query().OnlyX(ctx)
		require.Equal(t, bid, b.ID)
		require.Equal(t, "world", b.Text)

		client.Item.CreateBulk(client.Item.Create().SetID(bid).SetText("hello")).
			OnConflictColumns(item.FieldID).
			Ignore().
			ExecX(ctx)
		require.Equal(t, bid, client.Item.Query().OnlyIDX(ctx))
	}

	ts := time.Unix(1623279251, 0)
	c1 := client.Card.Create().
		SetNumber("102030").
		SetCreateTime(ts).
		SetUpdateTime(ts).
		SaveX(ctx)

	// "DO UPDATE SET ... WHERE ..." does not support by MySQL.
	if strings.Contains(t.Name(), "Postgres") || strings.Contains(t.Name(), "SQLite") {
		err = client.Card.Create().
			SetNumber(c1.Number).
			OnConflict(
				sql.ConflictColumns(card.FieldNumber),
				sql.UpdateWhere(sql.NEQ(card.FieldCreateTime, ts)),
			).
			UpdateNewValues().
			Exec(ctx)
		// Only rows for which the "UpdateWhere" expression
		// returns true will be updated. That is, none.
		require.True(t, errors.Is(err, stdsql.ErrNoRows))

		id = client.Card.Create().
			SetNumber(c1.Number).
			OnConflict(
				sql.ConflictColumns(card.FieldNumber),
				sql.UpdateWhere(sql.EQ(card.FieldCreateTime, ts)),
			).
			UpdateNewValues().
			IDX(ctx)
	} else {
		id = client.Card.Create().
			SetNumber(c1.Number).
			OnConflictColumns(card.FieldNumber).
			UpdateNewValues().
			IDX(ctx)
	}

	// Ensure immutable fields were not changed during upsert.
	c2 := client.Card.GetX(ctx, id)
	require.Equal(t, c1.CreateTime.Unix(), c2.CreateTime.Unix())
	require.NotEqual(t, c1.UpdateTime.Unix(), c2.UpdateTime.Unix())

	// Ensure immutable fields were not changed during bulk upsert.
	l1 := client.License.Create().SetCreateTime(ts).SetUpdateTime(ts).SaveX(ctx)
	client.License.CreateBulk(client.License.Create().SetID(l1.ID)).
		OnConflictColumns(license.FieldID).
		UpdateNewValues().
		ExecX(ctx)
	l2 := client.License.GetX(ctx, l1.ID)
	require.Equal(t, l1.CreateTime.Unix(), l2.CreateTime.Unix())
	require.NotEqual(t, l1.UpdateTime.Unix(), l2.UpdateTime.Unix())

	c3 := client.Card.Create().SetName("a8m").SetNumber("405060").SaveX(ctx)
	client.Card.Create().SetNumber(c3.Number).OnConflictColumns(card.FieldNumber).ClearName().UpdateNewValues().ExecX(ctx)
	require.Empty(t, client.Card.GetX(ctx, c3.ID).Name)
	c3.Update().SetName("a8m").ExecX(ctx)
	client.Card.CreateBulk(client.Card.Create().SetNumber(c3.Number), client.Card.Create().SetNumber("708090").SetName("m8a")).
		OnConflictColumns(card.FieldNumber).
		UpdateNewValues().
		ExecX(ctx)
	require.Empty(t, client.Card.GetX(ctx, c3.ID).Name, "existing name fields should be cleared when not set (= set to nil)")
	require.NotEmpty(t, client.Card.Query().Where(card.Number("708090")).OnlyX(ctx).Name, "new record should set their name")

	// Conflict on a composite unique index.
	t1 := client.Task.Create().SetName("todo1").SetOwner("a8m").SetPriority(task.PriorityLow).SaveX(ctx)
	tid := client.Task.Create().
		SetName("todo1").
		SetOwner("a8m").
		SetPriority(task.PriorityHigh).
		OnConflictColumns(enttask.FieldName, enttask.FieldOwner).
		UpdatePriority().
		IDX(ctx)
	require.Equal(t, t1.ID, tid)
	require.Equal(t, task.PriorityHigh, client.Task.GetX(ctx, tid).Priority)
}

func Clone(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	f1 := client.File.Create().SetName("foo").SetSize(10).SaveX(ctx)
	f2 := client.File.Create().SetName("foo").SetSize(20).SaveX(ctx)
	base := client.File.Query().Where(file.Name("foo"))
	require.Equal(t, f1.Size, base.Clone().Where(file.Size(f1.Size)).OnlyX(ctx).Size)
	require.Equal(t, f2.Size, base.Clone().Where(file.Size(f2.Size)).OnlyX(ctx).Size)
	// ensure clone emits valid code.
	query := client.Pet.Query().Where(pet.Name("unknown")).QueryTeam()
	for i := 0; i < 10; i++ {
		_, err := query.Clone().Where(user.Name("unknown")).First(ctx)
		require.True(t, ent.IsNotFound(err), "should not return syntax error")
	}
}

func Paging(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		client.User.Create().SetName(fmt.Sprintf("name-%d", i)).SetAge(i).SaveX(ctx)
	}

	require.Equal(10, client.User.Query().CountX(ctx))
	require.Len(client.User.Query().Offset(5).AllX(ctx), 5)
	require.Len(client.User.Query().Offset(6).AllX(ctx), 4)
	require.Equal(
		[]int{7, 8},
		client.User.Query().
			Offset(6).
			Limit(2).
			Order(ent.Asc(user.FieldAge)).
			GroupBy(user.FieldAge).
			IntsX(ctx),
	)
	for i := 0; i < 10; i++ {
		require.Equal(i+1, client.User.Query().Order(ent.Asc(user.FieldAge)).Offset(i).Limit(1).AllX(ctx)[0].Age)
	}
}

func Select(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)

	t.Log("select one field")
	u := client.User.Create().SetName("foo").SetAge(30).SaveX(ctx)
	name := client.User.
		Query().
		Where(user.ID(u.ID)).
		Select(user.FieldName).
		StringX(ctx)
	require.Equal("foo", name)
	client.User.Create().SetName("bar").SetAge(30).AddFriends(u).SaveX(ctx)
	t.Log("select one field with ordering")
	names := client.User.
		Query().
		Order(ent.Asc(user.FieldName)).
		Select(user.FieldName).
		StringsX(ctx)
	require.Equal([]string{"bar", "foo"}, names)
	names = client.User.
		Query().
		Order(ent.Desc(user.FieldName)).
		Select(user.FieldName).
		StringsX(ctx)
	require.Equal([]string{"foo", "bar"}, names)
	client.User.Create().SetName("baz").SetAge(30).SaveX(ctx)
	names = client.User.
		Query().
		Order(ent.Asc(user.FieldName)).
		Select(user.FieldName).
		StringsX(ctx)
	require.Equal([]string{"bar", "baz", "foo"}, names)

	t.Log("select 2 fields")
	var v []struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}
	client.User.
		Query().
		Order(ent.Asc(user.FieldName)).
		Select(user.FieldAge, user.FieldName).
		ScanX(ctx, &v)
	require.Equal([]int{30, 30, 30}, []int{v[0].Age, v[1].Age, v[2].Age})
	require.Equal([]string{"bar", "baz", "foo"}, []string{v[0].Name, v[1].Name, v[2].Name})

	users := client.User.
		Query().
		Select(user.FieldAge).
		Where(user.Name("foo")).
		WithFriends(func(q *ent.UserQuery) {
			q.Select(user.FieldName)
		}).
		AllX(ctx)
	for i := range users {
		require.Empty(users[i].Name)
		require.NotZero(users[i].ID)
		require.NotZero(users[i].Age)
		for _, f := range users[i].Edges.Friends {
			require.NotEmpty(f.Name)
			require.NotZero(f.ID)
			require.Zero(f.Age)
		}
	}
	a8m := client.User.Create().SetName("Ariel").SetNickname("a8m").SetAge(30).SaveX(ctx)
	require.NotEmpty(a8m.ID)
	require.NotEmpty(a8m.Age)
	require.NotEmpty(a8m.Name)
	require.NotEmpty(a8m.Nickname)
	a8m = a8m.Update().SetAge(32).Select(user.FieldAge).SaveX(ctx)
	require.NotEmpty(a8m.ID)
	require.NotEmpty(a8m.Age)
	require.Empty(a8m.Name)
	require.Empty(a8m.Nickname)

	client.Pet.CreateBulk(
		client.Pet.Create().SetName("a"),
		client.Pet.Create().SetName("a"),
	).ExecX(ctx)
	names = client.Pet.Query().Select(pet.FieldName).StringsX(ctx)
	require.Equal([]string{"a", "a"}, names)
	names = client.Pet.Query().Unique(true).Select(pet.FieldName).StringsX(ctx)
	require.Equal([]string{"a"}, names)
	client.Pet.Delete().ExecX(ctx)

	pets := client.Pet.CreateBulk(
		client.Pet.Create().SetName("a"),
		client.Pet.Create().SetName("b"),
		client.Pet.Create().SetName("c"),
		client.Pet.Create().SetName("b"),
	).SaveX(ctx)
	client.User.Create().SetName("foo").SetAge(20).AddPets(pets[0], pets[1]).SaveX(ctx)
	client.User.Create().SetName("bar").SetAge(20).AddPets(pets[2], pets[3]).SaveX(ctx)
	names = client.Pet.Query().Order(ent.Asc(pet.FieldID)).Select(pet.FieldName).StringsX(ctx)
	require.Equal([]string{"a", "b", "c", "b"}, names)
	names = client.Pet.Query().Order(ent.Asc(pet.FieldName)).Select(pet.FieldName).StringsX(ctx)
	require.Equal([]string{"a", "b", "b", "c"}, names)
	names = client.Pet.Query().
		Order(func(s *sql.Selector) {
			// Join with user table for ordering by owner-name
			// and pet-name (edge + field ordering).
			t := sql.Table(user.Table)
			s.Join(t).On(s.C(pet.OwnerColumn), t.C(user.FieldID))
			s.OrderBy(t.C(user.FieldName), s.C(pet.FieldName))
		}).
		Select(pet.FieldName).
		StringsX(ctx)
	require.Equal([]string{"b", "c", "a", "b"}, names)

	var ps []*ent.Pet
	client.Pet.Query().Select().ScanX(ctx, &ps)
	require.Len(ps, 4, "support scanning nodes manually")

	lens := client.Pet.Query().
		Modify(func(s *sql.Selector) {
			s.Select("LENGTH(name)")
		}).
		IntsX(ctx)
	require.Equal([]int{1, 1, 1, 1}, lens)

	dlen := client.Pet.Query().
		Modify(func(s *sql.Selector) {
			s.SelectExpr(sql.ExprFunc(func(b *sql.Builder) {
				b.WriteString("LENGTH(name)").WriteOp(sql.OpMul).Arg(2)
			}))
		}).
		IntsX(ctx)
	require.Equal([]int{2, 2, 2, 2}, dlen)

	for i := range pets {
		pets[i].Update().SetName(pets[i].Name + pets[i].Name).ExecX(ctx)
	}
	n := client.Pet.Query().
		Modify(func(s *sql.Selector) {
			s.Select("SUM(LENGTH(name))")
		}).
		IntX(ctx)
	require.Equal(8, n)

	var (
		p1 []struct {
			ent.Pet
			NameLength int `sql:"length"`
		}
		p2 = client.Pet.Query().Order(ent.Asc(pet.FieldID)).AllX(ctx)
	)
	client.Pet.Query().
		Order(ent.Asc(pet.FieldID)).
		Modify(func(s *sql.Selector) {
			s.AppendSelect("LENGTH(name)")
		}).
		ScanX(ctx, &p1)
	for i := range p2 {
		require.Equal(p2[i].ID, p1[i].ID)
		require.Equal(p2[i].Age, p1[i].Age)
		require.Equal(p2[i].Name, p1[i].Name)
		require.Equal(len(p1[i].Name), p1[1].NameLength)
	}

	// Select count.
	names = client.Pet.Query().Order(ent.Asc(pet.FieldName)).Select(pet.FieldName).StringsX(ctx)
	require.Equal([]string{"aa", "bb", "bb", "cc"}, names)
	count := client.Pet.Query().Select(pet.FieldName).CountX(ctx)
	require.Equal(4, count)
	count = client.Pet.Query().Unique(true).Select(pet.FieldName).CountX(ctx)
	require.Equal(3, count)

	var (
		gs []struct {
			ent.Group
			UsersCount int `sql:"users_count"`
		}
		inf = client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
		hub = client.Group.Create().SetName("GitHub").SetExpire(time.Now()).SetInfo(inf).AddUsers(a8m).SaveX(ctx)
		lab = client.Group.Create().SetName("GitLab").SetExpire(time.Now()).SetInfo(inf).AddUsers(users...).SaveX(ctx)
	)
	client.Group.Query().
		Order(ent.Asc(group.FieldID)).
		Modify(func(s *sql.Selector) {
			t := sql.Table(group.UsersTable)
			s.LeftJoin(t).
				On(
					s.C(group.FieldID),
					t.C(group.UsersPrimaryKey[1]),
				).
				// Append the "users_count" column to the selected columns.
				AppendSelect(
					sql.As(sql.Count(t.C(group.UsersPrimaryKey[1])), "users_count"),
				).
				GroupBy(s.C(group.FieldID))
		}).
		ScanX(ctx, &gs)
	require.Len(gs, 2)
	require.Equal(hub.QueryUsers().CountX(ctx), gs[0].UsersCount)
	require.Equal(lab.QueryUsers().CountX(ctx), gs[1].UsersCount)

	// Select Subquery.
	t.Log("select subquery")
	i, err := client.User.
		Query().
		Modify(func(s *sql.Selector) {
			subQuery := sql.SelectExpr(sql.Raw("1")).As("s")
			s.Select("*").From(subQuery)
		}).
		Int(ctx)
	require.NoError(err)
	require.Equal(1, i)

	// Select with join.
	u = client.User.Create().SetName("crossworth").SetAge(28).SaveX(ctx)
	id := client.User.
		Query().
		Where(func(s *sql.Selector) {
			subQuery := sql.Select(user.FieldID).
				From(sql.Table(user.Table)).
				Where(sql.EQ(s.C(user.FieldName), "crossworth"))
			s.Join(subQuery).On(s.C(user.FieldID), subQuery.C(user.FieldID))
		}).
		OnlyIDX(ctx)
	require.Equal(u.ID, id)

	// Update modifiers.
	allUpper := func() bool {
		for _, name := range client.User.Query().Select(user.FieldName).StringsX(ctx) {
			if strings.ToUpper(name) != name {
				return false
			}
		}
		return true
	}
	require.False(allUpper(), "at least one name is not upper-cased")
	// Execute custom update modifier.
	client.User.Update().
		Modify(func(u *sql.UpdateBuilder) {
			u.Set(user.FieldName, sql.Expr(fmt.Sprintf("UPPER(%s)", user.FieldName)))
		}).
		ExecX(ctx)
	require.True(allUpper(), "at names must be upper-cased")
}

func Aggregate(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	a8m := client.User.Create().SetAge(1).SetName("a8m").SaveX(ctx)
	nat := client.User.Create().SetAge(1).SetName("nati").SetSpouse(a8m).SaveX(ctx)
	owners := []*ent.User{a8m, nat}
	for i := 1; i <= 10; i++ {
		client.Pet.Create().SetName(fmt.Sprintf("pet%d", i)).SetAge(float64(i)).SetOwner(owners[i%2]).SaveX(ctx)
	}
	s1 := client.Pet.Query().Aggregate(ent.Sum(pet.FieldAge)).IntX(ctx)
	require.Equal(t, 55, s1)
	s2 := client.Pet.Query().Where(pet.HasOwner()).Aggregate(ent.Sum(pet.FieldAge)).IntX(ctx)
	require.Equal(t, s1, s2)

	// Aggregate traversals.
	require.Equal(t, 30, a8m.QueryPets().Aggregate(ent.Sum(pet.FieldAge)).IntX(ctx))
	require.Equal(t, 25, nat.QueryPets().Aggregate(ent.Sum(pet.FieldAge)).IntX(ctx))
	require.Equal(t, 25, a8m.QuerySpouse().QueryPets().Aggregate(ent.Sum(pet.FieldAge)).IntX(ctx))
	require.Equal(t, 30, nat.QuerySpouse().QueryPets().Aggregate(ent.Sum(pet.FieldAge)).IntX(ctx))

	// Aggregate 2 fields.
	var vs1 []struct{ Sum, Count int }
	client.Pet.Query().
		Aggregate(
			ent.Sum(pet.FieldAge),
			ent.Count(),
		).
		ScanX(ctx, &vs1)
	require.Len(t, vs1, 1)
	require.Equal(t, 55, vs1[0].Sum)
	require.Equal(t, 10, vs1[0].Count)

	// Aggregate 4 fields.
	var vs2 []struct {
		Sum, Min, Max, Count int
		Avg                  float64
	}
	client.Pet.Query().
		Aggregate(
			ent.Sum(pet.FieldAge),
			ent.Min(pet.FieldAge),
			ent.Max(pet.FieldAge),
			ent.Mean(pet.FieldAge),
			ent.Count(),
		).
		ScanX(ctx, &vs2)
	require.Len(t, vs2, 1)
	require.Equal(t, 55, vs2[0].Sum)
	require.Equal(t, 1, vs2[0].Min)
	require.Equal(t, 10, vs2[0].Max)
	require.Equal(t, 10, vs2[0].Count)
	require.Equal(t, 5.5, vs2[0].Avg)
}

func ExecQuery(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	rows, err := client.QueryContext(ctx, "SELECT 1")
	require.NoError(err)
	require.True(rows.Next())
	require.NoError(rows.Close())
	tx, err := client.Tx(ctx)
	require.NoError(err)
	tx.Task.Create().ExecX(ctx)
	require.Equal(1, tx.Task.Query().CountX(ctx))
	rows, err = tx.QueryContext(ctx, "SELECT COUNT(*) FROM "+enttask.Table)
	require.NoError(err)
	count, err := sql.ScanInt(rows)
	require.NoError(err)
	require.NoError(rows.Close())
	require.Equal(1, count)
	require.NoError(tx.Commit())
}

func NillableRequired(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	client.Task.Create().ExecX(ctx)
	tk := client.Task.Query().OnlyX(ctx)
	require.NotNil(tk.CreatedAt, "field value should be populated by default by the database")
	require.False(reflect.ValueOf(tk.Update()).MethodByName("SetNillableCreatedAt").IsValid(), "immutable-nillable should not have SetNillable setter on update")
	tk = client.Task.Query().Select(enttask.FieldID, enttask.FieldPriority).OnlyX(ctx)
	require.Nil(tk.CreatedAt, "field should not be populated when it is not selected")
}

func Predicate(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	f1 := client.File.Create().SetName("1").SetSize(10).SaveX(ctx)
	f2 := client.File.Create().SetName("2").SetSize(20).SaveX(ctx)
	f3 := client.File.Create().SetName("3").SetSize(30).SaveX(ctx)
	f4 := client.File.Create().SetName("4").SetSize(40).SaveX(ctx)
	files := client.File.Query().
		Where(
			file.Or(
				file.Name(f1.Name),
				file.And(file.Name(f2.Name), file.Size(f2.Size)),
			),
		).
		Order(ent.Asc(file.FieldName)).
		AllX(ctx)
	require.Equal(f1.Name, files[0].Name)
	require.Equal(f2.Name, files[1].Name)

	match := client.File.Query().
		Where(file.Or(file.Name(f1.Name), file.Name(f2.Name))).
		Where(file.Size(f1.Size)).
		OnlyX(ctx)
	require.Equal(f1.Name, match.Name)

	match = client.File.Query().
		Where(file.Size(f2.Size)).
		Where(file.Or(file.Name(f1.Name), file.Name(f2.Name))).
		OnlyX(ctx)
	require.Equal(f2.Name, match.Name)

	files = client.File.Query().
		Where(file.Or(file.Size(f3.Size), file.Size(f4.Size))).
		Where(file.Or(file.Name(f3.Name), file.Name(f4.Name))).
		Where(file.Not(file.Or(file.Name(f1.Name), file.Size(f1.Size)))).
		Order(ent.Asc(file.FieldName)).
		AllX(ctx)
	require.Equal(f3.Name, files[0].Name)
	require.Equal(f4.Name, files[1].Name)

	files = client.File.Query().
		Where(
			file.Or(
				file.Name(f4.Name),
				file.And(file.Name(f3.Name), file.Size(f3.Size)),
			),
		).
		Order(ent.Asc(file.FieldName)).
		AllX(ctx)
	require.Equal(f3.Name, files[0].Name)
	require.Equal(f4.Name, files[1].Name)

	require.Zero(client.File.Query().Where(file.UserNotNil()).CountX(ctx))
	require.Equal(4, client.File.Query().Where(file.UserIsNil()).CountX(ctx))
	require.Zero(client.File.Query().Where(file.GroupNotNil()).CountX(ctx))
	require.Equal(4, client.File.Query().Where(file.GroupIsNil()).CountX(ctx))

	f1 = f1.Update().SetUser("a8m").SaveX(ctx)
	require.NotNil(f1.User)
	require.Equal("a8m", *f1.User)
	require.Equal(3, client.File.Query().Where(file.UserIsNil()).CountX(ctx))
	require.Equal(f1.Name, client.File.Query().Where(file.UserNotNil()).OnlyX(ctx).Name)
	f5 := client.File.Create().SetName("5").SetSize(40).SetUser("mashraki").SaveX(ctx)
	require.NotNil(f5.User)
	require.Equal("mashraki", *f5.User)
	require.Equal(3, client.File.Query().Where(file.UserIsNil()).CountX(ctx))
	require.Equal(2, client.File.Query().Where(file.UserNotNil()).CountX(ctx))

	require.Equal(5, client.File.Query().Where(file.GroupIsNil()).CountX(ctx))
	f4 = f4.Update().SetGroup("fbc").SaveX(ctx)
	require.Equal(1, client.File.Query().Where(file.GroupNotNil()).CountX(ctx))
	require.Equal(4, client.File.Query().Where(file.GroupIsNil()).CountX(ctx))
	require.Equal(
		5,
		client.File.Query().
			Where(
				file.Or(
					file.GroupIsNil(),
					file.And(
						file.GroupNotNil(),
						file.Name(f4.Name),
					),
				),
			).
			CountX(ctx),
	)

	inf := client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
	hub := client.Group.Create().SetName("GitHub").SetExpire(time.Now()).SetInfo(inf).SaveX(ctx)
	lab := client.Group.Create().SetName("GitLab").SetExpire(time.Now()).SetInfo(inf).SetActive(false).SaveX(ctx)
	require.Equal(hub.ID, client.Group.Query().Where(group.Active(true)).OnlyIDX(ctx))
	require.Equal(lab.ID, client.Group.Query().Where(group.Active(false)).OnlyIDX(ctx))
	require.Equal(hub.ID, client.Group.Query().Where(group.ActiveNEQ(false)).OnlyIDX(ctx))
	require.Equal(lab.ID, client.Group.Query().Where(group.ActiveNEQ(true)).OnlyIDX(ctx))
}

func AddValues(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	t.Log("add values to fields")
	cmt := client.Comment.Create().SetUniqueInt(1).SetUniqueFloat(1).SaveX(ctx)
	cmt = cmt.Update().AddUniqueInt(10).SaveX(ctx)
	require.Equal(11, cmt.UniqueInt)
	require.Equal(11, client.Comment.Query().OnlyX(ctx).UniqueInt, "should be updated in the database")
	t.Log("add values to null fields")
	cmt = cmt.Update().AddNillableInt(10).SaveX(ctx)
	require.Equal(10, *cmt.NillableInt)

	cmt1 := client.Comment.Create().SetUniqueInt(1).SetUniqueFloat(10).SaveX(ctx)
	err := cmt1.Update().AddUniqueInt(10).Exec(ctx)
	require.True(ent.IsConstraintError(err))
	cmt1 = cmt1.Update().AddUniqueInt(20).AddNillableInt(20).SaveX(ctx)
	require.Equal(21, cmt1.UniqueInt)
	require.Equal(20, *cmt1.NillableInt)

	cmt1 = cmt1.Update().AddUniqueInt(10).AddUniqueInt(-1).SaveX(ctx)
	require.Equal(30, cmt1.UniqueInt)
	require.Equal(30, client.Comment.GetX(ctx, cmt1.ID).UniqueInt)
}

func Delete(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()

	nd := client.Node.Create().SetValue(1e3).SaveX(ctx)
	err := client.Node.DeleteOneID(nd.ID).Exec(ctx)
	require.NoError(err)
	err = client.Node.DeleteOneID(nd.ID).Exec(ctx)
	require.True(ent.IsNotFound(err))

	for i := 0; i < 5; i++ {
		client.Node.Create().SetValue(i).ExecX(ctx)
	}
	affected, err := client.Node.Delete().Where(node.ValueGT(2)).Exec(ctx)
	require.NoError(err)
	require.Equal(2, affected)

	affected, err = client.Node.Delete().Exec(ctx)
	require.NoError(err)
	require.Equal(3, affected)

	info := client.GroupInfo.Create().SetDesc("group info").SaveX(ctx)
	hub := client.Group.Create().SetInfo(info).SetName("GitHub").SetExpire(time.Now().Add(time.Hour)).SaveX(ctx)
	err = client.GroupInfo.DeleteOne(info).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	// Group.DeleteOneID(id).Where(...), is identical to Group.Delete().Where(group.ID(id), ...),
	// but, in case the OpDelete is not an allowed operation, the DeleteOne can be used with Where.
	n, err := client.Group.Delete().
		Where(
			group.ID(hub.ID),
			group.ExpireLT(time.Now()), // Expired.
		).Exec(ctx)
	require.Zero(n)
	require.NoError(err)

	err = client.Group.DeleteOne(hub).
		Where(group.ExpireLT(time.Now())).
		Exec(ctx)
	require.True(ent.IsNotFound(err))
	hub.Update().SetExpire(time.Now().Add(-time.Hour)).ExecX(ctx)
	client.Group.DeleteOne(hub).
		Where(group.ExpireLT(time.Now())).
		ExecX(ctx)

	// The behavior described above it also applied to UpdateOne.
	hub = client.Group.Create().SetInfo(info).SetName("GitHub").SetExpire(time.Now().Add(time.Hour)).SaveX(ctx)
	err = hub.Update().
		SetActive(false).
		SetExpire(time.Time{}).
		Where(group.ExpireLT(time.Now())). // Expired.
		Exec(ctx)
	require.True(ent.IsNotFound(err))
}

func Relation(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	t.Log("querying group info")
	info, err := client.GroupInfo.
		Query().
		First(ctx)
	require.Nil(info)
	require.True(ent.IsNotFound(err))

	t.Log("creating group info")
	info = client.GroupInfo.
		Create().
		SetDesc("group info").
		SaveX(ctx)
	t.Logf("group info created: %v", info)

	t.Log("creating group")
	grp := client.Group.
		Create().
		SetInfo(info).
		SetName("Github").
		SetExpire(time.Now().Add(time.Hour)).
		SaveX(ctx)
	require.NotZero(grp.ID)
	require.Equal(grp.MaxUsers, 10)
	require.Equal(grp.Name, "Github")
	t.Logf("group created: %v", grp)

	t.Log("creating user")
	usr := client.User.
		Create().
		SetAge(20).
		SetName("a8m").
		AddGroups(grp).
		SaveX(ctx)
	require.NotZero(usr.ID)
	require.Equal(usr.Age, 20)
	require.Equal(usr.Name, "a8m")
	require.Equal(usr.Last, "unknown")
	t.Logf("user created: %v", usr)

	t.Log("querying assoc edges")
	groups := usr.QueryGroups().IDsX(ctx)
	require.NotEmpty(groups)
	require.Equal(grp.ID, groups[0])
	t.Log("querying inverse edge")
	users := grp.QueryUsers().IDsX(ctx)
	require.NotEmpty(users)
	require.Equal(usr.ID, users[0])

	t.Log("remove group edge")
	client.User.UpdateOne(usr).RemoveGroups(grp).ExecX(ctx)
	require.Empty(grp.QueryUsers().AllX(ctx))
	require.Empty(usr.QueryGroups().AllX(ctx))
	t.Logf("add group edge")
	client.User.UpdateOne(usr).AddGroups(grp).ExecX(ctx)
	require.NotEmpty(grp.QueryUsers().AllX(ctx))
	require.NotEmpty(usr.QueryGroups().AllX(ctx))
	t.Log("remove users inverse edge")
	client.Group.UpdateOne(grp).RemoveUsers(usr).ExecX(ctx)
	require.Empty(grp.QueryUsers().AllX(ctx))
	require.Empty(usr.QueryGroups().AllX(ctx))
	t.Logf("add group inverse edge")
	client.Group.UpdateOne(grp).AddUsers(usr).ExecX(ctx)
	require.NotEmpty(grp.QueryUsers().AllX(ctx))
	require.NotEmpty(usr.QueryGroups().AllX(ctx))

	t.Log("count vertices")
	require.Equal(1, client.User.Query().CountX(ctx))
	require.Equal(1, client.Group.Query().CountX(ctx))

	t.Log("get only vertices")
	require.NotNil(client.User.Query().OnlyX(ctx))
	require.NotNil(client.Group.Query().OnlyX(ctx))

	t.Log("get only ids")
	require.NotEmpty(client.User.Query().OnlyIDX(ctx))
	require.NotEmpty(client.Group.Query().OnlyIDX(ctx))

	t.Log("query spouse edge")
	require.Zero(client.User.Query().Where(user.HasSpouse()).CountX(ctx))
	neta := client.User.Create().SetName("neta").SetAge(18).SetSpouse(usr).SaveX(ctx)
	require.Equal(2, client.User.Query().Where(user.HasSpouse()).CountX(ctx))

	t.Log("check for singular error")
	_, err = client.User.Query().Only(ctx)
	require.True(ent.IsNotSingular(err))

	t.Log("query parent/children edges")
	require.False(usr.QueryParent().ExistX(ctx))
	require.Empty(usr.QueryChildren().AllX(ctx))
	child := client.User.Create().SetName("pedro").SetAge(7).SetParent(usr).SaveX(ctx)
	require.Equal(usr.Name, child.QueryParent().OnlyX(ctx).Name)
	require.Equal(child.Name, usr.QueryChildren().OnlyX(ctx).Name)
	require.False(usr.QueryParent().ExistX(ctx))

	t.Log("clear parent edge")
	brat := client.User.Create().SetName("brat").SetAge(19).SetParent(usr).SaveX(ctx)
	require.Equal(2, usr.QueryChildren().CountX(ctx))
	brat = client.User.UpdateOne(brat).ClearParent().SaveX(ctx)
	err = client.User.UpdateOne(brat).ClearParent().Exec(ctx)
	require.NoError(err)
	require.False(brat.QueryParent().ExistX(ctx))
	require.Equal(1, usr.QueryChildren().CountX(ctx))

	t.Log("delete child clears edge")
	brat = client.User.UpdateOne(brat).SetParent(usr).SaveX(ctx)
	require.Equal(2, usr.QueryChildren().CountX(ctx))
	client.User.DeleteOne(brat).ExecX(ctx)
	require.Equal(1, usr.QueryChildren().CountX(ctx))

	client.Group.UpdateOne(grp).AddBlocked(neta).ExecX(ctx)
	blocked := usr.QueryGroups().OnlyX(ctx).QueryBlocked().OnlyX(ctx)
	t.Log("blocked:", blocked)

	t.Log("query users with or condition")
	require.Len(client.User.Query().Where(user.Or(user.Name("a8m"), user.Name("neta"))).AllX(ctx), 2)
	require.Len(client.User.Query().Where(user.Or(user.Name("a8m"), user.Name("noam"))).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.Or(user.Name("alex"), user.Name("noam"))).AllX(ctx), 0)

	t.Log("query using the in predicate")
	require.Len(client.User.Query().Where(user.NameIn("a8m", "neta")).AllX(ctx), 2)
	require.Len(client.User.Query().Where(user.NameIn("a8m", "alex")).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.IDIn(neta.ID)).AllX(ctx), 1)

	t.Log("query existence")
	require.True(client.User.Query().Where(user.Name("a8m")).Exist(ctx))
	require.False(client.User.Query().Where(user.Name("alex")).Exist(ctx))

	t.Log("query using get")
	require.Equal(usr.Name, client.User.GetX(ctx, usr.ID).Name)
	uid, err := client.User.Query().Where(user.ID(usr.ID), user.Not(user.Name(usr.Name))).Only(ctx)
	require.Error(err)
	require.Nil(uid)

	t.Log("test validators")
	err = client.Group.Create().SetInfo(info).SetType("a").SetName("Gituhb").SetExpire(time.Now().Add(time.Hour)).Exec(ctx)
	require.Error(err, "type validator failed")
	err = client.Group.Create().SetInfo(info).SetType("pass").SetName("failed").SetExpire(time.Now().Add(time.Hour)).Exec(ctx)
	require.Error(err, "name validator failed")
	var checkerr schema.CheckError
	require.True(errors.As(err, &checkerr))
	require.EqualError(err, `ent: validator failed for field "Group.name": last name must begin with uppercase`)
	require.EqualError(checkerr, "last name must begin with uppercase")
	err = client.Group.Create().SetInfo(info).SetType("pass").SetName("Github20").SetExpire(time.Now().Add(time.Hour)).Exec(ctx)
	require.Error(err, "name validator failed")
	err = client.Group.Create().SetInfo(info).SetType("pass").SetName("Github").SetMaxUsers(-1).SetExpire(time.Now().Add(time.Hour)).Exec(ctx)
	require.Error(err, "max_users validator failed")
	err = client.Group.Update().SetMaxUsers(-10).Exec(ctx)
	require.Error(err, "max_users validator failed")
	err = client.Group.UpdateOne(grp).SetMaxUsers(-10).Exec(ctx)
	require.Error(err, "max_users validator failed")
	_, err = client.Group.Query().Select("unknown_field").String(ctx)
	require.EqualError(err, "ent: invalid field \"unknown_field\" for query")
	_, err = client.Group.Query().GroupBy("unknown_field").String(ctx)
	require.EqualError(err, "ent: invalid field \"unknown_field\" for query")
	_, err = client.User.Query().Order(ent.Asc("invalid")).Only(ctx)
	require.EqualError(err, "ent: unknown column \"invalid\" for table \"users\"")
	_, err = client.User.Query().Order(ent.Asc("invalid")).QueryFollowing().Only(ctx)
	require.EqualError(err, "ent: unknown column \"invalid\" for table \"users\"")
	_, err = client.User.Query().GroupBy("name").Aggregate(ent.Sum("invalid")).String(ctx)
	require.EqualError(err, "ent: unknown column \"invalid\" for table \"users\"")

	t.Log("query using edge-with predicate")
	require.Len(usr.QueryGroups().Where(group.HasInfoWith(groupinfo.Desc("group info"))).AllX(ctx), 1)
	require.Empty(usr.QueryGroups().Where(group.HasInfoWith(groupinfo.Desc("missing info"))).AllX(ctx))
	t.Log("query using edge-with predicate on inverse edges")
	require.Len(client.Group.Query().Where(group.Name("Github"), group.HasUsersWith(user.Name("a8m"))).AllX(ctx), 1)
	require.Empty(client.Group.Query().Where(group.Name("Github"), group.HasUsersWith(user.Name("alex"))).AllX(ctx))
	t.Logf("query path using edge-with predicate")
	require.Len(client.GroupInfo.Query().Where(groupinfo.HasGroupsWith(group.HasUsersWith(user.Name("a8m")))).AllX(ctx), 1)
	require.Empty(client.GroupInfo.Query().Where(groupinfo.HasGroupsWith(group.HasUsersWith(user.Name("alex")))).AllX(ctx))
	require.Len(client.GroupInfo.Query().Where(groupinfo.Or(groupinfo.Desc("group info"), groupinfo.HasGroupsWith(group.HasUsersWith(user.Name("alex"))))).AllX(ctx), 1)

	t.Log("query with ordering")
	u1 := client.User.Query().Order(ent.Asc(user.FieldName)).FirstIDX(ctx)
	u2 := client.User.Query().Order(ent.Desc(user.FieldName)).FirstIDX(ctx)
	require.NotEqual(u1, u2)
	u1 = client.User.Query().Order(ent.Asc(user.FieldLast), ent.Asc(user.FieldAge)).FirstIDX(ctx)
	u2 = client.User.Query().Order(ent.Asc(user.FieldLast), ent.Desc(user.FieldAge)).FirstIDX(ctx)
	require.NotEqual(u1, u2)
	u1 = client.User.Query().Order(ent.Asc(user.FieldName, user.FieldAge)).FirstIDX(ctx)
	u2 = client.User.Query().Order(ent.Asc(user.FieldName, user.FieldAge)).FirstIDX(ctx)
	require.Equal(u1, u2)

	t.Log("query path")
	require.Len(client.Group.Query().QueryUsers().AllX(ctx), 1)
	require.Empty(client.Group.Query().Where(group.Name("boring")).QueryUsers().AllX(ctx))
	require.Equal(neta.Name, usr.QueryGroups().Where(group.Name("Github")).QueryUsers().QuerySpouse().OnlyX(ctx).Name)
	require.Empty(client.GroupInfo.Query().Where(groupinfo.Desc("group info")).QueryGroups().Where(group.Name("boring")).AllX(ctx))
	require.Equal(child.Name, client.GroupInfo.Query().Where(groupinfo.Desc("group info")).QueryGroups().Where(group.Name("Github")).QueryUsers().QueryChildren().FirstX(ctx).Name)
	neta.Update().AddGroups(grp).ExecX(ctx)
	require.Equal(grp.ID, client.User.Query().QueryGroups().OnlyIDX(ctx))

	t.Log("query using string predicate")
	require.Len(client.User.Query().Where(user.NameIn("a8m", "neta", "pedro")).AllX(ctx), 3)
	require.Empty(client.User.Query().Where(user.NameNotIn("a8m", "neta", "pedro")).AllX(ctx))
	require.Empty(client.User.Query().Where(user.NameIn("alex", "rocket")).AllX(ctx))
	require.NotNil(client.User.Query().Where(user.HasParentWith(user.NameIn("a8m", "neta"))).OnlyX(ctx))
	require.Len(client.User.Query().Where(user.NameContains("a8")).AllX(ctx), 1)
	require.Equal(1, client.User.Query().Where(user.NameHasPrefix("a8")).CountX(ctx))
	require.Zero(client.User.Query().Where(user.NameHasPrefix("%a8%")).CountX(ctx))
	require.Equal(2, client.User.Query().Where(user.Or(user.NameHasPrefix("a8"), user.NameHasSuffix("eta"))).CountX(ctx))

	t.Log("group-by one field")
	names, err := client.User.Query().GroupBy(user.FieldName).Strings(ctx)
	require.NoError(err)
	sort.Strings(names)
	require.Equal([]string{"a8m", "neta", "pedro"}, names)
	ages, err := client.User.Query().GroupBy(user.FieldAge).Ints(ctx)
	require.NoError(err)
	require.Len(ages, 3)
	age, err := client.User.Query().Where(user.Name("alexsn")).GroupBy(user.FieldAge).Int(ctx)
	require.True(ent.IsNotFound(err))
	require.Zero(age)

	t.Log("group-by two fields with aggregation")
	client.User.Create().SetName(usr.Name).SetAge(usr.Age).ExecX(ctx)
	client.User.Create().SetName(neta.Name).SetAge(neta.Age).ExecX(ctx)
	child2 := client.User.Create().SetName(child.Name).SetAge(child.Age + 1).SaveX(ctx)
	var v []struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Sum   int    `json:"sum"`
		Count int    `json:"count"`
	}
	client.User.Query().
		GroupBy(user.FieldName, user.FieldAge).
		Aggregate(ent.Count(), ent.Sum(user.FieldAge)).
		ScanX(ctx, &v)
	require.Len(v, 4)
	sort.Slice(v, func(i, j int) bool {
		if v[i].Name != v[j].Name {
			return v[i].Name < v[j].Name
		}
		return v[i].Age < v[j].Age
	})
	for i, usr := range []*ent.User{usr, neta} {
		require.Equal(usr.Name, v[i].Name)
		require.Equal(usr.Age, v[i].Age)
		require.Equal(usr.Age*2, v[i].Sum)
		require.Equal(2, v[i].Count, "should have 2 vertices")
	}
	v = v[2:]
	for i, usr := range []*ent.User{child, child2} {
		require.Equal(usr.Name, v[i].Name)
		require.Equal(usr.Age, v[i].Age)
		require.Equal(usr.Age, v[i].Sum)
		require.Equal(1, v[i].Count)
	}

	t.Log("group by with .as modulator")
	var v2 []struct {
		Name  string `json:"name"`
		Total int    `json:"total"`
	}
	client.User.Query().GroupBy(user.FieldName).Aggregate(ent.As(ent.Count(), "total")).ScanX(ctx, &v2)
	require.Len(v2, 3)
	for i := range v2 {
		require.Equal(2, v2[i].Total)
	}

	t.Log("group by a relation")
	foo := client.User.Create().SetName("foo").SetAge(10).AddPets(
		client.Pet.Create().SetName("a").SetAge(10).SaveX(ctx),
		client.Pet.Create().SetName("b").SetAge(7).SaveX(ctx),
	).SaveX(ctx)
	bar := client.User.Create().SetName("bar").SetAge(10).AddPets(
		client.Pet.Create().SetName("c").SetAge(14).SaveX(ctx),
		client.Pet.Create().SetName("d").SetAge(1).SaveX(ctx),
	).SaveX(ctx)

	var v3 []struct {
		ID      int
		Name    string
		Average float64
	}
	client.User.Query().
		Where(user.IDIn(foo.ID, bar.ID)).
		Order(ent.Asc(user.FieldID)).
		GroupBy(user.FieldID, user.FieldName).
		Aggregate(func(s *sql.Selector) string {
			// Join with pet table and calculate the
			// average age of the pets of each user.
			t := sql.Table(pet.Table)
			s.Join(t).On(s.C(user.FieldID), t.C(pet.OwnerColumn))
			return sql.As(sql.Avg(t.C(pet.FieldAge)), "average")
		}).
		ScanX(ctx, &v3)
	require.Len(v3, 2)
	require.Equal(foo.ID, v3[0].ID)
	require.Equal(foo.Name, v3[0].Name)
	require.Equal(8.5, v3[0].Average)
	require.Equal(bar.ID, v3[1].ID)
	require.Equal(bar.Name, v3[1].Name)
	require.Equal(7.5, v3[1].Average)

	var v4 []struct {
		ID    int    `sql:"id"`
		Name  string `sql:"name"`
		Owner string `sql:"owner"`
	}
	client.Pet.Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(user.Table).As(user.Table)
			s.Join(t).On(s.C(pet.OwnerColumn), t.C(user.FieldID)) // owner_id = id for edge fields.
			s.AppendSelect(sql.As(t.C(user.FieldName), "owner"))
		}).
		Order(ent.Asc(pet.FieldID)).
		Select(pet.FieldID, pet.FieldName).
		ScanX(ctx, &v4)
	require.Equal(v4[0].Name, "a")
	require.Equal(v4[0].Owner, "foo")
	require.Equal(v4[1].Name, "b")
	require.Equal(v4[1].Owner, "foo")
	require.Equal(v4[2].Name, "c")
	require.Equal(v4[2].Owner, "bar")
	require.Equal(v4[3].Name, "d")
	require.Equal(v4[3].Owner, "bar")
}

func ClearFields(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	img := client.File.Create().SetName("foo").SetSize(100).SetUser("a8m").SetGroup("Github").SaveX(ctx)

	t.Log("clear one field")
	img = img.Update().ClearUser().SaveX(ctx)
	require.Nil(t, img.User)
	img = client.File.Query().OnlyX(ctx)
	require.Nil(t, img.User)
	require.Equal(t, "Github", img.Group)

	t.Log("clear many fields")
	img = img.Update().ClearUser().ClearGroup().SaveX(ctx)
	require.Nil(t, img.User)
	img = client.File.Query().OnlyX(ctx)
	require.Nil(t, img.User)
	require.Empty(t, img.Group)

	t.Log("revert previous set")
	img = img.Update().SetUser("a8m").ClearUser().SaveX(ctx)
	require.Nil(t, img.User)
}

func ClearEdges(t *testing.T, client *ent.Client) {
	ctx := context.Background()

	t.Log("clear o2m edges")
	ft := client.FileType.Create().SetName("photo").SaveX(ctx)
	client.File.CreateBulk(
		client.File.Create().SetName("A").SetSize(10).SetType(ft),
		client.File.Create().SetName("B").SetSize(20).SetType(ft),
	).ExecX(ctx)
	require.NotZero(t, ft.QueryFiles().CountX(ctx))
	ft = ft.Update().ClearFiles().SaveX(ctx)
	require.Zero(t, ft.QueryFiles().CountX(ctx))

	t.Log("clear m2m edges")
	a8m := client.User.Create().SetName("a8m").SetAge(30).SaveX(ctx)
	nat := client.User.Create().SetName("nati").SetAge(28).SaveX(ctx)
	inf := client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
	hub := client.Group.Create().SetName("GitHub").SetExpire(time.Now()).SetInfo(inf).AddUsers(a8m, nat).SaveX(ctx)
	lab := client.Group.Create().SetName("GitLab").SetExpire(time.Now()).SetInfo(inf).AddUsers(a8m, nat).SaveX(ctx)
	require.Equal(t, 2, a8m.QueryGroups().CountX(ctx))
	a8m.Update().ClearGroups().ExecX(ctx)
	require.Zero(t, a8m.QueryGroups().CountX(ctx))
	err := client.Group.Update().AddUsers(a8m).Exec(ctx)
	require.NoError(t, err, "return the user-edge back to groups")
	require.Equal(t, 2, a8m.QueryGroups().CountX(ctx))

	t.Log("clear m2m inverse-edges")
	require.Equal(t, 2, hub.QueryUsers().CountX(ctx))
	hub = hub.Update().ClearUsers().SaveX(ctx)
	require.Zero(t, hub.QueryUsers().CountX(ctx))
	require.Equal(t, 2, lab.QueryUsers().CountX(ctx))
	client.Group.Update().ClearUsers().ExecX(ctx)
	require.Zero(t, lab.QueryUsers().CountX(ctx))
	require.Zero(t, a8m.QueryGroups().CountX(ctx))
	require.Zero(t, nat.QueryGroups().CountX(ctx))

	t.Log("clear m2m bidi-edges")
	friends := client.User.CreateBulk(
		client.User.Create().SetName("f1").SetAge(30).AddFriends(a8m, nat),
		client.User.Create().SetName("f2").SetAge(30).AddFriends(a8m, nat),
		client.User.Create().SetName("f3").SetAge(30).AddFriends(a8m, nat),
	).SaveX(ctx)
	for i := range friends {
		require.Equal(t, 2, friends[i].QueryFriends().CountX(ctx))
	}
	require.Equal(t, 3, a8m.QueryFriends().CountX(ctx))
	require.Equal(t, 3, nat.QueryFriends().CountX(ctx))
	nat = nat.Update().ClearFriends().SaveX(ctx)
	require.Zero(t, nat.QueryFriends().CountX(ctx))
	require.Equal(t, 3, a8m.QueryFriends().CountX(ctx))
	for i := range friends {
		require.Equal(t, 1, friends[i].QueryFriends().CountX(ctx))
	}
	client.User.Update().ClearFriends().ExecX(ctx)
	require.Zero(t, client.User.Query().Where(user.HasFriends()).CountX(ctx))

	t.Log("clear m2m inverse-bidi-edges")
	a8m = a8m.Update().AddFollowing(friends...).SaveX(ctx)
	require.Equal(t, 3, a8m.QueryFollowing().CountX(ctx))
	require.Zero(t, a8m.QueryFollowers().CountX(ctx))
	nat = nat.Update().AddFollowers(friends...).SaveX(ctx)
	require.Zero(t, nat.QueryFollowing().CountX(ctx))
	require.Equal(t, 3, nat.QueryFollowers().CountX(ctx))
	for i := range friends {
		require.Equal(t, 1, friends[i].QueryFollowers().CountX(ctx))
		require.Equal(t, 1, friends[i].QueryFollowing().CountX(ctx))
	}
	nat.Update().ClearFollowing().ExecX(ctx)
	require.Equal(t, 3, nat.QueryFollowers().CountX(ctx), "expect no effect on followers")
	nat.Update().ClearFollowers().ExecX(ctx)
	require.Zero(t, nat.QueryFollowers().CountX(ctx))
	for i := range friends {
		require.Equal(t, 1, friends[i].QueryFollowers().CountX(ctx), "expect no effect to followers")
		require.Zero(t, friends[i].QueryFollowing().CountX(ctx))
	}
	a8m.Update().ClearFollowers().ExecX(ctx)
	require.Equal(t, 3, a8m.QueryFollowing().CountX(ctx), "expect no effect on following")
	a8m.Update().ClearFollowing().ExecX(ctx)
	require.Zero(t, a8m.QueryFollowing().CountX(ctx))
	for i := range friends {
		require.Zero(t, friends[i].QueryFollowers().CountX(ctx))
		require.Zero(t, friends[i].QueryFollowing().CountX(ctx))
	}

	t.Log("remove/clear and add edges")
	a8m = a8m.Update().AddFollowing(friends[0], friends[1]).SaveX(ctx)
	require.Equal(t, []int{friends[0].ID, friends[1].ID}, a8m.QueryFollowing().Order(ent.Asc(user.FieldID)).IDsX(ctx))
	a8m = a8m.Update().RemoveFollowing(friends[0], friends[1]).AddFollowing(friends[2]).SaveX(ctx)
	require.Equal(t, friends[2].ID, a8m.QueryFollowing().OnlyIDX(ctx))
	a8m = a8m.Update().ClearFollowing().AddFollowing(friends[0]).SaveX(ctx)
	require.Equal(t, friends[0].ID, a8m.QueryFollowing().OnlyIDX(ctx))
}

func UniqueConstraint(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()

	t.Log("unique constraint violation on 1 field")
	foo := client.User.Create().SetAge(1).SetName("foo").SetNickname("baz").SaveX(ctx)
	_, err := client.User.Create().SetAge(1).SetName("bar").SetNickname("baz").Save(ctx)
	require.True(ent.IsConstraintError(err))
	bar := client.User.Create().SetAge(1).SetName("bar").SetNickname("bar").SetPhone("1").SaveX(ctx)

	t.Log("unique constraint violation on 2 fields")
	err = client.User.Create().SetAge(1).SetName("baz").SetNickname("bar").SetPhone("1").Exec(ctx)
	require.True(ent.IsConstraintError(err))
	err = client.User.Create().SetAge(1).SetName("baz").SetNickname("qux").SetPhone("1").Exec(ctx)
	require.True(ent.IsConstraintError(err))
	err = client.User.Create().SetAge(1).SetName("baz").SetNickname("bar").SetPhone("2").Exec(ctx)
	require.True(ent.IsConstraintError(err))
	client.User.Create().SetAge(1).SetName("baz").SetNickname("qux").SetPhone("2").ExecX(ctx)
	err = client.User.UpdateOne(foo).SetNickname("bar").SetPhone("1").Exec(ctx)
	require.True(ent.IsConstraintError(err))
	err = client.User.UpdateOne(foo).SetNickname("bar").SetPhone("2").Exec(ctx)
	require.True(ent.IsConstraintError(err))

	t.Log("o2o unique constraint on creation")
	dan := client.User.Create().SetAge(1).SetName("dan").SetNickname("dan").SetSpouse(foo).SaveX(ctx)
	require.Equal(dan.Name, foo.QuerySpouse().OnlyX(ctx).Name)
	err = client.User.Create().SetAge(1).SetName("b").SetSpouse(foo).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	t.Log("o2m/m2o unique constraint on creation")
	c1 := client.User.Create().SetAge(1).SetName("c1").SetNickname("c1").SetParent(foo).SaveX(ctx)
	c2 := client.User.Create().SetAge(1).SetName("c2").SetNickname("c2").SetParent(foo).SaveX(ctx)
	err = client.User.Create().SetAge(10).SetName("z").SetNickname("z").AddChildren(c1).Exec(ctx)
	require.True(ent.IsConstraintError(err), "c1 already has a parent")
	err = client.User.Create().SetAge(10).SetName("z").SetNickname("z").AddChildren(c2).Exec(ctx)
	require.True(ent.IsConstraintError(err), "c2 already has a parent")
	err = client.User.Create().SetAge(10).SetName("z").SetNickname("z").AddChildren(c1, c2).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	inf := client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
	grp := client.Group.Create().SetName("Github").SetExpire(time.Now()).SetInfo(inf).SaveX(ctx)
	err = client.GroupInfo.Create().SetDesc("desc").AddGroups(grp).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	p1 := client.Pet.Create().SetName("p1").SetOwner(foo).SaveX(ctx)
	p2 := client.Pet.Create().SetName("p2").SetOwner(foo).SaveX(ctx)
	err = client.User.Create().SetAge(10).SetName("new-owner").AddPets(p1, p2).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	err = client.User.UpdateOne(c2).SetNickname(c1.Nickname).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	t.Log("o2o unique constraint on update")
	err = client.User.UpdateOne(bar).SetSpouse(foo).Exec(ctx)
	require.True(ent.IsConstraintError(err))
	err = client.User.UpdateOne(foo).SetSpouse(bar).Exec(ctx)
	require.True(ent.IsConstraintError(err))
	client.User.UpdateOne(bar).ClearSpouse().ExecX(ctx)
	client.User.UpdateOne(foo).ClearSpouse().SetSpouse(bar).ExecX(ctx)
	require.False(dan.QuerySpouse().ExistX(ctx))
	require.Equal(bar.Name, foo.QuerySpouse().OnlyX(ctx).Name)
	require.Equal(foo.Name, bar.QuerySpouse().OnlyX(ctx).Name)

	t.Log("o2m unique constraint on update")
	err = client.User.UpdateOne(bar).SetAge(1).SetName("new-owner").AddPets(p1).Exec(ctx)
	require.True(ent.IsConstraintError(err))
	err = client.User.UpdateOne(bar).SetAge(1).SetName("new-owner").AddPets(p1, p2).Exec(ctx)
	require.True(ent.IsConstraintError(err))

	t.Log("unique constraint violation when updating more than 1 vertex")
	err = client.User.Update().SetNickname("yada").Exec(ctx)
	require.True(ent.IsConstraintError(err))
	require.False(client.User.Query().Where(user.Nickname("yada")).ExistX(ctx))
	client.User.Update().Where(user.Nickname("dan")).SetNickname("yada").ExecX(ctx)
	require.False(client.User.Query().Where(user.Nickname("dan")).ExistX(ctx))
	require.True(client.User.Query().Where(user.Nickname("yada")).ExistX(ctx))

	t.Log("unique constraint on numeric fields")
	cm1 := client.Comment.Create().SetUniqueInt(42).SetUniqueFloat(math.Pi).SaveX(ctx)
	err = client.Comment.Create().SetUniqueInt(42).SetUniqueFloat(math.E).Exec(ctx)
	require.Error(err)
	err = client.Comment.Create().SetUniqueInt(7).SetUniqueFloat(math.Pi).Exec(ctx)
	require.Error(err)
	client.Comment.Create().SetUniqueInt(7).SetUniqueFloat(math.E).ExecX(ctx)
	err = cm1.Update().SetUniqueInt(7).Exec(ctx)
	require.Error(err)
	err = cm1.Update().SetUniqueFloat(math.E).Exec(ctx)
	require.Error(err)
}

type mocker struct{ mock.Mock }

func (m *mocker) onCommit(err error)   { m.Called(err) }
func (m *mocker) onRollback(err error) { m.Called(err) }
func (m *mocker) rHook() ent.RollbackHook {
	return func(next ent.Rollbacker) ent.Rollbacker {
		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
			err := next.Rollback(ctx, tx)
			m.onRollback(err)
			return err
		})
	}
}

func Tx(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	t.Run("Rollback", func(t *testing.T) {
		tx, err := client.Tx(ctx)
		require.NoError(t, err)
		var m mocker
		m.On("onRollback", nil).Once()
		defer m.AssertExpectations(t)
		tx.OnRollback(m.rHook())
		tx.Node.Create().ExecX(ctx)
		require.NoError(t, tx.Rollback())
		require.Zero(t, client.Node.Query().CountX(ctx), "rollback should discard all changes")
	})
	t.Run("Commit", func(t *testing.T) {
		tx, err := client.Tx(ctx)
		require.NoError(t, err)
		var m mocker
		m.On("onCommit", mock.Anything).Twice()
		defer m.AssertExpectations(t)
		tx.OnCommit(func(next ent.Committer) ent.Committer {
			return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
				err := next.Commit(ctx, tx)
				m.onCommit(err)
				return err
			})
		})
		nde := tx.Node.Create().SaveX(ctx)
		require.NoError(t, tx.Commit())
		require.Error(t, tx.Commit(), "should return an error on the second call")
		require.NotZero(t, client.Node.Query().CountX(ctx), "commit should save all changes")
		_, err = nde.QueryNext().Count(ctx)
		require.Error(t, err, "should not be able to query after tx was closed")
		require.Zero(t, nde.Unwrap().QueryNext().CountX(ctx), "should be able to query the entity after wrap")
	})
	t.Run("Nested", func(t *testing.T) {
		tx, err := client.Tx(ctx)
		require.NoError(t, err)
		var m mocker
		m.On("onRollback", nil).Once()
		defer m.AssertExpectations(t)
		tx.OnRollback(m.rHook())
		_, err = tx.Client().Tx(ctx)
		require.Error(t, err, "cannot start a transaction within a transaction")
		require.NoError(t, tx.Rollback())
	})
	t.Run("TxOptions Rollback", func(t *testing.T) {
		skip(t, "SQLite")
		tx, err := client.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
		require.NoError(t, err)
		var m mocker
		m.On("onRollback", nil).Once()
		defer m.AssertExpectations(t)
		tx.OnRollback(func(next ent.Rollbacker) ent.Rollbacker {
			return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
				err := next.Rollback(ctx, tx)
				m.onRollback(err)
				require.NotNil(t, ctx)
				return err
			})
		})
		err = tx.Item.Create().Exec(ctx)
		require.Error(t, err, "expect creation to fail in read-only tx")
		require.NoError(t, tx.Rollback())
	})
	t.Run("TxOptions Commit", func(t *testing.T) {
		skip(t, "SQLite")
		tx, err := client.BeginTx(ctx, &sql.TxOptions{Isolation: stdsql.LevelReadCommitted})
		require.NoError(t, err)
		var m mocker
		m.On("onCommit", nil).Once()
		defer m.AssertExpectations(t)
		tx.OnCommit(func(next ent.Committer) ent.Committer {
			return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
				err := next.Commit(ctx, tx)
				m.onCommit(err)
				require.NotNil(t, ctx)
				return err
			})
		})
		err = tx.Item.Create().Exec(ctx)
		require.NoError(t, tx.Commit())
		require.NoError(t, err)
	})
}

func DefaultValue(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	c1 := client.Card.Create().SetNumber("102030").SetName("Firstname Lastname").SaveX(ctx)
	ctime, mtime := c1.CreateTime, c1.UpdateTime
	require.False(t, ctime.IsZero())
	require.False(t, mtime.IsZero())
	c1 = c1.Update().SetName("F Lastname").SaveX(ctx)
	require.False(t, c1.CreateTime.IsZero())
	require.False(t, c1.UpdateTime.IsZero())
	require.False(t, mtime.Equal(c1.UpdateTime))

	// Enum default value
	usr := client.User.
		Create().
		SetAge(23).
		SetName("dario").
		SaveX(ctx)
	require.Equal(t, usr.Role, user.Role("user"))

	b := time.Now().Add(-time.Hour)
	n1 := client.Node.Create().SetValue(1).SetUpdatedAt(b).SaveX(ctx)
	require.NotNil(t, n1.UpdatedAt)
	require.WithinDuration(t, b, *n1.UpdatedAt, time.Second)
	n1 = n1.Update().SetValue(2).SaveX(ctx)
	require.NotNil(t, n1.UpdatedAt)
	require.False(t, b.Equal(*n1.UpdatedAt))
}

func ImmutableValue(t *testing.T, client *ent.Client) {
	tests := []struct {
		name    string
		updater func() any
	}{
		{
			name: "Update",
			updater: func() any {
				return client.Card.Update()
			},
		},
		{
			name: "UpdateOne",
			updater: func() any {
				return client.Card.Create().SetNumber("42").SaveX(context.Background()).Update()
			},
		},
	}
	for _, tc := range tests {
		v := reflect.ValueOf(tc.updater())
		require.False(t, v.MethodByName("SetCreatedAt").IsValid())
		require.False(t, v.MethodByName("SetNillableCreatedAt").IsValid())
		require.False(t, v.MethodByName("SetNumber").IsValid())
		require.True(t, v.MethodByName("SetName").IsValid())
	}
}

func Sensitive(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	usr := client.User.Create().SetName("foo").SetAge(20).SetPassword("secret-password").SaveX(ctx)
	require.Equal("secret-password", usr.Password)
	require.Contains(usr.String(), "password=<sensitive>")
	b, err := json.Marshal(usr)
	require.NoError(err)
	require.NotContains(string(b), "secret-password")
}

func EagerLoading(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)

	a8m := client.User.Create().SetName("a8m").SetAge(30).SaveX(ctx)
	nati := client.User.Create().SetName("nati").SetAge(28).SetSpouse(a8m).SaveX(ctx)
	alex := client.User.Create().SetName("alexsn").SetAge(35).AddFriends(a8m).SaveX(ctx)
	client.Pet.Create().SetName("xabi").ExecX(ctx)
	client.Pet.Create().SetName("pedro").SetOwner(a8m).SetTeam(nati).ExecX(ctx)
	client.Card.Create().SetNumber("102030").SetOwner(a8m).ExecX(ctx)

	inf := client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
	files := ent.Files{
		client.File.Create().SetName("a").SetSize(10).SaveX(ctx),
		client.File.Create().SetName("b").SetSize(10).SaveX(ctx),
		client.File.Create().SetName("c").SetSize(10).SaveX(ctx),
	}
	typ := client.FileType.Create().SetName("type").AddFiles(files...).SaveX(ctx)
	hub := client.Group.Create().SetName("GitHub").SetExpire(time.Now()).AddUsers(alex, a8m).SetInfo(inf).SaveX(ctx)
	lab := client.Group.Create().SetName("GitLab").SetExpire(time.Now()).AddUsers(nati, a8m).SetInfo(inf).AddFiles(files...).SaveX(ctx)

	t.Run("O2O", func(t *testing.T) {
		users := client.User.
			Query().
			Where(user.HasSpouse()).
			WithSpouse().
			WithCard().
			WithParent().
			Order(ent.Asc(user.FieldName)).
			AllX(ctx)
		require.Len(users, 2)
		require.NotNil(users[0].Edges.Spouse)
		require.NotNil(users[1].Edges.Spouse)
		require.NotNil(nati.Name, users[0].Edges.Spouse.Name)
		require.NotNil(a8m.Name, users[1].Edges.Spouse.Name)
		require.NotNil(users[0].Edges.Card)
		require.Nil(users[1].Edges.Card)

		edges := users[0].Edges
		pets, err := edges.PetsOrErr()
		require.True(ent.IsNotLoaded(err))
		require.Nil(pets)
		groups, err := edges.GroupsOrErr()
		require.True(ent.IsNotLoaded(err))
		require.Nil(groups)
		card, err := edges.CardOrErr()
		require.Nil(err)
		require.NotNil(card)
		spouse, err := edges.SpouseOrErr()
		require.Nil(err)
		require.NotNil(spouse)
		parent, err := edges.ParentOrErr()
		require.True(ent.IsNotFound(err), "loaded but was not found")
		require.Nil(parent)
	})

	t.Run("O2M", func(t *testing.T) {
		pets := client.Pet.Query().AllX(ctx)
		require.Nil(pets[0].Edges.Team)
		require.Nil(pets[0].Edges.Owner)
		require.Nil(pets[1].Edges.Team)
		require.Nil(pets[1].Edges.Owner)

		pedro := client.Pet.Query().Where(pet.HasOwner()).WithOwner().OnlyX(ctx)
		require.Nil(pedro.Edges.Team)
		require.NotNil(pedro.Edges.Owner)
		require.Equal(a8m.Name, pedro.Edges.Owner.Name)

		pedro = client.Pet.Query().Where(pet.HasOwner()).WithOwner().WithTeam().OnlyX(ctx)
		require.NotNil(pedro.Edges.Team)
		require.NotNil(pedro.Edges.Owner)
		require.Equal(a8m.Name, pedro.Edges.Owner.Name)
		require.Equal(nati.Name, pedro.Edges.Team.Name)
	})

	t.Run("M2O", func(t *testing.T) {
		a8m := client.User.Query().Where(user.ID(a8m.ID)).OnlyX(ctx)
		require.Empty(a8m.Edges.Pets)

		a8m = client.User.
			Query().
			Where(user.ID(a8m.ID)).
			WithPets(func(q *ent.PetQuery) {
				q.WithTeam().Order(ent.Asc(pet.FieldName))
			}).
			OnlyX(ctx)
		require.Len(a8m.Edges.Pets, 1)
		require.Equal("pedro", a8m.Edges.Pets[0].Name)
		require.Equal(nati.Name, a8m.Edges.Pets[0].Edges.Team.Name)

		a8m = client.User.
			Query().
			Where(user.ID(a8m.ID)).
			WithPets(func(q *ent.PetQuery) {
				q.Where(pet.Name("unknown"))
			}).
			OnlyX(ctx)
		require.Empty(a8m.Edges.Pets)
		require.NotNil(a8m.Edges.Pets)
	})

	t.Run("M2M", func(t *testing.T) {
		users := client.User.
			Query().
			WithFriends().
			WithGroups(func(q *ent.GroupQuery) {
				q.Order(ent.Desc(group.FieldName))
			}).
			Order(ent.Asc(user.FieldName)).
			AllX(ctx)
		require.Equal(a8m.Name, users[0].Name)
		require.Len(users[0].Edges.Groups, 2)
		require.Len(users[0].Edges.Friends, 1)
		require.Equal(alex.Name, users[0].Edges.Friends[0].Name)
		g1, g2 := users[0].Edges.Groups[0], users[0].Edges.Groups[1]
		require.Equal(lab.Name, g1.Name)
		require.Equal(hub.Name, g2.Name)

		groups := client.Group.
			Query().
			WithUsers(func(q *ent.UserQuery) {
				q.Order(ent.Asc(user.FieldName))
			}).
			Order(ent.Asc(group.FieldName)).
			AllX(ctx)
		require.Len(groups, 2)
		g1, g2 = groups[0], groups[1]
		require.Equal(hub.Name, g1.Name)
		require.Equal(lab.Name, g2.Name)
		require.Equal(a8m.Name, g1.Edges.Users[0].Name)
		require.Equal(alex.Name, g1.Edges.Users[1].Name)
		require.Equal(a8m.Name, g2.Edges.Users[0].Name)
		require.Equal(nati.Name, g2.Edges.Users[1].Name)
		require.Equal(g1.Edges.Users[0], g2.Edges.Users[0], "should share the same object")
	})

	t.Run("Graph", func(t *testing.T) {
		users := client.User.
			Query().
			WithSpouse().
			WithFriends().
			WithGroups(func(q *ent.GroupQuery) {
				q.WithInfo()
				q.WithFiles(func(q *ent.FileQuery) {
					q.WithType()
					q.Order(ent.Asc(file.FieldName))
				})
				q.Order(ent.Desc(group.FieldName))
			}).
			Order(ent.Asc(user.FieldName)).
			AllX(ctx)
		require.Equal(a8m.Name, users[0].Name)
		require.NotNil(users[0].Edges.Spouse)
		require.Equal(nati.Name, users[0].Edges.Spouse.Name)
		require.Len(users[0].Edges.Groups, 2)
		require.Len(users[0].Edges.Friends, 1)
		require.Equal(alex.Name, users[0].Edges.Friends[0].Name)

		require.Equal(alex.Name, users[1].Name)
		require.Len(users[1].Edges.Groups, 1)
		require.Equal(hub.Name, users[1].Edges.Groups[0].Name)

		require.Equal(nati.Name, users[2].Name)
		require.Len(users[2].Edges.Groups, 1)
		require.Equal(lab.Name, users[2].Edges.Groups[0].Name)

		g1, g2 := users[0].Edges.Groups[0], users[0].Edges.Groups[1]
		require.Equal(lab.Name, g1.Name)
		require.Equal(hub.Name, g2.Name)
		require.Equal(inf.Desc, g1.Edges.Info.Desc)
		require.Equal([]string{"a", "c"}, []string{g1.Edges.Files[0].Name, g1.Edges.Files[2].Name})
		for _, f := range g1.Edges.Files {
			require.NotNil(f.Edges.Type)
			require.Equal(typ.Name, f.Edges.Type.Name)
		}
	})

	t.Run("LimitRows/O2M", func(t *testing.T) {
		skip(t, "MySQL/5")
		client.Pet.Delete().ExecX(ctx)
		client.Pet.Create().SetName("nala").SetOwner(nati).ExecX(ctx)
		client.Pet.Create().SetName("xabi3").SetOwner(a8m).ExecX(ctx)
		client.Pet.Create().SetName("xabi2").SetOwner(a8m).ExecX(ctx)
		client.Pet.Create().SetName("xabi1").SetOwner(a8m).ExecX(ctx)
		client.Pet.Create().SetName("lola4").SetOwner(alex).ExecX(ctx)
		client.Pet.Create().SetName("lola3").SetOwner(alex).ExecX(ctx)
		client.Pet.Create().SetName("lola2").SetOwner(alex).ExecX(ctx)
		client.Pet.Create().SetName("lola1").SetOwner(alex).ExecX(ctx)

		users := client.User.Query().WithPets().Order(ent.Asc(user.FieldID)).AllX(ctx)
		require.Len(users[0].Edges.Pets, 3)
		require.Len(users[1].Edges.Pets, 1)
		require.Len(users[2].Edges.Pets, 4)

		users = client.User.
			Query().
			WithPets(func(q *ent.PetQuery) {
				q.Modify(limitRows(pet.OwnerColumn, 2))
			}).
			Order(ent.Asc(user.FieldID)).
			AllX(ctx)
		require.Len(users[0].Edges.Pets, 2)
		require.Equal(users[0].Edges.Pets[0].Name, "xabi3")
		require.Equal(users[0].Edges.Pets[1].Name, "xabi2")
		require.Len(users[1].Edges.Pets, 1)
		require.Equal(users[1].Edges.Pets[0].Name, "nala")
		require.Len(users[2].Edges.Pets, 2)
		require.Equal(users[2].Edges.Pets[0].Name, "lola4")
		require.Equal(users[2].Edges.Pets[1].Name, "lola3")

		users = client.User.
			Query().
			WithPets(func(q *ent.PetQuery) {
				q.Modify(limitRows(pet.OwnerColumn, 1, pet.FieldName))
			}).
			Order(ent.Asc(user.FieldID)).
			AllX(ctx)
		require.Len(users[0].Edges.Pets, 1)
		require.Equal(users[0].Edges.Pets[0].Name, "xabi1")
		require.Len(users[1].Edges.Pets, 1)
		require.Equal(users[1].Edges.Pets[0].Name, "nala")
		require.Len(users[2].Edges.Pets, 1)
		require.Equal(users[2].Edges.Pets[0].Name, "lola1")
	})

	t.Run("LimitRows/M2M", func(t *testing.T) {
		skip(t, "MySQL/5")
		users := client.User.Query().WithGroups().Order(ent.Asc(user.FieldID)).AllX(ctx)
		require.Len(users[0].Edges.Groups, 2)
		require.Len(users[1].Edges.Groups, 1)
		require.Len(users[2].Edges.Groups, 1)

		users = client.User.
			Query().
			WithGroups(func(q *ent.GroupQuery) {
				q.Modify(limitRows(user.GroupsPrimaryKey[0], 1))
			}).
			Order(ent.Asc(user.FieldID)).
			AllX(ctx)
		require.Len(users[0].Edges.Groups, 1)
		require.Equal(users[0].Edges.Groups[0].Name, "GitHub")
		require.Len(users[1].Edges.Groups, 1)
		require.Equal(users[1].Edges.Groups[0].Name, "GitLab")
		require.Len(users[2].Edges.Groups, 1)
		require.Equal(users[2].Edges.Groups[0].Name, "GitHub")

		client.Group.Create().SetName("BitBucket").SetExpire(time.Now()).AddUsers(alex, a8m).SetInfo(inf).SaveX(ctx)
		users = client.User.
			Query().
			WithGroups(func(q *ent.GroupQuery) {
				q.Modify(limitRows(user.GroupsPrimaryKey[0], 1, group.FieldName))
			}).
			Order(ent.Asc(user.FieldID)).
			AllX(ctx)
		require.Len(users[0].Edges.Groups, 1)
		require.Equal(users[0].Edges.Groups[0].Name, "BitBucket")
		require.Len(users[1].Edges.Groups, 1)
		require.Equal(users[1].Edges.Groups[0].Name, "GitLab")
		require.Len(users[2].Edges.Groups, 1)
		require.Equal(users[2].Edges.Groups[0].Name, "BitBucket")
	})
}

func limitRows(partitionBy string, limit int, orderBy ...string) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		if len(orderBy) == 0 {
			orderBy = append(orderBy, "id")
		}
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderBy(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

func NamedEagerLoading(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	a8m := client.User.Create().SetName("a8m").SetAge(30).SaveX(ctx)
	p1 := client.Pet.Create().SetName("pet1").SetOwner(a8m).SetTrained(true).SaveX(ctx)
	p2 := client.Pet.Create().SetName("pet2").SetOwner(a8m).SetTrained(false).SaveX(ctx)

	a8m = client.User.Query().
		WithNamedPets("Trained", func(q *ent.PetQuery) { q.Where(pet.Trained(true)) }).
		WithNamedPets("Untrained", func(q *ent.PetQuery) { q.Where(pet.Trained(false)) }).
		OnlyX(ctx)
	trained, err := a8m.NamedPets("Trained")
	require.NoError(t, err)
	require.Len(t, trained, 1)
	require.Equal(t, p1.ID, trained[0].ID)
	untrained, err := a8m.NamedPets("Untrained")
	require.NoError(t, err)
	require.Len(t, untrained, 1)
	require.Equal(t, p2.ID, untrained[0].ID)
	unknown, err := a8m.NamedPets("Unknown")
	require.True(t, ent.IsNotLoaded(err))
	require.Nil(t, unknown)
}

// writerFunc is an io.Writer implemented by the underlying func.
type writerFunc func(p []byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) { return f(p) }

func NoSchemaChanges(t *testing.T, client *ent.Client) {
	w := writerFunc(func(p []byte) (int, error) {
		stmt := strings.Trim(string(p), "\n;")
		ok := []string{"BEGIN", "COMMIT"}
		if strings.Contains(t.Name(), "SQLite") {
			ok = append(ok, "PRAGMA foreign_keys = off", "PRAGMA foreign_keys = on")
		}
		if !func() bool {
			for _, s := range ok {
				if s == stmt {
					return true
				}
			}
			return false
		}() {
			t.Errorf("expect no statement to execute. got: %q", stmt)
		}
		return len(p), nil
	})
	tables, err := sqlschema.CopyTables(migrate.Tables)
	require.NoError(t, err)
	err = migrate.Create(
		context.Background(),
		migrate.NewSchema(&sqlschema.WriteDriver{Writer: w, Driver: client.Driver()}),
		tables,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	require.NoError(t, err)
}

func Mutation(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	setName := func(ns interface{ SetName(string) }, name string) {
		ns.SetName(name)
	}
	ub := client.User.Create().SetAge(30)
	setName(ub.Mutation(), "a8m")
	pb := client.Pet.Create()
	setName(pb.Mutation(), "pedro")

	a8m := ub.SaveX(ctx)
	require.Equal(t, "a8m", a8m.Name)
	pedro := pb.SaveX(ctx)
	require.Equal(t, "pedro", pedro.Name)

	setUsers := func(ms ...*ent.UserMutation) {
		for _, m := range ms {
			m.SetName("boring")
			m.SetAge(30)
		}
	}
	uu := a8m.Update().AddPets(pedro)
	ub = client.User.Create()
	setUsers(ub.Mutation(), uu.Mutation())
	a8m = uu.SaveX(ctx)
	usr := ub.SaveX(ctx)
	require.Equal(t, "boring", a8m.Name)
	require.Equal(t, "boring", usr.Name)

	require.Equal(t, []int{usr.ID}, a8m.Update().AddFriends(usr).Mutation().FriendsIDs())
	require.Empty(t, a8m.Update().AddFriends(usr).RemoveFriends(usr).Mutation().FriendsIDs())
	require.Equal(t, []int{usr.ID}, a8m.Update().AddFriends(usr).RemoveFriends(a8m).Mutation().FriendsIDs())
	a8m.Update().AddFriends(usr).ExecX(ctx)

	t.Run("IDs", func(t *testing.T) {
		ids := client.User.Query().IDsX(ctx)
		u := client.User.Update().Where(user.IDIn(ids...)).AddAge(1)
		mids, err := u.Mutation().IDs(ctx)
		require.NoError(t, err)
		// Order can change between the 2 queries.
		sort.Ints(ids)
		sort.Ints(mids)
		require.Equal(t, ids, mids)
		u.ExecX(ctx)

		u = client.User.
			Update().
			AddAge(1).
			Where(
				user.Name(a8m.Name),
				user.HasPets(),
				user.HasPetsWith(
					pet.Name(pedro.Name),
				),
			)
		mids, err = u.Mutation().IDs(ctx)
		require.NoError(t, err)
		require.Len(t, mids, 1)
		require.Equal(t, a8m.ID, mids[0])
		u.ExecX(ctx)
	})

	t.Run("Predicate", func(t *testing.T) {
		updater := a8m.Update()
		updater.Mutation().Where(user.Name(a8m.Name))
		updater.SetName("mashraki")
		a8m, err := updater.Save(ctx)
		require.NoError(t, err, "predicate should not affect the returned object")
		require.Equal(t, "mashraki", a8m.Name)

		updater = a8m.Update()
		updater.Mutation().Where(user.Name(a8m.Name + a8m.Name))
		updater.SetName("a8m")
		a8m, err = updater.Save(ctx)
		require.True(t, ent.IsNotFound(err))
		require.Nil(t, a8m)
	})
}

// Test templates codegen.
var (
	_ = ent.CardExtension{}
	_ = ent.Card{}.StaticField
	_ = ent.Client{}.TemplateField
	_ = []filetype.State{filetype.StateOn, filetype.StateOff}
	_ = []filetype.Type{filetype.TypeJPG, filetype.TypePNG, filetype.TypeSVG}
)

func CreateBulk(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	cards := client.Card.CreateBulk(
		client.Card.Create().SetNumber("10").SetName("1st"),
		client.Card.Create().SetNumber("20").SetName("2nd"),
		client.Card.Create().SetNumber("30").SetName("3rd"),
	).SaveX(ctx)
	require.Equal(t, cards[0].ID, cards[1].ID-1)
	require.Equal(t, cards[1].ID, cards[2].ID-1)

	inf := client.GroupInfo.Create().SetDesc("group info").SaveX(ctx)
	groups := client.Group.CreateBulk(
		client.Group.Create().SetName("Github").SetExpire(time.Now()).SetInfo(inf),
		client.Group.Create().SetName("GitLab").SetExpire(time.Now()).SetInfo(inf),
	).SaveX(ctx)
	require.Equal(t, inf.ID, groups[0].QueryInfo().OnlyIDX(ctx))
	require.Equal(t, inf.ID, groups[1].QueryInfo().OnlyIDX(ctx))

	client.User.Use(
		func(next ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
				m.SetPassword("password")
				return next.Mutate(ctx, m)
			})
		},
		func(next ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
				name, _ := m.Name()
				m.SetNickname("@" + name)
				return next.Mutate(ctx, m)
			})
		},
	)

	users := client.User.CreateBulk(
		client.User.Create().SetName("a8m").SetAge(20).AddGroups(groups...),
		client.User.Create().SetName("nati").SetAge(20).SetCard(cards[0]).AddGroups(groups[0]),
	).SaveX(ctx)
	require.Equal(t, 2, users[0].QueryGroups().CountX(ctx))
	require.False(t, users[0].QueryCard().ExistX(ctx))
	require.Equal(t, "password", users[0].Password)
	require.Equal(t, "@a8m", users[0].Nickname)
	require.Equal(t, groups[0].ID, users[1].QueryGroups().OnlyIDX(ctx))
	require.Equal(t, cards[0].ID, users[1].QueryCard().OnlyIDX(ctx))
	require.Equal(t, "password", users[1].Password)
	require.Equal(t, "@nati", users[1].Nickname)

	pets := client.Pet.CreateBulk(
		client.Pet.Create().SetName("pedro").SetOwner(users[0]),
		client.Pet.Create().SetName("xabi").SetOwner(users[1]),
		client.Pet.Create().SetName("layla"),
	).SaveX(ctx)
	require.Equal(t, "pedro", pets[0].Name)
	require.Equal(t, users[0].ID, pets[0].QueryOwner().OnlyIDX(ctx))
	require.Equal(t, "xabi", pets[1].Name)
	require.Equal(t, users[1].ID, pets[1].QueryOwner().OnlyIDX(ctx))
	require.Equal(t, "layla", pets[2].Name)
	require.False(t, pets[2].QueryOwner().ExistX(ctx))
}

func ConstraintChecks(t *testing.T, client *ent.Client) {
	var cerr *ent.ConstraintError
	err := client.Pet.Create().SetName("orphan").SetOwnerID(0).Exec(context.Background())
	require.True(t, errors.As(err, &cerr))
	require.True(t, sqlgraph.IsForeignKeyConstraintError(err))
	require.False(t, sqlgraph.IsUniqueConstraintError(err))

	client.FileType.Create().SetName("a unique name").SaveX(context.Background())
	err = client.FileType.Create().SetName("a unique name").Exec(context.Background())
	require.True(t, errors.As(err, &cerr))
	require.False(t, sqlgraph.IsForeignKeyConstraintError(err))
	require.True(t, sqlgraph.IsUniqueConstraintError(err))
}

func Lock(t *testing.T, client *ent.Client) {
	skip(t, "SQLite", "MySQL/5", "Maria/10.2")
	ctx := context.Background()
	xabi := client.Pet.Create().SetName("Xabi").SaveX(ctx)

	t.Run("ForUpdate", func(t *testing.T) {
		tx1, err := client.Tx(ctx)
		require.NoError(t, err)
		tx2, err := client.Tx(ctx)
		require.NoError(t, err)
		tx3, err := client.Tx(ctx)
		require.NoError(t, err)
		p1 := tx1.Pet.Query().Where(pet.ID(xabi.ID)).ForUpdate().OnlyX(ctx)
		_, err = tx2.Pet.Query().Where(pet.ID(xabi.ID)).ForUpdate(sql.WithLockAction(sql.NoWait)).Only(ctx)
		switch name := t.Name(); {
		case strings.Contains(name, "Postgres"):
			err := err.(*pq.Error)
			require.EqualValues(t, "55P03", err.Code)
			require.EqualValues(t, `could not obtain lock on row in relation "pet"`, err.Message)
		case strings.Contains(name, "MySQL"):
			err := err.(*mysql.MySQLError)
			require.EqualValues(t, 3572, err.Number)
			require.EqualValues(t, "Statement aborted because lock(s) could not be acquired immediately and NOWAIT is set.", err.Message)
		case strings.Contains(name, "Maria"):
			err := err.(*mysql.MySQLError)
			require.EqualValues(t, 1205, err.Number)
			require.EqualValues(t, "Lock wait timeout exceeded; try restarting transaction", err.Message)
		}
		require.NoError(t, tx2.Rollback())
		p1.Update().SetName("updated").ExecX(ctx)
		require.NoError(t, tx1.Commit())
		tx3.Pet.Query().Where(pet.ID(xabi.ID)).ForUpdate().OnlyX(ctx)
		require.NoError(t, tx3.Rollback())
	})

	t.Run("ForShare", func(t *testing.T) {
		skip(t, "Maria")
		tx1, err := client.Tx(ctx)
		require.NoError(t, err)
		tx2, err := client.Tx(ctx)
		require.NoError(t, err)
		tx3, err := client.Tx(ctx)
		require.NoError(t, err)
		tx1.Pet.Query().Where(pet.ID(xabi.ID)).ForShare().OnlyX(ctx)
		tx2.Pet.Query().Where(pet.ID(xabi.ID)).ForShare().OnlyX(ctx)
		_, err = tx3.Pet.Query().
			Where(pet.ID(xabi.ID)).
			ForUpdate(
				sql.WithLockTables(pet.Table),
				sql.WithLockAction(sql.NoWait),
			).
			Only(ctx)
		require.Error(t, err)
		require.NoError(t, tx1.Rollback())
		require.NoError(t, tx2.Rollback())
		require.NoError(t, tx3.Rollback())
	})
}

func skip(t *testing.T, names ...string) {
	for _, n := range names {
		if strings.Contains(t.Name(), n) {
			t.Skipf("skip %s", n)
		}
	}
}

func drop(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	ctx := context.Background()
	client.Pet.Delete().ExecX(ctx)
	client.Item.Delete().ExecX(ctx)
	client.Task.Delete().ExecX(ctx)
	client.File.Delete().ExecX(ctx)
	client.Card.Delete().ExecX(ctx)
	client.Node.Delete().ExecX(ctx)
	client.User.Delete().ExecX(ctx)
	client.Group.Delete().ExecX(ctx)
	client.Comment.Delete().ExecX(ctx)
	client.GroupInfo.Delete().ExecX(ctx)
	client.FieldType.Delete().ExecX(ctx)
	client.FileType.Delete().ExecX(ctx)
}
