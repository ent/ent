// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// EncoderOfNative returns a value encoder of a native type.
func (encodeExtension) EncoderOfNative(typ reflect2.Type) jsoniter.ValEncoder {
	switch typ.Kind() {
	case reflect.Float64:
		return float64Encoder{typ}
	default:
		return nil
	}
}

// DecoratorOfNative decorates a value encoder of a native type.
func (encodeExtension) DecoratorOfNative(typ reflect2.Type, enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	switch typ.Kind() {
	case reflect.Bool, reflect.String:
		return enc
	case reflect.Int64, reflect.Int, reflect.Uint32:
		return typeEncoder{enc, int64Type}
	case reflect.Int32, reflect.Int8, reflect.Uint16:
		return typeEncoder{enc, int32Type}
	case reflect.Int16:
		return typeEncoder{enc, int16Type}
	case reflect.Uint64, reflect.Uint:
		return typeEncoder{enc, bigIntegerType}
	case reflect.Uint8:
		return typeEncoder{enc, byteType}
	case reflect.Float32:
		return typeEncoder{enc, floatType}
	case reflect.Float64:
		return typeEncoder{enc, doubleType}
	default:
		return nil
	}
}

// DecoderOfNative returns a value decoder of a native type.
func (decodeExtension) DecoderOfNative(typ reflect2.Type) jsoniter.ValDecoder {
	switch typ.Kind() {
	case reflect.Float64:
		return float64Decoder{typ}
	default:
		return nil
	}
}

// DecoratorOfNative decorates a value decoder of a native type.
func (decodeExtension) DecoratorOfNative(typ reflect2.Type, dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	switch typ.Kind() {
	case reflect.Bool:
		return dec
	case reflect.String:
		return typeDecoder{dec, typeCheckerFunc(func(Type) error { return nil })}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return typeDecoder{dec, integerTypes}
	case reflect.Float32:
		return typeDecoder{dec, floatTypes}
	case reflect.Float64:
		return typeDecoder{dec, doubleTypes}
	default:
		return nil
	}
}

type float64Encoder struct {
	reflect2.Type
}

func (enc float64Encoder) IsEmpty(ptr unsafe.Pointer) bool {
	return enc.UnsafeIndirect(ptr).(float64) == 0
}

func (enc float64Encoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	f := enc.UnsafeIndirect(ptr).(float64)
	switch {
	case math.IsNaN(f):
		stream.WriteString("NaN")
	case math.IsInf(f, 1):
		stream.WriteString("Infinity")
	case math.IsInf(f, -1):
		stream.WriteString("-Infinity")
	default:
		stream.WriteFloat64(f)
	}
}

type float64Decoder struct {
	reflect2.Type
}

var (
	integerTypes = Types{byteType, int16Type, int32Type, int64Type, bigIntegerType}
	floatTypes   = append(integerTypes, floatType, bigDecimal)
	doubleTypes  = append(floatTypes, doubleType)
)

func (dec float64Decoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	var val float64
	switch next := iter.WhatIsNext(); next {
	case jsoniter.NumberValue:
		val = iter.ReadFloat64()
	case jsoniter.StringValue:
		switch str := iter.ReadString(); str {
		case "NaN":
			val = math.NaN()
		case "Infinity":
			val = math.Inf(1)
		case "-Infinity":
			val = math.Inf(-1)
		default:
			iter.ReportError("decode float64", "invalid value "+str)
		}
	default:
		iter.ReportError("decode float64", fmt.Sprintf("unexpected value type: %d", next))
	}

	if iter.Error == nil || iter.Error == io.EOF {
		// nolint: gas
		dec.UnsafeSet(ptr, unsafe.Pointer(&val))
	}
}
