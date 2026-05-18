// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
)

// Blob defines the interface for blob storage operations.
// Implementations should return [io/fs.ErrNotExist] (or an error wrapping it)
// from NewReader when the requested key does not exist.
//
// Single-row SQL create builders write blob data to external storage before
// inserting the database row. If the row insertion fails (for example, due to
// a constraint violation), generated code attempts to delete the just-written blobs.
type Blob interface {
	// NewReader opens a reader for the given key.
	NewReader(ctx context.Context, key string) (io.ReadCloser, error)
	// NewWriter opens a writer for the given key.
	NewWriter(ctx context.Context, key string) (io.WriteCloser, error)
	// Delete removes the blob at the given key.
	// Implementations should return nil (not an error) if the key does not exist.
	Delete(ctx context.Context, key string) error
	// Close releases any resources held by the bucket.
	Close() error
}

// BlobOpener is a function that opens a [Blob] bucket for the given field name.
type BlobOpener func(context.Context, string) (Blob, error)

// BlobKey identifies a blob in storage by field name and key.
type BlobKey struct {
	Field string
	Key   string
}

// BlobQuerier queries existing blob keys from the database.
// [Blobs.Update] passes the mutated field names; [Blobs.Delete] passes nil
// to indicate all fields should be queried.
type BlobQuerier interface {
	QueryBlobKeys(ctx context.Context, fields []string) ([]BlobKey, error)
}

// BlobUpdateResult holds post-update blob operations.
type BlobUpdateResult struct {
	Rollback BlobOp // Deletes newly-written blobs. Call on SQL failure.
	Commit   BlobOp // Deletes old replaced blobs. Call after successful SQL commit.
}

// BlobOp is a deferred blob storage operation (e.g. rollback or commit).
type BlobOp func(context.Context) error

// BlobKeyFunc generates a storage key for a blob from its content.
type BlobKeyFunc func(context.Context, []byte) (string, error)

// Blobs orchestrates blob storage operations for a single mutation.
// Use [NewBlobs] to create, then call [Blobs.Set] or [Blobs.SetCleared]
// for each blob field, then [Blobs.Create] or [Blobs.Update].
type Blobs struct {
	opener BlobOpener
	inputs []blobInput
}

type blobInput struct {
	field   string
	data    []byte
	newKey  BlobKeyFunc
	apply   func(string)
	cleared bool
	clear   func()
}

// NewBlobs creates a blob orchestrator for the given opener.
func NewBlobs(opener BlobOpener) *Blobs {
	return &Blobs{opener: opener}
}

// Set adds a blob field to be written. The apply callback is called with
// the generated key to set it on the SQL spec and node.
func (b *Blobs) Set(f string, data []byte, key BlobKeyFunc, apply func(string)) {
	b.inputs = append(b.inputs, blobInput{field: f, data: data, newKey: key, apply: apply})
}

// SetCleared marks a blob field as cleared. The clear callback should
// remove the key column from the SQL spec.
func (b *Blobs) SetCleared(f string, clear func()) {
	b.inputs = append(b.inputs, blobInput{field: f, cleared: true, clear: clear})
}

// Create prepares inputs, writes blobs, and returns a rollback [BlobOp].
func (b *Blobs) Create(ctx context.Context) (BlobOp, error) {
	writes, err := b.prepare(ctx)
	if err != nil {
		return nil, err
	}
	return b.write(ctx, writes)
}

// Update prepares inputs, queries old keys, writes new blobs, and returns
// a [BlobUpdateResult] for post-SQL handling.
func (b *Blobs) Update(ctx context.Context, q BlobQuerier) (*BlobUpdateResult, error) {
	if len(b.inputs) == 0 {
		return noopBlobResult, nil
	}
	writes, err := b.prepare(ctx)
	if err != nil {
		return nil, err
	}
	mutated := make([]string, len(b.inputs))
	cleared := make(map[string]bool)
	for i := range b.inputs {
		mutated[i] = b.inputs[i].field
		if b.inputs[i].cleared {
			cleared[b.inputs[i].field] = true
		}
	}
	keys, err := q.QueryBlobKeys(ctx, mutated)
	if err != nil {
		return nil, fmt.Errorf("querying old blob keys: %w", err)
	}
	// Build a set of old keys per field to detect unchanged blobs.
	oldKeys := make(map[string]string, len(keys))
	for _, k := range keys {
		oldKeys[k.Field] = k.Key
	}
	// Filter out writes where the key is unchanged (same content).
	filtered := writes[:0]
	for _, wr := range writes {
		if oldKeys[wr.Field] == wr.Key {
			continue
		}
		filtered = append(filtered, wr)
	}
	rollback, err := b.write(ctx, filtered)
	if err != nil {
		return nil, err
	}
	// Collect orphaned blobs: old keys for fields that changed or were cleared.
	var orphaned []BlobKey
	for _, k := range keys {
		if cleared[k.Field] {
			orphaned = append(orphaned, k)
			continue
		}
		for _, wr := range writes {
			if wr.Field == k.Field && wr.Key != k.Key {
				orphaned = append(orphaned, k)
				break
			}
		}
	}
	return &BlobUpdateResult{
		Rollback: rollback,
		Commit:   b.deleteOp(orphaned),
	}, nil
}

// Delete queries existing blob keys and returns a [BlobOp] that removes
// them from storage. Use for delete mutations.
func (b *Blobs) Delete(ctx context.Context, q BlobQuerier) (BlobOp, error) {
	keys, err := q.QueryBlobKeys(ctx, nil)
	if err != nil {
		return nil, err
	}
	return b.deleteOp(keys), nil
}

type blobWrite struct {
	BlobKey
	data []byte
}

func (b *Blobs) prepare(ctx context.Context) ([]blobWrite, error) {
	var writes []blobWrite
	for _, inp := range b.inputs {
		if inp.cleared {
			inp.clear()
			continue
		}
		k, err := inp.newKey(ctx, inp.data)
		if err != nil {
			return nil, fmt.Errorf("generating blob key for %s: %w", inp.field, err)
		}
		if inp.apply != nil {
			inp.apply(k)
		}
		writes = append(writes, blobWrite{
			BlobKey: BlobKey{Field: inp.field, Key: k},
			data:    inp.data,
		})
	}
	return writes, nil
}

func (b *Blobs) write(ctx context.Context, writes []blobWrite) (BlobOp, error) {
	if len(writes) == 0 {
		return noOp, nil
	}
	w := NewBlobStore(b.opener)
	var written []BlobKey
	for _, wr := range writes {
		if err := w.write(ctx, wr.Field, wr.Key, wr.data); err != nil {
			var errs []error
			errs = append(errs, fmt.Errorf("writing blob for %s: %w", wr.Field, err))
			for _, k := range written {
				if derr := w.delete(ctx, k.Field, k.Key); derr != nil {
					errs = append(errs, derr)
				}
			}
			errs = append(errs, w.Close())
			return nil, errors.Join(errs...)
		}
		written = append(written, wr.BlobKey)
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.deleteOp(written), nil
}

func (b *Blobs) deleteOp(keys []BlobKey) BlobOp {
	if len(keys) == 0 {
		return noOp
	}
	return func(ctx context.Context) error {
		s := NewBlobStore(b.opener)
		var errs []error
		for _, k := range keys {
			if err := s.delete(ctx, k.Field, k.Key); err != nil {
				errs = append(errs, err)
			}
		}
		if err := s.Close(); err != nil {
			errs = append(errs, err)
		}
		return errors.Join(errs...)
	}
}

var (
	noOp           = func(context.Context) error { return nil }
	noopBlobResult = &BlobUpdateResult{Rollback: noOp, Commit: noOp}
)

// BlobReader returns a reader for the given key from the blob bucket.
// The returned reader closes both the underlying reader and the bucket.
// Returns nil, nil if the blob does not exist (fs.ErrNotExist).
func BlobReader(ctx context.Context, b Blob, key string) (io.ReadCloser, error) {
	switch r, err := b.NewReader(ctx, key); {
	case errors.Is(err, fs.ErrNotExist):
		return nil, b.Close()
	case err != nil:
		return nil, errors.Join(err, b.Close())
	default:
		return &blobReadCloser{ReadCloser: r, bucket: b}, nil
	}
}

type blobReadCloser struct {
	io.ReadCloser
	bucket Blob
}

func (r *blobReadCloser) Close() error {
	return errors.Join(r.ReadCloser.Close(), r.bucket.Close())
}

// BlobStore manages blob bucket lifecycles for read, write, and delete operations.
// It lazily opens buckets per field and reuses them for subsequent operations.
type BlobStore struct {
	opener  BlobOpener
	buckets map[string]Blob
}

// NewBlobStore creates a store that uses opener to lazily open buckets.
func NewBlobStore(opener BlobOpener) *BlobStore {
	return &BlobStore{buckets: make(map[string]Blob), opener: opener}
}

// Close closes all open buckets.
func (s *BlobStore) Close() error {
	var errs []error
	for _, b := range s.buckets {
		errs = append(errs, b.Close())
	}
	return errors.Join(errs...)
}

// write writes data to the blob at key for the given field.
func (s *BlobStore) write(ctx context.Context, field, key string, data []byte) error {
	b, err := s.bucket(ctx, field)
	if err != nil {
		return err
	}
	wr, err := b.NewWriter(ctx, key)
	if err != nil {
		return err
	}
	if _, err := wr.Write(data); err != nil {
		return errors.Join(err, wr.Close(), b.Delete(ctx, key))
	}
	return wr.Close()
}

// delete removes the blob at key for the given field.
func (s *BlobStore) delete(ctx context.Context, field, key string) error {
	b, err := s.bucket(ctx, field)
	if err != nil {
		return err
	}
	return b.Delete(ctx, key)
}

// Read reads the blob at key for the given field.
// Returns nil, nil if the blob does not exist (fs.ErrNotExist).
func (s *BlobStore) Read(ctx context.Context, field, key string) ([]byte, error) {
	b, err := s.bucket(ctx, field)
	if err != nil {
		return nil, err
	}
	rc, err := b.NewReader(ctx, key)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(rc)
	if closeErr := rc.Close(); closeErr != nil && err == nil {
		err = closeErr
	}
	return data, err
}

func (s *BlobStore) bucket(ctx context.Context, field string) (Blob, error) {
	if b, ok := s.buckets[field]; ok {
		return b, nil
	}
	if s.opener == nil {
		return nil, errors.New("ent: blob storage not configured (missing WithBlobOpeners)")
	}
	b, err := s.opener(ctx, field)
	if err != nil {
		return nil, err
	}
	s.buckets[field] = b
	return b, nil
}
