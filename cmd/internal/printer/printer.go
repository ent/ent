// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package printer

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"entgo.io/ent/entc/gen"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

// A Config controls the output of Fprint.
type Config struct {
	io.Writer
}

// Print prints a table description of the graph to the given writer.
func (p Config) Print(g *gen.Graph) {
	for _, n := range g.Nodes {
		p.node(n)
	}
}

// Fprint executes "pretty-printer" on the given writer.
func Fprint(w io.Writer, g *gen.Graph) {
	Config{Writer: w}.Print(g)
}

// node returns description of a type. The format of the description is:
//
//	Type:
//			<Fields Table>
//
//			<Edges Table>
func (p Config) node(t *gen.Type) {
	var (
		b      strings.Builder
		id     []*gen.Field
		header = []string{"Field", "Type", "Unique", "Optional", "Nillable", "Default", "UpdateDefault", "Immutable", "StructTag", "Validators", "Comment"}
		table  = tablewriter.NewTable(&b, tablewriter.WithHeaderAutoFormat(tw.Off), tablewriter.WithHeader(header))
	)
	b.WriteString(t.Name + ":\n")
	if t.ID != nil {
		id = append(id, t.ID)
	}
	for _, f := range append(id, t.Fields...) {
		v := reflect.ValueOf(*f)
		row := make([]string, len(header))
		for i := 0; i < len(row)-1; i++ {
			field := v.FieldByNameFunc(func(name string) bool {
				// The first field is mapped from "Name" to "Field".
				return name == "Name" && i == 0 || name == header[i]
			})
			row[i] = fmt.Sprint(field.Interface())
		}
		row[len(row)-1] = f.Comment()
		_ = table.Append(row)
	}
	_ = table.Render()
	table = tablewriter.NewTable(&b, tablewriter.WithHeaderAutoFormat(tw.Off),
		tablewriter.WithHeader([]string{"Edge", "Type", "Inverse", "BackRef", "Relation", "Unique", "Optional", "Comment"}))
	for _, e := range t.Edges {
		_ = table.Append([]string{
			e.Name,
			e.Type.Name,
			strconv.FormatBool(e.IsInverse()),
			e.Inverse,
			e.Rel.Type.String(),
			strconv.FormatBool(e.Unique),
			strconv.FormatBool(e.Optional),
			e.Comment(),
		})
	}

	if len(t.Edges) > 0 {
		_ = table.Render()
	}
	io.WriteString(p, strings.ReplaceAll(b.String(), "\n", "\n\t")+"\n")
}
