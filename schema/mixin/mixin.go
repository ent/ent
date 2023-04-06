// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Schema is the default implementation for the ent.Mixin interface.
// It should be embedded in end-user mixin as follows:
//
//	type M struct {
//		mixin.Schema
//	}
type Schema struct{}

// Fields of the mixin.
func (Schema) Fields() []ent.Field { return nil }

// Edges of the mixin.
func (Schema) Edges() []ent.Edge { return nil }

// Indexes of the mixin.
func (Schema) Indexes() []ent.Index { return nil }

// Hooks of the mixin.
func (Schema) Hooks() []ent.Hook { return nil }

// Interceptors of the schema.
func (Schema) Interceptors() []ent.Interceptor { return nil }

// Policy of the mixin.
func (Schema) Policy() ent.Policy { return nil }

// Annotations of the mixin.
func (Schema) Annotations() []schema.Annotation { return nil }

// schema mixin must implement `Mixin` interface.
var _ ent.Mixin = (*Schema)(nil)

// CreateTime adds created at time field.
type CreateTime struct{ Schema }

// Fields of the create time mixin.
func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Default(time.Now).
			Immutable(),
	}
}

// create time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*CreateTime)(nil)

// UpdateTime adds updated at time field.
type UpdateTime struct{ Schema }

// Fields of the update time mixin.
func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// update time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*UpdateTime)(nil)

// Time composes create/update time mixin.
type Time struct{ Schema }

// Fields of the time mixin.
func (Time) Fields() []ent.Field {
	return append(
		CreateTime{}.Fields(),
		UpdateTime{}.Fields()...,
	)
}

// time mixin must implement `Mixin` interface.
var _ ent.Mixin = (*Time)(nil)

// AnnotateFields adds field annotations to underlying mixin fields.
func AnnotateFields(m ent.Mixin, annotations ...schema.Annotation) ent.Mixin {
	return fieldAnnotator{Mixin: m, annotations: annotations}
}

// AnnotateEdges adds edge annotations to underlying mixin edges.
func AnnotateEdges(m ent.Mixin, annotations ...schema.Annotation) ent.Mixin {
	return edgeAnnotator{Mixin: m, annotations: annotations}
}

type fieldAnnotator struct {
	ent.Mixin
	annotations []schema.Annotation
}

func (a fieldAnnotator) Fields() []ent.Field {
	fields := a.Mixin.Fields()
	for i := range fields {
		desc := fields[i].Descriptor()
		desc.Annotations = append(desc.Annotations, a.annotations...)
	}
	return fields
}

type edgeAnnotator struct {
	ent.Mixin
	annotations []schema.Annotation
}

func (a edgeAnnotator) Edges() []ent.Edge {
	edges := a.Mixin.Edges()
	for i := range edges {
		desc := edges[i].Descriptor()
		desc.Annotations = append(desc.Annotations, a.annotations...)
	}
	return edges
}
