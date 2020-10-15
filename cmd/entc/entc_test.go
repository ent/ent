// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCmd(t *testing.T) {
	defer os.RemoveAll("ent")
	defer os.RemoveAll("entities")
	defer os.RemoveAll("schemas")
	cmd := exec.Command("go", "run", "github.com/facebook/ent/cmd/entc", "init", "User")
	stderr := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err := os.Stat("ent/generate.go")
	require.NoError(t, err)
	_, err = os.Stat("ent/schema/user.go")
	require.NoError(t, err)

	cmd = exec.Command(
		"go",
		"run",
		"github.com/facebook/ent/cmd/entc",
		"init",
		"--target",
		"schemas",
		"--entity",
		"entities",
		"User",
	)
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err = os.Stat("entities/generate.go")
	require.NoError(t, err)
	_, err = os.Stat("schemas/user.go")
	require.NoError(t, err)

	cmd = exec.Command(
		"go",
		"run",
		"github.com/facebook/ent/cmd/entc",
		"init",
		"--target",
		"samedir",
		"--entity",
		"samedir",
		"User",
	)
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.Error(t, cmd.Run(), "schemas and entities can't be generated in the same path. see: entc help init")

	cmd = exec.Command("go", "run", "github.com/facebook/ent/cmd/entc", "generate", "./ent/schema")
	stderr = bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	require.NoError(t, cmd.Run(), stderr.String())

	_, err = os.Stat("ent/user.go")
	require.NoError(t, err)
}
