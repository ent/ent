// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/card"
	"entgo.io/ent/examples/migration/ent/predicate"
	"entgo.io/ent/examples/migration/ent/user"
	"entgo.io/ent/schema/field"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx        *QueryContext
	order      []user.OrderOption
	inters     []Interceptor
	predicates []predicate.User
	withCards  *CardQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (q *UserQuery) Where(ps ...predicate.User) *UserQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *UserQuery) Limit(limit int) *UserQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *UserQuery) Offset(offset int) *UserQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *UserQuery) Unique(unique bool) *UserQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	q.order = append(q.order, o...)
	return q
}

// QueryCards chains the current query on the "cards" edge.
func (q *UserQuery) QueryCards() *CardQuery {
	query := (&CardClient{config: q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CardsTable, user.CardsColumn),
		)
		fromU = sqlgraph.SetNeighbors(q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (q *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *UserQuery) FirstX(ctx context.Context) *User {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (q *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (q *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (q *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (q *UserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*UserQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *UserQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *UserQuery) Exist(ctx context.Context) (bool, error) {
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
func (q *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *UserQuery) Clone() *UserQuery {
	if q == nil {
		return nil
	}
	return &UserQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]user.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.User{}, q.predicates...),
		withCards:  q.withCards.Clone(),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
}

// WithCards tells the query-builder to eager-load the nodes that are connected to
// the "cards" edge. The optional arguments are used to configure the query builder of the edge.
func (q *UserQuery) WithCards(opts ...func(*CardQuery)) *UserQuery {
	query := (&CardClient{config: q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	q.withCards = query
	return q
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
//	client.User.Query().
//		GroupBy(user.FieldAge).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
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
//	client.User.Query().
//		Select(user.FieldAge).
//		Scan(ctx, &v)
func (q *UserQuery) Select(fields ...string) *UserSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: q}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (q *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return q.Select().Aggregate(fns...)
}

func (q *UserQuery) prepareQuery(ctx context.Context) error {
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
		if !user.ValidColumn(f) {
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

func (q *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = q.querySpec()
		loadedTypes = [1]bool{
			q.withCards != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: q.config}
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
	if query := q.withCards; query != nil {
		if err := q.loadCards(ctx, query, nodes,
			func(n *User) { n.Edges.Cards = []*Card{} },
			func(n *User, e *Card) { n.Edges.Cards = append(n.Edges.Cards, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (q *UserQuery) loadCards(ctx context.Context, query *CardQuery, nodes []*User, init func(*User), assign func(*User, *Card)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(card.FieldOwnerID)
	}
	query.Where(predicate.Card(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.CardsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OwnerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "owner_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (q *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
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

func (q *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
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

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (q *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
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

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (q *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
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
