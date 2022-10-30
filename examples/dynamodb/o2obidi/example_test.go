// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/dynamodb/o2obidi/ent"
)

func Example_O2OBidi() {
	client, err := ent.Open("dynamodb", "")
	if err != nil {
		log.Fatalf("failed opening connection to dynamodb: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Output:
	// a8m
	// nati
	// 2
	// nati
}

func Do(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetID(1).
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	_, err = client.User.
		Create().
		SetID(2).
		SetAge(28).
		SetName("nati").
		SetSpouse(a8m).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}

	//// Query the spouse edge.
	//// Unlike `Only`, `OnlyX` panics if an error occurs.
	//spouse := nati.QuerySpouse().OnlyX(ctx)
	//fmt.Println(spouse.Name)
	//// Output: a8m
	//
	//spouse = a8m.QuerySpouse().OnlyX(ctx)
	//fmt.Println(spouse.Name)
	//// Output: nati
	//
	//// Query how many users have a spouse.
	//// Unlike `Count`, `CountX` panics if an error occurs.
	//count := client.User.
	//	Query().
	//	Where(user.HasSpouse()).
	//	CountX(ctx)
	//fmt.Println(count)
	//// Output: 2
	//
	//// Get the user, that has a spouse with name="a8m".
	//spouse = client.User.
	//	Query().
	//	Where(user.HasSpouseWith(user.Name("a8m"))).
	//	OnlyX(ctx)
	//fmt.Println(spouse.Name)
	//// Output: nati
	return nil
}
