// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
		field.String("some_field"),
	}
}

// MixinID holds the schema definition for the MixinID entity.
type MixinID struct {
	ent.Schema
}

// Fields of the MixinID.
func (MixinID) Fields() []ent.Field {
	return []ent.Field{
		field.String("mixin_field"),
	}
}

// Indexes of the MixinID
func (MixinID) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("id", "some_field"),
		index.Fields("id", "mixin_field"),
		index.Fields("id", "mixin_field", "some_field"),
	}
}

// Mixin of MixinID
func (MixinID) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
