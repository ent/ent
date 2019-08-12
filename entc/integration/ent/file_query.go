// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"fbc/ent/entc/integration/ent/file"
	"fbc/ent/entc/integration/ent/predicate"
	"fbc/ent/entc/integration/ent/user"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FileQuery is the builder for querying File entities.
type FileQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.File
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
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

// Order adds an order step to the query.
func (fq *FileQuery) Order(o ...Order) *FileQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryOwner chains the current query on the owner edge.
func (fq *FileQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: fq.config}
	switch fq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		t1 := sql.Table(user.Table)
		t2 := fq.sqlQuery()
		t2.Select(t2.C(file.OwnerColumn))
		query.sql = sql.Select(t1.Columns(user.Columns...)...).
			From(t1).
			Join(t2).
			On(t1.C(user.FieldID), t2.C(file.OwnerColumn))
	case dialect.Neptune:
		gremlin := fq.gremlinQuery()
		query.gremlin = gremlin.InE(user.FilesLabel).OutV()
	}
	return query
}

// Get returns a File entity by its id.
func (fq *FileQuery) Get(ctx context.Context, id string) (*File, error) {
	return fq.Where(file.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (fq *FileQuery) GetX(ctx context.Context, id string) *File {
	f, err := fq.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return f
}

// First returns the first File entity in the query. Returns *ErrNotFound when no file was found.
func (fq *FileQuery) First(ctx context.Context) (*File, error) {
	fs, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fs) == 0 {
		return nil, &ErrNotFound{file.Label}
	}
	return fs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FileQuery) FirstX(ctx context.Context) *File {
	f, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return f
}

// FirstID returns the first File id in the query. Returns *ErrNotFound when no id was found.
func (fq *FileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{file.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (fq *FileQuery) FirstXID(ctx context.Context) string {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only File entity in the query, returns an error if not exactly one entity was returned.
func (fq *FileQuery) Only(ctx context.Context) (*File, error) {
	fs, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fs) {
	case 1:
		return fs[0], nil
	case 0:
		return nil, &ErrNotFound{file.Label}
	default:
		return nil, &ErrNotSingular{file.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FileQuery) OnlyX(ctx context.Context) *File {
	f, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// OnlyID returns the only File id in the query, returns an error if not exactly one id was returned.
func (fq *FileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{file.Label}
	default:
		err = &ErrNotSingular{file.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (fq *FileQuery) OnlyXID(ctx context.Context) string {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Files.
func (fq *FileQuery) All(ctx context.Context) ([]*File, error) {
	switch fq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fq.sqlAll(ctx)
	case dialect.Neptune:
		return fq.gremlinAll(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (fq *FileQuery) AllX(ctx context.Context) []*File {
	fs, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fs
}

// IDs executes the query and returns a list of File ids.
func (fq *FileQuery) IDs(ctx context.Context) ([]string, error) {
	switch fq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fq.sqlIDs(ctx)
	case dialect.Neptune:
		return fq.gremlinIDs(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FileQuery) IDsX(ctx context.Context) []string {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FileQuery) Count(ctx context.Context) (int, error) {
	switch fq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fq.sqlCount(ctx)
	case dialect.Neptune:
		return fq.gremlinCount(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
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
	switch fq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fq.sqlExist(ctx)
	case dialect.Neptune:
		return fq.gremlinExist(ctx)
	default:
		return false, errors.New("ent: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FileQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FileQuery) Clone() *FileQuery {
	return &FileQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]Order{}, fq.order...),
		unique:     append([]string{}, fq.unique...),
		predicates: append([]predicate.File{}, fq.predicates...),
		// clone intermediate queries.
		sql:     fq.sql.Clone(),
		gremlin: fq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
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
//
func (fq *FileQuery) GroupBy(field string, fields ...string) *FileGroupBy {
	group := &FileGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	switch fq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = fq.sqlQuery()
	case dialect.Neptune:
		group.gremlin = fq.gremlinQuery()
	}
	return group
}

func (fq *FileQuery) sqlAll(ctx context.Context) ([]*File, error) {
	rows := &sql.Rows{}
	selector := fq.sqlQuery()
	if unique := fq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := fq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var fs Files
	if err := fs.FromRows(rows); err != nil {
		return nil, err
	}
	fs.config(fq.config)
	return fs, nil
}

func (fq *FileQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := fq.sqlQuery()
	unique := []string{file.FieldID}
	if len(fq.unique) > 0 {
		unique = fq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := fq.driver.Query(ctx, query, args, rows); err != nil {
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

func (fq *FileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (fq *FileQuery) sqlIDs(ctx context.Context) ([]string, error) {
	vs, err := fq.sqlAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range vs {
		ids = append(ids, v.ID)
	}
	return ids, nil
}

func (fq *FileQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(file.Table)
	selector := sql.Select(t1.Columns(file.Columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(file.Columns...)...)
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
		selector.Offset(*offset).Limit(math.MaxInt64)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (fq *FileQuery) gremlinIDs(ctx context.Context) ([]string, error) {
	res := &gremlin.Response{}
	query, bindings := fq.gremlinQuery().Query()
	if err := fq.driver.Exec(ctx, query, bindings, res); err != nil {
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

func (fq *FileQuery) gremlinAll(ctx context.Context) ([]*File, error) {
	res := &gremlin.Response{}
	query, bindings := fq.gremlinQuery().ValueMap(true).Query()
	if err := fq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fs Files
	if err := fs.FromResponse(res); err != nil {
		return nil, err
	}
	fs.config(fq.config)
	return fs, nil
}

func (fq *FileQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := fq.gremlinQuery().Count().Query()
	if err := fq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (fq *FileQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := fq.gremlinQuery().HasNext().Query()
	if err := fq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (fq *FileQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(file.Label)
	if fq.gremlin != nil {
		v = fq.gremlin.Clone()
	}
	for _, p := range fq.predicates {
		p(v)
	}
	if len(fq.order) > 0 {
		v.Order()
		for _, p := range fq.order {
			p(v)
		}
	}
	switch limit, offset := fq.limit, fq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt64)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := fq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// FileQuery is the builder for group-by File entities.
type FileGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FileGroupBy) Aggregate(fns ...Aggregate) *FileGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scan the result into the given value.
func (fgb *FileGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch fgb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fgb.sqlScan(ctx, v)
	case dialect.Neptune:
		return fgb.gremlinScan(ctx, v)
	default:
		return errors.New("fgb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FileGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FileGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (fgb *FileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fgb.sqlQuery().Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FileGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql
	columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
	columns = append(columns, fgb.fields...)
	for _, fn := range fgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(fgb.fields...)
}

func (fgb *FileGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := fgb.gremlinQuery().Query()
	if err := fgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(fgb.fields)+len(fgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (fgb *FileGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range fgb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range fgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return fgb.gremlin.Group().
		By(__.Values(fgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}
