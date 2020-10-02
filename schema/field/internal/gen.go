// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// gen is a codegen cmd for generating numeric build types from template.
package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/facebook/ent/schema/field"
)

func main() {
	buf, err := ioutil.ReadFile("internal/numeric.tmpl")
	if err != nil {
		log.Fatal("reading template file:", err)
	}
	intTmpl := template.Must(template.New("numeric").
		Funcs(template.FuncMap{"title": strings.Title, "hasPrefix": strings.HasPrefix}).
		Parse(string(buf)))
	b := &bytes.Buffer{}
	if err := intTmpl.Execute(b, struct {
		Ints, Floats []field.Type
	}{
		Ints: []field.Type{
			field.TypeInt,
			field.TypeUint,
			field.TypeInt8,
			field.TypeInt16,
			field.TypeInt32,
			field.TypeInt64,
			field.TypeUint8,
			field.TypeUint16,
			field.TypeUint32,
			field.TypeUint64,
		},
		Floats: []field.Type{
			field.TypeFloat64,
			field.TypeFloat32,
		},
	}); err != nil {
		log.Fatal("executing template:", err)
	}
	if buf, err = format.Source(b.Bytes()); err != nil {
		log.Fatal("formatting output:", err)
	}
	if err := ioutil.WriteFile("numeric.go", buf, 0644); err != nil {
		log.Fatal("writing go file:", err)
	}
}
