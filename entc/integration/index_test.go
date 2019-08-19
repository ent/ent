package integration

import (
	"context"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/ent"

	"github.com/stretchr/testify/require"
)

func Indexes(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)
	t.Log("prevent inserting 2 files with the same (name, user)")
	f1, err := client.File.Create().SetName("foo").SetSize(10).SetUser("bar").Save(ctx)
	require.NoError(err)
	require.Equal("foo", f1.Name)
	require.Equal("bar", *f1.User)
	f2, err := client.File.Create().SetName("foo").SetSize(10).SetUser("bar").Save(ctx)
	require.Nil(f2)
	require.Error(err)
	require.True(ent.IsConstraintFailure(err))

	t.Log("deletion should allow recreation")
	client.File.DeleteOne(f1).ExecX(ctx)
	f3, err := client.File.Create().SetName("foo").SetSize(10).SetUser("bar").Save(ctx)
	require.NoError(err)
	require.Equal("foo", f3.Name)
	require.Equal("bar", *f3.User)

	t.Log("allow inserting 2 files the same name and NULL user (optional field)")
	f4 := client.File.Create().SetName("foo").SetSize(10).SaveX(ctx)
	f5 := client.File.Create().SetName("foo").SetSize(10).SaveX(ctx)

	t.Log("index on edge sub-graph")
	a8m := client.User.Create().SetName("a8m").SetAge(18).SaveX(ctx)
	err = a8m.Update().AddFiles(f4).Exec(ctx)
	require.NoError(err)
	err = a8m.Update().AddFiles(f5).Exec(ctx)
	require.Error(err)
	require.True(ent.IsConstraintFailure(err))
	a8m.Update().RemoveFiles(f4).ExecX(ctx)

	t.Log("prevent inserting duplicates files in the same insert")
	err = a8m.Update().AddFiles(f4, f5).Exec(ctx)
	require.Error(err)
	require.True(ent.IsConstraintFailure(err))
	require.Zero(a8m.QueryFiles().CountX(ctx))

	t.Log("edge indexes should applied on the edge sub-graph")
	nati := client.User.Create().SetName("nati").SetAge(18).AddFiles(f5).SaveX(ctx)
	err = nati.Update().AddFiles(f4).Exec(ctx)
	require.Error(err)
	require.True(ent.IsConstraintFailure(err))
	err = a8m.Update().AddFiles(f4).Exec(ctx)
	require.NoError(err)

	require.Equal(1, a8m.QueryFiles().CountX(ctx))
	require.Equal(1, nati.QueryFiles().CountX(ctx))
}
