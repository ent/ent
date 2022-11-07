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
	"reflect"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/internal"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"golang.org/x/tools/go/packages"
)

// LoadGraph loads the schema package from the given schema path,
// and constructs a *gen.Graph.
func LoadGraph(schemaPath string, cfg *gen.Config) (*gen.Graph, error) {
	spec, err := (&load.Config{Path: schemaPath, BuildFlags: cfg.BuildFlags}).Load()
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

// Annotation is used to attach arbitrary metadata to the schema objects in codegen.
// Unlike schema annotations, being serializable to JSON raw value is not mandatory.
//
// Template extensions can retrieve this metadata and use it inside their execution.
// Read more about it in ent website: https://entgo.io/docs/templates/#annotations.
type Annotation = schema.Annotation

// Annotations appends the given annotations to the codegen config.
func Annotations(annotations ...Annotation) Option {
	return func(cfg *gen.Config) error {
		if cfg.Annotations == nil {
			cfg.Annotations = gen.Annotations{}
		}
		for _, ant := range annotations {
			name := ant.Name()
			if curr, ok := cfg.Annotations[name]; !ok {
				cfg.Annotations[name] = ant
			} else if m, ok := curr.(schema.Merger); ok {
				cfg.Annotations[name] = m.Merge(ant)
			} else {
				return fmt.Errorf("duplicate annotations with name %q", name)
			}
		}
		return nil
	}
}

// BuildFlags appends the given build flags to the codegen config.
func BuildFlags(flags ...string) Option {
	return func(cfg *gen.Config) error {
		cfg.BuildFlags = append(cfg.BuildFlags, flags...)
		return nil
	}
}

// BuildTags appends the given build tags as build flags to the codegen
// config.
func BuildTags(tags ...string) Option {
	return BuildFlags("-tags", strings.Join(tags, ","))
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

// Extension describes an Ent code generation extension that
// allows customizing the code generation and integrate with
// other tools and libraries (e.g. GraphQL, gRPC, OpenAPI) by
// registering hooks, templates and global annotations in one
// simple call.
//
//	ex, err := entgql.NewExtension(
//		entgql.WithConfig("../gqlgen.yml"),
//		entgql.WithSchema("../schema.graphql"),
//	)
//	if err != nil {
//		log.Fatalf("creating graphql extension: %v", err)
//	}
//	err = entc.Generate("./schema", &gen.Config{
//		Templates: entswag.Templates,
//	}, entc.Extensions(ex))
//	if err != nil {
//		log.Fatalf("running ent codegen: %v", err)
//	}
//
type Extension interface {
	// Hooks holds an optional list of Hooks to apply
	// on the graph before/after the code-generation.
	Hooks() []gen.Hook

	// Annotations injects global annotations to the gen.Config object that
	// can be accessed globally in all templates. Unlike schema annotations,
	// being serializable to JSON raw value is not mandatory.
	//
	//	{{- with $.Config.Annotations.GQL }}
	//		{{/* Annotation usage goes here. */}}
	//	{{- end }}
	//
	Annotations() []Annotation

	// Templates specifies a list of alternative templates
	// to execute or to override the default.
	Templates() []*gen.Template

	// Options specifies a list of entc.Options to evaluate on
	// the gen.Config before executing the code generation.
	Options() []Option
}

// Extensions evaluates the list of Extensions on the gen.Config.
func Extensions(extensions ...Extension) Option {
	return func(cfg *gen.Config) error {
		for _, ex := range extensions {
			cfg.Hooks = append(cfg.Hooks, ex.Hooks()...)
			cfg.Templates = append(cfg.Templates, ex.Templates()...)
			for _, opt := range ex.Options() {
				if err := opt(cfg); err != nil {
					return err
				}
			}
			if err := Annotations(ex.Annotations()...)(cfg); err != nil {
				return err
			}
		}
		return nil
	}
}

// DefaultExtension is the default implementation for entc.Extension.
//
// Embedding this type allow third-party packages to create extensions
// without implementing all methods.
//
//	type Extension struct {
//		entc.DefaultExtension
//	}
//
type DefaultExtension struct{}

// Hooks of the extensions.
func (DefaultExtension) Hooks() []gen.Hook { return nil }

// Annotations of the extensions.
func (DefaultExtension) Annotations() []Annotation { return nil }

// Templates of the extensions.
func (DefaultExtension) Templates() []*gen.Template { return nil }

// Options of the extensions.
func (DefaultExtension) Options() []Option { return nil }

var _ Extension = (*DefaultExtension)(nil)

// DependencyOption allows configuring optional dependencies using functional options.
type DependencyOption func(*gen.Dependency) error

// DependencyType sets the type of the struct field in
// the generated builders for the configured dependency.
func DependencyType(v any) DependencyOption {
	return func(d *gen.Dependency) error {
		if v == nil {
			return errors.New("nil dependency type")
		}
		t := reflect.TypeOf(v)
		tv := indirect(t)
		d.Type = &field.TypeInfo{
			Ident:   t.String(),
			PkgPath: tv.PkgPath(),
			RType: &field.RType{
				Kind:    t.Kind(),
				Name:    tv.Name(),
				Ident:   tv.String(),
				PkgPath: tv.PkgPath(),
			},
		}
		return nil
	}
}

// DependencyTypeInfo is similar to DependencyType, but
// allows setting the field.TypeInfo explicitly.
func DependencyTypeInfo(t *field.TypeInfo) DependencyOption {
	return func(d *gen.Dependency) error {
		if t == nil {
			return errors.New("nil dependency type info")
		}
		d.Type = t
		return nil
	}
}

// DependencyName sets the struct field and the option name
// of the dependency in the generated builders.
func DependencyName(name string) DependencyOption {
	return func(d *gen.Dependency) error {
		d.Field = name
		d.Option = name
		return nil
	}
}

// Dependency allows configuring optional dependencies as struct fields on the
// generated builders. For example:
//
//	opts := []entc.Option{
//		entc.Dependency(
//			entc.DependencyType(&http.Client{}),
//		),
//		entc.Dependency(
//			entc.DependencyName("DB"),
//			entc.DependencyType(&sql.DB{}),
//		)
//	}
//	if err := entc.Generate("./ent/path", &gen.Config{}, opts...); err != nil {
//		log.Fatalf("running ent codegen: %v", err)
//	}
//
func Dependency(opts ...DependencyOption) Option {
	return func(cfg *gen.Config) error {
		d := &gen.Dependency{}
		for _, opt := range opts {
			if err := opt(d); err != nil {
				return err
			}
		}
		if err := d.Build(); err != nil {
			return err
		}
		return Annotations(gen.Dependencies{d})(cfg)
	}
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
	if !errors.As(err, &packages.Error{}) && !internal.IsBuildError(err) {
		return err
	}
	// If the build error comes from the schema package.
	if err := internal.CheckDir(schemaPath); err != nil {
		return fmt.Errorf("schema failure: %w", err)
	}
	target := filepath.Join(cfg.Target, "internal/schema.go")
	return (&internal.Snapshot{Path: target, Config: cfg}).Restore()
}

// indirect returns the type at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
