// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"testing"

	"github.com/facebookincubator/ent/entc/load"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	require := require.New(t)
	typ, err := NewType(&Config{Package: "entc/gen"}, T1)
	require.NoError(err)
	require.NotNil(typ)
	require.Equal("T1", typ.Name)
	require.Equal("t1", typ.Label())
	require.Equal("t1", typ.Package())
	require.Equal("t", typ.Receiver())

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Fields: []*load.Field{
			{Unique: true, Default: true, Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.Error(err, "unique field can not have default")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Fields: []*load.Field{
			{Sensitive: true, Tag: `yaml:"pwd"`, Info: &field.TypeInfo{Type: field.TypeString}},
		},
	})
	require.Error(err, "sensitive field cannot have tags")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "foo", Unique: true, Info: &field.TypeInfo{Type: field.TypeInt}},
			{Name: "foo", Unique: true, Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.Error(err, "field foo redeclared")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "enums", Info: &field.TypeInfo{Type: field.TypeEnum}, Enums: []string{"A", "A"}},
		},
	})
	require.Error(err, "duplicate enums")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "enums", Info: &field.TypeInfo{Type: field.TypeEnum}, Enums: []string{""}},
		},
	})
	require.Error(err, "empty value for enums")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "", Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.Error(err, "empty field name")
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
		{"[16]byte", "b"},
	}
	for _, tt := range tests {
		typ := &Type{Name: tt.name, Config: &Config{Package: "entc/gen"}}
		require.Equal(t, tt.receiver, typ.Receiver())
	}
}

func TestType_MixedInWithDefaultOrValidator(t *testing.T) {
	position := &load.Position{
		MixedIn: true,
	}
	typ := &Type{
		Fields: []*Field{
			{Default: true, Position: position},
			{UpdateDefault: true, Position: position},
			{Validators: 1, Position: position},
		},
	}
	fields := typ.MixedInWithDefaultOrValidator()
	require.Equal(t, 3, len(fields))
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
	typ, err := NewType(&Config{}, &load.Schema{
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
	require.Error(t, err, "missing fields or edges")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"unknown"}})
	require.Error(t, err, "unknown field for index")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"text"}})
	require.Error(t, err, "index size exceeded")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"parent"}})
	require.Error(t, err, "missing edge")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"next"}})
	require.Error(t, err, "not an inverse edge for O2O relation")

	err = typ.AddIndex(&load.Index{Unique: true, Edges: []string{"prev", "owner"}})
	require.NoError(t, err, "valid index defined only on edges")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"prev"}})
	require.NoError(t, err, "valid index on O2O relation and field")

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"name"}, Edges: []string{"owner"}})
	require.NoError(t, err, "valid index on M2O relation and field")
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

func TestBuilderField(t *testing.T) {
	tests := []struct {
		name  string
		field string
	}{
		{"active", "active"},
		{"type", "_type"},
		{"config", "_config"},
	}
	for _, tt := range tests {
		require.Equal(t, tt.field, Edge{Name: tt.name}.BuilderField())
		require.Equal(t, tt.field, Field{Name: tt.name}.BuilderField())
	}
}

func TestEdge(t *testing.T) {
	u, g := &Type{Name: "User"}, &Type{Name: "Group"}
	groups := &Edge{Name: "groups", Type: g, Owner: u, Rel: Relation{Type: M2M}}
	users := &Edge{Name: "users", Inverse: "groups", Type: u, Owner: u, Rel: Relation{Type: M2M}}

	require.True(t, users.IsInverse())
	require.False(t, groups.IsInverse())

	require.Equal(t, "GroupsLabel", users.LabelConstant())
	require.Equal(t, "GroupsLabel", groups.LabelConstant())

	require.Equal(t, "UsersInverseLabel", users.InverseLabelConstant())
	require.Equal(t, "user_groups", users.Label())
	require.Equal(t, "user_groups", groups.Label())
}
