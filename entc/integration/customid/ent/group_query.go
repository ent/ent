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
	"entgo.io/ent/entc/integration/customid/ent/group"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/user"
	"entgo.io/ent/schema/field"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
	config
	ctx        *QueryContext
	order      []group.OrderOption
	inters     []Interceptor
	predicates []predicate.Group
	withUsers  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupQuery builder.
func (q *GroupQuery) Where(ps ...predicate.Group) *GroupQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *GroupQuery) Limit(limit int) *GroupQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *GroupQuery) Offset(offset int) *GroupQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *GroupQuery) Unique(unique bool) *GroupQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *GroupQuery) Order(o ...group.OrderOption) *GroupQuery {
	q.order = append(q.order, o...)
	return q
}

// QueryUsers chains the current query on the "users" edge.
func (q *GroupQuery) QueryUsers() *UserQuery {
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
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Group entity from the query.
// Returns a *NotFoundError when no Group was found.
func (q *GroupQuery) First(ctx context.Context) (*Group, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{group.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *GroupQuery) FirstX(ctx context.Context) *Group {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Group ID from the query.
// Returns a *NotFoundError when no Group ID was found.
func (q *GroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{group.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *GroupQuery) FirstIDX(ctx context.Context) int {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Group entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Group entity is found.
// Returns a *NotFoundError when no Group entities are found.
func (q *GroupQuery) Only(ctx context.Context) (*Group, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
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
func (q *GroupQuery) OnlyX(ctx context.Context) *Group {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Group ID in the query.
// Returns a *NotSingularError when more than one Group ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *GroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
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
func (q *GroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (q *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Group, *GroupQuery]()
	return withInterceptors[[]*Group](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *GroupQuery) AllX(ctx context.Context) []*Group {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Group IDs.
func (q *GroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(group.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *GroupQuery) IDsX(ctx context.Context) []int {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *GroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*GroupQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *GroupQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *GroupQuery) Exist(ctx context.Context) (bool, error) {
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
func (q *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *GroupQuery) Clone() *GroupQuery {
	if q == nil {
		return nil
	}
	return &GroupQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]group.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.Group{}, q.predicates...),
		withUsers:  q.withUsers.Clone(),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (q *GroupQuery) WithUsers(opts ...func(*UserQuery)) *GroupQuery {
	query := (&UserClient{config: q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	q.withUsers = query
	return q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (q *GroupQuery) GroupBy(field string, fields ...string) *GroupGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = group.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (q *GroupQuery) Select(fields ...string) *GroupSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &GroupSelect{GroupQuery: q}
	sbuild.label = group.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupSelect configured with the given aggregations.
func (q *GroupQuery) Aggregate(fns ...AggregateFunc) *GroupSelect {
	return q.Select().Aggregate(fns...)
}

func (q *GroupQuery) prepareQuery(ctx context.Context) error {
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
		if !group.ValidColumn(f) {
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

func (q *GroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Group, error) {
	var (
		nodes       = []*Group{}
		_spec       = q.querySpec()
		loadedTypes = [1]bool{
			q.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Group).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Group{config: q.config}
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
	if query := q.withUsers; query != nil {
		if err := q.loadUsers(ctx, query, nodes,
			func(n *Group) { n.Edges.Users = []*User{} },
			func(n *Group, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (q *GroupQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Group, init func(*Group), assign func(*Group, *User)) error {
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
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(group.UsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(group.UsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(group.UsersPrimaryKey[0]))
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

func (q *GroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *GroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for i := range fields {
			if fields[i] != group.FieldID {
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

func (q *GroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(group.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = group.Columns
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

func (q *GroupGroupBy) sqlScan(ctx context.Context, root *GroupQuery, v any) error {
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

func (q *GroupSelect) sqlScan(ctx context.Context, root *GroupQuery, v any) error {
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
