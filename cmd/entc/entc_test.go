// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCmd(t *testing.T) {
	defer os.RemoveAll("ent")
	cmd := exec.Command("go", "run", "github.com/facebookincubator/ent/cmd/entc", "init", "User")
	stderr := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err := os.Stat("ent/generate.go")
	require.NoError(t, err)
	_, err = os.Stat("ent/schema/user.go")
	require.NoError(t, err)

	cmd = exec.Command("go", "run", "github.com/facebookincubator/ent/cmd/entc", "generate", "./ent/schema")
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err = os.Stat("ent/user.go")
	require.NoError(t, err)
}

func TestCmdWithTargetArg(t *testing.T) {
	basePath := "entity"
	schemaDir := "schema"
	absPath, _ := filepath.Abs(filepath.Join(basePath, "/", schemaDir))

	defer os.RemoveAll(basePath)
	cmd := exec.Command("go", "run", "github.com/facebookincubator/ent/cmd/entc", "init", "--target", absPath, "User")
	stderr := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err := os.Stat(filepath.Join(basePath, "/generate.go"))
	require.NoError(t, err)
	_, err = os.Stat(filepath.Join(basePath, "/", schemaDir, "/user.go"))
	require.NoError(t, err)

	cmd = exec.Command("go", "run", "github.com/facebookincubator/ent/cmd/entc", "generate", absPath)
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err = os.Stat(filepath.Join(basePath, "/user.go"))
	require.NoError(t, err)
}
