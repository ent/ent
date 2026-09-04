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
	Driver  dialect.Driver
	Table   string
	Columns map[string]string // field name -> key column name
	// CheckRefs holds the fields whose keys more than one row can hold, the ones
	// declaring CheckRefs on the schema. Only these are looked up before a delete;
	// a key generated fresh per write has no other holder to find. An empty
	// CheckRefs makes [BlobSpec.ReferencedBlobKeys] a no-op.
	CheckRefs map[string]bool
	Predicate func(*sql.Selector)
}

// maxBlobKeysPerQuery bounds the keys per reference lookup, keeping the
// statement clear of driver placeholder limits on large deletes.
const maxBlobKeysPerQuery = 500

// ReferencedBlobKeys implements [ent.BlobRefChecker]. It returns the subset of keys
// some row holds in its field's key column. [BlobSpec.Predicate] is deliberately not
// applied: the question is which keys are in use anywhere, and callers time the lookup
// so the mutation's own rows already read the way they will settle.
//
// Only fields listed in [BlobSpec.CheckRefs] are looked up. Keys of any other field are
// reported as held by no row, which is what an unshared key strategy guarantees anyway.
//
// The lookup only asks whether a key is used, never how often, so an index on each
// blob key column lets it stop at the first match.
func (s *BlobSpec) ReferencedBlobKeys(ctx context.Context, keys []ent.BlobKey) ([]ent.BlobKey, error) {
	if len(keys) == 0 || len(s.CheckRefs) == 0 {
		return nil, nil
	}
	// Group the distinct keys by the column they would be stored in.
	lookup := make(map[string][]any)
	seen := make(map[ent.BlobKey]bool, len(keys))
	for _, k := range keys {
		if k.Key == "" || seen[k] || !s.CheckRefs[k.Field] {
			continue
		}
		seen[k] = true
		if col, ok := s.Columns[k.Field]; ok {
			lookup[col] = append(lookup[col], k.Key)
		}
	}
	var referenced []ent.BlobKey
	for field := range s.CheckRefs {
		col, ok := s.Columns[field]
		if !ok {
			continue
		}
		vals := lookup[col]
		for len(vals) > 0 {
			n := min(len(vals), maxBlobKeysPerQuery)
			found, err := s.selectReferenced(ctx, col, vals[:n])
			if err != nil {
				return nil, err
			}
			for _, key := range found {
				referenced = append(referenced, ent.BlobKey{Field: field, Key: key})
			}
			vals = vals[n:]
		}
	}
	return referenced, nil
}

// selectReferenced returns which of vals appear in col.
func (s *BlobSpec) selectReferenced(ctx context.Context, col string, vals []any) ([]string, error) {
	query, args := sql.Dialect(s.Driver.Dialect()).
		Select(col).
		Distinct().
		From(sql.Table(s.Table)).
		Where(sql.In(col, vals...)).
		Query()
	rows := &sql.Rows{}
	if err := s.Driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var found []string
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, err
		}
		found = append(found, key)
	}
	return found, rows.Err()
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
