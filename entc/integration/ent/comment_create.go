// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"fbc/ent/entc/integration/ent/comment"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	switch cc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return cc.sqlSave(ctx)
	case dialect.Neptune:
		return cc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	var (
		res sql.Result
		c   = &Comment{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(comment.Table).Default(cc.driver.Dialect())
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CommentCreate) gremlinSave(ctx context.Context) (*Comment, error) {
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Comment{config: cc.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CommentCreate) gremlin() *dsl.Traversal {
	v := g.AddV(comment.Label)
	return v.ValueMap(true)
}
