// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package blob provides a gocloud.dev/blob adapter for ent's generated Blob
// interface. It is used by the integration tests and serves as a reference
// implementation for wiring blob storage into ent-generated code.
package blob

import (
	"context"
	"io"
	"io/fs"

	goblob "gocloud.dev/blob"
	"gocloud.dev/gcerrors"
)

// GoBucket wraps a gocloud.dev/blob.Bucket and implements the generated Blob interface.
type GoBucket struct {
	b          *goblob.Bucket
	readerOpts *goblob.ReaderOptions
	writerOpts *goblob.WriterOptions
}

// OpenBucket opens a gocloud.dev/blob bucket by URL and returns it as a [*GoBucket].
func OpenBucket(ctx context.Context, url string) (*GoBucket, error) {
	b, err := goblob.OpenBucket(ctx, url)
	if err != nil {
		return nil, err
	}
	return &GoBucket{b: b}, nil
}

// Prefixed returns a new GoBucket scoped to the given key prefix.
// The original bucket is consumed and must not be used after this call.
func (b *GoBucket) Prefixed(prefix string) *GoBucket {
	return &GoBucket{
		b:          goblob.PrefixedBucket(b.b, prefix),
		readerOpts: b.readerOpts,
		writerOpts: b.writerOpts,
	}
}

// WithReaderOptions returns a new GoBucket that uses the given reader options.
func (b *GoBucket) WithReaderOptions(opts *goblob.ReaderOptions) *GoBucket {
	return &GoBucket{b: b.b, writerOpts: b.writerOpts, readerOpts: opts}
}

// WithWriterOptions returns a new GoBucket that uses the given writer options.
func (b *GoBucket) WithWriterOptions(opts *goblob.WriterOptions) *GoBucket {
	return &GoBucket{b: b.b, writerOpts: opts, readerOpts: b.readerOpts}
}

// NewReader opens a reader for the blob stored at key.
// It returns [fs.ErrNotExist] if the key does not exist.
func (b *GoBucket) NewReader(ctx context.Context, key string) (io.ReadCloser, error) {
	switch r, err := b.b.NewReader(ctx, key, b.readerOpts); {
	case gcerrors.Code(err) == gcerrors.NotFound:
		return nil, fs.ErrNotExist
	case err != nil:
		return nil, err
	default:
		return r, nil
	}
}

// NewWriter opens a writer for the blob stored at key.
func (b *GoBucket) NewWriter(ctx context.Context, key string) (io.WriteCloser, error) {
	return b.b.NewWriter(ctx, key, b.writerOpts)
}

// Delete removes the blob at key. Returns nil if the key does not exist.
func (b *GoBucket) Delete(ctx context.Context, key string) error {
	err := b.b.Delete(ctx, key)
	if gcerrors.Code(err) == gcerrors.NotFound {
		return nil
	}
	return err
}

// Close releases the underlying gocloud bucket resources.
func (b *GoBucket) Close() error {
	return b.b.Close()
}
