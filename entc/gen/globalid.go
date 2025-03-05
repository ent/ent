// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"entgo.io/ent/dialect/entsql"
)

const incrementIdent = "const IncrementStarts"

// IncrementStarts holds the autoincrement start value for each type.
type IncrementStarts map[string]int

// IncrementStartAnnotation assigns a unique range to each node in the graph.
func IncrementStartAnnotation(g *Graph) error {
	// To ensure we keep the existing type ranges, load the current global id sequence, if there is any.
	var (
		r        = make(IncrementStarts)
		path     = IncrementStartsFilePath(g.Target)
		buf, err = os.ReadFile(path)
	)
	switch {
	case os.IsNotExist(err): // first time generation
	case err != nil:
		return err
	default:
		if ok, _ := g.FeatureEnabled(FeatureSnapshot.Name); ok {
			if err = ResolveIncrementStartsConflict(g.Target); err != nil {
				return err
			}
			buf, err = os.ReadFile(path)
			if err != nil {
				return err
			}
		}
		var (
			matches = make([][]byte, 0, 2)
			lines   = bytes.Split(buf, []byte("\n"))
		)
		for i := 0; i < len(lines); i++ {
			if l := lines[i]; bytes.HasPrefix(l, []byte(incrementIdent)) {
				matches = append(matches, l)
			}
		}
		if len(matches) != 1 {
			return fmt.Errorf("expect to have exactly 1 ranges in %s", path)
		}
		var (
			line  = matches[0]
			start = bytes.IndexByte(line, '"')
			end   = bytes.LastIndexByte(line, '"')
		)
		if start == -1 || start >= end {
			return fmt.Errorf("unexpected line %s", line)
		}
		l, err := strconv.Unquote(string(line[start : end+1]))
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(l), &r); err != nil {
			return fmt.Errorf("unmarshal ranges: %w", err)
		}
	}
	// Range over all nodes and assign the increment starting value.
	var (
		need    []*Type
		lastIdx = -1
	)
	for _, n := range g.Nodes {
		a := n.EntSQL()
		if a == nil {
			a = &entsql.Annotation{}
		}
		switch v, ok := r[n.Table()]; {
		case a.IncrementStart != nil:
			// In case the start value is defined by an annotation already, it has precedence.
			r[n.Table()] = *a.IncrementStart
		case ok:
			// In case this node has no annotation but an existing entry in the increments file.
			a.IncrementStart = &v
		default:
			// In case this is a new node, it gets the next free increment range (highest value << 32).
			need = append(need, n)
		}
		if v, ok := r[n.Table()]; ok {
			lastIdx = max(lastIdx, v/(1<<32-1))
		}
		if err := setAnnotation(n, a); err != nil {
			return err
		}
	}
	// Compute new ranges and write them back to the file.
	for i, n := range need {
		r[n.Table()] = (lastIdx + i + 1) << 32
		a := n.EntSQL()
		a.IncrementStart = func(i int) *int { return &i }(r[n.Table()]) // copy to not override previous values
		if err := setAnnotation(n, a); err != nil {
			return err
		}
	}
	// Ensure increment ranges are exactly of size 1<<32 with no overlaps.
	d := make(map[int]string)
	for t, s := range r {
		switch t1, ok := d[s]; {
		case ok:
			return fmt.Errorf("duplicated increment start value %d for types %s and %s", s, t1, t)
		case s%(1<<32) != 0:
			return fmt.Errorf(
				"unexpected increment start value %d for type %s, expected multiple of %d (1<<32)", s, t, 1<<32,
			)
		}
		d[s] = t
	}
	if g.Annotations == nil {
		g.Annotations = make(Annotations)
	}
	g.Annotations[r.Name()] = r
	return nil
}

// Name implements Annotation interface.
func (IncrementStarts) Name() string {
	return "IncrementStarts"
}

// WriteToDisk writes the increment starts to the disk.
func (i IncrementStarts) WriteToDisk(target string) error {
	initTemplates()
	p := IncrementStartsFilePath(target)
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return err
	}
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	return templates.Lookup("internal/globalid").
		Execute(f, &Config{
			Target:      target,
			Annotations: Annotations{"IncrementStarts": i},
		})
}

func IncrementStartsFilePath(dir string) string {
	return filepath.Join(dir, "internal", "globalid.go")
}

// ResolveIncrementStartsConflict resolves git/mercurial conflicts by "accepting theirs".
func ResolveIncrementStartsConflict(dir string) error {
	// Expect 2 ranges in the file, accept the second one, since this is the remote content.
	p := IncrementStartsFilePath(dir)
	fi, err := os.Stat(p)
	if err != nil {
		return err
	}
	c, err := os.ReadFile(p)
	if err != nil {
		return err
	}
	var (
		fixed             [][]byte
		conflict, skipped bool
		lines             = bytes.Split(c, []byte("\n"))
	)
	for _, l := range lines {
		switch {
		case bytes.HasPrefix(l, []byte("<<<<<<<")):
			conflict = true
		case bytes.HasPrefix(l, []byte("=======")), bytes.HasPrefix(l, []byte(">>>>>>>")):
		case bytes.HasPrefix(l, []byte(incrementIdent)) && conflict && !skipped:
			skipped = true
		default:
			fixed = append(fixed, l)
		}
	}
	return os.WriteFile(p, bytes.Join(fixed, []byte("\n")), fi.Mode())
}

func ToMap(a *entsql.Annotation) (map[string]any, error) {
	buf, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	m := make(map[string]any)
	if err = json.Unmarshal(buf, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func setAnnotation(n *Type, a *entsql.Annotation) error {
	m, err := ToMap(a)
	if err != nil {
		return err
	}
	n.Annotations.Set(a.Name(), m)
	return nil
}
