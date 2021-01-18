// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/fieldtype"
	"github.com/facebook/ent/entc/integration/ent/role"
	"github.com/facebook/ent/entc/integration/ent/schema"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	mutation *FieldTypeMutation
	hooks    []Hook
}

// SetInt sets the "int" field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.mutation.SetInt(i)
	return ftc
}

// SetInt8 sets the "int8" field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetInt8(i)
	return ftc
}

// SetInt16 sets the "int16" field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetInt16(i)
	return ftc
}

// SetInt32 sets the "int32" field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetInt32(i)
	return ftc
}

// SetInt64 sets the "int64" field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetInt64(i)
	return ftc
}

// SetOptionalInt sets the "optional_int" field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt(i)
	return ftc
}

// SetNillableOptionalInt sets the "optional_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the "optional_int8" field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt8(i)
	return ftc
}

// SetNillableOptionalInt8 sets the "optional_int8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the "optional_int16" field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt16(i)
	return ftc
}

// SetNillableOptionalInt16 sets the "optional_int16" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the "optional_int32" field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt32(i)
	return ftc
}

// SetNillableOptionalInt32 sets the "optional_int32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the "optional_int64" field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt64(i)
	return ftc
}

// SetNillableOptionalInt64 sets the "optional_int64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNillableInt sets the "nillable_int" field.
func (ftc *FieldTypeCreate) SetNillableInt(i int) *FieldTypeCreate {
	ftc.mutation.SetNillableInt(i)
	return ftc
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt(*i)
	}
	return ftc
}

// SetNillableInt8 sets the "nillable_int8" field.
func (ftc *FieldTypeCreate) SetNillableInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetNillableInt8(i)
	return ftc
}

// SetNillableNillableInt8 sets the "nillable_int8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt8(*i)
	}
	return ftc
}

// SetNillableInt16 sets the "nillable_int16" field.
func (ftc *FieldTypeCreate) SetNillableInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetNillableInt16(i)
	return ftc
}

// SetNillableNillableInt16 sets the "nillable_int16" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt16(*i)
	}
	return ftc
}

// SetNillableInt32 sets the "nillable_int32" field.
func (ftc *FieldTypeCreate) SetNillableInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetNillableInt32(i)
	return ftc
}

// SetNillableNillableInt32 sets the "nillable_int32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt32(*i)
	}
	return ftc
}

// SetNillableInt64 sets the "nillable_int64" field.
func (ftc *FieldTypeCreate) SetNillableInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetNillableInt64(i)
	return ftc
}

// SetNillableNillableInt64 sets the "nillable_int64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt64(*i)
	}
	return ftc
}

// SetValidateOptionalInt32 sets the "validate_optional_int32" field.
func (ftc *FieldTypeCreate) SetValidateOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetValidateOptionalInt32(i)
	return ftc
}

// SetNillableValidateOptionalInt32 sets the "validate_optional_int32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetValidateOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalUint sets the "optional_uint" field.
func (ftc *FieldTypeCreate) SetOptionalUint(u uint) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint(u)
	return ftc
}

// SetNillableOptionalUint sets the "optional_uint" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint(u *uint) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint(*u)
	}
	return ftc
}

// SetOptionalUint8 sets the "optional_uint8" field.
func (ftc *FieldTypeCreate) SetOptionalUint8(u uint8) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint8(u)
	return ftc
}

// SetNillableOptionalUint8 sets the "optional_uint8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint8(u *uint8) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint8(*u)
	}
	return ftc
}

// SetOptionalUint16 sets the "optional_uint16" field.
func (ftc *FieldTypeCreate) SetOptionalUint16(u uint16) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint16(u)
	return ftc
}

// SetNillableOptionalUint16 sets the "optional_uint16" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint16(u *uint16) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint16(*u)
	}
	return ftc
}

// SetOptionalUint32 sets the "optional_uint32" field.
func (ftc *FieldTypeCreate) SetOptionalUint32(u uint32) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint32(u)
	return ftc
}

// SetNillableOptionalUint32 sets the "optional_uint32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint32(u *uint32) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint32(*u)
	}
	return ftc
}

// SetOptionalUint64 sets the "optional_uint64" field.
func (ftc *FieldTypeCreate) SetOptionalUint64(u uint64) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint64(u)
	return ftc
}

// SetNillableOptionalUint64 sets the "optional_uint64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint64(u *uint64) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint64(*u)
	}
	return ftc
}

// SetDuration sets the "duration" field.
func (ftc *FieldTypeCreate) SetDuration(t time.Duration) *FieldTypeCreate {
	ftc.mutation.SetDuration(t)
	return ftc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDuration(t *time.Duration) *FieldTypeCreate {
	if t != nil {
		ftc.SetDuration(*t)
	}
	return ftc
}

// SetState sets the "state" field.
func (ftc *FieldTypeCreate) SetState(f fieldtype.State) *FieldTypeCreate {
	ftc.mutation.SetState(f)
	return ftc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableState(f *fieldtype.State) *FieldTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// SetOptionalFloat sets the "optional_float" field.
func (ftc *FieldTypeCreate) SetOptionalFloat(f float64) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat(f)
	return ftc
}

// SetNillableOptionalFloat sets the "optional_float" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat(*f)
	}
	return ftc
}

// SetOptionalFloat32 sets the "optional_float32" field.
func (ftc *FieldTypeCreate) SetOptionalFloat32(f float32) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat32(f)
	return ftc
}

// SetNillableOptionalFloat32 sets the "optional_float32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat32(f *float32) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat32(*f)
	}
	return ftc
}

// SetDatetime sets the "datetime" field.
func (ftc *FieldTypeCreate) SetDatetime(t time.Time) *FieldTypeCreate {
	ftc.mutation.SetDatetime(t)
	return ftc
}

// SetNillableDatetime sets the "datetime" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDatetime(t *time.Time) *FieldTypeCreate {
	if t != nil {
		ftc.SetDatetime(*t)
	}
	return ftc
}

// SetDecimal sets the "decimal" field.
func (ftc *FieldTypeCreate) SetDecimal(f float64) *FieldTypeCreate {
	ftc.mutation.SetDecimal(f)
	return ftc
}

// SetNillableDecimal sets the "decimal" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDecimal(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetDecimal(*f)
	}
	return ftc
}

// SetDir sets the "dir" field.
func (ftc *FieldTypeCreate) SetDir(h http.Dir) *FieldTypeCreate {
	ftc.mutation.SetDir(h)
	return ftc
}

// SetNillableDir sets the "dir" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDir(h *http.Dir) *FieldTypeCreate {
	if h != nil {
		ftc.SetDir(*h)
	}
	return ftc
}

// SetNdir sets the "ndir" field.
func (ftc *FieldTypeCreate) SetNdir(h http.Dir) *FieldTypeCreate {
	ftc.mutation.SetNdir(h)
	return ftc
}

// SetNillableNdir sets the "ndir" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNdir(h *http.Dir) *FieldTypeCreate {
	if h != nil {
		ftc.SetNdir(*h)
	}
	return ftc
}

// SetStr sets the "str" field.
func (ftc *FieldTypeCreate) SetStr(ss sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetStr(ss)
	return ftc
}

// SetNullStr sets the "null_str" field.
func (ftc *FieldTypeCreate) SetNullStr(ss sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetNullStr(ss)
	return ftc
}

// SetLink sets the "link" field.
func (ftc *FieldTypeCreate) SetLink(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLink(s)
	return ftc
}

// SetLinkOther sets the "link_other" field.
func (ftc *FieldTypeCreate) SetLinkOther(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLinkOther(s)
	return ftc
}

// SetNullLink sets the "null_link" field.
func (ftc *FieldTypeCreate) SetNullLink(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetNullLink(s)
	return ftc
}

// SetActive sets the "active" field.
func (ftc *FieldTypeCreate) SetActive(s schema.Status) *FieldTypeCreate {
	ftc.mutation.SetActive(s)
	return ftc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableActive(s *schema.Status) *FieldTypeCreate {
	if s != nil {
		ftc.SetActive(*s)
	}
	return ftc
}

// SetNullActive sets the "null_active" field.
func (ftc *FieldTypeCreate) SetNullActive(s schema.Status) *FieldTypeCreate {
	ftc.mutation.SetNullActive(s)
	return ftc
}

// SetNillableNullActive sets the "null_active" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullActive(s *schema.Status) *FieldTypeCreate {
	if s != nil {
		ftc.SetNullActive(*s)
	}
	return ftc
}

// SetDeleted sets the "deleted" field.
func (ftc *FieldTypeCreate) SetDeleted(sb sql.NullBool) *FieldTypeCreate {
	ftc.mutation.SetDeleted(sb)
	return ftc
}

// SetDeletedAt sets the "deleted_at" field.
func (ftc *FieldTypeCreate) SetDeletedAt(st sql.NullTime) *FieldTypeCreate {
	ftc.mutation.SetDeletedAt(st)
	return ftc
}

// SetIP sets the "ip" field.
func (ftc *FieldTypeCreate) SetIP(n net.IP) *FieldTypeCreate {
	ftc.mutation.SetIP(n)
	return ftc
}

// SetNullInt64 sets the "null_int64" field.
func (ftc *FieldTypeCreate) SetNullInt64(si sql.NullInt64) *FieldTypeCreate {
	ftc.mutation.SetNullInt64(si)
	return ftc
}

// SetSchemaInt sets the "schema_int" field.
func (ftc *FieldTypeCreate) SetSchemaInt(s schema.Int) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt(s)
	return ftc
}

// SetNillableSchemaInt sets the "schema_int" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt(s *schema.Int) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaInt(*s)
	}
	return ftc
}

// SetSchemaInt8 sets the "schema_int8" field.
func (ftc *FieldTypeCreate) SetSchemaInt8(s schema.Int8) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt8(s)
	return ftc
}

// SetNillableSchemaInt8 sets the "schema_int8" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt8(s *schema.Int8) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaInt8(*s)
	}
	return ftc
}

// SetSchemaInt64 sets the "schema_int64" field.
func (ftc *FieldTypeCreate) SetSchemaInt64(s schema.Int64) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt64(s)
	return ftc
}

// SetNillableSchemaInt64 sets the "schema_int64" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt64(s *schema.Int64) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaInt64(*s)
	}
	return ftc
}

// SetSchemaFloat sets the "schema_float" field.
func (ftc *FieldTypeCreate) SetSchemaFloat(s schema.Float64) *FieldTypeCreate {
	ftc.mutation.SetSchemaFloat(s)
	return ftc
}

// SetNillableSchemaFloat sets the "schema_float" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaFloat(s *schema.Float64) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaFloat(*s)
	}
	return ftc
}

// SetSchemaFloat32 sets the "schema_float32" field.
func (ftc *FieldTypeCreate) SetSchemaFloat32(s schema.Float32) *FieldTypeCreate {
	ftc.mutation.SetSchemaFloat32(s)
	return ftc
}

// SetNillableSchemaFloat32 sets the "schema_float32" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaFloat32(s *schema.Float32) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaFloat32(*s)
	}
	return ftc
}

// SetNullFloat sets the "null_float" field.
func (ftc *FieldTypeCreate) SetNullFloat(sf sql.NullFloat64) *FieldTypeCreate {
	ftc.mutation.SetNullFloat(sf)
	return ftc
}

// SetRole sets the "role" field.
func (ftc *FieldTypeCreate) SetRole(r role.Role) *FieldTypeCreate {
	ftc.mutation.SetRole(r)
	return ftc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableRole(r *role.Role) *FieldTypeCreate {
	if r != nil {
		ftc.SetRole(*r)
	}
	return ftc
}

// SetMAC sets the "mac" field.
func (ftc *FieldTypeCreate) SetMAC(s schema.MAC) *FieldTypeCreate {
	ftc.mutation.SetMAC(s)
	return ftc
}

// SetUUID sets the "uuid" field.
func (ftc *FieldTypeCreate) SetUUID(u uuid.UUID) *FieldTypeCreate {
	ftc.mutation.SetUUID(u)
	return ftc
}

// Mutation returns the FieldTypeMutation object of the builder.
func (ftc *FieldTypeCreate) Mutation() *FieldTypeMutation {
	return ftc.mutation
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	var (
		err  error
		node *FieldType
	)
	ftc.defaults()
	if len(ftc.hooks) == 0 {
		if err = ftc.check(); err != nil {
			return nil, err
		}
		node, err = ftc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FieldTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftc.check(); err != nil {
				return nil, err
			}
			ftc.mutation = mutation
			node, err = ftc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftc.hooks) - 1; i >= 0; i-- {
			mut = ftc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ftc *FieldTypeCreate) defaults() {
	if _, ok := ftc.mutation.Str(); !ok {
		v := fieldtype.DefaultStr()
		ftc.mutation.SetStr(v)
	}
	if _, ok := ftc.mutation.NullStr(); !ok {
		v := fieldtype.DefaultNullStr()
		ftc.mutation.SetNullStr(v)
	}
	if _, ok := ftc.mutation.IP(); !ok {
		v := fieldtype.DefaultIP()
		ftc.mutation.SetIP(v)
	}
	if _, ok := ftc.mutation.Role(); !ok {
		v := fieldtype.DefaultRole
		ftc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FieldTypeCreate) check() error {
	if _, ok := ftc.mutation.Int(); !ok {
		return &ValidationError{Name: "int", err: errors.New("ent: missing required field \"int\"")}
	}
	if _, ok := ftc.mutation.Int8(); !ok {
		return &ValidationError{Name: "int8", err: errors.New("ent: missing required field \"int8\"")}
	}
	if _, ok := ftc.mutation.Int16(); !ok {
		return &ValidationError{Name: "int16", err: errors.New("ent: missing required field \"int16\"")}
	}
	if _, ok := ftc.mutation.Int32(); !ok {
		return &ValidationError{Name: "int32", err: errors.New("ent: missing required field \"int32\"")}
	}
	if _, ok := ftc.mutation.Int64(); !ok {
		return &ValidationError{Name: "int64", err: errors.New("ent: missing required field \"int64\"")}
	}
	if v, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		if err := fieldtype.ValidateOptionalInt32Validator(v); err != nil {
			return &ValidationError{Name: "validate_optional_int32", err: fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %w", err)}
		}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := fieldtype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if v, ok := ftc.mutation.Ndir(); ok {
		if err := fieldtype.NdirValidator(string(v)); err != nil {
			return &ValidationError{Name: "ndir", err: fmt.Errorf("ent: validator failed for field \"ndir\": %w", err)}
		}
	}
	if v, ok := ftc.mutation.Link(); ok {
		if err := fieldtype.LinkValidator(v.String()); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf("ent: validator failed for field \"link\": %w", err)}
		}
	}
	if _, ok := ftc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New("ent: missing required field \"role\"")}
	}
	if v, ok := ftc.mutation.Role(); ok {
		if err := fieldtype.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	if v, ok := ftc.mutation.MAC(); ok {
		if err := fieldtype.MACValidator(v.String()); err != nil {
			return &ValidationError{Name: "mac", err: fmt.Errorf("ent: validator failed for field \"mac\": %w", err)}
		}
	}
	return nil
}

func (ftc *FieldTypeCreate) sqlSave(ctx context.Context) (*FieldType, error) {
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ftc *FieldTypeCreate) createSpec() (*FieldType, *sqlgraph.CreateSpec) {
	var (
		_node = &FieldType{config: ftc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fieldtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fieldtype.FieldID,
			},
		}
	)
	if value, ok := ftc.mutation.Int(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldInt,
		})
		_node.Int = value
	}
	if value, ok := ftc.mutation.Int8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldInt8,
		})
		_node.Int8 = value
	}
	if value, ok := ftc.mutation.Int16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fieldtype.FieldInt16,
		})
		_node.Int16 = value
	}
	if value, ok := ftc.mutation.Int32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldInt32,
		})
		_node.Int32 = value
	}
	if value, ok := ftc.mutation.Int64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldInt64,
		})
		_node.Int64 = value
	}
	if value, ok := ftc.mutation.OptionalInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldOptionalInt,
		})
		_node.OptionalInt = value
	}
	if value, ok := ftc.mutation.OptionalInt8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldOptionalInt8,
		})
		_node.OptionalInt8 = value
	}
	if value, ok := ftc.mutation.OptionalInt16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fieldtype.FieldOptionalInt16,
		})
		_node.OptionalInt16 = value
	}
	if value, ok := ftc.mutation.OptionalInt32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldOptionalInt32,
		})
		_node.OptionalInt32 = value
	}
	if value, ok := ftc.mutation.OptionalInt64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldOptionalInt64,
		})
		_node.OptionalInt64 = value
	}
	if value, ok := ftc.mutation.NillableInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldNillableInt,
		})
		_node.NillableInt = &value
	}
	if value, ok := ftc.mutation.NillableInt8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldNillableInt8,
		})
		_node.NillableInt8 = &value
	}
	if value, ok := ftc.mutation.NillableInt16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: fieldtype.FieldNillableInt16,
		})
		_node.NillableInt16 = &value
	}
	if value, ok := ftc.mutation.NillableInt32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldNillableInt32,
		})
		_node.NillableInt32 = &value
	}
	if value, ok := ftc.mutation.NillableInt64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldNillableInt64,
		})
		_node.NillableInt64 = &value
	}
	if value, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: fieldtype.FieldValidateOptionalInt32,
		})
		_node.ValidateOptionalInt32 = value
	}
	if value, ok := ftc.mutation.OptionalUint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: fieldtype.FieldOptionalUint,
		})
		_node.OptionalUint = value
	}
	if value, ok := ftc.mutation.OptionalUint8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: fieldtype.FieldOptionalUint8,
		})
		_node.OptionalUint8 = value
	}
	if value, ok := ftc.mutation.OptionalUint16(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: fieldtype.FieldOptionalUint16,
		})
		_node.OptionalUint16 = value
	}
	if value, ok := ftc.mutation.OptionalUint32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fieldtype.FieldOptionalUint32,
		})
		_node.OptionalUint32 = value
	}
	if value, ok := ftc.mutation.OptionalUint64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: fieldtype.FieldOptionalUint64,
		})
		_node.OptionalUint64 = value
	}
	if value, ok := ftc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := ftc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: fieldtype.FieldState,
		})
		_node.State = value
	}
	if value, ok := ftc.mutation.OptionalFloat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fieldtype.FieldOptionalFloat,
		})
		_node.OptionalFloat = value
	}
	if value, ok := ftc.mutation.OptionalFloat32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: fieldtype.FieldOptionalFloat32,
		})
		_node.OptionalFloat32 = value
	}
	if value, ok := ftc.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fieldtype.FieldDatetime,
		})
		_node.Datetime = value
	}
	if value, ok := ftc.mutation.Decimal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fieldtype.FieldDecimal,
		})
		_node.Decimal = value
	}
	if value, ok := ftc.mutation.Dir(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldDir,
		})
		_node.Dir = value
	}
	if value, ok := ftc.mutation.Ndir(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldNdir,
		})
		_node.Ndir = &value
	}
	if value, ok := ftc.mutation.Str(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldStr,
		})
		_node.Str = value
	}
	if value, ok := ftc.mutation.NullStr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldNullStr,
		})
		_node.NullStr = &value
	}
	if value, ok := ftc.mutation.Link(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldLink,
		})
		_node.Link = value
	}
	if value, ok := ftc.mutation.LinkOther(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fieldtype.FieldLinkOther,
		})
		_node.LinkOther = value
	}
	if value, ok := ftc.mutation.NullLink(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldNullLink,
		})
		_node.NullLink = &value
	}
	if value, ok := ftc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fieldtype.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := ftc.mutation.NullActive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fieldtype.FieldNullActive,
		})
		_node.NullActive = &value
	}
	if value, ok := ftc.mutation.Deleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fieldtype.FieldDeleted,
		})
		_node.Deleted = value
	}
	if value, ok := ftc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fieldtype.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ftc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: fieldtype.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := ftc.mutation.NullInt64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldNullInt64,
		})
		_node.NullInt64 = value
	}
	if value, ok := ftc.mutation.SchemaInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fieldtype.FieldSchemaInt,
		})
		_node.SchemaInt = value
	}
	if value, ok := ftc.mutation.SchemaInt8(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: fieldtype.FieldSchemaInt8,
		})
		_node.SchemaInt8 = value
	}
	if value, ok := ftc.mutation.SchemaInt64(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: fieldtype.FieldSchemaInt64,
		})
		_node.SchemaInt64 = value
	}
	if value, ok := ftc.mutation.SchemaFloat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fieldtype.FieldSchemaFloat,
		})
		_node.SchemaFloat = value
	}
	if value, ok := ftc.mutation.SchemaFloat32(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: fieldtype.FieldSchemaFloat32,
		})
		_node.SchemaFloat32 = value
	}
	if value, ok := ftc.mutation.NullFloat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fieldtype.FieldNullFloat,
		})
		_node.NullFloat = value
	}
	if value, ok := ftc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: fieldtype.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := ftc.mutation.MAC(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fieldtype.FieldMAC,
		})
		_node.MAC = value
	}
	if value, ok := ftc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fieldtype.FieldUUID,
		})
		_node.UUID = value
	}
	return _node, _spec
}

// FieldTypeCreateBulk is the builder for creating many FieldType entities in bulk.
type FieldTypeCreateBulk struct {
	config
	builders []*FieldTypeCreate
}

// Save creates the FieldType entities in the database.
func (ftcb *FieldTypeCreateBulk) Save(ctx context.Context) ([]*FieldType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FieldType, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FieldTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftcb *FieldTypeCreateBulk) SaveX(ctx context.Context) []*FieldType {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
