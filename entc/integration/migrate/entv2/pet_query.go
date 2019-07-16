// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"
	"math"

	"fbc/ent/entc/integration/migrate/entv2/pet"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/__"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// PetQuery is the builder for querying Pet entities.
type PetQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []ent.Predicate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (pq *PetQuery) Where(ps ...ent.Predicate) *PetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PetQuery) Limit(limit int) *PetQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PetQuery) Offset(offset int) *PetQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PetQuery) Order(o ...Order) *PetQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// Get returns a Pet entity by its id.
func (pq *PetQuery) Get(ctx context.Context, id string) (*Pet, error) {
	return pq.Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (pq *PetQuery) GetX(ctx context.Context, id string) *Pet {
	pe, err := pq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pe
}

// First returns the first Pet entity in the query. Returns *ErrNotFound when no pet was found.
func (pq *PetQuery) First(ctx context.Context) (*Pet, error) {
	pes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pes) == 0 {
		return nil, &ErrNotFound{pet.Label}
	}
	return pes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PetQuery) FirstX(ctx context.Context) *Pet {
	pe, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pe
}

// FirstID returns the first Pet id in the query. Returns *ErrNotFound when no id was found.
func (pq *PetQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{pet.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PetQuery) FirstXID(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Pet entity in the query, returns an error if not exactly one entity was returned.
func (pq *PetQuery) Only(ctx context.Context) (*Pet, error) {
	pes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pes) {
	case 1:
		return pes[0], nil
	case 0:
		return nil, &ErrNotFound{pet.Label}
	default:
		return nil, &ErrNotSingular{pet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PetQuery) OnlyX(ctx context.Context) *Pet {
	pe, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// OnlyID returns the only Pet id in the query, returns an error if not exactly one id was returned.
func (pq *PetQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{pet.Label}
	default:
		err = &ErrNotSingular{pet.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (pq *PetQuery) OnlyXID(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pets.
func (pq *PetQuery) All(ctx context.Context) ([]*Pet, error) {
	switch pq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pq.sqlAll(ctx)
	case dialect.Neptune:
		return pq.gremlinAll(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (pq *PetQuery) AllX(ctx context.Context) []*Pet {
	pes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pes
}

// IDs executes the query and returns a list of Pet ids.
func (pq *PetQuery) IDs(ctx context.Context) ([]string, error) {
	switch pq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pq.sqlIDs(ctx)
	case dialect.Neptune:
		return pq.gremlinIDs(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PetQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PetQuery) Count(ctx context.Context) (int, error) {
	switch pq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pq.sqlCount(ctx)
	case dialect.Neptune:
		return pq.gremlinCount(ctx)
	default:
		return 0, errors.New("entv2: unsupported dialect")
	}
}

// CountX is like Count, but panics if an error occurs.
func (pq *PetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PetQuery) Exist(ctx context.Context) (bool, error) {
	switch pq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pq.sqlExist(ctx)
	case dialect.Neptune:
		return pq.gremlinExist(ctx)
	default:
		return false, errors.New("entv2: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PetQuery) Clone() *PetQuery {
	return &PetQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]Order{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]ent.Predicate{}, pq.predicates...),
		// clone intermediate queries.
		sql:     pq.sql.Clone(),
		gremlin: pq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (pq *PetQuery) GroupBy(field string, fields ...string) *PetGroupBy {
	group := &PetGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	switch pq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = pq.sqlQuery()
	case dialect.Neptune:
		group.gremlin = pq.gremlinQuery()
	}
	return group
}

func (pq *PetQuery) sqlAll(ctx context.Context) ([]*Pet, error) {
	rows := &sql.Rows{}
	selector := pq.sqlQuery()
	if unique := pq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := pq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var pes Pets
	if err := pes.FromRows(rows); err != nil {
		return nil, err
	}
	pes.config(pq.config)
	return pes, nil
}

func (pq *PetQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := pq.sqlQuery()
	unique := []string{pet.FieldID}
	if len(pq.unique) > 0 {
		unique = pq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := pq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("entv2: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("entv2: failed reading count: %v", err)
	}
	return n, nil
}

func (pq *PetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entv2: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PetQuery) sqlIDs(ctx context.Context) ([]string, error) {
	vs, err := pq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (pq *PetQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(pet.Table)
	selector := sql.Select(t1.Columns(pet.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(pet.Columns...)...)
	}
	for _, p := range pq.predicates {
		p.SQL(selector)
	}
	for _, p := range pq.order {
		p.SQL(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (pq *PetQuery) gremlinIDs(ctx context.Context) ([]string, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
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

func (pq *PetQuery) gremlinAll(ctx context.Context) ([]*Pet, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().ValueMap(true).Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var pes Pets
	if err := pes.FromResponse(res); err != nil {
		return nil, err
	}
	pes.config(pq.config)
	return pes, nil
}

func (pq *PetQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().Count().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pq *PetQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := pq.gremlinQuery().HasNext().Query()
	if err := pq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (pq *PetQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(pet.Label)
	if pq.gremlin != nil {
		v = pq.gremlin.Clone()
	}
	for _, p := range pq.predicates {
		p.Gremlin(v)
	}
	if len(pq.order) > 0 {
		v.Order()
		for _, p := range pq.order {
			p.Gremlin(v)
		}
	}
	switch limit, offset := pq.limit, pq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt64)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := pq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// PetQuery is the builder for group-by Pet entities.
type PetGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PetGroupBy) Aggregate(fns ...Aggregate) *PetGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PetGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch pgb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pgb.sqlScan(ctx, v)
	case dialect.Neptune:
		return pgb.gremlinScan(ctx, v)
	default:
		return errors.New("pgb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PetGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("entv2: PetGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PetGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("entv2: PetGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PetGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("entv2: PetGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PetGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PetGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("entv2: PetGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PetGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PetGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

func (pgb *PetGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := pgb.gremlinQuery().Query()
	if err := pgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(pgb.fields)+len(pgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (pgb *PetGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range pgb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range pgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return pgb.gremlin.Group().
		By(__.Values(pgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}
