// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"fbc/ent/entc/integration/ent/group"
	"fbc/ent/entc/integration/ent/groupinfo"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// GroupInfoQuery is the builder for querying GroupInfo entities.
type GroupInfoQuery struct {
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
func (giq *GroupInfoQuery) Where(ps ...ent.Predicate) *GroupInfoQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit adds a limit step to the query.
func (giq *GroupInfoQuery) Limit(limit int) *GroupInfoQuery {
	giq.limit = &limit
	return giq
}

// Offset adds an offset step to the query.
func (giq *GroupInfoQuery) Offset(offset int) *GroupInfoQuery {
	giq.offset = &offset
	return giq
}

// Order adds an order step to the query.
func (giq *GroupInfoQuery) Order(o ...Order) *GroupInfoQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// QueryGroups chains the current query on the groups edge.
func (giq *GroupInfoQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: giq.config}
	switch giq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		t1 := sql.Table(group.Table)
		t2 := giq.sqlQuery()
		t2.Select(t2.C(groupinfo.FieldID))
		query.sql = sql.Select().
			From(t1).
			Join(t2).
			On(t1.C(groupinfo.GroupsColumn), t2.C(groupinfo.FieldID))
	case dialect.Neptune:
		gremlin := giq.gremlinQuery()
		query.gremlin = gremlin.InE(group.InfoLabel).OutV()
	}
	return query
}

// Get returns a GroupInfo entity by its id.
func (giq *GroupInfoQuery) Get(ctx context.Context, id string) (*GroupInfo, error) {
	return giq.Where(groupinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (giq *GroupInfoQuery) GetX(ctx context.Context, id string) *GroupInfo {
	gi, err := giq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gi
}

// First returns the first GroupInfo entity in the query. Returns *ErrNotFound when no groupinfo was found.
func (giq *GroupInfoQuery) First(ctx context.Context) (*GroupInfo, error) {
	gis, err := giq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gis) == 0 {
		return nil, &ErrNotFound{groupinfo.Label}
	}
	return gis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstX(ctx context.Context) *GroupInfo {
	gi, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gi
}

// FirstID returns the first GroupInfo id in the query. Returns *ErrNotFound when no id was found.
func (giq *GroupInfoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = giq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{groupinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstXID(ctx context.Context) string {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GroupInfo entity in the query, returns an error if not exactly one entity was returned.
func (giq *GroupInfoQuery) Only(ctx context.Context) (*GroupInfo, error) {
	gis, err := giq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gis) {
	case 1:
		return gis[0], nil
	case 0:
		return nil, &ErrNotFound{groupinfo.Label}
	default:
		return nil, &ErrNotSingular{groupinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyX(ctx context.Context) *GroupInfo {
	gi, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gi
}

// OnlyID returns the only GroupInfo id in the query, returns an error if not exactly one id was returned.
func (giq *GroupInfoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = giq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{groupinfo.Label}
	default:
		err = &ErrNotSingular{groupinfo.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyXID(ctx context.Context) string {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupInfos.
func (giq *GroupInfoQuery) All(ctx context.Context) ([]*GroupInfo, error) {
	switch giq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return giq.sqlAll(ctx)
	case dialect.Neptune:
		return giq.gremlinAll(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (giq *GroupInfoQuery) AllX(ctx context.Context) []*GroupInfo {
	gis, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gis
}

// IDs executes the query and returns a list of GroupInfo ids.
func (giq *GroupInfoQuery) IDs(ctx context.Context) ([]string, error) {
	switch giq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return giq.sqlIDs(ctx)
	case dialect.Neptune:
		return giq.gremlinIDs(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GroupInfoQuery) IDsX(ctx context.Context) []string {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GroupInfoQuery) Count(ctx context.Context) (int, error) {
	switch giq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return giq.sqlCount(ctx)
	case dialect.Neptune:
		return giq.gremlinCount(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// CountX is like Count, but panics if an error occurs.
func (giq *GroupInfoQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GroupInfoQuery) Exist(ctx context.Context) (bool, error) {
	switch giq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return giq.sqlExist(ctx)
	case dialect.Neptune:
		return giq.gremlinExist(ctx)
	default:
		return false, errors.New("ent: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GroupInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GroupInfoQuery) Clone() *GroupInfoQuery {
	return &GroupInfoQuery{
		config:     giq.config,
		limit:      giq.limit,
		offset:     giq.offset,
		order:      append([]Order{}, giq.order...),
		unique:     append([]string{}, giq.unique...),
		predicates: append([]ent.Predicate{}, giq.predicates...),
		// clone intermediate queries.
		sql:     giq.sql.Clone(),
		gremlin: giq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		GroupBy(groupinfo.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) GroupBy(field string, fields ...string) *GroupInfoGroupBy {
	group := &GroupInfoGroupBy{config: giq.config}
	group.fields = append([]string{field}, fields...)
	switch giq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = giq.sqlQuery()
	case dialect.Neptune:
		group.gremlin = giq.gremlinQuery()
	}
	return group
}

func (giq *GroupInfoQuery) sqlAll(ctx context.Context) ([]*GroupInfo, error) {
	rows := &sql.Rows{}
	selector := giq.sqlQuery()
	if unique := giq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := giq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var gis GroupInfos
	if err := gis.FromRows(rows); err != nil {
		return nil, err
	}
	gis.config(giq.config)
	return gis, nil
}

func (giq *GroupInfoQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := giq.sqlQuery()
	unique := []string{groupinfo.FieldID}
	if len(giq.unique) > 0 {
		unique = giq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := giq.driver.Query(ctx, query, args, rows); err != nil {
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

func (giq *GroupInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := giq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (giq *GroupInfoQuery) sqlIDs(ctx context.Context) ([]string, error) {
	vs, err := giq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (giq *GroupInfoQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(groupinfo.Table)
	selector := sql.Select(t1.Columns(groupinfo.Columns...)...).From(t1)
	if giq.sql != nil {
		selector = giq.sql
		selector.Select(selector.Columns(groupinfo.Columns...)...)
	}
	for _, p := range giq.predicates {
		p.SQL(selector)
	}
	for _, p := range giq.order {
		p.SQL(selector)
	}
	if offset := giq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := giq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (giq *GroupInfoQuery) gremlinIDs(ctx context.Context) ([]string, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
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

func (giq *GroupInfoQuery) gremlinAll(ctx context.Context) ([]*GroupInfo, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().ValueMap(true).Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var gis GroupInfos
	if err := gis.FromResponse(res); err != nil {
		return nil, err
	}
	gis.config(giq.config)
	return gis, nil
}

func (giq *GroupInfoQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().Count().Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (giq *GroupInfoQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().HasNext().Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (giq *GroupInfoQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(groupinfo.Label)
	if giq.gremlin != nil {
		v = giq.gremlin.Clone()
	}
	for _, p := range giq.predicates {
		p.Gremlin(v)
	}
	if len(giq.order) > 0 {
		v.Order()
		for _, p := range giq.order {
			p.Gremlin(v)
		}
	}
	switch limit, offset := giq.limit, giq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt64)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := giq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// GroupInfoQuery is the builder for group-by GroupInfo entities.
type GroupInfoGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GroupInfoGroupBy) Aggregate(fns ...Aggregate) *GroupInfoGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the group-by query and scan the result into the given value.
func (gigb *GroupInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch gigb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gigb.sqlScan(ctx, v)
	case dialect.Neptune:
		return gigb.gremlinScan(ctx, v)
	default:
		return errors.New("gigb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := gigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := gigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gigb *GroupInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gigb.sqlQuery().Query()
	if err := gigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gigb *GroupInfoGroupBy) sqlQuery() *sql.Selector {
	selector := gigb.sql
	columns := make([]string, 0, len(gigb.fields)+len(gigb.fns))
	columns = append(columns, gigb.fields...)
	for _, fn := range gigb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(gigb.fields...)
}

func (gigb *GroupInfoGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := gigb.gremlinQuery().Query()
	if err := gigb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(gigb.fields)+len(gigb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (gigb *GroupInfoGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range gigb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range gigb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return gigb.gremlin.Group().
		By(__.Values(gigb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}
