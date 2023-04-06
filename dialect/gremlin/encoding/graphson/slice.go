// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// DecoratorOfSlice decorates a value encoder of a slice type.
func (encodeExtension) DecoratorOfSlice(typ reflect2.Type, enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	encoder := typeEncoder{ValEncoder: enc}
	sliceType := typ.(reflect2.SliceType)
	if sliceType.Elem().Kind() == reflect.Uint8 {
		encoder.Type = byteBufferType
	} else {
		encoder.Type = listType
	}
	return sliceEncoder{sliceType, encoder}
}

// DecoratorOfArray decorates a value encoder of an array type.
func (encodeExtension) DecoratorOfArray(enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	return typeEncoder{enc, listType}
}

// DecoderOfSlice returns a value decoder of a slice type.
func (ext decodeExtension) DecoderOfSlice(typ reflect2.Type) jsoniter.ValDecoder {
	sliceType := typ.(reflect2.SliceType)
	elemType := sliceType.Elem()
	if elemType.Kind() == reflect.Uint8 {
		return nil
	}
	return sliceDecoder{
		sliceType: sliceType,
		elemDec:   ext.LazyDecoderOf(elemType),
	}
}

// DecoderOfArray returns a value decoder of an array type.
func (ext decodeExtension) DecoderOfArray(typ reflect2.Type) jsoniter.ValDecoder {
	arrayType := typ.(reflect2.ArrayType)
	return arrayDecoder{
		arrayType: arrayType,
		elemDec:   ext.LazyDecoderOf(arrayType.Elem()),
	}
}

// DecoratorOfSlice decorates a value decoder of a slice type.
func (ext decodeExtension) DecoratorOfSlice(typ reflect2.Type, dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	if typ.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		return typeDecoder{dec, byteBufferType}
	}
	return typeDecoder{dec, listType}
}

// DecoratorOfArray decorates a value decoder of an array type.
func (ext decodeExtension) DecoratorOfArray(dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	return typeDecoder{dec, listType}
}

type sliceEncoder struct {
	sliceType reflect2.SliceType
	jsoniter.ValEncoder
}

func (enc sliceEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	if enc.sliceType.UnsafeIsNil(ptr) {
		stream.WriteNil()
	} else {
		enc.ValEncoder.Encode(ptr, stream)
	}
}

type sliceDecoder struct {
	sliceType reflect2.SliceType
	elemDec   jsoniter.ValDecoder
}

func (dec sliceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	dec.decode(ptr, iter)
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("decoding slice %s: %w", dec.sliceType, iter.Error)
	}
}

func (dec sliceDecoder) decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	sliceType := dec.sliceType
	if iter.ReadNil() {
		sliceType.UnsafeSetNil(ptr)
		return
	}

	sliceType.UnsafeSet(ptr, sliceType.UnsafeMakeSlice(0, 0))
	var length int

	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		idx := length
		length++
		sliceType.UnsafeGrow(ptr, length)
		elem := sliceType.UnsafeGetIndex(ptr, idx)
		dec.elemDec.Decode(elem, iter)
		return iter.Error == nil
	})
}

type arrayDecoder struct {
	arrayType reflect2.ArrayType
	elemDec   jsoniter.ValDecoder
}

func (dec arrayDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	dec.decode(ptr, iter)
	if iter.Error != nil && iter.Error != io.EOF {
		iter.Error = fmt.Errorf("decoding array %s: %w", dec.arrayType, iter.Error)
	}
}

func (dec arrayDecoder) decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var (
		arrayType = dec.arrayType
		length    int
	)
	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		if length < arrayType.Len() {
			idx := length
			length++
			elem := arrayType.UnsafeGetIndex(ptr, idx)
			dec.elemDec.Decode(elem, iter)
		} else {
			iter.Skip()
		}
		return iter.Error == nil
	})
}
