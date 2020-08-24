// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/facebook/ent/dialect"

	"github.com/stretchr/testify/require"
)

func TestWriteDriver(t *testing.T) {
	b := &bytes.Buffer{}
	w := WriteDriver{Driver: nopDriver{}, Writer: b}
	ctx := context.Background()
	tx, err := w.Tx(ctx)
	require.NoError(t, err)
	err = tx.Query(ctx, "SELECT `name` FROM `users`", nil, nil)
	require.NoError(t, err)
	err = tx.Query(ctx, "SELECT `name` FROM `users`", nil, nil)
	require.NoError(t, err)
	err = tx.Exec(ctx, "ALTER TABLE `users` ADD COLUMN `age` int", nil, nil)
	require.NoError(t, err)
	err = tx.Exec(ctx, "ALTER TABLE `users` ADD COLUMN `NAME` varchar(100);", nil, nil)
	require.NoError(t, err)
	err = tx.Query(ctx, "SELECT `name` FROM `users`", nil, nil)
	require.NoError(t, err)
	require.NoError(t, tx.Commit())
	lines := strings.Split(b.String(), "\n")
	require.Equal(t, "BEGIN;", lines[0])
	require.Equal(t, "ALTER TABLE `users` ADD COLUMN `age` int;", lines[1])
	require.Equal(t, "ALTER TABLE `users` ADD COLUMN `NAME` varchar(100);", lines[2])
	require.Equal(t, "COMMIT;", lines[3])
	require.Empty(t, lines[4], "file ends with blank line")
}

type nopDriver struct {
	dialect.Driver
}

func (nopDriver) Exec(context.Context, string, interface{}, interface{}) error {
	return nil
}

func (nopDriver) Query(context.Context, string, interface{}, interface{}) error {
	return nil
}
