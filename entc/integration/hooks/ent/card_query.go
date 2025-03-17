// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/hooks/ent/card"
	"entgo.io/ent/entc/integration/hooks/ent/predicate"
	"entgo.io/ent/entc/integration/hooks/ent/user"
	"entgo.io/ent/schema/field"
)

// CardQuery is the builder for querying Card entities.
type CardQuery struct {
	config
	ctx        *QueryContext
	order      []card.OrderOption
	inters     []Interceptor
	predicates []predicate.Card
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CardQuery builder.
func (q *CardQuery) Where(ps ...predicate.Card) *CardQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *CardQuery) Limit(limit int) *CardQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *CardQuery) Offset(offset int) *CardQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *CardQuery) Unique(unique bool) *CardQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *CardQuery) Order(o ...card.OrderOption) *CardQuery {
	q.order = append(q.order, o...)
	return q
}

// QueryOwner chains the current query on the "owner" edge.
func (q *CardQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, card.OwnerTable, card.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Card entity from the query.
// Returns a *NotFoundError when no Card was found.
func (q *CardQuery) First(ctx context.Context) (*Card, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{card.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *CardQuery) FirstX(ctx context.Context) *Card {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Card ID from the query.
// Returns a *NotFoundError when no Card ID was found.
func (q *CardQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{card.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *CardQuery) FirstIDX(ctx context.Context) int {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Card entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Card entity is found.
// Returns a *NotFoundError when no Card entities are found.
func (q *CardQuery) Only(ctx context.Context) (*Card, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{card.Label}
	default:
		return nil, &NotSingularError{card.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *CardQuery) OnlyX(ctx context.Context) *Card {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Card ID in the query.
// Returns a *NotSingularError when more than one Card ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *CardQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{card.Label}
	default:
		err = &NotSingularError{card.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *CardQuery) OnlyIDX(ctx context.Context) int {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cards.
func (q *CardQuery) All(ctx context.Context) ([]*Card, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Card, *CardQuery]()
	return withInterceptors[[]*Card](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *CardQuery) AllX(ctx context.Context) []*Card {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Card IDs.
func (q *CardQuery) IDs(ctx context.Context) (ids []int, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(card.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *CardQuery) IDsX(ctx context.Context) []int {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *CardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*CardQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *CardQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *CardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryExist)
	switch _, err := q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (q *CardQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *CardQuery) Clone() *CardQuery {
	if q == nil {
		return nil
	}
	return &CardQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]card.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.Card{}, q.predicates...),
		withOwner:  q.withOwner.Clone(),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (q *CardQuery) WithOwner(opts ...func(*UserQuery)) *CardQuery {
	query := (&UserClient{config: q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	q.withOwner = query
	return q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Number string `json:"number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Card.Query().
//		GroupBy(card.FieldNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *CardQuery) GroupBy(field string, fields ...string) *CardGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CardGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = card.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Number string `json:"number,omitempty"`
//	}
//
//	client.Card.Query().
//		Select(card.FieldNumber).
//		Scan(ctx, &v)
func (q *CardQuery) Select(fields ...string) *CardSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &CardSelect{CardQuery: q}
	sbuild.label = card.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CardSelect configured with the given aggregations.
func (q *CardQuery) Aggregate(fns ...AggregateFunc) *CardSelect {
	return q.Select().Aggregate(fns...)
}

func (q *CardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, q); err != nil {
				return err
			}
		}
	}
	for _, f := range q.ctx.Fields {
		if !card.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if q.path != nil {
		prev, err := q.path(ctx)
		if err != nil {
			return err
		}
		q.sql = prev
	}
	return nil
}

func (q *CardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Card, error) {
	var (
		nodes       = []*Card{}
		withFKs     = q.withFKs
		_spec       = q.querySpec()
		loadedTypes = [1]bool{
			q.withOwner != nil,
		}
	)
	if q.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, card.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Card).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Card{config: q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := q.withOwner; query != nil {
		if err := q.loadOwner(ctx, query, nodes, nil,
			func(n *Card, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (q *CardQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Card, init func(*Card), assign func(*Card, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Card)
	for i := range nodes {
		if nodes[i].user_cards == nil {
			continue
		}
		fk := *nodes[i].user_cards
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_cards" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (q *CardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *CardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for i := range fields {
			if fields[i] != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (q *CardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(card.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = card.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if q.sql != nil {
		selector = q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if q.ctx.Unique != nil && *q.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range q.predicates {
		p(selector)
	}
	for _, p := range q.order {
		p(selector)
	}
	if offset := q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CardGroupBy is the group-by builder for Card entities.
type CardGroupBy struct {
	selector
	build *CardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CardGroupBy) Aggregate(fns ...AggregateFunc) *CardGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CardQuery, *CardGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (q *CardGroupBy) sqlScan(ctx context.Context, root *CardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(q.fns))
	for _, fn := range q.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*q.flds)+len(q.fns))
		for _, f := range *q.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*q.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := q.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CardSelect is the builder for selecting fields of Card entities.
type CardSelect struct {
	*CardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CardSelect) Aggregate(fns ...AggregateFunc) *CardSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CardQuery, *CardSelect](ctx, cs.CardQuery, cs, cs.inters, v)
}

func (q *CardSelect) sqlScan(ctx context.Context, root *CardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(q.fns))
	for _, fn := range q.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*q.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := q.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
