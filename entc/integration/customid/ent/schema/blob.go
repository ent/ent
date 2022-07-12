// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// Blob holds the schema definition for the Blob entity.
type Blob struct {
	ent.Schema
}

// Fields of the Blob.
func (Blob) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Annotations(entsql.Annotation{
				Default: "uuid_generate_v4()",
			}).
			Unique(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.Int("count").
			Default(0),
	}
}

// Edges of the Blob.
func (Blob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Blob.Type).
			Unique(),
		edge.To("links", Blob.Type).
			Through("blob_links", BlobLink.Type),
	}
}

// BlobLink holds the edge schema definition for blob links.
type BlobLink struct {
	ent.Schema
}

// Annotations of the BlobLink.
func (BlobLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("blob_id", "link_id"),
	}
}

// Fields of the BlobLink.
func (BlobLink) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.UUID("blob_id", uuid.UUID{}),
		field.UUID("link_id", uuid.UUID{}),
	}
}

// Edges of the BlobLink.
func (BlobLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blob", Blob.Type).
			Field("blob_id").
			Required().
			Unique(),
		edge.To("link", Blob.Type).
			Field("link_id").
			Required().
			Unique(),
	}
}
