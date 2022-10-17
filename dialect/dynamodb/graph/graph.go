// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/dynamodb"
	"entgo.io/ent/schema/field"
)

type (
	// CreateSpec holds the information for creating
	// a node in the graph.
	CreateSpec struct {
		Table  string
		ID     *FieldSpec
		Fields []*FieldSpec
		Edges  []*EdgeSpec
	}

	// FieldSpec holds the information for updating a field in the database.
	FieldSpec struct {
		Key   string
		Type  field.Type
		Value interface{}
	}

	// EdgeTarget holds the information for the target nodes of an edge.
	EdgeTarget struct {
		Nodes  []interface{}
		IDSpec *FieldSpec
	}

	// EdgeSpec holds the information for updating a field in the database.
	EdgeSpec struct {
		Rel     Rel
		Inverse bool
		Table   string
		Keys    []string
		Bidi    bool        // bidirectional edge.
		Target  *EdgeTarget // target nodes.
	}

	// EdgeSpecs used for perform common operations on list of edges.
	EdgeSpecs []*EdgeSpec

	// NodeSpec defines the information for querying and
	// decoding nodes in the graph.
	NodeSpec struct {
		Table string
		Keys  []string
		ID    *FieldSpec
	}

	// Rel is a relation type of edge.
	Rel int
)

// Relation types.
const (
	_   Rel = iota // Unknown.
	O2O            // One to one / has one.
	O2M            // One to many / has many.
	M2O            // Many to one (inverse perspective for O2M).
	M2M            // Many to many.
)

// String returns the relation name.
func (r Rel) String() (s string) {
	switch r {
	case O2O:
		s = "O2O"
	case O2M:
		s = "O2M"
	case M2O:
		s = "M2O"
	case M2M:
		s = "M2M"
	default:
		s = "Unknown"
	}
	return s
}

type (
	graph struct {
		tx dialect.ExecQuerier
	}

	creator struct {
		graph
		*CreateSpec
	}
)

// CreateNode applies the CreateSpec on the graph.
func CreateNode(ctx context.Context, drv dialect.Driver, spec *CreateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx}
	cr := &creator{CreateSpec: spec, graph: gr}
	if err := cr.node(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func (c *creator) node(ctx context.Context, tx dialect.ExecQuerier) error {
	createTable := dynamodb.CreateTable(c.Table)
	var res dynamodb.CreateTableResult
	op, args := createTable.Op()
	if err := tx.Exec(ctx, op, args, &res); err != nil {
		return err
	}
	return nil
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}
