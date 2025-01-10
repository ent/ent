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
type IncrementStarts map[string]int64

// IncrementStartAnnotation assigns a unique range to each node in the graph.
func IncrementStartAnnotation(g *Graph) error {
	// To ensure we keep the existing type ranges, load the current global id sequence, if there is any.
	var (
		r        = make(IncrementStarts)
		path     = rangesFilePath(g.Target)
		buf, err = os.ReadFile(path)
	)
	switch {
	case os.IsNotExist(err): // first time generation
	case err != nil:
		return err
	default:
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
		need []*Type
		last int64
	)
	for _, n := range g.Nodes {
		if n.Annotations == nil {
			n.Annotations = make(Annotations)
		}
		a := n.EntSQL()
		if a == nil {
			a = &entsql.Annotation{}
			n.Annotations[a.Name()] = a
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
		last = max(last, r[n.Table()])
	}
	// Compute new ranges and write them back to the file.
	s := len(g.Nodes) - len(need) // number of nodes with existing increment values
	for i, n := range need {
		r[n.Table()] = last + int64(s+i)<<32
		a := n.EntSQL()
		a.IncrementStart = func(i int64) *int64 { return &i }(r[n.Table()]) // copy to not override previous values
		n.Annotations[a.Name()] = a
	}
	// Ensure increment ranges are exactly of size 1<<32 with no overlaps.
	d := make(map[int64]string)
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

func rangesFilePath(dir string) string {
	return filepath.Join(dir, "internal", "globalid.go")
}
