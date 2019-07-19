package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeUnsupportedType(t *testing.T) {
	_, err := Marshal(func() {})
	assert.Error(t, err)
}
