package gen

import dyschema "entgo.io/ent/dialect/dynamodb/schema"

// DyAttribute converts template Field to DynamoDB attribute.
func (f Field) DyAttribute() *dyschema.Attribute {
	return &dyschema.Attribute{
		Name: f.Name,
		Type: f.Type.Type,
	}
}

// AttributeConstant returns the attribute name of the relation column.
func (e Edge) AttributeConstant() string { return pascal(e.Name) + "Attribute" }
