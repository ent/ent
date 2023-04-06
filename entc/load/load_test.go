// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"testing"

	"entgo.io/ent/schema/field"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	cfg := &Config{Path: "./testdata/valid"}
	spec, err := cfg.Load()
	require.NoError(t, err)
	require.Len(t, spec.Schemas, 3)
	require.Equal(t, "entgo.io/ent/entc/load/testdata/valid", spec.PkgPath)

	require.Equal(t, "Group", spec.Schemas[0].Name, "ordered alphabetically")
	require.Equal(t, "Tag", spec.Schemas[1].Name)
	require.Equal(t, "User", spec.Schemas[2].Name)
}

func TestLoadWrongPath(t *testing.T) {
	cfg := &Config{Path: "./boring"}
	plg, err := cfg.Load()
	require.Error(t, err)
	require.Nil(t, plg)
}

func TestLoadSpecific(t *testing.T) {
	cfg := &Config{Path: "./testdata/valid", Names: []string{"User"}}
	spec, err := cfg.Load()
	require.NoError(t, err)
	require.Len(t, spec.Schemas, 1)
	require.Equal(t, "User", spec.Schemas[0].Name)
	require.Equal(t, "entgo.io/ent/entc/load/testdata/valid", spec.PkgPath)
}

func TestLoadNoSchema(t *testing.T) {
	cfg := &Config{Path: "./testdata/invalid"}
	schemas, err := cfg.Load()
	require.Error(t, err)
	require.Empty(t, schemas)
}

func TestLoadSchemaFailure(t *testing.T) {
	cfg := &Config{Path: "./testdata/failure"}
	spec, err := cfg.Load()
	require.Error(t, err)
	require.Nil(t, spec)
}

func TestLoadBaseSchema(t *testing.T) {
	cfg := &Config{Path: "./testdata/base"}
	spec, err := cfg.Load()
	require.NoError(t, err)
	require.Len(t, spec.Schemas, 1)
	require.Len(t, spec.Schemas[0].Fields, 2, "embedded base schema")
	f1 := spec.Schemas[0].Fields[0]
	require.Equal(t, "base_field", f1.Name)
	require.Equal(t, field.TypeInt, f1.Info.Type)
	f2 := spec.Schemas[0].Fields[1]
	require.Equal(t, "user_field", f2.Name)
	require.Equal(t, field.TypeString, f2.Info.Type)
}

func TestLoadTags(t *testing.T) {
	all, err := (&Config{
		Path: "./testdata/buildflags",
	}).Load()
	require.NoError(t, err)

	require.Len(t, all.Schemas, 2)
	require.Equal(t, "Group", all.Schemas[0].Name, "ordered alphabetically")
	require.Equal(t, "User", all.Schemas[1].Name)

	notags, err := (&Config{
		Path:       "./testdata/buildflags",
		BuildFlags: []string{"-tags", "hidegroups"},
	}).Load()
	require.NoError(t, err)

	require.Len(t, notags.Schemas, 1)
	require.Equal(t, "User", notags.Schemas[0].Name)

	require.Equal(t, all.Schemas[1], notags.Schemas[0])
}

func TestLoadCycleError(t *testing.T) {
	cfg := &Config{Path: "./testdata/cycle"}
	spec, err := cfg.Load()
	require.Nil(t, spec)
	require.EqualError(t, err, `entc/load: parse schema dir: import cycle not allowed: import stack: [entgo.io/ent/entc/load/testdata/cycle entgo.io/ent/entc/load/testdata/cycle/fakent entgo.io/ent/entc/load/testdata/cycle]
To resolve this issue, move the custom types used by the generated code to a separate package: "Enum", "Used"`)
}
