// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"errors"
	"io"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

// graphson encoding type / value keys
const (
	TypeKey  = "@type"
	ValueKey = "@value"
)

// typeEncoder adds graphson type information to a value encoder.
type typeEncoder struct {
	jsoniter.ValEncoder
	Type Type
}

// Encode belongs to jsoniter.ValEncoder interface.
func (enc typeEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteObjectStart()
	stream.WriteObjectField(TypeKey)
	stream.WriteString(enc.Type.String())
	stream.WriteMore()
	stream.WriteObjectField(ValueKey)
	enc.ValEncoder.Encode(ptr, stream)
	stream.WriteObjectEnd()
}

type (
	// typeDecoder decorates a value decoder and adds graphson type verification.
	typeDecoder struct {
		jsoniter.ValDecoder
		typeChecker
	}

	// typeChecker defines an interface for graphson type verification.
	typeChecker interface {
		CheckType(Type) error
	}

	// typeCheckerFunc allows the use of functions as type checkers.
	typeCheckerFunc func(Type) error

	// typeValue defines a graphson type / value pair.
	typeValue struct {
		Type  Type
		Value jsoniter.RawMessage
	}
)

// Decode belongs to jsoniter.ValDecoder interface.
func (dec typeDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.WhatIsNext() != jsoniter.ObjectValue {
		dec.ValDecoder.Decode(ptr, iter)
		return
	}

	data := iter.SkipAndReturnBytes()
	if iter.Error != nil && iter.Error != io.EOF {
		return
	}

	var tv typeValue
	if err := jsoniter.Unmarshal(data, &tv); err != nil {
		iter.ReportError("unmarshal type value", err.Error())
		return
	}

	if err := dec.CheckType(tv.Type); err != nil {
		iter.ReportError("check type", err.Error())
		return
	}

	it := config.BorrowIterator(tv.Value)
	defer config.ReturnIterator(it)

	dec.ValDecoder.Decode(ptr, it)
	if it.Error != nil && it.Error != io.EOF {
		iter.ReportError("decode value", it.Error.Error())
	}
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (tv *typeValue) UnmarshalJSON(data []byte) error {
	var v struct {
		Type  *Type               `json:"@type"`
		Value jsoniter.RawMessage `json:"@value"`
	}

	if err := jsoniter.Unmarshal(data, &v); err != nil {
		return err
	}
	if v.Type == nil || v.Value == nil {
		return errors.New("missing type or value")
	}

	tv.Type = *v.Type
	tv.Value = v.Value
	return nil
}

// CheckType implements typeChecker interface.
func (f typeCheckerFunc) CheckType(typ Type) error {
	return f(typ)
}
