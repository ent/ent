// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/compositetypes/ent/schema"
	"entgo.io/ent/examples/compositetypes/ent/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetAddress sets the "address" field.
func (m *UserCreate) SetAddress(v *schema.Address) *UserCreate {
	m.mutation.SetAddress(v)
	return m
}

// Mutation returns the UserMutation object of the builder.
func (m *UserCreate) Mutation() *UserMutation {
	return m.mutation
}

// Save creates the User in the database.
func (c *UserCreate) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, c.sqlSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *UserCreate) SaveX(ctx context.Context) *User {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *UserCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *UserCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *UserCreate) check() error {
	if _, ok := c.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "User.address"`)}
	}
	return nil
}

func (c *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	_node, _spec := c.createSpec()
	if err := sqlgraph.CreateNode(ctx, c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	c.mutation.id = &_node.ID
	c.mutation.done = true
	return _node, nil
}

func (c *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: c.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := c.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (c *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if c.err != nil {
		return nil, c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(c.builders))
	nodes := make([]*User, len(c.builders))
	mutators := make([]Mutator, len(c.builders))
	for i := range c.builders {
		func(i int, root context.Context) {
			builder := c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (c *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *UserCreateBulk) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}
