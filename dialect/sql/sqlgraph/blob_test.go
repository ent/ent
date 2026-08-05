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

func TestBlobSpecCountBlobKeyRefs(t *testing.T) {
	tests := []struct {
		name    string
		dialect string
		keys    []ent.BlobKey
		expect  func(sqlmock.Sqlmock)
		want    map[ent.BlobKey]int
	}{
		{
			name:    "counts rows per key",
			dialect: dialect.MySQL,
			keys: []ent.BlobKey{
				{Field: "content", Key: "k1"},
				{Field: "content", Key: "k2"},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape("SELECT `content_key`, COUNT(*) FROM `documents` WHERE `content_key` IN (?, ?) GROUP BY `content_key`")).
					WithArgs("k1", "k2").
					WillReturnRows(sqlmock.NewRows([]string{"content_key", "count"}).
						AddRow("k1", 2).
						AddRow("k2", 1))
			},
			want: map[ent.BlobKey]int{
				{Field: "content", Key: "k1"}: 2,
				{Field: "content", Key: "k2"}: 1,
			},
		},
		{
			name:    "keys held by no row are absent",
			dialect: dialect.SQLite,
			keys:    []ent.BlobKey{{Field: "content", Key: "gone"}},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape("SELECT `content_key`, COUNT(*) FROM `documents` WHERE `content_key` IN (?) GROUP BY `content_key`")).
					WithArgs("gone").
					WillReturnRows(sqlmock.NewRows([]string{"content_key", "count"}))
			},
			want: map[ent.BlobKey]int{},
		},
		{
			name:    "duplicate keys are looked up once",
			dialect: dialect.Postgres,
			keys: []ent.BlobKey{
				{Field: "content", Key: "k1"},
				{Field: "content", Key: "k1"},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(escape(`SELECT "content_key", COUNT(*) FROM "documents" WHERE "content_key" IN ($1) GROUP BY "content_key"`)).
					WithArgs("k1").
					WillReturnRows(sqlmock.NewRows([]string{"content_key", "count"}).AddRow("k1", 2))
			},
			want: map[ent.BlobKey]int{{Field: "content", Key: "k1"}: 2},
		},
		{
			name:    "unknown fields and empty keys are skipped",
			dialect: dialect.MySQL,
			keys: []ent.BlobKey{
				{Field: "nosuchfield", Key: "k1"},
				{Field: "content", Key: ""},
			},
			expect: func(sqlmock.Sqlmock) {},
			want:   map[ent.BlobKey]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.expect(mock)
			spec := &BlobSpec{
				Driver:  sql.OpenDB(tt.dialect, db),
				Table:   "documents",
				Columns: map[string]string{"content": "content_key"},
			}
			counts, err := spec.CountBlobKeyRefs(context.Background(), tt.keys)
			require.NoError(t, err)
			require.Equal(t, tt.want, counts)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// A key set larger than maxBlobKeysPerQuery is split across statements so the
// query stays clear of driver placeholder limits.
func TestBlobSpecCountBlobKeyRefsChunks(t *testing.T) {
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
	mock.ExpectQuery("COUNT").
		WithArgs(first...).
		WillReturnRows(sqlmock.NewRows([]string{"content_key", "count"}).AddRow(keys[0].Key, 3))
	mock.ExpectQuery(escape("SELECT `content_key`, COUNT(*) FROM `documents` WHERE `content_key` IN (?) GROUP BY `content_key`")).
		WithArgs(keys[maxBlobKeysPerQuery].Key).
		WillReturnRows(sqlmock.NewRows([]string{"content_key", "count"}).AddRow(keys[maxBlobKeysPerQuery].Key, 1))
	spec := &BlobSpec{
		Driver:  sql.OpenDB(dialect.MySQL, db),
		Table:   "documents",
		Columns: map[string]string{"content": "content_key"},
	}
	counts, err := spec.CountBlobKeyRefs(context.Background(), keys)
	require.NoError(t, err)
	require.Equal(t, 3, counts[keys[0]])
	require.Equal(t, 1, counts[keys[maxBlobKeysPerQuery]])
	require.NoError(t, mock.ExpectationsWereMet())
}

// Empty input must not reach the database.
func TestBlobSpecCountBlobKeyRefsNoop(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	spec := &BlobSpec{
		Driver:  sql.OpenDB(dialect.MySQL, db),
		Table:   "documents",
		Columns: map[string]string{"content": "content_key"},
	}
	counts, err := spec.CountBlobKeyRefs(context.Background(), nil)
	require.NoError(t, err)
	require.Empty(t, counts)

	spec.Columns = nil
	counts, err = spec.CountBlobKeyRefs(context.Background(), []ent.BlobKey{{Field: "content", Key: "k"}})
	require.NoError(t, err)
	require.Empty(t, counts)
	require.NoError(t, mock.ExpectationsWereMet())
}
