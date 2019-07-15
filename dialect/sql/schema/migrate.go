package schema

import (
	"context"
	"crypto/md5"
	"fmt"
	"math"
	"sort"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"
	"fbc/ent/field"
)

const (
	// TypeTable holds the table name holding the type information.
	TypeTable = "ent_types"
	// MaxTypes defines the max number of types can be created when
	// defining universal ids. The left 16-bits are reserved.
	MaxTypes = math.MaxUint16
)

// MigrateOption allows for managing schema configuration using functional options.
type MigrateOption func(m *Migrate)

// WithGlobalUniqueID sets the universal ids options to the migration.
func WithGlobalUniqueID(b bool) MigrateOption {
	return func(o *Migrate) {
		o.universalID = b
	}
}

// Migrate runs the migrations logic for the SQL dialects.
type Migrate struct {
	sqlDialect
	universalID bool     // global unique id flag.
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
		switch exist, err := m.tableExist(ctx, tx, t.Name); {
		case err != nil:
			return err
		case exist:
			curr, err := m.table(ctx, tx, t.Name)
			if err != nil {
				return err
			}
			change, err := m.changeSet(curr, t)
			if err != nil {
				return err
			}
			if len(change.add) != 0 || len(change.modify) != 0 {
				b := sql.AlterTable(curr.Name)
				for _, c := range change.add {
					b.AddColumn(m.cBuilder(c))
				}
				for _, c := range change.modify {
					b.ModifyColumn(m.cBuilder(c))
				}
				query, args := b.Query()
				if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
					return fmt.Errorf("alter table %q: %v", t.Name, err)
				}
			}
			if len(change.indexes) > 0 {
				panic("missing implementation")
			}
		default: // !exist
			query, args := m.tBuilder(t).Query()
			if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
				return fmt.Errorf("create table %q: %v", t.Name, err)
			}
			// if global unique identifier is enabled and it's not a relation table,
			// allocate a range for the table pk.
			if m.universalID && len(t.PrimaryKey) == 1 {
				if err := m.allocPKRange(ctx, tx, t); err != nil {
					return err
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
			fk.Symbol = symbol(fk.Symbol)
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
		b := sql.AlterTable(t.Name)
		for _, fk := range fks {
			b.AddForeignKey(fk.DSL())
		}
		query, args := b.Query()
		if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
			return fmt.Errorf("create foreign keys for %q: %v", t.Name, err)
		}
	}
	return nil
}

// changes to apply on existing table.
type changes struct {
	add     []*Column
	modify  []*Column
	indexes []*Index
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
	// columns.
	for _, c1 := range new.Columns {
		switch c2, ok := curr.column(c1.Name); {
		case !ok:
			change.add = append(change.add, c1)
		case c1.Unique != c2.Unique:
			return nil, fmt.Errorf("changing column cardinality for %q is invalid", c1.Name)
		case m.cType(c1) != m.cType(c2):
			if !c2.ConvertibleTo(c1) {
				return nil, fmt.Errorf("changing column type for %q is invalid (%s != %s)", c1.Name, m.cType(c1), m.cType(c2))
			}
			fallthrough
		case c1.Charset != "" && c1.Charset != c2.Charset || c1.Collation != "" && c1.Charset != c2.Collation:
			change.modify = append(change.modify, c1)
		}
	}
	// indexes.
	for _, idx1 := range new.Indexes {
		switch idx2, ok := curr.index(idx1.Name); {
		case !ok:
			change.indexes = append(change.indexes, idx1)
		case idx1.Unique != idx2.Unique:
			return nil, fmt.Errorf("changing index %q uniqness is invalid", idx1.Name)
		}
	}
	return change, nil
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
		if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
			return fmt.Errorf("create types table: %v", err)
		}
		return nil
	}
	rows := &sql.Rows{}
	query, args := sql.Select("type").From(sql.Table(TypeTable)).OrderBy(sql.Asc("id")).Query()
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
		query, args := sql.Insert(TypeTable).Columns("type").Values(t.Name).Query()
		if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
			return fmt.Errorf("insert into type: %v", err)
		}
		id = len(m.typeRanges)
		m.typeRanges = append(m.typeRanges, t.Name)
	}
	// set the id offset for table.
	return m.setRange(ctx, tx, t.Name, id<<32)
}

// symbol makes sure the symbol length is not longer than the maxlength in MySQL standard (64).
func symbol(name string) string {
	if len(name) <= 64 {
		return name
	}
	return fmt.Sprintf("%s_%x", name[:31], md5.Sum([]byte(name)))
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
	if !rows.Next() {
		return false, fmt.Errorf("no rows returned")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return false, fmt.Errorf("scanning count")
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
	// table and column builder per dialect.
	cType(*Column) string
	tBuilder(*Table) *sql.TableBuilder
	cBuilder(*Column) *sql.ColumnBuilder
}
