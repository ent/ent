// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
)

// ctxVarsKey is the key used for attaching and reading the context variables.
type ctxVarsKey struct{}

// sessionVars holds sessions/transactions variables to set before every statement.
type sessionVars struct {
	vars []sessionVar
}

type sessionVar struct {
	name, value string
}

// WithVar returns a new context that holds the session variable to be executed before every query.
// PostgreSQL variables set within a transaction are local to that transaction.
// Variables set outside a transaction are reset after execution or closing query rows.
// The name must be a trusted SQL variable name.
func WithVar(ctx context.Context, name, value string) context.Context {
	sv, _ := ctx.Value(ctxVarsKey{}).(sessionVars)
	sv.vars = append(sv.vars, sessionVar{name: name, value: value})
	return context.WithValue(ctx, ctxVarsKey{}, sv)
}

// VarFromContext returns the session variable value from the context.
func VarFromContext(ctx context.Context, name string) (string, bool) {
	sv, _ := ctx.Value(ctxVarsKey{}).(sessionVars)
	for _, s := range sv.vars {
		if s.name == name {
			return s.value, true
		}
	}
	return "", false
}

// WithIntVar calls WithVar with the string representation of the value.
func WithIntVar(ctx context.Context, name string, value int) context.Context {
	return WithVar(ctx, name, strconv.Itoa(value))
}

type varQuery struct {
	query string
	args  []any
}

func (s sessionVars) queries(driver string, local bool) ([]varQuery, []varQuery, error) {
	var first, queries, reset []varQuery
	seen := make(map[string]struct{}, len(s.vars))
	postgres := driver == dialect.Postgres
	for _, v := range s.vars {
		name, value := v.name, v.value
		if postgres {
			var err error
			name, value, err = postgresVar(name, value)
			if err != nil {
				return nil, nil, err
			}
		}
		if postgres && local {
			statement, err := postgresTxVar(name, value)
			if err != nil {
				return nil, nil, err
			}
			if statement != "" {
				first = append(first, varQuery{query: statement})
				continue
			}
		}
		var query, cleanup varQuery
		if postgres {
			query = varQuery{"SELECT set_config($1, $2, $3)", []any{name, value, local}}
			cleanup = varQuery{"SELECT set_config($1, NULL, false)", []any{name}}
		} else {
			query = varQuery{fmt.Sprintf("SET %s = ?", name), []any{value}}
			cleanup = varQuery{query: fmt.Sprintf("SET %s = NULL", name)}
		}
		queries = append(queries, query)
		if _, ok := seen[name]; ok {
			continue
		}
		reset = append(reset, cleanup)
		seen[name] = struct{}{}
	}
	return append(first, queries...), reset, nil
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
