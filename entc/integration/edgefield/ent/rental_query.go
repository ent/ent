// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RentalQuery is the builder for querying Rental entities.
type RentalQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Rental
	withUser   *UserQuery
	withCar    *CarQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RentalQuery builder.
func (rq *RentalQuery) Where(ps ...predicate.Rental) *RentalQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RentalQuery) Limit(limit int) *RentalQuery {
	rq.limit = &limit
	return rq
}

// Offset to start from.
func (rq *RentalQuery) Offset(offset int) *RentalQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RentalQuery) Unique(unique bool) *RentalQuery {
	rq.unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RentalQuery) Order(o ...OrderFunc) *RentalQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUser chains the current query on the "user" edge.
func (rq *RentalQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rental.Table, rental.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rental.UserTable, rental.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCar chains the current query on the "car" edge.
func (rq *RentalQuery) QueryCar() *CarQuery {
	query := (&CarClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rental.Table, rental.FieldID, selector),
			sqlgraph.To(car.Table, car.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rental.CarTable, rental.CarColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Rental entity from the query.
// Returns a *NotFoundError when no Rental was found.
func (rq *RentalQuery) First(ctx context.Context) (*Rental, error) {
	nodes, err := rq.Limit(1).All(newQueryContext(ctx, TypeRental, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rental.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RentalQuery) FirstX(ctx context.Context) *Rental {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Rental ID from the query.
// Returns a *NotFoundError when no Rental ID was found.
func (rq *RentalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(newQueryContext(ctx, TypeRental, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rental.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RentalQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Rental entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Rental entity is found.
// Returns a *NotFoundError when no Rental entities are found.
func (rq *RentalQuery) Only(ctx context.Context) (*Rental, error) {
	nodes, err := rq.Limit(2).All(newQueryContext(ctx, TypeRental, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rental.Label}
	default:
		return nil, &NotSingularError{rental.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RentalQuery) OnlyX(ctx context.Context) *Rental {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Rental ID in the query.
// Returns a *NotSingularError when more than one Rental ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RentalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(newQueryContext(ctx, TypeRental, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rental.Label}
	default:
		err = &NotSingularError{rental.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RentalQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Rentals.
func (rq *RentalQuery) All(ctx context.Context) ([]*Rental, error) {
	ctx = newQueryContext(ctx, TypeRental, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Rental, *RentalQuery]()
	return withInterceptors[[]*Rental](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RentalQuery) AllX(ctx context.Context) []*Rental {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Rental IDs.
func (rq *RentalQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeRental, "IDs")
	if err := rq.Select(rental.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RentalQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RentalQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeRental, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RentalQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RentalQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RentalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeRental, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RentalQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RentalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RentalQuery) Clone() *RentalQuery {
	if rq == nil {
		return nil
	}
	return &RentalQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		predicates: append([]predicate.Rental{}, rq.predicates...),
		withUser:   rq.withUser.Clone(),
		withCar:    rq.withCar.Clone(),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RentalQuery) WithUser(opts ...func(*UserQuery)) *RentalQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithCar tells the query-builder to eager-load the nodes that are connected to
// the "car" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RentalQuery) WithCar(opts ...func(*CarQuery)) *RentalQuery {
	query := (&CarClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withCar = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Rental.Query().
//		GroupBy(rental.FieldDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RentalQuery) GroupBy(field string, fields ...string) *RentalGroupBy {
	rq.fields = append([]string{field}, fields...)
	grbuild := &RentalGroupBy{build: rq}
	grbuild.flds = &rq.fields
	grbuild.label = rental.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//	}
//
//	client.Rental.Query().
//		Select(rental.FieldDate).
//		Scan(ctx, &v)
func (rq *RentalQuery) Select(fields ...string) *RentalSelect {
	rq.fields = append(rq.fields, fields...)
	sbuild := &RentalSelect{RentalQuery: rq}
	sbuild.label = rental.Label
	sbuild.flds, sbuild.scan = &rq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RentalSelect configured with the given aggregations.
func (rq *RentalQuery) Aggregate(fns ...AggregateFunc) *RentalSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RentalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.fields {
		if !rental.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RentalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Rental, error) {
	var (
		nodes       = []*Rental{}
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withUser != nil,
			rq.withCar != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Rental).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Rental{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Rental, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withCar; query != nil {
		if err := rq.loadCar(ctx, query, nodes, nil,
			func(n *Rental, e *Car) { n.Edges.Car = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RentalQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Rental, init func(*Rental), assign func(*Rental, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Rental)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RentalQuery) loadCar(ctx context.Context, query *CarQuery, nodes []*Rental, init func(*Rental), assign func(*Rental, *Car)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Rental)
	for i := range nodes {
		fk := nodes[i].CarID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(car.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "car_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RentalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.fields
	if len(rq.fields) > 0 {
		_spec.Unique = rq.unique != nil && *rq.unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RentalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rental.Table,
			Columns: rental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rental.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rental.FieldID)
		for i := range fields {
			if fields[i] != rental.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RentalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(rental.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = rental.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RentalGroupBy is the group-by builder for Rental entities.
type RentalGroupBy struct {
	selector
	build *RentalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RentalGroupBy) Aggregate(fns ...AggregateFunc) *RentalGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RentalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeRental, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RentalQuery, *RentalGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RentalGroupBy) sqlScan(ctx context.Context, root *RentalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RentalSelect is the builder for selecting fields of Rental entities.
type RentalSelect struct {
	*RentalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RentalSelect) Aggregate(fns ...AggregateFunc) *RentalSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RentalSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeRental, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RentalQuery, *RentalSelect](ctx, rs.RentalQuery, rs, rs.inters, v)
}

func (rs *RentalSelect) sqlScan(ctx context.Context, root *RentalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
