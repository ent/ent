// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/facebook/ent/entc/gen"

	"github.com/olekukonko/tablewriter"
)

// printer is a table printer for ent graphs.
type printer struct {
	io.Writer
}

// Print prints a table description of the graph to the given writer.
func (p printer) Print(g *gen.Graph) {
	for _, n := range g.Nodes {
		p.node(n)
	}
}

// node returns description of a type. The format of the description is:
//
//	Type:
//			<Fields Table>
//
//			<Edges Table>
//
func (p printer) node(t *gen.Type) {
	var (
		b      strings.Builder
		table  = tablewriter.NewWriter(&b)
		header = []string{"Field", "Type", "Unique", "Optional", "Nillable", "Default", "UpdateDefault", "Immutable", "StructTag", "Validators"}
	)
	b.WriteString(t.Name + ":\n")
	table.SetAutoFormatHeaders(false)
	table.SetHeader(header)
	for _, f := range append([]*gen.Field{t.ID}, t.Fields...) {
		v := reflect.ValueOf(*f)
		row := make([]string, len(header))
		for i := range row {
			field := v.FieldByNameFunc(func(name string) bool {
				// The first field is mapped from "Name" to "Field".
				return name == "Name" && i == 0 || name == header[i]
			})
			row[i] = fmt.Sprint(field.Interface())
		}
		table.Append(row)
	}
	table.Render()
	table = tablewriter.NewWriter(&b)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Edge", "Type", "Inverse", "BackRef", "Relation", "Unique", "Optional"})
	for _, e := range t.Edges {
		table.Append([]string{
			e.Name,
			e.Type.Name,
			strconv.FormatBool(e.IsInverse()),
			e.Inverse,
			e.Rel.Type.String(),
			strconv.FormatBool(e.Unique),
			strconv.FormatBool(e.Optional),
		})
	}
	if table.NumLines() > 0 {
		table.Render()
	}
	io.WriteString(p, strings.ReplaceAll(b.String(), "\n", "\n\t")+"\n")
}
