package schema

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"entgo.io/ent/schema/field"
)

// Table schema definition for DynamoDB dialects.
type Table struct {
	Name          string
	Attributes    []*Attribute
	PrimaryKey    []*KeySchema
	ReadCapacity  int
	WriteCapacity int
}

// Attribute schema definition for DynamoDB dialects.
type Attribute struct {
	Name string
	Type field.Type
}

type KeySchema struct {
	AttributeName string
	KeyType       KeyType
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeHash  KeyType = "HASH"
	KeyTypeRange KeyType = "RANGE"
)

// dynamoType returns DynamoDB attribute type for the given Column type.
func (c *Attribute) dynamoType() types.ScalarAttributeType {
	switch c.Type {
	case field.TypeInt:
		return types.ScalarAttributeTypeN
	case field.TypeString:
		return types.ScalarAttributeTypeS
	case field.TypeBytes:
		return types.ScalarAttributeTypeB
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
}
