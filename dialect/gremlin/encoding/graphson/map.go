// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// EncoderOfMap returns a value encoder of a map type.
func (ext encodeExtension) EncoderOfMap(typ reflect2.Type) jsoniter.ValEncoder {
	mapType := typ.(reflect2.MapType)
	return &mapEncoder{
		mapType: mapType,
		keyEnc:  ext.LazyEncoderOf(mapType.Key()),
		elemEnc: ext.LazyEncoderOf(mapType.Elem()),
	}
}

// DecoratorOfMap decorates a value encoder of a map type.
func (encodeExtension) DecoratorOfMap(enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	return typeEncoder{enc, mapType}
}

type mapEncoder struct {
	mapType reflect2.MapType
	keyEnc  jsoniter.ValEncoder
	elemEnc jsoniter.ValEncoder
}

func (enc *mapEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	iter := enc.mapType.UnsafeIterate(ptr)
	if !iter.HasNext() {
		stream.WriteEmptyArray()
		return
	}

	stream.WriteArrayStart()
	for {
		key, elem := iter.UnsafeNext()
		enc.keyEnc.Encode(key, stream)
		stream.WriteMore()
		enc.elemEnc.Encode(elem, stream)
		if !iter.HasNext() {
			break
		}
		stream.WriteMore()
	}
	stream.WriteArrayEnd()
}

func (enc *mapEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return !enc.mapType.UnsafeIterate(ptr).HasNext()
}

// DecoderOfMap returns a value decoder of a map type.
func (ext decodeExtension) DecoderOfMap(typ reflect2.Type) jsoniter.ValDecoder {
	mapType := typ.(reflect2.MapType)
	keyType, elemType := mapType.Key(), mapType.Elem()
	return &mapDecoder{
		mapType:  mapType,
		keyType:  keyType,
		elemType: elemType,
		keyDec:   ext.LazyDecoderOf(keyType),
		elemDec:  ext.LazyDecoderOf(elemType),
	}
}

// DecoratorOfMap decorates a value decoder of a map type.
func (decodeExtension) DecoratorOfMap(dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	return typeDecoder{dec, mapType}
}

type mapDecoder struct {
	mapType  reflect2.MapType
	keyType  reflect2.Type
	elemType reflect2.Type
	keyDec   jsoniter.ValDecoder
	elemDec  jsoniter.ValDecoder
}

func (dec *mapDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	mapType := dec.mapType
	if mapType.UnsafeIsNil(ptr) {
		mapType.UnsafeSet(ptr, mapType.UnsafeMakeMap(0))
	}

	var key unsafe.Pointer
	if !iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		if key == nil {
			key = dec.keyType.UnsafeNew()
			dec.keyDec.Decode(key, iter)
			return iter.Error == nil
		}

		elem := dec.elemType.UnsafeNew()
		dec.elemDec.Decode(elem, iter)
		if iter.Error != nil {
			return false
		}

		mapType.UnsafeSetIndex(ptr, key, elem)
		key = nil
		return true
	}) {
		return
	}

	if key != nil {
		iter.ReportError("decode map", "odd number of map items")
	}
}
