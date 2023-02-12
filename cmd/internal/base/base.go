// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package base defines shared basic pieces of the ent command.
package base

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"entgo.io/ent/cmd/internal/printer"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"

	"github.com/spf13/cobra"
)

// IDType is a custom ID implementation for pflag.
type IDType field.Type

// Set implements the Set method of the flag.Value interface.
func (t *IDType) Set(s string) error {
	switch s {
	case field.TypeInt.String():
		*t = IDType(field.TypeInt)
	case field.TypeInt64.String():
		*t = IDType(field.TypeInt64)
	case field.TypeUint.String():
		*t = IDType(field.TypeUint)
	case field.TypeUint64.String():
		*t = IDType(field.TypeUint64)
	case field.TypeString.String():
		*t = IDType(field.TypeString)
	default:
		return fmt.Errorf("invalid type %q", s)
	}
	return nil
}

// Type returns the type representation of the id option for help command.
func (IDType) Type() string {
	return fmt.Sprintf("%v", []field.Type{
		field.TypeInt,
		field.TypeInt64,
		field.TypeUint,
		field.TypeUint64,
		field.TypeString,
	})
}

// String returns the default value for the help command.
func (IDType) String() string {
	return field.TypeInt.String()
}

// InitCmd returns the init command for ent/c packages.
func InitCmd() *cobra.Command {
	c := NewCmd()
	c.Use = "init [flags] [schemas]"
	c.Short = "initialize an environment with zero or more schemas"
	c.Example = examples(
		"ent init Example",
		"ent init --target entv1/schema User Group",
		"ent init --template ./path/to/file.tmpl User",
	)
	c.Deprecated = "use `ent new` instead"
	return c
}

// NewCmd returns the new command for ent/c packages.
func NewCmd() *cobra.Command {
	var target, tmplPath string
	cmd := &cobra.Command{
		Use:   "new [flags] [schemas]",
		Short: "new an environment with zero or more schemas",
		Example: examples(
			"ent new Example",
			"ent new --target entv1/schema User Group",
			"ent new --template ./path/to/file.tmpl User",
		),
		Args: func(_ *cobra.Command, names []string) error {
			for _, name := range names {
				if !unicode.IsUpper(rune(name[0])) {
					return errors.New("schema names must begin with uppercase")
				}
			}
			return nil
		},
		Run: func(cmd *cobra.Command, names []string) {
			var (
				err  error
				tmpl *template.Template
			)
			if tmplPath != "" {
				tmpl, err = template.ParseFiles(tmplPath)
			} else {
				tmpl, err = template.New("schema").Parse(defaultTemplate)
			}
			if err != nil {
				log.Fatalln(fmt.Errorf("ent/new: could not parse template %w", err))
			}
			if err := newEnv(target, names, tmpl); err != nil {
				log.Fatalln(fmt.Errorf("ent/new: %w", err))
			}
		},
	}
	cmd.Flags().StringVar(&target, "target", defaultSchema, "target directory for schemas")
	cmd.Flags().StringVar(&tmplPath, "template", "", "template to use for new schemas")
	return cmd
}

// DescribeCmd returns the describe command for ent/c packages.
func DescribeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "describe [flags] path",
		Short: "printer a description of the graph schema",
		Example: examples(
			"ent describe ./ent/schema",
			"ent describe github.com/a8m/x",
		),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, path []string) {
			graph, err := entc.LoadGraph(path[0], &gen.Config{})
			if err != nil {
				log.Fatalln(err)
			}
			printer.Fprint(os.Stdout, graph)
		},
	}
}

// GenerateCmd returns the generate command for ent/c packages.
func GenerateCmd(postRun ...func(*gen.Config)) *cobra.Command {
	var (
		cfg       gen.Config
		storage   string
		features  []string
		templates []string
		idtype    = IDType(field.TypeInt)
		cmd       = &cobra.Command{
			Use:   "generate [flags] path",
			Short: "generate go code for the schema directory",
			Example: examples(
				"ent generate ./ent/schema",
				"ent generate github.com/a8m/x",
			),
			Args: cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, path []string) {
				opts := []entc.Option{
					entc.Storage(storage),
					entc.FeatureNames(features...),
				}
				for _, tmpl := range templates {
					typ := "dir"
					if parts := strings.SplitN(tmpl, "=", 2); len(parts) > 1 {
						typ, tmpl = parts[0], parts[1]
					}
					switch typ {
					case "dir":
						opts = append(opts, entc.TemplateDir(tmpl))
					case "file":
						opts = append(opts, entc.TemplateFiles(tmpl))
					case "glob":
						opts = append(opts, entc.TemplateGlob(tmpl))
					default:
						log.Fatalln("unsupported template type", typ)
					}
				}
				// If the target directory is not inferred from
				// the schema path, resolve its package path.
				if cfg.Target != "" {
					pkgPath, err := PkgPath(DefaultConfig, cfg.Target)
					if err != nil {
						log.Fatalln(err)
					}
					cfg.Package = pkgPath
				}
				cfg.IDType = &field.TypeInfo{Type: field.Type(idtype)}
				if err := entc.Generate(path[0], &cfg, opts...); err != nil {
					log.Fatalln(err)
				}
				for _, fn := range postRun {
					fn(&cfg)
				}
			},
		}
	)
	cmd.Flags().Var(&idtype, "idtype", "type of the id field")
	cmd.Flags().StringVar(&storage, "storage", "sql", "storage driver to support in codegen")
	cmd.Flags().StringVar(&cfg.Header, "header", "", "override codegen header")
	cmd.Flags().StringVar(&cfg.Target, "target", "", "target directory for codegen")
	cmd.Flags().StringSliceVarP(&features, "feature", "", nil, "extend codegen with additional features")
	cmd.Flags().StringSliceVarP(&templates, "template", "", nil, "external templates to execute")
	return cmd
}

// newEnv create an new environment for ent codegen.
func newEnv(target string, names []string, tmpl *template.Template) error {
	if err := createDir(target); err != nil {
		return fmt.Errorf("create dir %s: %w", target, err)
	}
	for _, name := range names {
		if err := gen.ValidSchemaName(name); err != nil {
			return fmt.Errorf("new schema %s: %w", name, err)
		}
		if fileExists(target, name) {
			return fmt.Errorf("new schema %s: already exists", name)
		}
		b := bytes.NewBuffer(nil)
		if err := tmpl.Execute(b, name); err != nil {
			return fmt.Errorf("executing template %s: %w", name, err)
		}
		newFileTarget := filepath.Join(target, strings.ToLower(name+".go"))
		if err := os.WriteFile(newFileTarget, b.Bytes(), 0644); err != nil {
			return fmt.Errorf("writing file %s: %w", newFileTarget, err)
		}
	}
	return nil
}

func createDir(target string) error {
	_, err := os.Stat(target)
	if err == nil || !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(target, os.ModePerm); err != nil {
		return fmt.Errorf("creating schema directory: %w", err)
	}
	if target != defaultSchema {
		return nil
	}
	if err := os.WriteFile("ent/generate.go", []byte(genFile), 0644); err != nil {
		return fmt.Errorf("creating generate.go file: %w", err)
	}
	return nil
}

func fileExists(target, name string) bool {
	var _, err = os.Stat(filepath.Join(target, strings.ToLower(name+".go")))

	return err == nil
}

const (
	// default schema package path.
	defaultSchema = "ent/schema"
	// ent/generate.go file used for "go generate" command.
	genFile = "package ent\n\n//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema\n"
	// schema template for the "init" command.
	defaultTemplate = `package schema

import "entgo.io/ent"

// {{ . }} holds the schema definition for the {{ . }} entity.
type {{ . }} struct {
	ent.Schema
}

// Fields of the {{ . }}.
func ({{ . }}) Fields() []ent.Field {
	return nil
}

// Edges of the {{ . }}.
func ({{ . }}) Edges() []ent.Edge {
	return nil
}
`
)

// examples formats the given examples to the cli.
func examples(ex ...string) string {
	for i := range ex {
		ex[i] = "  " + ex[i] // indent each row with 2 spaces.
	}
	return strings.Join(ex, "\n")
}
