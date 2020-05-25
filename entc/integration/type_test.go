// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"math"
	"net/http"
	"testing"
	"time"

	"github.com/facebookincubator/ent/entc/integration/ent"

	"github.com/stretchr/testify/require"
)

func Types(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)

	ft := client.FieldType.Create().
		SetInt(1).
		SetInt8(8).
		SetInt16(16).
		SetInt32(32).
		SetInt64(64).
		SaveX(ctx)

	require.NotEmpty(t, ft.ID)
	require.Equal(1, ft.Int)
	require.Equal(int8(8), ft.Int8)
	require.Equal(int16(16), ft.Int16)
	require.Equal(int32(32), ft.Int32)
	require.Equal(int64(64), ft.Int64)

	ft = client.FieldType.Create().
		SetInt(1).
		SetInt8(math.MinInt8).
		SetInt16(math.MinInt16).
		SetInt32(math.MinInt16).
		SetInt64(math.MinInt16).
		SetOptionalInt8(math.MinInt8).
		SetOptionalInt16(math.MinInt16).
		SetOptionalInt32(math.MinInt32).
		SetOptionalInt64(math.MinInt64).
		SetNillableInt8(math.MinInt8).
		SetNillableInt16(math.MinInt16).
		SetNillableInt32(math.MinInt32).
		SetNillableInt64(math.MinInt64).
		SaveX(ctx)

	require.Equal(int8(math.MinInt8), ft.OptionalInt8)
	require.Equal(int16(math.MinInt16), ft.OptionalInt16)
	require.Equal(int32(math.MinInt32), ft.OptionalInt32)
	require.Equal(int64(math.MinInt64), ft.OptionalInt64)
	require.Equal(int8(math.MinInt8), *ft.NillableInt8)
	require.Equal(int16(math.MinInt16), *ft.NillableInt16)
	require.Equal(int32(math.MinInt32), *ft.NillableInt32)
	require.Equal(int64(math.MinInt64), *ft.NillableInt64)

	ft = ft.Update().
		SetInt(1).
		SetInt8(math.MaxInt8).
		SetInt16(math.MaxInt16).
		SetInt32(math.MaxInt16).
		SetInt64(math.MaxInt16).
		SetOptionalInt8(math.MaxInt8).
		SetOptionalInt16(math.MaxInt16).
		SetOptionalInt32(math.MaxInt32).
		SetOptionalInt64(math.MaxInt64).
		SetNillableInt8(math.MaxInt8).
		SetNillableInt16(math.MaxInt16).
		SetNillableInt32(math.MaxInt32).
		SetNillableInt64(math.MaxInt64).
		SetDatetime(time.Now()).
		SetDecimal(10.20).
		SetDir("dir").
		SaveX(ctx)

	require.Equal(int8(math.MaxInt8), ft.OptionalInt8)
	require.Equal(int16(math.MaxInt16), ft.OptionalInt16)
	require.Equal(int32(math.MaxInt32), ft.OptionalInt32)
	require.Equal(int64(math.MaxInt64), ft.OptionalInt64)
	require.Equal(int8(math.MaxInt8), *ft.NillableInt8)
	require.Equal(int16(math.MaxInt16), *ft.NillableInt16)
	require.Equal(int32(math.MaxInt32), *ft.NillableInt32)
	require.Equal(int64(math.MaxInt64), *ft.NillableInt64)
	require.Equal(10.20, ft.Decimal)
	require.Equal(http.Dir("dir"), ft.Dir)
	require.False(ft.Datetime.IsZero())
}
