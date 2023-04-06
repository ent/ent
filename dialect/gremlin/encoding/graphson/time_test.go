// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeEncoding(t *testing.T) {
	const ms = 1481750076295
	ts := time.Unix(0, ms*time.Millisecond.Nanoseconds())

	for _, v := range []any{ts, &ts} {
		got, err := MarshalToString(v)
		require.NoError(t, err)
		assert.JSONEq(t, `{ "@type": "g:Timestamp", "@value": 1481750076295 }`, got)
	}

	strs := []string{
		`{ "@type": "g:Timestamp", "@value": 1481750076295 }`,
		`{ "@type": "g:Date", "@value": 1481750076295 }`,
	}

	for _, str := range strs {
		var v time.Time
		err := UnmarshalFromString(str, &v)
		assert.NoError(t, err)
		assert.True(t, ts.Equal(v))
	}
}
