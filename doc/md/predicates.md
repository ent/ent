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
SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users` WHERE `users`.`id` IN (SELECT `cars`.`id` FROM `cars` WHERE `cars`.`model` = 'Tesla')

-- `JOIN` version.
SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users` JOIN `cars` ON `users`.`id` = `cars`.`owner_id` WHERE `cars`.`model` = 'Tesla'

-- `EXISTS` version.
SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name` FROM `users` WHERE EXISTS (SELECT * FROM `cars` WHERE `cars`.`model` = 'Tesla' AND `users`.`id` = `cars`.`owner_id`)
```

#### Get all pets where pet name contains a specific pattern 
The generated code provides the `HasPrefix`, `HasSuffix`, `Contains`, and `ContainsFold` predicates for pattern matching. However, in order to use the `LIKE` operator with a custom pattern, use the following example.

```go
pets := client.Pet.Query().
	Where(func(s *sql.Selector){
		s.Where(sql.Like(pet.Name,"_B%"))
	}).
	AllX(ctx)
```

The above code will produce the following SQL query:

```sql
SELECT DISTINCT `pets`.`id`, `pets`.`owner_id`, `pets`.`name`, `pets`.`age`, `pets`.`species` FROM `pets` WHERE `name` LIKE 'B'
```
