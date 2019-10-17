// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
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

// table loads the current table description from the database.
func (d *Postgres) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(dialect.Postgres).
		Select("column_name", "data_type", "character_maximum_length", "is_nullable", "column_default").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.EQ("table_schema", sql.Raw("CURRENT_SCHEMA()")).And().EQ("table_name", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("postgres: reading table description %v", err)
	}
	// call `Close` in cases of failures (`Close` is idempotent).
	defer rows.Close()
	t := NewTable(name)
	for rows.Next() {
		c := &Column{}
		if err := d.scanColumn(c, rows); err != nil {
			return nil, fmt.Errorf("postgres: %v", err)
		}
		t.AddColumn(c)
	}
	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("postgres: closing rows %v", err)
	}
	// TODO: populate PK/UNI information for columns and tables and scan indexes.
	//
	// Get PK and UNI columns of a table:
	//
	//	SELECT a.attname                           AS column,
	//		format_type(a.atttypid, a.atttypmod)   AS data_type,
	//		i.indisprimary                         AS primary,
	//		i.indisunique                          AS unique
	//	FROM pg_index i
	//	join pg_attribute a ON a.attrelid = i.indrelid AND a.attnum = ANY (i.indkey)
	//	join pg_stat_user_tables t ON t.relid = i.indrelid
	//	WHERE t.schemaname = CURRENT_SCHEMA()
	//		AND i.indrelid = '<TABLE>' :: regclass;
	//
	//   column | data_type | primary | unique
	//	--------+-----------+---------+--------
	//	 a1     | integer   | t       | t
	//	 a2     | integer   | t       | t
	// 	 a0     | integer   | f       | t
	//
	return t, nil
}

// maxCharSize defines the maximum size of limited character types in Postgres (10 MB).
const maxCharSize = 10 << 20

// scanColumn scans the information a column from column description.
func (d *Postgres) scanColumn(c *Column, rows *sql.Rows) error {
	var (
		maxlen   sql.NullInt64
		nullable sql.NullString
		defaults sql.NullString
	)
	if err := rows.Scan(&c.Name, &c.typ, &maxlen, &nullable, &defaults); err != nil {
		return fmt.Errorf("scanning column description: %v", err)
	}
	if nullable.Valid {
		c.Nullable = nullable.String == "YES"
	}
	switch c.typ {
	case "boolean":
		c.Type = field.TypeBool
	case "smallint":
		c.Type = field.TypeInt16
	case "integer":
		c.Type = field.TypeInt32
	case "bigint":
		c.Type = field.TypeInt64
	case "real":
		c.Type = field.TypeFloat32
	case "double precision":
		c.Type = field.TypeFloat64
	case "text":
		c.Type = field.TypeString
		c.Size = maxCharSize + 1
	case "character":
		c.Type = field.TypeString
		c.Size = maxlen.Int64
	case "timestamp with time zone":
		c.Type = field.TypeTime
	case "bytea":
		c.Type = field.TypeBytes
	case "jsonb":
		c.Type = field.TypeJSON
	}
	switch {
	case !defaults.Valid:
		return nil
	case strings.Contains(defaults.String, "::"):
		parts := strings.Split(defaults.String, "::")
		defaults.String = strings.Trim(parts[0], "'")
		fallthrough
	default:
		return c.ScanDefault(defaults.String)
	}
}
