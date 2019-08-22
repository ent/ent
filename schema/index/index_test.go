package index_test

import (
	"testing"

	"github.com/facebookincubator/ent/schema/index"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	idx := index.Fields("name", "address")
	require.Empty(t, idx.Edges())
	require.False(t, idx.IsUnique())
	require.Equal(t, []string{"name", "address"}, idx.Fields())

	idx = index.Fields("name", "address").Unique()
	require.Empty(t, idx.Edges())
	require.True(t, idx.IsUnique())
	require.Equal(t, []string{"name", "address"}, idx.Fields())

	idx = index.Fields("name", "address").FromEdges("parent", "type").Unique()
	require.Equal(t, []string{"parent", "type"}, idx.Edges())
	require.True(t, idx.IsUnique())
	require.Equal(t, []string{"name", "address"}, idx.Fields())
}
