---
title: A few examples for Ent hooks  
author: Yoni Davidson
authorURL: "https://github.com/yonidavidson"
authorImageURL: "https://avatars0.githubusercontent.com/u/5472778"
authorTwitter: yonidavidson
image: "https://entgo.io/images/assets/ogent/1.png"
---
Hooks are a popular subject in our blog posts
[building-observable-ent-application-with-prometheus](https://entgo.io/blog/2021/08/12/building-observable-ent-application-with-prometheus)
[sync-to-external-data-systems-using-hooks](https://entgo.io/blog/2021/11/1/sync-to-external-data-systems-using-hooks)
but their reactive nature makes them a tool that is usually affiliated with more advanced users.
My goal in this blog post is to show you a few use cases that hooks can help you solve when building servers and provide you with a few go-to recipes.
These examples can be later reviewed in this [repo](https://github.com/yonidavidson/ent-hooks-examples).

### So first, let us review in a few words what are hooks?

[Hooks](https://entgo.io/docs/hooks) are a feature of Ent that allows adding custom logic before and after operations that change the data entities.

A mutation is an operation that changes something in the database.
There are 5 types of mutations:
1. Create.
2. UpdateOne.
3. Update.
4. DeleteOne.
5. Delete.

Hooks are functions that get an [ent.Mutator](https://pkg.go.dev/entgo.io/ent#Mutator) and return a mutator back.
They function similar to the popular [HTTP middleware pattern](https://github.com/go-chi/chi#middleware-handlers).

```go
package example

import (
	"context"

	"entgo.io/ent"
)

func exampleHook() ent.Hook {
	//use this to init your hook
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

In Ent, there are two types of mutation hooks - schema hooks and runtime hooks. Schema hooks are mainly used for defining
custom mutation logic on a specific entity type, for example, syncing entity creation to another system. Runtime hooks, on the other hand, are used
to define more global logic for adding things like logging, metrics, tracing, etc.

In this post, I will focus mainly on schemas hooks.

### So why use them?

In ent, we try and push the entity's logic down to the schema level, this allows us to define everything about an entity close to its creation. Entity properties could be:
1. Valid values.
2. Types.
3. Allowed mutations.
4. Privacy.
5. Required Side effect (heavy compute for example).

---
Many times, it’s more comfortable to define that in the RPC level, a controller provides a context and all the mutations
and validation are done before changing the value.
Yet, as the system gets bigger and more developers are working on it making sure that these validations and calculations are done
whenever the entity is changing becomes harder and requires more pre-known knowledge about each entity to work with and
can cause unintentional bugs or inconsistencies in the data that are only discovered in runtime.
The way ent, based on the experience of bigger companies like Facebook, decided to handle it is to keep this information
as close as possible to the entity, therefore hooks are just that, a piece of code that runs before and after an entity 
is changing and defines that additional behavior.

## Let us do some examples

For this examples we use a simple schema of users and dogs.
Each user can have multiple pets and each pet has an owner.
We have also added a cache entity that will be used later.
as can be seen here:
```go
// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("connection_string").
			NotEmpty(),
		field.String("password").
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
```go
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
```go
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

During each example I'll provide a test and a recipe, I really suggest starting with a test to see that the hook does what you need it to do,
Unlike many other parts of the codebase, hooks do one thing only and should be simple to test.

### Keeping our secrets safe:

Many times, we need to store strings that include secrets, a good example can be a connection
string to a database, usually in the form of ```mysql://root:pass@localhost:3306)```, here, for example, it’s very
easy to see  that we don't want the password to unintentionally be retrieved by a query unless required by the server
(for example to connect to a database).
The required behavior is best described by this test:
```go
func TestUserConnectionStringHook(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite,
		"file:TestSchemaConfHooks?mode=memory&cache=shared&_fk=1",
	)
	u := c.User.Create().SetName("Yoni").SetConnectionString("mysql://root:pass@localhost:3306)").SaveX(ctx)
	require.Equal(t, "mysql://root:****@localhost:3306)", u.ConnectionString)
	require.Equal(t, "pass", u.Password)
	require.Equal(t, "mysql://root:pass@localhost:3306)", u.FullConnectionString())
}
```

As you can see, inserting a connection string updates the string to have `****` instead of the password and stores the 
password in the ```Password``` field that is defined as [sensitive](https://entgo.io/docs/schema-fields#sensitive-fields).
To have  the connection string in full valid form we created a utility function ```FullConnectionString()```.

#### Recipe:

First thing, add the hook to your schema:
```go
// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(clearConnectionString,
			hook.And(hook.Or(hook.HasOp(ent.OpUpdateOne), hook.HasOp(ent.OpCreate)), hook.HasFields(user.FieldConnectionString)),
		),
	}
}
```
* Do not forget to run go generate or the hook will not be registered.

A few hook helpers are used here, ```clearConnectionString``` is the name of the hook function, and it will run if a user 
is Updated or Created and the mutation has the field ```connection_string``` (if we change only the name of the user there is no need to do anything right?).

The hook:
```go
func clearConnectionString(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
		cs, ok := m.ConnectionString()
		if !ok {
			return next.Mutate(ctx, m)
		}
		sp := strings.Split(cs, "@")
		if len(sp) != 2 {
			return next.Mutate(ctx, m)
		}
		sp = strings.Split(sp[0], ":")
		if len(sp) != 3 {
			return next.Mutate(ctx, m)
		}
		pass := sp[2]
		m.SetPassword(pass)
		m.SetConnectionString(strings.ReplaceAll(cs, pass, "****"))
		return next.Mutate(ctx, m)
	})
} 
```
We validate that the mutation has a connection string (if not just let the mutation continue).
A bit of string acrobatics, and we have the password extracted from the string, we store it in another field - the ```password``` and
change the ```connection_string``` with a masked version. 

To get the full connection string we add a utility function 
```go
package ent

import "strings"

// FullConnectionString returns the connection string with an unmasked password.
func (u *User) FullConnectionString() string {
	return strings.ReplaceAll(u.ConnectionString, "****", u.Password)
}
```

BTW, you can see that it's OK to add files to the ent directory (not only in schema folder) since the ent generate command only appends/changes ent
generated files.


### Special Validators:

Ent provides [validators](https://entgo.io/docs/schema-fields#validators) for fields, yet sometimes the logic can contain
dependencies that are more applicative and require a broader context.
For example, in our use case, we would like to make sure that dogs names don’t start with the first 2 letters of the owners name 
(Best practice when you don't want your dog to run in circles every time someone calls your name).

We start with a simple test
```go
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
```go
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
[edge field](https://entgo.io/docs/schema-edges#edge-field), this provides the hooks a trigger on when the edge's id changes.

The hook:
```go
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
		if owner.Name[0:1] == dn[0:1] {
			return nil, errors.New("invalid dog name")
		}
		return next.Mutate(ctx, m)
	})
}
```
A simple name validator and if we see a problem we cancel the mutation by returning an error before ```next.Mutate(ctx,m)```.

### Offloading Long Computations:

Some cases, we have fields in our database that require long computation, we would like to make sure that this data is always updated once
a change in the data requires it.

A test for it will look like this:
```go
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
This test is a bit more complicated so lets review it line by line:
First, I create a new ```cache syncer``` (remember we have a cache entity in our schema).

This ```Cache Syncer``` can do heavy computations  while running in a different  context.
We create an ent client and inject our ```Cache Syncer```, it is disabled by default.
After that, we start it and provide it with the ent client, the reason for that is to allow it to change our data (modify our cache entity).
We create a few entities and update the dog's name, that will trigger a cache sync.
Finally, we close the ```Cache Syncer```, this will block until all pending hooks are completed.
The assertion validates a value of "walks" larger than 0 (not -1).

#### Recipe:

First, add a hook:
```go
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

The hooks itself will query for the owners cache ID and call the ```Cache Syncer``` with that ID:
```go
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

For injecting the dependency in ent we follow [injecting external dependencies](https://entgo.io/docs/code-gen#external-dependencies)
section and update our  ```entc.go```:

```go
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

The interface just exposes the sync command:
```go
package hook

import "context"

type Syncer interface {
	Sync(ctx context.Context, cacheID int)
}
```

Full code for the cache syncer can be found [here](https://github.com/yonidavidson/ent-hooks-examples/blob/main/cache/syncer.go).

### Wrapping Up

In this post we reviewed a few examples for using hooks to embed 3 types of known problems we have when writing servers:
1. Keeping secrets safe.
2. Validating rules over our data  before saving in our DB.
3. Offloading heavy computes and store in cache.

I hope this will help you next time you approach this type of problems using ent.

Have questions? Need help with getting started? Feel free to join our [Ent Discord Server](https://discord.gg/qZmPgTE6RX).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
- join our [Slack channel](https://entgo.io/docs/slack/)

:::
