// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "entc"}
	cmd.AddCommand(
		func() *cobra.Command {
			var (
				target string
				cmd    = &cobra.Command{
					Use:   "init [flags] [schemas]",
					Short: "initialize an environment with zero or more schemas",
					Example: examples(
						"entc init Example",
						"entc init --target entv1/schema User Group",
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
						if err := initEnv(target, names); err != nil {
							log.Fatalln(err)
						}
					},
				}
			)
			cmd.Flags().StringVar(&target, "target", defaultSchema, "target directory for schemas")
			return cmd
		}(),
		&cobra.Command{
			Use:   "describe [flags] path",
			Short: "print a description of the graph schema",
			Example: examples(
				"entc describe ./ent/schema",
				"entc describe github.com/a8m/x",
			),
			Args: cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, path []string) {
				graph, err := entc.LoadGraph(path[0], &gen.Config{})
				if err != nil {
					log.Fatalln(err)
				}
				printer{os.Stdout}.Print(graph)
			},
		},
		func() *cobra.Command {
			var (
				cfg       gen.Config
				storage   string
				features  []string
				templates []string
				idtype    = idType(field.TypeInt)
				cmd       = &cobra.Command{
					Use:   "generate [flags] path",
					Short: "generate go code for the schema directory",
					Example: examples(
						"entc generate ./ent/schema",
						"entc generate github.com/a8m/x",
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
		}(),
	)
	_ = cmd.Execute()
}

// custom implementation for pflag.
type idType field.Type

// Set implements the Set method of the flag.Value interface.
func (t *idType) Set(s string) error {
	switch s {
	case field.TypeInt.String():
		*t = idType(field.TypeInt)
	case field.TypeInt64.String():
		*t = idType(field.TypeInt64)
	case field.TypeUint.String():
		*t = idType(field.TypeUint)
	case field.TypeUint64.String():
		*t = idType(field.TypeUint64)
	case field.TypeString.String():
		*t = idType(field.TypeString)
	default:
		return errors.New("invalid type")
	}
	return nil
}

// Type returns the type representation of the id option for help command.
func (idType) Type() string {
	return fmt.Sprintf("%v", []field.Type{
		field.TypeInt,
		field.TypeInt64,
		field.TypeUint,
		field.TypeUint64,
		field.TypeString,
	})
}

// String returns the default value for the help command.
func (idType) String() string {
	return field.TypeInt.String()
}

// initEnv initialize an environment for ent codegen.
func initEnv(target string, names []string) error {
	if err := createDir(target); err != nil {
		return err
	}
	for _, name := range names {
		b := bytes.NewBuffer(nil)
		if err := tmpl.Execute(b, name); err != nil {
			log.Fatalln(err)
		}
		target := filepath.Join(target, strings.ToLower(name+".go"))
		if err := ioutil.WriteFile(target, b.Bytes(), 0644); err != nil {
			log.Fatalln(err)
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
	if err := ioutil.WriteFile("ent/generate.go", []byte(genFile), 0644); err != nil {
		return fmt.Errorf("creating generate.go file: %w", err)
	}
	return nil
}

// schema template for the "init" command.
var tmpl = template.Must(template.New("schema").
	Parse(`package schema

import "github.com/facebook/ent"

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
`))

const (
	// default schema package path.
	defaultSchema = "ent/schema"
	// ent/generate.go file used for "go generate" command.
	genFile = "package ent\n\n//go:generate go run github.com/facebook/ent/cmd/entc generate ./schema\n"
)

// examples formats the given examples to the cli.
func examples(ex ...string) string {
	for i := range ex {
		ex[i] = "  " + ex[i] // indent each row with 2 spaces.
	}
	return strings.Join(ex, "\n")
}
