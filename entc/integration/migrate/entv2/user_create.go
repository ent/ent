// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age      *int
	name     *string
	phone    *string
	buffer   *[]byte
	title    *string
	new_name *string
	blob     *[]byte
	state    *user.State
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

// SetNewName sets the new_name field.
func (uc *UserCreate) SetNewName(s string) *UserCreate {
	uc.new_name = &s
	return uc
}

// SetNillableNewName sets the new_name field if the given value is not nil.
func (uc *UserCreate) SetNillableNewName(s *string) *UserCreate {
	if s != nil {
		uc.SetNewName(*s)
	}
	return uc
}

// SetBlob sets the blob field.
func (uc *UserCreate) SetBlob(b []byte) *UserCreate {
	uc.blob = &b
	return uc
}

// SetState sets the state field.
func (uc *UserCreate) SetState(u user.State) *UserCreate {
	uc.state = &u
	return uc
}

// SetNillableState sets the state field if the given value is not nil.
func (uc *UserCreate) SetNillableState(u *user.State) *UserCreate {
	if u != nil {
		uc.SetState(*u)
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
	if uc.state != nil {
		if err := user.StateValidator(*uc.state); err != nil {
			return nil, fmt.Errorf("entv2: validator failed for field \"state\": %v", err)
		}
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
		builder = sql.Dialect(uc.driver.Dialect())
		u       = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(user.Table).Default()
	if value := uc.age; value != nil {
		insert.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uc.name; value != nil {
		insert.Set(user.FieldName, *value)
		u.Name = *value
	}
	if value := uc.phone; value != nil {
		insert.Set(user.FieldPhone, *value)
		u.Phone = *value
	}
	if value := uc.buffer; value != nil {
		insert.Set(user.FieldBuffer, *value)
		u.Buffer = *value
	}
	if value := uc.title; value != nil {
		insert.Set(user.FieldTitle, *value)
		u.Title = *value
	}
	if value := uc.new_name; value != nil {
		insert.Set(user.FieldNewName, *value)
		u.NewName = *value
	}
	if value := uc.blob; value != nil {
		insert.Set(user.FieldBlob, *value)
		u.Blob = *value
	}
	if value := uc.state; value != nil {
		insert.Set(user.FieldState, *value)
		u.State = *value
	}
	id, err := insertLastID(ctx, tx, insert.Returning(user.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
