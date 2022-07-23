// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/pet"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// PetQuery is the builder for querying Pet entities.
type PetQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Pet
	// eager-loading edges.
	withTeam  *UserQuery
	withOwner *UserQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the PetQuery builder.
func (pq *PetQuery) Where(ps ...predicate.Pet) *PetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PetQuery) Limit(limit int) *PetQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PetQuery) Offset(offset int) *PetQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PetQuery) Unique(unique bool) *PetQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PetQuery) Order(o ...OrderFunc) *PetQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryTeam chains the current query on the "team" edge.
func (pq *PetQuery) QueryTeam() *UserQuery {
	query := &UserQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.TeamLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (pq *PetQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := pq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.PetsLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first Pet entity from the query.
// Returns a *NotFoundError when no Pet was found.
func (pq *PetQuery) First(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PetQuery) FirstX(ctx context.Context) *Pet {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pet ID from the query.
// Returns a *NotFoundError when no Pet ID was found.
func (pq *PetQuery) FirstID(ctx context.Context) (id string, err error) {
	ids, err := pq.Limit(1).IDs(ctx)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PetQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pet entity is found.
// Returns a *NotFoundError when no Pet entities are found.
func (pq *PetQuery) Only(ctx context.Context) (*Pet, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pet.Label}
	default:
		return nil, &NotSingularError{pet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PetQuery) OnlyX(ctx context.Context) *Pet {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pet ID in the query.
// Returns a *NotSingularError when more than one Pet ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PetQuery) OnlyID(ctx context.Context) (id string, err error) {
	ids, err := pq.Limit(2).IDs(ctx)
	if err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pet.Label}
	default:
		err = &NotSingularError{pet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PetQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pets.
func (pq *PetQuery) All(ctx context.Context) ([]*Pet, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PetQuery) AllX(ctx context.Context) []*Pet {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pet IDs.
func (pq *PetQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := pq.Select(pet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PetQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PetQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PetQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PetQuery) Clone() *PetQuery {
	if pq == nil {
		return nil
	}
	return &PetQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		predicates: append([]predicate.Pet{}, pq.predicates...),
		withTeam:   pq.withTeam.Clone(),
		withOwner:  pq.withOwner.Clone(),
		// clone intermediate query.
		gremlin: pq.gremlin.Clone(),
		path:    pq.path,
		unique:  pq.unique,
	}
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithTeam(opts ...func(*UserQuery)) *PetQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withTeam = query
	return pq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PetQuery) WithOwner(opts ...func(*UserQuery)) *PetQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Age float64 `json:"age,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pet.Query().
//		GroupBy(pet.FieldAge).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PetQuery) GroupBy(field string, fields ...string) *PetGroupBy {
	grbuild := &PetGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.gremlinQuery(ctx), nil
	}
	grbuild.label = pet.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Age float64 `json:"age,omitempty"`
//	}
//
//	client.Pet.Query().
//		Select(pet.FieldAge).
//		Scan(ctx, &v)
//
func (pq *PetQuery) Select(fields ...string) *PetSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PetSelect{PetQuery: pq}
	selbuild.label = pet.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

func (pq *PetQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.gremlin = prev
	}
	return nil
}

func (pq *PetQuery) gremlinAll(ctx context.Context) ([]*Pet, error) {
	res := &gremlin.Response{}
	traversal := pq.gremlinQuery(ctx)
	if len(pq.fields) > 0 {
		fields := make([]interface{}, len(pq.fields))
		for i, f := range pq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var pes Pets
	if err := pes.FromResponse(res); err != nil {
		return nil, err
	}
	pes.config(pq.config)
	return pes, nil
}

func (pq *PetQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery(ctx).Count().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pq *PetQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery(ctx).HasNext().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (pq *PetQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(pet.Label)
	if pq.gremlin != nil {
		v = pq.gremlin.Clone()
	}
	for _, p := range pq.predicates {
		p(v)
	}
	if len(pq.order) > 0 {
		v.Order()
		for _, p := range pq.order {
			p(v)
		}
	}
	switch limit, offset := pq.limit, pq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := pq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// PetGroupBy is the group-by builder for Pet entities.
type PetGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PetGroupBy) Aggregate(fns ...AggregateFunc) *PetGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PetGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.gremlin = query
	return pgb.gremlinScan(ctx, v)
}

func (pgb *PetGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := pgb.gremlinQuery().Query()
	if err := pgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(pgb.fields)+len(pgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (pgb *PetGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range pgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range pgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return pgb.gremlin.Group().
		By(__.Values(pgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// PetSelect is the builder for selecting fields of Pet entities.
type PetSelect struct {
	*PetQuery
	selector
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PetSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.gremlin = ps.PetQuery.gremlinQuery(ctx)
	return ps.gremlinScan(ctx, v)
}

func (ps *PetSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(ps.fields) == 1 {
		if ps.fields[0] != pet.FieldID {
			traversal = ps.gremlin.Values(ps.fields...)
		} else {
			traversal = ps.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(ps.fields))
		for i, f := range ps.fields {
			fields[i] = f
		}
		traversal = ps.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ps.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ps.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
