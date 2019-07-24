// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/migrate/entv2/group"
	"fbc/ent/entc/integration/migrate/entv2/predicate"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
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
	switch gu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gu.sqlSave(ctx)
	case dialect.Neptune:
		vertices, err := gu.gremlinSave(ctx)
		return len(vertices), err
	default:
		return 0, errors.New("entv2: unsupported dialect")
	}
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

func (gu *GroupUpdate) gremlinSave(ctx context.Context) ([]*Group, error) {
	res := &gremlin.Response{}
	query, bindings := gu.gremlin().Query()
	if err := gu.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	var grs Groups
	grs.config(gu.config)
	if err := grs.FromResponse(res); err != nil {
		return nil, err
	}
	return grs, nil
}

func (gu *GroupUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(group.Label)
	for _, p := range gu.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	id string
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	switch guo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return guo.sqlSave(ctx)
	case dialect.Neptune:
		return guo.gremlinSave(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
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
		id = gr.id()
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

func (guo *GroupUpdateOne) gremlinSave(ctx context.Context) (*Group, error) {
	res := &gremlin.Response{}
	query, bindings := guo.gremlin(guo.id).Query()
	if err := guo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gr := &Group{config: guo.config}
	if err := gr.FromResponse(res); err != nil {
		return nil, err
	}
	return gr, nil
}

func (guo *GroupUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
