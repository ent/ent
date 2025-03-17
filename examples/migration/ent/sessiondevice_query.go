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
	"entgo.io/ent/examples/migration/ent/predicate"
	"entgo.io/ent/examples/migration/ent/session"
	"entgo.io/ent/examples/migration/ent/sessiondevice"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SessionDeviceQuery is the builder for querying SessionDevice entities.
type SessionDeviceQuery struct {
	config
	ctx          *QueryContext
	order        []sessiondevice.OrderOption
	inters       []Interceptor
	predicates   []predicate.SessionDevice
	withSessions *SessionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SessionDeviceQuery builder.
func (q *SessionDeviceQuery) Where(ps ...predicate.SessionDevice) *SessionDeviceQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *SessionDeviceQuery) Limit(limit int) *SessionDeviceQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *SessionDeviceQuery) Offset(offset int) *SessionDeviceQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *SessionDeviceQuery) Unique(unique bool) *SessionDeviceQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *SessionDeviceQuery) Order(o ...sessiondevice.OrderOption) *SessionDeviceQuery {
	q.order = append(q.order, o...)
	return q
}

// QuerySessions chains the current query on the "sessions" edge.
func (q *SessionDeviceQuery) QuerySessions() *SessionQuery {
	query := (&SessionClient{config: q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sessiondevice.Table, sessiondevice.FieldID, selector),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sessiondevice.SessionsTable, sessiondevice.SessionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SessionDevice entity from the query.
// Returns a *NotFoundError when no SessionDevice was found.
func (q *SessionDeviceQuery) First(ctx context.Context) (*SessionDevice, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sessiondevice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *SessionDeviceQuery) FirstX(ctx context.Context) *SessionDevice {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SessionDevice ID from the query.
// Returns a *NotFoundError when no SessionDevice ID was found.
func (q *SessionDeviceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sessiondevice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *SessionDeviceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SessionDevice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SessionDevice entity is found.
// Returns a *NotFoundError when no SessionDevice entities are found.
func (q *SessionDeviceQuery) Only(ctx context.Context) (*SessionDevice, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sessiondevice.Label}
	default:
		return nil, &NotSingularError{sessiondevice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *SessionDeviceQuery) OnlyX(ctx context.Context) *SessionDevice {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SessionDevice ID in the query.
// Returns a *NotSingularError when more than one SessionDevice ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *SessionDeviceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sessiondevice.Label}
	default:
		err = &NotSingularError{sessiondevice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *SessionDeviceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SessionDevices.
func (q *SessionDeviceQuery) All(ctx context.Context) ([]*SessionDevice, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SessionDevice, *SessionDeviceQuery]()
	return withInterceptors[[]*SessionDevice](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *SessionDeviceQuery) AllX(ctx context.Context) []*SessionDevice {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SessionDevice IDs.
func (q *SessionDeviceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(sessiondevice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *SessionDeviceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *SessionDeviceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*SessionDeviceQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *SessionDeviceQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *SessionDeviceQuery) Exist(ctx context.Context) (bool, error) {
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
func (q *SessionDeviceQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SessionDeviceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *SessionDeviceQuery) Clone() *SessionDeviceQuery {
	if q == nil {
		return nil
	}
	return &SessionDeviceQuery{
		config:       q.config,
		ctx:          q.ctx.Clone(),
		order:        append([]sessiondevice.OrderOption{}, q.order...),
		inters:       append([]Interceptor{}, q.inters...),
		predicates:   append([]predicate.SessionDevice{}, q.predicates...),
		withSessions: q.withSessions.Clone(),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
}

// WithSessions tells the query-builder to eager-load the nodes that are connected to
// the "sessions" edge. The optional arguments are used to configure the query builder of the edge.
func (q *SessionDeviceQuery) WithSessions(opts ...func(*SessionQuery)) *SessionDeviceQuery {
	query := (&SessionClient{config: q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	q.withSessions = query
	return q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IPAddress string `json:"ip_address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SessionDevice.Query().
//		GroupBy(sessiondevice.FieldIPAddress).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *SessionDeviceQuery) GroupBy(field string, fields ...string) *SessionDeviceGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SessionDeviceGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = sessiondevice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IPAddress string `json:"ip_address,omitempty"`
//	}
//
//	client.SessionDevice.Query().
//		Select(sessiondevice.FieldIPAddress).
//		Scan(ctx, &v)
func (q *SessionDeviceQuery) Select(fields ...string) *SessionDeviceSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &SessionDeviceSelect{SessionDeviceQuery: q}
	sbuild.label = sessiondevice.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SessionDeviceSelect configured with the given aggregations.
func (q *SessionDeviceQuery) Aggregate(fns ...AggregateFunc) *SessionDeviceSelect {
	return q.Select().Aggregate(fns...)
}

func (q *SessionDeviceQuery) prepareQuery(ctx context.Context) error {
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
		if !sessiondevice.ValidColumn(f) {
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

func (q *SessionDeviceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SessionDevice, error) {
	var (
		nodes       = []*SessionDevice{}
		_spec       = q.querySpec()
		loadedTypes = [1]bool{
			q.withSessions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SessionDevice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SessionDevice{config: q.config}
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
	if query := q.withSessions; query != nil {
		if err := q.loadSessions(ctx, query, nodes,
			func(n *SessionDevice) { n.Edges.Sessions = []*Session{} },
			func(n *SessionDevice, e *Session) { n.Edges.Sessions = append(n.Edges.Sessions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (q *SessionDeviceQuery) loadSessions(ctx context.Context, query *SessionQuery, nodes []*SessionDevice, init func(*SessionDevice), assign func(*SessionDevice, *Session)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*SessionDevice)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(session.FieldDeviceID)
	}
	query.Where(predicate.Session(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sessiondevice.SessionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DeviceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "device_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (q *SessionDeviceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *SessionDeviceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sessiondevice.Table, sessiondevice.Columns, sqlgraph.NewFieldSpec(sessiondevice.FieldID, field.TypeUUID))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sessiondevice.FieldID)
		for i := range fields {
			if fields[i] != sessiondevice.FieldID {
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

func (q *SessionDeviceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(sessiondevice.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = sessiondevice.Columns
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

// SessionDeviceGroupBy is the group-by builder for SessionDevice entities.
type SessionDeviceGroupBy struct {
	selector
	build *SessionDeviceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdgb *SessionDeviceGroupBy) Aggregate(fns ...AggregateFunc) *SessionDeviceGroupBy {
	sdgb.fns = append(sdgb.fns, fns...)
	return sdgb
}

// Scan applies the selector query and scans the result into the given value.
func (sdgb *SessionDeviceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sdgb.build.ctx, ent.OpQueryGroupBy)
	if err := sdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SessionDeviceQuery, *SessionDeviceGroupBy](ctx, sdgb.build, sdgb, sdgb.build.inters, v)
}

func (q *SessionDeviceGroupBy) sqlScan(ctx context.Context, root *SessionDeviceQuery, v any) error {
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

// SessionDeviceSelect is the builder for selecting fields of SessionDevice entities.
type SessionDeviceSelect struct {
	*SessionDeviceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sds *SessionDeviceSelect) Aggregate(fns ...AggregateFunc) *SessionDeviceSelect {
	sds.fns = append(sds.fns, fns...)
	return sds
}

// Scan applies the selector query and scans the result into the given value.
func (sds *SessionDeviceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sds.ctx, ent.OpQuerySelect)
	if err := sds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SessionDeviceQuery, *SessionDeviceSelect](ctx, sds.SessionDeviceQuery, sds, sds.inters, v)
}

func (q *SessionDeviceSelect) sqlScan(ctx context.Context, root *SessionDeviceQuery, v any) error {
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
