// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTypeEncode(t *testing.T) {
	var got bytes.Buffer
	stream := config.BorrowStream(&got)
	defer config.ReturnStream(stream)

	typ, val := int32Type, 42
	ptr := unsafe.Pointer(&val)

	var m mocker
	m.On("Encode", ptr, stream).
		Run(func(args mock.Arguments) {
			stream := args.Get(1).(*jsoniter.Stream)
			stream.WriteInt(val)
		}).
		Once()
	defer m.AssertExpectations(t)

	typeEncoder{&m, typ}.Encode(ptr, stream)
	require.NoError(t, stream.Flush())
	want := fmt.Sprintf(`{"@type": "%s", "@value": %d}`, typ, val)
	assert.JSONEq(t, want, got.String())
}

func TestTypeDecode(t *testing.T) {
	typ, val := int64Type, 84
	ptr := unsafe.Pointer(&val)

	data := fmt.Sprintf(`{"@value": %d, "@type": "%s"}`, val, typ)
	iter := config.BorrowIterator([]byte(data))
	defer config.ReturnIterator(iter)

	m := &mocker{}
	m.On("CheckType", typ).
		Return(nil).
		Once()
	m.On("Decode", ptr, mock.Anything).
		Run(func(args mock.Arguments) {
			iter := args.Get(1).(*jsoniter.Iterator)
			assert.Equal(t, val, iter.ReadInt())
		}).
		Once()
	defer m.AssertExpectations(t)

	typeDecoder{m, m}.Decode(ptr, iter)
	assert.NoError(t, iter.Error)
}

func TestTypeDecodeBadType(t *testing.T) {
	typ, val := int64Type, 55
	ptr := unsafe.Pointer(&val)

	m := &mocker{}
	m.On("CheckType", typ).Return(errors.New("bad type")).Once()
	defer m.AssertExpectations(t)

	data := fmt.Sprintf(`{"@type": "%s", "@value": %d}`, typ, val)
	iter := config.BorrowIterator([]byte(data))
	defer config.ReturnIterator(iter)

	typeDecoder{m, m}.Decode(ptr, iter)
	require.Error(t, iter.Error)
	assert.Contains(t, iter.Error.Error(), "bad type")
}

func TestTypeDecodeDuplicateField(t *testing.T) {
	data := `{"@type": "gx:Byte", "@value": 33, "@type": "g:Int32"}`
	iter := config.BorrowIterator([]byte(data))
	defer config.ReturnIterator(iter)
	var ptr unsafe.Pointer

	m := &mocker{}
	m.On("CheckType", mock.MatchedBy(func(typ Type) bool { return typ == int32Type })).
		Return(nil).
		Once()
	m.On("Decode", ptr, mock.Anything).
		Run(func(args mock.Arguments) {
			args.Get(1).(*jsoniter.Iterator).Skip()
			require.NoError(t, iter.Error)
		}).
		Once()
	defer m.AssertExpectations(t)

	typeDecoder{m, m}.Decode(ptr, iter)
	assert.NoError(t, iter.Error)
}

func TestTypeDecodeMissingField(t *testing.T) {
	data := `{"@type": "g:Int32"}`
	iter := config.BorrowIterator([]byte(data))
	defer config.ReturnIterator(iter)

	m := &mocker{}
	defer m.AssertExpectations(t)

	typeDecoder{m, m}.Decode(nil, iter)
	require.Error(t, iter.Error)
	assert.Contains(t, iter.Error.Error(), "missing type or value")
}

func TestTypeDecodeSyntaxError(t *testing.T) {
	data := `{"@type": "gx:Int16", "@value", 65000}`
	iter := config.BorrowIterator([]byte(data))
	defer config.ReturnIterator(iter)

	m := &mocker{}
	defer m.AssertExpectations(t)

	typeDecoder{m, m}.Decode(nil, iter)
	assert.Error(t, iter.Error)
}
