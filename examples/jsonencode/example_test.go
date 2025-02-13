// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"entgo.io/ent/examples/jsonencode/ent"
	_ "github.com/mattn/go-sqlite3"
)

func Example_jsonEncode() {
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

	a8m := client.User.Create().SetName("a8m").SetAge(10).SaveX(ctx)
	buf, err := json.Marshal(a8m)
	if err != nil {
		log.Fatalf("failed marshaling user: %v", err)
	}
	fmt.Println(string(buf))

	xabi := client.Pet.Create().SetName("xabi").SetAge(1).SetOwner(a8m).SaveX(ctx)
	buf, err = json.Marshal(xabi)
	if err != nil {
		log.Fatalf("failed marshaling pet: %v", err)
	}
	fmt.Println(string(buf))

	users := client.User.Query().WithPets().AllX(ctx)
	buf, err = json.Marshal(users)
	if err != nil {
		log.Fatalf("failed marshaling users: %v", err)
	}
	fmt.Println(string(buf))

	pets := client.Pet.Query().WithOwner().AllX(ctx)
	buf, err = json.Marshal(pets)
	if err != nil {
		log.Fatalf("failed marshaling pets: %v", err)
	}
	fmt.Println(string(buf))

	// Output:
	// {"id":1,"age":10,"name":"a8m"}
	// {"id":1,"age":1,"name":"xabi","owner_id":1}
	// [{"id":1,"age":10,"name":"a8m","pets":[{"id":1,"age":1,"name":"xabi","owner_id":1}]}]
	// [{"id":1,"age":1,"name":"xabi","owner_id":1,"owner":{"id":1,"age":10,"name":"a8m"}}]
}
