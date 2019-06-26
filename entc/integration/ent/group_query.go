// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"fbc/ent/entc/integration/ent/file"
	"fbc/ent/entc/integration/ent/group"
	"fbc/ent/entc/integration/ent/groupinfo"
	"fbc/ent/entc/integration/ent/user"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/__"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
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
func (gq *GroupQuery) Where(ps ...ent.Predicate) *GroupQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GroupQuery) Limit(limit int) *GroupQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GroupQuery) Offset(offset int) *GroupQuery {
	gq.offset = &offset
	return gq
}

// Order adds an order step to the query.
func (gq *GroupQuery) Order(o ...Order) *GroupQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryFiles chains the current query on the files edge.
func (gq *GroupQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: gq.config}
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		t1 := sql.Table(file.Table)
		t2 := gq.sqlQuery()
		t2.Select(t2.C(group.FieldID))
		query.sql = sql.Select().
			From(t1).
			Join(t2).
			On(t1.C(group.FilesColumn), t2.C(group.FieldID))
	case dialect.Neptune:
		gremlin := gq.gremlinQuery()
		query.gremlin = gremlin.OutE(group.FilesLabel).InV()
	}
	return query
}

// QueryBlocked chains the current query on the blocked edge.
func (gq *GroupQuery) QueryBlocked() *UserQuery {
	query := &UserQuery{config: gq.config}
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		t1 := sql.Table(user.Table)
		t2 := gq.sqlQuery()
		t2.Select(t2.C(group.FieldID))
		query.sql = sql.Select().
			From(t1).
			Join(t2).
			On(t1.C(group.BlockedColumn), t2.C(group.FieldID))
	case dialect.Neptune:
		gremlin := gq.gremlinQuery()
		query.gremlin = gremlin.OutE(group.BlockedLabel).InV()
	}
	return query
}

// QueryUsers chains the current query on the users edge.
func (gq *GroupQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: gq.config}
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		t1 := sql.Table(user.Table)
		t2 := gq.sqlQuery()
		t2.Select(t2.C(group.FieldID))
		t3 := sql.Table(group.UsersTable)
		t4 := sql.Select(t3.C(group.UsersPrimaryKey[0])).
			From(t3).
			Join(t2).
			On(t3.C(group.UsersPrimaryKey[1]), t2.C(group.FieldID))
		query.sql = sql.Select().
			From(t1).
			Join(t4).
			On(t1.C(user.FieldID), t4.C(group.UsersPrimaryKey[0]))
	case dialect.Neptune:
		gremlin := gq.gremlinQuery()
		query.gremlin = gremlin.InE(user.GroupsLabel).OutV()
	}
	return query
}

// QueryInfo chains the current query on the info edge.
func (gq *GroupQuery) QueryInfo() *GroupInfoQuery {
	query := &GroupInfoQuery{config: gq.config}
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		t1 := sql.Table(groupinfo.Table)
		t2 := gq.sqlQuery()
		t2.Select(t2.C(group.InfoColumn))
		query.sql = sql.Select(t1.Columns(groupinfo.Columns...)...).
			From(t1).
			Join(t2).
			On(t1.C(groupinfo.FieldID), t2.C(group.InfoColumn))
	case dialect.Neptune:
		gremlin := gq.gremlinQuery()
		query.gremlin = gremlin.OutE(group.InfoLabel).InV()
	}
	return query
}

// Get returns a Group entity by its id.
func (gq *GroupQuery) Get(ctx context.Context, id string) (*Group, error) {
	return gq.Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (gq *GroupQuery) GetX(ctx context.Context, id string) *Group {
	gr, err := gq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gr
}

// First returns the first Group entity in the query. Returns *ErrNotFound when no group was found.
func (gq *GroupQuery) First(ctx context.Context) (*Group, error) {
	grs, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(grs) == 0 {
		return nil, &ErrNotFound{group.Label}
	}
	return grs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GroupQuery) FirstX(ctx context.Context) *Group {
	gr, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gr
}

// FirstID returns the first Group id in the query. Returns *ErrNotFound when no id was found.
func (gq *GroupQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{group.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (gq *GroupQuery) FirstXID(ctx context.Context) string {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Group entity in the query, returns an error if not exactly one entity was returned.
func (gq *GroupQuery) Only(ctx context.Context) (*Group, error) {
	grs, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(grs) {
	case 1:
		return grs[0], nil
	case 0:
		return nil, &ErrNotFound{group.Label}
	default:
		return nil, &ErrNotSingular{group.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GroupQuery) OnlyX(ctx context.Context) *Group {
	gr, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// OnlyID returns the only Group id in the query, returns an error if not exactly one id was returned.
func (gq *GroupQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{group.Label}
	default:
		err = &ErrNotSingular{group.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (gq *GroupQuery) OnlyXID(ctx context.Context) string {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (gq *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gq.sqlAll(ctx)
	case dialect.Neptune:
		return gq.gremlinAll(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (gq *GroupQuery) AllX(ctx context.Context) []*Group {
	grs, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return grs
}

// IDs executes the query and returns a list of Group ids.
func (gq *GroupQuery) IDs(ctx context.Context) ([]string, error) {
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gq.sqlIDs(ctx)
	case dialect.Neptune:
		return gq.gremlinIDs(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GroupQuery) IDsX(ctx context.Context) []string {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GroupQuery) Count(ctx context.Context) (int, error) {
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gq.sqlCount(ctx)
	case dialect.Neptune:
		return gq.gremlinCount(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// CountX is like Count, but panics if an error occurs.
func (gq *GroupQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GroupQuery) Exist(ctx context.Context) (bool, error) {
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gq.sqlExist(ctx)
	case dialect.Neptune:
		return gq.gremlinExist(ctx)
	default:
		return false, errors.New("ent: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Active bool `json:"active,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Group.Query().
//		GroupBy(group.FieldActive).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GroupQuery) GroupBy(field string, fields ...string) *GroupGroupBy {
	group := &GroupGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	switch gq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = gq.sqlQuery()
	case dialect.Neptune:
		group.gremlin = gq.gremlinQuery()
	}
	return group
}

func (gq *GroupQuery) sqlAll(ctx context.Context) ([]*Group, error) {
	rows := &sql.Rows{}
	selector := gq.sqlQuery()
	if unique := gq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := gq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var grs Groups
	if err := grs.FromRows(rows); err != nil {
		return nil, err
	}
	grs.config(gq.config)
	return grs, nil
}

func (gq *GroupQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := gq.sqlQuery()
	unique := []string{group.FieldID}
	if len(gq.unique) > 0 {
		unique = gq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := gq.driver.Query(ctx, query, args, rows); err != nil {
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

func (gq *GroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (gq *GroupQuery) sqlIDs(ctx context.Context) ([]string, error) {
	vs, err := gq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (gq *GroupQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(group.Table)
	selector := sql.Select(t1.Columns(group.Columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(group.Columns...)...)
	}
	for _, p := range gq.predicates {
		p.SQL(selector)
	}
	for _, p := range gq.order {
		p.SQL(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (gq *GroupQuery) gremlinIDs(ctx context.Context) ([]string, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
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

func (gq *GroupQuery) gremlinAll(ctx context.Context) ([]*Group, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().ValueMap(true).Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var grs Groups
	if err := grs.FromResponse(res); err != nil {
		return nil, err
	}
	grs.config(gq.config)
	return grs, nil
}

func (gq *GroupQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().Count().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (gq *GroupQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery().HasNext().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (gq *GroupQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(group.Label)
	if gq.gremlin != nil {
		v = gq.gremlin.Clone()
	}
	for _, p := range gq.predicates {
		p.Gremlin(v)
	}
	if len(gq.order) > 0 {
		v.Order()
		for _, p := range gq.order {
			p.Gremlin(v)
		}
	}
	switch limit, offset := gq.limit, gq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt64)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := gq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// GroupQuery is the builder for group-by Group entities.
type GroupGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GroupGroupBy) Aggregate(fns ...Aggregate) *GroupGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scan the result into the given value.
func (ggb *GroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch ggb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ggb.sqlScan(ctx, v)
	case dialect.Neptune:
		return ggb.gremlinScan(ctx, v)
	default:
		return errors.New("ggb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GroupGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ggb *GroupGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GroupGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GroupGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ggb *GroupGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GroupGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GroupGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ggb *GroupGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GroupGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GroupGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ggb *GroupGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GroupGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GroupGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ggb.sqlQuery().Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GroupGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql
	columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
	columns = append(columns, ggb.fields...)
	for _, fn := range ggb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(ggb.fields...)
}

func (ggb *GroupGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ggb.gremlinQuery().Query()
	if err := ggb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ggb.fields)+len(ggb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ggb *GroupGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ggb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ggb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ggb.gremlin.Group().
		By(__.Values(ggb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}
