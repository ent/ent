// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package entc provides an interface for interacting with
// entc (ent codegen) as a package rather than an executable.
package entc

import (
	"errors"
	"fmt"
	"go/token"
	"path"
	"path/filepath"
	"strings"

	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/entc/internal"
	"github.com/facebook/ent/entc/load"

	"golang.org/x/tools/go/packages"
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
	return generate(schemaPath, cfg)
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

// FeatureNames enables sets of features by their names.
func FeatureNames(names ...string) Option {
	return func(cfg *gen.Config) error {
		for _, name := range names {
			for _, feat := range gen.AllFeatures {
				if name == feat.Name {
					cfg.Features = append(cfg.Features, feat)
				}
			}
		}
		return nil
	}
}

// TemplateFiles parses the named files and associates the resulting templates
// with codegen templates.
func TemplateFiles(filenames ...string) Option {
	return templateOption(func(t *gen.Template) (*gen.Template, error) {
		return t.ParseFiles(filenames...)
	})
}

// TemplateGlob parses the template definitions from the files identified
// by the pattern and associates the resulting templates with codegen templates.
func TemplateGlob(pattern string) Option {
	return templateOption(func(t *gen.Template) (*gen.Template, error) {
		return t.ParseGlob(pattern)
	})
}

// TemplateDir parses the template definitions from the files in the directory
// and associates the resulting templates with codegen templates.
func TemplateDir(path string) Option {
	return templateOption(func(t *gen.Template) (*gen.Template, error) {
		return t.ParseDir(path)
	})
}

// templateOption ensures the template instantiate
// once for config and execute the given Option.
func templateOption(next func(t *gen.Template) (*gen.Template, error)) Option {
	return func(cfg *gen.Config) (err error) {
		tmpl, err := next(gen.NewTemplate("external"))
		if err != nil {
			return err
		}
		cfg.Templates = append(cfg.Templates, tmpl)
		return nil
	}
}

// generate loads the given schema and run codegen.
func generate(schemaPath string, cfg *gen.Config) error {
	graph, err := LoadGraph(schemaPath, cfg)
	if err != nil {
		if err := mayRecover(err, schemaPath, cfg); err != nil {
			return err
		}
		if graph, err = LoadGraph(schemaPath, cfg); err != nil {
			return err
		}
	}
	if err := normalizePkg(cfg); err != nil {
		return err
	}
	return graph.Gen()
}

func mayRecover(err error, schemaPath string, cfg *gen.Config) error {
	if enabled, _ := cfg.FeatureEnabled(gen.FeatureSnapshot.Name); !enabled {
		return err
	}
	if errors.As(err, &packages.Error{}) || !internal.IsBuildError(err) {
		return err
	}
	// If the build error comes from the schema package.
	if err := internal.CheckDir(schemaPath); err != nil {
		return fmt.Errorf("schema failure: %w", err)
	}
	target := filepath.Join(cfg.Target, "internal/schema.go")
	return (&internal.Snapshot{Path: target, Config: cfg}).Restore()
}
