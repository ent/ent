// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sessionvars

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"strings"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestPostgresWithVarNames(t *testing.T) {
	drv := openPostgres(t)
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
		{"plain-search-path", "search_path", "public", "SELECT current_setting('search_path')", "public"},
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

func TestPostgresVarCleanup(t *testing.T) {
	for _, tt := range []struct{ name, key string }{
		{"app.user", "app.user"},
		{"app." + strings.Repeat("x", 100), "app." + strings.Repeat("x", 63)},
	} {
		name, key := tt.name, tt.key
		for _, operation := range []string{"query", "exec"} {
			t.Run(key+"/"+operation, func(t *testing.T) {
				drv := openPostgres(t)
				ctx := context.Background()
				var before int
				require.NoError(t, drv.DB().QueryRow("SELECT pg_backend_pid()").Scan(&before))
				if operation == "query" {
					var rows sql.Rows
					require.NoError(t, drv.Query(sql.WithVar(ctx, name, "alice"), "SELECT current_setting($1)", []any{key}, &rows))
					require.True(t, rows.Next())
					var value string
					require.NoError(t, rows.Scan(&value))
					require.Equal(t, "alice", value)
					require.NoError(t, rows.Close())
				} else {
					_, err := drv.DB().Exec("CREATE TEMP TABLE effects(id int)")
					require.NoError(t, err)
					require.NoError(t, drv.Exec(sql.WithVar(ctx, name, "alice"), "INSERT INTO effects VALUES (1)", []any{}, nil))
					var count int
					require.NoError(t, drv.DB().QueryRow("SELECT count(*) FROM effects").Scan(&count))
					require.Equal(t, 1, count)
				}
				var after int
				var value string
				require.NoError(t, drv.DB().QueryRow("SELECT pg_backend_pid(), current_setting($1)", key).Scan(&after, &value))
				require.Equal(t, before, after)
				require.Empty(t, value)
			})
		}
	}
}

func TestPostgresTransactionVars(t *testing.T) {
	for _, tt := range []struct{ key, value string }{
		{"transaction_isolation", "serializable"},
		{"transaction_read_only", "off"},
		{"transaction_deferrable", "on"},
	} {
		for _, commit := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/commit=%t", tt.key, commit), func(t *testing.T) {
				drv := openPostgres(t)
				var original string
				require.NoError(t, drv.DB().QueryRow("SELECT current_setting($1)", tt.key).Scan(&original))
				tx, err := drv.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: tt.key == "transaction_read_only"})
				require.NoError(t, err)
				t.Cleanup(func() { tx.Rollback() })
				ctx := sql.WithVar(context.Background(), "app.actor", "alice")
				ctx = sql.WithVar(ctx, tt.key, tt.value)
				var rows sql.Rows
				require.NoError(t, tx.Query(ctx, "SELECT current_setting($1)", []any{tt.key}, &rows))
				require.True(t, rows.Next())
				var value string
				require.NoError(t, rows.Scan(&value))
				require.Equal(t, tt.value, value)
				require.NoError(t, rows.Close())
				if tt.key != "transaction_deferrable" {
					// Reapplying the same context must remain valid after a query has run.
					require.NoError(t, tx.Query(ctx, "SELECT current_setting($1)", []any{tt.key}, &rows))
					require.True(t, rows.Next())
					require.NoError(t, rows.Scan(&value))
					require.Equal(t, tt.value, value)
					require.NoError(t, rows.Close())
				}
				if commit {
					require.NoError(t, tx.Commit())
				} else {
					require.NoError(t, tx.Rollback())
				}
				require.NoError(t, drv.DB().QueryRow("SELECT current_setting($1)", tt.key).Scan(&value))
				require.Equal(t, original, value)
			})
		}
	}
}

func openPostgres(t *testing.T) *sql.Driver {
	t.Helper()
	drv, err := sql.Open("postgres", "host=localhost port=5437 user=postgres dbname=test password=pass sslmode=disable")
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, drv.Close()) })
	drv.DB().SetMaxOpenConns(1)
	return drv
}

func TestPostgresBuiltinConfig(t *testing.T) {
	drv := openPostgres(t)
	db := drv.DB()
	_, err := db.Exec(`CREATE SCHEMA ent_vars_untrusted`)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := db.Exec(`RESET search_path; DROP SCHEMA ent_vars_untrusted CASCADE`)
		require.NoError(t, err)
	})
	_, err = db.Exec(`
		CREATE FUNCTION ent_vars_untrusted.set_config(text, text, text) RETURNS text
		LANGUAGE plpgsql AS $$ BEGIN RAISE EXCEPTION 'untrusted set_config called'; END $$;
		CREATE FUNCTION ent_vars_untrusted.set_config(text, text, boolean) RETURNS text
		LANGUAGE plpgsql AS $$ BEGIN RAISE EXCEPTION 'untrusted set_config called'; END $$;
		SET search_path = ent_vars_untrusted, pg_catalog;
	`)
	require.NoError(t, err)
	ctx := sql.WithVar(context.Background(), "app.current_tenant", "expected")
	var rows sql.Rows
	require.NoError(t, drv.Query(ctx, "SELECT pg_catalog.current_setting($1)", []any{"app.current_tenant"}, &rows))
	closed := false
	t.Cleanup(func() {
		if !closed {
			require.NoError(t, rows.Close())
		}
	})
	require.True(t, rows.Next())
	var tenant string
	require.NoError(t, rows.Scan(&tenant))
	require.Equal(t, "expected", tenant)
	require.NoError(t, rows.Close())
	closed = true
	require.NoError(t, db.QueryRow("SELECT pg_catalog.current_setting($1)", "app.current_tenant").Scan(&tenant))
	require.Empty(t, tenant)
}

func TestPostgresVarPrivileges(t *testing.T) {
	drv := openPostgres(t)
	_, err := drv.DB().Exec("CREATE ROLE ent_sessionvars_setter; GRANT SET ON PARAMETER session_preload_libraries TO ent_sessionvars_setter")
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := drv.DB().Exec("REVOKE SET ON PARAMETER session_preload_libraries FROM ent_sessionvars_setter; DROP ROLE ent_sessionvars_setter")
		require.NoError(t, err)
	})
	for _, transaction := range []bool{false, true} {
		t.Run(fmt.Sprintf("transaction=%t", transaction), func(t *testing.T) {
			conn := openPostgres(t)
			_, err := conn.DB().Exec("SET ROLE ent_sessionvars_setter")
			require.NoError(t, err)
			ctx := sql.WithVar(context.Background(), "session_preload_libraries", "")
			if transaction {
				tx, err := conn.Tx(ctx)
				require.NoError(t, err)
				t.Cleanup(func() { _ = tx.Rollback() })
				require.NoError(t, tx.Exec(ctx, "SELECT 1", []any{}, nil))
				require.NoError(t, tx.Commit())
			} else {
				require.NoError(t, conn.Exec(ctx, "SELECT 1", []any{}, nil))
			}
		})
	}
}

func TestPostgresVarSnapshot(t *testing.T) {
	drv := openPostgres(t)
	writer := openPostgres(t)
	_, err := drv.DB().Exec("CREATE TABLE ent_sessionvars_snapshot (value int); INSERT INTO ent_sessionvars_snapshot VALUES (1)")
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := drv.DB().Exec("DROP TABLE ent_sessionvars_snapshot")
		require.NoError(t, err)
	})
	ctx := sql.WithVar(context.Background(), "app.actor", "alice")
	tx, err := drv.BeginTx(ctx, &sql.TxOptions{Isolation: stdsql.LevelRepeatableRead})
	require.NoError(t, err)
	t.Cleanup(func() { _ = tx.Rollback() })
	require.NoError(t, tx.Exec(ctx, "SET LOCAL application_name = 'snapshot-probe'", []any{}, nil))
	_, err = writer.DB().Exec("UPDATE ent_sessionvars_snapshot SET value=2")
	require.NoError(t, err)
	var rows sql.Rows
	require.NoError(t, tx.Query(ctx, "SELECT value FROM ent_sessionvars_snapshot", []any{}, &rows))
	require.True(t, rows.Next())
	var value int
	require.NoError(t, rows.Scan(&value))
	require.NoError(t, rows.Close())
	require.Equal(t, 2, value)
	require.NoError(t, tx.Rollback())
}

func TestPostgresVarValues(t *testing.T) {
	const value = "O'Reilly\\team; SELECT 1 --\n\"quoted\" 雪"
	for _, mode := range []string{"on", "off"} {
		t.Run(mode, func(t *testing.T) {
			drv := openPostgres(t)
			_, err := drv.DB().Exec("SET standard_conforming_strings = " + mode)
			require.NoError(t, err)
			var rows sql.Rows
			require.NoError(t, drv.Query(sql.WithVar(context.Background(), "app.actor", value), "SELECT current_setting('app.actor')", []any{}, &rows))
			require.True(t, rows.Next())
			var got string
			require.NoError(t, rows.Scan(&got))
			require.NoError(t, rows.Close())
			require.Equal(t, value, got)
		})
	}
}
