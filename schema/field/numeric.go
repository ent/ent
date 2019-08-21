package field

import "errors"

//go:generate go run gen/gen.go

// Int returns a new Field with type int.
func Int(name string) *intBuilder { return &intBuilder{Field{typ: TypeInt, name: name}} }

// Uint returns a new Field with type uint.
func Uint(name string) *uintBuilder { return &uintBuilder{Field{typ: TypeUint, name: name}} }

// Int8 returns a new Field with type int8.
func Int8(name string) *int8Builder { return &int8Builder{Field{typ: TypeInt8, name: name}} }

// Int16 returns a new Field with type int16.
func Int16(name string) *int16Builder { return &int16Builder{Field{typ: TypeInt16, name: name}} }

// Int32 returns a new Field with type int32.
func Int32(name string) *int32Builder { return &int32Builder{Field{typ: TypeInt32, name: name}} }

// Int64 returns a new Field with type int64.
func Int64(name string) *int64Builder { return &int64Builder{Field{typ: TypeInt64, name: name}} }

// Uint8 returns a new Field with type uint8.
func Uint8(name string) *uint8Builder { return &uint8Builder{Field{typ: TypeUint8, name: name}} }

// Uint16 returns a new Field with type uint16.
func Uint16(name string) *uint16Builder { return &uint16Builder{Field{typ: TypeUint16, name: name}} }

// Uint32 returns a new Field with type uint32.
func Uint32(name string) *uint32Builder { return &uint32Builder{Field{typ: TypeUint32, name: name}} }

// Uint64 returns a new Field with type uint64.
func Uint64(name string) *uint64Builder { return &uint64Builder{Field{typ: TypeUint64, name: name}} }

// Float returns a new Field with type float64.
func Float(name string) *float64Builder { return &float64Builder{Field{typ: TypeFloat64, name: name}} }

// Float32 returns a new Field with type float32.
func Float32(name string) *float32Builder { return &float32Builder{Field{typ: TypeFloat32, name: name}} }

// intBuilder is the builder for int field.
type intBuilder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *intBuilder) Unique() *intBuilder {
	b.unique = true
	return b
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

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *intBuilder) Nillable() *intBuilder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *intBuilder) Comment(c string) *intBuilder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *intBuilder) Optional() *intBuilder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *intBuilder) Immutable() *intBuilder {
	b.immutable = true
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

// uintBuilder is the builder for uint field.
type uintBuilder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *uintBuilder) Unique() *uintBuilder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *uintBuilder) Range(i, j uint) *uintBuilder {
	b.validators = append(b.validators, func(v uint) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *uintBuilder) Min(i uint) *uintBuilder {
	b.validators = append(b.validators, func(v uint) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *uintBuilder) Max(i uint) *uintBuilder {
	b.validators = append(b.validators, func(v uint) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *uintBuilder) Positive() *uintBuilder {
	return b.Min(1)
}

// Default sets the default value of the field.
func (b *uintBuilder) Default(i uint) *uintBuilder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *uintBuilder) Nillable() *uintBuilder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *uintBuilder) Comment(c string) *uintBuilder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *uintBuilder) Optional() *uintBuilder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *uintBuilder) Immutable() *uintBuilder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *uintBuilder) StructTag(s string) *uintBuilder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *uintBuilder) Validate(fn func(uint) error) *uintBuilder {
	b.validators = append(b.validators, fn)
	return b
}

// int8Builder is the builder for int8 field.
type int8Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *int8Builder) Unique() *int8Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *int8Builder) Range(i, j int8) *int8Builder {
	b.validators = append(b.validators, func(v int8) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *int8Builder) Min(i int8) *int8Builder {
	b.validators = append(b.validators, func(v int8) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *int8Builder) Max(i int8) *int8Builder {
	b.validators = append(b.validators, func(v int8) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *int8Builder) Positive() *int8Builder {
	return b.Min(1)
}

// Negative adds a maximum value validator with the value of -1. Operation fails if the validator fails.
func (b *int8Builder) Negative() *int8Builder {
	return b.Max(-1)
}

// Default sets the default value of the field.
func (b *int8Builder) Default(i int8) *int8Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *int8Builder) Nillable() *int8Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *int8Builder) Comment(c string) *int8Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *int8Builder) Optional() *int8Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *int8Builder) Immutable() *int8Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *int8Builder) StructTag(s string) *int8Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *int8Builder) Validate(fn func(int8) error) *int8Builder {
	b.validators = append(b.validators, fn)
	return b
}

// int16Builder is the builder for int16 field.
type int16Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *int16Builder) Unique() *int16Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *int16Builder) Range(i, j int16) *int16Builder {
	b.validators = append(b.validators, func(v int16) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *int16Builder) Min(i int16) *int16Builder {
	b.validators = append(b.validators, func(v int16) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *int16Builder) Max(i int16) *int16Builder {
	b.validators = append(b.validators, func(v int16) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *int16Builder) Positive() *int16Builder {
	return b.Min(1)
}

// Negative adds a maximum value validator with the value of -1. Operation fails if the validator fails.
func (b *int16Builder) Negative() *int16Builder {
	return b.Max(-1)
}

// Default sets the default value of the field.
func (b *int16Builder) Default(i int16) *int16Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *int16Builder) Nillable() *int16Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *int16Builder) Comment(c string) *int16Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *int16Builder) Optional() *int16Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *int16Builder) Immutable() *int16Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *int16Builder) StructTag(s string) *int16Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *int16Builder) Validate(fn func(int16) error) *int16Builder {
	b.validators = append(b.validators, fn)
	return b
}

// int32Builder is the builder for int32 field.
type int32Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *int32Builder) Unique() *int32Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *int32Builder) Range(i, j int32) *int32Builder {
	b.validators = append(b.validators, func(v int32) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *int32Builder) Min(i int32) *int32Builder {
	b.validators = append(b.validators, func(v int32) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *int32Builder) Max(i int32) *int32Builder {
	b.validators = append(b.validators, func(v int32) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *int32Builder) Positive() *int32Builder {
	return b.Min(1)
}

// Negative adds a maximum value validator with the value of -1. Operation fails if the validator fails.
func (b *int32Builder) Negative() *int32Builder {
	return b.Max(-1)
}

// Default sets the default value of the field.
func (b *int32Builder) Default(i int32) *int32Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *int32Builder) Nillable() *int32Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *int32Builder) Comment(c string) *int32Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *int32Builder) Optional() *int32Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *int32Builder) Immutable() *int32Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *int32Builder) StructTag(s string) *int32Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *int32Builder) Validate(fn func(int32) error) *int32Builder {
	b.validators = append(b.validators, fn)
	return b
}

// int64Builder is the builder for int64 field.
type int64Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *int64Builder) Unique() *int64Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *int64Builder) Range(i, j int64) *int64Builder {
	b.validators = append(b.validators, func(v int64) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *int64Builder) Min(i int64) *int64Builder {
	b.validators = append(b.validators, func(v int64) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *int64Builder) Max(i int64) *int64Builder {
	b.validators = append(b.validators, func(v int64) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *int64Builder) Positive() *int64Builder {
	return b.Min(1)
}

// Negative adds a maximum value validator with the value of -1. Operation fails if the validator fails.
func (b *int64Builder) Negative() *int64Builder {
	return b.Max(-1)
}

// Default sets the default value of the field.
func (b *int64Builder) Default(i int64) *int64Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *int64Builder) Nillable() *int64Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *int64Builder) Comment(c string) *int64Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *int64Builder) Optional() *int64Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *int64Builder) Immutable() *int64Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *int64Builder) StructTag(s string) *int64Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *int64Builder) Validate(fn func(int64) error) *int64Builder {
	b.validators = append(b.validators, fn)
	return b
}

// uint8Builder is the builder for uint8 field.
type uint8Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *uint8Builder) Unique() *uint8Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *uint8Builder) Range(i, j uint8) *uint8Builder {
	b.validators = append(b.validators, func(v uint8) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *uint8Builder) Min(i uint8) *uint8Builder {
	b.validators = append(b.validators, func(v uint8) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *uint8Builder) Max(i uint8) *uint8Builder {
	b.validators = append(b.validators, func(v uint8) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *uint8Builder) Positive() *uint8Builder {
	return b.Min(1)
}

// Default sets the default value of the field.
func (b *uint8Builder) Default(i uint8) *uint8Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *uint8Builder) Nillable() *uint8Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *uint8Builder) Comment(c string) *uint8Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *uint8Builder) Optional() *uint8Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *uint8Builder) Immutable() *uint8Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *uint8Builder) StructTag(s string) *uint8Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *uint8Builder) Validate(fn func(uint8) error) *uint8Builder {
	b.validators = append(b.validators, fn)
	return b
}

// uint16Builder is the builder for uint16 field.
type uint16Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *uint16Builder) Unique() *uint16Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *uint16Builder) Range(i, j uint16) *uint16Builder {
	b.validators = append(b.validators, func(v uint16) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *uint16Builder) Min(i uint16) *uint16Builder {
	b.validators = append(b.validators, func(v uint16) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *uint16Builder) Max(i uint16) *uint16Builder {
	b.validators = append(b.validators, func(v uint16) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *uint16Builder) Positive() *uint16Builder {
	return b.Min(1)
}

// Default sets the default value of the field.
func (b *uint16Builder) Default(i uint16) *uint16Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *uint16Builder) Nillable() *uint16Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *uint16Builder) Comment(c string) *uint16Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *uint16Builder) Optional() *uint16Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *uint16Builder) Immutable() *uint16Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *uint16Builder) StructTag(s string) *uint16Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *uint16Builder) Validate(fn func(uint16) error) *uint16Builder {
	b.validators = append(b.validators, fn)
	return b
}

// uint32Builder is the builder for uint32 field.
type uint32Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *uint32Builder) Unique() *uint32Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *uint32Builder) Range(i, j uint32) *uint32Builder {
	b.validators = append(b.validators, func(v uint32) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *uint32Builder) Min(i uint32) *uint32Builder {
	b.validators = append(b.validators, func(v uint32) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *uint32Builder) Max(i uint32) *uint32Builder {
	b.validators = append(b.validators, func(v uint32) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *uint32Builder) Positive() *uint32Builder {
	return b.Min(1)
}

// Default sets the default value of the field.
func (b *uint32Builder) Default(i uint32) *uint32Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *uint32Builder) Nillable() *uint32Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *uint32Builder) Comment(c string) *uint32Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *uint32Builder) Optional() *uint32Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *uint32Builder) Immutable() *uint32Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *uint32Builder) StructTag(s string) *uint32Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *uint32Builder) Validate(fn func(uint32) error) *uint32Builder {
	b.validators = append(b.validators, fn)
	return b
}

// uint64Builder is the builder for uint64 field.
type uint64Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *uint64Builder) Unique() *uint64Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *uint64Builder) Range(i, j uint64) *uint64Builder {
	b.validators = append(b.validators, func(v uint64) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *uint64Builder) Min(i uint64) *uint64Builder {
	b.validators = append(b.validators, func(v uint64) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *uint64Builder) Max(i uint64) *uint64Builder {
	b.validators = append(b.validators, func(v uint64) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 1. Operation fails if the validator fails.
func (b *uint64Builder) Positive() *uint64Builder {
	return b.Min(1)
}

// Default sets the default value of the field.
func (b *uint64Builder) Default(i uint64) *uint64Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *uint64Builder) Nillable() *uint64Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *uint64Builder) Comment(c string) *uint64Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *uint64Builder) Optional() *uint64Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *uint64Builder) Immutable() *uint64Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *uint64Builder) StructTag(s string) *uint64Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *uint64Builder) Validate(fn func(uint64) error) *uint64Builder {
	b.validators = append(b.validators, fn)
	return b
}

// float64Builder is the builder for float fields.
type float64Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *float64Builder) Unique() *float64Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *float64Builder) Range(i, j float64) *float64Builder {
	b.validators = append(b.validators, func(v float64) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *float64Builder) Min(i float64) *float64Builder {
	b.validators = append(b.validators, func(v float64) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *float64Builder) Max(i float64) *float64Builder {
	b.validators = append(b.validators, func(v float64) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 0.000001. Operation fails if the validator fails.
func (b *float64Builder) Positive() *float64Builder {
	return b.Min(1e-06)
}

// Negative adds a maximum value validator with the value of -0.000001. Operation fails if the validator fails.
func (b *float64Builder) Negative() *float64Builder {
	return b.Max(-1e-06)
}

// Default sets the default value of the field.
func (b *float64Builder) Default(i float64) *float64Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *float64Builder) Nillable() *float64Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *float64Builder) Comment(c string) *float64Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *float64Builder) Optional() *float64Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *float64Builder) Immutable() *float64Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *float64Builder) StructTag(s string) *float64Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *float64Builder) Validate(fn func(float64) error) *float64Builder {
	b.validators = append(b.validators, fn)
	return b
}

// float32Builder is the builder for float fields.
type float32Builder struct {
	Field
}

// Unique makes the field unique within all vertices of this type.
func (b *float32Builder) Unique() *float32Builder {
	b.unique = true
	return b
}

// Range adds a range validator for this field where the given value needs to be in the range of [i, j].
func (b *float32Builder) Range(i, j float32) *float32Builder {
	b.validators = append(b.validators, func(v float32) error {
		if v < i || v > j {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Min adds a minimum value validator for this field. Operation fails if the validator fails.
func (b *float32Builder) Min(i float32) *float32Builder {
	b.validators = append(b.validators, func(v float32) error {
		if v < i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Max adds a maximum value validator for this field. Operation fails if the validator fails.
func (b *float32Builder) Max(i float32) *float32Builder {
	b.validators = append(b.validators, func(v float32) error {
		if v > i {
			return errors.New("value out of range")
		}
		return nil
	})
	return b
}

// Positive adds a minimum value validator with the value of 0.000001. Operation fails if the validator fails.
func (b *float32Builder) Positive() *float32Builder {
	return b.Min(1e-06)
}

// Negative adds a maximum value validator with the value of -0.000001. Operation fails if the validator fails.
func (b *float32Builder) Negative() *float32Builder {
	return b.Max(-1e-06)
}

// Default sets the default value of the field.
func (b *float32Builder) Default(i float32) *float32Builder {
	b.value = i
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *float32Builder) Nillable() *float32Builder {
	b.nillable = true
	return b
}

// Comment sets the comment of the field.
func (b *float32Builder) Comment(c string) *float32Builder {
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *float32Builder) Optional() *float32Builder {
	b.optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *float32Builder) Immutable() *float32Builder {
	b.immutable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *float32Builder) StructTag(s string) *float32Builder {
	b.tag = s
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *float32Builder) Validate(fn func(float32) error) *float32Builder {
	b.validators = append(b.validators, fn)
	return b
}
