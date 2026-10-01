// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package generics

import (
	"database/sql"
	"time"

	"entgo.io/ent"
	otheruuid "entgo.io/ent/entc/gen/testdata/generics/uuid"
	versioned "entgo.io/ent/entc/gen/testdata/generics/versioned/v2"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

type Text[T any] string
type Pair[A, B any] string

type Record struct{ ent.Schema }

func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(Text[uuid.UUID]("")),
		field.String("builtin").GoType(Text[string]("")),
		field.String("timestamp").GoType(Text[time.Time]("")).DefaultFunc(func() Text[time.Time] { return "now" }),
		field.String("nested").GoType(Text[Text[uuid.UUID]]("")),
		field.String("pair").GoType(Pair[uuid.UUID, otheruuid.ID]("")),
		field.String("versioned").GoType(Text[versioned.Value]("")),
		field.Time("nullable_time").GoType(sql.Null[time.Time]{}).Optional(),
		field.Time("nullable_pointer").GoType(&sql.Null[time.Time]{}).Optional().Nillable(),
		field.JSON("values", map[uuid.UUID][]*Text[time.Time]{}).Optional(),
	}
}

func (Record) Edges() []ent.Edge {
	return []ent.Edge{edge.To("items", Item.Type)}
}

type Item struct{ ent.Schema }

func (Item) Fields() []ent.Field {
	return []ent.Field{field.String("other").GoType(Text[otheruuid.ID](""))}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{edge.From("record", Record.Type).Ref("items").Unique()}
}

// UUID also checks collisions between a type argument's package and a node.
type UUID struct{ ent.Schema }
