// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"github.com/stretchr/testify/require"
)

func BenchmarkGraph_Gen(b *testing.B) {
	target := filepath.Join(os.TempDir(), "ent")
	require.NoError(b, os.MkdirAll(target, os.ModePerm), "creating tmpdir")
	defer os.RemoveAll(target)
	storage, err := gen.NewStorage("sql")
	require.NoError(b, err)
	graph, err := entc.LoadGraph("../integration/ent/schema", &gen.Config{
		Storage: storage,
		IDType:  &field.TypeInfo{Type: field.TypeInt},
		Target:  target,
		Package: "github.com/facebook/ent/entc/integration/ent",
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("template").
				Funcs(gen.Funcs).
				ParseGlob("../integration/ent/template/*.tmpl")),
		},
	})
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := graph.Gen()
		require.NoError(b, err)
	}
}
