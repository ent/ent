// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	uuidc "entgo.io/ent/entc/integration/customid/uuidcompatible"
	"entgo.io/ent/schema/field"
)

type LinkInformation struct {
	Name string
	Link string
}

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the IntSid.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuidc.UUIDC{}).Default(uuidc.NewUUIDC),
		field.JSON("link_information", map[string]LinkInformation{}).
			Default(map[string]LinkInformation{
				"ent": {
					Name: "ent",
					Link: "https://entgo.io/",
				},
			}),
	}
}
