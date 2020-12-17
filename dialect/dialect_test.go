// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dialect_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/facebook/ent/dialect"

	"github.com/stretchr/testify/require"
)

type driver struct {
	dialect.Driver
}

func (driver) Dialect() string { return "driver" }

func TestAlwaysFail(t *testing.T) {
	d := driver{}
	fd := dialect.AlwaysFail(d, fmt.Errorf("error"))
	_, err := fd.Tx(context.Background())
	require.Error(t, err)
	require.EqualError(t, err, "error")
	require.Equal(t, d.Dialect(), fd.Dialect())
}
