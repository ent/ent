// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	fd := field.Int("age").
		Positive().
		Comment("comment").
		Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeInt, fd.Info.Type)
	assert.Len(t, fd.Validators, 1)
	assert.Equal(t, "comment", fd.Comment)

	fd = field.Int("age").
		Default(10).
		Min(10).
		Max(20).
		Descriptor()
	assert.NotNil(t, fd.Default)
	assert.Equal(t, 10, fd.Default)
	assert.Len(t, fd.Validators, 2)

	fd = field.Int("age").
		Range(20, 40).
		Nillable().
		SchemaType(map[string]string{
			dialect.SQLite:   "numeric",
			dialect.Postgres: "int_type",
		}).
		Descriptor()
	assert.Nil(t, fd.Default)
	assert.True(t, fd.Nillable)
	assert.False(t, fd.Immutable)
	assert.Len(t, fd.Validators, 1)
	assert.Equal(t, "numeric", fd.SchemaType[dialect.SQLite])
	assert.Equal(t, "int_type", fd.SchemaType[dialect.Postgres])

	assert.Equal(t, field.TypeInt8, field.Int8("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeInt16, field.Int16("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeInt32, field.Int32("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeInt64, field.Int64("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeUint, field.Uint("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeUint8, field.Uint8("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeUint16, field.Uint16("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeUint32, field.Uint32("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeUint64, field.Uint64("age").Descriptor().Info.Type)

	type Count int
	fd = field.Int("active").GoType(Count(0)).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "field_test.Count", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Count", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Int("count").GoType(&sql.NullInt64{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "*sql.NullInt64", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "*sql.NullInt64", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Int("count").GoType(false).Descriptor()
	assert.EqualError(t, fd.Err, `GoType must be a "int" type or ValueScanner`)
	fd = field.Int("count").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Int("count").GoType(new(Count)).Descriptor()
	assert.Error(t, fd.Err)
}

func TestInt_DefaultFunc(t *testing.T) {
	type CustomInt int

	f1 := func() int { return 1000 }
	fd := field.Int("id").DefaultFunc(f1).GoType(CustomInt(0)).Descriptor()
	assert.Error(t, fd.Err, "expect type (func() field_test.CustomInt) for default value")

	fd = field.Int("dir").GoType(CustomInt(0)).DefaultFunc(f1).Descriptor()
	assert.Error(t, fd.Err, "`var _ CustomInt = f2()` should fail")

	fd = field.Int("id").DefaultFunc(f1).Descriptor()
	assert.NoError(t, fd.Err)
}

func TestFloat(t *testing.T) {
	f := field.Float("age").Comment("comment").Positive()
	fd := f.Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeFloat64, fd.Info.Type)
	assert.Len(t, fd.Validators, 1)
	assert.Equal(t, "comment", fd.Comment)

	f = field.Float("age").Min(2.5).Max(5)
	fd = f.Descriptor()
	assert.Len(t, fd.Validators, 2)
	assert.Equal(t, field.TypeFloat32, field.Float32("age").Descriptor().Info.Type)

	type Count float64
	fd = field.Float("active").GoType(Count(0)).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "field_test.Count", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Count", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Float("count").GoType(&sql.NullFloat64{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "*sql.NullFloat64", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "*sql.NullFloat64", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Float("count").GoType(1).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Float("count").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Float("count").GoType(new(Count)).Descriptor()
	assert.Error(t, fd.Err)
}

func TestBool(t *testing.T) {
	fd := field.Bool("active").Default(true).Comment("comment").Immutable().Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBool, fd.Info.Type)
	assert.NotNil(t, fd.Default)
	assert.True(t, fd.Immutable)
	assert.Equal(t, true, fd.Default)
	assert.Equal(t, "comment", fd.Comment)

	type Status bool
	fd = field.Bool("active").GoType(Status(false)).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "field_test.Status", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Status", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Bool("deleted").GoType(&sql.NullBool{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "*sql.NullBool", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "*sql.NullBool", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Bool("active").GoType(1).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Bool("active").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Bool("active").GoType(new(Status)).Descriptor()
	assert.Error(t, fd.Err)
}

type Pair struct {
	K, V []byte
}

func (*Pair) Scan(interface{}) error      { return nil }
func (Pair) Value() (driver.Value, error) { return nil, nil }

func TestBytes(t *testing.T) {
	fd := field.Bytes("active").
		Unique().
		Default([]byte("{}")).
		Comment("comment").
		Validate(func(bytes []byte) error {
			return nil
		}).
		Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.True(t, fd.Unique)
	assert.Equal(t, field.TypeBytes, fd.Info.Type)
	assert.NotNil(t, fd.Default)
	assert.Equal(t, []byte("{}"), fd.Default)
	assert.Equal(t, "comment", fd.Comment)
	assert.Len(t, fd.Validators, 1)

	fd = field.Bytes("ip").GoType(net.IP("127.0.0.1")).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "net.IP", fd.Info.Ident)
	assert.Equal(t, "net", fd.Info.PkgPath)
	assert.Equal(t, "net.IP", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Bytes("blob").GoType(sql.NullString{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "sql.NullString", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullString", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Bytes("uuid").GoType(uuid.UUID{}).DefaultFunc(uuid.New).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "uuid.UUID", fd.Info.Ident)
	assert.Equal(t, "github.com/google/uuid", fd.Info.PkgPath)
	assert.Equal(t, "uuid.UUID", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())
	assert.NotEmpty(t, fd.Default.(func() uuid.UUID)())

	fd = field.Bytes("uuid").
		GoType(uuid.UUID{}).
		DefaultFunc(uuid.New).
		Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "uuid.UUID", fd.Info.String())
	fd = field.Bytes("pair").
		GoType(&Pair{}).
		Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "*field_test.Pair", fd.Info.String())

	fd = field.Bytes("blob").GoType(1).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Bytes("blob").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Bytes("blob").GoType(new(net.IP)).Descriptor()
	assert.Error(t, fd.Err)
}

func TestBytes_DefaultFunc(t *testing.T) {
	f1 := func() net.IP { return net.IP("0.0.0.0") }
	fd := field.Bytes("ip").GoType(net.IP("127.0.0.1")).DefaultFunc(f1).Descriptor()
	assert.NoError(t, fd.Err)

	var _ []byte = f1()
	fd = field.Bytes("ip").DefaultFunc(f1).Descriptor()
	assert.NoError(t, fd.Err)

	f2 := func() []byte { return []byte("0.0.0.0") }
	var _ net.IP = f2()
	fd = field.Bytes("ip").GoType(net.IP("127.0.0.1")).DefaultFunc(f2).Descriptor()
	assert.NoError(t, fd.Err)

	f3 := func() []uint8 { return []uint8("0.0.0.0") }
	var _ net.IP = f3()
	fd = field.Bytes("ip").GoType(net.IP("127.0.0.1")).DefaultFunc(f3).Descriptor()
	assert.NoError(t, fd.Err)
	fd = field.Bytes("ip").DefaultFunc(f3).Descriptor()
	assert.NoError(t, fd.Err)

	f4 := func() net.IPMask { return net.IPMask("ffff:ff80::") }
	fd = field.Bytes("ip").GoType(net.IP("127.0.0.1")).DefaultFunc(f4).Descriptor()
	assert.Error(t, fd.Err, "`var _ net.IP = f4()` should fail")
}

func TestString_DefaultFunc(t *testing.T) {
	f1 := func() http.Dir { return "/tmp" }
	fd := field.String("dir").GoType(http.Dir("/tmp")).DefaultFunc(f1).Descriptor()
	assert.NoError(t, fd.Err)

	fd = field.String("dir").DefaultFunc(f1).Descriptor()
	assert.Error(t, fd.Err, "`var _ string = f1()` should fail")

	f2 := func() string { return "/tmp" }
	fd = field.String("dir").GoType(http.Dir("/tmp")).DefaultFunc(f2).Descriptor()
	assert.Error(t, fd.Err, "`var _ http.Dir = f2()` should fail")

	f3 := func() sql.NullString { return sql.NullString{} }
	fd = field.String("str").GoType(sql.NullString{}).DefaultFunc(f3).Descriptor()
	assert.NoError(t, fd.Err)

	type S string
	f4 := func() S { return "" }
	fd = field.String("str").GoType(http.Dir("/tmp")).DefaultFunc(f4).Descriptor()
	assert.Error(t, fd.Err, "`var _ http.Dir = f4()` should fail")
}

type VString string

func (s *VString) Scan(interface{}) error {
	return nil
}

func (s VString) Value() (driver.Value, error) {
	return "", nil
}

func TestString(t *testing.T) {
	fd := field.String("name").
		DefaultFunc(func() string {
			return "Ent"
		}).
		Comment("comment").
		Descriptor()

	assert.Equal(t, "name", fd.Name)
	assert.Equal(t, field.TypeString, fd.Info.Type)
	assert.Equal(t, "Ent", fd.Default.(func() string)())
	assert.Equal(t, "comment", fd.Comment)

	re := regexp.MustCompile("[a-zA-Z0-9]")
	f := field.String("name").Unique().Match(re).Validate(func(string) error { return nil }).Sensitive()
	fd = f.Descriptor()
	assert.Equal(t, field.TypeString, fd.Info.Type)
	assert.Equal(t, "name", fd.Name)
	assert.True(t, fd.Unique)
	assert.Len(t, fd.Validators, 2)
	assert.True(t, fd.Sensitive)

	fd = field.String("name").GoType(http.Dir("dir")).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "http.Dir", fd.Info.Ident)
	assert.Equal(t, "net/http", fd.Info.PkgPath)
	assert.Equal(t, "http.Dir", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.String("name").GoType(http.MethodOptions).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "string", fd.Info.Ident)
	assert.Equal(t, "", fd.Info.PkgPath)
	assert.Equal(t, "string", fd.Info.String())
	assert.False(t, fd.Info.Nillable)

	fd = field.String("nullable_name").GoType(&sql.NullString{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "*sql.NullString", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "*sql.NullString", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())
	assert.False(t, fd.Info.Stringer())
	assert.True(t, fd.Info.RType.TypeEqual(reflect.TypeOf(&sql.NullString{})))

	fd = field.String("nullable_name").GoType(VString("")).Descriptor()
	assert.True(t, fd.Info.Valuer())
	assert.True(t, fd.Info.ValueScanner())
	assert.False(t, fd.Info.Stringer())

	type tURL struct {
		field.ValueScanner
		*url.URL
	}
	fd = field.String("nullable_url").GoType(&tURL{}).Descriptor()
	assert.Equal(t, "*field_test.tURL", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "*field_test.tURL", fd.Info.String())
	assert.True(t, fd.Info.ValueScanner())
	assert.True(t, fd.Info.Stringer())

	fd = field.String("name").GoType(1).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.String("name").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.String("name").GoType(new(http.Dir)).Descriptor()
	assert.Error(t, fd.Err)
}

func TestTime(t *testing.T) {
	now := time.Now()
	fd := field.Time("created_at").
		Default(func() time.Time {
			return now
		}).
		Comment("comment").
		Descriptor()
	assert.Equal(t, "created_at", fd.Name)
	assert.Equal(t, field.TypeTime, fd.Info.Type)
	assert.Equal(t, "time.Time", fd.Info.Type.String())
	assert.NotNil(t, fd.Default)
	assert.Equal(t, now, fd.Default.(func() time.Time)())
	assert.Equal(t, "comment", fd.Comment)

	fd = field.Time("updated_at").
		UpdateDefault(func() time.Time {
			return now
		}).
		Descriptor()
	assert.Equal(t, "updated_at", fd.Name)
	assert.Equal(t, now, fd.UpdateDefault.(func() time.Time)())

	type Time time.Time
	fd = field.Time("deleted_at").GoType(Time{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "field_test.Time", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Time", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Time("deleted_at").GoType(&sql.NullTime{}).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "*sql.NullTime", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "*sql.NullTime", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Time("active").GoType(1).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Time("active").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err)
	fd = field.Time("active").GoType(new(Time)).Descriptor()
	assert.Error(t, fd.Err)
}

func TestJSON(t *testing.T) {
	fd := field.JSON("name", map[string]string{}).
		Optional().
		Comment("comment").
		Descriptor()
	assert.True(t, fd.Optional)
	assert.Empty(t, fd.Info.PkgPath)
	assert.Equal(t, "name", fd.Name)
	assert.Equal(t, field.TypeJSON, fd.Info.Type)
	assert.Equal(t, "map[string]string", fd.Info.String())
	assert.Equal(t, "comment", fd.Comment)
	assert.True(t, fd.Info.Nillable)
	assert.False(t, fd.Info.RType.IsPtr())

	type T struct{ S string }
	fd = field.JSON("name", &T{}).
		Descriptor()
	assert.True(t, fd.Info.Nillable)
	assert.Equal(t, "*field_test.T", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.True(t, fd.Info.RType.IsPtr())
	assert.Equal(t, "T", fd.Info.RType.Name)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.RType.PkgPath)

	fd = field.JSON("dir", http.Dir("dir")).
		Optional().
		Descriptor()
	assert.True(t, fd.Optional)
	assert.Equal(t, field.TypeJSON, fd.Info.Type)
	assert.Equal(t, "dir", fd.Name)
	assert.Equal(t, "net/http", fd.Info.PkgPath)
	assert.Equal(t, "http.Dir", fd.Info.String())
	assert.False(t, fd.Info.Nillable)

	fd = field.Strings("strings").
		Optional().
		Descriptor()
	assert.True(t, fd.Optional)
	assert.Empty(t, fd.Info.PkgPath)
	assert.Equal(t, "strings", fd.Name)
	assert.Equal(t, field.TypeJSON, fd.Info.Type)
	assert.Equal(t, "[]string", fd.Info.String())

	fd = field.JSON("values", &url.Values{}).Descriptor()
	assert.Equal(t, "net/url", fd.Info.PkgPath)
	fd = field.JSON("values", []url.Values{}).Descriptor()
	assert.Equal(t, "net/url", fd.Info.PkgPath)
	fd = field.JSON("values", []*url.Values{}).Descriptor()
	assert.Equal(t, "net/url", fd.Info.PkgPath)
	fd = field.JSON("values", map[string]url.Values{}).Descriptor()
	assert.Equal(t, "net/url", fd.Info.PkgPath)
	fd = field.JSON("values", map[string]*url.Values{}).Descriptor()
	assert.Equal(t, "net/url", fd.Info.PkgPath)
}

func TestField_Tag(t *testing.T) {
	fd := field.Bool("expired").
		StructTag(`json:"expired,omitempty"`).
		Descriptor()
	assert.Equal(t, `json:"expired,omitempty"`, fd.Tag)
}

type Role string

func (Role) Values() []string {
	return []string{"admin", "owner"}
}

type RoleInt int32

func (RoleInt) Values() []string {
	return []string{"unknown", "admin", "owner"}
}

func (i RoleInt) String() string {
	switch i {
	case 1:
		return "admin"
	case 2:
		return "owner"
	default:
		return "unknown"
	}
}

func (i RoleInt) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *RoleInt) Scan(val interface{}) error {
	switch v := val.(type) {
	case string:
		switch v {
		case "admin":
			*i = 1
		case "owner":
			*i = 2
		default:
			*i = 0
		}
	default:
		return errors.New("bad enum value")
	}

	return nil
}

func TestField_Enums(t *testing.T) {
	fd := field.Enum("role").
		Values(
			"user",
			"admin",
			"master",
		).
		Default("user").
		Comment("comment").
		Descriptor()
	assert.Equal(t, "role", fd.Name)
	assert.Equal(t, "user", fd.Enums[0].V)
	assert.Equal(t, "admin", fd.Enums[1].V)
	assert.Equal(t, "master", fd.Enums[2].V)
	assert.Equal(t, "user", fd.Default)
	assert.Equal(t, "comment", fd.Comment)

	fd = field.Enum("role").
		NamedValues("USER", "user").
		Default("user").
		Descriptor()
	assert.Equal(t, "role", fd.Name)
	assert.Equal(t, "USER", fd.Enums[0].N)
	assert.Equal(t, "user", fd.Enums[0].V)
	assert.Equal(t, "user", fd.Default)

	fd = field.Enum("role").GoType(Role("")).Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "field_test.Role", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Role", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())
	assert.Equal(t, "admin", fd.Enums[0].V)
	assert.Equal(t, "owner", fd.Enums[1].V)
	assert.False(t, fd.Info.Stringer())

	fd = field.Enum("role").GoType(RoleInt(0)).Descriptor()
	assert.Equal(t, "field_test.RoleInt", fd.Info.Ident)
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.RoleInt", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())
	assert.Equal(t, "unknown", fd.Enums[0].V)
	assert.Equal(t, "admin", fd.Enums[1].V)
	assert.Equal(t, "owner", fd.Enums[2].V)
	assert.True(t, fd.Info.Stringer())
}

func TestField_UUID(t *testing.T) {
	fd := field.UUID("id", uuid.UUID{}).
		Unique().
		Default(uuid.New).
		Comment("comment").
		Nillable().
		Descriptor()
	assert.Equal(t, "id", fd.Name)
	assert.True(t, fd.Unique)
	assert.Equal(t, "uuid.UUID", fd.Info.String())
	assert.Equal(t, "github.com/google/uuid", fd.Info.PkgPath)
	assert.NotNil(t, fd.Default)
	assert.NotEmpty(t, fd.Default.(func() uuid.UUID)())
	assert.Equal(t, "comment", fd.Comment)
	assert.True(t, fd.Nillable)

	fd = field.UUID("id", &uuid.UUID{}).
		Descriptor()
	assert.Equal(t, "github.com/google/uuid", fd.Info.PkgPath)

	fd = field.UUID("id", uuid.UUID{}).
		Default(uuid.UUID{}).
		Descriptor()
	assert.EqualError(t, fd.Err, "expect type (func() uuid.UUID) for uuid default value")
}

type custom struct {
}

func (c *custom) Scan(_ interface{}) (err error) {
	return nil
}

func (c custom) Value() (driver.Value, error) {
	return nil, nil
}

func TestField_Other(t *testing.T) {
	fd := field.Other("other", &custom{}).
		Unique().
		Default(&custom{}).
		SchemaType(map[string]string{dialect.Postgres: "varchar"}).
		Descriptor()
	assert.NoError(t, fd.Err)
	assert.Equal(t, "other", fd.Name)
	assert.True(t, fd.Unique)
	assert.Equal(t, "*field_test.custom", fd.Info.String())
	assert.Equal(t, "entgo.io/ent/schema/field_test", fd.Info.PkgPath)
	assert.NotNil(t, fd.Default)

	fd = field.Other("other", &custom{}).
		Descriptor()
	assert.Error(t, fd.Err, "missing SchemaType option")

	fd = field.Other("other", &custom{}).
		SchemaType(map[string]string{dialect.Postgres: "varchar"}).
		Default(func() *custom { return &custom{} }).
		Descriptor()
	assert.NoError(t, fd.Err)

	fd = field.Other("other", custom{}).
		SchemaType(map[string]string{dialect.Postgres: "varchar"}).
		Default(func() custom { return custom{} }).
		Descriptor()
	assert.NoError(t, fd.Err)

	fd = field.Other("other", &custom{}).
		SchemaType(map[string]string{dialect.Postgres: "varchar"}).
		Default(func() custom { return custom{} }).
		Descriptor()
	assert.Error(t, fd.Err, "invalid default value")
}

func TestTypeString(t *testing.T) {
	typ := field.TypeBool
	assert.Equal(t, "bool", typ.String())
	typ = field.TypeInvalid
	assert.Equal(t, "invalid", typ.String())
	typ = 21
	assert.Equal(t, "invalid", typ.String())
}

func TestTypeNumeric(t *testing.T) {
	typ := field.TypeBool
	assert.False(t, typ.Numeric())
	typ = field.TypeUint8
	assert.True(t, typ.Numeric())
}

func TestTypeValid(t *testing.T) {
	typ := field.TypeBool
	assert.True(t, typ.Valid())
	typ = 0
	assert.False(t, typ.Valid())
	typ = 21
	assert.False(t, typ.Valid())
}

func TestTypeConstName(t *testing.T) {
	typ := field.TypeJSON
	assert.Equal(t, "TypeJSON", typ.ConstName())
	typ = field.TypeInt
	assert.Equal(t, "TypeInt", typ.ConstName())
	typ = field.TypeInt64
	assert.Equal(t, "TypeInt64", typ.ConstName())
	typ = field.TypeOther
	assert.Equal(t, "TypeOther", typ.ConstName())
	typ = 21
	assert.Equal(t, "invalid", typ.ConstName())
}
