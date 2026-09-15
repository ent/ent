// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	ChineseString           = "金额"
	ChineseStringSnakeCased = "金额"

	DanishString           = "MÆNGDEPenge"
	DanishStringSnakeCased = "mængde_penge"

	EnglishString           = "TotalAmount"
	EnglishStringSnakeCased = "total_amount"

	GreekString           = "ΧρηματικόΠοσό"
	GreekStringSnakeCased = "χρηματικό_ποσό"
)

func TestSnakeCasing(t *testing.T) {
	require := require.New(t)

	require.Equal(ChineseStringSnakeCased, snake(ChineseString))
	require.Equal(DanishStringSnakeCased, snake(DanishString))
	require.Equal(EnglishStringSnakeCased, snake(EnglishString))
	require.Equal(GreekStringSnakeCased, snake(GreekString))
}
