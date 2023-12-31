# Versioned Migration Example

The full reference example for https://entgo.io/docs/versioned-migrations#create-a-migration-files-generator.

### Migration directory

Versioned migration files exists under `ent/migrate/migrations`.

### Changes to the Ent schema

1\. Change the `ent/schema`.

2\. Run `go generate ./ent`

### Generate a new migration file

```bash
atlas migrate diff <migration_name> \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://mysql/8/ent"
```

### Run migration linting

```bash
atlas migrate lint \
  --dev-url="mysql://root:pass@localhost:3306/test" \
  --dir="file://ent/migrate/migrations" \
  --latest=1
```
