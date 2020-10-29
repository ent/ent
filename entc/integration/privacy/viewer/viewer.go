// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package viewer

import (
	"context"

	"github.com/facebook/ent/entc/integration/privacy/ent"
	"github.com/facebook/ent/entc/integration/privacy/ent/team"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Admin
	Edit
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Teams(context.Context) ([]string, error) // Team to query (tenant == team).
	Admin() bool                             // If viewer is admin.
	Can(Role) bool                           // If viewer is able to apply role action.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	User *ent.User // Actual user.
	Role Role      // Attached roles.
}

func (v UserViewer) Teams(ctx context.Context) ([]string, error) {
	return v.User.QueryTeams().Select(team.FieldName).Strings(ctx)
}

func (v UserViewer) Can(r Role) bool {
	if v.Admin() {
		return true
	}
	return v.Role&r != 0
}

func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

// AppViewer describes an app-viewer.
type AppViewer struct {
	Role Role // Attached roles.
}

func (v AppViewer) Teams(context.Context) ([]string, error) {
	return nil, nil
}

func (v AppViewer) Can(r Role) bool {
	if v.Admin() {
		return true
	}
	return v.Role&r != 0
}

func (v AppViewer) Admin() bool {
	return v.Role&Admin != 0
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
