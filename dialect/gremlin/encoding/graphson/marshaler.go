// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// DecoratorOfMarshaler decorates a value encoder of a Marshaler interface.
func (ext encodeExtension) DecoratorOfMarshaler(typ reflect2.Type, enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	if typ == marshalerType {
		enc := marshalerEncoder{enc, typ}
		return directMarshalerEncoder{enc}
	}
	if typ.Implements(marshalerType) {
		return marshalerEncoder{enc, typ}
	}
	ptrType := reflect2.PtrTo(typ)
	if ptrType.Implements(marshalerType) {
		ptrEnc := ext.LazyEncoderOf(ptrType)
		enc := marshalerEncoder{ptrEnc, ptrType}
		return referenceEncoder{enc}
	}
	return nil
}

// DecoderOfUnmarshaler returns a value decoder of an Unmarshaler interface.
func (decodeExtension) DecoderOfUnmarshaler(typ reflect2.Type) jsoniter.ValDecoder {
	ptrType := reflect2.PtrTo(typ)
	if ptrType.Implements(unmarshalerType) {
		return referenceDecoder{
			unmarshalerDecoder{ptrType},
		}
	}
	return nil
}

// DecoratorOfUnmarshaler decorates a value encoder of an Unmarshaler interface.
func (decodeExtension) DecoratorOfUnmarshaler(typ reflect2.Type, dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	if reflect2.PtrTo(typ).Implements(unmarshalerType) {
		return dec
	}
	return nil
}

var (
	marshalerType   = reflect2.TypeOfPtr((*Marshaler)(nil)).Elem()
	unmarshalerType = reflect2.TypeOfPtr((*Unmarshaler)(nil)).Elem()
)

type marshalerEncoder struct {
	jsoniter.ValEncoder
	reflect2.Type
}

func (enc marshalerEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	marshaler := enc.Type.UnsafeIndirect(ptr).(Marshaler)
	enc.encode(marshaler, stream)
}

func (enc marshalerEncoder) encode(marshaler Marshaler, stream *jsoniter.Stream) {
	data, err := marshaler.MarshalGraphson()
	if err != nil {
		stream.Error = fmt.Errorf("graphson: error calling MarshalGraphson for type %s: %w", enc.Type, err)
		return
	}
	if !config.Valid(data) {
		stream.Error = fmt.Errorf("graphson: syntax error when marshaling type %s", enc.Type)
		return
	}
	_, stream.Error = stream.Write(data)
}

type directMarshalerEncoder struct {
	marshalerEncoder
}

func (enc directMarshalerEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	marshaler := *(*Marshaler)(ptr)
	enc.encode(marshaler, stream)
}

type referenceEncoder struct {
	jsoniter.ValEncoder
}

func (enc referenceEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	// nolint: gas
	enc.ValEncoder.Encode(unsafe.Pointer(&ptr), stream)
}

func (enc referenceEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	// nolint: gas
	return enc.ValEncoder.IsEmpty(unsafe.Pointer(&ptr))
}

type unmarshalerDecoder struct {
	reflect2.Type
}

func (dec unmarshalerDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	bytes := iter.SkipAndReturnBytes()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}

	unmarshaler := dec.UnsafeIndirect(ptr).(Unmarshaler)
	if err := unmarshaler.UnmarshalGraphson(bytes); err != nil {
		iter.ReportError(
			"unmarshal graphson",
			fmt.Sprintf(
				"graphson: error calling UnmarshalGraphson for type %s: %s",
				dec.Type, err,
			),
		)
	}
}

type referenceDecoder struct {
	jsoniter.ValDecoder
}

func (dec referenceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	// nolint: gas
	dec.ValDecoder.Decode(unsafe.Pointer(&ptr), iter)
}
