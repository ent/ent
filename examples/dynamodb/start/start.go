package main

import (
	"context"
	"log"

	"entgo.io/ent/examples/dynamodb/start/ent"
)

func main() {
	client, err := ent.Open("dynamodb", "http://localhost:8000")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
