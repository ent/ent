package main

import (
	"fmt"

	"fbc/ent/entc/gen"
	"fbc/ent/entc/plugin"
)

// Gen is the required plugin symbol.
var Gen = plugin.GeneratorFunc(func(graph *gen.Graph) error {
	for _, n := range graph.Nodes {
		fmt.Println(n.Name)
	}
	return nil
})
