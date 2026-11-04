// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sessionvars

import (
	"context"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestPostgresWithVarNames(t *testing.T) {
	drv, err := sql.Open("postgres", "host=localhost port=5437 user=postgres dbname=test password=pass sslmode=disable")
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, drv.Close()) })
	drv.DB().SetMaxOpenConns(1)
	ctx := context.Background()
	for _, name := range []string{`"EntTenantA"`, `"Ent,Comma"`, `"Ent""Quote"`} {
		_, err := drv.DB().ExecContext(ctx, "CREATE SCHEMA "+name)
		require.NoError(t, err)
		t.Cleanup(func() {
			_, err := drv.DB().ExecContext(ctx, "DROP SCHEMA "+name)
			require.NoError(t, err)
		})
	}
	var original string
	require.NoError(t, drv.DB().QueryRowContext(ctx, "SHOW search_path").Scan(&original))
	for _, tt := range []struct {
		name, key, value, query, want string
	}{
		{"mixed-case", "search_path", "EntTenantA", "SELECT current_schema()", "EntTenantA"},
		{"comma", "search_path", "Ent,Comma", "SELECT current_schema()", "Ent,Comma"},
		{"quote", "search_path", `Ent"Quote`, "SELECT current_schema()", `Ent"Quote`},
		{"quoted-search-path", `"SEARCH_PATH"`, "EntTenantA", "SELECT current_schema()", "EntTenantA"},
		{"quoted-custom-name", ` "app" . "user" `, "O'Reilly\\team; SELECT 1 --", "SELECT current_setting('app.user')", "O'Reilly\\team; SELECT 1 --"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var rows sql.Rows
			require.NoError(t, drv.Query(sql.WithVar(ctx, tt.key, tt.value), tt.query, []any{}, &rows))
			closed := false
			t.Cleanup(func() {
				if !closed {
					require.NoError(t, rows.Close())
				}
			})
			require.True(t, rows.Next())
			var got sql.NullString
			require.NoError(t, rows.Scan(&got))
			require.True(t, got.Valid)
			require.Equal(t, tt.want, got.String)
			require.NoError(t, rows.Close())
			closed = true
			var restored string
			require.NoError(t, drv.DB().QueryRowContext(ctx, "SHOW search_path").Scan(&restored))
			require.Equal(t, original, restored)
		})
	}
}
