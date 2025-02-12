// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package base

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func tempGoModule(t *testing.T, modPath string) string {
	t.Helper()

	tmpDir := t.TempDir()
	gomod := filepath.Join(tmpDir, "go.mod")

	f, err := os.Create(gomod)
	if err != nil {
		t.Fatalf("failed to create %s: %v", gomod, err)
	}
	defer f.Close()

	fmt.Fprintf(f, "module %s\n", modPath)

	return tmpDir
}

func TestPkgPath(t *testing.T) {
	root := tempGoModule(t, "example.com/foo")
	pkgPath, err := PkgPath(filepath.Join(root, "ent"))
	require.NoError(t, err)
	require.Equal(t, "example.com/foo/ent", pkgPath)

	pkgPath, err = PkgPath(filepath.Join(root, "x", "y", "ent"))
	require.NoError(t, err)
	require.Equal(t, "example.com/foo/x/y/ent", pkgPath)

	root = t.TempDir() // no go.mod
	pkgPath, err = PkgPath(filepath.Join(root, "ent"))
	require.Error(t, err)
	require.Empty(t, pkgPath)
}
