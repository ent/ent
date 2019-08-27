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

// A Descriptor for field configuration.
type Descriptor struct {
	Tag           string        // struct tag.
	Size          int           // varchar size.
	Name          string        // field name.
	Type          Type          // field type.
	Charset       string        // string charset.
	Unique        bool          // unique index of field.
	Nillable      bool          // nillable struct field.
	Optional      bool          // nullable field in database.
	Immutable     bool          // create-only field.
	Default       interface{}   // default value on create.
	UpdateDefault interface{}   // default value on update.
	Validators    []interface{} // validator functions.
}

// String returns a new Field with type string.
func String(name string) *stringBuilder {
	return &stringBuilder{desc: &Descriptor{Type: TypeString, Name: name}}
}

// Text returns a new string field without limitation on the size.
// In MySQL, it is the "longtext" type, but in SQLite and Gremlin it has not effect.
func Text(name string) *stringBuilder {
	return &stringBuilder{desc: &Descriptor{Type: TypeString, Name: name, Size: math.MaxInt32}}
}

// Bytes returns a new Field with type bytes/buffer.
// In MySQL and SQLite, it is the "BLOB" type, and it does not support for Gremlin.
func Bytes(name string) *bytesBuilder {
	return &bytesBuilder{desc: &Descriptor{Type: TypeBytes, Name: name}}
}

// Bool returns a new Field with type bool.
func Bool(name string) *boolBuilder {
	return &boolBuilder{desc: &Descriptor{Type: TypeBool, Name: name}}
}

// Time returns a new Field with type timestamp.
func Time(name string) *timeBuilder {
	return &timeBuilder{desc: &Descriptor{Type: TypeTime, Name: name}}
}

// stringBuilder is the builder for string fields.
type stringBuilder struct {
	desc *Descriptor
}

// Unique makes the field unique within all vertices of this type.
func (b *stringBuilder) Unique() *stringBuilder {
	b.desc.Unique = true
	return b
}

// Match adds a regex matcher for this field. Operation fails if the regex fails.
func (b *stringBuilder) Match(re *regexp.Regexp) *stringBuilder {
	b.desc.Validators = append(b.desc.Validators, func(v string) error {
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
	b.desc.Validators = append(b.desc.Validators, func(v string) error {
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
	b.desc.Size = i
	b.desc.Validators = append(b.desc.Validators, func(v string) error {
		if len(v) > i {
			return errors.New("value is less than the required length")
		}
		return nil
	})
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *stringBuilder) Validate(fn func(string) error) *stringBuilder {
	b.desc.Validators = append(b.desc.Validators, fn)
	return b
}

// Default sets the default value of the field.
func (b *stringBuilder) Default(s string) *stringBuilder {
	b.desc.Default = s
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *stringBuilder) Nillable() *stringBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *stringBuilder) Optional() *stringBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *stringBuilder) Immutable() *stringBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *stringBuilder) Comment(c string) *stringBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *stringBuilder) StructTag(s string) *stringBuilder {
	b.desc.Tag = s
	return b
}

// SetCharset sets the character set attribute for character fields.
// For example, utf8 or utf8mb4 in MySQL.
func (b *stringBuilder) Charset(s string) *stringBuilder {
	b.desc.Charset = s
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *stringBuilder) Descriptor() *Descriptor {
	return b.desc
}

// timeBuilder is the builder for time fields.
type timeBuilder struct {
	desc *Descriptor
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *timeBuilder) Nillable() *timeBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *timeBuilder) Optional() *timeBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *timeBuilder) Immutable() *timeBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *timeBuilder) Comment(c string) *timeBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *timeBuilder) StructTag(s string) *timeBuilder {
	b.desc.Tag = s
	return b
}

// Default sets the function that is applied to set default value
// of the field on creation. For example:
//
//	field.Time("created_at").
//		Default(time.Now)
//
func (b *timeBuilder) Default(f func() time.Time) *timeBuilder {
	b.desc.Default = f
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *timeBuilder) Descriptor() *Descriptor {
	return b.desc
}

// boolBuilder is the builder for boolean fields.
type boolBuilder struct {
	desc *Descriptor
}

// Default sets the default value of the field.
func (b *boolBuilder) Default(v bool) *boolBuilder {
	b.desc.Default = v
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *boolBuilder) Nillable() *boolBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *boolBuilder) Optional() *boolBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *boolBuilder) Immutable() *boolBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *boolBuilder) Comment(c string) *boolBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *boolBuilder) StructTag(s string) *boolBuilder {
	b.desc.Tag = s
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *boolBuilder) Descriptor() *Descriptor {
	return b.desc
}

// bytesBuilder is the builder for bytes fields.
type bytesBuilder struct {
	desc *Descriptor
}

// Default sets the default value of the field.
func (b *bytesBuilder) Default(v []byte) *bytesBuilder {
	b.desc.Default = v
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *bytesBuilder) Nillable() *bytesBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *bytesBuilder) Optional() *bytesBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *bytesBuilder) Immutable() *bytesBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *bytesBuilder) Comment(c string) *bytesBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *bytesBuilder) StructTag(s string) *bytesBuilder {
	b.desc.Tag = s
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *bytesBuilder) Descriptor() *Descriptor {
	return b.desc
}
