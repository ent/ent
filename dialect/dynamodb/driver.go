// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"entgo.io/ent/dialect"
)

// Driver is a dialect.Driver implementation for DynamoDB based databases.
type Driver struct {
	Client
	dialect string
}

// NewDriver creates a new Driver with the given Conn and dialect.
func NewDriver(dialect string, c Client) *Driver {
	return &Driver{dialect: dialect, Client: c}
}

// Open returns a dialect.Driver that implements the ent/dialect.Driver interface.
func Open(dialect, source string) (*Driver, error) {
	var (
		awsCfg                aws.Config
		err                   error
		endpointUrlOptionFunc config.LoadOptionsFunc
	)
	if source != "" {
		endpointUrlOptionFunc = config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: source}, nil
				}),
		)
		awsCfg, err = config.LoadDefaultConfig(context.TODO(), endpointUrlOptionFunc)
	} else {
		awsCfg, err = config.LoadDefaultConfig(context.TODO())
	}
	if err != nil {
		return nil, err
	}
	// Using the Config value, create the DynamoDB client
	dynamoDBClient := dynamodb.NewFromConfig(awsCfg)
	return NewDriver(dialect, Client{dynamoDBClient}), nil
}

// Dialect implements the dialect.Dialect method.
func (d Driver) Dialect() string {
	return d.dialect
}

// Tx starts and returns a transaction.
func (d *Driver) Tx(ctx context.Context) (dialect.Tx, error) {
	return dialect.NopTx(d), nil
}

// Close is a noop operation when the dialect is dynamodb
func (d *Driver) Close() error {
	return nil
}

// Exec implements the dialect.Exec method.
func (c Client) Exec(ctx context.Context, op string, args, v any) error {
	return c.run(ctx, op, args, v)
}

// Query implements the dialect.Query method.
func (c Client) Query(ctx context.Context, op string, args, v any) error {
	return c.run(ctx, op, args, v)
}

// run decides which SDK command to call
func (c Client) run(ctx context.Context, op string, args, v interface{}) error {
	switch op {
	case CreateTableOperation:
		createTableArgs := args.(*CreateTableArgs)
		_, err := c.CreateTable(ctx, createTableArgs.Opts)
		if err != nil {
			var inUseEx *types.ResourceInUseException
			if !errors.As(err, &inUseEx) {
				return err
			}
		}
		waiter := dynamodb.NewTableExistsWaiter(c)
		err = waiter.Wait(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(createTableArgs.Name),
		}, 30*time.Second)
		return err
	case PutItemOperation:
		putItemArgs := args.(*PutItemArgs)
		_, err := c.PutItem(ctx, putItemArgs.Opts)
		return err
	default:
		return fmt.Errorf("%s operation is unsupported", op)
	}
}

// Client wrap a DynamoDB client from AWS SDK to implement dialect.ExecQuerier
type Client struct {
	*dynamodb.Client
}

var (
	_ dialect.Driver = (*Driver)(nil)
)
