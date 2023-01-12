// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"fmt"
	"go/token"
	"sort"
	"strings"

	"entgo.io/ent/schema/field"
)

const (
	goImportSeparator = "/"
)

var importTracker = newImportTracker()

// ImportTracker track the imports needed for a package
type ImportTracker struct {
	registerKeyWords map[string]struct{}
	pathToName       map[string]string
	nameToPath       map[string]string
}

func newImportTracker() *ImportTracker {
	registerKeyWords := map[string]struct{}{
		"go": {},
	}

	return &ImportTracker{
		registerKeyWords: registerKeyWords,
		pathToName:       make(map[string]string),
		nameToPath:       make(map[string]string),
	}
}

func (i *ImportTracker) Empty() (empty string) {
	i.pathToName = make(map[string]string)
	i.nameToPath = make(map[string]string)
	return
}

func (i *ImportTracker) TypeIdent(t *field.TypeInfo) string {
	path := t.PkgPath
	if len(path) == 0 {
		return t.String()
	}

	lcName := i.LocalNameOf(path)
	fmt.Println(lcName, path)
	if len(lcName) == 0 {
		return t.String()
	}

	return strings.ReplaceAll(t.String(), t.PkgName, lcName)
}

func (i *ImportTracker) ImportLine(path string) string {
	alias := i.pathToName[path]
	if strings.HasSuffix(path, alias) {
		alias = ""
	}

	return alias + " \"" + path + "\""
}

func (i *ImportTracker) ImportLines() []string {
	importPaths := []string{}
	for path := range i.pathToName {
		importPaths = append(importPaths, path)
	}

	sort.Sort(sort.StringSlice(importPaths))
	out := []string{}
	for _, path := range importPaths {
		out = append(out, i.ImportLine(path))
	}

	return out
}

// PathOf returns the path that a given localName is referring to within the
// body of a file.
func (i *ImportTracker) PathOf(localName string) (string, bool) {
	name, ok := i.nameToPath[localName]
	return name, ok
}

// LocalNameOf returns the name you would use to refer to the package at the
// specified path within the body of a file.
func (i *ImportTracker) LocalNameOf(path string) string {
	return i.pathToName[path]
}

func (i *ImportTracker) AddField(f *Field) (empty string) {
	return i.AddPath(f.Type.PkgPath)
}

// AddImport register a full import line with path and alias. rename another alias
// if duplicate and return false
func (i *ImportTracker) AddImport(alias, path string) bool {
	isDuplicated := false
	if existPath, exists := i.PathOf(alias); exists {
		if path != existPath {
			defer i.AddPath(path)

			i.removePath(path)
			isDuplicated = true
		}
	}

	i.nameToPath[alias] = path
	i.pathToName[path] = alias
	return isDuplicated
}

func (i *ImportTracker) removePath(path string) {
	name := i.LocalNameOf(path)
	delete(i.pathToName, path)
	delete(i.nameToPath, name)
}

func (i *ImportTracker) AddPath(paths ...string) (empty string) {
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}

		if _, found := i.pathToName[path]; found {
			continue
		}

		name := i.localName(path)
		i.pathToName[path] = name
		i.nameToPath[name] = path
	}

	return
}

func (i *ImportTracker) localName(pkg string) string {
	return golangTrackerLocalName(i, pkg)
}

func golangTrackerLocalName(tracker *ImportTracker, path string) string {
	dirs := strings.Split(path, goImportSeparator)
	for n := len(dirs) - 1; n >= 0; n-- {
		name := strings.Join(dirs[n:], "")
		name = strings.Replace(name, "_", "", -1)
		// These characters commonly appear in import paths for go
		// packages, but aren't legal go names. So we'll sanitize.
		name = strings.Replace(name, ".", "", -1)
		name = strings.Replace(name, "-", "", -1)
		if _, found := tracker.PathOf(name); found {
			// This name collides with some other package
			continue
		}

		if _, ok := tracker.registerKeyWords[name]; ok {
			continue
		}

		// If the import name is a Go keyword prefix with an underscore.
		if token.Lookup(name).IsKeyword() {
			name = "_" + name
		}

		return name
	}

	panic("can't find import for " + path)
}
