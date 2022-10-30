package gen

import (
	dyschema "entgo.io/ent/dialect/dynamodb/schema"
	"entgo.io/ent/schema/field"
)

// DyTables returns the schema definitions of DynamoDB tables for the graph.
func (g *Graph) DyTables() (all []*dyschema.Table, err error) {
	tables := make(map[string]*dyschema.Table)
	for _, n := range g.Nodes {
		table := dyschema.NewTable(n.Table())
		table.AddAttribute(n.ID.DyAttribute())
		for _, f := range n.Fields {
			if !f.IsEdgeField() {
				table.AddAttribute(f.DyAttribute())
			}
		}
		tables[table.Name] = table
		all = append(all, table)
	}

	// Append primary key to tables.
	for _, n := range g.Nodes {
		table := tables[n.Table()]
		for _, idx := range n.Indexes {
			// Assume each index contains at most two columns:
			// The first column is the partition key for DynamoDB table
			// The second column (if exists) is the sort key.
			if len(idx.Columns) == 1 {
				table.AddKeySchema(&dyschema.KeySchema{
					AttributeName: idx.Columns[0],
					KeyType:       dyschema.KeyTypeHash,
				})
			} else if len(idx.Columns) == 2 {
				table.AddKeySchema(&dyschema.KeySchema{
					AttributeName: idx.Columns[0],
					KeyType:       dyschema.KeyTypeHash,
				})
				table.AddKeySchema(&dyschema.KeySchema{
					AttributeName: idx.Columns[1],
					KeyType:       dyschema.KeyTypeRange,
				})
			}
		}
		for _, e := range n.Edges {
			if e.IsInverse() {
				continue
			}
			if e.Rel.Type == M2M {
				// If there is an edge schema for the association (i.e. edge.Through).
				if e.Through != nil || e.Ref != nil && e.Ref.Through != nil {
					continue
				}
				c1 := &dyschema.Attribute{Name: e.Rel.Columns[0], Type: field.TypeInt}
				if ref := n.ID; ref.UserDefined {
					c1.Type = ref.Type.Type
				}
				c2 := &dyschema.Attribute{Name: e.Rel.Columns[1], Type: field.TypeInt}
				if ref := e.Type.ID; ref.UserDefined {
					c2.Type = ref.Type.Type
				}
				all = append(all, &dyschema.Table{
					Name:       e.Rel.Table,
					Attributes: []*dyschema.Attribute{c1, c2},
					PrimaryKey: []*dyschema.KeySchema{
						{
							AttributeName: c1.Name,
							KeyType:       dyschema.KeyTypeHash,
						},
						{
							AttributeName: c2.Name,
							KeyType:       dyschema.KeyTypeRange,
						},
					},
				})
			}
		}
	}
	return all, nil
}
