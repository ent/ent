// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/internal"
	"entgo.io/ent/entc/load"
	"github.com/stretchr/testify/require"
)

func TestIncrementStartAnnotation(t *testing.T) {
	var (
		p = func(i int) *int { return &i }
		a = &entsql.Annotation{IncrementStart: p(100)}
		s = []*load.Schema{
			{
				Name:        "T1",
				Annotations: map[string]any{a.Name(): must(gen.ToMap(a))},
			},
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
	a = &entsql.Annotation{IncrementStart: p(1 << 32)}
	s[0].Annotations[a.Name()] = must(gen.ToMap(a))
	g, err = gen.NewGraph(c, s...)
	require.NoError(t, err)
	require.NotNil(t, g)

	// Duplicated increment starting values are not allowed.
	s = append(s, &load.Schema{Name: "T2"}, &load.Schema{
		Name:        "T3",
		Annotations: map[string]any{a.Name(): must(gen.ToMap(a))},
	})
	g, err = gen.NewGraph(c, s...)
	require.ErrorContains(t, err, "duplicated increment start value 4294967296 for types")
	require.Nil(t, g)

	// Respects existing increment starting values loaded from file.
	s = []*load.Schema{{Name: "A"}, {Name: "B"}, {Name: "C"}}
	c.Target = t.TempDir()
	is := gen.IncrementStarts{"bs": 0, "as": 1 << 32, "cs": 2 << 32}
	require.NoError(t, is.WriteToDisk(c.Target))
	g, err = gen.NewGraph(c, s...)
	require.NoError(t, err)
	require.Equal(t, is, g.Annotations[is.Name()])
}

func TestResolveConflicts(t *testing.T) {
	var (
		c = &gen.Config{
			Package:  "entc/gen",
			Target:   t.TempDir(),
			Features: []gen.Feature{gen.FeatureGlobalID},
		}
		s     = []*load.Schema{{Name: "A"}, {Name: "B"}, {Name: "C"}, {Name: "D"}}
		cflct = fmt.Sprintf(`
package internal

<<<<<<< HEAD:globalid.go
const IncrementStarts = %s
=======
const IncrementStarts = %s
>>>>>>> 1234567:globalid.go
`,
			// We added the c table which got the range 2<<32
			marshal(t, gen.IncrementStarts{"bs": 0, "as": 1 << 32, "cs": 2 << 32}),
			// In the meantime someone else added the d table which got the
			// range 2<<32 as well, and they merged before we did.
			marshal(t, gen.IncrementStarts{"bs": 0, "as": 1 << 32, "ds": 2 << 32}),
		)
		p = filepath.Join(c.Target, "internal", "globalid.go")
	)
	require.NoError(t, os.MkdirAll(filepath.Dir(p), 0755))
	require.NoError(t, os.WriteFile(p, []byte(cflct), 0644))

	// Expect an error when there is a file conflict.
	_, err := gen.NewGraph(c, s...)
	require.Error(t, err)
	// Conflict is resolved to "accept theirs".
	require.NoError(t, gen.ResolveIncrementStartsConflict(c.Target))
	require.NoError(t, internal.CheckDir(filepath.Dir(p)))
	g, err := gen.NewGraph(c, s...)
	require.NoError(t, err)
	// Expect the conflict to be resolved with the remote table d keeping
	// its range and our newly added table c gets the next one (3<<32).
	require.Equal(t,
		gen.IncrementStarts{"bs": 0, "as": 1 << 32, "cs": 3 << 32, "ds": 2 << 32},
		g.Annotations[(&gen.IncrementStarts{}).Name()],
	)
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func marshal(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	require.NoError(t, err)
	return strconv.Quote(string(b))
}
