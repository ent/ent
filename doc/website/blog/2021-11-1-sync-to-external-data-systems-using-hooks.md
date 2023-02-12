---
title: Sync Changes to External Data Systems using Ent Hooks
author: Ariel Mashraki
authorURL: https://github.com/a8m
authorImageURL: "https://avatars0.githubusercontent.com/u/7413593"
authorTwitter: arielmashraki
image: https://entgo.io/images/assets/sync-hook/share.png
---

One of the common questions we get from the Ent community is how to synchronize objects or references between the
database backing an Ent application (e.g. MySQL or PostgreSQL) with external services. For example, users would like
to create or delete a record from within their CRM when a user is created or deleted in Ent, publish a message to a
[Pub/Sub system](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern) when an entity is updated, or verify
references to blobs in object storage such as AWS S3 or Google Cloud Storage.

Ensuring consistency between two separate data systems is not a simple task.  When we want to propagate, for example,
the deletion of a record in one system to another, there is no obvious way to guarantee that the two systems will end in
a synchronized state, since one of them may fail, and the network link between them may be slow or down. Having said
that, and especially with the prominence of microservices-architectures, these problems have become more common, and
distributed systems researchers have come up with patterns to solve them, such as the
[Saga Pattern](https://microservices.io/patterns/data/saga.html).

The application of these patterns is usually complex and difficult, and so in many cases architects do not go after a
"perfect" design, and instead go after simpler solutions that involve either the acceptance of some inconsistency
between the systems or background reconciliation procedures.

In this post, we will not discuss how to solve distributed transactions or implement the Saga pattern with Ent.
Instead, we will limit our scope to study how to hook into Ent mutations before and after they occur, and run our
custom logic there.

### Propagating Mutations to External Systems

In our example, we are going to create a simple `User` schema with 2 immutable string fields, `"name"` and
`"avatar_url"`. Let's run the `ent init` command for creating a skeleton schema for our `User`:

```shell
go run entgo.io/ent/cmd/ent new User
```

Then, add the `name` and the `avatar_url` fields and run `go generate` to generate the assets.

```go title="ent/schema/user.go"
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Immutable(),
		field.String("avatar_url").
			Immutable(),
	}
}
```

```shell
go generate ./ent
```

### The Problem

The `avatar_url` field defines a URL to an image in a bucket on our object storage (e.g. AWS S3). For the purpose of
this discussion  we want to make sure that:

- When a user is created, an image with the URL stored in `"avatar_url"` exists in our bucket.
- Orphan images are deleted from the bucket. This means that when a user is deleted from our system, its avatar image
  is deleted as well.
  
For interacting with blobs, we will use the [`gocloud.dev/blob`](https://gocloud.dev/howto/blob) package. This package
provides abstraction for reading, writing, deleting and listing blobs in a bucket. Similar to the `database/sql`
package, it allows interacting with variety of object storages with the same API by configuring its driver URL. 
For example:

```go
// Open an in-memory bucket. 
if bucket, err := blob.OpenBucket(ctx, "mem://photos/"); err != nil {
	log.Fatal("failed opening in-memory bucket:", err)
}

// Open an S3 bucket named photos.
if bucket, err := blob.OpenBucket(ctx, "s3://photos"); err != nil {
	log.Fatal("failed opening s3 bucket:", err)
}

// Open a bucket named photos in Google Cloud Storage.
if bucket, err := blob.OpenBucket(ctx, "gs://my-bucket"); err != nil {
	log.Fatal("failed opening gs bucket:", err)
}
```

### Schema Hooks

[Hooks](https://entgo.io/docs/hooks) are a powerful feature of Ent that allows adding custom logic before and after
operations that mutate the graph.

Hooks can be either defined dynamically using `client.Use` (called "Runtime Hooks"), or explicitly on the schema
(called "Schema Hooks") as follows:

```go
// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		EnsureImageExists(),
		DeleteOrphans(),
	}
}
```

As you can imagine, the `EnsureImageExists` hook will be responsible for ensuring that when a user is created, their
avatar URL exists in the bucket, and the `DeleteOrphans` will ensure that orphan images are deleted. Let's start
writing them.

```go title="ent/schema/hooks.go"
func EnsureImageExists() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			avatarURL, exists := m.AvatarURL()
			if !exists {
				return nil, errors.New("avatar field is missing")
			}
			// TODO:
			// 1. Verify that "avatarURL" points to a real object in the bucket.
			// 2. Otherwise, fail.
			return next.Mutate(ctx, m)
		})
	}
	// Limit the hook only to "Create" operations.
	return hook.On(hk, ent.OpCreate)
}

func DeleteOrphans() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			id, exists := m.ID()
			if !exists {
				return nil, errors.New("id field is missing")
			}
			// TODO:
			// 1. Get the AvatarURL field of the deleted user.
			// 2. Cascade the deletion to object storage.
			return next.Mutate(ctx, m)
		})
	}
	// Limit the hook only to "DeleteOne" operations.
	return hook.On(hk, ent.OpDeleteOne)
}
```

Now, you may ask yourself, _how do we access the blob client from the mutations hooks?_ You are going to find out in
the next section.

### Injecting Dependencies

The [entc.Dependency](https://entgo.io/docs/code-gen/#external-dependencies) option allows extending the generated
builders with external dependencies as struct fields, and provides options for injecting them on client initialization.

To inject a `blob.Bucket` to be available inside our hooks, we can follow the tutorial about external dependencies in
[the website](https://entgo.io/docs/code-gen/#external-dependencies), and define the
[`gocloud.dev/blob.Bucket`](https://pkg.go.dev/gocloud.dev/blob#Bucket) as a dependency.

```go title="ent/entc.go" {3-6}
func main() {
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyName("Bucket"),
			entc.DependencyType(&blob.Bucket{}),
		),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

Next, re-run code generation:

```shell
go generate ./ent
```

We can now access the Bucket API from all generated builders. Let's finish the implementations of the above hooks.

```go title="ent/schema/hooks.go"
// EnsureImageExists ensures the avatar_url points
// to a real object in the bucket.
func EnsureImageExists() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			avatarURL, exists := m.AvatarURL()
			if !exists {
				return nil, errors.New("avatar field is missing")
			}
			switch exists, err := m.Bucket.Exists(ctx, avatarURL); {
			case err != nil:
				return nil, fmt.Errorf("check key existence: %w", err)
			case !exists:
				return nil, fmt.Errorf("key %q does not exist in the bucket", avatarURL)
			default:
				return next.Mutate(ctx, m)
			}
		})
	}
	return hook.On(hk, ent.OpCreate)
}

// DeleteOrphans cascades the user deletion to the bucket.
// Hence, when a user is deleted, its avatar image is deleted
// as well.
func DeleteOrphans() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			id, exists := m.ID()
			if !exists {
				return nil, errors.New("id field is missing")
			}
			u, err := m.Client().User.Get(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("getting deleted user: %w", err)
			}
			if err := m.Bucket.Delete(ctx, u.AvatarURL); err != nil {
				return nil, fmt.Errorf("deleting user avatar from bucket: %w", err)
			}
			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpDeleteOne)
}
```

Now, it's time to test our hooks! Let's write a testable example that verifies that our 2 hooks work as expected.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/a8m/ent-sync-example/ent"
	_ "github.com/a8m/ent-sync-example/ent/runtime"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/memblob"
)

func Example_SyncCreate() {
	ctx := context.Background()
	// Open an in-memory bucket.
	bucket, err := blob.OpenBucket(ctx, "mem://photos/")
	if err != nil {
		log.Fatal("failed opening bucket:", err)
	}
	client, err := ent.Open(
		dialect.SQLite,
		"file:ent?mode=memory&cache=shared&_fk=1",
		// Inject the blob.Bucket on client initialization.
		ent.Bucket(bucket),
	)
	if err != nil {
		log.Fatal("failed opening connection to sqlite:", err)
	}
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources:", err)
	}
	if err := client.User.Create().SetName("a8m").SetAvatarURL("a8m.png").Exec(ctx); err == nil {
		log.Fatal("expect user creation to fail because the image does not exist in the bucket")
	}
	if err := bucket.WriteAll(ctx, "a8m.png", []byte{255, 255, 255}, nil); err != nil {
		log.Fatalf("failed uploading image to the bucket: %v", err)
	}
	fmt.Printf("%q\n", keys(ctx, bucket))

	// User creation should pass as image was uploaded to the bucket.
	u := client.User.Create().SetName("a8m").SetAvatarURL("a8m.png").SaveX(ctx)

	// Deleting a user, should delete also its image from the bucket.
	client.User.DeleteOne(u).ExecX(ctx)
	fmt.Printf("%q\n", keys(ctx, bucket))

	// Output:
	// ["a8m.png"]
	// []
}
```

### Wrapping Up

Great! We have configured Ent to extend our generated code and inject the `blob.Bucket` as an 
[External Dependency](https://entgo.io/docs/code-gen#external-dependencies). Next, we defined two mutation hooks and
used the `blob.Bucket` API to ensure our product constraints are satisfied.

The code for this example is available at [github.com/a8m/ent-sync-example](https://github.com/a8m/ent-sync-example).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
