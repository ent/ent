// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// SubjectStudent holds the schema for the subject entity.
type SubjectStudent struct {
	ent.Schema
}

// Fields of the SubjectStudent.
func (SubjectStudent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),

		field.String("note").Optional().Nillable(),

		field.UUID("subject_id", uuid.UUID{}),
		field.UUID("student_id", uuid.UUID{}),
	}
}

// Edges of the SubjectStudent.
func (SubjectStudent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subject", Subject.Type).Field("subject_id").Unique().Required().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("student", Student.Type).Field("student_id").Unique().Required().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

// Indexes of the SubjectStudent.
func (SubjectStudent) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("subject", "student").Unique(),
	}
}
