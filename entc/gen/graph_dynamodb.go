package gen

import dyschema "entgo.io/ent/dialect/dynamodb/schema"

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
			// Assume each index contains two columns:
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
	}
	return all, nil
}
