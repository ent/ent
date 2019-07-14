// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
)

// FieldType is the model entity for the FieldType schema.
type FieldType struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Int holds the value of the "int" field.
	Int int `json:"int,omitempty"`
	// Int8 holds the value of the "int8" field.
	Int8 int8 `json:"int8,omitempty"`
	// Int16 holds the value of the "int16" field.
	Int16 int16 `json:"int16,omitempty"`
	// Int32 holds the value of the "int32" field.
	Int32 int32 `json:"int32,omitempty"`
	// Int64 holds the value of the "int64" field.
	Int64 int64 `json:"int64,omitempty"`
	// OptionalInt holds the value of the "optional_int" field.
	OptionalInt int `json:"optional_int,omitempty"`
	// OptionalInt8 holds the value of the "optional_int8" field.
	OptionalInt8 int8 `json:"optional_int8,omitempty"`
	// OptionalInt16 holds the value of the "optional_int16" field.
	OptionalInt16 int16 `json:"optional_int16,omitempty"`
	// OptionalInt32 holds the value of the "optional_int32" field.
	OptionalInt32 int32 `json:"optional_int32,omitempty"`
	// OptionalInt64 holds the value of the "optional_int64" field.
	OptionalInt64 int64 `json:"optional_int64,omitempty"`
	// NullableInt holds the value of the "nullable_int" field.
	NullableInt *int `json:"nullable_int,omitempty"`
	// NullableInt8 holds the value of the "nullable_int8" field.
	NullableInt8 *int8 `json:"nullable_int8,omitempty"`
	// NullableInt16 holds the value of the "nullable_int16" field.
	NullableInt16 *int16 `json:"nullable_int16,omitempty"`
	// NullableInt32 holds the value of the "nullable_int32" field.
	NullableInt32 *int32 `json:"nullable_int32,omitempty"`
	// NullableInt64 holds the value of the "nullable_int64" field.
	NullableInt64 *int64 `json:"nullable_int64,omitempty"`
}

// FromResponse scans the gremlin response data into FieldType.
func (ft *FieldType) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vft struct {
		ID            string `json:"id,omitempty"`
		Int           int    `json:"int,omitempty"`
		Int8          int8   `json:"int8,omitempty"`
		Int16         int16  `json:"int16,omitempty"`
		Int32         int32  `json:"int32,omitempty"`
		Int64         int64  `json:"int64,omitempty"`
		OptionalInt   int    `json:"optional_int,omitempty"`
		OptionalInt8  int8   `json:"optional_int8,omitempty"`
		OptionalInt16 int16  `json:"optional_int16,omitempty"`
		OptionalInt32 int32  `json:"optional_int32,omitempty"`
		OptionalInt64 int64  `json:"optional_int64,omitempty"`
		NullableInt   *int   `json:"nullable_int,omitempty"`
		NullableInt8  *int8  `json:"nullable_int8,omitempty"`
		NullableInt16 *int16 `json:"nullable_int16,omitempty"`
		NullableInt32 *int32 `json:"nullable_int32,omitempty"`
		NullableInt64 *int64 `json:"nullable_int64,omitempty"`
	}
	if err := vmap.Decode(&vft); err != nil {
		return err
	}
	ft.ID = vft.ID
	ft.Int = vft.Int
	ft.Int8 = vft.Int8
	ft.Int16 = vft.Int16
	ft.Int32 = vft.Int32
	ft.Int64 = vft.Int64
	ft.OptionalInt = vft.OptionalInt
	ft.OptionalInt8 = vft.OptionalInt8
	ft.OptionalInt16 = vft.OptionalInt16
	ft.OptionalInt32 = vft.OptionalInt32
	ft.OptionalInt64 = vft.OptionalInt64
	ft.NullableInt = vft.NullableInt
	ft.NullableInt8 = vft.NullableInt8
	ft.NullableInt16 = vft.NullableInt16
	ft.NullableInt32 = vft.NullableInt32
	ft.NullableInt64 = vft.NullableInt64
	return nil
}

// FromRows scans the sql response data into FieldType.
func (ft *FieldType) FromRows(rows *sql.Rows) error {
	var vft struct {
		ID            int
		Int           int
		Int8          int8
		Int16         int16
		Int32         int32
		Int64         int64
		OptionalInt   sql.NullInt64
		OptionalInt8  sql.NullInt64
		OptionalInt16 sql.NullInt64
		OptionalInt32 sql.NullInt64
		OptionalInt64 sql.NullInt64
		NullableInt   sql.NullInt64
		NullableInt8  sql.NullInt64
		NullableInt16 sql.NullInt64
		NullableInt32 sql.NullInt64
		NullableInt64 sql.NullInt64
	}
	// the order here should be the same as in the `fieldtype.Columns`.
	if err := rows.Scan(
		&vft.ID,
		&vft.Int,
		&vft.Int8,
		&vft.Int16,
		&vft.Int32,
		&vft.Int64,
		&vft.OptionalInt,
		&vft.OptionalInt8,
		&vft.OptionalInt16,
		&vft.OptionalInt32,
		&vft.OptionalInt64,
		&vft.NullableInt,
		&vft.NullableInt8,
		&vft.NullableInt16,
		&vft.NullableInt32,
		&vft.NullableInt64,
	); err != nil {
		return err
	}
	ft.ID = strconv.Itoa(vft.ID)
	ft.Int = vft.Int
	ft.Int8 = vft.Int8
	ft.Int16 = vft.Int16
	ft.Int32 = vft.Int32
	ft.Int64 = vft.Int64
	ft.OptionalInt = int(vft.OptionalInt.Int64)
	ft.OptionalInt8 = int8(vft.OptionalInt8.Int64)
	ft.OptionalInt16 = int16(vft.OptionalInt16.Int64)
	ft.OptionalInt32 = int32(vft.OptionalInt32.Int64)
	ft.OptionalInt64 = vft.OptionalInt64.Int64
	if vft.NullableInt.Valid {
		ft.NullableInt = new(int)
		*ft.NullableInt = int(vft.NullableInt.Int64)
	}
	if vft.NullableInt8.Valid {
		ft.NullableInt8 = new(int8)
		*ft.NullableInt8 = int8(vft.NullableInt8.Int64)
	}
	if vft.NullableInt16.Valid {
		ft.NullableInt16 = new(int16)
		*ft.NullableInt16 = int16(vft.NullableInt16.Int64)
	}
	if vft.NullableInt32.Valid {
		ft.NullableInt32 = new(int32)
		*ft.NullableInt32 = int32(vft.NullableInt32.Int64)
	}
	if vft.NullableInt64.Valid {
		ft.NullableInt64 = new(int64)
		*ft.NullableInt64 = vft.NullableInt64.Int64
	}
	return nil
}

// Update returns a builder for updating this FieldType.
// Note that, you need to call FieldType.Unwrap() before calling this method, if this FieldType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FieldType) Update() *FieldTypeUpdateOne {
	return (&FieldTypeClient{ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ft *FieldType) Unwrap() *FieldType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FieldType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FieldType) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("FieldType(")
	buf.WriteString(fmt.Sprintf("id=%v", ft.ID))
	buf.WriteString(fmt.Sprintf(", int=%v", ft.Int))
	buf.WriteString(fmt.Sprintf(", int8=%v", ft.Int8))
	buf.WriteString(fmt.Sprintf(", int16=%v", ft.Int16))
	buf.WriteString(fmt.Sprintf(", int32=%v", ft.Int32))
	buf.WriteString(fmt.Sprintf(", int64=%v", ft.Int64))
	buf.WriteString(fmt.Sprintf(", optional_int=%v", ft.OptionalInt))
	buf.WriteString(fmt.Sprintf(", optional_int8=%v", ft.OptionalInt8))
	buf.WriteString(fmt.Sprintf(", optional_int16=%v", ft.OptionalInt16))
	buf.WriteString(fmt.Sprintf(", optional_int32=%v", ft.OptionalInt32))
	buf.WriteString(fmt.Sprintf(", optional_int64=%v", ft.OptionalInt64))
	if v := ft.NullableInt; v != nil {
		buf.WriteString(fmt.Sprintf(", nullable_int=%v", *v))
	}
	if v := ft.NullableInt8; v != nil {
		buf.WriteString(fmt.Sprintf(", nullable_int8=%v", *v))
	}
	if v := ft.NullableInt16; v != nil {
		buf.WriteString(fmt.Sprintf(", nullable_int16=%v", *v))
	}
	if v := ft.NullableInt32; v != nil {
		buf.WriteString(fmt.Sprintf(", nullable_int32=%v", *v))
	}
	if v := ft.NullableInt64; v != nil {
		buf.WriteString(fmt.Sprintf(", nullable_int64=%v", *v))
	}
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (ft *FieldType) id() int {
	id, _ := strconv.Atoi(ft.ID)
	return id
}

// FieldTypes is a parsable slice of FieldType.
type FieldTypes []*FieldType

// FromResponse scans the gremlin response data into FieldTypes.
func (ft *FieldTypes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vft []struct {
		ID            string `json:"id,omitempty"`
		Int           int    `json:"int,omitempty"`
		Int8          int8   `json:"int8,omitempty"`
		Int16         int16  `json:"int16,omitempty"`
		Int32         int32  `json:"int32,omitempty"`
		Int64         int64  `json:"int64,omitempty"`
		OptionalInt   int    `json:"optional_int,omitempty"`
		OptionalInt8  int8   `json:"optional_int8,omitempty"`
		OptionalInt16 int16  `json:"optional_int16,omitempty"`
		OptionalInt32 int32  `json:"optional_int32,omitempty"`
		OptionalInt64 int64  `json:"optional_int64,omitempty"`
		NullableInt   *int   `json:"nullable_int,omitempty"`
		NullableInt8  *int8  `json:"nullable_int8,omitempty"`
		NullableInt16 *int16 `json:"nullable_int16,omitempty"`
		NullableInt32 *int32 `json:"nullable_int32,omitempty"`
		NullableInt64 *int64 `json:"nullable_int64,omitempty"`
	}
	if err := vmap.Decode(&vft); err != nil {
		return err
	}
	for _, v := range vft {
		*ft = append(*ft, &FieldType{
			ID:            v.ID,
			Int:           v.Int,
			Int8:          v.Int8,
			Int16:         v.Int16,
			Int32:         v.Int32,
			Int64:         v.Int64,
			OptionalInt:   v.OptionalInt,
			OptionalInt8:  v.OptionalInt8,
			OptionalInt16: v.OptionalInt16,
			OptionalInt32: v.OptionalInt32,
			OptionalInt64: v.OptionalInt64,
			NullableInt:   v.NullableInt,
			NullableInt8:  v.NullableInt8,
			NullableInt16: v.NullableInt16,
			NullableInt32: v.NullableInt32,
			NullableInt64: v.NullableInt64,
		})
	}
	return nil
}

// FromRows scans the sql response data into FieldTypes.
func (ft *FieldTypes) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vft := &FieldType{}
		if err := vft.FromRows(rows); err != nil {
			return err
		}
		*ft = append(*ft, vft)
	}
	return nil
}

func (ft FieldTypes) config(cfg config) {
	for i := range ft {
		ft[i].config = cfg
	}
}
