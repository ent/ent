// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package load is the interface for loading schema package into a Go program.
package load

import (
	"bytes"
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
	"github.com/facebookincubator/ent/entc/load/internal"

	"golang.org/x/tools/go/packages"
)

// A SchemaSpec holds an ent.schema package that created by Load.
type SchemaSpec struct {
	// Schemas are the schema descriptors.
	Schemas []*Schema
	// PkgPath is the package path of the schema.
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

// Load loads the schemas package and build the Go plugin with this info.
func (c *Config) Load() (*SchemaSpec, error) {
	pkgPath, err := c.load()
	if err != nil {
		return nil, fmt.Errorf("load schemas dir: %v", err)
	}
	if len(c.Names) == 0 {
		return nil, fmt.Errorf("no schema found in: %s", c.Path)
	}
	b := bytes.NewBuffer(nil)
	err = buildTmpl.ExecuteTemplate(b, "main", struct {
		*Config
		Package string
	}{c, pkgPath})
	if err != nil {
		return nil, fmt.Errorf("execute template: %v", err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format template: %v", err)
	}
	if err := os.MkdirAll(".entc", os.ModePerm); err != nil {
		return nil, err
	}
	target := fmt.Sprintf(".entc/%s.go", filename(pkgPath))
	if err := ioutil.WriteFile(target, buf, 0644); err != nil {
		return nil, fmt.Errorf("write file %s: %v", target, err)
	}
	defer os.RemoveAll(".entc")
	out, err := run(target)
	if err != nil {
		return nil, err
	}
	spec := &SchemaSpec{PkgPath: pkgPath}
	for _, line := range strings.Split(out, "\n") {
		schema, err := UnmarshalSchema([]byte(line))
		if err != nil {
			return nil, fmt.Errorf("unmarshal schema %s: %v", line, err)
		}
		spec.Schemas = append(spec.Schemas, schema)
	}
	return spec, nil
}

// entInterface holds the reflect.Type of ent.Interface.
var entInterface = reflect.TypeOf(struct{ ent.Interface }{}).Field(0).Type

// load loads the schemas info.
func (c *Config) load() (string, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, c.Path, entInterface.PkgPath())
	if err != nil {
		return "", fmt.Errorf("loading package: %v", err)
	}
	if len(pkgs) < 2 {
		return "", fmt.Errorf("missing package information for: %s", c.Path)
	}
	entPkg, pkg := pkgs[0], pkgs[1]
	if len(pkg.Errors) != 0 {
		return "", pkg.Errors[0]
	}
	if pkgs[0].PkgPath != entInterface.PkgPath() {
		entPkg, pkg = pkgs[1], pkgs[0]
	}
	names := make([]string, 0)
	iface := entPkg.Types.Scope().Lookup(entInterface.Name()).Type().Underlying().(*types.Interface)
	for k, v := range pkg.TypesInfo.Defs {
		typ, ok := v.(*types.TypeName)
		if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
			continue
		}
		spec, ok := k.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return "", fmt.Errorf("invalid declaration %T for %s", k.Obj.Decl, k.Name)
		}
		if _, ok := spec.Type.(*ast.StructType); !ok {
			return "", fmt.Errorf("invalid spec type %T for %s", spec.Type, k.Name)
		}
		names = append(names, k.Name)
	}
	if len(c.Names) == 0 {
		c.Names = names
	}
	sort.Strings(c.Names)
	return pkg.PkgPath, nil
}

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -pkg=internal -o=internal/bindata.go -modtime=1 ./template/... schema.go

var buildTmpl = templates()

func templates() *template.Template {
	tmpl := template.New("templates").Funcs(template.FuncMap{"base": filepath.Base})
	tmpl = template.Must(tmpl.Parse(string(internal.MustAsset("template/main.tmpl"))))
	// Turns the schema file and its imports into templates.
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
	f, err := parser.ParseFile(fset, name, string(internal.MustAsset(name)), parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse file: %s: %v", name, err)
	}
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
			for _, spec := range decl.Specs {
				imports = append(imports, spec.(*ast.ImportSpec).Path.Value)
			}
			continue
		}
		if err := format.Node(&code, fset, decl); err != nil {
			return nil, fmt.Errorf("format node: %v", err)
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
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("entc/load: %s", stderr)
	}
	return stdout.String(), nil
}
