// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/ent/comment"
	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	predicates []predicate.Comment
}

// Where adds a new predicate for the builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	switch cu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return cu.sqlSave(ctx)
	case dialect.Neptune:
		vertices, err := cu.gremlinSave(ctx)
		return len(vertices), err
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(comment.FieldID).From(sql.Table(comment.Table))
	for _, p := range cu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := cu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (cu *CommentUpdate) gremlinSave(ctx context.Context) ([]*Comment, error) {
	res := &gremlin.Response{}
	query, bindings := cu.gremlin().Query()
	if err := cu.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	var cs Comments
	cs.config(cu.config)
	if err := cs.FromResponse(res); err != nil {
		return nil, err
	}
	return cs, nil
}

func (cu *CommentUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(comment.Label)
	for _, p := range cu.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	id string
}

// Save executes the query and returns the updated entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	switch cuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return cuo.sqlSave(ctx)
	case dialect.Neptune:
		return cuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (c *Comment, err error) {
	selector := sql.Select(comment.Columns...).From(sql.Table(comment.Table))
	comment.ID(cuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		c = &Comment{config: cuo.config}
		if err := c.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Comment: %v", err)
		}
		id = c.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: Comment not found with id: %v", cuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Comment with the same id: %v", cuo.id)
	}

	tx, err := cuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CommentUpdateOne) gremlinSave(ctx context.Context) (*Comment, error) {
	res := &gremlin.Response{}
	query, bindings := cuo.gremlin(cuo.id).Query()
	if err := cuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Comment{config: cuo.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CommentUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
