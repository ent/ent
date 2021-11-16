// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

var (
	typeEncoders = map[string]jsoniter.ValEncoder{}
	typeDecoders = map[string]jsoniter.ValDecoder{}
)

// RegisterTypeEncoder register type encoder for typ.
func RegisterTypeEncoder(typ string, enc jsoniter.ValEncoder) {
	typeEncoders[typ] = enc
}

// RegisterTypeDecoder register type decoder for typ.
func RegisterTypeDecoder(typ string, dec jsoniter.ValDecoder) {
	typeDecoders[typ] = dec
}

type registeredEncoder struct{ jsoniter.ValEncoder }

// EncoderOfRegistered returns a value encoder of a registered type.
func (encodeExtension) EncoderOfRegistered(typ reflect2.Type) jsoniter.ValEncoder {
	enc := typeEncoders[typ.String()]
	if enc != nil {
		return registeredEncoder{enc}
	}
	if typ.Kind() == reflect.Ptr {
		ptrType := typ.(reflect2.PtrType)
		enc := typeEncoders[ptrType.Elem().String()]
		if enc != nil {
			return registeredEncoder{
				ValEncoder: &jsoniter.OptionalEncoder{
					ValueEncoder: enc,
				},
			}
		}
	}
	return nil
}

// DecoratorOfRegistered decorates a value encoder of a registered type.
func (encodeExtension) DecoratorOfRegistered(enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	if _, ok := enc.(registeredEncoder); ok {
		return enc
	}
	return nil
}

type registeredDecoder struct{ jsoniter.ValDecoder }

// DecoderOfRegistered returns a value decoder of a registered type.
func (decodeExtension) DecoderOfRegistered(typ reflect2.Type) jsoniter.ValDecoder {
	dec := typeDecoders[typ.String()]
	if dec != nil {
		return registeredDecoder{dec}
	}
	if typ.Kind() == reflect.Ptr {
		ptrType := typ.(reflect2.PtrType)
		dec := typeDecoders[ptrType.Elem().String()]
		if dec != nil {
			return registeredDecoder{
				ValDecoder: &jsoniter.OptionalDecoder{
					ValueType:    ptrType.Elem(),
					ValueDecoder: dec,
				},
			}
		}
	}
	return nil
}

// DecoratorOfRegistered decorates a value decoder of a registered type.
func (decodeExtension) DecoratorOfRegistered(dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	if _, ok := dec.(registeredDecoder); ok {
		return dec
	}
	return nil
}
