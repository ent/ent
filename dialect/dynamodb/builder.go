// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// RootBuilder is the constructor for all DynamoDB operation builders.
type RootBuilder struct{}

func (d RootBuilder) CreateTable(name string) *CreateTableBuilder {
	return &CreateTableBuilder{
		tableName: name,
	}
}

func (d RootBuilder) PutItem(tableName string) *PutItemBuilder {
	return &PutItemBuilder{
		tableName: tableName,
	}
}

const (
	CreateTableOperation = "CreateTable"
	PutItemOperation     = "PutItem"
)

type (
	// CreateTableArgs contains input for CreateTable operation.
	CreateTableArgs struct {
		Name string
		Opts *dynamodb.CreateTableInput
	}

	// CreateTableBuilder is the builder for CreateTableArgs.
	CreateTableBuilder struct {
		attributeDefinitions []types.AttributeDefinition
		keySchema            []types.KeySchemaElement
		provisiondThroughput *types.ProvisionedThroughput
		tableName            string
	}
)

// CreateTable returns a builder for CreateTable operation.
func CreateTable(name string) *CreateTableBuilder {
	return &CreateTableBuilder{
		tableName: name,
	}
}

// AddAttribute adds one attribute to the table.
func (c *CreateTableBuilder) AddAttribute(attributeName string, attributeType types.ScalarAttributeType) *CreateTableBuilder {
	c.attributeDefinitions = append(c.attributeDefinitions, types.AttributeDefinition{
		AttributeName: aws.String(attributeName),
		AttributeType: attributeType,
	})
	return c
}

// AddKeySchemaElement adds one element to the key schema of the table.
func (c *CreateTableBuilder) AddKeySchemaElement(attributeName string, keyType types.KeyType) *CreateTableBuilder {
	c.keySchema = append(c.keySchema, types.KeySchemaElement{
		AttributeName: aws.String(attributeName),
		KeyType:       keyType,
	})
	return c
}

// SetProvisionedThroughput sets the provisioned throughput of the table.
func (c *CreateTableBuilder) SetProvisionedThroughput(readCap, writeCap int) *CreateTableBuilder {
	c.provisiondThroughput = &types.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(int64(readCap)),
		WriteCapacityUnits: aws.Int64(int64(writeCap)),
	}
	return c
}

// Op returns name and input for CreateTable operation.
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

type (
	// PutItemArgs contains input for PutItem operation.
	PutItemArgs struct {
		Opts *dynamodb.PutItemInput
	}

	// PutItemBuilder is the builder for PutItemArgs.
	PutItemBuilder struct {
		item      map[string]types.AttributeValue
		tableName string
	}
)

// PutItem returns a builder for PutItem operation.
func PutItem(tableName string) *PutItemBuilder {
	return &PutItemBuilder{
		tableName: tableName,
	}
}

// SetItem provides the data of the item to be put to that table.
func (p *PutItemBuilder) SetItem(i map[string]types.AttributeValue) *PutItemBuilder {
	p.item = i
	return p
}

// Op returns name and input for PutItem operation.
func (p *PutItemBuilder) Op() (string, interface{}) {
	return PutItemOperation, &PutItemArgs{
		Opts: &dynamodb.PutItemInput{
			TableName: aws.String(p.tableName),
			Item:      p.item,
		},
	}
}
