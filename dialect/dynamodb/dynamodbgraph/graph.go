// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package dynamodbgraph

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	sdk "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

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
		Rel        Rel
		Inverse    bool
		Table      string
		Attributes []string
		Bidi       bool        // bidirectional edge.
		Target     *EdgeTarget // target nodes.
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
		dynamodb.RootBuilder
	}

	creator struct {
		graph
		*CreateSpec
		data map[string]types.AttributeValue
	}
)

// CreateNode applies the CreateSpec on the graph.
func CreateNode(ctx context.Context, drv dialect.Driver, spec *CreateSpec) (err error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx, RootBuilder: dynamodb.RootBuilder{}}
	cr := &creator{
		CreateSpec: spec,
		graph:      gr,
		data:       make(map[string]types.AttributeValue),
	}
	if err = cr.node(ctx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// node is the controller to create a single node in the graph.
func (c *creator) node(ctx context.Context) (err error) {
	if err = c.insert(ctx); err != nil {
		return err
	}
	edges := EdgeSpecs(c.CreateSpec.Edges).GroupRel()
	if err = c.graph.addFKEdges(ctx, []interface{}{c.ID.Value}, append(edges[O2M], edges[O2O]...)); err != nil {
		return err
	}
	if err = c.graph.addM2MEdges(ctx, []interface{}{c.ID.Value}, edges[M2M]); err != nil {
		return err
	}
	return nil
}

// insert returns potential errors during process of marshaling CreateSpec
// to DynamoBD attributes and build steps in dynamodb.PutItemBuilder.
func (c *creator) insert(ctx context.Context) (err error) {
	edges, fields, putItemBuilder := EdgeSpecs(c.CreateSpec.Edges).GroupRel(), c.CreateSpec.Fields, c.PutItem(c.Table)
	// ID field is not included in CreateSpec.Fields
	if c.CreateSpec.ID != nil {
		fields = append(fields, c.CreateSpec.ID)
	}
	if err = setTableAttributes(fields, edges, c.data); err != nil {
		return err
	}
	putItemBuilder.SetItem(c.data)
	op, args := putItemBuilder.Op()
	return c.tx.Exec(ctx, op, args, nil)
}

// setTableAttributes is shared between updater and creator.
func setTableAttributes(fields []*FieldSpec, edges map[Rel][]*EdgeSpec, data map[string]types.AttributeValue) (err error) {
	for _, f := range fields {
		if data[f.Key], err = attributevalue.Marshal(f.Value); err != nil {
			return err
		}
	}
	for _, e := range edges[M2O] {
		if data[e.Attributes[0]], err = attributevalue.Marshal(e.Target.Nodes[0]); err != nil {
			return err
		}
	}
	for _, e := range edges[O2O] {
		if e.Inverse || e.Bidi {
			if data[e.Attributes[0]], err = attributevalue.Marshal(e.Target.Nodes[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *graph) addFKEdges(ctx context.Context, ids []interface{}, edges []*EdgeSpec) (err error) {
	if len(ids) > 1 && len(edges) != 0 {
		// O2M and O2O edges are defined by a FK in the "other" collection.
		// Therefore, ids[i+1] will override ids[i] which is invalid.
		return fmt.Errorf("unable to link FK edge to more than 1 node: %v", ids)
	}
	id := ids[0]
	for _, edge := range edges {
		if edge.Rel == O2O && edge.Inverse {
			continue
		}
		for _, n := range edge.Target.Nodes {
			keyVal, err := attributevalue.Marshal(n)
			if err != nil {
				return fmt.Errorf("key type not supported: %v has type %T", n, n)
			}
			query, err := g.Update(edge.Table).
				WithKey(edge.Target.IDSpec.Key, keyVal).
				Set(edge.Attributes[0], id).
				Where(dynamodb.NotExist(edge.Attributes[0])).
				Query(types.ReturnValueAllNew)
			if err != nil {
				return fmt.Errorf("build update query for table %s: %w", edge.Table, err)
			}
			op, input := query.Op()
			var output sdk.UpdateItemOutput
			if err := g.tx.Exec(ctx, op, input, &output); err != nil {
				return fmt.Errorf("add %s edge for table %s: %w", edge.Rel, edge.Table, err)
			}
		}
	}
	return nil
}

func (g *graph) addM2MEdges(ctx context.Context, ids []interface{}, edges []*EdgeSpec) (err error) {
	if len(edges) == 0 {
		return nil
	}
	batchWrite := dynamodb.BatchWriteItem()
	for _, e := range edges {
		m2mTable := e.Table
		fromIds, toIds := e.Target.Nodes, ids
		if e.Inverse {
			fromIds, toIds = toIds, fromIds
		}
		toAttr, fromAttr := e.Attributes[0], e.Attributes[1]
		for _, fromId := range fromIds {
			for _, toId := range toIds {
				data := make(map[string]types.AttributeValue)
				if data[toAttr], err = attributevalue.Marshal(toId); err != nil {
					return fmt.Errorf("add m2m edge: %w", err)
				}
				if data[fromAttr], err = attributevalue.Marshal(fromId); err != nil {
					return fmt.Errorf("add m2m edge: %w", err)
				}
				batchWrite.Append(m2mTable, g.PutItem(m2mTable).SetItem(data))
			}
		}
	}
	op, input := batchWrite.Op()
	var output sdk.BatchWriteItemOutput
	if err := g.tx.Exec(ctx, op, input, &output); err != nil {
		return fmt.Errorf("add m2m edge: %w", err)
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

// GroupRel groups edges by their relation type.
func (es EdgeSpecs) GroupRel() map[Rel][]*EdgeSpec {
	edges := make(map[Rel][]*EdgeSpec)
	for _, edge := range es {
		edges[edge.Rel] = append(edges[edge.Rel], edge)
	}
	return edges
}
