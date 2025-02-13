// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"entgo.io/ent/examples/o2o2types/ent"

	_ "github.com/mattn/go-sqlite3"
)

func Example_o2o2Types() {
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
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// user: User(id=1, age=30, name=Mashraki)
	// card: Card(id=1, expired=Sun Dec  8 15:04:05 2019, number=1020)
	// card: Card(id=1, expired=Sun Dec  8 15:04:05 2019, number=1020)
	// owner: User(id=1, age=30, name=Mashraki)
}

func Do(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Mashraki").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	fmt.Println("user:", a8m)
	expired, err := time.Parse(time.RFC3339, "2019-12-08T15:04:05Z")
	if err != nil {
		return err
	}
	card1, err := client.Card.
		Create().
		SetOwner(a8m).
		SetNumber("1020").
		SetExpired(expired).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	fmt.Println("card:", card1)
	// Only returns the card of the user,
	// and expects that there's only one.
	card2, err := a8m.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %w", err)
	}
	fmt.Println("card:", card2)
	// The Card entity is able to query its owner using
	// its back-reference.
	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %w", err)
	}
	fmt.Println("owner:", owner)
	return nil
}
