// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"fbc/ent/entc/integration/ent/file"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	size *int
	name *string
}

// SetSize sets the size field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.size = &i
	return fc
}

// SetName sets the name field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.name = &s
	return fc
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if fc.size == nil {
		return nil, errors.New("ent: missing required field \"size\"")
	}
	if err := file.SizeValidator(*fc.size); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
	}
	if fc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	switch fc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fc.sqlSave(ctx)
	case dialect.Neptune:
		return fc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	var (
		res sql.Result
		f   = &File{config: fc.config}
	)
	tx, err := fc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(file.Table).Default(fc.driver.Dialect())
	if fc.size != nil {
		builder.Set(file.FieldSize, *fc.size)
		f.Size = *fc.size
	}
	if fc.name != nil {
		builder.Set(file.FieldName, *fc.name)
		f.Name = *fc.name
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	f.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return f, nil
}

func (fc *FileCreate) gremlinSave(ctx context.Context) (*File, error) {
	res := &gremlin.Response{}
	query, bindings := fc.gremlin().Query()
	if err := fc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	f := &File{config: fc.config}
	if err := f.FromResponse(res); err != nil {
		return nil, err
	}
	return f, nil
}

func (fc *FileCreate) gremlin() *dsl.Traversal {
	v := g.AddV(file.Label)
	if fc.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fc.size)
	}
	if fc.name != nil {
		v.Property(dsl.Single, file.FieldName, *fc.name)
	}
	return v.ValueMap(true)
}
