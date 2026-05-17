// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ent

import (
	"context"
	"errors"
	"io"
	"io/fs"
)

// Blob defines the interface for blob storage operations.
// Implementations should return [io/fs.ErrNotExist] (or an error wrapping it)
// from NewReader when the requested key does not exist.
//
// Blob data is written to external storage before the database row is created.
// If the row insertion fails (e.g. constraint violation, connectivity issue),
// the already-written blob is not automatically cleaned up. This is by design—
// users should implement their own garbage collection for orphaned blobs if needed.
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

// blobReadCloser wraps an io.ReadCloser to also close the bucket on Close.
type blobReadCloser struct {
	io.ReadCloser
	bucket Blob
}

func (r *blobReadCloser) Close() error {
	return errors.Join(r.ReadCloser.Close(), r.bucket.Close())
}

// BlobWriter returns a writer for the given key in the blob bucket.
// The returned writer closes both the underlying writer and the bucket.
func BlobWriter(ctx context.Context, b Blob, key string) (io.WriteCloser, error) {
	w, err := b.NewWriter(ctx, key)
	if err != nil {
		return nil, errors.Join(err, b.Close())
	}
	return &blobWriteCloser{WriteCloser: w, bucket: b}, nil
}

// blobWriteCloser wraps an io.WriteCloser to also close the bucket on Close.
type blobWriteCloser struct {
	io.WriteCloser
	bucket Blob
}

func (w *blobWriteCloser) Close() error {
	return errors.Join(w.WriteCloser.Close(), w.bucket.Close())
}

// BlobBulkWriter manages blob bucket lifecycles for write operations.
// It lazily opens buckets per field and provides a Close method
// to release all resources when done.
type BlobBulkWriter struct {
	opener  func(context.Context, string) (Blob, error)
	buckets map[string]Blob
}

// NewBlobBulkWriter creates a writer that uses opener to lazily open buckets.
func NewBlobBulkWriter(opener func(context.Context, string) (Blob, error)) *BlobBulkWriter {
	return &BlobBulkWriter{
		buckets: make(map[string]Blob),
		opener:  opener,
	}
}

// Close closes all open buckets.
func (w *BlobBulkWriter) Close() error {
	var errs []error
	for _, b := range w.buckets {
		errs = append(errs, b.Close())
	}
	return errors.Join(errs...)
}

// Write writes r to the blob at key for the given field. The bucket is opened
// lazily on first use and reused for subsequent writes to the same field.
func (w *BlobBulkWriter) Write(ctx context.Context, field, key string, r io.Reader) error {
	b, err := w.bucket(ctx, field)
	if err != nil {
		return err
	}
	wr, err := b.NewWriter(ctx, key)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wr, r); err != nil {
		return errors.Join(err, wr.Close())
	}
	return wr.Close()
}

// Delete removes the blob at key for the given field. The bucket is opened
// lazily on first use and reused for subsequent deletes to the same field.
func (w *BlobBulkWriter) Delete(ctx context.Context, field, key string) error {
	b, err := w.bucket(ctx, field)
	if err != nil {
		return err
	}
	return b.Delete(ctx, key)
}

func (w *BlobBulkWriter) bucket(ctx context.Context, field string) (Blob, error) {
	if b, ok := w.buckets[field]; ok {
		return b, nil
	}
	b, err := w.opener(ctx, field)
	if err != nil {
		return nil, err
	}
	w.buckets[field] = b
	return b, nil
}

// BlobBulkReader manages blob bucket lifecycles for read operations.
// It lazily opens buckets per field and provides a Close method
// to release all resources when done.
type BlobBulkReader struct {
	opener  func(context.Context, string) (Blob, error)
	buckets map[string]Blob
}

// NewBlobBulkReader creates a reader that uses opener to lazily open buckets.
func NewBlobBulkReader(opener func(context.Context, string) (Blob, error)) *BlobBulkReader {
	return &BlobBulkReader{
		buckets: make(map[string]Blob),
		opener:  opener,
	}
}

// Close closes all open buckets.
func (r *BlobBulkReader) Close() error {
	var errs []error
	for _, b := range r.buckets {
		errs = append(errs, b.Close())
	}
	return errors.Join(errs...)
}

// Read reads the blob at key for the given field. The bucket is opened
// lazily on first use and reused for subsequent reads to the same field.
// Returns nil, nil if the blob does not exist (fs.ErrNotExist).
func (r *BlobBulkReader) Read(ctx context.Context, field, key string) ([]byte, error) {
	b, err := r.bucket(ctx, field)
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

func (r *BlobBulkReader) bucket(ctx context.Context, field string) (Blob, error) {
	if b, ok := r.buckets[field]; ok {
		return b, nil
	}
	b, err := r.opener(ctx, field)
	if err != nil {
		return nil, err
	}
	r.buckets[field] = b
	return b, nil
}
