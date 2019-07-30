// Package plugin provides a way to extend entc via plugins.
// Plugin can be either a Go plugin that is loaded and executed
// by entc runtime or a standalone program.
package plugin

import (
	"os"
	"path/filepath"
	"plugin"

	"fbc/ent/entc/gen"
	"fbc/ent/entc/internal/build"

	"github.com/pkg/errors"
)

// Symbol is the expected symbol name in the provided plugin.
// The plugin.Symbol need to be a type Generator. For example:
//
//	package main
//
//	var Gen = GeneratorFunc(func(graph *gen.Graph) error {
//		return nil
//	})
//
const Symbol = "Gen"

// Generator is the interface that wrap the Gen method executed by entc.
type Generator interface {
	Gen(*gen.Graph) error
}

// The GeneratorFunc type is an adapter to allow the use of ordinary functions as Generator.
// If f is a function with the appropriate signature, GeneratorFunc(f) is a Generator that calls f.
type GeneratorFunc func(*gen.Graph) error

// Gen calls f(g).
func (f GeneratorFunc) Gen(g *gen.Graph) error { return f(g) }

// LoadGraph loads the given schema package from the given path
// and construct a *gen.Graph. The path can be either a package
// path (e.g github.com/a8m/x) or a filepath.
//
// The second argument is an optional config for the graph creation.
//
// This function used to create a standalone plugin programs that
// want to interact with the ent schemas. An example for usage:
//
//	package main
//
//	import (
//		"log"
//
//		"fbc/ent/entc/plugin"
//	)
//
//	func main() {
//		graph, err := plugin.LoadGraph("./ent/schema", gen.Config{})
//		if err != nil {
//			log.Fatal(err)
//		}
//		for _, node := range graph.Nodes {
//			log.Println(node.Name)
//		}
//	}
//
func LoadGraph(path string, cfg gen.Config) (*gen.Graph, error) {
	plg, err := (&build.Config{Path: path}).Build()
	if err != nil {
		return nil, err
	}
	defer os.Remove(plg.Path)
	schemas, err := plg.Load()
	if err != nil {
		return nil, err
	}
	cfg.Schema = plg.PkgPath
	cfg.Package = filepath.Dir(plg.PkgPath)
	return gen.NewGraph(cfg, schemas...)
}

// MustLoadGraph is like LoadGraph but panics if LoadGraph returns an error.
// It simplifies safe initialization of global variables holding a *gen.Graph.
func MustLoadGraph(path string, cfg gen.Config) *gen.Graph {
	graph, err := LoadGraph(path, cfg)
	if err != nil {
		panic(err)
	}
	return graph
}

// Exec loads and executes the provided plugin with
// the provided *gen/Graph.
//
// It returns an error if the plugin is invalid or
// it's not fulfilling the entc/plugin interface.
func Exec(path string, graph *gen.Graph) error {
	plg, err := plugin.Open(path)
	if err != nil {
		return errors.WithMessagef(err, "open plugin %s", path)
	}
	sym, err := plg.Lookup(Symbol)
	if err != nil {
		return errors.WithMessagef(err, "find symbol (%q) in plugin", Symbol)
	}
	g, ok := sym.(Generator)
	if !ok {
		return errors.Errorf("exported symbol %q does not implement the entc/plugin.Generator", Symbol)
	}
	return g.Gen(graph)
}
