---
id: predicates
title: Predicates
---

## Field Predicates

- **Bool**:
  - =, !=
- **Numeric**:
  - =, !=, >, <, >=, <=,
  - IN, NOT IN
- **Time**:
  - =, !=, >, <, >=, <=
  - IN, NOT IN
- **String**:
  - =, !=, >, <, >=, <=
  - IN, NOT IN
  - Contains, HasPrefix, HasSuffix
  - ContainsFold, EqualFold (**SQL** specific)
- **JSON**
  - =, !=
  - =, !=, >, <, >=, <= on nested values (JSON path).
  - Contains on nested values (JSON path).
  - HasKey, Len&lt;P>
  - `null` checks for nested values (JSON path).
- **Optional** fields:
  - IsNil, NotNil

## Edge Predicates

- **HasEdge**. For example, for edge named `owner` of type `Pet`, use:

  ```go
   client.Pet.
		Query().
		Where(pet.HasOwner()).
		All(ctx)
  ```

- **HasEdgeWith**. Add list of predicates for edge predicate.

  ```go
   client.Pet.
		Query().
		Where(pet.HasOwnerWith(user.Name("a8m"))).
		All(ctx)
  ```


## Negation (NOT)

```go
client.Pet.
	Query().
	Where(pet.Not(pet.NameHasPrefix("Ari"))).
	All(ctx)
```

## Disjunction (OR)

```go
client.Pet.
	Query().
	Where(
		pet.Or(
			pet.HasOwner(),
			pet.Not(pet.HasFriends()),
		)
	).
	All(ctx)
```

## Conjunction (AND)

```go
client.Pet.
	Query().
	Where(
		pet.And(
			pet.HasOwner(),
			pet.Not(pet.HasFriends()),
		)
	).
	All(ctx)
```

## Custom Predicates

Custom predicates can be useful if you want to write your own dialect-specific logic or to control the executed queries.

#### Get all pets of users 1, 2 and 3

```go
pets := client.Pet.
	Query().
	Where(func(s *sql.Selector) {
		s.Where(sql.InInts(pet.FieldOwnerID, 1, 2, 3))
	}).
	AllX(ctx)
```
The above code will produce the following SQL query:
```sql
SELECT DISTINCT `pets`.`id`, `pets`.`owner_id` FROM `pets` WHERE `owner_id` IN (1, 2, 3)
```

#### Count the number of users whose JSON field named `URL` contains the `Scheme` key

```go
count := client.User.
	Query().
	Where(func(s *sql.Selector) {
		s.Where(sqljson.HasKey(user.FieldURL, sqljson.Path("Scheme")))
	}).
	CountX(ctx)
```

The above code will produce the following SQL query:

```sql
-- PostgreSQL
SELECT COUNT(DISTINCT "users"."id") FROM "users" WHERE "url"->'Scheme' IS NOT NULL

-- SQLite and MySQL
SELECT COUNT(DISTINCT `users`.`id`) FROM `users` WHERE JSON_EXTRACT(`url`, "$.Scheme") IS NOT NULL
```

#### Get all users with a `"Tesla"` car

Consider an ent query such as: 

```go
users := client.User.Query().
	Where(user.HasCarWith(car.Model("Tesla"))).
	AllX(ctx)
```

This query can be rephrased in 3 different forms: `IN`, `EXISTS` and `JOIN`.

```go
// `IN` version.
users := client.User.Query().
	Where(func(s *sql.Selector) {
		t := sql.Table(car.Table)
        s.Where(
            sql.In(
                s.C(user.FieldID),
                sql.Select(t.C(user.FieldID)).From(t).Where(sql.EQ(t.C(car.FieldModel), "Tesla")),
            ),
        )
	}).
	AllX(ctx)

// `JOIN` version.
users := client.User.Query().
	Where(func(s *sql.Selector) {
		t := sql.Table(car.Table)
		s.Join(t).On(s.C(user.FieldID), t.C(car.FieldOwnerID))
		s.Where(sql.EQ(t.C(car.FieldModel), "Tesla"))
	}).
	AllX(ctx)

// `EXISTS` version.
users := client.User.Query().
	Where(func(s *sql.Selector) {
		t := sql.Table(car.Table)
		p := sql.And(
            sql.EQ(t.C(car.FieldModel), "Tesla"),
			sql.ColumnsEQ(s.C(user.FieldID), t.C(car.FieldOwnerID)),
		)
		s.Where(sql.Exists(sql.Select().From(t).Where(p)))
	}).
	AllX(ctx)
```

The above code will produce the following SQL query:

```sql
-- `IN` version.
SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users` WHERE `users`.`id` IN (SELECT `cars`.`owner_id` FROM `cars` WHERE `cars`.`model` = 'Tesla')

-- `JOIN` version.
SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users` JOIN `cars` ON `users`.`id` = `cars`.`owner_id` WHERE `cars`.`model` = 'Tesla'

-- `EXISTS` version.
SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users` WHERE EXISTS (SELECT * FROM `cars` WHERE `cars`.`model` = 'Tesla' AND `users`.`id` = `cars`.`owner_id`)
```

#### Get all pets where pet name contains a specific pattern

The generated code provides the `HasPrefix`, `HasSuffix`, `Contains`, and `ContainsFold` predicates for pattern matching.
However, in order to use the `LIKE` operator with a custom pattern, use the following example.

```go
pets := client.Pet.Query().
	Where(func(s *sql.Selector){
		s.Where(sql.Like(pet.Name,"_B%"))
	}).
	AllX(ctx)
```

The above code will produce the following SQL query:

```sql
SELECT DISTINCT `pets`.`id`, `pets`.`owner_id`, `pets`.`name`, `pets`.`age`, `pets`.`species` FROM `pets` WHERE `name` LIKE '_B%'
```

#### Custom SQL functions

In order to use built-in SQL functions such as `DATE()`, use one of the following options:

1\. Pass a dialect-aware predicate function using the `sql.P` option:

```go
users := client.User.Query().
	Select(user.FieldID).
	Where(func(s *sql.Selector) {
		s.Where(sql.P(func(b *sql.Builder) {
			b.WriteString("DATE(").Ident("last_login_at").WriteByte(')').WriteOp(OpGTE).Arg(value)
		}))
	}).
	AllX(ctx)
```

The above code will produce the following SQL query:

```sql
SELECT `id` FROM `users` WHERE DATE(`last_login_at`) >= ?
```

2\. Inline a predicate expression using the `ExprP()` option:

```go
users := client.User.Query().
	Select(user.FieldID).
	Where(func(s *sql.Selector) {
		s.Where(sql.ExprP("DATE(last_login_at) >= ?", value))
	}).
	AllX(ctx)
```

The above code will produce the same SQL query:

```sql
SELECT `id` FROM `users` WHERE DATE(`last_login_at`) >= ?
```

## JSON predicates

JSON predicates are not generated by default as part of the code generation. However, ent provides an official package
named [`sqljson`](https://pkg.go.dev/entgo.io/ent/dialect/sql/sqljson) for applying predicates on JSON columns using the
[custom predicates option](#custom-predicates).

#### Compare a JSON value

```go
sqljson.ValueEQ(user.FieldData, data)

sqljson.ValueEQ(user.FieldURL, "https", sqljson.Path("Scheme"))

sqljson.ValueNEQ(user.FieldData, content, sqljson.DotPath("attributes[1].body.content"))

sqljson.ValueGTE(user.FieldData, status.StatusBadRequest, sqljson.Path("response", "status"))
```

#### Check for the presence of a JSON key

```go
sqljson.HasKey(user.FieldData, sqljson.Path("attributes", "[1]", "body"))

sqljson.HasKey(user.FieldData, sqljson.DotPath("attributes[1].body"))
```

Note that, a key with the `null` literal as a value also matches this operation.

#### Check JSON `null` literals

```go
sqljson.ValueIsNull(user.FieldData)

sqljson.ValueIsNull(user.FieldData, sqljson.Path("attributes"))

sqljson.ValueIsNull(user.FieldData, sqljson.DotPath("attributes[1].body"))
```

Note that, the `ValueIsNull` returns true if the value is JSON `null`,
but not database `NULL`.

#### Compare the length of a JSON array

```go
sqljson.LenEQ(user.FieldAttrs, 2)

sql.Or(
	sqljson.LenGT(user.FieldData, 10, sqljson.Path("attributes")),
	sqljson.LenLT(user.FieldData, 20, sqljson.Path("attributes")),
)
```

#### Check if a JSON value contains another value

```go
sqljson.ValueContains(user.FieldData, data)

sqljson.ValueContains(user.FieldData, attrs, sqljson.Path("attributes"))

sqljson.ValueContains(user.FieldData, code, sqljson.DotPath("attributes[0].status_code"))
```

#### Check if a JSON string value contains a given substring or has a given suffix or prefix

```go
sqljson.StringContains(user.FieldURL, "github", sqljson.Path("host"))

sqljson.StringHasSuffix(user.FieldURL, ".com", sqljson.Path("host"))

sqljson.StringHasPrefix(user.FieldData, "20", sqljson.DotPath("attributes[0].status_code"))
```

#### Check if a JSON value is equal to any of the values in a list

```go
sqljson.ValueIn(user.FieldURL, []any{"https", "ftp"}, sqljson.Path("Scheme"))

sqljson.ValueNotIn(user.FieldURL, []any{"github", "gitlab"}, sqljson.Path("Host"))
```

## Comparing Fields

The `dialect/sql` package provides a set of comparison functions that can be used to compare fields in a query.

```go
client.Order.Query().
	Where(
		sql.FieldsEQ(order.FieldTotal, order.FieldTax),
        sql.FieldsNEQ(order.FieldTotal, order.FieldDiscount),
	).
	All(ctx)

client.Order.Query().
	Where(
		order.Or(
			sql.FieldsGT(order.FieldTotal, order.FieldTax),
			sql.FieldsLT(order.FieldTotal, order.FieldDiscount),
		),
	).
	All(ctx)
```

