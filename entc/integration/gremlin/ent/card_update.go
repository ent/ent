// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"
	errors "errors"
	fmt "fmt"
	time "time"

	gremlin "entgo.io/ent/dialect/gremlin"
	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	g "entgo.io/ent/dialect/gremlin/graph/dsl/g"
	p "entgo.io/ent/dialect/gremlin/graph/dsl/p"
	card "entgo.io/ent/entc/integration/gremlin/ent/card"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
	spec "entgo.io/ent/entc/integration/gremlin/ent/spec"
	user "entgo.io/ent/entc/integration/gremlin/ent/user"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CardUpdate) SetUpdateTime(t time.Time) *CardUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetBalance sets the "balance" field.
func (cu *CardUpdate) SetBalance(f float64) *CardUpdate {
	cu.mutation.ResetBalance()
	cu.mutation.SetBalance(f)
	return cu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (cu *CardUpdate) SetNillableBalance(f *float64) *CardUpdate {
	if f != nil {
		cu.SetBalance(*f)
	}
	return cu
}

// AddBalance adds f to the "balance" field.
func (cu *CardUpdate) AddBalance(f float64) *CardUpdate {
	cu.mutation.AddBalance(f)
	return cu
}

// SetName sets the "name" field.
func (cu *CardUpdate) SetName(s string) *CardUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CardUpdate) SetNillableName(s *string) *CardUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *CardUpdate) ClearName() *CardUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *CardUpdate) SetOwnerID(id string) *CardUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cu *CardUpdate) SetNillableOwnerID(id *string) *CardUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *CardUpdate) SetOwner(u *User) *CardUpdate {
	return cu.SetOwnerID(u.ID)
}

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (cu *CardUpdate) AddSpecIDs(ids ...string) *CardUpdate {
	cu.mutation.AddSpecIDs(ids...)
	return cu
}

// AddSpec adds the "spec" edges to the Spec entity.
func (cu *CardUpdate) AddSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddSpecIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cu *CardUpdate) Mutation() *CardMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *CardUpdate) ClearOwner() *CardUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// ClearSpec clears all "spec" edges to the Spec entity.
func (cu *CardUpdate) ClearSpec() *CardUpdate {
	cu.mutation.ClearSpec()
	return cu
}

// RemoveSpecIDs removes the "spec" edge to Spec entities by IDs.
func (cu *CardUpdate) RemoveSpecIDs(ids ...string) *CardUpdate {
	cu.mutation.RemoveSpecIDs(ids...)
	return cu
}

// RemoveSpec removes "spec" edges to Spec entities.
func (cu *CardUpdate) RemoveSpec(s ...*Spec) *CardUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveSpecIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, CardMutation](ctx, cu.gremlinSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CardUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CardUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	return nil
}

func (cu *CardUpdate) gremlinSave(ctx context.Context) (int, error) {
	if err := cu.check(); err != nil {
		return 0, err
	}
	res := &gremlin.Response{}
	query, bindings := cu.gremlin().Query()
	if err := cu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	cu.mutation.done = true
	return res.ReadInt()
}

func (cu *CardUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V().HasLabel(card.Label)
	for _, p := range cu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cu.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, card.FieldUpdateTime, value)
	}
	if value, ok := cu.mutation.Balance(); ok {
		v.Property(dsl.Single, card.FieldBalance, value)
	}
	if value, ok := cu.mutation.AddedBalance(); ok {
		v.Property(dsl.Single, card.FieldBalance, __.Union(__.Values(card.FieldBalance), __.Constant(value)).Sum())
	}
	if value, ok := cu.mutation.Name(); ok {
		v.Property(dsl.Single, card.FieldName, value)
	}
	var properties []any
	if cu.mutation.NameCleared() {
		properties = append(properties, card.FieldName)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if cu.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cu.mutation.OwnerIDs() {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for _, id := range cu.mutation.RemovedSpecIDs() {
		tr := rv.Clone().InE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cu.mutation.SpecIDs() {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CardUpdateOne) SetUpdateTime(t time.Time) *CardUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetBalance sets the "balance" field.
func (cuo *CardUpdateOne) SetBalance(f float64) *CardUpdateOne {
	cuo.mutation.ResetBalance()
	cuo.mutation.SetBalance(f)
	return cuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableBalance(f *float64) *CardUpdateOne {
	if f != nil {
		cuo.SetBalance(*f)
	}
	return cuo
}

// AddBalance adds f to the "balance" field.
func (cuo *CardUpdateOne) AddBalance(f float64) *CardUpdateOne {
	cuo.mutation.AddBalance(f)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CardUpdateOne) SetName(s string) *CardUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableName(s *string) *CardUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *CardUpdateOne) ClearName() *CardUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *CardUpdateOne) SetOwnerID(id string) *CardUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cuo *CardUpdateOne) SetNillableOwnerID(id *string) *CardUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (cuo *CardUpdateOne) AddSpecIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.AddSpecIDs(ids...)
	return cuo
}

// AddSpec adds the "spec" edges to the Spec entity.
func (cuo *CardUpdateOne) AddSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddSpecIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cuo *CardUpdateOne) Mutation() *CardMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *CardUpdateOne) ClearOwner() *CardUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// ClearSpec clears all "spec" edges to the Spec entity.
func (cuo *CardUpdateOne) ClearSpec() *CardUpdateOne {
	cuo.mutation.ClearSpec()
	return cuo
}

// RemoveSpecIDs removes the "spec" edge to Spec entities by IDs.
func (cuo *CardUpdateOne) RemoveSpecIDs(ids ...string) *CardUpdateOne {
	cuo.mutation.RemoveSpecIDs(ids...)
	return cuo
}

// RemoveSpec removes "spec" edges to Spec entities.
func (cuo *CardUpdateOne) RemoveSpec(s ...*Spec) *CardUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveSpecIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Card entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	cuo.defaults()
	return withHooks[*Card, CardMutation](ctx, cuo.gremlinSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CardUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := card.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CardUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Card.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *CardUpdateOne) gremlinSave(ctx context.Context) (*Card, error) {
	if err := cuo.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	query, bindings := cuo.gremlin(id).Query()
	if err := cuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	cuo.mutation.done = true
	c := &Card{config: cuo.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cuo *CardUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := cuo.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, card.FieldUpdateTime, value)
	}
	if value, ok := cuo.mutation.Balance(); ok {
		v.Property(dsl.Single, card.FieldBalance, value)
	}
	if value, ok := cuo.mutation.AddedBalance(); ok {
		v.Property(dsl.Single, card.FieldBalance, __.Union(__.Values(card.FieldBalance), __.Constant(value)).Sum())
	}
	if value, ok := cuo.mutation.Name(); ok {
		v.Property(dsl.Single, card.FieldName, value)
	}
	var properties []any
	if cuo.mutation.NameCleared() {
		properties = append(properties, card.FieldName)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if cuo.mutation.OwnerCleared() {
		tr := rv.Clone().InE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cuo.mutation.OwnerIDs() {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for _, id := range cuo.mutation.RemovedSpecIDs() {
		tr := rv.Clone().InE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range cuo.mutation.SpecIDs() {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	if len(cuo.fields) > 0 {
		fields := make([]any, 0, len(cuo.fields)+1)
		fields = append(fields, true)
		for _, f := range cuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
