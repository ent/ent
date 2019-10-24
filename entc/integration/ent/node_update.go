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

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/node"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	value       *int
	addvalue    *int
	clearvalue  bool
	prev        map[string]struct{}
	next        map[string]struct{}
	clearedPrev bool
	clearedNext bool
	predicates  []predicate.Node
}

// Where adds a new predicate for the builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.predicates = append(nu.predicates, ps...)
	return nu
}

// SetValue sets the value field.
func (nu *NodeUpdate) SetValue(i int) *NodeUpdate {
	nu.value = &i
	nu.addvalue = nil
	return nu
}

// SetNillableValue sets the value field if the given value is not nil.
func (nu *NodeUpdate) SetNillableValue(i *int) *NodeUpdate {
	if i != nil {
		nu.SetValue(*i)
	}
	return nu
}

// AddValue adds i to value.
func (nu *NodeUpdate) AddValue(i int) *NodeUpdate {
	if nu.addvalue == nil {
		nu.addvalue = &i
	} else {
		*nu.addvalue += i
	}
	return nu
}

// ClearValue clears the value of value.
func (nu *NodeUpdate) ClearValue() *NodeUpdate {
	nu.value = nil
	nu.clearvalue = true
	return nu
}

// SetPrevID sets the prev edge to Node by id.
func (nu *NodeUpdate) SetPrevID(id string) *NodeUpdate {
	if nu.prev == nil {
		nu.prev = make(map[string]struct{})
	}
	nu.prev[id] = struct{}{}
	return nu
}

// SetNillablePrevID sets the prev edge to Node by id if the given value is not nil.
func (nu *NodeUpdate) SetNillablePrevID(id *string) *NodeUpdate {
	if id != nil {
		nu = nu.SetPrevID(*id)
	}
	return nu
}

// SetPrev sets the prev edge to Node.
func (nu *NodeUpdate) SetPrev(n *Node) *NodeUpdate {
	return nu.SetPrevID(n.ID)
}

// SetNextID sets the next edge to Node by id.
func (nu *NodeUpdate) SetNextID(id string) *NodeUpdate {
	if nu.next == nil {
		nu.next = make(map[string]struct{})
	}
	nu.next[id] = struct{}{}
	return nu
}

// SetNillableNextID sets the next edge to Node by id if the given value is not nil.
func (nu *NodeUpdate) SetNillableNextID(id *string) *NodeUpdate {
	if id != nil {
		nu = nu.SetNextID(*id)
	}
	return nu
}

// SetNext sets the next edge to Node.
func (nu *NodeUpdate) SetNext(n *Node) *NodeUpdate {
	return nu.SetNextID(n.ID)
}

// ClearPrev clears the prev edge to Node.
func (nu *NodeUpdate) ClearPrev() *NodeUpdate {
	nu.clearedPrev = true
	return nu
}

// ClearNext clears the next edge to Node.
func (nu *NodeUpdate) ClearNext() *NodeUpdate {
	nu.clearedNext = true
	return nu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	if len(nu.prev) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"prev\"")
	}
	if len(nu.next) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"next\"")
	}
	switch nu.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		return nu.sqlSave(ctx)
	case dialect.Gremlin:
		return nu.gremlinSave(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(node.FieldID).From(sql.Table(node.Table))
	for _, p := range nu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = nu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := nu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(node.Table).Where(sql.InInts(node.FieldID, ids...))
	)
	if value := nu.value; value != nil {
		builder.Set(node.FieldValue, *value)
	}
	if value := nu.addvalue; value != nil {
		builder.Add(node.FieldValue, *value)
	}
	if nu.clearvalue {
		builder.SetNull(node.FieldValue)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if nu.clearedPrev {
		query, args := sql.Update(node.PrevTable).
			SetNull(node.PrevColumn).
			Where(sql.InInts(node.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(nu.prev) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(nu.prev)[0])
			if serr != nil {
				return 0, rollback(tx, err)
			}
			query, args := sql.Update(node.PrevTable).
				Set(node.PrevColumn, eid).
				Where(sql.EQ(node.FieldID, id).And().IsNull(node.PrevColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(nu.prev) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"prev\" %v already connected to a different \"Node\"", keys(nu.prev))})
			}
		}
	}
	if nu.clearedNext {
		query, args := sql.Update(node.NextTable).
			SetNull(node.NextColumn).
			Where(sql.InInts(node.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(nu.next) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(nu.next)[0])
			if serr != nil {
				return 0, rollback(tx, err)
			}
			query, args := sql.Update(node.NextTable).
				Set(node.NextColumn, id).
				Where(sql.EQ(node.FieldID, eid).And().IsNull(node.NextColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(nu.next) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"next\" %v already connected to a different \"Node\"", keys(nu.next))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (nu *NodeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := nu.gremlin().Query()
	if err := nu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (nu *NodeUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(node.Label)
	for _, p := range nu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := nu.value; value != nil {
		v.Property(dsl.Single, node.FieldValue, *value)
	}
	if value := nu.addvalue; value != nil {
		v.Property(dsl.Single, node.FieldValue, __.Union(__.Values(node.FieldValue), __.Constant(*value)).Sum())
	}
	var properties []interface{}
	if nu.clearvalue {
		properties = append(properties, node.FieldValue)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if nu.clearedPrev {
		tr := rv.Clone().InE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range nu.prev {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if nu.clearedNext {
		tr := rv.Clone().OutE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range nu.next {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ErrConstraintFailed{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	id          string
	value       *int
	addvalue    *int
	clearvalue  bool
	prev        map[string]struct{}
	next        map[string]struct{}
	clearedPrev bool
	clearedNext bool
}

// SetValue sets the value field.
func (nuo *NodeUpdateOne) SetValue(i int) *NodeUpdateOne {
	nuo.value = &i
	nuo.addvalue = nil
	return nuo
}

// SetNillableValue sets the value field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableValue(i *int) *NodeUpdateOne {
	if i != nil {
		nuo.SetValue(*i)
	}
	return nuo
}

// AddValue adds i to value.
func (nuo *NodeUpdateOne) AddValue(i int) *NodeUpdateOne {
	if nuo.addvalue == nil {
		nuo.addvalue = &i
	} else {
		*nuo.addvalue += i
	}
	return nuo
}

// ClearValue clears the value of value.
func (nuo *NodeUpdateOne) ClearValue() *NodeUpdateOne {
	nuo.value = nil
	nuo.clearvalue = true
	return nuo
}

// SetPrevID sets the prev edge to Node by id.
func (nuo *NodeUpdateOne) SetPrevID(id string) *NodeUpdateOne {
	if nuo.prev == nil {
		nuo.prev = make(map[string]struct{})
	}
	nuo.prev[id] = struct{}{}
	return nuo
}

// SetNillablePrevID sets the prev edge to Node by id if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillablePrevID(id *string) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetPrevID(*id)
	}
	return nuo
}

// SetPrev sets the prev edge to Node.
func (nuo *NodeUpdateOne) SetPrev(n *Node) *NodeUpdateOne {
	return nuo.SetPrevID(n.ID)
}

// SetNextID sets the next edge to Node by id.
func (nuo *NodeUpdateOne) SetNextID(id string) *NodeUpdateOne {
	if nuo.next == nil {
		nuo.next = make(map[string]struct{})
	}
	nuo.next[id] = struct{}{}
	return nuo
}

// SetNillableNextID sets the next edge to Node by id if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableNextID(id *string) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetNextID(*id)
	}
	return nuo
}

// SetNext sets the next edge to Node.
func (nuo *NodeUpdateOne) SetNext(n *Node) *NodeUpdateOne {
	return nuo.SetNextID(n.ID)
}

// ClearPrev clears the prev edge to Node.
func (nuo *NodeUpdateOne) ClearPrev() *NodeUpdateOne {
	nuo.clearedPrev = true
	return nuo
}

// ClearNext clears the next edge to Node.
func (nuo *NodeUpdateOne) ClearNext() *NodeUpdateOne {
	nuo.clearedNext = true
	return nuo
}

// Save executes the query and returns the updated entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	if len(nuo.prev) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"prev\"")
	}
	if len(nuo.next) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"next\"")
	}
	switch nuo.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		return nuo.sqlSave(ctx)
	case dialect.Gremlin:
		return nuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	n, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (n *Node, err error) {
	selector := sql.Select(node.Columns...).From(sql.Table(node.Table))
	node.ID(nuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = nuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		n = &Node{config: nuo.config}
		if err := n.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Node: %v", err)
		}
		id = n.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Node with id: %v", nuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Node with the same id: %v", nuo.id)
	}

	tx, err := nuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(node.Table).Where(sql.InInts(node.FieldID, ids...))
	)
	if value := nuo.value; value != nil {
		builder.Set(node.FieldValue, *value)
		n.Value = *value
	}
	if value := nuo.addvalue; value != nil {
		builder.Add(node.FieldValue, *value)
		n.Value += *value
	}
	if nuo.clearvalue {
		var value int
		n.Value = value
		builder.SetNull(node.FieldValue)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if nuo.clearedPrev {
		query, args := sql.Update(node.PrevTable).
			SetNull(node.PrevColumn).
			Where(sql.InInts(node.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(nuo.prev) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(nuo.prev)[0])
			if serr != nil {
				return nil, rollback(tx, err)
			}
			query, args := sql.Update(node.PrevTable).
				Set(node.PrevColumn, eid).
				Where(sql.EQ(node.FieldID, id).And().IsNull(node.PrevColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(nuo.prev) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"prev\" %v already connected to a different \"Node\"", keys(nuo.prev))})
			}
		}
	}
	if nuo.clearedNext {
		query, args := sql.Update(node.NextTable).
			SetNull(node.NextColumn).
			Where(sql.InInts(node.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(nuo.next) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(nuo.next)[0])
			if serr != nil {
				return nil, rollback(tx, err)
			}
			query, args := sql.Update(node.NextTable).
				Set(node.NextColumn, id).
				Where(sql.EQ(node.FieldID, eid).And().IsNull(node.NextColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(nuo.next) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"next\" %v already connected to a different \"Node\"", keys(nuo.next))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return n, nil
}

func (nuo *NodeUpdateOne) gremlinSave(ctx context.Context) (*Node, error) {
	res := &gremlin.Response{}
	query, bindings := nuo.gremlin(nuo.id).Query()
	if err := nuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	n := &Node{config: nuo.config}
	if err := n.FromResponse(res); err != nil {
		return nil, err
	}
	return n, nil
}

func (nuo *NodeUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := nuo.value; value != nil {
		v.Property(dsl.Single, node.FieldValue, *value)
	}
	if value := nuo.addvalue; value != nil {
		v.Property(dsl.Single, node.FieldValue, __.Union(__.Values(node.FieldValue), __.Constant(*value)).Sum())
	}
	var properties []interface{}
	if nuo.clearvalue {
		properties = append(properties, node.FieldValue)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if nuo.clearedPrev {
		tr := rv.Clone().InE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range nuo.prev {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if nuo.clearedNext {
		tr := rv.Clone().OutE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range nuo.next {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
