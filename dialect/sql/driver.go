// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
)

// Driver is a dialect.Driver implementation for SQL based databases.
type Driver struct {
	Conn
	dialect string
}

// NewDriver creates a new Driver with the given Conn and dialect.
func NewDriver(dialect string, c Conn) *Driver {
	drv := &Driver{dialect: dialect, Conn: c}
	drv.Conn.dialect = drv.Dialect()
	return drv
}

// Open wraps the database/sql.Open method and returns a dialect.Driver that implements the an ent/dialect.Driver interface.
func Open(dialect, source string) (*Driver, error) {
	db, err := sql.Open(dialect, source)
	if err != nil {
		return nil, err
	}
	return NewDriver(dialect, Conn{db, dialect}), nil
}

// OpenDB wraps the given database/sql.DB method with a Driver.
func OpenDB(dialect string, db *sql.DB) *Driver {
	return NewDriver(dialect, Conn{db, dialect})
}

// DB returns the underlying *sql.DB instance.
func (d Driver) DB() *sql.DB {
	return d.ExecQuerier.(*sql.DB)
}

// Dialect implements the dialect.Dialect method.
func (d Driver) Dialect() string {
	// If the underlying driver is wrapped with a telemetry driver.
	for _, name := range []string{dialect.MySQL, dialect.SQLite, dialect.Postgres} {
		if strings.HasPrefix(d.dialect, name) {
			return name
		}
	}
	return d.dialect
}

// Tx starts and returns a transaction.
func (d *Driver) Tx(ctx context.Context) (dialect.Tx, error) {
	return d.BeginTx(ctx, nil)
}

// BeginTx starts a transaction with options.
func (d *Driver) BeginTx(ctx context.Context, opts *TxOptions) (dialect.Tx, error) {
	tx, err := d.DB().BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Tx{
		Conn: Conn{tx, d.Dialect()},
		Tx:   tx,
	}, nil
}

// Close closes the underlying connection.
func (d *Driver) Close() error { return d.DB().Close() }

// Tx implements dialect.Tx interface.
type Tx struct {
	Conn
	driver.Tx
}

// ExecQuerier wraps the standard Exec and Query methods.
type ExecQuerier interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

// Conn implements dialect.ExecQuerier given ExecQuerier.
type Conn struct {
	ExecQuerier
	dialect string
}

// Exec implements the dialect.Exec method.
func (c Conn) Exec(ctx context.Context, query string, args, v any) (rerr error) {
	argv, ok := args.([]any)
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect []any for args", v)
	}
	ex, cf, err := c.maySetVars(ctx)
	if err != nil {
		return err
	}
	if cf != nil {
		defer func() { rerr = errors.Join(rerr, cf()) }()
	}
	switch v := v.(type) {
	case nil:
		if _, err := ex.ExecContext(ctx, query, argv...); err != nil {
			return err
		}
	case *sql.Result:
		res, err := ex.ExecContext(ctx, query, argv...)
		if err != nil {
			return err
		}
		*v = res
	default:
		return fmt.Errorf("dialect/sql: invalid type %T. expect *sql.Result", v)
	}
	return nil
}

// Query implements the dialect.Query method.
func (c Conn) Query(ctx context.Context, query string, args, v any) error {
	vr, ok := v.(*Rows)
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect *sql.Rows", v)
	}
	argv, ok := args.([]any)
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect []any for args", args)
	}
	ex, cf, err := c.maySetVars(ctx)
	if err != nil {
		return err
	}
	rows, err := ex.QueryContext(ctx, query, argv...)
	if err != nil {
		if cf != nil {
			err = errors.Join(err, cf())
		}
		return err
	}
	*vr = Rows{rows}
	if cf != nil {
		vr.ColumnScanner = rowsWithCloser{rows, cf}
	}
	return nil
}

// maySetVars sets the session variables before executing a query.
func (c Conn) maySetVars(ctx context.Context) (ExecQuerier, func() error, error) {
	sv, _ := ctx.Value(ctxVarsKey{}).(sessionVars)
	if len(sv.vars) == 0 {
		return c, nil, nil
	}
	_, local := c.ExecQuerier.(*sql.Tx)
	queries, reset, err := sv.queries(c.dialect, local)
	if err != nil {
		return nil, nil, err
	}
	var ex ExecQuerier
	var cf func() error
	switch e := c.ExecQuerier.(type) {
	case *sql.Tx:
		ex = e
	case *sql.DB:
		conn, err := e.Conn(ctx)
		if err != nil {
			return nil, nil, err
		}
		ex, cf = conn, conn.Close
	}
	if close := cf; close != nil && len(reset) > 0 {
		cf = func() error {
			var err error
			for _, q := range reset {
				_, e := ex.ExecContext(ctx, q.query, q.args...)
				err = errors.Join(err, e)
			}
			return errors.Join(err, close())
		}
	}
	for _, q := range queries {
		if _, err := ex.ExecContext(ctx, q.query, q.args...); err != nil {
			if cf != nil {
				err = errors.Join(err, cf())
			}
			return nil, nil, err
		}
	}
	return ex, cf, nil
}

var _ dialect.Driver = (*Driver)(nil)

type (
	// Rows wraps the sql.Rows to avoid locks copy.
	Rows struct{ ColumnScanner }
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
	// NullTime represents a time.Time that may be null.
	NullTime = sql.NullTime
	// TxOptions holds the transaction options to be used in DB.BeginTx.
	TxOptions = sql.TxOptions
)

// NullScanner implements the sql.Scanner interface such that it
// can be used as a scan destination, similar to the types above.
type NullScanner struct {
	S     sql.Scanner
	Valid bool // Valid is true if the Scan value is not NULL.
}

// Scan implements the Scanner interface.
func (n *NullScanner) Scan(value any) error {
	n.Valid = value != nil
	if n.Valid {
		return n.S.Scan(value)
	}
	return nil
}

// ColumnScanner is the interface that wraps the standard
// sql.Rows methods used for scanning database rows.
type ColumnScanner interface {
	Close() error
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...any) error
}

// rowsWithCloser wraps the ColumnScanner interface with a custom Close hook.
type rowsWithCloser struct {
	ColumnScanner
	closer func() error
}

// Close closes the underlying ColumnScanner and calls the custom closer.
func (r rowsWithCloser) Close() error {
	err := r.ColumnScanner.Close()
	return errors.Join(err, r.closer())
}
