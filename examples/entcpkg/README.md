# entcpkg example

An example of using `entc` (ent codegen) as package rather than an executable.

In this example, we have a file named `entc.go` under the `./ent` directory, that holds the
configuration for the codegen:

```go
// +build ignore

package main

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Header: `
			// Copyright 2019-present Facebook Inc. All rights reserved.
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated (@generated) by entc, DO NOT EDIT.
		`,
		IDType: &field.TypeInfo{Type: field.TypeInt},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

As you can see, the file is tagged with `// +build ignore` in order to not include it
in the `ent` package. In order to run the codegen, run the file itself (using `go run`)
or run `go generate ./ent`. The `generate.go` file holds the `go run command`:

```go
package ent

//go:generate go run entc.go
```

The `generate.go` file is preferred if you have many `generate` pragmas in your project.
