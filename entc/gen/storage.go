// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/sql"
)

// A SchemaMode defines what type of schema feature a storage driver support.
type SchemaMode uint

const (
	// Unique defines field and edge uniqueness support.
	Unique SchemaMode = 1 << iota

	// Indexes defines indexes support.
	Indexes

	// Cascade defines cascading operations (e.g. cascade deletion).
	Cascade

	// Migrate defines static schema and migration support (e.g. SQL-based).
	Migrate
)

// Support reports whether m support the given mode.
func (m SchemaMode) Support(mode SchemaMode) bool { return m&mode != 0 }

// Storage driver type for codegen.
type Storage struct {
	Name       string             // storage name.
	Builder    reflect.Type       // query builder type.
	Dialects   []string           // supported dialects.
	IdentName  string             // identifier name (fields and funcs).
	Imports    []string           // import packages needed.
	SchemaMode SchemaMode         // schema mode support.
	Ops        func(*Field) []Op  // storage specific operations.
	OpCode     func(Op) string    // operation code for predicates.
	Init       func(*Graph) error // optional init function.
}

// StorageDrivers holds the storage driver options for entc.
var drivers = []*Storage{
	{
		Name:      "sql",
		IdentName: "SQL",
		Builder:   reflect.TypeOf(&sql.Selector{}),
		Dialects:  []string{"dialect.SQLite", "dialect.MySQL", "dialect.Postgres"},
		Imports: []string{
			"database/sql/driver",
			"entgo.io/ent/dialect/sql",
			"entgo.io/ent/dialect/sql/sqlgraph",
			"entgo.io/ent/dialect/sql/sqljson",
			"entgo.io/ent/schema/field",
		},
		SchemaMode: Unique | Indexes | Cascade | Migrate,
		Ops: func(f *Field) []Op {
			if f.IsString() && f.ConvertedToBasic() {
				return []Op{EqualFold, ContainsFold}
			}
			return nil
		},
		OpCode: opCodes(sqlCode[:]),
		Init: func(g *Graph) error {
			var with, without []string
			for _, n := range g.Nodes {
				if s, err := n.TableSchema(); err == nil && s != "" {
					with = append(with, n.Name)
				} else {
					without = append(without, n.Name)
				}
			}
			switch {
			case len(with) == 0:
				return nil
			case len(without) > 0:
				return fmt.Errorf("missing schema annotation for %s", strings.Join(without, ", "))
			default:
				if !g.featureEnabled(FeatureSchemaConfig) {
					g.Features = append(g.Features, FeatureSchemaConfig)
				}
				if !g.featureEnabled(featureMultiSchema) {
					g.Features = append(g.Features, featureMultiSchema)
				}
				return nil
			}
		},
	},
	{
		Name:      "gremlin",
		IdentName: "Gremlin",
		Builder:   reflect.TypeOf(&dsl.Traversal{}),
		Dialects:  []string{"dialect.Gremlin"},
		Imports: []string{
			"entgo.io/ent/dialect/gremlin",
			"entgo.io/ent/dialect/gremlin/graph/dsl",
			"entgo.io/ent/dialect/gremlin/graph/dsl/__",
			"entgo.io/ent/dialect/gremlin/graph/dsl/g",
			"entgo.io/ent/dialect/gremlin/graph/dsl/p",
			"entgo.io/ent/dialect/gremlin/encoding/graphson",
		},
		SchemaMode: Unique,
		OpCode:     opCodes(gremlinCode[:]),
		Init:       func(*Graph) error { return nil }, // Noop.
	},
}

// NewStorage returns the storage driver type from the given string.
// It fails if the provided string is not a valid option. this function
// is here in order to remove the validation logic from entc command line.
func NewStorage(s string) (*Storage, error) {
	for _, d := range drivers {
		if s == d.Name {
			return d, nil
		}
	}
	return nil, fmt.Errorf("entc/gen: invalid storage driver %q", s)
}

// String implements the fmt.Stringer interface for template usage.
func (s *Storage) String() string { return s.Name }

var (
	// exceptional operation names in sql.
	sqlCode = [...]string{
		IsNil:  "IsNull",
		NotNil: "NotNull",
	}
	// exceptional operation names in gremlin.
	gremlinCode = [...]string{
		IsNil:     "HasNot",
		NotNil:    "Has",
		In:        "Within",
		NotIn:     "Without",
		Contains:  "Containing",
		HasPrefix: "StartingWith",
		HasSuffix: "EndingWith",
	}
)

func opCodes(codes []string) func(Op) string {
	return func(o Op) string {
		if int(o) < len(codes) && codes[o] != "" {
			return codes[o]
		}
		return o.Name()
	}
}

// TableSchemas returns all table schemas in ent/schema (intentionally exported).
func (g *Graph) TableSchemas() ([]string, error) {
	all := make(map[string]struct{})
	for _, n := range g.Nodes {
		s, err := n.TableSchema()
		if err != nil {
			return nil, err
		}
		all[s] = struct{}{}
		for _, e := range n.Edges {
			// {{- if and $e.M2M (not $e.Inverse) (not $e.Through) }}
			if e.M2M() && !e.IsInverse() && e.Through == nil {
				s, err := e.TableSchema()
				if err != nil {
					return nil, err
				}
				all[s] = struct{}{}
			}
		}
	}
	names := make([]string, 0, len(all))
	for s := range all {
		names = append(names, s)
	}
	sort.Strings(names)
	return names, nil
}

// TableSchema returns the schema name of where the type table resides (intentionally exported).
func (t *Type) TableSchema() (string, error) {
	switch ant := t.EntSQL(); {
	case ant == nil || ant.Schema == "":
		return "", fmt.Errorf("atlas: missing schema annotation for node %q", t.Name)
	default:
		return ant.Schema, nil
	}
}

// TableSchema returns the schema name of where the type table resides (intentionally exported).
func (e *Edge) TableSchema() (string, error) {
	switch ant := e.EntSQL(); {
	case ant == nil || ant.Schema == "":
		return e.Owner.TableSchema()
	default:
		return ant.Schema, nil
	}
}
