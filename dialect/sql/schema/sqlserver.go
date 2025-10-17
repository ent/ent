// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
)

// SQLServer type constants for MS SQL Server.
const (
	TypeBit             = "bit"
	TypeTinyInt         = "tinyint"
	TypeSmallInt        = "smallint"
	TypeInt             = "int"
	TypeBigInt          = "bigint"
	TypeReal            = "real"
	TypeFloat           = "float"
	TypeVarBinary       = "varbinary"
	TypeVarBinaryMAX    = "varbinary(MAX)"
	TypeNVarChar        = "nvarchar"
	TypeNVarCharMAX     = "nvarchar(MAX)"
	TypeDateTime2       = "datetime2"
	TypeUniqueIdentifier = "uniqueidentifier"
)

// SQLServer adapter for Atlas migration engine.
// Note: Full Atlas support for SQL Server migrations is not yet available.
// This provides basic type mapping and dialect support.
type SQLServer struct {
	dialect.Driver
	schema  string
	version string
}

// init loads the SQL Server version from the database for later use in the migration process.
func (d *SQLServer) init(ctx context.Context) error {
	if d.version != "" {
		return nil // already initialized.
	}
	rows := &sql.Rows{}
	if err := d.Query(ctx, "SELECT @@VERSION", []any{}, rows); err != nil {
		return fmt.Errorf("sqlserver: querying server version %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return fmt.Errorf("sqlserver: version variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("sqlserver: scanning server version: %w", err)
	}
	d.version = version
	return nil
}

// tableExist checks if a table exists in the database and current schema.
func (d *SQLServer) tableExist(ctx context.Context, conn dialect.ExecQuerier, name string) (bool, error) {
	query, args := sql.Dialect(dialect.SQLServer).
		Select(sql.Count("*")).From(sql.Table("TABLES").Schema("INFORMATION_SCHEMA")).
		Where(sql.And(
			d.matchSchema(),
			sql.EQ("TABLE_NAME", name),
		)).Query()
	return exist(ctx, conn, query, args...)
}

// matchSchema returns the predicate for matching table schema.
func (d *SQLServer) matchSchema(columns ...string) *sql.Predicate {
	column := "TABLE_SCHEMA"
	if len(columns) > 0 {
		column = columns[0]
	}
	if d.schema != "" {
		return sql.EQ(column, d.schema)
	}
	return sql.EQ(column, sql.Raw("SCHEMA_NAME()"))
}

func (d *SQLServer) atOpen(conn dialect.ExecQuerier) (migrate.Driver, error) {
	// Atlas does not yet have native SQL Server support.
	// Return an error indicating this limitation.
	return nil, fmt.Errorf("sqlserver: Atlas migration engine does not yet support SQL Server. Please use manual migrations or alternative tools")
}

func (d *SQLServer) atTable(t1 *Table, t2 *schema.Table) {
	if t1.Annotation != nil {
		setAtChecks(t1, t2)
	}
}

func (d *SQLServer) supportsDefault(c *Column) bool {
	// SQL Server supports default values for most standard types.
	// Some limitations exist for text, ntext, image, timestamp, and rowversion types.
	switch c.Type {
	case field.TypeTime:
		// SQL Server supports default values for datetime types.
		return true
	default:
		return c.supportDefault()
	}
}

func (d *SQLServer) atTypeC(c1 *Column, c2 *schema.Column) error {
	if c1.SchemaType != nil && c1.SchemaType[dialect.SQLServer] != "" {
		// Use the custom schema type as-is
		c2.Type.Type = &schema.UnsupportedType{T: c1.SchemaType[dialect.SQLServer]}
		return nil
	}
	var t schema.Type
	switch c1.Type {
	case field.TypeBool:
		t = &schema.BoolType{T: TypeBit}
	case field.TypeInt8:
		t = &schema.IntegerType{T: TypeTinyInt}
	case field.TypeUint8:
		t = &schema.IntegerType{T: TypeTinyInt}
	case field.TypeInt16:
		t = &schema.IntegerType{T: TypeSmallInt}
	case field.TypeUint16:
		t = &schema.IntegerType{T: TypeSmallInt}
	case field.TypeInt32:
		t = &schema.IntegerType{T: TypeInt}
	case field.TypeUint32:
		t = &schema.IntegerType{T: TypeInt}
	case field.TypeInt, field.TypeInt64:
		t = &schema.IntegerType{T: TypeBigInt}
	case field.TypeUint, field.TypeUint64:
		t = &schema.IntegerType{T: TypeBigInt}
	case field.TypeBytes:
		size := int64(math.MaxUint16)
		if c1.Size > 0 {
			size = c1.Size
		}
		if size <= math.MaxInt32 {
			t = &schema.BinaryType{T: TypeVarBinary, Size: intPtr(int(size))}
		} else {
			t = &schema.BinaryType{T: TypeVarBinaryMAX}
		}
	case field.TypeJSON:
		// SQL Server stores JSON as NVARCHAR(MAX)
		t = &schema.StringType{T: TypeNVarCharMAX}
	case field.TypeString:
		size := c1.Size
		if size == 0 {
			size = DefaultStringLen
		}
		if size <= 4000 {
			t = &schema.StringType{T: TypeNVarChar, Size: int(size)}
		} else {
			t = &schema.StringType{T: TypeNVarCharMAX}
		}
	case field.TypeFloat32:
		t = &schema.FloatType{T: c1.scanTypeOr(TypeReal)}
	case field.TypeFloat64:
		t = &schema.FloatType{T: c1.scanTypeOr(TypeFloat)}
	case field.TypeTime:
		t = &schema.TimeType{T: c1.scanTypeOr(TypeDateTime2)}
	case field.TypeEnum:
		// SQL Server doesn't have native enum type, use NVARCHAR
		size := int64(DefaultStringLen)
		if c1.Size > 0 {
			size = c1.Size
		}
		t = &schema.StringType{T: TypeNVarChar, Size: int(size)}
	case field.TypeUUID:
		t = &schema.UnsupportedType{T: TypeUniqueIdentifier}
	default:
		// Use as-is for unknown types
		t = &schema.UnsupportedType{T: c1.typ}
	}
	c2.Type.Type = t
	return nil
}

func (d *SQLServer) atUniqueC(t1 *Table, c1 *Column, t2 *schema.Table, c2 *schema.Column) {
	// For UNIQUE columns, SQL Server creates an implicit index with a system-generated name.
	// Check if an explicit unique index is defined for this column.
	for _, idx := range t1.Indexes {
		// Index also defined explicitly, and will be added in atIndexes.
		if idx.Unique && len(idx.Columns) == 1 && idx.Columns[0].Name == c1.Name {
			return
		}
	}
	// Create unique index with a predictable name
	indexName := fmt.Sprintf("UQ_%s_%s", t1.Name, c1.Name)
	t2.AddIndexes(schema.NewUniqueIndex(indexName).AddColumns(c2))
}

func (d *SQLServer) atIncrementC(t *schema.Table, c *schema.Column) {
	// SQL Server uses IDENTITY for auto-increment columns
	// This would require custom Atlas attributes which aren't available yet
	// For now, we'll just mark it in a comment-like way
	if c.Default == nil {
		// In a full implementation, this would add IDENTITY(1,1) attribute
		c.SetComment("AUTO_INCREMENT")
	}
}

func (d *SQLServer) atIncrementT(t *schema.Table, v int64) {
	// Set the identity seed value for the table
	// This would require custom Atlas attributes which aren't available yet
	if v >= 0 {
		t.SetComment(fmt.Sprintf("IDENTITY_SEED=%d", v))
	}
}

func (d *SQLServer) atIndex(idx1 *Index, t2 *schema.Table, idx2 *schema.Index) error {
	for _, c1 := range idx1.Columns {
		c2, ok := t2.Column(c1.Name)
		if !ok {
			return fmt.Errorf("unexpected index %q column: %q", idx1.Name, c1.Name)
		}
		part := &schema.IndexPart{C: c2}
		idx2.AddParts(part)
	}
	// Index type and predicate handling would require custom Atlas attributes
	return nil
}

func (*SQLServer) atTypeRangeSQL(ts ...string) string {
	for i := range ts {
		ts[i] = fmt.Sprintf("(N'%s')", ts[i])
	}
	return fmt.Sprintf("INSERT INTO [%s] ([type]) VALUES %s", TypeTable, strings.Join(ts, ", "))
}

// intPtr returns a pointer to an int value
func intPtr(i int) *int {
return &i
}
