// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/json/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	url     **url.URL
	raw     *json.RawMessage
	dirs    *[]http.Dir
	ints    *[]int
	floats  *[]float64
	strings *[]string
}

// SetURL sets the url field.
func (uc *UserCreate) SetURL(u *url.URL) *UserCreate {
	uc.url = &u
	return uc
}

// SetRaw sets the raw field.
func (uc *UserCreate) SetRaw(jm json.RawMessage) *UserCreate {
	uc.raw = &jm
	return uc
}

// SetDirs sets the dirs field.
func (uc *UserCreate) SetDirs(h []http.Dir) *UserCreate {
	uc.dirs = &h
	return uc
}

// SetInts sets the ints field.
func (uc *UserCreate) SetInts(i []int) *UserCreate {
	uc.ints = &i
	return uc
}

// SetFloats sets the floats field.
func (uc *UserCreate) SetFloats(f []float64) *UserCreate {
	uc.floats = &f
	return uc
}

// SetStrings sets the strings field.
func (uc *UserCreate) SetStrings(s []string) *UserCreate {
	uc.strings = &s
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
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
	builder := sql.Dialect(uc.driver.Dialect()).
		Insert(user.Table).
		Default()
	if value := uc.url; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldURL, buf)
		u.URL = *value
	}
	if value := uc.raw; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldRaw, buf)
		u.Raw = *value
	}
	if value := uc.dirs; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldDirs, buf)
		u.Dirs = *value
	}
	if value := uc.ints; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldInts, buf)
		u.Ints = *value
	}
	if value := uc.floats; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldFloats, buf)
		u.Floats = *value
	}
	if value := uc.strings; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldStrings, buf)
		u.Strings = *value
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
