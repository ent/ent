package gen

import (
	"strings"
	"testing"

	"fbc/ent/field"

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

func TestField(t *testing.T) {
	f := &Field{Type: field.TypeTime}
	require.True(t, f.IsTime())
	require.Equal(t, "time.Now()", f.ExampleCode())

	require.Equal(t, "1", Field{Type: field.TypeInt}.ExampleCode())
	require.Equal(t, "true", Field{Type: field.TypeBool}.ExampleCode())
	require.Equal(t, "1", Field{Type: field.TypeFloat64}.ExampleCode())
	require.Equal(t, "\"string\"", Field{Type: field.TypeString}.ExampleCode())
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

func TestField_DefaultConstant(t *testing.T) {
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
		require.Equal(t, tt.constant, typ.DefaultConstant())
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
				ID:   &Field{Name: "id", Type: field.TypeInt},
				Fields: []*Field{
					{Name: "name", Type: field.TypeString, Validators: 1},
					{Name: "age", Type: field.TypeInt, Nullable: true},
				},
			},
			out: `
User:
	+-------+--------+--------+----------+----------+------------+-----------+------------+
	| Field |  Type  | Unique | Optional | Nullable | HasDefault | StructTag | Validators |
	+-------+--------+--------+----------+----------+------------+-----------+------------+
	| id    | int    | false  | false    | false    | false      |           |          0 |
	| name  | string | false  | false    | false    | false      |           |          1 |
	| age   | int    | false  | false    | true     | false      |           |          0 |
	+-------+--------+--------+----------+----------+------------+-----------+------------+
	
`,
		},
		{
			typ: &Type{
				Name: "User",
				ID:   &Field{Name: "id", Type: field.TypeInt},
				Edges: []*Edge{
					{Name: "groups", Type: &Type{Name: "Group"}, Rel: Relation{Type: M2M}, Optional: true},
					{Name: "spouse", Type: &Type{Name: "User"}, Unique: true, Rel: Relation{Type: O2O}},
				},
			},
			out: `
User:
	+-------+------+--------+----------+----------+------------+-----------+------------+
	| Field | Type | Unique | Optional | Nullable | HasDefault | StructTag | Validators |
	+-------+------+--------+----------+----------+------------+-----------+------------+
	| id    | int  | false  | false    | false    | false      |           |          0 |
	+-------+------+--------+----------+----------+------------+-----------+------------+
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
				ID:   &Field{Name: "id", Type: field.TypeInt},
				Fields: []*Field{
					{Name: "name", Type: field.TypeString, Validators: 1},
					{Name: "age", Type: field.TypeInt, Nullable: true},
				},
				Edges: []*Edge{
					{Name: "groups", Type: &Type{Name: "Group"}, Rel: Relation{Type: M2M}, Optional: true},
					{Name: "spouse", Type: &Type{Name: "User"}, Unique: true, Rel: Relation{Type: O2O}},
				},
			},
			out: `
User:
	+-------+--------+--------+----------+----------+------------+-----------+------------+
	| Field |  Type  | Unique | Optional | Nullable | HasDefault | StructTag | Validators |
	+-------+--------+--------+----------+----------+------------+-----------+------------+
	| id    | int    | false  | false    | false    | false      |           |          0 |
	| name  | string | false  | false    | false    | false      |           |          1 |
	| age   | int    | false  | false    | true     | false      |           |          0 |
	+-------+--------+--------+----------+----------+------------+-----------+------------+
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
