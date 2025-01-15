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

type encodeExtension struct {
	jsoniter.DummyExtension
}

// Marshal returns the graphson encoding of v.
func Marshal(v any) ([]byte, error) {
	return config.Marshal(v)
}

// MarshalToString returns the graphson encoding of v as string.
func MarshalToString(v any) (string, error) {
	return config.MarshalToString(v)
}

// Encoder defines a graphson encoder.
type Encoder interface {
	Encode(any) error
}

// NewEncoder create a graphson encoder.
func NewEncoder(w io.Writer) Encoder {
	return config.NewEncoder(w)
}

// Marshaler is the interface implemented by types that
// can marshal themselves as graphson.
type Marshaler interface {
	MarshalGraphson() ([]byte, error)
}

// UpdateStructDescriptor decorates struct field encoders for graphson tags.
func (ext encodeExtension) UpdateStructDescriptor(desc *jsoniter.StructDescriptor) {
	for _, binding := range desc.Fields {
		if tag, ok := binding.Field.Tag().Lookup("graphson"); ok && tag != "-" {
			if enc := ext.DecoratorOfStructField(binding.Encoder, tag); enc != nil {
				binding.Encoder = enc
			}
		}
	}
}

// CreateEncoder returns a value encoder for type.
func (ext encodeExtension) CreateEncoder(typ reflect2.Type) jsoniter.ValEncoder {
	if enc := ext.EncoderOfRegistered(typ); enc != nil {
		return enc
	}
	if enc := ext.EncoderOfNative(typ); enc != nil {
		return enc
	}
	switch typ.Kind() {
	case reflect.Map:
		return ext.EncoderOfMap(typ)
	default:
		return nil
	}
}

// DecorateEncoder decorates an passed in value encoder for type.
func (ext encodeExtension) DecorateEncoder(typ reflect2.Type, enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	if enc := ext.DecoratorOfRegistered(enc); enc != nil {
		return enc
	}
	if enc := ext.DecoratorOfMarshaler(typ, enc); enc != nil {
		return enc
	}
	if enc := ext.DecoratorOfTyper(typ, enc); enc != nil {
		return enc
	}
	if enc := ext.DecoratorOfNative(typ, enc); enc != nil {
		return enc
	}
	switch typ.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Struct:
		return enc
	case reflect.Array:
		return ext.DecoratorOfArray(enc)
	case reflect.Slice:
		return ext.DecoratorOfSlice(typ, enc)
	case reflect.Map:
		return ext.DecoratorOfMap(enc)
	default:
		return ext.EncoderOfError("graphson: unsupported type: %s", typ.String())
	}
}
