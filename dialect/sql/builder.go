// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Package sql provides wrappers around the standard database/sql package
// to allow the generated code to interact with a statically-typed API.
//
// Users that are interacting with this package should be aware that the
// following builders don't check the given SQL syntax nor validate or escape
// user-inputs. ~All validations are expected to be happened in the generated
// ent package.
package sql

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"entgo.io/ent/dialect"
)

// Querier wraps the basic Query method that is implemented
// by the different builders in this file.
type Querier interface {
	// Query returns the query representation of the element
	// and its arguments (if any).
	Query() (string, []any)
}

// querierErr allowed propagate Querier's inner error
type querierErr interface {
	Err() error
}

// ColumnBuilder is a builder for column definition in table creation.
type ColumnBuilder struct {
	Builder
	typ    string             // column type.
	name   string             // column name.
	attr   string             // extra attributes.
	modify bool               // modify existing.
	fk     *ForeignKeyBuilder // foreign-key constraint.
	check  func(*Builder)     // column checks.
}

// Column returns a new ColumnBuilder with the given name.
//
//	sql.Column("group_id").Type("int").Attr("UNIQUE")
func Column(name string) *ColumnBuilder { return &ColumnBuilder{name: name} }

// Type sets the column type.
func (c *ColumnBuilder) Type(t string) *ColumnBuilder {
	c.typ = t
	return c
}

// Attr sets an extra attribute for the column, like UNIQUE or AUTO_INCREMENT.
func (c *ColumnBuilder) Attr(attr string) *ColumnBuilder {
	if c.attr != "" && attr != "" {
		c.attr += " "
	}
	c.attr += attr
	return c
}

// Constraint adds the CONSTRAINT clause to the ADD COLUMN statement in SQLite.
func (c *ColumnBuilder) Constraint(fk *ForeignKeyBuilder) *ColumnBuilder {
	c.fk = fk
	return c
}

// Check adds a CHECK clause to the ADD COLUMN statement.
func (c *ColumnBuilder) Check(check func(*Builder)) *ColumnBuilder {
	c.check = check
	return c
}

// Query returns query representation of a Column.
func (c *ColumnBuilder) Query() (string, []any) {
	c.Ident(c.name)
	if c.typ != "" {
		if c.postgres() && c.modify {
			c.WriteString(" TYPE")
		}
		c.Pad().WriteString(c.typ)
	}
	if c.attr != "" {
		c.Pad().WriteString(c.attr)
	}
	if c.fk != nil {
		c.WriteString(" CONSTRAINT " + c.fk.symbol)
		c.Pad().Join(c.fk.ref)
		for _, action := range c.fk.actions {
			c.Pad().WriteString(action)
		}
	}
	if c.check != nil {
		c.WriteString(" CHECK ")
		c.Wrap(c.check)
	}
	return c.String(), c.args
}

// TableBuilder is a query builder for `CREATE TABLE` statement.
type TableBuilder struct {
	Builder
	name        string           // table name.
	exists      bool             // check existence.
	charset     string           // table charset.
	collation   string           // table collation.
	options     string           // table options.
	columns     []Querier        // table columns.
	primary     []string         // primary key.
	constraints []Querier        // foreign keys and indices.
	checks      []func(*Builder) // check constraints.
}

// CreateTable returns a query builder for the `CREATE TABLE` statement.
//
//	CreateTable("users").
//		Columns(
//			Column("id").Type("int").Attr("auto_increment"),
//			Column("name").Type("varchar(255)"),
//		).
//		PrimaryKey("id")
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

// Columns appends a list of columns to the builder.
func (t *TableBuilder) Columns(columns ...*ColumnBuilder) *TableBuilder {
	t.columns = make([]Querier, 0, len(columns))
	for i := range columns {
		t.columns = append(t.columns, columns[i])
	}
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
		// Erase the constraint symbol/name.
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

// Checks adds CHECK clauses to the CREATE TABLE statement.
func (t *TableBuilder) Checks(checks ...func(*Builder)) *TableBuilder {
	t.checks = append(t.checks, checks...)
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

// Options appends additional options to the statement (MySQL only).
func (t *TableBuilder) Options(s string) *TableBuilder {
	t.options = s
	return t
}

// Query returns query representation of a `CREATE TABLE` statement.
//
// CREATE TABLE [IF NOT EXISTS] name
//
//	(table definition)
//	[charset and collation]
func (t *TableBuilder) Query() (string, []any) {
	t.WriteString("CREATE TABLE ")
	if t.exists {
		t.WriteString("IF NOT EXISTS ")
	}
	t.Ident(t.name)
	t.Wrap(func(b *Builder) {
		b.JoinComma(t.columns...)
		if len(t.primary) > 0 {
			b.Comma().WriteString("PRIMARY KEY")
			b.Wrap(func(b *Builder) {
				b.IdentComma(t.primary...)
			})
		}
		if len(t.constraints) > 0 {
			b.Comma().JoinComma(t.constraints...)
		}
		for _, check := range t.checks {
			check(b.Comma())
		}
	})
	if t.charset != "" {
		t.WriteString(" CHARACTER SET " + t.charset)
	}
	if t.collation != "" {
		t.WriteString(" COLLATE " + t.collation)
	}
	if t.options != "" {
		t.WriteString(" " + t.options)
	}
	return t.String(), t.args
}

// DescribeBuilder is a query builder for `DESCRIBE` statement.
type DescribeBuilder struct {
	Builder
	name string // table name.
}

// Describe returns a query builder for the `DESCRIBE` statement.
//
//	Describe("users")
func Describe(name string) *DescribeBuilder { return &DescribeBuilder{name: name} }

// Query returns query representation of a `DESCRIBE` statement.
func (t *DescribeBuilder) Query() (string, []any) {
	t.WriteString("DESCRIBE ")
	t.Ident(t.name)
	return t.String(), nil
}

// TableAlter is a query builder for `ALTER TABLE` statement.
type TableAlter struct {
	Builder
	name    string    // table to alter.
	Queries []Querier // columns and foreign-keys to add.
}

// AlterTable returns a query builder for the `ALTER TABLE` statement.
//
//	AlterTable("users").
//		AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
//		AddForeignKey(ForeignKey().Columns("group_id").
//			Reference(Reference().Table("groups").Columns("id")).OnDelete("CASCADE")),
//		)
func AlterTable(name string) *TableAlter { return &TableAlter{name: name} }

// AddColumn appends the `ADD COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) AddColumn(c *ColumnBuilder) *TableAlter {
	t.Queries = append(t.Queries, &Wrapper{"ADD COLUMN %s", c})
	return t
}

// ModifyColumn appends the `MODIFY/ALTER COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) ModifyColumn(c *ColumnBuilder) *TableAlter {
	switch {
	case t.postgres():
		c.modify = true
		t.Queries = append(t.Queries, &Wrapper{"ALTER COLUMN %s", c})
	default:
		t.Queries = append(t.Queries, &Wrapper{"MODIFY COLUMN %s", c})
	}
	return t
}

// RenameColumn appends the `RENAME COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) RenameColumn(old, new string) *TableAlter {
	t.Queries = append(t.Queries, Raw(fmt.Sprintf("RENAME COLUMN %s TO %s", t.Quote(old), t.Quote(new))))
	return t
}

// ModifyColumns calls ModifyColumn with each of the given builders.
func (t *TableAlter) ModifyColumns(cs ...*ColumnBuilder) *TableAlter {
	for _, c := range cs {
		t.ModifyColumn(c)
	}
	return t
}

// DropColumn appends the `DROP COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) DropColumn(c *ColumnBuilder) *TableAlter {
	t.Queries = append(t.Queries, &Wrapper{"DROP COLUMN %s", c})
	return t
}

// ChangeColumn appends the `CHANGE COLUMN` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) ChangeColumn(name string, c *ColumnBuilder) *TableAlter {
	prefix := fmt.Sprintf("CHANGE COLUMN %s", t.Quote(name))
	t.Queries = append(t.Queries, &Wrapper{prefix + " %s", c})
	return t
}

// RenameIndex appends the `RENAME INDEX` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) RenameIndex(curr, new string) *TableAlter {
	t.Queries = append(t.Queries, Raw(fmt.Sprintf("RENAME INDEX %s TO %s", t.Quote(curr), t.Quote(new))))
	return t
}

// DropIndex appends the `DROP INDEX` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) DropIndex(name string) *TableAlter {
	t.Queries = append(t.Queries, Raw(fmt.Sprintf("DROP INDEX %s", t.Quote(name))))
	return t
}

// AddIndex appends the `ADD INDEX` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) AddIndex(idx *IndexBuilder) *TableAlter {
	b := &Builder{dialect: t.dialect}
	b.WriteString("ADD ")
	if idx.unique {
		b.WriteString("UNIQUE ")
	}
	b.WriteString("INDEX ")
	b.Ident(idx.name)
	b.Wrap(func(b *Builder) {
		b.IdentComma(idx.columns...)
	})
	t.Queries = append(t.Queries, b)
	return t
}

// AddForeignKey adds a foreign key constraint to the `ALTER TABLE` statement.
func (t *TableAlter) AddForeignKey(fk *ForeignKeyBuilder) *TableAlter {
	t.Queries = append(t.Queries, &Wrapper{"ADD CONSTRAINT %s", fk})
	return t
}

// DropConstraint appends the `DROP CONSTRAINT` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) DropConstraint(ident string) *TableAlter {
	t.Queries = append(t.Queries, Raw(fmt.Sprintf("DROP CONSTRAINT %s", t.Quote(ident))))
	return t
}

// DropForeignKey appends the `DROP FOREIGN KEY` clause to the given `ALTER TABLE` statement.
func (t *TableAlter) DropForeignKey(ident string) *TableAlter {
	t.Queries = append(t.Queries, Raw(fmt.Sprintf("DROP FOREIGN KEY %s", t.Quote(ident))))
	return t
}

// Query returns query representation of the `ALTER TABLE` statement.
//
//	ALTER TABLE name
//		[alter_specification]
func (t *TableAlter) Query() (string, []any) {
	t.WriteString("ALTER TABLE ")
	t.Ident(t.name)
	t.Pad()
	t.JoinComma(t.Queries...)
	return t.String(), t.args
}

// IndexAlter is a query builder for `ALTER INDEX` statement.
type IndexAlter struct {
	Builder
	name    string    // index to alter.
	Queries []Querier // alter options.
}

// AlterIndex returns a query builder for the `ALTER INDEX` statement.
//
//	AlterIndex("old_key").
//		Rename("new_key")
func AlterIndex(name string) *IndexAlter { return &IndexAlter{name: name} }

// Rename appends the `RENAME TO` clause to the `ALTER INDEX` statement.
func (i *IndexAlter) Rename(name string) *IndexAlter {
	i.Queries = append(i.Queries, Raw(fmt.Sprintf("RENAME TO %s", i.Quote(name))))
	return i
}

// Query returns query representation of the `ALTER INDEX` statement.
//
//	ALTER INDEX name
//		[alter_specification]
func (i *IndexAlter) Query() (string, []any) {
	i.WriteString("ALTER INDEX ")
	i.Ident(i.name)
	i.Pad()
	i.JoinComma(i.Queries...)
	return i.String(), i.args
}

// ForeignKeyBuilder is the builder for the foreign-key constraint clause.
type ForeignKeyBuilder struct {
	Builder
	symbol  string
	columns []string
	actions []string
	ref     *ReferenceBuilder
}

// ForeignKey returns a builder for the foreign-key constraint clause in create/alter table statements.
//
//	ForeignKey().
//		Columns("group_id").
//		Reference(Reference().Table("groups").Columns("id")).
//		OnDelete("CASCADE")
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
func (fk *ForeignKeyBuilder) Query() (string, []any) {
	if fk.symbol != "" {
		fk.Ident(fk.symbol).Pad()
	}
	fk.WriteString("FOREIGN KEY")
	fk.Wrap(func(b *Builder) {
		b.IdentComma(fk.columns...)
	})
	fk.Pad().Join(fk.ref)
	for _, action := range fk.actions {
		fk.Pad().WriteString(action)
	}
	return fk.String(), fk.args
}

// ReferenceBuilder is a builder for the reference clause in constraints. For example, in foreign key creation.
type ReferenceBuilder struct {
	Builder
	table   string   // referenced table.
	columns []string // referenced columns.
}

// Reference creates a reference builder for the reference_option clause.
//
//	Reference().Table("groups").Columns("id")
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
func (r *ReferenceBuilder) Query() (string, []any) {
	r.WriteString("REFERENCES ")
	r.Ident(r.table)
	r.Wrap(func(b *Builder) {
		b.IdentComma(r.columns...)
	})
	return r.String(), r.args
}

// IndexBuilder is a builder for `CREATE INDEX` statement.
type IndexBuilder struct {
	Builder
	name    string
	unique  bool
	exists  bool
	table   string
	method  string
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
func CreateIndex(name string) *IndexBuilder {
	return &IndexBuilder{name: name}
}

// IfNotExists appends the `IF NOT EXISTS` clause to the `CREATE INDEX` statement.
func (i *IndexBuilder) IfNotExists() *IndexBuilder {
	i.exists = true
	return i
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

// Using sets the method to create the index with.
func (i *IndexBuilder) Using(method string) *IndexBuilder {
	i.method = method
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
func (i *IndexBuilder) Query() (string, []any) {
	i.WriteString("CREATE ")
	if i.unique {
		i.WriteString("UNIQUE ")
	}
	i.WriteString("INDEX ")
	if i.exists {
		i.WriteString("IF NOT EXISTS ")
	}
	i.Ident(i.name)
	i.WriteString(" ON ")
	i.Ident(i.table)
	switch i.dialect {
	case dialect.Postgres:
		if i.method != "" {
			i.WriteString(" USING ").Ident(i.method)
		}
		i.Wrap(func(b *Builder) {
			b.IdentComma(i.columns...)
		})
	case dialect.MySQL:
		i.Wrap(func(b *Builder) {
			b.IdentComma(i.columns...)
		})
		if i.method != "" {
			i.WriteString(" USING " + i.method)
		}
	default:
		i.Wrap(func(b *Builder) {
			b.IdentComma(i.columns...)
		})
	}
	return i.String(), nil
}

// DropIndexBuilder is a builder for `DROP INDEX` statement.
type DropIndexBuilder struct {
	Builder
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
//	SQLite/PostgreSQL:
//
//		DropIndex("index_name")
func DropIndex(name string) *DropIndexBuilder {
	return &DropIndexBuilder{name: name}
}

// Table defines the table for the index.
func (d *DropIndexBuilder) Table(table string) *DropIndexBuilder {
	d.table = table
	return d
}

// Query returns query representation of a reference clause.
//
//	DROP INDEX index_name [ON table_name]
func (d *DropIndexBuilder) Query() (string, []any) {
	d.WriteString("DROP INDEX ")
	d.Ident(d.name)
	if d.table != "" {
		d.WriteString(" ON ")
		d.Ident(d.table)
	}
	return d.String(), nil
}

// InsertBuilder is a builder for `INSERT INTO` statement.
type InsertBuilder struct {
	Builder
	table     string
	schema    string
	columns   []string
	defaults  bool
	returning []string
	values    [][]any
	conflict  *conflict
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

// Schema sets the database name for the insert table.
func (i *InsertBuilder) Schema(name string) *InsertBuilder {
	i.schema = name
	return i
}

// Set is a syntactic sugar API for inserting only one row.
func (i *InsertBuilder) Set(column string, v any) *InsertBuilder {
	i.columns = append(i.columns, column)
	if len(i.values) == 0 {
		i.values = append(i.values, []any{v})
	} else {
		i.values[0] = append(i.values[0], v)
	}
	return i
}

// Columns appends columns to the INSERT statement.
func (i *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	i.columns = append(i.columns, columns...)
	return i
}

// Values append a value tuple for the insert statement.
func (i *InsertBuilder) Values(values ...any) *InsertBuilder {
	i.values = append(i.values, values)
	return i
}

// Default sets the default values clause based on the dialect type.
func (i *InsertBuilder) Default() *InsertBuilder {
	i.defaults = true
	return i
}

// Returning adds the `RETURNING` clause to the insert statement.
// Supported by SQLite and PostgreSQL.
func (i *InsertBuilder) Returning(columns ...string) *InsertBuilder {
	i.returning = columns
	return i
}

type (
	// conflict holds the configuration for the
	// `ON CONFLICT` / `ON DUPLICATE KEY` clause.
	conflict struct {
		target struct {
			constraint string
			columns    []string
			where      *Predicate
		}
		action struct {
			nothing bool
			where   *Predicate
			update  []func(*UpdateSet)
		}
	}

	// ConflictOption allows configuring the
	// conflict config using functional options.
	ConflictOption func(*conflict)
)

// ConflictColumns sets the unique constraints that trigger the conflict
// resolution on insert to perform an upsert operation. The columns must
// have a unique constraint applied to trigger this behaviour.
//
//	sql.Insert("users").
//		Columns("id", "name").
//		Values(1, "Mashraki").
//		OnConflict(
//			sql.ConflictColumns("id"),
//			sql.ResolveWithNewValues(),
//		)
func ConflictColumns(names ...string) ConflictOption {
	return func(c *conflict) {
		c.target.columns = names
	}
}

// ConflictConstraint allows setting the constraint
// name (i.e. `ON CONSTRAINT <name>`) for PostgreSQL.
//
//	sql.Insert("users").
//		Columns("id", "name").
//		Values(1, "Mashraki").
//		OnConflict(
//			sql.ConflictConstraint("users_pkey"),
//			sql.ResolveWithNewValues(),
//		)
func ConflictConstraint(name string) ConflictOption {
	return func(c *conflict) {
		c.target.constraint = name
	}
}

// ConflictWhere allows inference of partial unique indexes. See, PostgreSQL
// doc: https://www.postgresql.org/docs/current/sql-insert.html#SQL-ON-CONFLICT
func ConflictWhere(p *Predicate) ConflictOption {
	return func(c *conflict) {
		c.target.where = p
	}
}

// UpdateWhere allows setting the update condition. Only rows
// for which this expression returns true will be updated.
func UpdateWhere(p *Predicate) ConflictOption {
	return func(c *conflict) {
		c.action.where = p
	}
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported by SQLite and PostgreSQL.
//
//	sql.Insert("users").
//		Columns("id", "name").
//		Values(1, "Mashraki").
//		OnConflict(
//			sql.ConflictColumns("id"),
//			sql.DoNothing()
//		)
func DoNothing() ConflictOption {
	return func(c *conflict) {
		c.action.nothing = true
	}
}

// ResolveWithIgnore sets each column to itself to force an update and return the ID,
// otherwise does not change any data. This may still trigger update hooks in the database.
//
//	sql.Insert("users").
//		Columns("id").
//		Values(1).
//		OnConflict(
//			sql.ConflictColumns("id"),
//			sql.ResolveWithIgnore()
//		)
//
//	// Output:
//	// MySQL: INSERT INTO `users` (`id`) VALUES(1) ON DUPLICATE KEY UPDATE `id` = `users`.`id`
//	// PostgreSQL: INSERT INTO "users" ("id") VALUES(1) ON CONFLICT ("id") DO UPDATE SET "id" = "users"."id
func ResolveWithIgnore() ConflictOption {
	return func(c *conflict) {
		c.action.update = append(c.action.update, func(u *UpdateSet) {
			for _, c := range u.columns {
				u.SetIgnore(c)
			}
		})
	}
}

// ResolveWithNewValues updates columns using the new values proposed
// for insertion using the special EXCLUDED/VALUES table.
//
//	sql.Insert("users").
//		Columns("id", "name").
//		Values(1, "Mashraki").
//		OnConflict(
//			sql.ConflictColumns("id"),
//			sql.ResolveWithNewValues()
//		)
//
//	// Output:
//	// MySQL: INSERT INTO `users` (`id`, `name`) VALUES(1, 'Mashraki) ON DUPLICATE KEY UPDATE `id` = VALUES(`id`), `name` = VALUES(`name`),
//	// PostgreSQL: INSERT INTO "users" ("id") VALUES(1) ON CONFLICT ("id") DO UPDATE SET "id" = "excluded"."id, "name" = "excluded"."name"
func ResolveWithNewValues() ConflictOption {
	return func(c *conflict) {
		c.action.update = append(c.action.update, func(u *UpdateSet) {
			for _, c := range u.columns {
				u.SetExcluded(c)
			}
		})
	}
}

// ResolveWith allows setting a custom function to set the `UPDATE` clause.
//
//	Insert("users").
//		Columns("id", "name").
//		Values(1, "Mashraki").
//		OnConflict(
//			ConflictColumns("name"),
//			ResolveWith(func(u *UpdateSet) {
//				u.SetIgnore("id")
//				u.SetNull("created_at")
//				u.Set("name", Expr(u.Excluded().C("name")))
//			}),
//		)
func ResolveWith(fn func(*UpdateSet)) ConflictOption {
	return func(c *conflict) {
		c.action.update = append(c.action.update, fn)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	sql.Insert("users").
//		Columns("id", "name").
//		Values(1, "Mashraki").
//		OnConflict(
//			sql.ConflictColumns("id"),
//			sql.ResolveWithNewValues()
//		)
func (i *InsertBuilder) OnConflict(opts ...ConflictOption) *InsertBuilder {
	if i.conflict == nil {
		i.conflict = &conflict{}
	}
	for _, opt := range opts {
		opt(i.conflict)
	}
	return i
}

// UpdateSet describes a set of changes of the `DO UPDATE` clause.
type UpdateSet struct {
	*UpdateBuilder
	columns []string
}

// Table returns the table the `UPSERT` statement is executed on.
func (u *UpdateSet) Table() *SelectTable {
	return Dialect(u.UpdateBuilder.dialect).Table(u.UpdateBuilder.table)
}

// Columns returns all columns in the `INSERT` statement.
func (u *UpdateSet) Columns() []string {
	return u.columns
}

// UpdateColumns returns all columns in the `UPDATE` statement.
func (u *UpdateSet) UpdateColumns() []string {
	return append(u.UpdateBuilder.nulls, u.UpdateBuilder.columns...)
}

// Set sets a column to a given value.
func (u *UpdateSet) Set(column string, v any) *UpdateSet {
	u.UpdateBuilder.Set(column, v)
	return u
}

// Add adds a numeric value to the given column.
func (u *UpdateSet) Add(column string, v any) *UpdateSet {
	u.UpdateBuilder.Add(column, v)
	return u
}

// SetNull sets a column as null value.
func (u *UpdateSet) SetNull(column string) *UpdateSet {
	u.UpdateBuilder.SetNull(column)
	return u
}

// SetIgnore sets the column to itself. For example, "id" = "users"."id".
func (u *UpdateSet) SetIgnore(name string) *UpdateSet {
	return u.Set(name, Expr(u.Table().C(name)))
}

// SetExcluded sets the column name to its EXCLUDED/VALUES value.
// For example, "c" = "excluded"."c", or `c` = VALUES(`c`).
func (u *UpdateSet) SetExcluded(name string) *UpdateSet {
	switch u.UpdateBuilder.Dialect() {
	case dialect.MySQL:
		u.UpdateBuilder.Set(name, ExprFunc(func(b *Builder) {
			b.WriteString("VALUES(").Ident(name).WriteByte(')')
		}))
	default:
		t := Dialect(u.UpdateBuilder.dialect).Table("excluded")
		u.UpdateBuilder.Set(name, Expr(t.C(name)))
	}
	return u
}

// Query returns query representation of an `INSERT INTO` statement.
func (i *InsertBuilder) Query() (string, []any) {
	query, args, _ := i.QueryErr()
	return query, args
}

// QueryErr returns query representation of an `INSERT INTO`
// statement and any error occurred in building the statement.
func (i *InsertBuilder) QueryErr() (string, []any, error) {
	b := i.Builder.clone()
	b.WriteString("INSERT INTO ")
	b.writeSchema(i.schema)
	b.Ident(i.table).Pad()
	if i.defaults && len(i.columns) == 0 {
		i.writeDefault(&b)
	} else {
		b.WriteByte('(').IdentComma(i.columns...).WriteByte(')')
		b.WriteString(" VALUES ")
		for j, v := range i.values {
			if j > 0 {
				b.Comma()
			}
			b.WriteByte('(').Args(v...).WriteByte(')')
		}
	}
	if i.conflict != nil {
		i.writeConflict(&b)
	}
	joinReturning(i.returning, &b)
	return b.String(), b.args, b.Err()
}

func (i *InsertBuilder) writeDefault(b *Builder) {
	switch i.Dialect() {
	case dialect.MySQL:
		b.WriteString("VALUES ()")
	case dialect.SQLite, dialect.Postgres:
		b.WriteString("DEFAULT VALUES")
	}
}

func (i *InsertBuilder) writeConflict(b *Builder) {
	switch i.Dialect() {
	case dialect.MySQL:
		b.WriteString(" ON DUPLICATE KEY UPDATE ")
		// Fallback to ResolveWithIgnore() as MySQL
		// does not support the "DO NOTHING" clause.
		if i.conflict.action.nothing {
			i.OnConflict(ResolveWithIgnore())
		}
	case dialect.SQLite, dialect.Postgres:
		b.WriteString(" ON CONFLICT")
		switch t := i.conflict.target; {
		case t.constraint != "" && len(t.columns) != 0:
			b.AddError(fmt.Errorf("duplicate CONFLICT clauses: %q, %q", t.constraint, t.columns))
		case t.constraint != "":
			b.WriteString(" ON CONSTRAINT ").Ident(t.constraint)
		case len(t.columns) != 0:
			b.WriteString(" (").IdentComma(t.columns...).WriteByte(')')
		}
		if p := i.conflict.target.where; p != nil {
			b.WriteString(" WHERE ").Join(p)
		}
		if i.conflict.action.nothing {
			b.WriteString(" DO NOTHING")
			return
		}
		b.WriteString(" DO UPDATE SET ")
	}
	if len(i.conflict.action.update) == 0 {
		b.AddError(errors.New("missing action for 'DO UPDATE SET' clause"))
	}
	u := &UpdateSet{UpdateBuilder: Dialect(i.dialect).Update(i.table), columns: i.columns}
	u.Builder = *b
	for _, f := range i.conflict.action.update {
		f(u)
	}
	u.writeSetter(b)
	if p := i.conflict.action.where; p != nil {
		p.qualifier = i.table
		b.WriteString(" WHERE ").Join(p)
	}
}

// UpdateBuilder is a builder for `UPDATE` statement.
type UpdateBuilder struct {
	Builder
	table     string
	schema    string
	where     *Predicate
	nulls     []string
	columns   []string
	returning []string
	values    []any
	order     []any
	limit     *int
	prefix    Queries
}

// Update creates a builder for the `UPDATE` statement.
//
//	Update("users").Set("name", "foo").Set("age", 10)
func Update(table string) *UpdateBuilder { return &UpdateBuilder{table: table} }

// Schema sets the database name for the updated table.
func (u *UpdateBuilder) Schema(name string) *UpdateBuilder {
	u.schema = name
	return u
}

// Set sets a column to a given value. If `Set` was called before with
// the same column name, it overrides the value of the previous call.
func (u *UpdateBuilder) Set(column string, v any) *UpdateBuilder {
	for i := range u.columns {
		if column == u.columns[i] {
			u.values[i] = v
			return u
		}
	}
	u.columns = append(u.columns, column)
	u.values = append(u.values, v)
	return u
}

// Add adds a numeric value to the given column. Note that, calling Set(c)
// after Add(c) will erase previous calls with c from the builder.
func (u *UpdateBuilder) Add(column string, v any) *UpdateBuilder {
	u.columns = append(u.columns, column)
	u.values = append(u.values, ExprFunc(func(b *Builder) {
		b.WriteString("COALESCE")
		b.Wrap(func(b *Builder) {
			b.Ident(Table(u.table).C(column)).Comma().WriteByte('0')
		})
		b.WriteString(" + ")
		b.Arg(v)
	}))
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
		u.where = And(u.where, p)
	} else {
		u.where = p
	}
	return u
}

// FromSelect makes it possible to update entities that match the sub-query.
func (u *UpdateBuilder) FromSelect(s *Selector) *UpdateBuilder {
	u.Where(s.where)
	if t := s.Table(); t != nil {
		u.table = t.name
	}
	return u
}

// Empty reports whether this builder does not contain update changes.
func (u *UpdateBuilder) Empty() bool {
	return len(u.columns) == 0 && len(u.nulls) == 0
}

// OrderBy appends the `ORDER BY` clause to the `UPDATE` statement.
// Supported by SQLite and MySQL.
func (u *UpdateBuilder) OrderBy(columns ...string) *UpdateBuilder {
	if u.postgres() {
		u.AddError(errors.New("ORDER BY is not supported by PostgreSQL"))
		return u
	}
	for i := range columns {
		u.order = append(u.order, columns[i])
	}
	return u
}

// Limit appends the `LIMIT` clause to the `UPDATE` statement.
// Supported by SQLite and MySQL.
func (u *UpdateBuilder) Limit(limit int) *UpdateBuilder {
	if u.postgres() {
		u.AddError(errors.New("LIMIT is not supported by PostgreSQL"))
		return u
	}
	u.limit = &limit
	return u
}

// Prefix prefixes the UPDATE statement with list of statements.
func (u *UpdateBuilder) Prefix(stmts ...Querier) *UpdateBuilder {
	u.prefix = append(u.prefix, stmts...)
	return u
}

// Returning adds the `RETURNING` clause to the insert statement.
// Supported by SQLite and PostgreSQL.
func (u *UpdateBuilder) Returning(columns ...string) *UpdateBuilder {
	u.returning = columns
	return u
}

// Query returns query representation of an `UPDATE` statement.
func (u *UpdateBuilder) Query() (string, []any) {
	b := u.Builder.clone()
	if len(u.prefix) > 0 {
		b.join(u.prefix, " ")
		b.Pad()
	}
	b.WriteString("UPDATE ")
	b.writeSchema(u.schema)
	b.Ident(u.table).WriteString(" SET ")
	u.writeSetter(&b)
	if u.where != nil {
		b.WriteString(" WHERE ")
		b.Join(u.where)
	}
	joinReturning(u.returning, &b)
	joinOrder(u.order, &b)
	if u.limit != nil {
		b.WriteString(" LIMIT ")
		b.WriteString(strconv.Itoa(*u.limit))
	}
	return b.String(), b.args
}

// writeSetter writes the "SET" clause for the UPDATE statement.
func (u *UpdateBuilder) writeSetter(b *Builder) {
	for i, c := range u.nulls {
		if i > 0 {
			b.Comma()
		}
		b.Ident(c).WriteString(" = NULL")
	}
	if len(u.nulls) > 0 && len(u.columns) > 0 {
		b.Comma()
	}
	for i, c := range u.columns {
		if i > 0 {
			b.Comma()
		}
		b.Ident(c).WriteString(" = ")
		switch v := u.values[i].(type) {
		case Querier:
			b.Join(v)
		default:
			b.Arg(v)
		}
	}
}

// DeleteBuilder is a builder for `DELETE` statement.
type DeleteBuilder struct {
	Builder
	table  string
	schema string
	where  *Predicate
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
func Delete(table string) *DeleteBuilder { return &DeleteBuilder{table: table} }

// Schema sets the database name for the table whose row will be deleted.
func (d *DeleteBuilder) Schema(name string) *DeleteBuilder {
	d.schema = name
	return d
}

// Where appends a where predicate to the `DELETE` statement.
func (d *DeleteBuilder) Where(p *Predicate) *DeleteBuilder {
	if d.where != nil {
		d.where = And(d.where, p)
	} else {
		d.where = p
	}
	return d
}

// FromSelect makes it possible to delete a sub query.
func (d *DeleteBuilder) FromSelect(s *Selector) *DeleteBuilder {
	d.Where(s.where)
	if t := s.Table(); t != nil {
		d.table = t.name
	}
	return d
}

// Query returns query representation of a `DELETE` statement.
func (d *DeleteBuilder) Query() (string, []any) {
	d.WriteString("DELETE FROM ")
	d.writeSchema(d.schema)
	d.Ident(d.table)
	if d.where != nil {
		d.WriteString(" WHERE ")
		d.Join(d.where)
	}
	return d.String(), d.args
}

// Predicate is a where predicate.
type Predicate struct {
	Builder
	depth int
	fns   []func(*Builder)
}

// P creates a new predicate.
//
//	P().EQ("name", "a8m").And().EQ("age", 30)
func P(fns ...func(*Builder)) *Predicate {
	return &Predicate{fns: fns}
}

// ExprP creates a new predicate from the given expression.
//
//	ExprP("A = ? AND B > ?", args...)
func ExprP(exr string, args ...any) *Predicate {
	return P(func(b *Builder) {
		b.Join(Expr(exr, args...))
	})
}

// Or combines all given predicates with OR between them.
//
//	Or(EQ("name", "foo"), EQ("name", "bar"))
func Or(preds ...*Predicate) *Predicate {
	p := P()
	return p.Append(func(b *Builder) {
		p.mayWrap(preds, b, "OR")
	})
}

// False appends the FALSE keyword to the predicate.
//
//	Delete().From("users").Where(False())
func False() *Predicate {
	return P().False()
}

// False appends FALSE to the predicate.
func (p *Predicate) False() *Predicate {
	return p.Append(func(b *Builder) {
		b.WriteString("FALSE")
	})
}

// Not wraps the given predicate with the not predicate.
//
//	Not(Or(EQ("name", "foo"), EQ("name", "bar")))
func Not(pred *Predicate) *Predicate {
	return P().Not().Append(func(b *Builder) {
		b.Wrap(func(b *Builder) {
			b.Join(pred)
		})
	})
}

// Not appends NOT to the predicate.
func (p *Predicate) Not() *Predicate {
	return p.Append(func(b *Builder) {
		b.WriteString("NOT ")
	})
}

// ColumnsOp returns a new predicate between 2 columns.
func ColumnsOp(col1, col2 string, op Op) *Predicate {
	return P().ColumnsOp(col1, col2, op)
}

// ColumnsOp appends the given predicate between 2 columns.
func (p *Predicate) ColumnsOp(col1, col2 string, op Op) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col1)
		b.WriteOp(op)
		b.Ident(col2)
	})
}

// And combines all given predicates with AND between them.
func And(preds ...*Predicate) *Predicate {
	p := P()
	return p.Append(func(b *Builder) {
		p.mayWrap(preds, b, "AND")
	})
}

// IsTrue appends a predicate that checks if the column value is truthy.
func IsTrue(col string) *Predicate {
	return P().IsTrue(col)
}

// IsTrue appends a predicate that checks if the column value is truthy.
func (p *Predicate) IsTrue(col string) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col)
	})
}

// IsFalse appends a predicate that checks if the column value is falsey.
func IsFalse(col string) *Predicate {
	return P().IsFalse(col)
}

// IsFalse appends a predicate that checks if the column value is falsey.
func (p *Predicate) IsFalse(col string) *Predicate {
	return p.Append(func(b *Builder) {
		b.WriteString("NOT ").Ident(col)
	})
}

// EQ returns a "=" predicate.
func EQ(col string, value any) *Predicate {
	return P().EQ(col, value)
}

// EQ appends a "=" predicate.
func (p *Predicate) EQ(col string, arg any) *Predicate {
	// A small optimization to avoid passing
	// arguments when it can be avoided.
	switch arg := arg.(type) {
	case bool:
		if arg {
			return IsTrue(col)
		}
		return IsFalse(col)
	default:
		return p.Append(func(b *Builder) {
			b.Ident(col)
			b.WriteOp(OpEQ)
			p.arg(b, arg)
		})
	}
}

// ColumnsEQ appends a "=" predicate between 2 columns.
func ColumnsEQ(col1, col2 string) *Predicate {
	return P().ColumnsEQ(col1, col2)
}

// ColumnsEQ appends a "=" predicate between 2 columns.
func (p *Predicate) ColumnsEQ(col1, col2 string) *Predicate {
	return p.ColumnsOp(col1, col2, OpEQ)
}

// NEQ returns a "<>" predicate.
func NEQ(col string, value any) *Predicate {
	return P().NEQ(col, value)
}

// NEQ appends a "<>" predicate.
func (p *Predicate) NEQ(col string, arg any) *Predicate {
	// A small optimization to avoid passing
	// arguments when it can be avoided.
	switch arg := arg.(type) {
	case bool:
		if arg {
			return IsFalse(col)
		}
		return IsTrue(col)
	default:
		return p.Append(func(b *Builder) {
			b.Ident(col)
			b.WriteOp(OpNEQ)
			p.arg(b, arg)
		})
	}
}

// ColumnsNEQ appends a "<>" predicate between 2 columns.
func ColumnsNEQ(col1, col2 string) *Predicate {
	return P().ColumnsNEQ(col1, col2)
}

// ColumnsNEQ appends a "<>" predicate between 2 columns.
func (p *Predicate) ColumnsNEQ(col1, col2 string) *Predicate {
	return p.ColumnsOp(col1, col2, OpNEQ)
}

// LT returns a "<" predicate.
func LT(col string, value any) *Predicate {
	return P().LT(col, value)
}

// LT appends a "<" predicate.
func (p *Predicate) LT(col string, arg any) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col)
		p.WriteOp(OpLT)
		p.arg(b, arg)
	})
}

// ColumnsLT appends a "<" predicate between 2 columns.
func ColumnsLT(col1, col2 string) *Predicate {
	return P().ColumnsLT(col1, col2)
}

// ColumnsLT appends a "<" predicate between 2 columns.
func (p *Predicate) ColumnsLT(col1, col2 string) *Predicate {
	return p.ColumnsOp(col1, col2, OpLT)
}

// LTE returns a "<=" predicate.
func LTE(col string, value any) *Predicate {
	return P().LTE(col, value)
}

// LTE appends a "<=" predicate.
func (p *Predicate) LTE(col string, arg any) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col)
		p.WriteOp(OpLTE)
		p.arg(b, arg)
	})
}

// ColumnsLTE appends a "<=" predicate between 2 columns.
func ColumnsLTE(col1, col2 string) *Predicate {
	return P().ColumnsLTE(col1, col2)
}

// ColumnsLTE appends a "<=" predicate between 2 columns.
func (p *Predicate) ColumnsLTE(col1, col2 string) *Predicate {
	return p.ColumnsOp(col1, col2, OpLTE)
}

// GT returns a ">" predicate.
func GT(col string, value any) *Predicate {
	return P().GT(col, value)
}

// GT appends a ">" predicate.
func (p *Predicate) GT(col string, arg any) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col)
		p.WriteOp(OpGT)
		p.arg(b, arg)
	})
}

// ColumnsGT appends a ">" predicate between 2 columns.
func ColumnsGT(col1, col2 string) *Predicate {
	return P().ColumnsGT(col1, col2)
}

// ColumnsGT appends a ">" predicate between 2 columns.
func (p *Predicate) ColumnsGT(col1, col2 string) *Predicate {
	return p.ColumnsOp(col1, col2, OpGT)
}

// GTE returns a ">=" predicate.
func GTE(col string, value any) *Predicate {
	return P().GTE(col, value)
}

// GTE appends a ">=" predicate.
func (p *Predicate) GTE(col string, arg any) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col)
		p.WriteOp(OpGTE)
		p.arg(b, arg)
	})
}

// ColumnsGTE appends a ">=" predicate between 2 columns.
func ColumnsGTE(col1, col2 string) *Predicate {
	return P().ColumnsGTE(col1, col2)
}

// ColumnsGTE appends a ">=" predicate between 2 columns.
func (p *Predicate) ColumnsGTE(col1, col2 string) *Predicate {
	return p.ColumnsOp(col1, col2, OpGTE)
}

// NotNull returns the `IS NOT NULL` predicate.
func NotNull(col string) *Predicate {
	return P().NotNull(col)
}

// NotNull appends the `IS NOT NULL` predicate.
func (p *Predicate) NotNull(col string) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col).WriteString(" IS NOT NULL")
	})
}

// IsNull returns the `IS NULL` predicate.
func IsNull(col string) *Predicate {
	return P().IsNull(col)
}

// IsNull appends the `IS NULL` predicate.
func (p *Predicate) IsNull(col string) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col).WriteString(" IS NULL")
	})
}

// In returns the `IN` predicate.
func In(col string, args ...any) *Predicate {
	return P().In(col, args...)
}

// In appends the `IN` predicate.
func (p *Predicate) In(col string, args ...any) *Predicate {
	// If no arguments were provided, append the FALSE constant, since
	// we cannot apply "IN ()". This will make this predicate falsy.
	if len(args) == 0 {
		return p.False()
	}
	return p.Append(func(b *Builder) {
		b.Ident(col).WriteOp(OpIn)
		b.Wrap(func(b *Builder) {
			if s, ok := args[0].(*Selector); ok {
				b.Join(s)
			} else {
				b.Args(args...)
			}
		})
	})
}

// InInts returns the `IN` predicate for ints.
func InInts(col string, args ...int) *Predicate {
	return P().InInts(col, args...)
}

// InValues adds the `IN` predicate for slice of driver.Value.
func InValues(col string, args ...driver.Value) *Predicate {
	return P().InValues(col, args...)
}

// InInts adds the `IN` predicate for ints.
func (p *Predicate) InInts(col string, args ...int) *Predicate {
	iface := make([]any, len(args))
	for i := range args {
		iface[i] = args[i]
	}
	return p.In(col, iface...)
}

// InValues adds the `IN` predicate for slice of driver.Value.
func (p *Predicate) InValues(col string, args ...driver.Value) *Predicate {
	iface := make([]any, len(args))
	for i := range args {
		iface[i] = args[i]
	}
	return p.In(col, iface...)
}

// NotIn returns the `Not IN` predicate.
func NotIn(col string, args ...any) *Predicate {
	return P().NotIn(col, args...)
}

// NotIn appends the `Not IN` predicate.
func (p *Predicate) NotIn(col string, args ...any) *Predicate {
	// If no arguments were provided, append the NOT FALSE constant, since
	// we cannot apply "NOT IN ()". This will make this predicate truthy.
	if len(args) == 0 {
		return Not(p.False())
	}
	return p.Append(func(b *Builder) {
		b.Ident(col).WriteOp(OpNotIn)
		b.Wrap(func(b *Builder) {
			if s, ok := args[0].(*Selector); ok {
				b.Join(s)
			} else {
				b.Args(args...)
			}
		})
	})
}

// Exists returns the `Exists` predicate.
func Exists(query Querier) *Predicate {
	return P().Exists(query)
}

// Exists appends the `EXISTS` predicate with the given query.
func (p *Predicate) Exists(query Querier) *Predicate {
	return p.Append(func(b *Builder) {
		b.WriteString("EXISTS ")
		b.Wrap(func(b *Builder) {
			b.Join(query)
		})
	})
}

// NotExists returns the `NotExists` predicate.
func NotExists(query Querier) *Predicate {
	return P().NotExists(query)
}

// NotExists appends the `NOT EXISTS` predicate with the given query.
func (p *Predicate) NotExists(query Querier) *Predicate {
	return p.Append(func(b *Builder) {
		b.WriteString("NOT EXISTS ")
		b.Wrap(func(b *Builder) {
			b.Join(query)
		})
	})
}

// Like returns the `LIKE` predicate.
func Like(col, pattern string) *Predicate {
	return P().Like(col, pattern)
}

// Like appends the `LIKE` predicate.
func (p *Predicate) Like(col, pattern string) *Predicate {
	return p.Append(func(b *Builder) {
		b.Ident(col).WriteOp(OpLike)
		b.Arg(pattern)
	})
}

// escape escapes w with the default escape character ('/'),
// to be used by the pattern matching functions below.
// The second return value indicates if w was escaped or not.
func escape(w string) (string, bool) {
	var n int
	for i := range w {
		if c := w[i]; c == '%' || c == '_' || c == '\\' {
			n++
		}
	}
	// No characters to escape.
	if n == 0 {
		return w, false
	}
	var b strings.Builder
	b.Grow(len(w) + n)
	for _, c := range w {
		if c == '%' || c == '_' || c == '\\' {
			b.WriteByte('\\')
		}
		b.WriteRune(c)
	}
	return b.String(), true
}

func (p *Predicate) escapedLike(col, left, right, word string) *Predicate {
	return p.Append(func(b *Builder) {
		w, escaped := escape(word)
		b.Ident(col).WriteOp(OpLike)
		b.Arg(left + w + right)
		if p.dialect == dialect.SQLite && escaped {
			p.WriteString(" ESCAPE ").Arg("\\")
		}
	})
}

// HasPrefix is a helper predicate that checks prefix using the LIKE predicate.
func HasPrefix(col, prefix string) *Predicate {
	return P().HasPrefix(col, prefix)
}

// HasPrefix is a helper predicate that checks prefix using the LIKE predicate.
func (p *Predicate) HasPrefix(col, prefix string) *Predicate {
	return p.escapedLike(col, "", "%", prefix)
}

// ColumnsHasPrefix appends a new predicate that checks if the given column begins with the other column (prefix).
func ColumnsHasPrefix(col, prefixC string) *Predicate {
	return P().ColumnsHasPrefix(col, prefixC)
}

// ColumnsHasPrefix appends a new predicate that checks if the given column begins with the other column (prefix).
func (p *Predicate) ColumnsHasPrefix(col, prefixC string) *Predicate {
	return p.Append(func(b *Builder) {
		switch p.dialect {
		case dialect.MySQL:
			b.Ident(col)
			b.WriteOp(OpLike)
			b.S("CONCAT(REPLACE(REPLACE(").Ident(prefixC).S(", '_', '\\_'), '%', '\\%'), '%')")
		case dialect.Postgres, dialect.SQLite:
			b.Ident(col)
			b.WriteOp(OpLike)
			b.S("(REPLACE(REPLACE(").Ident(prefixC).S(", '_', '\\_'), '%', '\\%') || '%')")
			if p.dialect == dialect.SQLite {
				p.WriteString(" ESCAPE ").Arg("\\")
			}
		default:
			b.AddError(fmt.Errorf("ColumnsHasPrefix: unsupported dialect: %q", p.dialect))
		}
	})
}

// HasSuffix is a helper predicate that checks suffix using the LIKE predicate.
func HasSuffix(col, suffix string) *Predicate { return P().HasSuffix(col, suffix) }

// HasSuffix is a helper predicate that checks suffix using the LIKE predicate.
func (p *Predicate) HasSuffix(col, suffix string) *Predicate {
	return p.escapedLike(col, "%", "", suffix)
}

// EqualFold is a helper predicate that applies the "=" predicate with case-folding.
func EqualFold(col, sub string) *Predicate { return P().EqualFold(col, sub) }

// EqualFold is a helper predicate that applies the "=" predicate with case-folding.
func (p *Predicate) EqualFold(col, sub string) *Predicate {
	return p.Append(func(b *Builder) {
		f := &Func{}
		f.SetDialect(b.dialect)
		switch b.dialect {
		case dialect.MySQL:
			// We assume the CHARACTER SET is configured to utf8mb4,
			// because this how it is defined in dialect/sql/schema.
			b.Ident(col).WriteString(" COLLATE utf8mb4_general_ci = ")
			b.Arg(strings.ToLower(sub))
		case dialect.Postgres:
			b.Ident(col).WriteString(" ILIKE ")
			w, _ := escape(sub)
			b.Arg(strings.ToLower(w))
		default: // SQLite.
			f.Lower(col)
			b.WriteString(f.String())
			b.WriteOp(OpEQ)
			b.Arg(strings.ToLower(sub))
		}
	})
}

// Contains is a helper predicate that checks substring using the LIKE predicate.
func Contains(col, sub string) *Predicate { return P().Contains(col, sub) }

// Contains is a helper predicate that checks substring using the LIKE predicate.
func (p *Predicate) Contains(col, substr string) *Predicate {
	return p.escapedLike(col, "%", "%", substr)
}

// ContainsFold is a helper predicate that checks substring using the LIKE predicate with case-folding.
func ContainsFold(col, sub string) *Predicate { return P().ContainsFold(col, sub) }

// ContainsFold is a helper predicate that applies the LIKE predicate with case-folding.
func (p *Predicate) ContainsFold(col, substr string) *Predicate {
	return p.Append(func(b *Builder) {
		w, escaped := escape(substr)
		switch b.dialect {
		case dialect.MySQL:
			// We assume the CHARACTER SET is configured to utf8mb4,
			// because this how it is defined in dialect/sql/schema.
			b.Ident(col).WriteString(" COLLATE utf8mb4_general_ci LIKE ")
			b.Arg("%" + strings.ToLower(w) + "%")
		case dialect.Postgres:
			b.Ident(col).WriteString(" ILIKE ")
			b.Arg("%" + strings.ToLower(w) + "%")
		default: // SQLite.
			var f Func
			f.SetDialect(b.dialect)
			f.Lower(col)
			b.WriteString(f.String()).WriteString(" LIKE ")
			b.Arg("%" + strings.ToLower(w) + "%")
			if escaped {
				p.WriteString(" ESCAPE ").Arg("\\")
			}
		}
	})
}

// CompositeGT returns a composite ">" predicate
func CompositeGT(columns []string, args ...any) *Predicate {
	return P().CompositeGT(columns, args...)
}

// CompositeLT returns a composite "<" predicate
func CompositeLT(columns []string, args ...any) *Predicate {
	return P().CompositeLT(columns, args...)
}

func (p *Predicate) compositeP(operator string, columns []string, args ...any) *Predicate {
	return p.Append(func(b *Builder) {
		b.Wrap(func(nb *Builder) {
			nb.IdentComma(columns...)
		})
		b.WriteString(operator)
		b.WriteString("(")
		b.Args(args...)
		b.WriteString(")")
	})
}

// CompositeGT returns a composite ">" predicate.
func (p *Predicate) CompositeGT(columns []string, args ...any) *Predicate {
	const operator = " > "
	return p.compositeP(operator, columns, args...)
}

// CompositeLT appends a composite "<" predicate.
func (p *Predicate) CompositeLT(columns []string, args ...any) *Predicate {
	const operator = " < "
	return p.compositeP(operator, columns, args...)
}

// Append appends a new function to the predicate callbacks.
// The callback list are executed on call to Query.
func (p *Predicate) Append(f func(*Builder)) *Predicate {
	p.fns = append(p.fns, f)
	return p
}

// Query returns query representation of a predicate.
func (p *Predicate) Query() (string, []any) {
	if p.Len() > 0 || len(p.args) > 0 {
		p.Reset()
		p.args = nil
	}
	for _, f := range p.fns {
		f(&p.Builder)
	}
	return p.String(), p.args
}

// arg calls Builder.Arg, but wraps `a` with parens in case of a Selector.
func (*Predicate) arg(b *Builder, a any) {
	switch a.(type) {
	case *Selector:
		b.Wrap(func(b *Builder) {
			b.Arg(a)
		})
	default:
		b.Arg(a)
	}
}

// clone returns a shallow clone of p.
func (p *Predicate) clone() *Predicate {
	if p == nil {
		return p
	}
	return &Predicate{fns: append([]func(*Builder){}, p.fns...)}
}

func (p *Predicate) mayWrap(preds []*Predicate, b *Builder, op string) {
	switch n := len(preds); {
	case n == 1:
		b.Join(preds[0])
		return
	case n > 1 && p.depth != 0:
		b.WriteByte('(')
		defer b.WriteByte(')')
	}
	for i := range preds {
		preds[i].depth = p.depth + 1
		if i > 0 {
			b.WriteByte(' ')
			b.WriteString(op)
			b.WriteByte(' ')
		}
		if len(preds[i].fns) > 1 {
			b.Wrap(func(b *Builder) {
				b.Join(preds[i])
			})
		} else {
			b.Join(preds[i])
		}
	}
}

// Func represents an SQL function.
type Func struct {
	Builder
	fns []func(*Builder)
}

// Lower wraps the given column with the LOWER function.
//
//	P().EQ(sql.Lower("name"), "a8m")
func Lower(ident string) string {
	f := &Func{}
	f.Lower(ident)
	return f.String()
}

// Lower wraps the given ident with the LOWER function.
func (f *Func) Lower(ident string) {
	f.byName("LOWER", ident)
}

// Count wraps the ident with the COUNT aggregation function.
func Count(ident string) string {
	f := &Func{}
	f.Count(ident)
	return f.String()
}

// Count wraps the ident with the COUNT aggregation function.
func (f *Func) Count(ident string) {
	f.byName("COUNT", ident)
}

// Max wraps the ident with the MAX aggregation function.
func Max(ident string) string {
	f := &Func{}
	f.Max(ident)
	return f.String()
}

// Max wraps the ident with the MAX aggregation function.
func (f *Func) Max(ident string) {
	f.byName("MAX", ident)
}

// Min wraps the ident with the MIN aggregation function.
func Min(ident string) string {
	f := &Func{}
	f.Min(ident)
	return f.String()
}

// Min wraps the ident with the MIN aggregation function.
func (f *Func) Min(ident string) {
	f.byName("MIN", ident)
}

// Sum wraps the ident with the SUM aggregation function.
func Sum(ident string) string {
	f := &Func{}
	f.Sum(ident)
	return f.String()
}

// Sum wraps the ident with the SUM aggregation function.
func (f *Func) Sum(ident string) {
	f.byName("SUM", ident)
}

// Avg wraps the ident with the AVG aggregation function.
func Avg(ident string) string {
	f := &Func{}
	f.Avg(ident)
	return f.String()
}

// Avg wraps the ident with the AVG aggregation function.
func (f *Func) Avg(ident string) {
	f.byName("AVG", ident)
}

// byName wraps an identifier with a function name.
func (f *Func) byName(fn, ident string) {
	f.Append(func(b *Builder) {
		f.WriteString(fn)
		f.Wrap(func(b *Builder) {
			b.Ident(ident)
		})
	})
}

// Append appends a new function to the function callbacks.
// The callback list are executed on call to String.
func (f *Func) Append(fn func(*Builder)) *Func {
	f.fns = append(f.fns, fn)
	return f
}

// String implements the fmt.Stringer.
func (f *Func) String() string {
	for _, fn := range f.fns {
		fn(&f.Builder)
	}
	return f.Builder.String()
}

// As suffixed the given column with an alias (`a` AS `b`).
func As(ident string, as string) string {
	b := &Builder{}
	b.fromIdent(ident)
	b.Ident(ident).Pad().WriteString("AS")
	b.Pad().Ident(as)
	return b.String()
}

// Distinct prefixed the given columns with the `DISTINCT` keyword (DISTINCT `id`).
func Distinct(idents ...string) string {
	b := &Builder{}
	if len(idents) > 0 {
		b.fromIdent(idents[0])
	}
	b.WriteString("DISTINCT")
	b.Pad().IdentComma(idents...)
	return b.String()
}

// TableView is a view that returns a table view. Can be a Table, Selector or a View (WITH statement).
type TableView interface {
	view()
	// C returns a formatted string prefixed
	// with the table view qualifier.
	C(string) string
}

// queryView allows using Querier (expressions) in the FROM clause.
type queryView struct{ Querier }

func (*queryView) view() {}

func (q *queryView) C(column string) string {
	if tv, ok := q.Querier.(TableView); ok {
		return tv.C(column)
	}
	return column
}

// SelectTable is a table selector.
type SelectTable struct {
	Builder
	as     string
	name   string
	schema string
	quote  bool
}

// Table returns a new table selector.
//
//	t1 := Table("users").As("u")
//	return Select(t1.C("name"))
func Table(name string) *SelectTable {
	return &SelectTable{quote: true, name: name}
}

// Schema sets the schema name of the table.
func (s *SelectTable) Schema(name string) *SelectTable {
	s.schema = name
	return s
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
	b := &Builder{dialect: s.dialect}
	if s.as == "" {
		b.writeSchema(s.schema)
	}
	b.Ident(name).WriteByte('.').Ident(column)
	return b.String()
}

// Columns returns a list of formatted strings for the table columns.
func (s *SelectTable) Columns(columns ...string) []string {
	names := make([]string, 0, len(columns))
	for _, c := range columns {
		names = append(names, s.C(c))
	}
	return names
}

// Unquote makes the table name to be formatted as raw string (unquoted).
// It is useful when you don't want to query tables under the current database.
// For example: "INFORMATION_SCHEMA.TABLE_CONSTRAINTS" in MySQL.
func (s *SelectTable) Unquote() *SelectTable {
	s.quote = false
	return s
}

// ref returns the table reference.
func (s *SelectTable) ref() string {
	if !s.quote {
		return s.name
	}
	b := &Builder{dialect: s.dialect}
	b.writeSchema(s.schema)
	b.Ident(s.name)
	if s.as != "" {
		b.WriteString(" AS ")
		b.Ident(s.as)
	}
	return b.String()
}

// implement the table view.
func (*SelectTable) view() {}

// join table option.
type join struct {
	on    *Predicate
	kind  string
	table TableView
}

// clone a joiner.
func (j join) clone() join {
	if sel, ok := j.table.(*Selector); ok {
		j.table = sel.Clone()
	}
	j.on = j.on.clone()
	return j
}

// Selector is a builder for the `SELECT` statement.
type Selector struct {
	Builder
	// ctx stores contextual data typically from
	// generated code such as alternate table schemas.
	ctx       context.Context
	as        string
	selection []selection
	from      []TableView
	joins     []join
	collected [][]*Predicate
	where     *Predicate
	or        bool
	not       bool
	order     []any
	group     []string
	having    *Predicate
	limit     *int
	offset    *int
	distinct  bool
	setOps    []setOp
	prefix    Queries
	lock      *LockOptions
}

// New returns a new Selector with the same dialect and context.
func (s *Selector) New() *Selector {
	c := Dialect(s.dialect).Select()
	if s.ctx != nil {
		c = c.WithContext(s.ctx)
	}
	return c
}

// WithContext sets the context into the *Selector.
func (s *Selector) WithContext(ctx context.Context) *Selector {
	if ctx == nil {
		panic("nil context")
	}
	s.ctx = ctx
	return s
}

// Context returns the Selector context or Background
// if nil.
func (s *Selector) Context() context.Context {
	if s.ctx != nil {
		return s.ctx
	}
	return context.Background()
}

// Select returns a new selector for the `SELECT` statement.
//
//	t1 := Table("users").As("u")
//	t2 := Select().From(Table("groups")).Where(EQ("user_id", 10)).As("g")
//	return Select(t1.C("id"), t2.C("name")).
//			From(t1).
//			Join(t2).
//			On(t1.C("id"), t2.C("user_id"))
func Select(columns ...string) *Selector {
	return (&Selector{}).Select(columns...)
}

// SelectExpr is like Select, but supports passing arbitrary
// expressions for SELECT clause.
func SelectExpr(exprs ...Querier) *Selector {
	return (&Selector{}).SelectExpr(exprs...)
}

// selection represents a column or an expression selection.
type selection struct {
	x  Querier
	c  string
	as string
}

// Select changes the columns selection of the SELECT statement.
// Empty selection means all columns *.
func (s *Selector) Select(columns ...string) *Selector {
	s.selection = make([]selection, len(columns))
	for i := range columns {
		s.selection[i] = selection{c: columns[i]}
	}
	return s
}

// SelectDistinct selects distinct columns.
func (s *Selector) SelectDistinct(columns ...string) *Selector {
	return s.Select(columns...).Distinct()
}

// AppendSelect appends additional columns to the SELECT statement.
func (s *Selector) AppendSelect(columns ...string) *Selector {
	for i := range columns {
		s.selection = append(s.selection, selection{c: columns[i]})
	}
	return s
}

// AppendSelectAs appends additional column to the SELECT statement with the given alias.
func (s *Selector) AppendSelectAs(column, as string) *Selector {
	s.selection = append(s.selection, selection{c: column, as: as})
	return s
}

// SelectExpr changes the columns selection of the SELECT statement
// with custom list of expressions.
func (s *Selector) SelectExpr(exprs ...Querier) *Selector {
	s.selection = make([]selection, len(exprs))
	for i := range exprs {
		s.selection[i] = selection{x: exprs[i]}
	}
	return s
}

// AppendSelectExpr appends additional expressions to the SELECT statement.
func (s *Selector) AppendSelectExpr(exprs ...Querier) *Selector {
	for i := range exprs {
		s.selection = append(s.selection, selection{x: exprs[i]})
	}
	return s
}

// AppendSelectExprAs appends additional expressions to the SELECT statement with the given name.
func (s *Selector) AppendSelectExprAs(expr Querier, as string) *Selector {
	x := expr
	if _, ok := expr.(*raw); !ok {
		x = ExprFunc(func(b *Builder) {
			b.S("(").Join(expr).S(")")
		})
	}
	s.selection = append(s.selection, selection{
		x:  x,
		as: as,
	})
	return s
}

// FindSelection returns all occurrences in the selection that match the given column name.
// For example, for column "a" the following match: a, "a", "t"."a", "t"."b" AS "a".
func (s *Selector) FindSelection(name string) (matches []string) {
	matchC := func(qualified string) bool {
		switch ident, pg := s.isIdent(qualified), s.postgres(); {
		case !ident:
			if i := strings.IndexRune(qualified, '.'); i > 0 {
				return qualified[i+1:] == name
			}
		case ident && pg:
			if i := strings.Index(qualified, `"."`); i > 0 {
				return s.unquote(qualified[i+2:]) == name
			}
		case ident:
			if i := strings.Index(qualified, "`.`"); i > 0 {
				return s.unquote(qualified[i+2:]) == name
			}
		}
		return false
	}
	for _, c := range s.selection {
		switch {
		// Match aliases.
		case c.as != "":
			if ident := s.isIdent(c.as); !ident && c.as == name || ident && s.unquote(c.as) == name {
				matches = append(matches, c.as)
			}
		// Match qualified columns.
		case c.c != "" && s.isQualified(c.c) && matchC(c.c):
			matches = append(matches, c.c)
		// Match unqualified columns.
		case c.c != "" && (c.c == name || s.isIdent(c.c) && s.unquote(c.c) == name):
			matches = append(matches, c.c)
		}
	}
	return matches
}

// SelectedColumns returns the selected columns in the Selector.
func (s *Selector) SelectedColumns() []string {
	columns := make([]string, 0, len(s.selection))
	for i := range s.selection {
		if c := s.selection[i].c; c != "" {
			columns = append(columns, c)
		}
	}
	return columns
}

// UnqualifiedColumns returns an unqualified version of the
// selected columns in the Selector. e.g. "t1"."c" => "c".
func (s *Selector) UnqualifiedColumns() []string {
	columns := make([]string, 0, len(s.selection))
	for i := range s.selection {
		c := s.selection[i].c
		if c == "" {
			continue
		}
		if s.isIdent(c) {
			parts := strings.FieldsFunc(c, func(r rune) bool {
				return r == '`' || r == '"'
			})
			if n := len(parts); n > 0 && parts[n-1] != "" {
				c = parts[n-1]
			}
		}
		columns = append(columns, c)
	}
	return columns
}

// From sets the source of `FROM` clause.
func (s *Selector) From(t TableView) *Selector {
	s.from = nil
	return s.AppendFrom(t)
}

// AppendFrom appends a new TableView to the `FROM` clause.
func (s *Selector) AppendFrom(t TableView) *Selector {
	s.from = append(s.from, t)
	if st, ok := t.(state); ok {
		st.SetDialect(s.dialect)
	}
	return s
}

// FromExpr sets the expression of `FROM` clause.
func (s *Selector) FromExpr(x Querier) *Selector {
	s.from = nil
	return s.AppendFromExpr(x)
}

// AppendFromExpr appends an expression (Queries) to the `FROM` clause.
func (s *Selector) AppendFromExpr(x Querier) *Selector {
	s.from = append(s.from, &queryView{Querier: x})
	if st, ok := x.(state); ok {
		st.SetDialect(s.dialect)
	}
	return s
}

// Distinct adds the DISTINCT keyword to the `SELECT` statement.
func (s *Selector) Distinct() *Selector {
	s.distinct = true
	return s
}

// SetDistinct sets explicitly if the returned rows are distinct or indistinct.
func (s *Selector) SetDistinct(v bool) *Selector {
	s.distinct = v
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

// CollectPredicates indicates the appended predicated should be collected
// and not appended to the `WHERE` clause.
func (s *Selector) CollectPredicates() *Selector {
	s.collected = append(s.collected, []*Predicate{})
	return s
}

// CollectedPredicates returns the collected predicates.
func (s *Selector) CollectedPredicates() []*Predicate {
	if len(s.collected) == 0 {
		return nil
	}
	return s.collected[len(s.collected)-1]
}

// UncollectedPredicates stop collecting predicates.
func (s *Selector) UncollectedPredicates() *Selector {
	if len(s.collected) > 0 {
		s.collected = s.collected[:len(s.collected)-1]
	}
	return s
}

// Where sets or appends the given predicate to the statement.
func (s *Selector) Where(p *Predicate) *Selector {
	if len(s.collected) > 0 {
		s.collected[len(s.collected)-1] = append(s.collected[len(s.collected)-1], p)
		return s
	}
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
		s.where = And(s.where, p)
	}
	return s
}

// P returns the predicate of a selector.
func (s *Selector) P() *Predicate {
	return s.where
}

// SetP sets explicitly the predicate function for the selector and clear its previous state.
func (s *Selector) SetP(p *Predicate) *Selector {
	s.where = p
	s.or = false
	s.not = false
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
	if len(s.from) == 0 {
		return nil
	}
	return selectTable(s.from[0])
}

// selectTable returns a *SelectTable from the given TableView.
func selectTable(t TableView) *SelectTable {
	if t == nil {
		return nil
	}
	switch view := t.(type) {
	case *SelectTable:
		return view
	case *Selector:
		if len(view.from) == 0 {
			return nil
		}
		return selectTable(view.from[0])
	case *queryView, *WithBuilder:
		return nil
	default:
		panic(fmt.Sprintf("unexpected TableView %T", t))
	}
}

// TableName returns the name of the selected table or alias of selector.
func (s *Selector) TableName() string {
	switch view := s.from[0].(type) {
	case *SelectTable:
		return view.name
	case *Selector:
		return view.as
	default:
		panic(fmt.Sprintf("unhandled TableView type %T", s.from))
	}
}

// HasJoins reports if the selector has any JOINs.
func (s *Selector) HasJoins() bool {
	return len(s.joins) > 0
}

// JoinedTable returns the first joined table with the given name.
func (s *Selector) JoinedTable(name string) (*SelectTable, bool) {
	for _, j := range s.joins {
		if t := selectTable(j.table); t != nil && t.name == name {
			return t, true
		}
	}
	return nil, false
}

// JoinedTableView returns the first joined TableView with the given name or alias.
func (s *Selector) JoinedTableView(name string) (TableView, bool) {
	for _, j := range s.joins {
		switch t := j.table.(type) {
		case *SelectTable:
			if t.name == name || t.as == name {
				return t, true
			}
		case *Selector:
			if t.as == name {
				return t, true
			}
			for _, t2 := range t.from {
				if t3 := selectTable(t2); t3 != nil && (t3.name == name || t3.as == name) {
					return t3, true
				}
			}
		}
	}
	return nil, false
}

// Join appends a `JOIN` clause to the statement.
func (s *Selector) Join(t TableView) *Selector {
	return s.join("JOIN", t)
}

// LeftJoin appends a `LEFT JOIN` clause to the statement.
func (s *Selector) LeftJoin(t TableView) *Selector {
	return s.join("LEFT JOIN", t)
}

// RightJoin appends a `RIGHT JOIN` clause to the statement.
func (s *Selector) RightJoin(t TableView) *Selector {
	return s.join("RIGHT JOIN", t)
}

// FullJoin appends a `FULL JOIN` clause to the statement.
func (s *Selector) FullJoin(t TableView) *Selector {
	return s.join("FULL JOIN", t)
}

// join adds a join table to the selector with the given kind.
func (s *Selector) join(kind string, t TableView) *Selector {
	s.joins = append(s.joins, join{
		kind:  kind,
		table: t,
	})
	switch view := t.(type) {
	case *SelectTable:
		if view.as == "" {
			view.as = "t" + strconv.Itoa(len(s.joins))
		}
	case *Selector:
		if view.as == "" {
			view.as = "t" + strconv.Itoa(len(s.joins))
		}
	}
	if st, ok := t.(state); ok {
		st.SetDialect(s.dialect)
	}
	return s
}

type (
	// setOp represents a set/compound operation.
	setOp struct {
		Type      setOpType // Set operation type.
		All       bool      // Quantifier was set to ALL (defaults to DISTINCT).
		TableView           // Query or table to operate on.
	}
	// setOpType is a set operation type.
	setOpType string
)

const (
	setOpTypeUnion     setOpType = "UNION"
	setOpTypeExcept    setOpType = "EXCEPT"
	setOpTypeIntersect setOpType = "INTERSECT"
)

// Union appends the UNION (DISTINCT) clause to the query.
func (s *Selector) Union(t TableView) *Selector {
	if s1, ok := t.(*Selector); ok && s == s1 {
		s.AddError(errors.New("self UNION is not supported. Create a clone or a new selector instead"))
		return s
	}
	s.setOps = append(s.setOps, setOp{
		Type:      setOpTypeUnion,
		TableView: t,
	})
	return s
}

// UnionAll appends the UNION ALL clause to the query.
func (s *Selector) UnionAll(t TableView) *Selector {
	s.setOps = append(s.setOps, setOp{
		Type:      setOpTypeUnion,
		All:       true,
		TableView: t,
	})
	return s
}

// UnionDistinct appends the UNION DISTINCT clause to the query.
// Deprecated: use Union instead as by default, duplicate rows
// are eliminated unless ALL is specified.
func (s *Selector) UnionDistinct(t TableView) *Selector {
	return s.Union(t)
}

// Except appends the EXCEPT clause to the query.
func (s *Selector) Except(t TableView) *Selector {
	s.setOps = append(s.setOps, setOp{
		Type:      setOpTypeExcept,
		TableView: t,
	})
	return s
}

// ExceptAll appends the EXCEPT ALL clause to the query.
func (s *Selector) ExceptAll(t TableView) *Selector {
	if s.sqlite() {
		s.AddError(errors.New("EXCEPT ALL is not supported by SQLite"))
	} else {
		s.setOps = append(s.setOps, setOp{
			Type:      setOpTypeExcept,
			All:       true,
			TableView: t,
		})
	}
	return s
}

// Intersect appends the INTERSECT clause to the query.
func (s *Selector) Intersect(t TableView) *Selector {
	s.setOps = append(s.setOps, setOp{
		Type:      setOpTypeIntersect,
		TableView: t,
	})
	return s
}

// IntersectAll appends the INTERSECT ALL clause to the query.
func (s *Selector) IntersectAll(t TableView) *Selector {
	if s.sqlite() {
		s.AddError(errors.New("INTERSECT ALL is not supported by SQLite"))
	} else {
		s.setOps = append(s.setOps, setOp{
			Type:      setOpTypeIntersect,
			All:       true,
			TableView: t,
		})
	}
	return s
}

// Prefix prefixes the query with list of queries.
func (s *Selector) Prefix(queries ...Querier) *Selector {
	s.prefix = append(s.prefix, queries...)
	return s
}

// C returns a formatted string for a selected column from this statement.
func (s *Selector) C(column string) string {
	// Skip formatting qualified columns.
	if s.isQualified(column) {
		return column
	}
	if s.as != "" {
		b := &Builder{dialect: s.dialect}
		b.Ident(s.as)
		b.WriteByte('.')
		b.Ident(column)
		return b.String()
	}
	return s.Table().C(column)
}

// Columns returns a list of formatted strings for a selected columns from this statement.
func (s *Selector) Columns(columns ...string) []string {
	names := make([]string, 0, len(columns))
	for _, c := range columns {
		names = append(names, s.C(c))
	}
	return names
}

// OnP sets or appends the given predicate for the `ON` clause of the statement.
func (s *Selector) OnP(p *Predicate) *Selector {
	if len(s.joins) > 0 {
		join := &s.joins[len(s.joins)-1]
		switch {
		case join.on == nil:
			join.on = p
		default:
			join.on = And(join.on, p)
		}
	}
	return s
}

// On sets the `ON` clause for the `JOIN` operation.
func (s *Selector) On(c1, c2 string) *Selector {
	s.OnP(P(func(builder *Builder) {
		builder.Ident(c1).WriteOp(OpEQ).Ident(c2)
	}))
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
		b := &Builder{}
		b.IdentComma(columns...)
		column = b.String()
	}
	s.Select(Count(column))
	return s
}

// LockAction tells the transaction what to do in case of
// requesting a row that is locked by other transaction.
type LockAction string

const (
	// NoWait means never wait and returns an error.
	NoWait LockAction = "NOWAIT"
	// SkipLocked means never wait and skip.
	SkipLocked LockAction = "SKIP LOCKED"
)

// LockStrength defines the strength of the lock (see the list below).
type LockStrength string

// A list of all locking clauses.
const (
	LockShare       LockStrength = "SHARE"
	LockUpdate      LockStrength = "UPDATE"
	LockNoKeyUpdate LockStrength = "NO KEY UPDATE"
	LockKeyShare    LockStrength = "KEY SHARE"
)

type (
	// LockOptions defines a SELECT statement
	// lock for protecting concurrent updates.
	LockOptions struct {
		// Strength of the lock.
		Strength LockStrength
		// Action of the lock.
		Action LockAction
		// Tables are an option tables.
		Tables []string
		// custom clause for locking.
		clause string
	}
	// LockOption allows configuring the LockOptions using functional options.
	LockOption func(*LockOptions)
)

// WithLockAction sets the Action of the lock.
func WithLockAction(action LockAction) LockOption {
	return func(c *LockOptions) {
		c.Action = action
	}
}

// WithLockTables sets the Tables of the lock.
func WithLockTables(tables ...string) LockOption {
	return func(c *LockOptions) {
		c.Tables = tables
	}
}

// WithLockClause allows providing a custom clause for
// locking the statement. For example, in MySQL <= 8.22:
//
//	Select().
//	From(Table("users")).
//	ForShare(
//		WithLockClause("LOCK IN SHARE MODE"),
//	)
func WithLockClause(clause string) LockOption {
	return func(c *LockOptions) {
		c.clause = clause
	}
}

// For sets the lock configuration for suffixing the `SELECT`
// statement with the `FOR [SHARE | UPDATE] ...` clause.
func (s *Selector) For(l LockStrength, opts ...LockOption) *Selector {
	if s.Dialect() == dialect.SQLite {
		s.AddError(errors.New("sql: SELECT .. FOR UPDATE/SHARE not supported in SQLite"))
	}
	s.lock = &LockOptions{Strength: l}
	for _, opt := range opts {
		opt(s.lock)
	}
	return s
}

// ForShare sets the lock configuration for suffixing the
// `SELECT` statement with the `FOR SHARE` clause.
func (s *Selector) ForShare(opts ...LockOption) *Selector {
	return s.For(LockShare, opts...)
}

// ForUpdate sets the lock configuration for suffixing the
// `SELECT` statement with the `FOR UPDATE` clause.
func (s *Selector) ForUpdate(opts ...LockOption) *Selector {
	return s.For(LockUpdate, opts...)
}

// Clone returns a duplicate of the selector, including all associated steps. It can be
// used to prepare common SELECT statements and use them differently after the clone is made.
func (s *Selector) Clone() *Selector {
	if s == nil {
		return nil
	}
	joins := make([]join, len(s.joins))
	for i := range s.joins {
		joins[i] = s.joins[i].clone()
	}
	return &Selector{
		Builder:   s.Builder.clone(),
		ctx:       s.ctx,
		as:        s.as,
		or:        s.or,
		not:       s.not,
		from:      s.from,
		limit:     s.limit,
		offset:    s.offset,
		distinct:  s.distinct,
		where:     s.where.clone(),
		having:    s.having.clone(),
		joins:     append([]join{}, joins...),
		group:     append([]string{}, s.group...),
		order:     append([]any{}, s.order...),
		selection: append([]selection{}, s.selection...),
	}
}

// Asc adds the ASC suffix for the given column.
func Asc(column string) string {
	b := &Builder{}
	b.Ident(column).WriteString(" ASC")
	return b.String()
}

// Desc adds the DESC suffix for the given column.
func Desc(column string) string {
	b := &Builder{}
	b.Ident(column).WriteString(" DESC")
	return b.String()
}

// DescExpr returns a new expression where the DESC suffix is added.
func DescExpr(x Querier) Querier {
	return ExprFunc(func(b *Builder) {
		b.Join(x)
		b.WriteString(" DESC")
	})
}

// OrderBy appends the `ORDER BY` clause to the `SELECT` statement.
func (s *Selector) OrderBy(columns ...string) *Selector {
	for i := range columns {
		s.order = append(s.order, columns[i])
	}
	return s
}

// OrderColumns returns the ordered columns in the Selector.
// Note, this function skips columns selected with expressions.
func (s *Selector) OrderColumns() []string {
	columns := make([]string, 0, len(s.order))
	for i := range s.order {
		if c, ok := s.order[i].(string); ok {
			columns = append(columns, c)
		}
	}
	return columns
}

// OrderExpr appends the `ORDER BY` clause to the `SELECT`
// statement with custom list of expressions.
func (s *Selector) OrderExpr(exprs ...Querier) *Selector {
	for i := range exprs {
		s.order = append(s.order, exprs[i])
	}
	return s
}

// OrderExprFunc appends the `ORDER BY` expression that evaluates
// the given function.
func (s *Selector) OrderExprFunc(f func(*Builder)) *Selector {
	return s.OrderExpr(
		Dialect(s.Dialect()).Expr(f),
	)
}

// ClearOrder clears the ORDER BY clause to be empty.
func (s *Selector) ClearOrder() *Selector {
	s.order = nil
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
func (s *Selector) Query() (string, []any) {
	b := s.Builder.clone()
	s.joinPrefix(&b)
	b.WriteString("SELECT ")
	if s.distinct {
		b.WriteString("DISTINCT ")
	}
	if len(s.selection) > 0 {
		s.joinSelect(&b)
	} else {
		b.WriteString("*")
	}
	if len(s.from) > 0 {
		b.WriteString(" FROM ")
	}
	for i, from := range s.from {
		if i > 0 {
			b.Comma()
		}
		switch t := from.(type) {
		case *SelectTable:
			t.SetDialect(s.dialect)
			b.WriteString(t.ref())
		case *Selector:
			t.SetDialect(s.dialect)
			b.Wrap(func(b *Builder) {
				b.Join(t)
			})
			if t.as != "" {
				b.WriteString(" AS ")
				b.Ident(t.as)
			}
		case *WithBuilder:
			t.SetDialect(s.dialect)
			b.Ident(t.Name())
		case *queryView:
			b.Join(t.Querier)
		}
	}
	for _, join := range s.joins {
		b.WriteString(" " + join.kind + " ")
		switch view := join.table.(type) {
		case *SelectTable:
			view.SetDialect(s.dialect)
			b.WriteString(view.ref())
		case *Selector:
			view.SetDialect(s.dialect)
			b.Wrap(func(b *Builder) {
				b.Join(view)
			})
			b.WriteString(" AS ")
			b.Ident(view.as)
		case *WithBuilder:
			view.SetDialect(s.dialect)
			b.Ident(view.Name())
		}
		if join.on != nil {
			b.WriteString(" ON ")
			b.Join(join.on)
		}
	}
	if s.where != nil {
		b.WriteString(" WHERE ")
		b.Join(s.where)
	}
	if len(s.group) > 0 {
		b.WriteString(" GROUP BY ")
		b.IdentComma(s.group...)
	}
	if s.having != nil {
		b.WriteString(" HAVING ")
		b.Join(s.having)
	}
	if len(s.setOps) > 0 {
		s.joinSetOps(&b)
	}
	joinOrder(s.order, &b)
	if s.limit != nil {
		b.WriteString(" LIMIT ")
		b.WriteString(strconv.Itoa(*s.limit))
	}
	if s.offset != nil {
		b.WriteString(" OFFSET ")
		b.WriteString(strconv.Itoa(*s.offset))
	}
	s.joinLock(&b)
	s.total = b.total
	s.AddError(b.Err())
	return b.String(), b.args
}

func (s *Selector) joinPrefix(b *Builder) {
	if len(s.prefix) > 0 {
		b.join(s.prefix, " ")
		b.Pad()
	}
}

func (s *Selector) joinLock(b *Builder) {
	if s.lock == nil {
		return
	}
	b.Pad()
	if s.lock.clause != "" {
		b.WriteString(s.lock.clause)
		return
	}
	b.WriteString("FOR ").WriteString(string(s.lock.Strength))
	if len(s.lock.Tables) > 0 {
		b.WriteString(" OF ").IdentComma(s.lock.Tables...)
	}
	if s.lock.Action != "" {
		b.Pad().WriteString(string(s.lock.Action))
	}
}

func (s *Selector) joinSetOps(b *Builder) {
	for _, op := range s.setOps {
		b.WriteString(" " + string(op.Type) + " ")
		if op.All {
			b.WriteString("ALL ")
		}
		switch view := op.TableView.(type) {
		case *SelectTable:
			view.SetDialect(s.dialect)
			b.WriteString(view.ref())
		case *Selector:
			view.SetDialect(s.dialect)
			b.Join(view)
			if view.as != "" {
				b.WriteString(" AS ")
				b.Ident(view.as)
			}
		}
	}
}

func joinOrder(order []any, b *Builder) {
	if len(order) == 0 {
		return
	}
	b.WriteString(" ORDER BY ")
	for i := range order {
		if i > 0 {
			b.Comma()
		}
		switch r := order[i].(type) {
		case string:
			b.Ident(r)
		case Querier:
			b.Join(r)
		}
	}
}

func joinReturning(columns []string, b *Builder) {
	if len(columns) == 0 || (!b.postgres() && !b.sqlite()) {
		return
	}
	b.WriteString(" RETURNING ")
	b.IdentComma(columns...)
}

func (s *Selector) joinSelect(b *Builder) {
	for i, sc := range s.selection {
		if i > 0 {
			b.Comma()
		}
		switch {
		case sc.c != "":
			b.Ident(sc.c)
		case sc.x != nil:
			b.Join(sc.x)
		}
		if sc.as != "" {
			b.WriteString(" AS ")
			b.Ident(sc.as)
		}
	}
}

// implement the table view interface.
func (*Selector) view() {}

// WithBuilder is the builder for the `WITH` statement.
type WithBuilder struct {
	Builder
	recursive bool
	ctes      []struct {
		name    string
		columns []string
		s       *Selector
	}
}

// With returns a new builder for the `WITH` statement.
//
//	n := Queries{
//		With("users_view").As(Select().From(Table("users"))),
//		Select().From(Table("users_view")),
//	}
//	return n.Query()
func With(name string, columns ...string) *WithBuilder {
	return &WithBuilder{
		ctes: []struct {
			name    string
			columns []string
			s       *Selector
		}{
			{name: name, columns: columns},
		},
	}
}

// WithRecursive returns a new builder for the `WITH RECURSIVE` statement.
//
//	n := Queries{
//		WithRecursive("users_view").As(Select().From(Table("users"))),
//		Select().From(Table("users_view")),
//	}
//	return n.Query()
func WithRecursive(name string, columns ...string) *WithBuilder {
	w := With(name, columns...)
	w.recursive = true
	return w
}

// Name returns the name of the view.
func (w *WithBuilder) Name() string {
	return w.ctes[0].name
}

// As sets the view sub query.
func (w *WithBuilder) As(s *Selector) *WithBuilder {
	w.ctes[len(w.ctes)-1].s = s
	return w
}

// With appends another named CTE to the statement.
func (w *WithBuilder) With(name string, columns ...string) *WithBuilder {
	w.ctes = append(w.ctes, With(name, columns...).ctes...)
	return w
}

// C returns a formatted string for the WITH column.
func (w *WithBuilder) C(column string) string {
	b := &Builder{dialect: w.dialect}
	b.Ident(w.Name()).WriteByte('.').Ident(column)
	return b.String()
}

// Query returns query representation of a `WITH` clause.
func (w *WithBuilder) Query() (string, []any) {
	w.WriteString("WITH ")
	if w.recursive {
		w.WriteString("RECURSIVE ")
	}
	for i, cte := range w.ctes {
		if i > 0 {
			w.Comma()
		}
		w.Ident(cte.name)
		if len(cte.columns) > 0 {
			w.WriteByte('(')
			w.IdentComma(cte.columns...)
			w.WriteByte(')')
		}
		w.WriteString(" AS ")
		w.Wrap(func(b *Builder) {
			b.Join(cte.s)
		})
	}
	return w.String(), w.args
}

// implement the table view interface.
func (*WithBuilder) view() {}

// WindowBuilder represents a builder for a window clause.
// Note that window functions support is limited and used
// only to query rows-limited edges in pagination.
type WindowBuilder struct {
	Builder
	fn        func(*Builder) // e.g. ROW_NUMBER(), RANK()
	partition func(*Builder)
	order     []any
}

// RowNumber returns a new window clause with the ROW_NUMBER() as a function.
// Using this function will assign each row a number, from 1 to N, in the
// order defined by the ORDER BY clause in the window spec.
func RowNumber() *WindowBuilder {
	return Window(func(b *Builder) {
		b.WriteString("ROW_NUMBER()")
	})
}

// Window returns a new window clause with a custom selector allowing
// for custom window functions.
//
//	Window(func(b *Builder) {
//		b.WriteString(Sum(posts.C("duration")))
//	}).PartitionBy("author_id").OrderBy("id"), "duration").
func Window(fn func(*Builder)) *WindowBuilder {
	return &WindowBuilder{fn: fn}
}

// PartitionBy indicates to divide the query rows into groups by the given columns.
// Note that, standard SQL spec allows partition only by columns, and in order to
// use the "expression" version, use the PartitionByExpr.
func (w *WindowBuilder) PartitionBy(columns ...string) *WindowBuilder {
	w.partition = func(b *Builder) {
		b.IdentComma(columns...)
	}
	return w
}

// PartitionExpr indicates to divide the query rows into groups by the given expression.
func (w *WindowBuilder) PartitionExpr(x Querier) *WindowBuilder {
	w.partition = func(b *Builder) {
		b.Join(x)
	}
	return w
}

// OrderBy indicates how to sort rows in each partition.
func (w *WindowBuilder) OrderBy(columns ...string) *WindowBuilder {
	for i := range columns {
		w.order = append(w.order, columns[i])
	}
	return w
}

// OrderExpr appends the `ORDER BY` clause to the window
// partition with custom list of expressions.
func (w *WindowBuilder) OrderExpr(exprs ...Querier) *WindowBuilder {
	for i := range exprs {
		w.order = append(w.order, exprs[i])
	}
	return w
}

// Query returns query representation of the window function.
func (w *WindowBuilder) Query() (string, []any) {
	w.fn(&w.Builder)
	w.WriteString(" OVER ")
	w.Wrap(func(b *Builder) {
		if w.partition != nil {
			b.WriteString("PARTITION BY ")
			w.partition(b)
		}
		joinOrder(w.order, b)
	})
	return w.Builder.String(), w.args
}

// Wrapper wraps a given Querier with different format.
// Used to prefix/suffix other queries.
type Wrapper struct {
	format  string
	wrapped Querier
}

// Query returns query representation of a wrapped Querier.
func (w *Wrapper) Query() (string, []any) {
	query, args := w.wrapped.Query()
	return fmt.Sprintf(w.format, query), args
}

// SetDialect calls SetDialect on the wrapped query.
func (w *Wrapper) SetDialect(name string) {
	if s, ok := w.wrapped.(state); ok {
		s.SetDialect(name)
	}
}

// Dialect calls Dialect on the wrapped query.
func (w *Wrapper) Dialect() string {
	if s, ok := w.wrapped.(state); ok {
		return s.Dialect()
	}
	return ""
}

// Total returns the total number of arguments so far.
func (w *Wrapper) Total() int {
	if s, ok := w.wrapped.(state); ok {
		return s.Total()
	}
	return 0
}

// SetTotal sets the value of the total arguments.
// Used to pass this information between sub queries/expressions.
func (w *Wrapper) SetTotal(total int) {
	if s, ok := w.wrapped.(state); ok {
		s.SetTotal(total)
	}
}

// Raw returns a raw SQL query that is placed as-is in the query.
func Raw(s string) Querier { return &raw{s} }

type raw struct{ s string }

func (r *raw) Query() (string, []any) { return r.s, nil }

// Expr returns an SQL expression that implements the Querier interface.
func Expr(exr string, args ...any) Querier { return &expr{s: exr, args: args} }

type expr struct {
	s    string
	args []any
}

func (e *expr) Query() (string, []any) { return e.s, e.args }

// ExprFunc returns an expression function that implements the Querier interface.
//
//	Update("users").
//		Set("x", ExprFunc(func(b *Builder) {
//			// The sql.Builder config (argc and dialect)
//			// was set before the function was executed.
//			b.Ident("x").WriteOp(OpAdd).Arg(1)
//		}))
func ExprFunc(fn func(*Builder)) Querier {
	return &exprFunc{fn: fn}
}

type exprFunc struct {
	Builder
	fn func(*Builder)
}

func (e *exprFunc) Query() (string, []any) {
	b := e.Builder.clone()
	e.fn(&b)
	return b.Query()
}

// Queries are list of queries join with space between them.
type Queries []Querier

// Query returns query representation of Queriers.
func (n Queries) Query() (string, []any) {
	b := &Builder{}
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
	sb        *strings.Builder // underlying builder.
	dialect   string           // configured dialect.
	args      []any            // query parameters.
	total     int              // total number of parameters in query tree.
	errs      []error          // errors that added during the query construction.
	qualifier string           // qualifier to prefix identifiers (e.g. table name).
}

// Quote quotes the given identifier with the characters based
// on the configured dialect. It defaults to "`".
func (b *Builder) Quote(ident string) string {
	quote := "`"
	switch {
	case b.postgres():
		// If it was quoted with the wrong
		// identifier character.
		if strings.Contains(ident, "`") {
			return strings.ReplaceAll(ident, "`", `"`)
		}
		quote = `"`
	// An identifier for unknown dialect.
	case b.dialect == "" && strings.ContainsAny(ident, "`\""):
		return ident
	}
	return quote + ident + quote
}

// Ident appends the given string as an identifier.
func (b *Builder) Ident(s string) *Builder {
	switch {
	case len(s) == 0:
	case !strings.HasSuffix(s, "*") && !b.isIdent(s) && !isFunc(s) && !isModifier(s) && !isAlias(s):
		if b.qualifier != "" {
			b.WriteString(b.Quote(b.qualifier)).WriteByte('.')
		}
		b.WriteString(b.Quote(s))
	case (isFunc(s) || isModifier(s) || isAlias(s)) && b.postgres():
		// Modifiers and aggregation functions that
		// were called without dialect information.
		b.WriteString(strings.ReplaceAll(s, "`", `"`))
	default:
		b.WriteString(s)
	}
	return b
}

// IdentComma calls Ident on all arguments and adds a comma between them.
func (b *Builder) IdentComma(s ...string) *Builder {
	for i := range s {
		if i > 0 {
			b.Comma()
		}
		b.Ident(s[i])
	}
	return b
}

// String returns the accumulated string.
func (b *Builder) String() string {
	if b.sb == nil {
		return ""
	}
	return b.sb.String()
}

// WriteByte wraps the Buffer.WriteByte to make it chainable with other methods.
func (b *Builder) WriteByte(c byte) *Builder {
	if b.sb == nil {
		b.sb = &strings.Builder{}
	}
	b.sb.WriteByte(c)
	return b
}

// WriteString wraps the Buffer.WriteString to make it chainable with other methods.
func (b *Builder) WriteString(s string) *Builder {
	if b.sb == nil {
		b.sb = &strings.Builder{}
	}
	b.sb.WriteString(s)
	return b
}

// S is a short version of WriteString.
func (b *Builder) S(s string) *Builder {
	return b.WriteString(s)
}

// Len returns the number of accumulated bytes.
func (b *Builder) Len() int {
	if b.sb == nil {
		return 0
	}
	return b.sb.Len()
}

// Reset resets the Builder to be empty.
func (b *Builder) Reset() *Builder {
	if b.sb != nil {
		b.sb.Reset()
	}
	return b
}

// AddError appends an error to the builder errors.
func (b *Builder) AddError(err error) *Builder {
	// allowed nil error make build process easier
	if err != nil {
		b.errs = append(b.errs, err)
	}
	return b
}

func (b *Builder) writeSchema(schema string) {
	if schema != "" && b.dialect != dialect.SQLite {
		b.Ident(schema).WriteByte('.')
	}
}

// Err returns a concatenated error of all errors encountered during
// the query-building, or were added manually by calling AddError.
func (b *Builder) Err() error {
	if len(b.errs) == 0 {
		return nil
	}
	br := strings.Builder{}
	for i := range b.errs {
		if i > 0 {
			br.WriteString("; ")
		}
		br.WriteString(b.errs[i].Error())
	}
	return errors.New(br.String())
}

// An Op represents an operator.
type Op int

// Predicate and arithmetic operators.
const (
	OpEQ      Op = iota // =
	OpNEQ               // <>
	OpGT                // >
	OpGTE               // >=
	OpLT                // <
	OpLTE               // <=
	OpIn                // IN
	OpNotIn             // NOT IN
	OpLike              // LIKE
	OpIsNull            // IS NULL
	OpNotNull           // IS NOT NULL
	OpAdd               // +
	OpSub               // -
	OpMul               // *
	OpDiv               // / (Quotient)
	OpMod               // % (Reminder)
)

var ops = [...]string{
	OpEQ:      "=",
	OpNEQ:     "<>",
	OpGT:      ">",
	OpGTE:     ">=",
	OpLT:      "<",
	OpLTE:     "<=",
	OpIn:      "IN",
	OpNotIn:   "NOT IN",
	OpLike:    "LIKE",
	OpIsNull:  "IS NULL",
	OpNotNull: "IS NOT NULL",
	OpAdd:     "+",
	OpSub:     "-",
	OpMul:     "*",
	OpDiv:     "/",
	OpMod:     "%",
}

// WriteOp writes an operator to the builder.
func (b *Builder) WriteOp(op Op) *Builder {
	switch {
	case op >= OpEQ && op <= OpLike || op >= OpAdd && op <= OpMod:
		b.Pad().WriteString(ops[op]).Pad()
	case op == OpIsNull || op == OpNotNull:
		b.Pad().WriteString(ops[op])
	default:
		panic(fmt.Sprintf("invalid op %d", op))
	}
	return b
}

type (
	// StmtInfo holds an information regarding
	// the statement
	StmtInfo struct {
		// The Dialect of the SQL driver.
		Dialect string
	}
	// ParamFormatter wraps the FormatPram function.
	ParamFormatter interface {
		// The FormatParam function lets users define
		// custom placeholder formatting for their types.
		// For example, formatting the default placeholder
		// from '?' to 'ST_GeomFromWKB(?)' for MySQL dialect.
		FormatParam(placeholder string, info *StmtInfo) string
	}
)

// Arg appends an input argument to the builder.
func (b *Builder) Arg(a any) *Builder {
	switch v := a.(type) {
	case nil:
		b.WriteString("NULL")
		return b
	case *raw:
		b.WriteString(v.s)
		return b
	case Querier:
		b.Join(v)
		return b
	}
	// Default placeholder param (MySQL and SQLite).
	format := "?"
	if b.postgres() {
		// Postgres' arguments are referenced using the syntax $n.
		// $1 refers to the 1st argument, $2 to the 2nd, and so on.
		format = "$" + strconv.Itoa(b.total+1)
	}
	if f, ok := a.(ParamFormatter); ok {
		format = f.FormatParam(format, &StmtInfo{
			Dialect: b.dialect,
		})
	}
	return b.Argf(format, a)
}

// Args appends a list of arguments to the builder.
func (b *Builder) Args(a ...any) *Builder {
	for i := range a {
		if i > 0 {
			b.Comma()
		}
		b.Arg(a[i])
	}
	return b
}

// Argf appends an input argument to the builder
// with the given format. For example:
//
//	FormatArg("JSON(?)", b).
//	FormatArg("ST_GeomFromText(?)", geom)
func (b *Builder) Argf(format string, a any) *Builder {
	switch a := a.(type) {
	case nil:
		b.WriteString("NULL")
		return b
	case *raw:
		b.WriteString(a.s)
		return b
	case Querier:
		b.Join(a)
		return b
	}
	b.total++
	b.args = append(b.args, a)
	b.WriteString(format)
	return b
}

// Comma adds a comma to the query.
func (b *Builder) Comma() *Builder {
	return b.WriteString(", ")
}

// Pad adds a space to the query.
func (b *Builder) Pad() *Builder {
	return b.WriteByte(' ')
}

// Join joins a list of Queries to the builder.
func (b *Builder) Join(qs ...Querier) *Builder {
	return b.join(qs, "")
}

// JoinComma joins a list of Queries and adds comma between them.
func (b *Builder) JoinComma(qs ...Querier) *Builder {
	return b.join(qs, ", ")
}

// join a list of Queries to the builder with a given separator.
func (b *Builder) join(qs []Querier, sep string) *Builder {
	for i, q := range qs {
		if i > 0 {
			b.WriteString(sep)
		}
		st, ok := q.(state)
		if ok {
			st.SetDialect(b.dialect)
			st.SetTotal(b.total)
		}
		query, args := q.Query()
		b.WriteString(query)
		b.args = append(b.args, args...)
		b.total += len(args)
		if qe, ok := q.(querierErr); ok {
			if err := qe.Err(); err != nil {
				b.AddError(err)
			}
		}
	}
	return b
}

// Wrap gets a callback, and wraps its result with parentheses.
func (b *Builder) Wrap(f func(*Builder)) *Builder {
	nb := &Builder{dialect: b.dialect, total: b.total, sb: &strings.Builder{}}
	nb.WriteByte('(')
	f(nb)
	nb.WriteByte(')')
	b.WriteString(nb.String())
	b.args = append(b.args, nb.args...)
	b.total = nb.total
	return b
}

// Nested gets a callback, and wraps its result with parentheses.
//
// Deprecated: Use Builder.Wrap instead.
func (b *Builder) Nested(f func(*Builder)) *Builder {
	return b.Wrap(f)
}

// SetDialect sets the builder dialect. It's used for garnering dialect specific queries.
func (b *Builder) SetDialect(dialect string) {
	b.dialect = dialect
}

// Dialect returns the dialect of the builder.
func (b Builder) Dialect() string {
	return b.dialect
}

// Total returns the total number of arguments so far.
func (b Builder) Total() int {
	return b.total
}

// SetTotal sets the value of the total arguments.
// Used to pass this information between sub queries/expressions.
func (b *Builder) SetTotal(total int) {
	b.total = total
}

// Query implements the Querier interface.
func (b Builder) Query() (string, []any) {
	return b.String(), b.args
}

// clone returns a shallow clone of a builder.
func (b Builder) clone() Builder {
	c := Builder{dialect: b.dialect, total: b.total, sb: &strings.Builder{}}
	if len(b.args) > 0 {
		c.args = append(c.args, b.args...)
	}
	if b.sb != nil {
		c.sb.WriteString(b.sb.String())
	}
	return c
}

// postgres reports if the builder dialect is PostgreSQL.
func (b Builder) postgres() bool {
	return b.Dialect() == dialect.Postgres
}

// sqlite reports if the builder dialect is SQLite.
func (b Builder) sqlite() bool {
	return b.Dialect() == dialect.SQLite
}

// fromIdent sets the builder dialect from the identifier format.
func (b *Builder) fromIdent(ident string) {
	if strings.Contains(ident, `"`) {
		b.SetDialect(dialect.Postgres)
	}
	// otherwise, use the default.
}

// isIdent reports if the given string is a dialect identifier.
func (b *Builder) isIdent(s string) bool {
	switch {
	case b.postgres():
		return strings.Contains(s, `"`)
	default:
		return strings.Contains(s, "`")
	}
}

// unquote database identifiers.
func (b *Builder) unquote(s string) string {
	switch pg := b.postgres(); {
	case len(s) < 2:
	case !pg && s[0] == '`' && s[len(s)-1] == '`', pg && s[0] == '"' && s[len(s)-1] == '"':
		if u, err := strconv.Unquote(s); err == nil {
			return u
		}
	}
	return s
}

// isQualified reports if the given string is a qualified identifier.
func (b *Builder) isQualified(s string) bool {
	ident, pg := b.isIdent(s), b.postgres()
	return !ident && len(s) > 2 && strings.ContainsRune(s[1:len(s)-1], '.') || // <qualifier>.<column>
		ident && pg && strings.Contains(s, `"."`) || // "qualifier"."column"
		ident && !pg && strings.Contains(s, "`.`") // `qualifier`.`column`
}

// state wraps the all methods for setting and getting
// update state between all queries in the query tree.
type state interface {
	Dialect() string
	SetDialect(string)
	Total() int
	SetTotal(int)
}

// DialectBuilder prefixes all root builders with the `Dialect` constructor.
type DialectBuilder struct {
	dialect string
}

// Dialect creates a new DialectBuilder with the given dialect name.
func Dialect(name string) *DialectBuilder {
	return &DialectBuilder{name}
}

// String builds a dialect-aware expression string from the given callback.
func (d *DialectBuilder) String(f func(*Builder)) string {
	b := &Builder{}
	b.SetDialect(d.dialect)
	f(b)
	return b.String()
}

// Expr builds a dialect-aware expression from the given callback.
func (d *DialectBuilder) Expr(f func(*Builder)) Querier {
	return Expr(d.String(f))
}

// Describe creates a DescribeBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		Describe("users")
func (d *DialectBuilder) Describe(name string) *DescribeBuilder {
	b := Describe(name)
	b.SetDialect(d.dialect)
	return b
}

// CreateTable creates a TableBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		CreateTable("users").
//			Columns(
//				Column("id").Type("int").Attr("auto_increment"),
//				Column("name").Type("varchar(255)"),
//			).
//			PrimaryKey("id")
func (d *DialectBuilder) CreateTable(name string) *TableBuilder {
	b := CreateTable(name)
	b.SetDialect(d.dialect)
	return b
}

// AlterTable creates a TableAlter for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		AlterTable("users").
//		AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
//		AddForeignKey(ForeignKey().Columns("group_id").
//			Reference(Reference().Table("groups").Columns("id")).
//			OnDelete("CASCADE"),
//		)
func (d *DialectBuilder) AlterTable(name string) *TableAlter {
	b := AlterTable(name)
	b.SetDialect(d.dialect)
	return b
}

// AlterIndex creates an IndexAlter for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		AlterIndex("old").
//		Rename("new")
func (d *DialectBuilder) AlterIndex(name string) *IndexAlter {
	b := AlterIndex(name)
	b.SetDialect(d.dialect)
	return b
}

// Column creates a ColumnBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres)..
//		Column("group_id").Type("int").Attr("UNIQUE")
func (d *DialectBuilder) Column(name string) *ColumnBuilder {
	b := Column(name)
	b.SetDialect(d.dialect)
	return b
}

// Insert creates a InsertBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		Insert("users").Columns("age").Values(1)
func (d *DialectBuilder) Insert(table string) *InsertBuilder {
	b := Insert(table)
	b.SetDialect(d.dialect)
	return b
}

// Update creates a UpdateBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		Update("users").Set("name", "foo")
func (d *DialectBuilder) Update(table string) *UpdateBuilder {
	b := Update(table)
	b.SetDialect(d.dialect)
	return b
}

// Delete creates a DeleteBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		Delete().From("users")
func (d *DialectBuilder) Delete(table string) *DeleteBuilder {
	b := Delete(table)
	b.SetDialect(d.dialect)
	return b
}

// Select creates a Selector for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		Select().From(Table("users"))
func (d *DialectBuilder) Select(columns ...string) *Selector {
	b := Select(columns...)
	b.SetDialect(d.dialect)
	return b
}

// SelectExpr is like Select, but supports passing arbitrary
// expressions for SELECT clause.
//
//	Dialect(dialect.Postgres).
//		SelectExpr(expr...).
//		From(Table("users"))
func (d *DialectBuilder) SelectExpr(exprs ...Querier) *Selector {
	b := SelectExpr(exprs...)
	b.SetDialect(d.dialect)
	return b
}

// Table creates a SelectTable for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		Table("users").As("u")
func (d *DialectBuilder) Table(name string) *SelectTable {
	b := Table(name)
	b.SetDialect(d.dialect)
	return b
}

// With creates a WithBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		With("users_view").
//		As(Select().From(Table("users")))
func (d *DialectBuilder) With(name string) *WithBuilder {
	b := With(name)
	b.SetDialect(d.dialect)
	return b
}

// CreateIndex creates a IndexBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		CreateIndex("unique_name").
//		Unique().
//		Table("users").
//		Columns("first", "last")
func (d *DialectBuilder) CreateIndex(name string) *IndexBuilder {
	b := CreateIndex(name)
	b.SetDialect(d.dialect)
	return b
}

// DropIndex creates a DropIndexBuilder for the configured dialect.
//
//	Dialect(dialect.Postgres).
//		DropIndex("name")
func (d *DialectBuilder) DropIndex(name string) *DropIndexBuilder {
	b := DropIndex(name)
	b.SetDialect(d.dialect)
	return b
}

func isAlias(s string) bool {
	return strings.Contains(s, " AS ") || strings.Contains(s, " as ")
}

func isFunc(s string) bool {
	return strings.Contains(s, "(") && strings.Contains(s, ")")
}

func isModifier(s string) bool {
	for _, m := range [...]string{"DISTINCT", "ALL", "WITH ROLLUP"} {
		if strings.HasPrefix(s, m) {
			return true
		}
	}
	return false
}
