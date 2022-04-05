// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"os"
	"path/filepath"
)

var (
	// FeaturePrivacy provides a feature-flag for the privacy extension for ent.
	FeaturePrivacy = Feature{
		Name:        "privacy",
		Stage:       Alpha,
		Default:     false,
		Description: "Privacy provides a privacy layer for ent through the schema configuration",
		cleanup: func(c *Config) error {
			return os.RemoveAll(filepath.Join(c.Target, "privacy"))
		},
	}

	// FeatureEntQL provides a feature-flag for the entql extension for ent.
	FeatureEntQL = Feature{
		Name:        "entql",
		Stage:       Experimental,
		Default:     false,
		Description: "EntQL provides a generic filtering capability at runtime",
		cleanup: func(c *Config) error {
			return os.RemoveAll(filepath.Join(c.Target, "entql.go"))
		},
	}

	// FeatureSnapshot stores a snapshot of ent/schema and auto-solve merge-conflict (issue #852).
	FeatureSnapshot = Feature{
		Name:        "schema/snapshot",
		Stage:       Experimental,
		Default:     false,
		Description: "Schema snapshot stores a snapshot of ent/schema and auto-solve merge-conflict (issue #852)",
		GraphTemplates: []GraphTemplate{
			{
				Name:   "internal/schema",
				Format: "internal/schema.go",
			},
		},
		cleanup: func(c *Config) error {
			return remove(filepath.Join(c.Target, "internal"), "schema.go")
		},
	}

	// FeatureSchemaConfig allows users to pass init time alternate schema names
	// for each ent model. This is useful if your SQL tables are spread out against
	// multiple databases.
	FeatureSchemaConfig = Feature{
		Name:        "sql/schemaconfig",
		Stage:       Experimental,
		Default:     false,
		Description: "Allows alternate schema names for each ent model. Useful if SQL tables are spread out against multiple databases",
		GraphTemplates: []GraphTemplate{
			{
				Name:   "dialect/sql/internal/schemaconfig",
				Format: "internal/schemaconfig.go",
			},
		},
		cleanup: func(c *Config) error {
			return remove(filepath.Join(c.Target, "internal"), "schemaconfig.go")
		},
	}

	// FeatureLock provides a feature-flag for sql locking extension for ent.
	FeatureLock = Feature{
		Name:        "sql/lock",
		Stage:       Experimental,
		Default:     false,
		Description: "Allows users to use row-level locking in SQL using the 'FOR {UPDATE|SHARE}' clauses",
	}

	// FeatureModifier provides a feature-flag for adding query modifiers.
	FeatureModifier = Feature{
		Name:        "sql/modifier",
		Stage:       Experimental,
		Default:     false,
		Description: "Allows users to attach custom modifiers to queries",
	}

	// FeatureExecQuery provides a feature-flag for exposing the ExecContext/QueryContext methods of the underlying SQL drivers.
	FeatureExecQuery = Feature{
		Name:        "sql/execquery",
		Stage:       Experimental,
		Default:     false,
		Description: "Allows users to execute statements using the ExecContext/QueryContext methods of the underlying driver",
	}

	// FeatureUpsert provides a feature-flag for adding upsert (ON CONFLICT) capabilities to create builders.
	FeatureUpsert = Feature{
		Name:        "sql/upsert",
		Stage:       Experimental,
		Default:     false,
		Description: "Allows users to configure the `ON CONFLICT`/`ON DUPLICATE KEY` clause for `INSERT` statements",
	}

	FeatureVersionedMigration = Feature{
		Name:        "sql/versioned-migration",
		Stage:       Experimental,
		Default:     false,
		Description: "Allows users to work with versioned migrations / migration files",
	}

	// AllFeatures holds a list of all feature-flags.
	AllFeatures = []Feature{
		FeaturePrivacy,
		FeatureEntQL,
		FeatureSnapshot,
		FeatureSchemaConfig,
		FeatureLock,
		FeatureModifier,
		FeatureExecQuery,
		FeatureUpsert,
		FeatureVersionedMigration,
	}
)

// FeatureStage describes the stage of the codegen feature.
type FeatureStage int

const (
	_ FeatureStage = iota

	// Experimental features are in development, and actively being tested in the
	// integration environment.
	Experimental

	// Alpha features are features whose initial development was finished, tested
	// on the infra of the ent team, but we expect breaking-changes to their APIs.
	Alpha

	// Beta features are Alpha features that were added to the entgo.io
	// documentation, and no breaking-changes are expected for them.
	Beta

	// Stable features are Beta features that were running for a while on ent
	// infra.
	Stable
)

// A Feature of the ent codegen.
type Feature struct {
	// Name of the feature.
	Name string

	// Stage of the feature.
	Stage FeatureStage

	// Default values indicates if this feature is enabled by default.
	Default bool

	// A Description of this feature.
	Description string

	// Templates defines list of templates for extending or overriding the default
	// templates. In order to write the template output to a standalone file, use
	// the GraphTemplates below.
	Templates []*Template

	// GraphTemplates defines optional templates to be executed on the graph
	// and will their output will be written to the configured destination.
	GraphTemplates []GraphTemplate

	// cleanup used to cleanup all changes when a feature-flag is removed.
	// e.g. delete files from previous codegen runs.
	cleanup func(*Config) error
}

// remove file (if exists) and its dir if it's empty.
func remove(dir, file string) error {
	if err := os.Remove(filepath.Join(dir, file)); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	infos, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	if len(infos) == 0 {
		return os.Remove(dir)
	}
	return nil
}
