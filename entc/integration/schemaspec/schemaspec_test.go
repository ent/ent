
package schema

import (
	"os"
	"path/filepath"
	"testing"

	"entgo.io/ent/entc/integration/schemaspec/ent"
)

func TestGeneratesWithSchemaSpec(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getting cwd: %v", err)
	}
	err = ent.RunEntc(filepath.Join(cwd, "ent"))
	if err != nil {
		t.Fatalf("expected no error, received: %v", err)
	}
}
