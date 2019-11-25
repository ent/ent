// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect"
)

// Driver is a dialect.Driver implementation for SQL based databases.
type Driver struct {
	conn
	dialect string
}

// Open wraps the database/sql.Open method and returns a dialect.Driver that implements the an ent/dialect.Driver interface.
func Open(driver, source string) (*Driver, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	return &Driver{conn{db}, driver}, nil
}

// OpenDB wraps the given database/sql.DB method with a Driver.
func OpenDB(driver string, db *sql.DB) *Driver {
	return &Driver{conn{db}, driver}
}

// DB returns the underlying *sql.DB instance.
func (d Driver) DB() *sql.DB {
	return d.conn.ExecQuerier.(*sql.DB)
}

// Dialect implements the dialect.Dialect method.
func (d Driver) Dialect() string {
	// if the underlying driver is wrapped with opencensus driver.
	for _, name := range []string{dialect.MySQL, dialect.SQLite} {
		if strings.HasPrefix(d.dialect, name) {
			return name
		}
	}
	return d.dialect
}

// Tx starts and returns a transaction.
func (d *Driver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.BeginTx(ctx, &sql.TxOptions{})
}

// BeginTx starts a transaction with options.
func (d *Driver) BeginTx(ctx context.Context, opts *TxOptions) (dialect.Tx, error) {
	tx, err := d.ExecQuerier.(*sql.DB).BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx{conn{tx}}, nil
}

// Close closes the underlying connection.
func (d *Driver) Close() error { return d.ExecQuerier.(*sql.DB).Close() }

// Tx wraps the sql.Tx for implementing the dialect.Tx interface.
type Tx struct {
	conn
}

// Commit commits the transaction.
func (t *Tx) Commit() error { return t.ExecQuerier.(*sql.Tx).Commit() }

// Rollback rollback the transaction.
func (t *Tx) Rollback() error { return t.ExecQuerier.(*sql.Tx).Rollback() }

// ExecQuerier wraps the standard Exec and Query methods.
type ExecQuerier interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// shared connection ExecQuerier between Driver and Tx.
type conn struct {
	ExecQuerier
}

// Exec implements the dialect.Exec method.
func (c *conn) Exec(ctx context.Context, query string, args, v interface{}) error {
	vr, ok := v.(*sql.Result)
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect *sql.Result", v)
	}
	argv, ok := args.([]interface{})
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect []interface{} for args", v)
	}
	res, err := c.ExecContext(ctx, query, argv...)
	if err != nil {
		return err
	}
	*vr = res
	return nil
}

// Exec implements the dialect.Query method.
func (c *conn) Query(ctx context.Context, query string, args, v interface{}) error {
	vr, ok := v.(*Rows)
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect *sql.Rows", v)
	}
	argv, ok := args.([]interface{})
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect []interface{} for args", args)
	}
	rows, err := c.QueryContext(ctx, query, argv...)
	if err != nil {
		return err
	}
	*vr = Rows{rows}
	return nil
}

var _ dialect.Driver = (*Driver)(nil)

type (
	// Rows wraps the sql.Rows to avoid locks copy.
	Rows struct{ *sql.Rows }
	// Result is an alias to sql.Result.
	Result = sql.Result
	// NullBool is an alias to sql.NullBool.
	NullBool = sql.NullBool
	// NullInt64 is an alias to sql.NullInt64.
	NullInt64 = sql.NullInt64
	// NullString is an alias to sql.NullString.
	NullString = sql.NullString
	// NullFloat64 is an alias to sql.NullFloat64.
	NullFloat64 = sql.NullFloat64
	// TxOptions holds the transaction options to be used in DB.BeginTx.
	TxOptions = sql.TxOptions
)
