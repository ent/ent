// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package internal

import "context"

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig struct {
	Group      string // Group table.
	GroupUsers string // Group-users->User table.
	Pet        string // Pet table.
	User       string // User table.
}

type schemaCtxKey struct{}

// SchemaConfigFromContext returns a SchemaConfig stored inside a context, or empty if there isn't one.
func SchemaConfigFromContext(ctx context.Context) SchemaConfig {
	config, _ := ctx.Value(schemaCtxKey{}).(SchemaConfig)
	return config
}

// NewSchemaConfigContext returns a new context with the given SchemaConfig attached.
func NewSchemaConfigContext(parent context.Context, config SchemaConfig) context.Context {
	return context.WithValue(parent, schemaCtxKey{}, config)
}
