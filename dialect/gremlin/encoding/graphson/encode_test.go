// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeUnsupportedType(t *testing.T) {
	_, err := Marshal(func() {})
	assert.Error(t, err)
}
