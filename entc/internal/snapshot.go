// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

// Snapshot describes the schema snapshot restore.
type Snapshot struct {
	Path   string      // Path to snapshot.
	Config *gen.Config // Config of codegen.
}

// Restore restores the generated package from the latest schema snapshot.
// If there is a conflict between upstream and local snapshots, it is merged
// before running the code generation.
func (s *Snapshot) Restore() error {
	buf, err := os.ReadFile(s.Path)
	if err != nil {
		return fmt.Errorf("unable to read snapshot schema %w", err)
	}
	snap, err := s.parseSnapshot(buf)
	if err != nil {
		return err
	}
	s.Config.Schema = snap.Schema
	s.Config.Package = snap.Package
	s.addFeatures(snap)
	graph, err := gen.NewGraph(s.Config, snap.Schemas...)
	if err != nil {
		return err
	}
	return graph.Gen()
}

// schemaIdent holds the schema identifier in snapshot file.
const schemaIdent = "const Schema"

// parseSnapshot parses the given buffer and extract the generated snapshot.
// If it encounters a merge-conflict, it will resolve it by merging the relevant
// parts for the codegen.
func (s *Snapshot) parseSnapshot(buf []byte) (*gen.Snapshot, error) {
	var (
		conflict bool
		matches  = make([][]byte, 0, 2)
		lines    = bytes.Split(buf, []byte("\n"))
	)
	for i := 0; i < len(lines); i++ {
		switch line := lines[i]; {
		case bytes.HasPrefix(line, []byte(schemaIdent)):
			matches = append(matches, line)
		case bytes.HasPrefix(line, []byte(conflictMarker)):
			conflict = true
		}
	}
	switch n := len(matches); {
	case n == 0:
		return nil, fmt.Errorf("schema snapshot was not found in %s", s.Path)
	case n > 1 && !conflict:
		return nil, fmt.Errorf("expect to have exactly 1 snapshot in %s", s.Path)
	}
	line, err := trim(matches[0])
	if err != nil {
		return nil, err
	}
	local := &gen.Snapshot{}
	if err := json.Unmarshal(line, &local); err != nil {
		return nil, fmt.Errorf("unmarshal snapshot %v: %w", local, err)
	}
	if !conflict || len(matches) == 1 {
		return local, nil
	}
	// In case of merge-conflict, we merge the 2 schemas.
	line, err = trim(matches[1])
	if err != nil {
		return nil, err
	}
	other := &gen.Snapshot{}
	if err := json.Unmarshal(line, &other); err != nil {
		return nil, fmt.Errorf("unmarshal snapshot %v: %w", local, err)
	}
	merge(local, other)
	return local, nil
}

// addFeatures adds the features in the snapshot to the codegen config.
func (s *Snapshot) addFeatures(snap *gen.Snapshot) {
	add := make(map[string]gen.Feature)
	for _, name := range snap.Features {
		for _, feat := range gen.AllFeatures {
			if name == feat.Name {
				add[name] = feat
			}
		}
	}
	for _, feat := range s.Config.Features {
		delete(add, feat.Name)
	}
	for _, feat := range add {
		s.Config.Features = append(s.Config.Features, feat)
	}
}

// merge the "other"/"upstream" snapshot to the "local" version.
func merge(local, other *gen.Snapshot) {
	if local.Schema == "" {
		local.Schema = other.Schema
	}
	if local.Package == "" {
		local.Package = other.Package
	}
	locals := make(map[string]*load.Schema, len(local.Schemas))
	for _, schema := range local.Schemas {
		locals[schema.Name] = schema
	}
	// Merge "other" schemas.
	for _, schema := range other.Schemas {
		switch match, ok := locals[schema.Name]; {
		case !ok:
			local.Schemas = append(local.Schemas, schema)
		default:
			mergeSchema(match, schema)
		}
	}
	// Merge codegen features.
	features := make(map[string]struct{}, len(local.Features))
	for _, feat := range local.Features {
		features[feat] = struct{}{}
	}
	for _, feat := range other.Features {
		if _, ok := features[feat]; !ok {
			local.Features = append(local.Features, feat)
		}
	}
}

// mergeSchema merges to "local" the additional information in
// the "other" schema, that may be necessary for code-generation.
func mergeSchema(local, other *load.Schema) {
	if local.Config.Table == "" {
		local.Config.Table = other.Config.Table
	}
	if local.Annotations == nil && other.Annotations != nil {
		local.Annotations = make(map[string]any)
	}
	for ant := range other.Annotations {
		if _, ok := local.Annotations[ant]; !ok {
			local.Annotations[ant] = other.Annotations[ant]
		}
	}
	fields := make(map[string]*load.Field, len(local.Fields))
	for _, f := range local.Fields {
		fields[f.Name] = f
	}
	for _, f := range other.Fields {
		switch match, ok := fields[f.Name]; {
		case !ok:
			local.Fields = append(local.Fields, f)
		default:
			mergeField(match, f)
		}
	}
	edges := make(map[string]*load.Edge, len(local.Edges))
	for _, e := range local.Edges {
		edges[e.Name] = e
	}
	for _, e := range other.Edges {
		switch match, ok := edges[e.Name]; {
		case !ok:
			local.Edges = append(local.Edges, e)
		default:
			mergeEdge(match, e)
		}
	}
}

// mergeField merges to "local" the additional information in
// the "other" field, that may be necessary for code-generation.
func mergeField(local, other *load.Field) {
	if local.Annotations == nil && other.Annotations != nil {
		local.Annotations = make(map[string]any)
	}
	for ant := range other.Annotations {
		if _, ok := local.Annotations[ant]; !ok {
			local.Annotations[ant] = other.Annotations[ant]
		}
	}
	if !local.Immutable && other.Immutable {
		local.Immutable = other.Immutable
	}
}

// mergeEdge merges to "local" the additional information in
// the "other" edge, that may be necessary for code-generation.
func mergeEdge(local, other *load.Edge) {
	if local.Annotations == nil && other.Annotations != nil {
		local.Annotations = make(map[string]any)
	}
	for ant := range other.Annotations {
		if _, ok := local.Annotations[ant]; !ok {
			local.Annotations[ant] = other.Annotations[ant]
		}
	}
}

// IsBuildError reports if the given error is an error from the Go command (e.g. syntax error).
func IsBuildError(err error) bool {
	if strings.HasPrefix(err.Error(), "entc/load: #") {
		return true
	}
	for _, s := range []string{
		"syntax error",
		"previous declaration",
		"invalid character",
		"could not import",
		"found '<<'",
	} {
		if strings.Contains(err.Error(), s) {
			return true
		}
	}
	return false
}

func trim(line []byte) ([]byte, error) {
	start := bytes.IndexByte(line, '"')
	end := bytes.LastIndexByte(line, '"')
	if start == -1 || start >= end {
		return nil, fmt.Errorf("unexpected snapshot line %s", line)
	}
	l, err := strconv.Unquote(string(line[start : end+1]))
	if err != nil {
		return nil, err
	}
	return []byte(l), nil
}
