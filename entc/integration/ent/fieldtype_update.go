// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/ent/fieldtype"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// FieldTypeUpdate is the builder for updating FieldType entities.
type FieldTypeUpdate struct {
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
	predicates     []ent.Predicate
}

// Where adds a new predicate for the builder.
func (ftu *FieldTypeUpdate) Where(ps ...ent.Predicate) *FieldTypeUpdate {
	ftu.predicates = append(ftu.predicates, ps...)
	return ftu
}

// SetInt sets the int field.
func (ftu *FieldTypeUpdate) SetInt(i int) *FieldTypeUpdate {
	ftu.int = &i
	return ftu
}

// SetInt8 sets the int8 field.
func (ftu *FieldTypeUpdate) SetInt8(i int8) *FieldTypeUpdate {
	ftu.int8 = &i
	return ftu
}

// SetInt16 sets the int16 field.
func (ftu *FieldTypeUpdate) SetInt16(i int16) *FieldTypeUpdate {
	ftu.int16 = &i
	return ftu
}

// SetInt32 sets the int32 field.
func (ftu *FieldTypeUpdate) SetInt32(i int32) *FieldTypeUpdate {
	ftu.int32 = &i
	return ftu
}

// SetInt64 sets the int64 field.
func (ftu *FieldTypeUpdate) SetInt64(i int64) *FieldTypeUpdate {
	ftu.int64 = &i
	return ftu
}

// SetOptionalInt sets the optional_int field.
func (ftu *FieldTypeUpdate) SetOptionalInt(i int) *FieldTypeUpdate {
	ftu.optional_int = &i
	return ftu
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt(i *int) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt(*i)
	}
	return ftu
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftu *FieldTypeUpdate) SetOptionalInt8(i int8) *FieldTypeUpdate {
	ftu.optional_int8 = &i
	return ftu
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt8(i *int8) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt8(*i)
	}
	return ftu
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftu *FieldTypeUpdate) SetOptionalInt16(i int16) *FieldTypeUpdate {
	ftu.optional_int16 = &i
	return ftu
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt16(i *int16) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt16(*i)
	}
	return ftu
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftu *FieldTypeUpdate) SetOptionalInt32(i int32) *FieldTypeUpdate {
	ftu.optional_int32 = &i
	return ftu
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt32(i *int32) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt32(*i)
	}
	return ftu
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftu *FieldTypeUpdate) SetOptionalInt64(i int64) *FieldTypeUpdate {
	ftu.optional_int64 = &i
	return ftu
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt64(i *int64) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt64(*i)
	}
	return ftu
}

// SetNullableInt sets the nullable_int field.
func (ftu *FieldTypeUpdate) SetNullableInt(i int) *FieldTypeUpdate {
	ftu.nullable_int = &i
	return ftu
}

// SetNillableNullableInt sets the nullable_int field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNullableInt(i *int) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNullableInt(*i)
	}
	return ftu
}

// SetNullableInt8 sets the nullable_int8 field.
func (ftu *FieldTypeUpdate) SetNullableInt8(i int8) *FieldTypeUpdate {
	ftu.nullable_int8 = &i
	return ftu
}

// SetNillableNullableInt8 sets the nullable_int8 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNullableInt8(i *int8) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNullableInt8(*i)
	}
	return ftu
}

// SetNullableInt16 sets the nullable_int16 field.
func (ftu *FieldTypeUpdate) SetNullableInt16(i int16) *FieldTypeUpdate {
	ftu.nullable_int16 = &i
	return ftu
}

// SetNillableNullableInt16 sets the nullable_int16 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNullableInt16(i *int16) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNullableInt16(*i)
	}
	return ftu
}

// SetNullableInt32 sets the nullable_int32 field.
func (ftu *FieldTypeUpdate) SetNullableInt32(i int32) *FieldTypeUpdate {
	ftu.nullable_int32 = &i
	return ftu
}

// SetNillableNullableInt32 sets the nullable_int32 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNullableInt32(i *int32) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNullableInt32(*i)
	}
	return ftu
}

// SetNullableInt64 sets the nullable_int64 field.
func (ftu *FieldTypeUpdate) SetNullableInt64(i int64) *FieldTypeUpdate {
	ftu.nullable_int64 = &i
	return ftu
}

// SetNillableNullableInt64 sets the nullable_int64 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNullableInt64(i *int64) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNullableInt64(*i)
	}
	return ftu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ftu *FieldTypeUpdate) Save(ctx context.Context) (int, error) {
	switch ftu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftu.sqlSave(ctx)
	case dialect.Neptune:
		vertices, err := ftu.gremlinSave(ctx)
		return len(vertices), err
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FieldTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FieldTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FieldTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftu *FieldTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(fieldtype.FieldID).From(sql.Table(fieldtype.Table))
	for _, p := range ftu.predicates {
		p.SQL(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ftu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := ftu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(fieldtype.Table).Where(sql.InInts(fieldtype.FieldID, ids...))
	)
	if ftu.int != nil {
		update = true
		builder.Set(fieldtype.FieldInt, *ftu.int)
	}
	if ftu.int8 != nil {
		update = true
		builder.Set(fieldtype.FieldInt8, *ftu.int8)
	}
	if ftu.int16 != nil {
		update = true
		builder.Set(fieldtype.FieldInt16, *ftu.int16)
	}
	if ftu.int32 != nil {
		update = true
		builder.Set(fieldtype.FieldInt32, *ftu.int32)
	}
	if ftu.int64 != nil {
		update = true
		builder.Set(fieldtype.FieldInt64, *ftu.int64)
	}
	if ftu.optional_int != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt, *ftu.optional_int)
	}
	if ftu.optional_int8 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt8, *ftu.optional_int8)
	}
	if ftu.optional_int16 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt16, *ftu.optional_int16)
	}
	if ftu.optional_int32 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt32, *ftu.optional_int32)
	}
	if ftu.optional_int64 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt64, *ftu.optional_int64)
	}
	if ftu.nullable_int != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt, *ftu.nullable_int)
	}
	if ftu.nullable_int8 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt8, *ftu.nullable_int8)
	}
	if ftu.nullable_int16 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt16, *ftu.nullable_int16)
	}
	if ftu.nullable_int32 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt32, *ftu.nullable_int32)
	}
	if ftu.nullable_int64 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt64, *ftu.nullable_int64)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (ftu *FieldTypeUpdate) gremlinSave(ctx context.Context) ([]*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftu.gremlin().Query()
	if err := ftu.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	var fts FieldTypes
	fts.config(ftu.config)
	if err := fts.FromResponse(res); err != nil {
		return nil, err
	}
	return fts, nil
}

func (ftu *FieldTypeUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(fieldtype.Label)
	for _, p := range ftu.predicates {
		p.Gremlin(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if ftu.int != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, *ftu.int)
	}
	if ftu.int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt8, *ftu.int8)
	}
	if ftu.int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt16, *ftu.int16)
	}
	if ftu.int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt32, *ftu.int32)
	}
	if ftu.int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, *ftu.int64)
	}
	if ftu.optional_int != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, *ftu.optional_int)
	}
	if ftu.optional_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, *ftu.optional_int8)
	}
	if ftu.optional_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, *ftu.optional_int16)
	}
	if ftu.optional_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, *ftu.optional_int32)
	}
	if ftu.optional_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, *ftu.optional_int64)
	}
	if ftu.nullable_int != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt, *ftu.nullable_int)
	}
	if ftu.nullable_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt8, *ftu.nullable_int8)
	}
	if ftu.nullable_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt16, *ftu.nullable_int16)
	}
	if ftu.nullable_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt32, *ftu.nullable_int32)
	}
	if ftu.nullable_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt64, *ftu.nullable_int64)
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FieldTypeUpdateOne is the builder for updating a single FieldType entity.
type FieldTypeUpdateOne struct {
	config
	id             string
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
func (ftuo *FieldTypeUpdateOne) SetInt(i int) *FieldTypeUpdateOne {
	ftuo.int = &i
	return ftuo
}

// SetInt8 sets the int8 field.
func (ftuo *FieldTypeUpdateOne) SetInt8(i int8) *FieldTypeUpdateOne {
	ftuo.int8 = &i
	return ftuo
}

// SetInt16 sets the int16 field.
func (ftuo *FieldTypeUpdateOne) SetInt16(i int16) *FieldTypeUpdateOne {
	ftuo.int16 = &i
	return ftuo
}

// SetInt32 sets the int32 field.
func (ftuo *FieldTypeUpdateOne) SetInt32(i int32) *FieldTypeUpdateOne {
	ftuo.int32 = &i
	return ftuo
}

// SetInt64 sets the int64 field.
func (ftuo *FieldTypeUpdateOne) SetInt64(i int64) *FieldTypeUpdateOne {
	ftuo.int64 = &i
	return ftuo
}

// SetOptionalInt sets the optional_int field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt(i int) *FieldTypeUpdateOne {
	ftuo.optional_int = &i
	return ftuo
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt(i *int) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt(*i)
	}
	return ftuo
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt8(i int8) *FieldTypeUpdateOne {
	ftuo.optional_int8 = &i
	return ftuo
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt8(i *int8) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt8(*i)
	}
	return ftuo
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt16(i int16) *FieldTypeUpdateOne {
	ftuo.optional_int16 = &i
	return ftuo
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt16(i *int16) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt16(*i)
	}
	return ftuo
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt32(i int32) *FieldTypeUpdateOne {
	ftuo.optional_int32 = &i
	return ftuo
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt32(i *int32) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt32(*i)
	}
	return ftuo
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt64(i int64) *FieldTypeUpdateOne {
	ftuo.optional_int64 = &i
	return ftuo
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt64(i *int64) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt64(*i)
	}
	return ftuo
}

// SetNullableInt sets the nullable_int field.
func (ftuo *FieldTypeUpdateOne) SetNullableInt(i int) *FieldTypeUpdateOne {
	ftuo.nullable_int = &i
	return ftuo
}

// SetNillableNullableInt sets the nullable_int field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNullableInt(i *int) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNullableInt(*i)
	}
	return ftuo
}

// SetNullableInt8 sets the nullable_int8 field.
func (ftuo *FieldTypeUpdateOne) SetNullableInt8(i int8) *FieldTypeUpdateOne {
	ftuo.nullable_int8 = &i
	return ftuo
}

// SetNillableNullableInt8 sets the nullable_int8 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNullableInt8(i *int8) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNullableInt8(*i)
	}
	return ftuo
}

// SetNullableInt16 sets the nullable_int16 field.
func (ftuo *FieldTypeUpdateOne) SetNullableInt16(i int16) *FieldTypeUpdateOne {
	ftuo.nullable_int16 = &i
	return ftuo
}

// SetNillableNullableInt16 sets the nullable_int16 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNullableInt16(i *int16) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNullableInt16(*i)
	}
	return ftuo
}

// SetNullableInt32 sets the nullable_int32 field.
func (ftuo *FieldTypeUpdateOne) SetNullableInt32(i int32) *FieldTypeUpdateOne {
	ftuo.nullable_int32 = &i
	return ftuo
}

// SetNillableNullableInt32 sets the nullable_int32 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNullableInt32(i *int32) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNullableInt32(*i)
	}
	return ftuo
}

// SetNullableInt64 sets the nullable_int64 field.
func (ftuo *FieldTypeUpdateOne) SetNullableInt64(i int64) *FieldTypeUpdateOne {
	ftuo.nullable_int64 = &i
	return ftuo
}

// SetNillableNullableInt64 sets the nullable_int64 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNullableInt64(i *int64) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNullableInt64(*i)
	}
	return ftuo
}

// Save executes the query and returns the updated entity.
func (ftuo *FieldTypeUpdateOne) Save(ctx context.Context) (*FieldType, error) {
	switch ftuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftuo.sqlSave(ctx)
	case dialect.Neptune:
		return ftuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FieldTypeUpdateOne) SaveX(ctx context.Context) *FieldType {
	ft, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// Exec executes the query on the entity.
func (ftuo *FieldTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FieldTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftuo *FieldTypeUpdateOne) sqlSave(ctx context.Context) (ft *FieldType, err error) {
	selector := sql.Select(fieldtype.Columns...).From(sql.Table(fieldtype.Table))
	fieldtype.ID(ftuo.id).SQL(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ftuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		ft = &FieldType{config: ftuo.config}
		if err := ft.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into FieldType: %v", err)
		}
		id = ft.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: FieldType not found with id: %v", ftuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one FieldType with the same id: %v", ftuo.id)
	}

	tx, err := ftuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(fieldtype.Table).Where(sql.InInts(fieldtype.FieldID, ids...))
	)
	if ftuo.int != nil {
		update = true
		builder.Set(fieldtype.FieldInt, *ftuo.int)
		ft.Int = *ftuo.int
	}
	if ftuo.int8 != nil {
		update = true
		builder.Set(fieldtype.FieldInt8, *ftuo.int8)
		ft.Int8 = *ftuo.int8
	}
	if ftuo.int16 != nil {
		update = true
		builder.Set(fieldtype.FieldInt16, *ftuo.int16)
		ft.Int16 = *ftuo.int16
	}
	if ftuo.int32 != nil {
		update = true
		builder.Set(fieldtype.FieldInt32, *ftuo.int32)
		ft.Int32 = *ftuo.int32
	}
	if ftuo.int64 != nil {
		update = true
		builder.Set(fieldtype.FieldInt64, *ftuo.int64)
		ft.Int64 = *ftuo.int64
	}
	if ftuo.optional_int != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt, *ftuo.optional_int)
		ft.OptionalInt = *ftuo.optional_int
	}
	if ftuo.optional_int8 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt8, *ftuo.optional_int8)
		ft.OptionalInt8 = *ftuo.optional_int8
	}
	if ftuo.optional_int16 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt16, *ftuo.optional_int16)
		ft.OptionalInt16 = *ftuo.optional_int16
	}
	if ftuo.optional_int32 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt32, *ftuo.optional_int32)
		ft.OptionalInt32 = *ftuo.optional_int32
	}
	if ftuo.optional_int64 != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt64, *ftuo.optional_int64)
		ft.OptionalInt64 = *ftuo.optional_int64
	}
	if ftuo.nullable_int != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt, *ftuo.nullable_int)
		*ft.NullableInt = *ftuo.nullable_int
	}
	if ftuo.nullable_int8 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt8, *ftuo.nullable_int8)
		*ft.NullableInt8 = *ftuo.nullable_int8
	}
	if ftuo.nullable_int16 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt16, *ftuo.nullable_int16)
		*ft.NullableInt16 = *ftuo.nullable_int16
	}
	if ftuo.nullable_int32 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt32, *ftuo.nullable_int32)
		*ft.NullableInt32 = *ftuo.nullable_int32
	}
	if ftuo.nullable_int64 != nil {
		update = true
		builder.Set(fieldtype.FieldNullableInt64, *ftuo.nullable_int64)
		*ft.NullableInt64 = *ftuo.nullable_int64
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FieldTypeUpdateOne) gremlinSave(ctx context.Context) (*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftuo.gremlin(ftuo.id).Query()
	if err := ftuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FieldType{config: ftuo.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FieldTypeUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if ftuo.int != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, *ftuo.int)
	}
	if ftuo.int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt8, *ftuo.int8)
	}
	if ftuo.int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt16, *ftuo.int16)
	}
	if ftuo.int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt32, *ftuo.int32)
	}
	if ftuo.int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, *ftuo.int64)
	}
	if ftuo.optional_int != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, *ftuo.optional_int)
	}
	if ftuo.optional_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, *ftuo.optional_int8)
	}
	if ftuo.optional_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, *ftuo.optional_int16)
	}
	if ftuo.optional_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, *ftuo.optional_int32)
	}
	if ftuo.optional_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, *ftuo.optional_int64)
	}
	if ftuo.nullable_int != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt, *ftuo.nullable_int)
	}
	if ftuo.nullable_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt8, *ftuo.nullable_int8)
	}
	if ftuo.nullable_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt16, *ftuo.nullable_int16)
	}
	if ftuo.nullable_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt32, *ftuo.nullable_int32)
	}
	if ftuo.nullable_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldNullableInt64, *ftuo.nullable_int64)
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
