// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

// DefaultConfig for loading Go packages.
var DefaultConfig = &packages.Config{Mode: packages.NeedName}

// PkgPath returns the Go package name for given target path.
// Even if the existing path is not exist yet in the filesystem.
//
// If packages.Config is nil, DefaultConfig will be used to load packages.
func PkgPath(config *packages.Config, target string) (string, error) {
	if config == nil {
		config = DefaultConfig
	}
	pathCheck, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}
	var parts []string
	if _, err := os.Stat(pathCheck); os.IsNotExist(err) {
		parts = append(parts, filepath.Base(pathCheck))
		pathCheck = filepath.Dir(pathCheck)
	}
	// Try maximum 2 directories above the given
	// target to find the root package or module.
	for i := 0; i < 2; i++ {
		pkgs, err := packages.Load(config, pathCheck)
		if err != nil {
			return "", fmt.Errorf("load package info: %v", err)
		}
		if len(pkgs) == 0 || len(pkgs[0].Errors) != 0 {
			parts = append(parts, filepath.Base(pathCheck))
			pathCheck = filepath.Dir(pathCheck)
			continue
		}
		pkgPath := pkgs[0].PkgPath
		for j := len(parts) - 1; j >= 0; j-- {
			pkgPath = path.Join(pkgPath, parts[j])
		}
		return pkgPath, nil
	}
	return "", fmt.Errorf("root package or module was not found for: %s", target)
}
