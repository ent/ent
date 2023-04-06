// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package cycle

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/load/testdata/cycle/fakent"
	"entgo.io/ent/schema/field"
)

// User holds the user schema.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("used", &Used{}),
		field.Enum("e").
			GoType(Enum(0)),
	}
}

type (
	Used        struct{}
	NotUsed     struct{}
	notExported struct{}
	Enum        int
)

func (Enum) Values() []string { return nil }

// The cause for cycle.
var _ fakent.Hook = nil
