// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package load is the interface for loading an ent/schema package into a Go program.
package load

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"entgo.io/ent"

	"golang.org/x/tools/go/ast/astutil"
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
	spec, pos, err := c.load()
	if err != nil {
		return nil, fmt.Errorf("entc/load: parse schema dir: %w", err)
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
	out, err := gorun(target, c.BuildFlags)
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
	for _, s := range spec.Schemas {
		s.Pos = pos[s.Name]
	}
	return spec, nil
}

// entInterface holds the reflect.Type of ent.Interface.
var entInterface = reflect.TypeOf(struct{ ent.Interface }{}).Field(0).Type

// load the ent/schema info.
func (c *Config) load() (*SchemaSpec, map[string]string, error) {
	pkgs, err := packages.Load(&packages.Config{
		BuildFlags: c.BuildFlags,
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
	}, c.Path, entInterface.PkgPath())
	if err != nil {
		return nil, nil, fmt.Errorf("loading package: %w", err)
	}
	if len(pkgs) < 2 {
		// Check if the package loading failed due to Go-related
		// errors, such as 'missing go.sum entry'.
		if err := golist(c.Path, c.BuildFlags); err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("missing package information for: %s", c.Path)
	}
	entPkg, pkg := pkgs[0], pkgs[1]
	if len(pkg.Errors) != 0 {
		return nil, nil, c.loadError(pkg.Errors[0])
	}
	if len(entPkg.Errors) != 0 {
		return nil, nil, entPkg.Errors[0]
	}
	if pkgs[0].PkgPath != entInterface.PkgPath() {
		entPkg, pkg = pkgs[1], pkgs[0]
	}
	names := make(map[string]string)
	iface := entPkg.Types.Scope().Lookup(entInterface.Name()).Type().Underlying().(*types.Interface)
	for k, v := range pkg.TypesInfo.Defs {
		typ, ok := v.(*types.TypeName)
		if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
			continue
		}
		spec, ok := k.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return nil, nil, fmt.Errorf("invalid declaration %T for %s", k.Obj.Decl, k.Name)
		}
		if _, ok := spec.Type.(*ast.StructType); !ok {
			return nil, nil, fmt.Errorf("invalid spec type %T for %s", spec.Type, k.Name)
		}
		p := pkg.Fset.Position(spec.Pos())
		names[k.Name] = fmt.Sprintf("%s:%d", p.Filename, p.Line)
	}
	if len(c.Names) == 0 {
		c.Names = slices.Sorted(maps.Keys(names))
	} else {
		sort.Strings(c.Names)
	}
	return &SchemaSpec{PkgPath: pkg.PkgPath, Module: pkg.Module}, names, nil
}

func (c *Config) loadError(perr packages.Error) (err error) {
	if strings.Contains(perr.Msg, "import cycle not allowed") {
		if cause := c.cycleCause(); cause != "" {
			perr.Msg += "\n" + cause
		}
	}
	err = perr
	if perr.Pos == "" {
		// Strip "-:" prefix in case of empty position.
		err = errors.New(perr.Msg)
	}
	return err
}

func (c *Config) cycleCause() (cause string) {
	dir, err := parser.ParseDir(token.NewFileSet(), c.Path, nil, 0)
	// Ignore reporting in case of parsing
	// error, or there no packages to parse.
	if err != nil || len(dir) == 0 {
		return
	}
	// Find the package that contains the schema, or
	// extract the first package if there is only one.
	pkg := dir[filepath.Base(c.Path)]
	if pkg == nil {
		for _, v := range dir {
			pkg = v
			break
		}
	}
	// Package local declarations used by schema fields.
	locals := make(map[string]bool)
	for _, f := range pkg.Files {
		for _, d := range f.Decls {
			g, ok := d.(*ast.GenDecl)
			if !ok || g.Tok != token.TYPE {
				continue
			}
			for _, s := range g.Specs {
				ts, ok := s.(*ast.TypeSpec)
				if !ok || !ts.Name.IsExported() {
					continue
				}
				// Non-struct types such as "type Role int".
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					locals[ts.Name.Name] = true
					continue
				}
				var embedSchema bool
				astutil.Apply(st.Fields, func(c *astutil.Cursor) bool {
					f, ok := c.Node().(*ast.Field)
					if ok {
						switch x := f.Type.(type) {
						case *ast.SelectorExpr:
							if x.Sel.Name == "Schema" || x.Sel.Name == "Mixin" {
								embedSchema = true
							}
						case *ast.Ident:
							// A common pattern is to create local base schema to be embedded by other schemas.
							if name := strings.ToLower(x.Name); name == "schema" || name == "mixin" {
								embedSchema = true
							}
						}
					}
					// Stop traversing the AST in case an ~ent.Schema is embedded.
					return !embedSchema
				}, nil)
				if !embedSchema {
					locals[ts.Name.Name] = true
				}
			}
		}
	}
	// No local declarations to report.
	if len(locals) == 0 {
		return
	}
	// Usage of local declarations by schema fields.
	goTypes := make(map[string]bool)
	for _, f := range pkg.Files {
		for _, d := range f.Decls {
			f, ok := d.(*ast.FuncDecl)
			if !ok || f.Name.Name != "Fields" || f.Type.Params.NumFields() != 0 || f.Type.Results.NumFields() != 1 {
				continue
			}
			astutil.Apply(f.Body, func(cursor *astutil.Cursor) bool {
				i, ok := cursor.Node().(*ast.Ident)
				if ok && locals[i.Name] {
					goTypes[i.Name] = true
				}
				return true
			}, nil)
		}
	}
	names := make([]string, 0, len(goTypes))
	for k := range goTypes {
		names = append(names, strconv.Quote(k))
	}
	sort.Strings(names)
	if len(names) > 0 {
		cause = fmt.Sprintf("To resolve this issue, move the custom types used by the generated code to a separate package: %s", strings.Join(names, ", "))
	}
	return
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
func gorun(target string, buildFlags []string) (string, error) {
	s, err := gocmd("run", target, buildFlags)
	if err != nil {
		return "", fmt.Errorf("entc/load: %s", err)
	}
	return s, nil
}

// golist checks if 'go list' can be executed on the given target.
func golist(target string, buildFlags []string) error {
	_, err := gocmd("list", target, buildFlags)
	return err
}

// goCmd runs a go command and returns its output.
func gocmd(command, target string, buildFlags []string) (string, error) {
	args := []string{command}
	args = append(args, buildFlags...)
	args = append(args, target)
	cmd := exec.Command("go", args...)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", errors.New(strings.TrimSpace(stderr.String()))
	}
	return stdout.String(), nil
}
