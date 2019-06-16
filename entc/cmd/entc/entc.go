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

	"fbc/ent/entc/plugin"

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
				graph, err := plugin.LoadGraph(path[0])
				failOnErr(err)
				graph.Describe(os.Stdout)
			},
		},
		func() *cobra.Command {
			var (
				plugins        []string
				header, target string
				cmd            = &cobra.Command{
					Use:   "generate [flags] path",
					Short: "generate go code for the schema directory",
					Example: examples(
						"entc generate ./ent/schema",
						"entc generate github.com/a8m/x",
					),
					Args: cobra.ExactArgs(1),
					Run: func(cmd *cobra.Command, path []string) {
						graph, err := plugin.LoadGraph(path[0])
						failOnErr(err)

						if target == "" {
							abs, err := filepath.Abs(path[0])
							failOnErr(err)
							target = filepath.Dir(abs)
						}
						graph.Target = target
						graph.Header = header
						failOnErr(graph.Gen())

						// execute additional plugins.
						for _, plg := range plugins {
							failOnErr(plugin.Exec(plg, graph))
						}
					},
				}
			)
			cmd.Flags().StringVar(&header, "header", "", "override codegen header")
			cmd.Flags().StringVar(&target, "target", "", "target directory for codegen")
			cmd.Flags().StringSliceVarP(&plugins, "plugin", "", nil, "specifies additional plugin to execute")
			return cmd
		}(),
	)
	cmd.Execute()
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
