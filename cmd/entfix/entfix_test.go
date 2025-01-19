package main

import (
	"bytes"
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
	var (
		ctx = context.Background()
		buf = new(bytes.Buffer)
	)
	out = buf

	cli := &GlobalID{
		Dialect: dialect.SQLite,
		DSN:     fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", t.Name()),
		Path:    t.TempDir(),
	}

	// Prints details, requires confirmation.
	in = strings.NewReader("no\n")
	require.NoError(t, cli.Run(ctx))
	require.True(t, strings.HasPrefix(buf.String(), "IMPORTANT INFORMATION\n\n"))
	require.True(t, strings.HasSuffix(buf.String(), "Aborted.\n"))
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

	in = strings.NewReader("yes\n")
	require.NoError(t, cli.Run(ctx))
	require.True(t, strings.HasSuffix(buf.String(), "Success! Please run code generation to complete the process.\n"))
	c, err := os.ReadFile(filepath.Join(cli.Path, "internal", "globalid.go"))
	require.NoError(t, err)
	require.Contains(t,
		string(c),
		fmt.Sprintf(`const IncrementStarts = "{\"a\":%d,\"b\":%d,\"y\":%d,\"z\":%d}"`, 2<<32, 3<<32, 1<<32, 0),
	)
}
