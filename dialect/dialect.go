// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dialect

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

// Dialect names for external usage.
const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
	Gremlin  = "gremlin"
)

// ExecQuerier wraps the 2 database operations.
type ExecQuerier interface {
	// Exec executes a query that does not return records. For example, in SQL, INSERT or UPDATE.
	// It scans the result into the pointer v. For SQL drivers, it is dialect/sql.Result.
	Exec(ctx context.Context, query string, args, v any) error
	// Query executes a query that returns rows, typically a SELECT in SQL.
	// It scans the result into the pointer v. For SQL drivers, it is *dialect/sql.Rows.
	Query(ctx context.Context, query string, args, v any) error
}

// Driver is the interface that wraps all necessary operations for ent clients.
type Driver interface {
	ExecQuerier
	// Tx starts and returns a new transaction.
	// The provided context is used until the transaction is committed or rolled back.
	Tx(context.Context) (Tx, error)
	// Close closes the underlying connection.
	Close() error
	// Dialect returns the dialect name of the driver.
	Dialect() string
}

// Tx wraps the Exec and Query operations in transaction.
type Tx interface {
	ExecQuerier
	driver.Tx
}

type nopTx struct {
	Driver
}

func (nopTx) Commit() error   { return nil }
func (nopTx) Rollback() error { return nil }

// NopTx returns a Tx with a no-op Commit / Rollback methods wrapping
// the provided Driver d.
func NopTx(d Driver) Tx {
	return nopTx{d}
}

type ContextLogger func(context.Context, ...any)

// DebugDriver is a driver that logs all driver operations.
type DebugDriver struct {
	Driver               // underlying driver.
	log    ContextLogger // log function. defaults to log.Println.
	errLog ContextLogger // errLog function log the queries and errors when encounters any errors. defaults to log.Println.
}

func DefaultContextLogger(_ context.Context, v ...any) {
	log.Println(v...)
}

func logError(ctx context.Context, logger ContextLogger, err error, msg ...string) {
	logger(ctx, fmt.Sprintf("an error occurred: %v in %v", err, msg))
}

func timeSubToMilliseconds(t1, t2 time.Time) float64 {
	return float64(t2.Sub(t1).Nanoseconds()) / 1000000
}

// Debug gets a driver, optional logging and error logging(the second logger) functions, and returns
// a new debugged-driver that prints all outgoing operations.
func Debug(d Driver, logger ...func(...any)) Driver {
	logf := log.Println
	errLogf := log.Println
	if len(logger) == 1 {
		logf = logger[0]
	}
	if len(logger) == 2 {
		errLogf = logger[1]
	}
	drv := &DebugDriver{d, func(_ context.Context, v ...any) { logf(v...) }, func(_ context.Context, v ...any) { errLogf(v...) }}
	return drv
}

// DebugWithContext gets a driver, optional logging and error logging(the second logger) functions, and returns
// a new debugged-driver that prints all outgoing operations with context.
func DebugWithContext(d Driver, logger ...func(context.Context, ...any)) Driver {
	logf := DefaultContextLogger
	errLogf := DefaultContextLogger
	if len(logger) == 1 {
		logf = logger[0]
	}
	if len(logger) == 2 {
		errLogf = logger[1]
	}
	drv := &DebugDriver{d, logf, errLogf}
	return drv
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *DebugDriver) Exec(ctx context.Context, query string, args, v any) error {
	qLog := fmt.Sprintf("driver.Exec: query=%v args=%v", query, args)
	d.log(ctx, qLog)
	err := d.Driver.Exec(ctx, query, args, v)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return err
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *DebugDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	st := time.Now()
	rs, err := drv.ExecContext(ctx, query, args...)
	et := time.Now()
	qLog := fmt.Sprintf("driver.ExecContext: latency= %.6fms query=%v args=%v", timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return rs, err
}

// Query logs its params and calls the underlying driver Query method.
func (d *DebugDriver) Query(ctx context.Context, query string, args, v any) error {
	st := time.Now()
	err := d.Driver.Query(ctx, query, args, v)
	et := time.Now()
	qLog := fmt.Sprintf("driver.Query: latency= %.6fms query=%v args=%v", timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return err
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *DebugDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	st := time.Now()
	rs, err := drv.QueryContext(ctx, query, args...)
	et := time.Now()
	qLog := fmt.Sprintf("driver.QueryContext: latency= %.6fms query=%v args=%v", timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return rs, err
}

// Tx adds an log-id for the transaction and calls the underlying driver Tx command.
func (d *DebugDriver) Tx(ctx context.Context) (Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		logError(ctx, d.errLog, err)
		return nil, err
	}
	id := uuid.New().String()
	d.log(ctx, fmt.Sprintf("driver.Tx(%s): started", id))
	return &DebugTx{tx, id, d.log, d.errLog, ctx}, nil
}

// BeginTx adds an log-id for the transaction and calls the underlying driver BeginTx command if it is supported.
func (d *DebugDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	drv, ok := d.Driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (Tx, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.BeginTx is not supported")
	}
	tx, err := drv.BeginTx(ctx, opts)
	if err != nil {
		logError(ctx, d.errLog, err)
		return nil, err
	}
	id := uuid.New().String()
	d.log(ctx, fmt.Sprintf("driver.BeginTx(%s): started", id))
	return &DebugTx{tx, id, d.log, d.errLog, ctx}, nil
}

// DebugTx is a transaction implementation that logs all transaction operations.
type DebugTx struct {
	Tx                     // underlying transaction.
	id     string          // transaction logging id.
	log    ContextLogger   // log function. defaults to fmt.Println.
	errLog ContextLogger   // errLog function log the queries and errors when encounters any errors. defaults to fmt.Println.
	ctx    context.Context // underlying transaction context.
}

// Exec logs its params and calls the underlying transaction Exec method.
func (d *DebugTx) Exec(ctx context.Context, query string, args, v any) error {
	st := time.Now()
	err := d.Tx.Exec(ctx, query, args, v)
	et := time.Now()
	qLog := fmt.Sprintf("Tx(%s).Exec: latency= %.6fms query=%v args=%v", d.id, timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return err
}

// ExecContext logs its params and calls the underlying transaction ExecContext method if it is supported.
func (d *DebugTx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Tx.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.ExecContext is not supported")
	}
	st := time.Now()
	rs, err := drv.ExecContext(ctx, query, args...)
	et := time.Now()
	qLog := fmt.Sprintf("Tx(%s).ExecContext: latency= %.6fms query=%v args=%v", d.id, timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return rs, err
}

// Query logs its params and calls the underlying transaction Query method.
func (d *DebugTx) Query(ctx context.Context, query string, args, v any) error {
	st := time.Now()
	err := d.Tx.Query(ctx, query, args, v)
	et := time.Now()
	qLog := fmt.Sprintf("Tx(%s).ExecContext: latency= %.6fms query=%v args=%v", d.id, timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return err
}

// QueryContext logs its params and calls the underlying transaction QueryContext method if it is supported.
func (d *DebugTx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Tx.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.QueryContext is not supported")
	}
	st := time.Now()
	rs, err := drv.QueryContext(ctx, query, args...)
	et := time.Now()
	qLog := fmt.Sprintf("Tx(%s).QueryContext: latency= %.6fms query=%v args=%v", d.id, timeSubToMilliseconds(st, et), query, args)
	d.log(ctx, qLog)
	if err != nil {
		logError(ctx, d.errLog, err, qLog)
	}
	return rs, err
}

// Commit logs this step and calls the underlying transaction Commit method.
func (d *DebugTx) Commit() error {
	st := time.Now()
	err := d.Tx.Commit()
	et := time.Now()
	qLog := fmt.Sprintf("Tx(%s): committed. latency= %.6fms", d.id, timeSubToMilliseconds(st, et))
	d.log(d.ctx, qLog)
	if err != nil {
		logError(d.ctx, d.errLog, err, qLog)
	}
	return err
}

// Rollback logs this step and calls the underlying transaction Rollback method.
func (d *DebugTx) Rollback() error {
	qLog := fmt.Sprintf("Tx(%s): rollbacked", d.id)
	d.log(d.ctx, qLog)
	err := d.Tx.Rollback()
	if err != nil {
		logError(d.ctx, d.errLog, err, qLog)
	}
	return err
}
