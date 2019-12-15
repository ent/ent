// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENTV1_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

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

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetNickname("string").
		SetAddress("string").
		SetRenamed("string").
		SetBlob(nil).
		SetState(user.StateLoggedIn).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.

	// Output:
}
