// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
	"github.com/facebookincubator/ent/schema/field"
)

// StreetCreate is the builder for creating a Street entity.
type StreetCreate struct {
	config
	name *string
	city map[int]struct{}
}

// SetName sets the name field.
func (sc *StreetCreate) SetName(s string) *StreetCreate {
	sc.name = &s
	return sc
}

// SetCityID sets the city edge to City by id.
func (sc *StreetCreate) SetCityID(id int) *StreetCreate {
	if sc.city == nil {
		sc.city = make(map[int]struct{})
	}
	sc.city[id] = struct{}{}
	return sc
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (sc *StreetCreate) SetNillableCityID(id *int) *StreetCreate {
	if id != nil {
		sc = sc.SetCityID(*id)
	}
	return sc
}

// SetCity sets the city edge to City.
func (sc *StreetCreate) SetCity(c *City) *StreetCreate {
	return sc.SetCityID(c.ID)
}

// Save creates the Street in the database.
func (sc *StreetCreate) Save(ctx context.Context) (*Street, error) {
	if sc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(sc.city) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"city\"")
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreetCreate) SaveX(ctx context.Context) *Street {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StreetCreate) sqlSave(ctx context.Context) (*Street, error) {
	var (
		s    = &Street{config: sc.config}
		spec = &sqlgraph.CreateSpec{
			Table: street.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: street.FieldID,
			},
		}
	)
	if value := sc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: street.FieldName,
		})
		s.Name = *value
	}
	if nodes := sc.city; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}

	id := spec.ID.Value.(int64)
	s.ID = int(id)

	return s, nil
}
