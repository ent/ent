// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

// DecoratorOfInterface decorates a value decoder of an interface type.
func (decodeExtension) DecoratorOfInterface(typ reflect2.Type, dec jsoniter.ValDecoder) jsoniter.ValDecoder {
	if _, ok := typ.(*reflect2.UnsafeEFaceType); ok {
		return efaceDecoder{typ, dec}
	}
	return dec
}

type efaceDecoder struct {
	typ reflect2.Type
	jsoniter.ValDecoder
}

func (dec efaceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch next := iter.WhatIsNext(); next {
	case jsoniter.StringValue, jsoniter.BoolValue, jsoniter.NilValue:
		dec.ValDecoder.Decode(ptr, iter)
	case jsoniter.ObjectValue:
		dec.decode(ptr, iter)
	default:
		iter.ReportError("decode empty interface", fmt.Sprintf("unexpected value type: %d", next))
	}
}

func (dec efaceDecoder) decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	data := iter.SkipAndReturnBytes()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}

	rtype, err := dec.reflectBytes(data)
	if err != nil {
		iter.ReportError("decode empty interface", err.Error())
		return
	}

	it := config.BorrowIterator(data)
	defer config.ReturnIterator(it)

	var val any
	if rtype != nil {
		val = rtype.New()
		it.ReadVal(val)
		val = rtype.Indirect(val)
	} else {
		if jsoniter.Get(data, TypeKey).LastError() == nil {
			vk := jsoniter.Get(data, ValueKey)
			if vk.LastError() == nil {
				val = vk.GetInterface()
			}
		}
		if val == nil {
			val = it.Read()
		}
	}

	if it.Error != nil && it.Error != io.EOF {
		iter.ReportError("decode empty interface", it.Error.Error())
		return
	}

	// nolint: gas
	dec.typ.UnsafeSet(ptr, unsafe.Pointer(&val))
}

func (dec efaceDecoder) reflectBytes(data []byte) (reflect2.Type, error) {
	typ := Type(jsoniter.Get(data, TypeKey).ToString())
	rtype := dec.reflectType(typ)
	if rtype != nil {
		return rtype, nil
	}

	switch typ {
	case listType:
		return dec.reflectSlice(data)
	case mapType:
		return dec.reflectMap(data)
	default:
		return nil, nil
	}
}

func (efaceDecoder) reflectType(typ Type) reflect2.Type {
	switch typ {
	case doubleType:
		return reflect2.TypeOf(float64(0))
	case floatType:
		return reflect2.TypeOf(float32(0))
	case byteType:
		return reflect2.TypeOf(uint8(0))
	case int16Type:
		return reflect2.TypeOf(int16(0))
	case int32Type:
		return reflect2.TypeOf(int32(0))
	case int64Type, bigIntegerType:
		return reflect2.TypeOf(int64(0))
	case byteBufferType:
		return reflect2.TypeOf([]byte{})
	default:
		return nil
	}
}

func (efaceDecoder) reflectSlice(data []byte) (reflect2.Type, error) {
	var elem any
	if err := Unmarshal(data, &[...]*any{&elem}); err != nil {
		return nil, fmt.Errorf("cannot read first list element: %w", err)
	}

	if elem == nil {
		return reflect2.TypeOf([]any{}), nil
	}

	sliceType := reflect.SliceOf(reflect.TypeOf(elem))
	return reflect2.Type2(sliceType), nil
}

func (efaceDecoder) reflectMap(data []byte) (reflect2.Type, error) {
	var key, elem any
	if err := Unmarshal(
		bytes.Replace(data, []byte(mapType), []byte(listType), 1),
		&[...]*any{&key, &elem},
	); err != nil {
		return nil, fmt.Errorf("cannot unmarshal first map item: %w", err)
	}

	if key == nil {
		return reflect2.TypeOf(map[any]any{}), nil
	} else if elem == nil {
		return nil, errors.New("expect map element, but found only key")
	}

	mapType := reflect.MapOf(reflect.TypeOf(key), reflect.TypeOf(elem))
	return reflect2.Type2(mapType), nil
}
