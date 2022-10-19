// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	CreateTableOperation = "CreateTable"
)

// CreateTableArgs contains input for CreateTable operation
type CreateTableArgs struct {
	Name string
	Opts *dynamodb.CreateTableInput
}

// CreateTableBuilder is the builder for CreateTableArgs
type CreateTableBuilder struct {
	attributeDefinitions []types.AttributeDefinition
	keySchema            []types.KeySchemaElement
	provisiondThroughput *types.ProvisionedThroughput
	tableName            string
}

// CreateTable returns a builder for CreateTable operation
func CreateTable(name string) *CreateTableBuilder {
	return &CreateTableBuilder{
		tableName: name,
	}
}

// AddAttribute adds one attribute to the table
func (c *CreateTableBuilder) AddAttribute(attributeName string, attributeType types.ScalarAttributeType) *CreateTableBuilder {
	c.attributeDefinitions = append(c.attributeDefinitions, types.AttributeDefinition{
		AttributeName: aws.String(attributeName),
		AttributeType: attributeType,
	})
	return c
}

// AddKeySchemaElement adds one element to the key schema of the table
func (c *CreateTableBuilder) AddKeySchemaElement(attributeName string, keyType types.KeyType) *CreateTableBuilder {
	c.keySchema = append(c.keySchema, types.KeySchemaElement{
		AttributeName: aws.String(attributeName),
		KeyType:       keyType,
	})
	return c
}

// SetProvisionedThroughput sets the provisioned throughput of the table
func (c *CreateTableBuilder) SetProvisionedThroughput(readCap, writeCap int) *CreateTableBuilder {
	c.provisiondThroughput = &types.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(int64(readCap)),
		WriteCapacityUnits: aws.Int64(int64(writeCap)),
	}
	return c
}

// Op returns name and input for CreateTable operation
func (c *CreateTableBuilder) Op() (string, interface{}) {
	return CreateTableOperation, &CreateTableArgs{
		Name: c.tableName,
		Opts: &dynamodb.CreateTableInput{
			TableName:             aws.String(c.tableName),
			AttributeDefinitions:  c.attributeDefinitions,
			KeySchema:             c.keySchema,
			ProvisionedThroughput: c.provisiondThroughput,
		},
	}
}
