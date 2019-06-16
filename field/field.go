package field

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Type is a field type.
type Type uint

// Field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeTime
	TypeString
	TypeInt
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt64
	TypeUint
	TypeUint8
	TypeUint16
	TypeUint32
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

// Numeric reports of the given type is a numeric type.
func (t Type) Numeric() bool { return t >= TypeInt && t < endTypes }

// ConstName returns the constant name of a type. It's used by entc for printing the constant name in templates.
func (t Type) ConstName() string {
	if t == TypeTime {
		return "TypeTime"
	}
	return "Type" + strings.Title(t.String())
}

var typeNames = [...]string{
	TypeInvalid: "invalid",
	TypeBool:    "bool",
	TypeTime:    "time.Time",
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

// Field represents a field on a graph vertex.
type Field struct {
	typ        Type
	tag        string
	name       string
	comment    string
	unique     bool
	nullable   bool
	optional   bool
	value      interface{}
	matchers   []*regexp.Regexp
	validators []interface{}
}

// Int returns a new Field with type int.
func Int(name string) *intBuilder { return &intBuilder{Field{typ: TypeInt, name: name}} }

// Float returns a new Field with type float.
func Float(name string) *floatBuilder { return &floatBuilder{Field{typ: TypeFloat64, name: name}} }

// String returns a new Field with type string.
func String(name string) *stringBuilder { return &stringBuilder{Field{typ: TypeString, name: name}} }

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

// IsNullable returns if this field is an nullable field. Basically, wraps the value with pointer.
func (f Field) IsNullable() bool { return f.nullable }

// IsOptional returns is this field is an optional field.
func (f Field) IsOptional() bool { return f.optional }

// IsUnique returns is this field is a unique field.
func (f Field) IsUnique() bool { return f.unique }

// Validators returns the field matchers.
func (f Field) Validators() []interface{} { return f.validators }

// Tag returns the struct tag of the field.
func (f Field) Tag() string { return f.tag }

// intBuilder is the builder for int field.
type intBuilder struct {
	Field
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *intBuilder) Range(i, j int) *intBuilder {
	b.validators = append(b.validators, func(v int) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *intBuilder) Min(i int) *intBuilder {
	b.validators = append(b.validators, func(v int) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *intBuilder) Max(i int) *intBuilder {
	b.validators = append(b.validators, func(v int) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *intBuilder) Positive() *intBuilder {
	return b.Min(1)
}

// Negative adds a maximum value validator with the value of -1. Operation fails if the validator fails.
func (b *intBuilder) Negative() *intBuilder {
	return b.Max(-1)
}

// Default sets the default value of the field.
func (b *intBuilder) Default(i int) *intBuilder {
	b.value = i
	return b
}

// Nullable indicates that this field is nullable.
// Unlike "Optional", nullable fields are pointers in the generated field.
func (b *intBuilder) Nullable() *intBuilder {
	b.nullable = true
	return b
}

// Comment sets the comment of the field.
func (b *intBuilder) Comment(c string) *intBuilder {
	b.comment = c
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *intBuilder) Optional() *intBuilder {
	b.optional = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *intBuilder) StructTag(s string) *intBuilder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *intBuilder) Validate(fn func(int) error) *intBuilder {
	b.validators = append(b.validators, fn)
	return b
}

// floatBuilder is the builder for float fields.
type floatBuilder struct {
	Field
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *floatBuilder) Range(i, j float64) *floatBuilder {
	b.validators = append(b.validators, func(v float64) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *floatBuilder) Min(i float64) *floatBuilder {
	b.validators = append(b.validators, func(v float64) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *floatBuilder) Max(i float64) *floatBuilder {
	b.validators = append(b.validators, func(v float64) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 0.000001. Operation fails if the validator fails.
func (b *floatBuilder) Positive() *floatBuilder {
	return b.Min(1e-06)
}

// Negative adds a maximum value validator with the value of -0.000001. Operation fails if the validator fails.
func (b *floatBuilder) Negative() *floatBuilder {
	return b.Max(-1e-06)
}

// Default sets the default value of the field.
func (b *floatBuilder) Default(i float64) *floatBuilder {
	b.value = i
	return b
}

// Nullable indicates that this field is nullable.
// Unlike "Optional", nullable fields are pointers in the generated field.
func (b *floatBuilder) Nullable() *floatBuilder {
	b.nullable = true
	return b
}

// Comment sets the comment of the field.
func (b *floatBuilder) Comment(c string) *floatBuilder {
	b.comment = c
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *floatBuilder) Optional() *floatBuilder {
	b.optional = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *floatBuilder) StructTag(s string) *floatBuilder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *floatBuilder) Validate(fn func(float64) error) *floatBuilder {
	b.validators = append(b.validators, fn)
	return b
}

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

// Nullable indicates that this field is nullable.
// Unlike "Optional", nullable fields are pointers in the generated field.
func (b *stringBuilder) Nullable() *stringBuilder {
	b.nullable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *stringBuilder) Optional() *stringBuilder {
	b.optional = true
	return b
}

// Comment sets the comment of the field.
func (b *stringBuilder) Comment(c string) *stringBuilder {
	b.comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *stringBuilder) StructTag(s string) *stringBuilder {
	b.tag = s
	return b
}

// timeBuilder is the builder for time fields.
type timeBuilder struct {
	Field
}

// Nullable indicates that this field is nullable.
// Unlike "Optional", nullable fields are pointers in the generated field.
func (b *timeBuilder) Nullable() *timeBuilder {
	b.nullable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *timeBuilder) Optional() *timeBuilder {
	b.optional = true
	return b
}

// Comment sets the comment of the field.
func (b *timeBuilder) Comment(c string) *timeBuilder {
	b.comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *timeBuilder) StructTag(s string) *timeBuilder {
	b.tag = s
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

// Nullable indicates that this field is nullable.
// Unlike "Optional", nullable fields are pointers in the generated field.
func (b *boolBuilder) Nullable() *boolBuilder {
	b.nullable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *boolBuilder) Optional() *boolBuilder {
	b.optional = true
	return b
}

// Comment sets the comment of the field.
func (b *boolBuilder) Comment(c string) *boolBuilder {
	b.comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *boolBuilder) StructTag(s string) *boolBuilder {
	b.tag = s
	return b
}
