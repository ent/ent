package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"fbc/ent/dialect"
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
	tx, err := d.ExecQuerier.(*sql.DB).BeginTx(ctx, &sql.TxOptions{})
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

// shared connection ExecQuerier between Gremlin and Tx.
type conn struct {
	ExecQuerier
}

// Exec implements the dialect.Exec method.
func (c *conn) Exec(ctx context.Context, query string, args interface{}, v interface{}) error {
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
func (c *conn) Query(ctx context.Context, query string, args interface{}, v interface{}) error {
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
)

// Note:
// NullTime is a modified copy of database/sql.NullTime from Go 1.13,
// It should be replaced with standard library code when Go 1.13 is released.

// NullTime represents a time.Time that may be null.
// NullTime implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (n *NullTime) Scan(v interface{}) error {
	if v, ok := v.(time.Time); ok {
		n.Time = v
		n.Valid = true
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
