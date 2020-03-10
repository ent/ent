// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import jsoniter "github.com/json-iterator/go"

// DecoratorOfStructField decorates a struct field value encoder.
func (encodeExtension) DecoratorOfStructField(enc jsoniter.ValEncoder, tag string) jsoniter.ValEncoder {
	typ, _ := parseTag(tag)
	if typ == "" {
		return nil
	}

	encoder, ok := enc.(typeEncoder)
	if !ok {
		encoder = typeEncoder{ValEncoder: enc}
	}
	encoder.Type = Type(typ)

	return encoder
}

// DecoratorOfStructField decorates a struct field value decoder.
func (decodeExtension) DecoratorOfStructField(dec jsoniter.ValDecoder, tag string) jsoniter.ValDecoder {
	typ, _ := parseTag(tag)
	if typ == "" {
		return nil
	}

	decoder, ok := dec.(typeDecoder)
	if !ok {
		decoder = typeDecoder{ValDecoder: dec}
	}
	decoder.typeChecker = Type(typ)

	return decoder
}
