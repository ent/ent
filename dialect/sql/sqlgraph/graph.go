// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// sqlgraph provides graph abstraction capabilities on top
// of sql-based databases for ent codegen.
package sqlgraph

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"sort"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
)

// Rel is a relation type of an edge.
type Rel int

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

// A ConstraintError represents an error from mutation that violates a specific constraint.
type ConstraintError struct {
	msg string
}

func (e ConstraintError) Error() string { return e.msg }

// A Step provides a path-step information to the traversal functions.
type Step struct {
	// From is the source of the step.
	From struct {
		// V can be either one vertex or set of vertices.
		// It can be a pre-processed step (sql.Query) or a simple Go type (integer or string).
		V interface{}
		// Table holds the table name of V (from).
		Table string
		// Column to join with. Usually the "id" column.
		Column string
	}
	// Edge holds the edge information for getting the neighbors.
	Edge struct {
		// Rel of the edge.
		Rel Rel
		// Table name of where this edge columns reside.
		Table string
		// Columns of the edge.
		// In O2O and M2O, it holds the foreign-key column. Hence, len == 1.
		// In M2M, it holds the primary-key columns of the join table. Hence, len == 2.
		Columns []string
		// Inverse indicates if the edge is an inverse edge.
		Inverse bool
	}
	// To is the dest of the path (the neighbors).
	To struct {
		// Table holds the table name of the neighbors (to).
		Table string
		// Column to join with. Usually the "id" column.
		Column string
	}
}

// StepOption allows configuring Steps using functional options.
type StepOption func(*Step)

// From sets the source of the step.
func From(table, column string, v ...interface{}) StepOption {
	return func(s *Step) {
		s.From.Table = table
		s.From.Column = column
		if len(v) > 0 {
			s.From.V = v[0]
		}
	}
}

// To sets the destination of the step.
func To(table, column string) StepOption {
	return func(s *Step) {
		s.To.Table = table
		s.To.Column = column
	}
}

// Edge sets the edge info for getting the neighbors.
func Edge(rel Rel, inverse bool, table string, columns ...string) StepOption {
	return func(s *Step) {
		s.Edge.Rel = rel
		s.Edge.Table = table
		s.Edge.Columns = columns
		s.Edge.Inverse = inverse
	}
}

// NewStep gets list of options and returns a configured step.
//
//	NewStep(
//		From("table", "pk", V),
//		To("table", "pk"),
//		Edge("name", O2M, "fk"),
//	)
//
func NewStep(opts ...StepOption) *Step {
	s := &Step{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Neighbors returns a Selector for evaluating the path-step
// and getting the neighbors of one vertex.
func Neighbors(dialect string, s *Step) (q *sql.Selector) {
	builder := sql.Dialect(dialect)
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1, pk2 := s.Edge.Columns[1], s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1, pk2 = pk2, pk1
		}
		to := builder.Table(s.To.Table)
		join := builder.Table(s.Edge.Table)
		match := builder.Select(join.C(pk1)).
			From(join).
			Where(sql.EQ(join.C(pk2), s.From.V))
		q = builder.Select().
			From(to).
			Join(match).
			On(to.C(s.To.Column), match.C(pk1))
	case r == M2O || (r == O2O && s.Edge.Inverse):
		t1 := builder.Table(s.To.Table)
		t2 := builder.Select(s.Edge.Columns[0]).
			From(builder.Table(s.Edge.Table)).
			Where(sql.EQ(s.From.Column, s.From.V))
		q = builder.Select().
			From(t1).
			Join(t2).
			On(t1.C(s.From.Column), t2.C(s.Edge.Columns[0]))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		q = builder.Select().
			From(builder.Table(s.To.Table)).
			Where(sql.EQ(s.Edge.Columns[0], s.From.V))
	}
	return q
}

// SetNeighbors returns a Selector for evaluating the path-step
// and getting the neighbors of set of vertices.
func SetNeighbors(dialect string, s *Step) (q *sql.Selector) {
	set := s.From.V.(*sql.Selector)
	builder := sql.Dialect(dialect)
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1, pk2 := s.Edge.Columns[1], s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1, pk2 = pk2, pk1
		}
		to := builder.Table(s.To.Table)
		set.Select(set.C(s.From.Column))
		join := builder.Table(s.Edge.Table)
		match := builder.Select(join.C(pk1)).
			From(join).
			Join(set).
			On(join.C(pk2), set.C(s.From.Column))
		q = builder.Select().
			From(to).
			Join(match).
			On(to.C(s.To.Column), match.C(pk1))
	case r == M2O || (r == O2O && s.Edge.Inverse):
		t1 := builder.Table(s.To.Table)
		set.Select(set.C(s.Edge.Columns[0]))
		q = builder.Select().
			From(t1).
			Join(set).
			On(t1.C(s.To.Column), set.C(s.Edge.Columns[0]))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		t1 := builder.Table(s.To.Table)
		set.Select(set.C(s.From.Column))
		q = builder.Select().
			From(t1).
			Join(set).
			On(t1.C(s.Edge.Columns[0]), set.C(s.From.Column))
	}
	return q
}

// HasNeighbors applies on the given Selector a neighbors check.
func HasNeighbors(q *sql.Selector, s *Step) {
	builder := sql.Dialect(q.Dialect())
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1 := s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1 = s.Edge.Columns[1]
		}
		from := q.Table()
		join := builder.Table(s.Edge.Table)
		q.Where(
			sql.In(
				from.C(s.From.Column),
				builder.Select(join.C(pk1)).From(join),
			),
		)
	case r == M2O || (r == O2O && s.Edge.Inverse):
		from := q.Table()
		q.Where(sql.NotNull(from.C(s.Edge.Columns[0])))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		from := q.Table()
		to := builder.Table(s.Edge.Table)
		q.Where(
			sql.In(
				from.C(s.From.Column),
				builder.Select(to.C(s.Edge.Columns[0])).
					From(to).
					Where(sql.NotNull(to.C(s.Edge.Columns[0]))),
			),
		)
	}
}

// HasNeighborsWith applies on the given Selector a neighbors check.
// The given predicate applies its filtering on the selector.
func HasNeighborsWith(q *sql.Selector, s *Step, pred func(*sql.Selector)) {
	builder := sql.Dialect(q.Dialect())
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1, pk2 := s.Edge.Columns[1], s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1, pk2 = pk2, pk1
		}
		from := q.Table()
		to := builder.Table(s.To.Table)
		edge := builder.Table(s.Edge.Table)
		join := builder.Select(edge.C(pk2)).
			From(edge).
			Join(to).
			On(edge.C(pk1), to.C(s.To.Column))
		matches := builder.Select().From(to)
		pred(matches)
		join.FromSelect(matches)
		q.Where(sql.In(from.C(s.From.Column), join))
	case r == M2O || (r == O2O && s.Edge.Inverse):
		from := q.Table()
		to := builder.Table(s.To.Table)
		matches := builder.Select(to.C(s.To.Column)).
			From(to)
		pred(matches)
		q.Where(sql.In(from.C(s.Edge.Columns[0]), matches))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		from := q.Table()
		to := builder.Table(s.Edge.Table)
		matches := builder.Select(to.C(s.Edge.Columns[0])).
			From(to)
		pred(matches)
		q.Where(sql.In(from.C(s.From.Column), matches))
	}
}

type (
	// FieldSpec holds the information for updating a field
	// column in the database.
	FieldSpec struct {
		Column string
		Type   field.Type
		Value  driver.Value // value to be stored.
	}

	// EdgeTarget holds the information for the target nodes
	// of an edge.
	EdgeTarget struct {
		Nodes  []driver.Value
		IDSpec *FieldSpec
	}

	// EdgeSpec holds the information for updating a field
	// column in the database.
	EdgeSpec struct {
		Rel     Rel
		Inverse bool
		Table   string
		Columns []string
		Bidi    bool        // bidirectional edge.
		Target  *EdgeTarget // target nodes.
	}

	// EdgeSpecs used for perform common operations on list of edges.
	EdgeSpecs []*EdgeSpec

	// NodeSpec defines the information for querying and
	// decoding nodes in the graph.
	NodeSpec struct {
		Table   string
		Columns []string
		ID      *FieldSpec
	}
)

// CreateSpec holds the information for creating
// a node in the graph.
type CreateSpec struct {
	Table  string
	ID     *FieldSpec
	Fields []*FieldSpec
	Edges  []*EdgeSpec
}

// CreateNode applies the CreateSpec on the graph.
func CreateNode(ctx context.Context, drv dialect.Driver, spec *CreateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx, builder: sql.Dialect(drv.Dialect())}
	cr := &creator{CreateSpec: spec, graph: gr}
	if err := cr.node(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

type (
	// EdgeMut defines edge mutations.
	EdgeMut struct {
		Add   []*EdgeSpec
		Clear []*EdgeSpec
	}

	// FieldMut defines field mutations.
	FieldMut struct {
		Set   []*FieldSpec // field = ?
		Add   []*FieldSpec // field = field + ?
		Clear []*FieldSpec // field = NULL
	}

	// UpdateSpec holds the information for updating one
	// or more nodes in the graph.
	UpdateSpec struct {
		Node      *NodeSpec
		Edges     EdgeMut
		Fields    FieldMut
		Predicate func(*sql.Selector)

		ScanValues []interface{}
		Assign     func(...interface{}) error
	}
)

// UpdateNode applies the UpdateSpec on one node in the graph.
func UpdateNode(ctx context.Context, drv dialect.Driver, spec *UpdateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx, builder: sql.Dialect(drv.Dialect())}
	cr := &updater{UpdateSpec: spec, graph: gr}
	if err := cr.node(ctx, tx); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

// UpdateNodes applies the UpdateSpec on a set of nodes in the graph.
func UpdateNodes(ctx context.Context, drv dialect.Driver, spec *UpdateSpec) (int, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return 0, err
	}
	gr := graph{tx: tx, builder: sql.Dialect(drv.Dialect())}
	cr := &updater{UpdateSpec: spec, graph: gr}
	affected, err := cr.nodes(ctx, tx)
	if err != nil {
		return 0, rollback(tx, err)
	}
	return affected, tx.Commit()
}

// DeleteSpec holds the information for delete one
// or more nodes in the graph.
type DeleteSpec struct {
	Node      *NodeSpec
	Predicate func(*sql.Selector)
}

// DeleteNodes applies the DeleteSpec on the graph.
func DeleteNodes(ctx context.Context, drv dialect.Driver, spec *DeleteSpec) (int, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Dialect(drv.Dialect())
	)
	selector := builder.Select().
		From(builder.Table(spec.Node.Table))
	if pred := spec.Predicate; pred != nil {
		pred(selector)
	}
	query, args := builder.Delete(spec.Node.Table).FromSelect(selector).Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return 0, rollback(tx, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, rollback(tx, err)
	}
	return int(affected), tx.Commit()
}

// QuerySpec holds the information for querying
// nodes in the graph.
type QuerySpec struct {
	Node *NodeSpec     // Nodes info.
	From *sql.Selector // Optional query source (from path).

	Limit     int
	Offset    int
	Unique    bool
	Order     func(*sql.Selector)
	Predicate func(*sql.Selector)

	ScanValues func() []interface{}
	Assign     func(...interface{}) error
}

// QueryNodes query the nodes in the graph query and scans them to the given values.
func QueryNodes(ctx context.Context, drv dialect.Driver, spec *QuerySpec) error {
	builder := sql.Dialect(drv.Dialect())
	qr := &query{graph: graph{builder: builder}, QuerySpec: spec}
	return qr.nodes(ctx, drv)
}

// CountNodes counts the nodes in the given graph query.
func CountNodes(ctx context.Context, drv dialect.Driver, spec *QuerySpec) (int, error) {
	builder := sql.Dialect(drv.Dialect())
	qr := &query{graph: graph{builder: builder}, QuerySpec: spec}
	return qr.count(ctx, drv)
}

type query struct {
	graph
	*QuerySpec
}

func (q *query) nodes(ctx context.Context, drv dialect.Driver) error {
	rows := &sql.Rows{}
	query, args := q.selector().Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := q.ScanValues()
		if err := rows.Scan(values...); err != nil {
			return err
		}
		if err := q.Assign(values...); err != nil {
			return err
		}
	}
	return nil
}

func (q *query) count(ctx context.Context, drv dialect.Driver) (int, error) {
	rows := &sql.Rows{}
	selector := q.selector()
	selector.Count(selector.C(q.Node.ID.Column))
	if q.Unique {
		selector.SetDistinct(false)
		selector.Count(sql.Distinct(selector.C(q.Node.ID.Column)))
	}
	query, args := selector.Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	return sql.ScanInt(rows)
}

func (q *query) selector() *sql.Selector {
	selector := q.builder.Select().From(q.builder.Table(q.Node.Table))
	if q.From != nil {
		selector = q.From
	}
	selector.Select(selector.Columns(q.Node.Columns...)...)
	if pred := q.Predicate; pred != nil {
		pred(selector)
	}
	if order := q.Order; order != nil {
		order(selector)
	}
	if q.Offset != 0 {
		// Limit is mandatory for the offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(q.Offset).Limit(math.MaxInt32)
	}
	if q.Limit != 0 {
		selector.Limit(q.Limit)
	}
	if q.Unique {
		selector.Distinct()
	}
	return selector
}

type updater struct {
	graph
	*UpdateSpec
}

func (u *updater) node(ctx context.Context, tx dialect.ExecQuerier) error {
	var (
		// id holds the PK of the node used for linking
		// it with the other nodes.
		id         = u.Node.ID.Value
		addEdges   = EdgeSpecs(u.Edges.Add).GroupRel()
		clearEdges = EdgeSpecs(u.Edges.Clear).GroupRel()
	)
	update := u.builder.Update(u.Node.Table).Where(sql.EQ(u.Node.ID.Column, id))
	if err := u.setTableColumns(update, addEdges, clearEdges); err != nil {
		return err
	}
	if !update.Empty() {
		var res sql.Result
		query, args := update.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return err
		}
	}
	if err := u.setExternalEdges(ctx, []driver.Value{id}, addEdges, clearEdges); err != nil {
		return err
	}
	selector := u.builder.Select(u.Node.Columns...).
		From(u.builder.Table(u.Node.Table)).
		Where(sql.EQ(u.Node.ID.Column, u.Node.ID.Value))
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tx.Query(ctx, query, args, rows); err != nil {
		return err
	}
	return u.scan(rows)
}

func (u *updater) nodes(ctx context.Context, tx dialect.ExecQuerier) (int, error) {
	var (
		ids        []driver.Value
		addEdges   = EdgeSpecs(u.Edges.Add).GroupRel()
		clearEdges = EdgeSpecs(u.Edges.Clear).GroupRel()
	)
	selector := u.builder.Select(u.Node.ID.Column).
		From(u.builder.Table(u.Node.Table))
	if pred := u.Predicate; pred != nil {
		pred(selector)
	}
	query, args := selector.Query()
	rows := &sql.Rows{}
	if err := u.tx.Query(ctx, query, args, rows); err != nil {
		return 0, fmt.Errorf("querying table %s: %v", u.Node.Table, err)
	}
	defer rows.Close()
	if err := sql.ScanSlice(rows, &ids); err != nil {
		return 0, fmt.Errorf("scan node ids: %v", err)
	}
	if err := rows.Close(); err != nil {
		return 0, err
	}
	if len(ids) == 0 {
		return 0, nil
	}
	update := u.builder.Update(u.Node.Table).Where(matchID(u.Node.ID.Column, ids))
	if err := u.setTableColumns(update, addEdges, clearEdges); err != nil {
		return 0, err
	}
	if !update.Empty() {
		var res sql.Result
		query, args := update.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, err
		}
	}
	if err := u.setExternalEdges(ctx, ids, addEdges, clearEdges); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (u *updater) setExternalEdges(ctx context.Context, ids []driver.Value, addEdges, clearEdges map[Rel][]*EdgeSpec) error {
	if err := u.graph.clearM2MEdges(ctx, ids, clearEdges[M2M]); err != nil {
		return err
	}
	if err := u.graph.addM2MEdges(ctx, ids, addEdges[M2M]); err != nil {
		return err
	}
	if err := u.graph.clearFKEdges(ctx, ids, append(clearEdges[O2M], clearEdges[O2O]...)); err != nil {
		return err
	}
	if err := u.graph.addFKEdges(ctx, ids, append(addEdges[O2M], addEdges[O2O]...)); err != nil {
		return err
	}
	return nil
}

// setTableColumns sets the table columns and foreign_keys used in insert.
func (u *updater) setTableColumns(update *sql.UpdateBuilder, addEdges, clearEdges map[Rel][]*EdgeSpec) error {
	// Avoid multiple assignments to the same column.
	setEdges := make(map[string]bool)
	for _, e := range addEdges[M2O] {
		setEdges[e.Columns[0]] = true
	}
	for _, e := range addEdges[O2O] {
		if e.Inverse || e.Bidi {
			setEdges[e.Columns[0]] = true
		}
	}
	for _, fi := range u.Fields.Clear {
		update.SetNull(fi.Column)
	}
	for _, e := range clearEdges[M2O] {
		if col := e.Columns[0]; !setEdges[col] {
			update.SetNull(col)
		}
	}
	for _, e := range clearEdges[O2O] {
		col := e.Columns[0]
		if (e.Inverse || e.Bidi) && !setEdges[col] {
			update.SetNull(col)
		}
	}
	err := setTableColumns(u.Fields.Set, addEdges, func(column string, value driver.Value) {
		update.Set(column, value)
	})
	if err != nil {
		return err
	}
	for _, fi := range u.Fields.Add {
		update.Add(fi.Column, fi.Value)
	}
	return nil
}

func (u *updater) scan(rows *sql.Rows) error {
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("record with id %v not found in table %s", u.Node.ID.Value, u.Node.Table)
	}
	if err := rows.Scan(u.ScanValues...); err != nil {
		return fmt.Errorf("failed scanning rows: %v", err)
	}
	return u.Assign(u.ScanValues...)
}

type creator struct {
	graph
	*CreateSpec
}

func (c *creator) node(ctx context.Context, tx dialect.ExecQuerier) error {
	var (
		edges  = EdgeSpecs(c.Edges).GroupRel()
		insert = c.builder.Insert(c.Table).Default()
	)
	// Set and create the node.
	if err := c.setTableColumns(insert, edges); err != nil {
		return err
	}
	if err := c.insert(ctx, tx, insert); err != nil {
		return fmt.Errorf("insert node to table %q: %v", c.Table, err)
	}
	if err := c.graph.addM2MEdges(ctx, []driver.Value{c.ID.Value}, edges[M2M]); err != nil {
		return err
	}
	if err := c.graph.addFKEdges(ctx, []driver.Value{c.ID.Value}, append(edges[O2M], edges[O2O]...)); err != nil {
		return err
	}
	return nil
}

// setTableColumns sets the table columns and foreign_keys used in insert.
func (c *creator) setTableColumns(insert *sql.InsertBuilder, edges map[Rel][]*EdgeSpec) error {
	err := setTableColumns(c.Fields, edges, func(column string, value driver.Value) {
		insert.Set(column, value)
	})
	return err
}

// insert inserts the node to its table and sets its ID if it wasn't provided by the user.
func (c *creator) insert(ctx context.Context, tx dialect.ExecQuerier, insert *sql.InsertBuilder) error {
	var res sql.Result
	// If the id field was provided by the user.
	if c.ID.Value != nil {
		insert.Set(c.ID.Column, c.ID.Value)
		query, args := insert.Query()
		return tx.Exec(ctx, query, args, &res)
	}
	id, err := insertLastID(ctx, tx, insert.Returning(c.ID.Column))
	if err != nil {
		return err
	}
	c.ID.Value = id
	return nil
}

// GroupRel groups edges by their relation type.
func (es EdgeSpecs) GroupRel() map[Rel][]*EdgeSpec {
	edges := make(map[Rel][]*EdgeSpec)
	for _, edge := range es {
		edges[edge.Rel] = append(edges[edge.Rel], edge)
	}
	return edges
}

// GroupTable groups edges by their table name.
func (es EdgeSpecs) GroupTable() map[string][]*EdgeSpec {
	edges := make(map[string][]*EdgeSpec)
	for _, edge := range es {
		edges[edge.Table] = append(edges[edge.Table], edge)
	}
	return edges
}

// The common operations shared between the different builders.
//
// M2M edges reside in join tables and require INSERT and DELETE
// queries for adding or removing edges respectively.
//
// O2M and non-inverse O2O edges also reside in external tables,
// but use UPDATE queries (fk = ?, fk = NULL).
type graph struct {
	tx      dialect.ExecQuerier
	builder *sql.DialectBuilder
}

func (g *graph) clearM2MEdges(ctx context.Context, ids []driver.Value, edges EdgeSpecs) error {
	var (
		res sql.Result
		// Delete all M2M edges from the same type at once.
		// The EdgeSpec is the same for all members in a group.
		tables = edges.GroupTable()
	)
	for _, table := range sortedKeys(tables) {
		edges := tables[table]
		preds := make([]*sql.Predicate, 0, len(edges))
		for _, edge := range edges {
			pk1, pk2 := ids, edge.Target.Nodes
			if edge.Inverse {
				pk1, pk2 = pk2, pk1
			}
			preds = append(preds, matchIDs(edge.Columns[0], pk1, edge.Columns[1], pk2))
			if edge.Bidi {
				preds = append(preds, matchIDs(edge.Columns[0], pk2, edge.Columns[1], pk1))
			}
		}
		query, args := g.builder.Delete(table).Where(sql.Or(preds...)).Query()
		if err := g.tx.Exec(ctx, query, args, &res); err != nil {
			return fmt.Errorf("remove m2m edge for table %s: %v", table, err)
		}
	}
	return nil
}

func (g *graph) addM2MEdges(ctx context.Context, ids []driver.Value, edges EdgeSpecs) error {
	var (
		res sql.Result
		// Insert all M2M edges from the same type at once.
		// The EdgeSpec is the same for all members in a group.
		tables = edges.GroupTable()
	)
	for _, table := range sortedKeys(tables) {
		edges := tables[table]
		insert := g.builder.Insert(table).Columns(edges[0].Columns...)
		for _, edge := range edges {
			pk1, pk2 := ids, edge.Target.Nodes
			if edge.Inverse {
				pk1, pk2 = pk2, pk1
			}
			for _, pair := range product(pk1, pk2) {
				insert.Values(pair[0], pair[1])
				if edge.Bidi {
					insert.Values(pair[1], pair[0])
				}
			}
		}
		query, args := insert.Query()
		if err := g.tx.Exec(ctx, query, args, &res); err != nil {
			return fmt.Errorf("add m2m edge for table %s: %v", table, err)
		}
	}
	return nil
}

func (g *graph) clearFKEdges(ctx context.Context, ids []driver.Value, edges []*EdgeSpec) error {
	for _, edge := range edges {
		if edge.Rel == O2O && edge.Inverse {
			continue
		}
		// O2O relations can be cleared without
		// passing the target ids.
		pred := matchID(edge.Columns[0], ids)
		if nodes := edge.Target.Nodes; len(nodes) > 0 {
			pred = matchIDs(edge.Target.IDSpec.Column, edge.Target.Nodes, edge.Columns[0], ids)
		}
		query, args := g.builder.Update(edge.Table).
			SetNull(edge.Columns[0]).
			Where(pred).
			Query()
		var res sql.Result
		if err := g.tx.Exec(ctx, query, args, &res); err != nil {
			return fmt.Errorf("add %s edge for table %s: %v", edge.Rel, edge.Table, err)
		}
	}
	return nil
}

func (g *graph) addFKEdges(ctx context.Context, ids []driver.Value, edges []*EdgeSpec) error {
	id := ids[0]
	if len(ids) > 1 && len(edges) != 0 {
		// O2M and O2O edges are defined by a FK in the "other" table.
		// Therefore, ids[i+1] will override ids[i] which is invalid.
		return fmt.Errorf("unable to link FK edge to more than 1 node: %v", ids)
	}
	for _, edge := range edges {
		if edge.Rel == O2O && edge.Inverse {
			continue
		}
		p := sql.EQ(edge.Target.IDSpec.Column, edge.Target.Nodes[0])
		// Use "IN" predicate instead of list of "OR"
		// in case of more than on nodes to connect.
		if len(edge.Target.Nodes) > 1 {
			p = sql.InValues(edge.Target.IDSpec.Column, edge.Target.Nodes...)
		}
		query, args := g.builder.Update(edge.Table).
			Set(edge.Columns[0], id).
			Where(sql.And(p, sql.IsNull(edge.Columns[0]))).
			Query()
		var res sql.Result
		if err := g.tx.Exec(ctx, query, args, &res); err != nil {
			return fmt.Errorf("add %s edge for table %s: %v", edge.Rel, edge.Table, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if ids := edge.Target.Nodes; int(affected) < len(ids) {
			return &ConstraintError{msg: fmt.Sprintf("one of %v is already connected to a different %s", ids, edge.Columns[0])}
		}
	}
	return nil
}

// setTableColumns is shared between updater and creator.
func setTableColumns(fields []*FieldSpec, edges map[Rel][]*EdgeSpec, set func(string, driver.Value)) (err error) {
	for _, fi := range fields {
		value := fi.Value
		if fi.Type == field.TypeJSON {
			if value, err = json.Marshal(value); err != nil {
				return fmt.Errorf("marshal value for column %s: %v", fi.Column, err)
			}
		}
		set(fi.Column, value)
	}
	for _, e := range edges[M2O] {
		set(e.Columns[0], e.Target.Nodes[0])
	}
	for _, e := range edges[O2O] {
		if e.Inverse || e.Bidi {
			set(e.Columns[0], e.Target.Nodes[0])
		}
	}
	return nil
}

// insertLastID invokes the insert query on the transaction and returns the LastInsertID.
func insertLastID(ctx context.Context, tx dialect.ExecQuerier, insert *sql.InsertBuilder) (int64, error) {
	query, args := insert.Query()
	// PostgreSQL does not support the LastInsertId() method of sql.Result
	// on Exec, and should be extracted manually using the `RETURNING` clause.
	if insert.Dialect() == dialect.Postgres {
		rows := &sql.Rows{}
		if err := tx.Query(ctx, query, args, rows); err != nil {
			return 0, err
		}
		defer rows.Close()
		return sql.ScanInt64(rows)
	}
	// MySQL, SQLite, etc.
	var res sql.Result
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}

func sortedKeys(m map[string][]*EdgeSpec) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func matchID(column string, pk []driver.Value) *sql.Predicate {
	if len(pk) > 1 {
		return sql.InValues(column, pk...)
	}
	return sql.EQ(column, pk[0])
}

func matchIDs(column1 string, pk1 []driver.Value, column2 string, pk2 []driver.Value) *sql.Predicate {
	p := matchID(column1, pk1)
	if len(pk2) > 1 {
		// Use "IN" predicate instead of list of "OR"
		// in case of more than on nodes to connect.
		return p.And().InValues(column2, pk2...)
	}
	return p.And().EQ(column2, pk2[0])
}

// cartesian product of 2 id sets.
func product(a, b []driver.Value) [][2]driver.Value {
	c := make([][2]driver.Value, 0, len(a)*len(b))
	for i := range a {
		for j := range b {
			c = append(c, [2]driver.Value{a[i], b[j]})
		}
	}
	return c
}
