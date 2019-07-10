// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"strconv"

	"fbc/ent/entc/integration/migrate/entv2/user"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age   *int
	name  *string
	phone *string
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.age = &i
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetPhone sets the phone field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.phone = &s
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("entv2: missing required field \"age\"")
	}
	if uc.name == nil {
		return nil, errors.New("entv2: missing required field \"name\"")
	}
	if uc.phone == nil {
		return nil, errors.New("entv2: missing required field \"phone\"")
	}
	switch uc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return uc.sqlSave(ctx)
	case dialect.Neptune:
		return uc.gremlinSave(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		res sql.Result
		u   = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(user.Table).Default(uc.driver.Dialect())
	if uc.age != nil {
		builder.Set(user.FieldAge, *uc.age)
		u.Age = *uc.age
	}
	if uc.name != nil {
		builder.Set(user.FieldName, *uc.name)
		u.Name = *uc.name
	}
	if uc.phone != nil {
		builder.Set(user.FieldPhone, *uc.phone)
		u.Phone = *uc.phone
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserCreate) gremlinSave(ctx context.Context) (*User, error) {
	res := &gremlin.Response{}
	query, bindings := uc.gremlin().Query()
	if err := uc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	u := &User{config: uc.config}
	if err := u.FromResponse(res); err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserCreate) gremlin() *dsl.Traversal {
	v := g.AddV(user.Label)
	if uc.age != nil {
		v.Property(dsl.Single, user.FieldAge, *uc.age)
	}
	if uc.name != nil {
		v.Property(dsl.Single, user.FieldName, *uc.name)
	}
	if uc.phone != nil {
		v.Property(dsl.Single, user.FieldPhone, *uc.phone)
	}
	return v.ValueMap(true)
}
