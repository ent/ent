package schema

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"entgo.io/ent/schema/field"
)

// Table schema definition for DynamoDB dialects.
type Table struct {
	Name          string
	attributes    map[string]*Attribute
	Attributes    []*Attribute
	PrimaryKey    []*KeySchema
	ReadCapacity  int
	WriteCapacity int
}

// NewTable returns a new table with the given name.
func NewTable(name string) *Table {
	return &Table{
		Name:       name,
		attributes: make(map[string]*Attribute),
	}
}

// AddAttribute adds a new attribute to the table.
func (t *Table) AddAttribute(a *Attribute) *Table {
	t.attributes[a.Name] = a
	t.Attributes = append(t.Attributes, a)
	return t
}

// AddKeySchema adds a new key schema to the primary key.
func (t *Table) AddKeySchema(ks *KeySchema) *Table {
	t.PrimaryKey = append(t.PrimaryKey, ks)
	return t
}

// Attribute schema definition for DynamoDB dialects.
type Attribute struct {
	Name string
	Type field.Type
}

// KeySchema schema definition for DynamoDB KeySchemaElement.
type KeySchema struct {
	AttributeName string
	KeyType       KeyType
}

// KeyType defines if the key element is HASH (partition key)
// or RANGE (sort key).
type KeyType string

// Enum values for KeyType.
const (
	KeyTypeHash  KeyType = "HASH"
	KeyTypeRange KeyType = "RANGE"
)

// dynamoType returns DynamoDB attribute type for the given Column type.
func (c *Attribute) dynamoType() types.ScalarAttributeType {
	switch c.Type {
	case field.TypeInt:
		return types.ScalarAttributeTypeN
	case field.TypeString, field.TypeTime:
		return types.ScalarAttributeTypeS
	case field.TypeBytes:
		return types.ScalarAttributeTypeB
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
}
