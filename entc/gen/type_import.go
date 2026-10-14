// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"maps"
	"path"
	"slices"
	"strconv"

	"entgo.io/ent/schema/field"
)

// resolveTypeImports assigns consistent aliases to generic type imports across
// the graph. Shared files may refer to types from several different schemas.
func resolveTypeImports(g *Graph) error {
	names := maps.Clone(importPkg)
	// Storage imports are also reserved before templates are initialized.
	for _, driver := range drivers {
		for _, pkg := range driver.Imports {
			names[path.Base(pkg)] = pkg
		}
	}
	names["predicate"] = g.Package + "/predicate"
	if g.Schema != "" {
		names[path.Base(g.Schema)] = g.Schema
	}
	var types []*field.TypeInfo
	seen := make(map[*field.TypeInfo]bool)
	for _, n := range g.Nodes {
		for _, f := range append([]*Field{n.ID}, n.Fields...) {
			if f == nil || f.Type == nil || seen[f.Type] {
				continue
			}
			t := f.Type
			seen[t] = true
			if len(t.PkgImports) > 0 {
				types = append(types, t)
				continue
			}
			// Preserve the aliases already used by non-generic field types.
			if t.PkgPath != "" {
				name := t.PkgName
				if name == "" {
					name = path.Base(t.PkgPath)
				}
				if names[name] == "" {
					names[name] = t.PkgPath
				}
			}
		}
	}
	paths := make(map[string]string)
	for _, name := range slices.Sorted(maps.Keys(names)) {
		paths[names[name]] = name
	}
	for _, t := range types {
		replacements := make(map[string]string)
		imports := make(map[string]string)
		for _, pkg := range slices.Sorted(maps.Keys(t.PkgImports)) {
			name := t.PkgImports[pkg]
			alias := paths[pkg]
			if alias == "" {
				alias = name
				for i := 2; names[alias] != ""; i++ {
					alias = name + strconv.Itoa(i)
				}
				names[alias], paths[pkg] = pkg, alias
			}
			replacements[name] = alias
			imports[pkg] = alias
		}
		expr, err := parser.ParseExpr(t.Ident)
		if err != nil {
			return fmt.Errorf("parse type %q: %w", t.Ident, err)
		}
		ast.Inspect(expr, func(n ast.Node) bool {
			if s, ok := n.(*ast.SelectorExpr); ok {
				if id, ok := s.X.(*ast.Ident); ok && replacements[id.Name] != "" {
					id.Name = replacements[id.Name]
				}
			}
			return true
		})
		var b bytes.Buffer
		if err := format.Node(&b, token.NewFileSet(), expr); err != nil {
			return err
		}
		t.Ident = b.String()
		t.PkgImports = imports
		if name := imports[t.PkgPath]; name != "" {
			t.PkgName = name
		}
	}
	return nil
}
