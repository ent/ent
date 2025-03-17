// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/rls/ent/predicate"
	"entgo.io/ent/examples/rls/ent/tenant"
	"entgo.io/ent/schema/field"
)

// TenantQuery is the builder for querying Tenant entities.
type TenantQuery struct {
	config
	ctx        *QueryContext
	order      []tenant.OrderOption
	inters     []Interceptor
	predicates []predicate.Tenant
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TenantQuery builder.
func (q *TenantQuery) Where(ps ...predicate.Tenant) *TenantQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *TenantQuery) Limit(limit int) *TenantQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *TenantQuery) Offset(offset int) *TenantQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *TenantQuery) Unique(unique bool) *TenantQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *TenantQuery) Order(o ...tenant.OrderOption) *TenantQuery {
	q.order = append(q.order, o...)
	return q
}

// First returns the first Tenant entity from the query.
// Returns a *NotFoundError when no Tenant was found.
func (q *TenantQuery) First(ctx context.Context) (*Tenant, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tenant.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *TenantQuery) FirstX(ctx context.Context) *Tenant {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tenant ID from the query.
// Returns a *NotFoundError when no Tenant ID was found.
func (q *TenantQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tenant.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *TenantQuery) FirstIDX(ctx context.Context) int {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tenant entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tenant entity is found.
// Returns a *NotFoundError when no Tenant entities are found.
func (q *TenantQuery) Only(ctx context.Context) (*Tenant, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tenant.Label}
	default:
		return nil, &NotSingularError{tenant.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *TenantQuery) OnlyX(ctx context.Context) *Tenant {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tenant ID in the query.
// Returns a *NotSingularError when more than one Tenant ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *TenantQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tenant.Label}
	default:
		err = &NotSingularError{tenant.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *TenantQuery) OnlyIDX(ctx context.Context) int {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tenants.
func (q *TenantQuery) All(ctx context.Context) ([]*Tenant, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tenant, *TenantQuery]()
	return withInterceptors[[]*Tenant](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *TenantQuery) AllX(ctx context.Context) []*Tenant {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tenant IDs.
func (q *TenantQuery) IDs(ctx context.Context) (ids []int, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(tenant.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *TenantQuery) IDsX(ctx context.Context) []int {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *TenantQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*TenantQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *TenantQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *TenantQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryExist)
	switch _, err := q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (q *TenantQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TenantQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *TenantQuery) Clone() *TenantQuery {
	if q == nil {
		return nil
	}
	return &TenantQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]tenant.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.Tenant{}, q.predicates...),
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
//	client.Tenant.Query().
//		GroupBy(tenant.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *TenantQuery) GroupBy(field string, fields ...string) *TenantGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TenantGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = tenant.Label
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
//	client.Tenant.Query().
//		Select(tenant.FieldName).
//		Scan(ctx, &v)
func (q *TenantQuery) Select(fields ...string) *TenantSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &TenantSelect{TenantQuery: q}
	sbuild.label = tenant.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TenantSelect configured with the given aggregations.
func (q *TenantQuery) Aggregate(fns ...AggregateFunc) *TenantSelect {
	return q.Select().Aggregate(fns...)
}

func (q *TenantQuery) prepareQuery(ctx context.Context) error {
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
		if !tenant.ValidColumn(f) {
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

func (q *TenantQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tenant, error) {
	var (
		nodes = []*Tenant{}
		_spec = q.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tenant).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tenant{config: q.config}
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

func (q *TenantQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *TenantQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tenant.Table, tenant.Columns, sqlgraph.NewFieldSpec(tenant.FieldID, field.TypeInt))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tenant.FieldID)
		for i := range fields {
			if fields[i] != tenant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
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

func (q *TenantQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(tenant.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = tenant.Columns
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

// TenantGroupBy is the group-by builder for Tenant entities.
type TenantGroupBy struct {
	selector
	build *TenantQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TenantGroupBy) Aggregate(fns ...AggregateFunc) *TenantGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TenantGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TenantQuery, *TenantGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (q *TenantGroupBy) sqlScan(ctx context.Context, root *TenantQuery, v any) error {
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

// TenantSelect is the builder for selecting fields of Tenant entities.
type TenantSelect struct {
	*TenantQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TenantSelect) Aggregate(fns ...AggregateFunc) *TenantSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TenantSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TenantQuery, *TenantSelect](ctx, ts.TenantQuery, ts, ts.inters, v)
}

func (q *TenantSelect) sqlScan(ctx context.Context, root *TenantQuery, v any) error {
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
