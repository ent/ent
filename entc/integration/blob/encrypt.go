// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package blob

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

// tenantKey is the context key for the tenant name.
type tenantKey struct{}

// WithTenant returns a context carrying the given tenant name.
func WithTenant(ctx context.Context, tenant string) context.Context {
	return context.WithValue(ctx, tenantKey{}, tenant)
}

// TenantFrom extracts the tenant name from the context.
// Returns "" if no tenant is set.
func TenantFrom(ctx context.Context) string {
	v, _ := ctx.Value(tenantKey{}).(string)
	return v
}

// Encrypted wraps a GoBucket and transparently encrypts/decrypts blob data
// using AES-CTR with a per-tenant derived key.
//
// The encryption key is derived from the tenant name (from context) combined
// with a master seed using SHA-256. This ensures that data written by one
// tenant cannot be decrypted by another tenant.
//
// A random IV is prepended to the ciphertext on write and read back on open.
//
// Usage:
//
//	bucket, _ := blob.OpenBucket(ctx, url)
//	enc := blob.NewEncrypted(bucket, masterSeed)
//	// Use enc as the Blob implementation in BlobOpeners.
//	// Ensure the context carries the tenant: blob.WithTenant(ctx, "acme")
type Encrypted struct {
	inner *GoBucket
	seed  []byte
}

// NewEncrypted creates an encrypting wrapper around the given bucket.
// The seed is combined with the tenant name (from context) at each operation
// to derive a per-tenant AES-256 key via SHA-256.
func NewEncrypted(bucket *GoBucket, seed []byte) *Encrypted {
	return &Encrypted{inner: bucket, seed: seed}
}

// deriveKey produces a 32-byte AES-256 key from the tenant + seed.
func (e *Encrypted) deriveKey(tenant string) (cipher.Block, error) {
	if tenant == "" {
		return nil, errors.New("blob: encryption requires a tenant in context (use blob.WithTenant)")
	}
	h := sha256.New()
	h.Write([]byte(tenant))
	h.Write(e.seed)
	b, err := aes.NewCipher(h.Sum(nil)) // 32 bytes → AES-256
	if err != nil {
		return nil, fmt.Errorf("blob: deriving encryption key: %w", err)
	}
	return b, nil
}

// NewReader opens and decrypts the blob at key. The first [aes.BlockSize] bytes
// are treated as the IV; the rest is decrypted with AES-CTR using the
// tenant-derived key.
// Returns [fs.ErrNotExist] if the key does not exist.
func (e *Encrypted) NewReader(ctx context.Context, key string) (io.ReadCloser, error) {
	block, err := e.deriveKey(TenantFrom(ctx))
	if err != nil {
		return nil, err
	}
	rc, err := e.inner.NewReader(ctx, key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	switch _, err := io.ReadFull(rc, iv); {
	case err == io.EOF, err == io.ErrUnexpectedEOF:
		return nil, errors.Join(fmt.Errorf("blob: ciphertext too short for key %q", key), rc.Close())
	case err != nil:
		return nil, errors.Join(err, rc.Close())
	}
	return &decryptionReader{
		Reader: cipher.StreamReader{S: cipher.NewCTR(block, iv), R: rc},
		closer: rc,
	}, nil
}

// NewWriter opens an encrypting writer for the blob at key. A random IV is
// written first, followed by AES-CTR encrypted data using the tenant-derived key.
func (e *Encrypted) NewWriter(ctx context.Context, key string) (io.WriteCloser, error) {
	block, err := e.deriveKey(TenantFrom(ctx))
	if err != nil {
		return nil, err
	}
	wc, err := e.inner.NewWriter(ctx, key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, errors.Join(fmt.Errorf("blob: generating IV: %w", err), wc.Close())
	}
	if _, err := wc.Write(iv); err != nil {
		return nil, errors.Join(fmt.Errorf("blob: writing IV: %w", err), wc.Close())
	}
	return cipher.StreamWriter{S: cipher.NewCTR(block, iv), W: wc}, nil
}

// Close releases the underlying bucket resources.
func (e *Encrypted) Close() error {
	return e.inner.Close()
}

// Delete removes the blob at key. Encryption is irrelevant for deletion.
func (e *Encrypted) Delete(ctx context.Context, key string) error {
	return e.inner.Delete(ctx, key)
}

// decryptionReader decrypts on Read and closes the underlying storage decryptionReader.
type decryptionReader struct {
	io.Reader
	closer io.Closer
}

func (r *decryptionReader) Close() error {
	return r.closer.Close()
}
