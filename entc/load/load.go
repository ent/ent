// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package load is the interface for loading schema package into a Go program.
package load

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/facebookincubator/ent"

	"github.com/pkg/errors"
	"golang.org/x/tools/go/packages"
)

// A SchemaSpec holds an ent.schema package that created by Load.
type SchemaSpec struct {
	// Schemas are the schema descriptors.
	Schemas []*Schema
	// PkgPath is the path where the schema package reside.
	// Note that path can be either a package path (e.g. github.com/a8m/x)
	// or a filepath (e.g. ./ent/schema).
	PkgPath string
}

// Config holds the configuration for package building.
type Config struct {
	// Path is the path for the schema package.
	Path string
	// Names are the schema names to run the code generation on.
	// Empty means all schemas in the directory.
	Names []string
}

// Build loads the schemas package and build the Go plugin with this info.
func (c *Config) Load() (*SchemaSpec, error) {
	pkgPath, err := c.load()
	if err != nil {
		return nil, errors.WithMessage(err, "load schemas dir")
	}
	if len(c.Names) == 0 {
		return nil, errors.Errorf("no schema found in: %s", c.Path)
	}
	b := bytes.NewBuffer(nil)
	err = buildTmpl.ExecuteTemplate(b, "main", struct {
		*Config
		Package string
	}{c, pkgPath})
	if err != nil {
		return nil, errors.WithMessage(err, "execute template")
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return nil, errors.WithMessage(err, "format template")
	}
	target := fmt.Sprintf("%s.go", filename(pkgPath))
	if err := ioutil.WriteFile(target, buf, 0644); err != nil {
		return nil, errors.WithMessagef(err, "write file %s", target)
	}
	defer os.Remove(target)
	out, err := run(target)
	if err != nil {
		return nil, err
	}
	spec := &SchemaSpec{PkgPath: pkgPath}
	for _, line := range strings.Split(out, "\n") {
		schema := &Schema{}
		if err := json.Unmarshal([]byte(line), schema); err != nil {
			return nil, errors.WithMessagef(err, "unmarshal schema %s", line)
		}
		spec.Schemas = append(spec.Schemas, schema)
	}
	return spec, nil
}

// load loads the schemas info.
func (c *Config) load() (string, error) {
	// get the ent package info statically instead of dealing with string constants
	// in the code, since import is handled by goimports and renaming should be easy.
	entface := reflect.TypeOf(struct{ ent.Interface }{}).Field(0).Type
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, c.Path, entface.PkgPath())
	if err != nil {
		return "", err
	}
	entPkg, pkg := pkgs[0], pkgs[1]
	if pkgs[0].PkgPath != entface.PkgPath() {
		entPkg, pkg = pkgs[1], pkgs[0]
	}
	names := make([]string, 0)
	iface := entPkg.Types.Scope().Lookup(entface.Name()).Type().Underlying().(*types.Interface)
	for k, v := range pkg.TypesInfo.Defs {
		typ, ok := v.(*types.TypeName)
		if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
			continue
		}
		names = append(names, k.Name)
	}
	if len(c.Names) == 0 {
		c.Names = names
	}
	sort.Strings(c.Names)
	return pkg.PkgPath, err
}

//go:generate go-bindata -pkg=load ./template/... schema.go

var buildTmpl = templates()

func templates() *template.Template {
	tmpl := template.New("templates").Funcs(template.FuncMap{"base": filepath.Base})
	tmpl = template.Must(tmpl.Parse(string(MustAsset("template/main.tmpl"))))
	// turns the schema file and its imports into templates.
	tmpls, err := schemaTemplates()
	if err != nil {
		panic(err)
	}
	for _, t := range tmpls {
		tmpl = template.Must(tmpl.Parse(t))
	}
	return tmpl
}

// schemaTemplates returns the templates needed for loading the schema.go file.
func schemaTemplates() ([]string, error) {
	const name = "schema.go"
	var (
		imports []string
		code    bytes.Buffer
		fset    = token.NewFileSet()
	)
	f, err := parser.ParseFile(fset, name, string(MustAsset(name)), parser.AllErrors)
	if err != nil {
		return nil, errors.WithMessagef(err, "parse file: %s", name)
	}
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
			for _, spec := range decl.Specs {
				imports = append(imports, spec.(*ast.ImportSpec).Path.Value)
			}
			continue
		}
		if err := format.Node(&code, fset, decl); err != nil {
			return nil, errors.WithMessage(err, "format node")
		}
		code.WriteByte('\n')
	}
	return []string{
		fmt.Sprintf(`{{ define "schema" }} %s {{ end }}`, code.String()),
		fmt.Sprintf(`{{ define "imports" }} %s {{ end }}`, strings.Join(imports, "\n")),
	}, nil
}

func filename(pkg string) string {
	name := strings.Replace(pkg, "/", "_", -1)
	return fmt.Sprintf("entc_%s_%d", name, time.Now().Unix())
}

// run 'go run' command and return its output.
func run(target string) (string, error) {
	cmd := exec.Command("go", "run", target)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("entc/load: %s", stderr)
	}
	return stdout.String(), nil
}
