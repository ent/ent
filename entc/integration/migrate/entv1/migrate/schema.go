// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"context"
	"fmt"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql/schema"
)

// SQLDialect wraps the dialect.Driver with additional migration methods.
type SQLDriver interface {
	Create(context.Context, ...*schema.Table) error
}

// Schema is the API for creating, migrating and dropping a schema.
type Schema struct {
	drv SQLDriver
}

// NewSchema creates a new schema client.
func NewSchema(drv dialect.Driver) *Schema {
	s := &Schema{}
	switch drv.Dialect() {
	case dialect.MySQL:
		s.drv = &schema.MySQL{Driver: drv}
	case dialect.SQLite:
		s.drv = &schema.SQLite{Driver: drv}
	}
	return s
}

// Create creates all schema resources.
func (s *Schema) Create(ctx context.Context) error {
	if s.drv == nil {
		return fmt.Errorf("entv1/migrate: dialect does not support migration")
	}
	return s.drv.Create(ctx, Tables...)
}
