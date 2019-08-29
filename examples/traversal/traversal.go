// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/examples/traversal/ent/user"

	"github.com/facebookincubator/ent/examples/traversal/ent/pet"

	"github.com/facebookincubator/ent/examples/traversal/ent"
	"github.com/facebookincubator/ent/examples/traversal/ent/group"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:o2o2types?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	client := ent.NewClient(ent.Driver(db))
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Gen(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err := Traverse(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err := Traverse2(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Gen(ctx context.Context, client *ent.Client) error {
	hub, err := client.Group.
		Create().
		SetName("Github").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating the group: %v", err)
	}
	// Create the admin of the group.
	// Unlike `Save`, `SaveX` panics if an error occurs.
	dan := client.User.
		Create().
		SetAge(29).
		SetName("Dan").
		AddManage(hub).
		SaveX(ctx)

	// Create "Ariel" and its pets.
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		AddGroups(hub).
		AddFriends(dan).
		SaveX(ctx)
	pedro := client.Pet.
		Create().
		SetName("Pedro").
		SetOwner(a8m).
		SaveX(ctx)
	xabi := client.Pet.
		Create().
		SetName("Xabi").
		SetOwner(a8m).
		SaveX(ctx)

	// Create "Alex" and its pets.
	alex := client.User.
		Create().
		SetAge(37).
		SetName("Alex").
		SaveX(ctx)
	coco := client.Pet.
		Create().
		SetName("Coco").
		SetOwner(alex).
		AddFriends(pedro).
		SaveX(ctx)

	fmt.Println("Pets created:", pedro, xabi, coco)
	// Output:
	// Pets created: Pet(id=1, name=Pedro) Pet(id=2, name=Xabi) Pet(id=3, name=Coco)
	return nil
}

func Traverse(ctx context.Context, client *ent.Client) error {
	owner, err := client.Group. // GroupClient.
					Query().                     // Query builder.
					Where(group.Name("Github")). // Filter only Github group (only 1).
					QueryAdmin().                // Getting Dan.
					QueryFriends().              // Getting Dan's friends: [Ariel].
					QueryPets().                 // Their pets: [Pedro, Xabi].
					QueryFriends().              // Pedro's friends: [Coco], Xabi's friends: [].
					QueryOwner().                // Coco's owner: Alex.
					Only(ctx)                    // Expect only one entity to return in the query.
	if err != nil {
		return fmt.Errorf("failed querying the owner: %v", err)
	}
	fmt.Println(owner)
	// Output:
	// User(id=3, age=37, name=Alex)
	return nil
}

func Traverse2(ctx context.Context, client *ent.Client) error {
	pets, err := client.Pet.
		Query().
		Where(
			pet.HasOwnerWith(
				user.HasFriendsWith(
					user.HasManage(),
				),
			),
		).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying the pets: %v", err)
	}
	fmt.Println(pets)
	// Output:
	// [Pet(id=1, name=Pedro) Pet(id=2, name=Xabi)]
	return nil
}
