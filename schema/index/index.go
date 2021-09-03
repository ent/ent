// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package index

import "entgo.io/ent/schema"

// A Descriptor for index configuration.
type Descriptor struct {
	Unique      bool                // unique index.
	Edges       []string            // edge columns.
	Fields      []string            // field columns.
	StorageKey  string              // custom index name.
	Annotations []schema.Annotation // index annotations.
}

// Builder for indexes on vertex columns and edges in the graph.
type Builder struct {
	desc *Descriptor
}

// Fields creates an index on the given vertex fields.
// Note that indexes are implemented only for SQL dialects, and does not support gremlin.
//
//	func (T) Indexes() []ent.Index {
//
//		// Unique index on 2 fields.
//		index.Fields("first", "last").
//			Unique(),
//
//		// Unique index of field under specific edge.
//		index.Fields("name").
//			Edges("parent").
//			Unique(),
//
//	}
//
func Fields(fields ...string) *Builder {
	return &Builder{desc: &Descriptor{Fields: fields}}
}

// Edges creates an index on the given vertex edge fields.
// Note that indexes are implemented only for SQL dialects, and does not support gremlin.
//
//	func (T) Indexes() []ent.Index {
//
//		// Unique index of field under 2 edges.
//		index.Fields("name").
//			Edges("parent", "type").
//			Unique(),
//
//	}
//
func Edges(edges ...string) *Builder {
	return &Builder{desc: &Descriptor{Edges: edges}}
}

// Fields sets the fields of the index.
//
//	func (T) Indexes() []ent.Index {
//
//		// Unique "name" and "age" fields under the "parent" edge.
//		index.Edges("parent").
//			Fields("name", "age").
//			Unique(),
//
//	}
func (b *Builder) Fields(fields ...string) *Builder {
	b.desc.Fields = fields
	return b
}

// Edges sets the fields index to be unique under the set of edges (sub-graph). For example:
//
//	func (T) Indexes() []ent.Index {
//
//		// Unique "name" field under the "parent" edge.
//		index.Fields("name").
//			Edges("parent").
//			Unique(),
//	}
//
func (b *Builder) Edges(edges ...string) *Builder {
	b.desc.Edges = edges
	return b
}

// Unique sets the index to be a unique index.
// Note that defining a uniqueness on optional fields won't prevent
// duplicates if one of the column contains NULL values.
func (b *Builder) Unique() *Builder {
	b.desc.Unique = true
	return b
}

// StorageKey sets the storage key of the index. In SQL dialects, it's the index name.
func (b *Builder) StorageKey(key string) *Builder {
	b.desc.StorageKey = key
	return b
}

// Annotations adds a list of annotations to the index object to be used by codegen extensions.
//
//	func (T) Indexes() []ent.Index {
//
//		// Partial index on name where the entity is not deleted.
//		index.Fields("name").
//			Annotations(entsql.Prefix(100))
//
//	}
//
func (b *Builder) Annotations(annotations ...schema.Annotation) *Builder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Descriptor interface.
func (b *Builder) Descriptor() *Descriptor {
	return b.desc
}
