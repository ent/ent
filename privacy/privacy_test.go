// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package privacy_test

import (
	"context"
	"errors"
	"testing"

	"github.com/facebook/ent"
	"github.com/facebook/ent/privacy"
	"github.com/stretchr/testify/assert"
)

func TestDecisionContext(t *testing.T) {
	deny := privacy.DecisionContext(context.Background(), privacy.Deny)
	err := privacy.Policies{}.EvalMutation(deny, nil)
	assert.True(t, errors.Is(err, privacy.Deny))

	allow := privacy.DecisionContext(context.Background(), privacy.Allow)
	err = privacy.Policies{}.EvalQuery(allow, nil)
	assert.NoError(t, err)

	decision, ok := privacy.DecisionFromContext(deny)
	assert.True(t, ok)
	assert.True(t, errors.Is(decision, privacy.Deny))

	pc := policyFunc(func(context.Context) error { panic("invalid") })
	policy := privacy.NewPolicies(pc)
	err = policy.EvalQuery(allow, nil)
	assert.NoError(t, err)
	err = policy.EvalMutation(deny, nil)
	assert.True(t, errors.Is(err, privacy.Deny))
}

func TestNewPolicies(t *testing.T) {
	type ctxKey string
	var (
		key = ctxKey("i")
		pc  = policyFunc(func(ctx context.Context) error {
			*(ctx.Value(key).(*int))++
			return nil
		})
		policy = privacy.NewPolicies(pc, pc, pc, pc)
	)

	ctx := context.WithValue(context.Background(), key, new(int))
	err := policy.EvalMutation(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, 4, *(ctx.Value(key).(*int)))
	err = policy.EvalQuery(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, 8, *(ctx.Value(key).(*int)))
}

type policyFunc func(context.Context) error

func (f policyFunc) Policy() ent.Policy {
	return f
}

func (f policyFunc) EvalQuery(ctx context.Context, _ ent.Query) error {
	return f(ctx)
}

func (f policyFunc) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return f(ctx)
}
