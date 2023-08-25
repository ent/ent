// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/socialprofile"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// SocialProfileCreate is the builder for creating a SocialProfile entity.
type SocialProfileCreate struct {
	config
	mutation *SocialProfileMutation
	hooks    []Hook
}

// SetDesc sets the "desc" field.
func (spc *SocialProfileCreate) SetDesc(s string) *SocialProfileCreate {
	spc.mutation.SetDesc(s)
	return spc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (spc *SocialProfileCreate) SetUserID(id string) *SocialProfileCreate {
	spc.mutation.SetUserID(id)
	return spc
}

// SetUser sets the "user" edge to the User entity.
func (spc *SocialProfileCreate) SetUser(u *User) *SocialProfileCreate {
	return spc.SetUserID(u.ID)
}

// Mutation returns the SocialProfileMutation object of the builder.
func (spc *SocialProfileCreate) Mutation() *SocialProfileMutation {
	return spc.mutation
}

// Save creates the SocialProfile in the database.
func (spc *SocialProfileCreate) Save(ctx context.Context) (*SocialProfile, error) {
	return withHooks(ctx, spc.gremlinSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *SocialProfileCreate) SaveX(ctx context.Context) *SocialProfile {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *SocialProfileCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *SocialProfileCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *SocialProfileCreate) check() error {
	if _, ok := spc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "SocialProfile.desc"`)}
	}
	if _, ok := spc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "SocialProfile.user"`)}
	}
	return nil
}

func (spc *SocialProfileCreate) gremlinSave(ctx context.Context) (*SocialProfile, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := spc.gremlin().Query()
	if err := spc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &SocialProfile{config: spc.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	spc.mutation.id = &rnode.ID
	spc.mutation.done = true
	return rnode, nil
}

func (spc *SocialProfileCreate) gremlin() *dsl.Traversal {
	v := g.AddV(socialprofile.Label)
	if value, ok := spc.mutation.Desc(); ok {
		v.Property(dsl.Single, socialprofile.FieldDesc, value)
	}
	for _, id := range spc.mutation.UserIDs() {
		v.AddE(user.SocialProfilesLabel).From(g.V(id)).InV()
	}
	return v.ValueMap(true)
}

// SocialProfileCreateBulk is the builder for creating many SocialProfile entities in bulk.
type SocialProfileCreateBulk struct {
	config
	err      error
	builders []*SocialProfileCreate
}
