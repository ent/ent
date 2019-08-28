// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	size  *int
	name  *string
	text  *string
	user  *string
	group *string
	owner map[string]struct{}
	_type map[string]struct{}
}

// SetSize sets the size field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.size = &i
	return fc
}

// SetNillableSize sets the size field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetName sets the name field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.name = &s
	return fc
}

// SetText sets the text field.
func (fc *FileCreate) SetText(s string) *FileCreate {
	fc.text = &s
	return fc
}

// SetNillableText sets the text field if the given value is not nil.
func (fc *FileCreate) SetNillableText(s *string) *FileCreate {
	if s != nil {
		fc.SetText(*s)
	}
	return fc
}

// SetUser sets the user field.
func (fc *FileCreate) SetUser(s string) *FileCreate {
	fc.user = &s
	return fc
}

// SetNillableUser sets the user field if the given value is not nil.
func (fc *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		fc.SetUser(*s)
	}
	return fc
}

// SetGroup sets the group field.
func (fc *FileCreate) SetGroup(s string) *FileCreate {
	fc.group = &s
	return fc
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fc *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		fc.SetGroup(*s)
	}
	return fc
}

// SetOwnerID sets the owner edge to User by id.
func (fc *FileCreate) SetOwnerID(id string) *FileCreate {
	if fc.owner == nil {
		fc.owner = make(map[string]struct{})
	}
	fc.owner[id] = struct{}{}
	return fc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the owner edge to User.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetTypeID sets the type edge to FileType by id.
func (fc *FileCreate) SetTypeID(id string) *FileCreate {
	if fc._type == nil {
		fc._type = make(map[string]struct{})
	}
	fc._type[id] = struct{}{}
	return fc
}

// SetNillableTypeID sets the type edge to FileType by id if the given value is not nil.
func (fc *FileCreate) SetNillableTypeID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetTypeID(*id)
	}
	return fc
}

// SetType sets the type edge to FileType.
func (fc *FileCreate) SetType(f *FileType) *FileCreate {
	return fc.SetTypeID(f.ID)
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if fc.size == nil {
		v := file.DefaultSize
		fc.size = &v
	}
	if err := file.SizeValidator(*fc.size); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
	}
	if fc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(fc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if len(fc._type) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"type\"")
	}
	switch fc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fc.sqlSave(ctx)
	case dialect.Gremlin:
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
	if fc.text != nil {
		builder.Set(file.FieldText, *fc.text)
		f.Text = *fc.text
	}
	if fc.user != nil {
		builder.Set(file.FieldUser, *fc.user)
		f.User = fc.user
	}
	if fc.group != nil {
		builder.Set(file.FieldGroup, *fc.group)
		f.Group = *fc.group
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
	if len(fc.owner) > 0 {
		for eid := range fc.owner {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := sql.Update(file.OwnerTable).
				Set(file.OwnerColumn, eid).
				Where(sql.EQ(file.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(fc._type) > 0 {
		for eid := range fc._type {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := sql.Update(file.TypeTable).
				Set(file.TypeColumn, eid).
				Where(sql.EQ(file.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
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
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(file.Label)
	if fc.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fc.size)
	}
	if fc.name != nil {
		v.Property(dsl.Single, file.FieldName, *fc.name)
	}
	if fc.text != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldText, *fc.text).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldText, *fc.text)),
		})
		v.Property(dsl.Single, file.FieldText, *fc.text)
	}
	if fc.user != nil {
		v.Property(dsl.Single, file.FieldUser, *fc.user)
	}
	if fc.group != nil {
		v.Property(dsl.Single, file.FieldGroup, *fc.group)
	}
	for id := range fc.owner {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	for id := range fc._type {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}
