// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package sql

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
)

// ColumnJSONOptions is an optional extension to ColumnScanner that supplies
// JSON decoding options for a result column by name. ScanSlice uses these options
// when the destination requires JSON decoding; directly scannable destinations
// retain their normal SQL scanning behavior.
//
// Options are passed directly to jsonv2.Unmarshal. Nil uses v2 defaults.
// Implementations should return encoding/json.DefaultOptionsV1 explicitly for
// columns requiring v1 compatibility, including unknown columns if appropriate.
// Scanners without this extension retain v1 compatibility.
// It requires the jsonv2 experiment, enabled by default in Go 1.27.
type ColumnJSONOptions interface {
	JSONOptions(column string) []jsonv2.Options
}

// JSONOptions forwards options from the underlying rows, falling back to v1
// compatibility when the underlying scanner does not implement ColumnJSONOptions.
func (r *Rows) JSONOptions(column string) []jsonv2.Options {
	return columnJSONOptions(r.ColumnScanner, column)
}

func (r rowsWithCloser) JSONOptions(column string) []jsonv2.Options {
	return columnJSONOptions(r.ColumnScanner, column)
}

func columnJSONOptions(rows ColumnScanner, column string) []jsonv2.Options {
	if r, ok := rows.(ColumnJSONOptions); ok {
		return r.JSONOptions(column)
	}
	return defaultJSONScanOptions()
}

func scanColumns(rows ColumnScanner, names []string) []scanColumn {
	provider, ok := rows.(ColumnJSONOptions)
	if !ok {
		return defaultScanColumns(names)
	}
	columns := make([]scanColumn, len(names))
	for i, name := range names {
		columns[i] = scanColumn{name: name, opts: provider.JSONOptions(name)}
	}
	return columns
}

type jsonScanOptions = []jsonv2.Options

func defaultJSONScanOptions() jsonScanOptions {
	return []jsonv2.Options{json.DefaultOptionsV1()}
}

func (c scanColumn) unmarshal(data []byte, v any) error {
	return jsonv2.Unmarshal(data, v, c.opts...)
}
