// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// PetUserName represents a user/pet name returned from the pet_user_names view.
type PetUserName struct {
	ent.View
}

// Fields of the PetUserName.
func (PetUserName) Fields() []ent.Field {
	return []ent.Field{
		// Skip adding the "id" field as
		// it is not needed for this view.
		field.String("name"),
	}
}
