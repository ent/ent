// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"net/url"
	"os"

	"fbc/ent/dialect/gremlin"
)

// endpoint for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="http://localhost:8182" go test -v
//
var endpoint *gremlin.Endpoint

func init() {
	if e, ok := os.LookupEnv("ENT_INTEGRATION_ENDPOINT"); ok {
		if u, err := url.Parse(e); err == nil {
			endpoint = &gremlin.Endpoint{u}
		}
	}
}

func ExampleBoring() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(gremlin.NewDriver(conn)))

	// creating vertices for the boring's edges.

	// create boring vertex with its edges.
	b := client.Boring.
		Create().
		SaveX(ctx)
	log.Println("boring created:", b)

	// query edges.

	// Output:
}
