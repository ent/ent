// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanSlice(t *testing.T) {
	rows := &mockRows{
		columns: []string{"name"},
		values:  [][]interface{}{{"foo"}, {"bar"}},
	}
	var v0 []string
	require.NoError(t, ScanSlice(rows, &v0))
	require.Equal(t, []string{"foo", "bar"}, v0)

	rows = &mockRows{
		columns: []string{"age"},
		values:  [][]interface{}{{1}, {2}},
	}
	var v1 []int
	require.NoError(t, ScanSlice(rows, &v1))
	require.Equal(t, []int{1, 2}, v1)

	rows = &mockRows{
		columns: []string{"name", "COUNT(*)"},
		values:  [][]interface{}{{"foo", 1}, {"bar", 2}},
	}
	var v2 []struct {
		Name  string
		Count int
	}
	require.NoError(t, ScanSlice(rows, &v2))
	require.Equal(t, "foo", v2[0].Name)
	require.Equal(t, "bar", v2[1].Name)
	require.Equal(t, 1, v2[0].Count)
	require.Equal(t, 2, v2[1].Count)

	rows = &mockRows{
		columns: []string{"nick_name", "COUNT(*)"},
		values:  [][]interface{}{{"foo", 1}, {"bar", 2}},
	}
	var v3 []struct {
		Count int
		Name  string `json:"nick_name"`
	}
	require.NoError(t, ScanSlice(rows, &v3))
	require.Equal(t, "foo", v3[0].Name)
	require.Equal(t, "bar", v3[1].Name)
	require.Equal(t, 1, v3[0].Count)
	require.Equal(t, 2, v3[1].Count)

	rows = &mockRows{
		columns: []string{"nick_name", "COUNT(*)"},
		values:  [][]interface{}{{"foo", 1}, {"bar", 2}},
	}
	var v4 []*struct {
		Count   int
		Name    string `json:"nick_name"`
		Ignored string `json:"string"`
	}
	require.NoError(t, ScanSlice(rows, &v4))
	require.Equal(t, "foo", v4[0].Name)
	require.Equal(t, "bar", v4[1].Name)
	require.Equal(t, 1, v4[0].Count)
	require.Equal(t, 2, v4[1].Count)

	rows = &mockRows{
		columns: []string{"nick_name", "COUNT(*)"},
		values:  [][]interface{}{{"foo", 1}, {"bar", 2}},
	}
	var v5 []*struct {
		Count int
		Name  string `json:"name" sql:"nick_name"`
	}
	require.NoError(t, ScanSlice(rows, &v5))
	require.Equal(t, "foo", v5[0].Name)
	require.Equal(t, "bar", v5[1].Name)
	require.Equal(t, 1, v5[0].Count)
	require.Equal(t, 2, v5[1].Count)
}

type mockRows struct {
	columns []string
	values  [][]interface{}
}

func (m mockRows) Columns() ([]string, error) { return m.columns, nil }

func (m mockRows) Next() bool { return len(m.values) > 0 }

func (m *mockRows) Scan(vs ...interface{}) error {
	if len(m.values) == 0 {
		return sql.ErrNoRows
	}
	row := m.values[0]
	m.values = m.values[1:]
	for i := range vs {
		reflect.Indirect(reflect.ValueOf(vs[i])).Set(reflect.ValueOf(row[i]))
	}
	return nil
}
