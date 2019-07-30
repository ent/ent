// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"fbc/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleBoring() {
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
	// creating vertices for the boring's edges.

	// create boring vertex with its edges.
	b := client.Boring.
		Create().
		SaveX(ctx)
	log.Println("boring created:", b)

	// query edges.

	// Output:
}
