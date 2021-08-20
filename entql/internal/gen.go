// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// A codegen cmd for generating builder types from template.
package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"entgo.io/ent/schema/field"
)

func main() {
	buf, err := os.ReadFile("internal/types.tmpl")
	if err != nil {
		log.Fatal("reading template file:", err)
	}
	tmpl := template.Must(template.New("types").
		Funcs(template.FuncMap{
			"ops":   ops,
			"title": strings.Title,
			"ident": ident,
			"type":  typ,
		}).
		Parse(string(buf)))
	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, struct {
		Types []field.Type
	}{
		Types: []field.Type{
			field.TypeBool,
			field.TypeBytes,
			field.TypeTime,
			field.TypeUint,
			field.TypeUint8,
			field.TypeUint16,
			field.TypeUint32,
			field.TypeUint64,
			field.TypeInt,
			field.TypeInt8,
			field.TypeInt16,
			field.TypeInt32,
			field.TypeInt64,
			field.TypeFloat32,
			field.TypeFloat64,
			field.TypeString,
			field.TypeUUID,
			field.TypeOther,
		},
	}); err != nil {
		log.Fatal("executing template:", err)
	}
	if buf, err = format.Source(b.Bytes()); err != nil {
		log.Fatal("formatting output:", err)
	}
	if err := os.WriteFile("types.go", buf, 0644); err != nil {
		log.Fatal("writing go file:", err)
	}
}

func ops(t field.Type) []string {
	switch t {
	case field.TypeBool, field.TypeBytes, field.TypeUUID, field.TypeOther:
		return []string{"EQ", "NEQ"}
	default:
		return []string{"EQ", "NEQ", "LT", "LTE", "GT", "GTE"}
	}
}

func ident(t field.Type) string {
	switch t {
	case field.TypeBytes:
		return "bytes"
	case field.TypeTime:
		return "time"
	case field.TypeUUID:
		return "value"
	case field.TypeOther:
		return "other"
	default:
		return t.String()
	}
}

func typ(t field.Type) string {
	if t == field.TypeUUID || t == field.TypeOther {
		return "driver.Valuer"
	}
	return t.String()
}
