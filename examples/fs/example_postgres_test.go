// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/file"

	entSQL "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func Example_PostgreSQL_RecursiveTraversal() {

	for _, port := range []int{5434, 5433, 5432, 5431, 5430} {

		db, err := sql.Open("pgx", fmt.Sprintf("postgresql://postgres:pass@127.0.0.1:%d/test", port))
		if err != nil {
			log.Fatalf("failed opening connection to PostgreSQL(port %d): %v", port, err)
		}
		driver := entSQL.OpenDB(dialect.Postgres, db)
		client := ent.NewClient(ent.Driver(driver))

		defer client.Close()
		ctx := context.Background()
		// Run the auto migration tool.
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources PostgreSQL(port %d): %v", port, err)
		}

		// Clean old tests
		client.Debug().File.Delete().Exec(ctx)

		// Add multiple files in the following tree structure:
		//
		//	a/
		//	├─ b/
		//	│  ├─ ba
		//	│  ├─ bb
		//	│  └─ bc (deleted)
		//	├─ c/ (deleted)
		//	│  ├─ ca
		//	│  └─ cb
		//	└─ d (deleted)
		//
		a := client.File.Create().SetName("a").SaveX(ctx)
		b := client.File.Create().SetName("b").SetParent(a).SaveX(ctx)
		client.File.Create().SetName("ba").SetParent(b).SaveX(ctx)
		client.File.Create().SetName("bb").SetParent(b).SaveX(ctx)
		client.File.Create().SetName("bc").SetParent(b).SetDeleted(true).SaveX(ctx)
		c := client.File.Create().SetName("c").SetParent(a).SetDeleted(true).SaveX(ctx)
		client.File.Create().SetName("ca").SetParent(c).SaveX(ctx)
		client.File.Create().SetName("cb").SetParent(c).SaveX(ctx)
		client.File.Create().SetName("d").SetParent(a).SetDeleted(true).SaveX(ctx)

		// Query undeleted files:
		//
		//	a/
		//	└─ b/
		//	   ├─ ba
		//	   └─ bb
		//
		// First query to check SQL is valid for database enginge
		client.File.Query().
			Where(func(s *entSQL.Selector) {
				t1, t2 := entSQL.Table(file.Table), entSQL.Table(file.Table)
				with := entSQL.WithRecursive("undeleted", file.FieldID, file.FieldParentID)
				with.As(
					// The initial `SELECT` statement executed once at the start,
					// and produces the initial row or rows for the recursion.
					entSQL.Select(t1.Columns(file.FieldID, file.FieldParentID)...).
						From(t1).
						Where(
							entSQL.And(
								entSQL.IsNull(t1.C(file.FieldParentID)),
								entSQL.EQ(t1.C(file.FieldDeleted), false),
							),
						).
						OrderBy(entSQL.Asc(file.FieldID)).Limit(1).
						// Merge the `SELECT` statement above with the following:
						UnionAll(
							// A `SELECT` statement that produces additional rows and recurses by referring
							// to the CTE name (e.g. "undeleted"), and ends when there are no more new rows.
							entSQL.Select(t2.Columns(file.FieldID, file.FieldParentID)...).
								From(t2).
								Join(with).
								On(t2.C(file.FieldParentID), with.C(file.FieldID)).
								Where(
									entSQL.EQ(t1.C(file.FieldDeleted), false),
								).
								OrderBy(entSQL.Asc(file.FieldDeleted)).Limit(2),
						),
				)
				// Join the root `SELECT` query with the CTE result (`WITH` clause).
				s.Prefix(with).Join(with).On(s.C(file.FieldID), with.C(file.FieldID))
			}).
			Limit(3).
			Order(func(s *entSQL.Selector) {
				s.OrderBy(entSQL.Asc(file.FieldName))
			}).
			Select(file.FieldName).
			StringsX(ctx)

		// Second query to check to check output

		names := client.File.Query().
			Where(func(s *entSQL.Selector) {
				t1, t2 := entSQL.Table(file.Table), entSQL.Table(file.Table)
				with := entSQL.WithRecursive("undeleted", file.FieldID, file.FieldParentID)
				with.As(
					// The initial `SELECT` statement executed once at the start,
					// and produces the initial row or rows for the recursion.
					entSQL.Select(t1.Columns(file.FieldID, file.FieldParentID)...).
						From(t1).
						Where(
							entSQL.And(
								entSQL.IsNull(t1.C(file.FieldParentID)),
								entSQL.EQ(t1.C(file.FieldDeleted), false),
							),
						).
						// Merge the `SELECT` statement above with the following:
						Union(
							// A `SELECT` statement that produces additional rows and recurses by referring
							// to the CTE name (e.g. "undeleted"), and ends when there are no more new rows.
							entSQL.Select(t2.Columns(file.FieldID, file.FieldParentID)...).
								From(t2).
								Join(with).
								On(t2.C(file.FieldParentID), with.C(file.FieldID)).
								Where(
									entSQL.EQ(t1.C(file.FieldDeleted), false),
								),
						),
				)
				// Join the root `SELECT` query with the CTE result (`WITH` clause).
				s.Prefix(with).Join(with).On(s.C(file.FieldID), with.C(file.FieldID))
			}).
			Order(func(s *entSQL.Selector) {
				s.OrderBy(entSQL.Asc(file.FieldName))
			}).
			Select(file.FieldName).
			StringsX(ctx)

		fmt.Printf("PostgreSQL(port %d) %q\n", port, names)
	}

	// Output:
	// PostgreSQL(port 5434) ["a" "b" "ba" "bb"]
	// PostgreSQL(port 5433) ["a" "b" "ba" "bb"]
	// PostgreSQL(port 5432) ["a" "b" "ba" "bb"]
	// PostgreSQL(port 5431) ["a" "b" "ba" "bb"]
	// PostgreSQL(port 5430) ["a" "b" "ba" "bb"]
}
