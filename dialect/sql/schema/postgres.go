// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Postgres is a postgres migration driver.
type Postgres struct {
	dialect.Driver
	version string
}

// init loads the Postgres version from the database for later use in the migration process.
// It returns an error if the server version is lower than v10.
func (d *Postgres) init(ctx context.Context, tx dialect.Tx) error {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW server_version_num", []interface{}{}, rows); err != nil {
		return fmt.Errorf("postgres: querying server version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("postgres: server_version_num variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("postgres: scanning version: %v", err)
	}
	if len(version) < 6 {
		return fmt.Errorf("postgres: malformed version: %s", version)
	}
	d.version = fmt.Sprintf("%s.%s.%s", version[:2], version[2:4], version[4:])
	if compareVersions(d.version, "10.0.0") == -1 {
		return fmt.Errorf("postgres: unsupported version: %s", d.version)
	}
	return nil
}

// tableExist checks if a table exists in the database and current schema.
func (d *Postgres) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("table_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

// tableExist checks if a foreign-key exists in the current schema.
func (d *Postgres) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Dialect(dialect.Postgres).
		Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("constraint_type", "FOREIGN KEY").And().EQ("constraint_name", name)).Query()
	return exist(ctx, tx, query, args...)
}

// setRange sets restart the identity column to the given offset. Used by the universal-id option.
func (d *Postgres) setRange(ctx context.Context, tx dialect.Tx, name string, value int) error {
	return tx.Exec(ctx, fmt.Sprintf("ALTER TABLE %s ALTER COLUMN id RESTART WITH %d", name, value), []interface{}{}, new(sql.Result))
}
