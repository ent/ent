// Copyright 2024-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
)

// YDB implements the dialect.Dialect interface for YDB.
type YDB struct {
	*Driver
}

// init registers the YDB driver with the SQL dialect.
func init() {
	Register(dialect.YDB, &YDB{})
}

// Placeholder returns the placeholder for the i'th argument in the query.
func (YDB) Placeholder() string {
	return "$"
}

// Array returns the placeholder for array values in the query.
func (YDB) Array() string {
	return "?"
}

// Drivers returns the list of supported drivers by YDB.
func (YDB) Drivers() []string {
	return []string{"ydb"}
}

// QueryBuilder returns the query builder for YDB.
func (d *YDB) QueryBuilder(b *Builder) *Builder {
	b.Quote = func(s string) string {
		return fmt.Sprintf("`%s`", s)
	}
	return b
}

// Schema returns the schema name of the database.
func (YDB) Schema() string {
	return "ydb"
}

// Version returns the version of the database.
func (d *YDB) Version(ctx context.Context) (*sql.DB, string, error) {
	if db, ok := d.DB().(*sql.DB); ok {
		var version string
		if err := db.QueryRowContext(ctx, "SELECT version()").Scan(&version); err != nil {
			return db, "", fmt.Errorf("ydb: failed getting version: %w", err)
		}
		return db, version, nil
	}
	return nil, "", fmt.Errorf("ydb: failed getting version: not a *sql.DB")
}

// YDBType represents a YDB data type.
type YDBType string

const (
	// YDBTypeInt8 represents YDB INT8 type.
	YDBTypeInt8 YDBType = "Int8"
	// YDBTypeInt16 represents YDB INT16 type.
	YDBTypeInt16 YDBType = "Int16"
	// YDBTypeInt32 represents YDB INT32 type.
	YDBTypeInt32 YDBType = "Int32"
	// YDBTypeInt64 represents YDB INT64 type.
	YDBTypeInt64 YDBType = "Int64"
	// YDBTypeUint8 represents YDB UINT8 type.
	YDBTypeUint8 YDBType = "Uint8"
	// YDBTypeUint16 represents YDB UINT16 type.
	YDBTypeUint16 YDBType = "Uint16"
	// YDBTypeUint32 represents YDB UINT32 type.
	YDBTypeUint32 YDBType = "Uint32"
	// YDBTypeUint64 represents YDB UINT64 type.
	YDBTypeUint64 YDBType = "Uint64"
	// YDBTypeFloat represents YDB FLOAT type.
	YDBTypeFloat YDBType = "Float"
	// YDBTypeDouble represents YDB DOUBLE type.
	YDBTypeDouble YDBType = "Double"
	// YDBTypeString represents YDB STRING type.
	YDBTypeString YDBType = "String"
	// YDBTypeBytes represents YDB BYTES type.
	YDBTypeBytes YDBType = "Bytes"
	// YDBTypeTimestamp represents YDB TIMESTAMP type.
	YDBTypeTimestamp YDBType = "Timestamp"
	// YDBTypeDate represents YDB DATE type.
	YDBTypeDate YDBType = "Date"
	// YDBTypeDateTime represents YDB DATETIME type.
	YDBTypeDateTime YDBType = "Datetime"
	// YDBTypeInterval represents YDB INTERVAL type.
	YDBTypeInterval YDBType = "Interval"
	// YDBTypeBool represents YDB BOOL type.
	YDBTypeBool YDBType = "Bool"
	// YDBTypeJSON represents YDB JSON type.
	YDBTypeJSON YDBType = "Json"
)

// ConvertType converts Go type to YDB type.
func (YDB) ConvertType(t interface{}) (YDBType, error) {
	switch t.(type) {
	case int8:
		return YDBTypeInt8, nil
	case int16:
		return YDBTypeInt16, nil
	case int32:
		return YDBTypeInt32, nil
	case int64:
		return YDBTypeInt64, nil
	case uint8:
		return YDBTypeUint8, nil
	case uint16:
		return YDBTypeUint16, nil
	case uint32:
		return YDBTypeUint32, nil
	case uint64:
		return YDBTypeUint64, nil
	case float32:
		return YDBTypeFloat, nil
	case float64:
		return YDBTypeDouble, nil
	case string:
		return YDBTypeString, nil
	case []byte:
		return YDBTypeBytes, nil
	case bool:
		return YDBTypeBool, nil
	default:
		return "", fmt.Errorf("unsupported type: %T", t)
	}
}

// FormatError formats YDB error messages.
func (d *YDB) FormatError(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if strings.Contains(msg, "not found") {
		return &NotFoundError{err}
	}
	return &Error{err}
}

// YDBDriver wraps the standard SQL driver to add YDB-specific functionality.
type YDBDriver struct {
	*Driver
}

// Open opens a new YDB connection.
func Open(driverName, dataSourceName string) (*YDBDriver, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	drv := &Driver{DB: db, Dialect: &YDB{}}
	return &YDBDriver{drv}, nil
}

// ExecContext executes a query that doesn't return rows.
// For example, in SQL, INSERT or UPDATE.
func (d *YDBDriver) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// YDB-specific query modifications if needed
	query = d.modifyQuery(query)
	return d.Driver.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
func (d *YDBDriver) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// YDB-specific query modifications if needed
	query = d.modifyQuery(query)
	return d.Driver.QueryContext(ctx, query, args...)
}

// modifyQuery modifies the SQL query to be compatible with YDB syntax.
func (d *YDBDriver) modifyQuery(query string) string {
	// Replace standard SQL syntax with YDB-specific syntax where needed
	// For example, replace LIMIT with TOP, adjust JOIN syntax, etc.
	query = strings.Replace(query, "LIMIT", "TOP", -1)
	return query
}

// BeginTx starts a new transaction with the given options.
func (d *YDBDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	if opts == nil {
		opts = &sql.TxOptions{
			Isolation: sql.LevelSerializable,
			ReadOnly:  false,
		}
	}
	return d.Driver.BeginTx(ctx, opts)
}
