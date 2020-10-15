// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"database/sql"
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

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/entc/integration/ent"
	"github.com/facebook/ent/entc/integration/ent/enttest"
	"github.com/facebook/ent/entc/integration/ent/file"
	"github.com/facebook/ent/entc/integration/ent/filetype"
	"github.com/facebook/ent/entc/integration/ent/group"
	"github.com/facebook/ent/entc/integration/ent/groupinfo"
	"github.com/facebook/ent/entc/integration/ent/hook"
	"github.com/facebook/ent/entc/integration/ent/migrate"
	"github.com/facebook/ent/entc/integration/ent/node"
	"github.com/facebook/ent/entc/integration/ent/pet"
	"github.com/facebook/ent/entc/integration/ent/schema"
	"github.com/facebook/ent/entc/integration/ent/user"
	"github.com/stretchr/testify/mock"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestSQLite(t *testing.T) {
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
	t.Parallel()
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
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
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5433} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port), opts)
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
		NoSchemaChanges,
		Tx,
		Indexes,
		Types,
		Clone,
		EntQL,
		Sanity,
		Paging,
		Select,
		Delete,
		Relation,
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
		Mutation,
		CreateBulk,
	}
)

func Sanity(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()
	usr := client.User.Create().SetName("foo").SetAge(20).SaveX(ctx)
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
	// update fields.
	client.User.Update().Where(user.ID(child.ID)).SetName("Ariel").SaveX(ctx)
	client.User.Query().Where(user.Name("Ariel")).OnlyX(ctx)
	// update edges.
	require.Empty(child.QueryPets().AllX(ctx))
	require.NoError(client.Pet.UpdateOne(pt).ClearOwner().Exec(ctx))
	client.User.Update().Where(user.ID(child.ID)).AddPets(pt).SaveX(ctx)
	require.NotEmpty(child.QueryPets().AllX(ctx))
	client.User.Update().Where(user.ID(child.ID)).RemovePets(pt).SaveX(ctx)
	require.Empty(child.QueryPets().AllX(ctx))
	// remove edges.
	client.User.Update().ClearSpouse().SaveX(ctx)
	require.Empty(client.User.Query().Where(user.HasSpouse()).AllX(ctx))
	client.User.Update().AddFriends(child).RemoveGroups(grp).Where(user.ID(usr.ID)).SaveX(ctx)
	require.NotEmpty(child.QueryGroups().AllX(ctx))
	require.Empty(usr.QueryGroups().AllX(ctx))
	require.Len(child.QueryFriends().AllX(ctx), 1)
	require.Len(usr.QueryFriends().AllX(ctx), 1)
	// update one vertex.
	usr = client.User.UpdateOne(usr).SetName("baz").AddGroups(grp).SaveX(ctx)
	require.Equal("baz", usr.Name)
	require.NotEmpty(usr.QueryGroups().AllX(ctx))
	// update unknown vertex.
	_, err := client.User.UpdateOneID(usr.ID + usr.ID).SetName("foo").Save(ctx)
	require.Error(err)
	require.True(ent.IsNotFound(err))

	// grouping.
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
	client.User.Create().SetName("bar").SetAge(30).SaveX(ctx)
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
		client.Node.Create().SetValue(i).SaveX(ctx)
	}
	affected, err := client.Node.Delete().Where(node.ValueGT(2)).Exec(ctx)
	require.NoError(err)
	require.Equal(2, affected)

	affected, err = client.Node.Delete().Exec(ctx)
	require.NoError(err)
	require.Equal(3, affected)
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
	_, err = client.User.UpdateOne(brat).ClearParent().Save(ctx)
	require.NoError(err)
	require.False(brat.QueryParent().ExistX(ctx))
	require.Equal(1, usr.QueryChildren().CountX(ctx))

	t.Log("delete child clears edge")
	brat = client.User.UpdateOne(brat).SetParent(usr).SaveX(ctx)
	require.Equal(2, usr.QueryChildren().CountX(ctx))
	client.User.DeleteOne(brat).ExecX(ctx)
	require.Equal(1, usr.QueryChildren().CountX(ctx))

	client.Group.UpdateOne(grp).AddBlocked(neta).SaveX(ctx)
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
	_, err = client.Group.Create().SetInfo(info).SetType("a").SetName("Gituhb").SetExpire(time.Now().Add(time.Hour)).Save(ctx)
	require.Error(err, "type validator failed")
	_, err = client.Group.Create().SetInfo(info).SetType("pass").SetName("failed").SetExpire(time.Now().Add(time.Hour)).Save(ctx)
	require.Error(err, "name validator failed")
	var checkerr schema.CheckError
	require.True(errors.As(err, &checkerr))
	require.EqualError(err, "ent: validator failed for field \"name\": last name must begin with uppercase")
	require.EqualError(checkerr, "last name must begin with uppercase")
	_, err = client.Group.Create().SetInfo(info).SetType("pass").SetName("Github20").SetExpire(time.Now().Add(time.Hour)).Save(ctx)
	require.Error(err, "name validator failed")
	_, err = client.Group.Create().SetInfo(info).SetType("pass").SetName("Github").SetMaxUsers(-1).SetExpire(time.Now().Add(time.Hour)).Save(ctx)
	require.Error(err, "max_users validator failed")
	_, err = client.Group.Update().SetMaxUsers(-10).Save(ctx)
	require.Error(err, "max_users validator failed")
	_, err = client.Group.UpdateOne(grp).SetMaxUsers(-10).Save(ctx)
	require.Error(err, "max_users validator failed")
	_, err = client.Group.Query().Select("unknown_field").String(ctx)
	require.EqualError(err, "invalid field \"unknown_field\" for selection")
	_, err = client.Group.Query().GroupBy("unknown_field").String(ctx)
	require.EqualError(err, "invalid field \"unknown_field\" for group-by")
	_, err = client.User.Query().Order(ent.Asc("invalid")).Only(ctx)
	require.EqualError(err, "invalid field \"invalid\" for ordering")
	_, err = client.User.Query().Order(ent.Asc("invalid")).QueryFollowing().Only(ctx)
	require.EqualError(err, "invalid field \"invalid\" for ordering")
	_, err = client.User.Query().GroupBy("name").Aggregate(ent.Sum("invalid")).String(ctx)
	require.EqualError(err, "invalid field \"invalid\" for grouping")

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

	t.Log("query using string predicate")
	require.Len(client.User.Query().Where(user.NameIn("a8m", "neta", "pedro")).AllX(ctx), 3)
	require.Empty(client.User.Query().Where(user.NameNotIn("a8m", "neta", "pedro")).AllX(ctx))
	require.Empty(client.User.Query().Where(user.NameIn("alex", "rocket")).AllX(ctx))
	require.NotNil(client.User.Query().Where(user.HasParentWith(user.NameIn("a8m", "neta"))).OnlyX(ctx))
	require.Len(client.User.Query().Where(user.NameContains("a8")).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.NameHasPrefix("a8")).AllX(ctx), 1)
	require.Len(client.User.Query().Where(user.Or(user.NameHasPrefix("a8"), user.NameHasSuffix("eta"))).AllX(ctx), 2)

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
	client.User.Create().SetName(usr.Name).SetAge(usr.Age).SaveX(ctx)
	client.User.Create().SetName(neta.Name).SetAge(neta.Age).SaveX(ctx)
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
	).SaveX(ctx)
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
	a8m.Update().ClearGroups().SaveX(ctx)
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
	_, err = client.User.Create().SetAge(1).SetName("baz").SetNickname("bar").SetPhone("1").Save(ctx)
	require.True(ent.IsConstraintError(err))
	_, err = client.User.Create().SetAge(1).SetName("baz").SetNickname("qux").SetPhone("1").Save(ctx)
	require.True(ent.IsConstraintError(err))
	_, err = client.User.Create().SetAge(1).SetName("baz").SetNickname("bar").SetPhone("2").Save(ctx)
	require.True(ent.IsConstraintError(err))
	client.User.Create().SetAge(1).SetName("baz").SetNickname("qux").SetPhone("2").SaveX(ctx)
	_, err = client.User.UpdateOne(foo).SetNickname("bar").SetPhone("1").Save(ctx)
	require.True(ent.IsConstraintError(err))
	_, err = client.User.UpdateOne(foo).SetNickname("bar").SetPhone("2").Save(ctx)
	require.True(ent.IsConstraintError(err))

	t.Log("o2o unique constraint on creation")
	dan := client.User.Create().SetAge(1).SetName("dan").SetNickname("dan").SetSpouse(foo).SaveX(ctx)
	require.Equal(dan.Name, foo.QuerySpouse().OnlyX(ctx).Name)
	_, err = client.User.Create().SetAge(1).SetName("b").SetSpouse(foo).Save(ctx)
	require.True(ent.IsConstraintError(err))

	t.Log("o2m/m2o unique constraint on creation")
	c1 := client.User.Create().SetAge(1).SetName("c1").SetNickname("c1").SetParent(foo).SaveX(ctx)
	c2 := client.User.Create().SetAge(1).SetName("c2").SetNickname("c2").SetParent(foo).SaveX(ctx)
	_, err = client.User.Create().SetAge(10).SetName("z").SetNickname("z").AddChildren(c1).Save(ctx)
	require.True(ent.IsConstraintError(err), "c1 already has a parent")
	_, err = client.User.Create().SetAge(10).SetName("z").SetNickname("z").AddChildren(c2).Save(ctx)
	require.True(ent.IsConstraintError(err), "c2 already has a parent")
	_, err = client.User.Create().SetAge(10).SetName("z").SetNickname("z").AddChildren(c1, c2).Save(ctx)
	require.True(ent.IsConstraintError(err))

	inf := client.GroupInfo.Create().SetDesc("desc").SaveX(ctx)
	grp := client.Group.Create().SetName("Github").SetExpire(time.Now()).SetInfo(inf).SaveX(ctx)
	_, err = client.GroupInfo.Create().SetDesc("desc").AddGroups(grp).Save(ctx)
	require.True(ent.IsConstraintError(err))

	p1 := client.Pet.Create().SetName("p1").SetOwner(foo).SaveX(ctx)
	p2 := client.Pet.Create().SetName("p2").SetOwner(foo).SaveX(ctx)
	_, err = client.User.Create().SetAge(10).SetName("new-owner").AddPets(p1, p2).Save(ctx)
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
	_, err = client.User.UpdateOne(bar).SetAge(1).SetName("new-owner").AddPets(p1).Save(ctx)
	require.True(ent.IsConstraintError(err))
	_, err = client.User.UpdateOne(bar).SetAge(1).SetName("new-owner").AddPets(p1, p2).Save(ctx)
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
	_, err = client.Comment.Create().SetUniqueInt(42).SetUniqueFloat(math.E).Save(ctx)
	require.Error(err)
	_, err = client.Comment.Create().SetUniqueInt(7).SetUniqueFloat(math.Pi).Save(ctx)
	require.Error(err)
	_ = client.Comment.Create().SetUniqueInt(7).SetUniqueFloat(math.E).SaveX(ctx)
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
		tx.Node.Create().SaveX(ctx)
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
	t.Run("TxOptions", func(t *testing.T) {
		if strings.Contains(t.Name(), "SQLite") {
			t.Skip("SQLite does not support TxOptions.ReadOnly")
		}
		tx, err := client.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
		require.NoError(t, err)
		var m mocker
		m.On("onRollback", nil).Once()
		defer m.AssertExpectations(t)
		tx.OnRollback(m.rHook())
		_, err = tx.Item.Create().Save(ctx)
		require.Error(t, err)
		require.NoError(t, tx.Rollback())
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
}

func ImmutableValue(t *testing.T, client *ent.Client) {
	tests := []struct {
		name    string
		updater func() interface{}
	}{
		{
			name: "Update",
			updater: func() interface{} {
				return client.Card.Update()
			},
		},
		{
			name: "UpdateOne",
			updater: func() interface{} {
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
	client.Pet.Create().SetName("xabi").SaveX(ctx)
	client.Pet.Create().SetName("pedro").SetOwner(a8m).SetTeam(nati).SaveX(ctx)
	client.Card.Create().SetNumber("102030").SetOwner(a8m).SaveX(ctx)

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
}

// writerFunc is an io.Writer implemented by the underlying func.
type writerFunc func(p []byte) (int, error)

func (f writerFunc) Write(p []byte) (int, error) { return f(p) }

func NoSchemaChanges(t *testing.T, client *ent.Client) {
	w := writerFunc(func(p []byte) (int, error) {
		stmt := strings.Trim(string(p), "\n;")
		if stmt != "BEGIN" && stmt != "COMMIT" {
			t.Errorf("expect no statement to execute. got: %q", stmt)
		}
		return len(p), nil
	})
	err := client.Schema.WriteTo(context.Background(), w, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
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
	uu := a8m.Update()
	ub = client.User.Create()
	setUsers(ub.Mutation(), uu.Mutation())
	a8m = uu.SaveX(ctx)
	usr := ub.SaveX(ctx)
	require.Equal(t, "boring", a8m.Name)
	require.Equal(t, "boring", usr.Name)
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
