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
	"entgo.io/ent/examples/edgeindex/ent/city"
	"entgo.io/ent/examples/edgeindex/ent/predicate"
	"entgo.io/ent/examples/edgeindex/ent/street"
	"entgo.io/ent/schema/field"
)

// CityQuery is the builder for querying City entities.
type CityQuery struct {
	config
	ctx         *QueryContext
	order       []city.OrderOption
	inters      []Interceptor
	predicates  []predicate.City
	withStreets *StreetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CityQuery builder.
func (q *CityQuery) Where(ps ...predicate.City) *CityQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *CityQuery) Limit(limit int) *CityQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *CityQuery) Offset(offset int) *CityQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *CityQuery) Unique(unique bool) *CityQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *CityQuery) Order(o ...city.OrderOption) *CityQuery {
	q.order = append(q.order, o...)
	return q
}

// QueryStreets chains the current query on the "streets" edge.
func (q *CityQuery) QueryStreets() *StreetQuery {
	query := (&StreetClient{config: q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(city.Table, city.FieldID, selector),
			sqlgraph.To(street.Table, street.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, city.StreetsTable, city.StreetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first City entity from the query.
// Returns a *NotFoundError when no City was found.
func (q *CityQuery) First(ctx context.Context) (*City, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{city.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *CityQuery) FirstX(ctx context.Context) *City {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first City ID from the query.
// Returns a *NotFoundError when no City ID was found.
func (q *CityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{city.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *CityQuery) FirstIDX(ctx context.Context) int {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single City entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one City entity is found.
// Returns a *NotFoundError when no City entities are found.
func (q *CityQuery) Only(ctx context.Context) (*City, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{city.Label}
	default:
		return nil, &NotSingularError{city.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *CityQuery) OnlyX(ctx context.Context) *City {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only City ID in the query.
// Returns a *NotSingularError when more than one City ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *CityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{city.Label}
	default:
		err = &NotSingularError{city.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *CityQuery) OnlyIDX(ctx context.Context) int {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cities.
func (q *CityQuery) All(ctx context.Context) ([]*City, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*City, *CityQuery]()
	return withInterceptors[[]*City](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *CityQuery) AllX(ctx context.Context) []*City {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of City IDs.
func (q *CityQuery) IDs(ctx context.Context) (ids []int, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(city.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *CityQuery) IDsX(ctx context.Context) []int {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *CityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*CityQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *CityQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *CityQuery) Exist(ctx context.Context) (bool, error) {
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
func (q *CityQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *CityQuery) Clone() *CityQuery {
	if q == nil {
		return nil
	}
	return &CityQuery{
		config:      q.config,
		ctx:         q.ctx.Clone(),
		order:       append([]city.OrderOption{}, q.order...),
		inters:      append([]Interceptor{}, q.inters...),
		predicates:  append([]predicate.City{}, q.predicates...),
		withStreets: q.withStreets.Clone(),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
}

// WithStreets tells the query-builder to eager-load the nodes that are connected to
// the "streets" edge. The optional arguments are used to configure the query builder of the edge.
func (q *CityQuery) WithStreets(opts ...func(*StreetQuery)) *CityQuery {
	query := (&StreetClient{config: q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	q.withStreets = query
	return q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.City.Query().
//		GroupBy(city.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *CityQuery) GroupBy(field string, fields ...string) *CityGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CityGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = city.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.City.Query().
//		Select(city.FieldName).
//		Scan(ctx, &v)
func (q *CityQuery) Select(fields ...string) *CitySelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &CitySelect{CityQuery: q}
	sbuild.label = city.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CitySelect configured with the given aggregations.
func (q *CityQuery) Aggregate(fns ...AggregateFunc) *CitySelect {
	return q.Select().Aggregate(fns...)
}

func (q *CityQuery) prepareQuery(ctx context.Context) error {
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
		if !city.ValidColumn(f) {
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

func (q *CityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*City, error) {
	var (
		nodes       = []*City{}
		_spec       = q.querySpec()
		loadedTypes = [1]bool{
			q.withStreets != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*City).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &City{config: q.config}
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
	if query := q.withStreets; query != nil {
		if err := q.loadStreets(ctx, query, nodes,
			func(n *City) { n.Edges.Streets = []*Street{} },
			func(n *City, e *Street) { n.Edges.Streets = append(n.Edges.Streets, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (q *CityQuery) loadStreets(ctx context.Context, query *StreetQuery, nodes []*City, init func(*City), assign func(*City, *Street)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*City)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Street(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(city.StreetsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.city_streets
		if fk == nil {
			return fmt.Errorf(`foreign-key "city_streets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "city_streets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (q *CityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *CityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(city.Table, city.Columns, sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, city.FieldID)
		for i := range fields {
			if fields[i] != city.FieldID {
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

func (q *CityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(city.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = city.Columns
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

// CityGroupBy is the group-by builder for City entities.
type CityGroupBy struct {
	selector
	build *CityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CityGroupBy) Aggregate(fns ...AggregateFunc) *CityGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CityQuery, *CityGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (q *CityGroupBy) sqlScan(ctx context.Context, root *CityQuery, v any) error {
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

// CitySelect is the builder for selecting fields of City entities.
type CitySelect struct {
	*CityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CitySelect) Aggregate(fns ...AggregateFunc) *CitySelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CityQuery, *CitySelect](ctx, cs.CityQuery, cs, cs.inters, v)
}

func (q *CitySelect) sqlScan(ctx context.Context, root *CityQuery, v any) error {
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
