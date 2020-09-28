// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// EncoderOfError returns a value encoder which always fails to encode.
func (encodeExtension) EncoderOfError(format string, args ...interface{}) jsoniter.ValEncoder {
	return decoratorOfError(format, args...)
}

// DecoderOfError returns a value decoder which always fails to decode.
func (decodeExtension) DecoderOfError(format string, args ...interface{}) jsoniter.ValDecoder {
	return decoratorOfError(format, args...)
}

func decoratorOfError(format string, args ...interface{}) errorCodec {
	err := errors.Errorf(format, args...)
	return errorCodec{err}
}

type errorCodec struct{ error }

func (ec errorCodec) Encode(_ unsafe.Pointer, stream *jsoniter.Stream) {
	if stream.Error == nil {
		stream.Error = ec.error
	}
}

func (errorCodec) IsEmpty(unsafe.Pointer) bool {
	return false
}

func (ec errorCodec) Decode(_ unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.Error == nil {
		iter.Error = ec.error
	}
}
