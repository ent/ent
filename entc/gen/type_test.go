// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"strings"
	"testing"

	"github.com/facebookincubator/ent/entc/load"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	require := require.New(t)
	typ, err := NewType(Config{Package: "entc/gen"}, T1)
	require.NoError(err)
	require.NotNil(typ)
	require.Equal("T1", typ.Name)
	require.Equal("t1", typ.Label())
	require.Equal("t1", typ.Package())
	require.Equal("t", typ.Receiver())

	typ, err = NewType(Config{Package: "entc/gen"}, &load.Schema{
		Fields: []*load.Field{
			{Unique: true, Default: true, Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.Error(err, "unique field can not have default")
	require.Nil(typ)

	typ, err = NewType(Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "foo", Unique: true, Info: &field.TypeInfo{Type: field.TypeInt}},
			{Name: "foo", Unique: true, Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.Error(err, "field foo redeclared")
	require.Nil(typ)

	typ, err = NewType(Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "enums", Info: &field.TypeInfo{Type: field.TypeEnum}, Enums: []string{"A", "A"}},
		},
	})
	require.Error(err, "duplicate enums")
	require.Nil(typ)

	typ, err = NewType(Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "enums", Info: &field.TypeInfo{Type: field.TypeEnum}, Enums: []string{""}},
		},
	})
	require.Error(err, "empty value for enums")
	require.Nil(typ)
}

func TestType_Label(t *testing.T) {
	tests := []struct {
		name  string
		label string
	}{
		{"User", "user"},
		{"UserInfo", "user_info"},
		{"PHBOrg", "phb_org"},
		{"UserID", "user_id"},
		{"HTTPCode", "http_code"},
	}
	for _, tt := range tests {
		typ := &Type{Name: tt.name}
		require.Equal(t, tt.label, typ.Label())
	}
}

func TestType_Table(t *testing.T) {
	tests := []struct {
		name  string
		label string
	}{
		{"User", "users"},
		{"Device", "devices"},
		{"UserInfo", "user_infos"},
		{"PHBOrg", "phb_orgs"},
		{"HTTPCode", "http_codes"},
	}
	for _, tt := range tests {
		typ := &Type{Name: tt.name}
		require.Equal(t, tt.label, typ.Table())
	}
}

func TestType_Receiver(t *testing.T) {
	tests := []struct {
		name     string
		receiver string
	}{
		{"User", "u"},
		{"Group", "gr"},
		{"UserData", "ud"},
		{"UserInfo", "ui"},
		{"User_Info", "ui"},
		{"PHBUser", "pu"},
		{"PHBOrg", "po"},
		{"DomainSpecificLang", "dospla"},
		{"[]byte", "b"},
	}
	for _, tt := range tests {
		typ := &Type{Name: tt.name, Config: Config{Package: "entc/gen"}}
		require.Equal(t, tt.receiver, typ.Receiver())
	}
}

func TestType_TagTypes(t *testing.T) {
	typ := &Type{
		Fields: []*Field{
			{StructTag: `json:"age"`},
			{StructTag: `json:"name,omitempty`},
			{StructTag: `json:"name,omitempty" sql:"nothing"`},
			{StructTag: `sql:"nothing" yaml:"ignore"`},
			{StructTag: `sql:"nothing" yaml:"ignore"`},
			{StructTag: `invalid`},
			{StructTag: `"invalid"`},
		},
	}
	tags := typ.TagTypes()
	require.Equal(t, []string{"json", "sql", "yaml"}, tags)
}

func TestType_Package(t *testing.T) {
	tests := []struct {
		name string
		pkg  string
	}{
		{"User", "user"},
		{"UserInfo", "userinfo"},
		{"PHBOrg", "phborg"},
		{"UserID", "userid"},
		{"HTTPCode", "httpcode"},
	}
	for _, tt := range tests {
		typ := &Type{Name: tt.name}
		require.Equal(t, tt.pkg, typ.Package())
	}
}

func TestType_AddIndex(t *testing.T) {
	size := int64(1024)
	typ, err := NewType(Config{}, &load.Schema{
		Name: "User",
		Fields: []*load.Field{
			{Name: "name", Info: &field.TypeInfo{Type: field.TypeString}},
			{Name: "text", Info: &field.TypeInfo{Type: field.TypeString}, Size: &size},
		},
	})
	require.NoError(t, err)
	typ.Edges = append(typ.Edges,
		&Edge{Name: "next", Rel: Relation{Type: O2O, Columns: []string{"prev_id"}}},
		&Edge{Name: "prev", Inverse: "next", Rel: Relation{Type: O2O, Columns: []string{"prev_id"}}},
		&Edge{Name: "owner", Inverse: "files", Rel: Relation{Type: M2O, Columns: []string{"file_id"}}},
	)

	err = typ.AddIndex(&load.Index{Unique: true})
	require.Error(t, err, "missing fields")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"unknown"}})
	require.Error(t, err, "unknown field for index")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"text"}})
	require.Error(t, err, "index size exceeded")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"parent"}})
	require.Error(t, err, "missing edge")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"next"}})
	require.Error(t, err, "not an inverse edge for O2O relation")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"prev"}})
	require.NoError(t, err, "valid index on O2O relation and field")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"owner"}})
	require.NoError(t, err, "valid index on M2O relation and field")
}

func TestField(t *testing.T) {
	f := &Field{Type: &field.TypeInfo{Type: field.TypeTime}}
	require.True(t, f.IsTime())
	require.Equal(t, "time.Now()", f.ExampleCode())

	require.Equal(t, "1", Field{Type: &field.TypeInfo{Type: field.TypeInt}}.ExampleCode())
	require.Equal(t, "true", Field{Type: &field.TypeInfo{Type: field.TypeBool}}.ExampleCode())
	require.Equal(t, "1", Field{Type: &field.TypeInfo{Type: field.TypeFloat64}}.ExampleCode())
	require.Equal(t, "\"string\"", Field{Type: &field.TypeInfo{Type: field.TypeString}}.ExampleCode())
}

func TestField_Constant(t *testing.T) {
	tests := []struct {
		name     string
		constant string
	}{
		{"user", "FieldUser"},
		{"user_id", "FieldUserID"},
		{"user_name", "FieldUserName"},
	}
	for _, tt := range tests {
		typ := &Field{Name: tt.name}
		require.Equal(t, tt.constant, typ.Constant())
	}
}

func TestField_DefaultName(t *testing.T) {
	tests := []struct {
		name     string
		constant string
	}{
		{"active", "DefaultActive"},
		{"expired_at", "DefaultExpiredAt"},
		{"group_name", "DefaultGroupName"},
	}
	for _, tt := range tests {
		typ := &Field{Name: tt.name}
		require.Equal(t, tt.constant, typ.DefaultName())
	}
}

func TestEdge(t *testing.T) {
	u, g := &Type{Name: "User"}, &Type{Name: "Group"}
	groups := &Edge{Name: "groups", Type: g, Owner: u, Rel: Relation{Type: M2M}}
	users := &Edge{Name: "users", Inverse: "groups", Type: u, Owner: u, Rel: Relation{Type: M2M}}

	require.True(t, users.IsInverse())
	require.False(t, groups.IsInverse())

	require.Equal(t, "GroupsLabel", users.Constant())
	require.Equal(t, "GroupsLabel", groups.Constant())

	require.Equal(t, "UsersInverseLabel", users.InverseConstant())
	require.Equal(t, "user_groups", users.Label())
	require.Equal(t, "user_groups", groups.Label())
}

func TestType_Describe(t *testing.T) {
	tests := []struct {
		typ *Type
		out string
	}{
		{
			typ: &Type{
				Name: "User",
				ID:   &Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
				Fields: []*Field{
					{Name: "name", Type: &field.TypeInfo{Type: field.TypeString}, Validators: 1},
					{Name: "age", Type: &field.TypeInfo{Type: field.TypeInt}, Nillable: true},
					{Name: "created_at", Type: &field.TypeInfo{Type: field.TypeTime}, Nillable: true, Immutable: true},
				},
			},
			out: `
User:
	+------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	|   Field    |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators |
	+------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	| id         | int       | false  | false    | false    | false   | false         | false     |           |          0 |
	| name       | string    | false  | false    | false    | false   | false         | false     |           |          1 |
	| age        | int       | false  | false    | true     | false   | false         | false     |           |          0 |
	| created_at | time.Time | false  | false    | true     | false   | false         | true      |           |          0 |
	+------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	
`,
		},
		{
			typ: &Type{
				Name: "User",
				ID:   &Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
				Edges: []*Edge{
					{Name: "groups", Type: &Type{Name: "Group"}, Rel: Relation{Type: M2M}, Optional: true},
					{Name: "spouse", Type: &Type{Name: "User"}, Unique: true, Rel: Relation{Type: O2O}},
				},
			},
			out: `
User:
	+-------+------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	| Field | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators |
	+-------+------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	| id    | int  | false  | false    | false    | false   | false         | false     |           |          0 |
	+-------+------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	+--------+-------+---------+---------+----------+--------+----------+
	|  Edge  | Type  | Inverse | BackRef | Relation | Unique | Optional |
	+--------+-------+---------+---------+----------+--------+----------+
	| groups | Group | false   |         | M2M      | false  | true     |
	| spouse | User  | false   |         | O2O      | true   | false    |
	+--------+-------+---------+---------+----------+--------+----------+
	
`,
		},
		{
			typ: &Type{
				Name: "User",
				ID:   &Field{Name: "id", Type: &field.TypeInfo{Type: field.TypeInt}},
				Fields: []*Field{
					{Name: "name", Type: &field.TypeInfo{Type: field.TypeString}, Validators: 1},
					{Name: "age", Type: &field.TypeInfo{Type: field.TypeInt}, Nillable: true},
				},
				Edges: []*Edge{
					{Name: "groups", Type: &Type{Name: "Group"}, Rel: Relation{Type: M2M}, Optional: true},
					{Name: "spouse", Type: &Type{Name: "User"}, Unique: true, Rel: Relation{Type: O2O}},
				},
			},
			out: `
User:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable | StructTag | Validators |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	| id    | int    | false  | false    | false    | false   | false         | false     |           |          0 |
	| name  | string | false  | false    | false    | false   | false         | false     |           |          1 |
	| age   | int    | false  | false    | true     | false   | false         | false     |           |          0 |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------+------------+
	+--------+-------+---------+---------+----------+--------+----------+
	|  Edge  | Type  | Inverse | BackRef | Relation | Unique | Optional |
	+--------+-------+---------+---------+----------+--------+----------+
	| groups | Group | false   |         | M2M      | false  | true     |
	| spouse | User  | false   |         | O2O      | true   | false    |
	+--------+-------+---------+---------+----------+--------+----------+
	
`,
		},
	}
	for _, tt := range tests {
		b := &strings.Builder{}
		tt.typ.Describe(b)
		assert.Equal(t, tt.out, "\n"+b.String())
	}
}
