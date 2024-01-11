// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package privacy

import (
	"context"
	"errors"
	"testing"

	"entgo.io/ent/entc/integration/privacy/ent"
	"entgo.io/ent/entc/integration/privacy/ent/enttest"
	"entgo.io/ent/entc/integration/privacy/ent/privacy"
	"entgo.io/ent/entc/integration/privacy/ent/task"
	"entgo.io/ent/entc/integration/privacy/rule"
	"entgo.io/ent/entc/integration/privacy/viewer"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestPrivacyRules(t *testing.T) {
	client := enttest.Open(t, "sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	defer client.Close()
	logf := rule.SetMutationLogFunc(func(string, ...any) {
		require.FailNow(t, "hook called on privacy deny")
	})
	ctx := context.Background()
	_, err := client.Team.Create().SetName("ent").Save(ctx)
	require.True(t, errors.Is(err, privacy.Deny), "policy requires viewer context")
	view := viewer.NewContext(ctx, viewer.AppViewer{
		Role: viewer.View,
	})
	_, err = client.Team.CreateBulk(
		client.Team.Create().SetName("ent"),
		client.Team.Create().SetName("ent-contrib"),
	).Save(view)
	require.True(t, errors.Is(err, privacy.Deny), "team policy requires admin user")
	rule.SetMutationLogFunc(logf)

	admin := viewer.NewContext(ctx, viewer.AppViewer{
		Role: viewer.Admin,
	})
	teams := client.Team.CreateBulk(
		client.Team.Create().SetName("ent"),
		client.Team.Create().SetName("ent-contrib"),
	).SaveX(admin)

	_, err = client.User.Create().SetName("a8m").AddTeams(teams[0]).Save(view)
	require.True(t, errors.Is(err, privacy.Deny), "user creation requires admin user")
	a8m := client.User.Create().SetName("a8m").AddTeams(teams[0], teams[1]).SaveX(admin)
	nat := client.User.Create().SetName("nati").AddTeams(teams[1]).SaveX(admin)

	_, err = client.Task.Create().SetTitle("task 1").AddTeams(teams[0]).SetOwner(a8m).Save(ctx)
	require.True(t, errors.Is(err, privacy.Deny), "task creation requires viewer/owner match")

	a8mctx := viewer.NewContext(view, &viewer.UserViewer{User: a8m, Role: viewer.View | viewer.Edit})
	client.Task.Create().SetTitle("task 1").AddTeams(teams[0]).SetOwner(a8m).SaveX(a8mctx)
	_, err = client.Task.Create().SetTitle("task 2").AddTeams(teams[1]).SetOwner(nat).Save(a8mctx)
	require.True(t, errors.Is(err, privacy.Deny), "task creation requires viewer/owner match")

	natctx := viewer.NewContext(view, &viewer.UserViewer{User: nat, Role: viewer.View | viewer.Edit})
	client.Task.Create().SetTitle("task 2").AddTeams(teams[1]).SetOwner(nat).SaveX(natctx)

	tasks := client.Task.Query().AllX(a8mctx)
	require.Len(t, tasks, 2, "returned tasks from teams 1, 2")
	task2 := client.Task.Query().OnlyX(natctx)
	require.Equal(t, "task 2", task2.Title, "returned tasks must be from the same team")

	task3 := client.Task.Create().SetTitle("multi-team-task (1, 2)").AddTeams(teams...).SetOwner(a8m).SaveX(a8mctx)
	_, err = task3.Update().SetStatus(task.StatusClosed).Save(natctx)
	require.True(t, errors.Is(err, privacy.Deny), "viewer 2 is not allowed to change the task status")

	// DecisionContext returns a new context from the parent with a decision attached to it.
	task3.Update().SetStatus(task.StatusClosed).SaveX(privacy.DecisionContext(natctx, privacy.Allow))
	task3.Update().SetStatus(task.StatusClosed).SaveX(a8mctx)
	// Update description is allowed for other users in the team.
	task3.Update().SetDescription("boring description").SaveX(natctx)
	task3.Update().SetDescription("boring description").SaveX(a8mctx)

	// Create some notes (if not readonly can be updated by anyone).
	note1 := client.Note.Create().SetTitle("note 1").SetReadonly(true).SaveX(a8mctx)
	note2 := client.Note.Create().SetTitle("note 2").SetReadonly(false).SaveX(a8mctx)
	note3 := client.Note.Create().SetTitle("note 3").SetReadonly(false).SaveX(a8mctx)

	requireIDsInHook := func(expected []int, message string) func(ctx context.Context, next ent.Mutator, m *ent.NoteMutation) (ent.Value, error) {
		return func(ctx context.Context, next ent.Mutator, m *ent.NoteMutation) (ent.Value, error) {
			ids, err := m.IDs(ctx)
			require.NoError(t, err, "expected no error in hook")
			require.Len(t, ids, len(expected), message)
			for i, id := range ids {
				require.Equal(t, expected[i], id, message)
			}
			return next.Mutate(ctx, m)
		}
	}

	// Update note 2 is allowed
	rule.SetNoteMockHook(requireIDsInHook([]int{note2.ID}, "expected only note 2 in hook"))
	note2.Update().SetDescription("boring description").SaveX(a8mctx)

	// Update note 1 is not allowed
	rule.SetNoteMockHook(requireIDsInHook([]int{}, "expected no notes in hook"))
	_, err = note1.Update().SetDescription("boring description").Save(a8mctx)
	require.True(t, ent.IsNotFound(err), "viewer is not allowed to update the note")

	// Update many should only affect note 2 and 3
	rule.SetNoteMockHook(requireIDsInHook([]int{note2.ID, note3.ID}, "expected one note 2 and 3 in hook"))
	n := client.Note.Update().SetDescription("boring description").SaveX(a8mctx)
	require.Equal(t, 2, n, "expected one note to be updated")

	// Delete note 1 is not allowed
	rule.SetNoteMockHook(requireIDsInHook([]int{}, "expected no notes in hook"))
	err = client.Note.DeleteOne(note1).Exec(a8mctx)
	require.True(t, ent.IsNotFound(err), "viewer is not allowed to delete the note")

	// Delete note 2 is allowed
	rule.SetNoteMockHook(requireIDsInHook([]int{note2.ID}, "expected only note 2 in hook"))
	client.Note.DeleteOne(note2).ExecX(a8mctx)

	// Delete many should only affect note 3
	rule.SetNoteMockHook(requireIDsInHook([]int{note3.ID}, "expected only note 3 in hook"))
	n = client.Note.Delete().ExecX(a8mctx)
	require.Equal(t, 1, n, "expected one note to be deleted")

}
