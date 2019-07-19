package gremlin

import (
	"context"
	"flag"
	"log"
	"os"
	"time"
)

func ExampleClient_Query() {
	addr := flag.String("gremlin-server", os.Getenv("GREMLIN_SERVER"), "gremlin server address")
	flag.Parse()

	if *addr == "" {
		log.Fatal("missing gremlin server address")
	}

	client, err := NewHTTPClient(*addr, nil)
	if err != nil {
		log.Fatalf("creating client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rsp, err := client.Query(ctx, "g.E()")
	if err != nil {
		log.Fatalf("executing query: %v", err)
	}

	edges, err := rsp.ReadEdges()
	if err != nil {
		log.Fatalf("unmashal edges")
	}

	for _, e := range edges {
		log.Println(e.String())
	}
	// - Output:
}
