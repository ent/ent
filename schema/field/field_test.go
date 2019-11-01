// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field_test

import (
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/facebookincubator/ent/schema/field"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	f := field.Int("age").Positive()
	fd := f.Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeInt, fd.Info.Type)
	assert.Len(t, fd.Validators, 1)

	f = field.Int("age").Default(10).Min(10).Max(20)
	fd = f.Descriptor()
	assert.NotNil(t, fd.Default)
	assert.Equal(t, 10, fd.Default)
	assert.Len(t, fd.Validators, 2)

	f = field.Int("age").Range(20, 40).Nillable()
	fd = f.Descriptor()
	assert.Nil(t, fd.Default)
	assert.True(t, fd.Nillable)
	assert.False(t, fd.Immutable)
	assert.Len(t, fd.Validators, 1)

	assert.Equal(t, field.TypeInt8, field.Int8("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeInt16, field.Int16("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeInt32, field.Int32("age").Descriptor().Info.Type)
	assert.Equal(t, field.TypeInt64, field.Int64("age").Descriptor().Info.Type)
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
}

func TestBool(t *testing.T) {
	f := field.Bool("active").Default(true).Immutable()
	fd := f.Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBool, fd.Info.Type)
	assert.NotNil(t, fd.Default)
	assert.True(t, fd.Immutable)
	assert.Equal(t, true, fd.Default)
}

func TestBytes(t *testing.T) {
	f := field.Bytes("active").Default([]byte("{}"))
	fd := f.Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBytes, fd.Info.Type)
	assert.NotNil(t, fd.Default)
	assert.Equal(t, []byte("{}"), fd.Default)
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
}

func TestJSON(t *testing.T) {
	fd := field.JSON("name", map[string]string{}).
		Optional().
		Descriptor()
	require.True(t, fd.Optional)
	require.Empty(t, fd.Info.PkgPath)
	require.Equal(t, "name", fd.Name)
	require.Equal(t, field.TypeJSON, fd.Info.Type)
	require.Equal(t, "map[string]string", fd.Info.String())

	fd = field.JSON("dir", http.Dir("dir")).
		Optional().
		Descriptor()
	require.True(t, fd.Optional)
	require.Equal(t, field.TypeJSON, fd.Info.Type)
	require.Equal(t, "dir", fd.Name)
	require.Equal(t, "net/http", fd.Info.PkgPath)
	require.Equal(t, "http.Dir", fd.Info.String())

	fd = field.Strings("strings").
		Optional().
		Descriptor()
	require.True(t, fd.Optional)
	require.Empty(t, fd.Info.PkgPath)
	require.Equal(t, "strings", fd.Name)
	require.Equal(t, field.TypeJSON, fd.Info.Type)
	require.Equal(t, "[]string", fd.Info.String())
}

func TestField_Tag(t *testing.T) {
	fd := field.Bool("expired").
		StructTag(`json:"expired,omitempty"`).
		Descriptor()
	assert.Equal(t, `json:"expired,omitempty"`, fd.Tag)
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
	require.Equal(t, "role", fd.Name)
	require.Equal(t, []string{"user", "admin", "master"}, fd.Enums)
	require.Equal(t, "user", fd.Default)
}
