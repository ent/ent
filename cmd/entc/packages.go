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
	abs, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}
	pathCheck := abs
	if _, err := os.Stat(pathCheck); os.IsNotExist(err) {
		pathCheck = filepath.Dir(abs)
	}
	pkgs, err := packages.Load(config, pathCheck)
	if err != nil {
		return "", fmt.Errorf("load package info: %v", err)
	}
	if len(pkgs) == 0 {
		return "", fmt.Errorf("no package was found for: %s", pathCheck)
	}
	if errs := pkgs[0].Errors; len(errs) != 0 {
		return "", errs[0]
	}
	pkgPath := pkgs[0].PkgPath
	if abs != pathCheck {
		pkgPath = path.Join(pkgPath, filepath.Base(abs))
	}
	return pkgPath, nil
}
