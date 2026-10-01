// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package gen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"entgo.io/ent"
	"entgo.io/ent/entc/gen/testdata/jsonv2/schema"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"github.com/stretchr/testify/require"
)

func TestJSONOptions(t *testing.T) {
	// Exercise the schema serialization boundary before generating a client.
	var schemas []*load.Schema
	for _, s := range []ent.Interface{schema.Record{}, schema.Configured{}} {
		data, err := load.MarshalSchema(s)
		require.NoError(t, err)
		loaded, err := load.UnmarshalSchema(data)
		require.NoError(t, err)
		schemas = append(schemas, loaded)
	}
	require.False(t, schemas[0].Fields[1].JSONOptions)
	require.True(t, schemas[0].Fields[2].JSONOptions)

	dir := t.TempDir()
	root, err := filepath.Abs("../..")
	require.NoError(t, err)
	mod := fmt.Sprintf("module jsonv2test\n\ngo 1.25.0\n\nrequire entgo.io/ent v0.0.0\n\nreplace entgo.io/ent => %q\n", filepath.ToSlash(root))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "go.mod"), []byte(mod), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "doc.go"), []byte("package jsonv2test\n"), 0644))
	graph, err := NewGraph(&Config{
		Schema:   "entgo.io/ent/entc/gen/testdata/jsonv2/schema",
		Package:  "jsonv2test/ent",
		Target:   filepath.Join(dir, "ent"),
		Storage:  drivers[0],
		Features: []Feature{FeatureUpsert},
	}, schemas...)
	require.NoError(t, err)
	require.True(t, graph.Nodes[0].Fields[2].HasJSONOptions())
	tables, err := graph.Tables()
	require.NoError(t, err)
	for _, column := range tables[0].Columns[2:] {
		require.Equal(t, field.TypeJSON, column.Type)
	}
	runGo := func(args ...string) {
		t.Helper()
		cmd := exec.Command("go", args...)
		cmd.Dir = dir
		cmd.Env = append(os.Environ(), "GOWORK=off")
		out, err := cmd.CombinedOutput()
		require.NoError(t, err, "%s", out)
	}
	runGo("run", "-mod=mod", "entgo.io/ent/cmd/ent", "generate", "--feature", "sql/upsert", "--target", "./ent", "entgo.io/ent/entc/gen/testdata/jsonv2/schema")
	test, err := os.ReadFile("testdata/jsonv2/client_test.go.txt")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "ent", "jsonv2_test.go"), test, 0644))
	runGo("test", "-mod=mod", "./ent")

	_, err = NewGraph(&Config{Storage: drivers[1]}, schemas...)
	require.ErrorContains(t, err, "only supported by SQL storage")
}
