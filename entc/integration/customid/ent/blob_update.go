// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/predicate"
	"github.com/google/uuid"
)

// BlobUpdate is the builder for updating Blob entities.
type BlobUpdate struct {
	config
	uuid       *uuid.UUID
	predicates []predicate.Blob
}

// Where adds a new predicate for the builder.
func (bu *BlobUpdate) Where(ps ...predicate.Blob) *BlobUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetUUID sets the uuid field.
func (bu *BlobUpdate) SetUUID(u uuid.UUID) *BlobUpdate {
	bu.uuid = &u
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BlobUpdate) Save(ctx context.Context) (int, error) {
	return bu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlobUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlobUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlobUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BlobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(bu.driver.Dialect())
		selector = builder.Select(blob.FieldID).From(builder.Table(blob.Table))
	)
	for _, p := range bu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = bu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := bu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(blob.Table)
	)
	idface := make([]interface{}, len(ids))
	for i := range ids {
		idface[i] = ids[i]
	}
	updater = updater.Where(sql.In(blob.FieldID, idface...))
	if value := bu.uuid; value != nil {
		updater.Set(blob.FieldUUID, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// BlobUpdateOne is the builder for updating a single Blob entity.
type BlobUpdateOne struct {
	config
	id   uuid.UUID
	uuid *uuid.UUID
}

// SetUUID sets the uuid field.
func (buo *BlobUpdateOne) SetUUID(u uuid.UUID) *BlobUpdateOne {
	buo.uuid = &u
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BlobUpdateOne) Save(ctx context.Context) (*Blob, error) {
	return buo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlobUpdateOne) SaveX(ctx context.Context) *Blob {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BlobUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlobUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BlobUpdateOne) sqlSave(ctx context.Context) (b *Blob, err error) {
	var (
		builder  = sql.Dialect(buo.driver.Dialect())
		selector = builder.Select(blob.Columns...).From(builder.Table(blob.Table))
	)
	blob.ID(buo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = buo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		b = &Blob{config: buo.config}
		if err := b.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Blob: %v", err)
		}
		id = b.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Blob with id: %v", buo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Blob with the same id: %v", buo.id)
	}

	tx, err := buo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(blob.Table)
	)
	idface := make([]interface{}, len(ids))
	for i := range ids {
		idface[i] = ids[i]
	}
	updater = updater.Where(sql.In(blob.FieldID, idface...))
	if value := buo.uuid; value != nil {
		updater.Set(blob.FieldUUID, *value)
		b.UUID = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return b, nil
}
