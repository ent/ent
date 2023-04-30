# Versioned Migration Example

The full reference example for https://entgo.io/docs/versioned-migrations#create-a-migration-files-generator.

### Migration directory

Versioned migration files exists under `ent/migrate/migrations` and follows the `golang-migrate` format.

### Changes to the Ent schema

1\. Change the `ent/schema`.

2\. Run `go generate ./ent`

### Generate new versioned migration

1\. Create a dev-database container if there is no one.

```shell
docker run --name migration --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=pass -e MYSQL_DATABASE=test -d mysql
```

2\. Generate a new versioned migration file:

```go
go run -mod=mod ent/migrate/main.go <name>
```

### Run migration linting

```bash
atlas migrate lint \
  --dev-url="mysql://root:pass@localhost:3306/test" \
  --dir="file://ent/migrate/migrations" \
  --latest=1
```
