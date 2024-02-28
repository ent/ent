// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqlclient"
	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

// Atlas atlas migration engine.
type Atlas struct {
	atDriver   migrate.Driver
	sqlDialect sqlDialect

	legacy      bool // if the legacy migration engine instead of Atlas should be used
	withFixture bool // deprecated: with fks rename fixture
	sum         bool // deprecated: sum file generation will be required

	indent          string // plan indentation
	errNoPlan       bool   // no plan error enabled
	universalID     bool   // global unique ids
	dropColumns     bool   // drop deleted columns
	dropIndexes     bool   // drop deleted indexes
	withForeignKeys bool   // with foreign keys
	mode            Mode
	hooks           []Hook              // hooks to apply before creation
	diffHooks       []DiffHook          // diff hooks to run when diffing current and desired
	diffOptions     []schema.DiffOption // diff options to pass to the diff engine
	applyHook       []ApplyHook         // apply hooks to run when applying the plan
	skip            ChangeKind          // what changes to skip and not apply
	dir             migrate.Dir         // the migration directory to read from
	fmt             migrate.Formatter   // how to format the plan into migration files

	driver  dialect.Driver // driver passed in when not using an atlas URL
	url     *url.URL       // url of database connection
	dialect string         // Ent dialect to use when generating migration files

	types []string // pre-existing pk range allocation for global unique id
}

// Diff compares the state read from a database connection or migration directory with the state defined by the Ent
// schema. Changes will be written to new migration files.
func Diff(ctx context.Context, u, name string, tables []*Table, opts ...MigrateOption) (err error) {
	m, err := NewMigrateURL(u, opts...)
	if err != nil {
		return err
	}
	return m.NamedDiff(ctx, name, tables...)
}

// NewMigrate creates a new Atlas form the given dialect.Driver.
func NewMigrate(drv dialect.Driver, opts ...MigrateOption) (*Atlas, error) {
	a := &Atlas{driver: drv, withForeignKeys: true, mode: ModeInspect, sum: true}
	for _, opt := range opts {
		opt(a)
	}
	a.dialect = a.driver.Dialect()
	if err := a.init(); err != nil {
		return nil, err
	}
	return a, nil
}

// NewMigrateURL create a new Atlas from the given url.
func NewMigrateURL(u string, opts ...MigrateOption) (*Atlas, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	a := &Atlas{url: parsed, withForeignKeys: true, mode: ModeInspect, sum: true}
	for _, opt := range opts {
		opt(a)
	}
	if a.dialect == "" {
		a.dialect = parsed.Scheme
	}
	if err := a.init(); err != nil {
		return nil, err
	}
	return a, nil
}

// Create creates all schema resources in the database. It works in an "append-only"
// mode, which means, it only creates tables, appends columns to tables or modifies column types.
//
// Column can be modified by turning into a NULL from NOT NULL, or having a type conversion not
// resulting data altering. From example, changing varchar(255) to varchar(120) is invalid, but
// changing varchar(120) to varchar(255) is valid. For more info, see the convert function below.
func (a *Atlas) Create(ctx context.Context, tables ...*Table) (err error) {
	a.setupTables(tables)
	var creator Creator = CreateFunc(a.create)
	if a.legacy {
		m, err := a.legacyMigrate()
		if err != nil {
			return err
		}
		creator = CreateFunc(m.create)
	}
	for i := len(a.hooks) - 1; i >= 0; i-- {
		creator = a.hooks[i](creator)
	}
	return creator.Create(ctx, tables...)
}

// Diff compares the state read from the connected database with the state defined by Ent.
// Changes will be written to migration files by the configured Planner.
func (a *Atlas) Diff(ctx context.Context, tables ...*Table) error {
	return a.NamedDiff(ctx, "changes", tables...)
}

// NamedDiff compares the state read from the connected database with the state defined by Ent.
// Changes will be written to migration files by the configured Planner.
func (a *Atlas) NamedDiff(ctx context.Context, name string, tables ...*Table) error {
	if a.dir == nil {
		return errors.New("no migration directory given")
	}
	opts := []migrate.PlannerOption{migrate.WithFormatter(a.fmt)}
	if a.sum {
		// Validate the migration directory before proceeding.
		if err := migrate.Validate(a.dir); err != nil {
			return fmt.Errorf("validating migration directory: %w", err)
		}
	} else {
		opts = append(opts, migrate.DisableChecksum())
	}
	a.setupTables(tables)
	// Set up connections.
	if a.driver != nil {
		var err error
		a.sqlDialect, err = a.entDialect(ctx, a.driver)
		if err != nil {
			return err
		}
		a.atDriver, err = a.sqlDialect.atOpen(a.sqlDialect)
		if err != nil {
			return err
		}
	} else {
		c, err := sqlclient.OpenURL(ctx, a.url)
		if err != nil {
			return err
		}
		defer c.Close()
		a.sqlDialect, err = a.entDialect(ctx, entsql.OpenDB(a.dialect, c.DB))
		if err != nil {
			return err
		}
		a.atDriver = c.Driver
	}
	defer func() {
		a.sqlDialect = nil
		a.atDriver = nil
	}()
	if err := a.sqlDialect.init(ctx); err != nil {
		return err
	}
	if a.universalID {
		tables = append(tables, NewTypesTable())
	}
	var (
		err  error
		plan *migrate.Plan
	)
	switch a.mode {
	case ModeInspect:
		plan, err = a.planInspect(ctx, a.sqlDialect, name, tables)
	case ModeReplay:
		plan, err = a.planReplay(ctx, name, tables)
	default:
		return fmt.Errorf("unknown migration mode: %q", a.mode)
	}
	switch {
	case err != nil:
		return err
	case len(plan.Changes) == 0:
		if a.errNoPlan {
			return migrate.ErrNoPlan
		}
		return nil
	default:
		return migrate.NewPlanner(nil, a.dir, opts...).WritePlan(plan)
	}
}

func (a *Atlas) cleanSchema(ctx context.Context, name string, err0 error) (err error) {
	defer func() {
		if err0 != nil {
			err = fmt.Errorf("%v: %w", err0, err)
		}
	}()
	s, err := a.atDriver.InspectSchema(ctx, name, nil)
	if err != nil {
		return err
	}
	drop := make([]schema.Change, len(s.Tables))
	for i, t := range s.Tables {
		drop[i] = &schema.DropTable{T: t, Extra: []schema.Clause{&schema.IfExists{}}}
	}
	return a.atDriver.ApplyChanges(ctx, drop)
}

// VerifyTableRange ensures, that the defined autoincrement starting value is set for each table as defined by the
// TypTable. This is necessary for MySQL versions < 8.0. In those versions the defined starting value for AUTOINCREMENT
// columns was stored in memory, and when a server restarts happens and there are no rows yet in a table, the defined
// starting value is lost, which will result in incorrect behavior when working with global unique ids. Calling this
// method on service start ensures the information are correct and are set again, if they aren't. For MySQL versions > 8
// calling this method is only required once after the upgrade.
func (a *Atlas) VerifyTableRange(ctx context.Context, tables []*Table) error {
	if a.driver != nil {
		var err error
		a.sqlDialect, err = a.entDialect(ctx, a.driver)
		if err != nil {
			return err
		}
	} else {
		c, err := sqlclient.OpenURL(ctx, a.url)
		if err != nil {
			return err
		}
		defer c.Close()
		a.sqlDialect, err = a.entDialect(ctx, entsql.OpenDB(a.dialect, c.DB))
		if err != nil {
			return err
		}
	}
	defer func() {
		a.sqlDialect = nil
	}()
	vr, ok := a.sqlDialect.(verifyRanger)
	if !ok {
		return nil
	}
	types, err := a.loadTypes(ctx, a.sqlDialect)
	if err != nil {
		// In most cases this means the table does not exist, which in turn
		// indicates the user does not use global unique ids.
		return err
	}
	for _, t := range tables {
		id := indexOf(types, t.Name)
		if id == -1 {
			continue
		}
		if err := vr.verifyRange(ctx, a.sqlDialect, t, int64(id<<32)); err != nil {
			return err
		}
	}
	return nil
}

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
func WithDiffHook(hooks ...DiffHook) MigrateOption {
	return func(a *Atlas) {
		a.diffHooks = append(a.diffHooks, hooks...)
	}
}

// WithDiffOptions adds a list of options to pass to the diff engine.
func WithDiffOptions(opts ...schema.DiffOption) MigrateOption {
	return func(a *Atlas) {
		a.diffOptions = append(a.diffOptions, opts...)
	}
}

// WithSkipChanges allows skipping/filtering list of changes
// returned by the Differ before executing migration planning.
//
//	SkipChanges(schema.DropTable|schema.DropColumn)
func WithSkipChanges(skip ChangeKind) MigrateOption {
	return func(a *Atlas) {
		a.skip = skip
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
func WithApplyHook(hooks ...ApplyHook) MigrateOption {
	return func(a *Atlas) {
		a.applyHook = append(a.applyHook, hooks...)
	}
}

// WithAtlas is an opt-out option for v0.11 indicating the migration
// should be executed using the deprecated legacy engine.
// Note, in future versions, this option is going to be removed
// and the Atlas (https://atlasgo.io) based migration engine should be used.
//
// Deprecated: The legacy engine will be removed.
func WithAtlas(b bool) MigrateOption {
	return func(a *Atlas) {
		a.legacy = !b
	}
}

// WithDir sets the atlas migration directory to use to store migration files.
func WithDir(dir migrate.Dir) MigrateOption {
	return func(a *Atlas) {
		a.dir = dir
	}
}

// WithFormatter sets atlas formatter to use to write changes to migration files.
func WithFormatter(fmt migrate.Formatter) MigrateOption {
	return func(a *Atlas) {
		a.fmt = fmt
	}
}

// WithDialect configures the Ent dialect to use when migrating for an Atlas supported dialect flavor.
// As an example, Ent can work with TiDB in MySQL dialect and Atlas can handle TiDB migrations.
func WithDialect(d string) MigrateOption {
	return func(a *Atlas) {
		a.dialect = d
	}
}

// WithSumFile instructs atlas to generate a migration directory integrity sum file.
//
// Deprecated: generating the sum file is now opt-out. This method will be removed in future versions.
func WithSumFile() MigrateOption {
	return func(a *Atlas) {}
}

// DisableChecksum instructs atlas to skip migration directory integrity sum file generation.
//
// Deprecated: generating the sum file will no longer be optional in future versions.
func DisableChecksum() MigrateOption {
	return func(a *Atlas) {
		a.sum = false
	}
}

// WithMigrationMode instructs atlas how to compute the current state of the schema. This can be done by either
// replaying (ModeReplay) the migration directory on the connected database, or by inspecting (ModeInspect) the
// connection. Currently, ModeReplay is opt-in, and ModeInspect is the default. In future versions, ModeReplay will
// become the default behavior. This option has no effect when using online migrations.
func WithMigrationMode(mode Mode) MigrateOption {
	return func(a *Atlas) {
		a.mode = mode
	}
}

// Mode to compute the current state.
type Mode uint

const (
	// ModeReplay computes the current state by replaying the migration directory on the connected database.
	ModeReplay = iota
	// ModeInspect computes the current state by inspecting the connected database.
	ModeInspect
)

// StateReader returns an atlas migrate.StateReader returning the state as described by the Ent table slice.
func (a *Atlas) StateReader(tables ...*Table) migrate.StateReaderFunc {
	return func(ctx context.Context) (*schema.Realm, error) {
		if a.sqlDialect == nil {
			drv, err := a.entDialect(ctx, a.driver)
			if err != nil {
				return nil, err
			}
			a.sqlDialect = drv
		}
		ts, err := a.tables(tables)
		if err != nil {
			return nil, err
		}
		return &schema.Realm{Schemas: []*schema.Schema{{Tables: ts}}}, nil
	}
}

// atBuilder must be implemented by the different drivers in
// order to convert a dialect/sql/schema to atlas/sql/schema.
type atBuilder interface {
	atOpen(dialect.ExecQuerier) (migrate.Driver, error)
	atTable(*Table, *schema.Table)
	supportsDefault(*Column) bool
	atTypeC(*Column, *schema.Column) error
	atUniqueC(*Table, *Column, *schema.Table, *schema.Column)
	atIncrementC(*schema.Table, *schema.Column)
	atIncrementT(*schema.Table, int64)
	atIndex(*Index, *schema.Table, *schema.Index) error
	atTypeRangeSQL(t ...string) string
}

// init initializes the configuration object based on the options passed in.
func (a *Atlas) init() error {
	skip := DropIndex | DropColumn
	if a.skip != NoChange {
		skip = a.skip
	}
	if a.dropIndexes {
		skip &= ^DropIndex
	}
	if a.dropColumns {
		skip &= ^DropColumn
	}
	if skip != NoChange {
		a.diffHooks = append(a.diffHooks, filterChanges(skip))
	}
	if !a.withForeignKeys {
		a.diffHooks = append(a.diffHooks, withoutForeignKeys)
	}
	if a.dir != nil && a.fmt == nil {
		switch a.dir.(type) {
		case *sqltool.GooseDir:
			a.fmt = sqltool.GooseFormatter
		case *sqltool.DBMateDir:
			a.fmt = sqltool.DBMateFormatter
		case *sqltool.FlywayDir:
			a.fmt = sqltool.FlywayFormatter
		case *sqltool.LiquibaseDir:
			a.fmt = sqltool.LiquibaseFormatter
		default: // migrate.LocalDir, sqltool.GolangMigrateDir and custom ones
			a.fmt = sqltool.GolangMigrateFormatter
		}
	}
	if a.mode == ModeReplay {
		// ModeReplay requires a migration directory.
		if a.dir == nil {
			return errors.New("sql/schema: WithMigrationMode(ModeReplay) requires versioned migrations: WithDir()")
		}
		// ModeReplay requires sum file generation.
		if !a.sum {
			return errors.New("sql/schema: WithMigrationMode(ModeReplay) requires migration directory integrity file")
		}
	}
	return nil
}

// create is the Atlas engine based online migration.
func (a *Atlas) create(ctx context.Context, tables ...*Table) (err error) {
	if a.universalID {
		tables = append(tables, NewTypesTable())
	}
	if a.driver != nil {
		a.sqlDialect, err = a.entDialect(ctx, a.driver)
		if err != nil {
			return err
		}
	} else {
		c, err := sqlclient.OpenURL(ctx, a.url)
		if err != nil {
			return err
		}
		defer c.Close()
		a.sqlDialect, err = a.entDialect(ctx, entsql.OpenDB(a.dialect, c.DB))
		if err != nil {
			return err
		}
	}
	defer func() { a.sqlDialect = nil }()
	if err := a.sqlDialect.init(ctx); err != nil {
		return err
	}
	// Open a transaction for backwards compatibility,
	// even if the migration is not transactional.
	tx, err := a.sqlDialect.Tx(ctx)
	if err != nil {
		return err
	}
	a.atDriver, err = a.sqlDialect.atOpen(tx)
	if err != nil {
		return err
	}
	defer func() { a.atDriver = nil }()
	if err := func() error {
		plan, err := a.planInspect(ctx, tx, "changes", tables)
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
		for i := len(a.applyHook) - 1; i >= 0; i-- {
			applier = a.applyHook[i](applier)
		}
		return applier.Apply(ctx, tx, plan)
	}(); err != nil {
		err = fmt.Errorf("sql/schema: %w", err)
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}
	return tx.Commit()
}

// planInspect creates the current state by inspecting the connected database, computing the current state of the Ent schema
// and proceeds to diff the changes to create a migration plan.
func (a *Atlas) planInspect(ctx context.Context, conn dialect.ExecQuerier, name string, tables []*Table) (*migrate.Plan, error) {
	current, err := a.atDriver.InspectSchema(ctx, "", &schema.InspectOptions{
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
	var types []string
	if a.universalID {
		types, err = a.loadTypes(ctx, conn)
		if err != nil && !errors.Is(err, errTypeTableNotFound) {
			return nil, err
		}
		a.types = types
	}
	realm, err := a.StateReader(tables...).ReadState(ctx)
	if err != nil {
		return nil, err
	}
	desired := realm.Schemas[0]
	desired.Name, desired.Attrs = current.Name, current.Attrs
	return a.diff(ctx, name, current, desired, a.types[len(types):])
}

func (a *Atlas) planReplay(ctx context.Context, name string, tables []*Table) (*migrate.Plan, error) {
	// We consider a database clean if there are no tables in the connected schema.
	s, err := a.atDriver.InspectSchema(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	if len(s.Tables) > 0 {
		return nil, &migrate.NotCleanError{Reason: fmt.Sprintf("found table %q", s.Tables[0].Name)}
	}
	// Replay the migration directory on the database.
	ex, err := migrate.NewExecutor(a.atDriver, a.dir, &migrate.NopRevisionReadWriter{})
	if err != nil {
		return nil, err
	}
	if err := ex.ExecuteN(ctx, 0); err != nil && !errors.Is(err, migrate.ErrNoPendingFiles) {
		return nil, a.cleanSchema(ctx, "", err)
	}
	// Inspect the current schema (migration directory).
	current, err := a.atDriver.InspectSchema(ctx, "", nil)
	if err != nil {
		return nil, a.cleanSchema(ctx, "", err)
	}
	var types []string
	if a.universalID {
		if types, err = a.loadTypes(ctx, a.sqlDialect); err != nil && !errors.Is(err, errTypeTableNotFound) {
			return nil, a.cleanSchema(ctx, "", err)
		}
		a.types = types
	}
	if err := a.cleanSchema(ctx, "", nil); err != nil {
		return nil, fmt.Errorf("clean schemas after migration replaying: %w", err)
	}
	desired, err := a.tables(tables)
	if err != nil {
		return nil, err
	}
	// In case of replay mode, normalize the desired state (i.e. ent/schema).
	if nr, ok := a.atDriver.(schema.Normalizer); ok {
		ns, err := nr.NormalizeSchema(ctx, schema.New(current.Name).AddTables(desired...))
		if err != nil {
			return nil, err
		}
		if len(ns.Tables) != len(desired) {
			return nil, fmt.Errorf("unexpected number of tables after normalization: %d != %d", len(ns.Tables), len(desired))
		}
		// Ensure all tables exist in the normalized format and the order is preserved.
		for i, t := range desired {
			d, ok := ns.Table(t.Name)
			if !ok {
				return nil, fmt.Errorf("table %q not found after normalization", t.Name)
			}
			desired[i] = d
		}
	}
	return a.diff(ctx, name, current,
		&schema.Schema{Name: current.Name, Attrs: current.Attrs, Tables: desired}, a.types[len(types):],
		// For BC reason, we omit the schema qualifier from the migration scripts,
		// but that is currently limiting versioned migration to a single schema.
		func(opts *migrate.PlanOptions) {
			var noQualifier string
			opts.SchemaQualifier = &noQualifier
		},
	)
}

func (a *Atlas) diff(ctx context.Context, name string, current, desired *schema.Schema, newTypes []string, opts ...migrate.PlanOption) (*migrate.Plan, error) {
	changes, err := (&diffDriver{a.atDriver, a.diffHooks}).SchemaDiff(current, desired, a.diffOptions...)
	if err != nil {
		return nil, err
	}
	filtered := make([]schema.Change, 0, len(changes))
	for _, c := range changes {
		switch c.(type) {
		// Select only table creation and modification. The reason we may encounter this, even though specific tables
		// are passed to Inspect, is if the MySQL system variable 'lower_case_table_names' is set to 1. In such a case,
		// the given tables will be returned from inspection because MySQL compares case-insensitive, but they won't
		// match when compare them in code.
		case *schema.AddTable, *schema.ModifyTable:
			filtered = append(filtered, c)
		}
	}
	if a.indent != "" {
		opts = append(opts, func(opts *migrate.PlanOptions) {
			opts.Indent = a.indent
		})
	}
	plan, err := a.atDriver.PlanChanges(ctx, name, filtered, opts...)
	if err != nil {
		return nil, err
	}
	if len(newTypes) > 0 {
		plan.Changes = append(plan.Changes, &migrate.Change{
			Cmd:     a.sqlDialect.atTypeRangeSQL(newTypes...),
			Comment: fmt.Sprintf("add pk ranges for %s tables", strings.Join(newTypes, ",")),
		})
	}
	return plan, nil
}

var errTypeTableNotFound = errors.New("ent_type table not found")

// loadTypes loads the currently saved range allocations from the TypeTable.
func (a *Atlas) loadTypes(ctx context.Context, conn dialect.ExecQuerier) ([]string, error) {
	// Fetch pre-existing type allocations.
	exists, err := a.sqlDialect.tableExist(ctx, conn, TypeTable)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errTypeTableNotFound
	}
	rows := &entsql.Rows{}
	query, args := entsql.Dialect(a.dialect).
		Select("type").From(entsql.Table(TypeTable)).OrderBy(entsql.Asc("id")).Query()
	if err := conn.Query(ctx, query, args, rows); err != nil {
		return nil, fmt.Errorf("query types table: %w", err)
	}
	defer rows.Close()
	var types []string
	if err := entsql.ScanSlice(rows, &types); err != nil {
		return nil, err
	}
	return types, nil
}

type db struct{ dialect.ExecQuerier }

func (d *db) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	rows := &entsql.Rows{}
	if err := d.ExecQuerier.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	return rows.ColumnScanner.(*sql.Rows), nil
}

func (d *db) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	var r sql.Result
	if err := d.ExecQuerier.Exec(ctx, query, args, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// tables converts an Ent table slice to an atlas table slice
func (a *Atlas) tables(tables []*Table) ([]*schema.Table, error) {
	ts := make([]*schema.Table, len(tables))
	for i, et := range tables {
		at := schema.NewTable(et.Name)
		if et.Comment != "" {
			at.SetComment(et.Comment)
		}
		a.sqlDialect.atTable(et, at)
		if a.universalID && et.Name != TypeTable && len(et.PrimaryKey) == 1 {
			r, err := a.pkRange(et)
			if err != nil {
				return nil, err
			}
			a.sqlDialect.atIncrementT(at, r)
		}
		if err := a.aColumns(et, at); err != nil {
			return nil, err
		}
		if err := a.aIndexes(et, at); err != nil {
			return nil, err
		}
		ts[i] = at
	}
	for i, t1 := range tables {
		t2 := ts[i]
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
			for _, t2 := range ts {
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
	return ts, nil
}

func (a *Atlas) aColumns(et *Table, at *schema.Table) error {
	for _, c1 := range et.Columns {
		c2 := schema.NewColumn(c1.Name).
			SetNull(c1.Nullable)
		if c1.Collation != "" {
			c2.SetCollation(c1.Collation)
		}
		if c1.Comment != "" {
			c2.SetComment(c1.Comment)
		}
		if err := a.sqlDialect.atTypeC(c1, c2); err != nil {
			return err
		}
		if err := a.atDefault(c1, c2); err != nil {
			return err
		}
		if c1.Unique && (len(et.PrimaryKey) != 1 || et.PrimaryKey[0] != c1) {
			a.sqlDialect.atUniqueC(et, c1, at, c2)
		}
		if c1.Increment {
			a.sqlDialect.atIncrementC(at, c2)
		}
		at.AddColumns(c2)
	}
	return nil
}

func (a *Atlas) atDefault(c1 *Column, c2 *schema.Column) error {
	if c1.Default == nil || !a.sqlDialect.supportsDefault(c1) {
		return nil
	}
	switch x := c1.Default.(type) {
	case Expr:
		if len(x) > 1 && (x[0] != '(' || x[len(x)-1] != ')') {
			x = "(" + x + ")"
		}
		c2.SetDefault(&schema.RawExpr{X: string(x)})
	case map[string]Expr:
		d, ok := x[a.sqlDialect.Dialect()]
		if !ok {
			return nil
		}
		if len(d) > 1 && (d[0] != '(' || d[len(d)-1] != ')') {
			d = "(" + d + ")"
		}
		c2.SetDefault(&schema.RawExpr{X: string(d)})
	default:
		switch {
		case c1.Type == field.TypeJSON:
			s, ok := c1.Default.(string)
			if !ok {
				return fmt.Errorf("invalid default value for JSON column %q: %v", c1.Name, c1.Default)
			}
			c2.SetDefault(&schema.Literal{V: strings.ReplaceAll(s, "'", "''")})
		default:
			// Keep backwards compatibility with the old default value format.
			x := fmt.Sprint(c1.Default)
			if v, ok := c1.Default.(string); ok && c1.Type != field.TypeUUID && c1.Type != field.TypeTime {
				// Escape single quote by replacing each with 2.
				x = fmt.Sprintf("'%s'", strings.ReplaceAll(v, "'", "''"))
			}
			c2.SetDefault(&schema.RawExpr{X: x})
		}
	}
	return nil
}

func (a *Atlas) aIndexes(et *Table, at *schema.Table) error {
	// Primary-key index.
	pk := make([]*schema.Column, 0, len(et.PrimaryKey))
	for _, c1 := range et.PrimaryKey {
		c2, ok := at.Column(c1.Name)
		if !ok {
			return fmt.Errorf("unexpected primary-key column: %q", c1.Name)
		}
		pk = append(pk, c2)
	}
	// CreateFunc might clear the primary keys.
	if len(pk) > 0 {
		at.SetPrimaryKey(schema.NewPrimaryKey(pk...))
	}
	// Rest of indexes.
	for _, idx1 := range et.Indexes {
		idx2 := schema.NewIndex(idx1.Name).
			SetUnique(idx1.Unique)
		if err := a.sqlDialect.atIndex(idx1, at, idx2); err != nil {
			return err
		}
		desc := descIndexes(idx1)
		for _, p := range idx2.Parts {
			p.Desc = desc[p.C.Name]
		}
		at.AddIndexes(idx2)
	}
	return nil
}

// setupTables ensures the table is configured properly, like table columns
// are linked to their indexes, and PKs columns are defined.
func (a *Atlas) setupTables(tables []*Table) {
	for _, t := range tables {
		if t.columns == nil {
			t.columns = make(map[string]*Column, len(t.Columns))
		}
		for _, c := range t.Columns {
			t.columns[c.Name] = c
		}
		for _, idx := range t.Indexes {
			idx.Name = a.symbol(idx.Name)
			for _, c := range idx.Columns {
				c.indexes.append(idx)
			}
		}
		for _, pk := range t.PrimaryKey {
			c := t.columns[pk.Name]
			c.Key = PrimaryKey
			pk.Key = PrimaryKey
		}
		for _, fk := range t.ForeignKeys {
			fk.Symbol = a.symbol(fk.Symbol)
			for i := range fk.Columns {
				fk.Columns[i].foreign = fk
			}
		}
	}
}

// symbol makes sure the symbol length is not longer than the maxlength in the dialect.
func (a *Atlas) symbol(name string) string {
	size := 64
	if a.dialect == dialect.Postgres {
		size = 63
	}
	if len(name) <= size {
		return name
	}
	return fmt.Sprintf("%s_%x", name[:size-33], md5.Sum([]byte(name)))
}

// entDialect returns the Ent dialect as configured by the dialect option.
func (a *Atlas) entDialect(ctx context.Context, drv dialect.Driver) (sqlDialect, error) {
	var d sqlDialect
	switch a.dialect {
	case dialect.MySQL:
		d = &MySQL{Driver: drv}
	case dialect.SQLite:
		d = &SQLite{Driver: drv, WithForeignKeys: a.withForeignKeys}
	case dialect.Postgres:
		d = &Postgres{Driver: drv}
	default:
		return nil, fmt.Errorf("sql/schema: unsupported dialect %q", a.dialect)
	}
	if err := d.init(ctx); err != nil {
		return nil, err
	}
	return d, nil
}

func (a *Atlas) pkRange(et *Table) (int64, error) {
	idx := indexOf(a.types, et.Name)
	// If the table re-created, re-use its range from
	// the past. Otherwise, allocate a new id-range.
	if idx == -1 {
		if len(a.types) > MaxTypes {
			return 0, fmt.Errorf("max number of types exceeded: %d", MaxTypes)
		}
		idx = len(a.types)
		a.types = append(a.types, et.Name)
	}
	return int64(idx << 32), nil
}

func setAtChecks(et *Table, at *schema.Table) {
	if check := et.Annotation.Check; check != "" {
		at.AddChecks(&schema.Check{
			Expr: check,
		})
	}
	if checks := et.Annotation.Checks; len(et.Annotation.Checks) > 0 {
		names := make([]string, 0, len(checks))
		for name := range checks {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			at.AddChecks(&schema.Check{
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

// driver decorates the atlas migrate.Driver and adds "diff hooking" and functionality.
type diffDriver struct {
	migrate.Driver
	hooks []DiffHook // hooks to apply
}

// RealmDiff creates the diff between two realms. Since Ent does not care about Realms,
// not even schema changes, calling this method raises an error.
func (r *diffDriver) RealmDiff(_, _ *schema.Realm, _ ...schema.DiffOption) ([]schema.Change, error) {
	return nil, errors.New("sqlDialect does not support working with realms")
}

// SchemaDiff creates the diff between two schemas, but includes "diff hooks".
func (r *diffDriver) SchemaDiff(from, to *schema.Schema, opts ...schema.DiffOption) ([]schema.Change, error) {
	var d Differ = DiffFunc(func(current, desired *schema.Schema) ([]schema.Change, error) {
		return r.Driver.SchemaDiff(current, desired, opts...)
	})
	for i := len(r.hooks) - 1; i >= 0; i-- {
		d = r.hooks[i](d)
	}
	return d.Diff(from, to)
}

// legacyMigrate returns a configured legacy migration engine (before Atlas) to keep backwards compatibility.
//
// Deprecated: Will be removed alongside legacy migration support.
func (a *Atlas) legacyMigrate() (*Migrate, error) {
	m := &Migrate{
		universalID:     a.universalID,
		dropColumns:     a.dropColumns,
		dropIndexes:     a.dropIndexes,
		withFixture:     a.withFixture,
		withForeignKeys: a.withForeignKeys,
		hooks:           a.hooks,
		atlas:           a,
	}
	switch a.dialect {
	case dialect.MySQL:
		m.sqlDialect = &MySQL{Driver: a.driver}
	case dialect.SQLite:
		m.sqlDialect = &SQLite{Driver: a.driver, WithForeignKeys: a.withForeignKeys}
	case dialect.Postgres:
		m.sqlDialect = &Postgres{Driver: a.driver}
	default:
		return nil, fmt.Errorf("sql/schema: unsupported dialect %q", a.dialect)
	}
	return m, nil
}

// removeAttr is a temporary patch due to compiler errors we get by using the generic
// schema.RemoveAttr function (<autogenerated>:1: internal compiler error: panic: ...).
// Can be removed in Go 1.20. See: https://github.com/golang/go/issues/54302.
func removeAttr(attrs []schema.Attr, t reflect.Type) []schema.Attr {
	f := make([]schema.Attr, 0, len(attrs))
	for _, a := range attrs {
		if reflect.TypeOf(a) != t {
			f = append(f, a)
		}
	}
	return f
}
