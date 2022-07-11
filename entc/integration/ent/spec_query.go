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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/spec"
	"entgo.io/ent/schema/field"
)

// SpecQuery is the builder for querying Spec entities.
type SpecQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Spec
	// eager-loading edges.
	withCard  *CardQuery
	modifiers []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SpecQuery builder.
func (sq *SpecQuery) Where(ps ...predicate.Spec) *SpecQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// When runs the provided builder(s) if and only if condition is true.
func (sq *SpecQuery) When(condition bool, action func(builder *SpecQuery)) *SpecQuery {
	if condition {
		action(sq)
	}

	return sq
}

// Limit adds a limit step to the query.
func (sq *SpecQuery) Limit(limit int) *SpecQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SpecQuery) Offset(offset int) *SpecQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SpecQuery) Unique(unique bool) *SpecQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SpecQuery) Order(o ...OrderFunc) *SpecQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCard chains the current query on the "card" edge.
func (sq *SpecQuery) QueryCard() *CardQuery {
	query := &CardQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(spec.Table, spec.FieldID, selector),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, spec.CardTable, spec.CardPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Spec entity from the query.
// Returns a *NotFoundError when no Spec was found.
func (sq *SpecQuery) First(ctx context.Context) (*Spec, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{spec.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SpecQuery) FirstX(ctx context.Context) *Spec {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Spec ID from the query.
// Returns a *NotFoundError when no Spec ID was found.
func (sq *SpecQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{spec.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SpecQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Spec entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Spec entity is found.
// Returns a *NotFoundError when no Spec entities are found.
func (sq *SpecQuery) Only(ctx context.Context) (*Spec, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{spec.Label}
	default:
		return nil, &NotSingularError{spec.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SpecQuery) OnlyX(ctx context.Context) *Spec {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Spec ID in the query.
// Returns a *NotSingularError when more than one Spec ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SpecQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{spec.Label}
	default:
		err = &NotSingularError{spec.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SpecQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Specs.
func (sq *SpecQuery) All(ctx context.Context) ([]*Spec, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SpecQuery) AllX(ctx context.Context) []*Spec {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Spec IDs.
func (sq *SpecQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(spec.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SpecQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SpecQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SpecQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SpecQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SpecQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpecQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SpecQuery) Clone() *SpecQuery {
	if sq == nil {
		return nil
	}
	return &SpecQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		predicates: append([]predicate.Spec{}, sq.predicates...),
		withCard:   sq.withCard.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithCard tells the query-builder to eager-load the nodes that are connected to
// the "card" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SpecQuery) WithCard(opts ...func(*CardQuery)) *SpecQuery {
	query := &CardQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withCard = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (sq *SpecQuery) GroupBy(field string, fields ...string) *SpecGroupBy {
	grbuild := &SpecGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = spec.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (sq *SpecQuery) Select(fields ...string) *SpecSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SpecSelect{SpecQuery: sq}
	selbuild.label = spec.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

func (sq *SpecQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !spec.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SpecQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Spec, error) {
	var (
		nodes       = []*Spec{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withCard != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Spec).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Spec{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withCard; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*Spec)
		nids := make(map[int]map[*Spec]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Card = []*Card{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(spec.CardTable)
			s.Join(joinT).On(s.C(card.FieldID), joinT.C(spec.CardPrimaryKey[1]))
			s.Where(sql.InValues(joinT.C(spec.CardPrimaryKey[0]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(spec.CardPrimaryKey[0]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Spec]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "card" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Card = append(kn.Edges.Card, n)
			}
		}
	}

	return nodes, nil
}

func (sq *SpecQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SpecQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sq *SpecQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   spec.Table,
			Columns: spec.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: spec.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, spec.FieldID)
		for i := range fields {
			if fields[i] != spec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SpecQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(spec.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = spec.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sq *SpecQuery) ForUpdate(opts ...sql.LockOption) *SpecQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sq *SpecQuery) ForShare(opts ...sql.LockOption) *SpecQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sq *SpecQuery) Modify(modifiers ...func(s *sql.Selector)) *SpecSelect {
	sq.modifiers = append(sq.modifiers, modifiers...)
	return sq.Select()
}

// SpecGroupBy is the group-by builder for Spec entities.
type SpecGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SpecGroupBy) Aggregate(fns ...AggregateFunc) *SpecGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SpecGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SpecGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgb.fields {
		if !spec.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SpecGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SpecSelect is the builder for selecting fields of Spec entities.
type SpecSelect struct {
	*SpecQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SpecSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SpecQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SpecSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ss *SpecSelect) Modify(modifiers ...func(s *sql.Selector)) *SpecSelect {
	ss.modifiers = append(ss.modifiers, modifiers...)
	return ss
}
