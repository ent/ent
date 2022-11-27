// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"ariga.io/atlas/sql/migrate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"github.com/stretchr/testify/require"
)

func TestWriteDriver(t *testing.T) {
	b := &bytes.Buffer{}
	w := NewWriteDriver(dialect.MySQL, b)
	ctx := context.Background()
	tx, err := w.Tx(ctx)
	require.NoError(t, err)
	err = tx.Query(ctx, "SELECT `name` FROM `users`", nil, nil)
	require.EqualError(t, err, "query is not supported by the WriteDriver")
	err = tx.Exec(ctx, "ALTER TABLE `users` ADD COLUMN `age` int", nil, nil)
	require.NoError(t, err)
	err = tx.Exec(ctx, "ALTER TABLE `users` ADD COLUMN `NAME` varchar(100);", nil, nil)
	require.NoError(t, err)
	require.NoError(t, tx.Commit())
	lines := strings.Split(b.String(), "\n")
	require.Len(t, lines, 3)
	require.Equal(t, "ALTER TABLE `users` ADD COLUMN `age` int;", lines[0])
	require.Equal(t, "ALTER TABLE `users` ADD COLUMN `NAME` varchar(100);", lines[1])
	require.Empty(t, lines[2], "file ends with blank line")

	b.Reset()
	query, args := sql.Update("users").Schema("test").Set("a", 1).Set("b", "a").Set("c", "'c'").Set("d", true).Where(sql.EQ("p", 0.2)).Query()
	err = w.Exec(ctx, query, args, nil)
	require.NoError(t, err)
	require.Equal(t, "UPDATE `test`.`users` SET `a` = 1, `b` = 'a', `c` = '''c''', `d` = 1 WHERE `p` = 0.2;\n", b.String())

	b.Reset()
	query, args = sql.Dialect(dialect.MySQL).Update("users").Schema("test").Set("a", "{}").Where(sqljson.ValueIsNull("a")).Query()
	err = w.Exec(ctx, query, args, nil)
	require.NoError(t, err)
	require.Equal(t, "UPDATE `test`.`users` SET `a` = '{}' WHERE JSON_CONTAINS(`a`, 'null', '$');\n", b.String())

	b.Reset()
	w = NewWriteDriver(dialect.Postgres, b)
	query, args = sql.Dialect(dialect.Postgres).Update("users").Set("a", 1).Set("b", time.Now()).Query()
	err = w.Exec(ctx, query, args, nil)
	require.NoError(t, err)
	require.Equal(t, `UPDATE "users" SET "a" = 1, "b" = {{ TIME_VALUE }};`+"\n", b.String())

	b.Reset()
	err = w.Exec(ctx, `INSERT INTO "users" (name) VALUES("a8m") RETURNING id`, nil, nil)
	require.NoError(t, err)
	require.Equal(t, `INSERT INTO "users" (name) VALUES("a8m") RETURNING id;`+"\n", b.String())
}

func TestDirWriter(t *testing.T) {
	p := t.TempDir()
	dir, err := migrate.NewLocalDir(p)
	require.NoError(t, err)
	w := &DirWriter{Dir: dir}
	drv := NewWriteDriver(dialect.MySQL, w)
	require.NoError(t, drv.Exec(context.Background(), "UPDATE `test`.`users` SET `a` = ?", []any{1}, nil))
	w.Change("Comment 1.")
	require.NoError(t, drv.Exec(context.Background(), "UPDATE `test`.`users` SET `b` = ?", []any{2}, nil))
	w.Change("Comment 2.")
	require.NoError(t, w.Flush("migration_file"))
	files, err := os.ReadDir(p)
	require.NoError(t, err)
	require.Len(t, files, 2)
	require.Contains(t, files[0].Name(), "_migration_file.sql")
	buf, err := os.ReadFile(filepath.Join(p, files[0].Name()))
	require.NoError(t, err)
	require.Equal(t, "-- Comment 1.\nUPDATE `test`.`users` SET `a` = 1;\n-- Comment 2.\nUPDATE `test`.`users` SET `b` = 2;\n", string(buf))
	require.Equal(t, "atlas.sum", files[1].Name())
}
