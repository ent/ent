// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package field

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
)

// JSONOptions adds options for marshaling and unmarshaling the field, including
// slice appends and projected scans. Options override the field's v1 compatibility
// defaults, with later options taking precedence. Use jsonv2.DefaultOptionsV2()
// to select full v2 semantics.
// JSONOptions requires the jsonv2 experiment (enabled by default in Go 1.27)
// and is supported only by SQL storage.
//
//	field.JSON("info", &Info{}).
//		JSONOptions(jsonv2.DefaultOptionsV2(), jsonv2.RejectUnknownMembers(true)).
//		Optional()
func (b *jsonBuilder) JSONOptions(opts ...jsonv2.Options) *jsonBuilder {
	if len(opts) == 0 {
		return b
	}
	current, ok := b.desc.JSONOptions.([]jsonv2.Options)
	if !ok {
		current = []jsonv2.Options{json.DefaultOptionsV1()}
	}
	b.desc.JSONOptions = append(current, opts...)
	return b
}

// JSONOptions adds JSON encoding options while retaining the slice builder.
// See jsonBuilder.JSONOptions for the option semantics.
func (b *sliceBuilder[T]) JSONOptions(opts ...jsonv2.Options) *sliceBuilder[T] {
	b.jsonBuilder.JSONOptions(opts...)
	return b
}
