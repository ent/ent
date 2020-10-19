// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckDir(t *testing.T) {
	target := filepath.Join(os.TempDir(), "entvcs")
	require.NoError(t, os.MkdirAll(target, os.ModePerm), "creating tmpdir")
	defer os.RemoveAll(target)
	err := ioutil.WriteFile(filepath.Join(target, "a.go"), []byte(`package schema`), 0644)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(target, "b.go"), []byte(`package schema

type User struct {
<<<<<<< local
        ent.Schema
=======
        schema
>>>>>>> other
}
`), 0644)
	require.NoError(t, err)
	err = CheckDir(target)
	require.Error(t, err)
	expected := fmt.Sprintf("vcs conflict %s/b.go:4", target)
	require.EqualError(t, err, expected)
}
