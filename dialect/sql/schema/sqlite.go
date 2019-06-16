package schema

import (
	"context"
	"fmt"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"
)

// SQLite is an SQLite migration driver.
type SQLite struct {
	dialect.Driver
}

// Create creates all tables resources in the database.
func (d *SQLite) Create(ctx context.Context, tables ...*Table) error {
	tx, err := d.Tx(ctx)
	if err != nil {
		return err
	}
	on, err := d.fkEnabled(ctx, tx)
	if err != nil {
		return fmt.Errorf("sql/sqlite: check foreign_keys pragma: %v", err)
	}
	if !on {
		// foreign_keys pragma is off, either enable it by execute "PRAGMA foreign_keys=ON"
		// or add the following parameter in the connection string "_fk=1".
		return fmt.Errorf("sql/sqlite: foreign_keys pragma is off: missing %q is the connection string", "_fk=1")
	}
	for _, t := range tables {
		exist, err := d.tableExist(ctx, tx, t.Name)
		if err != nil {
			return rollback(tx, err)
		}
		if exist {
			continue
		}
		query, args := t.SQLite().Query()
		if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
			err = fmt.Errorf("sql/sqlite: create table %q: %v", t.Name, err)
			return rollback(tx, err)
		}
	}
	return tx.Commit()
}

func (d *SQLite) tableExist(ctx context.Context, tx dialect.Tx, name string) (bool, error) {
	query, args := sql.Select().Count().
		From(sql.Table("sqlite_master")).
		Where(sql.EQ("type", "table").And().EQ("name", name)).
		Query()
	return d.exist(ctx, tx, query, args...)
}

func (d *SQLite) fkEnabled(ctx context.Context, tx dialect.Tx) (bool, error) {
	return d.exist(ctx, tx, "PRAGMA foreign_keys")
}

func (d *SQLite) exist(ctx context.Context, tx dialect.Tx, query string, args ...interface{}) (bool, error) {
	rows := &sql.Rows{}
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return false, fmt.Errorf("dialect/sqlite: reading schema information %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return false, fmt.Errorf("dialect/sqlite: no rows returned")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return false, fmt.Errorf("dialect/sqlite: scanning count")
	}
	return n > 0, nil
}
