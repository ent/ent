// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

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
	u0 := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u0)

	// create group vertex with its edges.
	gr := client.Group.
		Create().
		SetName("string").
		AddUsers(u0).
		SaveX(ctx)
	log.Println("group created:", gr)

	// query edges.
	u0, err = gr.QueryUsers().First(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	log.Println("users found:", u0)

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

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.

	// Output:
}
