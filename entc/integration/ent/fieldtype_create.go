// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"fbc/ent/entc/integration/ent/fieldtype"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	int            *int
	int8           *int8
	int16          *int16
	int32          *int32
	int64          *int64
	optional_int   *int
	optional_int8  *int8
	optional_int16 *int16
	optional_int32 *int32
	optional_int64 *int64
	nullable_int   *int
	nullable_int8  *int8
	nullable_int16 *int16
	nullable_int32 *int32
	nullable_int64 *int64
}

// SetInt sets the int field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.int = &i
	return ftc
}

// SetInt8 sets the int8 field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.int8 = &i
	return ftc
}

// SetInt16 sets the int16 field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.int16 = &i
	return ftc
}

// SetInt32 sets the int32 field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.int32 = &i
	return ftc
}

// SetInt64 sets the int64 field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.int64 = &i
	return ftc
}

// SetOptionalInt sets the optional_int field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.optional_int = &i
	return ftc
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.optional_int8 = &i
	return ftc
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.optional_int16 = &i
	return ftc
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.optional_int32 = &i
	return ftc
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.optional_int64 = &i
	return ftc
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNullableInt sets the nullable_int field.
func (ftc *FieldTypeCreate) SetNullableInt(i int) *FieldTypeCreate {
	ftc.nullable_int = &i
	return ftc
}

// SetNillableNullableInt sets the nullable_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNullableInt(*i)
	}
	return ftc
}

// SetNullableInt8 sets the nullable_int8 field.
func (ftc *FieldTypeCreate) SetNullableInt8(i int8) *FieldTypeCreate {
	ftc.nullable_int8 = &i
	return ftc
}

// SetNillableNullableInt8 sets the nullable_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNullableInt8(*i)
	}
	return ftc
}

// SetNullableInt16 sets the nullable_int16 field.
func (ftc *FieldTypeCreate) SetNullableInt16(i int16) *FieldTypeCreate {
	ftc.nullable_int16 = &i
	return ftc
}

// SetNillableNullableInt16 sets the nullable_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNullableInt16(*i)
	}
	return ftc
}

// SetNullableInt32 sets the nullable_int32 field.
func (ftc *FieldTypeCreate) SetNullableInt32(i int32) *FieldTypeCreate {
	ftc.nullable_int32 = &i
	return ftc
}

// SetNillableNullableInt32 sets the nullable_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNullableInt32(*i)
	}
	return ftc
}

// SetNullableInt64 sets the nullable_int64 field.
func (ftc *FieldTypeCreate) SetNullableInt64(i int64) *FieldTypeCreate {
	ftc.nullable_int64 = &i
	return ftc
}

// SetNillableNullableInt64 sets the nullable_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNullableInt64(*i)
	}
	return ftc
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	if ftc.int == nil {
		return nil, errors.New("ent: missing required field \"int\"")
	}
	if ftc.int8 == nil {
		return nil, errors.New("ent: missing required field \"int8\"")
	}
	if ftc.int16 == nil {
		return nil, errors.New("ent: missing required field \"int16\"")
	}
	if ftc.int32 == nil {
		return nil, errors.New("ent: missing required field \"int32\"")
	}
	if ftc.int64 == nil {
		return nil, errors.New("ent: missing required field \"int64\"")
	}
	switch ftc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftc.sqlSave(ctx)
	case dialect.Neptune:
		return ftc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FieldTypeCreate) sqlSave(ctx context.Context) (*FieldType, error) {
	var (
		res sql.Result
		ft  = &FieldType{config: ftc.config}
	)
	tx, err := ftc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(fieldtype.Table).Default(ftc.driver.Dialect())
	if ftc.int != nil {
		builder.Set(fieldtype.FieldInt, *ftc.int)
		ft.Int = *ftc.int
	}
	if ftc.int8 != nil {
		builder.Set(fieldtype.FieldInt8, *ftc.int8)
		ft.Int8 = *ftc.int8
	}
	if ftc.int16 != nil {
		builder.Set(fieldtype.FieldInt16, *ftc.int16)
		ft.Int16 = *ftc.int16
	}
	if ftc.int32 != nil {
		builder.Set(fieldtype.FieldInt32, *ftc.int32)
		ft.Int32 = *ftc.int32
	}
	if ftc.int64 != nil {
		builder.Set(fieldtype.FieldInt64, *ftc.int64)
		ft.Int64 = *ftc.int64
	}
	if ftc.optional_int != nil {
		builder.Set(fieldtype.FieldOptionalInt, *ftc.optional_int)
		ft.OptionalInt = *ftc.optional_int
	}
	if ftc.optional_int8 != nil {
		builder.Set(fieldtype.FieldOptionalInt8, *ftc.optional_int8)
		ft.OptionalInt8 = *ftc.optional_int8
	}
	if ftc.optional_int16 != nil {
		builder.Set(fieldtype.FieldOptionalInt16, *ftc.optional_int16)
		ft.OptionalInt16 = *ftc.optional_int16
	}
	if ftc.optional_int32 != nil {
		builder.Set(fieldtype.FieldOptionalInt32, *ftc.optional_int32)
		ft.OptionalInt32 = *ftc.optional_int32
	}
	if ftc.optional_int64 != nil {
		builder.Set(fieldtype.FieldOptionalInt64, *ftc.optional_int64)
		ft.OptionalInt64 = *ftc.optional_int64
	}
	if ftc.nullable_int != nil {
		builder.Set(fieldtype.FieldNullableInt, *ftc.nullable_int)
		ft.NullableInt = ftc.nullable_int
	}
	if ftc.nullable_int8 != nil {
		builder.Set(fieldtype.FieldNullableInt8, *ftc.nullable_int8)
		ft.NullableInt8 = ftc.nullable_int8
	}
	if ftc.nullable_int16 != nil {
		builder.Set(fieldtype.FieldNullableInt16, *ftc.nullable_int16)
		ft.NullableInt16 = ftc.nullable_int16
	}
	if ftc.nullable_int32 != nil {
		builder.Set(fieldtype.FieldNullableInt32, *ftc.nullable_int32)
		ft.NullableInt32 = ftc.nullable_int32
	}
	if ftc.nullable_int64 != nil {
		builder.Set(fieldtype.FieldNullableInt64, *ftc.nullable_int64)
		ft.NullableInt64 = ftc.nullable_int64
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	ft.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftc *FieldTypeCreate) gremlinSave(ctx context.Context) (*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftc.gremlin().Query()
	if err := ftc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FieldType{config: ftc.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftc *FieldTypeCreate) gremlin() *dsl.Traversal {
	v := g.AddV(fieldtype.Label)
	if ftc.int != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, *ftc.int)
	}
	if ftc.int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt8, *ftc.int8)
	}
	if ftc.int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt16, *ftc.int16)
	}
	if ftc.int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt32, *ftc.int32)
	}
	if ftc.int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, *ftc.int64)
	}
	if ftc.optional_int != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, *ftc.optional_int)
	}
	if ftc.optional_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, *ftc.optional_int8)
	}
	if ftc.optional_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, *ftc.optional_int16)
	}
	if ftc.optional_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, *ftc.optional_int32)
	}
	if ftc.optional_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, *ftc.optional_int64)
	}
	if ftc.nullable_int != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt, *ftc.nullable_int)
	}
	if ftc.nullable_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt8, *ftc.nullable_int8)
	}
	if ftc.nullable_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt16, *ftc.nullable_int16)
	}
	if ftc.nullable_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt32, *ftc.nullable_int32)
	}
	if ftc.nullable_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt64, *ftc.nullable_int64)
	}
	return v.ValueMap(true)
}
