// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field

import (
	"fmt"
	"reflect"
	"strings"
)

// A Type represents a field type.
type Type uint8

// List of field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeTime
	TypeJSON
	TypeUUID
	TypeBytes
	TypeEnum
	TypeString
	TypeOther
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt
	TypeInt64
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint
	TypeUint64
	TypeFloat32
	TypeFloat64
	endTypes
)

// String returns the string representation of a type.
func (t Type) String() string {
	if t < endTypes {
		return typeNames[t]
	}
	return typeNames[TypeInvalid]
}

// Numeric reports if the given type is a numeric type.
func (t Type) Numeric() bool {
	return t >= TypeInt8 && t < endTypes
}

// Float reports if the given type is a float type.
func (t Type) Float() bool {
	return t == TypeFloat32 || t == TypeFloat64
}

// Integer reports if the given type is an integral type.
func (t Type) Integer() bool {
	return t.Numeric() && !t.Float()
}

// Valid reports if the given type if known type.
func (t Type) Valid() bool {
	return t > TypeInvalid && t < endTypes
}

// ConstName returns the constant name of an info type.
// It's used by entc for printing the constant name in templates.
func (t Type) ConstName() string {
	switch {
	case !t.Valid():
		return typeNames[TypeInvalid]
	case int(t) < len(constNames) && constNames[t] != "":
		return constNames[t]
	default:
		return "Type" + strings.Title(typeNames[t])
	}
}

// TypeInfo holds the information regarding field type.
// Used by complex types like JSON and  Bytes.
type TypeInfo struct {
	Type     Type
	Ident    string
	PkgPath  string // import path.
	PkgName  string // local package name.
	Nillable bool   // slices or pointers.
	RType    *RType
}

// String returns the string representation of a type.
func (t TypeInfo) String() string {
	switch {
	case t.Ident != "":
		return t.Ident
	case t.Type < endTypes:
		return typeNames[t.Type]
	default:
		return typeNames[TypeInvalid]
	}
}

// Valid reports if the given type if known type.
func (t TypeInfo) Valid() bool {
	return t.Type.Valid()
}

// Numeric reports if the given type is a numeric type.
func (t TypeInfo) Numeric() bool {
	return t.Type.Numeric()
}

// ConstName returns the const name of the info type.
func (t TypeInfo) ConstName() string {
	return t.Type.ConstName()
}

// ValueScanner indicates if this type implements the ValueScanner interface.
func (t TypeInfo) ValueScanner() bool {
	return t.RType.Implements(valueScannerType)
}

// Validator indicates if this type implements the Validator interface.
func (t TypeInfo) Validator() bool {
	return t.RType.Implements(validatorType)
}

// Valuer indicates if this type implements the driver.Valuer interface.
func (t TypeInfo) Valuer() bool {
	return t.RType.Implements(valuerType)
}

// Comparable reports whether values of this type are comparable.
func (t TypeInfo) Comparable() bool {
	switch t.Type {
	case TypeBool, TypeTime, TypeUUID, TypeEnum, TypeString:
		return true
	case TypeOther:
		// Always accept custom types as comparable on the database side.
		// In the future, we should consider adding an interface to let
		// custom types tell if they are comparable or not (see #1304).
		return true
	default:
		return t.Numeric()
	}
}

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

// Stringer indicates if this type implements the Stringer interface.
func (t TypeInfo) Stringer() bool {
	return t.RType.Implements(stringerType)
}

var (
	typeNames = [...]string{
		TypeInvalid: "invalid",
		TypeBool:    "bool",
		TypeTime:    "time.Time",
		TypeJSON:    "json.RawMessage",
		TypeUUID:    "[16]byte",
		TypeBytes:   "[]byte",
		TypeEnum:    "string",
		TypeString:  "string",
		TypeOther:   "other",
		TypeInt:     "int",
		TypeInt8:    "int8",
		TypeInt16:   "int16",
		TypeInt32:   "int32",
		TypeInt64:   "int64",
		TypeUint:    "uint",
		TypeUint8:   "uint8",
		TypeUint16:  "uint16",
		TypeUint32:  "uint32",
		TypeUint64:  "uint64",
		TypeFloat32: "float32",
		TypeFloat64: "float64",
	}
	constNames = [...]string{
		TypeJSON:  "TypeJSON",
		TypeUUID:  "TypeUUID",
		TypeTime:  "TypeTime",
		TypeEnum:  "TypeEnum",
		TypeBytes: "TypeBytes",
		TypeOther: "TypeOther",
	}
)

// RType holds a serializable reflect.Type information of
// Go object. Used by the entc package.
type RType struct {
	Name    string // reflect.Type.Name
	Ident   string // reflect.Type.String
	Kind    reflect.Kind
	PkgPath string
	Methods map[string]struct{ In, Out []*RType }
	// Used only for in-package checks.
	rtype reflect.Type
}

// TypeEqual reports if the underlying type is equal to the RType (after pointer indirections).
func (r *RType) TypeEqual(t reflect.Type) bool {
	tv := indirect(t)
	return r.Name == tv.Name() && r.Kind == t.Kind() && r.PkgPath == tv.PkgPath()
}

// RType returns the string value of the indirect reflect.Type.
func (r *RType) String() string {
	if r.rtype != nil {
		return r.rtype.String()
	}
	return r.Ident
}

// IsPtr reports if the reflect-type is a pointer type.
func (r *RType) IsPtr() bool {
	return r != nil && r.Kind == reflect.Ptr
}

// Implements reports whether the RType ~implements the given interface type.
func (r *RType) Implements(typ reflect.Type) bool {
	if r == nil {
		return false
	}
	n := typ.NumMethod()
	for i := 0; i < n; i++ {
		m0 := typ.Method(i)
		m1, ok := r.Methods[m0.Name]
		if !ok || len(m1.In) != m0.Type.NumIn() || len(m1.Out) != m0.Type.NumOut() {
			return false
		}
		in := m0.Type.NumIn()
		for j := 0; j < in; j++ {
			if !m1.In[j].TypeEqual(m0.Type.In(j)) {
				return false
			}
		}
		out := m0.Type.NumOut()
		for j := 0; j < out; j++ {
			if !m1.Out[j].TypeEqual(m0.Type.Out(j)) {
				return false
			}
		}
	}
	return true
}
