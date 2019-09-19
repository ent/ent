package field

import "strings"

type Type uint8

const (
	TypeInvalid Type = iota
	TypeBool
	TypeTime
	TypeJSON
	TypeBytes
	TypeString
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

// ConstName returns the constant name of a info type.
// It's used by entc for printing the constant name in templates.
func (t Type) ConstName() string {
	switch t {
	case TypeJSON:
		return "TypeJSON"
	case TypeTime:
		return "TypeTime"
	case TypeBytes:
		return "TypeBytes"
	default:
		return "Type" + strings.Title(typeNames[t])
	}
}

type TypeInfo struct {
	Type     Type
	Ident    string
	PkgPath  string
	Nillable bool // slices or pointers.
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
	return t.Type > TypeInvalid && t.Type < endTypes
}

// Numeric reports if the given type is a numeric type.
func (t TypeInfo) Numeric() bool {
	return t.Type.Numeric()
}

var typeNames = [...]string{
	TypeInvalid: "invalid",
	TypeBool:    "bool",
	TypeTime:    "time.Time",
	TypeJSON:    "json.RawMessage",
	TypeBytes:   "[]byte",
	TypeString:  "string",
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
