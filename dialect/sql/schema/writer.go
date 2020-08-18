// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"io"
	"strings"

	"github.com/facebook/ent/dialect"
)

// WriteDriver is a driver that writes all driver exec operations to its writer.
type WriteDriver struct {
	dialect.Driver // underlying driver.
	io.Writer      // target for exec statements.
}

// Exec writes its query and calls the underlying driver Exec method.
func (w *WriteDriver) Exec(_ context.Context, query string, _, _ interface{}) error {
	if !strings.HasSuffix(query, ";") {
		query += ";"
	}
	_, err := io.WriteString(w, query+"\n")
	return err
}

// Tx writes the transaction start.
func (w *WriteDriver) Tx(context.Context) (dialect.Tx, error) {
	if _, err := io.WriteString(w, "BEGIN;\n"); err != nil {
		return nil, err
	}
	return w, nil
}

// Commit writes the transaction commit.
func (w *WriteDriver) Commit() error {
	_, err := io.WriteString(w, "COMMIT;\n")
	return err
}

// Rollback writes the transaction rollback.
func (w *WriteDriver) Rollback() error {
	_, err := io.WriteString(w, "ROLLBACK;\n")
	return err
}
