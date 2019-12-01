// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestScanSlice(t *testing.T) {
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("foo").
		AddRow("bar")
	var v0 []string
	require.NoError(t, scanSlice(rows, &v0))
	require.Equal(t, []string{"foo", "bar"}, v0)

	rows = sqlmock.NewRows([]string{"age"}).
		AddRow(1).
		AddRow(2)
	var v1 []int
	require.NoError(t, scanSlice(rows, &v1))
	require.Equal(t, []int{1, 2}, v1)

	rows = sqlmock.NewRows([]string{"name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v2 []struct {
		Name  string
		Count int
	}
	require.NoError(t, scanSlice(rows, &v2))
	require.Equal(t, "foo", v2[0].Name)
	require.Equal(t, "bar", v2[1].Name)
	require.Equal(t, 1, v2[0].Count)
	require.Equal(t, 2, v2[1].Count)

	rows = sqlmock.NewRows([]string{"nick_name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v3 []struct {
		Count int
		Name  string `json:"nick_name"`
	}
	require.NoError(t, scanSlice(rows, &v3))
	require.Equal(t, "foo", v3[0].Name)
	require.Equal(t, "bar", v3[1].Name)
	require.Equal(t, 1, v3[0].Count)
	require.Equal(t, 2, v3[1].Count)

	rows = sqlmock.NewRows([]string{"nick_name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v4 []*struct {
		Count   int
		Name    string `json:"nick_name"`
		Ignored string `json:"string"`
	}
	require.NoError(t, scanSlice(rows, &v4))
	require.Equal(t, "foo", v4[0].Name)
	require.Equal(t, "bar", v4[1].Name)
	require.Equal(t, 1, v4[0].Count)
	require.Equal(t, 2, v4[1].Count)

	rows = sqlmock.NewRows([]string{"nick_name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v5 []*struct {
		Count int
		Name  string `json:"name" sql:"nick_name"`
	}
	require.NoError(t, scanSlice(rows, &v5))
	require.Equal(t, "foo", v5[0].Name)
	require.Equal(t, "bar", v5[1].Name)
	require.Equal(t, 1, v5[0].Count)
	require.Equal(t, 2, v5[1].Count)
}

func TestScanSlicePtr(t *testing.T) {
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("foo").
		AddRow("bar")
	var v0 []*string
	require.NoError(t, scanSlice(rows, &v0))
	require.Equal(t, "foo", *v0[0])
	require.Equal(t, "bar", *v0[1])

	rows = sqlmock.NewRows([]string{"age"}).
		AddRow(1).
		AddRow(2)
	var v1 []**int
	require.NoError(t, scanSlice(rows, &v1))
	require.Equal(t, 1, **v1[0])
	require.Equal(t, 2, **v1[1])

	rows = sqlmock.NewRows([]string{"age", "name"}).
		AddRow(1, "a8m").
		AddRow(2, "nati")
	var v2 []*struct {
		Age  *int
		Name **string
	}
	require.NoError(t, scanSlice(rows, &v2))
	require.Equal(t, 1, *v2[0].Age)
	require.Equal(t, "a8m", **v2[0].Name)
	require.Equal(t, 2, *v2[1].Age)
	require.Equal(t, "nati", **v2[1].Name)
}

func scanSlice(mrows *sqlmock.Rows, v interface{}) error {
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery("").WillReturnRows(mrows)
	rows, _ := db.Query("")
	return ScanSlice(rows, v)
}
