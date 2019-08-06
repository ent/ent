// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"

	"fbc/ent/entc/integration/migrate/entv2/group"

	"fbc/ent/dialect/sql"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	return gc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	var (
		res sql.Result
		gr  = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(group.Table).Default(gc.driver.Dialect())
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
