// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/pet"
	"entgo.io/ent/examples/migration/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescOwnerID is the schema descriptor for owner_id field.
	cardDescOwnerID := cardFields[1].Descriptor()
	// card.DefaultOwnerID holds the default value on creation for the owner_id field.
	card.DefaultOwnerID = cardDescOwnerID.Default.(int)
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescOwnerID is the schema descriptor for owner_id field.
	petDescOwnerID := petFields[3].Descriptor()
	// pet.DefaultOwnerID holds the default value on creation for the owner_id field.
	pet.DefaultOwnerID = petDescOwnerID.Default.(int)
	// petDescID is the schema descriptor for id field.
	petDescID := petFields[0].Descriptor()
	// pet.DefaultID holds the default value on creation for the id field.
	pet.DefaultID = petDescID.Default.(func() uuid.UUID)
}
