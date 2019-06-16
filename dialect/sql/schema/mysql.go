package schema

import (
	"context"
	"crypto/md5"
	"fmt"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"
)

// MySQL is a mysql migration driver.
type MySQL struct {
	dialect.Driver
}

// Create creates all tables resources in the database.
func (d *MySQL) Create(ctx context.Context, tables ...*Table) error {
	tx, err := d.Tx(ctx)
	if err != nil {
		return err
	}
	for _, t := range tables {
		exist, err := d.tableExist(ctx, tx, t.Name)
		if err != nil {
			return rollback(tx, err)
		}
		if exist {
			continue
		}
		query, args := t.DSL().Query()
		if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
			return rollback(tx, fmt.Errorf("sql/mysql: create table %q: %v", t.Name, err))
		}
	}
	// create foreign keys after table was created, because circular foreign-key constraints are possible.
	for _, t := range tables {
		if len(t.ForeignKeys) == 0 {
			continue
		}
		fks := make([]*ForeignKey, 0, len(t.ForeignKeys))
		for _, fk := range t.ForeignKeys {
			fk.Symbol = symbol(fk.Symbol)
			exist, err := d.fkExist(ctx, tx, fk.Symbol)
			if err != nil {
				return rollback(tx, err)
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
			return rollback(tx, fmt.Errorf("sql/mysql: create foreign keys for %q: %v", t.Name, err))
		}
	}
	return tx.Commit()
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
