// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRawMessageEncoding(t *testing.T) {
	var s struct{ M RawMessage }
	got, err := MarshalToString(s)
	require.NoError(t, err)
	assert.Equal(t, `{"M":null}`, got)

	s.M = []byte(`"155a"`)
	got, err = MarshalToString(s)
	require.NoError(t, err)
	assert.JSONEq(t, `{"M": "155a"}`, got)

	err = (*RawMessage)(nil).UnmarshalGraphson(s.M)
	assert.Error(t, err)

	s.M = nil
	err = UnmarshalFromString(got, &s)
	require.NoError(t, err)
	assert.Equal(t, `"155a"`, string(s.M))
}
