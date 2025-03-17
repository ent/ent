// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/privacytenant/ent/group"
	"entgo.io/ent/examples/privacytenant/ent/predicate"
	"entgo.io/ent/examples/privacytenant/ent/tenant"
	"entgo.io/ent/examples/privacytenant/ent/user"
	"entgo.io/ent/schema/field"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
	config
	ctx        *QueryContext
	order      []group.OrderOption
	inters     []Interceptor
	predicates []predicate.Group
	withTenant *TenantQuery
	withUsers  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupQuery builder.
func (_q *GroupQuery) Where(ps ...predicate.Group) *GroupQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *GroupQuery) Limit(limit int) *GroupQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *GroupQuery) Offset(offset int) *GroupQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *GroupQuery) Unique(unique bool) *GroupQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *GroupQuery) Order(o ...group.OrderOption) *GroupQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryTenant chains the current query on the "tenant" edge.
func (_q *GroupQuery) QueryTenant() *TenantQuery {
	query := (&TenantClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(tenant.Table, tenant.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, group.TenantTable, group.TenantColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (_q *GroupQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Group entity from the query.
// Returns a *NotFoundError when no Group was found.
func (_q *GroupQuery) First(ctx context.Context) (*Group, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{group.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *GroupQuery) FirstX(ctx context.Context) *Group {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Group ID from the query.
// Returns a *NotFoundError when no Group ID was found.
func (_q *GroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{group.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *GroupQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Group entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Group entity is found.
// Returns a *NotFoundError when no Group entities are found.
func (_q *GroupQuery) Only(ctx context.Context) (*Group, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{group.Label}
	default:
		return nil, &NotSingularError{group.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *GroupQuery) OnlyX(ctx context.Context) *Group {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Group ID in the query.
// Returns a *NotSingularError when more than one Group ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *GroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = &NotSingularError{group.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *GroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (_q *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Group, *GroupQuery]()
	return withInterceptors[[]*Group](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *GroupQuery) AllX(ctx context.Context) []*Group {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Group IDs.
func (_q *GroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(group.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *GroupQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *GroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*GroupQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *GroupQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *GroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *GroupQuery) Clone() *GroupQuery {
	if _q == nil {
		return nil
	}
	return &GroupQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]group.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Group{}, _q.predicates...),
		withTenant: _q.withTenant.Clone(),
		withUsers:  _q.withUsers.Clone(),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// WithTenant tells the query-builder to eager-load the nodes that are connected to
// the "tenant" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *GroupQuery) WithTenant(opts ...func(*TenantQuery)) *GroupQuery {
	query := (&TenantClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withTenant = query
	return _q
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *GroupQuery) WithUsers(opts ...func(*UserQuery)) *GroupQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withUsers = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TenantID int `json:"tenant_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Group.Query().
//		GroupBy(group.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *GroupQuery) GroupBy(field string, fields ...string) *GroupGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = group.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TenantID int `json:"tenant_id,omitempty"`
//	}
//
//	client.Group.Query().
//		Select(group.FieldTenantID).
//		Scan(ctx, &v)
func (_q *GroupQuery) Select(fields ...string) *GroupSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &GroupSelect{GroupQuery: _q}
	sbuild.label = group.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupSelect configured with the given aggregations.
func (_q *GroupQuery) Aggregate(fns ...AggregateFunc) *GroupSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *GroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !group.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	if group.Policy == nil {
		return errors.New("ent: uninitialized group.Policy (forgotten import ent/runtime?)")
	}
	if err := group.Policy.EvalQuery(ctx, _q); err != nil {
		return err
	}
	return nil
}

func (_q *GroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Group, error) {
	var (
		nodes       = []*Group{}
		_spec       = _q.querySpec()
		loadedTypes = [2]bool{
			_q.withTenant != nil,
			_q.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Group).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Group{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := _q.withTenant; query != nil {
		if err := _q.loadTenant(ctx, query, nodes, nil,
			func(n *Group, e *Tenant) { n.Edges.Tenant = e }); err != nil {
			return nil, err
		}
	}
	if query := _q.withUsers; query != nil {
		if err := _q.loadUsers(ctx, query, nodes,
			func(n *Group) { n.Edges.Users = []*User{} },
			func(n *Group, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *GroupQuery) loadTenant(ctx context.Context, query *TenantQuery, nodes []*Group, init func(*Group), assign func(*Group, *Tenant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Group)
	for i := range nodes {
		fk := nodes[i].TenantID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tenant.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tenant_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (_q *GroupQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Group, init func(*Group), assign func(*Group, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Group)
	nids := make(map[int]map[*Group]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(group.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(group.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(group.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(group.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Group]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (_q *GroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *GroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for i := range fields {
			if fields[i] != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if _q.withTenant != nil {
			_spec.Node.AddColumnOnce(group.FieldTenantID)
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *GroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(group.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = group.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupGroupBy is the group-by builder for Group entities.
type GroupGroupBy struct {
	selector
	build *GroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GroupGroupBy) Aggregate(fns ...AggregateFunc) *GroupGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, ent.OpQueryGroupBy)
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupQuery, *GroupGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GroupGroupBy) sqlScan(ctx context.Context, root *GroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupSelect is the builder for selecting fields of Group entities.
type GroupSelect struct {
	*GroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GroupSelect) Aggregate(fns ...AggregateFunc) *GroupSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, ent.OpQuerySelect)
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupQuery, *GroupSelect](ctx, gs.GroupQuery, gs, gs.inters, v)
}

func (gs *GroupSelect) sqlScan(ctx context.Context, root *GroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
