// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/migrate/entv2/pet"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (pu *PetUpdate) Where(ps ...ent.Predicate) *PetUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	switch pu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pu.sqlSave(ctx)
	case dialect.Neptune:
		vertices, err := pu.gremlinSave(ctx)
		return len(vertices), err
	default:
		return 0, errors.New("entv2: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(pet.FieldID).From(sql.Table(pet.Table))
	for _, p := range pu.predicates {
		p.SQL(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = pu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("entv2: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := pu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (pu *PetUpdate) gremlinSave(ctx context.Context) ([]*Pet, error) {
	res := &gremlin.Response{}
	query, bindings := pu.gremlin().Query()
	if err := pu.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	var pes Pets
	pes.config(pu.config)
	if err := pes.FromResponse(res); err != nil {
		return nil, err
	}
	return pes, nil
}

func (pu *PetUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(pet.Label)
	for _, p := range pu.predicates {
		p.Gremlin(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	id string
}

// Save executes the query and returns the updated entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	switch puo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return puo.sqlSave(ctx)
	case dialect.Neptune:
		return puo.gremlinSave(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	pe, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (pe *Pet, err error) {
	selector := sql.Select(pet.Columns...).From(sql.Table(pet.Table))
	pet.ID(puo.id).SQL(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = puo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		pe = &Pet{config: puo.config}
		if err := pe.FromRows(rows); err != nil {
			return nil, fmt.Errorf("entv2: failed scanning row into Pet: %v", err)
		}
		id = pe.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("entv2: Pet not found with id: %v", puo.id)
	case n > 1:
		return nil, fmt.Errorf("entv2: more than one Pet with the same id: %v", puo.id)
	}

	tx, err := puo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}

func (puo *PetUpdateOne) gremlinSave(ctx context.Context) (*Pet, error) {
	res := &gremlin.Response{}
	query, bindings := puo.gremlin(puo.id).Query()
	if err := puo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	pe := &Pet{config: puo.config}
	if err := pe.FromResponse(res); err != nil {
		return nil, err
	}
	return pe, nil
}

func (puo *PetUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
