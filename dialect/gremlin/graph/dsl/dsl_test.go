// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dsl_test

import (
	"strconv"
	"testing"

	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"

	"github.com/stretchr/testify/require"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		input     dsl.Querier
		wantQuery string
		wantBinds dsl.Bindings
	}{
		{
			input:     g.V(5),
			wantQuery: "g.V($0)",
			wantBinds: dsl.Bindings{"$0": 5},
		},
		{
			input:     g.V(2).Both("knows"),
			wantQuery: "g.V($0).both($1)",
			wantBinds: dsl.Bindings{"$0": 2, "$1": "knows"},
		},
		{
			input:     g.V(49).BothE("knows").OtherV().ValueMap(),
			wantQuery: "g.V($0).bothE($1).otherV().valueMap()",
			wantBinds: dsl.Bindings{"$0": 49, "$1": "knows"},
		},
		{
			input:     g.AddV("person").Property("name", "a8m").Next(),
			wantQuery: "g.addV($0).property($1, $2).next()",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "name", "$2": "a8m"},
		},
		{
			input: dsl.Each([]any{1, 2, 3}, func(it *dsl.Traversal) *dsl.Traversal {
				return g.V(it)
			}),
			wantQuery: "[$0, $1, $2].each { g.V(it) }",
			wantBinds: dsl.Bindings{"$0": 1, "$1": 2, "$2": 3},
		},
		{
			input: dsl.Each([]any{g.V(1).Next()}, func(it *dsl.Traversal) *dsl.Traversal {
				return it.ID()
			}),
			wantQuery: "[g.V($0).next()].each { it.id() }",
			wantBinds: dsl.Bindings{"$0": 1},
		},
		{
			input:     g.AddV("person").AddE("knows").To(g.V(2)),
			wantQuery: "g.addV($0).addE($1).to(g.V($2))",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "knows", "$2": 2},
		},
		{
			input: func() *dsl.Traversal {
				v1 := g.V(2).Next()
				v2 := g.AddV("person").Property("name", "a8m")
				e1 := g.V(v1).AddE("knows").To(v2)
				return dsl.Group(v1, v2, e1)
			}(),
			wantQuery: "t0 = g.V($0).next(); t1 = g.addV($1).property($2, $3); t2 = g.V(t0).addE($4).to(t1); t2",
			wantBinds: dsl.Bindings{"$0": 2, "$1": "person", "$2": "name", "$3": "a8m", "$4": "knows"},
		},
		{
			input: func() *dsl.Traversal {
				v1 := g.AddV("person")
				each := dsl.Each([]any{1, 2, 3}, func(it *dsl.Traversal) *dsl.Traversal {
					return g.V(v1).AddE("knows").To(g.V(it)).Next()
				})
				return dsl.Group(v1, each)
			}(),
			wantQuery: "t0 = g.addV($0); t1 = [$1, $2, $3].each { g.V(t0).addE($4).to(g.V(it)).next() }; t1",
			wantBinds: dsl.Bindings{"$0": "person", "$1": 1, "$2": 2, "$3": 3, "$4": "knows"},
		},
		{
			input: g.V().HasLabel("person").
				Choose(__.Values("age").Is(p.LTE(20))),
			wantQuery: "g.V().hasLabel($0).choose(__.values($1).is(lte($2)))",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "age", "$2": 20},
		},
		{
			input:     g.AddV("person").Property("name", "a8m").Properties(),
			wantQuery: "g.addV($0).property($1, $2).properties()",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "name", "$2": "a8m"},
		},
		{
			input: func() *dsl.Traversal {
				v1 := g.AddV("person").Next()
				e1 := g.V(v1).AddE("knows").To(g.V(2).Next())
				return dsl.Group(v1, e1, g.V(v1).ValueMap(true))
			}(),
			wantQuery: "t0 = g.addV($0).next(); t1 = g.V(t0).addE($1).to(g.V($2).next()); t2 = g.V(t0).valueMap($3); t2",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "knows", "$2": 2, "$3": true},
		},
		{
			input: func() *dsl.Traversal {
				vs := g.V().HasLabel("person").ToList()
				edge := g.V(vs).AddE("assoc").To(g.V(1)).Iterate()
				each := dsl.Each(vs, func(it *dsl.Traversal) *dsl.Traversal {
					return g.V(1).AddE("inverse").To(it).Next()
				})
				return dsl.Group(vs, edge, each)
			}(),
			wantQuery: "t0 = g.V().hasLabel($0).toList(); t1 = g.V(t0).addE($1).to(g.V($2)).iterate(); t2 = t0.each { g.V($3).addE($4).to(it).next() }; t2",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "assoc", "$2": 1, "$3": 1, "$4": "inverse"},
		},
		{
			input:     g.V().Where(__.Or(__.Has("age", 29), __.Has("age", 30))),
			wantQuery: "g.V().where(__.or(__.has($0, $1), __.has($2, $3)))",
			wantBinds: dsl.Bindings{"$0": "age", "$1": 29, "$2": "age", "$3": 30},
		},
		{
			input:     g.V().Has("name", p.Containing("le")).Has("name", p.StartingWith("A")),
			wantQuery: `g.V().has($0, containing($1)).has($2, startingWith($3))`,
			wantBinds: dsl.Bindings{"$0": "name", "$1": "le", "$2": "name", "$3": "A"},
		},
		{
			input:     g.AddV().Property(dsl.Single, "age", 32).ValueMap(),
			wantQuery: "g.addV().property(single, $0, $1).valueMap()",
			wantBinds: dsl.Bindings{"$0": "age", "$1": 32},
		},
		{
			input:     g.V().Count(),
			wantQuery: "g.V().count()",
			wantBinds: dsl.Bindings{},
		},
		{
			input:     g.V().HasNot("age"),
			wantQuery: "g.V().hasNot($0)",
			wantBinds: dsl.Bindings{"$0": "age"},
		},
		{
			input: func() *dsl.Traversal {
				v := g.V().HasID(1)
				u := v.Clone().InE().Drop()
				return dsl.Join(v, u)
			}(),
			wantQuery: "g.V().hasId($0); g.V().hasId($1).inE().drop()",
			wantBinds: dsl.Bindings{"$0": 1, "$1": 1},
		},
		{
			input: func() *dsl.Traversal {
				v := g.V().HasID(1)
				u := v.Clone().InE().Drop()
				w := u.Clone()
				return dsl.Join(v, u, w)
			}(),
			wantQuery: "g.V().hasId($0); g.V().hasId($1).inE().drop(); g.V().hasId($2).inE().drop()",
			wantBinds: dsl.Bindings{"$0": 1, "$1": 1, "$2": 1},
		},
		{
			input:     g.V().OutE("knows").Where(__.InV().Has("name", "a8m")).OutV(),
			wantQuery: "g.V().outE($0).where(__.inV().has($1, $2)).outV()",
			wantBinds: dsl.Bindings{"$0": "knows", "$1": "name", "$2": "a8m"},
		},
		{
			input:     g.V().Has("name", p.Within("a8m", "alex")),
			wantQuery: "g.V().has($0, within($1, $2))",
			wantBinds: dsl.Bindings{"$0": "name", "$1": "a8m", "$2": "alex"},
		},
		{
			input:     g.V().HasID(p.Within(1, 2)),
			wantQuery: "g.V().hasId(within($0, $1))",
			wantBinds: dsl.Bindings{"$0": 1, "$1": 2},
		},
		{
			input:     g.V().HasID(p.Without(1, 2)),
			wantQuery: "g.V().hasId(without($0, $1))",
			wantBinds: dsl.Bindings{"$0": 1, "$1": 2},
		},
		{
			input:     g.V().Order().By("name"),
			wantQuery: "g.V().order().by($0)",
			wantBinds: dsl.Bindings{"$0": "name"},
		},
		{
			input:     g.V().Order().By("name", dsl.Incr),
			wantQuery: "g.V().order().by($0, incr)",
			wantBinds: dsl.Bindings{"$0": "name"},
		},
		{
			input:     g.V().Order().By("name", dsl.Incr).Undo(),
			wantQuery: "g.V().order()",
			wantBinds: dsl.Bindings{},
		},
		{
			input:     g.V().OutE("knows").Where(__.InV().Has("name", "a8m")).Undo(),
			wantQuery: "g.V().outE($0)",
			wantBinds: dsl.Bindings{"$0": "knows"},
		},
		{
			input:     g.V().Has("name").Group().By("name").By("age").Select(dsl.Values),
			wantQuery: "g.V().has($0).group().by($1).by($2).select(values)",
			wantBinds: dsl.Bindings{"$0": "name", "$1": "name", "$2": "age"},
		},
		{
			input:     g.V().Fold().Unfold(),
			wantQuery: "g.V().fold().unfold()",
			wantBinds: dsl.Bindings{},
		},
		{
			input: g.V().Has("person", "name", "a8m").Count().Coalesce(
				__.Is(p.NEQ(0)).Constant("unique constraint failed"),
				g.AddV("person").Property("name", "a8m").ValueMap(true),
			),
			wantQuery: "g.V().has($0, $1, $2).count().coalesce(__.is(neq($3)).constant($4), g.addV($5).property($6, $7).valueMap($8))",
			wantBinds: dsl.Bindings{"$0": "person", "$1": "name", "$2": "a8m", "$3": 0, "$4": "unique constraint failed", "$5": "person", "$6": "name", "$7": "a8m", "$8": true},
		},
		{
			input:     g.V().Has("age").Property("age", __.Union(__.Values("age"), __.Constant(10)).Sum()).ValueMap(),
			wantQuery: "g.V().has($0).property($1, __.union(__.values($2), __.constant($3)).sum()).valueMap()",
			wantBinds: dsl.Bindings{"$0": "age", "$1": "age", "$2": "age", "$3": 10},
		},
		{
			input:     g.V().Has("age").SideEffect(__.Properties("name").Drop()).ValueMap(),
			wantQuery: "g.V().has($0).sideEffect(__.properties($1).drop()).valueMap()",
			wantBinds: dsl.Bindings{"$0": "age", "$1": "name"},
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			query, bindings := tt.input.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantBinds, bindings)
		})
	}
}
