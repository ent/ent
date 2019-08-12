package index

// Index represents an index on a vertex columns.
type Index struct {
	unique bool
	edge   string
	fields []string
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
//			FromEdge("parent").
//			Unique(),
//	}
//
func Fields(fields ...string) *Index {
	return &Index{fields: fields}
}

// Fields returns the field names of the given index.
func (i Index) Fields() []string {
	return i.fields
}

// FromEdge sets the fields index to be unique under the given edge (sub -graph). For example:
//
//	func (T) Indexes() []ent.Index {
//
//		// Unique "name" field under the "parent" edge.
//		index.Fields("name").
//			FromEdge("parent").
//			Unique(),
//	}
//
func (i *Index) FromEdge(edge string) *Index {
	i.edge = edge
	return i
}

// Edge returns the edge name of the given index.
func (i Index) Edge() string {
	return i.edge
}

// Unique sets the index to be a unique index.
// Note that defining a uniqueness on optional fields won't prevent
// duplicates if one of the column contains NULL values.
func (i *Index) Unique() *Index {
	i.unique = true
	return i
}

// IsUnique indicates if this index is a unique index.
func (i *Index) IsUnique() bool {
	return i.unique
}
