// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"testing"

	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"

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
	require.Equal("_m", typ.Receiver())

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Fields: []*load.Field{
			{Name: "foo", Unique: true, Default: true, Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.EqualError(err, "unique field \"foo\" cannot have default value", "unique field can not have default")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Fields: []*load.Field{
			{Name: "foo", Sensitive: true, Tag: `yaml:"pwd"`, Info: &field.TypeInfo{Type: field.TypeString}},
		},
	})
	require.EqualError(err, "sensitive field \"foo\" cannot have struct tags", "sensitive field cannot have tags")

	typ, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Fields: []*load.Field{
			{Name: "id", Info: &field.TypeInfo{Type: field.TypeString}, Annotations: dict("EntSQL", dict("collation", "utf8_ci_bin"))},
		},
	})
	require.NoError(err)
	require.NotNil(typ)
	require.NotNil(t, typ.ID)
	pkCol := typ.ID.PK()
	require.NotNil(pkCol)
	require.Equal("utf8_ci_bin", pkCol.Collation)

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "foo", Unique: true, Info: &field.TypeInfo{Type: field.TypeInt}},
			{Name: "foo", Unique: true, Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.EqualError(err, "field \"foo\" redeclared for type \"T\"", "field foo redeclared")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "enums", Info: &field.TypeInfo{Type: field.TypeEnum}, Enums: []struct{ N, V string }{{V: "v"}, {V: "v"}}},
		},
	})
	require.EqualError(err, "duplicate values \"v\" for enum field \"enums\"", "duplicate enums")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "enums", Info: &field.TypeInfo{Type: field.TypeEnum}, Enums: []struct{ N, V string }{{}}},
		},
	})
	require.EqualError(err, "\"enums\" field value cannot be empty", "empty value for enums")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "", Info: &field.TypeInfo{Type: field.TypeInt}},
		},
	})
	require.EqualError(err, "field name cannot be empty", "empty field name")

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "id", Info: &field.TypeInfo{Type: field.TypeInt}, Optional: true},
		},
	})
	require.EqualError(err, "id field cannot be optional", "id field cannot be optional")

	typ, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "T",
		Fields: []*load.Field{
			{Name: "id", Info: &field.TypeInfo{Type: field.TypeString}, ValueScanner: true},
		},
	})
	require.NoError(err)
	require.True(typ.HasValueScanner())

	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{Name: "Type"})
	require.EqualError(err, "schema lowercase name conflicts with Go keyword \"type\"")
	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{Name: "Int"})
	require.EqualError(err, "schema lowercase name conflicts with Go predeclared identifier \"int\"")
	_, err = NewType(&Config{Package: "entc/gen"}, &load.Schema{Name: "Value"})
	require.EqualError(err, "schema name conflicts with ent predeclared identifier \"Value\"")
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
		{"UserIDs", "user_ids"},
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

func TestField_EnumName(t *testing.T) {
	tests := []struct {
		name string
		enum string
	}{
		{"GIF", "TypeGIF"},
		{"SVG", "TypeSVG"},
		{"PNG", "TypePNG"},
		{"MP4", "TypeMP4"},
		{"unknown", "TypeUnknown"},
		{"user_data", "TypeUserData"},
		{"test user", "TypeTestUser"},
	}
	for _, tt := range tests {
		require.Equal(t, tt.enum, Field{Name: "Type"}.EnumName(tt.name))
	}
}

func TestType_WithRuntimeMixin(t *testing.T) {
	position := &load.Position{MixedIn: true}
	typ := &Type{
		ID: &Field{},
		Fields: []*Field{
			{Default: true, Position: position},
			{UpdateDefault: true, Position: position},
			{Validators: 1, Position: position},
		},
	}
	require.True(t, typ.RuntimeMixin())
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

	err = typ.AddIndex(&load.Index{Unique: true, Fields: []string{"id"}})
	require.NoError(t, err, "valid index for ID field")

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

func TestField_incremental(t *testing.T) {
	tests := []struct {
		annotations map[string]any
		def         bool
		expected    bool
	}{
		{dict("EntSQL", nil), false, false},
		{dict("EntSQL", nil), true, true},
		{dict("EntSQL", dict("incremental", true)), false, true},
		{dict("EntSQL", dict("incremental", false)), true, false},
	}
	for _, tt := range tests {
		typ := &Field{Annotations: tt.annotations}
		require.Equal(t, tt.expected, typ.incremental(tt.def))
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
		{"SSOCert", "_SSOCert"},
		{"driver", "_driver"},
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

func TestValidSchemaName(t *testing.T) {
	err := ValidSchemaName("Config")
	require.Error(t, err)
	err = ValidSchemaName("Mutation")
	require.Error(t, err)
	err = ValidSchemaName("Boring")
	require.NoError(t, err)
	err = ValidSchemaName("Order")
	require.NoError(t, err)
}

func TestField_Blob(t *testing.T) {
	require := require.New(t)

	// Test creating a type with blob-stored fields.
	typ, err := NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "Document",
		Fields: []*load.Field{
			{
				Name:     "content",
				Info:     &field.TypeInfo{Type: field.TypeBlob},
				Optional: true,
				Comment:  "blob content",
			},
			{
				Name: "thumbnail",
				Info: &field.TypeInfo{Type: field.TypeBlob},
			},
			{
				Name: "title",
				Info: &field.TypeInfo{Type: field.TypeString},
			},
		},
	})
	require.NoError(err)
	require.NotNil(typ)
	require.Equal("Document", typ.Name)

	// Find blob fields.
	require.True(typ.HasBlobFields())
	blobFields := typ.BlobFields()
	require.Len(blobFields, 2)

	// First blob field: content.
	f0 := blobFields[0]
	require.Equal("content", f0.Name)
	require.True(f0.IsBlob())
	require.True(f0.Optional)

	// Second blob field: thumbnail.
	f1 := blobFields[1]
	require.Equal("thumbnail", f1.Name)
	require.True(f1.IsBlob())

	// Non-blob field should not be blob-stored.
	titleField := typ.Fields[2]
	require.Equal("title", titleField.Name)
	require.False(titleField.IsBlob())

	// MutationFields should exclude only lazy blob fields.
	mutFields := typ.MutationFields()
	for _, mf := range mutFields {
		require.False(mf.IsBlobLazy(), "lazy blob field %s should not be in MutationFields", mf.Name)
	}
	// Non-lazy blob fields (content, thumbnail) should be in MutationFields
	// because their mutation struct fields are used by the blob hook.
	var blobInMut int
	for _, mf := range mutFields {
		if mf.IsBlob() {
			blobInMut++
		}
	}
	require.Equal(2, blobInMut)

	// Type without blob fields.
	typ2, err := NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "Simple",
		Fields: []*load.Field{
			{Name: "name", Info: &field.TypeInfo{Type: field.TypeString}},
		},
	})
	require.NoError(err)
	require.False(typ2.HasBlobFields())
	require.Empty(typ2.BlobFields())
}

func TestField_BlobScanType(t *testing.T) {
	require := require.New(t)

	typ, err := NewType(&Config{Package: "entc/gen"}, &load.Schema{
		Name: "Doc",
		Fields: []*load.Field{
			{
				Name: "data",
				Info: &field.TypeInfo{Type: field.TypeBlob},
			},
		},
	})
	require.NoError(err)
	f := typ.Fields[0]
	require.True(f.IsBlob())
	require.True(f.IsBlobNoColumn())
	require.False(f.IsBlobLazy())
	// Non-lazy blob fields appear in MutationFields (for blob hook usage),
	// but are excluded from SQL columns by IsBlobNoColumn.
	require.NotEmpty(typ.MutationFields())
}
