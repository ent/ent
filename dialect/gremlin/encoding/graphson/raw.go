// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"errors"
)

// RawMessage is a raw encoded graphson value.
type RawMessage []byte

// RawMessage must implement Marshaler/Unmarshaler interfaces.
var (
	_ Marshaler   = (*RawMessage)(nil)
	_ Unmarshaler = (*RawMessage)(nil)
)

// MarshalGraphson returns m as the graphson encoding of m.
func (m RawMessage) MarshalGraphson() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalGraphson sets *m to a copy of data.
func (m *RawMessage) UnmarshalGraphson(data []byte) error {
	if m == nil {
		return errors.New("graphson.RawMessage: UnmarshalGraphson on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
