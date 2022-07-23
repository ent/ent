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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/sid"
	"entgo.io/ent/schema/field"
)

// IntSIDQuery is the builder for querying IntSID entities.
type IntSIDQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.IntSID
	// eager-loading edges.
	withParent   *IntSIDQuery
	withChildren *IntSIDQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IntSIDQuery builder.
func (isq *IntSIDQuery) Where(ps ...predicate.IntSID) *IntSIDQuery {
	isq.predicates = append(isq.predicates, ps...)
	return isq
}

// Limit adds a limit step to the query.
func (isq *IntSIDQuery) Limit(limit int) *IntSIDQuery {
	isq.limit = &limit
	return isq
}

// Offset adds an offset step to the query.
func (isq *IntSIDQuery) Offset(offset int) *IntSIDQuery {
	isq.offset = &offset
	return isq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (isq *IntSIDQuery) Unique(unique bool) *IntSIDQuery {
	isq.unique = &unique
	return isq
}

// Order adds an order step to the query.
func (isq *IntSIDQuery) Order(o ...OrderFunc) *IntSIDQuery {
	isq.order = append(isq.order, o...)
	return isq
}

// QueryParent chains the current query on the "parent" edge.
func (isq *IntSIDQuery) QueryParent() *IntSIDQuery {
	query := &IntSIDQuery{config: isq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(intsid.Table, intsid.FieldID, selector),
			sqlgraph.To(intsid.Table, intsid.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, intsid.ParentTable, intsid.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(isq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (isq *IntSIDQuery) QueryChildren() *IntSIDQuery {
	query := &IntSIDQuery{config: isq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(intsid.Table, intsid.FieldID, selector),
			sqlgraph.To(intsid.Table, intsid.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, intsid.ChildrenTable, intsid.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(isq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IntSID entity from the query.
// Returns a *NotFoundError when no IntSID was found.
func (isq *IntSIDQuery) First(ctx context.Context) (*IntSID, error) {
	nodes, err := isq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{intsid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (isq *IntSIDQuery) FirstX(ctx context.Context) *IntSID {
	node, err := isq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IntSID ID from the query.
// Returns a *NotFoundError when no IntSID ID was found.
func (isq *IntSIDQuery) FirstID(ctx context.Context) (id sid.ID, err error) {
	ids, err := isq.Limit(1).IDs(ctx)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{intsid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (isq *IntSIDQuery) FirstIDX(ctx context.Context) sid.ID {
	id, err := isq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IntSID entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IntSID entity is found.
// Returns a *NotFoundError when no IntSID entities are found.
func (isq *IntSIDQuery) Only(ctx context.Context) (*IntSID, error) {
	nodes, err := isq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{intsid.Label}
	default:
		return nil, &NotSingularError{intsid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (isq *IntSIDQuery) OnlyX(ctx context.Context) *IntSID {
	node, err := isq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IntSID ID in the query.
// Returns a *NotSingularError when more than one IntSID ID is found.
// Returns a *NotFoundError when no entities are found.
func (isq *IntSIDQuery) OnlyID(ctx context.Context) (id sid.ID, err error) {
	ids, err := isq.Limit(2).IDs(ctx)
	if err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{intsid.Label}
	default:
		err = &NotSingularError{intsid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (isq *IntSIDQuery) OnlyIDX(ctx context.Context) sid.ID {
	id, err := isq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IntSIDs.
func (isq *IntSIDQuery) All(ctx context.Context) ([]*IntSID, error) {
	if err := isq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return isq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (isq *IntSIDQuery) AllX(ctx context.Context) []*IntSID {
	nodes, err := isq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IntSID IDs.
func (isq *IntSIDQuery) IDs(ctx context.Context) ([]sid.ID, error) {
	var ids []sid.ID
	if err := isq.Select(intsid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (isq *IntSIDQuery) IDsX(ctx context.Context) []sid.ID {
	ids, err := isq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (isq *IntSIDQuery) Count(ctx context.Context) (int, error) {
	if err := isq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return isq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (isq *IntSIDQuery) CountX(ctx context.Context) int {
	count, err := isq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (isq *IntSIDQuery) Exist(ctx context.Context) (bool, error) {
	if err := isq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return isq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (isq *IntSIDQuery) ExistX(ctx context.Context) bool {
	exist, err := isq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IntSIDQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (isq *IntSIDQuery) Clone() *IntSIDQuery {
	if isq == nil {
		return nil
	}
	return &IntSIDQuery{
		config:       isq.config,
		limit:        isq.limit,
		offset:       isq.offset,
		order:        append([]OrderFunc{}, isq.order...),
		predicates:   append([]predicate.IntSID{}, isq.predicates...),
		withParent:   isq.withParent.Clone(),
		withChildren: isq.withChildren.Clone(),
		// clone intermediate query.
		sql:    isq.sql.Clone(),
		path:   isq.path,
		unique: isq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (isq *IntSIDQuery) WithParent(opts ...func(*IntSIDQuery)) *IntSIDQuery {
	query := &IntSIDQuery{config: isq.config}
	for _, opt := range opts {
		opt(query)
	}
	isq.withParent = query
	return isq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (isq *IntSIDQuery) WithChildren(opts ...func(*IntSIDQuery)) *IntSIDQuery {
	query := &IntSIDQuery{config: isq.config}
	for _, opt := range opts {
		opt(query)
	}
	isq.withChildren = query
	return isq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (isq *IntSIDQuery) GroupBy(field string, fields ...string) *IntSIDGroupBy {
	grbuild := &IntSIDGroupBy{config: isq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return isq.sqlQuery(ctx), nil
	}
	grbuild.label = intsid.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (isq *IntSIDQuery) Select(fields ...string) *IntSIDSelect {
	isq.fields = append(isq.fields, fields...)
	selbuild := &IntSIDSelect{IntSIDQuery: isq}
	selbuild.label = intsid.Label
	selbuild.flds, selbuild.scan = &isq.fields, selbuild.Scan
	return selbuild
}

func (isq *IntSIDQuery) prepareQuery(ctx context.Context) error {
	for _, f := range isq.fields {
		if !intsid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if isq.path != nil {
		prev, err := isq.path(ctx)
		if err != nil {
			return err
		}
		isq.sql = prev
	}
	return nil
}

func (isq *IntSIDQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IntSID, error) {
	var (
		nodes       = []*IntSID{}
		withFKs     = isq.withFKs
		_spec       = isq.querySpec()
		loadedTypes = [2]bool{
			isq.withParent != nil,
			isq.withChildren != nil,
		}
	)
	if isq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, intsid.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*IntSID).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &IntSID{config: isq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, isq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := isq.withParent; query != nil {
		ids := make([]sid.ID, 0, len(nodes))
		nodeids := make(map[sid.ID][]*IntSID)
		for i := range nodes {
			if nodes[i].int_sid_parent == nil {
				continue
			}
			fk := *nodes[i].int_sid_parent
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(intsid.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "int_sid_parent" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := isq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[sid.ID]*IntSID)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*IntSID{}
		}
		query.withFKs = true
		query.Where(predicate.IntSID(func(s *sql.Selector) {
			s.Where(sql.InValues(intsid.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.int_sid_parent
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "int_sid_parent" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "int_sid_parent" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	return nodes, nil
}

func (isq *IntSIDQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := isq.querySpec()
	_spec.Node.Columns = isq.fields
	if len(isq.fields) > 0 {
		_spec.Unique = isq.unique != nil && *isq.unique
	}
	return sqlgraph.CountNodes(ctx, isq.driver, _spec)
}

func (isq *IntSIDQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := isq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (isq *IntSIDQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   intsid.Table,
			Columns: intsid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: intsid.FieldID,
			},
		},
		From:   isq.sql,
		Unique: true,
	}
	if unique := isq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := isq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, intsid.FieldID)
		for i := range fields {
			if fields[i] != intsid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := isq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := isq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := isq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := isq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (isq *IntSIDQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(isq.driver.Dialect())
	t1 := builder.Table(intsid.Table)
	columns := isq.fields
	if len(columns) == 0 {
		columns = intsid.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if isq.sql != nil {
		selector = isq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if isq.unique != nil && *isq.unique {
		selector.Distinct()
	}
	for _, p := range isq.predicates {
		p(selector)
	}
	for _, p := range isq.order {
		p(selector)
	}
	if offset := isq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := isq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IntSIDGroupBy is the group-by builder for IntSID entities.
type IntSIDGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (isgb *IntSIDGroupBy) Aggregate(fns ...AggregateFunc) *IntSIDGroupBy {
	isgb.fns = append(isgb.fns, fns...)
	return isgb
}

// Scan applies the group-by query and scans the result into the given value.
func (isgb *IntSIDGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := isgb.path(ctx)
	if err != nil {
		return err
	}
	isgb.sql = query
	return isgb.sqlScan(ctx, v)
}

func (isgb *IntSIDGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range isgb.fields {
		if !intsid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := isgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := isgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (isgb *IntSIDGroupBy) sqlQuery() *sql.Selector {
	selector := isgb.sql.Select()
	aggregation := make([]string, 0, len(isgb.fns))
	for _, fn := range isgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(isgb.fields)+len(isgb.fns))
		for _, f := range isgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(isgb.fields...)...)
}

// IntSIDSelect is the builder for selecting fields of IntSID entities.
type IntSIDSelect struct {
	*IntSIDQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (iss *IntSIDSelect) Scan(ctx context.Context, v interface{}) error {
	if err := iss.prepareQuery(ctx); err != nil {
		return err
	}
	iss.sql = iss.IntSIDQuery.sqlQuery(ctx)
	return iss.sqlScan(ctx, v)
}

func (iss *IntSIDSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := iss.sql.Query()
	if err := iss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
