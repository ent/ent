// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package rule

import (
	"context"
	"fmt"
	"sync"

	"github.com/facebook/ent/entc/integration/privacy/ent"
	"github.com/facebook/ent/entc/integration/privacy/ent/hook"
	"github.com/facebook/ent/entc/integration/privacy/ent/predicate"
	"github.com/facebook/ent/entc/integration/privacy/ent/privacy"
	"github.com/facebook/ent/entc/integration/privacy/ent/task"
	"github.com/facebook/ent/entc/integration/privacy/ent/team"
	"github.com/facebook/ent/entc/integration/privacy/ent/user"
	"github.com/facebook/ent/entc/integration/privacy/viewer"
)

// DenyUpdateRule is a mutation rule that denies the update-many operation.
func DenyUpdateRule() privacy.MutationRule {
	return privacy.DenyMutationOperationRule(ent.OpUpdate)
}

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		return privacy.Skip
	})
}

// DenyIfNotAdmin is a rule that returns deny decision if the viewer not admin.
func DenyIfNotAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if !view.Admin() {
			return privacy.Denyf("viewer-context is not admin")
		}
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// AllowUserCreateIfAdmin is a rule that allows user creation only if the viewer is admin.
func AllowUserCreateIfAdmin() privacy.MutationRule {
	rule := privacy.UserMutationRuleFunc(func(ctx context.Context, _ *ent.UserMutation) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule, that may accept or reject this operation.
		return privacy.Skip
	})
	return privacy.OnMutationOperation(rule, ent.OpCreate)
}

// AllowTaskCreateIfOwner is a rule that allows creating task only if the creator is also the user.
func AllowTaskCreateIfOwner() privacy.MutationRule {
	rule := privacy.TaskMutationRuleFunc(func(ctx context.Context, m *ent.TaskMutation) error {
		view, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
		if !ok {
			return privacy.Skip
		}
		id, exists := m.OwnerID()
		if exists && view.User.ID == id {
			return privacy.Allow
		}
		// Skip to the next privacy rule, that may accept or reject this operation.
		return privacy.Skip
	})
	return privacy.OnMutationOperation(rule, ent.OpCreate)
}

// FilterTeamRule is a query rule that filters out tasks and users that are not in the team.
func FilterTeamRule() privacy.QueryRule {
	type TeamsFilter interface {
		WhereHasTeamsWith(...predicate.Team)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		teams, err := view.Teams(ctx)
		if err != nil {
			return privacy.Denyf("getting team names: %w", err)
		}
		tf, ok := f.(TeamsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		tf.WhereHasTeamsWith(team.NameIn(teams...))
		return privacy.Skip
	})
}

// DenyIfStatusChangedByOther is a mutation rule that returns a deny decision if the
// task status was changed by someone that is not the owner of the task, or an admin.
func DenyIfStatusChangedByOther() privacy.MutationRule {
	policy := privacy.TaskMutationRuleFunc(func(ctx context.Context, m *ent.TaskMutation) error {
		// Skip if the mutation does not change the task status.
		if _, exists := m.Status(); !exists {
			return privacy.Skip
		}
		view, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
		// Skip if the the viewer is an admin (or an app).
		if !ok || view.Admin() {
			return privacy.Skip
		}
		id, ok := m.ID()
		if !ok {
			return fmt.Errorf("missing task id")
		}
		owner, err := m.Client().User.Query().Where(user.HasTasksWith(task.ID(id))).Only(ctx)
		if err != nil {
			return err
		}
		// Deny the mutation, if the viewer is not the owner.
		if owner.ID != view.User.ID {
			return privacy.Denyf("viewer %d is not allowed to change the task status", view.User.ID)
		}
		return privacy.Skip
	})
	return privacy.OnMutationOperation(policy, ent.OpUpdateOne)
}

// AllowIfViewerInTheSameTeam returns allow decision if viewer on the same team as the task.
func AllowIfViewerInTheSameTeam() privacy.MutationRule {
	policy := privacy.TaskMutationRuleFunc(func(ctx context.Context, m *ent.TaskMutation) error {
		view, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
		// Skip if the the viewer is an admin (or an app).
		if !ok || view.Admin() {
			return privacy.Skip
		}
		teams, err := view.Teams(ctx)
		if err != nil {
			return privacy.Denyf("getting team names: %w", err)
		}
		id, ok := m.ID()
		if !ok {
			return fmt.Errorf("missing task id")
		}
		// Query should return an error if the viewer
		// does not belong to the task namespace/team.
		if _, err = m.Client().Task.Query().
			Where(
				task.ID(id),
				task.HasTeamsWith(team.NameIn(teams...)),
			).
			Only(ctx); err != nil {
			return err
		}
		return privacy.Allow
	})
	return privacy.OnMutationOperation(policy, ent.OpUpdateOne)
}

var logger = struct {
	logf func(string, ...interface{})
	sync.RWMutex
}{
	logf: func(string, ...interface{}) {},
}

// SetMutationLogFunc overrides the logging function used by LogPlanetMutationHook.
func SetMutationLogFunc(f func(string, ...interface{})) func(string, ...interface{}) {
	logger.Lock()
	defer logger.Unlock()
	logf := logger.logf
	logger.logf = f
	return logf
}

// LogPlanetMutationHook returns a hook logging planet mutations.
func LogTaskMutationHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.TaskFunc(func(ctx context.Context, m *ent.TaskMutation) (ent.Value, error) {
			value, err := next.Mutate(ctx, m)
			logger.RLock()
			defer logger.RUnlock()
			logger.logf("task mutation: type %s, value %v, err %v", m.Op(), value, err)
			return value, err
		})
	}
}
