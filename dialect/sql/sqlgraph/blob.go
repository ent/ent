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

// maxBlobKeysPerQuery bounds the keys per reference-count query, keeping the
// statement clear of driver placeholder limits on large deletes.
const maxBlobKeysPerQuery = 500

// CountBlobKeyRefs implements [ent.BlobRefCounter]. For each key it counts the
// rows holding it in its field's key column, across the whole table — [BlobSpec.Predicate]
// is deliberately not applied, since the question is who else refers to the key.
// Keys absent from the result are held by no row.
func (s *BlobSpec) CountBlobKeyRefs(ctx context.Context, keys []ent.BlobKey) (map[ent.BlobKey]int, error) {
	if len(keys) == 0 || len(s.Columns) == 0 {
		return nil, nil
	}
	// Group the distinct keys by the column they would be stored in.
	lookup := make(map[string][]any)
	seen := make(map[ent.BlobKey]bool, len(keys))
	for _, k := range keys {
		if k.Key == "" || seen[k] {
			continue
		}
		seen[k] = true
		if col, ok := s.Columns[k.Field]; ok {
			lookup[col] = append(lookup[col], k.Key)
		}
	}
	counts := make(map[ent.BlobKey]int)
	for field, col := range s.Columns {
		vals := lookup[col]
		for len(vals) > 0 {
			n := min(len(vals), maxBlobKeysPerQuery)
			if err := s.countRefs(ctx, field, col, vals[:n], counts); err != nil {
				return nil, err
			}
			vals = vals[n:]
		}
	}
	return counts, nil
}

// countRefs accumulates the number of rows holding each of vals in col.
func (s *BlobSpec) countRefs(ctx context.Context, field, col string, vals []any, counts map[ent.BlobKey]int) error {
	query, args := sql.Dialect(s.Driver.Dialect()).
		Select(col, sql.Count("*")).
		From(sql.Table(s.Table)).
		Where(sql.In(col, vals...)).
		GroupBy(col).
		Query()
	rows := &sql.Rows{}
	if err := s.Driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			key string
			n   int
		)
		if err := rows.Scan(&key, &n); err != nil {
			return err
		}
		counts[ent.BlobKey{Field: field, Key: key}] += n
	}
	return rows.Err()
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
