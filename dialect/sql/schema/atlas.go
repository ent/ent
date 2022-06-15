// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"sort"
	"strings"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqltool"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

type (
	// Differ is the interface that wraps the Diff method.
	Differ interface {
		// Diff returns a list of changes that construct a migration plan.
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

// WithSkipChanges allows skipping/filtering list of changes
// returned by the Differ before executing migration planning.
//
//	SkipChanges(schema.DropTable|schema.DropColumn)
//
func WithSkipChanges(skip ChangeKind) MigrateOption {
	return func(m *Migrate) {
		m.atlas.skip = skip
	}
}

// A ChangeKind denotes the kind of schema change.
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

// Is reports whether c is match the given change kind.
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

func withoutForeignKeys(next Differ) Differ {
	return DiffFunc(func(current, desired *schema.Schema) ([]schema.Change, error) {
		changes, err := next.Diff(current, desired)
		if err != nil {
			return nil, err
		}
		for _, c := range changes {
			switch c := c.(type) {
			case *schema.AddTable:
				c.T.ForeignKeys = nil
			case *schema.ModifyTable:
				c.T.ForeignKeys = nil
				filtered := make([]schema.Change, 0, len(c.Changes))
				for _, change := range c.Changes {
					switch change.(type) {
					case *schema.AddForeignKey, *schema.DropForeignKey, *schema.ModifyForeignKey:
						continue
					default:
						filtered = append(filtered, change)
					}
				}
				c.Changes = filtered
			}
		}
		return changes, nil
	})
}

type (
	// Applier is the interface that wraps the Apply method.
	Applier interface {
		// Apply applies the given migrate.Plan on the database.
		Apply(context.Context, dialect.ExecQuerier, *migrate.Plan) error
	}

	// The ApplyFunc type is an adapter to allow the use of ordinary function as Applier.
	// If f is a function with the appropriate signature, ApplyFunc(f) is an Applier that calls f.
	ApplyFunc func(context.Context, dialect.ExecQuerier, *migrate.Plan) error

	// ApplyHook defines the "migration applying middleware". A function that gets an Applier and returns an Applier.
	ApplyHook func(Applier) Applier
)

// Apply calls f(ctx, tables...).
func (f ApplyFunc) Apply(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
	return f(ctx, conn, plan)
}

// WithApplyHook adds a list of ApplyHook to the schema migration.
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

// WithDir sets the atlas migration directory to use to store migration files.
func WithDir(dir migrate.Dir) MigrateOption {
	return func(m *Migrate) {
		m.atlas.dir = dir
	}
}

// WithFormatter sets atlas formatter to use to write changes to migration files.
func WithFormatter(fmt migrate.Formatter) MigrateOption {
	return func(m *Migrate) {
		m.atlas.fmt = fmt
	}
}

// WithSumFile instructs atlas to generate a migration directory integrity sum file as well.
func WithSumFile() MigrateOption {
	return func(m *Migrate) {
		m.atlas.genSum = true
	}
}

// WithUniversalID instructs atlas to use a file based type store when
// global unique ids are enabled. For more information see the setupAtlas method on Migrate.
//
// ATTENTION:
// The file based PK range store is not backward compatible, since the allocated ranges were computed
// dynamically when computing the diff between a deployed database and the current schema. In cases where there
// exist multiple deployments, the allocated ranges for the same type might be different from each other,
// depending on when the deployment took part.
func WithUniversalID() MigrateOption {
	return func(m *Migrate) {
		m.universalID = true
		m.atlas.typeStoreConsent = true
	}
}

type (
	// atlasOptions describes the options for atlas.
	atlasOptions struct {
		enabled          bool
		diff             []DiffHook
		apply            []ApplyHook
		skip             ChangeKind
		dir              migrate.Dir
		fmt              migrate.Formatter
		genSum           bool
		typeStoreConsent bool
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
		atTypeRangeSQL(t ...string) string
	}
)

var errConsent = errors.New("sql/schema: use WithUniversalID() instead of WithGlobalUniqueID(true) when using WithDir(): https://entgo.io/docs/migrate#universal-ids")

func (m *Migrate) setupAtlas() error {
	// Using one of the Atlas options, opt-in to Atlas migration.
	if !m.atlas.enabled && (m.atlas.skip != NoChange || len(m.atlas.diff) > 0 || len(m.atlas.apply) > 0) || m.atlas.dir != nil {
		m.atlas.enabled = true
	}
	if !m.atlas.enabled {
		return nil
	}
	if m.withFixture {
		return errors.New("sql/schema: WithFixture(true) does not work in Atlas migration")
	}
	skip := DropIndex | DropColumn
	if m.atlas.skip != NoChange {
		skip = m.atlas.skip
	}
	if m.dropIndexes {
		skip &= ^DropIndex
	}
	if m.dropColumns {
		skip &= ^DropColumn
	}
	if skip != NoChange {
		m.atlas.diff = append(m.atlas.diff, filterChanges(skip))
	}
	if !m.withForeignKeys {
		m.atlas.diff = append(m.atlas.diff, withoutForeignKeys)
	}
	if m.atlas.dir != nil && m.atlas.fmt == nil {
		m.atlas.fmt = sqltool.GolangMigrateFormatter
	}
	if m.universalID && m.atlas.dir != nil {
		// If global unique ids and a migration directory is given, enable the file based type store for pk ranges.
		m.typeStore = &dirTypeStore{dir: m.atlas.dir}
		// To guard the user against a possible bug due to backward incompatibility, the file based type store must
		// be enabled by an option. For more information see the comment of WithUniversalID function.
		if !m.atlas.typeStoreConsent {
			return errConsent
		}
		m.atlas.diff = append(m.atlas.diff, m.ensureTypeTable)
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
		plan, err := m.atDiff(ctx, tx, "", tables...)
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

func (m *Migrate) atDiff(ctx context.Context, conn dialect.ExecQuerier, name string, tables ...*Table) (*migrate.Plan, error) {
	drv, err := m.atOpen(conn)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	tt, err := m.aTables(ctx, m, conn, tables)
	if err != nil {
		return nil, err
	}
	// Diff changes.
	var differ Differ = DiffFunc(drv.SchemaDiff)
	for i := len(m.atlas.diff) - 1; i >= 0; i-- {
		differ = m.atlas.diff[i](differ)
	}
	changes, err := differ.Diff(current, &schema.Schema{Name: current.Name, Attrs: current.Attrs, Tables: tt})
	if err != nil {
		return nil, err
	}
	// Plan changes.
	return drv.PlanChanges(ctx, name, changes)
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
		if c1.Unique && (len(t1.PrimaryKey) != 1 || t1.PrimaryKey[0] != c1) {
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
		desc := descIndexes(idx1)
		for _, p := range idx2.Parts {
			p.Desc = desc[p.C.Name]
		}
		t2.AddIndexes(idx2)
	}
	return nil
}

func (m *Migrate) ensureTypeTable(next Differ) Differ {
	return DiffFunc(func(current, desired *schema.Schema) ([]schema.Change, error) {
		// If there is a types table but no types file yet, the user most likely
		// switched from online migration to migration files.
		if len(m.dbTypeRanges) == 0 {
			var (
				at = schema.NewTable(TypeTable)
				et = NewTable(TypeTable).
					AddPrimary(&Column{Name: "id", Type: field.TypeUint, Increment: true}).
					AddColumn(&Column{Name: "type", Type: field.TypeString, Unique: true})
			)
			m.atTable(et, at)
			if err := m.aColumns(m, et, at); err != nil {
				return nil, err
			}
			if err := m.aIndexes(m, et, at); err != nil {
				return nil, err
			}
			desired.Tables = append(desired.Tables, at)
		}
		// If there is a drift between the types stored in the database and the ones stored in the file,
		// stop diffing, as this is potentially destructive. This will most likely happen on the first diffing
		// after moving from online-migration to versioned migrations if the "old" ent types are not in sync with
		// the deterministic ones computed by the new engine.
		if len(m.dbTypeRanges) > 0 && len(m.fileTypeRanges) > 0 && !equal(m.fileTypeRanges, m.dbTypeRanges) {
			return nil, fmt.Errorf(
				"type allocation range drift detected: %v <> %v: see %s for more information",
				m.dbTypeRanges, m.fileTypeRanges,
				"https://entgo.io/docs/versioned-migrations#moving-from-auto-migration-to-versioned-migrations",
			)
		}
		changes, err := next.Diff(current, desired)
		if err != nil {
			return nil, err
		}
		if len(m.dbTypeRanges) > 0 && len(m.fileTypeRanges) == 0 {
			// Override the types file created in the diff process with the "old" allocated types ranges.
			if err := m.typeStore.(*dirTypeStore).save(m.dbTypeRanges); err != nil {
				return nil, err
			}
			// Change the type range allocations since they will be added to the migration files when
			// writing the migration plan to migration files.
			m.typeRanges = m.dbTypeRanges
		}
		return changes, nil
	})
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

// descIndexes returns a map holding the DESC mapping if exist.
func descIndexes(idx *Index) map[string]bool {
	descs := make(map[string]bool)
	if idx.Annotation == nil {
		return descs
	}
	// If DESC (without a column) was defined on the
	// annotation, map it to the single column index.
	if idx.Annotation.Desc && len(idx.Columns) == 1 {
		descs[idx.Columns[0].Name] = idx.Annotation.Desc
	}
	for column, desc := range idx.Annotation.DescColumns {
		descs[column] = desc
	}
	return descs
}

const entTypes = ".ent_types"

// dirTypeStore stores and read pk information from a text file stored alongside generated versioned migrations.
// This behaviour is enabled automatically when using versioned migrations.
type dirTypeStore struct {
	dir migrate.Dir
}

const atlasDirective = "atlas:sum ignore\n"

// load the types from the types file.
func (s *dirTypeStore) load(context.Context, dialect.ExecQuerier) ([]string, error) {
	f, err := s.dir.Open(entTypes)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("reading types file: %w", err)
	}
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	defer f.Close()
	c, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("reading types file: %w", err)
	}
	return strings.Split(strings.TrimPrefix(string(c), atlasDirective), ","), nil
}

// add a new type entry to the types file.
func (s *dirTypeStore) add(ctx context.Context, conn dialect.ExecQuerier, t string) error {
	ts, err := s.load(ctx, conn)
	if err != nil {
		return fmt.Errorf("adding type %q: %w", t, err)
	}
	return s.save(append(ts, t))
}

// save takes the given allocation range and writes them to the types file.
// The types file will be overridden.
func (s *dirTypeStore) save(ts []string) error {
	if err := s.dir.WriteFile(entTypes, []byte(atlasDirective+strings.Join(ts, ","))); err != nil {
		return fmt.Errorf("writing types file: %w", err)
	}
	return nil
}

var _ typeStore = (*dirTypeStore)(nil)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
