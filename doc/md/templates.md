---
id: templates
title: External Templates
---

`entc` accepts external [Go templates](https://golang.org/pkg/text/template) to execute using the `--template` flag.
If the template name is already defined by `entc`, it will override the existing one. Otherwise, it will write the
execution output to a file with the same name as the template. For example:

`node.tmpl` - This template example will be written in a file named: `ent/node.go`.
```gotemplate
{{ define "node" }}

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

## Examples
A custom template for implementing the `Node` API for GraphQL - 
[Github](https://github.com/facebookincubator/ent/blob/master/entc/integration/template/ent/template/node.tmpl).

## Documentation

Templates are executed on either a specific node-type or the entire schema graph. For API
documentation, see the <a target="_blank" href="https://godoc.org/github.com/facebookincubator/ent/entc/gen">GoDoc<a>.
