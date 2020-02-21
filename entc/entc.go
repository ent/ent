// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package entc provides an interface for interacting with
// entc (ent codegen) as a package rather than an executable.
package entc

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/facebookincubator/ent/entc/gen"
	"github.com/facebookincubator/ent/entc/load"
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
func Generate(schemaPath string, cfg *gen.Config, options ...Option) error {
	if cfg.Target == "" {
		abs, err := filepath.Abs(schemaPath)
		if err != nil {
			return err
		}
		// default target-path for codegen is one dir above
		// the schema.
		cfg.Target = filepath.Dir(abs)
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
	graph, err := LoadGraph(schemaPath, cfg)
	if err != nil {
		return err
	}
	return graph.Gen()
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

// TemplateFiles parses the named files and associates the resulting templates
// with codegen templates.
func TemplateFiles(filenames ...string) Option {
	return templateOption(func(cfg *gen.Config) (err error) {
		cfg.Template, err = cfg.Template.ParseFiles(filenames...)
		return
	})
}

// TemplateGlob parses the template definitions from the files identified
// by the pattern and associates the resulting templates with codegen templates.
func TemplateGlob(pattern string) Option {
	return templateOption(func(cfg *gen.Config) (err error) {
		cfg.Template, err = cfg.Template.ParseGlob(pattern)
		return
	})
}

// TemplateDir parses the template definitions from the files in the directory
// and associates the resulting templates with codegen templates.
func TemplateDir(path string) Option {
	return templateOption(func(cfg *gen.Config) error {
		return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("load template: %v", err)
			}
			if info.IsDir() {
				return nil
			}
			cfg.Template, err = cfg.Template.ParseFiles(path)
			return err
		})
	})
}

// templateOption ensures the template instantiate
// once for config and execute the given Option.
func templateOption(next Option) Option {
	return func(cfg *gen.Config) (err error) {
		if cfg.Template == nil {
			cfg.Template = template.New("external").Funcs(gen.Funcs)
		}
		return next(cfg)
	}
}
