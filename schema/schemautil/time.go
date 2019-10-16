package schemautil

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// CreateTimeMixin adds created at time field.
type CreateTimeMixin struct{}

// Fields of the create time mixin.
func (CreateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*CreateTimeMixin)(nil)

// UpdateTimeMixin adds updated at time field.
type UpdateTimeMixin struct{}

// Fields of the update time mixin.
func (UpdateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable(),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*UpdateTimeMixin)(nil)

// TimeMixin composes create/update time mixin.
type TimeMixin struct{}

// Fields of the time mixin.
func (TimeMixin) Fields() []ent.Field {
	return append(
		CreateTimeMixin{}.Fields(),
		UpdateTimeMixin{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*TimeMixin)(nil)
