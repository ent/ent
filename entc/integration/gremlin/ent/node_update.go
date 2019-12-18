// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/node"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
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
	return nu.gremlinSave(ctx)
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
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
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
	return nuo.gremlinSave(ctx)
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
