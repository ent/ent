// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCodec(t *testing.T) {
	codec := errorCodec{errors.New("codec error")}
	assert.False(t, codec.IsEmpty(nil))

	var buf bytes.Buffer
	stream := config.BorrowStream(&buf)
	defer config.ReturnStream(stream)
	codec.Encode(nil, stream)
	assert.Empty(t, buf.Bytes())
	assert.EqualError(t, stream.Error, codec.Error())

	iter := config.BorrowIterator([]byte{})
	defer config.ReturnIterator(iter)
	codec.Decode(nil, iter)
	assert.EqualError(t, iter.Error, codec.Error())
}
