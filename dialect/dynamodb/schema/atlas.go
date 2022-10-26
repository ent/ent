package schema

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/dynamodb"
)

// Atlas is the migration engine.
type Atlas struct {
	dialect string
	driver  dialect.Driver
	builder dynamodb.RootBuilder
}

// NewMigrate creates a new Atlas for DynamoDB dialect.
func NewMigrate(drv dialect.Driver, opts ...MigrateOption) (*Atlas, error) {
	a := &Atlas{driver: drv, builder: dynamodb.RootBuilder{}}
	for _, opt := range opts {
		opt(a)
	}
	a.dialect = a.driver.Dialect()
	return a, nil
}

// Create reads variables in migrate/schema.go
// and creates all schema resources in the database.
func (a *Atlas) Create(ctx context.Context, tables ...*Table) (err error) {
	for _, t := range tables {
		ct := a.builder.CreateTable(t.Name)
		for _, ks := range t.PrimaryKey {
			// The map of t.attributes is empty.
			for _, a := range t.Attributes {
				if a.Name == ks.AttributeName {
					ct.AddAttribute(a.Name, a.dynamoType())
					ct.AddKeySchemaElement(ks.AttributeName, types.KeyType(ks.KeyType))
				}
			}
		}
		// ProvisionedThroughput is required. Use hardcoded values for now.
		ct.SetProvisionedThroughput(10, 10)
		op, args := ct.Op()
		if err := a.driver.Exec(ctx, op, args, nil); err != nil {
			return err
		}
	}
	return nil
}

// MigrateOption allows configuring Atlas using functional arguments.
type MigrateOption func(*Atlas)

// WithGlobalUniqueID is a noop fuction for now.
func WithGlobalUniqueID(b bool) MigrateOption {
	return func(a *Atlas) {
		return
	}
}

// WithDropColumn is a noop fuction for now.
func WithDropColumn(b bool) MigrateOption {
	return func(a *Atlas) {
		return
	}
}

// WithDropIndex is a noop fuction for now.
func WithDropIndex(b bool) MigrateOption {
	return func(a *Atlas) {
		return
	}
}

// WithForeignKeys is a noop fuction for now.
func WithForeignKeys(b bool) MigrateOption {
	return func(a *Atlas) {
		return
	}
}
