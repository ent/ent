// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field_test

import (
	"database/sql"
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
	"time"

	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type GenericText[T any] string
type GenericPair[A, B any] string
type GenericAlias = GenericText[time.Time]

func TestGoTypeGeneric(t *testing.T) {
	for _, tt := range []struct {
		value   any
		ident   string
		imports map[string]string
	}{
		{
			value: GenericText[string](""),
			ident: "field_test.GenericText[string]",
		},
		{
			value:   GenericText[time.Time](""),
			ident:   "field_test.GenericText[time.Time]",
			imports: map[string]string{"time": "time"},
		},
		{
			value:   GenericAlias(""),
			ident:   "field_test.GenericText[time.Time]",
			imports: map[string]string{"time": "time"},
		},
		{
			value:   GenericText[uuid.UUID](""),
			ident:   "field_test.GenericText[uuid.UUID]",
			imports: map[string]string{"github.com/google/uuid": "uuid"},
		},
		{
			value:   GenericPair[uuid.UUID, time.Time](""),
			ident:   "field_test.GenericPair[uuid.UUID,time.Time]",
			imports: map[string]string{"github.com/google/uuid": "uuid", "time": "time"},
		},
		{
			value:   GenericText[GenericText[uuid.UUID]](""),
			ident:   "field_test.GenericText[field_test.GenericText[uuid.UUID]]",
			imports: map[string]string{"github.com/google/uuid": "uuid"},
		},
		{
			value:   GenericText[map[url.URL][]*time.Time](""),
			ident:   "field_test.GenericText[map[url.URL][]*time.Time]",
			imports: map[string]string{"net/url": "url", "time": "time"},
		},
		{
			value: GenericText[struct {
				Value uuid.UUID `json:"example.com/value"`
			}](""),
			ident:   `field_test.GenericText[struct { Value uuid.UUID "json:\"example.com/value\"" }]`,
			imports: map[string]string{"github.com/google/uuid": "uuid"},
		},
	} {
		t.Run(tt.ident, func(t *testing.T) {
			d := field.String("value").GoType(tt.value).Descriptor()
			require.NoError(t, d.Err)
			require.Equal(t, tt.ident, d.Info.Ident)
			require.Equal(t, "field_test", d.Info.PkgName)
			require.Equal(t, "entgo.io/ent/schema/field_test", d.Info.PkgPath)
			if tt.imports == nil {
				tt.imports = make(map[string]string)
			}
			tt.imports[d.Info.PkgPath] = "field_test"
			require.Equal(t, tt.imports, d.Info.PkgImports)
			// Import metadata must survive the schema loader's JSON boundary.
			data, err := json.Marshal(d.Info)
			require.NoError(t, err)
			var info field.TypeInfo
			require.NoError(t, json.Unmarshal(data, &info))
			require.Equal(t, d.Info.PkgImports, info.PkgImports)
			// Runtime type identity is independent of source-level aliases.
			require.True(t, info.RType.TypeEqual(reflect.TypeOf(tt.value)))
		})
	}
}

func TestGenericCompositeTypes(t *testing.T) {
	for _, tt := range []struct {
		value any
		ident string
	}{
		{new(GenericText[uuid.UUID]), "*field_test.GenericText[uuid.UUID]"},
		{[]GenericText[uuid.UUID]{}, "[]field_test.GenericText[uuid.UUID]"},
		{[2]GenericText[uuid.UUID]{}, "[2]field_test.GenericText[uuid.UUID]"},
		{map[string]*GenericText[uuid.UUID]{}, "map[string]*field_test.GenericText[uuid.UUID]"},
		{map[uuid.UUID]GenericText[time.Time]{}, "map[uuid.UUID]field_test.GenericText[time.Time]"},
	} {
		t.Run(tt.ident, func(t *testing.T) {
			d := field.JSON("value", tt.value).Descriptor()
			require.NoError(t, d.Err)
			require.Equal(t, tt.ident, d.Info.Ident)
			require.Equal(t, "uuid", d.Info.PkgImports["github.com/google/uuid"])
		})
	}
	d := field.Time("value").GoType(&sql.Null[time.Time]{}).Descriptor()
	require.NoError(t, d.Err)
	require.Equal(t, "*sql.Null[time.Time]", d.Info.Ident)
	require.True(t, d.Info.ValueScanner())
}
