// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CheckDir checks the given dir and reports if there are any VCS conflicts.
func CheckDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && dir != path {
			return filepath.SkipDir
		}
		return checkFile(path)
	})
}

// conflictMarker holds the default marker string for
// both Git and Mercurial (default length is 7).
const conflictMarker = "<<<<<<<"

// checkFile checks the given file line by line
// and reports if it contains any VCS conflicts.
func checkFile(path string) error {
	fi, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fi.Close()

	scan := bufio.NewScanner(fi)
	scan.Split(bufio.ScanLines)
	for i := 0; scan.Scan(); i++ {
		if l := scan.Text(); strings.HasPrefix(l, conflictMarker) {
			return fmt.Errorf("vcs conflict %s:%d", path, i+1)
		}
	}
	return nil
}
