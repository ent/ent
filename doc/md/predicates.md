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

Custom predicates can be useful if you want to write your own dialect-specific logic.

```go
pets := client.Pet.
	Query().
	Where(predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.InInts(pet.OwnerColumn, 1, 2, 3))
	})).
	AllX(ctx)

users := client.User.
	Query().
	Where(predicate.User(func(s *sql.Selector) {
		s.Where(sqljson.HasKey(user.FieldURL, sqljson.Path("Scheme")))
	})).
	AllX(ctx)
```
