// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	created_at *time.Time
	updated_at *time.Time
	number     *string
	owner      map[string]struct{}
}

// SetCreatedAt sets the created_at field.
func (cc *CardCreate) SetCreatedAt(t time.Time) *CardCreate {
	cc.created_at = &t
	return cc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cc *CardCreate) SetNillableCreatedAt(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the updated_at field.
func (cc *CardCreate) SetUpdatedAt(t time.Time) *CardCreate {
	cc.updated_at = &t
	return cc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdatedAt(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetNumber sets the number field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.number = &s
	return cc
}

// SetOwnerID sets the owner edge to User by id.
func (cc *CardCreate) SetOwnerID(id string) *CardCreate {
	if cc.owner == nil {
		cc.owner = make(map[string]struct{})
	}
	cc.owner[id] = struct{}{}
	return cc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(id *string) *CardCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	if cc.created_at == nil {
		v := card.DefaultCreatedAt()
		cc.created_at = &v
	}
	if cc.updated_at == nil {
		v := card.DefaultUpdatedAt()
		cc.updated_at = &v
	}
	if cc.number == nil {
		return nil, errors.New("ent: missing required field \"number\"")
	}
	if err := card.NumberValidator(*cc.number); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"number\": %v", err)
	}
	if len(cc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	switch cc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return cc.sqlSave(ctx)
	case dialect.Gremlin:
		return cc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	var (
		res sql.Result
		c   = &Card{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Dialect(cc.driver.Dialect()).
		Insert(card.Table).
		Default()
	if value := cc.created_at; value != nil {
		builder.Set(card.FieldCreatedAt, *value)
		c.CreatedAt = *value
	}
	if value := cc.updated_at; value != nil {
		builder.Set(card.FieldUpdatedAt, *value)
		c.UpdatedAt = *value
	}
	if value := cc.number; value != nil {
		builder.Set(card.FieldNumber, *value)
		c.Number = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = strconv.FormatInt(id, 10)
	if len(cc.owner) > 0 {
		eid, err := strconv.Atoi(keys(cc.owner)[0])
		if err != nil {
			return nil, err
		}
		query, args := sql.Update(card.OwnerTable).
			Set(card.OwnerColumn, eid).
			Where(sql.EQ(card.FieldID, id).And().IsNull(card.OwnerColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(cc.owner) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"owner\" %v already connected to a different \"Card\"", keys(cc.owner))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CardCreate) gremlinSave(ctx context.Context) (*Card, error) {
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Card{config: cc.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CardCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(card.Label)
	if cc.created_at != nil {
		v.Property(dsl.Single, card.FieldCreatedAt, *cc.created_at)
	}
	if cc.updated_at != nil {
		v.Property(dsl.Single, card.FieldUpdatedAt, *cc.updated_at)
	}
	if cc.number != nil {
		v.Property(dsl.Single, card.FieldNumber, *cc.number)
	}
	for id := range cc.owner {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}
