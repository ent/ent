// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age    *int
	name   *string
	phone  *string
	buffer *[]byte
	title  *string
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

// SetBuffer sets the buffer field.
func (uc *UserCreate) SetBuffer(b []byte) *UserCreate {
	uc.buffer = &b
	return uc
}

// SetTitle sets the title field.
func (uc *UserCreate) SetTitle(s string) *UserCreate {
	uc.title = &s
	return uc
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uc *UserCreate) SetNillableTitle(s *string) *UserCreate {
	if s != nil {
		uc.SetTitle(*s)
	}
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
	if uc.buffer == nil {
		v := user.DefaultBuffer
		uc.buffer = &v
	}
	if uc.title == nil {
		v := user.DefaultTitle
		uc.title = &v
	}
	return uc.sqlSave(ctx)
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
	if uc.buffer != nil {
		builder.Set(user.FieldBuffer, *uc.buffer)
		u.Buffer = *uc.buffer
	}
	if uc.title != nil {
		builder.Set(user.FieldTitle, *uc.title)
		u.Title = *uc.title
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
