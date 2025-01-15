## Use PostgreSQL Views in Ent Schema Using Atlas `composite_schema`

This example demonstrates how to define an `ent.View` but keep its definition and creation externally using the Atlas
`composite_schema` data source. This option allows users controlling the way views are created, their dependencies, and
the SQL definition itself, that might not be supported by the build-in SQL builders.

The second approach is to define an `ent.View` with its SQL definition (`AS ...`) in the Ent schema. The big advantage
of the non-composite_schema approach is that the `CREATE VIEW` correctness is checked during migration and not during queries.
For example, if one of the `ent.Field`s defined in your ent/schema does not exist in your SQL definition, PostgreSQL will return the
following error:

```text
create "clean_users" view: pq: CREATE VIEW specifies more column names than columns
```