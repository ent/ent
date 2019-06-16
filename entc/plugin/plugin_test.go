package plugin

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecInvalid(t *testing.T) {
	err := Exec("./testdata/notfound", nil)
	require.Error(t, err, "plugin not found")

	dest := "invalid.so"
	require.NoError(t, buildPlg("./testdata/invalid", dest))
	defer os.Remove(dest)

	err = Exec(dest, nil)
	require.Error(t, err, "does not implement the entc/plugin interface")
}

func TestExecValid(t *testing.T) {
	err := Exec("./testdata/notfound", nil)
	require.Error(t, err, "plugin not found")

	dest := "valid.so"
	require.NoError(t, buildPlg("./testdata/valid", dest))
	defer os.Remove(dest)

	err = Exec(dest, nil)
	require.NoError(t, err)
}

func buildPlg(src, dest string) error {
	out := bytes.NewBuffer(nil)
	cmd := exec.Command("go", "build", "-o", dest, "-buildmode", "plugin", src)
	cmd.Stderr = out
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("entc/plugin: %s", out)
	}
	return nil
}
