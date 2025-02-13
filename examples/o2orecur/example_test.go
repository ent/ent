// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/o2orecur/ent"
	"entgo.io/ent/examples/o2orecur/ent/node"

	_ "github.com/mattn/go-sqlite3"
)

func Example_o2oRecur() {
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
	// 1 2 3 4 5
	// true
}

func Do(ctx context.Context, client *ent.Client) error {
	head, err := client.Node.
		Create().
		SetValue(1).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the head: %w", err)
	}
	curr := head
	// Generate the following linked-list: 1<->2<->3<->4<->5.
	for i := 0; i < 4; i++ {
		curr, err = client.Node.
			Create().
			SetValue(curr.Value + 1).
			SetPrev(curr).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	// Loop over the list and print it. `FirstX` panics if an error occur.
	for curr = head; curr != nil; curr = curr.QueryNext().FirstX(ctx) {
		fmt.Printf(" %d", curr.Value)
	}
	// Output: 1 2 3 4 5

	// Make the linked-list circular:
	// The tail of the list, has no "next".
	tail, err := client.Node.
		Query().
		Where(node.Not(node.HasNext())).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("getting the tail of the list: %v", tail)
	}
	tail, err = tail.Update().SetNext(head).Save(ctx)
	if err != nil {
		return err
	}
	// Check that the change actually applied:
	prev, err := head.QueryPrev().Only(ctx)
	if err != nil {
		return fmt.Errorf("getting head's prev: %w", err)
	}
	fmt.Printf("\n%v", prev.Value == tail.Value)
	// Output: true

	return nil
}
