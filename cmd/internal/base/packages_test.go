// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package base

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

func TestPkgPath(t *testing.T) {
	root := t.TempDir()
	for name, content := range map[string]string{
		"go.mod": "module golang.org/x\n",
		"x.go":   "package x",
		"y/y.go": "package y",
	} {
		require.NoError(t, os.MkdirAll(filepath.Join(root, filepath.Dir(name)), 0755))
		require.NoError(t, os.WriteFile(filepath.Join(root, name), []byte(content), 0644))
	}
	config := &packages.Config{Mode: packages.NeedName, Dir: root}
	target := filepath.Join(config.Dir, "ent")
	pkgPath, err := PkgPath(config, target)
	require.NoError(t, err)
	require.Equal(t, "golang.org/x/ent", pkgPath)

	config.Dir = filepath.Join(root, "y")
	target = filepath.Join(config.Dir, "ent")
	pkgPath, err = PkgPath(config, target)
	require.NoError(t, err)
	require.Equal(t, "golang.org/x/y/ent", pkgPath)

	target = filepath.Join(config.Dir, "z/ent")
	pkgPath, err = PkgPath(config, target)
	require.NoError(t, err)
	require.Equal(t, "golang.org/x/y/z/ent", pkgPath)

	target = filepath.Join(config.Dir, "z/e/n/t")
	pkgPath, err = PkgPath(config, target)
	require.Error(t, err)
	require.Empty(t, pkgPath)
}
