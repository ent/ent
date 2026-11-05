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
	"strconv"
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

// ctyVarsKey is the key used for attaching and reading the context variables.
type ctxVarsKey struct{}

// sessionVars holds sessions/transactions variables to set before every statement.
type sessionVars struct {
	vars []struct{ k, v string }
}

// WithVar returns a new context that holds the session variable to be executed before every query.
// PostgreSQL variables set within a transaction are local to that transaction.
// Variables set outside a transaction are reset after execution or closing query rows.
// The name must be a trusted SQL variable name.
func WithVar(ctx context.Context, name, value string) context.Context {
	sv, _ := ctx.Value(ctxVarsKey{}).(sessionVars)
	sv.vars = append(sv.vars, struct {
		k, v string
	}{
		k: name,
		v: value,
	})
	return context.WithValue(ctx, ctxVarsKey{}, sv)
}

// VarFromContext returns the session variable value from the context.
func VarFromContext(ctx context.Context, name string) (string, bool) {
	sv, _ := ctx.Value(ctxVarsKey{}).(sessionVars)
	for _, s := range sv.vars {
		if s.k == name {
			return s.v, true
		}
	}
	return "", false
}

// WithIntVar calls WithVar with the string representation of the value.
func WithIntVar(ctx context.Context, name string, value int) context.Context {
	return WithVar(ctx, name, strconv.Itoa(value))
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

type varQuery struct {
	query string
	args  []any
}

// maySetVars sets the session variables before executing a query.
func (c Conn) maySetVars(ctx context.Context) (ExecQuerier, func() error, error) {
	sv, _ := ctx.Value(ctxVarsKey{}).(sessionVars)
	if len(sv.vars) == 0 {
		return c, nil, nil
	}
	_, local := c.ExecQuerier.(*sql.Tx)
	var first, queries, reset []varQuery
	seen := make(map[string]struct{}, len(sv.vars))
	for _, s := range sv.vars {
		key := s.k
		q := varQuery{fmt.Sprintf("SET %s = ?", s.k), []any{s.v}}
		cleanup := varQuery{fmt.Sprintf("SET %s = NULL", s.k), nil}
		if c.dialect == dialect.Postgres {
			name, value, err := postgresVar(s.k, s.v)
			if err != nil {
				return nil, nil, err
			}
			key = name
			if local {
				statement, err := postgresTxVar(name, value)
				if err != nil {
					return nil, nil, err
				}
				if statement != "" {
					first = append(first, varQuery{query: statement})
					continue
				}
			}
			q = varQuery{"SELECT set_config($1, $2, $3)", []any{name, value, local}}
			cleanup = varQuery{"SELECT set_config($1, NULL, false)", []any{name}}
		}
		queries = append(queries, q)
		if _, ok := seen[key]; !ok {
			reset = append(reset, cleanup)
			seen[key] = struct{}{}
		}
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
	for _, q := range append(first, queries...) {
		if _, err := ex.ExecContext(ctx, q.query, q.args...); err != nil {
			if cf != nil {
				err = errors.Join(err, cf())
			}
			return nil, nil, err
		}
	}
	return ex, cf, nil
}

// Transaction characteristics must be set before any SELECT acquires a snapshot.
// Both the name and value below come from fixed allowlists, not SQL input.
func postgresTxVar(name, value string) (string, error) {
	name = postgresLower(name)
	switch name {
	case "transaction_isolation":
		switch value = postgresLower(value); value {
		case "read uncommitted", "read committed", "repeatable read", "serializable":
		default:
			return "", fmt.Errorf("dialect/sql: invalid transaction isolation %q", value)
		}
	case "transaction_read_only", "transaction_deferrable":
		var ok bool
		value, ok = postgresBool(value)
		if !ok {
			return "", fmt.Errorf("dialect/sql: invalid transaction boolean")
		}
	default:
		return "", nil
	}
	return fmt.Sprintf("SET LOCAL %s = '%s'", name, value), nil
}

func postgresBool(value string) (string, bool) {
	switch postgresLower(strings.TrimSpace(value)) {
	case "1", "on", "t", "tr", "tru", "true", "y", "ye", "yes":
		return "on", true
	case "0", "of", "off", "f", "fa", "fal", "fals", "false", "n", "no":
		return "off", true
	default:
		return "", false
	}
}

// postgresVar converts a trusted SQL variable name to a configuration key.
// Identifier-list settings treat the entire value as one identifier, as SET does.
func postgresVar(name, value string) (string, string, error) {
	if strings.ContainsRune(name, 0) {
		return "", "", fmt.Errorf("dialect/sql: invalid PostgreSQL variable name %q", name)
	}
	var key strings.Builder
	rest := name
	for {
		rest = strings.TrimLeft(rest, " \t\n\r\f")
		part, tail, ok := postgresVarIdent(rest)
		if !ok {
			return "", "", fmt.Errorf("dialect/sql: invalid PostgreSQL variable name %q", name)
		}
		if key.Len() > 0 {
			key.WriteByte('.')
		}
		key.WriteString(part)
		rest = strings.TrimLeft(tail, " \t\n\r\f")
		if rest == "" {
			break
		}
		if rest[0] != '.' {
			return "", "", fmt.Errorf("dialect/sql: invalid PostgreSQL variable name %q", name)
		}
		rest = rest[1:]
	}
	name = key.String()
	switch postgresLower(name) {
	case "search_path", "temp_tablespaces", "session_preload_libraries",
		"shared_preload_libraries", "local_preload_libraries",
		"unix_socket_directories", "output_plugin_libraries":
		value = `"` + strings.ReplaceAll(value, `"`, `""`) + `"`
	}
	return name, value, nil
}

// postgresVarIdent reads one identifier without folding quoted characters.
func postgresVarIdent(s string) (part, rest string, ok bool) {
	if s == "" {
		return "", "", false
	}
	if s[0] == '"' {
		var quoted strings.Builder
		rest = s[1:]
		for {
			prefix, tail, found := strings.Cut(rest, `"`)
			if !found {
				return "", "", false
			}
			quoted.WriteString(prefix)
			if !strings.HasPrefix(tail, `"`) {
				return quoted.String(), tail, quoted.Len() > 0
			}
			quoted.WriteByte('"')
			rest = tail[1:]
		}
	}
	if !postgresIdentStart(s[0]) {
		return "", "", false
	}
	i := 1
	for i < len(s) && (postgresIdentStart(s[i]) || s[i] >= '0' && s[i] <= '9' || s[i] == '$') {
		i++
	}
	return postgresLower(s[:i]), s[i:], true
}

// postgresLower folds ASCII letters without changing non-ASCII identifier bytes.
func postgresLower(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			continue
		}
		b := []byte(s)
		for j := i; j < len(b); j++ {
			if b[j] >= 'A' && b[j] <= 'Z' {
				b[j] += 'a' - 'A'
			}
		}
		return string(b)
	}
	return s
}

func postgresIdentStart(c byte) bool {
	return c == '_' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= 0x80
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
