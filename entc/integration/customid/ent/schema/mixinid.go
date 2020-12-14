package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebook/ent/schema/mixin"
	"github.com/google/uuid"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("someOtherField"),
	}
}

// MixinId holds the schema definition for the MixinId entity.
type MixinId struct {
	ent.Schema
}

// Fields of the MixinId.
func (MixinId) Fields() []ent.Field {
	return []ent.Field{
		field.String("testField"),
	}
}

// Edges of the MixinId.
func (MixinId) Edges() []ent.Edge {
	return nil
}

func (MixinId) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (MixinId) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
