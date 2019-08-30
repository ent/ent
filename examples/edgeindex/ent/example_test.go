// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleCity() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the city's edges.
	s0 := client.Street.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("street created:", s0)

	// create city vertex with its edges.
	c := client.City.
		Create().
		SetName("string").
		AddStreets(s0).
		SaveX(ctx)
	log.Println("city created:", c)

	// query edges.
	s0, err = c.QueryStreets().First(ctx)
	if err != nil {
		log.Fatalf("failed querying streets: %v", err)
	}
	log.Println("streets found:", s0)

	// Output:
}
func ExampleStreet() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the street's edges.

	// create street vertex with its edges.
	s := client.Street.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("street created:", s)

	// query edges.

	// Output:
}
