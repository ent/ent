// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"fbc/ent/entc/integration/ent/file"
	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	size       *int
	name       *string
	predicates []predicate.File
}

// Where adds a new predicate for the builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetSize sets the size field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.size = &i
	return fu
}

// SetName sets the name field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.name = &s
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	if fu.size != nil {
		if err := file.SizeValidator(*fu.size); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	switch fu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fu.sqlSave(ctx)
	case dialect.Neptune:
		return fu.gremlinSave(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(file.FieldID).From(sql.Table(file.Table))
	for _, p := range fu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = fu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := fu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(file.Table).Where(sql.InInts(file.FieldID, ids...))
	)
	if fu.size != nil {
		update = true
		builder.Set(file.FieldSize, *fu.size)
	}
	if fu.name != nil {
		update = true
		builder.Set(file.FieldName, *fu.name)
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

func (fu *FileUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := fu.gremlin().Query()
	if err := fu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (fu *FileUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(file.Label)
	for _, p := range fu.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if fu.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fu.size)
	}
	if fu.name != nil {
		v.Property(dsl.Single, file.FieldName, *fu.name)
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	id   string
	size *int
	name *string
}

// SetSize sets the size field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.size = &i
	return fuo
}

// SetName sets the name field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.name = &s
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	if fuo.size != nil {
		if err := file.SizeValidator(*fuo.size); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	switch fuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fuo.sqlSave(ctx)
	case dialect.Neptune:
		return fuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (f *File, err error) {
	selector := sql.Select(file.Columns...).From(sql.Table(file.Table))
	file.ID(fuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = fuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		f = &File{config: fuo.config}
		if err := f.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into File: %v", err)
		}
		id = f.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: File not found with id: %v", fuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one File with the same id: %v", fuo.id)
	}

	tx, err := fuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(file.Table).Where(sql.InInts(file.FieldID, ids...))
	)
	if fuo.size != nil {
		update = true
		builder.Set(file.FieldSize, *fuo.size)
		f.Size = *fuo.size
	}
	if fuo.name != nil {
		update = true
		builder.Set(file.FieldName, *fuo.name)
		f.Name = *fuo.name
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
	return f, nil
}

func (fuo *FileUpdateOne) gremlinSave(ctx context.Context) (*File, error) {
	res := &gremlin.Response{}
	query, bindings := fuo.gremlin(fuo.id).Query()
	if err := fuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	f := &File{config: fuo.config}
	if err := f.FromResponse(res); err != nil {
		return nil, err
	}
	return f, nil
}

func (fuo *FileUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if fuo.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fuo.size)
	}
	if fuo.name != nil {
		v.Property(dsl.Single, file.FieldName, *fuo.name)
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
