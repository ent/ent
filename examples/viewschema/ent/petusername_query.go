// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/viewschema/ent/petusername"
	"entgo.io/ent/examples/viewschema/ent/predicate"
)

// PetUserNameQuery is the builder for querying PetUserName entities.
type PetUserNameQuery struct {
	config
	ctx        *QueryContext
	order      []petusername.OrderOption
	inters     []Interceptor
	predicates []predicate.PetUserName
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PetUserNameQuery builder.
func (q *PetUserNameQuery) Where(ps ...predicate.PetUserName) *PetUserNameQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *PetUserNameQuery) Limit(limit int) *PetUserNameQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *PetUserNameQuery) Offset(offset int) *PetUserNameQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *PetUserNameQuery) Unique(unique bool) *PetUserNameQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *PetUserNameQuery) Order(o ...petusername.OrderOption) *PetUserNameQuery {
	q.order = append(q.order, o...)
	return q
}

// First returns the first PetUserName entity from the query.
// Returns a *NotFoundError when no PetUserName was found.
func (q *PetUserNameQuery) First(ctx context.Context) (*PetUserName, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{petusername.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *PetUserNameQuery) FirstX(ctx context.Context) *PetUserName {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single PetUserName entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PetUserName entity is found.
// Returns a *NotFoundError when no PetUserName entities are found.
func (q *PetUserNameQuery) Only(ctx context.Context) (*PetUserName, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{petusername.Label}
	default:
		return nil, &NotSingularError{petusername.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *PetUserNameQuery) OnlyX(ctx context.Context) *PetUserName {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of PetUserNames.
func (q *PetUserNameQuery) All(ctx context.Context) ([]*PetUserName, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PetUserName, *PetUserNameQuery]()
	return withInterceptors[[]*PetUserName](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *PetUserNameQuery) AllX(ctx context.Context) []*PetUserName {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (q *PetUserNameQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*PetUserNameQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *PetUserNameQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *PetUserNameQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryExist)
	switch _, err := q.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (q *PetUserNameQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PetUserNameQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *PetUserNameQuery) Clone() *PetUserNameQuery {
	if q == nil {
		return nil
	}
	return &PetUserNameQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]petusername.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.PetUserName{}, q.predicates...),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
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
//	client.PetUserName.Query().
//		GroupBy(petusername.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *PetUserNameQuery) GroupBy(field string, fields ...string) *PetUserNameGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PetUserNameGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = petusername.Label
	grbuild.scan = grbuild.Scan
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
//	client.PetUserName.Query().
//		Select(petusername.FieldName).
//		Scan(ctx, &v)
func (q *PetUserNameQuery) Select(fields ...string) *PetUserNameSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &PetUserNameSelect{PetUserNameQuery: q}
	sbuild.label = petusername.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PetUserNameSelect configured with the given aggregations.
func (q *PetUserNameQuery) Aggregate(fns ...AggregateFunc) *PetUserNameSelect {
	return q.Select().Aggregate(fns...)
}

func (q *PetUserNameQuery) prepareQuery(ctx context.Context) error {
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
		if !petusername.ValidColumn(f) {
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

func (q *PetUserNameQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PetUserName, error) {
	var (
		nodes = []*PetUserName{}
		_spec = q.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PetUserName).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PetUserName{config: q.config}
		nodes = append(nodes, node)
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
	return nodes, nil
}

func (q *PetUserNameQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *PetUserNameQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(petusername.Table, petusername.Columns, nil)
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
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

func (q *PetUserNameQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(petusername.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = petusername.Columns
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

// PetUserNameGroupBy is the group-by builder for PetUserName entities.
type PetUserNameGroupBy struct {
	selector
	build *PetUserNameQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pungb *PetUserNameGroupBy) Aggregate(fns ...AggregateFunc) *PetUserNameGroupBy {
	pungb.fns = append(pungb.fns, fns...)
	return pungb
}

// Scan applies the selector query and scans the result into the given value.
func (pungb *PetUserNameGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pungb.build.ctx, ent.OpQueryGroupBy)
	if err := pungb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetUserNameQuery, *PetUserNameGroupBy](ctx, pungb.build, pungb, pungb.build.inters, v)
}

func (q *PetUserNameGroupBy) sqlScan(ctx context.Context, root *PetUserNameQuery, v any) error {
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

// PetUserNameSelect is the builder for selecting fields of PetUserName entities.
type PetUserNameSelect struct {
	*PetUserNameQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (puns *PetUserNameSelect) Aggregate(fns ...AggregateFunc) *PetUserNameSelect {
	puns.fns = append(puns.fns, fns...)
	return puns
}

// Scan applies the selector query and scans the result into the given value.
func (puns *PetUserNameSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, puns.ctx, ent.OpQuerySelect)
	if err := puns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PetUserNameQuery, *PetUserNameSelect](ctx, puns.PetUserNameQuery, puns, puns.inters, v)
}

func (q *PetUserNameSelect) sqlScan(ctx context.Context, root *PetUserNameQuery, v any) error {
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
