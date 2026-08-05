---
id: blob-storage
title: Blob Storage
---

## Overview

Ent's `field.Blob` type stores field data in external blob storage (e.g., S3, GCS, local filesystem)
rather than in the database. The database only stores a **key** (a short string) referencing the blob.
This keeps database rows small while supporting arbitrarily large binary or text payloads.

## Quick Start

### 1. Define a schema with blob fields

```go
package schema

import (
  "crypto"

  "entgo.io/ent"
  "entgo.io/ent/schema/field"
)

type Document struct {
  ent.Schema
}

func (Document) Fields() []ent.Field {
  return []ent.Field{
    field.String("name").Unique(),
    field.Blob("content"),
    field.Blob("avatar").Optional(),
  }
}
```

### 2. Configure blob storage

Each entity type with blob fields requires a `BlobOpener` — a function that opens a
blob bucket for a given field name:

```go
import (
  "context"

  "myapp/ent"
  "myapp/ent/document"

  _ "gocloud.dev/blob/s3blob" // or fileblob, gcsblob, etc.
)

client := ent.NewClient(
  ent.Driver(drv),
  ent.WithBlobOpeners(ent.BlobOpeners{
    Document: func(ctx context.Context, field string) (ent.Blob, error) {
      switch field {
      case document.FieldContent:
        return blob.OpenBucket(ctx, "s3://my-bucket/content")
      case document.FieldAvatar:
        return blob.OpenBucket(ctx, "s3://my-bucket/avatars")
      default:
        return nil, fmt.Errorf("unknown blob field: %s", field)
      }
    },
  }),
)
```

### 3. Use the generated API

```go
// Create — pass []byte directly.
doc := client.Document.Create().
  SetName("readme").
  SetContent([]byte("Hello, World!")).
  SaveX(ctx)

// Read — the struct field holds the loaded value.
fmt.Println(string(doc.Content)) // "Hello, World!"

// Update
doc = doc.Update().
  SetContent([]byte("Updated content")).
  SaveX(ctx)
```

## `field.Blob` API

| Method | Description |
|--------|-------------|
| `Optional()` | Field is nullable; not required on create. |
| `Nillable()` | Struct field is a pointer (distinguishes zero value from unset). |
| `Immutable()` | Field can only be set on create, not updated. |
| `Lazy()` | Mutation accepts `io.Reader`; struct field omitted; use Reader method to read. |
| `HashKey(h)` | Content-addressable key via hash (default: `crypto.SHA256`). |
| `UUIDKey()` | Random UUID v7 key per write. |
| `DualWrite(...)` | Migration mode: write to both blob storage and database column. |
| `GoType(typ)` | Override the default `[]byte` Go type. |
| `ValueScanner(vs)` | Custom codec between Go type and raw bytes (required for non-`[]byte`/`string` GoType). |
| `StorageKey(key)` | Override the database column name. |
| `StructTag(s)` | Set the struct tag on the generated field. |
| `Comment(c)` | Set the field comment. |
| `Annotations(...)` | Attach codegen annotations. |
| `Deprecated(...)` | Mark the field as deprecated. |

## Key Strategies

Every blob field requires a **key function** that determines the storage key for a given
piece of data. By default (when neither `UUIDKey` nor `HashKey` is called), blobs use
SHA-256 content hashing.

### HashKey (default)

Content-addressable storage: the data is hashed to produce the key. Identical content
always maps to the same key, enabling deduplication and write-skip optimizations on update.

```go
field.Blob("content").HashKey(crypto.SHA256)
```

On update, if the new content produces the same hash as the existing key, the write to
blob storage is skipped entirely.

### UUIDKey

Each write generates a new random UUID (v7) as the storage key. This guarantees uniqueness
but does not deduplicate.

```go
field.Blob("content").UUIDKey()
```

## Lazy Fields

By default, blob data is automatically loaded into the entity struct field when the row is
scanned from the database. For large blobs where you don't always need the data in memory,
use `Lazy()`:

```go
field.Blob("content").Lazy()
```

With `Lazy()`:
- The **mutation builder** accepts an `io.Reader` (which is fully buffered before writing).
- The entity **struct field is omitted** — data is not loaded on scan.
- A **Reader method** (e.g., `ContentReader`) is generated to explicitly open a reader from storage.

```go
// Create with io.Reader.
doc := client.Document.Create().
  SetName("large-file").
  SetContent(bytes.NewReader(largeData)).
  SaveX(ctx)

// Read explicitly via the Reader method.
rc, err := doc.ContentReader(ctx)
if err != nil { ... }
defer rc.Close()
data, _ := io.ReadAll(rc)
```

## DualWrite (Migration Mode)

`DualWrite()` preserves the original bytes column alongside the blob key column. This is
useful when migrating an existing column to blob storage:

```go
field.Blob("payload").
  DualWrite(map[string]string{
    dialect.MySQL:    "json",
    dialect.Postgres: "jsonb",
    dialect.SQLite:   "json",
  })
```

In DualWrite mode:
- **Writes** go to both blob storage and the database column.
- **Reads** prefer blob storage (if a key exists) and fall back to the database column.

The optional `columnType` argument overrides the database column type per dialect to prevent
schema drift when migrating from an existing column definition.

## Custom GoType

Blob fields default to `[]byte`. You can override the Go type:

```go
// String type (automatic conversion).
field.Blob("description").GoType("")

// Custom struct with a ValueScanner.
field.Blob("config").
  GoType(&MyConfig{}).
  ValueScanner(configScanner{})
```

When using a custom `GoType` other than `string`, a `ValueScanner` must be provided
to encode/decode between the Go type and the raw bytes stored in blob storage.

## Blob Interface

Any blob storage backend must implement the `ent.Blob` interface:

```go
type Blob interface {
  NewReader(ctx context.Context, key string) (io.ReadCloser, error)
  NewWriter(ctx context.Context, key string) (io.WriteCloser, error)
  Delete(ctx context.Context, key string) error
  Close() error
}
```

- `NewReader` should return `fs.ErrNotExist` (or a wrapping error) when the key does not exist.
- `Delete` should return `nil` (not an error) when the key does not exist.

The [Go CDK `blob` package](https://gocloud.dev/howto/blob/) provides implementations
for S3, GCS, Azure, and local filesystem that satisfy this interface via a thin adapter.

## Lifecycle and Cleanup

### Shared objects

A blob key is not private to the row that wrote it. With `HashKey` (the default), the key
*is* the content, so any two rows holding identical content point at one object in storage.
That is what makes deduplication work — and it means an object may only be removed once
the last row referencing it lets go.

Every cleanup path therefore counts references before deleting: it queries the key columns
for the keys it is about to remove, and skips the ones other rows still hold. A blob written
ahead of a failed INSERT is left alone if it turns out to be the object an existing row
already points at.

The count is taken while the statement (or transaction) is still open, so a row inserted
concurrently between the count and the delete is not seen. Closing that window entirely
requires reference counting or a periodic sweep of the bucket against the key columns.

### Create

Blobs are written **before** the database row is inserted. If the SQL INSERT fails
(e.g., constraint violation), the generated code deletes the just-written blobs — except
any that rows in the database already reference.

### Update

On update:
1. New blob data is written to storage.
2. The SQL UPDATE executes.
3. On success, old blobs are deleted unless another row still references them.
4. On SQL failure, newly-written blobs are rolled back (deleted) under the same rule.

When using `HashKey`, if the content hasn't changed (same hash), the write is skipped entirely.

### Delete

Generated delete builders query existing blob keys before deleting the row, then remove
the blobs from storage after a successful SQL DELETE — keeping any object that rows outside
the deleted set still reference.

## OnConflict (Upsert)

Blob fields work with `OnConflict` / upsert builders. Blobs are written to storage
**before** the SQL executes. If the INSERT succeeds (no conflict), the new key is stored.
If there is a conflict, behavior depends on the conflict action and key strategy.

The following per-field methods are generated on the upsert builder:

- **`Update<Field>()`** — Sets the blob key column (and data column for DualWrite fields) to
  the value provided on create.
- **`Clear<Field>()`** — Nulls the blob key column (and data column for DualWrite). Only
  generated for optional fields. **Note:** this does not delete the old blob from storage —
  the previously-referenced blob becomes orphaned. Use a regular Update with `Clear<Field>()`
  if you need storage cleanup.

```go
client.Document.Create().
  SetName("readme").
  SetContent([]byte("data")).
  SetAttachment([]byte("file")).
  OnConflict().
  UpdateContent().       // update only the content blob
  UpdateAttachment().    // update only the attachment blob
  ExecX(ctx)
```

You can also use `UpdateNewValues()` to update all fields (including all blob key columns)
at once.

### With HashKey (content-addressable)

Since the key is derived from content, identical data always produces the same key:

- **`Update<Field>()`** — The SQL updates the key column to the new key. If the content
  matches what the row already held, the key is unchanged and it is a no-op in storage.
  If the content differs, the row moves to the new key and the old object is left behind —
  `OnConflict` does not query old keys the way a regular Update does.
- **`Ignore()`** — The SQL sets each column to itself. If the content differs from the
  existing row's, the blob written to storage has no database reference.
- **`DoNothing()`** — No row is inserted, which surfaces as `sql.ErrNoRows`. The blob
  written to storage is deleted, unless the conflicting row already references that same
  key — the common case when the content is identical.

Upserting identical content is idempotent in storage: the write lands on the key that is
already there, and cleanup leaves it in place because the existing row references it.

```go
client.Document.Create().
  SetName("readme").
  SetContent([]byte("data")).
  OnConflict().
  UpdateContent().
  ExecX(ctx)
```

### With UUIDKey (random keys)

Each write generates a new unique key. This means the blob written before the SQL is always
at a **new** key that didn't previously exist:

- **`Update<Field>()`** — The SQL updates the key column to the new UUID key. The old blob
  at the previous key becomes orphaned and is **not** automatically cleaned up (OnConflict
  does not query old keys like a regular Update does).
- **`Ignore()`** — The SQL succeeds without changing anything; the row keeps its existing
  key. The newly-written blob is orphaned in storage with no database reference.
- **`DoNothing()`** — No row is inserted, which surfaces as `sql.ErrNoRows`. Because a fresh
  UUID is referenced by nobody, the newly-written blob is deleted.

For these reasons, `UUIDKey` is **not recommended** with `OnConflict`. If you need upsert
semantics, prefer `HashKey` which is inherently idempotent.
