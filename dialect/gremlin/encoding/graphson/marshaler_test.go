// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMarshalerEncode(t *testing.T) {
	want := []byte(`{"@type": "g:Int32", "@value": 42}`)
	m := &mocker{}
	call := m.On("MarshalGraphson").Return(want, nil)
	defer m.AssertExpectations(t)

	tests := []any{m, &m, func() *Marshaler { marshaler := Marshaler(m); return &marshaler }(), Marshaler(nil)}
	call.Times(len(tests) - 1)

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc), func(t *testing.T) {
			got, err := Marshal(tc)
			assert.NoError(t, err)
			if !reflect2.IsNil(tc) {
				assert.Equal(t, want, got)
			} else {
				assert.Equal(t, []byte("null"), got)
			}
		})
	}
}

func TestMarshalerError(t *testing.T) {
	errStr := "marshaler error"
	m := &mocker{}
	m.On("MarshalGraphson").Return(nil, errors.New(errStr)).Once()
	defer m.AssertExpectations(t)

	_, err := Marshal(m)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), errStr)
}

func TestBadMarshaler(t *testing.T) {
	m := &mocker{}
	m.On("MarshalGraphson").Return([]byte(`{"@type": "g:Int32", "@value":`), nil).Once()
	defer m.AssertExpectations(t)
	_, err := Marshal(m)
	assert.Error(t, err)
}

func TestUnmarshalerDecode(t *testing.T) {
	data := `{"@type": "g:UUID", "@value": "cb682578-9d92-4499-9ebc-5c6aa73c5397"}`
	var value string

	m := &mocker{}
	m.On("UnmarshalGraphson", mock.Anything).
		Run(func(args mock.Arguments) {
			data := args.Get(0).([]byte)
			value = jsoniter.Get(data, "@value").ToString()
		}).
		Return(nil).
		Once()
	defer m.AssertExpectations(t)

	err := UnmarshalFromString(data, m)
	require.NoError(t, err)
	assert.Equal(t, "cb682578-9d92-4499-9ebc-5c6aa73c5397", value)
}

func TestUnmarshalerError(t *testing.T) {
	errStr := "unmarshaler error"
	m := &mocker{}
	m.On("UnmarshalGraphson", mock.Anything).Return(errors.New(errStr)).Once()
	defer m.AssertExpectations(t)

	err := Unmarshal([]byte(`{}`), m)
	require.Error(t, err)
	assert.Contains(t, err.Error(),
		fmt.Sprintf("graphson: error calling UnmarshalGraphson for type %s: %s",
			reflect.TypeOf(m), errStr,
		),
	)
}

func TestUnmarshalBadInput(t *testing.T) {
	m := &mocker{}
	defer m.AssertExpectations(t)
	err := UnmarshalFromString(`{"@type"}`, m)
	assert.Error(t, err)
}
