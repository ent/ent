---
id: feature-flags
title: Feature Flags
sidebar_label: Feature Flags
---

The framework provides a collection of code-generation features that be added or removed using flags.

## Usage

Feature flags can be provided either by CLI flags or as arguments to the `gen` package. 

#### CLI

```console
go run -mod=mod entgo.io/ent/cmd/ent generate --feature privacy,entql ./ent/schema
```

#### Go

```go
// +build ignore

package main

import (
	"log"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
		},
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("static").
				Funcs(template.FuncMap{"title": strings.ToTitle}).
				ParseFiles("template/static.tmpl")),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

## List of Features

### Auto-Solve Merge Conflicts

The `schema/snapshot` option tells `entc` (ent codegen) to store a snapshot of the latest schema in an internal package,
and use it to automatically solve merge conflicts when user's schema can't be built.

This option can be added to a project using the `--feature schema/snapshot` flag, but please see
[ent/ent/issues/852](https://github.com/ent/ent/issues/852) to get more context about it.

### Privacy Layer

The privacy layer allows configuring privacy policy for queries and mutations of entities in the database.

This option can be added to a project using the `--feature privacy` flag, and you can learn more about in the
[privacy](privacy.mdx) documentation.

### EntQL Filtering

The `entql` option provides a generic and dynamic filtering capability at runtime for the different query builders.

This option can be added to a project using the `--feature entql` flag, and you can learn more about in the
[privacy](privacy.mdx#multi-tenancy) documentation.

### Named Edges

The `namedges` option provides an API for preloading edges with custom names.

This option can be added to a project using the `--feature namedges` flag, and you can learn more about in the
[Eager Loading](eager-load.mdx) documentation.

### Bidirectional Edge Refs

The `bidiedges` option guides Ent to set two-way references when eager-loading (O2M/O2O) edges.

This option can be added to a project using the `--feature bidiedges` flag.

:::note
Users that use the standard encoding/json.MarshalJSON should detach the circular references before calling `json.Marshal`.
:::

### Schema Config

The `sql/schemaconfig` option lets you pass alternate SQL database names to models. This is useful when your models don't all live under one database and are spread out across different schemas.

This option can be added to a project using the `--feature sql/schemaconfig` flag. Once you generate the code, you can now use a new option as such: 

```go
c, err := ent.Open(dialect, conn, ent.AlternateSchema(ent.SchemaConfig{
	User: "usersdb",
	Car: "carsdb",
}))
c.User.Query().All(ctx) // SELECT * FROM `usersdb`.`users`
c.Car.Query().All(ctx) 	// SELECT * FROM `carsdb`.`cars`
```

### Row-level Locks

The `sql/lock` option lets configure row-level locking using the SQL `SELECT ... FOR {UPDATE | SHARE}` syntax.

This option can be added to a project using the `--feature sql/lock` flag.

```go
tx, err := client.Tx(ctx)
if err != nil {
	log.Fatal(err)
}

tx.Pet.Query().
	Where(pet.Name(name)).
	ForUpdate().
	Only(ctx)

tx.Pet.Query().
	Where(pet.ID(id)).
	ForShare(
		sql.WithLockTables(pet.Table),
		sql.WithLockAction(sql.NoWait),
	).
	Only(ctx)
```

### Custom SQL Modifiers

The `sql/modifier` option lets add custom SQL modifiers to the builders and mutate the statements before they are executed.

This option can be added to a project using the `--feature sql/modifier` flag.

#### Modify Example 1

```go
client.Pet.
	Query().
	Modify(func(s *sql.Selector) {
		s.Select("SUM(LENGTH(name))")
	}).
	IntX(ctx)
```

The above code will produce the following SQL query:

```sql
SELECT SUM(LENGTH(name)) FROM `pet`
```

#### Select and Scan Dynamic Values

If you work with SQL modifiers and need to scan dynamic values not present in your Ent schema definition, such as
aggregation or custom ordering, you can apply `AppendSelect`/`AppendSelectAs` to the `sql.Selector`. You can later
access their values using the `Value` method defined on each entity:

```go {6,11}
const as = "name_length"

// Query the entity with the dynamic value.
p := client.Pet.Query().
	Modify(func(s *sql.Selector) {
		s.AppendSelectAs("LENGTH(name)", as)
	}).
	FirstX(ctx)

// Read the value from the entity.
n, err := p.Value(as)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Name length: %d == %d", n, len(p.Name))
```

#### Modify Example 2

```go
var p1 []struct {
	ent.Pet
	NameLength int `sql:"length"`
}

client.Pet.Query().
	Order(ent.Asc(pet.FieldID)).
	Modify(func(s *sql.Selector) {
		s.AppendSelect("LENGTH(name)")
	}).
	ScanX(ctx, &p1)
```

The above code will produce the following SQL query:

```sql
SELECT `pet`.*, LENGTH(name) FROM `pet` ORDER BY `pet`.`id` ASC
```

#### Modify Example 3

```go
var v []struct {
	Count     int       `json:"count"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

client.User.
	Query().
	Where(
        user.CreatedAtGT(x),
        user.CreatedAtLT(y),
	).
	Modify(func(s *sql.Selector) {
		s.Select(
			sql.As(sql.Count("*"), "count"),
			sql.As(sql.Sum("price"), "price"),
			sql.As("DATE(created_at)", "created_at"),
		).
		GroupBy("DATE(created_at)").
		OrderBy(sql.Desc("DATE(created_at)"))
	}).
	ScanX(ctx, &v)
```

The above code will produce the following SQL query:

```sql
SELECT
    COUNT(*) AS `count`,
    SUM(`price`) AS `price`,
    DATE(created_at) AS `created_at`
FROM
    `users`
WHERE
    `created_at` > x AND `created_at` < y
GROUP BY
    DATE(created_at)
ORDER BY
    DATE(created_at) DESC
```

#### Modify Example 4

```go
var gs []struct {
	ent.Group
	UsersCount int `sql:"users_count"`
}

client.Group.Query().
	Order(ent.Asc(group.FieldID)).
	Modify(func(s *sql.Selector) {
		t := sql.Table(group.UsersTable)
		s.LeftJoin(t).
			On(
				s.C(group.FieldID),
				t.C(group.UsersPrimaryKey[1]),
			).
			// Append the "users_count" column to the selected columns.
			AppendSelect(
				sql.As(sql.Count(t.C(group.UsersPrimaryKey[1])), "users_count"),
			).
			GroupBy(s.C(group.FieldID))
	}).
	ScanX(ctx, &gs)
```

The above code will produce the following SQL query:

```sql
SELECT
    `groups`.*,
    COUNT(`t1`.`group_id`) AS `users_count`
FROM
    `groups` LEFT JOIN `user_groups` AS `t1`
ON
    `groups`.`id` = `t1`.`group_id`
GROUP BY
    `groups`.`id`
ORDER BY
    `groups`.`id` ASC
```


#### Modify Example 5

```go
client.User.Update().
	Modify(func(s *sql.UpdateBuilder) {
		s.Set(user.FieldName, sql.Expr(fmt.Sprintf("UPPER(%s)", user.FieldName)))
	}).
	ExecX(ctx)
```

The above code will produce the following SQL query:

```sql
UPDATE `users` SET `name` = UPPER(`name`)
```

#### Modify Example 6

```go
client.User.Update().
	Modify(func(u *sql.UpdateBuilder) {
		u.Set(user.FieldID, sql.ExprFunc(func(b *sql.Builder) {
			b.Ident(user.FieldID).WriteOp(sql.OpAdd).Arg(1)
		}))
		u.OrderBy(sql.Desc(user.FieldID))
	}).
	ExecX(ctx)
```

The above code will produce the following SQL query:

```sql
UPDATE `users` SET `id` = `id` + 1 ORDER BY `id` DESC
```

#### Modify Example 7

Append elements to the `values` array in a JSON column:

```go
client.User.Update().
	Modify(func(u *sql.UpdateBuilder) {
        sqljson.Append(u, user.FieldTags, []string{"tag1", "tag2"}, sqljson.Path("values"))
	}).
	ExecX(ctx)
```

The above code will produce the following SQL query:

```sql
UPDATE `users` SET `tags` = CASE
    WHEN (JSON_TYPE(JSON_EXTRACT(`tags`, '$.values')) IS NULL OR JSON_TYPE(JSON_EXTRACT(`tags`, '$.values')) = 'NULL')
    THEN JSON_SET(`tags`, '$.values', JSON_ARRAY(?, ?))
    ELSE JSON_ARRAY_APPEND(`tags`, '$.values', ?, '$.values', ?) END
    WHERE `id` = ?
```

### SQL Raw API

The `sql/execquery` option allows executing statements using the `ExecContext`/`QueryContext` methods of the underlying
driver. For full documentation, see: [DB.ExecContext](https://pkg.go.dev/database/sql#DB.ExecContext), and
[DB.QueryContext](https://pkg.go.dev/database/sql#DB.QueryContext).

```go
// From ent.Client.
if _, err := client.ExecContext(ctx, "TRUNCATE t1"); err != nil {
	return err
}

// From ent.Tx.
tx, err := client.Tx(ctx)
if err != nil {
	return err
}
if err := tx.User.Create().Exec(ctx); err != nil {
	return err
}
if _, err := tx.ExecContext("SAVEPOINT user_created"); err != nil {
	return err
}
// ...
```

:::warning Note
Statements executed using `ExecContext`/`QueryContext` do not go through Ent, and may skip fundamental layers in your
application such as hooks, privacy (authorization), and validators.
:::

### Upsert

The `sql/upsert` option lets configure upsert and bulk-upsert logic using the SQL `ON CONFLICT` / `ON DUPLICATE KEY`
syntax. For full documentation, go to the [Upsert API](crud.mdx#upsert-one).

This option can be added to a project using the `--feature sql/upsert` flag.

```go
// Use the new values that were set on create.
id, err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflict().
	UpdateNewValues().
	ID(ctx)

// In PostgreSQL, the conflict target is required.
err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflictColumns(user.FieldName).
	UpdateNewValues().
	Exec(ctx)

// Bulk upsert is also supported.
client.User.
	CreateBulk(builders...).
	OnConflict(
		sql.ConflictWhere(...),
		sql.UpdateWhere(...),
	).
	UpdateNewValues().
	Exec(ctx)

// INSERT INTO "users" (...) VALUES ... ON CONFLICT WHERE ... DO UPDATE SET ... WHERE ...
```

### Globally Unique ID

By default, SQL primary-keys start from 1 for each table; which means that multiple entities of different types
can share the same ID. Unlike AWS Neptune, where node IDs are UUIDs.

This does not work well if you work with [GraphQL](https://graphql.org/learn/schema/#scalar-types), which requires 
the object ID to be unique.

To enable the Universal-IDs support for your project, simply use the `--feature sql/globalid` flag.

:::warning Note
If you have used the `migrate.WithGlobalUniqueID(true)` migration option in the past, please read 
[this guide](globalid-migrate) before you switch your project to use the new globalid feature.
:::

**How does it work?** `ent` migration allocates a 1<<32 range for the IDs of each entity (table),
and store this information alongside your generated code (`internal/globalid.go`). For example, type `A` will have the 
range of `[1,4294967296)` for its IDs, and type `B` will have the range of `[4294967296,8589934592)`, etc.

Note that if this option is enabled, the maximum number of possible tables is **65535**. 
