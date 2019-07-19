// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"log"
	"net/url"
	"os"

	"fbc/ent/dialect/gremlin"
)

// endpoint for the database. In order to run the tests locally, run the following command:
//
//	 ENTV2_INTEGRATION_ENDPOINT="http://localhost:8182" go test -v
//
var endpoint *gremlin.Endpoint

func init() {
	if e, ok := os.LookupEnv("ENTV2_INTEGRATION_ENDPOINT"); ok {
		if u, err := url.Parse(e); err == nil {
			endpoint = &gremlin.Endpoint{u}
		}
	}
}

func ExampleGroup() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(gremlin.NewDriver(conn)))

	// creating vertices for the group's edges.

	// create group vertex with its edges.
	gr := client.Group.
		Create().
		SaveX(ctx)
	log.Println("group created:", gr)

	// query edges.

	// Output:
}
func ExamplePet() {
	if endpoint == nil {
		return
	}
	ctx := context.Background()
	conn, err := gremlin.NewClient(gremlin.Config{Endpoint: *endpoint})
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	client := NewClient(Driver(gremlin.NewDriver(conn)))

	// creating vertices for the pet's edges.

	// create pet vertex with its edges.
	pe := client.Pet.
		Create().
		SaveX(ctx)
	log.Println("pet created:", pe)

	// query edges.

	// Output:
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
		SetPhone("string").
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.

	// Output:
}
