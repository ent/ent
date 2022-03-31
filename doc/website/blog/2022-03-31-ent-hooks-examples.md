---
title: How I use Ent Schema Hooks
author: Yoni Davidson
authorURL: "https://github.com/yonidavidson"
authorImageURL: "https://avatars0.githubusercontent.com/u/5472778"
authorTwitter: yonidavidson
---
Despite being one of the most powerful features of Ent, Schema [hooks](https://entgo.io/docs/hooks) 
are often overlooked by many users. We've covered hooks in previous blog posts (such as 
[Building Observable Ent Applications with Prometheus](/blog/2021/08/12/building-observable-ent-application-with-prometheus)
and [Sync Changes to External Data Systems using Ent Hooks](/blog/2021/11/1/sync-to-external-data-systems-using-hooks)),
but learning from my personal experience building real applications with Ent, I thought it would be
beneficial to put them in the spotlight once more.

### What are hooks?

[Hooks](https://entgo.io/docs/hooks) are an Ent feature that allow adding custom logic before and after operations that change the data entities.

Hooks are functions that get and return an [ent.Mutator](https://pkg.go.dev/entgo.io/ent#Mutator).
They function similar to the popular HTTP middleware pattern.

```go
package example

import (
	"context"

	"entgo.io/ent"
)

func exampleHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Do something before mutation.
			v, err := next.Mutate(ctx, m)
			if err != nil {
				// Do something if error after mutation.
			}
			// Do something after mutation.
			return v, err
		})
	}
}
```

[Schema hooks](https://entgo.io/docs/hooks#schema-hooks) are mainly used for defining custom mutation logic on a specific entity type, for example,
syncing entity creation to another system.

### Schema as code

In Ent, we try and push the entity's logic to the schema level. This way, the logic can be near the schema and not located
in different places in a system. Entity properties could be:
1. Field type (Text, Int, Boolean)
2. Field valid values ("Not empty","positive")
3. Immutability status for each field
4. Privacy
5. Required Side effects (heavy compute for example)
---
Oftentimes, it’s more comfortable to define these properties in the RPC level: a controller provides a context and all the mutations
and validation are done before changing the value.
Yet, as the system grows and more developers are working on it simultaneously, making sure that these validations and calculations are done
whenever the entity is changing becomes much more challenging. This requires more pre-known knowledge about each entity and
can cause unintentional bugs or inconsistencies in the data that are only discovered in runtime.
Inspired by the experience of bigger companies like Facebook, Ent decided to handle this by keeping this information
as close as possible to the entity definition. Therefore hooks are just that - a piece of code that runs before and after an entity 
is changing.

## Let's do some examples

For our examples we use a simple schema of Users and Dogs, each user can have multiple pets and each pet has an owner 
as can be seen here:
```go title="ent/schema/user.go"
// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("phone_number").
			NotEmpty(),
		field.String("last_digits").
			Sensitive().
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Dog.Type),
		edge.To("cache", Cache.Type).Unique(),
	}
}
```
```go title="ent/schema/dog.go"
// Dog holds the schema definition for the Dog entity.
type Dog struct {
	ent.Schema
}

// Fields of the Dog.
func (Dog) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Int("owner_id").
			Optional(),
	}
}

// Edges of the Dog.
func (Dog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Field("owner_id").
			Unique(),
	}
}
```

For each example I'll provide a test and a recipe, I highly suggest starting with a test to see that the hook does what you need it to do.
Hooks should do one thing only and therfore be simple to test.

### Keeping our secrets safe:

Many times, we need to store strings that include secrets, a good example can be a phone number.
In our case the privacy concern is rather simple. We require to hide the last 4 digits of each phone number.
The expected behavior is best described by this test:
```go title="ent/schema/user_test.go"
func TestUserPhoneNumberHook(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite, "file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1")
	u := c.User.Create().SetName("Yoni").SetPhoneNumber("315-194-6020").SaveX(ctx)
	require.Equal(t, "315-194-****", u.PhoneNumber)
	require.Equal(t, "6020", u.LastDigits)
	require.Equal(t, "315-194-6020", u.FullPhoneNumber())
}
```

As you can see, inserting a phone number updates the string to have `****` instead of the last digits and stores them 
in the `last_digits` field that is defined as [sensitive](https://entgo.io/docs/schema-fields#sensitive-fields).
To have the full phone number we created a utility function `FullPhoneNumber()`.

#### Recipe:

First thing, add the hook to your schema:
```go title="ent/schema/user.go"
// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(maskPhoneNumber,
			hook.And(hook.Or(hook.HasOp(ent.OpUpdateOne), hook.HasOp(ent.OpCreate)), hook.HasFields(user.FieldPhoneNumber)),
		),
	}
}
```
* Run `go generate ./ent` or the hook will not be registered.

A few [hook helpers](https://entgo.io/docs/hooks/#hook-helpers) are used here, `maskPhoneNumber` is the name of the hook function, and it will run if a user 
is updated or created and the mutation has the field `phone_number` (if we change only the name of the user there is no need to do anything right?).

The hook:
```go title="ent/schema/user.go"
func maskPhoneNumber(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		cs, ok := m.PhoneNumber()
		if !ok {
			return next.Mutate(ctx, m)
		}
		sp := strings.Split(cs, "-")
		if len(sp) != 3 {
			return next.Mutate(ctx, m)
		}
		m.SetLastDigits(sp[2])
		sp[2] = "****"
		m.SetPhoneNumber(strings.Join(sp[0:3], "-"))
		return next.Mutate(ctx, m)
	})
}
```
We validate that the mutation has a `phone_number`(if not just let the mutation continue).
A bit of string acrobatics, and we have the `last_digits` extracted from the string. We store it in another field - the `last_digits` and
mask the `phone_number`. 

To get the full phone number we add a utility function 
```go title="ent/user_phone_number.go"
package ent

import "strings"

// FullPhoneNumber returns the phone number unmasked.
func (u *User) FullPhoneNumber() string {
	return strings.ReplaceAll(u.PhoneNumber, "****", u.LastDigits)
}
```

It's valid to add files to the Ent directory (not only in schema folder) since the Ent generate command only appends/changes Ent
generated files.


### Special Validators:

Ent provides [validators](https://entgo.io/docs/schema-fields#validators) for fields, yet sometimes the logic can contain
dependencies that are more applicative and require broader context.
For example, in our use case, we would like to make sure that dog's names don’t start with the first 2 letters of the owner's name 
(best practice when you don't want your dog to run in circles every time someone calls your name).

We start with a simple test:
```go title="ent/schema/dog_test.go"
func TestDogNameValidationHook(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite,
		"file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1",
	)
	u := c.User.Create().SetName("Yoni").
		SetConnectionString("mysql://root:pass@localhost:3306)").
		SaveX(ctx)
	_, err := c.Dog.Create().SetName("Yolo").SetOwner(u).Save(ctx)
	require.Error(t, err)
}
```

In this case, you can see that since "Yoni" and "Yolo" share the same 2 letters "Yo" the action fails.

#### Recipe:

First, register the hook:
```go title="ent/schema/dog.go"
// Hooks of the Dog.
func (Dog) Hooks() []ent.Hook {
	return []ent.Hook{		
		hook.If(validateName,
			hook.HasFields(dog.FieldOwnerID),
		),
	}
}
```

Since we want our hook to call every time we change the owner (to make sure the name is matched against), I am using an
[edge field](https://entgo.io/docs/schema-edges#edge-field). This provides the hooks a trigger when the edge's ID changes.

The hook:
```go title="ent/schema/dog.go"
func validateName(next ent.Mutator) ent.Mutator {
	return hook.DogFunc(func(ctx context.Context, m *gen.DogMutation) (ent.Value, error) {
		owID, ok := m.OwnerID()
		if !ok {
			return next.Mutate(ctx, m)
		}
		owner, err := m.Client().User.Query().Where(user.ID(owID)).Only(ctx)
		if err != nil {
			return next.Mutate(ctx, m)
		}
		dn, ok := m.Name()
		if !ok {
			return next.Mutate(ctx, m)
		}
		if owner.Name[0:2] == dn[0:2] {
			return nil, errors.New("invalid dog name")
		}
		return next.Mutate(ctx, m)
	})
}
```
A simple name validator and if we see a problem we cancel the mutation by returning an error before `next.Mutate(ctx,m)`.

### Offloading Long Computations:

In some cases, we have fields in our database that require long computation, we would like to make sure that this data is always updated once
a change in the data requires it.

A test for it will look like this:
```go title="ent/schema/dog_test.go"
func TestCacheHook(t *testing.T) {
	ctx := context.Background()
	cs := cache.NewSyncer()
	c := enttest.Open(t, dialect.SQLite,
		"file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1",
		enttest.WithOptions(ent.CacheSyncer(cs)),
	)
	cs.Start(ctx, c)
	cl := c.Cache.Create().SetWalks(-1).SaveX(ctx)
	d := c.Dog.Create().SetName("Karashindo").SaveX(ctx)
	u := c.User.Create().SetName("Yoni").
		SetCache(cl).
		AddPets(d).
		SetConnectionString("mysql://root:pass@localhost:3306)").
		SaveX(ctx)
	c.Dog.UpdateOne(d).SetName("Fortuna").ExecX(ctx)
	cs.Close()
	cl = u.QueryCache().OnlyX(ctx)
	require.True(t, cl.Walks > 0)
}
```
This test is a bit more complicated, so let's review it line by line:
First, I create a new `cache syncer` (remember we have a cache entity in our schema).

This `Cache Syncer` can do heavy computations while running in a different context.
We create an Ent client and inject our `Cache Syncer`, it is disabled by default.
After that, we start it and provide it with the Ent client, the reason for that is to allow it to change our data (modify our cache entity).
We create a few entities and update the dog's name, that will trigger a cache sync.
Finally, we close the `Cache Syncer`, this will block until all pending hooks are completed.
The assertion validates a value of "walks" larger than 0 (not -1).

#### Recipe:

Let's add a cache layer for our calculations:
```go title="ent/schema/cache.go"
// Cache holds the schema definition for the Cache entity.
type Cache struct {
	ent.Schema
}

// Fields of the Cache.
func (Cache) Fields() []ent.Field {
	return []ent.Field{
		field.Int("walks"),
	}
}
```

Now the hook:
```go title="ent/schema/dog.go"
// Hooks of the Dog.
func (Dog) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(syncCache,
			hook.HasOp(ent.OpUpdateOne),
		),		
	}
}
```

The Sync cache will trigger once a dog is updated.

For injecting the dependency in ent we follow [injecting external dependencies](https://entgo.io/docs/code-gen#external-dependencies)
section and update our  `entc.go`:

```go title="ent/entc.go"
func main() {
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyName("CacheSyncer"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "hook.Syncer",
				PkgPath: "github.com/yonidavidson/ent-hooks-examples/hook",
			}),
		),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

The interface exposes the sync command:
```go title="hook/syncer.go"
package hook

import "context"

type Syncer interface {
	Sync(ctx context.Context, cacheID int)
}
```


The hooks themselves will query for the Owner's cache ID and call the `Cache Syncer` with that ID:
```go title="ent/schema/dog.go"
func syncCache(next ent.Mutator) ent.Mutator {
	return hook.DogFunc(func(ctx context.Context, m *gen.DogMutation) (ent.Value, error) {
		cacheID, err := m.Client().Dog.Query().QueryOwner().QueryCache().OnlyID(ctx)
		if err != nil {
			return next.Mutate(ctx, m)
		}
		v, err := next.Mutate(ctx, m)
		if err == nil {
			m.Client().CacheSyncer.Sync(ctx, cacheID)
		}
		return v, err
	})
}
```

Full code for the cache syncer can be found [here](https://github.com/yonidavidson/ent-hooks-examples/blob/main/cache/syncer.go).

### Wrapping Up

In this post we reviewed a few examples for using hooks to embed three types of known problems we have when writing servers:
1. Keeping secrets safe
2. Data validation rules on Create
3. Offload heavy computation to a background task and store it in cache

I hope this will help you next time you approach these type of problems while using Ent.

Have questions? Need help with getting started? Feel free to join our [Ent Discord Server](https://discord.gg/qZmPgTE6RX).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
- join our [Slack channel](https://entgo.io/docs/slack/)

:::
