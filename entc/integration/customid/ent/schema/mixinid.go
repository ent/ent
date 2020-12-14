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
		field.String("someOtherField"),
	}
}

// MixinID holds the schema definition for the MixinID entity.
type MixinID struct {
	ent.Schema
}

// Fields of the MixinID.
func (MixinID) Fields() []ent.Field {
	return []ent.Field{
		field.String("testField"),
	}
}

// Edges of the MixinID.
func (MixinID) Edges() []ent.Edge {
	return nil
}

func (MixinID) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (MixinID) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
