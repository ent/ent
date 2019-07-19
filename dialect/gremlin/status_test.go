package gremlin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusText(t *testing.T) {
	assert.NotEmpty(t, StatusText(StatusSuccess))
	assert.Empty(t, StatusText(4242))
}
