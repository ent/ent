// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/integration/blob"
	"entgo.io/ent/entc/integration/ent"
	"entgo.io/ent/entc/integration/ent/document"
	"entgo.io/ent/entc/integration/ent/enttest"
	"entgo.io/ent/entc/integration/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	_ "gocloud.dev/blob/fileblob"
)

// readBlob reads all data from a blob field reader and closes it.
func readBlob(t *testing.T, rc io.ReadCloser, err error) []byte {
	t.Helper()
	require.NoError(t, err)
	if rc == nil {
		return nil
	}
	data, readErr := io.ReadAll(rc)
	require.NoError(t, readErr)
	require.NoError(t, rc.Close())
	return data
}

// blobContent is a shorthand for readBlob(t, entity.Field(ctx)).
func blobContent(t *testing.T, fn func(context.Context) (io.ReadCloser, error), ctx context.Context) []byte {
	t.Helper()
	rc, err := fn(ctx)
	return readBlob(t, rc, err)
}

// blobDir creates a temp directory with subdirectories for blob fields.
func blobDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "documents"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "thumbnails"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "attachments"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "metadata"), 0o755))
	return dir
}

// newBlobOpeners returns BlobOpeners that route to the correct
// file:// bucket based on the field name, using absolute paths under dir.
func newBlobOpeners(dir string) ent.BlobOpeners {
	return ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			switch field {
			case document.FieldContent:
				return blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
			case document.FieldThumbnail:
				return blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "thumbnails"))
			case document.FieldAttachment:
				return blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "attachments"))
			case document.FieldMetadata:
				return blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "metadata"))
			default:
				return nil, fmt.Errorf("unknown blob field: %s", field)
			}
		},
	}
}

// setupBlob creates a temp directory with subdirectories for each blob field,
// opens an in-memory SQLite client with auto-migration, and registers cleanup.
func setupBlob(t *testing.T, opts ...ent.Option) (*ent.Client, context.Context, string) {
	t.Helper()
	dir := blobDir(t)
	allOpts := append([]ent.Option{ent.WithBlobOpeners(newBlobOpeners(dir))}, opts...)
	entOpts := []enttest.Option{
		enttest.WithMigrateOptions(migrate.WithDropIndex(true), migrate.WithDropColumn(true)),
		enttest.WithOptions(allOpts...),
	}
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1",
		entOpts...,
	)
	t.Cleanup(func() {
		client.Close()
	})
	return client, context.Background(), dir
}

func TestBlobCreateAndRead(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	data := []byte("Hello from blob integration test!")
	doc := client.Document.Create().
		SetName("test-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Read the blob back through the entity method.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, data, got)
}

func TestBlobQueryAndRead(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	data := []byte("queried blob content")
	created := client.Document.Create().
		SetName("query-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Query it back from the database (no content column, just ID/name).
	queried := client.Document.GetX(ctx, created.ID)

	// Read blob content from the queried entity.
	got := blobContent(t, queried.ContentReader, ctx)
	require.Equal(t, data, got)
}

func TestBlobUpdateData(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	v1 := []byte("version 1")
	doc := client.Document.Create().
		SetName("update-doc").
		SetContent(bytes.NewReader(v1)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Update blob data through mutation (overwrites same key).
	v2 := []byte("version 2 - updated via mutation")
	doc = doc.Update().
		SetContent(bytes.NewReader(v2)).
		SaveX(ctx)

	// Read the new blob content.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, v2, got)
}

func TestBlobRequiredValidation(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Creating a document without required blob fields should fail.
	_, err := client.Document.Create().
		SetName("no-blob-doc").
		Save(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "missing required field")
}

func TestBlobMultipleDocuments(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	contents := []string{
		"document one content",
		"document two content - larger payload with more data",
		"document three",
	}

	var ids []int
	for i, c := range contents {
		doc := client.Document.Create().
			SetName("multi-" + string(rune('a'+i))).
			SetContent(strings.NewReader(c)).
			SetThumbnail(bytes.NewReader([]byte("thumb"))).
			SetAttachment([]byte("att")).
			SaveX(ctx)
		ids = append(ids, doc.ID)
	}

	// Read each document and verify content is correct.
	for i, id := range ids {
		doc := client.Document.GetX(ctx, id)
		got := blobContent(t, doc.ContentReader, ctx)
		require.Equal(t, contents[i], string(got), "document %d content mismatch", i)
	}
}

func TestBlobBulkCreate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	bulk := make([]*ent.DocumentCreate, 5)
	for i := range bulk {
		bulk[i] = client.Document.Create().
			SetName(strings.Repeat("bulk-", i+1)).
			SetContent(bytes.NewReader([]byte(strings.Repeat("x", i+1)))).
			SetThumbnail(bytes.NewReader([]byte("thumb"))).
			SetAttachment([]byte("att"))
	}
	docs, err := client.Document.CreateBulk(bulk...).Save(ctx)
	require.NoError(t, err)
	require.Len(t, docs, 5)

	// Verify each document's blob can be read back with correct data.
	for i, doc := range docs {
		got := blobContent(t, doc.ContentReader, ctx)
		require.Equal(t, []byte(strings.Repeat("x", i+1)), got)
	}
}

func TestBlobThumbnailCreateAndRead(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	thumbData := []byte("fake-png-thumbnail-data")
	doc := client.Document.Create().
		SetName("doc-with-thumb").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader(thumbData)).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	got := blobContent(t, doc.ThumbnailReader, ctx)
	require.Equal(t, thumbData, got)
}

func TestBlobBothFields(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	contentData := []byte("document body content")
	thumbData := []byte("thumbnail image bytes")
	doc := client.Document.Create().
		SetName("doc-both").
		SetContent(bytes.NewReader(contentData)).
		SetThumbnail(bytes.NewReader(thumbData)).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Read content.
	cGot := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, contentData, cGot)

	// Read thumbnail.
	tGot := blobContent(t, doc.ThumbnailReader, ctx)
	require.Equal(t, thumbData, tGot)

	// Update only thumbnail, content should remain unchanged.
	newThumb := []byte("updated thumbnail")
	doc = doc.Update().
		SetThumbnail(bytes.NewReader(newThumb)).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	cGot2 := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, contentData, cGot2)

	tGot2 := blobContent(t, doc.ThumbnailReader, ctx)
	require.Equal(t, newThumb, tGot2)
}

func TestBlobBulkCreateBothFields(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	bulk := make([]*ent.DocumentCreate, 3)
	for i := range bulk {
		bulk[i] = client.Document.Create().
			SetName(strings.Repeat("bulk-both-", i+1)).
			SetContent(bytes.NewReader([]byte(strings.Repeat("c", i+1)))).
			SetThumbnail(bytes.NewReader([]byte(strings.Repeat("t", i+1)))).
			SetAttachment([]byte("att"))
	}
	docs, err := client.Document.CreateBulk(bulk...).Save(ctx)
	require.NoError(t, err)
	require.Len(t, docs, 3)

	for i, doc := range docs {
		cGot := blobContent(t, doc.ContentReader, ctx)
		require.Equal(t, []byte(strings.Repeat("c", i+1)), cGot)

		tGot := blobContent(t, doc.ThumbnailReader, ctx)
		require.Equal(t, []byte(strings.Repeat("t", i+1)), tGot)
	}
}

func TestBlobWriter(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create a document with content via mutation.
	doc := client.Document.Create().
		SetName("writer-doc").
		SetContent(strings.NewReader("initial")).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Overwrite via ContentWriter (bypasses mutation pipeline).
	w, err := doc.ContentWriter(ctx)
	require.NoError(t, err)
	_, err = io.Copy(w, strings.NewReader("overwritten via writer"))
	require.NoError(t, err)
	require.NoError(t, w.Close())

	// Read back should reflect the overwritten data.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("overwritten via writer"), got)
}

func TestBlobReader(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	data := []byte("streaming read test data")
	doc := client.Document.Create().
		SetName("reader-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Read back via Content (io.ReadCloser).
	r, err := doc.ContentReader(ctx)
	require.NoError(t, err)
	defer r.Close()

	got, err := io.ReadAll(r)
	require.NoError(t, err)
	require.Equal(t, data, got)
}

func TestBlobWriterThenReader(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create a document with initial content, then overwrite via Writer.
	doc := client.Document.Create().
		SetName("write-then-read").
		SetContent(strings.NewReader("initial")).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Write via writer.
	w, err := doc.ContentWriter(ctx)
	require.NoError(t, err)
	_, err = w.Write([]byte("hello "))
	require.NoError(t, err)
	_, err = w.Write([]byte("world"))
	require.NoError(t, err)
	require.NoError(t, w.Close())

	// Read via reader.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("hello world"), got)
}

func TestBlobWriterOptions(t *testing.T) {
	// Verify that a custom opener with WriterOptions works for roundtrip.
	dir := blobDir(t)
	openerWithOpts := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
			if err != nil {
				return nil, err
			}
			return b.WithWriterOptions(nil), nil
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(openerWithOpts))

	data := []byte("content with writer options")
	doc := client.Document.Create().
		SetName("opts-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, data, got)
}

func TestBlobWriterOptionsApplied(t *testing.T) {
	// Verify that WriterOptions are applied on both create and update.
	dir := blobDir(t)
	openerWithOpts := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
			if err != nil {
				return nil, err
			}
			return b.WithWriterOptions(nil), nil
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(openerWithOpts))

	data := []byte("content with writer options")
	doc := client.Document.Create().
		SetName("opts-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Verify the data was written successfully and can be read back.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, data, got)

	// Update also uses WriterOpts.
	v2 := []byte("updated with writer options")
	doc = doc.Update().SetContent(bytes.NewReader(v2)).SaveX(ctx)
	got = blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, v2, got)
}

func TestBlobEncryption(t *testing.T) {
	// Demonstrate per-tenant encryption using a master seed + tenant from context.
	// Each tenant derives a unique AES-256 key via SHA-256(tenant || seed).
	// Data written by one tenant cannot be decrypted by another.
	dir := blobDir(t)

	// Master seed shared across all tenants (kept secret server-side).
	masterSeed := []byte("super-secret-master-seed-for-test")

	encryptedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			var subdir string
			switch field {
			case document.FieldContent:
				subdir = "documents"
			case document.FieldThumbnail:
				subdir = "thumbnails"
			case document.FieldAttachment:
				subdir = "attachments"
			default:
				return nil, fmt.Errorf("unknown blob field: %s", field)
			}
			b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, subdir))
			if err != nil {
				return nil, err
			}
			// Wrap the bucket with per-tenant encryption.
			return blob.NewEncrypted(b, masterSeed), nil
		},
	}

	// Tenant "acme" writes encrypted data.
	client, _, _ := setupBlob(t, ent.WithBlobOpeners(encryptedOpeners))
	acmeCtx := blob.WithTenant(context.Background(), "acme")

	plaintext := []byte("top secret document content for acme")
	doc := client.Document.Create().
		SetName("encrypted-doc").
		SetContent(bytes.NewReader(plaintext)).
		SetThumbnail(bytes.NewReader([]byte("secret-thumb"))).
		SetAttachment([]byte("att")).
		SaveX(acmeCtx)

	// Reading with the same tenant returns decrypted plaintext.
	got := blobContent(t, doc.ContentReader, acmeCtx)
	require.Equal(t, plaintext, got)

	gotThumb := blobContent(t, doc.ThumbnailReader, acmeCtx)
	require.Equal(t, []byte("secret-thumb"), gotThumb)

	// A different tenant ("evil") cannot decrypt acme's data — the derived key
	// differs, so AES-CTR produces garbage (not the original plaintext).
	evilCtx := blob.WithTenant(context.Background(), "evil")
	evilData := blobContent(t, doc.ContentReader, evilCtx)
	require.NotEqual(t, plaintext, evilData, "different tenant must not read acme's plaintext")

	// No tenant in context → error.
	noTenantCtx := context.Background()
	_, err := doc.ContentReader(noTenantCtx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "requires a tenant")

	// Update works through encryption too.
	v2 := []byte("updated secret content")
	doc = doc.Update().SetContent(bytes.NewReader(v2)).SaveX(acmeCtx)
	got = blobContent(t, doc.ContentReader, acmeCtx)
	require.Equal(t, v2, got)

	// ContentWriter also encrypts with tenant key.
	w, err := doc.ContentWriter(acmeCtx)
	require.NoError(t, err)
	_, err = w.Write([]byte("written via writer"))
	require.NoError(t, err)
	require.NoError(t, w.Close())

	got = blobContent(t, doc.ContentReader, acmeCtx)
	require.Equal(t, []byte("written via writer"), got)
}

func TestBlobPrefix(t *testing.T) {
	dir := blobDir(t)
	openers := newBlobOpeners(dir)
	prefixedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			if field == "content" {
				b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
				if err != nil {
					return nil, err
				}
				return b.Prefixed("tenant-1/"), nil
			}
			return openers.Document(ctx, field)
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(prefixedOpeners))

	data := []byte("prefixed blob content")
	doc := client.Document.Create().
		SetName("prefixed-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Read through the entity — uses the same opener.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, data, got)

	// Create a second client with the default opener — reading should return nil (not found).
	defaultClient := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithOptions(ent.WithBlobOpeners(openers)),
	)
	t.Cleanup(func() { defaultClient.Close() })
	doc2 := defaultClient.Document.Query().OnlyX(ctx)
	got = blobContent(t, doc2.ContentReader, ctx)
	require.Nil(t, got, "expected nil when reading without prefix")
}

func TestBlobPrefixUpdate(t *testing.T) {
	dir := blobDir(t)
	openers := newBlobOpeners(dir)
	prefixedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			if field == "content" {
				b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
				if err != nil {
					return nil, err
				}
				return b.Prefixed("tenant-2/"), nil
			}
			return openers.Document(ctx, field)
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(prefixedOpeners))

	data := []byte("initial")
	doc := client.Document.Create().
		SetName("prefix-update-doc").
		SetContent(bytes.NewReader(data)).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	updated := []byte("updated under prefix")
	doc = doc.Update().SetContent(bytes.NewReader(updated)).SaveX(ctx)

	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, updated, got)
}

func TestBlobPrefixBulkCreate(t *testing.T) {
	dir := blobDir(t)
	openers := newBlobOpeners(dir)
	prefixedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			if field == "content" {
				b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
				if err != nil {
					return nil, err
				}
				return b.Prefixed("bulk-tenant/"), nil
			}
			return openers.Document(ctx, field)
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(prefixedOpeners))

	docs := client.Document.CreateBulk(
		client.Document.Create().SetName("bulk-p1").SetContent(strings.NewReader("b1")).SetThumbnail(bytes.NewReader([]byte("t1"))).SetAttachment([]byte("att")),
		client.Document.Create().SetName("bulk-p2").SetContent(strings.NewReader("b2")).SetThumbnail(bytes.NewReader([]byte("t2"))).SetAttachment([]byte("att")),
	).SaveX(ctx)

	for i, doc := range docs {
		got := blobContent(t, doc.ContentReader, ctx)
		require.Equal(t, []byte(fmt.Sprintf("b%d", i+1)), got)
	}
}

func TestBlobPrefixWriterReader(t *testing.T) {
	dir := blobDir(t)
	openers := newBlobOpeners(dir)
	prefixedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			if field == "content" {
				b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, "documents"))
				if err != nil {
					return nil, err
				}
				return b.Prefixed("rw-tenant/"), nil
			}
			return openers.Document(ctx, field)
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(prefixedOpeners))

	doc := client.Document.Create().
		SetName("prefix-rw-doc").
		SetContent(strings.NewReader("initial")).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Write via ContentWriter.
	w, err := doc.ContentWriter(ctx)
	require.NoError(t, err)
	_, err = io.Copy(w, strings.NewReader("writer-prefixed"))
	require.NoError(t, err)
	require.NoError(t, w.Close())

	// Read via Content.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("writer-prefixed"), got)
}

func TestBlobDualWriteCreate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	data := []byte("dual-write attachment data")
	doc := client.Document.Create().
		SetName("dualwrite-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment(data).
		SaveX(ctx)

	// 1. The public struct field holds the bytes.
	require.Equal(t, data, doc.Attachment)

	// 2. Read via AttachmentReader — reads from blob storage.
	got := blobContent(t, doc.AttachmentReader, ctx)
	require.Equal(t, data, got)
}

func TestBlobDualWriteQueryFallback(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	data := []byte("attachment for query fallback test")
	created := client.Document.Create().
		SetName("dualwrite-query").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment(data).
		SaveX(ctx)

	// Query back — the entity has both the bytes column and the blob key.
	queried := client.Document.GetX(ctx, created.ID)

	// The struct field is populated from the SQL column.
	require.Equal(t, data, queried.Attachment)

	// Reading via AttachmentReader also works (reads from blob).
	got := blobContent(t, queried.AttachmentReader, ctx)
	require.Equal(t, data, got)
}

func TestBlobDualWriteUpdate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	v1 := []byte("attachment v1")
	doc := client.Document.Create().
		SetName("dualwrite-update").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment(v1).
		SaveX(ctx)

	// Update attachment.
	v2 := []byte("attachment v2 - updated")
	doc = doc.Update().
		SetAttachment(v2).
		SaveX(ctx)

	// The struct field holds the updated bytes.
	require.Equal(t, v2, doc.Attachment)

	// Read updated data from blob.
	got := blobContent(t, doc.AttachmentReader, ctx)
	require.Equal(t, v2, got)
}

func TestBlobDualWriteBulkCreate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	bulk := make([]*ent.DocumentCreate, 3)
	for i := range bulk {
		bulk[i] = client.Document.Create().
			SetName(fmt.Sprintf("bulk-dw-%d", i)).
			SetContent(bytes.NewReader([]byte("c"))).
			SetThumbnail(bytes.NewReader([]byte("t"))).
			SetAttachment([]byte(strings.Repeat("a", i+1)))
	}
	docs, err := client.Document.CreateBulk(bulk...).Save(ctx)
	require.NoError(t, err)
	require.Len(t, docs, 3)

	for i, doc := range docs {
		// Struct field has the data.
		require.Equal(t, []byte(strings.Repeat("a", i+1)), doc.Attachment)

		// Blob reader also works.
		got := blobContent(t, doc.AttachmentReader, ctx)
		require.Equal(t, []byte(strings.Repeat("a", i+1)), got)
	}
}

func TestBlobLoadOnScanCreate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	meta := []byte(`{"version": 1, "tags": ["test"]}`)
	doc := client.Document.Create().
		SetName("loadonscan-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(ctx)

	// The struct field is populated on create (from the mutation value).
	require.Equal(t, meta, doc.Metadata)

	// Reading via MetadataReader also works.
	got := blobContent(t, doc.MetadataReader, ctx)
	require.Equal(t, meta, got)
}

func TestBlobLoadOnScanQuery(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	meta := []byte(`{"loaded": true}`)
	created := client.Document.Create().
		SetName("loadonscan-query").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(ctx)

	// Query back — LoadOnScan auto-loads the metadata from blob storage.
	queried := client.Document.GetX(ctx, created.ID)
	require.Equal(t, meta, queried.Metadata)
}

func TestBlobLoadOnScanUpdate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	v1 := []byte(`{"v": 1}`)
	doc := client.Document.Create().
		SetName("loadonscan-update").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(v1).
		SaveX(ctx)

	v2 := []byte(`{"v": 2}`)
	doc = doc.Update().SetMetadata(v2).SaveX(ctx)
	require.Equal(t, v2, doc.Metadata)

	// Query confirms the update.
	queried := client.Document.GetX(ctx, doc.ID)
	require.Equal(t, v2, queried.Metadata)
}

func TestBlobLoadOnScanNil(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create without metadata (optional).
	doc := client.Document.Create().
		SetName("loadonscan-nil").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// The struct field is nil when no metadata was set.
	require.Nil(t, doc.Metadata)

	// Query also returns nil.
	queried := client.Document.GetX(ctx, doc.ID)
	require.Nil(t, queried.Metadata)
}

func TestBlobLoadOnScanHook(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Register a hook that observes the metadata field in the mutation.
	var hookSeen []byte
	client.Document.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if dm, ok := m.(*ent.DocumentMutation); ok {
				if v, exists := dm.Metadata(); exists {
					hookSeen = v
				}
			}
			return next.Mutate(ctx, m)
		})
	})

	meta := []byte(`{"hook": true}`)
	client.Document.Create().
		SetName("hook-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(ctx)

	// The hook saw the metadata value.
	require.Equal(t, meta, hookSeen)
}

func TestBlobLoadOnScanBulkCreate(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	bulk := make([]*ent.DocumentCreate, 3)
	for i := range bulk {
		bulk[i] = client.Document.Create().
			SetName(fmt.Sprintf("bulk-los-%d", i)).
			SetContent(bytes.NewReader([]byte("c"))).
			SetThumbnail(bytes.NewReader([]byte("t"))).
			SetAttachment([]byte("att")).
			SetMetadata([]byte(fmt.Sprintf(`{"i": %d}`, i)))
	}
	docs, err := client.Document.CreateBulk(bulk...).Save(ctx)
	require.NoError(t, err)
	require.Len(t, docs, 3)

	for i, doc := range docs {
		// Struct field is populated on bulk create.
		require.Equal(t, []byte(fmt.Sprintf(`{"i": %d}`, i)), doc.Metadata)

		// Blob reader also works.
		got := blobContent(t, doc.MetadataReader, ctx)
		require.Equal(t, []byte(fmt.Sprintf(`{"i": %d}`, i)), got)
	}
}

func TestBlobLoadOnScanQueryAll(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create multiple documents with metadata.
	for i := range 4 {
		client.Document.Create().
			SetName(fmt.Sprintf("queryall-%d", i)).
			SetContent(bytes.NewReader([]byte("c"))).
			SetThumbnail(bytes.NewReader([]byte("t"))).
			SetAttachment([]byte("att")).
			SetMetadata([]byte(fmt.Sprintf(`{"n": %d}`, i))).
			SaveX(ctx)
	}

	// Query all — each entity should have its metadata auto-loaded.
	docs := client.Document.Query().
		Order(ent.Asc(document.FieldName)).
		AllX(ctx)
	require.Len(t, docs, 4)
	for i, doc := range docs {
		require.Equal(t, []byte(fmt.Sprintf(`{"n": %d}`, i)), doc.Metadata)
	}
}

func TestBlobLoadOnScanClear(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	meta := []byte(`{"clear": true}`)
	doc := client.Document.Create().
		SetName("clear-meta").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(ctx)
	require.Equal(t, meta, doc.Metadata)

	// Clear the metadata.
	doc = doc.Update().ClearMetadata().SaveX(ctx)

	// After clearing, the struct field is nil (metadata_key was not set on the update).
	require.Nil(t, doc.Metadata)

	// Query back — also nil.
	queried := client.Document.GetX(ctx, doc.ID)
	require.Nil(t, queried.Metadata)
}

func TestBlobLoadOnScanUpdateHook(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	doc := client.Document.Create().
		SetName("update-hook-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata([]byte(`{"v": 1}`)).
		SaveX(ctx)

	// Register hook that observes metadata on update.
	var hookSeen []byte
	client.Document.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if dm, ok := m.(*ent.DocumentMutation); ok {
				if v, exists := dm.Metadata(); exists {
					hookSeen = v
				}
			}
			return next.Mutate(ctx, m)
		})
	})

	v2 := []byte(`{"v": 2}`)
	doc.Update().SetMetadata(v2).SaveX(ctx)
	require.Equal(t, v2, hookSeen)
}

func TestBlobLoadOnScanNillableSet(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	meta := []byte(`{"nillable": true}`)
	doc := client.Document.Create().
		SetName("nillable-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetNillableMetadata(&meta).
		SaveX(ctx)
	require.Equal(t, meta, doc.Metadata)

	// SetNillableMetadata with nil does nothing.
	doc = doc.Update().SetNillableMetadata(nil).SaveX(ctx)
	// The metadata should remain unchanged (loaded from blob).
	require.Equal(t, meta, doc.Metadata)
}

func TestBlobLoadOnScanWriter(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	meta := []byte(`{"writer": false}`)
	doc := client.Document.Create().
		SetName("writer-meta-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(ctx)

	// Overwrite metadata via MetadataWriter (bypasses mutation pipeline).
	w, err := doc.MetadataWriter(ctx)
	require.NoError(t, err)
	newMeta := []byte(`{"writer": true}`)
	_, err = w.Write(newMeta)
	require.NoError(t, err)
	require.NoError(t, w.Close())

	// Read back should reflect the overwritten data.
	got := blobContent(t, doc.MetadataReader, ctx)
	require.Equal(t, newMeta, got)
}

func TestBlobLoadOnScanReaderNilKey(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create without metadata — key is nil.
	doc := client.Document.Create().
		SetName("nil-key-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// MetadataReader should return an error when key is nil.
	_, err := doc.MetadataReader(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nil or empty")

	// MetadataWriter should also return an error.
	_, err = doc.MetadataWriter(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nil or empty")
}

func TestBlobDualWriteHook(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Register a hook that observes the attachment field.
	var hookSeen []byte
	client.Document.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if dm, ok := m.(*ent.DocumentMutation); ok {
				if v, exists := dm.Attachment(); exists {
					hookSeen = v
				}
			}
			return next.Mutate(ctx, m)
		})
	})

	att := []byte("hook-attachment-data")
	client.Document.Create().
		SetName("dw-hook-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment(att).
		SaveX(ctx)

	require.Equal(t, att, hookSeen)
}

func TestBlobDualWriteNillableSet(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	att := []byte("initial-attachment")
	doc := client.Document.Create().
		SetName("dw-nillable-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment(att).
		SaveX(ctx)

	// SetNillableAttachment with nil does nothing.
	doc = doc.Update().SetNillableAttachment(nil).SaveX(ctx)
	require.Equal(t, att, doc.Attachment)

	// SetNillableAttachment with a value updates.
	newAtt := []byte("updated-attachment")
	doc = doc.Update().SetNillableAttachment(&newAtt).SaveX(ctx)
	require.Equal(t, newAtt, doc.Attachment)

	// Blob storage also has the new value.
	got := blobContent(t, doc.AttachmentReader, ctx)
	require.Equal(t, newAtt, got)
}

func TestBlobInlineKeyCheckContent(t *testing.T) {
	// Verify that ContentReader/Writer errors when the entity has no key set.
	// This tests the inlined nil-and-empty key check.
	client, ctx, _ := setupBlob(t)

	doc := client.Document.Create().
		SetName("key-check-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// The entity has a key set, so this should work fine.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("content"), got)

	// Construct a bare Document without keys to test the guard.
	bare := &ent.Document{}
	_, err := bare.ContentReader(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nil or empty")

	_, err = bare.ContentWriter(ctx)
	require.Error(t, err)
	require.Contains(t, err.Error(), "nil or empty")
}

func TestBlobLoadOnScanWithPrefix(t *testing.T) {
	dir := blobDir(t)
	prefixedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			var subdir string
			switch field {
			case document.FieldContent:
				subdir = "documents"
			case document.FieldThumbnail:
				subdir = "thumbnails"
			case document.FieldAttachment:
				subdir = "attachments"
			case document.FieldMetadata:
				subdir = "metadata"
			default:
				return nil, fmt.Errorf("unknown blob field: %s", field)
			}
			b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, subdir))
			if err != nil {
				return nil, err
			}
			return b.Prefixed("tenant-x/"), nil
		},
	}
	client, ctx, _ := setupBlob(t, ent.WithBlobOpeners(prefixedOpeners))

	meta := []byte(`{"prefixed": true}`)
	doc := client.Document.Create().
		SetName("prefixed-los-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(ctx)

	// Struct field populated on create.
	require.Equal(t, meta, doc.Metadata)

	// Query auto-loads from prefixed blob.
	queried := client.Document.GetX(ctx, doc.ID)
	require.Equal(t, meta, queried.Metadata)

	// Update also works through prefix.
	v2 := []byte(`{"prefixed": "v2"}`)
	doc = doc.Update().SetMetadata(v2).SaveX(ctx)
	require.Equal(t, v2, doc.Metadata)
}

func TestBlobLoadOnScanWithEncryption(t *testing.T) {
	dir := blobDir(t)
	masterSeed := []byte("encryption-seed-for-loadonscan")

	encryptedOpeners := ent.BlobOpeners{
		Document: func(ctx context.Context, field string) (ent.Blob, error) {
			var subdir string
			switch field {
			case document.FieldContent:
				subdir = "documents"
			case document.FieldThumbnail:
				subdir = "thumbnails"
			case document.FieldAttachment:
				subdir = "attachments"
			case document.FieldMetadata:
				subdir = "metadata"
			default:
				return nil, fmt.Errorf("unknown blob field: %s", field)
			}
			b, err := blob.OpenBucket(ctx, "file://"+filepath.Join(dir, subdir))
			if err != nil {
				return nil, err
			}
			return blob.NewEncrypted(b, masterSeed), nil
		},
	}
	client, _, _ := setupBlob(t, ent.WithBlobOpeners(encryptedOpeners))

	acmeCtx := blob.WithTenant(context.Background(), "acme")
	meta := []byte(`{"encrypted": true, "tenant": "acme"}`)
	doc := client.Document.Create().
		SetName("enc-los-doc").
		SetContent(bytes.NewReader([]byte("content"))).
		SetThumbnail(bytes.NewReader([]byte("thumb"))).
		SetAttachment([]byte("att")).
		SetMetadata(meta).
		SaveX(acmeCtx)

	// Struct field populated on create.
	require.Equal(t, meta, doc.Metadata)

	// Query auto-loads encrypted metadata (same tenant).
	queried := client.Document.GetX(acmeCtx, doc.ID)
	require.Equal(t, meta, queried.Metadata)

	// Update round-trips through encryption.
	v2 := []byte(`{"encrypted": true, "v": 2}`)
	doc = doc.Update().SetMetadata(v2).SaveX(acmeCtx)
	require.Equal(t, v2, doc.Metadata)
}

func TestBlobDeleteRemovesBlobs(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create a document with all blob fields populated.
	doc := client.Document.Create().
		SetName("delete-doc").
		SetContent(bytes.NewReader([]byte("content-data"))).
		SetThumbnail(bytes.NewReader([]byte("thumb-data"))).
		SetAttachment([]byte("att-data")).
		SetMetadata([]byte("meta-data")).
		SaveX(ctx)

	// Verify blobs are readable before delete.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("content-data"), got)
	got = blobContent(t, doc.ThumbnailReader, ctx)
	require.Equal(t, []byte("thumb-data"), got)
	got = blobContent(t, doc.AttachmentReader, ctx)
	require.Equal(t, []byte("att-data"), got)
	got = blobContent(t, doc.MetadataReader, ctx)
	require.Equal(t, []byte("meta-data"), got)

	// Delete the entity.
	client.Document.DeleteOne(doc).ExecX(ctx)

	// Verify entity is gone from database.
	count := client.Document.Query().CountX(ctx)
	require.Equal(t, 0, count)

	// Verify blobs were removed: readers return nil (not found).
	got = blobContent(t, doc.ContentReader, ctx)
	require.Nil(t, got, "content blob should be deleted")
	got = blobContent(t, doc.ThumbnailReader, ctx)
	require.Nil(t, got, "thumbnail blob should be deleted")
	got = blobContent(t, doc.AttachmentReader, ctx)
	require.Nil(t, got, "attachment blob should be deleted")
	got = blobContent(t, doc.MetadataReader, ctx)
	require.Nil(t, got, "metadata blob should be deleted")
}

func TestBlobDeleteBulkRemovesBlobs(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create multiple documents, save references.
	var docs []*ent.Document
	for i := range 3 {
		doc := client.Document.Create().
			SetName(fmt.Sprintf("bulk-del-%d", i)).
			SetContent(bytes.NewReader([]byte(fmt.Sprintf("c%d", i)))).
			SetThumbnail(bytes.NewReader([]byte(fmt.Sprintf("t%d", i)))).
			SetAttachment([]byte("att")).
			SaveX(ctx)
		docs = append(docs, doc)
	}

	// Verify blobs are readable.
	for i, doc := range docs {
		got := blobContent(t, doc.ContentReader, ctx)
		require.Equal(t, []byte(fmt.Sprintf("c%d", i)), got)
	}

	// Bulk delete all documents.
	n := client.Document.Delete().ExecX(ctx)
	require.Equal(t, 3, n)

	// Verify blobs were removed: readers return nil.
	for _, doc := range docs {
		got := blobContent(t, doc.ContentReader, ctx)
		require.Nil(t, got, "blob should be deleted after bulk delete")
	}
}

func TestBlobDeleteWithPredicateOnlyDeletesMatching(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create two documents.
	keep := client.Document.Create().
		SetName("keep-me").
		SetContent(bytes.NewReader([]byte("keep-content"))).
		SetThumbnail(bytes.NewReader([]byte("keep-thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)
	del := client.Document.Create().
		SetName("delete-me").
		SetContent(bytes.NewReader([]byte("delete-content"))).
		SetThumbnail(bytes.NewReader([]byte("delete-thumb"))).
		SetAttachment([]byte("att")).
		SaveX(ctx)

	// Delete only the second one.
	n := client.Document.Delete().Where(document.Name("delete-me")).ExecX(ctx)
	require.Equal(t, 1, n)

	// One document remains.
	remaining := client.Document.Query().OnlyX(ctx)
	require.Equal(t, "keep-me", remaining.Name)

	// The kept document's blobs are still readable.
	got := blobContent(t, keep.ContentReader, ctx)
	require.Equal(t, []byte("keep-content"), got)

	// The deleted document's blobs are gone.
	got = blobContent(t, del.ContentReader, ctx)
	require.Nil(t, got, "deleted doc's content blob should be removed")
	got = blobContent(t, del.ThumbnailReader, ctx)
	require.Nil(t, got, "deleted doc's thumbnail blob should be removed")
}

func TestBlobDeleteTxCommitRemovesBlobs(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create a document with blobs.
	doc := client.Document.Create().
		SetName("tx-commit-doc").
		SetContent(bytes.NewReader([]byte("tx-content"))).
		SetThumbnail(bytes.NewReader([]byte("tx-thumb"))).
		SetAttachment([]byte("tx-att")).
		SaveX(ctx)

	// Verify blobs exist.
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("tx-content"), got)

	// Delete inside a transaction and commit.
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	err = tx.Document.DeleteOne(doc).Exec(ctx)
	require.NoError(t, err)

	// Before commit, blobs should still exist (cleanup is deferred).
	got = blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("tx-content"), got, "blob should still exist before commit")
	got = blobContent(t, doc.ThumbnailReader, ctx)
	require.Equal(t, []byte("tx-thumb"), got, "blob should still exist before commit")

	// Commit the transaction.
	require.NoError(t, tx.Commit())

	// After commit, blobs should be deleted.
	got = blobContent(t, doc.ContentReader, ctx)
	require.Nil(t, got, "content blob should be deleted after tx commit")
	got = blobContent(t, doc.ThumbnailReader, ctx)
	require.Nil(t, got, "thumbnail blob should be deleted after tx commit")
	got = blobContent(t, doc.AttachmentReader, ctx)
	require.Nil(t, got, "attachment blob should be deleted after tx commit")
}

func TestBlobDeleteTxRollbackPreservesBlobs(t *testing.T) {
	client, ctx, _ := setupBlob(t)

	// Create a document with blobs.
	doc := client.Document.Create().
		SetName("tx-rollback-doc").
		SetContent(bytes.NewReader([]byte("keep-content"))).
		SetThumbnail(bytes.NewReader([]byte("keep-thumb"))).
		SetAttachment([]byte("keep-att")).
		SaveX(ctx)

	// Delete inside a transaction but rollback.
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	err = tx.Document.DeleteOne(doc).Exec(ctx)
	require.NoError(t, err)

	// Rollback the transaction.
	require.NoError(t, tx.Rollback())

	// After rollback, the entity and its blobs should still exist.
	exists := client.Document.Query().Where(document.ID(doc.ID)).ExistX(ctx)
	require.True(t, exists, "document should still exist after rollback")
	got := blobContent(t, doc.ContentReader, ctx)
	require.Equal(t, []byte("keep-content"), got, "content blob should be preserved after rollback")
	got = blobContent(t, doc.ThumbnailReader, ctx)
	require.Equal(t, []byte("keep-thumb"), got, "thumbnail blob should be preserved after rollback")
	got = blobContent(t, doc.AttachmentReader, ctx)
	require.Equal(t, []byte("keep-att"), got, "attachment blob should be preserved after rollback")
}
