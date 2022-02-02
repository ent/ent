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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent"
	"entgo.io/ent/entc/integration/ent/fieldtype"
	"entgo.io/ent/entc/integration/ent/role"
	"entgo.io/ent/entc/integration/ent/schema"
	"entgo.io/ent/entc/integration/ent/schema/task"
	enttask "entgo.io/ent/entc/integration/ent/task"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Types(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	require := require.New(t)

	link, err := url.Parse("localhost")
	require.NoError(err)

	bigint := schema.NewBigInt(0)
	require.NoError(bigint.Scan("1000"))

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
	require.Nil(ft.NullLink)
	require.Nil(ft.NilPair)
	require.Nil(ft.Deleted)

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
		SetNullStr(&sql.NullString{String: "not-default", Valid: true}).
		SetLink(schema.Link{URL: link}).
		SetLinkOther(&schema.Link{URL: link}).
		SetNullLink(&schema.Link{URL: link}).
		SetRole(role.Admin).
		SetPriority(role.High).
		SetDuration(time.Hour).
		SetPair(schema.Pair{K: []byte("K"), V: []byte("V")}).
		SetNilPair(&schema.Pair{K: []byte("K"), V: []byte("V")}).
		SetStringArray([]string{"foo", "bar", "baz"}).
		SetBigInt(bigint).
		SetRawData([]byte{1, 2, 3}).
		SaveX(ctx)

	require.Equal(int8(math.MinInt8), ft.OptionalInt8)
	require.Equal(int16(math.MinInt16), ft.OptionalInt16)
	require.Equal(int32(math.MinInt32), ft.OptionalInt32)
	require.Equal(int64(math.MinInt64), ft.OptionalInt64)
	require.Equal(int8(math.MinInt8), *ft.NillableInt8)
	require.Equal(int16(math.MinInt16), *ft.NillableInt16)
	require.Equal(int32(math.MinInt32), *ft.NillableInt32)
	require.Equal(int64(math.MinInt64), *ft.NillableInt64)
	require.Equal([]byte{1, 2, 3}, ft.RawData)
	require.Equal(http.Dir("dir"), ft.Dir)
	require.NotNil(*ft.Ndir)
	require.Equal(http.Dir("ndir"), *ft.Ndir)
	require.Equal("default", ft.Str.String)
	require.Equal("not-default", ft.NullStr.String)
	require.Equal("localhost", ft.Link.String())
	require.Equal("localhost", ft.LinkOther.String())
	require.Equal("localhost", ft.NullLink.String())
	require.Equal(net.IP("127.0.0.1").String(), ft.IP.String())
	mac, err := net.ParseMAC("3b:b3:6b:3c:10:79")
	require.Equal(role.Admin, ft.Role)
	require.Equal(role.High, ft.Priority)
	require.NoError(err)
	dt, err := time.Parse(time.RFC3339, "1906-01-02T00:00:00+00:00")
	require.NoError(err)
	require.Equal(schema.Pair{K: []byte("K"), V: []byte("V")}, ft.Pair)
	require.Equal(&schema.Pair{K: []byte("K"), V: []byte("V")}, ft.NilPair)
	require.EqualValues([]string{"foo", "bar", "baz"}, ft.StringArray)
	require.Equal("1000", ft.BigInt.String())
	exists, err := client.FieldType.Query().Where(fieldtype.DurationLT(time.Hour * 2)).Exist(ctx)
	require.NoError(err)
	require.True(exists)
	exists, err = client.FieldType.Query().Where(fieldtype.DurationLT(time.Hour)).Exist(ctx)
	require.NoError(err)
	require.False(exists)
	require.Equal("127.0.0.1", ft.LinkOtherFunc.String())
	require.False(ft.DeletedAt.Time.IsZero())

	ft = client.FieldType.UpdateOne(ft).AddOptionalUint64(10).SaveX(ctx)
	require.EqualValues(10, ft.OptionalUint64)
	ft = client.FieldType.UpdateOne(ft).AddOptionalUint64(20).SetOptionalUint64(5).SaveX(ctx)
	require.EqualValues(5, ft.OptionalUint64)
	ft = client.FieldType.UpdateOne(ft).AddOptionalUint64(-5).SaveX(ctx)
	require.Zero(ft.OptionalUint64)

	err = client.FieldType.Create().
		SetInt(1).
		SetInt8(8).
		SetInt16(16).
		SetInt32(32).
		SetInt64(64).
		SetRawData(make([]byte, 40)).
		Exec(ctx)
	require.Error(err, "MaxLen validator should reject this operation")
	err = client.FieldType.Create().
		SetInt(1).
		SetInt8(8).
		SetInt16(16).
		SetInt32(32).
		SetInt64(64).
		SetRawData(make([]byte, 2)).
		Exec(ctx)
	require.Error(err, "MinLen validator should reject this operation")
	ft = ft.Update().
		SetInt(1).
		SetInt8(math.MaxInt8).
		SetInt16(math.MaxInt16).
		SetInt32(math.MaxInt16).
		SetOptionalInt8(math.MaxInt8).
		SetOptionalInt16(math.MaxInt16).
		SetOptionalInt32(math.MaxInt32).
		SetOptionalInt64(math.MaxInt64).
		SetNillableInt8(math.MaxInt8).
		SetNillableInt16(math.MaxInt16).
		SetNillableInt32(math.MaxInt32).
		SetNillableInt64(math.MaxInt64).
		SetDatetime(dt).
		SetDecimal(10.20).
		SetDir("dir").
		SetNdir("ndir").
		SetStr(sql.NullString{String: "str", Valid: true}).
		SetNullStr(&sql.NullString{String: "str", Valid: true}).
		SetLink(schema.Link{URL: link}).
		SetNullLink(&schema.Link{URL: link}).
		SetLinkOther(&schema.Link{URL: link}).
		SetSchemaInt(64).
		SetSchemaInt8(8).
		SetSchemaInt64(64).
		SetMAC(schema.MAC{HardwareAddr: mac}).
		SetPair(schema.Pair{K: []byte("K1"), V: []byte("V1")}).
		SetNilPair(&schema.Pair{K: []byte("K1"), V: []byte("V1")}).
		SetStringArray([]string{"qux"}).
		AddBigInt(bigint).
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
	require.True(dt.Equal(ft.Datetime))
	require.Equal(http.Dir("dir"), ft.Dir)
	require.NotNil(*ft.Ndir)
	require.Equal(http.Dir("ndir"), *ft.Ndir)
	require.Equal("str", ft.Str.String)
	require.Equal("str", ft.NullStr.String)
	require.Equal("localhost", ft.Link.String())
	require.Equal("localhost", ft.LinkOther.String())
	require.Equal("localhost", ft.NullLink.String())
	require.Equal(schema.Int(64), ft.SchemaInt)
	require.Equal(schema.Int8(8), ft.SchemaInt8)
	require.Equal(schema.Int64(64), ft.SchemaInt64)
	require.Equal(mac.String(), ft.MAC.String())
	require.Equal(schema.Pair{K: []byte("K1"), V: []byte("V1")}, ft.Pair)
	require.Equal(&schema.Pair{K: []byte("K1"), V: []byte("V1")}, ft.NilPair)
	require.EqualValues([]string{"qux"}, ft.StringArray)
	require.Nil(ft.NillableUUID)
	require.Equal(uuid.UUID{}, ft.OptionalUUID)
	require.Equal("2000", ft.BigInt.String())
	require.EqualValues(100, ft.Int64, "UpdateDefault sets the value to 100")
	require.EqualValues(100, ft.Duration, "UpdateDefault sets the value to 100ns")
	require.False(ft.DeletedAt.Time.IsZero())

	err = client.Task.CreateBulk(
		client.Task.Create().SetPriority(task.PriorityLow),
		client.Task.Create().SetPriority(task.PriorityMid),
		client.Task.Create().SetPriority(task.PriorityHigh),
	).Exec(ctx)
	require.NoError(err)
	err = client.Task.Create().SetPriority(task.Priority(10)).Exec(ctx)
	require.Error(err)

	tasks := client.Task.Query().Order(ent.Asc(enttask.FieldPriority)).AllX(ctx)
	require.Equal(task.PriorityLow, tasks[0].Priority)
	require.Equal(task.PriorityMid, tasks[1].Priority)
	require.Equal(task.PriorityHigh, tasks[2].Priority)

	tasks = client.Task.Query().Order(ent.Desc(enttask.FieldPriority)).AllX(ctx)
	require.Equal(task.PriorityLow, tasks[2].Priority)
	require.Equal(task.PriorityMid, tasks[1].Priority)
	require.Equal(task.PriorityHigh, tasks[0].Priority)
}
