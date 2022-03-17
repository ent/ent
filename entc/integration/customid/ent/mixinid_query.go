// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/mixinid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MixinIDQuery is the builder for querying MixinID entities.
type MixinIDQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MixinID
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MixinIDQuery builder.
func (miq *MixinIDQuery) Where(ps ...predicate.MixinID) *MixinIDQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit adds a limit step to the query.
func (miq *MixinIDQuery) Limit(limit int) *MixinIDQuery {
	miq.limit = &limit
	return miq
}

// Offset adds an offset step to the query.
func (miq *MixinIDQuery) Offset(offset int) *MixinIDQuery {
	miq.offset = &offset
	return miq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (miq *MixinIDQuery) Unique(unique bool) *MixinIDQuery {
	miq.unique = &unique
	return miq
}

// Order adds an order step to the query.
func (miq *MixinIDQuery) Order(o ...OrderFunc) *MixinIDQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// First returns the first MixinID entity from the query.
// Returns a *NotFoundError when no MixinID was found.
func (miq *MixinIDQuery) First(ctx context.Context) (*MixinID, error) {
	nodes, err := miq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mixinid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MixinIDQuery) FirstX(ctx context.Context) *MixinID {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MixinID ID from the query.
// Returns a *NotFoundError when no MixinID ID was found.
func (miq *MixinIDQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = miq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mixinid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MixinIDQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MixinID entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MixinID entity is found.
// Returns a *NotFoundError when no MixinID entities are found.
func (miq *MixinIDQuery) Only(ctx context.Context) (*MixinID, error) {
	nodes, err := miq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mixinid.Label}
	default:
		return nil, &NotSingularError{mixinid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MixinIDQuery) OnlyX(ctx context.Context) *MixinID {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MixinID ID in the query.
// Returns a *NotSingularError when more than one MixinID ID is found.
// Returns a *NotFoundError when no entities are found.
func (miq *MixinIDQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = miq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = &NotSingularError{mixinid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MixinIDQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MixinIDs.
func (miq *MixinIDQuery) All(ctx context.Context) ([]*MixinID, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return miq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (miq *MixinIDQuery) AllX(ctx context.Context) []*MixinID {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MixinID IDs.
func (miq *MixinIDQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := miq.Select(mixinid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MixinIDQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MixinIDQuery) Count(ctx context.Context) (int, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return miq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MixinIDQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MixinIDQuery) Exist(ctx context.Context) (bool, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return miq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MixinIDQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MixinIDQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MixinIDQuery) Clone() *MixinIDQuery {
	if miq == nil {
		return nil
	}
	return &MixinIDQuery{
		config:     miq.config,
		limit:      miq.limit,
		offset:     miq.offset,
		order:      append([]OrderFunc{}, miq.order...),
		predicates: append([]predicate.MixinID{}, miq.predicates...),
		// clone intermediate query.
		sql:    miq.sql.Clone(),
		path:   miq.path,
		unique: miq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SomeField string `json:"some_field,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MixinID.Query().
//		GroupBy(mixinid.FieldSomeField).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (miq *MixinIDQuery) GroupBy(field string, fields ...string) *MixinIDGroupBy {
	grbuild := &MixinIDGroupBy{config: miq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return miq.sqlQuery(ctx), nil
	}
	grbuild.label = mixinid.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SomeField string `json:"some_field,omitempty"`
//	}
//
//	client.MixinID.Query().
//		Select(mixinid.FieldSomeField).
//		Scan(ctx, &v)
//
func (miq *MixinIDQuery) Select(fields ...string) *MixinIDSelect {
	miq.fields = append(miq.fields, fields...)
	selbuild := &MixinIDSelect{MixinIDQuery: miq}
	selbuild.label = mixinid.Label
	selbuild.flds, selbuild.scan = &miq.fields, selbuild.Scan
	return selbuild
}

func (miq *MixinIDQuery) prepareQuery(ctx context.Context) error {
	for _, f := range miq.fields {
		if !mixinid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MixinIDQuery) sqlAll(ctx context.Context) ([]*MixinID, error) {
	var (
		nodes = []*MixinID{}
		_spec = miq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MixinID{config: miq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (miq *MixinIDQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	_spec.Node.Columns = miq.fields
	if len(miq.fields) > 0 {
		_spec.Unique = miq.unique != nil && *miq.unique
	}
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MixinIDQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := miq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (miq *MixinIDQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mixinid.Table,
			Columns: mixinid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		},
		From:   miq.sql,
		Unique: true,
	}
	if unique := miq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := miq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mixinid.FieldID)
		for i := range fields {
			if fields[i] != mixinid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (miq *MixinIDQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(mixinid.Table)
	columns := miq.fields
	if len(columns) == 0 {
		columns = mixinid.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if miq.unique != nil && *miq.unique {
		selector.Distinct()
	}
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector)
	}
	if offset := miq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MixinIDGroupBy is the group-by builder for MixinID entities.
type MixinIDGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MixinIDGroupBy) Aggregate(fns ...AggregateFunc) *MixinIDGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the group-by query and scans the result into the given value.
func (migb *MixinIDGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := migb.path(ctx)
	if err != nil {
		return err
	}
	migb.sql = query
	return migb.sqlScan(ctx, v)
}

func (migb *MixinIDGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range migb.fields {
		if !mixinid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := migb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (migb *MixinIDGroupBy) sqlQuery() *sql.Selector {
	selector := migb.sql.Select()
	aggregation := make([]string, 0, len(migb.fns))
	for _, fn := range migb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(migb.fields)+len(migb.fns))
		for _, f := range migb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(migb.fields...)...)
}

// MixinIDSelect is the builder for selecting fields of MixinID entities.
type MixinIDSelect struct {
	*MixinIDQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mis *MixinIDSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mis.prepareQuery(ctx); err != nil {
		return err
	}
	mis.sql = mis.MixinIDQuery.sqlQuery(ctx)
	return mis.sqlScan(ctx, v)
}

func (mis *MixinIDSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mis.sql.Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
