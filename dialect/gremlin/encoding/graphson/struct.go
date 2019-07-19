package graphson

import (
	"github.com/json-iterator/go"
)

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
