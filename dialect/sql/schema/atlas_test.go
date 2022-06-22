// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"ariga.io/atlas/sql/migrate"
	"github.com/stretchr/testify/require"
)

func TestDirTypeStore(t *testing.T) {
	ex := []string{"a", "b", "c"}
	p := t.TempDir()
	d, err := migrate.NewLocalDir(p)
	require.NoError(t, err)

	s := &dirTypeStore{d}
	require.NoError(t, s.save(ex))
	require.FileExists(t, filepath.Join(p, entTypes))
	c, err := os.ReadFile(filepath.Join(p, entTypes))
	require.NoError(t, err)
	require.Contains(t, string(c), atlasDirective)

	ac, err := s.load(context.Background(), nil)
	require.NoError(t, err)
	require.Equal(t, ex, ac)
}
