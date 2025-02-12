// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package base

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

// This was pretty much ripped from https://github.com/golang/go/blob/master/src/cmd/go/internal/modload/init.go#L1581
func findModuleRoot(dir string) string {
	if dir == "" {
		panic("dir not set")
	}
	dir = filepath.Clean(dir)

	// Look for enclosing go.mod.
	for {
		if fi, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil && !fi.IsDir() {
			return dir
		}
		d := filepath.Dir(dir)
		if d == dir {
			break
		}
		dir = d
	}
	return ""
}

func PkgPath(target string) (string, error) {
	targetPath, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}

	modRoot := findModuleRoot(targetPath)
	if modRoot == "" {
		return "", fmt.Errorf("go module was not found for: %s", target)
	}

	goModPath := filepath.Join(modRoot, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", goModPath, err)
	}

	modPath := modfile.ModulePath(data)
	if modPath == "" {
		return "", fmt.Errorf("failed get module path from %s", goModPath)
	}

	rel, err := filepath.Rel(modRoot, targetPath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path from module root: %w", err)
	}

	return path.Join(modPath, rel), nil
}
