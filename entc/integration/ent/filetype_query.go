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
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// FileTypeQuery is the builder for querying FileType entities.
type FileTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FileType
	// eager-loading edges.
	withFiles *FileQuery
	modifiers []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileTypeQuery builder.
func (ftq *FileTypeQuery) Where(ps ...predicate.FileType) *FileTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FileTypeQuery) Limit(limit int) *FileTypeQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FileTypeQuery) Offset(offset int) *FileTypeQuery {
	ftq.offset = &offset
	return ftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftq *FileTypeQuery) Unique(unique bool) *FileTypeQuery {
	ftq.unique = &unique
	return ftq
}

// Order adds an order step to the query.
func (ftq *FileTypeQuery) Order(o ...OrderFunc) *FileTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// QueryFiles chains the current query on the "files" edge.
func (ftq *FileTypeQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: ftq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ftq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filetype.Table, filetype.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, filetype.FilesTable, filetype.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ftq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FileType entity from the query.
// Returns a *NotFoundError when no FileType was found.
func (ftq *FileTypeQuery) First(ctx context.Context) (*FileType, error) {
	nodes, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstX(ctx context.Context) *FileType {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileType ID from the query.
// Returns a *NotFoundError when no FileType ID was found.
func (ftq *FileTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FileType entity is found.
// Returns a *NotFoundError when no FileType entities are found.
func (ftq *FileTypeQuery) Only(ctx context.Context) (*FileType, error) {
	nodes, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filetype.Label}
	default:
		return nil, &NotSingularError{filetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyX(ctx context.Context) *FileType {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileType ID in the query.
// Returns a *NotSingularError when more than one FileType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FileTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filetype.Label}
	default:
		err = &NotSingularError{filetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileTypes.
func (ftq *FileTypeQuery) All(ctx context.Context) ([]*FileType, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ftq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FileTypeQuery) AllX(ctx context.Context) []*FileType {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileType IDs.
func (ftq *FileTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ftq.Select(filetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FileTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FileTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ftq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FileTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FileTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ftq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FileTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FileTypeQuery) Clone() *FileTypeQuery {
	if ftq == nil {
		return nil
	}
	return &FileTypeQuery{
		config:     ftq.config,
		limit:      ftq.limit,
		offset:     ftq.offset,
		order:      append([]OrderFunc{}, ftq.order...),
		predicates: append([]predicate.FileType{}, ftq.predicates...),
		withFiles:  ftq.withFiles.Clone(),
		// clone intermediate query.
		sql:    ftq.sql.Clone(),
		path:   ftq.path,
		unique: ftq.unique,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (ftq *FileTypeQuery) WithFiles(opts ...func(*FileQuery)) *FileTypeQuery {
	query := &FileQuery{config: ftq.config}
	for _, opt := range opts {
		opt(query)
	}
	ftq.withFiles = query
	return ftq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileType.Query().
//		GroupBy(filetype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FileTypeQuery) GroupBy(field string, fields ...string) *FileTypeGroupBy {
	grbuild := &FileTypeGroupBy{config: ftq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ftq.sqlQuery(ctx), nil
	}
	grbuild.label = filetype.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.FileType.Query().
//		Select(filetype.FieldName).
//		Scan(ctx, &v)
//
func (ftq *FileTypeQuery) Select(fields ...string) *FileTypeSelect {
	ftq.fields = append(ftq.fields, fields...)
	selbuild := &FileTypeSelect{FileTypeQuery: ftq}
	selbuild.label = filetype.Label
	selbuild.flds, selbuild.scan = &ftq.fields, selbuild.Scan
	return selbuild
}

func (ftq *FileTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ftq.fields {
		if !filetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.sql = prev
	}
	return nil
}

func (ftq *FileTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FileType, error) {
	var (
		nodes       = []*FileType{}
		_spec       = ftq.querySpec()
		loadedTypes = [1]bool{
			ftq.withFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FileType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FileType{config: ftq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ftq.modifiers) > 0 {
		_spec.Modifiers = ftq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ftq.withFiles; query != nil {
		if err := ftq.loadFiles(ctx, query, nodes,
			func(n *FileType) { n.Edges.Files = []*File{} },
			func(n *FileType, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ftq *FileTypeQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*FileType, init func(*FileType), assign func(*FileType, *File)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*FileType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.File(func(s *sql.Selector) {
		s.Where(sql.InValues(filetype.FilesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.file_type_files
		if fk == nil {
			return fmt.Errorf(`foreign-key "file_type_files" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "file_type_files" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ftq *FileTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftq.querySpec()
	if len(ftq.modifiers) > 0 {
		_spec.Modifiers = ftq.modifiers
	}
	_spec.Node.Columns = ftq.fields
	if len(ftq.fields) > 0 {
		_spec.Unique = ftq.unique != nil && *ftq.unique
	}
	return sqlgraph.CountNodes(ctx, ftq.driver, _spec)
}

func (ftq *FileTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ftq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ftq *FileTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filetype.Table,
			Columns: filetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filetype.FieldID,
			},
		},
		From:   ftq.sql,
		Unique: true,
	}
	if unique := ftq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ftq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetype.FieldID)
		for i := range fields {
			if fields[i] != filetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ftq *FileTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ftq.driver.Dialect())
	t1 := builder.Table(filetype.Table)
	columns := ftq.fields
	if len(columns) == 0 {
		columns = filetype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ftq.unique != nil && *ftq.unique {
		selector.Distinct()
	}
	for _, m := range ftq.modifiers {
		m(selector)
	}
	for _, p := range ftq.predicates {
		p(selector)
	}
	for _, p := range ftq.order {
		p(selector)
	}
	if offset := ftq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (ftq *FileTypeQuery) ForUpdate(opts ...sql.LockOption) *FileTypeQuery {
	if ftq.driver.Dialect() == dialect.Postgres {
		ftq.Unique(false)
	}
	ftq.modifiers = append(ftq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return ftq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (ftq *FileTypeQuery) ForShare(opts ...sql.LockOption) *FileTypeQuery {
	if ftq.driver.Dialect() == dialect.Postgres {
		ftq.Unique(false)
	}
	ftq.modifiers = append(ftq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return ftq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ftq *FileTypeQuery) Modify(modifiers ...func(s *sql.Selector)) *FileTypeSelect {
	ftq.modifiers = append(ftq.modifiers, modifiers...)
	return ftq.Select()
}

// FileTypeGroupBy is the group-by builder for FileType entities.
type FileTypeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FileTypeGroupBy) Aggregate(fns ...AggregateFunc) *FileTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ftgb *FileTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ftgb.path(ctx)
	if err != nil {
		return err
	}
	ftgb.sql = query
	return ftgb.sqlScan(ctx, v)
}

func (ftgb *FileTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ftgb.fields {
		if !filetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ftgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ftgb *FileTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ftgb.sql.Select()
	aggregation := make([]string, 0, len(ftgb.fns))
	for _, fn := range ftgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ftgb.fields)+len(ftgb.fns))
		for _, f := range ftgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ftgb.fields...)...)
}

// FileTypeSelect is the builder for selecting fields of FileType entities.
type FileTypeSelect struct {
	*FileTypeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FileTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	fts.sql = fts.FileTypeQuery.sqlQuery(ctx)
	return fts.sqlScan(ctx, v)
}

func (fts *FileTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fts.sql.Query()
	if err := fts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fts *FileTypeSelect) Modify(modifiers ...func(s *sql.Selector)) *FileTypeSelect {
	fts.modifiers = append(fts.modifiers, modifiers...)
	return fts
}
