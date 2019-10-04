// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleGroup() {
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
	// creating vertices for the group's edges.

	// create group vertex with its edges.
	gr := client.Group.
		Create().
		SetMaxUsers(1).
		SaveX(ctx)
	log.Println("group created:", gr)

	// query edges.

	// Output:
}
func ExamplePet() {
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
	// creating vertices for the pet's edges.

	// create pet vertex with its edges.
	pe := client.Pet.
		Create().
		SetAge(1).
		SetLicensedAt(time.Now()).
		SaveX(ctx)
	log.Println("pet created:", pe)

	// query edges.

	// Output:
}
func ExampleUser() {
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
	// creating vertices for the user's edges.
	pe0 := client.Pet.
		Create().
		SetAge(1).
		SetLicensedAt(time.Now()).
		SaveX(ctx)
	log.Println("pet created:", pe0)
	u1 := client.User.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u1)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetName("string").
		AddPets(pe0).
		AddFriends(u1).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	pe0, err = u.QueryPets().First(ctx)
	if err != nil {
		log.Fatalf("failed querying pets: %v", err)
	}
	log.Println("pets found:", pe0)

	u1, err = u.QueryFriends().First(ctx)
	if err != nil {
		log.Fatalf("failed querying friends: %v", err)
	}
	log.Println("friends found:", u1)

	// Output:
}
