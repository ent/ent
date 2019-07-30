// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"fbc/ent/entc/integration/plugin/ent/boring"
	"fbc/ent/entc/integration/plugin/ent/predicate"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// BoringQuery is the builder for querying Boring entities.
type BoringQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Boring
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (bq *BoringQuery) Where(ps ...predicate.Boring) *BoringQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BoringQuery) Limit(limit int) *BoringQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BoringQuery) Offset(offset int) *BoringQuery {
	bq.offset = &offset
	return bq
}

// Order adds an order step to the query.
func (bq *BoringQuery) Order(o ...Order) *BoringQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// Get returns a Boring entity by its id.
func (bq *BoringQuery) Get(ctx context.Context, id string) (*Boring, error) {
	return bq.Where(boring.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (bq *BoringQuery) GetX(ctx context.Context, id string) *Boring {
	b, err := bq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// First returns the first Boring entity in the query. Returns *ErrNotFound when no boring was found.
func (bq *BoringQuery) First(ctx context.Context) (*Boring, error) {
	bs, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, &ErrNotFound{boring.Label}
	}
	return bs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BoringQuery) FirstX(ctx context.Context) *Boring {
	b, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return b
}

// FirstID returns the first Boring id in the query. Returns *ErrNotFound when no id was found.
func (bq *BoringQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{boring.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bq *BoringQuery) FirstXID(ctx context.Context) string {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Boring entity in the query, returns an error if not exactly one entity was returned.
func (bq *BoringQuery) Only(ctx context.Context) (*Boring, error) {
	bs, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(bs) {
	case 1:
		return bs[0], nil
	case 0:
		return nil, &ErrNotFound{boring.Label}
	default:
		return nil, &ErrNotSingular{boring.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BoringQuery) OnlyX(ctx context.Context) *Boring {
	b, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// OnlyID returns the only Boring id in the query, returns an error if not exactly one id was returned.
func (bq *BoringQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{boring.Label}
	default:
		err = &ErrNotSingular{boring.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (bq *BoringQuery) OnlyXID(ctx context.Context) string {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Borings.
func (bq *BoringQuery) All(ctx context.Context) ([]*Boring, error) {
	switch bq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bq.sqlAll(ctx)
	case dialect.Neptune:
		return bq.gremlinAll(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (bq *BoringQuery) AllX(ctx context.Context) []*Boring {
	bs, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return bs
}

// IDs executes the query and returns a list of Boring ids.
func (bq *BoringQuery) IDs(ctx context.Context) ([]string, error) {
	switch bq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bq.sqlIDs(ctx)
	case dialect.Neptune:
		return bq.gremlinIDs(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BoringQuery) IDsX(ctx context.Context) []string {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BoringQuery) Count(ctx context.Context) (int, error) {
	switch bq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bq.sqlCount(ctx)
	case dialect.Neptune:
		return bq.gremlinCount(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// CountX is like Count, but panics if an error occurs.
func (bq *BoringQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BoringQuery) Exist(ctx context.Context) (bool, error) {
	switch bq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bq.sqlExist(ctx)
	case dialect.Neptune:
		return bq.gremlinExist(ctx)
	default:
		return false, errors.New("ent: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BoringQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BoringQuery) Clone() *BoringQuery {
	return &BoringQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]Order{}, bq.order...),
		unique:     append([]string{}, bq.unique...),
		predicates: append([]predicate.Boring{}, bq.predicates...),
		// clone intermediate queries.
		sql:     bq.sql.Clone(),
		gremlin: bq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (bq *BoringQuery) GroupBy(field string, fields ...string) *BoringGroupBy {
	group := &BoringGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	switch bq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = bq.sqlQuery()
	case dialect.Neptune:
		group.gremlin = bq.gremlinQuery()
	}
	return group
}

func (bq *BoringQuery) sqlAll(ctx context.Context) ([]*Boring, error) {
	rows := &sql.Rows{}
	selector := bq.sqlQuery()
	if unique := bq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := bq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var bs Borings
	if err := bs.FromRows(rows); err != nil {
		return nil, err
	}
	bs.config(bq.config)
	return bs, nil
}

func (bq *BoringQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := bq.sqlQuery()
	unique := []string{boring.FieldID}
	if len(bq.unique) > 0 {
		unique = bq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := bq.driver.Query(ctx, query, args, rows); err != nil {
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

func (bq *BoringQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bq *BoringQuery) sqlIDs(ctx context.Context) ([]string, error) {
	vs, err := bq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (bq *BoringQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(boring.Table)
	selector := sql.Select(t1.Columns(boring.Columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(boring.Columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (bq *BoringQuery) gremlinIDs(ctx context.Context) ([]string, error) {
	res := &gremlin.Response{}
	query, bindings := bq.gremlinQuery().Query()
	if err := bq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	vertices, err := res.ReadVertices()
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(vertices))
	for _, vertex := range vertices {
		ids = append(ids, vertex.ID.(string))
	}
	return ids, nil
}

func (bq *BoringQuery) gremlinAll(ctx context.Context) ([]*Boring, error) {
	res := &gremlin.Response{}
	query, bindings := bq.gremlinQuery().ValueMap(true).Query()
	if err := bq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var bs Borings
	if err := bs.FromResponse(res); err != nil {
		return nil, err
	}
	bs.config(bq.config)
	return bs, nil
}

func (bq *BoringQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := bq.gremlinQuery().Count().Query()
	if err := bq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (bq *BoringQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := bq.gremlinQuery().HasNext().Query()
	if err := bq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (bq *BoringQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(boring.Label)
	if bq.gremlin != nil {
		v = bq.gremlin.Clone()
	}
	for _, p := range bq.predicates {
		p(v)
	}
	if len(bq.order) > 0 {
		v.Order()
		for _, p := range bq.order {
			p(v)
		}
	}
	switch limit, offset := bq.limit, bq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt64)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := bq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// BoringQuery is the builder for group-by Boring entities.
type BoringGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BoringGroupBy) Aggregate(fns ...Aggregate) *BoringGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bgb *BoringGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch bgb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bgb.sqlScan(ctx, v)
	case dialect.Neptune:
		return bgb.gremlinScan(ctx, v)
	default:
		return errors.New("bgb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BoringGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bgb *BoringGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BoringGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BoringGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bgb *BoringGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BoringGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BoringGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bgb *BoringGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BoringGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BoringGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bgb *BoringGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BoringGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BoringGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BoringGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bgb.sqlQuery().Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BoringGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql
	columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
	columns = append(columns, bgb.fields...)
	for _, fn := range bgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(bgb.fields...)
}

func (bgb *BoringGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := bgb.gremlinQuery().Query()
	if err := bgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(bgb.fields)+len(bgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (bgb *BoringGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range bgb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range bgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return bgb.gremlin.Group().
		By(__.Values(bgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}
