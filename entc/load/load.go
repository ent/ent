// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package load is the interface for loading an ent/schema package into a Go program.
package load

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"text/template"
	"time"

	"entgo.io/ent"

	"golang.org/x/tools/go/packages"
)

type (
	// A SchemaSpec holds a serializable version of an ent.Schema
	// and its Go package and module information.
	SchemaSpec struct {
		// Schemas defines the loaded schema descriptors.
		Schemas []*Schema

		// PkgPath is the package path of the loaded
		// ent.Schema package.
		PkgPath string

		// Module defines the module information for
		// the user schema package if exists.
		Module *packages.Module
	}

	// Config holds the configuration for loading an ent/schema package.
	Config struct {
		// Path is the path for the schema package.
		Path string
		// Names are the schema names to load. Empty means all schemas in the directory.
		Names []string
		// BuildFlags are forwarded to the package.Config when
		// loading the schema package.
		BuildFlags []string
	}
)

// Load loads the schemas package and build the Go plugin with this info.
func (c *Config) Load() (*SchemaSpec, error) {
	spec, err := c.load()
	if err != nil {
		return nil, fmt.Errorf("entc/load: load schema dir: %w", err)
	}
	if len(c.Names) == 0 {
		return nil, fmt.Errorf("entc/load: no schema found in: %s", c.Path)
	}
	var b bytes.Buffer
	err = buildTmpl.ExecuteTemplate(&b, "main", struct {
		*Config
		Package string
	}{
		Config:  c,
		Package: spec.PkgPath,
	})
	if err != nil {
		return nil, fmt.Errorf("entc/load: execute template: %w", err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("entc/load: format template: %w", err)
	}
	if err := os.MkdirAll(".entc", os.ModePerm); err != nil {
		return nil, err
	}
	target := fmt.Sprintf(".entc/%s.go", filename(spec.PkgPath))
	if err := os.WriteFile(target, buf, 0644); err != nil {
		return nil, fmt.Errorf("entc/load: write file %s: %w", target, err)
	}
	defer os.RemoveAll(".entc")
	out, err := run(target, c.BuildFlags)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(out, "\n") {
		schema, err := UnmarshalSchema([]byte(line))
		if err != nil {
			return nil, fmt.Errorf("entc/load: unmarshal schema %s: %w", line, err)
		}
		spec.Schemas = append(spec.Schemas, schema)
	}
	return spec, nil
}

// entInterface holds the reflect.Type of ent.Interface.
var entInterface = reflect.TypeOf(struct{ ent.Interface }{}).Field(0).Type

// load loads the schemas info.
func (c *Config) load() (*SchemaSpec, error) {
	pkgs, err := packages.Load(&packages.Config{
		BuildFlags: c.BuildFlags,
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
	}, c.Path, entInterface.PkgPath())
	if err != nil {
		return nil, fmt.Errorf("loading package: %w", err)
	}
	if len(pkgs) < 2 {
		return nil, fmt.Errorf("missing package information for: %s", c.Path)
	}
	entPkg, pkg := pkgs[0], pkgs[1]
	if len(pkg.Errors) != 0 {
		return nil, pkg.Errors[0]
	}
	if pkgs[0].PkgPath != entInterface.PkgPath() {
		entPkg, pkg = pkgs[1], pkgs[0]
	}
	var names []string
	iface := entPkg.Types.Scope().Lookup(entInterface.Name()).Type().Underlying().(*types.Interface)
	for k, v := range pkg.TypesInfo.Defs {
		typ, ok := v.(*types.TypeName)
		if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
			continue
		}
		spec, ok := k.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return nil, fmt.Errorf("invalid declaration %T for %s", k.Obj.Decl, k.Name)
		}
		if _, ok := spec.Type.(*ast.StructType); !ok {
			return nil, fmt.Errorf("invalid spec type %T for %s", spec.Type, k.Name)
		}
		names = append(names, k.Name)
	}
	if len(c.Names) == 0 {
		c.Names = names
	}
	sort.Strings(c.Names)
	return &SchemaSpec{PkgPath: pkg.PkgPath, Module: pkg.Module}, nil
}

var (
	//go:embed template/main.tmpl schema.go
	files     embed.FS
	buildTmpl = templates()
)

func templates() *template.Template {
	tmpls, err := schemaTemplates()
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.New("templates").
		ParseFS(files, "template/main.tmpl"))
	for _, t := range tmpls {
		tmpl = template.Must(tmpl.Parse(t))
	}
	return tmpl
}

// schemaTemplates turns the schema.go file and its import block into templates.
func schemaTemplates() ([]string, error) {
	var (
		imports []string
		code    bytes.Buffer
		fset    = token.NewFileSet()
		src, _  = files.ReadFile("schema.go")
	)
	f, err := parser.ParseFile(fset, "schema.go", src, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse schema file: %w", err)
	}
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok && decl.Tok == token.IMPORT {
			for _, spec := range decl.Specs {
				imports = append(imports, spec.(*ast.ImportSpec).Path.Value)
			}
			continue
		}
		if err := format.Node(&code, fset, decl); err != nil {
			return nil, fmt.Errorf("format node: %w", err)
		}
		code.WriteByte('\n')
	}
	return []string{
		fmt.Sprintf(`{{ define "schema" }} %s {{ end }}`, code.String()),
		fmt.Sprintf(`{{ define "imports" }} %s {{ end }}`, strings.Join(imports, "\n")),
	}, nil
}

func filename(pkg string) string {
	name := strings.ReplaceAll(pkg, "/", "_")
	return fmt.Sprintf("entc_%s_%d", name, time.Now().Unix())
}

// run 'go run' command and return its output.
func run(target string, buildFlags []string) (string, error) {
	args := []string{"run"}
	args = append(args, buildFlags...)
	args = append(args, target)
	cmd := exec.Command("go", args...)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("entc/load: %s", stderr)
	}
	return stdout.String(), nil
}
