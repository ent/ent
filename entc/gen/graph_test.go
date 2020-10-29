// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/facebook/ent/entc/load"
	"github.com/facebook/ent/schema/field"

	"github.com/stretchr/testify/require"
)

var (
	T1 = &load.Schema{
		Name: "T1",
		Fields: []*load.Field{
			{Name: "age", Info: &field.TypeInfo{Type: field.TypeInt}, Optional: true},
			{Name: "expired_at", Info: &field.TypeInfo{Type: field.TypeTime}, Nillable: true, Optional: true},
			{Name: "name", Info: &field.TypeInfo{Type: field.TypeString}, Default: true},
		},
		Edges: []*load.Edge{
			{Name: "t2", Type: "T2", Required: true},
			{Name: "t1", Type: "T1", Unique: true},
			// Bidirectional unique edge (unique/"has-a" in both sides).
			{Name: "t2_o2o", Type: "T2", Unique: true},
			// Unidirectional non-unique edge ("has-many"). The reference is on the "many" side.
			// For example: A user "has-many" books, but a book "has-an" owner (and only one).
			{Name: "o2m", Type: "T2"},
			// Unidirectional unique edge ("has-one").
			// For example: A user "has-an" address (and only one), but an address "has-many" users.
			{Name: "m2o", Type: "T2", Unique: true},
			// Bidirectional unique edge ("has-one" in T1 side, and "has-many" in T2 side).
			{Name: "t2_m2o", Type: "T2", Unique: true},
			// Bidirectional non-unique edge ("has-many" in T1 side, and "has-one" in T2 side).
			{Name: "t2_o2m", Type: "T2"},
			// Bidirectional non-unique edge ("has-many" in both side).
			{Name: "t2_m2m", Type: "T2"},
			// Unidirectional non-unique edge for the same type.
			{Name: "t1_m2m", Type: "T1"},
		},
	}
	T2 = &load.Schema{
		Name:        "T2",
		Annotations: dict("GQL", map[string]string{"Name": "T2"}),
		Fields: []*load.Field{
			{Name: "active", Info: &field.TypeInfo{Type: field.TypeBool}},
		},
		Edges: []*load.Edge{
			{Name: "t1", Type: "T1", RefName: "t2", Inverse: true},
			{Name: "t1_o2o", Type: "T1", RefName: "t2_o2o", Unique: true, Inverse: true},
			{Name: "t1_o2m", Type: "T1", RefName: "t2_m2o", Inverse: true},
			{Name: "t1_m2o", Type: "T1", RefName: "t2_o2m", Unique: true, Inverse: true},
			{Name: "t1_m2m", Type: "T1", RefName: "t2_m2m", Inverse: true},
			{Name: "t2_m2m_from", Type: "T2", Ref: &load.Edge{Name: "t2_m2m_to", Type: "T2", Annotations: dict("GQL", map[string]string{"Name": "To"})}, Inverse: true, Annotations: dict("GQL", map[string]string{"Name": "From"})},
		},
	}
)

func TestNewGraph(t *testing.T) {
	require := require.New(t)
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, T1)
	require.Error(err, "should fail due to missing types")

	graph, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, T1, T2)
	require.NoError(err)
	require.NotNil(graph)
	require.Len(graph.Nodes, 2)

	t1 := graph.Nodes[0]

	// check fields.
	require.Equal("T1", t1.Name)
	require.Len(t1.Fields, 3)
	for i, name := range []string{"age", "expired_at", "name"} {
		require.Equal(name, t1.Fields[i].Name)
	}
	for i, typ := range []string{"int", "time.Time", "string"} {
		require.Equal(typ, t1.Fields[i].Type.String())
	}
	for i, optional := range []bool{true, true, false} {
		require.Equal(optional, t1.Fields[i].Optional)
	}
	for i, nullable := range []bool{false, true, false} {
		require.Equal(nullable, t1.Fields[i].Nillable)
	}
	for i, value := range []bool{false, false, true} {
		require.Equal(value, t1.Fields[i].Default)
	}

	// check edges.
	require.Len(t1.Edges, 9)
	for i, name := range []string{"t2", "t1"} {
		require.Equal(name, t1.Edges[i].Name)
	}
	for i, typ := range []*Type{graph.Nodes[1], graph.Nodes[0]} {
		require.Equal(typ, t1.Edges[i].Type, "edge should point to the right type")
	}
	for i, optional := range []bool{false, true} {
		require.Equal(optional, t1.Edges[i].Optional)
	}
	for i, unique := range []bool{false, true} {
		require.Equal(unique, t1.Edges[i].Unique)
	}
	for i, inverse := range []bool{false, false} {
		require.Equal(inverse, t1.Edges[i].IsInverse())
	}

	t2 := graph.Nodes[1]
	require.Equal(map[string]string{"Name": "T2"}, t2.Annotations["GQL"])
	f1, e1 := t2.Fields[0], t2.Edges[0]
	require.Equal("bool", f1.Type.String())
	require.Equal("active", f1.Name)
	require.Equal("t1", e1.Name)
	require.True(e1.IsInverse())
	require.Equal("t2", e1.Inverse)
	require.Equal(graph.Nodes[0], e1.Type)

	require.Equal("t2_m2m_from", t2.Edges[5].Name)
	require.Equal("t2_m2m_to", t2.Edges[6].Name)
	require.Equal(map[string]string{"Name": "From"}, t2.Edges[5].Annotations["GQL"])
	require.Equal(map[string]string{"Name": "To"}, t2.Edges[6].Annotations["GQL"])
}

func TestNewGraphRequiredLoop(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, &load.Schema{
		Name: "T1",
		Edges: []*load.Edge{
			{Name: "parent", Type: "T1", Unique: true, Required: true},
			{Name: "children", Type: "T1", Inverse: true, RefName: "parent", Required: true},
		},
	})
	require.Error(t, err, "require loop")

	_, err = NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]},
		&load.Schema{
			Name: "User",
			Edges: []*load.Edge{
				{Name: "pets", Type: "Pet", Required: true},
			},
		},
		&load.Schema{
			Name: "Pet",
			Edges: []*load.Edge{
				{Name: "owner", Type: "User", Inverse: true, RefName: "pets", Unique: true, Required: true},
			},
		})
	require.Error(t, err, "require loop")
}

func TestNewGraphBadInverse(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]},
		&load.Schema{
			Name: "User",
			Edges: []*load.Edge{
				{Name: "pets", Type: "Pet"},
				{Name: "groups", Type: "Group"},
			},
		},
		&load.Schema{
			Name: "Pet",
			Edges: []*load.Edge{
				{Name: "owner", Type: "User", Unique: true, Required: true, RefName: "pets", Inverse: true},
			},
		},
		&load.Schema{
			Name: "Group",
			Edges: []*load.Edge{
				{Name: "users", Type: "User", RefName: "pets", Inverse: true},
			},
		})
	require.Errorf(t, err, "mismatch type for back-reference")
}

func TestNewGraphDuplicateEdges(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]},
		&load.Schema{
			Name: "User",
			Edges: []*load.Edge{
				{Name: "groups", Type: "Group"},
				{Name: "groups", Type: "Group", RefName: "owner", Inverse: true},
			},
		},
		&load.Schema{
			Name: "Group",
			Edges: []*load.Edge{
				{Name: "users", Type: "User", RefName: "groups", Inverse: true},
				{Name: "owner", Type: "User", Unique: true},
			},
		})
	require.EqualError(t, err, `entc/gen: User schema contains multiple "groups" edges`)
}

func TestRelation(t *testing.T) {
	require := require.New(t)
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, T1)
	require.Error(err, "should fail due to missing types")

	graph, err := NewGraph(&Config{Package: "entc/gen"}, T1, T2)
	require.NoError(err)
	require.NotNil(graph)
	require.Len(graph.Nodes, 2)

	t1, t2 := graph.Nodes[0], graph.Nodes[1]
	// unidirectional one 2 one.
	require.Equal(O2O, t1.Edges[1].Rel.Type)
	// bidirectional one to one.
	require.Equal(O2O, t1.Edges[2].Rel.Type)
	require.Equal(O2O, t2.Edges[1].Rel.Type)
	// unidirectional one 2 many.
	require.Equal(O2M, t1.Edges[3].Rel.Type)
	// unidirectional many 2 one.
	require.Equal(M2O, t1.Edges[4].Rel.Type)
	// bidirectional many 2 one.
	require.Equal(M2O, t1.Edges[5].Rel.Type)
	require.Equal(O2M, t2.Edges[2].Rel.Type)
	// bidirectional one 2 many.
	require.Equal(O2M, t1.Edges[6].Rel.Type)
	require.Equal(M2O, t2.Edges[3].Rel.Type)
	// bidirectional many 2 many.
	require.Equal(M2M, t1.Edges[7].Rel.Type)
	require.Equal(M2M, t2.Edges[4].Rel.Type)
	// unidirectional many 2 many.
	require.Equal(M2M, t1.Edges[8].Rel.Type)
}

func TestFKColumns(t *testing.T) {
	user := &load.Schema{
		Name: "User",
		Edges: []*load.Edge{
			{Name: "pets", Type: "Pet"},
			{Name: "pet", Type: "Pet", Unique: true},
			{Name: "parent", Type: "User", Unique: true},
		},
	}
	require := require.New(t)
	graph, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, user, &load.Schema{Name: "Pet"})
	require.NoError(err)
	t1 := graph.Nodes[0]
	require.Equal(Relation{Type: O2M, Table: "pets", Columns: []string{"user_pets"}}, t1.Edges[0].Rel)
	require.Equal(Relation{Type: M2O, Table: "users", Columns: []string{"user_pet"}}, t1.Edges[1].Rel)
	require.Equal(Relation{Type: O2O, Table: "users", Columns: []string{"user_parent"}}, t1.Edges[2].Rel)

	// Adding inverse edges.
	graph, err = NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, user,
		&load.Schema{
			Name: "Pet",
			Edges: []*load.Edge{
				{Name: "owner", Type: "User", RefName: "pets", Inverse: true, Unique: true},
				{Name: "team", Type: "User", RefName: "pet", Inverse: true},
			},
		},
	)
	require.NoError(err)
	t1, t2 := graph.Nodes[0], graph.Nodes[1]
	require.Equal(Relation{Type: O2M, Table: "pets", Columns: []string{"user_pets"}}, t1.Edges[0].Rel)
	require.Equal(Relation{Type: M2O, Table: "users", Columns: []string{"user_pet"}}, t1.Edges[1].Rel)
	require.Equal(Relation{Type: M2O, Table: "pets", Columns: []string{"user_pets"}}, t2.Edges[0].Rel)
	require.Equal(Relation{Type: O2M, Table: "users", Columns: []string{"user_pet"}}, t2.Edges[1].Rel)
}

func TestGraph_Gen(t *testing.T) {
	require := require.New(t)
	target := filepath.Join(os.TempDir(), "ent")
	require.NoError(os.MkdirAll(target, os.ModePerm), "creating tmpdir")
	defer os.RemoveAll(target)
	external := MustParse(NewTemplate("external").Parse("package external"))
	graph, err := NewGraph(&Config{
		Package:   "entc/gen",
		Target:    target,
		Storage:   drivers[0],
		Templates: []*Template{external},
		IDType:    &field.TypeInfo{Type: field.TypeInt},
	}, &load.Schema{
		Name: "T1",
		Fields: []*load.Field{
			{Name: "age", Info: &field.TypeInfo{Type: field.TypeInt}, Optional: true},
			{Name: "expired_at", Info: &field.TypeInfo{Type: field.TypeTime}, Nillable: true, Optional: true},
			{Name: "name", Info: &field.TypeInfo{Type: field.TypeString}},
		},
		Edges: []*load.Edge{
			{Name: "t1", Type: "T1", Unique: true},
		},
	})
	require.NoError(err)
	require.NotNil(graph)
	require.NoError(graph.Gen())
	// ensure graph files were generated.
	for _, name := range []string{"ent", "client", "config"} {
		_, err := os.Stat(fmt.Sprintf("%s/%s.go", target, name))
		require.NoError(err)
	}
	// ensure entity files were generated.
	for _, format := range []string{"%s", "%s_create", "%s_update", "%s_delete", "%s_query"} {
		_, err := os.Stat(fmt.Sprintf(fmt.Sprintf("%s/%s.go", target, format), "t1"))
		require.NoError(err)
	}
	_, err = os.Stat(target + "/external.go")
	require.NoError(err)
}

func ensureStructTag(name string) Hook {
	return func(next Generator) Generator {
		return GenerateFunc(func(g *Graph) error {
			// Ensure all fields have a specific tag.
			for _, node := range g.Nodes {
				for _, f := range node.Fields {
					tag := reflect.StructTag(f.StructTag)
					if _, ok := tag.Lookup(name); !ok {
						return fmt.Errorf("struct tag %q is missing for field %s.%s", name, node.Name, f.Name)
					}
				}
			}
			return next.Generate(g)
		})
	}
}

func TestGraph_Hooks(t *testing.T) {
	require := require.New(t)
	graph, err := NewGraph(&Config{
		Package: "entc/gen",
		Storage: drivers[0],
		IDType:  &field.TypeInfo{Type: field.TypeInt},
		Hooks:   []Hook{ensureStructTag("yaml")},
	}, &load.Schema{
		Name: "T1",
		Fields: []*load.Field{
			{Name: "age", Info: &field.TypeInfo{Type: field.TypeInt}, Optional: true},
			{Name: "expired_at", Info: &field.TypeInfo{Type: field.TypeTime}, Nillable: true, Optional: true},
			{Name: "name", Info: &field.TypeInfo{Type: field.TypeString}},
		},
		Edges: []*load.Edge{
			{Name: "t1", Type: "T1", Unique: true},
		},
	})
	require.NoError(err)
	require.NotNil(graph)
	require.EqualError(graph.Gen(), `struct tag "yaml" is missing for field T1.age`)
}
