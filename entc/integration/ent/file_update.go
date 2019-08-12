// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"fbc/ent/entc/integration/ent/file"
	"fbc/ent/entc/integration/ent/predicate"
	"fbc/ent/entc/integration/ent/user"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	size         *int
	name         *string
	user         *string
	group        *string
	owner        map[string]struct{}
	clearedOwner bool
	predicates   []predicate.File
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

// SetNillableSize sets the size field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// SetName sets the name field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.name = &s
	return fu
}

// SetUser sets the user field.
func (fu *FileUpdate) SetUser(s string) *FileUpdate {
	fu.user = &s
	return fu
}

// SetNillableUser sets the user field if the given value is not nil.
func (fu *FileUpdate) SetNillableUser(s *string) *FileUpdate {
	if s != nil {
		fu.SetUser(*s)
	}
	return fu
}

// SetGroup sets the group field.
func (fu *FileUpdate) SetGroup(s string) *FileUpdate {
	fu.group = &s
	return fu
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fu *FileUpdate) SetNillableGroup(s *string) *FileUpdate {
	if s != nil {
		fu.SetGroup(*s)
	}
	return fu
}

// SetOwnerID sets the owner edge to User by id.
func (fu *FileUpdate) SetOwnerID(id string) *FileUpdate {
	if fu.owner == nil {
		fu.owner = make(map[string]struct{})
	}
	fu.owner[id] = struct{}{}
	return fu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fu *FileUpdate) SetNillableOwnerID(id *string) *FileUpdate {
	if id != nil {
		fu = fu.SetOwnerID(*id)
	}
	return fu
}

// SetOwner sets the owner edge to User.
func (fu *FileUpdate) SetOwner(u *User) *FileUpdate {
	return fu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (fu *FileUpdate) ClearOwner() *FileUpdate {
	fu.clearedOwner = true
	return fu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	if fu.size != nil {
		if err := file.SizeValidator(*fu.size); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	if len(fu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
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
	if fu.user != nil {
		update = true
		builder.Set(file.FieldUser, *fu.user)
	}
	if fu.group != nil {
		update = true
		builder.Set(file.FieldGroup, *fu.group)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if fu.clearedOwner {
		query, args := sql.Update(file.OwnerTable).
			SetNull(file.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(fu.owner) > 0 {
		for eid := range fu.owner {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := sql.Update(file.OwnerTable).
				Set(file.OwnerColumn, eid).
				Where(sql.InInts(file.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
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
		rv  = v.Clone()
		trs []*dsl.Traversal
	)
	if fu.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fu.size)
	}
	if fu.name != nil {
		v.Property(dsl.Single, file.FieldName, *fu.name)
	}
	if fu.user != nil {
		v.Property(dsl.Single, file.FieldUser, *fu.user)
	}
	if fu.group != nil {
		v.Property(dsl.Single, file.FieldGroup, *fu.group)
	}
	if fu.clearedOwner {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range fu.owner {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	id           string
	size         *int
	name         *string
	user         *string
	group        *string
	owner        map[string]struct{}
	clearedOwner bool
}

// SetSize sets the size field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.size = &i
	return fuo
}

// SetNillableSize sets the size field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// SetName sets the name field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.name = &s
	return fuo
}

// SetUser sets the user field.
func (fuo *FileUpdateOne) SetUser(s string) *FileUpdateOne {
	fuo.user = &s
	return fuo
}

// SetNillableUser sets the user field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUser(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetUser(*s)
	}
	return fuo
}

// SetGroup sets the group field.
func (fuo *FileUpdateOne) SetGroup(s string) *FileUpdateOne {
	fuo.group = &s
	return fuo
}

// SetNillableGroup sets the group field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableGroup(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetGroup(*s)
	}
	return fuo
}

// SetOwnerID sets the owner edge to User by id.
func (fuo *FileUpdateOne) SetOwnerID(id string) *FileUpdateOne {
	if fuo.owner == nil {
		fuo.owner = make(map[string]struct{})
	}
	fuo.owner[id] = struct{}{}
	return fuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableOwnerID(id *string) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetOwnerID(*id)
	}
	return fuo
}

// SetOwner sets the owner edge to User.
func (fuo *FileUpdateOne) SetOwner(u *User) *FileUpdateOne {
	return fuo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (fuo *FileUpdateOne) ClearOwner() *FileUpdateOne {
	fuo.clearedOwner = true
	return fuo
}

// Save executes the query and returns the updated entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	if fuo.size != nil {
		if err := file.SizeValidator(*fuo.size); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"size\": %v", err)
		}
	}
	if len(fuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
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
	if fuo.user != nil {
		update = true
		builder.Set(file.FieldUser, *fuo.user)
		f.User = fuo.user
	}
	if fuo.group != nil {
		update = true
		builder.Set(file.FieldGroup, *fuo.group)
		f.Group = *fuo.group
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if fuo.clearedOwner {
		query, args := sql.Update(file.OwnerTable).
			SetNull(file.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(fuo.owner) > 0 {
		for eid := range fuo.owner {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := sql.Update(file.OwnerTable).
				Set(file.OwnerColumn, eid).
				Where(sql.InInts(file.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
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
		rv  = v.Clone()
		trs []*dsl.Traversal
	)
	if fuo.size != nil {
		v.Property(dsl.Single, file.FieldSize, *fuo.size)
	}
	if fuo.name != nil {
		v.Property(dsl.Single, file.FieldName, *fuo.name)
	}
	if fuo.user != nil {
		v.Property(dsl.Single, file.FieldUser, *fuo.user)
	}
	if fuo.group != nil {
		v.Property(dsl.Single, file.FieldGroup, *fuo.group)
	}
	if fuo.clearedOwner {
		tr := rv.Clone().InE(user.FilesLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range fuo.owner {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
