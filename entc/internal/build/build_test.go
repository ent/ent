package build

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuild(t *testing.T) {
	cfg := &Config{Path: "./testdata/valid"}
	plg, err := cfg.Build()
	require.NoError(t, err)
	schemas, err := plg.Load()
	require.NoError(t, err)
	require.Len(t, schemas, 3)
	require.Equal(t, "fbc/ent/entc/internal/build/testdata/valid", plg.PkgPath)
}

func TestBuildWrongPath(t *testing.T) {
	cfg := &Config{Path: "./boring"}
	plg, err := cfg.Build()
	require.Error(t, err)
	require.Nil(t, plg)
}

func TestBuildSpecific(t *testing.T) {
	cfg := &Config{Path: "./testdata/valid", Names: []string{"User"}}
	plg, err := cfg.Build()
	require.NoError(t, err)
	schemas, err := plg.Load()
	require.NoError(t, err)
	require.Len(t, schemas, 1)
}

func TestBuildNoSchema(t *testing.T) {
	cfg := &Config{Path: "./testdata/invalid"}
	plg, err := cfg.Build()
	require.Error(t, err)
	require.Nil(t, plg)
}
