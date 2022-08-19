// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTypeCheckType(t *testing.T) {
	assert.NoError(t, int32Type.CheckType(int32Type))
	assert.Error(t, int32Type.CheckType(int64Type))
}

func TestTypesCheckType(t *testing.T) {
	assert.NoError(t, Types{int16Type, int32Type, int64Type}.CheckType(int32Type))
	assert.Error(t, Types{floatType, doubleType}.CheckType(bigIntegerType))
}

func TestTypesString(t *testing.T) {
	assert.Equal(t, "[]", Types{}.String())
	assert.Equal(t, "[gx:Byte]", Types{byteType}.String())
	assert.Equal(t, "[gx:Int16,g:Int32,g:Int64]", Types{int16Type, int32Type, int64Type}.String())
}

type vertex struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

func (vertex) GraphsonType() Type {
	return Type("g:Vertex")
}

type mockVertex struct {
	mock.Mock `json:"-"`
	ID        int    `json:"id"`
	Label     string `json:"label"`
}

func (m *mockVertex) GraphsonType() Type {
	return m.Called().Get(0).(Type)
}

func TestEncodeTyper(t *testing.T) {
	m := &mockVertex{ID: 42, Label: "person"}
	m.On("GraphsonType").Return(Type("g:Vertex")).Twice()
	defer m.AssertExpectations(t)

	v := vertex{ID: m.ID, Label: m.Label}
	var vv Typer = v

	want := `{
	   "@type": "g:Vertex",
	   "@value": {
		   "id": {
			   "@type": "g:Int64",
			   "@value": 42
		   },
		   "label": "person"
	   }
	}`

	for _, tc := range []any{m, &m, v, vv, &vv} {
		got, err := MarshalToString(tc)
		assert.NoError(t, err)
		assert.JSONEq(t, want, got)
	}
}

func TestDecodeTyper(t *testing.T) {
	var m mockVertex
	m.On("GraphsonType").Return(Type("g:Vertex")).Once()
	defer m.AssertExpectations(t)

	in := `{
	   "@type": "g:Vertex",
	   "@value": {
		   "id": {
			   "@type": "g:Int64",
			   "@value": 55
		   },
		   "label": "user"
	   }
	}`

	err := UnmarshalFromString(in, &m)
	assert.NoError(t, err)
	assert.Equal(t, 55, m.ID)
	assert.Equal(t, "user", m.Label)
}
