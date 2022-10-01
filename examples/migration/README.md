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

### Run linting

```bash
go run -mod=mod ariga.io/atlas/cmd/atlas@master migrate lint \
  --dev-url="mysql://root:pass@localhost:3306/test" \
  --dir="file://ent/migrate/migrations" \
  --dir-format="golang-migrate" \
  --latest=1
```
