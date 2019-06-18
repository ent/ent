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
	for _, t := range tables {
		switch exist, err := d.tableExist(ctx, tx, t.Name); {
		case err != nil:
			return err
		case exist:
			curr, err := d.table(ctx, tx, t.Name)
			if err != nil {
				return err
			}
			changes, err := changeSet(curr, t)
			if err != nil {
				return err
			}
			if len(changes.Columns) > 0 {
				b := sql.AlterTable(curr.Name)
				for _, c := range changes.Columns {
					b.AddColumn(c.DSL())
				}
				query, args := b.Query()
				if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
					return fmt.Errorf("alter table %q: %v", t.Name, err)
				}
			}
			if len(changes.Indexes) > 0 {
				panic("missing implementation")
			}
		default: // !exist
			query, args := t.DSL().Query()
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

func (d *MySQL) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	return d.exist(
		ctx,
		tx,
		"SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = (SELECT DATABASE()) AND TABLE_NAME = ?",
		name,
	)
}

func (d *MySQL) fkExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	return d.exist(
		ctx,
		tx,
		`SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS WHERE TABLE_SCHEMA=(SELECT DATABASE()) AND CONSTRAINT_TYPE="FOREIGN KEY" AND CONSTRAINT_NAME = ?`,
		name,
	)
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
	if err := tx.Query(ctx, "DESCRIBE "+name, []interface{}{}, rows); err != nil {
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

// changeSet returns a dummy table represents the change set that need
// to be applied on the table. it fails if one of the changes is invalid.
func changeSet(curr, new *Table) (*Table, error) {
	changes := &Table{}
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
			changes.Columns = append(changes.Columns, c1)
		case c1.Type != c2.Type:
			return nil, fmt.Errorf("changing column type for %q is invalid", c1.Name)
		case c1.Unique != c2.Unique:
			return nil, fmt.Errorf("changing column cardinality for %q is invalid", c1.Name)
		}
	}
	// indexes.
	for _, idx1 := range new.Indexes {
		switch idx2, ok := curr.index(idx1.Name); {
		case !ok:
			changes.Indexes = append(changes.Indexes, idx1)
		case idx1.Unique != idx2.Unique:
			return nil, fmt.Errorf("changing index %q uniqness is invalid", idx1.Name)
		}
	}
	return changes, nil
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
