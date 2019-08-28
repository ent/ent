// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"

	"github.com/facebookincubator/ent/dialect/sql"
)

// CityCreate is the builder for creating a City entity.
type CityCreate struct {
	config
	name    *string
	streets map[int]struct{}
}

// SetName sets the name field.
func (cc *CityCreate) SetName(s string) *CityCreate {
	cc.name = &s
	return cc
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cc *CityCreate) AddStreetIDs(ids ...int) *CityCreate {
	if cc.streets == nil {
		cc.streets = make(map[int]struct{})
	}
	for i := range ids {
		cc.streets[ids[i]] = struct{}{}
	}
	return cc
}

// AddStreets adds the streets edges to Street.
func (cc *CityCreate) AddStreets(s ...*Street) *CityCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStreetIDs(ids...)
}

// Save creates the City in the database.
func (cc *CityCreate) Save(ctx context.Context) (*City, error) {
	if cc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CityCreate) SaveX(ctx context.Context) *City {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CityCreate) sqlSave(ctx context.Context) (*City, error) {
	var (
		res sql.Result
		c   = &City{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(city.Table).Default(cc.driver.Dialect())
	if cc.name != nil {
		builder.Set(city.FieldName, *cc.name)
		c.Name = *cc.name
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = int(id)
	if len(cc.streets) > 0 {
		p := sql.P()
		for eid := range cc.streets {
			p.Or().EQ(street.FieldID, eid)
		}
		query, args := sql.Update(city.StreetsTable).
			Set(city.StreetsColumn, id).
			Where(sql.And(p, sql.IsNull(city.StreetsColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(cc.streets) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"streets\" %v already connected to a different \"City\"", keys(cc.streets))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
