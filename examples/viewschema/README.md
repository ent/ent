## Define PostgreSQL Views in Ent Schema

This example demonstrates how to define an `ent.View` with its SQL definition (`AS ...`) specified in the Ent schema.

The second approach is to define an `ent.View` but keep its definition and creation externally using the Atlas `composite_schema`
data source (see `examples/viewcomposite` for more information).

The main advantage of this example approach is that the `CREATE VIEW` correctness is checked during migration and not during queries.
For example, if one of the `ent.Field`s defined in your ent/schema does not exist in your SQL definition, PostgreSQL will return the
following error:

```text
create "clean_users" view: pq: CREATE VIEW specifies more column names than columns
```