// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field

import (
	"database/sql/driver"
	"errors"
	"math"
	"reflect"
	"regexp"
	"time"
)

// A Descriptor for field configuration.
type Descriptor struct {
	Tag           string        // struct tag.
	Size          int           // varchar size.
	Name          string        // field name.
	Info          *TypeInfo     // field type info.
	Unique        bool          // unique index of field.
	Nillable      bool          // nillable struct field.
	Optional      bool          // nullable field in database.
	Immutable     bool          // create-only field.
	Default       interface{}   // default value on create.
	UpdateDefault interface{}   // default value on update.
	Validators    []interface{} // validator functions.
	StorageKey    string        // sql column or gremlin property.
	Enums         []string      // enum values.
	Sensitive     bool          // sensitive info string field.
}

// String returns a new Field with type string.
func String(name string) *stringBuilder {
	return &stringBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeString},
	}}
}

// Text returns a new string field without limitation on the size.
// In MySQL, it is the "longtext" type, but in SQLite and Gremlin it has not effect.
func Text(name string) *stringBuilder {
	return &stringBuilder{&Descriptor{
		Name: name,
		Size: math.MaxInt32,
		Info: &TypeInfo{Type: TypeString},
	}}
}

// Bytes returns a new Field with type bytes/buffer.
// In MySQL and SQLite, it is the "BLOB" type, and it does not support for Gremlin.
func Bytes(name string) *bytesBuilder {
	return &bytesBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeBytes, Nillable: true},
	}}
}

// Bool returns a new Field with type bool.
func Bool(name string) *boolBuilder {
	return &boolBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeBool},
	}}
}

// Time returns a new Field with type timestamp.
func Time(name string) *timeBuilder {
	return &timeBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeTime, PkgPath: "time"},
	}}
}

// JSON returns a new Field with type json that is serialized to the given object.
// For example:
//
//	field.JSON("dirs", []http.Dir{}).
//		Optional()
//
//
//	field.JSON("info", &Info{}).
//		Optional()
//
func JSON(name string, typ interface{}) *jsonBuilder {
	t := reflect.TypeOf(typ)
	info := &TypeInfo{
		Type:    TypeJSON,
		Ident:   t.String(),
		PkgPath: t.PkgPath(),
	}
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map:
		info.Nillable = true
	}
	return &jsonBuilder{&Descriptor{
		Name: name,
		Info: info,
	}}
}

// Strings returns a new JSON Field with type []string.
func Strings(name string) *jsonBuilder {
	return JSON(name, []string{})
}

// Ints returns a new JSON Field with type []int.
func Ints(name string) *jsonBuilder {
	return JSON(name, []int{})
}

// Floats returns a new JSON Field with type []float.
func Floats(name string) *jsonBuilder {
	return JSON(name, []float64{})
}

// Enum returns a new Field with type enum. An example for defining enum is as follows:
//
//	field.Enum("state").
//		Values(
//			"on",
//			"off",
//		).
//		Default("on")
//
func Enum(name string) *enumBuilder {
	return &enumBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeEnum},
	}}
}

// UUID returns a new Field with type UUID. An example for defining UUID field is as follows:
//
//	field.UUID("id", uuid.New())
//
func UUID(name string, typ driver.Valuer) *uuidBuilder {
	rt := reflect.TypeOf(typ)
	return &uuidBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{
			Type:     TypeUUID,
			Nillable: true,
			Ident:    rt.String(),
			PkgPath:  rt.PkgPath(),
		},
	}}
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

// Sensitive fields not printable and not serializable.
func (b *stringBuilder) Sensitive() *stringBuilder {
	b.desc.Sensitive = true
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

// NotEmpty adds a length validator for this field.
// Operation fails if the length of the string is zero.
func (b *stringBuilder) NotEmpty() *stringBuilder {
	return b.MinLen(1)
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

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *stringBuilder) StorageKey(key string) *stringBuilder {
	b.desc.StorageKey = key
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

// UpdateDefault sets the function that is applied to set default value
// of the field on update. For example:
//
//	field.Time("updated_at").
//		Default(time.Now).
//		UpdateDefault(time.Now),
//
func (b *timeBuilder) UpdateDefault(f func() time.Time) *timeBuilder {
	b.desc.UpdateDefault = f
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *timeBuilder) StorageKey(key string) *timeBuilder {
	b.desc.StorageKey = key
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

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *boolBuilder) StorageKey(key string) *boolBuilder {
	b.desc.StorageKey = key
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

// MaxLen sets the max-length of the bytes type in the database.
// In MySQL, this affects the BLOB type (tiny 2^8-1, regular 2^16-1, medium 2^24-1, long 2^32-1).
// In SQLite, it does not have any effect on the type size, which is default to 1B bytes.
func (b *bytesBuilder) MaxLen(i int) *bytesBuilder {
	b.desc.Size = i
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *bytesBuilder) StorageKey(key string) *bytesBuilder {
	b.desc.StorageKey = key
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *bytesBuilder) Descriptor() *Descriptor {
	return b.desc
}

// jsonBuilder is the builder for json fields.
type jsonBuilder struct {
	desc *Descriptor
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *jsonBuilder) StorageKey(key string) *jsonBuilder {
	b.desc.StorageKey = key
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *jsonBuilder) Optional() *jsonBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *jsonBuilder) Immutable() *jsonBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *jsonBuilder) Comment(c string) *jsonBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *jsonBuilder) StructTag(s string) *jsonBuilder {
	b.desc.Tag = s
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *jsonBuilder) Descriptor() *Descriptor {
	return b.desc
}

// enumBuilder is the builder for enum fields.
type enumBuilder struct {
	desc *Descriptor
}

// Value sets the numeric value of the enum. Defaults to its index in the declaration.
func (b *enumBuilder) Values(values ...string) *enumBuilder {
	b.desc.Enums = values
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *enumBuilder) StorageKey(key string) *enumBuilder {
	b.desc.StorageKey = key
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *enumBuilder) Optional() *enumBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *enumBuilder) Immutable() *enumBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *enumBuilder) Comment(c string) *enumBuilder {
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *enumBuilder) Nillable() *enumBuilder {
	b.desc.Nillable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *enumBuilder) StructTag(s string) *enumBuilder {
	b.desc.Tag = s
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *enumBuilder) Descriptor() *Descriptor {
	return b.desc
}

// uuidBuilder is the builder for uuid fields.
type uuidBuilder struct {
	desc *Descriptor
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *uuidBuilder) StorageKey(key string) *uuidBuilder {
	b.desc.StorageKey = key
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *uuidBuilder) Optional() *uuidBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *uuidBuilder) Immutable() *uuidBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *uuidBuilder) Comment(c string) *uuidBuilder {
	return b
}

// StructTag sets the struct tag of the field.
func (b *uuidBuilder) StructTag(s string) *uuidBuilder {
	b.desc.Tag = s
	return b
}

// Default sets the function that is applied to set default value
// of the field on creation. Codegen fails if the default function
// doesn't return the same concrete that was set for the UUID type.
//
//	field.UUID("id", uuid.UUID{}).
//		Default(uuid.New)
//
func (b *uuidBuilder) Default(fn interface{}) *uuidBuilder {
	b.desc.Default = fn
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *uuidBuilder) Descriptor() *Descriptor {
	return b.desc
}
