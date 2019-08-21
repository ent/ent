package field

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Type is a field type.
type Type uint

// Field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeTime
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

func (t Type) String() string {
	if int(t) < len(typeNames) {
		return typeNames[t]
	}
	return "type" + strconv.Itoa(int(t))
}

// Valid reports if the given type if known type.
func (t Type) Valid() bool { return t > TypeInvalid && t < endTypes }

// Numeric reports if the given type is a numeric type.
func (t Type) Numeric() bool { return t >= TypeInt && t < endTypes }

// Slice reports if the given type is a slice type.
func (t Type) Slice() bool { return t == TypeBytes }

// ConstName returns the constant name of a type. It's used by entc for printing the constant name in templates.
func (t Type) ConstName() string {
	switch t {
	case TypeTime:
		return "TypeTime"
	case TypeBytes:
		return "TypeBytes"
	default:
		return "Type" + strings.Title(t.String())
	}
}

// Bits returns the size of the type in bits.
// It panics if the type is not numeric type.
func (t Type) Bits() int {
	if !t.Numeric() {
		panic("schema/field: Bits of non-numeric type")
	}
	return bits[t]
}

var (
	typeNames = [...]string{
		TypeInvalid: "invalid",
		TypeBool:    "bool",
		TypeTime:    "time.Time",
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
	bits = [...]int{
		TypeInt:     strconv.IntSize,
		TypeInt8:    8,
		TypeInt16:   16,
		TypeInt32:   32,
		TypeInt64:   64,
		TypeUint:    strconv.IntSize,
		TypeUint8:   8,
		TypeUint16:  16,
		TypeUint32:  32,
		TypeUint64:  64,
		TypeFloat32: 32,
		TypeFloat64: 64,
	}
)

// Field represents a field on a graph vertex.
type Field struct {
	typ        Type
	tag        string
	size       int
	name       string
	charset    string
	unique     bool
	nillable   bool
	optional   bool
	immutable  bool
	value      interface{}
	validators []interface{}
}

// String returns a new Field with type string.
func String(name string) *stringBuilder { return &stringBuilder{Field{typ: TypeString, name: name}} }

// Text returns a new string field without limitation on the size.
// In MySQL, it is the "longtext" type, but in SQLite and Gremlin it has not effect.
func Text(name string) *stringBuilder {
	return &stringBuilder{Field{typ: TypeString, name: name, size: math.MaxInt32}}
}

// Bytes returns a new Field with type bytes/buffer.
// In MySQL and SQLite, it is the "BLOB" type, and it does not support for Gremlin.
func Bytes(name string) *bytesBuilder { return &bytesBuilder{Field{typ: TypeBytes, name: name}} }

// Bool returns a new Field with type bool.
func Bool(name string) *boolBuilder { return &boolBuilder{Field{typ: TypeBool, name: name}} }

// Time returns a new Field with type timestamp.
func Time(name string) *timeBuilder { return &timeBuilder{Field{typ: TypeTime, name: name}} }

// Type returns the field type.
func (f Field) Type() Type { return f.typ }

// Name returns the field name.
func (f Field) Name() string { return f.name }

// HasDefault returns is this field has a default value.
func (f Field) HasDefault() bool { return f.value != nil }

// Value returns the default value of the field.
func (f Field) Value() interface{} { return f.value }

// IsNillable returns if this field is an nillable field. Basically, wraps the value with pointer.
func (f Field) IsNillable() bool { return f.nillable }

// IsOptional returns is this field is an optional field.
func (f Field) IsOptional() bool { return f.optional }

// IsImmutable returns is this field is an immutable field.
func (f Field) IsImmutable() bool { return f.immutable }

// IsUnique returns is this field is a unique field.
func (f Field) IsUnique() bool { return f.unique }

// Validators returns the field matchers.
func (f Field) Validators() []interface{} { return f.validators }

// Tag returns the struct tag of the field.
func (f Field) Tag() string { return f.tag }

// stringBuilder is the builder for string fields.
type stringBuilder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *stringBuilder) Unique() *stringBuilder {
	b.unique = true
	return b
}

// Match adds a regex matcher for this field. Operation fails if the regex fails.
func (b *stringBuilder) Match(re *regexp.Regexp) *stringBuilder {
	b.validators = append(b.validators, func(v string) error {
		if !re.MatchString(v) {
			return errors.New("value does not match validation")
		}
		return nil
	})
	return b
}

// MinLen adds a length validator for this field.
// Operation fails if the length of the string is less than the given value.
func (b *stringBuilder) MinLen(i int) *stringBuilder {
	b.validators = append(b.validators, func(v string) error {
		if len(v) < i {
			return errors.New("value is less than the required length")
		}
		return nil
	})
	return b
}

// MaxLen adds a length validator for this field.
// Operation fails if the length of the string is greater than the given value.
func (b *stringBuilder) MaxLen(i int) *stringBuilder {
	b.size = i
	b.validators = append(b.validators, func(v string) error {
		if len(v) > i {
			return errors.New("value is less than the required length")
		}
		return nil
	})
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *stringBuilder) Validate(fn func(string) error) *stringBuilder {
	b.validators = append(b.validators, fn)
	return b
}

// Default sets the default value of the field.
func (b *stringBuilder) Default(s string) *stringBuilder {
	b.value = s
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *stringBuilder) Nillable() *stringBuilder {
	b.nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *stringBuilder) Optional() *stringBuilder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *stringBuilder) Immutable() *stringBuilder {
	b.immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *stringBuilder) Comment(c string) *stringBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *stringBuilder) StructTag(s string) *stringBuilder {
	b.tag = s
	return b
}

// SetCharset sets the character set attribute for character fields.
// For example, utf8 or utf8mb4 in MySQL.
func (b *stringBuilder) SetCharset(s string) *stringBuilder {
	b.charset = s
	return b
}

// Size returns the maximum size of a string.
// In SQL dialects this is parameter for varchar.
func (b stringBuilder) Size() int { return b.size }

// Charset returns the character set of the field.
func (b stringBuilder) Charset() string { return b.charset }

// timeBuilder is the builder for time fields.
type timeBuilder struct {
	Field
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *timeBuilder) Nillable() *timeBuilder {
	b.nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *timeBuilder) Optional() *timeBuilder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *timeBuilder) Immutable() *timeBuilder {
	b.immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *timeBuilder) Comment(c string) *timeBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *timeBuilder) StructTag(s string) *timeBuilder {
	b.tag = s
	return b
}

// Default sets the function that is applied to set default value
// of the field on creation. For example:
//
//	field.Time("created_at").
//		Default(time.Now)
//
func (b *timeBuilder) Default(f func() time.Time) *timeBuilder {
	b.value = f
	return b
}

// boolBuilder is the builder for boolean fields.
type boolBuilder struct {
	Field
}

// Default sets the default value of the field.
func (b *boolBuilder) Default(v bool) *boolBuilder {
	b.value = v
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *boolBuilder) Nillable() *boolBuilder {
	b.nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *boolBuilder) Optional() *boolBuilder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *boolBuilder) Immutable() *boolBuilder {
	b.immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *boolBuilder) Comment(c string) *boolBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *boolBuilder) StructTag(s string) *boolBuilder {
	b.tag = s
	return b
}

// bytesBuilder is the builder for bytes fields.
type bytesBuilder struct {
	Field
}

// Default sets the default value of the field.
func (b *bytesBuilder) Default(v []byte) *bytesBuilder {
	b.value = v
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *bytesBuilder) Nillable() *bytesBuilder {
	b.nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *bytesBuilder) Optional() *bytesBuilder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *bytesBuilder) Immutable() *bytesBuilder {
	b.immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *bytesBuilder) Comment(c string) *bytesBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *bytesBuilder) StructTag(s string) *bytesBuilder {
	b.tag = s
	return b
}

// Charseter is the interface that wraps the Charset method.
type Charseter interface {
	Charset() string
}

// Sizer is the interface that wraps the Size method.
type Sizer interface {
	Size() int
}
