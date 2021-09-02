---
id: extensions
title: Extensions
---

### Introduction

The Ent [Extension API](https://pkg.go.dev/entgo.io/ent/entc#Extension)
facilitates the creation of code-generation extensions that bundle together [codegen hooks](code-gen.md#code-generation-hooks),
[templates](templates.md) and [annotations](templates.md#annotations) to create reusable components
that add new rich functionality to Ent's core. For example, Ent's [entgql plugin](https://pkg.go.dev/entgo.io/contrib/entgql#Extension)
exposes an `Extension` that automatically generates GraphQL servers from an Ent schema.

### Defining a New Extension

All extension's must implement the [Extension](https://pkg.go.dev/entgo.io/ent/entc#Extension) interface:

```go
type Extension interface {
	// Hooks holds an optional list of Hooks to apply
	// on the graph before/after the code-generation.
	Hooks() []gen.Hook

	// Annotations injects global annotations to the gen.Config object that
	// can be accessed globally in all templates. Unlike schema annotations,
	// being serializable to JSON raw value is not mandatory.
	//
	//	{{- with $.Config.Annotations.GQL }}
	//		{{/* Annotation usage goes here. */}}
	//	{{- end }}
	//
	Annotations() []Annotation

	// Templates specifies a list of alternative templates
	// to execute or to override the default.
	Templates() []*gen.Template

	// Options specifies a list of entc.Options to evaluate on
	// the gen.Config before executing the code generation.
	Options() []Option
}
```
To simplify the development of new extensions, developers can embed [entc.DefaultExtension](https://pkg.go.dev/entgo.io/ent/entc#DefaultExtension)
to create extensions  without implementing all methods:

```go
package hello

// GreetExtension implements entc.Extension.
type GreetExtension {
	entc.DefaultExtension
}
```

### Adding Templates

Ent supports adding [external templates](templates.md) that will be rendered during
code generation. To bundle such external templates on an extension, implement the `Templates`
method:
```gotemplate title="templates/greet.tmpl"
{{/* Tell Intellij/GoLand to enable the autocompletion based on the *gen.Graph type. */}}
{{/* gotype: entgo.io/ent/entc/gen.Graph */}}

{{ define "greet" }}

{{/* Add the base header for the generated file */}}
{{ $pkg := base $.Config.Package }}
{{ template "header" $ }}

{{/* Loop over all nodes and add the Greet method */}}
{{ range $n := $.Nodes }}
    {{ $receiver := $n.Receiver }}
    func ({{ $receiver }} *{{ $n.Name }}) Greet() string {
		return "Hello, {{ $n.Name }}"
    }
{{ end }}

{{ end }}
```
```go
func (*GreetExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("greet").ParseFiles("templates/greet.tmpl")),
	}
}
```

### Adding Global Annotations

Annotations are a convenient way to supply users of our extension with an API 
to modify the behavior of code generation. To add annotations to our extension,
implement the `Annotations` method. Let's say in our `GreetExtension` we want
to provide users with the ability to configure the greeting word in the generated
code:

```go
// GreetingWord implements entc.Annotation.
type GreetingWord string

// Name of the annotation. Used by the codegen templates.
func (GreetingWord) Name() string {
	return "GreetingWord"
}
```
Then add it to the `GreetExtension` struct:
```go
type GreetExtension struct {
	entc.DefaultExtension
	word GreetingWord
}
```
Next, implement the `Annotations` method:
```go

func (s *GreetExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		s.word,
	}
}
```
Now, from within your templates you can access the `GreetingWord` annotation:
```gotemplate
func ({{ $receiver }} *{{ $n.Name }}) Greet() string {
    return "{{ $.Annotations.GreetingWord }}, {{ $n.Name }}"
}
```

### Adding Hooks

The entc package provides an option to add a list of [hooks](code-gen.md#code-generation-hooks)
(middlewares) to the code-generation phase. This option is ideal for adding custom validators for the
schema, or for generating additional assets using the graph schema. To bundle
code generation hooks with your extension, implement the `Hooks` method:

```go
func (s *GreetExtension) Hooks() []gen.Hook {
    return []gen.Hook{
        DisallowTypeName("Shalom"),
    }
}

// DisallowTypeName ensures there is no ent.Schema with the given name in the graph.
func DisallowTypeName(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				if node.Name == name {
					return fmt.Errorf("entc: validation failed, type named %q not allowed.", name)
				}
			}
			return next.Generate(g)
		})
	}
}
```

### Using an Extension in Code Generation

To use an extension in our code-generation configuration, use `entc.Extensions`, a helper
method that returns an `entc.Option` that applies our chosen extensions:

```go title=ent/entc.go
//+build ignore

package main

import (
	"fmt"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(&GreetExtension{
			word: GreetingWord("Shalom"),
		}),
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```
