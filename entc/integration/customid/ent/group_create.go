// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/group"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	id    *string
	users map[int]struct{}
}

// SetID sets the id field.
func (gc *GroupCreate) SetID(s string) *GroupCreate {
	gc.id = &s
	return gc
}

// AddUserIDs adds the users edge to User by ids.
func (gc *GroupCreate) AddUserIDs(ids ...int) *GroupCreate {
	if gc.users == nil {
		gc.users = make(map[int]struct{})
	}
	for i := range ids {
		gc.users[ids[i]] = struct{}{}
	}
	return gc
}

// AddUsers adds the users edges to User.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	return gc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(gc.driver.Dialect())
		gr      = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(group.Table).Default()
	id, err := insertLastID(ctx, tx, insert.Returning(group.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = strconv.FormatInt(id, 10)
	if len(gc.users) > 0 {
		for eid := range gc.users {

			query, args := builder.Insert(group.UsersTable).
				Columns(group.UsersPrimaryKey[0], group.UsersPrimaryKey[1]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
