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
	"github.com/facebookincubator/ent/entc/load/internal"

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
	// schema types and their exported struct fields.
	fields map[string][]*StructField
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
		schema.StructFields = c.fields[schema.Name]
		spec.Schemas = append(spec.Schemas, schema)
	}
	return spec, nil
}

// entInterface represents the the ent.Interface type.
var entInterface = reflect.TypeOf(struct{ ent.Interface }{}).Field(0).Type

// load loads the schemas info.
func (c *Config) load() (string, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, c.Path, entInterface.PkgPath())
	if err != nil {
		return "", err
	}
	entPkg, pkg := pkgs[0], pkgs[1]
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
		specType, ok := spec.Type.(*ast.StructType)
		if !ok {
			return "", fmt.Errorf("invalid spec type %T for %s", spec.Type, k.Name)
		}
		if err := c.structFields(k.Name, v, specType); err != nil {
			return "", err
		}
		names = append(names, k.Name)
	}
	if len(c.Names) == 0 {
		c.Names = names
	}
	sort.Strings(c.Names)
	return pkg.PkgPath, err
}

// structFields loads schema type fields if exist.
func (c *Config) structFields(name string, obj types.Object, spec *ast.StructType) (err error) {
	typ, ok := obj.(*types.TypeName)
	if !ok {
		return
	}
	st, ok := typ.Type().Underlying().(*types.Struct)
	if !ok {
		return
	}
	if c.fields == nil {
		c.fields = make(map[string][]*StructField)
	}
	for i := 0; i < st.NumFields(); i++ {
		f := st.Field(i)
		// skip non-exported fields, because they
		// cannot be used outside the package.
		if !f.Exported() {
			continue
		}
		sf := &StructField{
			Tag:      st.Tag(i),
			Name:     f.Name(),
			Type:     f.Type().String(),
			Embedded: f.Embedded(),
			Comment:  strings.TrimSpace(spec.Fields.List[i].Comment.Text()),
		}
		switch typ := indirectType(f.Type()).(type) {
		case *types.Named:
			sf.PkgPath = typ.Obj().Pkg().Path()
			// skip fields used for schema definition.
			if sf.PkgPath == entInterface.PkgPath() {
				continue
			}
			c.fields[name] = append(c.fields[name], sf)
		default:
			if f.Embedded() {
				return fmt.Errorf("field %s for schema %q cannot be embbeded", f.Type(), name)
			}
			c.fields[name] = append(c.fields[name], sf)
		}
	}
	return
}

// indirectType returns the type at the end of indirection.
func indirectType(typ types.Type) types.Type {
	for {
		ptr, ok := typ.(*types.Pointer)
		if !ok {
			return typ
		}
		typ = ptr.Elem()
	}
}

//go:generate go-bindata -pkg=internal -o=internal/bindata.go ./template/... schema.go

var buildTmpl = templates()

func templates() *template.Template {
	tmpl := template.New("templates").Funcs(template.FuncMap{"base": filepath.Base})
	tmpl = template.Must(tmpl.Parse(string(internal.MustAsset("template/main.tmpl"))))
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
	f, err := parser.ParseFile(fset, name, string(internal.MustAsset(name)), parser.AllErrors)
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
