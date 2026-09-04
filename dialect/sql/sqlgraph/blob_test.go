// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestBlobSpecReferencedBlobKeys(t *testing.T) {
	tests := []struct {
		name    string
		dialect string
		keys    []ent.BlobKey
		expect  func(sqlmock.Sqlmock)
		want    []ent.BlobKey
	}{
		{
			name:    "keys held by a row",
			dialect: dialect.MySQL,
			keys: []ent.BlobKey{
				{Field: "content", Key: "k1"},
				{Field: "content", Key: "k2"},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape("SELECT DISTINCT `content_key` FROM `documents` WHERE `content_key` IN (?, ?)")).
					WithArgs("k1", "k2").
					WillReturnRows(sqlmock.NewRows([]string{"content_key"}).AddRow("k1"))
			},
			want: []ent.BlobKey{{Field: "content", Key: "k1"}},
		},
		{
			name:    "keys held by no row",
			dialect: dialect.SQLite,
			keys:    []ent.BlobKey{{Field: "content", Key: "gone"}},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape("SELECT DISTINCT `content_key` FROM `documents` WHERE `content_key` IN (?)")).
					WithArgs("gone").
					WillReturnRows(sqlmock.NewRows([]string{"content_key"}))
			},
			want: nil,
		},
		{
			// The mutation's own rows are not filtered out here -- callers time the
			// lookup so those rows already read the way they will settle.
			name:    "the mutation's predicate is not applied",
			dialect: dialect.Postgres,
			keys:    []ent.BlobKey{{Field: "content", Key: "k1"}},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape(`SELECT DISTINCT "content_key" FROM "documents" WHERE "content_key" IN ($1)`)).
					WithArgs("k1").
					WillReturnRows(sqlmock.NewRows([]string{"content_key"}).AddRow("k1"))
			},
			want: []ent.BlobKey{{Field: "content", Key: "k1"}},
		},
		{
			name:    "duplicate keys are looked up once",
			dialect: dialect.MySQL,
			keys: []ent.BlobKey{
				{Field: "content", Key: "k1"},
				{Field: "content", Key: "k1"},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape("SELECT DISTINCT `content_key` FROM `documents` WHERE `content_key` IN (?)")).
					WithArgs("k1").
					WillReturnRows(sqlmock.NewRows([]string{"content_key"}).AddRow("k1"))
			},
			want: []ent.BlobKey{{Field: "content", Key: "k1"}},
		},
		{
			name:    "unknown fields and empty keys are skipped",
			dialect: dialect.MySQL,
			keys: []ent.BlobKey{
				{Field: "nosuchfield", Key: "k1"},
				{Field: "content", Key: ""},
			},
			expect: func(sqlmock.Sqlmock) {},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.expect(mock)
			spec := &BlobSpec{
				Driver: sql.OpenDB(tt.dialect, db),
				Table:  "documents",
				// Set to prove ReferencedBlobKeys ignores it.
				Predicate: func(s *sql.Selector) { s.Where(sql.EQ("id", 7)) },
				Columns:   map[string]string{"content": "content_key", "meta": "meta_key"},
				CheckRefs: map[string]bool{"content": true},
			}
			referenced, err := spec.ReferencedBlobKeys(context.Background(), tt.keys)
			require.NoError(t, err)
			require.Equal(t, tt.want, referenced)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// A key set larger than maxBlobKeysPerQuery is split across statements so the
// query stays clear of driver placeholder limits.
func TestBlobSpecReferencedBlobKeysChunks(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	var (
		keys  = make([]ent.BlobKey, 0, maxBlobKeysPerQuery+1)
		first = make([]driver.Value, 0, maxBlobKeysPerQuery)
	)
	for i := range maxBlobKeysPerQuery + 1 {
		key := fmt.Sprintf("k%d", i)
		keys = append(keys, ent.BlobKey{Field: "content", Key: key})
		if i < maxBlobKeysPerQuery {
			first = append(first, key)
		}
	}
	// The first statement carries a full chunk, the leftover key its own. Matched
	// loosely — escape anchors the pattern, and the argument list asserts the size.
	mock.ExpectQuery("SELECT DISTINCT").
		WithArgs(first...).
		WillReturnRows(sqlmock.NewRows([]string{"content_key"}).AddRow(keys[0].Key))
	mock.ExpectQuery(escape("SELECT DISTINCT `content_key` FROM `documents` WHERE `content_key` IN (?)")).
		WithArgs(keys[maxBlobKeysPerQuery].Key).
		WillReturnRows(sqlmock.NewRows([]string{"content_key"}).AddRow(keys[maxBlobKeysPerQuery].Key))
	spec := &BlobSpec{
		Driver:    sql.OpenDB(dialect.MySQL, db),
		Table:     "documents",
		Columns:   map[string]string{"content": "content_key"},
		CheckRefs: map[string]bool{"content": true},
	}
	referenced, err := spec.ReferencedBlobKeys(context.Background(), keys)
	require.NoError(t, err)
	require.Equal(t, []ent.BlobKey{keys[0], keys[maxBlobKeysPerQuery]}, referenced)
	require.NoError(t, mock.ExpectationsWereMet())
}

// Empty input, and fields that did not opt in, must not reach the database.
func TestBlobSpecReferencedBlobKeysNoop(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	spec := &BlobSpec{
		Driver:    sql.OpenDB(dialect.MySQL, db),
		Table:     "documents",
		Columns:   map[string]string{"content": "content_key"},
		CheckRefs: map[string]bool{"content": true},
	}
	keys := []ent.BlobKey{{Field: "content", Key: "k"}}
	referenced, err := spec.ReferencedBlobKeys(context.Background(), nil)
	require.NoError(t, err)
	require.Empty(t, referenced)

	// No field opted in: the whole lookup is skipped.
	spec.CheckRefs = nil
	referenced, err = spec.ReferencedBlobKeys(context.Background(), keys)
	require.NoError(t, err)
	require.Empty(t, referenced)

	// Some other field opted in, but not this one.
	spec.CheckRefs = map[string]bool{"thumbnail": true}
	referenced, err = spec.ReferencedBlobKeys(context.Background(), keys)
	require.NoError(t, err)
	require.Empty(t, referenced)

	spec.CheckRefs = map[string]bool{"content": true}
	spec.Columns = nil
	referenced, err = spec.ReferencedBlobKeys(context.Background(), keys)
	require.NoError(t, err)
	require.Empty(t, referenced)
	require.NoError(t, mock.ExpectationsWereMet())
}
