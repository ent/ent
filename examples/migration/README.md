# Versioned Migration Example

The full reference example for https://entgo.io/docs/versioned-migrations#create-a-migration-files-generator.

### Migration directory

Versioned migration files exists under `ent/migrate/migrations` and follows the `golang-migrate` format.

### Changes to the Ent schema

1\. Change the `ent/schema`.

2\. Run `go generate ./ent`


### Generate new versioned migration

```go
go run -mod=mod ent/migrate/main.go <name>
```