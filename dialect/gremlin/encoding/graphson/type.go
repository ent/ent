// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// A Type is a graphson type.
type Type string

// graphson typed value types.
const (
	// core
	doubleType Type = "g:Double"
	floatType  Type = "g:Float"
	int32Type  Type = "g:Int32"
	int64Type  Type = "g:Int64"
	listType   Type = "g:List"
	mapType    Type = "g:Map"
	Timestamp  Type = "g:Timestamp"
	Date       Type = "g:Date"

	// extended
	bigIntegerType Type = "gx:BigInteger"
	bigDecimal     Type = "gx:BigDecimal"
	byteType       Type = "gx:Byte"
	byteBufferType Type = "gx:ByteBuffer"
	int16Type      Type = "gx:Int16"
)

// String implements fmt.Stringer interface.
func (typ Type) String() string {
	return string(typ)
}

// CheckType implements typeChecker interface.
func (typ Type) CheckType(other Type) error {
	if typ != other {
		return fmt.Errorf("expect type %s, but found %s", typ, other)
	}
	return nil
}

// Types is a slice of Type.
type Types []Type

// Contains reports whether a slice of types contains a particular type.
func (types Types) Contains(typ Type) bool {
	for i := range types {
		if types[i] == typ {
			return true
		}
	}
	return false
}

// String implements fmt.Stringer interface.
func (types Types) String() string {
	var builder strings.Builder
	builder.WriteByte('[')
	for i := range types {
		if i > 0 {
			builder.WriteByte(',')
		}
		builder.WriteString(types[i].String())
	}
	builder.WriteByte(']')
	return builder.String()
}

// CheckType implements typeChecker interface.
func (types Types) CheckType(typ Type) error {
	if !types.Contains(typ) {
		return fmt.Errorf("expect any of %s, but found %s", types, typ)
	}
	return nil
}

// Typer is the interface implemented by types that
// define an underlying graphson type.
type Typer interface {
	GraphsonType() Type
}

var typerType = reflect2.TypeOfPtr((*Typer)(nil)).Elem()

// DecoratorOfTyper decorates a value encoder of a Typer interface.
func (ext encodeExtension) DecoratorOfTyper(typ reflect2.Type, enc jsoniter.ValEncoder) jsoniter.ValEncoder {
	if typ.Kind() != reflect.Struct {
		return nil
	}
	if typ.Implements(typerType) {
		return typerEncoder{
			typeEncoder: typeEncoder{ValEncoder: enc},
			typerOf: func(ptr unsafe.Pointer) Typer {
				return typ.UnsafeIndirect(ptr).(Typer)
			},
		}
	}
	ptrType := reflect2.PtrTo(typ)
	if ptrType.Implements(typerType) {
		return typerEncoder{
			typeEncoder: typeEncoder{ValEncoder: enc},
			typerOf: func(ptr unsafe.Pointer) Typer {
				// nolint: gas
				return ptrType.UnsafeIndirect(unsafe.Pointer(&ptr)).(Typer)
			},
		}
	}
	return nil
}

// DecoratorOfTyper decorates a value decoder of a Typer interface.
func (ext decodeExtension) DecoratorOfTyper(typ reflect2.Type, dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	ptrType := reflect2.PtrTo(typ)
	if ptrType.Implements(typerType) {
		return typerDecoder{
			typeDecoder: typeDecoder{ValDecoder: dec},
			typerOf: func(ptr unsafe.Pointer) Typer {
				// nolint: gas
				return ptrType.UnsafeIndirect(unsafe.Pointer(&ptr)).(Typer)
			},
		}
	}
	return nil
}

type typerEncoder struct {
	typeEncoder
	typerOf func(unsafe.Pointer) Typer
}

func (enc typerEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	enc.typeEncoder.Type = enc.typerOf(ptr).GraphsonType()
	enc.typeEncoder.Encode(ptr, stream)
}

type typerDecoder struct {
	typeDecoder
	typerOf func(unsafe.Pointer) Typer
}

func (dec typerDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	dec.typeDecoder.typeChecker = dec.typerOf(ptr).GraphsonType()
	dec.typeDecoder.Decode(ptr, iter)
}
