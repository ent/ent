---
title: Extending Ent with the Extension API
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
---

A few months ago, [Ariel](https://github.com/a8m) made a silent but highly-impactful contribution
to Ent's core, the [Extension API](https://entgo.io/docs/extensions). While Ent has had extension capabilities (such as [Code-gen Hooks](https://entgo.io/docs/code-gen/#code-generation-hooks),
[External Templates](https://entgo.io/docs/templates/), and [Annotations](https://entgo.io/docs/templates/#annotations))
for a long time, there wasn't a convenient way to bundle together all of these moving parts into a 
coherent, self-contained component. The [Extension API](https://entgo.io/docs/extensions) which we 
discuss in the post does exactly that. 

Many open-source ecosystems thrive specifically because they excel at providing developers an
easy and structured way to extend a small, core system. Much criticism has been made of the 
Node.js ecosystem (even by its [original creator Ryan Dahl](https://www.youtube.com/watch?v=M3BM9TB-8yA))
but it is very hard to argue that the ease of publishing and consuming new `npm` modules
facilitated the explosion in its popularity. I've discussed on my personal blog how 
[protoc's plugin system works](https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/) 
and how that made the Protobuf ecosystem thrive. In short, ecosystems are only created under
modular designs. 

In our post today, we will explore Ent's `Extension` API by building a toy example. 

### Getting Started

The Extension API only works for projects use Ent's code-generation [as a Go package](https://entgo.io/docs/code-gen/#use-entc-as-a-package).
To set that up, after initializing your project, create a new file named `ent/entc.go`:
```go title=ent/entc.go
//+build ignore

package main

import (
    "log"

    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
    "entgo.io/ent/schema/field"
)

func main() {
    err := entc.Generate("./schema", &gen.Config{})
    if err != nil {
        log.Fatal("running ent codegen:", err)
    }
}
```
Next, modify `ent/generate.go` to invoke our `entc` file:
```go title=ent/generate.go
package ent

//go:generate go run entc.go
```

### Creating our Extension

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
to create extensions without implementing all methods. In `entc.go`, add: 
```go title=ent/entc.go
// ...

// GreetExtension implements entc.Extension.
type GreetExtension {
	entc.DefaultExtension
}
```

Currently, our extension doesn't do anything. Next, let's connect it to our code-generation config.
In `entc.go`, add our new extension to the `entc.Generate` invocation:

```go
err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(&GreetExtension{})
```

### Adding Templates

External templates can be bundled into extensions to enhance Ent's core code-generation
functionality. With our toy example, our goal is to add to each entity a generated method
name `Greet` that returns a greeting with the type's name when invoked. We're aiming for something 
like:

```go
func (u *User) Greet() string {
    return "Greetings, User"
}
```

To do this, let's add a new external template file and place it in `ent/templates/greet.tmpl`:
```gotemplate title="ent/templates/greet.tmpl"
{{ define "greet" }}

    {{/* Add the base header for the generated file */}}
    {{ $pkg := base $.Config.Package }}
    {{ template "header" $ }}

    {{/* Loop over all nodes and add the Greet method */}}
    {{ range $n := $.Nodes }}
        {{ $receiver := $n.Receiver }}
        func ({{ $receiver }} *{{ $n.Name }}) Greet() string {
            return "Greetings, {{ $n.Name }}"
        }
    {{ end }}
{{ end }}
```

Next, let's implement the `Templates` method:

```go title="ent/entc.go"
func (*GreetExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("greet").ParseFiles("templates/greet.tmpl")),
	}
}
```

Next, let's kick the tires on our extension. Add a new schema for the `User` type in a file
named `ent/schema/user.go`:

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email_address").
			Unique(),
	}
}
```

Next, run:
```shell
go generate ./...
```

Observe that a new file, `ent/greet.go`, was created, it contains:

```go title="ent/greet.go"
// Code generated by ent, DO NOT EDIT.

package ent

func (u *User) Greet() string {
	return "Greetings, User"
}
```

Great! Our extension was invoked from Ent's code-generation and produced the code
we wanted for our schema!

### Adding Annotations

Annotations provide a way to supply users of our extension with an API
to modify the behavior of code generation logic. To add annotations to our extension,
implement the `Annotations` method. Suppose that for our `GreetExtension` we want
to provide users with the ability to configure the greeting word in the generated
code:

```go
// GreetingWord implements entc.Annotation
type GreetingWord string

func (GreetingWord) Name() string {
	return "GreetingWord"
}
```
Next, we add a `word` field to our `GreetExtension` struct:
```go
type GreetExtension struct {
	entc.DefaultExtension
	Word GreetingWord
}
```
Next, implement the `Annotations` method:
```go
func (s *GreetExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		s.Word,
	}
}
```
Now, from within your templates you can access the `GreetingWord` annotation. Modify
`ent/templates/greet.tmpl` to use our new annotation:

```gotemplate
func ({{ $receiver }} *{{ $n.Name }}) Greet() string {
    return "{{ $.Annotations.GreetingWord }}, {{ $n.Name }}"
}
```
Next, modify the code-generation configuration to set the GreetingWord annotation:
```go title="ent/entc.go
err := entc.Generate("./schema",
    &gen.Config{},
    entc.Extensions(&GreetExtension{
        Word: GreetingWord("Shalom"),
    }),
)
```
To see our annotation control the generated code, re-run:
```shell
go generate ./...
```
Finally, observe that the generated `ent/greet.go` was updated:

```go
func (u *User) Greet() string {
	return "Shalom, User"
}
```

Hooray! We added an option to use an annotation to control the greeting word in the
generated `Greet` method!

### More Possibilities

In addition to templates and annotations, the Extension API allows developers to bundle 
`gen.Hook`s and `entc.Option`s in extensions to further control the behavior of your code-generation.
In this post we will not discuss these possibilities, but if you are interested in using them
head over to the [documentation](https://entgo.io/docs/extensions).

### Wrapping Up

In this post we explored via a toy example how to use the `Extension` API to create new
Ent code-generation extensions. As we've mentioned above, modular design that allows anyone
to extend the core functionality of software is critical to the success of any ecosystem.
We're seeing this claim start to realize with the Ent community, here's a list of some 
interesting projects that use the Extension API:
* [elk](https://github.com/masseelch/elk) - an extension to generate REST endpoints from Ent schemas.
* [entgql](https://github.com/ent/contrib/tree/master/entgql) - generate GraphQL servers from Ent schemas.
* [entviz](https://github.com/hedwigz/entviz) - generate ER diagrams from Ent schemas. 

And what about you? Do you have an idea for a useful Ent extension? I hope this post
demonstrated that with the new Extension API, it is not a difficult task. 

Have questions? Need help with getting started? Feel free to join our [Discord server](https://discord.gg/qZmPgTE6RX) or [Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
