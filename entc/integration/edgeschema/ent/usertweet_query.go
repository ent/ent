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
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// UserTweetQuery is the builder for querying UserTweet entities.
type UserTweetQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserTweet
	// eager-loading edges.
	withUser  *UserQuery
	withTweet *TweetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserTweetQuery builder.
func (utq *UserTweetQuery) Where(ps ...predicate.UserTweet) *UserTweetQuery {
	utq.predicates = append(utq.predicates, ps...)
	return utq
}

// Limit adds a limit step to the query.
func (utq *UserTweetQuery) Limit(limit int) *UserTweetQuery {
	utq.limit = &limit
	return utq
}

// Offset adds an offset step to the query.
func (utq *UserTweetQuery) Offset(offset int) *UserTweetQuery {
	utq.offset = &offset
	return utq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utq *UserTweetQuery) Unique(unique bool) *UserTweetQuery {
	utq.unique = &unique
	return utq
}

// Order adds an order step to the query.
func (utq *UserTweetQuery) Order(o ...OrderFunc) *UserTweetQuery {
	utq.order = append(utq.order, o...)
	return utq
}

// QueryUser chains the current query on the "user" edge.
func (utq *UserTweetQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: utq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertweet.Table, usertweet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usertweet.UserTable, usertweet.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTweet chains the current query on the "tweet" edge.
func (utq *UserTweetQuery) QueryTweet() *TweetQuery {
	query := &TweetQuery{config: utq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertweet.Table, usertweet.FieldID, selector),
			sqlgraph.To(tweet.Table, tweet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usertweet.TweetTable, usertweet.TweetColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserTweet entity from the query.
// Returns a *NotFoundError when no UserTweet was found.
func (utq *UserTweetQuery) First(ctx context.Context) (*UserTweet, error) {
	nodes, err := utq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usertweet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utq *UserTweetQuery) FirstX(ctx context.Context) *UserTweet {
	node, err := utq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserTweet ID from the query.
// Returns a *NotFoundError when no UserTweet ID was found.
func (utq *UserTweetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usertweet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (utq *UserTweetQuery) FirstIDX(ctx context.Context) int {
	id, err := utq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserTweet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserTweet entity is found.
// Returns a *NotFoundError when no UserTweet entities are found.
func (utq *UserTweetQuery) Only(ctx context.Context) (*UserTweet, error) {
	nodes, err := utq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usertweet.Label}
	default:
		return nil, &NotSingularError{usertweet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utq *UserTweetQuery) OnlyX(ctx context.Context) *UserTweet {
	node, err := utq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserTweet ID in the query.
// Returns a *NotSingularError when more than one UserTweet ID is found.
// Returns a *NotFoundError when no entities are found.
func (utq *UserTweetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usertweet.Label}
	default:
		err = &NotSingularError{usertweet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utq *UserTweetQuery) OnlyIDX(ctx context.Context) int {
	id, err := utq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserTweets.
func (utq *UserTweetQuery) All(ctx context.Context) ([]*UserTweet, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return utq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (utq *UserTweetQuery) AllX(ctx context.Context) []*UserTweet {
	nodes, err := utq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserTweet IDs.
func (utq *UserTweetQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := utq.Select(usertweet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utq *UserTweetQuery) IDsX(ctx context.Context) []int {
	ids, err := utq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utq *UserTweetQuery) Count(ctx context.Context) (int, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return utq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (utq *UserTweetQuery) CountX(ctx context.Context) int {
	count, err := utq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utq *UserTweetQuery) Exist(ctx context.Context) (bool, error) {
	if err := utq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return utq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (utq *UserTweetQuery) ExistX(ctx context.Context) bool {
	exist, err := utq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserTweetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utq *UserTweetQuery) Clone() *UserTweetQuery {
	if utq == nil {
		return nil
	}
	return &UserTweetQuery{
		config:     utq.config,
		limit:      utq.limit,
		offset:     utq.offset,
		order:      append([]OrderFunc{}, utq.order...),
		predicates: append([]predicate.UserTweet{}, utq.predicates...),
		withUser:   utq.withUser.Clone(),
		withTweet:  utq.withTweet.Clone(),
		// clone intermediate query.
		sql:    utq.sql.Clone(),
		path:   utq.path,
		unique: utq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTweetQuery) WithUser(opts ...func(*UserQuery)) *UserTweetQuery {
	query := &UserQuery{config: utq.config}
	for _, opt := range opts {
		opt(query)
	}
	utq.withUser = query
	return utq
}

// WithTweet tells the query-builder to eager-load the nodes that are connected to
// the "tweet" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTweetQuery) WithTweet(opts ...func(*TweetQuery)) *UserTweetQuery {
	query := &TweetQuery{config: utq.config}
	for _, opt := range opts {
		opt(query)
	}
	utq.withTweet = query
	return utq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserTweet.Query().
//		GroupBy(usertweet.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (utq *UserTweetQuery) GroupBy(field string, fields ...string) *UserTweetGroupBy {
	grbuild := &UserTweetGroupBy{config: utq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return utq.sqlQuery(ctx), nil
	}
	grbuild.label = usertweet.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserTweet.Query().
//		Select(usertweet.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (utq *UserTweetQuery) Select(fields ...string) *UserTweetSelect {
	utq.fields = append(utq.fields, fields...)
	selbuild := &UserTweetSelect{UserTweetQuery: utq}
	selbuild.label = usertweet.Label
	selbuild.flds, selbuild.scan = &utq.fields, selbuild.Scan
	return selbuild
}

func (utq *UserTweetQuery) prepareQuery(ctx context.Context) error {
	for _, f := range utq.fields {
		if !usertweet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if utq.path != nil {
		prev, err := utq.path(ctx)
		if err != nil {
			return err
		}
		utq.sql = prev
	}
	return nil
}

func (utq *UserTweetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserTweet, error) {
	var (
		nodes       = []*UserTweet{}
		_spec       = utq.querySpec()
		loadedTypes = [2]bool{
			utq.withUser != nil,
			utq.withTweet != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserTweet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserTweet{config: utq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, utq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := utq.withUser; query != nil {
		if err := utq.loadUser(ctx, query, nodes, nil,
			func(n *UserTweet, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := utq.withTweet; query != nil {
		if err := utq.loadTweet(ctx, query, nodes, nil,
			func(n *UserTweet, e *Tweet) { n.Edges.Tweet = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (utq *UserTweetQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserTweet, init func(*UserTweet), assign func(*UserTweet, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserTweet)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (utq *UserTweetQuery) loadTweet(ctx context.Context, query *TweetQuery, nodes []*UserTweet, init func(*UserTweet), assign func(*UserTweet, *Tweet)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserTweet)
	for i := range nodes {
		fk := nodes[i].TweetID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tweet.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tweet_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (utq *UserTweetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utq.querySpec()
	_spec.Node.Columns = utq.fields
	if len(utq.fields) > 0 {
		_spec.Unique = utq.unique != nil && *utq.unique
	}
	return sqlgraph.CountNodes(ctx, utq.driver, _spec)
}

func (utq *UserTweetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := utq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (utq *UserTweetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertweet.Table,
			Columns: usertweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertweet.FieldID,
			},
		},
		From:   utq.sql,
		Unique: true,
	}
	if unique := utq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := utq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertweet.FieldID)
		for i := range fields {
			if fields[i] != usertweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := utq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utq *UserTweetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utq.driver.Dialect())
	t1 := builder.Table(usertweet.Table)
	columns := utq.fields
	if len(columns) == 0 {
		columns = usertweet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utq.sql != nil {
		selector = utq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if utq.unique != nil && *utq.unique {
		selector.Distinct()
	}
	for _, p := range utq.predicates {
		p(selector)
	}
	for _, p := range utq.order {
		p(selector)
	}
	if offset := utq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserTweetGroupBy is the group-by builder for UserTweet entities.
type UserTweetGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utgb *UserTweetGroupBy) Aggregate(fns ...AggregateFunc) *UserTweetGroupBy {
	utgb.fns = append(utgb.fns, fns...)
	return utgb
}

// Scan applies the group-by query and scans the result into the given value.
func (utgb *UserTweetGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := utgb.path(ctx)
	if err != nil {
		return err
	}
	utgb.sql = query
	return utgb.sqlScan(ctx, v)
}

func (utgb *UserTweetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range utgb.fields {
		if !usertweet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := utgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (utgb *UserTweetGroupBy) sqlQuery() *sql.Selector {
	selector := utgb.sql.Select()
	aggregation := make([]string, 0, len(utgb.fns))
	for _, fn := range utgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(utgb.fields)+len(utgb.fns))
		for _, f := range utgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(utgb.fields...)...)
}

// UserTweetSelect is the builder for selecting fields of UserTweet entities.
type UserTweetSelect struct {
	*UserTweetQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uts *UserTweetSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uts.prepareQuery(ctx); err != nil {
		return err
	}
	uts.sql = uts.UserTweetQuery.sqlQuery(ctx)
	return uts.sqlScan(ctx, v)
}

func (uts *UserTweetSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uts.sql.Query()
	if err := uts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
