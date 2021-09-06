// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"testing"

	"entgo.io/ent/dialect"
)

func BenchmarkInsertBuilder_Default(b *testing.B) {
	for _, d := range []string{dialect.SQLite, dialect.MySQL, dialect.Postgres} {
		b.Run(d, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				Dialect(d).Insert("users").Default().Returning("id").Query()
			}
		})
	}
}

func BenchmarkInsertBuilder_Small(b *testing.B) {
	for _, d := range []string{dialect.SQLite, dialect.MySQL, dialect.Postgres} {
		b.Run(d, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				Dialect(d).Insert("users").
					Columns("id", "age", "first_name", "last_name", "nickname", "spouse_id", "created_at", "updated_at").
					Values(1, 30, "Ariel", "Mashraki", "a8m", 2, "2009-11-10 23:00:00", "2009-11-10 23:00:00").
					Returning("id").
					Query()
			}
		})
	}
}
