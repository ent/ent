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

	// AllFeatures holds a list of all feature-flags.
	AllFeatures = []Feature{
		FeaturePrivacy,
		FeatureEntQL,
	}
)

// FeatureStage describes the stage of the codegen feature.
type FeatureStage int

const (
	_ FeatureStage = iota

	// An Experimental feature is one that is in development,
	// and it's actively tested in the integration environment.
	Experimental

	// An Alpha feature is one that its initial development was
	// finished, it's tested on the infra of the ent team, but
	// we expect breaking-changes to its API.
	Alpha

	// A Beta feature is an Alpha feature that was added to the entgo.io
	// documentation, and no breaking-changes are expected for it.
	Beta

	// A Stable feature is a Beta feature that was running a while on ent infra.
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

	// cleanup used to cleanup all changes when a feature-flag is removed.
	// e.g. delete files from previous codegen runs.
	cleanup func(*Config) error
}
