// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package entc provides an interface for interacting with
// entc (ent codegen) as a package rather than an executable.
package entc

import (
	"fmt"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/entc/load"
	"github.com/facebook/ent/schema/field"
)

// LoadGraph loads the schema package from the given schema path,
// and constructs a *gen.Graph.
func LoadGraph(schemaPath string, cfg *gen.Config) (*gen.Graph, error) {
	spec, err := (&load.Config{Path: schemaPath}).Load()
	if err != nil {
		return nil, err
	}
	cfg.Schema = spec.PkgPath
	if cfg.Package == "" {
		// default package-path for codegen is one package
		// before the schema package (`<project>/ent/schema`).
		cfg.Package = path.Dir(spec.PkgPath)
	}
	return gen.NewGraph(cfg, spec.Schemas...)
}

// Generate runs the codegen on the schema path. The default target
// directory for the assets, is one directory above the schema path.
// Hence, if the schema package resides in "<project>/ent/schema",
// the base directory for codegen will be "<project>/ent".
//
// If no storage driver provided by option, SQL driver will be used.
//
//	entc.Generate("./ent/path", &gen.Config{
//		Header: "// Custom header",
//		IDType: &field.TypeInfo{Type: field.TypeInt},
//	})
//
func Generate(schemaPath string, cfg *gen.Config, options ...Option) (err error) {
	if cfg.Target == "" {
		abs, err := filepath.Abs(schemaPath)
		if err != nil {
			return err
		}
		// default target-path for codegen is one dir above
		// the schema.
		cfg.Target = filepath.Dir(abs)
	}
	if cfg.IDType == nil {
		cfg.IDType = &field.TypeInfo{Type: field.TypeInt}
	}
	for _, opt := range options {
		if err := opt(cfg); err != nil {
			return err
		}
	}
	if cfg.Storage == nil {
		driver, err := gen.NewStorage("sql")
		if err != nil {
			return err
		}
		cfg.Storage = driver
	}
	undo, err := gen.PrepareEnv(cfg)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = undo()
		}
	}()
	graph, err := LoadGraph(schemaPath, cfg)
	if err != nil {
		return err
	}
	if err := normalizePkg(cfg); err != nil {
		return err
	}
	return graph.Gen()
}

func normalizePkg(c *gen.Config) error {
	base := path.Base(c.Package)
	if strings.ContainsRune(base, '-') {
		base = strings.ReplaceAll(base, "-", "_")
		c.Package = path.Join(path.Dir(c.Package), base)
	}
	if !token.IsIdentifier(base) {
		return fmt.Errorf("invalid package identifier: %q", base)
	}
	return nil
}

// Option allows for managing codegen configuration using functional options.
type Option func(*gen.Config) error

// Storage sets the storage-driver type to support by the codegen.
func Storage(typ string) Option {
	return func(cfg *gen.Config) error {
		storage, err := gen.NewStorage(typ)
		if err != nil {
			return err
		}
		cfg.Storage = storage
		return nil
	}
}

// Funcs specifies external functions to add to the template execution.
func Funcs(funcMap template.FuncMap) Option {
	return func(cfg *gen.Config) error {
		if cfg.Funcs == nil {
			cfg.Funcs = funcMap
			return nil
		}
		for name, fn := range funcMap {
			cfg.Funcs[name] = fn
		}
		return nil
	}
}

// TemplateFiles parses the named files and associates the resulting templates
// with codegen templates.
func TemplateFiles(filenames ...string) Option {
	return templateOption(func(t *template.Template) (*template.Template, error) {
		return t.ParseFiles(filenames...)
	})
}

// TemplateGlob parses the template definitions from the files identified
// by the pattern and associates the resulting templates with codegen templates.
func TemplateGlob(pattern string) Option {
	return templateOption(func(t *template.Template) (*template.Template, error) {
		return t.ParseGlob(pattern)
	})
}

// TemplateDir parses the template definitions from the files in the directory
// and associates the resulting templates with codegen templates.
func TemplateDir(path string) Option {
	return templateOption(func(t *template.Template) (*template.Template, error) {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("load template: %v", err)
			}
			if info.IsDir() || strings.HasSuffix(path, ".go") {
				return nil
			}
			t, err = t.ParseFiles(path)
			return err
		})
		if err != nil {
			return nil, err
		}
		return t, nil
	})
}

// templateOption ensures the template instantiate
// once for config and execute the given Option.
func templateOption(next func(t *template.Template) (*template.Template, error)) Option {
	return func(cfg *gen.Config) (err error) {
		tmpl, err := next(template.New("external").Funcs(gen.Funcs).Funcs(cfg.Funcs))
		if err != nil {
			return err
		}
		cfg.Templates = append(cfg.Templates, tmpl)
		return nil
	}
}
