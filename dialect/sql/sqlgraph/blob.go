// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// BlobSpec configures SQL-level blob key queries and implements [ent.BlobQuerier].
type BlobSpec struct {
	Driver    dialect.Driver
	Table     string
	Columns   map[string]string // field name -> key column name
	Predicate func(*sql.Selector)
}

// QueryBlobKeys implements [ent.BlobQuerier].
// If fields is nil, all columns are queried (for deletes);
// otherwise only the named fields are queried.
func (s *BlobSpec) QueryBlobKeys(ctx context.Context, fields []string) ([]ent.BlobKey, error) {
	cols := s.Columns
	if len(fields) > 0 {
		cols = make(map[string]string, len(fields))
		for _, f := range fields {
			if c, ok := s.Columns[f]; ok {
				cols[f] = c
			}
		}
	}
	if len(cols) == 0 {
		return nil, nil
	}
	names := make([]string, 0, len(cols))
	colNames := make([]string, 0, len(cols))
	for field, col := range cols {
		names = append(names, field)
		colNames = append(colNames, col)
	}
	selector := sql.Dialect(s.Driver.Dialect()).
		Select(colNames...).
		From(sql.Table(s.Table))
	if s.Predicate != nil {
		s.Predicate(selector)
	}
	query, args := selector.Query()
	rows := &sql.Rows{}
	if err := s.Driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var keys []ent.BlobKey
	for rows.Next() {
		vals := make([]*string, len(colNames))
		ptrs := make([]any, len(colNames))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		for i, v := range vals {
			if v != nil && *v != "" {
				keys = append(keys, ent.BlobKey{Field: names[i], Key: *v})
			}
		}
	}
	return keys, rows.Err()
}
