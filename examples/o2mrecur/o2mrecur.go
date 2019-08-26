package main

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/examples/o2mrecur/ent"
	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:o2o2types?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	client := ent.NewClient(ent.Driver(db))
	ctx := context.Background()
	// run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	root, err := client.Node.
		Create().
		SetValue(2).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the root: %v", err)
	}

	// Add additional nodes to the tree:
	//
	//       2
	//     /   \
	//    1     4
	//        /   \
	//       3     5
	//

	// Unlike `Create`, `CreateX` panics if an error occurs.
	n1 := client.Node.
		Create().
		SetValue(1).
		SetParent(root).
		SaveX(ctx)
	n4 := client.Node.
		Create().
		SetValue(4).
		SetParent(root).
		SaveX(ctx)
	n3 := client.Node.
		Create().
		SetValue(3).
		SetParent(n4).
		SaveX(ctx)
	n5 := client.Node.
		Create().
		SetValue(5).
		SetParent(n4).
		SaveX(ctx)

	fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})
	// Output: Tree leafs [1 3 5]

	// Get all leafs (nodes without children).
	// Unlike `Int`, `IntX` panics if an error occurs.
	ints := client.Node.
		Query().                             // All nodes.
		Where(node.Not(node.HasChildren())). // Only leafs.
		Order(ent.Asc(node.FieldValue)).     // Order by their `value` field.
		GroupBy(node.FieldValue).            // Extract only the `value` field.
		IntsX(ctx)
	fmt.Println(ints)
	// Output: [1 3 5]

	// Get orphan nodes (nodes without parent).
	// Unlike `Only`, `OnlyX` panics if an error occurs.
	orphan := client.Node.
		Query().
		Where(node.Not(node.HasParent())).
		OnlyX(ctx)
	fmt.Println(orphan)
	// Output: Node(id=1, value=2)

	return nil
}
