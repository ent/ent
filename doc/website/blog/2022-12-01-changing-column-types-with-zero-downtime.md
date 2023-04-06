---
title: Changing a column’s type with zero-downtime using Atlas
author: Ronen Lubin (ronenlu)
authorURL: "https://github.com/ronenlu"
authorImageURL: "https://avatars.githubusercontent.com/u/63970571?v=4"
---
Changing a column's type in a database schema might seem trivial at first glance, but it is actually a risky operation
that can cause compatibility issues between the server and the database. In this blogpost,
I will explore how developers can perform this type of change without causing downtime to their application.

Recently, while working on a feature for [Ariga Cloud](https://atlasgo.io/cloud/getting-started),
I was required to change the type of an Ent field from an unstructured blob to a structured JSON field.
Changing the column type was necessary in order to use [JSON Predicates](https://entgo.io/docs/predicates/#json-predicates)
for more efficient queries.

The original schema looked like this on our cloud product’s schema visualization diagram:

![tutorial image 1](https://entgo.io/images/assets/migrate-column-type/users_table.png)

In our case, we couldn't just copy the data naively to the new column, since the data is not compatible
with the new column type (blob data may not be convertible to JSON).

In the past, it was considered acceptable to stop the server, migrate the database schema to the next version,
and only then start the server with the new version that is compatible with the new database schema.
Today, business requirements often dictate that applications must provide higher availability, leaving engineering teams
with the challenge of executing changes like this with zero-downtime.

The common pattern to satisfy this kind of requirement, as defined in "[Evolutionary Database Design](https://www.martinfowler.com/articles/evodb.html)" by Martin Fowler,
is to use a "transition phase".
> A transition phase is "a period of time when the database supports both the old access pattern and the new ones simultaneously.
This allows older systems time to migrate over to the new structures at their own pace", as illustrated by this diagram:

![tutorial image 2](https://www.martinfowler.com/articles/evodb/stages_refactoring.jpg)
Credit: martinfowler.com 

We planned the change in 5 simple steps, all of which are backward-compatible:
* Creating a new column named `meta_json` with the JSON type.
* Deploy a version of the application that performs dual-writes. Every new record or update is written to both the new column and the old column, while reads still happen from the old column.
* Backfill data from the old column to the new one.
* Deploy a version of the application that reads from the new column.
* Delete the old column.

### Versioned migrations
In our project we are using Ent’s [versioned migrations](https://entgo.io/docs/versioned-migrations) workflow for
managing the database schema. Versioned migrations provide teams with granular control on how changes to the application database schema are made.
This level of control will be very useful in implementing our plan. If your project uses [Automatic Migrations](https://entgo.io/docs/migrate) and you would like to follow along,
[first upgrade](https://entgo.io/docs/versioned/intro) your project to use versioned migrations.

:::note
The same can be done with automatic migrations as well by using the [Data Migrations](https://entgo.io/docs/data-migrations/#automatic-migrations) feature,
however this post is focusing on versioned migrations
:::

### Creating a JSON column with Ent:
First, we will add a new JSON Ent type to the user schema.

``` go title="types/types.go"
type Meta struct {
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
```
``` go title="ent/schema/user.go"
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Bytes("meta"),
        field.JSON("meta_json", &types.Meta{}).Optional(),
    }
}
```

Next, we run codegen to update the application schema:
``` shell
go generate ./...
```

Next, we run our [automatic migration planning](https://entgo.io/docs/versioned/auto-plan) script that generates a set of
migration files containing the necessary SQL statements to migrate the database to the newest version.
``` shell
go run -mod=mod ent/migrate/main.go add_json_meta_column
```

The resulted migration file describing the change:
``` sql
-- modify "users" table
ALTER TABLE `users` ADD COLUMN `meta_json` json NULL;
```

Now, we will apply the created migration file using [Atlas](https://atlasgo.io):
``` shell
atlas migrate apply \
  --dir "file://ent/migrate/migrations"
  --url mysql://root:pass@localhost:3306/ent
```

As a result, we have the following schema in our database:

![tutorial image 3](https://entgo.io/images/assets/migrate-column-type/users_table_add_column.png)

### Start writing to both columns

After generating the JSON type, we will start writing to the new column:
``` diff
-		err := client.User.Create().
-			SetMeta(input.Meta).
-			Exec(ctx)
+		var meta types.Meta
+		if err := json.Unmarshal(input.Meta, &meta); err != nil {
+			return nil, err
+		}
+		err := client.User.Create().
+			SetMetaJSON(&meta).
+			Exec(ctx)
```

To ensure that values written to the new column `meta_json` are replicated to the old column, we can utilize Ent’s
[Schema Hooks](https://entgo.io/docs/hooks/#schema-hooks) feature. This adds blank import `ent/runtime` in your main to
[register the hook](https://entgo.io/docs/hooks/#hooks-registration) and avoid circular import:
``` go
// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
					meta, ok := m.MetaJSON()
					if !ok {
						return next.Mutate(ctx, m)
					}
					if b, err := json.Marshal(meta); err != nil {
						return nil, err
					}
					m.SetMeta(b)
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}
```

After ensuring writes to both fields we can safely deploy to production.

### Backfill values from old column

Now in our production database we have two columns: one storing the meta object as a blob and another storing it as a JSON.
The second column may have null values since the JSON column was only added recently, therefore we need to backfill it with the old column’s values.

To do so, we manually create a SQL migration file that will fill values in the new JSON column from the old blob column.

:::note
You can also write Go code that generates this data migration file by using the [WriteDriver](https://entgo.io/docs/data-migrations#versioned-migrations).
:::

Create a new empty migration file:
``` shell
atlas migrate new --dir file://ent/migrate/migrations
```

For every row in the users table with a null JSON value (i.e: rows added before the creation of the new column), we try
to parse the meta object into a valid JSON. If we succeed, we will fill the `meta_json` column with the resulting value, otherwise we will mark it empty.

Our next step is to edit the migration file:
``` sql
UPDATE users
SET meta_json = CASE
        -- when meta is valid json stores it as is.
        WHEN JSON_VALID(cast(meta as char)) = 1 THEN cast(cast(meta as char) as json)
        -- if meta is not valid json, store it as an empty object.
        ELSE JSON_SET('{}')
    END
WHERE meta_json is null;
```

Rehash the migration directory after changing a migration file:
``` shell
atlas migrate hash --dir "file://ent/mirate/migrations"
```

We can test the migration file by executing all the previous migration files on a local database, seed it with temporary data, and
apply the last migration to ensure our migration file works as expected.

After testing we apply the migration file:
``` shell
atlas migrate apply \
  --dir "file://ent/migrate/migrations"
  --url mysql://root:pass@localhost:3306/ent 
```

Now, we will deploy to production once more.

### Redirect reads to the new column and delete old blob column

Now that we have values in the `meta_json` column, we can change the reads from the old field to the new field.

Instead of decoding the data from `user.meta` on each read, just use the `meta_json` field:
``` diff 	
-       var meta types.Meta
-       if err = json.Unmarshal(user.Meta, &meta); err != nil {
-	        return nil, err
-       }
-       if meta.CreateTime.Before(time.Unix(0, 0)) {
-	        return nil, errors.New("invalid create time")
-       }
+       if user.MetaJSON.CreateTime.Before(time.Unix(0, 0)) {
+	        return nil, errors.New("invalid create time")
+       }
```

After redirecting the reads we will deploy the changes to production.

### Delete the old column

It is now possible to remove the field describing the old column from the Ent schema, since as we are no longer using it.
``` diff
func (User) Fields() []ent.Field {
    return []ent.Field{
-     field.Bytes("meta"),
      field.JSON("meta_json", &types.Meta{}).Optional(),
    }
}

``` 

Generate the Ent schema again with the [Drop Column](https://entgo.io/docs/migrate/#drop-resources) feature enabled.
``` shell
go run -mod=mod ent/migrate/main.go drop_user_meta_column
```

Now that we have properly created our new field, redirected writes, backfilled it and dropped the old column -
we are ready for the final deployment. All that’s left is to merge our code into version control and deploy to production!

### Wrapping up

In this post, we discussed how to change a column type in the production database with zero downtime using Atlas’s version migrations integrated with Ent.

Have questions? Need help with getting started? Feel free to join
our [Ent Discord Server](https://discord.gg/qZmPgTE6RX).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
:::