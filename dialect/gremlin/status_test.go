// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusText(t *testing.T) {
	assert.NotEmpty(t, StatusText(StatusSuccess))
	assert.Empty(t, StatusText(4242))
}
