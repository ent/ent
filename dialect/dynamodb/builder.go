// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Oper wraps the basic stage implementation.
type Oper interface {
	// Op returns the database's operation representation.
	Op() (string, interface{})
}

// RootBuilder is the constructor for all DynamoDB operation builders.
type RootBuilder struct{}

// CreateTable returns a builder for CreateTable operation.
func (d RootBuilder) CreateTable(name string) *CreateTableBuilder {
	return &CreateTableBuilder{
		tableName: name,
	}
}

// PutItem returns a builder for PutItem operation.
func (d RootBuilder) PutItem(tableName string) *PutItemBuilder {
	return &PutItemBuilder{
		tableName: tableName,
	}
}

// Update returns a builder for UpdateItem operation.
func (d RootBuilder) Update(tableName string) *UpdateItemBuilder {
	return &UpdateItemBuilder{
		tableName:  tableName,
		expBuilder: expression.NewBuilder(),
		key:        make(map[string]types.AttributeValue),
	}
}

const (
	CreateTableOperation = "CreateTable"
	PutItemOperation     = "PutItem"
	UpdateItemOperation  = "UpdateItem"
	BatchWriteOperation  = "BatchWrite"
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

type (
	// UpdateItemBuilder is the builder for UpdateItemArg.
	UpdateItemBuilder struct {
		tableName    string
		key          map[string]types.AttributeValue
		expBuilder   expression.Builder
		exp          expression.Expression
		returnValues types.ReturnValue
	}
)

// WithKey selects which item to be updated for the UpdateItem operation.
func (u *UpdateItemBuilder) WithKey(k string, v types.AttributeValue) *UpdateItemBuilder {
	u.key[k] = v
	return u
}

// Set sets the attribute to be updated.
func (u *UpdateItemBuilder) Set(k string, v interface{}) *UpdateItemBuilder {
	setExp := expression.Set(expression.Name(k), expression.Value(v))
	u.expBuilder = u.expBuilder.WithUpdate(setExp)
	return u
}

// Where sets the expression for the UpdateItem operation.
func (u *UpdateItemBuilder) Where(p *Predicate) *UpdateItemBuilder {
	u.expBuilder = u.expBuilder.WithCondition(p.Query())
	return u
}

// Query returns potential errors during UpdateItemBuilder's build process.
func (u *UpdateItemBuilder) Query(rv types.ReturnValue) (*UpdateItemBuilder, error) {
	var err error
	u.returnValues = rv
	u.exp, err = u.expBuilder.Build()
	return u, err
}

// Op returns name and input for UpdateItem operation.
func (u *UpdateItemBuilder) Op() (string, interface{}) {
	return UpdateItemOperation, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(u.tableName),
		Key:                       u.key,
		ConditionExpression:       u.exp.Condition(),
		UpdateExpression:          u.exp.Update(),
		ExpressionAttributeNames:  u.exp.Names(),
		ExpressionAttributeValues: u.exp.Values(),
		ReturnValues:              u.returnValues,
	}
}

type (
	// BatchWriteItemBuilder is the builder for BatchWriteItem operation.
	BatchWriteItemBuilder struct {
		requestMap map[string][]Oper
	}

	// BatchWriteItemArgs contains input of BatchWriteItem operation.
	BatchWriteItemArgs struct {
		operationMap map[string][]Oper
	}
)

// BatchWriteItem returns a builder for BatchWriteItem operation.
func BatchWriteItem() *BatchWriteItemBuilder {
	return &BatchWriteItemBuilder{}
}

// Append appends a WriteRequest to BatchWriteItem.
func (b *BatchWriteItemBuilder) Append(tableName string, op Oper) *BatchWriteItemBuilder {
	b.requestMap[tableName] = append(b.requestMap[tableName], op)
	return b
}

// Op returns name and input for BatchWriteItem operation.
func (b *BatchWriteItemBuilder) Op() (string, interface{}) {
	return BatchWriteOperation, &BatchWriteItemArgs{
		operationMap: b.requestMap,
	}
}
