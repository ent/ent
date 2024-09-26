// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"text/template/parse"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type (
	// TypeTemplate specifies a template that is executed with
	// each Type object of the graph.
	TypeTemplate struct {
		Name           string             // template name.
		Cond           func(*Type) bool   // condition to apply the template.
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
			Cond:   notView,
			Format: pkgf("%s_create.go"),
			ExtendPatterns: []string{
				"dialect/*/create/fields/additional/*",
				"dialect/*/create_bulk/fields/additional/*",
			},
		},
		{
			Name:   "update",
			Cond:   notView,
			Format: pkgf("%s_update.go"),
		},
		{
			Name:   "delete",
			Cond:   notView,
			Format: pkgf("%s_delete.go"),
		},
		{
			Name:   "query",
			Format: pkgf("%s_query.go"),
			ExtendPatterns: []string{
				"dialect/*/query/fields/additional/*",
			},
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
				return fmt.Sprintf("%[1]s/%[1]s.go", t.PackageDir())
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
				"dialect/*/query/fields/init/*",
			},
		},
		{
			Name:   "tx",
			Format: "tx.go",
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
			Name:   "intercept",
			Format: "intercept/intercept.go",
			Skip: func(g *Graph) bool {
				return !g.featureEnabled(FeatureIntercept)
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
	}
	// template files that were deleted and should be removed by the codegen.
	deletedTemplates = []string{"config.go", "context.go"}
	// patterns for extending partial-templates (included by other templates).
	partialPatterns = [...]string{
		"client/additional/*",
		"client/additional/*/*",
		"config/*/*",
		"config/*/*/*",
		"create/additional/*",
		"delete/additional/*",
		"dialect/*/*/*/spec/*",
		"dialect/*/*/spec/*",
		"dialect/*/config/*/*",
		"dialect/*/import/additional/*",
		"dialect/*/query/selector/*",
		"dialect/sql/create/additional/*",
		"dialect/sql/create_bulk/additional/*",
		"dialect/sql/meta/constants/*",
		"dialect/sql/model/additional/*",
		"dialect/sql/model/edges/*",
		"dialect/sql/model/edges/fields/additional/*",
		"dialect/sql/model/fields/*",
		"dialect/sql/select/additional/*",
		"dialect/sql/predicate/edge/*/*",
		"dialect/sql/query/additional/*",
		"dialect/sql/query/all/nodes/*",
		"dialect/sql/query/from/*",
		"dialect/sql/query/path/*",
		"dialect/sql/query/*/*/*",
		"import/additional/*",
		"model/additional/*",
		"model/comment/additional/*",
		"model/edges/fields/additional/*",
		"tx/additional/*",
		"tx/additional/*/*",
		"update/additional/*",
		"query/additional/*",
		"privacy/additional/*",
		"privacy/additional/*/*",
	}
	// templates holds the Go templates for the code generation.
	templates *Template
	//go:embed template/*
	templateDir embed.FS
	// importPkg are the import packages used for code generation.
	// Extended by the function below on generation initialization.
	importPkg = map[string]string{
		"context": "context",
		"driver":  "database/sql/driver",
		"errors":  "errors",
		"fmt":     "fmt",
		"math":    "math",
		"strings": "strings",
		"time":    "time",
		"ent":     "entgo.io/ent",
		"dialect": "entgo.io/ent/dialect",
		"field":   "entgo.io/ent/schema/field",
	}
)

// notView reports if the given type is not a view.
func notView(t *Type) bool { return !t.IsView() }

func initTemplates() {
	templates = MustParse(NewTemplate("templates").
		ParseFS(templateDir, "template/*.tmpl", "template/*/*.tmpl", "template/*/*/*.tmpl", "template/*/*/*/*.tmpl"))
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
	FuncMap   template.FuncMap
	condition func(*Graph) bool
}

// NewTemplate creates an empty template with the standard codegen functions.
func NewTemplate(name string) *Template {
	t := &Template{Template: template.New(name)}
	return t.Funcs(Funcs)
}

// Funcs merges the given funcMap with the template functions.
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

// SkipIf allows registering a function to determine if the template needs to be skipped or not.
func (t *Template) SkipIf(cond func(*Graph) bool) *Template {
	t.condition = cond
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
			return fmt.Errorf("walk path %s: %w", path, err)
		}
		if info.IsDir() || strings.HasSuffix(path, ".go") {
			return nil
		}
		_, err = t.ParseFiles(path)
		return err
	})
	return t, err
}

// ParseFS is like ParseFiles or ParseGlob but reads from the file system fsys
// instead of the host operating system's file system.
func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error) {
	if _, err := t.Template.ParseFS(fsys, patterns...); err != nil {
		return nil, err
	}
	return t, nil
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

type (
	// Dependencies wraps a list of dependencies as codegen
	// annotation.
	Dependencies []*Dependency

	// Dependency allows configuring optional dependencies as struct fields on the
	// generated builders. For example:
	//
	//	DependencyAnnotation{
	//		Field:	"HTTPClient",
	//		Type:	"*http.Client",
	//		Option:	"WithClient",
	//	}
	//
	// Although the Dependency and the DependencyAnnotation are exported, used should
	// use the entc.Dependency option in order to build this annotation.
	Dependency struct {
		// Field defines the struct field name on the builders.
		// It defaults to the full type name. For example:
		//
		//	http.Client	=> HTTPClient
		//	net.Conn	=> NetConn
		//	url.URL		=> URL
		//
		Field string
		// Type defines the type identifier. For example, `*http.Client`.
		Type *field.TypeInfo
		// Option defines the name of the config option.
		// It defaults to the field name.
		Option string
	}
)

// Name describes the annotation name.
func (Dependencies) Name() string {
	return "Dependencies"
}

// Merge implements the schema.Merger interface.
func (d Dependencies) Merge(other schema.Annotation) schema.Annotation {
	if deps, ok := other.(Dependencies); ok {
		return append(d, deps...)
	}
	return d
}

var _ interface {
	schema.Annotation
	schema.Merger
} = (*Dependencies)(nil)

// Build builds the annotation and fails if it is invalid.
func (d *Dependency) Build() error {
	if d.Type == nil {
		return errors.New("entc/gen: missing dependency type")
	}
	if d.Field == "" {
		name, err := d.defaultName()
		if err != nil {
			return err
		}
		d.Field = name
	}
	if d.Option == "" {
		d.Option = d.Field
	}
	return nil
}

func (d *Dependency) defaultName() (string, error) {
	var pkg, name string
	switch parts := strings.Split(strings.TrimLeft(d.Type.Ident, "[]*"), "."); len(parts) {
	case 1:
		name = parts[0]
	case 2:
		name = parts[1]
		// Avoid stuttering.
		if !strings.EqualFold(parts[0], name) {
			pkg = parts[0]
		}
	default:
		return "", fmt.Errorf("entc/gen: unexpected number of parts: %q", parts)
	}
	if r := d.Type.RType; r != nil && (r.Kind == reflect.Array || r.Kind == reflect.Slice) {
		name = plural(name)
	}
	return pascal(pkg) + pascal(name), nil
}

func pkgf(s string) func(t *Type) string {
	return func(t *Type) string { return fmt.Sprintf(s, t.PackageDir()) }
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
