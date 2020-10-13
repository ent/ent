// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package privacy provides sets of types and helpers for writing privacy
// rules in user schemas, and deal with their evaluation at runtime.
package privacy

import (
	"context"
	"errors"

	"github.com/facebook/ent"
)

// List of policy decisions.
var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

type (
	// Policies combines multiple policies into a single policy.
	Policies []ent.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}

	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}

	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule
)

// NewPolicies creates an ent.Policy from list of mixin.Schema
// and ent.Schema that implement the ent.Policy interface.
func NewPolicies(schemas ...interface{ Policy() ent.Policy }) ent.Policy {
	policies := make(Policies, 0, len(schemas))
	for i := range schemas {
		if policy := schemas[i].Policy(); policy != nil {
			policies = append(policies, policy)
		}
	}
	return policies
}

// EvalQuery evaluates the query policies. If the Allow error is returned
// from one of the policies, it stops the evaluation with a nil error.
func (policies Policies) EvalQuery(ctx context.Context, q ent.Query) error {
	return policies.eval(ctx, func(policy ent.Policy) error {
		return policy.EvalQuery(ctx, q)
	})
}

// EvalMutation evaluates the mutation policies. If the Allow error is returned
// from one of the policies, it stops the evaluation with a nil error.
func (policies Policies) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policies.eval(ctx, func(policy ent.Policy) error {
		return policy.EvalMutation(ctx, m)
	})
}

func (policies Policies) eval(ctx context.Context, eval func(ent.Policy) error) error {
	if decision, ok := DecisionFromContext(ctx); ok {
		return decision
	}
	for _, policy := range policies {
		switch decision := eval(policy); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// EvalQuery evaluates a query against a query policy.
func (policies QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	for _, policy := range policies {
		switch decision := policy.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		default:
			return decision
		}
	}
	return nil
}

// EvalMutation evaluates a mutation against a mutation policy.
func (policies MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	for _, policy := range policies {
		switch decision := policy.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		default:
			return decision
		}
	}
	return nil
}

type decisionCtxKey struct{}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}
