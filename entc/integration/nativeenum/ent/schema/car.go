// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CarCar holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("transmission"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{}
}
