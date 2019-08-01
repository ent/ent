package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"fbc/ent/entc/gen"
	"fbc/ent/entc/load"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "entc"}
	cmd.AddCommand(
		&cobra.Command{
			Use:   "init [schemas]",
			Short: "initialize an environment with zero or more schemas",
			Example: examples(
				"entc init Example",
				"entc init User Group",
			),
			DisableFlagsInUseLine: true,
			Args: func(_ *cobra.Command, names []string) error {
				for _, name := range names {
					if !unicode.IsUpper(rune(name[0])) {
						return fmt.Errorf("schema names must begin with uppercase")
					}
				}
				return nil
			},
			Run: func(cmd *cobra.Command, names []string) {
				path := "ent/schema"
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
		},
		&cobra.Command{
			Use:   "describe [flags] path",
			Short: "print a description of the graph schema",
			Example: examples(
				"entc describe ./ent/schema",
				"entc describe github.com/a8m/x",
			),
			Args: cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, path []string) {
				graph, err := loadGraph(path[0], gen.Config{})
				failOnErr(err)
				graph.Describe(os.Stdout)
			},
		},
		func() *cobra.Command {
			var (
				cfg     gen.Config
				storage []string
				cmd     = &cobra.Command{
					Use:   "generate [flags] path",
					Short: "generate go code for the schema directory",
					Example: examples(
						"entc generate ./ent/schema",
						"entc generate github.com/a8m/x",
					),
					Args: cobra.ExactArgs(1),
					Run: func(cmd *cobra.Command, path []string) {
						if cfg.Target == "" {
							abs, err := filepath.Abs(path[0])
							failOnErr(err)
							cfg.Target = filepath.Dir(abs)
						}
						for _, s := range storage {
							sr, err := gen.NewStorage(s)
							failOnErr(err)
							cfg.Storage = append(cfg.Storage, sr)
						}
						graph, err := loadGraph(path[0], cfg)
						failOnErr(err)
						failOnErr(graph.Gen())
					},
				}
			)
			cmd.Flags().StringVar(&cfg.Header, "header", "", "override codegen header")
			cmd.Flags().StringVar(&cfg.Target, "target", "", "target directory for codegen")
			cmd.Flags().StringSliceVarP(&storage, "storage", "", []string{"sql"}, "list of storage drivers to support")
			return cmd
		}(),
	)
	cmd.Execute()
}

// loadGraph loads the given schema package from the given path
// and construct a *gen.Graph. The path can be either a package
// path (e.g github.com/a8m/x) or a filepath.
//
// The second argument is an optional config for the graph creation.
func loadGraph(path string, cfg gen.Config) (*gen.Graph, error) {
	spec, err := (&load.Config{Path: path}).Load()
	if err != nil {
		return nil, err
	}
	cfg.Schema = spec.PkgPath
	cfg.Package = filepath.Dir(spec.PkgPath)
	return gen.NewGraph(cfg, spec.Schemas...)
}

// schema template for the "init" command.
var tmpl = template.Must(template.New("schema").
	Parse(`package schema

import "fbc/ent"

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
