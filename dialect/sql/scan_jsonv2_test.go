// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package sql

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestScanSliceJSONOptions(t *testing.T) {
	type payload struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
	var values []*struct {
		Legacy     payload
		Modern     payload
		Configured payload
	}
	mock := sqlmock.NewRows([]string{"legacy", "modern", "configured"}).
		AddRow([]byte(`{"NAME":"first","NAME":"legacy"}`), `{"NAME":"modern"}`, `{"NAME":"configured","count":"42"}`).
		AddRow(nil, nil, nil)
	rows := &jsonOptionsRows{
		ColumnScanner: toRows(mock),
		options: func(column string) []jsonv2.Options {
			switch column {
			case "legacy":
				return []jsonv2.Options{json.DefaultOptionsV1()}
			case "configured":
				return []jsonv2.Options{jsonv2.MatchCaseInsensitiveNames(true), jsonv2.StringifyNumbers(true)}
			default:
				return nil
			}
		},
	}
	defer rows.Close()
	require.NoError(t, ScanSlice(rows, &values))
	require.Equal(t, "legacy", values[0].Legacy.Name)
	require.Empty(t, values[0].Modern.Name)
	require.Empty(t, values[1].Modern.Name)
	require.Equal(t, payload{Name: "configured", Count: 42}, values[0].Configured)
	require.Empty(t, values[1].Configured)
	require.Equal(t, []string{"legacy", "modern", "configured"}, rows.calls)

	var maps []map[string]int
	mock = sqlmock.NewRows([]string{"data"}).AddRow(`{"a":1}`).AddRow(nil)
	require.NoError(t, ScanSlice(&jsonOptionsRows{ColumnScanner: toRows(mock)}, &maps))
	require.Equal(t, []map[string]int{{"a": 1}, nil}, maps)

	var arrays [][]bool
	mock = sqlmock.NewRows([]string{"data"}).AddRow(`[true,false]`).AddRow(`[]`).AddRow(`null`)
	require.NoError(t, ScanSlice(&jsonOptionsRows{ColumnScanner: toRows(mock)}, &arrays))
	require.Equal(t, [][]bool{{true, false}, {}, nil}, arrays)

	for _, raw := range []string{`{"name":"a","name":"b"}`, "{\"name\":\"\xff\"}"} {
		mock = sqlmock.NewRows([]string{"modern"}).AddRow(raw)
		require.ErrorContains(t, ScanSlice(&jsonOptionsRows{ColumnScanner: toRows(mock)}, &values), "unmarshal field")
	}
	mock = sqlmock.NewRows([]string{"data"}).AddRow(`{"a":1,"a":2}`)
	require.ErrorContains(t, ScanSlice(&jsonOptionsRows{ColumnScanner: toRows(mock)}, &maps), "unmarshal column")
}

func TestScanSliceJSONOptionsProjection(t *testing.T) {
	// Column names, including aliases, determine the options independently of
	// the destination field order. Scalar projections keep SQL scan semantics.
	var projected []struct {
		Data map[string]int `sql:"payload"`
		Name string         `sql:"label"`
		Raw  []byte
	}
	rows := &jsonOptionsRows{
		ColumnScanner: toRows(sqlmock.NewRows([]string{"label", "raw", "payload"}).
			AddRow(`"quoted"`, []byte(`{"a":1,"a":2}`), `{"count":"42"}`)),
		options: func(column string) []jsonv2.Options {
			if column == "payload" {
				return []jsonv2.Options{jsonv2.StringifyNumbers(true)}
			}
			return nil
		},
	}
	defer rows.Close()
	require.NoError(t, ScanSlice(rows, &projected))
	require.Equal(t, `"quoted"`, projected[0].Name)
	require.Equal(t, []byte(`{"a":1,"a":2}`), projected[0].Raw)
	require.Equal(t, map[string]int{"count": 42}, projected[0].Data)
	require.Equal(t, []string{"label", "raw", "payload"}, rows.calls)

	var strings []string
	mock := sqlmock.NewRows([]string{"data"}).AddRow(`"value"`)
	require.NoError(t, ScanSlice(&jsonOptionsRows{ColumnScanner: toRows(mock)}, &strings))
	require.Equal(t, []string{`"value"`}, strings)
}

func TestScanSliceJSONOptionsWrappers(t *testing.T) {
	// Options must survive both of Ent's standard row wrappers.
	source := &jsonOptionsRows{
		ColumnScanner: toRows(sqlmock.NewRows([]string{"data"}).
			AddRow(`{"count":"42"}`).AddRow(`{"count":"43"}`)),
		options: func(string) []jsonv2.Options {
			return []jsonv2.Options{jsonv2.StringifyNumbers(true)}
		},
	}
	var closed bool
	closeErr := errors.New("close hook")
	rows := &Rows{ColumnScanner: rowsWithCloser{
		ColumnScanner: source,
		closer: func() error {
			closed = true
			return closeErr
		},
	}}
	var values []struct {
		Data struct {
			Count int `json:"count"`
		}
	}
	require.NoError(t, ScanSlice(rows, &values))
	require.Equal(t, 42, values[0].Data.Count)
	require.Equal(t, 43, values[1].Data.Count)
	require.Equal(t, []string{"data"}, source.calls, "read options once per result column")
	require.False(t, closed, "ScanSlice leaves closing rows to its caller")
	require.ErrorIs(t, rows.Close(), closeErr)
	require.True(t, closed)
}

func TestScanSliceWithoutJSONOptions(t *testing.T) {
	for _, tt := range []struct {
		name string
		wrap func(ColumnScanner) ColumnScanner
	}{
		{name: "plain", wrap: func(r ColumnScanner) ColumnScanner { return r }},
		{name: "Rows", wrap: func(r ColumnScanner) ColumnScanner { return &Rows{ColumnScanner: r} }},
		{name: "closer", wrap: func(r ColumnScanner) ColumnScanner {
			return rowsWithCloser{ColumnScanner: r, closer: func() error { return nil }}
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var values []struct{ Data struct{ Name string } }
			rows := tt.wrap(toRows(sqlmock.NewRows([]string{"data"}).AddRow(`{"NAME":"first","NAME":"last"}`)))
			defer rows.Close()
			require.NoError(t, ScanSlice(rows, &values))
			require.Equal(t, "last", values[0].Data.Name)
		})
	}
}

type jsonOptionsRows struct {
	ColumnScanner
	options func(string) []jsonv2.Options
	calls   []string
}

func (r *jsonOptionsRows) JSONOptions(column string) []jsonv2.Options {
	r.calls = append(r.calls, column)
	if r.options != nil {
		return r.options(column)
	}
	return nil
}
