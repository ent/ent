// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/version/ent"
	_ "entgo.io/ent/examples/version/ent/runtime"
	"entgo.io/ent/examples/version/ent/user"

	_ "github.com/mattn/go-sqlite3"
)

func Example_OptimisticLock() {
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

	usr := client.User.Create().SetStatus(user.StatusOnline).SaveX(ctx)
	fmt.Println(usr.ID, usr.Status)

	usrCopy := client.User.Query().OnlyX(ctx)
	usrCopy = usrCopy.Update().SetStatus(user.StatusOffline).SaveX(ctx)
	fmt.Println(usrCopy.ID, usrCopy.Status)

	// The operation fails because the user was updated by another process (usrCopy).
	_, err = usr.Update().SetStatus(user.StatusOffline).Save(ctx)
	fmt.Println(err)

	// Output:
	// 1 online
	// 1 offline
	// user 1 was changed by another process
}
