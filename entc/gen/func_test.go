// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHasImportFor(t *testing.T) {
	initTemplates()
	sqlStorage, err := NewStorage("sql")
	require.NoError(t, err)
	gremlinStorage, err := NewStorage("gremlin")
	require.NoError(t, err)

	// Imported by the storage itself.
	require.True(t, hasImportFor(sqlStorage, "sql"))
	require.True(t, hasImportFor(gremlinStorage, "dsl"))
	// Imported only by the other storage, so the name is free.
	require.False(t, hasImportFor(gremlinStorage, "sql"))
	require.False(t, hasImportFor(gremlinStorage, "sqlgraph"))
	require.False(t, hasImportFor(sqlStorage, "dsl"))
	// Imported by the shared import template under every storage.
	require.True(t, hasImportFor(gremlinStorage, "field"))
	require.True(t, hasImportFor(sqlStorage, "context"))
	// Unknown names and a nil storage behave like hasImport.
	require.False(t, hasImportFor(sqlStorage, "unknown"))
	require.Equal(t, hasImport("sql"), hasImportFor(nil, "sql"))
}
