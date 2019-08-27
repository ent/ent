// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
