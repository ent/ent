// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field_test

import (
	"database/sql"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/field"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	fd := field.Int("age").
		Positive().
		Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeInt, fd.Info.Type)
	assert.Len(t, fd.Validators, 1)

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
	assert.NoError(t, fd.Err())
	assert.Equal(t, "field_test.Count", fd.Info.Ident)
	assert.Equal(t, "github.com/facebook/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Count", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Int("count").GoType(&sql.NullInt64{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "sql.NullInt64", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullInt64", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Int("count").GoType(sql.NullInt64{}).Descriptor()
	assert.EqualError(t, fd.Err(), `GoType must be a "int" type or ValueScanner. Use *sql.NullInt64 instead`)
	fd = field.Int("count").GoType(false).Descriptor()
	assert.EqualError(t, fd.Err(), `GoType must be a "int" type or ValueScanner`)
	fd = field.Int("count").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Int("count").GoType(new(Count)).Descriptor()
	assert.Error(t, fd.Err())
}

func TestFloat(t *testing.T) {
	f := field.Float("age").Positive()
	fd := f.Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeFloat64, fd.Info.Type)
	assert.Len(t, fd.Validators, 1)

	f = field.Float("age").Min(2.5).Max(5)
	fd = f.Descriptor()
	assert.Len(t, fd.Validators, 2)
	assert.Equal(t, field.TypeFloat32, field.Float32("age").Descriptor().Info.Type)

	type Count float64
	fd = field.Float("active").GoType(Count(0)).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "field_test.Count", fd.Info.Ident)
	assert.Equal(t, "github.com/facebook/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Count", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Float("count").GoType(&sql.NullFloat64{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "sql.NullFloat64", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullFloat64", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Float("count").GoType(1).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Float("count").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Float("count").GoType(new(Count)).Descriptor()
	assert.Error(t, fd.Err())
}

func TestBool(t *testing.T) {
	fd := field.Bool("active").Default(true).Immutable().Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBool, fd.Info.Type)
	assert.NotNil(t, fd.Default)
	assert.True(t, fd.Immutable)
	assert.Equal(t, true, fd.Default)

	type Status bool
	fd = field.Bool("active").GoType(Status(false)).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "field_test.Status", fd.Info.Ident)
	assert.Equal(t, "github.com/facebook/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Status", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Bool("deleted").GoType(&sql.NullBool{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "sql.NullBool", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullBool", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Bool("active").GoType(1).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Bool("active").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Bool("active").GoType(new(Status)).Descriptor()
	assert.Error(t, fd.Err())
}

func TestBytes(t *testing.T) {
	fd := field.Bytes("active").Default([]byte("{}")).Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBytes, fd.Info.Type)
	assert.NotNil(t, fd.Default)
	assert.Equal(t, []byte("{}"), fd.Default)

	fd = field.Bytes("ip").GoType(net.IP("127.0.0.1")).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "net.IP", fd.Info.Ident)
	assert.Equal(t, "net", fd.Info.PkgPath)
	assert.Equal(t, "net.IP", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Bytes("blob").GoType(&sql.NullString{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "sql.NullString", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullString", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Bytes("blob").GoType(1).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Bytes("blob").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Bytes("blob").GoType(new(net.IP)).Descriptor()
	assert.Error(t, fd.Err())
}

func TestString(t *testing.T) {
	re := regexp.MustCompile("[a-zA-Z0-9]")
	f := field.String("name").Unique().Match(re).Validate(func(string) error { return nil }).Sensitive()
	fd := f.Descriptor()
	assert.Equal(t, field.TypeString, fd.Info.Type)
	assert.Equal(t, "name", fd.Name)
	assert.True(t, fd.Unique)
	assert.Len(t, fd.Validators, 2)
	assert.True(t, fd.Sensitive)

	fd = field.String("name").GoType(http.Dir("dir")).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "http.Dir", fd.Info.Ident)
	assert.Equal(t, "net/http", fd.Info.PkgPath)
	assert.Equal(t, "http.Dir", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.String("name").GoType(http.MethodOptions).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "string", fd.Info.Ident)
	assert.Equal(t, "", fd.Info.PkgPath)
	assert.Equal(t, "string", fd.Info.String())
	assert.False(t, fd.Info.Nillable)

	fd = field.String("nullable_name").GoType(&sql.NullString{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "sql.NullString", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullString", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())
	assert.False(t, fd.Info.Stringer())
	assert.True(t, fd.Info.RType.TypeEqual(reflect.TypeOf(sql.NullString{})))
	assert.True(t, fd.Info.RType.TypeEqual(reflect.TypeOf(&sql.NullString{})))

	type tURL struct {
		field.ValueScanner
		*url.URL
	}
	fd = field.String("nullable_url").GoType(&tURL{}).Descriptor()
	assert.Equal(t, "field_test.tURL", fd.Info.Ident)
	assert.Equal(t, "github.com/facebook/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.tURL", fd.Info.String())
	assert.True(t, fd.Info.ValueScanner())
	assert.True(t, fd.Info.Stringer())

	fd = field.String("name").GoType(1).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.String("name").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.String("name").GoType(new(http.Dir)).Descriptor()
	assert.Error(t, fd.Err())
}

func TestTime(t *testing.T) {
	now := time.Now()
	fd := field.Time("created_at").
		Default(func() time.Time {
			return now
		}).
		Descriptor()
	assert.Equal(t, "created_at", fd.Name)
	assert.Equal(t, field.TypeTime, fd.Info.Type)
	assert.Equal(t, "time.Time", fd.Info.Type.String())
	assert.NotNil(t, fd.Default)
	assert.Equal(t, now, fd.Default.(func() time.Time)())

	fd = field.Time("updated_at").
		UpdateDefault(func() time.Time {
			return now
		}).
		Descriptor()
	assert.Equal(t, "updated_at", fd.Name)
	assert.Equal(t, now, fd.UpdateDefault.(func() time.Time)())

	type Time time.Time
	fd = field.Time("deleted_at").GoType(Time{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "field_test.Time", fd.Info.Ident)
	assert.Equal(t, "github.com/facebook/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Time", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())

	fd = field.Time("deleted_at").GoType(&sql.NullTime{}).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "sql.NullTime", fd.Info.Ident)
	assert.Equal(t, "database/sql", fd.Info.PkgPath)
	assert.Equal(t, "sql.NullTime", fd.Info.String())
	assert.True(t, fd.Info.Nillable)
	assert.True(t, fd.Info.ValueScanner())

	fd = field.Time("active").GoType(1).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Time("active").GoType(struct{}{}).Descriptor()
	assert.Error(t, fd.Err())
	fd = field.Time("active").GoType(new(Time)).Descriptor()
	assert.Error(t, fd.Err())
}

func TestJSON(t *testing.T) {
	fd := field.JSON("name", map[string]string{}).
		Optional().
		Descriptor()
	assert.True(t, fd.Optional)
	assert.Empty(t, fd.Info.PkgPath)
	assert.Equal(t, "name", fd.Name)
	assert.Equal(t, field.TypeJSON, fd.Info.Type)
	assert.Equal(t, "map[string]string", fd.Info.String())

	fd = field.JSON("dir", http.Dir("dir")).
		Optional().
		Descriptor()
	assert.True(t, fd.Optional)
	assert.Equal(t, field.TypeJSON, fd.Info.Type)
	assert.Equal(t, "dir", fd.Name)
	assert.Equal(t, "net/http", fd.Info.PkgPath)
	assert.Equal(t, "http.Dir", fd.Info.String())

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

func TestField_Enums(t *testing.T) {
	fd := field.Enum("role").
		Values(
			"user",
			"admin",
			"master",
		).
		Default("user").
		Descriptor()
	assert.Equal(t, "role", fd.Name)
	assert.Equal(t, "user", fd.Enums[0].V)
	assert.Equal(t, "admin", fd.Enums[1].V)
	assert.Equal(t, "master", fd.Enums[2].V)
	assert.Equal(t, "user", fd.Default)

	fd = field.Enum("role").
		NamedValues("USER", "user").
		Default("user").
		Descriptor()
	assert.Equal(t, "role", fd.Name)
	assert.Equal(t, "USER", fd.Enums[0].N)
	assert.Equal(t, "user", fd.Enums[0].V)
	assert.Equal(t, "user", fd.Default)

	fd = field.Enum("role").
		ValueMap(map[string]string{"USER": "user"}).
		Default("user").
		Descriptor()
	assert.Equal(t, "role", fd.Name)
	assert.Equal(t, "USER", fd.Enums[0].N)
	assert.Equal(t, "user", fd.Enums[0].V)

	fd = field.Enum("role").GoType(Role("")).Descriptor()
	assert.NoError(t, fd.Err())
	assert.Equal(t, "field_test.Role", fd.Info.Ident)
	assert.Equal(t, "github.com/facebook/ent/schema/field_test", fd.Info.PkgPath)
	assert.Equal(t, "field_test.Role", fd.Info.String())
	assert.False(t, fd.Info.Nillable)
	assert.False(t, fd.Info.ValueScanner())
	assert.Equal(t, "admin", fd.Enums[0].V)
	assert.Equal(t, "owner", fd.Enums[1].V)
}

func TestField_UUID(t *testing.T) {
	fd := field.UUID("id", uuid.UUID{}).
		Unique().
		Default(uuid.New).
		Descriptor()
	assert.Equal(t, "id", fd.Name)
	assert.True(t, fd.Unique)
	assert.Equal(t, "uuid.UUID", fd.Info.String())
	assert.Equal(t, "github.com/google/uuid", fd.Info.PkgPath)
	assert.NotNil(t, fd.Default)
	assert.NotEmpty(t, fd.Default.(func() uuid.UUID)())

	fd = field.UUID("id", uuid.UUID{}).
		Default(uuid.UUID{}).
		Descriptor()
	assert.EqualError(t, fd.Err(), "expect type (func() uuid.UUID) for uuid default value")
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
	typ = 21
	assert.Equal(t, "invalid", typ.ConstName())
}
