// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"io"
	"reflect"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

type decodeExtension struct {
	jsoniter.DummyExtension
}

// Unmarshal parses the graphson encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v any) error {
	return config.Unmarshal(data, v)
}

// UnmarshalFromString parses the graphson encoded str and stores the result
// in the value pointed to by v.
func UnmarshalFromString(str string, v any) error {
	return config.UnmarshalFromString(str, v)
}

// Decoder defines a graphson decoder.
type Decoder interface {
	Decode(any) error
}

// NewDecoder create a graphson decoder.
func NewDecoder(r io.Reader) Decoder {
	return config.NewDecoder(r)
}

// Unmarshaler is the interface implemented by types
// that can unmarshal a graphson description of themselves.
type Unmarshaler interface {
	UnmarshalGraphson([]byte) error
}

// UpdateStructDescriptor decorates struct field encoders for graphson tags.
func (ext decodeExtension) UpdateStructDescriptor(desc *jsoniter.StructDescriptor) {
	for _, binding := range desc.Fields {
		if tag, ok := binding.Field.Tag().Lookup("graphson"); ok && tag != "-" {
			if dec := ext.DecoratorOfStructField(binding.Decoder, tag); dec != nil {
				binding.Decoder = dec
			}
		}
	}
}

// CreateDecoder returns a value decoder for type.
func (ext decodeExtension) CreateDecoder(typ reflect2.Type) jsoniter.ValDecoder {
	if dec := ext.DecoderOfRegistered(typ); dec != nil {
		return dec
	}
	if dec := ext.DecoderOfUnmarshaler(typ); dec != nil {
		return dec
	}
	if dec := ext.DecoderOfNative(typ); dec != nil {
		return dec
	}
	switch typ.Kind() {
	case reflect.Array:
		return ext.DecoderOfArray(typ)
	case reflect.Slice:
		return ext.DecoderOfSlice(typ)
	case reflect.Map:
		return ext.DecoderOfMap(typ)
	default:
		return nil
	}
}

// DecorateDecoder decorates an passed in value decoder for type.
func (ext decodeExtension) DecorateDecoder(typ reflect2.Type, dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	if dec := ext.DecoratorOfRegistered(dec); dec != nil {
		return dec
	}
	if dec := ext.DecoratorOfUnmarshaler(typ, dec); dec != nil {
		return dec
	}
	if dec := ext.DecoratorOfTyper(typ, dec); dec != nil {
		return dec
	}
	if dec := ext.DecoratorOfNative(typ, dec); dec != nil {
		return dec
	}
	switch typ.Kind() {
	case reflect.Ptr, reflect.Struct:
		return dec
	case reflect.Interface:
		return ext.DecoratorOfInterface(typ, dec)
	case reflect.Slice:
		return ext.DecoratorOfSlice(typ, dec)
	case reflect.Array:
		return ext.DecoratorOfArray(dec)
	case reflect.Map:
		return ext.DecoratorOfMap(dec)
	default:
		return ext.DecoderOfError("graphson: unsupported type: %s", typ.String())
	}
}
