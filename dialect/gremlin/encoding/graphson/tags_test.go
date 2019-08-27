// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTag(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		typ  string
		opts tagOptions
	}{
		{
			name: "Empty",
		},
		{
			name: "TypeOnly",
			tag:  "g:Int32",
			typ:  "g:Int32",
		},
		{
			name: "OptsOnly",
			tag:  ",opt1,opt2",
			opts: "opt1,opt2",
		},
		{
			name: "TypeAndOpts",
			tag:  "g:UUID,opt3,opt4",
			typ:  "g:UUID",
			opts: "opt3,opt4",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			typ, opts := parseTag(tc.tag)
			assert.Equal(t, tc.typ, typ)
			assert.Equal(t, tc.opts, opts)
		})
	}
}

func TestTagOptionsContains(t *testing.T) {
	_, opts := parseTag(",opt1,opt2,opt3")
	assert.True(t, opts.Contains("opt1"))
	assert.True(t, opts.Contains("opt2"))
	assert.True(t, opts.Contains("opt3"))
	assert.False(t, opts.Contains("opt4"))
	assert.False(t, opts.Contains("opt11"))
	assert.False(t, opts.Contains(""))
}
