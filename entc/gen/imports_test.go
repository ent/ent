// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"database/sql"
	"encoding/json"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"
	"strconv"
	"testing"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/imports"
)

// ImportsUser and ImportsGroup cover the field kinds whose generated code
// needs extra imports: time, JSON, enums, UUID IDs, edge fields and custom
// Go types backed by database/sql.
type ImportsUser struct{ ent.Schema }

func (ImportsUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(time.Now),
		field.JSON("labels", []string{}).Optional(),
		field.JSON("attrs", map[string]any{}).Optional(),
		field.Enum("state").Values("on", "off"),
		field.String("nick").GoType(sql.NullString{}).Optional(),
		field.UUID("group_id", uuid.UUID{}).Optional(),
	}
}

func (ImportsUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", ImportsGroup.Type).Unique().Field("group_id"),
		edge.To("friends", ImportsUser.Type),
	}
}

type ImportsGroup struct{ ent.Schema }

func (ImportsGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Int("size").Optional(),
	}
}

func (ImportsGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", ImportsUser.Type).Ref("group"),
	}
}

// TestGeneratedFilesDeclareTheirImports ensures that goimports never has to
// add an import to a generated file. A file that references a package it does
// not import sends goimports down its slow path, which parses every sibling
// file in the package directory; on large schemas that dominates codegen time.
func TestGeneratedFilesDeclareTheirImports(t *testing.T) {
	var schemas []*load.Schema
	for _, s := range []ent.Interface{ImportsUser{}, ImportsGroup{}} {
		b, err := load.MarshalSchema(s)
		require.NoError(t, err)
		schema := &load.Schema{}
		require.NoError(t, json.Unmarshal(b, schema))
		schemas = append(schemas, schema)
	}
	for _, storage := range drivers {
		t.Run(storage.Name, func(t *testing.T) {
			var features []Feature
			if storage.Name == "sql" {
				features = AllFeatures
			}
			target := filepath.Join(t.TempDir(), "ent")
			graph, err := NewGraph(&Config{
				Package:  "entc/gen",
				Schema:   "entgo.io/ent/entc/gen",
				Target:   target,
				Storage:  storage,
				IDType:   &field.TypeInfo{Type: field.TypeInt},
				Features: features,
			}, schemas...)
			require.NoError(t, err)
			assets, err := graph.assets()
			require.NoError(t, err)
			// goimports resolves references against sibling files on disk.
			require.NoError(t, assets.write())
			missing := make(map[string][]string)
			for path, content := range assets.files {
				before := importPaths(t, path, content)
				formatted, err := imports.Process(path, content, nil)
				require.NoError(t, err, path)
				var added []string
				for _, p := range importPaths(t, path, formatted) {
					if !contains(before, p) {
						added = append(added, p)
					}
				}
				if len(added) > 0 {
					rel, err := filepath.Rel(target, path)
					require.NoError(t, err)
					missing[rel] = added
				}
			}
			require.Empty(t, missing, "generated files reference packages they do not import")
		})
	}
}

func importPaths(t *testing.T, path string, src []byte) []string {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), path, src, parser.ImportsOnly)
	require.NoError(t, err, path)
	paths := make([]string, 0, len(f.Imports))
	for _, spec := range f.Imports {
		p, err := strconv.Unquote(spec.Path.Value)
		require.NoError(t, err, path)
		paths = append(paths, p)
	}
	sort.Strings(paths)
	return paths
}

func contains(paths []string, p string) bool {
	i := sort.SearchStrings(paths, p)
	return i < len(paths) && paths[i] == p
}
