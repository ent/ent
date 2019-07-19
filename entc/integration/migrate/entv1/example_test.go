// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"log"
	"net/url"
	"os"

	"fbc/ent/dialect/gremlin"
)

// endpoint for the database. In order to run the tests locally, run the following command:
//
//	 ENTV1_INTEGRATION_ENDPOINT="http://localhost:8182" go test -v
//
var endpoint *gremlin.Endpoint

func init() {
	if e, ok := os.LookupEnv("ENTV1_INTEGRATION_ENDPOINT"); ok {
		if u, err := url.Parse(e); err == nil {
			endpoint = &gremlin.Endpoint{u}
		}
	}
}

func ExampleUser() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(gremlin.NewDriver(conn)))

	// creating vertices for the user's edges.

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetAge(1).
		SetName("string").
		SetAddress("string").
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.

	// Output:
}
