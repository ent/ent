// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/file"

	_ "github.com/mattn/go-sqlite3"
)

func Example_recursiveTraversal() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
	names := client.File.Query().
		Where(func(s *sql.Selector) {
			t1, t2 := sql.Table(file.Table), sql.Table(file.Table)
			with := sql.WithRecursive("undeleted", file.FieldID, file.FieldParentID)
			with.As(
				// The initial `SELECT` statement executed once at the start,
				// and produces the initial row or rows for the recursion.
				sql.Select(t1.Columns(file.FieldID, file.FieldParentID)...).
					From(t1).
					Where(
						sql.And(
							sql.IsNull(t1.C(file.FieldParentID)),
							sql.EQ(t1.C(file.FieldDeleted), false),
						),
					).
					// Merge the `SELECT` statement above with the following:
					UnionAll(
						// A `SELECT` statement that produces additional rows and recurses by referring
						// to the CTE name (e.g. "undeleted"), and ends when there are no more new rows.
						sql.Select(t2.Columns(file.FieldID, file.FieldParentID)...).
							From(t2).
							Join(with).
							On(t2.C(file.FieldParentID), with.C(file.FieldID)).
							Where(
								sql.EQ(t1.C(file.FieldDeleted), false),
							),
					),
			)
			// Join the root `SELECT` query with the CTE result (`WITH` clause).
			s.Prefix(with).Join(with).On(s.C(file.FieldID), with.C(file.FieldID))
		}).
		Select(file.FieldName).
		StringsX(ctx)
	fmt.Printf("%q\n", names)

	// Output:
	// ["a" "b" "ba" "bb"]
}
