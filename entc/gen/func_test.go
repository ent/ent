// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen_test

import (
	"testing"

	"entgo.io/ent/entc/gen"
	"github.com/stretchr/testify/require"
)

func TestInflectRules(t *testing.T) {
	plural := gen.Funcs["plural"].(func(string) string)
	singular := gen.Funcs["singular"].(func(string) string)

	// Uncountable words keep their form in both directions.
	for _, w := range []string{"chassis", "equipment", "series"} {
		require.Equal(t, w, gen.Rules.Pluralize(w), "Pluralize(%q)", w)
		require.Equal(t, w, gen.Rules.Singularize(w), "Singularize(%q)", w)
		require.Equal(t, w, singular(w), "singular(%q)", w)
	}

	// The `plural` template function still appends "Slice" to distinguish the
	// entity slice type from the entity itself.
	require.Equal(t, "chassisSlice", plural("chassis"))

	// Regular and irregular words are unaffected.
	require.Equal(t, "users", gen.Rules.Pluralize("user"))
	require.Equal(t, "people", gen.Rules.Pluralize("person"))
	require.Equal(t, "user", gen.Rules.Singularize("users"))

	// The ruleset is exposed, so callers can add their own rules before
	// running code generation, and the template functions observe them.
	gen.Rules.AddIrregular("wug", "wugz")
	require.Equal(t, "wugz", plural("wug"))
	require.Equal(t, "wug", singular("wugz"))
}
