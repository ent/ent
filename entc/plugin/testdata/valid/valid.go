package main

import (
	"fbc/ent/entc/gen"
)

// Generator implements the plugin.Generator interface.
type Generator struct{}

// Gen implementation.
func (Generator) Gen(*gen.Graph) error {
	// logic goes here.
	return nil
}

// Gen is the required plugin symbol.
var Gen = Generator{}
