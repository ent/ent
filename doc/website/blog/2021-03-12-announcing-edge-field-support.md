---
title: Announcing Edge-field Support in v0.7.0
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
---
Over the past few months, there has been much discussion in the Ent project [issues](https://github.com/ent/ent/issues) about adding support for the retrieval of the foreign key field when retrieving entities with One-to-One or One-to-Many edges.  We are happy to announce that as of [v0.7.0](https://github.com/ent/ent/releases/tag/v0.7.0) ent supports this feature.

### Before Edge-field Support

Prior to merging this branch, a user that wanted to retrieve the foreign-key field for an entity needed to use eager-loading. Suppose our schema looked like this:

```go
// ent/schema/user.go:

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pets", Pet.Type).
			Ref("owner"),
	}
}

// ent/schema/pet.go

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			Unique().
			Required(),
	}
}
```

The schema describes two related entities: `User` and `Pet`, with a One-to-Many edge between them: a user can own many pets and a pet can have one owner.

When retrieving pets from the data storage, it is common for developers to want to access the foreign-key field on the pet. However, because this field is created implicitly from the `owner` edge it was automatically accessible when retrieving an entity. To retrieve this from the storage a developer needed to do something like:

```go
func Test(t *testing.T) {
    ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer c.Close()
	
	// Create the User
	u := c.User.Create().
		SetUserName("rotem").
		SaveX(ctx)

	// Create the Pet
	p := c.Pet.
		Create().
		SetOwner(u). // Associate with the user
		SetName("donut").
		SaveX(ctx)

	petWithOwnerId := c.Pet.Query().
		Where(pet.ID(p.ID)).
		WithOwner(func(query *ent.UserQuery) {
			query.Select(user.FieldID)
		}).
		OnlyX(ctx)
	fmt.Println(petWithOwnerId.Edges.Owner.ID)
	// Output: 1
}
```

Aside from being very verbose, retrieving the pet with the owner this way was inefficient in-terms of database queries. If we execute the query with the `.Debug()` we can see the DB queries ent generates to satisfy this call:

```sql
SELECT DISTINCT `pets`.`id`, `pets`.`name`, `pets`.`pet_owner` FROM `pets` WHERE `pets`.`id` = ? LIMIT 2 
SELECT DISTINCT `users`.`id` FROM `users` WHERE `users`.`id` IN (?)
```

In this example, Ent first retrieves the Pet with an ID of `1`, then redundantly fetches the `id` field from the `users` table for users with an ID of `1`.

### With Edge-field Support

[Edge-field support](https://entgo.io/docs/schema-edges/#edge-field) greatly simplifies and improves the efficiency of this flow. With this feature, developers can define the foreign key field as part of the schemas `Fields()`, and by using the `.Field(..)` modifier on the edge definition instruct Ent to expose and map the foreign column to this field.  So, in our example schema, we would modify it to be:

```go
// user.go stays the same

// pet.go
// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Int("owner_id"), // <-- explicitly add the field we want to contain the FK
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			Field("owner_id"). // <-- tell ent which field holds the reference to the owner
			Unique().
			Required(),
	}
}
```

In order to update our client code we need to re-run code generation:

```sql
go generate ./...
```

We can now modify our query to be much simpler:

```go
func Test(t *testing.T) {
	ctx := context.Background()
	c := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer c.Close()

	u := c.User.Create().
		SetUserName("rotem").
		SaveX(ctx)

	p := c.Pet.Create().
		SetOwner(u).
		SetName("donut").
		SaveX(ctx)

	petWithOwnerId := c.Pet.GetX(ctx, p.ID) // <-- Simply retrieve the Pet

	fmt.Println(petWithOwnerId.OwnerID)
	// Output: 1
}
```

Running with the `.Debug()` modifier we can see that the DB queries make more sense now:

```sql
SELECT DISTINCT `pets`.`id`, `pets`.`name`, `pets`.`owner_id` FROM `pets` WHERE `pets`.`id` = ? LIMIT 2
```

Hooray 🎉!

### Migrating Existing Schemas to Edge Fields

If you are already using Ent with an existing schema, you may already have O2M relations whose foreign-key columns already exist in your database.  Depending on how you configured your schema, chances are that they may be stored in a column by a different name than the field you are now adding. For instance, you want to create an `owner_id` field, but Ent auto-created the column foreign-key column as `pet_owner`.

To check what column name Ent is using for this field you can look in the `./ent/migrate/schema.go` file:

```go
PetsColumns = []*schema.Column{
	{Name: "id", Type: field.TypeInt, Increment: true},
	{Name: "name", Type: field.TypeString},
	{Name: "pet_owner", Type: field.TypeInt, Nullable: true}, // <-- this is our FK
}
```

To allow for a smooth migration, you must explicitly tell Ent to keep using the existing column name. You can do this by using the `StorageKey` modifier (either on the field or on the edge). For example:

```go
// In schema/pet.go:

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Int("owner_id").
			StorageKey("pet_owner"), // <-- explicitly set the column name
	}
}
```

In the near future we plan to implement Schema Versioning, which will store the history of schema changes alongside the code. Having this information will allow ent to support such migrations in an automatic and predictable way.

### Wrapping Up

Edge-field support is readily available and can be installed by `go get -u entgo.io/ent@v0.7.0`.

Many thanks 🙏 to all the good people who took the time to give feedback and helped design this feature properly: [Alex Snast](https://github.com/alexsn), [Ruben de Vries](https://github.com/rubensayshi), [Marwan Sulaiman](https://github.com/marwan-at-work), [Andy Day](https://github.com/adayNU), [Sebastian Fekete](https://github.com/aight8) and [Joe Harvey](https://github.com/errorhandler).

### For more Ent news and updates:

- Follow us on [twitter.com/entgo_io](https://twitter.com/entgo_io)
- Subscribe to our [newsletter](https://entgo.substack.com/)
- Join us on #ent on the [Gophers slack](https://app.slack.com/client/T029RQSE6/C01FMSQDT53)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
