// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/traversal/ent"
	"entgo.io/ent/examples/traversal/ent/group"
	"entgo.io/ent/examples/traversal/ent/pet"
	"entgo.io/ent/examples/traversal/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func Example_Traversal() {
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
	if err := Gen(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err := Traverse(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err := Traverse2(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Generate a group of entities in a transaction.
	if err := GenTx(ctx, client); err != nil {
		log.Fatal(err)
	}
	// Wrap an existing function ("Gen") in a transaction.
	if err := WrapGen(ctx, client); err != nil {
		log.Fatal(err)
	}
	// WithTx helper.
	if err := WithTx(ctx, client, func(tx *ent.Tx) error {
		return Gen(ctx, tx.Client())
	}); err != nil {
		log.Fatal(err)
	}
	// Output:
	// Pets created: Pet(id=1, name=Pedro) Pet(id=2, name=Xabi) Pet(id=3, name=Coco)
	// User(id=3, age=37, name=Alex)
	// [Pet(id=1, name=Pedro) Pet(id=2, name=Xabi)]
	// User(id=5, age=30, name=Ariel)
	// Pets created: Pet(id=4, name=Pedro) Pet(id=5, name=Xabi) Pet(id=6, name=Coco)
	// Pets created: Pet(id=7, name=Pedro) Pet(id=8, name=Xabi) Pet(id=9, name=Coco)
}

func Gen(ctx context.Context, client *ent.Client) error {
	hub, err := client.Group.
		Create().
		SetName("Github").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating the group: %w", err)
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
					Where(group.Name("Github")). // Filter only GitHub group (only 1).
					QueryAdmin().                // Getting Dan.
					QueryFriends().              // Getting Dan's friends: [Ariel].
					QueryPets().                 // Their pets: [Pedro, Xabi].
					QueryFriends().              // Pedro's friends: [Coco], Xabi's friends: [].
					QueryOwner().                // Coco's owner: Alex.
					Only(ctx)                    // Expect only one entity to return in the query.
	if err != nil {
		return fmt.Errorf("failed querying the owner: %w", err)
	}
	fmt.Println(owner)
	// Output:
	// User(id=3, age=37, name=Alex)
	return nil
}

// Traverse2 example from the doc.
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
		return fmt.Errorf("failed querying the pets: %w", err)
	}
	fmt.Println(pets)
	// Output:
	// [Pet(id=1, name=Pedro) Pet(id=2, name=Xabi)]
	return nil
}

// GenTx example from the doc.
func GenTx(ctx context.Context, client *ent.Client) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	hub, err := tx.Group.
		Create().
		SetName("Github").
		Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("failed creating the group: %w", err))
	}
	// Create the admin of the group.
	dan, err := tx.User.
		Create().
		SetAge(29).
		SetName("Dan").
		AddManage(hub).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	// Create user "Ariel".
	a8m, err := tx.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		AddGroups(hub).
		AddFriends(dan).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	fmt.Println(a8m)
	// Output:
	// User(id=1, age=30, name=Ariel)
	return tx.Commit()
}

// WrapGen wraps the existing "Gen" function in a transaction.
func WrapGen(ctx context.Context, client *ent.Client) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	txClient := tx.Client()
	// Use the same "Gen" as above, but give it the transactional client; no code changes to "Gen".
	if err := Gen(ctx, txClient); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// WithTx example from the doc.
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
