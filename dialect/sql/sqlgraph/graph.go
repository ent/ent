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

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"
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
			On(t1.C(s.To.Column), t2.C(s.Edge.Columns[0]))
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

type (
	// CreateSpec holds the information for creating
	// a node in the graph.
	CreateSpec struct {
		Table  string
		ID     *FieldSpec
		Fields []*FieldSpec
		Edges  []*EdgeSpec
	}
	// BatchCreateSpec holds the information for creating
	// multiple nodes in the graph.
	BatchCreateSpec struct {
		Nodes []*CreateSpec
	}
)

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

// BatchCreate applies the BatchCreateSpec on the graph.
func BatchCreate(ctx context.Context, drv dialect.Driver, spec *BatchCreateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	gr := graph{tx: tx, builder: sql.Dialect(drv.Dialect())}
	cr := &creator{BatchCreateSpec: spec, graph: gr}
	if err := cr.nodes(ctx, tx); err != nil {
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

// NotFoundError returns when trying to update an
// entity and it was not found in the database.
type NotFoundError struct {
	table string
	id    driver.Value
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("record with id %v not found in table %s", e.id, e.table)
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

// QueryNodes queries the nodes in the graph query and scans them to the given values.
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

// EdgeQuerySpec holds the information for querying
// edges in the graph.
type EdgeQuerySpec struct {
	Edge       *EdgeSpec
	Predicate  func(*sql.Selector)
	ScanValues func() [2]interface{}
	Assign     func(out, in interface{}) error
}

// QueryEdges queries the edges in the graph and scans the result with the given dest function.
func QueryEdges(ctx context.Context, drv dialect.Driver, spec *EdgeQuerySpec) error {
	if len(spec.Edge.Columns) != 2 {
		return fmt.Errorf("sqlgraph: edge query requires 2 columns (out, in)")
	}
	out, in := spec.Edge.Columns[0], spec.Edge.Columns[1]
	if spec.Edge.Inverse {
		out, in = in, out
	}
	selector := sql.Dialect(drv.Dialect()).
		Select(out, in).
		From(sql.Table(spec.Edge.Table))
	if p := spec.Predicate; p != nil {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		values := spec.ScanValues()
		if err := rows.Scan(values[0], values[1]); err != nil {
			return err
		}
		if err := spec.Assign(values[0], values[1]); err != nil {
			return err
		}
	}
	return rows.Err()
}

type query struct {
	graph
	*QuerySpec
}

func (q *query) nodes(ctx context.Context, drv dialect.Driver) error {
	rows := &sql.Rows{}
	selector, err := q.selector()
	if err != nil {
		return err
	}
	query, args := selector.Query()
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
	return rows.Err()
}

func (q *query) count(ctx context.Context, drv dialect.Driver) (int, error) {
	rows := &sql.Rows{}
	selector, err := q.selector()
	if err != nil {
		return 0, err
	}
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

func (q *query) selector() (*sql.Selector, error) {
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
	if err := selector.Err(); err != nil {
		return nil, err
	}
	return selector, nil
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
		if err := rows.Err(); err != nil {
			return err
		}
		return &NotFoundError{table: u.Node.Table, id: u.Node.ID.Value}
	}
	if err := rows.Scan(u.ScanValues...); err != nil {
		return fmt.Errorf("failed scanning rows: %v", err)
	}
	return u.Assign(u.ScanValues...)
}

type creator struct {
	graph
	*CreateSpec
	*BatchCreateSpec
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

func (c *creator) nodes(ctx context.Context, tx dialect.ExecQuerier) error {
	if len(c.Nodes) == 0 {
		return nil
	}
	columns := make(map[string]struct{})
	values := make([]map[string]driver.Value, len(c.Nodes))
	for i, node := range c.Nodes {
		if i > 0 && node.Table != c.Nodes[i-1].Table {
			return fmt.Errorf("more than 1 table for batch insert: %q != %q", node.Table, c.Nodes[i-1].Table)
		}
		values[i] = make(map[string]driver.Value)
		if node.ID.Value != nil {
			columns[node.ID.Column] = struct{}{}
			values[i][node.ID.Column] = node.ID.Value
		}
		edges := EdgeSpecs(node.Edges).GroupRel()
		err := setTableColumns(node.Fields, edges, func(column string, value driver.Value) {
			columns[column] = struct{}{}
			values[i][column] = value
		})
		if err != nil {
			return err
		}
	}
	for column := range columns {
		for i := range values {
			switch _, exists := values[i][column]; {
			case column == c.Nodes[i].ID.Column && !exists:
				// If the ID value was provided to one of the nodes, it should be
				// provided to all others because this affects the way we calculate
				// their values in MySQL and SQLite dialects.
				return fmt.Errorf("incosistent id values for batch insert")
			case !exists:
				// Assign NULL values for empty placeholders.
				values[i][column] = nil
			}
		}
	}
	sorted := keys(columns)
	insert := c.builder.Insert(c.Nodes[0].Table).Default().Columns(sorted...)
	for i := range values {
		vs := make([]interface{}, len(sorted))
		for j, c := range sorted {
			vs[j] = values[i][c]
		}
		insert.Values(vs...)
	}
	if err := c.batchInsert(ctx, tx, insert); err != nil {
		return fmt.Errorf("insert nodes to table %q: %v", c.Nodes[0].Table, err)
	}
	if err := c.batchAddM2M(ctx, c.BatchCreateSpec); err != nil {
		return err
	}
	// FKs that exist in different tables can't be updated in batch (using the CASE
	// statement), because we rely on RowsAffected to check if the FK column is NULL.
	for _, node := range c.Nodes {
		edges := EdgeSpecs(node.Edges).GroupRel()
		if err := c.graph.addFKEdges(ctx, []driver.Value{node.ID.Value}, append(edges[O2M], edges[O2O]...)); err != nil {
			return err
		}
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

// batchInsert inserts a batch of nodes to their table and sets their ID if it wasn't provided by the user.
func (c *creator) batchInsert(ctx context.Context, tx dialect.ExecQuerier, insert *sql.InsertBuilder) error {
	ids, err := insertLastIDs(ctx, tx, insert.Returning(c.Nodes[0].ID.Column))
	if err != nil {
		return err
	}
	for i, node := range c.Nodes {
		node.ID.Value = ids[i]
	}
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

// FilterRel returns edges for the given relation type.
func (es EdgeSpecs) FilterRel(r Rel) EdgeSpecs {
	edges := make([]*EdgeSpec, 0, len(es))
	for _, edge := range es {
		if edge.Rel == r {
			edges = append(edges, edge)
		}
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
		// Remove all M2M edges from the same type at once.
		// The EdgeSpec is the same for all members in a group.
		tables = edges.GroupTable()
	)
	for _, table := range edgeKeys(tables) {
		edges := tables[table]
		preds := make([]*sql.Predicate, 0, len(edges))
		for _, edge := range edges {
			fromC, toC := edge.Columns[0], edge.Columns[1]
			if edge.Inverse {
				fromC, toC = toC, fromC
			}
			// If there are no specific edges (to target-nodes) to remove,
			// clear all edges that go out (or come in) from the nodes.
			if len(edge.Target.Nodes) == 0 {
				preds = append(preds, matchID(fromC, ids))
				if edge.Bidi {
					preds = append(preds, matchID(toC, ids))
				}
			} else {
				pk1, pk2 := ids, edge.Target.Nodes
				preds = append(preds, matchIDs(fromC, pk1, toC, pk2))
				if edge.Bidi {
					preds = append(preds, matchIDs(toC, pk1, fromC, pk2))
				}
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
	for _, table := range edgeKeys(tables) {
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

func (g *graph) batchAddM2M(ctx context.Context, spec *BatchCreateSpec) error {
	tables := make(map[string]*sql.InsertBuilder)
	for _, node := range spec.Nodes {
		edges := EdgeSpecs(node.Edges).FilterRel(M2M)
		for t, edges := range edges.GroupTable() {
			insert, ok := tables[t]
			if !ok {
				insert = g.builder.Insert(t).Columns(edges[0].Columns...)
			}
			tables[t] = insert
			if len(edges) != 1 {
				return fmt.Errorf("expect exactly 1 edge-spec per table, but got %d", len(edges))
			}
			edge := edges[0]
			pk1, pk2 := []driver.Value{node.ID.Value}, edge.Target.Nodes
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
	}
	for _, table := range insertKeys(tables) {
		var (
			res         sql.Result
			query, args = tables[table].Query()
		)
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
		// Setting the FK value of the "other" table
		// without clearing it before, is not allowed.
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
			buf, err := json.Marshal(value)
			if err != nil {
				return fmt.Errorf("marshal value for column %s: %v", fi.Column, err)
			}
			// If the underlying driver does not support JSON types,
			// driver.DefaultParameterConverter will convert it to uint8.
			value = json.RawMessage(buf)
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

// insertLastIDs invokes the batch insert query on the transaction and returns the LastInsertID of all entities.
func insertLastIDs(ctx context.Context, tx dialect.ExecQuerier, insert *sql.InsertBuilder) (ids []int64, err error) {
	query, args := insert.Query()
	// PostgreSQL does not support the LastInsertId() method of sql.Result
	// on Exec, and should be extracted manually using the `RETURNING` clause.
	if insert.Dialect() == dialect.Postgres {
		rows := &sql.Rows{}
		if err := tx.Query(ctx, query, args, rows); err != nil {
			return nil, err
		}
		defer rows.Close()
		return ids, sql.ScanSlice(rows, &ids)
	}
	// MySQL, SQLite, etc.
	var res sql.Result
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	ids = make([]int64, 0, affected)
	switch insert.Dialect() {
	case dialect.SQLite:
		id -= affected - 1
		fallthrough
	case dialect.MySQL:
		for i := int64(0); i < affected; i++ {
			ids = append(ids, id+i)
		}
	}
	return ids, nil
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}

func edgeKeys(m map[string][]*EdgeSpec) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func insertKeys(m map[string]*sql.InsertBuilder) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func keys(m map[string]struct{}) []string {
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
		return sql.And(p, sql.InValues(column2, pk2...))
	}
	return sql.And(p, sql.EQ(column2, pk2[0]))
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
