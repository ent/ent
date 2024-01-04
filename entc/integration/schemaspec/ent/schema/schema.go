// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/load"
)

var AllSchema = []ent.Interface{
	User{},
}

// NewLoadSpec creates a new SchemaSpec from this package
func NewLoadSpec() (*load.SchemaSpec, error) {
	// load them into entgo.io/ent/entc/load.Schema
	schemas := []*load.Schema{}
	for _, o := range AllSchema {
		b, err := load.MarshalSchema(o)
		if err != nil {
			return nil, err
		}
		loaded, err := load.UnmarshalSchema(b)
		if err != nil {
			return nil, err
		}
		schemas = append(schemas, loaded)
	}

	return &load.SchemaSpec{
		Schemas: schemas,
		PkgPath: "entgo.io/ent/entc/integration/schemaspec/ent/schema",
	}, nil
}
