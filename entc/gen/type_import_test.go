// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"entgo.io/ent/entc/load"

	"github.com/stretchr/testify/require"
)

func TestGraph_GenericTypes(t *testing.T) {
	spec, err := (&load.Config{Path: "./testdata/generics"}).Load()
	require.NoError(t, err)
	target := t.TempDir()
	g, err := NewGraph(&Config{
		Schema:  spec.PkgPath,
		Target:  target,
		Package: "example.com/genericent",
		Storage: drivers[0],
	}, spec.Schemas...)
	require.NoError(t, err)
	require.NoError(t, g.Gen())

	// Compile every generated package, including mutations and edge builders.
	root, err := filepath.Abs("../..")
	require.NoError(t, err)
	mod, err := os.ReadFile(filepath.Join(root, "go.mod"))
	require.NoError(t, err)
	mod = bytes.Replace(mod, []byte("module entgo.io/ent"), []byte("module example.com/genericent"), 1)
	mod = append(mod, []byte("\nrequire entgo.io/ent v0.0.0\nreplace entgo.io/ent => "+root+"\n")...)
	require.NoError(t, os.WriteFile(filepath.Join(target, "go.mod"), mod, 0600))
	sum, err := os.ReadFile(filepath.Join(root, "go.sum"))
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(target, "go.sum"), sum, 0600))
	cmd := exec.Command("go", "test", "-mod=readonly", "./...")
	cmd.Dir = target
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "%s", output)
}
