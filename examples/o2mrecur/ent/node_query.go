// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"
	"github.com/facebookincubator/ent/examples/o2mrecur/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
)

// NodeQuery is the builder for querying Node entities.
type NodeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Node
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (nq *NodeQuery) Where(ps ...predicate.Node) *NodeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NodeQuery) Limit(limit int) *NodeQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NodeQuery) Offset(offset int) *NodeQuery {
	nq.offset = &offset
	return nq
}

// Order adds an order step to the query.
func (nq *NodeQuery) Order(o ...Order) *NodeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryParent chains the current query on the parent edge.
func (nq *NodeQuery) QueryParent() *NodeQuery {
	query := &NodeQuery{config: nq.config}
	t1 := sql.Table(node.Table)
	t2 := nq.sqlQuery()
	t2.Select(t2.C(node.ParentColumn))
	query.sql = sql.Select(t1.Columns(node.Columns...)...).
		From(t1).
		Join(t2).
		On(t1.C(node.FieldID), t2.C(node.ParentColumn))
	return query
}

// QueryChildren chains the current query on the children edge.
func (nq *NodeQuery) QueryChildren() *NodeQuery {
	query := &NodeQuery{config: nq.config}
	t1 := sql.Table(node.Table)
	t2 := nq.sqlQuery()
	t2.Select(t2.C(node.FieldID))
	query.sql = sql.Select().
		From(t1).
		Join(t2).
		On(t1.C(node.ChildrenColumn), t2.C(node.FieldID))
	return query
}

// Get returns a Node entity by its id.
func (nq *NodeQuery) Get(ctx context.Context, id int) (*Node, error) {
	return nq.Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (nq *NodeQuery) GetX(ctx context.Context, id int) *Node {
	n, err := nq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return n
}

// First returns the first Node entity in the query. Returns *ErrNotFound when no node was found.
func (nq *NodeQuery) First(ctx context.Context) (*Node, error) {
	ns, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ns) == 0 {
		return nil, &ErrNotFound{node.Label}
	}
	return ns[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NodeQuery) FirstX(ctx context.Context) *Node {
	n, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return n
}

// FirstID returns the first Node id in the query. Returns *ErrNotFound when no id was found.
func (nq *NodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{node.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (nq *NodeQuery) FirstXID(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Node entity in the query, returns an error if not exactly one entity was returned.
func (nq *NodeQuery) Only(ctx context.Context) (*Node, error) {
	ns, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ns) {
	case 1:
		return ns[0], nil
	case 0:
		return nil, &ErrNotFound{node.Label}
	default:
		return nil, &ErrNotSingular{node.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NodeQuery) OnlyX(ctx context.Context) *Node {
	n, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// OnlyID returns the only Node id in the query, returns an error if not exactly one id was returned.
func (nq *NodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{node.Label}
	default:
		err = &ErrNotSingular{node.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (nq *NodeQuery) OnlyXID(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nodes.
func (nq *NodeQuery) All(ctx context.Context) ([]*Node, error) {
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NodeQuery) AllX(ctx context.Context) []*Node {
	ns, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ns
}

// IDs executes the query and returns a list of Node ids.
func (nq *NodeQuery) IDs(ctx context.Context) ([]int, error) {
	return nq.sqlIDs(ctx)
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NodeQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NodeQuery) Count(ctx context.Context) (int, error) {
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NodeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NodeQuery) Exist(ctx context.Context) (bool, error) {
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NodeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NodeQuery) Clone() *NodeQuery {
	return &NodeQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]Order{}, nq.order...),
		unique:     append([]string{}, nq.unique...),
		predicates: append([]predicate.Node{}, nq.predicates...),
		// clone intermediate queries.
		sql: nq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Node.Query().
//		GroupBy(node.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nq *NodeQuery) GroupBy(field string, fields ...string) *NodeGroupBy {
	group := &NodeGroupBy{config: nq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = nq.sqlQuery()
	return group
}

func (nq *NodeQuery) sqlAll(ctx context.Context) ([]*Node, error) {
	rows := &sql.Rows{}
	selector := nq.sqlQuery()
	if unique := nq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := nq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ns Nodes
	if err := ns.FromRows(rows); err != nil {
		return nil, err
	}
	ns.config(nq.config)
	return ns, nil
}

func (nq *NodeQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := nq.sqlQuery()
	unique := []string{node.FieldID}
	if len(nq.unique) > 0 {
		unique = nq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := nq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (nq *NodeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (nq *NodeQuery) sqlIDs(ctx context.Context) ([]int, error) {
	vs, err := nq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (nq *NodeQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(node.Table)
	selector := sql.Select(t1.Columns(node.Columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(node.Columns...)...)
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NodeQuery is the builder for group-by Node entities.
type NodeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NodeGroupBy) Aggregate(fns ...Aggregate) *NodeGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scan the result into the given value.
func (ngb *NodeGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ngb *NodeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ngb *NodeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NodeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ngb *NodeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ngb *NodeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NodeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ngb *NodeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ngb *NodeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NodeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ngb *NodeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ngb *NodeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NodeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ngb *NodeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ngb *NodeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ngb.sqlQuery().Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NodeGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql
	columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
	columns = append(columns, ngb.fields...)
	for _, fn := range ngb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(ngb.fields...)
}
