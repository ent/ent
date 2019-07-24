// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/plugin/ent/boring"
	"fbc/ent/entc/integration/plugin/ent/predicate"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// BoringUpdate is the builder for updating Boring entities.
type BoringUpdate struct {
	config
	predicates []predicate.Boring
}

// Where adds a new predicate for the builder.
func (bu *BoringUpdate) Where(ps ...predicate.Boring) *BoringUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BoringUpdate) Save(ctx context.Context) (int, error) {
	switch bu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bu.sqlSave(ctx)
	case dialect.Neptune:
		vertices, err := bu.gremlinSave(ctx)
		return len(vertices), err
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BoringUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BoringUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BoringUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BoringUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(boring.FieldID).From(sql.Table(boring.Table))
	for _, p := range bu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = bu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := bu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (bu *BoringUpdate) gremlinSave(ctx context.Context) ([]*Boring, error) {
	res := &gremlin.Response{}
	query, bindings := bu.gremlin().Query()
	if err := bu.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	var bs Borings
	bs.config(bu.config)
	if err := bs.FromResponse(res); err != nil {
		return nil, err
	}
	return bs, nil
}

func (bu *BoringUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(boring.Label)
	for _, p := range bu.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// BoringUpdateOne is the builder for updating a single Boring entity.
type BoringUpdateOne struct {
	config
	id string
}

// Save executes the query and returns the updated entity.
func (buo *BoringUpdateOne) Save(ctx context.Context) (*Boring, error) {
	switch buo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return buo.sqlSave(ctx)
	case dialect.Neptune:
		return buo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BoringUpdateOne) SaveX(ctx context.Context) *Boring {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BoringUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BoringUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BoringUpdateOne) sqlSave(ctx context.Context) (b *Boring, err error) {
	selector := sql.Select(boring.Columns...).From(sql.Table(boring.Table))
	boring.ID(buo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = buo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		b = &Boring{config: buo.config}
		if err := b.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Boring: %v", err)
		}
		id = b.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: Boring not found with id: %v", buo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Boring with the same id: %v", buo.id)
	}

	tx, err := buo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return b, nil
}

func (buo *BoringUpdateOne) gremlinSave(ctx context.Context) (*Boring, error) {
	res := &gremlin.Response{}
	query, bindings := buo.gremlin(buo.id).Query()
	if err := buo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	b := &Boring{config: buo.config}
	if err := b.FromResponse(res); err != nil {
		return nil, err
	}
	return b, nil
}

func (buo *BoringUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
