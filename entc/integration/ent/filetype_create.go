// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/schema/field"
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
		ft   = &FileType{config: ftc.config}
		spec = &sqlgraph.CreateSpec{
			Table: filetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: filetype.FieldID,
			},
		}
	)
	if value := ftc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: filetype.FieldName,
		})
		ft.Name = *value
	}
	if nodes := ftc.files; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ftc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	ft.ID = strconv.FormatInt(id, 10)
	return ft, nil
}
