// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqlite"
)

type (
	// SQLite adapter for Atlas migration engine.
	SQLite struct {
		dialect.Driver
		WithForeignKeys bool
	}
	// SQLiteTx implements dialect.Tx.
	SQLiteTx struct {
		dialect.Tx
		commit   func() error // Override Commit to toggle foreign keys back on after Commit.
		rollback func() error // Override Rollback to toggle foreign keys back on after Rollback.
	}
)

// Tx implements opens a transaction.
func (d *SQLite) Tx(ctx context.Context) (dialect.Tx, error) {
	db := &db{d}
	if _, err := db.ExecContext(ctx, "PRAGMA foreign_keys = off"); err != nil {
		return nil, fmt.Errorf("sqlite: set 'foreign_keys = off': %w", err)
	}
	t, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	tx := &tx{t}
	cm, err := sqlite.CommitFunc(ctx, db, tx, true)
	if err != nil {
		return nil, err
	}
	return &SQLiteTx{Tx: t, commit: cm, rollback: sqlite.RollbackFunc(ctx, db, tx, true)}, nil
}

// Commit ensures foreign keys are toggled back on after commit.
func (tx *SQLiteTx) Commit() error {
	return tx.commit()
}

// Rollback ensures foreign keys are toggled back on after rollback.
func (tx *SQLiteTx) Rollback() error {
	return tx.rollback()
}

// init makes sure that foreign_keys support is enabled.
func (d *SQLite) init(ctx context.Context) error {
	on, err := exist(ctx, d, "PRAGMA foreign_keys")
	if err != nil {
		return fmt.Errorf("sqlite: check foreign_keys pragma: %w", err)
	}
	if !on {
		// foreign_keys pragma is off, either enable it by execute "PRAGMA foreign_keys=ON"
		// or add the following parameter in the connection string "_fk=1".
		return fmt.Errorf("sqlite: foreign_keys pragma is off: missing %q in the connection string", "_fk=1")
	}
	return nil
}

func (d *SQLite) tableExist(ctx context.Context, conn dialect.ExecQuerier, name string) (bool, error) {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_master")).
		Where(sql.And(
			sql.EQ("type", "table"),
			sql.EQ("name", name),
		)).
		Query()
	return exist(ctx, conn, query, args...)
}

func (d *SQLite) atOpen(conn dialect.ExecQuerier) (migrate.Driver, error) {
	return sqlite.Open(&db{ExecQuerier: conn})
}

func (d *SQLite) atTable(t1 *Table, t2 *schema.Table) {
	if t1.Annotation != nil {
		setAtChecks(t1, t2)
	}
}

func (d *SQLite) supportsDefault(*Column) bool {
	// SQLite supports default values for all standard types.
	return true
}

func (d *SQLite) atTypeC(c1 *Column, c2 *schema.Column) error {
	if c1.SchemaType != nil && c1.SchemaType[dialect.SQLite] != "" {
		t, err := sqlite.ParseType(strings.ToLower(c1.SchemaType[dialect.SQLite]))
		if err != nil {
			return err
		}
		c2.Type.Type = t
		return nil
	}
	var t schema.Type
	switch c1.Type {
	case field.TypeBool:
		t = &schema.BoolType{T: "bool"}
	case field.TypeInt8, field.TypeUint8, field.TypeInt16, field.TypeUint16, field.TypeInt32,
		field.TypeUint32, field.TypeUint, field.TypeInt, field.TypeInt64, field.TypeUint64:
		t = &schema.IntegerType{T: sqlite.TypeInteger}
	case field.TypeBytes:
		t = &schema.BinaryType{T: sqlite.TypeBlob}
	case field.TypeString, field.TypeEnum:
		// SQLite does not impose any length restrictions on
		// the length of strings, BLOBs or numeric values.
		t = &schema.StringType{T: sqlite.TypeText}
	case field.TypeFloat32, field.TypeFloat64:
		t = &schema.FloatType{T: sqlite.TypeReal}
	case field.TypeTime:
		t = &schema.TimeType{T: "datetime"}
	case field.TypeJSON:
		t = &schema.JSONType{T: "json"}
	case field.TypeUUID:
		t = &sqlite.UUIDType{T: "uuid"}
	case field.TypeOther:
		t = &schema.UnsupportedType{T: c1.typ}
	default:
		t, err := sqlite.ParseType(strings.ToLower(c1.typ))
		if err != nil {
			return err
		}
		c2.Type.Type = t
	}
	c2.Type.Type = t
	return nil
}

func (d *SQLite) atUniqueC(t1 *Table, c1 *Column, t2 *schema.Table, c2 *schema.Column) {
	// For UNIQUE columns, SQLite create an implicit index named
	// "sqlite_autoindex_<table>_<i>". Ent uses the PostgreSQL approach
	// in its migration, and name these indexes as "<table>_<column>_key".
	for _, idx := range t1.Indexes {
		// Index also defined explicitly, and will be add in atIndexes.
		if idx.Unique && d.atImplicitIndexName(idx, t1, c1) {
			return
		}
	}
	t2.AddIndexes(schema.NewUniqueIndex(fmt.Sprintf("%s_%s_key", t2.Name, c1.Name)).AddColumns(c2))
}

func (d *SQLite) atImplicitIndexName(idx *Index, t1 *Table, c1 *Column) bool {
	if idx.Name == c1.Name {
		return true
	}
	p := fmt.Sprintf("sqlite_autoindex_%s_", t1.Name)
	if !strings.HasPrefix(idx.Name, p) {
		return false
	}
	i, err := strconv.ParseInt(strings.TrimPrefix(idx.Name, p), 10, 64)
	return err == nil && i > 0
}

func (d *SQLite) atIncrementC(t *schema.Table, c *schema.Column) {
	if c.Default != nil {
		t.Attrs = removeAttr(t.Attrs, reflect.TypeOf(&sqlite.AutoIncrement{}))
	} else {
		c.AddAttrs(&sqlite.AutoIncrement{})
	}
}

func (d *SQLite) atIncrementT(t *schema.Table, v int64) {
	t.AddAttrs(&sqlite.AutoIncrement{Seq: v})
}

func (d *SQLite) atIndex(idx1 *Index, t2 *schema.Table, idx2 *schema.Index) error {
	for _, c1 := range idx1.Columns {
		c2, ok := t2.Column(c1.Name)
		if !ok {
			return fmt.Errorf("unexpected index %q column: %q", idx1.Name, c1.Name)
		}
		idx2.AddParts(&schema.IndexPart{C: c2})
	}
	if idx1.Annotation != nil && idx1.Annotation.Where != "" {
		idx2.AddAttrs(&sqlite.IndexPredicate{P: idx1.Annotation.Where})
	}
	return nil
}

func (*SQLite) atTypeRangeSQL(ts ...string) string {
	for i := range ts {
		ts[i] = fmt.Sprintf("('%s')", ts[i])
	}
	return fmt.Sprintf("INSERT INTO `%s` (`type`) VALUES %s", TypeTable, strings.Join(ts, ", "))
}

type tx struct {
	dialect.Tx
}

func (tx *tx) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	return rows.ColumnScanner.(*stdsql.Rows), nil
}

func (tx *tx) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	var r stdsql.Result
	if err := tx.Exec(ctx, query, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}
