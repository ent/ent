// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package sqlgraph

import (
	jsonv2 "encoding/json/v2"

	"entgo.io/ent/schema/field"
)

// SetJSONField appends a JSON field setter using the supplied encoding options.
// Nil options use encoding/json/v2 defaults.
func (u *CreateSpec) SetJSONField(column string, value any, opts ...jsonv2.Options) {
	u.Fields = append(u.Fields, jsonField(column, value, opts...))
}

// SetJSONField appends a JSON field setter using the supplied encoding options.
// Nil options use encoding/json/v2 defaults.
func (u *UpdateSpec) SetJSONField(column string, value any, opts ...jsonv2.Options) {
	u.Fields.Set = append(u.Fields.Set, jsonField(column, value, opts...))
}

func jsonField(column string, value any, opts ...jsonv2.Options) *FieldSpec {
	return &FieldSpec{
		Column:      column,
		Type:        field.TypeJSON,
		Value:       value,
		JSONMarshal: func(v any) ([]byte, error) { return jsonv2.Marshal(v, opts...) },
	}
}
