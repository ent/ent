package plugin

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlugin(t *testing.T) {
	plg := "printer.so"

	// build entc plugin.
	cmd := exec.Command("go", "build", "-o", plg, "-buildmode", "plugin", "./testdata")
	_, err := run(cmd)
	require.NoError(t, err)
	defer os.Remove(plg)

	// execute entc generate and expect the plugin to be executed.
	cmd = exec.Command("go", "run", "../../cmd/entc/entc.go", "generate", "--plugin", plg, "./ent/schema")
	out, err := run(cmd)
	require.NoError(t, err)
	require.Equal(t, "Boring\n", out, "printer plugin should print node names")

}

func run(cmd *exec.Cmd) (string, error) {
	out := bytes.NewBuffer(nil)
	cmd.Stderr = out
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("integration/plugin: %s", out)
	}
	return out.String(), nil
}
