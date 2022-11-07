// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package base

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages/packagestest"
)

func TestPkgPath(t *testing.T) { packagestest.TestAll(t, testPkgPath) }
func testPkgPath(t *testing.T, x packagestest.Exporter) {
	e := packagestest.Export(t, x, []packagestest.Module{
		{
			Name: "golang.org/x",
			Files: map[string]any{
				"x.go":   "package x",
				"y/y.go": "package y",
			},
		},
	})
	defer e.Cleanup()

	e.Config.Dir = filepath.Dir(e.File("golang.org/x", "x.go"))
	target := filepath.Join(e.Config.Dir, "ent")
	pkgPath, err := PkgPath(e.Config, target)
	require.NoError(t, err)
	require.Equal(t, "golang.org/x/ent", pkgPath)

	e.Config.Dir = filepath.Dir(e.File("golang.org/x", "y/y.go"))
	target = filepath.Join(e.Config.Dir, "ent")
	pkgPath, err = PkgPath(e.Config, target)
	require.NoError(t, err)
	require.Equal(t, "golang.org/x/y/ent", pkgPath)

	target = filepath.Join(e.Config.Dir, "z/ent")
	pkgPath, err = PkgPath(e.Config, target)
	require.NoError(t, err)
	require.Equal(t, "golang.org/x/y/z/ent", pkgPath)

	target = filepath.Join(e.Config.Dir, "z/e/n/t")
	pkgPath, err = PkgPath(e.Config, target)
	require.Error(t, err)
	require.Empty(t, pkgPath)
}
