package gen

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"fbc/ent/dialect/sql/schema"
	"fbc/ent/entc/load"
	"fbc/ent/field"
)

type (
	// Config for global generator configuration that similar for all nodes.
	Config struct {
		// Schema is the package path for the schema directory.
		Schema string
		// Target is the path for the directory that holding the generated code.
		Target string
		// Package name for the targeted directory that holds the generated code.
		Package string
		// Header is an optional header signature for generated files.
		Header string
		// Storage to support in codegen.
		Storage []*Storage
		// imports are the import packages used for code generation.
		imports map[string]string
	}
	// Graph holds the nodes/entities of the loaded graph schema. Note that, it doesn't
	// hold the edges of the graph. Instead, each Type holds the edges for other Types.
	Graph struct {
		Config
		// Nodes are list of Go types that mapped to the types in the loaded schema.
		Nodes []*Type
		// Schemas holds the raw interfaces for the loaded schemas.
		Schemas []*load.Schema
	}
)

// NewGraph creates a new Graph for the code generation from the given schema definitions.
// It fails if one of the schemas is invalid.
func NewGraph(c Config, schemas ...*load.Schema) (g *Graph, err error) {
	defer catch(&err)
	c.imports = imports()
	g = &Graph{c, make([]*Type, 0, len(schemas)), schemas}
	for _, schema := range schemas {
		g.addNode(schema)
	}
	for _, schema := range schemas {
		g.addEdges(schema)
	}
	for _, t := range g.Nodes {
		check(g.resolve(t), "resolve %q relations/references", t.Name)
	}
	return
}

// Gen generates the artifacts for the graph.
func (g *Graph) Gen() (err error) {
	defer catch(&err)
	for _, n := range g.Nodes {
		path := filepath.Join(g.Config.Target, n.Package())
		check(os.MkdirAll(path, os.ModePerm), "create dir %q", path)
		for _, tmpl := range Templates {
			b := bytes.NewBuffer(nil)
			check(templates.ExecuteTemplate(b, tmpl.Name, n), "execute template %q", tmpl.Name)
			target := filepath.Join(g.Config.Target, tmpl.Format(n))
			check(ioutil.WriteFile(target, b.Bytes(), 0644), "create file %q", target)
		}
	}
	for _, tmpl := range GraphTemplates {
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
		check(ioutil.WriteFile(target, b.Bytes(), 0644), "create file %q", target)
	}
	return run(exec.Command("goimports", "-w", g.Config.Target))
}

// Describe writes a description of the graph to the given writer.
func (g *Graph) Describe(w io.Writer) {
	for _, n := range g.Nodes {
		n.Describe(w)
	}
}

// addNode creates a new Type/Node/Ent to the graph.
func (g *Graph) addNode(schema *load.Schema) {
	t, err := NewType(g.Config, schema)
	check(err, "create type")
	g.Nodes = append(g.Nodes, t)
}

// addEdges adds the node edges to the graph.
func (g *Graph) addEdges(schema *load.Schema) {
	t, _ := g.typ(schema.Name)
	for _, e := range schema.Edges {
		typ, ok := g.typ(e.Type)
		expect(ok, "type %q does not exist for edge", e.Type)
		switch {
		// assoc only.
		case !e.Inverse:
			t.Edges = append(t.Edges, &Edge{
				Type:     typ,
				Name:     e.Name,
				Owner:    t,
				Unique:   e.Unique,
				Optional: !e.Required,
			})
		// inverse only.
		case e.Inverse && e.Ref == nil:
			expect(!e.Required, `inverse edge can not be required: %s.%s`, t.Name, e.Name)
			expect(e.RefName != "", `missing reference name for inverse edge: %s.%s`, t.Name, e.Name)
			t.Edges = append(t.Edges, &Edge{
				Type:     typ,
				Name:     e.Name,
				Owner:    typ,
				Inverse:  e.RefName,
				Unique:   e.Unique,
				Optional: !e.Required,
			})
		// inverse and assoc.
		case e.Inverse:
			ref := e.Ref
			expect(e.RefName == "", `reference name is derived from the assoc name: %s.%s <-> %s.%s`, t.Name, ref.Name, t.Name, e.Name)
			expect(ref.Type == t.Name, "assoc-inverse edge allowed only as o2o relation of the same type")
			t.Edges = append(t.Edges, &Edge{
				Type:     typ,
				Name:     e.Name,
				Owner:    t,
				Inverse:  ref.Name,
				Unique:   e.Unique,
				Optional: !e.Required,
			}, &Edge{
				Type:     typ,
				Owner:    t,
				Name:     ref.Name,
				Unique:   ref.Unique,
				Optional: !ref.Required,
			})
		default:
			panic(graphError{"edge must be either an assoc or inverse edge"})
		}
	}
}

// resolve resolves the type reference and relation of edges.
// It fails if one of the references is missing.
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
// 	M2O (The "One" side, keeps a reference to the "Many" side).
// 	 - A have a unique edge (E) to B, and B doesn't have a back-reference edge for E.
// 	 - A have a unique edge (E) to B, and B have a back-reference non-unique edge (E') for E.
//
// 	M2M
// 	 - A have an edge (E) to B (not unique), and B have a back-reference non-unique edge (E') for E.
// 	 - A have an edge (E) to A (not unique).
//
func (g *Graph) resolve(t *Type) error {
	for _, e := range t.Edges {
		switch {
		case e.IsInverse():
			ref, ok := e.Type.HasAssoc(e.Inverse)
			if !ok {
				return fmt.Errorf(`assoc is missing for inverse edge: %s.%s`, e.Type.Name, e.Name)
			}
			table := t.Table()
			// The name of the column is how we identify the other side. For example "A Parent has Children"
			// (Parent <-O2M-> Children), or "A User has Pets" (User <-O2M-> Pet). The Children/Pet hold the
			// relation, and they are identified the edge using how they call it in the inverse ("our parent")
			// even though that struct is called "User".
			column := snake(e.Name) + "_id"
			switch a, b := ref.Unique, e.Unique; {
			// If the relation column is in the inverse side/table. The rule is simple, if assoc is O2M,
			// then inverse is M2O and the relation is in its table.
			case a && b:
				e.Rel.Type, ref.Rel.Type = O2O, O2O
			case !a && b:
				e.Rel.Type, ref.Rel.Type = M2O, O2M

			// if the relation column is in the assoc side.
			case a && !b:
				e.Rel.Type, ref.Rel.Type = O2M, M2O
				table = e.Type.Table()
				column = snake(ref.Name) + "_id"

			case !a && !b:
				e.Rel.Type, ref.Rel.Type = M2M, M2M
				table = e.Type.Label() + "_" + ref.Name
				c1, c2 := ref.Owner.Label()+"_id", ref.Type.Label()+"_id"
				// if the relation is from the same type: User has Friends ([]User).
				// give the second column a different name (the relation name).
				if c1 == c2 {
					c2 = rules.Singularize(e.Name) + "_id"
				}
				e.Rel.Columns = append(e.Rel.Columns, c1, c2)
				ref.Rel.Columns = append(ref.Rel.Columns, c1, c2)
			}
			e.Rel.Table, ref.Rel.Table = table, table
			if !e.M2M() {
				e.Rel.Columns = []string{column}
				ref.Rel.Columns = []string{column}
			}
		// assoc with uninitialized relation.
		case !e.IsInverse() && e.Rel.Type == Unk:
			switch {
			case !e.Unique && e.Type == t:
				e.Rel.Type = M2M
				e.SelfRef = true
				e.Rel.Table = t.Label() + "_" + e.Name
				c1, c2 := e.Owner.Label()+"_id", rules.Singularize(e.Name)+"_id"
				e.Rel.Columns = append(e.Rel.Columns, c1, c2)
			case e.Unique && e.Type == t:
				e.Rel.Type = O2O
				e.SelfRef = true
				e.Rel.Table = t.Table()
			case e.Unique:
				e.Rel.Type = M2O
				e.Rel.Table = t.Table()
			default:
				e.Rel.Type = O2M
				e.Rel.Table = e.Type.Table()
			}
			if !e.M2M() {
				// Unlike assoc edges with inverse, we need to choose a unique name for the
				// column in order to no conflict with other types that point to this type.
				e.Rel.Columns = []string{fmt.Sprintf("%s_%s_id", t.Label(), snake(rules.Singularize(e.Name)))}
			}
		}
	}
	return nil
}

// Tables returns the schema definitions of SQL tables for the graph.
func (g *Graph) Tables() (all []*schema.Table) {
	nullable := true
	tables := make(map[string]*schema.Table)
	for _, n := range g.Nodes {
		table := schema.NewTable(n.Table()).AddPrimary(n.ID.Column())
		for _, f := range n.Fields {
			table.Columns = append(table.Columns, f.Column())
		}
		tables[table.Name] = table
		all = append(all, table)
	}
	for _, n := range g.Nodes {
		// foreign key + reference OR join table.
		for _, e := range n.Edges {
			if e.IsInverse() {
				continue
			}
			switch e.Rel.Type {
			case O2O, O2M:
				// "owner" is the table that owns the relations (we set the foreign-key on)
				// and "ref" is the referenced table.
				owner, ref := tables[e.Rel.Table], tables[n.Table()]
				column := &schema.Column{Name: e.Rel.Column(), Type: field.TypeInt, Unique: e.Rel.Type == O2O, Nullable: &nullable}
				owner.Columns = append(owner.Columns, column)
				owner.ForeignKeys = append(owner.ForeignKeys, &schema.ForeignKey{
					RefTable:   ref,
					OnDelete:   schema.SetNull,
					Columns:    []*schema.Column{column},
					RefColumns: []*schema.Column{ref.PrimaryKey[0]},
					Symbol:     fmt.Sprintf("%s_%s_%s", owner.Name, ref.Name, e.Name),
				})
			case M2O:
				ref, owner := tables[e.Type.Table()], tables[e.Rel.Table]
				column := &schema.Column{Name: e.Rel.Column(), Type: field.TypeInt, Nullable: &nullable}
				owner.Columns = append(owner.Columns, column)
				owner.ForeignKeys = append(owner.ForeignKeys, &schema.ForeignKey{
					RefTable:   ref,
					OnDelete:   schema.SetNull,
					Columns:    []*schema.Column{column},
					RefColumns: []*schema.Column{ref.PrimaryKey[0]},
					Symbol:     fmt.Sprintf("%s_%s_%s", owner.Name, ref.Name, e.Name),
				})
			case M2M:
				t1, t2 := tables[n.Table()], tables[e.Type.Table()]
				c1 := &schema.Column{Name: e.Rel.Columns[0], Type: field.TypeInt}
				c2 := &schema.Column{Name: e.Rel.Columns[1], Type: field.TypeInt}
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
	return
}

// migrateSupport reports if the codegen needs to support schema migratio.
func (g *Graph) migrateSupport() bool {
	for _, storage := range g.Storage {
		if storage.SchemaMode.Support(Migrate) {
			return true
		}
	}
	return false
}

func (g *Graph) typ(name string) (*Type, bool) {
	for _, n := range g.Nodes {
		if name == n.Name {
			return n, true
		}
	}
	return nil, false
}

func imports() map[string]string {
	var (
		specs = make(map[string]string)
		b     = bytes.NewBuffer([]byte("package main\n"))
	)
	check(templates.ExecuteTemplate(b, "import", Type{}), "load imports")
	f, err := parser.ParseFile(token.NewFileSet(), "", b, parser.ImportsOnly)
	check(err, "parse imports")
	for _, spec := range f.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		check(err, "unquote import path")
		specs[filepath.Base(path)] = path
	}
	for _, s := range drivers {
		for _, path := range s.Imports {
			specs[filepath.Base(path)] = path
		}
	}
	return specs
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

// run runs an exec command and returns the stderr if it failed.
func run(cmd *exec.Cmd) error {
	out := bytes.NewBuffer(nil)
	cmd.Stderr = out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("entc/gen: %s", out)
	}
	return nil
}
