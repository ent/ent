package sqlgraph

import (
	"strconv"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entql"
	"github.com/facebook/ent/schema/field"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGraph_AddE(t *testing.T) {
	g := &Graph{
		Nodes: []*Node{{Type: "user"}, {Type: "pet"}},
	}
	err := g.AddE("pets", &EdgeSpec{Rel: O2M}, "user", "pet")
	assert.NoError(t, err)
	err = g.AddE("owner", &EdgeSpec{Rel: O2M}, "pet", "user")
	assert.NoError(t, err)
	err = g.AddE("groups", &EdgeSpec{Rel: M2M}, "pet", "groups")
	assert.Error(t, err)
}

func TestGraph_EvalP(t *testing.T) {
	g := &Graph{
		Nodes: []*Node{
			{
				Type: "user",
				NodeSpec: NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "uid"},
				},
				Fields: map[string]*FieldSpec{
					"name": {Column: "name", Type: field.TypeString},
					"last": {Column: "last", Type: field.TypeString},
				},
			},
			{
				Type: "pet",
				NodeSpec: NodeSpec{
					Table: "pets",
					ID:    &FieldSpec{Column: "pid"},
				},
				Fields: map[string]*FieldSpec{
					"name": {Column: "name", Type: field.TypeString},
				},
			},
			{
				Type: "group",
				NodeSpec: NodeSpec{
					Table: "groups",
					ID:    &FieldSpec{Column: "gid"},
				},
				Fields: map[string]*FieldSpec{
					"name": {Column: "name", Type: field.TypeString},
				},
			},
		},
	}
	err := g.AddE("pets", &EdgeSpec{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}}, "user", "pet")
	require.NoError(t, err)
	err = g.AddE("owner", &EdgeSpec{Rel: M2O, Inverse: true, Table: "pets", Columns: []string{"owner_id"}}, "pet", "user")
	require.NoError(t, err)
	err = g.AddE("groups", &EdgeSpec{Rel: M2M, Table: "user_groups", Columns: []string{"user_id", "group_id"}}, "user", "group")
	require.NoError(t, err)
	err = g.AddE("users", &EdgeSpec{Rel: M2M, Inverse: true, Table: "user_groups", Columns: []string{"user_id", "group_id"}}, "group", "user")
	require.NoError(t, err)

	tests := []struct {
		s         *sql.Selector
		p         entql.P
		wantQuery string
		wantArgs  []interface{}
		wantErr   bool
	}{
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")),
			p:         entql.FieldHasPrefix("name", "a"),
			wantQuery: `SELECT * FROM "users" WHERE "name" LIKE $1`,
			wantArgs:  []interface{}{"a%"},
		},
		{
			s: sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")).
				Where(sql.EQ("age", 1)),
			p:         entql.FieldHasPrefix("name", "a"),
			wantQuery: `SELECT * FROM "users" WHERE "age" = $1 AND "name" LIKE $2`,
			wantArgs:  []interface{}{1, "a%"},
		},
		{
			s: sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")).
				Where(sql.EQ("age", 1)),
			p:         entql.FieldHasPrefix("name", "a"),
			wantQuery: `SELECT * FROM "users" WHERE "age" = $1 AND "name" LIKE $2`,
			wantArgs:  []interface{}{1, "a%"},
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")),
			p:         entql.EQ(entql.F("name"), entql.F("last")),
			wantQuery: `SELECT * FROM "users" WHERE "name" = "last"`,
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")),
			p:         entql.EQ(entql.F("name"), entql.F("last")),
			wantQuery: `SELECT * FROM "users" WHERE "name" = "last"`,
		},
		{
			s: sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")).
				Where(sql.EQ("foo", "bar")),
			p:         entql.Or(entql.FieldEQ("name", "foo"), entql.FieldEQ("name", "baz")),
			wantQuery: `SELECT * FROM "users" WHERE "foo" = $1 AND ("name" = $2 OR "name" = $3)`,
			wantArgs:  []interface{}{"bar", "foo", "baz"},
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")),
			p:         entql.HasEdge("pets"),
			wantQuery: `SELECT * FROM "users" WHERE "users"."uid" IN (SELECT "pets"."owner_id" FROM "pets" WHERE "pets"."owner_id" IS NOT NULL)`,
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")),
			p:         entql.HasEdge("groups"),
			wantQuery: `SELECT * FROM "users" WHERE "users"."uid" IN (SELECT "user_groups"."user_id" FROM "user_groups")`,
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")),
			p:         entql.HasEdgeWith("pets", entql.Or(entql.FieldEQ("name", "pedro"), entql.FieldEQ("name", "xabi"))),
			wantQuery: `SELECT * FROM "users" WHERE "users"."uid" IN (SELECT "pets"."owner_id" FROM "pets" WHERE "name" = $1 OR "name" = $2)`,
			wantArgs:  []interface{}{"pedro", "xabi"},
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")).Where(sql.EQ("active", true)),
			p:         entql.HasEdgeWith("groups", entql.Or(entql.FieldEQ("name", "GitHub"), entql.FieldEQ("name", "GitLab"))),
			wantQuery: `SELECT * FROM "users" WHERE "active" = $1 AND "users"."uid" IN (SELECT "user_groups"."user_id" FROM "user_groups" JOIN "groups" AS "t0" ON "user_groups"."group_id" = "t0"."gid" WHERE "name" = $2 OR "name" = $3)`,
			wantArgs:  []interface{}{true, "GitHub", "GitLab"},
		},
		{
			s:         sql.Dialect(dialect.Postgres).Select().From(sql.Table("users")).Where(sql.EQ("active", true)),
			p:         entql.And(entql.HasEdge("pets"), entql.HasEdge("groups"), entql.EQ(entql.F("name"), entql.F("uid"))),
			wantQuery: `SELECT * FROM "users" WHERE "active" = $1 AND ("users"."uid" IN (SELECT "pets"."owner_id" FROM "pets" WHERE "pets"."owner_id" IS NOT NULL) AND "users"."uid" IN (SELECT "user_groups"."user_id" FROM "user_groups") AND "name" = "uid")`,
			wantArgs:  []interface{}{true},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err = g.EvalP("user", tt.p, tt.s)
			require.Equal(t, tt.wantErr, err != nil, err)
			query, args := tt.s.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}
