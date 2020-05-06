package schema

import (
	"context"
	"fmt"
	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Re-usable functions that use information schema directly
// fkNames returns the foreign-key names of a column.
func fkNames(ctx context.Context, tx dialect.Tx, schema sql.Querier, table, column string) ([]string, error) {
	query, args := sql.Select("CONSTRAINT_NAME").From(sql.Table("INFORMATION_SCHEMA.KEY_COLUMN_USAGE").Unquote()).
		Where(sql.
			EQ("TABLE_NAME", table).
			And().EQ("COLUMN_NAME", column).
			// NULL for unique and primary-key constraints.
			And().NotNull("POSITION_IN_UNIQUE_CONSTRAINT").
			And().EQ("TABLE_SCHEMA", schema),
		).
		Query()
	var (
		names []string
		rows  = &sql.Rows{}
	)
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("MSSQL: reading constraint names %v", err)
	}
	defer rows.Close()
	if err := sql.ScanSlice(rows, &names); err != nil {
		return nil, err
	}
	return names, nil
}
