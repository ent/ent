// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

//go:build ignore

package main

import (
	"context"
	"flag"
	"log"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/ent/entc/integration/migrate/entv1"
)

func main() {
	var (
		drv string
		dsn string
		dir string
	)
	flag.StringVar(&drv, "driver", "", "driver to use")
	flag.StringVar(&dsn, "dsn", "", "dsn of development DB to connect to")
	flag.StringVar(&dir, "dir", "migrations", "path/to/migration/files")
	flag.Parse()
	// Create the migration directory.
	md, err := migrate.NewLocalDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	// Create the client.
	client, err := entv1.Open(drv, dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// Run the differ.
	if err := client.Schema.Diff(context.Background(), schema.WithDir(md)); err != nil {
		log.Fatalln(err)
	}
}
