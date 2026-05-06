// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"context"
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// BlobDeleteSpec holds the configuration for deleting blobs
// associated with entities matched by a predicate.
type BlobDeleteSpec struct {
	// Opener opens a blob bucket for the given field.
	Opener func(context.Context, string) (ent.Blob, error)
	// Predicate filters which rows' blob keys to collect.
	Predicate func(*sql.Selector)
	// Table is the SQL table name.
	Table string
	// Fields are the blob field names (e.g. "content", "thumbnail").
	// The corresponding key columns are derived as "<field>_key".
	Fields []string
}

// BlobDeletes queries blob key columns for rows matching the spec predicate
// and returns a cleanup function that deletes the blobs from storage.
// The cleanup function should be called only after the SQL rows have been
// successfully deleted to ensure keys are collected while rows still exist.
func BlobDeletes(ctx context.Context, drv dialect.Driver, spec *BlobDeleteSpec) (func() error, error) {
	if len(spec.Fields) == 0 {
		return func() error { return nil }, nil
	}
	columns := make([]string, len(spec.Fields))
	for i, f := range spec.Fields {
		columns[i] = f + "_key"
	}
	selector := sql.Dialect(drv.Dialect()).
		Select(columns...).
		From(sql.Table(spec.Table))
	if spec.Predicate != nil {
		spec.Predicate(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	type blobKey struct{ f, k string }
	var keys []blobKey
	for rows.Next() {
		vals := make([]*string, len(columns))
		ptrs := make([]any, len(columns))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		for i, v := range vals {
			if v != nil && *v != "" {
				keys = append(keys, blobKey{f: spec.Fields[i], k: *v})
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return func() error {
		if len(keys) == 0 {
			return nil
		}
		blobs := ent.NewBlobBulkWriter(spec.Opener)
		var errs []error
		for _, bk := range keys {
			if err := blobs.Delete(ctx, bk.f, bk.k); err != nil {
				errs = append(errs, err)
			}
		}
		if err := blobs.Close(); err != nil {
			errs = append(errs, err)
		}
		return errors.Join(errs...)
	}, nil
}
