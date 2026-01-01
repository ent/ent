// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
)

// SQLServer adapter for Atlas migration engine.
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
		return fmt.Errorf("querying server version %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return fmt.Errorf("@@VERSION variable was not found")
	}
	var version string
	if err := rows.Scan(&version); err != nil {
		return fmt.Errorf("scanning version: %w", err)
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
	// The open-source Atlas Go library (ariga.io/atlas) does not include a SQL Server driver.
	// SQL Server support in Atlas is a Pro/Cloud-only feature available through the Atlas CLI
	// but not exposed in the open-source Go module.
	// See: https://atlasgo.io/guides/mssql
	return nil, fmt.Errorf("sql/schema: atlas driver for SQL Server is not available in the open-source library; " +
		"use the Atlas CLI with 'atlas login' for SQL Server migrations, or use manual migrations")
}

func (d *SQLServer) atTable(t1 *Table, t2 *schema.Table) {
	if t1.Annotation != nil {
		setAtChecks(t1, t2)
	}
}

func (d *SQLServer) supportsDefault(*Column) bool {
	// SQL Server supports default values for all standard types.
	return true
}

func (d *SQLServer) atTypeC(c1 *Column, c2 *schema.Column) error {
	if c1.SchemaType != nil && c1.SchemaType[dialect.SQLServer] != "" {
		// For SQL Server, we set the type as unsupported since atlas doesn't have
		// a parser for SQL Server types yet.
		c2.Type.Type = &schema.UnsupportedType{T: c1.SchemaType[dialect.SQLServer]}
		return nil
	}
	var t schema.Type
	switch c1.Type {
	case field.TypeBool:
		t = &schema.BoolType{T: "bit"}
	case field.TypeUint8, field.TypeInt8:
		t = &schema.IntegerType{T: "tinyint"}
	case field.TypeInt16, field.TypeUint16:
		t = &schema.IntegerType{T: "smallint"}
	case field.TypeInt32, field.TypeUint32:
		t = &schema.IntegerType{T: "int"}
	case field.TypeInt, field.TypeUint, field.TypeInt64, field.TypeUint64:
		t = &schema.IntegerType{T: "bigint"}
	case field.TypeFloat32:
		t = &schema.FloatType{T: c1.scanTypeOr("real")}
	case field.TypeFloat64:
		t = &schema.FloatType{T: c1.scanTypeOr("float")}
	case field.TypeBytes:
		size := c1.Size
		if size == 0 {
			size = 8000 // SQL Server max for varbinary
		}
		if size > 8000 {
			t = &schema.BinaryType{T: "varbinary(max)"}
		} else {
			t = &schema.BinaryType{T: fmt.Sprintf("varbinary(%d)", size)}
		}
	case field.TypeUUID:
		t = &schema.UUIDType{T: "uniqueidentifier"}
	case field.TypeJSON:
		// SQL Server stores JSON as nvarchar(max)
		t = &schema.StringType{T: "nvarchar(max)"}
	case field.TypeString:
		size := c1.Size
		if size == 0 {
			size = DefaultStringLen
		}
		// SQL Server uses nvarchar for unicode strings
		if size > 4000 {
			t = &schema.StringType{T: "nvarchar(max)"}
		} else {
			t = &schema.StringType{T: "nvarchar", Size: int(size)}
		}
	case field.TypeTime:
		t = &schema.TimeType{T: c1.scanTypeOr("datetime2")}
	case field.TypeEnum:
		// SQL Server doesn't have native enum type, use nvarchar
		t = &schema.StringType{T: "nvarchar", Size: 255}
	case field.TypeOther:
		t = &schema.UnsupportedType{T: c1.typ}
	default:
		// For unknown types, mark as unsupported
		t = &schema.UnsupportedType{T: c1.typ}
	}
	c2.Type.Type = t
	return nil
}

func (d *SQLServer) atUniqueC(t1 *Table, c1 *Column, t2 *schema.Table, c2 *schema.Column) {
	// For UNIQUE columns, SQL Server creates an implicit index named
	// "UQ__<table>__<hash>".
	for _, idx := range t1.Indexes {
		// Index also defined explicitly, and will be added in atIndexes.
		if idx.Unique && d.atImplicitIndexName(idx, t1, c1) {
			return
		}
	}
	idx := schema.NewUniqueIndex(fmt.Sprintf("%s_%s_key", t1.Name, c1.Name)).AddColumns(c2)
	// Note: For nullable columns, SQL Server unique constraints by default don't allow
	// duplicate NULLs. A WHERE predicate would be needed but atlas/schema doesn't expose
	// a generic IndexPredicate - this is handled by the Atlas CLI instead.
	t2.AddIndexes(idx)
}

func (d *SQLServer) atImplicitIndexName(idx *Index, t1 *Table, c1 *Column) bool {
	// SQL Server generates index names like UQ__tablename__hash
	return strings.HasPrefix(idx.Name, "UQ__")
}

func (d *SQLServer) atIncrementC(t *schema.Table, c *schema.Column) {
	// SQL Server uses IDENTITY for auto-increment columns.
	// Note: The Atlas sqlserver package is not available in the open-source version,
	// so we cannot add the Identity attribute here. Migrations are handled via Atlas CLI.
}

func (d *SQLServer) atPrimaryKey(pk *schema.Index) {
	// SQL Server primary keys are CLUSTERED by default.
	// Note: The Atlas sqlserver package is not available in the open-source version,
	// so we cannot add the IndexType attribute here. Migrations are handled via Atlas CLI.
}

func (d *SQLServer) atIncrementT(t *schema.Table, v int64) {
	// SQL Server uses IDENTITY(start, increment) for auto-increment columns.
	// Note: The Atlas sqlserver package is not available in the open-source version,
	// so we cannot update the Identity seed here. Migrations are handled via Atlas CLI.
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
	// Note: The Atlas sqlserver package is not available in the open-source version,
	// so we cannot add IndexType attributes here. Migrations are handled via Atlas CLI.
	return nil
}

func (*SQLServer) atTypeRangeSQL(ts ...string) string {
	for i := range ts {
		ts[i] = fmt.Sprintf("('%s')", ts[i])
	}
	return fmt.Sprintf("INSERT INTO [%s] ([type]) VALUES %s", TypeTable, strings.Join(ts, ", "))
}

// verifyRange verifies the id range for a table if global unique id is enabled.
func (d *SQLServer) verifyRange(ctx context.Context, conn dialect.ExecQuerier, t *Table, expected int64) error {
	if expected == 0 {
		return nil
	}
	rows := &sql.Rows{}
	query := fmt.Sprintf("SELECT IDENT_CURRENT('[%s]')", t.Name)
	if err := conn.Query(ctx, query, []any{}, rows); err != nil {
		return fmt.Errorf("query identity value: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return nil
	}
	var current int64
	if err := rows.Scan(&current); err != nil {
		return fmt.Errorf("scan identity value: %w", err)
	}
	// If the current identity is less than expected, reseed.
	if current < expected {
		reseed := fmt.Sprintf("DBCC CHECKIDENT ('[%s]', RESEED, %d)", t.Name, expected)
		if err := conn.Exec(ctx, reseed, []any{}, nil); err != nil {
			return fmt.Errorf("reseed identity: %w", err)
		}
	}
	return nil
}
