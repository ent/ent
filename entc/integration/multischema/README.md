## Multi-schema support for versioned migration

Login first:

```shell
atlas login
```

Generate migrations:

```shell
atlas migrate diff --to ent://versioned/schema \
  --dev-url docker://mysql/8 \
  --format '{{ sql . "  " }}'
```

Apply migrations:

```shell
atlas migrate apply \
  --url mysql://root:pass@:3308/ \
  --dir file://versioned/migrate/migrations
```

Inspect/Visualize the schema:

```shell
atlas schema inspect \
  --url ent://versioned/schema \
  --dev-url docker://mysql/8 \
  -w
```
