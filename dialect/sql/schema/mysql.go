package schema

import (
	"context"
	"crypto/md5"
	"fmt"
	"sort"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"
)

// MySQL is a mysql migration driver.
type MySQL struct {
	dialect.Driver
}

// Create creates all schema resources in the database. It works in an "append-only"
// mode, which means, it won't delete or change any existing resource in the database.
func (d *MySQL) Create(ctx context.Context, tables ...*Table) error {
	tx, err := d.Tx(ctx)
	if err != nil {
		return err
	}
	if err := d.create(ctx, tx, tables...); err != nil {
		return rollback(tx, fmt.Errorf("dialect/mysql: %v", err))
	}
	return tx.Commit()
}

func (d *MySQL) create(ctx context.Context, tx dialect.Tx, tables ...*Table) error {
	version, err := d.version(ctx, tx)
	if err != nil {
		return err
	}
	for _, t := range tables {
		switch exist, err := d.tableExist(ctx, tx, t.Name); {
		case err != nil:
			return err
		case exist:
			curr, err := d.table(ctx, tx, t.Name)
			if err != nil {
				return err
			}
			change, err := changeSet(curr, t)
			if err != nil {
				return err
			}
			if len(change.add) != 0 || len(change.modify) != 0 {
				b := sql.AlterTable(curr.Name)
				for _, c := range change.add {
					b.AddColumn(c.MySQL(version))
				}
				for _, c := range change.modify {
					b.ModifyColumn(c.MySQL(version))
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
			query, args := t.MySQL(version).Query()
			if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
				return fmt.Errorf("create table %q: %v", t.Name, err)
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
			exist, err := d.fkExist(ctx, tx, fk.Symbol)
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

func (d *MySQL) version(ctx context.Context, tx dialect.Tx) (string, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, "SHOW VARIABLES LIKE 'version'", []interface{}{}, rows); err != nil {
		return "", fmt.Errorf("dialect/mysql: querying mysql version %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return "", fmt.Errorf("dialect/mysql: version variable was not found")
	}
	version := make([]string, 2)
	if err := rows.Scan(&version[0], &version[1]); err != nil {
		return "", fmt.Errorf("dialect/mysql: scanning mysql version: %v", err)
	}
	return version[1], nil
}

func (d *MySQL) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLES").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("TABLE_NAME", name)).Query()
	return d.exist(ctx, tx, query, args...)
}

func (d *MySQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select(sql.Count("*")).From(sql.Table("INFORMATION_SCHEMA.TABLE_CONSTRAINTS").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("CONSTRAINT_TYPE", "FOREIGN KEY").And().EQ("CONSTRAINT_NAME", name)).Query()
	return d.exist(ctx, tx, query, args...)
}

func (d *MySQL) exist(ctx context.Context, tx dialect.Tx, query string, args ...interface{}) (bool, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return false, fmt.Errorf("dialect/mysql: reading schema information %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return false, fmt.Errorf("dialect/mysql: no rows returned")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return false, fmt.Errorf("dialect/mysql: scanning count")
	}
	return n > 0, nil
}

// table loads the current table description from the database.
func (d *MySQL) table(ctx context.Context, tx dialect.Tx, name string) (*Table, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name").
		From(sql.Table("INFORMATION_SCHEMA.COLUMNS").Unquote()).
		Where(sql.EQ("TABLE_SCHEMA", sql.Raw("(SELECT DATABASE())")).And().EQ("TABLE_NAME", name)).Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("dialect/mysql: reading table description %v", err)
	}
	defer rows.Close()
	t := &Table{Name: name}
	for rows.Next() {
		c := &Column{}
		if err := c.ScanMySQL(rows); err != nil {
			return nil, fmt.Errorf("dialect/mysql: %v", err)
		}
		if c.PrimaryKey() {
			t.PrimaryKey = append(t.PrimaryKey, c)
		}
		t.Columns = append(t.Columns, c)
	}
	return t, nil
}

// changes to apply on existing table.
type changes struct {
	add     []*Column
	modify  []*Column
	indexes []*Index
}

// changeSet returns a changes object to be applied on existing table.
// It fails if one of the changes is invalid.
func changeSet(curr, new *Table) (*changes, error) {
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
		case c1.Type != c2.Type:
			return nil, fmt.Errorf("changing column type for %q is invalid (%s != %s)", c1.Name, c1.Type, c2.Type)
		case c1.Unique != c2.Unique:
			return nil, fmt.Errorf("changing column cardinality for %q is invalid", c1.Name)
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

// symbol makes sure the symbol length is not longer than the maxlength in MySQL standard (64).
func symbol(name string) string {
	if len(name) <= 64 {
		return name
	}
	return fmt.Sprintf("%s_%x", name[:31], md5.Sum([]byte(name)))
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}
