package index_test

import (
	"testing"

	"fbc/ent/schema/index"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	idx := index.Fields("name", "address")
	require.Empty(t, idx.Edge())
	require.False(t, idx.IsUnique())
	require.Equal(t, []string{"name", "address"}, idx.Fields())

	idx = index.Fields("name", "address").Unique()
	require.Empty(t, idx.Edge())
	require.True(t, idx.IsUnique())
	require.Equal(t, []string{"name", "address"}, idx.Fields())

	idx = index.Fields("name", "address").FromEdge("parent").Unique()
	require.Equal(t, "parent", idx.Edge())
	require.True(t, idx.IsUnique())
	require.Equal(t, []string{"name", "address"}, idx.Fields())
}
