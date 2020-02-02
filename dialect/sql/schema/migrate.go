// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"crypto/md5"
	"fmt"
	"math"
	"sort"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
)

const (
	// TypeTable defines the table name holding the type information.
	TypeTable = "ent_types"
	// MaxTypes defines the max number of types can be created when
	// defining universal ids. The left 16-bits are reserved.
	MaxTypes = math.MaxUint16
)

// MigrateOption allows for managing schema configuration using functional options.
type MigrateOption func(m *Migrate)

// WithGlobalUniqueID sets the universal ids options to the migration.
// Defaults to false.
func WithGlobalUniqueID(b bool) MigrateOption {
	return func(m *Migrate) {
		m.universalID = b
	}
}

// WithDropColumn sets the columns dropping option to the migration.
// Defaults to false.
func WithDropColumn(b bool) MigrateOption {
	return func(m *Migrate) {
		m.dropColumns = b
	}
}

// WithDropIndex sets the indexes dropping option to the migration.
// Defaults to false.
func WithDropIndex(b bool) MigrateOption {
	return func(m *Migrate) {
		m.dropIndexes = b
	}
}

// Migrate runs the migrations logic for the SQL dialects.
type Migrate struct {
	sqlDialect
	universalID bool     // global unique ids.
	dropColumns bool     // drop deleted columns.
	dropIndexes bool     // drop deleted indexes.
	typeRanges  []string // types order by their range.
}

// NewMigrate create a migration structure for the given SQL driver.
func NewMigrate(d dialect.Driver, opts ...MigrateOption) (*Migrate, error) {
	m := &Migrate{}
	switch d.Dialect() {
	case dialect.MySQL:
		m.sqlDialect = &MySQL{Driver: d}
	case dialect.SQLite:
		m.sqlDialect = &SQLite{Driver: d}
	case dialect.Postgres:
		m.sqlDialect = &Postgres{Driver: d}
	default:
		return nil, fmt.Errorf("sql/schema: unsupported dialect %q", d.Dialect())
	}
	for _, opt := range opts {
		opt(m)
	}
	return m, nil
}

// Create creates all schema resources in the database. It works in an "append-only"
// mode, which means, it only create tables, append column to tables or modifying column type.
//
// Column can be modified by turning into a NULL from NOT NULL, or having a type conversion not
// resulting data altering. From example, changing varchar(255) to varchar(120) is invalid, but
// changing varchar(120) to varchar(255) is valid. For more info, see the convert function below.
//
// Note that SQLite dialect does not support (this moment) the "append-only" mode describe above,
// since it's used only for testing.
func (m *Migrate) Create(ctx context.Context, tables ...*Table) error {
	tx, err := m.Tx(ctx)
	if err != nil {
		return err
	}
	if err := m.init(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	if m.universalID {
		if err := m.types(ctx, tx); err != nil {
			return rollback(tx, err)
		}
	}
	if err := m.create(ctx, tx, tables...); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func (m *Migrate) create(ctx context.Context, tx dialect.Tx, tables ...*Table) error {
	for _, t := range tables {
		m.setupTable(t)
		switch exist, err := m.tableExist(ctx, tx, t.Name); {
		case err != nil:
			return err
		case exist:
			curr, err := m.table(ctx, tx, t.Name)
			if err != nil {
				return err
			}
			if err := m.fixture(ctx, tx, curr, t); err != nil {
				return err
			}
			change, err := m.changeSet(curr, t)
			if err != nil {
				return err
			}
			if err := m.apply(ctx, tx, t.Name, change); err != nil {
				return err
			}
		default: // !exist
			query, args := m.tBuilder(t).Query()
			if err := tx.Exec(ctx, query, args, nil); err != nil {
				return fmt.Errorf("create table %q: %v", t.Name, err)
			}
			// if global unique identifier is enabled and it's not a relation table,
			// allocate a range for the table pk.
			if m.universalID && len(t.PrimaryKey) == 1 {
				if err := m.allocPKRange(ctx, tx, t); err != nil {
					return err
				}
			}
			// indexes.
			for _, idx := range t.Indexes {
				query, args := m.addIndex(idx, t.Name).Query()
				if err := tx.Exec(ctx, query, args, nil); err != nil {
					return fmt.Errorf("create index %q: %v", idx.Name, err)
				}
			}
		}
	}
	// create foreign keys after tables were created/altered,
	// because circular foreign-key constraints are possible.
	for _, t := range tables {
		if len(t.ForeignKeys) == 0 {
			continue
		}
		fks := make([]*ForeignKey, 0, len(t.ForeignKeys))
		for _, fk := range t.ForeignKeys {
			exist, err := m.fkExist(ctx, tx, fk.Symbol)
			if err != nil {
				return err
			}
			if !exist {
				fks = append(fks, fk)
			}
		}
		if len(fks) == 0 {
			continue
		}
		b := sql.Dialect(m.Dialect()).AlterTable(t.Name)
		for _, fk := range fks {
			b.AddForeignKey(fk.DSL())
		}
		query, args := b.Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("create foreign keys for %q: %v", t.Name, err)
		}
	}
	return nil
}

// apply applies changes on the given table.
func (m *Migrate) apply(ctx context.Context, tx dialect.Tx, table string, change *changes) error {
	// Constraints should be dropped before dropping columns, because if a column
	// is a part of multi-column constraints (like, unique index), ALTER TABLE
	// might fail if the intermediate state violates the constraints.
	if m.dropIndexes {
		if pr, ok := m.sqlDialect.(preparer); ok {
			if err := pr.prepare(ctx, tx, change, table); err != nil {
				return err
			}
		}
		for _, idx := range change.index.drop {
			if err := m.dropIndex(ctx, tx, idx, table); err != nil {
				return fmt.Errorf("drop index of table %q: %v", table, err)
			}
		}
	}
	b := sql.Dialect(m.Dialect()).AlterTable(table)
	for _, c := range change.column.add {
		b.AddColumn(m.addColumn(c))
	}
	for _, c := range change.column.modify {
		b.ModifyColumns(m.alterColumn(c)...)
	}
	if m.dropColumns {
		for _, c := range change.column.drop {
			b.DropColumn(sql.Dialect(m.Dialect()).Column(c.Name))
		}
	}
	// If there's actual action to execute on ALTER TABLE.
	if len(b.Queries) != 0 {
		query, args := b.Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("alter table %q: %v", table, err)
		}
	}
	for _, idx := range change.index.add {
		query, args := m.addIndex(idx, table).Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("create index %q: %v", table, err)
		}
	}
	return nil
}

// changes to apply on existing table.
type changes struct {
	// column changes.
	column struct {
		add    []*Column
		drop   []*Column
		modify []*Column
	}
	// index changes.
	index struct {
		add  Indexes
		drop Indexes
	}
}

// dropColumn returns the dropped column by name (if any).
func (c *changes) dropColumn(name string) (*Column, bool) {
	for _, col := range c.column.drop {
		if col.Name == name {
			return col, true
		}
	}
	return nil, false
}

// changeSet returns a changes object to be applied on existing table.
// It fails if one of the changes is invalid.
func (m *Migrate) changeSet(curr, new *Table) (*changes, error) {
	change := &changes{}
	// pks.
	if len(curr.PrimaryKey) != len(new.PrimaryKey) {
		return nil, fmt.Errorf("cannot change primary key for table: %q", curr.Name)
	}
	sort.Slice(new.PrimaryKey, func(i, j int) bool { return new.PrimaryKey[i].Name < new.PrimaryKey[j].Name })
	sort.Slice(curr.PrimaryKey, func(i, j int) bool { return curr.PrimaryKey[i].Name < curr.PrimaryKey[j].Name })
	for i := range curr.PrimaryKey {
		if curr.PrimaryKey[i].Name != new.PrimaryKey[i].Name {
			return nil, fmt.Errorf("cannot change primary key for table: %q", curr.Name)
		}
	}
	// add or modify columns.
	for _, c1 := range new.Columns {
		// ignore primary keys.
		if c1.PrimaryKey() {
			continue
		}
		switch c2, ok := curr.column(c1.Name); {
		case !ok:
			change.column.add = append(change.column.add, c1)
		// modify a non-unique column to unique.
		case c1.Unique && !c2.Unique:
			change.index.add.append(&Index{
				Name:    c1.Name,
				Unique:  true,
				Columns: []*Column{c1},
				columns: []string{c1.Name},
			})
		// modify a unique column to non-unique.
		case !c1.Unique && c2.Unique:
			idx, ok := curr.index(c2.Name)
			if !ok {
				return nil, fmt.Errorf("missing index to drop for column %q", c2.Name)
			}
			change.index.drop.append(idx)
		// extending column types.
		case m.cType(c1) != m.cType(c2):
			if !c2.ConvertibleTo(c1) {
				return nil, fmt.Errorf("changing column type for %q is invalid (%s != %s)", c1.Name, m.cType(c1), m.cType(c2))
			}
			fallthrough
		// change nullability of a column.
		case c1.Nullable != c2.Nullable:
			change.column.modify = append(change.column.modify, c1)
		}
	}

	// drop columns.
	for _, c1 := range curr.Columns {
		// if a column was dropped, multi-columns indexes that are associated with this column will
		// no longer behave the same. Therefore, these indexes should be dropped too. There's no need
		// to do it explicitly (here), because entc will remove them from the schema specification,
		// and they will be dropped in the block below.
		if _, ok := new.column(c1.Name); !ok {
			change.column.drop = append(change.column.drop, c1)
		}
	}

	// add or modify indexes.
	for _, idx1 := range new.Indexes {
		switch idx2, ok := curr.index(idx1.Name); {
		case !ok:
			change.index.add.append(idx1)
		// changing index cardinality require drop and create.
		case idx1.Unique != idx2.Unique:
			change.index.drop.append(idx2)
			change.index.add.append(idx1)
		}
	}

	// drop indexes.
	for _, idx1 := range curr.Indexes {
		_, ok1 := new.fk(idx1.Name)
		_, ok2 := new.index(idx1.Name)
		if !ok1 && !ok2 {
			change.index.drop.append(idx1)
		}
	}
	return change, nil
}

// fixture is a special migration code for renaming foreign-key columns (issue-#285).
func (m *Migrate) fixture(ctx context.Context, tx dialect.Tx, curr, new *Table) error {
	d, ok := m.sqlDialect.(fkRenamer)
	if !ok {
		return nil
	}
	rename := make(map[string]*Index)
	for _, fk := range new.ForeignKeys {
		ok, err := m.fkExist(ctx, tx, fk.Symbol)
		if err != nil {
			return fmt.Errorf("checking foreign-key existence %q: %v", fk.Symbol, err)
		}
		if !ok {
			continue
		}
		column, err := m.fkColumn(ctx, tx, fk)
		if err != nil {
			return err
		}
		newcol := fk.Columns[0]
		if column == newcol.Name {
			continue
		}
		query, args := d.renameColumn(curr, &Column{Name: column}, newcol).Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("rename column %q: %v", column, err)
		}
		prev, ok := curr.column(column)
		if !ok {
			continue
		}
		// Find all indexes that ~maybe need to be renamed.
		for _, idx := range prev.indexes {
			switch _, ok := new.index(idx.Name); {
			// Ignore indexes that exist in the schema, PKs.
			case ok || idx.primary:
			// Index that was created implicitly for a unique
			// column needs to be renamed to the column name.
			case d.isImplicitIndex(idx, prev):
				idx2 := &Index{Name: newcol.Name, Unique: true, Columns: []*Column{newcol}}
				query, args := d.renameIndex(curr, idx, idx2).Query()
				if err := tx.Exec(ctx, query, args, nil); err != nil {
					return fmt.Errorf("rename index %q: %v", prev.Name, err)
				}
				idx.Name = idx2.Name
			default:
				rename[idx.Name] = idx
			}
		}
		// Update the name of the loaded column, so `changeSet` won't create it.
		prev.Name = newcol.Name
	}
	// Go over the indexes that need to be renamed
	// and find their ~identical in the new schema.
	for _, idx := range rename {
	Find:
		// Find its ~identical in the new schema, and rename it
		// if it doesn't exist.
		for _, idx2 := range new.Indexes {
			if _, ok := curr.index(idx2.Name); ok {
				continue
			}
			if idx.sameAs(idx2) {
				query, args := d.renameIndex(curr, idx, idx2).Query()
				if err := tx.Exec(ctx, query, args, nil); err != nil {
					return fmt.Errorf("rename index %q: %v", idx.Name, err)
				}
				idx.Name = idx2.Name
				break Find
			}
		}
	}
	return nil
}

// types loads the type list from the database.
// If the table does not create, it will create one.
func (m *Migrate) types(ctx context.Context, tx dialect.Tx) error {
	exists, err := m.tableExist(ctx, tx, TypeTable)
	if err != nil {
		return err
	}
	if !exists {
		t := NewTable(TypeTable).
			AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}).
			AddColumn(&Column{Name: "type", Type: field.TypeString, Unique: true})
		query, args := m.tBuilder(t).Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("create types table: %v", err)
		}
		return nil
	}
	rows := &sql.Rows{}
	query, args := sql.Dialect(m.Dialect()).
		Select("type").From(sql.Table(TypeTable)).OrderBy(sql.Asc("id")).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return fmt.Errorf("query types table: %v", err)
	}
	defer rows.Close()
	return sql.ScanSlice(rows, &m.typeRanges)
}

func (m *Migrate) allocPKRange(ctx context.Context, tx dialect.Tx, t *Table) error {
	id := -1
	// if the table re-created, re-use its range from the past.
	for i, name := range m.typeRanges {
		if name == t.Name {
			id = i
			break
		}
	}
	// allocate a new id-range.
	if id == -1 {
		if len(m.typeRanges) > MaxTypes {
			return fmt.Errorf("max number of types exceeded: %d", MaxTypes)
		}
		query, args := sql.Dialect(m.Dialect()).
			Insert(TypeTable).Columns("type").Values(t.Name).Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("insert into type: %v", err)
		}
		id = len(m.typeRanges)
		m.typeRanges = append(m.typeRanges, t.Name)
	}
	// set the id offset for table.
	return m.setRange(ctx, tx, t.Name, id<<32)
}

// fkColumn returns the column name of a foreign-key.
func (m *Migrate) fkColumn(ctx context.Context, tx dialect.Tx, fk *ForeignKey) (string, error) {
	t1 := sql.Table("INFORMATION_SCHEMA.KEY_COLUMN_USAGE AS t1").Unquote().As("t1")
	t2 := sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS AS t2").Unquote().As("t2")
	query, args := sql.Dialect(m.Dialect()).
		Select("column_name").
		From(t1).
		Join(t2).
		On(t1.C("constraint_name"), t2.C("constraint_name")).
		Where(sql.And(
			sql.EQ(t2.C("constraint_type"), sql.Raw("'FOREIGN KEY'")),
			sql.EQ(t2.C("table_schema"), m.sqlDialect.(fkRenamer).tableSchema()),
			sql.EQ(t1.C("table_schema"), m.sqlDialect.(fkRenamer).tableSchema()),
			sql.EQ(t2.C("constraint_name"), fk.Symbol),
		)).
		Query()
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return "", fmt.Errorf("reading foreign-key %q column: %v", fk.Symbol, err)
	}
	defer rows.Close()
	column, err := sql.ScanString(rows)
	if err != nil {
		return "", fmt.Errorf("scanning foreign-key %q column: %v", fk.Symbol, err)
	}
	return column, nil
}

// setup ensures the table is configured properly, like table columns
// are linked to their indexes, and PKs columns are defined.
func (m *Migrate) setupTable(t *Table) {
	if t.columns == nil {
		t.columns = make(map[string]*Column, len(t.Columns))
	}
	for _, c := range t.Columns {
		t.columns[c.Name] = c
	}
	for _, idx := range t.Indexes {
		idx.Name = m.symbol(idx.Name)
		for _, c := range idx.Columns {
			c.indexes.append(idx)
		}
	}
	for _, pk := range t.PrimaryKey {
		c := t.columns[pk.Name]
		c.Key = PrimaryKey
		pk.Key = PrimaryKey
	}
	for _, fk := range t.ForeignKeys {
		fk.Symbol = m.symbol(fk.Symbol)
	}
}

// symbol makes sure the symbol length is not longer than the maxlength in the dialect.
func (m *Migrate) symbol(name string) string {
	size := 64
	if m.Dialect() == dialect.Postgres {
		size = 63
	}
	if len(name) <= size {
		return name
	}
	return fmt.Sprintf("%s_%x", name[:size-33], md5.Sum([]byte(name)))
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	err = fmt.Errorf("sql/schema: %v", err)
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}

// exist checks if the given COUNT query returns a value >= 1.
func exist(ctx context.Context, tx dialect.Tx, query string, args ...interface{}) (bool, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return false, fmt.Errorf("reading schema information %v", err)
	}
	defer rows.Close()
	n, err := sql.ScanInt(rows)
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

type sqlDialect interface {
	dialect.Driver
	init(context.Context, dialect.Tx) error
	table(context.Context, dialect.Tx, string) (*Table, error)
	tableExist(context.Context, dialect.Tx, string) (bool, error)
	fkExist(context.Context, dialect.Tx, string) (bool, error)
	setRange(context.Context, dialect.Tx, string, int) error
	dropIndex(context.Context, dialect.Tx, *Index, string) error
	// table, column and index builder per dialect.
	cType(*Column) string
	tBuilder(*Table) *sql.TableBuilder
	addColumn(*Column) *sql.ColumnBuilder
	alterColumn(*Column) []*sql.ColumnBuilder
	addIndex(*Index, string) *sql.IndexBuilder
}

type preparer interface {
	prepare(context.Context, dialect.Tx, *changes, string) error
}

// fkRenamer is used by the fixture migration (to solve #285),
// and it's implemented by the different dialects for renaming FKs.
type fkRenamer interface {
	tableSchema() sql.Querier
	isImplicitIndex(*Index, *Column) bool
	renameIndex(*Table, *Index, *Index) sql.Querier
	renameColumn(*Table, *Column, *Column) sql.Querier
}
