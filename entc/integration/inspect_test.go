// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.
package integration

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/stretchr/testify/require"
	"testing"
)

func Inspect(t *testing.T, drv *sql.Driver) {
	inspect, err := schema.NewInspect(drv)
	require.NoError(t, err)
	_, err = inspect.Tables(context.Background())
	require.NoError(t, err)
}
