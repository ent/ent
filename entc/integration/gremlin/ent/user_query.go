// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.User
	// intermediate query.
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.offset = &offset
	return uq
}

// Order adds an order step to the query.
func (uq *UserQuery) Order(o ...Order) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryCard chains the current query on the card edge.
func (uq *UserQuery) QueryCard() *CardQuery {
	query := &CardQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.CardLabel).InV()
	return query
}

// QueryPets chains the current query on the pets edge.
func (uq *UserQuery) QueryPets() *PetQuery {
	query := &PetQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.PetsLabel).InV()
	return query
}

// QueryFiles chains the current query on the files edge.
func (uq *UserQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.FilesLabel).InV()
	return query
}

// QueryGroups chains the current query on the groups edge.
func (uq *UserQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.GroupsLabel).InV()
	return query
}

// QueryFriends chains the current query on the friends edge.
func (uq *UserQuery) QueryFriends() *UserQuery {
	query := &UserQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.Both(user.FriendsLabel)
	return query
}

// QueryFollowers chains the current query on the followers edge.
func (uq *UserQuery) QueryFollowers() *UserQuery {
	query := &UserQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.InE(user.FollowingLabel).OutV()
	return query
}

// QueryFollowing chains the current query on the following edge.
func (uq *UserQuery) QueryFollowing() *UserQuery {
	query := &UserQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.FollowingLabel).InV()
	return query
}

// QueryTeam chains the current query on the team edge.
func (uq *UserQuery) QueryTeam() *PetQuery {
	query := &PetQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.TeamLabel).InV()
	return query
}

// QuerySpouse chains the current query on the spouse edge.
func (uq *UserQuery) QuerySpouse() *UserQuery {
	query := &UserQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.Both(user.SpouseLabel)
	return query
}

// QueryChildren chains the current query on the children edge.
func (uq *UserQuery) QueryChildren() *UserQuery {
	query := &UserQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.InE(user.ParentLabel).OutV()
	return query
}

// QueryParent chains the current query on the parent edge.
func (uq *UserQuery) QueryParent() *UserQuery {
	query := &UserQuery{config: uq.config}
	gremlin := uq.gremlinQuery()
	query.gremlin = gremlin.OutE(user.ParentLabel).InV()
	return query
}

// First returns the first User entity in the query. Returns *ErrNotFound when no user was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	us, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, &ErrNotFound{user.Label}
	}
	return us[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	u, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return u
}

// FirstID returns the first User id in the query. Returns *ErrNotFound when no id was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{user.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstXID(ctx context.Context) string {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only User entity in the query, returns an error if not exactly one entity was returned.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	us, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(us) {
	case 1:
		return us[0], nil
	case 0:
		return nil, &ErrNotFound{user.Label}
	default:
		return nil, &ErrNotSingular{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	u, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// OnlyID returns the only User id in the query, returns an error if not exactly one id was returned.
func (uq *UserQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{user.Label}
	default:
		err = &ErrNotSingular{user.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyXID(ctx context.Context) string {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	return uq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	us, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return us
}

// IDs executes the query and returns a list of User ids.
func (uq *UserQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []string {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	return uq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	return uq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	return &UserQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]Order{}, uq.order...),
		unique:     append([]string{}, uq.unique...),
		predicates: append([]predicate.User{}, uq.predicates...),
		// clone intermediate query.
		gremlin: uq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Age int `json:"age,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldAge).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	group := &UserGroupBy{config: uq.config}
	group.fields = append([]string{field}, fields...)
	group.gremlin = uq.gremlinQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Age int `json:"age,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldAge).
//		Scan(ctx, &v)
//
func (uq *UserQuery) Select(field string, fields ...string) *UserSelect {
	selector := &UserSelect{config: uq.config}
	selector.fields = append([]string{field}, fields...)
	selector.gremlin = uq.gremlinQuery()
	return selector
}

func (uq *UserQuery) gremlinAll(ctx context.Context) ([]*User, error) {
	res := &gremlin.Response{}
	query, bindings := uq.gremlinQuery().ValueMap(true).Query()
	if err := uq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var us Users
	if err := us.FromResponse(res); err != nil {
		return nil, err
	}
	us.config(uq.config)
	return us, nil
}

func (uq *UserQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := uq.gremlinQuery().Count().Query()
	if err := uq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (uq *UserQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := uq.gremlinQuery().HasNext().Query()
	if err := uq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (uq *UserQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(user.Label)
	if uq.gremlin != nil {
		v = uq.gremlin.Clone()
	}
	for _, p := range uq.predicates {
		p(v)
	}
	if len(uq.order) > 0 {
		v.Order()
		for _, p := range uq.order {
			p(v)
		}
	}
	switch limit, offset := uq.limit, uq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := uq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// UserGroupBy is the builder for group-by User entities.
type UserGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...Aggregate) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scan the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ugb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ugb *UserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ugb *UserGroupBy) StringsX(ctx context.Context) []string {
	v, err := ugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ugb *UserGroupBy) IntsX(ctx context.Context) []int {
	v, err := ugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ugb *UserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ugb *UserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ugb *UserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ugb *UserGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ugb.gremlinQuery().Query()
	if err := ugb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ugb.fields)+len(ugb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ugb *UserGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ugb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ugb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ugb.gremlin.Group().
		By(__.Values(ugb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// UserSelect is the builder for select fields of User entities.
type UserSelect struct {
	config
	fields []string
	// intermediate queries.
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scan the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v interface{}) error {
	return us.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (us *UserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := us.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (us *UserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (us *UserSelect) StringsX(ctx context.Context) []string {
	v, err := us.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (us *UserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (us *UserSelect) IntsX(ctx context.Context) []int {
	v, err := us.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (us *UserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (us *UserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := us.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (us *UserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (us *UserSelect) BoolsX(ctx context.Context) []bool {
	v, err := us.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (us *UserSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(us.fields) == 1 {
		if us.fields[0] != user.FieldID {
			traversal = us.gremlin.Values(us.fields...)
		} else {
			traversal = us.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(us.fields))
		for i, f := range us.fields {
			fields[i] = f
		}
		traversal = us.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := us.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(us.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
