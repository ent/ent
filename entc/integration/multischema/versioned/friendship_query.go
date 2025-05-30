// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package versioned

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/multischema/versioned/friendship"
	"entgo.io/ent/entc/integration/multischema/versioned/internal"
	"entgo.io/ent/entc/integration/multischema/versioned/predicate"
	"entgo.io/ent/entc/integration/multischema/versioned/user"
	"entgo.io/ent/schema/field"
)

// FriendshipQuery is the builder for querying Friendship entities.
type FriendshipQuery struct {
	config
	ctx        *QueryContext
	order      []friendship.OrderOption
	inters     []Interceptor
	predicates []predicate.Friendship
	withUser   *UserQuery
	withFriend *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FriendshipQuery builder.
func (_q *FriendshipQuery) Where(ps ...predicate.Friendship) *FriendshipQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *FriendshipQuery) Limit(limit int) *FriendshipQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *FriendshipQuery) Offset(offset int) *FriendshipQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *FriendshipQuery) Unique(unique bool) *FriendshipQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *FriendshipQuery) Order(o ...friendship.OrderOption) *FriendshipQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryUser chains the current query on the "user" edge.
func (_q *FriendshipQuery) QueryUser() *UserQuery {
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
			sqlgraph.From(friendship.Table, friendship.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.UserTable, friendship.UserColumn),
		)
		schemaConfig := _q.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Friendship
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFriend chains the current query on the "friend" edge.
func (_q *FriendshipQuery) QueryFriend() *UserQuery {
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
			sqlgraph.From(friendship.Table, friendship.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.FriendTable, friendship.FriendColumn),
		)
		schemaConfig := _q.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.Friendship
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Friendship entity from the query.
// Returns a *NotFoundError when no Friendship was found.
func (_q *FriendshipQuery) First(ctx context.Context) (*Friendship, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{friendship.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *FriendshipQuery) FirstX(ctx context.Context) *Friendship {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Friendship ID from the query.
// Returns a *NotFoundError when no Friendship ID was found.
func (_q *FriendshipQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{friendship.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *FriendshipQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Friendship entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Friendship entity is found.
// Returns a *NotFoundError when no Friendship entities are found.
func (_q *FriendshipQuery) Only(ctx context.Context) (*Friendship, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{friendship.Label}
	default:
		return nil, &NotSingularError{friendship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *FriendshipQuery) OnlyX(ctx context.Context) *Friendship {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Friendship ID in the query.
// Returns a *NotSingularError when more than one Friendship ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *FriendshipQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{friendship.Label}
	default:
		err = &NotSingularError{friendship.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *FriendshipQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Friendships.
func (_q *FriendshipQuery) All(ctx context.Context) ([]*Friendship, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Friendship, *FriendshipQuery]()
	return withInterceptors[[]*Friendship](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *FriendshipQuery) AllX(ctx context.Context) []*Friendship {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Friendship IDs.
func (_q *FriendshipQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(friendship.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *FriendshipQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *FriendshipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*FriendshipQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *FriendshipQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *FriendshipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("versioned: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *FriendshipQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FriendshipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *FriendshipQuery) Clone() *FriendshipQuery {
	if _q == nil {
		return nil
	}
	return &FriendshipQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]friendship.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Friendship{}, _q.predicates...),
		withUser:   _q.withUser.Clone(),
		withFriend: _q.withFriend.Clone(),
		// clone intermediate query.
		sql:       _q.sql.Clone(),
		path:      _q.path,
		modifiers: append([]func(*sql.Selector){}, _q.modifiers...),
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *FriendshipQuery) WithUser(opts ...func(*UserQuery)) *FriendshipQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withUser = query
	return _q
}

// WithFriend tells the query-builder to eager-load the nodes that are connected to
// the "friend" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *FriendshipQuery) WithFriend(opts ...func(*UserQuery)) *FriendshipQuery {
	query := (&UserClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withFriend = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Friendship.Query().
//		GroupBy(friendship.FieldWeight).
//		Aggregate(versioned.Count()).
//		Scan(ctx, &v)
func (_q *FriendshipQuery) GroupBy(field string, fields ...string) *FriendshipGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FriendshipGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = friendship.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//	}
//
//	client.Friendship.Query().
//		Select(friendship.FieldWeight).
//		Scan(ctx, &v)
func (_q *FriendshipQuery) Select(fields ...string) *FriendshipSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &FriendshipSelect{FriendshipQuery: _q}
	sbuild.label = friendship.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FriendshipSelect configured with the given aggregations.
func (_q *FriendshipQuery) Aggregate(fns ...AggregateFunc) *FriendshipSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *FriendshipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("versioned: uninitialized interceptor (forgotten import versioned/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	for _, f := range _q.ctx.Fields {
		if !friendship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("versioned: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	return nil
}

func (_q *FriendshipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Friendship, error) {
	var (
		nodes       = []*Friendship{}
		_spec       = _q.querySpec()
		loadedTypes = [2]bool{
			_q.withUser != nil,
			_q.withFriend != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Friendship).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Friendship{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = _q.schemaConfig.Friendship
	ctx = internal.NewSchemaConfigContext(ctx, _q.schemaConfig)
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
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
	if query := _q.withUser; query != nil {
		if err := _q.loadUser(ctx, query, nodes, nil,
			func(n *Friendship, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := _q.withFriend; query != nil {
		if err := _q.loadFriend(ctx, query, nodes, nil,
			func(n *Friendship, e *User) { n.Edges.Friend = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *FriendshipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Friendship, init func(*Friendship), assign func(*Friendship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Friendship)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (_q *FriendshipQuery) loadFriend(ctx context.Context, query *UserQuery, nodes []*Friendship, init func(*Friendship), assign func(*Friendship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Friendship)
	for i := range nodes {
		fk := nodes[i].FriendID
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
			return fmt.Errorf(`unexpected foreign-key "friend_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (_q *FriendshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Schema = _q.schemaConfig.Friendship
	ctx = internal.NewSchemaConfigContext(ctx, _q.schemaConfig)
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *FriendshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for i := range fields {
			if fields[i] != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if _q.withUser != nil {
			_spec.Node.AddColumnOnce(friendship.FieldUserID)
		}
		if _q.withFriend != nil {
			_spec.Node.AddColumnOnce(friendship.FieldFriendID)
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

func (_q *FriendshipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(friendship.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = friendship.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(_q.schemaConfig.Friendship)
	ctx = internal.NewSchemaConfigContext(ctx, _q.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range _q.modifiers {
		m(selector)
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

// Modify adds a query modifier for attaching custom logic to queries.
func (_q *FriendshipQuery) Modify(modifiers ...func(s *sql.Selector)) *FriendshipSelect {
	_q.modifiers = append(_q.modifiers, modifiers...)
	return _q.Select()
}

// FriendshipGroupBy is the group-by builder for Friendship entities.
type FriendshipGroupBy struct {
	selector
	build *FriendshipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (_g *FriendshipGroupBy) Aggregate(fns ...AggregateFunc) *FriendshipGroupBy {
	_g.fns = append(_g.fns, fns...)
	return _g
}

// Scan applies the selector query and scans the result into the given value.
func (_g *FriendshipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _g.build.ctx, ent.OpQueryGroupBy)
	if err := _g.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendshipQuery, *FriendshipGroupBy](ctx, _g.build, _g, _g.build.inters, v)
}

func (_g *FriendshipGroupBy) sqlScan(ctx context.Context, root *FriendshipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(_g.fns))
	for _, fn := range _g.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*_g.flds)+len(_g.fns))
		for _, f := range *_g.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*_g.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := _g.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FriendshipSelect is the builder for selecting fields of Friendship entities.
type FriendshipSelect struct {
	*FriendshipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (_s *FriendshipSelect) Aggregate(fns ...AggregateFunc) *FriendshipSelect {
	_s.fns = append(_s.fns, fns...)
	return _s
}

// Scan applies the selector query and scans the result into the given value.
func (_s *FriendshipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, _s.ctx, ent.OpQuerySelect)
	if err := _s.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendshipQuery, *FriendshipSelect](ctx, _s.FriendshipQuery, _s, _s.inters, v)
}

func (_s *FriendshipSelect) sqlScan(ctx context.Context, root *FriendshipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(_s.fns))
	for _, fn := range _s.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*_s.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := _s.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (_s *FriendshipSelect) Modify(modifiers ...func(s *sql.Selector)) *FriendshipSelect {
	_s.modifiers = append(_s.modifiers, modifiers...)
	return _s
}
