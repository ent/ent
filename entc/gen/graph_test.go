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

	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

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
	require.Equal("t2_m2m_to", t2.Edges[5].Inverse)
	require.Equal("t2_m2m_to", t2.Edges[6].Name)
	require.Empty(t2.Edges[6].Inverse)
	require.Equal(t2.Edges[6], t2.Edges[5].Ref)
	require.Equal(t2.Edges[5], t2.Edges[6].Ref)
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

func TestNewGraphDuplicateEdgeField(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]},
		&load.Schema{
			Name: "User",
			Fields: []*load.Field{
				{Name: "parent", Info: &field.TypeInfo{Type: field.TypeInt}},
			},
			Edges: []*load.Edge{
				{Name: "parent", Type: "User"},
			},
		})
	require.EqualError(t, err, `entc/gen: User schema cannot contain field and edge with the same name "parent"`)
}

func TestNewGraphThroughUndefinedType(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, &load.Schema{
		Name: "T1",
		Edges: []*load.Edge{
			{Name: "groups", Type: "T1", Required: true, Through: &struct{ N, T string }{N: "groups_edge", T: "T2"}},
		},
	})
	require.EqualError(t, err, `entc/gen: resolving edges: edge T1.groups defined with Through("groups_edge", T2.Type), but type T2 was not found`)
}

func TestNewGraphThroughInvalidRel(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, &load.Schema{
		Name: "T1",
		Edges: []*load.Edge{
			{Name: "groups", Type: "T1", Unique: true, Required: true, Through: &struct{ N, T string }{N: "groups_edge", T: "T2"}},
		},
	})
	require.EqualError(t, err, `entc/gen: resolving edges: edge T1.groups Through("groups_edge", T2.Type) is allowed only on M2M edges, but got: "O2O"`)
}

func TestNewGraphThroughDuplicates(t *testing.T) {
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]},
		&load.Schema{
			Name: "User",
			Edges: []*load.Edge{
				{Name: "groups", Type: "Group", Through: &struct{ N, T string }{N: "group_edges", T: "T1"}},
				{Name: "group_edges", Type: "Group"},
			},
		},
		&load.Schema{
			Name: "Group",
			Edges: []*load.Edge{
				{Name: "users", Type: "User", Inverse: true, RefName: "groups", Through: &struct{ N, T string }{N: "user_edges", T: "T1"}},
			},
		},
		&load.Schema{
			Name: "T1",
		},
	)
	require.EqualError(t, err, `entc/gen: resolving edges: edge User.groups defined with Through("group_edges", T1.Type), but schema User already has an edge named group_edges`)
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
	for i, r := range []Relation{
		{Type: O2M, Table: "pets", Columns: []string{"user_pets"}},
		{Type: M2O, Table: "users", Columns: []string{"user_pet"}},
		{Type: O2O, Table: "users", Columns: []string{"user_parent"}},
	} {
		require.Equal(r.Type, t1.Edges[i].Rel.Type)
		require.Equal(r.Table, t1.Edges[i].Rel.Table)
		require.Equal(r.Columns, t1.Edges[i].Rel.Columns)
	}

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
	for i, r := range []Relation{
		{Type: O2M, Table: "pets", Columns: []string{"user_pets"}},
		{Type: M2O, Table: "users", Columns: []string{"user_pet"}},
	} {
		require.Equal(r.Type, t1.Edges[i].Rel.Type)
		require.Equal(r.Table, t1.Edges[i].Rel.Table)
		require.Equal(r.Columns, t1.Edges[i].Rel.Columns)
	}
	for i, r := range []Relation{
		{Type: M2O, Table: "pets", Columns: []string{"user_pets"}},
		{Type: O2M, Table: "users", Columns: []string{"user_pet"}},
	} {
		require.Equal(r.Type, t2.Edges[i].Rel.Type)
		require.Equal(r.Table, t2.Edges[i].Rel.Table)
		require.Equal(r.Columns, t2.Edges[i].Rel.Columns)
	}
}

func TestAbortDuplicateFK(t *testing.T) {
	var (
		user = &load.Schema{
			Name: "User",
			Edges: []*load.Edge{
				{Name: "pets", Type: "Pet", StorageKey: &edge.StorageKey{Symbols: []string{"owner_id"}}},
				{Name: "cars", Type: "Car", StorageKey: &edge.StorageKey{Symbols: []string{"owner_id"}}},
			},
		}
		pet = &load.Schema{
			Name: "Pet",
			Fields: []*load.Field{
				{Name: "owner_id", Info: &field.TypeInfo{Type: field.TypeInt}, Nillable: true, Optional: true},
			},
			Edges: []*load.Edge{
				{Name: "owner", Type: "User", RefName: "pets", Inverse: true, Unique: true},
			},
		}
		car = &load.Schema{
			Name: "Car",
			Fields: []*load.Field{
				{Name: "owner_id", Info: &field.TypeInfo{Type: field.TypeInt}, Nillable: true, Optional: true},
			},
			Edges: []*load.Edge{
				{Name: "owner", Type: "User", RefName: "cars", Inverse: true, Unique: true},
			},
		}
	)
	g, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, user, pet, car)
	require.NoError(t, err)
	_, err = g.Tables()
	require.EqualError(t, err, `duplicate foreign-key symbol "owner_id" found in tables "cars" and "pets"`)
}

func TestEnsureCorrectFK(t *testing.T) {
	var (
		user = &load.Schema{
			Name: "User",
			Edges: []*load.Edge{
				{Name: "pets", Type: "Pet", StorageKey: &edge.StorageKey{Columns: []string{"owner_id"}}},
			},
		}
		pet = &load.Schema{
			Name: "Pet",
			Fields: []*load.Field{
				{Name: "owner_id", Info: &field.TypeInfo{Type: field.TypeInt}, Nillable: true, Optional: true},
			},
			Edges: []*load.Edge{
				{Name: "owner", Type: "User", RefName: "pets", Inverse: true, Unique: true},
			},
		}
	)
	_, err := NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, user, pet)
	require.EqualError(t, err, `entc/gen: set "User" foreign-keys: column "owner_id" definition on edge "pets" should be replaced with Field("owner_id") on its reference "owner"`)

	user.Edges[0].StorageKey = nil
	pet.Edges[0].Field = "owner_id"
	_, err = NewGraph(&Config{Package: "entc/gen", Storage: drivers[0]}, user, pet)
	require.NoError(t, err)
}

func TestGraph_Gen(t *testing.T) {
	require := require.New(t)
	target := filepath.Join(t.TempDir(), "ent")
	external := MustParse(NewTemplate("external").Parse("package external"))
	skipped := MustParse(NewTemplate("skipped").SkipIf(func(*Graph) bool { return true }).Parse("package external"))
	schemas := []*load.Schema{
		{
			Name: "T1",
			Fields: []*load.Field{
				{Name: "age", Info: &field.TypeInfo{Type: field.TypeInt}, Optional: true},
				{Name: "expired_at", Info: &field.TypeInfo{Type: field.TypeTime}, Nillable: true, Optional: true},
				{Name: "name", Info: &field.TypeInfo{Type: field.TypeString}},
			},
			Edges: []*load.Edge{
				{Name: "t1", Type: "T1", Unique: true},
			},
		},
		{Name: "T2"},
		{Name: "T3"},
	}
	graph, err := NewGraph(&Config{
		Package:   "entc/gen",
		Target:    target,
		Storage:   drivers[0],
		Templates: []*Template{external, skipped},
		IDType:    &field.TypeInfo{Type: field.TypeInt},
		Features:  AllFeatures,
	}, schemas...)
	require.NoError(err)
	require.NotNil(graph)
	require.NoError(graph.Gen())
	// Ensure globalid feature added annotations.
	a := IncrementStarts{"t1s": 0, "t2s": 1 << 32, "t3s": 2 << 32}
	require.Equal(a, graph.Annotations[a.Name()])
	for i, n := range graph.Nodes {
		require.Equal(i<<32, *n.EntSQL().IncrementStart)
	}
	// Ensure graph files were generated.
	for _, name := range []string{"ent", "client"} {
		_, err := os.Stat(fmt.Sprintf("%s/%s.go", target, name))
		require.NoError(err)
	}
	// Ensure entity files were generated.
	for _, format := range []string{"%s", "%s_create", "%s_update", "%s_delete", "%s_query"} {
		_, err := os.Stat(fmt.Sprintf(fmt.Sprintf("%s/%s.go", target, format), "t1"))
		require.NoError(err)
		_, err = os.Stat(fmt.Sprintf(fmt.Sprintf("%s/%s.go", target, format), "t2"))
		require.NoError(err)
	}
	_, err = os.Stat(filepath.Join(target, "external.go"))
	require.NoError(err)
	_, err = os.Stat(filepath.Join(target, "skipped.go"))
	require.True(os.IsNotExist(err))

	// Generated feature templates.
	_, err = os.Stat(filepath.Join(target, "internal", "schema.go"))
	require.NoError(err)
	_, err = os.Stat(filepath.Join(target, "internal", "schemaconfig.go"))
	require.NoError(err)
	c, err := os.ReadFile(filepath.Join(target, "internal", "globalid.go"))
	require.NoError(err)
	require.Contains(string(c), fmt.Sprintf(`"{\"t1s\":0,\"t2s\":%d,\"t3s\":%d}"`, 1<<32, 2<<32))
	// Rerun codegen with only one feature-flag.
	graph.Features = []Feature{FeatureSnapshot}
	require.NoError(graph.Gen())
	// Generated feature templates.
	_, err = os.Stat(filepath.Join(target, "internal", "schema.go"))
	require.NoError(err)
	_, err = os.Stat(filepath.Join(target, "internal", "schemaconfig.go"))
	require.True(os.IsNotExist(err))
	_, err = os.Stat(filepath.Join(target, "internal", "globalid.go"))
	require.True(os.IsNotExist(err))
	// Rerun codegen without any feature-flags.
	graph.Features = nil
	require.NoError(graph.Gen())
	_, err = os.Stat(filepath.Join(target, "internal"))
	require.True(os.IsNotExist(err))

	schemas = schemas[:1]
	graph, err = NewGraph(&Config{
		Package:   "entc/gen",
		Target:    target,
		Storage:   drivers[0],
		Templates: []*Template{external, skipped},
		IDType:    &field.TypeInfo{Type: field.TypeInt},
		Features:  AllFeatures,
	}, schemas...)
	require.NoError(err)
	require.NotNil(graph)
	require.NoError(graph.Gen())
	// Ensure entity files were generated.
	for _, format := range []string{"%s", "%s_create", "%s_update", "%s_delete", "%s_query"} {
		_, err := os.Stat(fmt.Sprintf(fmt.Sprintf("%s/%s.go", target, format), "t1"))
		require.NoError(err)
		_, err = os.Stat(fmt.Sprintf(fmt.Sprintf("%s/%s.go", target, format), "t2"))
		require.Error(err)
	}
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

func TestDependencyAnnotation_Build(t *testing.T) {
	tests := []struct {
		typ   *field.TypeInfo
		field string
	}{
		{
			typ: &field.TypeInfo{
				Ident: "*http.Client",
			},
			field: "HTTPClient",
		},
		{
			typ: &field.TypeInfo{
				Ident: "[]*http.Client",
				RType: &field.RType{
					Kind: reflect.Slice,
				},
			},
			field: "HTTPClients",
		},
		{
			typ: &field.TypeInfo{
				Ident: "[]*url.URL",
				RType: &field.RType{
					Kind: reflect.Slice,
				},
			},
			field: "URLs",
		},
		{
			typ: &field.TypeInfo{
				Ident: "*net.Conn",
			},
			field: "NetConn",
		},
	}
	for _, tt := range tests {
		d := &Dependency{Type: tt.typ}
		require.NoError(t, d.Build())
		require.Equal(t, tt.field, d.Field)
	}
}
