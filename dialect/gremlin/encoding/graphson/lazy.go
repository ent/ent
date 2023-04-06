// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"sync"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// LazyEncoderOf returns a lazy encoder for type.
func (encodeExtension) LazyEncoderOf(typ reflect2.Type) jsoniter.ValEncoder {
	return &lazyEncoder{resolve: func() jsoniter.ValEncoder {
		return config.EncoderOf(typ)
	}}
}

// LazyDecoderOf returns a lazy unique decoder for type.
func (decodeExtension) LazyDecoderOf(typ reflect2.Type) jsoniter.ValDecoder {
	return &lazyDecoder{resolve: func() jsoniter.ValDecoder {
		dec := config.DecoderOf(reflect2.PtrTo(typ))
		if td, ok := dec.(typeDecoder); ok {
			td.typeChecker = &uniqueType{elemChecker: td.typeChecker}
			dec = td
		}
		return dec
	}}
}

type lazyEncoder struct {
	jsoniter.ValEncoder
	resolve func() jsoniter.ValEncoder
	once    sync.Once
}

func (enc *lazyEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	enc.once.Do(func() { enc.ValEncoder = enc.resolve() })
	enc.ValEncoder.Encode(ptr, stream)
}

func (enc *lazyEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	enc.once.Do(func() { enc.ValEncoder = enc.resolve() })
	return enc.ValEncoder.IsEmpty(ptr)
}

type lazyDecoder struct {
	jsoniter.ValDecoder
	resolve func() jsoniter.ValDecoder
	once    sync.Once
}

func (dec *lazyDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	dec.once.Do(func() { dec.ValDecoder = dec.resolve() })
	dec.ValDecoder.Decode(ptr, iter)
}

type uniqueType struct {
	typ         Type
	once        sync.Once
	elemChecker typeChecker
}

func (u *uniqueType) CheckType(other Type) error {
	u.once.Do(func() { u.typ = other })
	if u.typ != other {
		return fmt.Errorf("expect type %s, but found %s", u.typ, other)
	}
	return u.elemChecker.CheckType(u.typ)
}
