---
id: templates
title: External Templates
---

`entc` accepts external [Go templates](https://golang.org/pkg/text/template) to execute using the `--template` flag.
If the template name is already defined by `entc`, it will override the existing one. Otherwise, it will write the
execution output to a file with the same name as the template. For example:

`stringer.tmpl` - This template example will be written in a file named: `ent/stringer.go`.

```gotemplate
{{ define "stringer" }}

{{/* Add the base header for the generated file */}}
{{ $pkg := base $.Config.Package }}
{{ template "header" $ }}

{{/* Loop over all nodes and add implement the "GoStringer" interface */}}
{{ range $n := $.Nodes }}
	{{ $receiver := $n.Receiver }}
	func ({{ $receiver }} *{{ $n.Name }}) GoString() string {
		if {{ $receiver }} == nil {
			return fmt.Sprintf("{{ $n.Name }}(nil)")
		}
		return {{ $receiver }}.String()
	}
{{ end }}

{{ end }}
```

`debug.tmpl` - This template example will be written in a file named: `ent/debug.go`.

```gotemplate
{{ define "debug" }}

{{/* A template that adds the functionality for running each client <T> in debug mode */}}

{{/* Add the base header for the generated file */}}
{{ $pkg := base $.Config.Package }}
{{ template "header" $ }}

{{/* Loop over all nodes and add option the "Debug" method */}}
{{ range $n := $.Nodes }}
	{{ $client := print $n.Name "Client" }}
	func (c *{{ $client }}) Debug() *{{ $client }} {
		if c.debug {
			return c
		}
		cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
		return &{{ $client }}{config: cfg}
	}
{{ end }}

{{ end }}
```

In order to override an existing template, use its name. For example:
```gotemplate
{{/* A template for adding additional fields to specific types. */}}
{{ define "model/fields/additional" }}
	{{- /* Add static fields to the "Card" entity. */}}
	{{- if eq $.Name "Card" }}
		// StaticField defined by templates.
		StaticField string `json:"static_field,omitempty"`
	{{- end }}
{{ end }}
```

## Annotations
Schema annotations allow attaching metadata to fields and edges and inject them to external templates.  
An annotation must be a Go type that is serializable to JSON raw value (e.g. struct, map or slice)
and implement the [Annotation](https://pkg.go.dev/github.com/facebook/ent/schema/field?tab=doc#Annotation) interface.

Here's an example of an annotation and its usage in schema and template:

1\. An annotation definition:
```go
package entgql

// Annotation annotates fields with metadata for templates.
type Annotation struct {
	// OrderField is the ordering field as defined in graphql schema.
	OrderField string
}

// Name implements ent.Annotation interface.
func (Annotation) Name() string {
	return "EntGQL"
}
```

2\. Annotation usage in ent/schema:

```go
// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("creation_date").
			Annotations(entgql.Annotation{
				OrderField: "CREATED_AT",
			}),
	}
}
```

3\. Annotation usage in external template:
```gotemplate
{{ range $node := $.Nodes }}
	{{ range $f := $node.Fields }}
		{{/* Get the annotation by its name. See: Annotation.Name */}}
		{{ if $annotation := $f.Annotations.EntGQL }}
			{{/* Get the field from the annotation. */}}
			{{ $orderField := $annotation.OrderField }}
		{{ end }}
	{{ end }}
{{ end }}
```


## Examples
- A custom template for implementing the `Node` API for GraphQL - 
[Github](https://github.com/facebook/ent/blob/master/entc/integration/template/ent/template/node.tmpl).

- An example for executing external templates with custom functions. See  [configuration](https://github.com/facebook/ent/blob/master/examples/entcpkg/ent/entc.go) and its
[README](https://github.com/facebook/ent/blob/master/examples/entcpkg) file.

## Documentation

Templates are executed on either a specific node-type or the entire schema graph. For API
documentation, see the <a target="_blank" href="https://pkg.go.dev/github.com/facebook/ent/entc/gen?tab=doc">GoDoc<a>.
