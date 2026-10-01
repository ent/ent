// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package sqljson_test

import (
	"encoding/json"
	"encoding/json/jsontext"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/stretchr/testify/require"
)

func TestAppendWithOptions(t *testing.T) {
	for _, d := range []string{dialect.SQLite, dialect.MySQL, dialect.Postgres} {
		t.Run(d, func(t *testing.T) {
			u := sql.Dialect(d).Update("t")
			sqljson.AppendWithOptions(u, "values", []any{[]string(nil), map[string]string(nil), true, nil, streamingValue{}}, nil)
			query, args := u.Query()
			require.NoError(t, u.Err())
			switch d {
			case dialect.Postgres:
				require.Equal(t, []any{`[[],{},true,null,"v2"]`, `[[],{},true,null,"v2"]`}, args)
			case dialect.SQLite:
				require.Equal(t, []any{`[[],{},true,null,"v2"]`, `[]`, `{}`, `true`, `null`, `"v2"`}, args)
			case dialect.MySQL:
				require.Contains(t, query, "THEN JSON_ARRAY(CAST(? AS JSON), CAST(? AS JSON), CAST(? AS JSON), CAST(? AS JSON), CAST(? AS JSON))")
				require.Equal(t, []any{`[]`, `{}`, `true`, `null`, `"v2"`, `[]`, `{}`, `true`, `null`, `"v2"`}, args)
			}
			for _, elems := range [][]any{nil, {"\xff"}, {json.RawMessage(`{"a":1,"a":2}`)}} {
				u := sql.Dialect(d).Update("t")
				sqljson.AppendWithOptions(u, "values", elems, nil)
				require.Error(t, u.Err())
			}
		})
	}
}

type streamingValue struct{}

func (streamingValue) MarshalJSONTo(enc *jsontext.Encoder) error {
	return enc.WriteToken(jsontext.String("v2"))
}
