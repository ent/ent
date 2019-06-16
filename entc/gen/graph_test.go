package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"fbc/ent"
	"fbc/ent/edge"
	"fbc/ent/field"

	"github.com/stretchr/testify/require"
)

type T1 struct {
	ent.Schema
}

func (T1) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Optional(),
		field.Time("expired_at").Nullable(),
		field.String("name").Default("hello"),
	}
}

func (T1) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("t2", T2.Type).Required(),
		edge.To("t1", T1.Type).Unique(),
		// Bidirectional unique edge (unique/"has-a" in both sides).
		edge.To("t2_o2o", T2.Type).Unique(),
		// Unidirectional non-unique edge ("has-many"). The reference is on the "many" side.
		// For example: A user "has-many" books, but a book "has-an" owner (and only one).
		edge.To("o2m", T2.Type),
		// Unidirectional unique edge ("has-one").
		// For example: A user "has-an" address (and only one), but an address "has-many" users.
		edge.To("m2o", T2.Type).Unique(),
		// Bidirectional unique edge ("has-one" in T1 side, and "has-many" in T2 side).
		edge.To("t2_m2o", T2.Type).Unique(),
		// Bidirectional non-unique edge ("has-many" in T1 side, and "has-one" in T2 side).
		edge.To("t2_o2m", T2.Type),
		// Bidirectional non-unique edge ("has-many" in both side).
		edge.To("t2_m2m", T2.Type),
		// Unidirectional non-unique edge for the same type.
		edge.To("t1_m2m", T1.Type),
	}
}

type T2 struct {
	ent.Schema
}

func (T2) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active"),
	}
}

func (T2) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("t1", T1.Type).Ref("t2"),
		edge.From("t1_o2o", T1.Type).Unique().Ref("t2_o2o"),
		edge.From("t1_o2m", T1.Type).Ref("t2_m2o"),
		edge.From("t1_m2o", T1.Type).Ref("t2_o2m").Unique(),
		edge.From("t1_m2m", T1.Type).Ref("t2_m2m"),
	}
}

func TestNewGraph(t *testing.T) {
	require := require.New(t)
	graph, err := NewGraph(Config{Package: "entc/gen"}, T1{})
	require.Error(err, "should fail due to missing types")

	graph, err = NewGraph(Config{Package: "entc/gen"}, T1{}, T2{})
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
	for i, optional := range []bool{true, false, false} {
		require.Equal(optional, t1.Fields[i].Optional)
	}
	for i, nullable := range []bool{false, true, false} {
		require.Equal(nullable, t1.Fields[i].Nullable)
	}
	for i, value := range []interface{}{nil, nil, "hello"} {
		require.Equal(value, t1.Fields[i].Default)
		require.Equal(value != nil, t1.Fields[i].HasDefault())
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
	f1, e1 := t2.Fields[0], t2.Edges[0]
	require.Equal("bool", f1.Type.String())
	require.Equal("active", f1.Name)
	require.Equal("t1", e1.Name)
	require.True(e1.IsInverse())
	require.Equal("t2", e1.Inverse)
	require.Equal(graph.Nodes[0], e1.Type)
}

func TestRelation(t *testing.T) {
	require := require.New(t)
	graph, err := NewGraph(Config{Package: "entc/gen"}, T1{})
	require.Error(err, "should fail due to missing types")

	graph, err = NewGraph(Config{Package: "entc/gen"}, T1{}, T2{})
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

func TestGraph_Gen(t *testing.T) {
	require := require.New(t)
	target := filepath.Join(os.TempDir(), "ent")
	require.NoError(os.MkdirAll(target, os.ModePerm), "creating tmpdir")
	defer os.Remove(target)
	graph, err := NewGraph(Config{Package: "entc/gen", Target: target}, T1{}, T2{})
	require.NoError(err)
	require.NotNil(graph)
	require.NoError(graph.Gen())
	// ensure graph files were generated.
	for _, name := range []string{"ent", "client", "config", "example_test"} {
		_, err := os.Stat(fmt.Sprintf("%s/%s.go", target, name))
		require.NoError(err)
	}
	// ensure entity files were generated.
	for _, format := range []string{"%s", "%s_create", "%s_update", "%s_delete", "%s_query"} {
		for _, name := range []string{"t1", "t2"} {
			_, err := os.Stat(fmt.Sprintf(fmt.Sprintf("%s/%s.go", target, format), name))
			require.NoError(err)
		}
	}
}
