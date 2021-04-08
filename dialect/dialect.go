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
	// Exec executes a query that doesn't return rows. For example, in SQL, INSERT or UPDATE.
	// It scans the result into the pointer v. In SQL, you it's usually sql.Result.
	Exec(ctx context.Context, query string, args, v interface{}) error
	// Query executes a query that returns rows, typically a SELECT in SQL.
	// It scans the result into the pointer v. In SQL, you it's usually *sql.Rows.
	Query(ctx context.Context, query string, args, v interface{}) error
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

// DebugDriver is a driver that logs all driver operations.
type DebugDriver struct {
	Driver                                 // underlying driver.
	log    func(context.Context, LogEntry) // log function. defaults to log.Println.
}

// Debug gets a driver and an optional logging function, and returns
// a new debugged-driver that prints all outgoing operations.
func Debug(d Driver, logger ...func(...interface{})) Driver {
	logf := log.Println
	if len(logger) == 1 {
		logf = logger[0]
	}
	drv := &DebugDriver{d, func(_ context.Context, v LogEntry) { logf(v) }}
	return drv
}

// DebugWithContext gets a driver and a logging function, and returns
// a new debugged-driver that prints all outgoing operations with context.
func DebugWithContext(d Driver, logger func(context.Context, LogEntry)) Driver {
	drv := &DebugDriver{d, logger}
	return drv
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *DebugDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	d.log(ctx, LogEntry{
		Action: DriverActionExec,
		Query:  query,
		Args:   args,
	})
	return d.Driver.Exec(ctx, query, args, v)
}

// Query logs its params and calls the underlying driver Query method.
func (d *DebugDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	d.log(ctx, LogEntry{
		Action: DriverActionQuery,
		Query:  query,
		Args:   args,
	})
	return d.Driver.Query(ctx, query, args, v)
}

// Tx adds an log-id for the transaction and calls the underlying driver Tx command.
func (d *DebugDriver) Tx(ctx context.Context) (Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	d.log(ctx, LogEntry{
		Action: DriverActionTx,
		TxID:   id,
	})
	return &DebugTx{tx, id, d.log, ctx, nil}, nil
}

// BeginTx adds an log-id for the transaction and calls the underlying driver BeginTx command if it's supported.
func (d *DebugDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	drv, ok := d.Driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (Tx, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.BeginTx is not supported")
	}
	tx, err := drv.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	d.log(ctx, LogEntry{
		Action: DriverActionBeginTx,
		TxID:   id,
	})
	return &DebugTx{tx, id, d.log, ctx, opts}, nil
}

// DebugTx is a transaction implementation that logs all transaction operations.
type DebugTx struct {
	Tx                                   // underlying transaction.
	id   string                          // transaction logging id.
	log  func(context.Context, LogEntry) // log function. defaults to fmt.Println.
	ctx  context.Context                 // underlying transaction context.
	opts *sql.TxOptions                  // underlying transaction options.
}

// Exec logs its params and calls the underlying transaction Exec method.
func (d *DebugTx) Exec(ctx context.Context, query string, args, v interface{}) error {
	d.log(ctx, LogEntry{
		Action:   DriverActionTx,
		TxAction: TxActionExec,
		TxID:     d.id,
		TxOpt:    d.opts,
		Query:    query,
		Args:     args,
	})
	return d.Tx.Exec(ctx, query, args, v)
}

// Query logs its params and calls the underlying transaction Query method.
func (d *DebugTx) Query(ctx context.Context, query string, args, v interface{}) error {
	d.log(ctx, LogEntry{
		Action:   DriverActionTx,
		TxAction: TxActionQuery,
		TxID:     d.id,
		TxOpt:    d.opts,
		Query:    query,
		Args:     args,
	})
	return d.Tx.Query(ctx, query, args, v)
}

// Commit logs this step and calls the underlying transaction Commit method.
func (d *DebugTx) Commit() error {
	d.log(d.ctx, LogEntry{
		Action:   DriverActionTx,
		TxAction: TxActionCommit,
		TxID:     d.id,
		TxOpt:    d.opts,
	})
	return d.Tx.Commit()
}

// Rollback logs this step and calls the underlying transaction Rollback method.
func (d *DebugTx) Rollback() error {
	d.log(d.ctx, LogEntry{
		Action:   DriverActionTx,
		TxAction: TxActionRollback,
		TxID:     d.id,
		TxOpt:    d.opts,
	})
	return d.Tx.Rollback()
}
