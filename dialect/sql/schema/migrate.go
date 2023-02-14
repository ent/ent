// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

const (
	// TypeTable defines the table name holding the type information.
	TypeTable = "ent_types"

	// MaxTypes defines the max number of types can be created when
	// defining universal ids. The left 16-bits are reserved.
	MaxTypes = math.MaxUint16
)

// NewTypesTable returns a new table for holding the global-id information.
func NewTypesTable() *Table {
	return NewTable(TypeTable).
		AddPrimary(&Column{Name: "id", Type: field.TypeUint, Increment: true}).
		AddColumn(&Column{Name: "type", Type: field.TypeString, Unique: true})
}

// MigrateOption allows configuring Atlas using functional arguments.
type MigrateOption func(*Atlas)

// WithGlobalUniqueID sets the universal ids options to the migration.
// Defaults to false.
func WithGlobalUniqueID(b bool) MigrateOption {
	return func(a *Atlas) {
		a.universalID = b
	}
}

// WithDropColumn sets the columns dropping option to the migration.
// Defaults to false.
func WithDropColumn(b bool) MigrateOption {
	return func(a *Atlas) {
		a.dropColumns = b
	}
}

// WithDropIndex sets the indexes dropping option to the migration.
// Defaults to false.
func WithDropIndex(b bool) MigrateOption {
	return func(a *Atlas) {
		a.dropIndexes = b
	}
}

// WithFixture sets the foreign-key renaming option to the migration when upgrading
// sqlDialect from v0.1.0 (issue-#285). Defaults to false.
//
// Deprecated: This option is no longer needed with the Atlas based
// migration engine, which now is the default.
func WithFixture(b bool) MigrateOption {
	return func(a *Atlas) {
		a.withFixture = b
	}
}

// WithForeignKeys enables creating foreign-key in ddl. Defaults to true.
func WithForeignKeys(b bool) MigrateOption {
	return func(a *Atlas) {
		a.withForeignKeys = b
	}
}

// WithHooks adds a list of hooks to the schema migration.
func WithHooks(hooks ...Hook) MigrateOption {
	return func(a *Atlas) {
		a.hooks = append(a.hooks, hooks...)
	}
}

type (
	// Creator is the interface that wraps the Create method.
	Creator interface {
		// Create creates the given tables in the database. See Migrate.Create for more details.
		Create(context.Context, ...*Table) error
	}

	// The CreateFunc type is an adapter to allow the use of ordinary function as Creator.
	// If f is a function with the appropriate signature, CreateFunc(f) is a Creator that calls f.
	CreateFunc func(context.Context, ...*Table) error

	// Hook defines the "create middleware". A function that gets a Creator and returns a Creator.
	// For example:
	//
	//	hook := func(next schema.Creator) schema.Creator {
	//		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
	//			fmt.Println("Tables:", tables)
	//			return next.Create(ctx, tables...)
	//		})
	//	}
	//
	Hook func(Creator) Creator
)

// Create calls f(ctx, tables...).
func (f CreateFunc) Create(ctx context.Context, tables ...*Table) error {
	return f(ctx, tables...)
}

// Migrate runs the migration logic for the SQL dialects.
//
// Deprecated: Use the new Atlas struct instead.
type Migrate struct {
	sqlDialect
	atlas *Atlas // Atlas this Migrate is based on

	universalID     bool     // global unique ids
	dropColumns     bool     // drop deleted columns
	dropIndexes     bool     // drop deleted indexes
	withFixture     bool     // with fks rename fixture
	withForeignKeys bool     // with foreign keys
	typeRanges      []string // types order by their range
	hooks           []Hook   // hooks to apply before creation
}

// Create creates all schema resources in the database. It works in an "append-only"
// mode, which means, it only creates tables, appends columns to tables or modifies column types.
//
// Column can be modified by turning into a NULL from NOT NULL, or having a type conversion not
// resulting data altering. From example, changing varchar(255) to varchar(120) is invalid, but
// changing varchar(120) to varchar(255) is valid. For more info, see the convert function below.
//
// Note that SQLite dialect does not support (this moment) the "append-only" mode describe above,
// since it's used only for testing.
func (m *Migrate) Create(ctx context.Context, tables ...*Table) error {
	m.setupTables(tables)
	var creator Creator = CreateFunc(m.create)
	for i := len(m.hooks) - 1; i >= 0; i-- {
		creator = m.hooks[i](creator)
	}
	return creator.Create(ctx, tables...)
}

func (m *Migrate) create(ctx context.Context, tables ...*Table) error {
	if err := m.init(ctx); err != nil {
		return err
	}
	tx, err := m.Tx(ctx)
	if err != nil {
		return err
	}
	if m.universalID {
		if err := m.types(ctx, tx); err != nil {
			return rollback(tx, err)
		}
	}
	if err := m.txCreate(ctx, tx, tables...); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func (m *Migrate) txCreate(ctx context.Context, tx dialect.Tx, tables ...*Table) error {
	for _, t := range tables {
		switch exist, err := m.tableExist(ctx, tx, t.Name); {
		case err != nil:
			return err
		case exist:
			curr, err := m.table(ctx, tx, t.Name)
			if err != nil {
				return err
			}
			if err := m.verify(ctx, tx, curr); err != nil {
				return err
			}
			if err := m.fixture(ctx, tx, curr, t); err != nil {
				return err
			}
			change, err := m.changeSet(curr, t)
			if err != nil {
				return fmt.Errorf("creating changeset for %q: %w", t.Name, err)
			}
			if err := m.apply(ctx, tx, t.Name, change); err != nil {
				return err
			}
		default: // !exist
			query, args := m.tBuilder(t).Query()
			if err := tx.Exec(ctx, query, args, nil); err != nil {
				return fmt.Errorf("create table %q: %w", t.Name, err)
			}
			// If global unique identifier is enabled, and it's not
			// a relation table, allocate a range for the table pk.
			if m.universalID && len(t.PrimaryKey) == 1 {
				if err := m.allocPKRange(ctx, tx, t); err != nil {
					return err
				}
			}
			// indexes.
			for _, idx := range t.Indexes {
				query, args := m.addIndex(idx, t.Name).Query()
				if err := tx.Exec(ctx, query, args, nil); err != nil {
					return fmt.Errorf("create index %q: %w", idx.Name, err)
				}
			}
		}
	}
	if !m.withForeignKeys {
		return nil
	}
	// Create foreign keys after tables were created/altered,
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
			return fmt.Errorf("create foreign keys for %q: %w", t.Name, err)
		}
	}
	return nil
}

// apply changes on the given table.
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
				return fmt.Errorf("drop index of table %q: %w", table, err)
			}
		}
	}
	var drop []*Column
	if m.dropColumns {
		drop = change.column.drop
	}
	queries := m.alterColumns(table, change.column.add, change.column.modify, drop)
	// If there's actual action to execute on ALTER TABLE.
	for i := range queries {
		query, args := queries[i].Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("alter table %q: %w", table, err)
		}
	}
	for _, idx := range change.index.add {
		query, args := m.addIndex(idx, table).Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("create index %q: %w", table, err)
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
	for i := range curr.PrimaryKey {
		if curr.PrimaryKey[i].Name != new.PrimaryKey[i].Name {
			return nil, fmt.Errorf("cannot change primary key for table: %q", curr.Name)
		}
	}
	// Add or modify columns.
	for _, c1 := range new.Columns {
		// Ignore primary keys.
		if c1.PrimaryKey() {
			continue
		}
		switch c2, ok := curr.column(c1.Name); {
		case !ok:
			change.column.add = append(change.column.add, c1)
		case !c2.Type.Valid():
			return nil, fmt.Errorf("invalid type %q for column %q", c2.typ, c2.Name)
		// Modify a non-unique column to unique.
		case c1.Unique && !c2.Unique:
			// Make sure the table does not have unique index for this column
			// before adding it to the changeset, because there are 2 ways to
			// configure uniqueness on sqlDialect.Field (using the Unique modifier or
			// adding rule on the Indexes option).
			if idx, ok := curr.index(c1.Name); !ok || !idx.Unique {
				change.index.add.append(&Index{
					Name:    c1.Name,
					Unique:  true,
					Columns: []*Column{c1},
					columns: []string{c1.Name},
				})
			}
		// Modify a unique column to non-unique.
		case !c1.Unique && c2.Unique:
			// If the uniqueness was defined on the Indexes option,
			// or was moved from the Unique modifier to the Indexes.
			if idx, ok := new.index(c1.Name); ok && idx.Unique {
				continue
			}
			idx, ok := curr.index(c2.Name)
			if !ok {
				return nil, fmt.Errorf("missing index to drop for unique column %q", c2.Name)
			}
			change.index.drop.append(idx)
		// Extending column types.
		case m.needsConversion(c2, c1):
			if !c2.ConvertibleTo(c1) {
				return nil, fmt.Errorf("changing column type for %q is invalid (%s != %s)", c1.Name, m.cType(c1), m.cType(c2))
			}
			fallthrough
		// Change nullability of a column.
		case c1.Nullable != c2.Nullable:
			change.column.modify = append(change.column.modify, c1)
		// Change default value.
		case c1.Default != nil && c2.Default == nil:
			change.column.modify = append(change.column.modify, c1)
		}
	}
	// Drop columns.
	for _, c1 := range curr.Columns {
		// If a column was dropped, multi-columns indexes that are associated with this column will
		// no longer behave the same. Therefore, these indexes should be dropped too. There's no need
		// to do it explicitly (here), because entc will remove them from the schema specification,
		// and they will be dropped in the block below.
		if _, ok := new.column(c1.Name); !ok {
			change.column.drop = append(change.column.drop, c1)
		}
	}
	// Add or modify indexes.
	for _, idx1 := range new.Indexes {
		switch idx2, ok := curr.index(idx1.Name); {
		case !ok:
			change.index.add.append(idx1)
		// Changing index cardinality require drop and create.
		case idx1.Unique != idx2.Unique:
			change.index.drop.append(idx2)
			change.index.add.append(idx1)
		default:
			im, ok := m.sqlDialect.(interface{ indexModified(old, new *Index) bool })
			// If the dialect supports comparing indexes.
			if ok && im.indexModified(idx2, idx1) {
				change.index.drop.append(idx2)
				change.index.add.append(idx1)
			}
		}
	}
	// Drop indexes.
	for _, idx := range curr.Indexes {
		if _, isFK := new.fk(idx.Name); !isFK && !new.hasIndex(idx.Name, idx.realname) {
			change.index.drop.append(idx)
		}
	}
	return change, nil
}

// fixture is a special migration code for renaming foreign-key columns (issue-#285).
func (m *Migrate) fixture(ctx context.Context, tx dialect.Tx, curr, new *Table) error {
	d, ok := m.sqlDialect.(fkRenamer)
	if !m.withFixture || !m.withForeignKeys || !ok {
		return nil
	}
	rename := make(map[string]*Index)
	for _, fk := range new.ForeignKeys {
		ok, err := m.fkExist(ctx, tx, fk.Symbol)
		if err != nil {
			return fmt.Errorf("checking foreign-key existence %q: %w", fk.Symbol, err)
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
			return fmt.Errorf("rename column %q: %w", column, err)
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
					return fmt.Errorf("rename index %q: %w", prev.Name, err)
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
					return fmt.Errorf("rename index %q: %w", idx.Name, err)
				}
				idx.Name = idx2.Name
				break Find
			}
		}
	}
	return nil
}

// verify that the auto-increment counter is correct for table with universal-id support.
func (m *Migrate) verify(ctx context.Context, tx dialect.Tx, t *Table) error {
	vr, ok := m.sqlDialect.(verifyRanger)
	if !ok || !m.universalID {
		return nil
	}
	id := indexOf(m.typeRanges, t.Name)
	if id == -1 {
		return nil
	}
	return vr.verifyRange(ctx, tx, t, int64(id<<32))
}

// types loads the type list from the type store. It will create the types table, if it does not exist yet.
func (m *Migrate) types(ctx context.Context, tx dialect.ExecQuerier) error {
	exists, err := m.tableExist(ctx, tx, TypeTable)
	if err != nil {
		return err
	}
	if !exists {
		t := NewTypesTable()
		query, args := m.tBuilder(t).Query()
		if err := tx.Exec(ctx, query, args, nil); err != nil {
			return fmt.Errorf("create types table: %w", err)
		}
		return nil
	}
	rows := &sql.Rows{}
	query, args := sql.Dialect(m.Dialect()).
		Select("type").From(sql.Table(TypeTable)).OrderBy(sql.Asc("id")).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return fmt.Errorf("query types table: %w", err)
	}
	defer rows.Close()
	return sql.ScanSlice(rows, &m.typeRanges)
}

func (m *Migrate) allocPKRange(ctx context.Context, conn dialect.ExecQuerier, t *Table) error {
	r, err := m.pkRange(ctx, conn, t)
	if err != nil {
		return err
	}
	return m.setRange(ctx, conn, t, r)
}

func (m *Migrate) pkRange(ctx context.Context, conn dialect.ExecQuerier, t *Table) (int64, error) {
	id := indexOf(m.typeRanges, t.Name)
	// If the table re-created, re-use its range from
	// the past. Otherwise, allocate a new id-range.
	if id == -1 {
		if len(m.typeRanges) > MaxTypes {
			return 0, fmt.Errorf("max number of types exceeded: %d", MaxTypes)
		}
		query, args := sql.Dialect(m.Dialect()).Insert(TypeTable).Columns("type").Values(t.Name).Query()
		if err := conn.Exec(ctx, query, args, nil); err != nil {
			return 0, fmt.Errorf("insert into ent_types: %w", err)
		}
		id = len(m.typeRanges)
		m.typeRanges = append(m.typeRanges, t.Name)
	}
	return int64(id << 32), nil
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
			m.sqlDialect.(fkRenamer).matchSchema(t2.C("table_schema")),
			m.sqlDialect.(fkRenamer).matchSchema(t1.C("table_schema")),
			sql.EQ(t2.C("constraint_name"), fk.Symbol),
		)).
		Query()
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return "", fmt.Errorf("reading foreign-key %q column: %w", fk.Symbol, err)
	}
	defer rows.Close()
	column, err := sql.ScanString(rows)
	if err != nil {
		return "", fmt.Errorf("scanning foreign-key %q column: %w", fk.Symbol, err)
	}
	return column, nil
}

// setup ensures the table is configured properly, like table columns
// are linked to their indexes, and PKs columns are defined.
func (m *Migrate) setupTables(tables []*Table) { m.atlas.setupTables(tables) }

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	err = fmt.Errorf("sql/schema: %w", err)
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

// exist checks if the given COUNT query returns a value >= 1.
func exist(ctx context.Context, conn dialect.ExecQuerier, query string, args ...any) (bool, error) {
	rows := &sql.Rows{}
	if err := conn.Query(ctx, query, args, rows); err != nil {
		return false, fmt.Errorf("reading schema information %w", err)
	}
	defer rows.Close()
	n, err := sql.ScanInt(rows)
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func indexOf(a []string, s string) int {
	for i := range a {
		if a[i] == s {
			return i
		}
	}
	return -1
}

type sqlDialect interface {
	atBuilder
	dialect.Driver
	init(context.Context) error
	table(context.Context, dialect.Tx, string) (*Table, error)
	tableExist(context.Context, dialect.ExecQuerier, string) (bool, error)
	fkExist(context.Context, dialect.Tx, string) (bool, error)
	setRange(context.Context, dialect.ExecQuerier, *Table, int64) error
	dropIndex(context.Context, dialect.Tx, *Index, string) error
	// table, column and index builder per dialect.
	cType(*Column) string
	tBuilder(*Table) *sql.TableBuilder
	addIndex(*Index, string) *sql.IndexBuilder
	alterColumns(table string, add, modify, drop []*Column) sql.Queries
	needsConversion(*Column, *Column) bool
}

type preparer interface {
	prepare(context.Context, dialect.Tx, *changes, string) error
}

// fkRenamer is used by the fixture migration (to solve #285),
// and it's implemented by the different dialects for renaming FKs.
type fkRenamer interface {
	matchSchema(...string) *sql.Predicate
	isImplicitIndex(*Index, *Column) bool
	renameIndex(*Table, *Index, *Index) sql.Querier
	renameColumn(*Table, *Column, *Column) sql.Querier
}

// verifyRanger wraps the method for verifying global-id range correctness.
type verifyRanger interface {
	verifyRange(context.Context, dialect.ExecQuerier, *Table, int64) error
}
