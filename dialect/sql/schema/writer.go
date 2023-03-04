// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"ariga.io/atlas/sql/migrate"
)

type (
	// WriteDriver is a driver that writes all driver exec operations to its writer.
	// Note that this driver is used only for printing or writing statements to SQL
	// files, and may require manual changes to the generated SQL statements.
	WriteDriver struct {
		dialect.Driver // optional driver for query calls.
		io.Writer      // target for exec statements.
		FormatFunc     func(string) (string, error)
	}
	// DirWriter implements the io.Writer interface
	// for writing to an Atlas managed directory.
	DirWriter struct {
		Dir       migrate.Dir       // target directory.
		Formatter migrate.Formatter // optional formatter.
		b         bytes.Buffer      // working buffer.
		changes   []*migrate.Change // changes to flush.
	}
)

// Write implements the io.Writer interface.
func (d *DirWriter) Write(p []byte) (int, error) { return d.b.Write(p) }

// Change converts all written statement so far into a migration
// change with the given comment.
func (d *DirWriter) Change(comment string) {
	// Trim semicolon and new line, because formatter adds it.
	d.changes = append(d.changes, &migrate.Change{Comment: comment, Cmd: strings.TrimRight(d.b.String(), ";\n")})
	d.b.Reset()
}

// Flush flushes the written statements to the directory.
func (d *DirWriter) Flush(name string) error {
	switch {
	case d.b.Len() != 0:
		return fmt.Errorf("writer has undocumented change. Use Change or FlushChange instead")
	case len(d.changes) == 0:
		return errors.New("writer has no changes to flush")
	default:
		return migrate.NewPlanner(nil, d.Dir, migrate.PlanFormat(d.Formatter)).
			WritePlan(&migrate.Plan{
				Name:    name,
				Changes: d.changes,
			})
	}
}

// FlushChange combines Change and Flush.
func (d *DirWriter) FlushChange(name, comment string) error {
	d.Change(comment)
	return d.Flush(name)
}

// NewWriteDriver creates a dialect.Driver that writes all driver exec statement to its writer.
func NewWriteDriver(dialect string, w io.Writer) *WriteDriver {
	return &WriteDriver{
		Writer: w,
		Driver: nopDriver{dialect: dialect},
	}
}

// Exec implements the dialect.Driver.Exec method.
func (w *WriteDriver) Exec(_ context.Context, query string, args, res any) error {
	if rr, ok := res.(*sql.Result); ok {
		*rr = noResult{}
	}
	if !strings.HasSuffix(query, ";") {
		query += ";"
	}
	if args != nil {
		args, ok := args.([]any)
		if !ok {
			return fmt.Errorf("unexpected args type: %T", args)
		}
		query = w.expandArgs(query, args)
	}
	_, err := io.WriteString(w, query+"\n")
	return err
}

// Query implements the dialect.Driver.Query method.
func (w *WriteDriver) Query(ctx context.Context, query string, args, res any) error {
	if strings.HasPrefix(query, "INSERT") || strings.HasPrefix(query, "UPDATE") {
		if err := w.Exec(ctx, query, args, nil); err != nil {
			return err
		}
		if rr, ok := res.(*sql.Rows); ok {
			*rr = sql.Rows{ColumnScanner: noRows{}}
		}
		return nil
	}
	switch w.Driver.(type) {
	case nil, nopDriver:
		return errors.New("query is not supported by the WriteDriver")
	default:
		return w.Driver.Query(ctx, query, args, res)
	}
}

// expandArgs combines to arguments and statement into a single statement to
// print or write into a file (before editing).
// Note, the output may be incorrect or unsafe SQL and require manual changes.
func (w *WriteDriver) expandArgs(query string, args []any) string {
	var (
		b    strings.Builder
		p    = w.placeholder()
		scan = w.scanPlaceholder()
	)
	for i := 0; i < len(query); i++ {
	Top:
		switch query[i] {
		case p:
			idx, size := scan(query[i+1:])
			// Unrecognized placeholder.
			if idx < 0 || idx >= len(args) {
				return query
			}
			i += size
			v, err := w.formatArg(args[idx])
			if err != nil {
				// Unexpected formatting error.
				return query
			}
			b.WriteString(v)
		// String or identifier.
		case '\'', '"', '`':
			for j := i + 1; j < len(query); j++ {
				switch query[j] {
				case '\\':
					j++
				case query[i]:
					b.WriteString(query[i : j+1])
					i = j
					break Top
				}
			}
			// Unexpected EOS.
			return query
		default:
			b.WriteByte(query[i])
		}
	}
	return b.String()
}

func (w *WriteDriver) scanPlaceholder() func(string) (int, int) {
	switch w.Dialect() {
	case dialect.Postgres:
		return func(s string) (int, int) {
			var i int
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				i++
			}
			idx, err := strconv.ParseInt(s[:i], 10, 64)
			if err != nil {
				return -1, 0
			}
			// Placeholders are 1-based.
			return int(idx) - 1, i
		}
	default:
		idx := -1
		return func(string) (int, int) {
			idx++
			return idx, 0
		}
	}
}

func (w *WriteDriver) placeholder() byte {
	if w.Dialect() == dialect.Postgres {
		return '$'
	}
	return '?'
}

func (w *WriteDriver) formatArg(v any) (string, error) {
	if w.FormatFunc != nil {
		return w.FormatFunc(fmt.Sprint(v))
	}
	switch v := v.(type) {
	case nil:
		return "NULL", nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%g", v), nil
	case bool:
		if v {
			return "1", nil
		} else {
			return "0", nil
		}
	case string:
		return "'" + strings.ReplaceAll(v, "'", "''") + "'", nil
	case json.RawMessage:
		return "'" + strings.ReplaceAll(string(v), "'", "''") + "'", nil
	case []byte:
		return "{{ BINARY_VALUE }}", nil
	case time.Time:
		return "{{ TIME_VALUE }}", nil
	case fmt.Stringer:
		return "'" + strings.ReplaceAll(v.String(), "'", "''") + "'", nil
	default:
		return "{{ VALUE }}", nil
	}
}

// Tx writes the transaction start.
func (w *WriteDriver) Tx(context.Context) (dialect.Tx, error) {
	return dialect.NopTx(w), nil
}

// noResult represents a zero result.
type noResult struct{}

func (noResult) LastInsertId() (int64, error) { return 0, nil }
func (noResult) RowsAffected() (int64, error) { return 0, nil }

// noRows represents no rows.
type noRows struct{ sql.ColumnScanner }

func (noRows) Close() error { return nil }
func (noRows) Err() error   { return nil }
func (noRows) Next() bool   { return false }

type nopDriver struct {
	dialect.Driver
	dialect string
}

func (d nopDriver) Dialect() string { return d.dialect }

func (nopDriver) Query(context.Context, string, any, any) error {
	return nil
}
