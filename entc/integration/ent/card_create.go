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
	create_time *time.Time
	update_time *time.Time
	number      *string
	name        *string
	owner       map[string]struct{}
}

// SetCreateTime sets the create_time field.
func (cc *CardCreate) SetCreateTime(t time.Time) *CardCreate {
	cc.create_time = &t
	return cc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (cc *CardCreate) SetNillableCreateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the update_time field.
func (cc *CardCreate) SetUpdateTime(t time.Time) *CardCreate {
	cc.update_time = &t
	return cc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetNumber sets the number field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.number = &s
	return cc
}

// SetName sets the name field.
func (cc *CardCreate) SetName(s string) *CardCreate {
	cc.name = &s
	return cc
}

// SetNillableName sets the name field if the given value is not nil.
func (cc *CardCreate) SetNillableName(s *string) *CardCreate {
	if s != nil {
		cc.SetName(*s)
	}
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
	if cc.create_time == nil {
		v := card.DefaultCreateTime()
		cc.create_time = &v
	}
	if cc.update_time == nil {
		v := card.DefaultUpdateTime()
		cc.update_time = &v
	}
	if cc.number == nil {
		return nil, errors.New("ent: missing required field \"number\"")
	}
	if err := card.NumberValidator(*cc.number); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"number\": %v", err)
	}
	if cc.name != nil {
		if err := card.NameValidator(*cc.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(cc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	switch cc.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
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
		res     sql.Result
		builder = sql.Dialect(cc.driver.Dialect())
		c       = &Card{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(card.Table).Default()
	if value := cc.create_time; value != nil {
		insert.Set(card.FieldCreateTime, *value)
		c.CreateTime = *value
	}
	if value := cc.update_time; value != nil {
		insert.Set(card.FieldUpdateTime, *value)
		c.UpdateTime = *value
	}
	if value := cc.number; value != nil {
		insert.Set(card.FieldNumber, *value)
		c.Number = *value
	}
	if value := cc.name; value != nil {
		insert.Set(card.FieldName, *value)
		c.Name = *value
	}
	id, err := insertLastID(ctx, tx, insert.Returning(card.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = strconv.FormatInt(id, 10)
	if len(cc.owner) > 0 {
		eid, err := strconv.Atoi(keys(cc.owner)[0])
		if err != nil {
			return nil, err
		}
		query, args := builder.Update(card.OwnerTable).
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
	if cc.create_time != nil {
		v.Property(dsl.Single, card.FieldCreateTime, *cc.create_time)
	}
	if cc.update_time != nil {
		v.Property(dsl.Single, card.FieldUpdateTime, *cc.update_time)
	}
	if cc.number != nil {
		v.Property(dsl.Single, card.FieldNumber, *cc.number)
	}
	if cc.name != nil {
		v.Property(dsl.Single, card.FieldName, *cc.name)
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
