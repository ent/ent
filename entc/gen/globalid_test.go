// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen_test

import (
	"testing"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"github.com/stretchr/testify/require"
)

func TestIncrementStartAnnotation(t *testing.T) {
	var (
		p = func(i int64) *int64 { return &i }
		a = &entsql.Annotation{IncrementStart: p(100)}
		s = []*load.Schema{
			{
				Name:        "T1",
				Annotations: gen.Annotations{a.Name(): a},
			},
			{Name: "T2"},
		}
		c = &gen.Config{
			Package: "entc/gen",
			Target:  t.TempDir(),
		}
	)
	// Arbitrary increment starting values allowed if feature is not enabled.
	g, err := gen.NewGraph(c, s...)
	require.NoError(t, err)
	require.NotNil(t, g)

	// Increments must be a multiple of 1<<32.
	c.Features = []gen.Feature{gen.FeatureGlobalID}
	g, err = gen.NewGraph(c, s...)
	require.EqualError(t, err, "unexpected increment start value 100 for type t1s, expected multiple of 4294967296 (1<<32)")
	require.Nil(t, g)
	a.IncrementStart = p(1 << 32)
	g, err = gen.NewGraph(c, s...)
	require.NoError(t, err)
	require.NotNil(t, g)

	// Duplicated increment starting values are not allowed.
	s = append(s, &load.Schema{
		Name:        "T3",
		Annotations: gen.Annotations{a.Name(): &entsql.Annotation{IncrementStart: p(1 << 32)}},
	})
	g, err = gen.NewGraph(c, s...)
	require.ErrorContains(t, err, "duplicated increment start value 4294967296 for types")
	require.Nil(t, g)
}
