// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package sqlgraph

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/json/jsontext"
	"errors"
	"testing"

	"entgo.io/ent/schema/field"
	"github.com/stretchr/testify/require"
)

func TestJSONFieldOptions(t *testing.T) {
	for _, tt := range []struct {
		name  string
		value any
		want  string
	}{
		{"nil slice", []string(nil), `[]`},
		{"nil map", map[string]string(nil), `{}`},
		{"html", "<tag>&", `"<tag>&"`},
		{"omitempty", struct {
			N int `json:"n,omitempty"`
		}{}, `{"n":0}`},
		{"marshaler", v2Marshaler{}, `"v2"`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			create := &CreateSpec{}
			create.SetJSONField("v2", tt.value)
			update := &UpdateSpec{}
			update.SetJSONField("v2", tt.value)
			for _, fields := range [][]*FieldSpec{create.Fields, update.Fields.Set} {
				var got driver.Value
				err := setTableColumns(fields, nil, func(_ string, value driver.Value) { got = value })
				require.NoError(t, err)
				require.Equal(t, json.RawMessage(tt.want), got)
			}
		})
	}
	create := &CreateSpec{}
	create.SetField("v1", field.TypeJSON, []string(nil))
	create.SetJSONField("v2", []string(nil))
	got := make(map[string]driver.Value)
	require.NoError(t, setTableColumns(create.Fields, nil, func(column string, value driver.Value) { got[column] = value }))
	require.Equal(t, json.RawMessage(`null`), got["v1"])
	require.Equal(t, json.RawMessage(`[]`), got["v2"])

	for _, value := range []any{"\xff", json.RawMessage(`{"a":1,"a":2}`), invalidV2Marshaler{}} {
		create := &CreateSpec{}
		create.SetJSONField("invalid", value)
		err := setTableColumns(create.Fields, nil, func(string, driver.Value) { t.Fatal("invalid JSON must not be stored") })
		require.ErrorContains(t, err, "marshal value for column invalid")
	}
}

type v2Marshaler struct{}

func (v2Marshaler) MarshalJSONTo(enc *jsontext.Encoder) error {
	return enc.WriteToken(jsontext.String("v2"))
}

type invalidV2Marshaler struct{}

func (invalidV2Marshaler) MarshalJSONTo(*jsontext.Encoder) error {
	return errors.New("invalid value")
}
