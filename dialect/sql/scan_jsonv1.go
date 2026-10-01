// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build !goexperiment.jsonv2

package sql

import "encoding/json"

// jsonScanOptions uses v1 semantics for all columns when the jsonv2 experiment is disabled.
type jsonScanOptions struct{}

func scanColumns(_ ColumnScanner, names []string) []scanColumn {
	return defaultScanColumns(names)
}

func defaultJSONScanOptions() jsonScanOptions {
	return jsonScanOptions{}
}

func (scanColumn) unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
