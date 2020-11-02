---
id: faq
title: Frequently Asked Questions (FAQ)
sidebar_label: FAQ
---

## Questions

[How to create a struct (or a mutation) level validator?](#how-to-create-a-mutation-level-validator)

[How to create an entity from a struct `T`?](#how-to-create-an-entity-from-a-struct-t)

## Answers

#### How to create a mutation level validator?

In order to implement a mutation-level validator, you can either use [schema hooks](hooks.md#schema-hooks) for validating
changes applied on one entity type, or use [transaction hooks](transactions.md#hooks) for validating mutations that being
applied on multiple entity types (e.g. a GraphQL mutation).

#### How to create an entity from a struct `T`?

The different builders don't support the option of setting the entity fields (or edges) from a given struct `T`.
The reason is that there's no way to distinguish between zero/real values when updating the database (for example, `&ent.T{Age: 0, Name: ""}`).
Setting these values, may set incorrect values in the database or update unnecessary columns.

However, the [external template](templates.md) option lets you extend the default code-generation assets by adding custom logic.
For example, in order to generate a method for each of the create-builders, that accepts a struct as an input and configure the builder,
use the following template:

```gotemplate
{{ range $n := $.Nodes }}
    {{ $builder := $n.CreateName }}
    {{ $receiver := receiver $builder }}

    func ({{ $receiver }} *{{ $builder }}) Set{{ $n.Name }}(input *{{ $n.Name }}) *{{ $builder }} {
        {{- range $f := $n.Fields }}
            {{- $setter := print "Set" $f.StructField }}
            {{ $receiver }}.{{ $setter }}(input.{{ $f.StructField }})
        {{- end }}
        return {{ $receiver }}
    }
{{ end }}
```