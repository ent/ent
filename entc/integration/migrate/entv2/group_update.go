// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"fbc/ent/entc/integration/migrate/entv2/group"
	"fbc/ent/entc/integration/migrate/entv2/predicate"

	"fbc/ent/dialect/sql"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	predicates []predicate.Group
}

// Where adds a new predicate for the builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	return gu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(group.FieldID).From(sql.Table(group.Table))
	for _, p := range gu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = gu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("entv2: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := gu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	id int
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return guo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	gr, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (gr *Group, err error) {
	selector := sql.Select(group.Columns...).From(sql.Table(group.Table))
	group.ID(guo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = guo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		gr = &Group{config: guo.config}
		if err := gr.FromRows(rows); err != nil {
			return nil, fmt.Errorf("entv2: failed scanning row into Group: %v", err)
		}
		id = gr.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("entv2: Group not found with id: %v", guo.id)
	case n > 1:
		return nil, fmt.Errorf("entv2: more than one Group with the same id: %v", guo.id)
	}

	tx, err := guo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
