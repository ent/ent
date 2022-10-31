// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"log"

	"entgo.io/ent/examples/dynamodb/m2m2types/ent"
)

func Example_M2M2Types() {
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
	// [Group(id=1, name=GitHub) Group(id=2, name=GitLab)]
	// [Group(id=1, name=GitHub)]
	// [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
}

func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	hub := client.Group.
		Create().
		SetID(1).
		SetName("GitHub").
		SaveX(ctx)
	lab := client.Group.
		Create().
		SetID(2).
		SetName("GitLab").
		SaveX(ctx)
	_ = client.User.
		Create().
		SetID(1).
		SetAge(30).
		SetName("a8m").
		AddGroups(hub, lab).
		SaveX(ctx)
	_ = client.User.
		Create().
		SetID(2).
		SetAge(28).
		SetName("nati").
		AddGroups(hub).
		SaveX(ctx)

	//// Query the edges.
	//groups, err := a8m.
	//	QueryGroups().
	//	All(ctx)
	//if err != nil {
	//	return fmt.Errorf("querying a8m groups: %w", err)
	//}
	//fmt.Println(groups)
	//// Output: [Group(id=1, name=GitHub) Group(id=2, name=GitLab)]
	//
	//groups, err = nati.
	//	QueryGroups().
	//	All(ctx)
	//if err != nil {
	//	return fmt.Errorf("querying nati groups: %w", err)
	//}
	//fmt.Println(groups)
	//// Output: [Group(id=1, name=GitHub)]
	//
	//// Traverse the graph.
	//users, err := a8m.
	//	QueryGroups().                                           // [hub, lab]
	//	Where(group.Not(group.HasUsersWith(user.Name("nati")))). // [lab]
	//	QueryUsers().                                            // [a8m]
	//	QueryGroups().                                           // [hub, lab]
	//	QueryUsers().                                            // [a8m, nati]
	//	All(ctx)
	//if err != nil {
	//	return fmt.Errorf("traversing the graph: %w", err)
	//}
	//fmt.Println(users)
	//// Output: [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
	return nil
}
