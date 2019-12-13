// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	name  *string
	files map[string]struct{}
}

// SetName sets the name field.
func (ftc *FileTypeCreate) SetName(s string) *FileTypeCreate {
	ftc.name = &s
	return ftc
}

// AddFileIDs adds the files edge to File by ids.
func (ftc *FileTypeCreate) AddFileIDs(ids ...string) *FileTypeCreate {
	if ftc.files == nil {
		ftc.files = make(map[string]struct{})
	}
	for i := range ids {
		ftc.files[ids[i]] = struct{}{}
	}
	return ftc
}

// AddFiles adds the files edges to File.
func (ftc *FileTypeCreate) AddFiles(f ...*File) *FileTypeCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddFileIDs(ids...)
}

// Save creates the FileType in the database.
func (ftc *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	if ftc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	return ftc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FileTypeCreate) sqlSave(ctx context.Context) (*FileType, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(ftc.driver.Dialect())
		ft      = &FileType{config: ftc.config}
	)
	tx, err := ftc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(filetype.Table).Default()
	if value := ftc.name; value != nil {
		insert.Set(filetype.FieldName, *value)
		ft.Name = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(filetype.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	ft.ID = strconv.FormatInt(id, 10)
	if len(ftc.files) > 0 {
		p := sql.P()
		for eid := range ftc.files {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			p.Or().EQ(file.FieldID, eid)
		}
		query, args := builder.Update(filetype.FilesTable).
			Set(filetype.FilesColumn, id).
			Where(sql.And(p, sql.IsNull(filetype.FilesColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(ftc.files) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"FileType\"", keys(ftc.files))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ft, nil
}
