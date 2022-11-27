// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package migratedata holds the functions for generating data migration files.
// It exists here for documentation and reference purpose only and has no runtime
// effect on the actual migration files.
package migratedata

import (
	"context"
	"fmt"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/examples/migration/ent"
	"entgo.io/ent/examples/migration/ent/user"
)

// BackfillUserTags was used to generate the migration file '20221126185750_backfill_user_tags.sql'.
// It exists here for documentation purpose only, and can be used as a reference for future data migrations.
func BackfillUserTags(dir *migrate.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.MySQL, w)))

	// Add defaults "foo" and "bar" tags for users without any.
	err := client.User.
		Update().
		Where(func(s *sql.Selector) {
			s.Where(
				sql.Or(
					sql.IsNull(user.FieldTags),
					sqljson.ValueIsNull(user.FieldTags),
				),
			)
		}).
		SetTags([]string{"foo", "bar"}).
		Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating backfill statement: %w", err)
	}
	// Document all changes until now with a custom comment.
	w.Change("Backfill NULL or null tags with a default value.")

	// Append the "org" special tag for users with a specific prefix or suffix.
	err = client.User.
		Update().
		Where(
			user.Or(
				user.NameHasPrefix("org-"),
				user.NameHasSuffix("-org"),
			),
			// Append to only those without this tag.
			func(s *sql.Selector) {
				s.Where(
					sql.Not(sqljson.ValueContains(user.FieldTags, "org")),
				)
			},
		).
		AppendTags([]string{"org"}).
		Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating backfill statement: %w", err)
	}
	// Document all changes until now with a custom comment.
	w.Change("Append the 'org' tag for organization accounts in case they don't have it.")

	// Write the content to the migration directory.
	return w.Flush("backfill_user_tags")
}

// BackfillUnknown back-fills all empty users' names with the default value 'Unknown'.
func BackfillUnknown(dir *migrate.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.MySQL, w)))

	// Change all empty names to 'unknown'.
	err := client.User.
		Update().
		Where(
			user.NameEQ(""),
		).
		SetName("Unknown").
		Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"unknown_names",
		"Backfill all empty user names with default value 'unknown'.",
	)
}

// SeedUsers add the initial users to the database.
func SeedUsers(dir *migrate.LocalDir) error {
	w := &schema.DirWriter{Dir: dir}
	client := ent.NewClient(ent.Driver(schema.NewWriteDriver(dialect.MySQL, w)))

	// The statement that generates the INSERT statement.
	err := client.User.CreateBulk(
		client.User.Create().SetName("a8m").SetAge(1).SetTags([]string{"foo"}),
		client.User.Create().SetName("nati").SetAge(1).SetTags([]string{"bar"}),
	).Exec(context.Background())
	if err != nil {
		return fmt.Errorf("failed generating statement: %w", err)
	}

	// Write the content to the migration directory.
	return w.FlushChange(
		"seed_users",
		"Add the initial users to the database.",
	)
}
