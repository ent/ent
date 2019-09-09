// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect"
)

// Querier wraps the basic Query method.
type Querier interface {
	// Query returns the query representation of the element and its arguments (if any).
	Query() (string, []interface{})
}

// Queries are list of queries join with space between them.
type Queries []Querier

// Query returns query representation of Queriers.
func (n Queries) Query() (string, []interface{}) {
	var b Builder
	for i := range n {
		if i > 0 {
			b.Pad()
		}
		query, args := n[i].Query()
		b.WriteString(query)
		b.args = append(b.args, args...)
	}
	return b.String(), b.args
}

// Builder is the base query builder for the sql dsl.
type Builder struct {
	bytes.Buffer
	args    []interface{}
	dialect string
}

// Append appends the given string as a quoted parameter
func (b *Builder) Append(s string) *Builder {
	switch {
	case len(s) == 0:
	case s != "*" && s[0] != '`' && !isFunc(s) && !isModifier(s):
		fmt.Fprintf(b, "`%s`", s)
	default:
		b.WriteString(s)
	}
	return b
}

// AppendComma appends calls Append on all arguments and adds a comma between them.
func (b *Builder) AppendComma(s ...string) *Builder {
	for i := range s {
		if i > 0 {
			b.Comma()
		}
		b.Append(s[i])
	}
	return b
}

// Arg appends an argument to the builder.
func (b *Builder) Arg(a interface{}) *Builder {
	switch a := a.(type) {
	case *raw:
		b.WriteString(a.s)
	default:
		b.WriteString("?")
		b.args = append(b.args, a)
	}
	return b
}

// Args appends a list of arguments to the builder.
func (b *Builder) Args(a ...interface{}) *Builder {
	for i := range a {
		if i > 0 {
			b.Comma()
		}
		b.Arg(a[i])
	}
	return b
}

// Comma adds a comma to the query.
func (b *Builder) Comma() *Builder {
	b.WriteString(", ")
	return b
}

// Pad adds a space to the query.
func (b *Builder) Pad() *Builder {
	b.WriteString(" ")
	return b
}

// Join joins a list of Queriers to the builder.
func (b *Builder) Join(n ...Querier) *Builder {
	for i := range n {
		query, args := n[i].Query()
		b.WriteString(query)
		b.args = append(b.args, args...)
	}
	return b
}

// JoinComma joins a list of Queriers and adds comma between them.
func (b *Builder) JoinComma(n ...Querier) *Builder {
	for i := range n {
		if i > 0 {
			b.Comma()
		}
		query, args := n[i].Query()
		b.WriteString(query)
		b.args = append(b.args, args...)
	}
	return b
}

// Nested gets a callback, and wraps its result with parentheses.
func (b *Builder) Nested(f func(*Builder)) *Builder {
	nb := &Builder{}
	nb.WriteString("(")
	f(nb)
	nb.WriteString(")")
	nb.WriteTo(b)
	b.args = append(b.args, nb.args...)
	return b
}

// clone returns a shallow clone of a builder.
func (b Builder) clone() Builder {
	c := Builder{args: append([]interface{}{}, b.args...)}
	c.Buffer.Write(c.Bytes())
	return c
}

// SetDialect sets the builder dialect. It's used for garnering dialect specific queries.
func (b *Builder) SetDialect(dialect string) *Builder {
	b.dialect = dialect
	return b
}

// Dialect returns the dialect of the builder.
func (b Builder) Dialect() string {
	return b.dialect
}

// Query implements the Querier interface.
func (b Builder) Query() (string, []interface{}) {
	return b.String(), b.args
}

// ColumnBuilder is a builder for column definition in table creation.
type ColumnBuilder struct {
	b    Builder
	typ  string // column type.
	name string // column name.
	attr string // extra attributes.
}

// Column returns a new ColumnBuilder with the given name.
//
//	sql.Column("group_id").Type("int").Attr("UNIQUE")
//
func Column(name string) *ColumnBuilder { return &ColumnBuilder{name: name} }

// Type sets the column type.
func (c *ColumnBuilder) Type(t string) *ColumnBuilder {
	c.typ = t
	return c
}

// Attr sets an extra attribute for the column, like UNIQUE or AUTO_INCREMENT.
func (c *ColumnBuilder) Attr(a string) *ColumnBuilder {
	if c.attr != "" && a != "" {
		c.attr += " "
	}
	c.attr += a
	return c
}

// Query returns query representation of a Column.
func (c *ColumnBuilder) Query() (string, []interface{}) {
	c.b.Append(c.name)
	if c.typ != "" {
		c.b.Pad().WriteString(c.typ)
	}
	if c.attr != "" {
		c.b.Pad().WriteString(c.attr)
	}
	return c.b.String(), c.b.args
}

// TableBuilder is a query builder for `CREATE TABLE` statement.
type TableBuilder struct {
	b           Builder
	name        string           // table name.
	exists      bool             // check existence.
	charset     string           // table charset.
	collation   string           // table collation.
	columns     []*ColumnBuilder // table columns.
	primary     []string         // primary key.
	constraints []Querier        // foreign keys and indices.
}

// CreateTable returns a query builder for the `CREATE TABLE` statement.
//
//	CreateTable("users").
//		Columns(
//			Column("id").Type("int").Attr("auto_increment"),
//			Column("name").Type("varchar(255)"),
//		).
//		PrimaryKey("id")
//
func CreateTable(name string) *TableBuilder { return &TableBuilder{name: name} }

// IfNotExists appends the `IF NOT EXISTS` clause to the `CREATE TABLE` statement.
func (t *TableBuilder) IfNotExists() *TableBuilder {
	t.exists = true
	return t
}

// Column appends the given column to the `CREATE TABLE` statement.
func (t *TableBuilder) Column(c *ColumnBuilder) *TableBuilder {
	t.columns = append(t.columns, c)
	return t
}

// Columns appends the a list of columns to the builder.
func (t *TableBuilder) Columns(c ...*ColumnBuilder) *TableBuilder {
	t.columns = append(t.columns, c...)
	return t
}

// PrimaryKey adds a column to the primary-key constraint in the statement.
func (t *TableBuilder) PrimaryKey(column ...string) *TableBuilder {
	t.primary = append(t.primary, column...)
	return t
}

// ForeignKeys adds a list of foreign-keys to the statement (without constraints).
func (t *TableBuilder) ForeignKeys(fks ...*ForeignKeyBuilder) *TableBuilder {
	queries := make([]Querier, len(fks))
	for i := range fks {
		// erase the constraint symbol/name.
		fks[i].symbol = ""
		queries[i] = fks[i]
	}
	t.constraints = append(t.constraints, queries...)
	return t
}

// Constraints adds a list of foreign-key constraints to the statement.
func (t *TableBuilder) Constraints(fks ...*ForeignKeyBuilder) *TableBuilder {
	queries := make([]Querier, len(fks))
	for i := range fks {
		queries[i] = &Wrapper{"CONSTRAINT %s", fks[i]}
	}
	t.constraints = append(t.constraints, queries...)
	return t
}

// Charset appends the `CHARACTER SET` clause to the statement. MySQL only.
func (t *TableBuilder) Charset(s string) *TableBuilder {
	t.charset = s
	return t
}

// Collate appends the `COLLATE` clause to the statement. MySQL only.
func (t *TableBuilder) Collate(s string) *TableBuilder {
	t.collation = s
	return t
}

// Query returns query representation of a `CREATE TABLE` statement.
func (t *TableBuilder) Query() (string, []interface{}) {
	t.b.WriteString("CREATE TABLE ")
	if t.exists {
		t.b.WriteString("IF NOT EXISTS ")
	}
	t.b.Append(t.name)
	t.b.Nested(func(b *Builder) {
		for i, c := range t.columns {
			if i > 0 {
				b.Comma()
			}
			b.Join(c)
		}
		if len(t.primary) > 0 {
			b.Comma().WriteString("PRIMARY KEY")
			b.Nested(func(b *Builder) {
				b.AppendComma(t.primary...)
			})
		}
		if len(t.constraints) > 0 {
			b.Comma().JoinComma(t.constraints...)
		}
	})
	if t.charset != "" {
		t.b.WriteString(" CHARACTER SET " + t.charset)
	}
	if t.collation != "" {
		t.b.WriteString(" COLLATE " + t.collation)
	}
	return t.b.String(), t.b.args
}

// DescribeBuilder is a query builder for `DESCRIBE` statement.
type DescribeBuilder struct {
	b    Builder
	name string // table name.
}

// Describe returns a query builder for the `DESCRIBE` statement.
//
//	Describe("users")
//
func Describe(name string) *DescribeBuilder { return &DescribeBuilder{name: name} }

// Query returns query representation of a `DESCRIBE` statement.
func (t *DescribeBuilder) Query() (string, []interface{}) {
	t.b.WriteString("DESCRIBE ")
	t.b.Append(t.name)
	return t.b.String(), nil
}

// TableAlter is a query builder for `ALTER TABLE` statement.
type TableAlter struct {
	b        Builder
	name     string    // table to alter.
	Queriers []Querier // columns and foreign-keys to add.
}

// AlterTable returns a query builder for the `ALTER TABLE` statement.
//
//	AlterTable("users").
//		AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
//		AddForeignKey(ForeignKey().Columns("group_id"). Reference(Reference().Table("groups").Columns("id")).OnDelete("CASCADE"))
//
func AlterTable(name string) *TableAlter { return &TableAlter{name: name} }

// AddColumn appends the `ADD COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) AddColumn(c *ColumnBuilder) *TableAlter {
	t.Queriers = append(t.Queriers, &Wrapper{"ADD COLUMN %s", c})
	return t
}

// Modify appends the `MODIFY COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) ModifyColumn(c *ColumnBuilder) *TableAlter {
	t.Queriers = append(t.Queriers, &Wrapper{"MODIFY COLUMN %s", c})
	return t
}

// DropColumn appends the `DROP COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) DropColumn(c *ColumnBuilder) *TableAlter {
	t.Queriers = append(t.Queriers, &Wrapper{"DROP COLUMN %s", c})
	return t
}

// AddForeignKey adds a foreign key constraint to the `ALTER TABLE` statement.
func (t *TableAlter) AddForeignKey(fk *ForeignKeyBuilder) *TableAlter {
	t.Queriers = append(t.Queriers, &Wrapper{"ADD CONSTRAINT %s", fk})
	return t
}

// Query returns query representation of the `ALTER TABLE` statement.
func (t *TableAlter) Query() (string, []interface{}) {
	t.b.WriteString("ALTER TABLE ")
	t.b.Append(t.name)
	t.b.Pad()
	t.b.JoinComma(t.Queriers...)
	return t.b.String(), t.b.args
}

// ForeignKeyBuilder is the builder for the foreign-key constraint clause.
type ForeignKeyBuilder struct {
	b       Builder
	symbol  string
	columns []string
	actions []string
	ref     *ReferenceBuilder
}

// ForeignKey returns a builder for the foreign-key constraint clause in create/alter table statements.
// 	ForeignKey().
// 		Columns("group_id").
//		Reference(Reference().Table("groups").Columns("id")).
//		OnDelete("CASCADE")
//
func ForeignKey(symbol ...string) *ForeignKeyBuilder {
	fk := &ForeignKeyBuilder{}
	if len(symbol) != 0 {
		fk.symbol = symbol[0]
	}
	return fk
}

// Symbol sets the symbol of the foreign key.
func (fk *ForeignKeyBuilder) Symbol(s string) *ForeignKeyBuilder {
	fk.symbol = s
	return fk
}

// Columns sets the columns of the foreign key in the source table.
func (fk *ForeignKeyBuilder) Columns(s ...string) *ForeignKeyBuilder {
	fk.columns = append(fk.columns, s...)
	return fk
}

// Reference sets the reference clause.
func (fk *ForeignKeyBuilder) Reference(r *ReferenceBuilder) *ForeignKeyBuilder {
	fk.ref = r
	return fk
}

// OnDelete sets the on delete action for this constraint.
func (fk *ForeignKeyBuilder) OnDelete(action string) *ForeignKeyBuilder {
	fk.actions = append(fk.actions, "ON DELETE "+action)
	return fk
}

// OnUpdate sets the on delete action for this constraint.
func (fk *ForeignKeyBuilder) OnUpdate(action string) *ForeignKeyBuilder {
	fk.actions = append(fk.actions, "ON UPDATE "+action)
	return fk
}

// Query returns query representation of a foreign key constraint.
func (fk *ForeignKeyBuilder) Query() (string, []interface{}) {
	if fk.symbol != "" {
		fk.b.Append(fk.symbol)
		fk.b.Pad()
	}
	fk.b.WriteString("FOREIGN KEY")
	fk.b.Nested(func(b *Builder) {
		b.AppendComma(fk.columns...)
	})
	fk.b.Pad()
	fk.b.Join(fk.ref)
	for _, action := range fk.actions {
		fk.b.Pad().WriteString(action)
	}
	return fk.b.String(), fk.b.args
}

// ReferenceBuilder is a builder for the reference clause in constraints. For example, in foreign key creation.
type ReferenceBuilder struct {
	b       Builder
	table   string   // referenced table.
	columns []string // referenced columns.
}

// Reference create a reference builder for the reference_option clause.
//
//	Reference().Table("groups").Columns("id")
//
func Reference() *ReferenceBuilder { return &ReferenceBuilder{} }

// Table sets the referenced table.
func (r *ReferenceBuilder) Table(s string) *ReferenceBuilder {
	r.table = s
	return r
}

// Columns sets the columns of the referenced table.
func (r *ReferenceBuilder) Columns(s ...string) *ReferenceBuilder {
	r.columns = append(r.columns, s...)
	return r
}

// Query returns query representation of a reference clause.
func (r *ReferenceBuilder) Query() (string, []interface{}) {
	r.b.WriteString("REFERENCES ")
	r.b.Append(r.table)
	r.b.Nested(func(b *Builder) {
		b.AppendComma(r.columns...)
	})
	return r.b.String(), r.b.args
}

// IndexBuilder is a builder for `CREATE INDEX` statement.
type IndexBuilder struct {
	b       Builder
	name    string
	unique  bool
	table   string
	columns []string
}

// CreateIndex creates a builder for the `CREATE INDEX` statement.
//
//	CreateIndex("index_name").
//		Unique().
//		Table("users").
//		Column("name")
//
// Or:
//
//	CreateIndex("index_name").
//		Unique().
//		Table("users").
//		Columns("name", "age")
//
func CreateIndex(name string) *IndexBuilder {
	return &IndexBuilder{name: name}
}

// Unique sets the index to be a unique index.
func (i *IndexBuilder) Unique() *IndexBuilder {
	i.unique = true
	return i
}

// Table defines the table for the index.
func (i *IndexBuilder) Table(table string) *IndexBuilder {
	i.table = table
	return i
}

// Column appends a column to the column list for the index.
func (i *IndexBuilder) Column(column string) *IndexBuilder {
	i.columns = append(i.columns, column)
	return i
}

// Columns appends the given columns to the column list for the index.
func (i *IndexBuilder) Columns(columns ...string) *IndexBuilder {
	i.columns = append(i.columns, columns...)
	return i
}

// Query returns query representation of a reference clause.
func (i *IndexBuilder) Query() (string, []interface{}) {
	i.b.WriteString("CREATE ")
	if i.unique {
		i.b.WriteString("UNIQUE ")
	}
	i.b.WriteString("INDEX ")
	i.b.Append(i.name)
	i.b.WriteString(" ON ")
	i.b.Append(i.table).Nested(func(b *Builder) {
		b.AppendComma(i.columns...)
	})
	return i.b.String(), nil
}

// DropIndexBuilder is a builder for `DROP INDEX` statement.
type DropIndexBuilder struct {
	b     Builder
	name  string
	table string
}

// DropIndex creates a builder for the `DROP INDEX` statement.
//
//	MySQL:
//
//		DropIndex("index_name").
//			Table("users").
//
// SQLite/PostgreSQL:
//
//		DropIndex("index_name")
//
func DropIndex(name string) *DropIndexBuilder {
	return &DropIndexBuilder{name: name}
}

// Table defines the table for the index.
func (d *DropIndexBuilder) Table(table string) *DropIndexBuilder {
	d.table = table
	return d
}

// Query returns query representation of a reference clause.
func (d *DropIndexBuilder) Query() (string, []interface{}) {
	d.b.WriteString("DROP INDEX ")
	d.b.Append(d.name)
	if d.table != "" {
		d.b.WriteString(" ON ")
		d.b.Append(d.table)
	}
	return d.b.String(), nil
}

// InsertBuilder is a builder for `INSERT INTO` statement.
type InsertBuilder struct {
	b        Builder
	table    string
	columns  []string
	defaults string
	values   [][]interface{}
}

// Insert creates a builder for the `INSERT INTO` statement.
//
//	Insert("users").
//		Columns("name", "age").
//		Values("a8m", 10).
//		Values("foo", 20)
//
// Note: Insert inserts all values in one batch.
func Insert(table string) *InsertBuilder { return &InsertBuilder{table: table} }

// Set is a syntactic sugar API for inserting only one row.
func (i *InsertBuilder) Set(column string, v interface{}) *InsertBuilder {
	i.columns = append(i.columns, column)
	if len(i.values) == 0 {
		i.values = append(i.values, []interface{}{v})
	} else {
		i.values[0] = append(i.values[0], v)
	}
	return i
}

// Columns sets the columns of the insert statement.
func (i *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	i.columns = append(i.columns, columns...)
	return i
}

// Values append a value tuple for the insert statement.
func (i *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	i.values = append(i.values, values)
	return i
}

// Default sets the default values clause based on the dialect type.
func (i *InsertBuilder) Default(d string) *InsertBuilder {
	switch d {
	case dialect.MySQL:
		i.defaults = "VALUES ()"
	case dialect.SQLite:
		i.defaults = "DEFAULT VALUES"
	}
	return i
}

// Query returns query representation of an `INSERT INTO` statement.
func (i *InsertBuilder) Query() (string, []interface{}) {
	i.b.WriteString("INSERT INTO ")
	if i.defaults != "" && len(i.columns) == 0 {
		return i.b.Append(i.table).Pad().String() + i.defaults, nil
	}
	i.b.Append(i.table).Pad().Nested(func(b *Builder) {
		b.AppendComma(i.columns...)
	})
	i.b.WriteString(" VALUES ")
	for j, v := range i.values {
		if j > 0 {
			i.b.Comma()
		}
		i.b.Nested(func(b *Builder) {
			b.Args(v...)
		})
	}
	return i.b.String(), i.b.args
}

// UpdateBuilder is a builder for `UPDATE` statement.
type UpdateBuilder struct {
	b       Builder
	table   string
	where   *Predicate
	nulls   []string
	columns []string
	values  []interface{}
}

// Update creates a builder for the `UPDATE` statement.
//
//	Update("users").Set("name", "foo").Set("age", 10)
//
func Update(table string) *UpdateBuilder { return &UpdateBuilder{table: table} }

// Set sets a column and a its value.
func (u *UpdateBuilder) Set(column string, v interface{}) *UpdateBuilder {
	u.columns = append(u.columns, column)
	u.values = append(u.values, v)
	return u
}

// Add adds a numeric value to the given column.
func (u *UpdateBuilder) Add(column string, v interface{}) *UpdateBuilder {
	u.columns = append(u.columns, column)
	var b Builder
	b.WriteString("COALESCE")
	b.Nested(func(b *Builder) {
		b.Append(column).Comma().Arg(0)
	})
	b.WriteString(" + ")
	b.Arg(v)
	u.values = append(u.values, b)
	return u
}

// SetNull sets a column as null value.
func (u *UpdateBuilder) SetNull(column string) *UpdateBuilder {
	u.nulls = append(u.nulls, column)
	return u
}

// Where adds a where predicate for update statement.
func (u *UpdateBuilder) Where(p *Predicate) *UpdateBuilder {
	if u.where != nil {
		u.where.merge(p)
	} else {
		u.where = p
	}
	return u
}

// Query returns query representation of an `UPDATE` statement.
func (u *UpdateBuilder) Query() (string, []interface{}) {
	u.b.WriteString("UPDATE ")
	u.b.Append(u.table).Pad().WriteString("SET ")
	for i, c := range u.nulls {
		if i > 0 {
			u.b.Comma()
		}
		u.b.Append(c).WriteString(" = NULL")
	}
	if len(u.nulls) > 0 && len(u.columns) > 0 {
		u.b.Comma()
	}
	for i, c := range u.columns {
		if i > 0 {
			u.b.Comma()
		}
		u.b.Append(c).WriteString(" = ")
		switch v := u.values[i].(type) {
		case Querier:
			u.b.Join(v)
		default:
			u.b.Arg(v)
		}
	}
	if u.where != nil {
		u.b.WriteString(" WHERE ")
		u.b.Join(u.where)
	}
	return u.b.String(), u.b.args
}

// DeleteBuilder is a builder for `DELETE` statement.
type DeleteBuilder struct {
	b     Builder
	table string
	where *Predicate
}

// Delete creates a builder for the `DELETE` statement.
//
//	Delete("users").
//		Where(
//			Or(
//				EQ("name", "foo").And().EQ("age", 10),
//				EQ("name", "bar").And().EQ("age", 20),
//				And(
//					EQ("name", "qux"),
//					EQ("age", 1).Or().EQ("age", 2),
//				),
//			),
//		)
//
func Delete(table string) *DeleteBuilder { return &DeleteBuilder{table: table} }

// Where appends a where predicate to the `DELETE` statement.
func (d *DeleteBuilder) Where(p *Predicate) *DeleteBuilder {
	if d.where != nil {
		d.where.merge(p)
	} else {
		d.where = p
	}
	return d
}

// FromSelect make it possible to delete a sub query.
func (d *DeleteBuilder) FromSelect(s *Selector) *DeleteBuilder {
	d.Where(s.where)
	if table, _ := s.from.(*SelectTable); table != nil {
		d.table = table.name
	}
	return d
}

// Query returns query representation of a `DELETE` statement.
func (d *DeleteBuilder) Query() (string, []interface{}) {
	d.b.WriteString("DELETE FROM ")
	d.b.Append(d.table)
	if d.where != nil {
		d.b.WriteString(" WHERE ")
		d.b.Join(d.where)
	}
	return d.b.String(), d.b.args
}

// Predicate is a where predicate.
type Predicate struct {
	b Builder
}

// P creates a new predicates.
//
//	P().EQ("name", "a8m").And().EQ("age", 30)
//
func P() *Predicate { return &Predicate{} }

// Or combines all given predicates with OR between them.
//
//	Or(EQ("name", "foo"), EQ("name", "bar"))
//
func Or(preds ...*Predicate) *Predicate {
	p := P()
	for i := range preds {
		p.Or().b.Nested(func(b *Builder) {
			b.Join(preds[i])
		})
	}
	return p
}

// Or appends an OR only if it's not a start of expression.
func (p *Predicate) Or() *Predicate {
	if p.b.Len() > 0 {
		p.b.WriteString(" OR ")
	}
	return p
}

// False appends the FALSE keyword to the predicate.
//
//	Delete().From("users").Where(False())
//
func False() *Predicate {
	return (&Predicate{}).False()
}

// False appends FALSE to the predicate.
func (p *Predicate) False() *Predicate {
	p.b.WriteString("FALSE")
	return p
}

// Not wraps the given predicate with the not predicate.
//
//	Not(Or(EQ("name", "foo"), EQ("name", "bar")))
//
func Not(pred *Predicate) *Predicate {
	p := P()
	p.Not().b.Nested(func(b *Builder) {
		b.Join(pred)
	})
	return p
}

// Not appends NOT to the predicate.
func (p *Predicate) Not() *Predicate {
	p.b.WriteString("NOT ")
	return p
}

// And combines all given predicates with AND between them.
func And(preds ...*Predicate) *Predicate {
	p := P()
	for i := range preds {
		p.And().b.Nested(func(b *Builder) {
			b.Join(preds[i])
		})
	}
	return p
}

// And appends And only if it's not a start of expression.
func (p *Predicate) And() *Predicate {
	if p.b.Len() > 0 {
		p.b.WriteString(" AND ")
	}
	return p
}

// EQ returns a "=" predicate.
func EQ(col string, value interface{}) *Predicate {
	return (&Predicate{}).EQ(col, value)
}

// EQ appends a "=" predicate.
func (p *Predicate) EQ(col string, arg interface{}) *Predicate {
	p.b.Append(col).WriteString(" = ")
	p.b.Arg(arg)
	return p
}

// NEQ returns a "<>" predicate.
func NEQ(col string, value interface{}) *Predicate {
	return (&Predicate{}).NEQ(col, value)
}

// NEQ appends a "<>" predicate.
func (p *Predicate) NEQ(col string, arg interface{}) *Predicate {
	p.b.Append(col).WriteString(" <> ")
	p.b.Arg(arg)
	return p
}

// LT returns a "<" predicate.
func LT(col string, value interface{}) *Predicate {
	return (&Predicate{}).LT(col, value)
}

// LT appends a "<" predicate.
func (p *Predicate) LT(col string, arg interface{}) *Predicate {
	p.b.Append(col).WriteString(" < ")
	p.b.Arg(arg)
	return p
}

// LTE returns a "<=" predicate.
func LTE(col string, value interface{}) *Predicate {
	return (&Predicate{}).LTE(col, value)
}

// LTE appends a "<=" predicate.
func (p *Predicate) LTE(col string, arg interface{}) *Predicate {
	p.b.Append(col).WriteString(" <= ")
	p.b.Arg(arg)
	return p
}

// GT returns a ">" predicate.
func GT(col string, value interface{}) *Predicate {
	return (&Predicate{}).GT(col, value)
}

// GT appends a ">" predicate.
func (p *Predicate) GT(col string, arg interface{}) *Predicate {
	p.b.Append(col).WriteString(" > ")
	p.b.Arg(arg)
	return p
}

// GTE returns a ">=" predicate.
func GTE(col string, value interface{}) *Predicate {
	return (&Predicate{}).GTE(col, value)
}

// GTE appends a ">=" predicate.
func (p *Predicate) GTE(col string, arg interface{}) *Predicate {
	p.b.Append(col).WriteString(" >= ")
	p.b.Arg(arg)
	return p
}

// NotNull returns the `IS NOT NULL` predicate.
func NotNull(col string) *Predicate {
	return (&Predicate{}).NotNull(col)
}

// NotNull appends the `IS NOT NULL` predicate.
func (p *Predicate) NotNull(col string) *Predicate {
	p.b.Append(col).WriteString(" IS NOT NULL")
	return p
}

// IsNull returns the `IS NULL` predicate.
func IsNull(col string) *Predicate {
	return (&Predicate{}).IsNull(col)
}

// IsNull appends the `IS NULL` predicate.
func (p *Predicate) IsNull(col string) *Predicate {
	p.b.Append(col).WriteString(" IS NULL")
	return p
}

// In returns the `IN` predicate.
func In(col string, args ...interface{}) *Predicate {
	return (&Predicate{}).In(col, args...)
}

// In appends the `IN` predicate.
func (p *Predicate) In(col string, args ...interface{}) *Predicate {
	if len(args) == 0 {
		return p
	}
	p.b.Append(col).WriteString(" IN ")
	p.b.Nested(func(b *Builder) {
		if s, ok := args[0].(*Selector); ok {
			b.Join(s)
		} else {
			b.Args(args...)
		}
	})
	return p
}

// InInts returns the `IN` predicate for ints.
func InInts(col string, args ...int) *Predicate {
	return (&Predicate{}).InInts(col, args...)
}

// InInts adds the `IN` predicate for ints.
func (p *Predicate) InInts(col string, args ...int) *Predicate {
	iface := make([]interface{}, len(args))
	for i := range args {
		iface[i] = args[i]
	}
	return p.In(col, iface...)
}

// NotIn returns the `Not IN` predicate.
func NotIn(col string, args ...interface{}) *Predicate {
	return (&Predicate{}).NotIn(col, args...)
}

// NotIn appends the `Not IN` predicate.
func (p *Predicate) NotIn(col string, args ...interface{}) *Predicate {
	p.b.Append(col).WriteString(" NOT IN ")
	p.b.Nested(func(b *Builder) {
		b.Args(args...)
	})
	return p
}

// Like returns the `LIKE` predicate.
func Like(col, pattern string) *Predicate {
	return (&Predicate{}).Like(col, pattern)
}

// Like appends the `LIKE` predicate.
func (p *Predicate) Like(col, pattern string) *Predicate {
	p.b.Append(col).WriteString(" LIKE ")
	p.b.Arg(pattern)
	return p
}

// HasPrefix is a helper predicate that checks prefix using the LIKE predicate.
func HasPrefix(col, prefix string) *Predicate {
	return (&Predicate{}).HasPrefix(col, prefix)
}

// HasPrefix is a helper predicate that checks prefix using the LIKE predicate.
func (p *Predicate) HasPrefix(col, prefix string) *Predicate {
	return p.Like(col, prefix+"%")
}

// HasSuffix is a helper predicate that checks suffix using the LIKE predicate.
func HasSuffix(col, suffix string) *Predicate { return (&Predicate{}).HasSuffix(col, suffix) }

// HasSuffix is a helper predicate that checks suffix using the LIKE predicate.
func (p *Predicate) HasSuffix(col, suffix string) *Predicate {
	return p.Like(col, "%"+suffix)
}

// EqualFold is a helper predicate that applies the "=" predicate with case-folding.
func EqualFold(col, sub string) *Predicate { return (&Predicate{}).EqualFold(col, sub) }

// EqualFold is a helper predicate that applies the "=" predicate with case-folding.
func (p *Predicate) EqualFold(col, sub string) *Predicate {
	return p.EQ(Lower(col), strings.ToLower(sub))
}

// Contains is a helper predicate that checks substring using the LIKE predicate.
func Contains(col, sub string) *Predicate { return (&Predicate{}).Contains(col, sub) }

// Contains is a helper predicate that checks substring using the LIKE predicate.
func (p *Predicate) Contains(col, sub string) *Predicate {
	return p.Like(col, "%"+sub+"%")
}

// ContainsFold is a helper predicate that checks substring using the LIKE predicate.
func ContainsFold(col, sub string) *Predicate { return (&Predicate{}).ContainsFold(col, sub) }

// ContainsFold is a helper predicate that applies the LIKE predicate with case-folding.
// The recommendation is to avoid using it, and to use a dialect specific feature, like
// `ILIKE` in PostgreSQL, and `COLLATE` clause in MySQL.
func (p *Predicate) ContainsFold(col, sub string) *Predicate {
	return p.Like(Lower(col), "%"+strings.ToLower(sub)+"%")
}

// Lower wraps the given column with the LOWER function.
//
//	P().EQ(sql.Lower("name"), "a8m")
//
func Lower(name string) string {
	var b Builder
	b.WriteString("LOWER")
	b.Nested(func(b *Builder) {
		b.Append(name)
	})
	return b.String()
}

// Upper wraps the given column with the UPPER function.
//
//	P().EQ(sql.Upper("name"), "a8m")
//
func Upper(name string) string {
	var b Builder
	b.WriteString("UPPER")
	b.Nested(func(b *Builder) {
		b.Append(name)
	})
	return b.String()
}

// Query returns query representation of a predicate.
func (p *Predicate) Query() (string, []interface{}) {
	return p.b.String(), p.b.args
}

// merge two predicates.
func (p *Predicate) merge(pred *Predicate) *Predicate {
	query, args := pred.Query()
	p.And().b.WriteString(query)
	p.b.args = append(p.b.args, args...)
	return p
}

// clone returns a shallow clone of p.
func (p *Predicate) clone() *Predicate {
	if p == nil {
		return p
	}
	return &Predicate{p.b.clone()}
}

// TableView is a view that returns a table view. Can ne a Table, Selector or a View (WITH statement).
type TableView interface {
	view()
}

// Count wraps the column with the COUNT aggregation function.
func Count(column string) string {
	return agg("COUNT", column)
}

// Max wraps the column with the MAX aggregation function.
func Max(column string) string {
	return agg("MAX", column)
}

// Min wraps the column with the MIN aggregation function.
func Min(column string) string {
	return agg("MIN", column)
}

// Sum wraps the column with the SUM aggregation function.
func Sum(column string) string {
	return agg("SUM", column)
}

// Avg wraps the column with the AVG aggregation function.
func Avg(column string) string {
	return agg("AVG", column)
}

// As suffixed the given column with an alias (`a` AS `b`).
func As(column string, as string) string {
	var b Builder
	b.Append(column).Pad().WriteString("AS")
	b.Pad().Append(as)
	return b.String()
}

// Distinct prefixed the given columns with the `DISTINCT` keyword (DISTINCT `id`).
func Distinct(columns ...string) string {
	var b Builder
	b.WriteString("DISTINCT")
	b.Pad().AppendComma(columns...)
	return b.String()
}

// SelectTable is a table selector.
type SelectTable struct {
	quote bool
	name  string
	as    string
}

// Table returns a new table selector.
//
//	t1 := Table("users").As("u")
//	return Select(t1.C("name"))
//
func Table(name string) *SelectTable {
	return &SelectTable{quote: true, name: name}
}

// As adds the AS clause to the table selector.
func (s *SelectTable) As(alias string) *SelectTable {
	s.as = alias
	return s
}

// C returns a formatted string for the table column.
func (s *SelectTable) C(column string) string {
	name := s.name
	if s.as != "" {
		name = s.as
	}
	return fmt.Sprintf("`%s`.`%s`", name, column)
}

// Columns returns a list of formatted strings for the table columns.
func (s *SelectTable) Columns(columns ...string) []string {
	var names []string
	for _, c := range columns {
		names = append(names, s.C(c))
	}
	return names
}

// Unquote makes the table name to be formatted as raw string (unquoted).
// It is useful whe you don't want to query tables under the current database.
// For example: "INFORMATION_SCHEMA.TABLE_CONSTRAINTS" in MySQL.
func (s *SelectTable) Unquote() *SelectTable {
	s.quote = false
	return s
}

// ref returns the table reference.
func (s *SelectTable) ref() string {
	switch {
	case !s.quote:
		return s.name
	case s.as == "":
		return fmt.Sprintf("`%s`", s.name)
	default:
		return fmt.Sprintf("`%s` AS `%s`", s.name, s.as)
	}
}

// implement the table view.
func (*SelectTable) view() {}

// join table option.
type join struct {
	on    string
	kind  string
	table TableView
}

// Selector a builder for the `SELECT` statement.
type Selector struct {
	as       string
	columns  []string
	from     TableView
	joins    []join
	where    *Predicate
	or       bool
	not      bool
	order    []string
	group    []string
	having   *Predicate
	limit    *int
	offset   *int
	distinct bool
}

// Select returns a new selector for the `SELECT` statement.
//
//	t1 := Table("users").As("u")
//	t2 := Select().From(Table("groups")).Where(EQ("user_id", 10)).As("g")
//	return Select(t1.C("id"), t2.C("name")).
//			From(t1).
//			Join(t2).
//			On(t1.C("id"), t2.C("user_id"))
//
func Select(columns ...string) *Selector {
	return (&Selector{}).Select(columns...)
}

// Select changes the columns selection of the SELECT statement.
// Empty selection means all columns *.
func (s *Selector) Select(columns ...string) *Selector {
	s.columns = columns
	return s
}

// From sets the source of `FROM` clause.
func (s *Selector) From(t TableView) *Selector {
	s.from = t
	return s
}

// Distinct adds the DISTINCT keyword to the `SELECT` statement.
func (s *Selector) Distinct() *Selector {
	s.distinct = true
	return s
}

// Limit adds the `LIMIT` clause to the `SELECT` statement.
func (s *Selector) Limit(limit int) *Selector {
	s.limit = &limit
	return s
}

// Offset adds the `OFFSET` clause to the `SELECT` statement.
func (s *Selector) Offset(offset int) *Selector {
	s.offset = &offset
	return s
}

// Where sets or appends the given predicate to the statement.
func (s *Selector) Where(p *Predicate) *Selector {
	if s.not {
		p = Not(p)
		s.not = false
	}
	switch {
	case s.where == nil:
		s.where = p
	case s.where != nil && s.or:
		s.where = Or(s.where, p)
		s.or = false
	default:
		s.where.merge(p)
	}
	return s
}

// FromSelect copies the predicate from a selector.
func (s *Selector) FromSelect(s2 *Selector) *Selector {
	s.where = s2.where
	return s
}

// Not sets the next coming predicate with not.
func (s *Selector) Not() *Selector {
	s.not = true
	return s
}

// Or sets the next coming predicate with OR operator (disjunction).
func (s *Selector) Or() *Selector {
	s.or = true
	return s
}

// Table returns the selected table.
func (s *Selector) Table() *SelectTable {
	return s.from.(*SelectTable)
}

// Join appends a `JOIN` clause to the statement.
func (s *Selector) Join(t TableView) *Selector {
	s.joins = append(s.joins, join{
		kind:  "JOIN",
		table: t,
	})
	switch view := t.(type) {
	case *SelectTable:
		if view.as == "" {
			view.as = "t0"
		}
	case *Selector:
		if view.as == "" {
			view.as = "t" + strconv.Itoa(len(s.joins))
		}
	}
	return s
}

// C returns a formatted string for a selected column from this statement.
func (s *Selector) C(column string) string {
	if s.as != "" {
		return fmt.Sprintf("`%s`.`%s`", s.as, column)
	}
	return s.Table().C(column)
}

// Columns returns a list of formatted strings for a selected columns from this statement.
func (s *Selector) Columns(columns ...string) []string {
	var names []string
	for _, c := range columns {
		names = append(names, s.C(c))
	}
	return names
}

// On sets the `ON` clause for the `JOIN` operation.
func (s *Selector) On(c1, c2 string) *Selector {
	if len(s.joins) > 0 {
		s.joins[len(s.joins)-1].on = fmt.Sprintf("%s = %s", c1, c2)
	}
	return s
}

// As give this selection an alias.
func (s *Selector) As(alias string) *Selector {
	s.as = alias
	return s
}

// Count sets the Select statement to be a `SELECT COUNT(*)`.
func (s *Selector) Count(columns ...string) *Selector {
	column := "*"
	if len(columns) > 0 {
		var b Builder
		b.AppendComma(columns...)
		column = b.String()
	}
	s.columns = []string{Count(column)}
	return s
}

// Clone returns a duplicate of the selector, including all associated steps. It can be
// used to prepare common SELECT statements and use them differently after the clone is made.
func (s *Selector) Clone() *Selector {
	if s == nil {
		return nil
	}
	return &Selector{
		as:       s.as,
		or:       s.or,
		not:      s.not,
		from:     s.from,
		limit:    s.limit,
		offset:   s.offset,
		distinct: s.distinct,
		where:    s.where.clone(),
		having:   s.having.clone(),
		joins:    append([]join{}, s.joins...),
		group:    append([]string{}, s.group...),
		order:    append([]string{}, s.order...),
		columns:  append([]string{}, s.columns...),
	}
}

// Asc adds the ASC suffix for the given column.
func Asc(column string) string {
	var b Builder
	b.Append(column).WriteString(" ASC")
	return b.String()
}

// Desc adds the DESC suffix for the given column.
func Desc(column string) string {
	var b Builder
	b.Append(column).WriteString(" DESC")
	return b.String()
}

// OrderBy appends the `ORDER BY` clause to the `SELECT` statement.
func (s *Selector) OrderBy(columns ...string) *Selector {
	s.order = append(s.order, columns...)
	return s
}

// GroupBy appends the `GROUP BY` clause to the `SELECT` statement.
func (s *Selector) GroupBy(columns ...string) *Selector {
	s.group = append(s.group, columns...)
	return s
}

// Having appends a predicate for the `HAVING` clause.
func (s *Selector) Having(p *Predicate) *Selector {
	s.having = p
	return s
}

// Query returns query representation of a `SELECT` statement.
func (s *Selector) Query() (string, []interface{}) {
	var b Builder
	b.WriteString("SELECT ")
	if s.distinct {
		b.WriteString("DISTINCT ")
	}
	if len(s.columns) > 0 {
		b.AppendComma(s.columns...)
	} else {
		b.WriteString("*")
	}
	b.WriteString(" FROM ")
	switch t := s.from.(type) {
	case *SelectTable:
		b.WriteString(t.ref())
	case *Selector:
		query, args := t.Query()
		b.WriteString(fmt.Sprintf("(%s) AS `%s`", query, t.as))
		b.args = append(b.args, args...)
	}
	for _, join := range s.joins {
		b.WriteString(fmt.Sprintf(" %s ", join.kind))
		switch view := join.table.(type) {
		case *SelectTable:
			b.WriteString(view.ref())
		case *Selector:
			query, args := view.Query()
			b.WriteString(fmt.Sprintf("(%s) AS `%s`", query, view.as))
			b.args = append(b.args, args...)
		}
		if join.on != "" {
			b.WriteString(" ON ")
			b.WriteString(join.on)
		}
	}
	if s.where != nil {
		b.WriteString(" WHERE ")
		query, args := s.where.Query()
		b.WriteString(query)
		b.args = append(b.args, args...)
	}
	if len(s.group) > 0 {
		b.WriteString(" GROUP BY ")
		b.AppendComma(s.group...)
	}
	if s.having != nil {
		b.WriteString(" HAVING ")
		query, args := s.where.Query()
		b.WriteString(query)
		b.args = append(b.args, args...)
	}
	if len(s.order) > 0 {
		b.WriteString(" ORDER BY ")
		b.AppendComma(s.order...)
	}
	if s.limit != nil {
		b.WriteString(" LIMIT ")
		b.Arg(*s.limit)
	}
	if s.offset != nil {
		b.WriteString(" OFFSET ")
		b.Arg(*s.offset)
	}
	return b.String(), b.args
}

// implement the table view interface.
func (*Selector) view() {}

// WithBuilder is the builder for the `WITH` statement.
type WithBuilder struct {
	b    Builder
	name string
	s    *Selector
}

// With returns a new builder for the `WITH` statement.
//
//	n := Queries{With("users_view").As(Select().From(Table("users"))), Select().From(Table("users_view"))}
//	return n.Query()
//
func With(name string) *WithBuilder {
	return &WithBuilder{name: name}
}

// Name returns the name of the view.
func (w *WithBuilder) Name() string { return w.name }

// As sets the view sub query.
func (w *WithBuilder) As(s *Selector) *WithBuilder {
	w.s = s
	return w
}

// Query returns query representation of a `WITH` clause.
func (w *WithBuilder) Query() (string, []interface{}) {
	w.b.WriteString("WITH " + w.name)
	w.b.WriteString(" AS ")
	w.b.Nested(func(b *Builder) {
		b.Join(w.s)
	})
	return w.b.String(), w.b.args
}

// implement the table view interface.
func (*WithBuilder) view() {}

// Wrapper wraps a given Querier with different format.
// Used to prefix/suffix other queries.
type Wrapper struct {
	format  string
	wrapped Querier
}

// Query returns query representation of a wrapped Querier.
func (w *Wrapper) Query() (string, []interface{}) {
	query, args := w.wrapped.Query()
	return fmt.Sprintf(w.format, query), args
}

// Raw returns a raw sql Querier that is placed as-is in the query.
func Raw(s string) Querier { return &raw{s} }

type raw struct{ s string }

func (r *raw) Query() (string, []interface{}) { return r.s, nil }

func isFunc(s string) bool {
	return strings.Contains(s, "(") && strings.Contains(s, ")")
}

func isModifier(s string) bool {
	for _, m := range []string{"DISTINCT", "ALL", "WITH ROLLUP"} {
		if strings.HasPrefix(s, m) {
			return true
		}
	}
	return false
}

func agg(fn, column string) string {
	var b Builder
	b.WriteString(fn)
	b.Nested(func(b *Builder) {
		b.Append(column)
	})
	return b.String()
}
