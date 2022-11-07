// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package edge

import (
	"reflect"

	"entgo.io/ent/schema"
)

// A Descriptor for edge configuration.
type Descriptor struct {
	Tag         string                 // struct tag.
	Type        string                 // edge type.
	Name        string                 // edge name.
	Field       string                 // edge field name (e.g. foreign-key).
	RefName     string                 // ref name; inverse only.
	Ref         *Descriptor            // edge reference; to/from of the same type.
	Through     *struct{ N, T string } // through type and name.
	Unique      bool                   // unique edge.
	Inverse     bool                   // inverse edge.
	Required    bool                   // required on creation.
	Immutable   bool                   // create only edge.
	StorageKey  *StorageKey            // optional storage-key configuration.
	Annotations []schema.Annotation    // edge annotations.
	Comment     string                 // edge comment.
}

// To defines an association edge between two vertices.
func To(name string, t any) *assocBuilder {
	return &assocBuilder{desc: &Descriptor{Name: name, Type: typ(t)}}
}

// From represents a reversed-edge between two vertices that has a back-reference to its source edge.
func From(name string, t any) *inverseBuilder {
	return &inverseBuilder{desc: &Descriptor{Name: name, Type: typ(t), Inverse: true}}
}

func typ(t any) string {
	if rt := reflect.TypeOf(t); rt.NumIn() > 0 {
		return rt.In(0).Name()
	}
	return ""
}

// assocBuilder is the builder for assoc edges.
type assocBuilder struct {
	desc *Descriptor
}

// Unique sets the edge type to be unique. Basically, it limits the edge to be one of the two:
// one2one or one2many. one2one applied if the inverse-edge is also unique.
func (b *assocBuilder) Unique() *assocBuilder {
	b.desc.Unique = true
	return b
}

// Required indicates that this edge is a required field on creation.
// Unlike fields, edges are optional by default.
func (b *assocBuilder) Required() *assocBuilder {
	b.desc.Required = true
	return b
}

// Immutable indicates that this edge cannot be updated.
func (b *assocBuilder) Immutable() *assocBuilder {
	b.desc.Immutable = true
	return b
}

// StructTag sets the struct tag of the assoc edge.
func (b *assocBuilder) StructTag(s string) *assocBuilder {
	b.desc.Tag = s
	return b
}

// From creates an inverse-edge with the same type.
func (b *assocBuilder) From(name string) *inverseBuilder {
	return &inverseBuilder{desc: &Descriptor{Name: name, Type: b.desc.Type, Inverse: true, Ref: b.desc}}
}

// Field is used to bind an edge (with a foreign-key) to a field in the schema.
//
//	field.Int("owner_id").
//		Optional()
//
//	edge.To("owner", User.Type).
//		Field("owner_id").
//		Unique(),
func (b *assocBuilder) Field(f string) *assocBuilder {
	b.desc.Field = f
	return b
}

// Through allows setting an "edge schema" to interact explicitly with M2M edges.
//
//	edge.To("friends", User.Type).
//		Through("friendships", Friendship.Type)
func (b *assocBuilder) Through(name string, t any) *assocBuilder {
	b.desc.Through = &struct{ N, T string }{N: name, T: typ(t)}
	return b
}

// Comment used to put annotations on the schema.
func (b *assocBuilder) Comment(c string) *assocBuilder {
	b.desc.Comment = c
	return b
}

// StorageKey sets the storage key of the edge.
//
//	edge.To("groups", Group.Type).
//		StorageKey(edge.Table("user_groups"), edge.Columns("user_id", "group_id"))
func (b *assocBuilder) StorageKey(opts ...StorageOption) *assocBuilder {
	if b.desc.StorageKey == nil {
		b.desc.StorageKey = &StorageKey{}
	}
	for i := range opts {
		opts[i](b.desc.StorageKey)
	}
	return b
}

// Annotations adds a list of annotations to the edge object to be used by
// codegen extensions.
//
//	edge.To("pets", Pet.Type).
//		Annotations(entgql.Bind())
func (b *assocBuilder) Annotations(annotations ...schema.Annotation) *assocBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Descriptor interface.
func (b *assocBuilder) Descriptor() *Descriptor {
	return b.desc
}

// inverseBuilder is the builder for inverse edges.
type inverseBuilder struct {
	desc *Descriptor
}

// Ref sets the referenced-edge of this inverse edge.
func (b *inverseBuilder) Ref(ref string) *inverseBuilder {
	b.desc.RefName = ref
	return b
}

// Unique sets the edge type to be unique. Basically, it limits the edge to be one of the two:
// one-2-one or one-2-many. one-2-one applied if the inverse-edge is also unique.
func (b *inverseBuilder) Unique() *inverseBuilder {
	b.desc.Unique = true
	return b
}

// Required indicates that this edge is a required field on creation.
// Unlike fields, edges are optional by default.
func (b *inverseBuilder) Required() *inverseBuilder {
	b.desc.Required = true
	return b
}

// Immutable indicates that this edge cannot be updated.
func (b *inverseBuilder) Immutable() *inverseBuilder {
	b.desc.Immutable = true
	return b
}

// StructTag sets the struct tag of the inverse edge.
func (b *inverseBuilder) StructTag(s string) *inverseBuilder {
	b.desc.Tag = s
	return b
}

// Comment used to put annotations on the schema.
func (b *inverseBuilder) Comment(c string) *inverseBuilder {
	b.desc.Comment = c
	return b
}

// Field is used to bind an edge (with a foreign-key) to a field in the schema.
//
//	field.Int("owner_id").
//		Optional()
//
//	edge.From("owner", User.Type).
//		Ref("pets").
//		Field("owner_id").
//		Unique(),
func (b *inverseBuilder) Field(f string) *inverseBuilder {
	b.desc.Field = f
	return b
}

// Through allows setting an "edge schema" to interact explicitly with M2M edges.
//
//	edge.From("liked_users", User.Type).
//		Ref("liked_tweets").
//		Through("likes", TweetLike.Type)
func (b *inverseBuilder) Through(name string, t any) *inverseBuilder {
	b.desc.Through = &struct{ N, T string }{N: name, T: typ(t)}
	return b
}

// Annotations adds a list of annotations to the edge object to be used by
// codegen extensions.
//
//	edge.From("owner", User.Type).
//		Ref("pets").
//		Unique().
//		Annotations(entgql.Bind())
func (b *inverseBuilder) Annotations(annotations ...schema.Annotation) *inverseBuilder {
	b.desc.Annotations = append(b.desc.Annotations, annotations...)
	return b
}

// Descriptor implements the ent.Descriptor interface.
func (b *inverseBuilder) Descriptor() *Descriptor {
	return b.desc
}

// StorageKey holds the configuration for edge storage-key.
type StorageKey struct {
	Table   string   // Table or label.
	Symbols []string // Symbols/names of the foreign-key constraints.
	Columns []string // Foreign-key columns.
}

// StorageOption allows for setting the storage configuration using functional options.
type StorageOption func(*StorageKey)

// Table sets the table name option for M2M edges.
func Table(name string) StorageOption {
	return func(key *StorageKey) {
		key.Table = name
	}
}

// Symbol sets the symbol/name of the foreign-key constraint for O2O, O2M and M2O edges.
// Note that, for M2M edges (2 columns and 2 constraints), use the edge.Symbols option.
func Symbol(symbol string) StorageOption {
	return func(key *StorageKey) {
		key.Symbols = []string{symbol}
	}
}

// Symbols sets the symbol/name of the foreign-key constraints for M2M edges.
// The 1st column defines the name of the "To" edge, and the 2nd defines
// the name of the "From" edge (inverse edge).
// Note that, for O2O, O2M and M2O edges, use the edge.Symbol option.
func Symbols(to, from string) StorageOption {
	return func(key *StorageKey) {
		key.Symbols = []string{to, from}
	}
}

// Column sets the foreign-key column name option for O2O, O2M and M2O edges.
// Note that, for M2M edges (2 columns), use the edge.Columns option.
func Column(name string) StorageOption {
	return func(key *StorageKey) {
		key.Columns = []string{name}
	}
}

// Columns sets the foreign-key column names option for M2M edges.
// The 1st column defines the name of the "To" edge, and the 2nd defines
// the name of the "From" edge (inverse edge).
// Note that, for O2O, O2M and M2O edges, use the edge.Column option.
func Columns(to, from string) StorageOption {
	return func(key *StorageKey) {
		key.Columns = []string{to, from}
	}
}
