// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"math"
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/facebook/ent/entc/integration/ent/task"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/ent"
	"github.com/facebook/ent/entc/integration/ent/role"
	"github.com/facebook/ent/entc/integration/ent/schema"

	"github.com/stretchr/testify/require"
)

func Types(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)

	link, err := url.Parse("localhost")
	require.NoError(err)

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
		SetDir("dir").
		SetNdir("ndir").
		SetStr(sql.NullString{String: "str", Valid: true}).
		SetNullStr(sql.NullString{String: "str", Valid: true}).
		SetLink(schema.Link{URL: link}).
		SetNullLink(schema.Link{URL: link}).
		SetRole(role.Admin).
		SaveX(ctx)

	require.Equal(int8(math.MinInt8), ft.OptionalInt8)
	require.Equal(int16(math.MinInt16), ft.OptionalInt16)
	require.Equal(int32(math.MinInt32), ft.OptionalInt32)
	require.Equal(int64(math.MinInt64), ft.OptionalInt64)
	require.Equal(int8(math.MinInt8), *ft.NillableInt8)
	require.Equal(int16(math.MinInt16), *ft.NillableInt16)
	require.Equal(int32(math.MinInt32), *ft.NillableInt32)
	require.Equal(int64(math.MinInt64), *ft.NillableInt64)
	require.Equal(http.Dir("dir"), ft.Dir)
	require.NotNil(*ft.Ndir)
	require.Equal(http.Dir("ndir"), *ft.Ndir)
	require.Equal("str", ft.Str.String)
	require.Equal("str", ft.NullStr.String)
	require.Equal("localhost", ft.Link.String())
	require.Equal("localhost", ft.NullLink.String())
	mac, err := net.ParseMAC("3b:b3:6b:3c:10:79")
	require.NoError(err)

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
		SetNdir("ndir").
		SetStr(sql.NullString{String: "str", Valid: true}).
		SetNullStr(sql.NullString{String: "str", Valid: true}).
		SetLink(schema.Link{URL: link}).
		SetNullLink(schema.Link{URL: link}).
		SetSchemaInt(64).
		SetSchemaInt8(8).
		SetSchemaInt64(64).
		SetMAC(schema.MAC{HardwareAddr: mac}).
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
	require.False(ft.Datetime.IsZero())
	require.Equal(http.Dir("dir"), ft.Dir)
	require.NotNil(*ft.Ndir)
	require.Equal(http.Dir("ndir"), *ft.Ndir)
	require.Equal("str", ft.Str.String)
	require.Equal("str", ft.NullStr.String)
	require.Equal("localhost", ft.Link.String())
	require.Equal("localhost", ft.NullLink.String())
	require.Equal(schema.Int(64), ft.SchemaInt)
	require.Equal(schema.Int8(8), ft.SchemaInt8)
	require.Equal(schema.Int64(64), ft.SchemaInt64)
	require.Equal(mac.String(), ft.MAC.String())

	_, err = client.Task.CreateBulk(
		client.Task.Create().SetPriority(schema.PriorityLow),
		client.Task.Create().SetPriority(schema.PriorityMid),
		client.Task.Create().SetPriority(schema.PriorityHigh),
	).Save(ctx)
	require.NoError(err)
	_, err = client.Task.Create().SetPriority(schema.Priority(10)).Save(ctx)
	require.Error(err)

	tasks := client.Task.Query().Order(ent.Asc(task.FieldPriority)).AllX(ctx)
	require.Equal(schema.PriorityLow, tasks[0].Priority)
	require.Equal(schema.PriorityMid, tasks[1].Priority)
	require.Equal(schema.PriorityHigh, tasks[2].Priority)

	tasks = client.Task.Query().Order(ent.Desc(task.FieldPriority)).AllX(ctx)
	require.Equal(schema.PriorityLow, tasks[2].Priority)
	require.Equal(schema.PriorityMid, tasks[1].Priority)
	require.Equal(schema.PriorityHigh, tasks[0].Priority)
}
