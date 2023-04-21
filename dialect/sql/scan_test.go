// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestScanSlice(t *testing.T) {
	mock := sqlmock.NewRows([]string{"name"}).
		AddRow("foo").
		AddRow("bar")
	var v0 []string
	require.NoError(t, ScanSlice(toRows(mock), &v0))
	require.Equal(t, []string{"foo", "bar"}, v0)

	mock = sqlmock.NewRows([]string{"age"}).
		AddRow(1).
		AddRow(2)
	var v1 []int
	require.NoError(t, ScanSlice(toRows(mock), &v1))
	require.Equal(t, []int{1, 2}, v1)

	mock = sqlmock.NewRows([]string{"name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v2 []struct {
		Name  string
		Count int
	}
	require.NoError(t, ScanSlice(toRows(mock), &v2))
	require.Equal(t, "foo", v2[0].Name)
	require.Equal(t, "bar", v2[1].Name)
	require.Equal(t, 1, v2[0].Count)
	require.Equal(t, 2, v2[1].Count)

	mock = sqlmock.NewRows([]string{"nick_name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v3 []struct {
		Count int
		Name  string `json:"nick_name"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v3))
	require.Equal(t, "foo", v3[0].Name)
	require.Equal(t, "bar", v3[1].Name)
	require.Equal(t, 1, v3[0].Count)
	require.Equal(t, 2, v3[1].Count)

	mock = sqlmock.NewRows([]string{"nick_name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v4 []*struct {
		Count   int
		Name    string `json:"nick_name"`
		Ignored string `json:"string"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v4))
	require.Equal(t, "foo", v4[0].Name)
	require.Equal(t, "bar", v4[1].Name)
	require.Equal(t, 1, v4[0].Count)
	require.Equal(t, 2, v4[1].Count)

	mock = sqlmock.NewRows([]string{"nick_name", "COUNT(*)"}).
		AddRow("foo", 1).
		AddRow("bar", 2)
	var v5 []*struct {
		Count int
		Name  string `json:"name" sql:"nick_name"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v5))
	require.Equal(t, "foo", v5[0].Name)
	require.Equal(t, "bar", v5[1].Name)
	require.Equal(t, 1, v5[0].Count)
	require.Equal(t, 2, v5[1].Count)

	mock = sqlmock.NewRows([]string{"age", "name"}).
		AddRow(1, nil).
		AddRow(nil, "a8m")
	var v6 []struct {
		Age  NullInt64
		Name NullString
	}
	require.NoError(t, ScanSlice(toRows(mock), &v6))
	require.EqualValues(t, 1, v6[0].Age.Int64)
	require.False(t, v6[0].Name.Valid)
	require.False(t, v6[1].Age.Valid)
	require.Equal(t, "a8m", v6[1].Name.String)

	u1, u2 := uuid.New().String(), uuid.New().String()
	mock = sqlmock.NewRows([]string{"ids"}).
		AddRow([]byte(u1)).
		AddRow([]byte(u2))
	var ids []uuid.UUID
	require.NoError(t, ScanSlice(toRows(mock), &ids))
	require.Equal(t, u1, ids[0].String())
	require.Equal(t, u2, ids[1].String())

	mock = sqlmock.NewRows([]string{"pids"}).
		AddRow([]byte(u1)).
		AddRow([]byte(u2))
	var pids []*uuid.UUID
	require.NoError(t, ScanSlice(toRows(mock), &pids))
	require.Equal(t, u1, pids[0].String())
	require.Equal(t, u2, pids[1].String())

	mock = sqlmock.NewRows([]string{"id", "first", "last"}).
		AddRow(10, "Ariel", "Mashraki")
	err := ScanSlice(toRows(mock), nil)
	require.EqualError(t, err, "sql/scan: ScanSlice(nil)")
	type P struct {
		_     int
		ID    int
		First string
		Last  string
	}
	var p []P
	err = ScanSlice(toRows(mock), p)
	require.EqualError(t, err, "sql/scan: ScanSlice(non-pointer []sql.P)")

	require.NoError(t, ScanSlice(toRows(mock), &p))
	require.Equal(t, 10, p[0].ID)
	require.Equal(t, "Ariel", p[0].First)
	require.Equal(t, "Mashraki", p[0].Last)

	var pp []struct{ _, id int }
	mock = sqlmock.NewRows([]string{"id"}).
		AddRow(10)
	err = ScanSlice(toRows(mock), &pp)
	require.EqualError(t, err, "sql/scan: missing struct field for column: id (id)")
	require.Empty(t, pp)
}

func TestScanSlice_CamelTags(t *testing.T) {
	mock := sqlmock.NewRows([]string{"nickName"}).
		AddRow("foo").
		AddRow("bar")
	var v []*struct {
		NickName string `json:"nickName"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v))
	require.Equal(t, "foo", v[0].NickName)
}

func TestScanJSON(t *testing.T) {
	mock := sqlmock.NewRows([]string{"v", "p"}).
		AddRow([]byte(`{"i": 1, "s":"a8m"}`), []byte(`{"i": 1, "s":"a8m"}`)).
		AddRow([]byte(`{"i": 2, "s":"tmr"}`), []byte(`{"i": 2, "s":"tmr"}`)).
		AddRow([]byte(nil), []byte(`null`)).
		AddRow(nil, nil)
	var v1 []*struct {
		V struct {
			I int    `json:"i"`
			S string `json:"s"`
		} `json:"v"`
		P *struct {
			I int    `json:"i"`
			S string `json:"s"`
		} `json:"p"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v1))
	require.Equal(t, 1, v1[0].V.I)
	require.Equal(t, "a8m", v1[0].V.S)
	require.Equal(t, v1[0].V, *v1[0].P)
	require.Equal(t, 2, v1[1].V.I)
	require.Equal(t, "tmr", v1[1].V.S)
	require.Equal(t, v1[1].V, *v1[1].P)
	require.Equal(t, 0, v1[2].V.I)
	require.Equal(t, "", v1[2].V.S)
	require.Nil(t, v1[2].P)
	require.Equal(t, 0, v1[3].V.I)
	require.Equal(t, "", v1[3].V.S)
	require.Nil(t, v1[3].P)

	mock = sqlmock.NewRows([]string{"v", "p"}).
		AddRow([]byte(`[1]`), []byte(`[1]`)).
		AddRow([]byte(`[2]`), []byte(`[2]`))
	var v2 []*struct {
		V []int  `json:"v"`
		P *[]int `json:"p"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v2))
	require.Equal(t, []int{1}, v2[0].V)
	require.Equal(t, v2[0].V, *v2[0].P)
	require.Equal(t, []int{2}, v2[1].V)
	require.Equal(t, v2[1].V, *v2[1].P)

	mock = sqlmock.NewRows([]string{"v", "p"}).
		AddRow([]byte(`null`), []byte(`{}`)).
		AddRow(nil, nil)
	var v3 []*struct {
		V json.RawMessage  `json:"v"`
		P *json.RawMessage `json:"p"`
	}
	require.NoError(t, ScanSlice(toRows(mock), &v3))
	require.Equal(t, json.RawMessage("null"), v3[0].V)
	require.Equal(t, json.RawMessage("{}"), *v3[0].P)
	require.Equal(t, json.RawMessage(nil), v3[1].V)
	require.Nil(t, v3[1].P)

	// Unmarshal errors.
	mock = sqlmock.NewRows([]string{"v", "p"}).
		AddRow([]byte(`{invalid}`), []byte(`{}`))
	require.EqualError(t, ScanSlice(toRows(mock), &v1), `unmarshal field "V": invalid character 'i' looking for beginning of object key string`)
	mock = sqlmock.NewRows([]string{"v", "p"}).
		AddRow([]byte(``), []byte(``))
	require.EqualError(t, ScanSlice(toRows(mock), &v1), `unmarshal field "V": unexpected end of JSON input`)
}

func TestScanNestedStruct(t *testing.T) {
	mock := sqlmock.NewRows([]string{"name", "age"}).
		AddRow("foo", 1).
		AddRow("bar", 2).
		AddRow("baz", nil)
	type T struct{ Name string }
	var v []struct {
		T
		Age int
	}
	require.NoError(t, ScanSlice(toRows(mock), &v))
	require.Equal(t, "foo", v[0].Name)
	require.Equal(t, 1, v[0].Age)
	require.Equal(t, "bar", v[1].Name)
	require.Equal(t, 2, v[1].Age)
	require.Equal(t, "baz", v[2].Name)
	require.Equal(t, 0, v[2].Age)

	mock = sqlmock.NewRows([]string{"name", "age"}).
		AddRow("foo", 1).
		AddRow("bar", nil)
	type T1 struct{ Name **string }
	var v1 []struct {
		T1
		Age *int
	}
	require.NoError(t, ScanSlice(toRows(mock), &v1))
	require.Equal(t, "foo", **v1[0].Name)
	require.Equal(t, "bar", **v1[1].Name)
	require.Equal(t, 1, *v1[0].Age)
	require.Nil(t, v1[1].Age)
}

func TestScanSlicePtr(t *testing.T) {
	mock := sqlmock.NewRows([]string{"name"}).
		AddRow("foo").
		AddRow("bar")
	var v0 []*string
	require.NoError(t, ScanSlice(toRows(mock), &v0))
	require.Equal(t, "foo", *v0[0])
	require.Equal(t, "bar", *v0[1])

	mock = sqlmock.NewRows([]string{"age"}).
		AddRow(1).
		AddRow(2)
	var v1 []**int
	require.NoError(t, ScanSlice(toRows(mock), &v1))
	require.Equal(t, 1, **v1[0])
	require.Equal(t, 2, **v1[1])

	mock = sqlmock.NewRows([]string{"age", "name"}).
		AddRow(1, "a8m").
		AddRow(2, "nati")
	var v2 []*struct {
		Age  *int
		Name **string
	}
	require.NoError(t, ScanSlice(toRows(mock), &v2))
	require.Equal(t, 1, *v2[0].Age)
	require.Equal(t, "a8m", **v2[0].Name)
	require.Equal(t, 2, *v2[1].Age)
	require.Equal(t, "nati", **v2[1].Name)
}

func TestScanInt64(t *testing.T) {
	mock := sqlmock.NewRows([]string{"age"}).
		AddRow("10").
		AddRow("20")
	n, err := ScanInt64(toRows(mock))
	require.Error(t, err)
	require.Zero(t, n)

	mock = sqlmock.NewRows([]string{"age", "count"}).
		AddRow("10", "1")
	n, err = ScanInt64(toRows(mock))
	require.Error(t, err)
	require.Zero(t, n)

	mock = sqlmock.NewRows([]string{"count"}).
		AddRow(10)
	n, err = ScanInt64(toRows(mock))
	require.NoError(t, err)
	require.EqualValues(t, 10, n)

	mock = sqlmock.NewRows([]string{"count"}).
		AddRow("10")
	n, err = ScanInt64(toRows(mock))
	require.NoError(t, err)
	require.EqualValues(t, 10, n)
}

func TestScanValue(t *testing.T) {
	mock := sqlmock.NewRows([]string{"count"}).
		AddRow(10)
	n, err := ScanValue(toRows(mock))
	require.NoError(t, err)
	require.EqualValues(t, 10, n)
}

func TestScanOne(t *testing.T) {
	mock := sqlmock.NewRows([]string{"name"}).
		AddRow("10").
		AddRow("20")
	err := ScanOne(toRows(mock), new(string))
	require.Error(t, err, "multiple lines")

	mock = sqlmock.NewRows([]string{"name"}).
		AddRow("10")
	err = ScanOne(toRows(mock), "")
	require.Error(t, err, "not a pointer")

	mock = sqlmock.NewRows([]string{"name"}).
		AddRow("10")
	var s string
	err = ScanOne(toRows(mock), &s)
	require.NoError(t, err)
	require.Equal(t, "10", s)
}

func TestInterface(t *testing.T) {
	mock := sqlmock.NewRows([]string{"age"}).
		AddRow("10").
		AddRow("20")
	var values []driver.Value
	err := ScanSlice(toRows(mock), &values)
	require.NoError(t, err)
	require.Equal(t, []driver.Value{"10", "20"}, values)

	mock = sqlmock.NewRows([]string{"age"}).
		AddRow(10).
		AddRow(20)
	values = values[:0:0]
	err = ScanSlice(toRows(mock), &values)
	require.NoError(t, err)
	require.Equal(t, []driver.Value{int64(10), int64(20)}, values)
}

func TestScanTypeOf(t *testing.T) {
	mock := &Rows{ColumnScanner: toRows(sqlmock.NewRows([]string{"age"}).AddRow(10))}
	tv := ScanTypeOf(mock, 0)
	require.IsType(t, (*any)(nil), tv)
}

func toRows(mrows *sqlmock.Rows) *sql.Rows {
	db, mock, _ := sqlmock.New()
	mock.ExpectQuery("").WillReturnRows(mrows)
	rows, _ := db.Query("")
	return rows
}
