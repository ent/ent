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

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	name         *string
	files        map[string]struct{}
	removedFiles map[string]struct{}
	predicates   []predicate.FileType
}

// Where adds a new predicate for the builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.predicates = append(ftu.predicates, ps...)
	return ftu
}

// SetName sets the name field.
func (ftu *FileTypeUpdate) SetName(s string) *FileTypeUpdate {
	ftu.name = &s
	return ftu
}

// AddFileIDs adds the files edge to File by ids.
func (ftu *FileTypeUpdate) AddFileIDs(ids ...string) *FileTypeUpdate {
	if ftu.files == nil {
		ftu.files = make(map[string]struct{})
	}
	for i := range ids {
		ftu.files[ids[i]] = struct{}{}
	}
	return ftu
}

// AddFiles adds the files edges to File.
func (ftu *FileTypeUpdate) AddFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFileIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (ftu *FileTypeUpdate) RemoveFileIDs(ids ...string) *FileTypeUpdate {
	if ftu.removedFiles == nil {
		ftu.removedFiles = make(map[string]struct{})
	}
	for i := range ids {
		ftu.removedFiles[ids[i]] = struct{}{}
	}
	return ftu
}

// RemoveFiles removes files edges to File.
func (ftu *FileTypeUpdate) RemoveFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {
	switch ftu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftu.sqlSave(ctx)
	case dialect.Gremlin:
		return ftu.gremlinSave(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftu *FileTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(filetype.FieldID).From(sql.Table(filetype.Table))
	for _, p := range ftu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ftu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := ftu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(filetype.Table).Where(sql.InInts(filetype.FieldID, ids...))
	)
	if value := ftu.name; value != nil {
		builder.Set(filetype.FieldName, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(ftu.removedFiles) > 0 {
		eids := make([]int, len(ftu.removedFiles))
		for eid := range ftu.removedFiles {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := sql.Update(filetype.FilesTable).
			SetNull(filetype.FilesColumn).
			Where(sql.InInts(filetype.FilesColumn, ids...)).
			Where(sql.InInts(file.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(ftu.files) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range ftu.files {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(file.FieldID, eid)
			}
			query, args := sql.Update(filetype.FilesTable).
				Set(filetype.FilesColumn, id).
				Where(sql.And(p, sql.IsNull(filetype.FilesColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(ftu.files) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"FileType\"", keys(ftu.files))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (ftu *FileTypeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftu.gremlin().Query()
	if err := ftu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (ftu *FileTypeUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(filetype.Label)
	for _, p := range ftu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := ftu.name; value != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, *value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, *value)),
		})
		v.Property(dsl.Single, filetype.FieldName, *value)
	}
	for id := range ftu.removedFiles {
		tr := rv.Clone().OutE(filetype.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range ftu.files {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ErrConstraintFailed{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	id           string
	name         *string
	files        map[string]struct{}
	removedFiles map[string]struct{}
}

// SetName sets the name field.
func (ftuo *FileTypeUpdateOne) SetName(s string) *FileTypeUpdateOne {
	ftuo.name = &s
	return ftuo
}

// AddFileIDs adds the files edge to File by ids.
func (ftuo *FileTypeUpdateOne) AddFileIDs(ids ...string) *FileTypeUpdateOne {
	if ftuo.files == nil {
		ftuo.files = make(map[string]struct{})
	}
	for i := range ids {
		ftuo.files[ids[i]] = struct{}{}
	}
	return ftuo
}

// AddFiles adds the files edges to File.
func (ftuo *FileTypeUpdateOne) AddFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFileIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (ftuo *FileTypeUpdateOne) RemoveFileIDs(ids ...string) *FileTypeUpdateOne {
	if ftuo.removedFiles == nil {
		ftuo.removedFiles = make(map[string]struct{})
	}
	for i := range ids {
		ftuo.removedFiles[ids[i]] = struct{}{}
	}
	return ftuo
}

// RemoveFiles removes files edges to File.
func (ftuo *FileTypeUpdateOne) RemoveFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFileIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {
	switch ftuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftuo.sqlSave(ctx)
	case dialect.Gremlin:
		return ftuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	ft, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftuo *FileTypeUpdateOne) sqlSave(ctx context.Context) (ft *FileType, err error) {
	selector := sql.Select(filetype.Columns...).From(sql.Table(filetype.Table))
	filetype.ID(ftuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ftuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		ft = &FileType{config: ftuo.config}
		if err := ft.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into FileType: %v", err)
		}
		id = ft.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: FileType not found with id: %v", ftuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one FileType with the same id: %v", ftuo.id)
	}

	tx, err := ftuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(filetype.Table).Where(sql.InInts(filetype.FieldID, ids...))
	)
	if value := ftuo.name; value != nil {
		builder.Set(filetype.FieldName, *value)
		ft.Name = *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(ftuo.removedFiles) > 0 {
		eids := make([]int, len(ftuo.removedFiles))
		for eid := range ftuo.removedFiles {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := sql.Update(filetype.FilesTable).
			SetNull(filetype.FilesColumn).
			Where(sql.InInts(filetype.FilesColumn, ids...)).
			Where(sql.InInts(file.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(ftuo.files) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range ftuo.files {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(file.FieldID, eid)
			}
			query, args := sql.Update(filetype.FilesTable).
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
			if int(affected) < len(ftuo.files) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"FileType\"", keys(ftuo.files))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FileTypeUpdateOne) gremlinSave(ctx context.Context) (*FileType, error) {
	res := &gremlin.Response{}
	query, bindings := ftuo.gremlin(ftuo.id).Query()
	if err := ftuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FileType{config: ftuo.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FileTypeUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := ftuo.name; value != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, *value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, *value)),
		})
		v.Property(dsl.Single, filetype.FieldName, *value)
	}
	for id := range ftuo.removedFiles {
		tr := rv.Clone().OutE(filetype.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range ftuo.files {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
