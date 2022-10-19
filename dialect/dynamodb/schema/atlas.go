package schema

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/dynamodb"
)

// MigrateOption is just to complement existing schema template for now.
type MigrateOption interface{}

// Atlas is the migration engine.
type Atlas struct {
	dialect string
	driver  dialect.Driver
}

// NewMigrate creates a new Atlas for DynamoDB dialect.
func NewMigrate(drv dialect.Driver, opts ...MigrateOption) (*Atlas, error) {
	a := &Atlas{driver: drv}
	for _, opt := range opts {
		opt(a)
	}
	a.dialect = a.driver.Dialect()
	return a, nil
}

// Create creates all schema resources in the database.
func (a *Atlas) Create(ctx context.Context, tables ...*Table) (err error) {
	for _, t := range tables {
		ct := dynamodb.CreateTable(t.Name)
		for _, a := range t.Attributes {
			ct.AddAttribute(a.Name, a.dynamoType())
		}
		for _, ks := range t.PrimaryKey {
			ct.AddKeySchemaElement(ks.AttributeName, types.KeyType(ks.KeyType))
		}
		ct.SetProvisionedThroughput(t.ReadCapacity, t.WriteCapacity)
		op, args := ct.Op()
		if err := a.driver.Exec(ctx, op, args, nil); err != nil {
			return err
		}
	}
	return nil
}
