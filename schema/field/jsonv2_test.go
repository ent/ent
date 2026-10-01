// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package field_test

import (
	jsonv2 "encoding/json/v2"
	"net/url"
	"testing"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/stretchr/testify/require"
)

func TestJSONOptionsBuilder(t *testing.T) {
	fd := field.JSON("urls", []*url.URL{}).
		JSONOptions(jsonv2.DefaultOptionsV2()).
		Optional().
		Default(func() []*url.URL { return nil }).
		StorageKey("links").
		Comment("URLs encoded with JSON options").
		Descriptor()
	require.NoError(t, fd.Err)
	require.Equal(t, field.TypeJSON, fd.Info.Type)
	require.Equal(t, "[]*url.URL", fd.Info.String())
	require.Equal(t, "net/url", fd.Info.PkgPath)
	require.True(t, fd.Info.Nillable)
	require.True(t, fd.Optional)
	require.Equal(t, "links", fd.StorageKey)
	require.NotNil(t, fd.Default)
	require.NotNil(t, fd.JSONOptions)
	require.Nil(t, field.JSON("urls", []*url.URL{}).Descriptor().JSONOptions)
	require.Nil(t, field.JSON("urls", []*url.URL{}).JSONOptions().Descriptor().JSONOptions)
	require.EqualError(t, field.JSON("invalid", nil).JSONOptions(jsonv2.DefaultOptionsV2()).Descriptor().Err, "expect a Go value as JSON type but got nil")
	require.Error(t, field.JSON("urls", []*url.URL{}).JSONOptions(jsonv2.DefaultOptionsV2()).Default([]string{}).Descriptor().Err)

	opts := []jsonv2.Options{jsonv2.Deterministic(true), jsonv2.FormatNilSliceAsNull(true)}
	b := field.JSON("tags", []string{}).JSONOptions(opts...)
	opts[0] = jsonv2.Deterministic(false)
	saved := b.Descriptor().JSONOptions.([]jsonv2.Options)
	v, ok := jsonv2.GetOption(jsonv2.JoinOptions(saved...), jsonv2.Deterministic)
	require.True(t, ok)
	require.True(t, v, "the descriptor must own its options slice")

	// Repeated calls append options, with later values overriding earlier ones.
	b.JSONOptions(jsonv2.Deterministic(false), jsonv2.FormatNilSliceAsNull(false))
	saved = b.Descriptor().JSONOptions.([]jsonv2.Options)
	v, ok = jsonv2.GetOption(jsonv2.JoinOptions(saved...), jsonv2.Deterministic)
	require.True(t, ok)
	require.False(t, v)
	data, err := jsonv2.Marshal([]string(nil), saved...)
	require.NoError(t, err)
	require.JSONEq(t, "[]", string(data))
}

func TestJSONOptionsCompatibility(t *testing.T) {
	for _, f := range []ent.Field{
		field.JSON("json", []string{}).JSONOptions(jsonv2.Deterministic(true)),
		field.Any("any").JSONOptions(jsonv2.Deterministic(true)),
		field.Strings("strings").JSONOptions(jsonv2.Deterministic(true)).Default([]string{}),
		field.Ints("ints").JSONOptions(jsonv2.Deterministic(true)).Default([]int{}),
		field.Floats("floats").JSONOptions(jsonv2.Deterministic(true)).Default([]float64{}),
	} {
		t.Run(f.Descriptor().Name, func(t *testing.T) {
			opts := f.Descriptor().JSONOptions.([]jsonv2.Options)
			data, err := jsonv2.Marshal([]string(nil), opts...)
			require.NoError(t, err)
			require.JSONEq(t, "null", string(data))
			var v struct{ Name string }
			require.NoError(t, jsonv2.Unmarshal([]byte(`{"NAME":"first","NAME":"last"}`), &v, opts...))
			require.Equal(t, "last", v.Name)
		})
	}

	// Full v2 defaults are an explicit option, overriding the compatibility bundle.
	opts := field.JSON("modern", []string{}).JSONOptions(jsonv2.DefaultOptionsV2()).Descriptor().JSONOptions.([]jsonv2.Options)
	data, err := jsonv2.Marshal([]string(nil), opts...)
	require.NoError(t, err)
	require.JSONEq(t, "[]", string(data))
	var v struct{ Name string }
	require.Error(t, jsonv2.Unmarshal([]byte(`{"NAME":"first","NAME":"last"}`), &v, opts...))
}
