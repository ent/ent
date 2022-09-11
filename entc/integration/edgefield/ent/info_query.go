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
	"entgo.io/ent/entc/integration/edgefield/ent/info"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// InfoQuery is the builder for querying Info entities.
type InfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Info
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InfoQuery builder.
func (iq *InfoQuery) Where(ps ...predicate.Info) *InfoQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InfoQuery) Limit(limit int) *InfoQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InfoQuery) Offset(offset int) *InfoQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InfoQuery) Unique(unique bool) *InfoQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *InfoQuery) Order(o ...OrderFunc) *InfoQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryUser chains the current query on the "user" edge.
func (iq *InfoQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(info.Table, info.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, info.UserTable, info.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Info entity from the query.
// Returns a *NotFoundError when no Info was found.
func (iq *InfoQuery) First(ctx context.Context) (*Info, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{info.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InfoQuery) FirstX(ctx context.Context) *Info {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Info ID from the query.
// Returns a *NotFoundError when no Info ID was found.
func (iq *InfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{info.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InfoQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Info entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Info entity is found.
// Returns a *NotFoundError when no Info entities are found.
func (iq *InfoQuery) Only(ctx context.Context) (*Info, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{info.Label}
	default:
		return nil, &NotSingularError{info.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InfoQuery) OnlyX(ctx context.Context) *Info {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Info ID in the query.
// Returns a *NotSingularError when more than one Info ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{info.Label}
	default:
		err = &NotSingularError{info.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Infos.
func (iq *InfoQuery) All(ctx context.Context) ([]*Info, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InfoQuery) AllX(ctx context.Context) []*Info {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Info IDs.
func (iq *InfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(info.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InfoQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InfoQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InfoQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InfoQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InfoQuery) Clone() *InfoQuery {
	if iq == nil {
		return nil
	}
	return &InfoQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]OrderFunc{}, iq.order...),
		predicates: append([]predicate.Info{}, iq.predicates...),
		withUser:   iq.withUser.Clone(),
		// clone intermediate query.
		sql:    iq.sql.Clone(),
		path:   iq.path,
		unique: iq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InfoQuery) WithUser(opts ...func(*UserQuery)) *InfoQuery {
	query := &UserQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Content json.RawMessage `json:"content,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Info.Query().
//		GroupBy(info.FieldContent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InfoQuery) GroupBy(field string, fields ...string) *InfoGroupBy {
	grbuild := &InfoGroupBy{config: iq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	grbuild.label = info.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Content json.RawMessage `json:"content,omitempty"`
//	}
//
//	client.Info.Query().
//		Select(info.FieldContent).
//		Scan(ctx, &v)
func (iq *InfoQuery) Select(fields ...string) *InfoSelect {
	iq.fields = append(iq.fields, fields...)
	selbuild := &InfoSelect{InfoQuery: iq}
	selbuild.label = info.Label
	selbuild.flds, selbuild.scan = &iq.fields, selbuild.Scan
	return selbuild
}

func (iq *InfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !info.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Info, error) {
	var (
		nodes       = []*Info{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Info).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Info{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withUser; query != nil {
		if err := iq.loadUser(ctx, query, nodes, nil,
			func(n *Info, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InfoQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Info, init func(*Info), assign func(*Info, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Info)
	for i := range nodes {
		fk := nodes[i].ID
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
			return fmt.Errorf(`unexpected foreign-key "id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *InfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.fields
	if len(iq.fields) > 0 {
		_spec.Unique = iq.unique != nil && *iq.unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InfoQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (iq *InfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   info.Table,
			Columns: info.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: info.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, info.FieldID)
		for i := range fields {
			if fields[i] != info.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(info.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = info.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.unique != nil && *iq.unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InfoGroupBy is the group-by builder for Info entities.
type InfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InfoGroupBy) Aggregate(fns ...AggregateFunc) *InfoGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *InfoGroupBy) Scan(ctx context.Context, v any) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

func (igb *InfoGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range igb.fields {
		if !info.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InfoGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// InfoSelect is the builder for selecting fields of Info entities.
type InfoSelect struct {
	*InfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *InfoSelect) Scan(ctx context.Context, v any) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.InfoQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

func (is *InfoSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
