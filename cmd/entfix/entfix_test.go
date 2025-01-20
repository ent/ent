package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestEntfix(t *testing.T) {
	ctx := context.Background()
	stdout, err := os.CreateTemp(t.TempDir(), "")
	require.NoError(t, err)
	oldStdout := os.Stdout
	os.Stdout = stdout
	t.Cleanup(func() {
		os.Stdout = oldStdout
		require.NoError(t, stdout.Close())
	})

	cli := &GlobalID{
		Dialect: dialect.SQLite,
		DSN:     fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name()),
		Path:    t.TempDir(),
	}

	// Prints details, requires confirmation.
	os.Stdin = mockStdin(t, "no\n")
	require.NoError(t, cli.Run(ctx))
	buf, err := os.ReadFile(stdout.Name())
	require.NoError(t, err)
	require.True(t, strings.HasPrefix(string(buf), "IMPORTANT INFORMATION\n\n"))
	require.True(t, strings.HasSuffix(string(buf), "Aborted.\n"))
	f, err := os.Open(cli.Path)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, f.Close()) })
	_, err = f.Readdirnames(1)
	require.ErrorIs(t, err, io.EOF)

	// If approved, converts the 'ent_types' table.
	db, err := sql.Open(cli.Dialect, cli.DSN)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, db.Close()) })
	drv := entsql.OpenDB(dialect.SQLite, db)
	m, err := schema.NewMigrate(drv, schema.WithGlobalUniqueID(true))
	require.NoError(t, err)
	require.NoError(t, m.Create(ctx))
	_, err = db.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (type) VALUES ('z'), ('y'), ('a'), ('b')", schema.TypeTable))
	require.NoError(t, err)

	os.Stdin = mockStdin(t, "yes\n")
	require.NoError(t, cli.Run(ctx))
	buf, err = os.ReadFile(stdout.Name())
	require.NoError(t, err)
	require.True(t, strings.HasSuffix(string(buf), "Success! Please run code generation to complete the process.\n"))
	c, err := os.ReadFile(filepath.Join(cli.Path, "internal", "globalid.go"))
	require.NoError(t, err)
	require.Contains(t,
		string(c),
		fmt.Sprintf(`const IncrementStarts = "{\"a\":%d,\"b\":%d,\"y\":%d,\"z\":%d}"`, 2<<32, 3<<32, 1<<32, 0),
	)
}

func mockStdin(t *testing.T, content string) *os.File {
	t.Helper()
	stdin, err := os.CreateTemp(t.TempDir(), "")
	require.NoError(t, err)
	_, err = stdin.WriteString(content)
	require.NoError(t, err)
	_, err = stdin.Seek(0, 0)
	require.NoError(t, err)
	oldStdin := os.Stdin
	t.Cleanup(func() {
		os.Stdin = oldStdin
		require.NoError(t, stdin.Close())
	})
	return stdin
}
