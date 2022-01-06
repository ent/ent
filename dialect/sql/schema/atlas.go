// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
)

type (
	// Differ is the interface that wraps the Diff method.
	Differ interface {
		// Diff creates the given tables in the database.
		Diff(current, desired *schema.Schema) ([]schema.Change, error)
	}

	// The DiffFunc type is an adapter to allow the use of ordinary function as Differ.
	// If f is a function with the appropriate signature, DiffFunc(f) is a Differ that calls f.
	DiffFunc func(current, desired *schema.Schema) ([]schema.Change, error)

	// DiffHook defines the "diff middleware". A function that gets a Differ and returns a Differ.
	DiffHook func(Differ) Differ
)

// Diff calls f(current, desired).
func (f DiffFunc) Diff(current, desired *schema.Schema) ([]schema.Change, error) {
	return f(current, desired)
}

// WithDiffHook adds a list of DiffHook to the schema migration.
//
//	schema.WithDiffHook(func(next schema.Differ) schema.Differ {
//		return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
//			// Code before standard diff.
//			changes, err := next.Diff(current, desired)
//			if err != nil {
//				return nil, err
//			}
//			// After diff, you can filter
//			// changes or return new ones.
//			return changes, nil
//		})
//	})
//
func WithDiffHook(hooks ...DiffHook) MigrateOption {
	return func(m *Migrate) {
		m.atlas.diff = append(m.atlas.diff, hooks...)
	}
}

// SkipChanges allows skipping/filtering list of changes
// returned by the differ before executing migration planning.
//
//	SkipChanges(schema.DropTable|schema.DropColumn)
//
func WithSkipChanges(skip ChangeKind) MigrateOption {
	return func(m *Migrate) {
		m.atlas.skip = skip
	}
}

// A Change of schema.
type ChangeKind uint

// List of change types.
const (
	NoChange  ChangeKind = 0
	AddSchema ChangeKind = 1 << (iota - 1)
	ModifySchema
	DropSchema
	AddTable
	ModifyTable
	DropTable
	AddColumn
	ModifyColumn
	DropColumn
	AddIndex
	ModifyIndex
	DropIndex
	AddForeignKey
	ModifyForeignKey
	DropForeignKey
	AddCheck
	ModifyCheck
	DropCheck
)

// Is reports whether c is match the given change king.
func (k ChangeKind) Is(c ChangeKind) bool {
	return k == c || k&c != 0
}

// filterChanges is a DiffHook for filtering changes before plan.
func filterChanges(skip ChangeKind) DiffHook {
	return func(next Differ) Differ {
		return DiffFunc(func(current, desired *schema.Schema) ([]schema.Change, error) {
			var f func([]schema.Change) []schema.Change
			f = func(changes []schema.Change) (keep []schema.Change) {
				var k ChangeKind
				for _, c := range changes {
					switch c := c.(type) {
					case *schema.AddSchema:
						k = AddSchema
					case *schema.ModifySchema:
						k = ModifySchema
						if !skip.Is(k) {
							c.Changes = f(c.Changes)
						}
					case *schema.DropSchema:
						k = DropSchema
					case *schema.AddTable:
						k = AddTable
					case *schema.ModifyTable:
						k = ModifyTable
						if !skip.Is(k) {
							c.Changes = f(c.Changes)
						}
					case *schema.DropTable:
						k = DropTable
					case *schema.AddColumn:
						k = AddColumn
					case *schema.ModifyColumn:
						k = ModifyColumn
					case *schema.DropColumn:
						k = DropColumn
					case *schema.AddIndex:
						k = AddIndex
					case *schema.ModifyIndex:
						k = ModifyIndex
					case *schema.DropIndex:
						k = DropIndex
					case *schema.AddForeignKey:
						k = AddIndex
					case *schema.ModifyForeignKey:
						k = ModifyForeignKey
					case *schema.DropForeignKey:
						k = DropForeignKey
					case *schema.AddCheck:
						k = AddCheck
					case *schema.ModifyCheck:
						k = ModifyCheck
					case *schema.DropCheck:
						k = DropCheck
					}
					if !skip.Is(k) {
						keep = append(keep, c)
					}
				}
				return
			}
			changes, err := next.Diff(current, desired)
			if err != nil {
				return nil, err
			}
			return f(changes), nil
		})
	}
}

type (
	// Applier is the interface that wraps the Apply method.
	Applier interface {
		// Diff creates the given tables in the database.
		Apply(context.Context, dialect.ExecQuerier, *migrate.Plan) error
	}

	// The ApplyFunc type is an adapter to allow the use of ordinary function as Applier.
	// If f is a function with the appropriate signature, ApplyFunc(f) is a Applier that calls f.
	ApplyFunc func(context.Context, dialect.ExecQuerier, *migrate.Plan) error

	// ApplyHook defines the "migration applying middleware". A function that gets a Applier and returns a Applier.
	ApplyHook func(Applier) Applier
)

// Apply calls f(ctx, tables...).
func (f ApplyFunc) Apply(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
	return f(ctx, conn, plan)
}

// func adds a list of ApplyHook to the schema migration.
//
//	schema.WithApplyHook(func(next schema.Applier) schema.Applier {
//		return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
//			// Example to hook into the apply process, or implement
//			// a custom applier.
//			//
//			//	for _, c := range plan.Changes {
//			//		fmt.Printf("%s: %s", c.Comment, c.Cmd)
//			//	}
//			//
//			return next.Apply(ctx, conn, plan)
//		})
//	})
//
func WithApplyHook(hooks ...ApplyHook) MigrateOption {
	return func(m *Migrate) {
		m.atlas.apply = append(m.atlas.apply, hooks...)
	}
}

// WithAtlas is an opt-in option for v0.10 indicates the migration
// should be executed using Atlas engine (i.e. https://atlasgo.io).
// Note, in future versions, this option is going to be replaced
// from opt-in to opt-out and the deprecation of this package.
func WithAtlas(b bool) MigrateOption {
	return func(m *Migrate) {
		m.atlas.enabled = b
	}
}

type (
	// atlasOptions describes the options for atlas.
	atlasOptions struct {
		enabled bool
		diff    []DiffHook
		apply   []ApplyHook
		skip    ChangeKind
	}

	// atBuilder must be implemented by the different drivers in
	// order to convert a dialect/sql/schema to atlas/sql/schema.
	atBuilder interface {
		atOpen(dialect.ExecQuerier) (migrate.Driver, error)
		atTable(*Table, *schema.Table)
		atTypeC(*Column, *schema.Column) error
		atUniqueC(*Table, *Column, *schema.Table, *schema.Column)
		atIncrementC(*schema.Table, *schema.Column)
		atIncrementT(*schema.Table, int64)
		atIndex(*Index, *schema.Table, *schema.Index) error
	}
)

func (m *Migrate) setupAtlas() error {
	// Using one of the Atlas options, opt-in to Atlas migration.
	if !m.atlas.enabled && (m.atlas.skip != NoChange || len(m.atlas.diff) > 0 || len(m.atlas.apply) > 0) {
		m.atlas.enabled = true
	}
	if !m.atlas.enabled {
		return nil
	}
	if !m.withForeignKeys {
		return errors.New("sql/schema: WithForeignKeys(false) does not work in Atlas migration")
	}
	if m.withFixture {
		return errors.New("sql/schema: WithFixture(true) does not work in Atlas migration")
	}
	k := DropIndex | DropColumn
	if m.atlas.skip != NoChange {
		k = m.atlas.skip
	}
	if m.dropIndexes {
		k |= ^DropIndex
	}
	if m.dropColumns {
		k |= ^DropColumn
	}
	if k == NoChange {
		m.atlas.diff = append(m.atlas.diff, filterChanges(k))
	}
	return nil
}

func (m *Migrate) atCreate(ctx context.Context, tables ...*Table) error {
	// Open a transaction for backwards compatibility,
	// even if the migration is not transactional.
	tx, err := m.Tx(ctx)
	if err != nil {
		return err
	}
	if err := func() error {
		if err := m.init(ctx, tx); err != nil {
			return err
		}
		if m.universalID {
			if err := m.types(ctx, tx); err != nil {
				return err
			}
		}
		drv, err := m.atOpen(tx)
		if err != nil {
			return err
		}
		current, err := drv.InspectSchema(ctx, "", &schema.InspectOptions{
			Tables: func() (t []string) {
				for i := range tables {
					t = append(t, tables[i].Name)
				}
				return t
			}(),
		})
		if err != nil {
			return err
		}
		tt, err := m.aTables(ctx, m, tx, tables)
		if err != nil {
			return err
		}
		// Diff changes.
		var differ Differ = DiffFunc(drv.SchemaDiff)
		for i := len(m.atlas.diff) - 1; i >= 0; i-- {
			differ = m.atlas.diff[i](differ)
		}
		changes, err := differ.Diff(current, &schema.Schema{Name: current.Name, Attrs: current.Attrs, Tables: tt})
		if err != nil {
			return err
		}
		// Plan changes.
		plan, err := drv.PlanChanges(ctx, "plan", changes)
		if err != nil {
			return err
		}
		// Apply plan (changes).
		var applier Applier = ApplyFunc(func(ctx context.Context, tx dialect.ExecQuerier, plan *migrate.Plan) error {
			for _, c := range plan.Changes {
				if err := tx.Exec(ctx, c.Cmd, c.Args, nil); err != nil {
					if c.Comment != "" {
						err = fmt.Errorf("%s: %w", c.Comment, err)
					}
					return err
				}
			}
			return nil
		})
		for i := len(m.atlas.apply) - 1; i >= 0; i-- {
			applier = m.atlas.apply[i](applier)
		}
		return applier.Apply(ctx, tx, plan)
	}(); err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

type db struct{ dialect.ExecQuerier }

func (d *db) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows := &entsql.Rows{}
	if err := d.ExecQuerier.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	return rows.ColumnScanner.(*sql.Rows), nil
}

func (d *db) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var r sql.Result
	if err := d.ExecQuerier.Exec(ctx, query, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (m *Migrate) aTables(ctx context.Context, b atBuilder, conn dialect.ExecQuerier, tables1 []*Table) ([]*schema.Table, error) {
	tables2 := make([]*schema.Table, len(tables1))
	for i, t1 := range tables1 {
		t2 := schema.NewTable(t1.Name)
		b.atTable(t1, t2)
		if m.universalID {
			r, err := m.pkRange(ctx, conn, t1)
			if err != nil {
				return nil, err
			}
			b.atIncrementT(t2, r)
		}
		if err := m.aColumns(b, t1, t2); err != nil {
			return nil, err
		}
		if err := m.aIndexes(b, t1, t2); err != nil {
			return nil, err
		}
		tables2[i] = t2
	}
	for i, t1 := range tables1 {
		t2 := tables2[i]
		for _, fk1 := range t1.ForeignKeys {
			fk2 := schema.NewForeignKey(fk1.Symbol).
				SetTable(t2).
				SetOnUpdate(schema.ReferenceOption(fk1.OnUpdate)).
				SetOnDelete(schema.ReferenceOption(fk1.OnDelete))
			for _, c1 := range fk1.Columns {
				c2, ok := t2.Column(c1.Name)
				if !ok {
					return nil, fmt.Errorf("unexpected fk %q column: %q", fk1.Symbol, c1.Name)
				}
				fk2.AddColumns(c2)
			}
			var refT *schema.Table
			for _, t2 := range tables2 {
				if t2.Name == fk1.RefTable.Name {
					refT = t2
					break
				}
			}
			if refT == nil {
				return nil, fmt.Errorf("unexpected fk %q ref-table: %q", fk1.Symbol, fk1.RefTable.Name)
			}
			fk2.SetRefTable(refT)
			for _, c1 := range fk1.RefColumns {
				c2, ok := refT.Column(c1.Name)
				if !ok {
					return nil, fmt.Errorf("unexpected fk %q ref-column: %q", fk1.Symbol, c1.Name)
				}
				fk2.AddRefColumns(c2)
			}
			t2.AddForeignKeys(fk2)
		}
	}
	return tables2, nil
}

func (m *Migrate) aColumns(b atBuilder, t1 *Table, t2 *schema.Table) error {
	for _, c1 := range t1.Columns {
		c2 := schema.NewColumn(c1.Name).
			SetNull(c1.Nullable)
		if c1.Collation != "" {
			c2.SetCollation(c1.Collation)
		}
		if err := b.atTypeC(c1, c2); err != nil {
			return err
		}
		if c1.Default != nil && c1.supportDefault() {
			// Has default and the database supports adding this default.
			x := fmt.Sprint(c1.Default)
			if v, ok := c1.Default.(string); ok && c1.Type != field.TypeUUID && c1.Type != field.TypeTime {
				// Escape single quote by replacing each with 2.
				x = fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''"))
			}
			c2.SetDefault(&schema.RawExpr{X: x})
		}
		if c1.Unique {
			b.atUniqueC(t1, c1, t2, c2)
		}
		if c1.Increment {
			b.atIncrementC(t2, c2)
		}
		t2.AddColumns(c2)
	}
	return nil
}

func (m *Migrate) aIndexes(b atBuilder, t1 *Table, t2 *schema.Table) error {
	// Primary-key index.
	pk := make([]*schema.Column, 0, len(t1.PrimaryKey))
	for _, c1 := range t1.PrimaryKey {
		c2, ok := t2.Column(c1.Name)
		if !ok {
			return fmt.Errorf("unexpected primary-key column: %q", c1.Name)
		}
		pk = append(pk, c2)
	}
	t2.SetPrimaryKey(schema.NewPrimaryKey(pk...))
	// Rest of indexes.
	for _, idx1 := range t1.Indexes {
		idx2 := schema.NewIndex(idx1.Name).
			SetUnique(idx1.Unique)
		if err := b.atIndex(idx1, t2, idx2); err != nil {
			return err
		}
		t2.AddIndexes(idx2)
	}
	return nil
}

func setAtChecks(t1 *Table, t2 *schema.Table) {
	if check := t1.Annotation.Check; check != "" {
		t2.AddChecks(&schema.Check{
			Expr: check,
		})
	}
	if checks := t1.Annotation.Checks; len(t1.Annotation.Checks) > 0 {
		names := make([]string, 0, len(checks))
		for name := range checks {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			t2.AddChecks(&schema.Check{
				Name: name,
				Expr: checks[name],
			})
		}
	}
}
