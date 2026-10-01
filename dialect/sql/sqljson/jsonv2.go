// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package sqljson

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

// AppendWithOptions is like Append, but encodes the elements using the supplied
// JSON options. Nil options use encoding/json/v2 defaults. It requires the
// jsonv2 experiment, enabled by default in Go 1.27.
func AppendWithOptions[T any](u *sql.UpdateBuilder, column string, elems []T, opts []jsonv2.Options, pathOpts ...Option) {
	values := make([]json.RawMessage, len(elems))
	for i := range elems {
		buf, err := jsonv2.Marshal(elems[i], opts...)
		if err != nil {
			u.AddError(fmt.Errorf("sqljson: marshal value for column %s: %w", column, err))
			return
		}
		values[i] = buf
	}
	Append(u, column, values, pathOpts...)
}
