// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"database/sql"
	"testing"
	"uuid"

	"entgo.io/ent/schema/field"

	googleuuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// Both the third-party ValueScanner type and the native uuid.UUID are
// accepted, distinguished by their package path. Pointer variants of
// both are supported as well.
func TestField_UUIDTypes(t *testing.T) {
	for _, tt := range []struct {
		pkgPath string
		typ     any
	}{
		{"github.com/google/uuid", googleuuid.UUID{}},
		{"github.com/google/uuid", &googleuuid.UUID{}},
		{"uuid", uuid.UUID{}},
		{"uuid", &uuid.UUID{}},
	} {
		desc := field.UUID("id", tt.typ).Descriptor()
		require.NoError(t, desc.Err)
		require.Equal(t, tt.pkgPath, desc.Info.PkgPath)
	}
}

func TestNativeUUIDNull(t *testing.T) {
	want := uuid.UUID{1}
	var got sql.Null[uuid.UUID]
	err := got.Scan(want.String())
	require.NoError(t, err)
	require.True(t, got.Valid)
	require.Equal(t, want, got.V)
	err = got.Scan(nil)
	require.NoError(t, err)
	require.False(t, got.Valid)
}
