// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"text/template/parse"

	"github.com/facebook/ent/entc/gen/internal"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -mode=420 -modtime=1 ./template/...

type (
	// TypeTemplate specifies a template that is executed with
	// each Type object of the graph.
	TypeTemplate struct {
		Name           string             // template name.
		Format         func(*Type) string // file name format.
		ExtendPatterns []string           // extend patterns.
	}
	// GraphTemplate specifies a template that is executed with
	// the Graph object.
	GraphTemplate struct {
		Name           string            // template name.
		Skip           func(*Graph) bool // skip condition (storage constraints or gated by a feature-flag).
		Format         string            // file name format.
		ExtendPatterns []string          // extend patterns.
	}
)

var (
	// Templates holds the template information for a file that the graph is generating.
	Templates = []TypeTemplate{
		{
			Name:   "create",
			Format: pkgf("%s_create.go"),
		},
		{
			Name:   "update",
			Format: pkgf("%s_update.go"),
		},
		{
			Name:   "delete",
			Format: pkgf("%s_delete.go"),
		},
		{
			Name:   "query",
			Format: pkgf("%s_query.go"),
		},
		{
			Name:   "model",
			Format: pkgf("%s.go"),
		},
		{
			Name:   "where",
			Format: pkgf("%s/where.go"),
			ExtendPatterns: []string{
				"where/additional/*",
			},
		},
		{
			Name: "meta",
			Format: func(t *Type) string {
				return fmt.Sprintf("%s/%s.go", t.Package(), t.Package())
			},
			ExtendPatterns: []string{
				"meta/additional/*",
			},
		},
	}
	// GraphTemplates holds the templates applied on the graph.
	GraphTemplates = []GraphTemplate{
		{
			Name:   "base",
			Format: "ent.go",
		},
		{
			Name:   "client",
			Format: "client.go",
			ExtendPatterns: []string{
				"client/fields/additional/*",
			},
		},
		{
			Name:   "context",
			Format: "context.go",
		},
		{
			Name:   "tx",
			Format: "tx.go",
		},
		{
			Name:   "config",
			Format: "config.go",
		},
		{
			Name:   "mutation",
			Format: "mutation.go",
		},
		{
			Name:   "migrate",
			Format: "migrate/migrate.go",
			Skip:   func(g *Graph) bool { return !g.SupportMigrate() },
		},
		{
			Name:   "schema",
			Format: "migrate/schema.go",
			Skip:   func(g *Graph) bool { return !g.SupportMigrate() },
		},
		{
			Name:   "predicate",
			Format: "predicate/predicate.go",
		},
		{
			Name:   "hook",
			Format: "hook/hook.go",
		},
		{
			Name:   "privacy",
			Format: "privacy/privacy.go",
			Skip: func(g *Graph) bool {
				return !g.featureEnabled(FeaturePrivacy)
			},
		},
		{
			Name:   "entql",
			Format: "entql.go",
			Skip: func(g *Graph) bool {
				return !g.featureEnabled(FeatureEntQL)
			},
		},
		{
			Name:   "runtime/ent",
			Format: "runtime.go",
		},
		{
			Name:   "enttest",
			Format: "enttest/enttest.go",
		},
		{
			Name:   "runtime/pkg",
			Format: "runtime/runtime.go",
		},
		{
			Name:   "internal/schema",
			Format: "internal/schema.go",
			Skip: func(g *Graph) bool {
				return !g.featureEnabled(FeatureSnapshot)
			},
		},
	}
	// templates holds the Go templates for the code generation.
	templates *Template
	// importPkg are the import packages used for code generation.
	importPkg = make(map[string]string)
)

func initTemplates() {
	templates = NewTemplate("templates")
	for _, asset := range internal.AssetNames() {
		templates = MustParse(templates.Parse(string(internal.MustAsset(asset))))
	}
	b := bytes.NewBuffer([]byte("package main\n"))
	check(templates.ExecuteTemplate(b, "import", Type{Config: &Config{}}), "load imports")
	f, err := parser.ParseFile(token.NewFileSet(), "", b, parser.ImportsOnly)
	check(err, "parse imports")
	for _, spec := range f.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		check(err, "unquote import path")
		importPkg[filepath.Base(path)] = path
	}
	for _, s := range drivers {
		for _, path := range s.Imports {
			importPkg[filepath.Base(path)] = path
		}
	}
}

// Template wraps the standard template.Template to
// provide additional functionality for ent extensions.
type Template struct {
	*template.Template
	FuncMap template.FuncMap
}

// NewTemplate creates an empty template with the standard codegen functions.
func NewTemplate(name string) *Template {
	t := &Template{Template: template.New(name)}
	return t.Funcs(Funcs)
}

// Funcs merges the given funcMap to the template functions.
func (t *Template) Funcs(funcMap template.FuncMap) *Template {
	t.Template.Funcs(funcMap)
	if t.FuncMap == nil {
		t.FuncMap = template.FuncMap{}
	}
	for name, f := range funcMap {
		if _, ok := t.FuncMap[name]; !ok {
			t.FuncMap[name] = f
		}
	}
	return t
}

// Parse parses text as a template body for t.
func (t *Template) Parse(text string) (*Template, error) {
	if _, err := t.Template.Parse(text); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseFiles parses a list of files as templates and associate them with t.
// Each file can be a standalone template.
func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
	if _, err := t.Template.ParseFiles(filenames...); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseGlob parses the files that match the given pattern as templates and
// associate them with t.
func (t *Template) ParseGlob(pattern string) (*Template, error) {
	if _, err := t.Template.ParseGlob(pattern); err != nil {
		return nil, err
	}
	return t, nil
}

// ParseDir walks on the given dir path and parses the given matches with aren't Go files.
func (t *Template) ParseDir(path string) (*Template, error) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walk path %s: %v", path, err)
		}
		if info.IsDir() || strings.HasSuffix(path, ".go") {
			return nil
		}
		_, err = t.ParseFiles(path)
		return err
	})
	return t, err
}

// AddParseTree adds the given parse tree to the template.
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error) {
	if _, err := t.Template.AddParseTree(name, tree); err != nil {
		return nil, err
	}
	return t, nil
}

// MustParse is a helper that wraps a call to a function returning (*Template, error)
// and panics if the error is non-nil.
func MustParse(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}

func pkgf(s string) func(t *Type) string {
	return func(t *Type) string { return fmt.Sprintf(s, t.Package()) }
}

// match reports if the given name matches the extended pattern.
func match(patterns []string, name string) bool {
	for _, pat := range patterns {
		matched, _ := filepath.Match(pat, name)
		if matched {
			return true
		}
	}
	return false
}
