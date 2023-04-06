// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package printer

import (
	"strings"
	"testing"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"

	"github.com/stretchr/testify/assert"
)

func TestPrinter_Print(t *testing.T) {
	tests := []struct {
		input *gen.Graph
		out   string
	}{
		{
			input: &gen.Graph{
				Nodes: []*gen.Type{
					{
						Name: "User",
						ID:   &gen.Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
						Fields: []*gen.Field{
							{Name: "name", Type: &field.TypeInfo{Type: field.TypeString}, Validators: 1},
							{Name: "age", Type: &field.TypeInfo{Type: field.TypeInt}, Nillable: true},
							{Name: "created_at", Type: &field.TypeInfo{Type: field.TypeTime}, Nillable: true, Immutable: true},
						},
					},
				},
			},
			out: `
User:
	+------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	|   Field    |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators | Comment |
	+------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| id         | int       | false  | false    | false    | false   | false         | false     |           |          0 |         |
	| name       | string    | false  | false    | false    | false   | false         | false     |           |          1 |         |
	| age        | int       | false  | false    | true     | false   | false         | false     |           |          0 |         |
	| created_at | time.Time | false  | false    | true     | false   | false         | true      |           |          0 |         |
	+------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	
`,
		},
		{
			input: &gen.Graph{
				Nodes: []*gen.Type{
					{
						Name: "User",
						ID:   &gen.Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
						Edges: []*gen.Edge{
							{Name: "groups", Type: &gen.Type{Name: "Group"}, Rel: gen.Relation{Type: gen.M2M}, Optional: true},
							{Name: "spouse", Type: &gen.Type{Name: "User"}, Unique: true, Rel: gen.Relation{Type: gen.O2O}},
						},
					},
				},
			},
			out: `
User:
	+-------+------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| Field | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators | Comment |
	+-------+------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| id    | int  | false  | false    | false    | false   | false         | false     |           |          0 |         |
	+-------+------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	+--------+-------+---------+---------+----------+--------+----------+---------+
	|  Edge  | Type  | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------+-------+---------+---------+----------+--------+----------+---------+
	| groups | Group | false   |         | M2M      | false  | true     |         |
	| spouse | User  | false   |         | O2O      | true   | false    |         |
	+--------+-------+---------+---------+----------+--------+----------+---------+
	
`,
		},
		{
			input: &gen.Graph{
				Nodes: []*gen.Type{
					{
						Name: "User",
						ID:   &gen.Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
						Fields: []*gen.Field{
							{Name: "name", Type: &field.TypeInfo{Type: field.TypeString}, Validators: 1},
							{Name: "age", Type: &field.TypeInfo{Type: field.TypeInt}, Nillable: true},
						},
						Edges: []*gen.Edge{
							{Name: "groups", Type: &gen.Type{Name: "Group"}, Rel: gen.Relation{Type: gen.M2M}, Optional: true},
							{Name: "spouse", Type: &gen.Type{Name: "User"}, Unique: true, Rel: gen.Relation{Type: gen.O2O}},
						},
					},
				},
			},
			out: `
User:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     |           |          0 |         |
	| name  | string | false  | false    | false    | false   | false         | false     |           |          1 |         |
	| age   | int    | false  | false    | true     | false   | false         | false     |           |          0 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	+--------+-------+---------+---------+----------+--------+----------+---------+
	|  Edge  | Type  | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------+-------+---------+---------+----------+--------+----------+---------+
	| groups | Group | false   |         | M2M      | false  | true     |         |
	| spouse | User  | false   |         | O2O      | true   | false    |         |
	+--------+-------+---------+---------+----------+--------+----------+---------+
	
`,
		},
		{
			input: &gen.Graph{
				Nodes: []*gen.Type{
					{
						Name: "User",
						ID:   &gen.Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
						Fields: []*gen.Field{
							{Name: "name", Type: &field.TypeInfo{Type: field.TypeString}, Validators: 1},
							{Name: "age", Type: &field.TypeInfo{Type: field.TypeInt}, Nillable: true},
						},
						Edges: []*gen.Edge{
							{Name: "groups", Type: &gen.Type{Name: "Group"}, Rel: gen.Relation{Type: gen.M2M}, Optional: true},
							{Name: "spouse", Type: &gen.Type{Name: "User"}, Unique: true, Rel: gen.Relation{Type: gen.O2O}},
						},
					},
					{
						Name: "Group",
						ID:   &gen.Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
						Fields: []*gen.Field{
							{Name: "name", Type: &field.TypeInfo{Type: field.TypeString}},
						},
						Edges: []*gen.Edge{
							{Name: "users", Type: &gen.Type{Name: "User"}, Rel: gen.Relation{Type: gen.M2M}, Optional: true},
						},
					},
				},
			},
			out: `
User:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     |           |          0 |         |
	| name  | string | false  | false    | false    | false   | false         | false     |           |          1 |         |
	| age   | int    | false  | false    | true     | false   | false         | false     |           |          0 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	+--------+-------+---------+---------+----------+--------+----------+---------+
	|  Edge  | Type  | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------+-------+---------+---------+----------+--------+----------+---------+
	| groups | Group | false   |         | M2M      | false  | true     |         |
	| spouse | User  | false   |         | O2O      | true   | false    |         |
	+--------+-------+---------+---------+----------+--------+----------+---------+
	
Group:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     |           |          0 |         |
	| name  | string | false  | false    | false    | false   | false         | false     |           |          0 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+---------+
	+-------+------+---------+---------+----------+--------+----------+---------+
	| Edge  | Type | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-------+------+---------+---------+----------+--------+----------+---------+
	| users | User | false   |         | M2M      | false  | true     |         |
	+-------+------+---------+---------+----------+--------+----------+---------+
	
`,
		},
	}
	for _, tt := range tests {
		b := &strings.Builder{}
		Fprint(b, tt.input)
		assert.Equal(t, tt.out, "\n"+b.String())
	}
}
