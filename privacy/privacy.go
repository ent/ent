// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package privacy provides sets of types and helpers for writing privacy
// rules in user schemas, and deal with their evaluation at runtime.
package privacy

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent"
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

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

type (
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

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule interface {
		QueryRule
		MutationRule
	}
)

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query a policy.
func (p Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return p.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutate a  policy.
func (p Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return p.Mutation.EvalMutation(ctx, m)
}

// NewPolicies creates an ent.Policy from list of mixin.Schema
// and ent.Schema that implement the ent.Policy interface.
//
// Note that, this is a runtime function used by the ent generated
// code and should not be used in ent/schemas as a privacy rule.
func NewPolicies(schemas ...interface{ Policy() ent.Policy }) ent.Policy {
	policies := make(Policies, 0, len(schemas))
	for i := range schemas {
		if policy := schemas[i].Policy(); policy != nil {
			policies = append(policies, policy)
		}
	}
	return policies
}

// Policies combines multiple policies into a single policy.
//
// Note that, this is a runtime type used by the ent generated
// code and should not be used in ent/schemas as a privacy rule.
type Policies []ent.Policy

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

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}
