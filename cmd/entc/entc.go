// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/facebookincubator/ent/entc"
	"github.com/facebookincubator/ent/entc/gen"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "entc"}
	cmd.AddCommand(
		func() *cobra.Command {
			var (
				path string
				cmd  = &cobra.Command{
					Use:   "init [flags] [schemas]",
					Short: "initialize an environment with zero or more schemas",
					Example: examples(
						"entc init Example",
						"entc init --target entv1/schema User Group",
					),
					Args: func(_ *cobra.Command, names []string) error {
						for _, name := range names {
							if !unicode.IsUpper(rune(name[0])) {
								return fmt.Errorf("schema names must begin with uppercase")
							}
						}
						return nil
					},
					Run: func(cmd *cobra.Command, names []string) {
						_, err := os.Stat(path)
						if os.IsNotExist(err) {
							err = os.MkdirAll(path, os.ModePerm)
						}
						failOnErr(err)
						for _, name := range names {
							b := bytes.NewBuffer(nil)
							failOnErr(tmpl.Execute(b, name))
							target := filepath.Join(path, strings.ToLower(name+".go"))
							failOnErr(ioutil.WriteFile(target, b.Bytes(), 0644))
						}
					},
				}
			)
			cmd.Flags().StringVar(&path, "target", "ent/schema", "target directory for schemas")
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
				failOnErr(err)
				p := printer{os.Stdout}
				p.Print(graph)
			},
		},
		func() *cobra.Command {
			var (
				cfg      gen.Config
				storage  string
				template []string
				idtype   = idType(field.TypeInt)
				cmd      = &cobra.Command{
					Use:   "generate [flags] path",
					Short: "generate go code for the schema directory",
					Example: examples(
						"entc generate ./ent/schema",
						"entc generate github.com/a8m/x",
					),
					Args: cobra.ExactArgs(1),
					Run: func(cmd *cobra.Command, path []string) {
						opts := []entc.Option{entc.Storage(storage)}
						for _, tmpl := range template {
							opts = append(opts, entc.TemplateDir(tmpl))
						}
						// If the target directory is not inferred from
						// the schema path, resolve its package path.
						if cfg.Target != "" {
							pkgPath, err := PkgPath(DefaultConfig, cfg.Target)
							failOnErr(err)
							cfg.Package = pkgPath
						}
						cfg.IDType = &field.TypeInfo{Type: field.Type(idtype)}
						err := entc.Generate(path[0], &cfg, opts...)
						failOnErr(err)
					},
				}
			)
			cmd.Flags().Var(&idtype, "idtype", "type of the id field")
			cmd.Flags().StringVar(&storage, "storage", "sql", "storage driver to support in codegen")
			cmd.Flags().StringVar(&cfg.Header, "header", "", "override codegen header")
			cmd.Flags().StringVar(&cfg.Target, "target", "", "target directory for codegen")
			cmd.Flags().StringSliceVarP(&template, "template", "", nil, "external templates to execute")
			return cmd
		}(),
		&cobra.Command{
			Use:   "clean [flags] path",
			Short: "delete generated go code",
			Example: examples(
				"entc clean ./ent",
			),
			Args: cobra.ExactArgs(1),
			Run:  clean,
		},
	)
	cmd.Execute()
}

// schema template for the "init" command.
var tmpl = template.Must(template.New("schema").
	Parse(`package schema

import "github.com/facebookincubator/ent"

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

func failOnErr(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		fmt.Fprint(os.Stderr, "\n")
		os.Exit(1)
	}
}

// examples formats the given examples to the cli.
func examples(ex ...string) string {
	for i := range ex {
		ex[i] = "  " + ex[i] // indent each row with 2 spaces.
	}
	return strings.Join(ex, "\n")
}

func clean(cmd *cobra.Command, path []string) {
	cleanup := make([]string, 0)
	cleanup = append(cleanup, filepath.Join(path[0], "migrate"))
	cleanup = append(cleanup, filepath.Join(path[0], "predicate"))
	cleanup = append(cleanup, filepath.Join(path[0], "client.go"))
	cleanup = append(cleanup, filepath.Join(path[0], "config.go"))
	cleanup = append(cleanup, filepath.Join(path[0], "context.go"))
	cleanup = append(cleanup, filepath.Join(path[0], "ent.go"))
	cleanup = append(cleanup, filepath.Join(path[0], "tx.go"))
	schemaDir, err := ioutil.ReadDir(filepath.Join(path[0], "schema"))
	failOnErr(err)
	for i := 0; i < len(schemaDir); i++ {
		fname := schemaDir[i].Name()
		ext := filepath.Ext(fname)
		if ext == ".go" {
			fnamePathWithoutExt := filepath.Join(path[0], strings.TrimRight(fname, ext))
			cleanup = append(cleanup, fnamePathWithoutExt)
			cleanup = append(cleanup, fnamePathWithoutExt+".go")
			filepath.Walk(path[0], func(walkpath string, info os.FileInfo, err error) error {
				match, _ := filepath.Match(fnamePathWithoutExt+"*", walkpath)
				if match {
					cleanup = append(cleanup, walkpath)
				}
				return nil
			})
		}
	}
	// remove files and directories
	for i := 0; i < len(cleanup); i++ {
		os.RemoveAll(cleanup[i])
	}
}
