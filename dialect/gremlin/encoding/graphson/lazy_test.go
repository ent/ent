// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"sync/atomic"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLazyEncode(t *testing.T) {
	var m mocker
	m.On("IsEmpty", mock.Anything).Return(false).Once()
	m.On("Encode", mock.Anything, mock.Anything).Once()
	defer m.AssertExpectations(t)

	var cnt uint32
	var enc jsoniter.ValEncoder = &lazyEncoder{resolve: func() jsoniter.ValEncoder {
		assert.Equal(t, uint32(1), atomic.AddUint32(&cnt, 1))
		return &m
	}}

	enc.IsEmpty(nil)
	enc.Encode(nil, nil)
}

func TestLazyDecode(t *testing.T) {
	var m mocker
	m.On("Decode", mock.Anything, mock.Anything).Times(3)
	defer m.AssertExpectations(t)

	var cnt uint32
	var dec jsoniter.ValDecoder = &lazyDecoder{resolve: func() jsoniter.ValDecoder {
		assert.Equal(t, uint32(1), atomic.AddUint32(&cnt, 1))
		return &m
	}}

	for i := 0; i < 3; i++ {
		dec.Decode(nil, nil)
	}
}
