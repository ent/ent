package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMime(t *testing.T) {
	str := "application/vnd.gremlin-v2.0+json"
	mime := NewMime(str)
	require.Len(t, mime, len(str)+1)
	assert.EqualValues(t, len(str), mime[0])
	assert.EqualValues(t, str, mime[1:])
	assert.Equal(t, str, mime.String())
}
