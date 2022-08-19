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
	"entgo.io/ent/entc/integration/ent/fieldtype"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
)

// FileQuery is the builder for querying File entities.
type FileQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.File
	withOwner      *UserQuery
	withType       *FileTypeQuery
	withField      *FieldTypeQuery
	withFKs        bool
	modifiers      []func(*sql.Selector)
	withNamedField map[string]*FieldTypeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileQuery builder.
func (fq *FileQuery) Where(ps ...predicate.File) *FileQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FileQuery) Limit(limit int) *FileQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FileQuery) Offset(offset int) *FileQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FileQuery) Unique(unique bool) *FileQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FileQuery) Order(o ...OrderFunc) *FileQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryOwner chains the current query on the "owner" edge.
func (fq *FileQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, file.OwnerTable, file.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryType chains the current query on the "type" edge.
func (fq *FileQuery) QueryType() *FileTypeQuery {
	query := &FileTypeQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, selector),
			sqlgraph.To(filetype.Table, filetype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, file.TypeTable, file.TypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryField chains the current query on the "field" edge.
func (fq *FileQuery) QueryField() *FieldTypeQuery {
	query := &FieldTypeQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, selector),
			sqlgraph.To(fieldtype.Table, fieldtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, file.FieldTable, file.FieldColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first File entity from the query.
// Returns a *NotFoundError when no File was found.
func (fq *FileQuery) First(ctx context.Context) (*File, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{file.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FileQuery) FirstX(ctx context.Context) *File {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first File ID from the query.
// Returns a *NotFoundError when no File ID was found.
func (fq *FileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{file.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FileQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single File entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one File entity is found.
// Returns a *NotFoundError when no File entities are found.
func (fq *FileQuery) Only(ctx context.Context) (*File, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{file.Label}
	default:
		return nil, &NotSingularError{file.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FileQuery) OnlyX(ctx context.Context) *File {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only File ID in the query.
// Returns a *NotSingularError when more than one File ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{file.Label}
	default:
		err = &NotSingularError{file.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FileQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Files.
func (fq *FileQuery) All(ctx context.Context) ([]*File, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FileQuery) AllX(ctx context.Context) []*File {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of File IDs.
func (fq *FileQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(file.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FileQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FileQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FileQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FileQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FileQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FileQuery) Clone() *FileQuery {
	if fq == nil {
		return nil
	}
	return &FileQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		predicates: append([]predicate.File{}, fq.predicates...),
		withOwner:  fq.withOwner.Clone(),
		withType:   fq.withType.Clone(),
		withField:  fq.withField.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithOwner(opts ...func(*UserQuery)) *FileQuery {
	query := &UserQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withOwner = query
	return fq
}

// WithType tells the query-builder to eager-load the nodes that are connected to
// the "type" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithType(opts ...func(*FileTypeQuery)) *FileQuery {
	query := &FileTypeQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withType = query
	return fq
}

// WithField tells the query-builder to eager-load the nodes that are connected to
// the "field" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithField(opts ...func(*FieldTypeQuery)) *FileQuery {
	query := &FieldTypeQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withField = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Size int `json:"size,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.File.Query().
//		GroupBy(file.FieldSize).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FileQuery) GroupBy(field string, fields ...string) *FileGroupBy {
	grbuild := &FileGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = file.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Size int `json:"size,omitempty"`
//	}
//
//	client.File.Query().
//		Select(file.FieldSize).
//		Scan(ctx, &v)
func (fq *FileQuery) Select(fields ...string) *FileSelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FileSelect{FileQuery: fq}
	selbuild.label = file.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

func (fq *FileQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !file.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*File, error) {
	var (
		nodes       = []*File{}
		withFKs     = fq.withFKs
		_spec       = fq.querySpec()
		loadedTypes = [3]bool{
			fq.withOwner != nil,
			fq.withType != nil,
			fq.withField != nil,
		}
	)
	if fq.withOwner != nil || fq.withType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, file.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*File).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &File{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withOwner; query != nil {
		if err := fq.loadOwner(ctx, query, nodes, nil,
			func(n *File, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withType; query != nil {
		if err := fq.loadType(ctx, query, nodes, nil,
			func(n *File, e *FileType) { n.Edges.Type = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withField; query != nil {
		if err := fq.loadField(ctx, query, nodes,
			func(n *File) { n.Edges.Field = []*FieldType{} },
			func(n *File, e *FieldType) { n.Edges.Field = append(n.Edges.Field, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fq.withNamedField {
		if err := fq.loadField(ctx, query, nodes,
			func(n *File) { n.appendNamedField(name) },
			func(n *File, e *FieldType) { n.appendNamedField(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FileQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*File, init func(*File), assign func(*File, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*File)
	for i := range nodes {
		if nodes[i].user_files == nil {
			continue
		}
		fk := *nodes[i].user_files
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
			return fmt.Errorf(`unexpected foreign-key "user_files" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FileQuery) loadType(ctx context.Context, query *FileTypeQuery, nodes []*File, init func(*File), assign func(*File, *FileType)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*File)
	for i := range nodes {
		if nodes[i].file_type_files == nil {
			continue
		}
		fk := *nodes[i].file_type_files
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(filetype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "file_type_files" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FileQuery) loadField(ctx context.Context, query *FieldTypeQuery, nodes []*File, init func(*File), assign func(*File, *FieldType)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*File)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.FieldType(func(s *sql.Selector) {
		s.Where(sql.InValues(file.FieldColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.file_field
		if fk == nil {
			return fmt.Errorf(`foreign-key "file_field" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "file_field" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *FileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fq *FileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: file.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for i := range fields {
			if fields[i] != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(file.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = file.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, m := range fq.modifiers {
		m(selector)
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fq *FileQuery) ForUpdate(opts ...sql.LockOption) *FileQuery {
	if fq.driver.Dialect() == dialect.Postgres {
		fq.Unique(false)
	}
	fq.modifiers = append(fq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fq *FileQuery) ForShare(opts ...sql.LockOption) *FileQuery {
	if fq.driver.Dialect() == dialect.Postgres {
		fq.Unique(false)
	}
	fq.modifiers = append(fq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fq *FileQuery) Modify(modifiers ...func(s *sql.Selector)) *FileSelect {
	fq.modifiers = append(fq.modifiers, modifiers...)
	return fq.Select()
}

// WithNamedField tells the query-builder to eager-load the nodes that are connected to the "field"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fq *FileQuery) WithNamedField(name string, opts ...func(*FieldTypeQuery)) *FileQuery {
	query := &FieldTypeQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	if fq.withNamedField == nil {
		fq.withNamedField = make(map[string]*FieldTypeQuery)
	}
	fq.withNamedField[name] = query
	return fq
}

// FileGroupBy is the group-by builder for File entities.
type FileGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FileGroupBy) Aggregate(fns ...AggregateFunc) *FileGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FileGroupBy) Scan(ctx context.Context, v any) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FileGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range fgb.fields {
		if !file.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FileGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FileSelect is the builder for selecting fields of File entities.
type FileSelect struct {
	*FileQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FileSelect) Scan(ctx context.Context, v any) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FileQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FileSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fs *FileSelect) Modify(modifiers ...func(s *sql.Selector)) *FileSelect {
	fs.modifiers = append(fs.modifiers, modifiers...)
	return fs
}
