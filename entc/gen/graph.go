// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package gen is the interface for generating loaded schemas into a Go package.
package gen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"text/template"
	"text/template/parse"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/entc/load"
	"github.com/facebookincubator/ent/schema/field"

	"golang.org/x/tools/imports"
)

type (
	// Config for global codegen to be shared between all nodes.
	Config struct {
		// Schema is the package path for the schema package.
		Schema string
		// Target is the filepath for the directory that holds the generated code.
		Target string
		// Package path for the targeted directory that holds the generated code.
		Package string
		// Header is an optional header signature for generated files.
		Header string
		// Storage to support in codegen.
		Storage *Storage
		// IDType specifies the type of the id field in the codegen.
		// The supported types are string and int, which also the default.
		IDType *field.TypeInfo
		// Template specifies an alternative template to execute or to override
		// the default. If nil, the default template is used.
		//
		// Note that, additional templates are executed on the Graph object and
		// the execution output is stored in a file derived by the template name.
		Template *template.Template
		// codegen state.
		state
	}
	// Graph holds the nodes/entities of the loaded graph schema. Note that, it doesn't
	// hold the edges of the graph. Instead, each Type holds the edges for other Types.
	Graph struct {
		*Config
		// Nodes are list of Go types that mapped to the types in the loaded schema.
		Nodes []*Type
		// Schemas holds the raw interfaces for the loaded schemas.
		Schemas []*load.Schema
	}
	// state holds the codegen state for the graph.
	state struct {
		// noSchemaImport indicates that the generated type package should
		// ignore importing the user's schema package. It's True if one of
		// the user's schemas uses hooks and circular imports is possible.
		noSchemaImport bool
	}
)

// NewGraph creates a new Graph for the code generation from the given schema definitions.
// It fails if one of the schemas is invalid.
func NewGraph(c *Config, schemas ...*load.Schema) (g *Graph, err error) {
	defer catch(&err)
	g = &Graph{c, make([]*Type, 0, len(schemas)), schemas}
	for _, schema := range schemas {
		g.addNode(schema)
	}
	for _, schema := range schemas {
		g.addEdges(schema)
	}
	for _, t := range g.Nodes {
		check(resolve(t), "resolve %q relations", t.Name)
	}
	for _, t := range g.Nodes {
		t.resolveFKs()
	}
	for _, schema := range schemas {
		g.addIndexes(schema)
	}
	return
}

// Gen generates the artifacts for the graph.
func (g *Graph) Gen() (err error) {
	defer catch(&err)
	var (
		written             []string
		templates, external = g.templates()
	)
	for _, n := range g.Nodes {
		path := filepath.Join(g.Config.Target, n.Package())
		check(os.MkdirAll(path, os.ModePerm), "create dir %q", path)
		for _, tmpl := range Templates {
			b := bytes.NewBuffer(nil)
			check(templates.ExecuteTemplate(b, tmpl.Name, n), "execute template %q", tmpl.Name)
			target := filepath.Join(g.Config.Target, tmpl.Format(n))
			check(ioutil.WriteFile(target, b.Bytes(), 0644), "write file %s", target)
			written = append(written, target)
		}
	}
	for _, tmpl := range append(GraphTemplates, external...) {
		if tmpl.Skip != nil && tmpl.Skip(g) {
			continue
		}
		if dir := filepath.Dir(tmpl.Format); dir != "." {
			path := filepath.Join(g.Config.Target, dir)
			check(os.MkdirAll(path, os.ModePerm), "create dir %q", path)
		}
		b := bytes.NewBuffer(nil)
		check(templates.ExecuteTemplate(b, tmpl.Name, g), "execute template %q", tmpl.Name)
		target := filepath.Join(g.Config.Target, tmpl.Format)
		check(ioutil.WriteFile(target, b.Bytes(), 0644), "write file %s", target)
		written = append(written, target)
	}
	// We can't run "imports" on files when the state is not completed.
	// Because, "goimports" will drop undefined package. Therefore, it's
	// suspended to the end of the writing.
	return formatFiles(written)
}

// addNode creates a new Type/Node/Ent to the graph.
func (g *Graph) addNode(schema *load.Schema) {
	t, err := NewType(g.Config, schema)
	check(err, "create type %s", schema.Name)
	g.Nodes = append(g.Nodes, t)
}

// addIndexes adds the indexes for the schema type.
func (g *Graph) addIndexes(schema *load.Schema) {
	typ, _ := g.typ(schema.Name)
	for _, idx := range schema.Indexes {
		check(typ.AddIndex(idx), "invalid index for schema %q", schema.Name)
	}
}

// addEdges adds the node edges to the graph.
func (g *Graph) addEdges(schema *load.Schema) {
	t, _ := g.typ(schema.Name)
	for _, e := range schema.Edges {
		typ, ok := g.typ(e.Type)
		expect(ok, "type %q does not exist for edge", e.Type)
		switch {
		// Assoc only.
		case !e.Inverse:
			t.Edges = append(t.Edges, &Edge{
				Type:      typ,
				Name:      e.Name,
				Owner:     t,
				Unique:    e.Unique,
				Optional:  !e.Required,
				StructTag: e.Tag,
			})
		// Inverse only.
		case e.Inverse && e.Ref == nil:
			expect(e.RefName != "", "missing reference name for inverse edge: %s.%s", t.Name, e.Name)
			t.Edges = append(t.Edges, &Edge{
				Type:      typ,
				Name:      e.Name,
				Owner:     typ,
				Inverse:   e.RefName,
				Unique:    e.Unique,
				Optional:  !e.Required,
				StructTag: e.Tag,
			})
		// Inverse and assoc.
		case e.Inverse:
			ref := e.Ref
			expect(e.RefName == "", "reference name is derived from the assoc name: %s.%s <-> %s.%s", t.Name, ref.Name, t.Name, e.Name)
			expect(ref.Type == t.Name, "assoc-inverse edge allowed only as o2o relation of the same type")
			t.Edges = append(t.Edges, &Edge{
				Type:      typ,
				Name:      e.Name,
				Owner:     t,
				Inverse:   ref.Name,
				Unique:    e.Unique,
				Optional:  !e.Required,
				StructTag: e.Tag,
			}, &Edge{
				Type:      typ,
				Owner:     t,
				Name:      ref.Name,
				Unique:    ref.Unique,
				Optional:  !ref.Required,
				StructTag: ref.Tag,
			})
		default:
			panic(graphError{"edge must be either an assoc or inverse edge"})
		}
	}
}

// resolve resolves the type reference and relation of edges.
// It fails if one of the references is missing or invalid.
//
// relation definitions between A and B, where A is the owner of
// the edge and B uses this edge as a back-reference:
//
// 	O2O
// 	 - A have a unique edge (E) to B, and B have a back-reference unique edge (E') for E.
// 	 - A have a unique edge (E) to A.
//
// 	O2M (The "Many" side, keeps a reference to the "One" side).
// 	 - A have an edge (E) to B (not unique), and B doesn't have a back-reference edge for E.
// 	 - A have an edge (E) to B (not unique), and B have a back-reference unique edge (E') for E.
//
// 	M2O (The "Many" side, holds the reference to the "One" side).
// 	 - A have a unique edge (E) to B, and B doesn't have a back-reference edge for E.
// 	 - A have a unique edge (E) to B, and B have a back-reference non-unique edge (E') for E.
//
// 	M2M
// 	 - A have an edge (E) to B (not unique), and B have a back-reference non-unique edge (E') for E.
// 	 - A have an edge (E) to A (not unique).
//
func resolve(t *Type) error {
	for _, e := range t.Edges {
		switch {
		case e.IsInverse():
			ref, ok := e.Type.HasAssoc(e.Inverse)
			if !ok {
				return fmt.Errorf("edge is missing for inverse edge: %s.%s", e.Type.Name, e.Name)
			}
			if !e.Optional && !ref.Optional {
				return fmt.Errorf("edges cannot be required in both directions: %s.%s <-> %s.%s", t.Name, e.Name, e.Type.Name, ref.Name)
			}
			if ref.Type != t {
				return fmt.Errorf("mismatch type for back-ref %q of %s.%s <-> %s.%s", e.Inverse, t.Name, e.Name, e.Type.Name, ref.Name)
			}
			table := t.Table()
			// Name the foreign-key column in a format that wouldn't change even if an inverse
			// edge is dropped (or added). The format is: "<Edge-Owner>_<Edge-Name>".
			column := fmt.Sprintf("%s_%s", e.Type.Label(), snake(ref.Name))
			switch a, b := ref.Unique, e.Unique; {
			// If the relation column is in the inverse side/table. The rule is simple, if assoc is O2M,
			// then inverse is M2O and the relation is in its table.
			case a && b:
				e.Rel.Type, ref.Rel.Type = O2O, O2O
			case !a && b:
				e.Rel.Type, ref.Rel.Type = M2O, O2M

			// If the relation column is in the assoc side.
			case a && !b:
				e.Rel.Type, ref.Rel.Type = O2M, M2O
				table = e.Type.Table()

			case !a && !b:
				e.Rel.Type, ref.Rel.Type = M2M, M2M
				table = e.Type.Label() + "_" + ref.Name
				c1, c2 := ref.Owner.Label()+"_id", ref.Type.Label()+"_id"
				// If the relation is from the same type: User has Friends ([]User).
				// give the second column a different name (the relation name).
				if c1 == c2 {
					c2 = rules.Singularize(e.Name) + "_id"
				}
				e.Rel.Columns = []string{c1, c2}
				ref.Rel.Columns = []string{c1, c2}
			}
			e.Rel.Table, ref.Rel.Table = table, table
			if !e.M2M() {
				e.Rel.Columns = []string{column}
				ref.Rel.Columns = []string{column}
			}
		// Assoc with uninitialized relation.
		case !e.IsInverse() && e.Rel.Type == Unk:
			switch {
			case !e.Unique && e.Type == t:
				e.Rel.Type = M2M
				e.Bidi = true
				e.Rel.Table = t.Label() + "_" + e.Name
				e.Rel.Columns = []string{e.Owner.Label() + "_id", rules.Singularize(e.Name) + "_id"}
			case e.Unique && e.Type == t:
				e.Rel.Type = O2O
				e.Bidi = true
				e.Rel.Table = t.Table()
			case e.Unique:
				e.Rel.Type = M2O
				e.Rel.Table = t.Table()
			default:
				e.Rel.Type = O2M
				e.Rel.Table = e.Type.Table()
			}
			if !e.M2M() {
				e.Rel.Columns = []string{fmt.Sprintf("%s_%s", t.Label(), snake(e.Name))}
			}
		}
	}
	return nil
}

// Tables returns the schema definitions of SQL tables for the graph.
func (g *Graph) Tables() (all []*schema.Table) {
	tables := make(map[string]*schema.Table)
	for _, n := range g.Nodes {
		table := schema.NewTable(n.Table()).AddPrimary(n.ID.PK())
		for _, f := range n.Fields {
			table.AddColumn(f.Column())
		}
		tables[table.Name] = table
		all = append(all, table)
	}
	for _, n := range g.Nodes {
		// Foreign key + reference OR join table.
		for _, e := range n.Edges {
			if e.IsInverse() {
				continue
			}
			switch e.Rel.Type {
			case O2O, O2M:
				// "owner" is the table that owns the relations (we set the foreign-key on)
				// and "ref" is the referenced table.
				owner, ref := tables[e.Rel.Table], tables[n.Table()]
				pk := ref.PrimaryKey[0]
				column := &schema.Column{Name: e.Rel.Column(), Size: pk.Size, Type: pk.Type, Unique: e.Rel.Type == O2O, Nullable: true}
				owner.AddColumn(column)
				owner.AddForeignKey(&schema.ForeignKey{
					RefTable:   ref,
					OnDelete:   schema.SetNull,
					Columns:    []*schema.Column{column},
					RefColumns: []*schema.Column{ref.PrimaryKey[0]},
					Symbol:     fmt.Sprintf("%s_%s_%s", owner.Name, ref.Name, e.Name),
				})
			case M2O:
				ref, owner := tables[e.Type.Table()], tables[e.Rel.Table]
				pk := ref.PrimaryKey[0]
				column := &schema.Column{Name: e.Rel.Column(), Size: pk.Size, Type: pk.Type, Nullable: true}
				owner.AddColumn(column)
				owner.AddForeignKey(&schema.ForeignKey{
					RefTable:   ref,
					OnDelete:   schema.SetNull,
					Columns:    []*schema.Column{column},
					RefColumns: []*schema.Column{ref.PrimaryKey[0]},
					Symbol:     fmt.Sprintf("%s_%s_%s", owner.Name, ref.Name, e.Name),
				})
			case M2M:
				t1, t2 := tables[n.Table()], tables[e.Type.Table()]
				c1 := &schema.Column{Name: e.Rel.Columns[0], Type: field.TypeInt}
				if ref := n.ID; ref.UserDefined {
					c1.Type = ref.Type.Type
				}
				c2 := &schema.Column{Name: e.Rel.Columns[1], Type: field.TypeInt}
				if ref := e.Type.ID; ref.UserDefined {
					c2.Type = ref.Type.Type
				}
				all = append(all, &schema.Table{
					Name:       e.Rel.Table,
					Columns:    []*schema.Column{c1, c2},
					PrimaryKey: []*schema.Column{c1, c2},
					ForeignKeys: []*schema.ForeignKey{
						{
							RefTable:   t1,
							OnDelete:   schema.Cascade,
							Columns:    []*schema.Column{c1},
							RefColumns: []*schema.Column{t1.PrimaryKey[0]},
							Symbol:     fmt.Sprintf("%s_%s", e.Rel.Table, c1.Name),
						},
						{
							RefTable:   t2,
							OnDelete:   schema.Cascade,
							Columns:    []*schema.Column{c2},
							RefColumns: []*schema.Column{t2.PrimaryKey[0]},
							Symbol:     fmt.Sprintf("%s_%s", e.Rel.Table, c2.Name),
						},
					},
				})
			}
		}
	}
	// Append indexes to tables after all columns were added (including relation columns).
	for _, n := range g.Nodes {
		table := tables[n.Table()]
		for _, idx := range n.Indexes {
			table.AddIndex(idx.Name, idx.Unique, idx.Columns)
		}
	}
	return
}

// SupportMigrate reports if the codegen supports schema migration.
func (g *Graph) SupportMigrate() bool {
	return g.Storage.SchemaMode.Support(Migrate)
}

func (g *Graph) typ(name string) (*Type, bool) {
	for _, n := range g.Nodes {
		if name == n.Name {
			return n, true
		}
	}
	return nil, false
}

// templates returns the template.Template for the code and external templates
// to execute on the Graph object if provided.
func (g *Graph) templates() (*template.Template, []GraphTemplate) {
	templates = template.Must(templates.Clone())
	if g.Template == nil {
		return templates, nil
	}
	external := make([]GraphTemplate, 0)
	for _, tmpl := range g.Template.Templates() {
		name := tmpl.Name()
		// Check that is not defined in the default templates
		// if it's not the root.
		if templates.Lookup(name) == nil && !parse.IsEmptyTree(tmpl.Root) {
			external = append(external, GraphTemplate{
				Name:   name,
				Format: snake(name) + ".go",
			})
		}
		templates = template.Must(templates.AddParseTree(name, tmpl.Tree))
	}
	return templates, external
}

// ModuleInfo returns the entc binary module version.
func (Config) ModuleInfo() (m debug.Module) {
	info, ok := debug.ReadBuildInfo()
	if ok {
		m = info.Main
	}
	return
}

// formatFiles runs "goimports" on given paths.
func formatFiles(paths []string) error {
	for _, path := range paths {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file %s: %v", path, err)
		}
		src, err := imports.Process(path, buf, nil)
		if err != nil {
			return fmt.Errorf("format file %s: %v", path, err)
		}
		if err := ioutil.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %v", path, err)
		}
	}
	return nil
}

// expect panic if the condition is false.
func expect(cond bool, msg string, args ...interface{}) {
	if !cond {
		panic(graphError{fmt.Sprintf(msg, args...)})
	}
}

// check panics if the error is not nil.
func check(err error, msg string, args ...interface{}) {
	if err != nil {
		args = append(args, err)
		panic(graphError{fmt.Sprintf(msg+": %s", args...)})
	}
}

type graphError struct {
	msg string
}

func (p graphError) Error() string { return fmt.Sprintf("entc/gen: %s", p.msg) }

func catch(err *error) {
	if e := recover(); e != nil {
		gerr, ok := e.(graphError)
		if !ok {
			panic(e)
		}
		*err = gerr
	}
}
