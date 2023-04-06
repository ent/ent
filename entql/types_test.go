// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package entql

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFielder(t *testing.T) {
	tests := []struct {
		input    Fielder
		expected string
	}{
		{
			input:    StringEQ("value"),
			expected: `field == "value"`,
		},
		{
			input: StringOr(
				StringEQ("a"),
				StringEQ("b"),
				StringEQ("c"),
			),
			expected: `(field == "a" || field == "b" || field == "c")`,
		},
		{
			input: StringAnd(
				StringEQ("a"),
				StringNot(
					StringOr(
						StringEQ("b"),
						StringGT("c"),
						StringNEQ("d"),
					),
				),
			),
			expected: `field == "a" && !((field == "b" || field > "c" || field != "d"))`,
		},
		{
			input:    IntGT(1),
			expected: `field > 1`,
		},
		{
			input:    IntGTE(1),
			expected: `field >= 1`,
		},
		{
			input:    IntLT(1),
			expected: `field < 1`,
		},
		{
			input:    IntLTE(1),
			expected: `field <= 1`,
		},
		{
			input:    IntGT(1),
			expected: `field > 1`,
		},
		{
			input:    IntNot(IntGTE(1)),
			expected: `!(field >= 1)`,
		},
		{
			input: BoolNot(
				BoolOr(
					BoolEQ(true),
					BoolEQ(false),
				),
			),
			expected: `!(field == true || field == false)`,
		},
	}
	for i := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := tests[i].input.Field("field")
			assert.Equal(t, tests[i].expected, p.String())
		})
	}
}
