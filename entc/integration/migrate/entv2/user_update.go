// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/migrate/entv2/user"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age        *int
	name       *string
	phone      *string
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...ent.Predicate) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.age = &i
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.phone = &s
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	switch uu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return uu.sqlSave(ctx)
	case dialect.Neptune:
		vertices, err := uu.gremlinSave(ctx)
		return len(vertices), err
	default:
		return 0, errors.New("entv2: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(user.FieldID).From(sql.Table(user.Table))
	for _, p := range uu.predicates {
		p.SQL(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = uu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := uu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if uu.age != nil {
		update = true
		builder.Set(user.FieldAge, *uu.age)
	}
	if uu.name != nil {
		update = true
		builder.Set(user.FieldName, *uu.name)
	}
	if uu.phone != nil {
		update = true
		builder.Set(user.FieldPhone, *uu.phone)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (uu *UserUpdate) gremlinSave(ctx context.Context) ([]*User, error) {
	res := &gremlin.Response{}
	query, bindings := uu.gremlin().Query()
	if err := uu.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	var us Users
	us.config(uu.config)
	if err := us.FromResponse(res); err != nil {
		return nil, err
	}
	return us, nil
}

func (uu *UserUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(user.Label)
	for _, p := range uu.predicates {
		p.Gremlin(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if uu.age != nil {
		v.Property(dsl.Single, user.FieldAge, *uu.age)
	}
	if uu.name != nil {
		v.Property(dsl.Single, user.FieldName, *uu.name)
	}
	if uu.phone != nil {
		v.Property(dsl.Single, user.FieldPhone, *uu.phone)
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id    string
	age   *int
	name  *string
	phone *string
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.phone = &s
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	switch uuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return uuo.sqlSave(ctx)
	case dialect.Neptune:
		return uuo.gremlinSave(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	selector := sql.Select(user.Columns...).From(sql.Table(user.Table))
	user.ID(uuo.id).SQL(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = uuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		u = &User{config: uuo.config}
		if err := u.FromRows(rows); err != nil {
			return nil, fmt.Errorf("entv2: failed scanning row into User: %v", err)
		}
		id = u.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("entv2: User not found with id: %v", uuo.id)
	case n > 1:
		return nil, fmt.Errorf("entv2: more than one User with the same id: %v", uuo.id)
	}

	tx, err := uuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if uuo.age != nil {
		update = true
		builder.Set(user.FieldAge, *uuo.age)
		u.Age = *uuo.age
	}
	if uuo.name != nil {
		update = true
		builder.Set(user.FieldName, *uuo.name)
		u.Name = *uuo.name
	}
	if uuo.phone != nil {
		update = true
		builder.Set(user.FieldPhone, *uuo.phone)
		u.Phone = *uuo.phone
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}

func (uuo *UserUpdateOne) gremlinSave(ctx context.Context) (*User, error) {
	res := &gremlin.Response{}
	query, bindings := uuo.gremlin(uuo.id).Query()
	if err := uuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	u := &User{config: uuo.config}
	if err := u.FromResponse(res); err != nil {
		return nil, err
	}
	return u, nil
}

func (uuo *UserUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if uuo.age != nil {
		v.Property(dsl.Single, user.FieldAge, *uuo.age)
	}
	if uuo.name != nil {
		v.Property(dsl.Single, user.FieldName, *uuo.name)
	}
	if uuo.phone != nil {
		v.Property(dsl.Single, user.FieldPhone, *uuo.phone)
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
