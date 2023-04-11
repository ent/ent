// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"errors"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"

	"entgo.io/ent/schema"
)

// String returns a new Field with type string.
func String(name string) *StringBuilder {
	return &StringBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeString},
	}}
}

// Text returns a new string field without limitation on the size.
// In MySQL, it is the "longtext" type, but in SQLite and Gremlin it has no effect.
func Text(name string) *StringBuilder {
	return &StringBuilder{&Descriptor{
		Name: name,
		Size: math.MaxInt32,
		Info: &TypeInfo{Type: TypeString},
	}}
}

// Bytes returns a new Field with type bytes/buffer.
// In MySQL and SQLite, it is the "BLOB" type, and it does not support for Gremlin.
func Bytes(name string) *BytesBuilder {
	return &BytesBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeBytes, Nillable: true},
	}}
}

// Bool returns a new Field with type bool.
func Bool(name string) *BoolBuilder {
	return &BoolBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeBool},
	}}
}

// Time returns a new Field with type timestamp.
func Time(name string) *TimeBuilder {
	return &TimeBuilder{&Descriptor{
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
func JSON(name string, typ any) *JsonBuilder {
	b := &JsonBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{
			Type: TypeJSON,
		},
	}}
	t := reflect.TypeOf(typ)
	if t == nil {
		b.desc.Err = errors.New("expect a Go value as JSON type but got nil")
		return b
	}
	b.desc.Info.Ident = t.String()
	b.desc.Info.PkgPath = t.PkgPath()
	b.desc.goType(typ)
	b.desc.checkGoType(t)
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map:
		b.desc.Info.Nillable = true
		b.desc.Info.PkgPath = pkgPath(t)
	}
	return b
}

// Strings returns a new JSON Field with type []string.
func Strings(name string) *JsonBuilder {
	return JSON(name, []string{})
}

// Ints returns a new JSON Field with type []int.
func Ints(name string) *JsonBuilder {
	return JSON(name, []int{})
}

// Floats returns a new JSON Field with type []float.
func Floats(name string) *JsonBuilder {
	return JSON(name, []float64{})
}

// Any returns a new JSON Field with type any. Although this field type can be
// useful for fields with dynamic data layout, it is strongly recommended to use
// JSON with json.RawMessage instead and implement custom marshaling.
func Any(name string) *JsonBuilder {
	const t = "any"
	return &JsonBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{
			Type:     TypeJSON,
			Ident:    t,
			Nillable: true,
			RType: &RType{
				Name:  t,
				Ident: t,
				Kind:  reflect.Interface,
			},
		},
	}}
}

// Enum returns a new Field with type enum. An example for defining enum is as follows:
//
//	field.Enum("state").
//		Values(
//			"on",
//			"off",
//		).
//		Default("on")
func Enum(name string) *EnumBuilder {
	return &EnumBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeEnum},
	}}
}

// UUID returns a new Field with type UUID. An example for defining UUID field is as follows:
//
//	field.UUID("id", uuid.New())
func UUID(name string, typ driver.Valuer) *UuidBuilder {
	rt := reflect.TypeOf(typ)
	b := &UuidBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{
			Type:    TypeUUID,
			Ident:   rt.String(),
			PkgPath: indirect(rt).PkgPath(),
		},
	}}
	b.desc.goType(typ)
	return b
}

// Other represents a field that is not a good fit for any of the standard field types.
//
// The second argument defines the GoType and must implement the ValueScanner interface.
// The SchemaType option must be set because the field type cannot be inferred.
// An example for defining Other field is as follows:
//
//	field.Other("link", &Link{}).
//		SchemaType(map[string]string{
//			dialect.MySQL:    "text",
//			dialect.Postgres: "varchar",
//		})
func Other(name string, typ driver.Valuer) *OtherBuilder {
	ob := &OtherBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{Type: TypeOther},
	}}
	ob.desc.goType(typ)
	return ob
}

// StringBuilder is the builder for string fields.
type StringBuilder struct {
	desc *Descriptor
}

// Unique makes the field unique within all vertices of this type.
func (b *StringBuilder) Unique() *StringBuilder {
	b.desc.Unique = true
	return b
}

// Sensitive fields not printable and not serializable.
func (b *StringBuilder) Sensitive() *StringBuilder {
	b.desc.Sensitive = true
	return b
}

// Match adds a regex matcher for this field. Operation fails if the regex fails.
func (b *StringBuilder) Match(re *regexp.Regexp) *StringBuilder {
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
func (b *StringBuilder) MinLen(i int) *StringBuilder {
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
func (b *StringBuilder) NotEmpty() *StringBuilder {
	return b.MinLen(1)
}

// MaxLen adds a length validator for this field.
// Operation fails if the length of the string is greater than the given value.
func (b *StringBuilder) MaxLen(i int) *StringBuilder {
	b.desc.Size = i
	b.desc.Validators = append(b.desc.Validators, func(v string) error {
		if len(v) > i {
			return errors.New("value is greater than the required length")
		}
		return nil
	})
	return b
}

// Validate adds a validator for this field. Operation fails if the validation fails.
func (b *StringBuilder) Validate(fn func(string) error) *StringBuilder {
	b.desc.Validators = append(b.desc.Validators, fn)
	return b
}

// Default sets the default value of the field.
func (b *StringBuilder) Default(s string) *StringBuilder {
	b.desc.Default = s
	return b
}

// DefaultFunc sets the function that is applied to set the default value
// of the field on creation. For example:
//
//	field.String("cuid").
//		DefaultFunc(cuid.New)
func (b *StringBuilder) DefaultFunc(fn any) *StringBuilder {
	if t := reflect.TypeOf(fn); t.Kind() != reflect.Func {
		b.desc.Err = fmt.Errorf("field.String(%q).DefaultFunc expects func but got %s", b.desc.Name, t.Kind())
	}
	b.desc.Default = fn
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated struct.
func (b *StringBuilder) Nillable() *StringBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *StringBuilder) Optional() *StringBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *StringBuilder) Immutable() *StringBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *StringBuilder) Comment(c string) *StringBuilder {
	b.desc.Comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *StringBuilder) StructTag(s string) *StringBuilder {
	b.desc.Tag = s
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *StringBuilder) StorageKey(key string) *StringBuilder {
	b.desc.StorageKey = key
	return b
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for string.
//
//	field.String("name").
//		SchemaType(map[string]string{
//			dialect.MySQL:    "text",
//			dialect.Postgres: "varchar",
//		})
func (b *StringBuilder) SchemaType(types map[string]string) *StringBuilder {
	b.desc.SchemaType = types
	return b
}

// GoType overrides the default Go type with a custom one.
// If the provided type implements the Validator interface
// and no validators have been set, the type validator will
// be used.
//
//	field.String("dir").
//		GoType(http.Dir("dir"))
func (b *StringBuilder) GoType(typ any) *StringBuilder {
	b.desc.goType(typ)
	return b
}

// ValueScanner provides an external value scanner for the given GoType.
// Using this option allow users to use field types that do not implement
// the sql.Scanner and driver.Valuer interfaces, such as slices and maps
// or types exist in external packages (e.g., url.URL).
func (b *StringBuilder) ValueScanner(vs any) *StringBuilder {
	b.desc.ValueScanner = vs
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
//
//	field.String("dir").
//		Annotations(
//			entgql.OrderField("DIR"),
//		)
func (b *StringBuilder) Annotations(annotations ...schema.Annotation) *StringBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *StringBuilder) Descriptor() *Descriptor {
	if b.desc.Default != nil {
		b.desc.checkDefaultFunc(stringType)
	}
	b.desc.checkGoType(stringType)
	return b.desc
}

// TimeBuilder is the builder for time fields.
type TimeBuilder struct {
	desc *Descriptor
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated struct.
func (b *TimeBuilder) Nillable() *TimeBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *TimeBuilder) Optional() *TimeBuilder {
	b.desc.Optional = true
	return b
}

// Immutable fields are fields that can be set only in the creation of the entity.
// i.e., no setters will be generated for the entity updaters (one and many).
func (b *TimeBuilder) Immutable() *TimeBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *TimeBuilder) Comment(c string) *TimeBuilder {
	b.desc.Comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *TimeBuilder) StructTag(s string) *TimeBuilder {
	b.desc.Tag = s
	return b
}

// Default sets the function that is applied to set default value
// of the field on creation. For example:
//
//	field.Time("created_at").
//		Default(time.Now)
func (b *TimeBuilder) Default(fn any) *TimeBuilder {
	b.desc.Default = fn
	return b
}

// UpdateDefault sets the function that is applied to set default value
// of the field on update. For example:
//
//	field.Time("updated_at").
//		Default(time.Now).
//		UpdateDefault(time.Now),
//
//	field.Time("deleted_at").
//		Optional().
//		GoType(&sql.NullTime{}).
//		UpdateDefault(NewNullTime),
func (b *TimeBuilder) UpdateDefault(fn any) *TimeBuilder {
	b.desc.UpdateDefault = fn
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *TimeBuilder) StorageKey(key string) *TimeBuilder {
	b.desc.StorageKey = key
	return b
}

// GoType overrides the default Go type with a custom one.
// If the provided type implements the Validator interface
// and no validators have been set, the type validator will
// be used.
//
//	field.Time("deleted_at").
//		GoType(&sql.NullTime{})
func (b *TimeBuilder) GoType(typ any) *TimeBuilder {
	b.desc.goType(typ)
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
//
//	field.Time("deleted_at").
//		Annotations(
//			entgql.OrderField("DELETED_AT"),
//		)
func (b *TimeBuilder) Annotations(annotations ...schema.Annotation) *TimeBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *TimeBuilder) Descriptor() *Descriptor {
	if b.desc.Default != nil {
		b.desc.checkDefaultFunc(timeType)
	}
	b.desc.checkGoType(timeType)
	return b.desc
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for time.
//
//	field.Time("created_at").
//		SchemaType(map[string]string{
//			dialect.MySQL:    "datetime",
//			dialect.Postgres: "time with time zone",
//		})
func (b *TimeBuilder) SchemaType(types map[string]string) *TimeBuilder {
	b.desc.SchemaType = types
	return b
}

// BoolBuilder is the builder for boolean fields.
type BoolBuilder struct {
	desc *Descriptor
}

// Default sets the default value of the field.
func (b *BoolBuilder) Default(v bool) *BoolBuilder {
	b.desc.Default = v
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated struct.
func (b *BoolBuilder) Nillable() *BoolBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *BoolBuilder) Optional() *BoolBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *BoolBuilder) Immutable() *BoolBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *BoolBuilder) Comment(c string) *BoolBuilder {
	b.desc.Comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *BoolBuilder) StructTag(s string) *BoolBuilder {
	b.desc.Tag = s
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *BoolBuilder) StorageKey(key string) *BoolBuilder {
	b.desc.StorageKey = key
	return b
}

// GoType overrides the default Go type with a custom one.
// If the provided type implements the Validator interface
// and no validators have been set, the type validator will
// be used.
//
//	field.Bool("deleted").
//		GoType(&sql.NullBool{})
func (b *BoolBuilder) GoType(typ any) *BoolBuilder {
	b.desc.goType(typ)
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
//
//	field.Bool("deleted").
//		Annotations(
//			entgql.OrderField("DELETED"),
//		)
func (b *BoolBuilder) Annotations(annotations ...schema.Annotation) *BoolBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *BoolBuilder) Descriptor() *Descriptor {
	b.desc.checkGoType(boolType)
	return b.desc
}

// BytesBuilder is the builder for bytes fields.
type BytesBuilder struct {
	desc *Descriptor
}

// Default sets the default value of the field.
func (b *BytesBuilder) Default(v []byte) *BytesBuilder {
	b.desc.Default = v
	return b
}

// DefaultFunc sets the function that is applied to set the default value
// of the field on creation. For example:
//
//	field.Bytes("cuid").
//		DefaultFunc(cuid.New)
func (b *BytesBuilder) DefaultFunc(fn any) *BytesBuilder {
	if t := reflect.TypeOf(fn); t.Kind() != reflect.Func {
		b.desc.Err = fmt.Errorf("field.Bytes(%q).DefaultFunc expects func but got %s", b.desc.Name, t.Kind())
	}
	b.desc.Default = fn
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated struct.
func (b *BytesBuilder) Nillable() *BytesBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *BytesBuilder) Optional() *BytesBuilder {
	b.desc.Optional = true
	return b
}

// Sensitive fields not printable and not serializable.
func (b *BytesBuilder) Sensitive() *BytesBuilder {
	b.desc.Sensitive = true
	return b
}

// Unique makes the field unique within all vertices of this type.
// Only supported in PostgreSQL.
func (b *BytesBuilder) Unique() *BytesBuilder {
	b.desc.Unique = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *BytesBuilder) Immutable() *BytesBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *BytesBuilder) Comment(c string) *BytesBuilder {
	b.desc.Comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *BytesBuilder) StructTag(s string) *BytesBuilder {
	b.desc.Tag = s
	return b
}

// MaxLen sets the max-length of the bytes type in the database.
// In MySQL, this affects the BLOB type (tiny 2^8-1, regular 2^16-1, medium 2^24-1, long 2^32-1).
// In SQLite, it does not have any effect on the type size, which is default to 1B bytes.
func (b *BytesBuilder) MaxLen(i int) *BytesBuilder {
	b.desc.Size = i
	b.desc.Validators = append(b.desc.Validators, func(buf []byte) error {
		if len(buf) > i {
			return errors.New("value is greater than the required length")
		}
		return nil
	})
	return b
}

// MinLen adds a length validator for this field.
// Operation fails if the length of the buffer is less than the given value.
func (b *BytesBuilder) MinLen(i int) *BytesBuilder {
	b.desc.Validators = append(b.desc.Validators, func(b []byte) error {
		if len(b) < i {
			return errors.New("value is less than the required length")
		}
		return nil
	})
	return b
}

// NotEmpty adds a length validator for this field.
// Operation fails if the length of the buffer is zero.
func (b *BytesBuilder) NotEmpty() *BytesBuilder {
	return b.MinLen(1)
}

// Validate adds a validator for this field. Operation fails if the validation fails.
//
//	field.Bytes("blob").
//		Validate(func(b []byte) error {
//			if len(b) % 2 == 0 {
//				return fmt.Errorf("ent/schema: blob length is even: %d", len(b))
//			}
//			return nil
//		})
func (b *BytesBuilder) Validate(fn func([]byte) error) *BytesBuilder {
	b.desc.Validators = append(b.desc.Validators, fn)
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *BytesBuilder) StorageKey(key string) *BytesBuilder {
	b.desc.StorageKey = key
	return b
}

// GoType overrides the default Go type with a custom one.
// If the provided type implements the Validator interface
// and no validators have been set, the type validator will
// be used.
//
//	field.Bytes("ip").
//		GoType(net.IP("127.0.0.1"))
func (b *BytesBuilder) GoType(typ any) *BytesBuilder {
	b.desc.goType(typ)
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
func (b *BytesBuilder) Annotations(annotations ...schema.Annotation) *BytesBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for bytes.
//
//	field.Bytes("blob").
//		SchemaType(map[string]string{
//			dialect.MySQL:	"tinyblob",
//			dialect.SQLite:	"tinyblob",
//		})
func (b *BytesBuilder) SchemaType(types map[string]string) *BytesBuilder {
	b.desc.SchemaType = types
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *BytesBuilder) Descriptor() *Descriptor {
	if b.desc.Default != nil {
		b.desc.checkDefaultFunc(bytesType)
	}
	b.desc.checkGoType(bytesType)
	return b.desc
}

// JsonBuilder is the builder for json fields.
type JsonBuilder struct {
	desc *Descriptor
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *JsonBuilder) StorageKey(key string) *JsonBuilder {
	b.desc.StorageKey = key
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *JsonBuilder) Optional() *JsonBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *JsonBuilder) Immutable() *JsonBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *JsonBuilder) Comment(c string) *JsonBuilder {
	b.desc.Comment = c
	return b
}

// Sensitive fields not printable and not serializable.
func (b *JsonBuilder) Sensitive() *JsonBuilder {
	b.desc.Sensitive = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *JsonBuilder) StructTag(s string) *JsonBuilder {
	b.desc.Tag = s
	return b
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for json.
//
//	field.JSON("json").
//		SchemaType(map[string]string{
//			dialect.MySQL:		"json",
//			dialect.Postgres:	"jsonb",
//		})
func (b *JsonBuilder) SchemaType(types map[string]string) *JsonBuilder {
	b.desc.SchemaType = types
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
func (b *JsonBuilder) Annotations(annotations ...schema.Annotation) *JsonBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Default sets the default value of the field. For example:
//
//	field.JSON("dirs", []http.Dir{}).
//		// A static default value.
//		Default([]http.Dir{"/tmp"})
//
//	field.JSON("dirs", []http.Dir{}).
//		// A function for generating the default value.
//		Default(DefaultDirs)
func (b *JsonBuilder) Default(v any) *JsonBuilder {
	b.desc.Default = v
	switch fieldT, defaultT := b.desc.Info.RType.rtype, reflect.TypeOf(v); {
	case fieldT == defaultT:
	case defaultT.Kind() == reflect.Func:
		b.desc.checkDefaultFunc(b.desc.Info.RType.rtype)
	default:
		b.desc.Err = fmt.Errorf("expect type (func() %[1]s) or (%[1]s) for other default value", b.desc.Info)
	}
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *JsonBuilder) Descriptor() *Descriptor {
	return b.desc
}

// EnumBuilder is the builder for enum fields.
type EnumBuilder struct {
	desc *Descriptor
}

// Values adds given values to the enum values.
//
//	field.Enum("priority").
//		Values("low", "mid", "high")
func (b *EnumBuilder) Values(values ...string) *EnumBuilder {
	for _, v := range values {
		b.desc.Enums = append(b.desc.Enums, struct{ N, V string }{N: v, V: v})
	}
	return b
}

// NamedValues adds the given name, value pairs to the enum value.
// The "name" defines the Go identifier of the enum, and the value
// defines the actual value in the database.
//
// NamedValues returns an error if given an odd number of arguments.
//
//	field.Enum("priority").
//		NamedValues(
//			"Low", "LOW",
//			"Mid", "MID",
//			"High", "HIGH",
//		)
func (b *EnumBuilder) NamedValues(namevalue ...string) *EnumBuilder {
	if len(namevalue)%2 == 1 {
		b.desc.Err = fmt.Errorf("Enum.NamedValues: odd argument count")
		return b
	}
	for i := 0; i < len(namevalue); i += 2 {
		b.desc.Enums = append(b.desc.Enums, struct{ N, V string }{N: namevalue[i], V: namevalue[i+1]})
	}
	return b
}

// Default sets the default value of the field.
func (b *EnumBuilder) Default(value string) *EnumBuilder {
	b.desc.Default = value
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *EnumBuilder) StorageKey(key string) *EnumBuilder {
	b.desc.StorageKey = key
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *EnumBuilder) Optional() *EnumBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *EnumBuilder) Immutable() *EnumBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *EnumBuilder) Comment(c string) *EnumBuilder {
	b.desc.Comment = c
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated struct.
func (b *EnumBuilder) Nillable() *EnumBuilder {
	b.desc.Nillable = true
	return b
}

// StructTag sets the struct tag of the field.
func (b *EnumBuilder) StructTag(s string) *EnumBuilder {
	b.desc.Tag = s
	return b
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for enum.
//
//	field.Enum("enum").
//		SchemaType(map[string]string{
//			dialect.Postgres: "EnumType",
//		})
func (b *EnumBuilder) SchemaType(types map[string]string) *EnumBuilder {
	b.desc.SchemaType = types
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
//
//	field.Enum("enum").
//		Annotations(
//			entgql.OrderField("ENUM"),
//		)
func (b *EnumBuilder) Annotations(annotations ...schema.Annotation) *EnumBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// EnumValues defines the interface for getting the enum values.
type EnumValues interface {
	Values() []string
}

// GoType overrides the default Go type with a custom one.
// If the provided type implements the Validator interface
// and no validators have been set, the type validator will
// be used.
//
//	field.Enum("enum").
//		GoType(role.Enum("role"))
func (b *EnumBuilder) GoType(ev EnumValues) *EnumBuilder {
	b.Values(ev.Values()...)
	b.desc.goType(ev)
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *EnumBuilder) Descriptor() *Descriptor {
	if b.desc.Info.RType != nil {
		// If an error already exists, let that be returned instead.
		// Otherwise, check that the underlying type is either a string or implements Stringer.
		if b.desc.Err == nil && b.desc.Info.RType.rtype.Kind() != reflect.String && !b.desc.Info.Stringer() {
			b.desc.Err = errors.New("enum values which implement ValueScanner must also implement Stringer")
		}
		b.desc.checkGoType(stringType)
	}
	return b.desc
}

// UuidBuilder is the builder for uuid fields.
type UuidBuilder struct {
	desc *Descriptor
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *UuidBuilder) StorageKey(key string) *UuidBuilder {
	b.desc.StorageKey = key
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated struct.
func (b *UuidBuilder) Nillable() *UuidBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *UuidBuilder) Optional() *UuidBuilder {
	b.desc.Optional = true
	return b
}

// Unique makes the field unique within all vertices of this type.
func (b *UuidBuilder) Unique() *UuidBuilder {
	b.desc.Unique = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *UuidBuilder) Immutable() *UuidBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *UuidBuilder) Comment(c string) *UuidBuilder {
	b.desc.Comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *UuidBuilder) StructTag(s string) *UuidBuilder {
	b.desc.Tag = s
	return b
}

// Default sets the function that is applied to set default value
// of the field on creation. Codegen fails if the default function
// doesn't return the same concrete that was set for the UUID type.
//
//	field.UUID("id", uuid.UUID{}).
//		Default(uuid.New)
func (b *UuidBuilder) Default(fn any) *UuidBuilder {
	typ := reflect.TypeOf(fn)
	if typ.Kind() != reflect.Func || typ.NumIn() != 0 || typ.NumOut() != 1 || typ.Out(0).String() != b.desc.Info.String() {
		b.desc.Err = fmt.Errorf("expect type (func() %s) for uuid default value", b.desc.Info)
	}
	b.desc.Default = fn
	return b
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for uuid.
//
//	field.UUID("id", uuid.New()).
//		SchemaType(map[string]string{
//			dialect.Postgres: "CustomUUID",
//		})
func (b *UuidBuilder) SchemaType(types map[string]string) *UuidBuilder {
	b.desc.SchemaType = types
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
//
//	field.UUID("id", uuid.New()).
//		Annotations(
//			entgql.OrderField("ID"),
//		)
func (b *UuidBuilder) Annotations(annotations ...schema.Annotation) *UuidBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *UuidBuilder) Descriptor() *Descriptor {
	b.desc.checkGoType(valueScannerType)
	return b.desc
}

// OtherBuilder is the builder for other fields.
type OtherBuilder struct {
	desc *Descriptor
}

// Unique makes the field unique within all vertices of this type.
func (b *OtherBuilder) Unique() *OtherBuilder {
	b.desc.Unique = true
	return b
}

// Sensitive fields not printable and not serializable.
func (b *OtherBuilder) Sensitive() *OtherBuilder {
	b.desc.Sensitive = true
	return b
}

// Default sets the default value of the field. For example:
//
//	field.Other("link", &Link{}).
//		SchemaType(map[string]string{
//			dialect.MySQL:    "text",
//			dialect.Postgres: "varchar",
//		}).
//		// A static default value.
//		Default(&Link{Addr: "0.0.0.0"})
//
//	field.Other("link", &Link{}).
//		SchemaType(map[string]string{
//			dialect.MySQL:    "text",
//			dialect.Postgres: "varchar",
//		}).
//		// A function for generating the default value.
//		Default(NewLink)
func (b *OtherBuilder) Default(v any) *OtherBuilder {
	b.desc.Default = v
	switch fieldT, defaultT := b.desc.Info.RType.rtype, reflect.TypeOf(v); {
	case fieldT == defaultT:
	case defaultT.Kind() == reflect.Func:
		b.desc.checkDefaultFunc(b.desc.Info.RType.rtype)
	default:
		b.desc.Err = fmt.Errorf("expect type (func() %[1]s) or (%[1]s) for other default value", b.desc.Info)
	}
	return b
}

// Nillable indicates that this field is a nillable.
// Unlike "Optional" only fields, "Nillable" fields are pointers in the generated field.
func (b *OtherBuilder) Nillable() *OtherBuilder {
	b.desc.Nillable = true
	return b
}

// Optional indicates that this field is optional on create.
// Unlike edges, fields are required by default.
func (b *OtherBuilder) Optional() *OtherBuilder {
	b.desc.Optional = true
	return b
}

// Immutable indicates that this field cannot be updated.
func (b *OtherBuilder) Immutable() *OtherBuilder {
	b.desc.Immutable = true
	return b
}

// Comment sets the comment of the field.
func (b *OtherBuilder) Comment(c string) *OtherBuilder {
	b.desc.Comment = c
	return b
}

// StructTag sets the struct tag of the field.
func (b *OtherBuilder) StructTag(s string) *OtherBuilder {
	b.desc.Tag = s
	return b
}

// StorageKey sets the storage key of the field.
// In SQL dialects is the column name and Gremlin is the property.
func (b *OtherBuilder) StorageKey(key string) *OtherBuilder {
	b.desc.StorageKey = key
	return b
}

// SchemaType overrides the default database type with a custom
// schema type (per dialect) for string.
//
//	field.Other("link", Link{}).
//		SchemaType(map[string]string{
//			dialect.MySQL:    "text",
//			dialect.Postgres: "varchar",
//		})
func (b *OtherBuilder) SchemaType(types map[string]string) *OtherBuilder {
	b.desc.SchemaType = types
	return b
}

// Annotations adds a list of annotations to the field object to be used by
// codegen extensions.
//
//	field.Other("link", &Link{}).
//		SchemaType(map[string]string{
//			dialect.MySQL:    "text",
//			dialect.Postgres: "varchar",
//		}).
//		Annotations(
//			entgql.OrderField("LINK"),
//		)
func (b *OtherBuilder) Annotations(annotations ...schema.Annotation) *OtherBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *OtherBuilder) Descriptor() *Descriptor {
	b.desc.checkGoType(valueScannerType)
	if len(b.desc.SchemaType) == 0 {
		b.desc.Err = fmt.Errorf("expect SchemaType to be set for other field")
	}
	return b.desc
}

// A Descriptor for field configuration.
type Descriptor struct {
	Tag           string                  // struct tag.
	Size          int                     // varchar size.
	Name          string                  // field name.
	Info          *TypeInfo               // field type info.
	ValueScanner  any                     // custom field codec.
	Unique        bool                    // unique index of field.
	Nillable      bool                    // nillable struct field.
	Optional      bool                    // nullable field in database.
	Immutable     bool                    // create only field.
	Default       any                     // default value on create.
	UpdateDefault any                     // default value on update.
	Validators    []any                   // validator functions.
	StorageKey    string                  // sql column or gremlin property.
	Enums         []struct{ N, V string } // enum values.
	Sensitive     bool                    // sensitive info string field.
	SchemaType    map[string]string       // override the schema type.
	Annotations   []schema.Annotation     // field annotations.
	Comment       string                  // field comment.
	Err           error
}

func (d *Descriptor) goType(typ any) {
	t := reflect.TypeOf(typ)
	tv := indirect(t)
	info := &TypeInfo{
		Type:    d.Info.Type,
		Ident:   t.String(),
		PkgPath: tv.PkgPath(),
		PkgName: pkgName(tv.String()),
		RType: &RType{
			rtype:   t,
			Kind:    t.Kind(),
			Name:    tv.Name(),
			Ident:   tv.String(),
			PkgPath: tv.PkgPath(),
			Methods: make(map[string]struct{ In, Out []*RType }, t.NumMethod()),
		},
	}
	methods(t, info.RType)
	switch t.Kind() {
	case reflect.Slice, reflect.Ptr, reflect.Map:
		info.Nillable = true
	}
	d.Info = info
}

func (d *Descriptor) checkGoType(expectType reflect.Type) {
	t := expectType
	if d.Info.RType != nil && d.Info.RType.rtype != nil {
		t = d.Info.RType.rtype
	}
	switch pt := reflect.PtrTo(t); {
	// An external ValueScanner.
	case d.ValueScanner != nil:
		vs := reflect.Indirect(reflect.ValueOf(d.ValueScanner)).Type()
		m1, ok1 := vs.MethodByName("Value")
		m2, ok2 := vs.MethodByName("ScanValue")
		m3, ok3 := vs.MethodByName("FromValue")
		switch {
		case !ok1, m1.Type.NumIn() != 2, m1.Type.In(1) != t,
			m1.Type.NumOut() != 2, m1.Type.Out(0) != valueType, m1.Type.Out(1) != errorType:
			d.Err = fmt.Errorf("ValueScanner must implement the Value method: func Value(%s) (driver.Valuer, error)", t)
		case !ok2, m2.Type.NumIn() != 1, m2.Type.NumOut() != 1, m2.Type.Out(0) != valueScannerType:
			d.Err = errors.New("ValueScanner must implement the ScanValue method: func ScanValue() field.ValueScanner")
		case !ok3, m3.Type.NumIn() != 2, m3.Type.In(1) != valueType, m3.Type.NumOut() != 2, m3.Type.Out(0) != t, m3.Type.Out(1) != errorType:
			d.Err = fmt.Errorf("ValueScanner must implement the FromValue method: func FromValue(driver.Valuer) (%s, error)", t)
		}
	// No GoType was provided.
	case d.Info.RType == nil:
	// A GoType without an external ValueScanner.
	case pt.Implements(valueScannerType), t.Implements(valueScannerType), t.Kind() == expectType.Kind() && t.ConvertibleTo(expectType):
	// There is a GoType, but it's not a ValueScanner.
	default:
		d.Err = fmt.Errorf("GoType must be a %q type, ValueScanner or provide an external ValueScanner", expectType)
	}
}

// pkgName returns the package name from a Go
// identifier with a package qualifier.
func pkgName(ident string) string {
	i := strings.LastIndexByte(ident, '.')
	if i == -1 {
		return ""
	}
	s := ident[:i]
	if i := strings.LastIndexAny(s, "]*"); i != -1 {
		s = s[i+1:]
	}
	return s
}

func methods(t reflect.Type, rtype *RType) {
	// For type T, add methods with
	// pointer receiver as well (*T).
	if t.Kind() != reflect.Ptr {
		t = reflect.PtrTo(t)
	}
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		m := t.Method(i)
		in := make([]*RType, m.Type.NumIn()-1)
		for j := range in {
			arg := m.Type.In(j + 1)
			in[j] = &RType{Name: arg.Name(), Ident: arg.String(), Kind: arg.Kind(), PkgPath: arg.PkgPath()}
		}
		out := make([]*RType, m.Type.NumOut())
		for j := range out {
			ret := m.Type.Out(j)
			out[j] = &RType{Name: ret.Name(), Ident: ret.String(), Kind: ret.Kind(), PkgPath: ret.PkgPath()}
		}
		rtype.Methods[m.Name] = struct{ In, Out []*RType }{in, out}
	}
}

func (d *Descriptor) checkDefaultFunc(expectType reflect.Type) {
	for _, typ := range []reflect.Type{reflect.TypeOf(d.Default), reflect.TypeOf(d.UpdateDefault)} {
		if typ == nil || typ.Kind() != reflect.Func || d.Err != nil {
			continue
		}
		err := fmt.Errorf("expect type (func() %s) for default value", d.Info)
		if typ.NumIn() != 0 || typ.NumOut() != 1 {
			d.Err = err
		}
		rtype := expectType
		if d.Info.RType != nil {
			rtype = d.Info.RType.rtype
		}
		if !typ.Out(0).AssignableTo(rtype) {
			d.Err = err
		}
	}
}

var (
	boolType         = reflect.TypeOf(false)
	bytesType        = reflect.TypeOf([]byte(nil))
	timeType         = reflect.TypeOf(time.Time{})
	stringType       = reflect.TypeOf("")
	valueType        = reflect.TypeOf((*driver.Value)(nil)).Elem()
	valuerType       = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
	errorType        = reflect.TypeOf((*error)(nil)).Elem()
	valueScannerType = reflect.TypeOf((*ValueScanner)(nil)).Elem()
	validatorType    = reflect.TypeOf((*Validator)(nil)).Elem()
)

// ValueScanner is the interface that groups the Value
// and the Scan methods implemented by custom Go types.
type ValueScanner interface {
	driver.Valuer
	sql.Scanner
}

// TypeValueScanner is the interface that groups all methods for
// attaching an external ValueScanner to a custom GoType.
type TypeValueScanner[T any] interface {
	// Value returns the driver.Valuer for the GoType.
	Value(T) (driver.Value, error)
	// ScanValue returns a new ValueScanner that functions as an
	// intermediate result between database value and GoType value.
	// For example, sql.NullString or sql.NullInt.
	ScanValue() ValueScanner
	// FromValue returns the field instance from the ScanValue
	// above after the database value was scanned.
	FromValue(driver.Value) (T, error)
}

// TextValueScanner returns a new TypeValueScanner that calls MarshalText
// for storing values in the database, and calls UnmarshalText for scanning
// database values into struct fields.
type TextValueScanner[T interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}] struct{}

// Value implements the TypeValueScanner.Value method.
func (TextValueScanner[T]) Value(v T) (driver.Value, error) {
	return v.MarshalText()
}

// ScanValue implements the TypeValueScanner.ScanValue method.
func (TextValueScanner[T]) ScanValue() ValueScanner {
	return &sql.NullString{}
}

// FromValue implements the TypeValueScanner.FromValue method.
func (TextValueScanner[T]) FromValue(v driver.Value) (tv T, err error) {
	s, ok := v.(*sql.NullString)
	if !ok {
		return tv, fmt.Errorf("unexpected input for FromValue: %T", v)
	}
	tv = newT(tv).(T)
	if s.Valid {
		err = tv.UnmarshalText([]byte(s.String))
	}
	return tv, err
}

// BinaryValueScanner returns a new TypeValueScanner that calls MarshalBinary
// for storing values in the database, and calls UnmarshalBinary for scanning
// database values into struct fields.
type BinaryValueScanner[T interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}] struct{}

// Value implements the TypeValueScanner.Value method.
func (BinaryValueScanner[T]) Value(v T) (driver.Value, error) {
	return v.MarshalBinary()
}

// ScanValue implements the TypeValueScanner.ScanValue method.
func (BinaryValueScanner[T]) ScanValue() ValueScanner {
	return &sql.NullString{}
}

// FromValue implements the TypeValueScanner.FromValue method.
func (BinaryValueScanner[T]) FromValue(v driver.Value) (tv T, err error) {
	s, ok := v.(*sql.NullString)
	if !ok {
		return tv, fmt.Errorf("unexpected input for FromValue: %T", v)
	}
	tv = newT(tv).(T)
	if s.Valid {
		err = tv.UnmarshalBinary([]byte(s.String))
	}
	return tv, err
}

// ValueScannerFunc is a wrapper for a function that implements the ValueScanner.
type ValueScannerFunc[T any, S ValueScanner] struct {
	V func(T) (driver.Value, error)
	S func(S) (T, error)
}

// Value implements the TypeValueScanner.Value method.
func (f ValueScannerFunc[T, S]) Value(t T) (driver.Value, error) {
	return f.V(t)
}

// ScanValue implements the TypeValueScanner.ScanValue method.
func (f ValueScannerFunc[T, S]) ScanValue() ValueScanner {
	var s S
	return newT(s).(S)
}

// FromValue implements the TypeValueScanner.FromValue method.
func (f ValueScannerFunc[T, S]) FromValue(v driver.Value) (tv T, err error) {
	s, ok := v.(S)
	if !ok {
		return tv, fmt.Errorf("unexpected input for FromValue: %T", v)
	}
	return f.S(s)
}

// newT ensures the type is initialized.
func newT(t any) any {
	if rt := reflect.TypeOf(t); rt.Kind() == reflect.Ptr {
		return reflect.New(rt.Elem()).Interface()
	}
	return t
}

// Validator interface wraps the Validate method. Custom GoTypes with
// this method will be validated when the entity is created or updated.
type Validator interface {
	Validate() error
}

// indirect returns the type at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func pkgPath(t reflect.Type) string {
	pkg := t.PkgPath()
	if pkg != "" {
		return pkg
	}
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map:
		return pkgPath(t.Elem())
	}
	return pkg
}
